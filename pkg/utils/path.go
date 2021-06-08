package utils

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	terrascanUtils "github.com/accurics/terrascan/pkg/utils"
)

func FindAllSubDirectories(basePath string, dirFilter []string) ([]string, error) {
	dirList := make([]string, 0)
	err := filepath.Walk(basePath, func(filePath string, fileInfo os.FileInfo, err error) error {
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

func FileFileByPrefix(basePath, prefix string) (string, error) {
	fileInfos, err := ioutil.ReadDir(basePath)
	if err != nil {
		return "", err
	}

	for i := range fileInfos {
		if strings.HasPrefix(fileInfos[i].Name(), prefix) {
			return fileInfos[i].Name(), nil
		}
	}

	// file with prefix not found
	return "", nil
}
