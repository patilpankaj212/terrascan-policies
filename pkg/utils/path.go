package utils

import (
	"os"
	"path/filepath"
	"strings"

	terrascanUtils "github.com/accurics/terrascan/pkg/utils"
)

// FindAllSubDirectories finds all the sub directories in a path and filters would any directories specified in dirFilter
func FindAllSubDirectories(basePath string, dirFilter []string) ([]string, error) {
	dirList := make([]string, 0)
	err := filepath.Walk(basePath, func(filePath string, fileInfo os.FileInfo, err error) error {
		// don't include the base path
		if filePath == basePath {
			return err
		}
		for i := range dirFilter {
			if !strings.Contains(filePath, dirFilter[i]) && fileInfo != nil && fileInfo.IsDir() {
				dirList = append(dirList, filePath)
			}
		}
		return err
	})
	return dirList, err
}

// FindAllMetadaFilesInDir finds all rule metadata json files in specified path
func FindAllMetadaFilesInDir(basePath string) ([]string, error) {
	fileList := make([]string, 0)
	allFiles, err := terrascanUtils.FindFilesBySuffixInDir(basePath, []string{"json"})
	if err != nil {
		return nil, err
	}
	for i := range allFiles {
		fileList = append(fileList, filepath.Join(basePath, *allFiles[i]))
	}
	return fileList, nil
}

// FileFileByPrefix finds file with specified prefix in a path
func FileFileByPrefix(basePath, prefix string) (string, error) {
	dirEntries, err := os.ReadDir(basePath)
	if err != nil {
		return "", err
	}

	for i := range dirEntries {
		if strings.HasPrefix(dirEntries[i].Name(), prefix) {
			return dirEntries[i].Name(), nil
		}
	}

	// file with prefix not found
	return "", nil
}
