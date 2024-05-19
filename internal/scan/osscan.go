package scan

import (
	"context"
	"fmt"
	"github.com/Ullaakut/nmap/v3"
	"gmap/internal/cache"
	"gmap/internal/fingerprints"
	"gmap/internal/log"
	"strings"
)

func GenTargetFingerPrint(target string) (string, error) {
	var osFp string
	// It should be ONLY ONE TARGET at a time
	scanner, err := nmap.NewScanner(
		context.Background(),
		nmap.WithTargets(target),
		nmap.WithOSDetection(),
		nmap.WithDebugging(4),
	)
	if err != nil {
		// log.Errorf("unable to create nmap scanner: %v", err)
		return "", err
	}
	result, warnings, err := scanner.Run()
	if len(*warnings) > 0 {
		log.Warnf("run finished with warnings: %s", *warnings)
	}
	if err != nil {
		// log.Errorf("unable to run nmap scan: %v", err)
		return "", err
	}
	for _, host := range result.Hosts {
		for _, osFps := range host.OS.Fingerprints {
			// return osFp.Fingerprint
			osFp = osFps.Fingerprint
			// log.Infof("OS Fingerprint Gathered")

		}
	}
	if osFp != "" {
		osFp = TrimOsFingerPrint(osFp)
	}
	// log.Infof("Scan finished: %d hosts up scanned in %fs", len(result.Hosts), result.Stats.Finished.Elapsed)
	return osFp, nil
}

func TrimOsFingerPrint(osFp string) string {
	// remove "OS:" from the beginning of each line
	lines := strings.Split(osFp, "\n")
	for i, line := range lines {
		if strings.HasPrefix(line, "OS:") {
			lines[i] = strings.TrimPrefix(line, "OS:")
		}
	}
	rawData := strings.Join(lines, "")
	// split by ")"
	resLines := strings.Split(rawData, ")")
	return strings.Join(resLines, ")\n")
}

func GenFinalFingerprint(target string, testTime int, vendor string, model string) string {
	log.Infof("generating fingerprint for %s %s, scan times: %d", vendor, model, testTime)
	for i := 0; i < testTime; i++ {
		osFp, err := GenTargetFingerPrint(target)
		if err != nil {
			log.Errorf("failed to scan target: %v", err)
		}
		if osFp != "" {
			cache.AddFingerPrintsCache(osFp)
		}
		log.Infof("scan %d done, %d left", i+1, testTime-i-1)
	}
	fpsString := cache.GetFingerPrintsCache()
	if len(fpsString) == 0 {
		log.Errorf("no fingerprint gathered, check if the host is up or the target is correct")
		return ""
	}
	fps, _ := fingerprints.DeserializeFpsFromStr(fpsString)
	finalFp := fingerprints.GenFinalFingerprint(fps)
	header := GenHeader(vendor, model)
	finalFp = header + finalFp
	return finalFp
}

func GenHeader(ven, mod string) string {
	/**
	# 2N VOIP doorbell
	Fingerprint 2N Helios IP VoIP doorbell
	Class 2N | embedded || specialized
	*/
	return fmt.Sprintf("# %s %s\nFingerprint %s %s\nClass %s | embedded || specialized\n", ven, mod, ven, mod, ven)
}
