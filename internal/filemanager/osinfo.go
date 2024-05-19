package filemanager

import (
	"errors"
	"gmap/internal/log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

func AddFpToOsInfo(fingerPrint string, dbPath string) error {
	filePath := dbPath
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(fingerPrint + "\n")
	if err != nil {
		return err
	}
	return nil
}

func DeleteFpFromOsInfo(fullNameString, dbPath string) error {
	filePath := dbPath
	data, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}
	dataStr := string(data)
	lines := strings.Split(dataStr, "\n\n")
	for i, line := range lines {
		if strings.Contains(line, fullNameString) {
			lines = append(lines[:i], lines[i+1:]...)
			break
		}

	}
	newData := strings.Join(lines, "\n\n")
	err = os.WriteFile(filePath, []byte(newData), 0644)
	return nil
}

func GetFpFromOsInfo(dbPath string) ([]string, error) {
	var res []string
	filePath := dbPath
	file, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err

	}
	dataStr := string(file)
	lines := strings.Split(dataStr, "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "Fingerprint") {
			res = append(res, strings.TrimPrefix(line, "Fingerprint "))
		}

	}
	return res, nil
}

func HasFpInOsInfo(fullNameString, dbPath string) bool {

	allFps, err := GetFpFromOsInfo(dbPath)
	if err != nil {
		log.Errorf("failed to get fingerprints from the database: %v", err)
	}
	// check if the fingerprint exists in the database
	hasFp := false
	for _, fp := range allFps {
		if fp == fullNameString {
			hasFp = true
			break
		}
	}
	return hasFp
}

func GetNmapOsDbPath() (string, error) {
	nmapExecPath, err := exec.LookPath("nmap")
	if err != nil {
		return "", errors.New("nmap is not installed")
	}
	goos := runtime.GOOS
	nmapDir := filepath.Dir(nmapExecPath)
	nmapDb := path.Join(nmapDir, "nmap-os-db")
	if _, err := os.Stat(nmapDb); os.IsNotExist(err) {
		if goos == "windows" {
			if strings.Contains(nmapDir, "scoop") {
				scoopNmapDir := filepath.Join(nmapDir, "../apps/nmap/")
				// log.Infof("scoop nmap dir: %s", scoopNmapDir)
				// find if there is nmap-os-db in the shims all sub directory
				err := filepath.Walk(scoopNmapDir, func(path string, info os.FileInfo, err error) error {
					if err != nil {
						return err
					}
					if strings.Contains(info.Name(), "nmap-os-db") {
						nmapDb = path
					}
					return nil
				})
				if err != nil {
					return "", errors.New("nmap-os-db file does not exist")
				}
			} else {
				return "", errors.New("nmap-os-db file does not exist")
			}
		} else if goos == "darwin" {
			return "", errors.New("nmap-os-db file does not exist")
		} else if goos == "linux" {
			return "", errors.New("nmap-os-db file does not exist")
		} else {
			return "", errors.New("nmap-os-db file does not exist")
		}
	}
	return nmapDb, nil
}

func init() {
	//if there is no os-info file, create one
	if _, err := os.Stat("./os-info"); os.IsNotExist(err) {
		file, err := os.Create("./os-info")
		if err != nil {
			log.Errorf("failed to create os-info file: %v", err)
		}
		defer file.Close()
	}
}
