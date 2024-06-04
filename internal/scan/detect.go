package scan

import (
	"context"
	"gmap/internal/log"
	"gmap/internal/nmap/v3"
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
				log.Infof("host: %s os: %s", host.Addresses[0].String(), os.Name)
			} else {
				res[host.Addresses[0].String()] = "unknown"
				log.Infof("host: %s os: unknown", host.Addresses[0].String())
			}
			break
		}
	}
	log.Infof("now detecting opened ports")
	scanner, err = nmap.NewScanner(
		context.Background(),
		nmap.WithTargets(targets...),
		nmap.WithPorts("1-65535"),
	)
	if err != nil {
		return nil, err
	}
	result, warnings, err = scanner.Run()
	if len(*warnings) > 0 {

	}
	if err != nil {
		return nil, err
	}
	for _, host := range result.Hosts {
		for _, port := range host.Ports {
			if port.State.State == "open" {
				log.Infof("host: %s port: %d state: %s", host.Addresses[0].String(), port.ID, port.State.State)
			}
		}
	}

	return res, nil
}
