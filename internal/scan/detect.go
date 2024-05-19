package scan

import (
	"context"
	"github.com/Ullaakut/nmap/v3"
	"gmap/internal/log"
	"strings"
)

func Detect(target string) ([]string, error) {
	// detect the alive hosts
	scanner, err := nmap.NewScanner(
		context.Background(),
		nmap.WithTargets(target),
		nmap.WithPingScan(),
	)
	if err != nil {
		return nil, err
	}
	log.Infof("scanning target: %s", target)
	result, warnings, err := scanner.Run()
	if len(*warnings) > 0 {
		for _, warning := range *warnings {
			log.Warnf("warning: %s", warning)
		}
	}
	if err != nil {
		return nil, err
	}
	var aliveHosts []string
	for _, host := range result.Hosts {
		aliveHosts = append(aliveHosts, host.Addresses[0].String())
	}
	return aliveHosts, nil
}

func DetectOs(target string) (map[string]string, error) {
	// detect the os of the target
	res := make(map[string]string)

	targets := strings.Split(target, ",")
	scanner, err := nmap.NewScanner(
		context.Background(),
		nmap.WithTargets(targets...),
		nmap.WithOSDetection(),
	)
	if err != nil {
		return nil, err
	}
	log.Infof("scanning target: %s", target)
	result, warnings, err := scanner.Run()
	if len(*warnings) > 0 {
		for _, warning := range *warnings {
			log.Warnf("warning: %s", warning)
		}
	}
	if err != nil {
		return nil, err
	}
	for _, host := range result.Hosts {
		for _, os := range host.OS.Matches {
			if os.Name != "" {
				res[host.Addresses[0].String()] = os.Name
			} else {
				res[host.Addresses[0].String()] = "unknown"
			}
			break
		}
	}
	return res, nil
}
