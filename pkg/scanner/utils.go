package scanner

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"

	log "github.com/sirupsen/logrus"
)

func ScanDirectoryForConfigFiles(directory string) ([]string, error) {
	var configFiles []string
	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Errorf("Error accessing path %q: %v", path, err)
			return err
		}
		if !info.IsDir() && filepath.Ext(info.Name()) == ".yaml" {
			configFiles = append(configFiles, path)
		}
		return nil
	})

	if err != nil {
		log.Errorf("Error scanning directory for config files: %v", err)
		return nil, err
	}

	return configFiles, nil
}

func ResolveALDirectoryAndConfigs(rootDir, changedFilePath string) (string, []string, error) {
	alDirPattern := regexp.MustCompile(`AL\d{5}`)
	relPath, err := filepath.Rel(rootDir, changedFilePath)
	if err != nil {
		log.Errorf("Error calculating relative path from rootDir to changedFilePath: %v", err)
		return "", nil, err
	}
	pathSegments := strings.Split(relPath, string(filepath.Separator))
	var alDir string
	for _, segment := range pathSegments {
		if alDirPattern.MatchString(segment) {
			alDir = filepath.Join(rootDir, segment)
			break
		}
	}
	if alDir == "" {
		errMsg := "Failed to find AL directory for changed file"
		log.Error(errMsg)
		return "", nil, err
	}
	configPaths, err := ScanDirectoryForConfigFiles(alDir)
	if err != nil {
		log.Errorf("Error scanning AL directory for config files: %v", err)
		return "", nil, err
	}
	return alDir, configPaths, nil
}
