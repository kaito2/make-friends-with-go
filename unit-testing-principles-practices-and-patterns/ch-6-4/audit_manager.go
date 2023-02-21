package main

import (
	"bufio"
	"fmt"
	"github.com/pkg/errors"
	"log"
	"os"
	"path"
	"strings"
	"time"
)

type AuditManager struct {
	maxEntriesPerFile int
	directoryName     string
}

func NewAuditManager(maxEntriesPerFile int, directoryName string) *AuditManager {
	return &AuditManager{maxEntriesPerFile: maxEntriesPerFile, directoryName: directoryName}
}

func (a *AuditManager) AddRecord(visitorName string, timeOfVisit time.Time) error {
	const failed = "failed to add visitor record"

	filePaths, err := os.ReadDir(a.directoryName)
	if err != nil {
		return errors.Wrap(err, failed)
	}
	sorted := make([]string, 0, len(filePaths))
	for _, filePath := range filePaths {
		// TODO: Assert filePath is `file`.
		sorted = append(sorted, filePath.Name())
	}
	log.Println("sorted: ", sorted)

	newRecord := fmt.Sprintf("%s;%s", visitorName, timeOfVisit.String())

	if len(sorted) == 0 {
		newFilePath := path.Join(a.directoryName, "audit_1.txt")
		newFile, err := os.Create(newFilePath)
		if err != nil {
			return errors.Wrap(err, failed)
		}
		defer newFile.Close()
		if _, err := newFile.WriteString(newRecord); err != nil {
			return errors.Wrap(err, failed)
		}
		return nil
	}

	currentFileName := sorted[len(sorted)-1]
	currentFilePath := path.Join(a.directoryName, currentFileName)
	currentFile, err := os.Open(currentFilePath)
	if err != nil {
		return errors.Wrap(err, failed)
	}
	defer currentFile.Close()

	scanner := bufio.NewScanner(currentFile)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if len(lines) < a.maxEntriesPerFile {
		lines = append(lines, newRecord)
		if err := os.WriteFile(currentFilePath, []byte(strings.Join(lines, "\r\n")), 0644); err != nil {
			return errors.Wrap(err, failed)
		}
	} else {
		currentFileIndex := len(sorted)
		newFileIndex := currentFileIndex + 1
		newFileName := fmt.Sprintf("audit_%d.txt", newFileIndex)
		newFilePath := path.Join(a.directoryName, newFileName)

		if err := os.WriteFile(newFilePath, []byte(newRecord), 0644); err != nil {
			return errors.Wrap(err, failed)
		}
	}

	return nil
}
