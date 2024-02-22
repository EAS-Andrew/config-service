package scanner

import (
	"config-service/pkg/logging"
	"os"
	"path/filepath"
	"regexp"
)

func TraverseDirectory(rootPath string) ([]string, error) {
	var dirs []string

	re, err := regexp.Compile(`AL\d{5}$`)
	if err != nil {
		logging.Log.Errorf("Error compiling regex: %v", err)
		return nil, err
	}

	err = filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			logging.Log.Errorf("Error accessing path %q: %v", path, err)
			return err
		}

		if info.IsDir() && re.MatchString(info.Name()) {
			dirs = append(dirs, path)
		}
		return nil
	})
	if err != nil {
		logging.Log.Errorf("Error during directory traversal: %v", err)
		return nil, err
	}
	return dirs, nil
}
