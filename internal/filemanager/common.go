package filemanager

import (
	"os"
	"strings"
)

func MergeFiles(file1, file2 string) error {
	// always using the first file as the base file
	startStr := "\n# GMAP_START\n"
	endStr := "# GMAP_END\n"
	originalData, err := os.ReadFile(file1)
	if err != nil {
		return err
	}

	// read the second file
	secondData, err := os.ReadFile(file2)
	if err != nil {
		return err
	}

	data1 := string(originalData)
	data2 := string(secondData)

	// find the start and end of the block and replace it with the new data
	startIndex := strings.Index(data1, startStr)
	endIndex := strings.Index(data1, endStr)

	fileFp, err := os.OpenFile(file1, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer fileFp.Close()
	if startIndex == -1 || endIndex == -1 {
		// if the block is not found, append the new data to the end of the file
		_, err = fileFp.WriteString(data1 + startStr + "\n" + data2 + endStr)
		return nil
	}
	// replace the block with the new data
	newData := data1[:startIndex] + startStr + "\n" + data2 + endStr + data1[endIndex+len(endStr):]

	_, err = fileFp.WriteString(newData)
	if err != nil {
		return err
	}
	return nil
}
