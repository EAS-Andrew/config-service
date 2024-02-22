package main

import (
	"config-service/pkg/logging"
	"config-service/pkg/parser"
	"config-service/pkg/scanner"
	"flag"
	"os"
)

func main() {
	var rootDir string
	var logLevel string
	flag.StringVar(&rootDir, "dir", "", "Root directory path to start scanning for configuration files")
	flag.StringVar(&logLevel, "logLevel", "error", "Log level (debug, info, warn, error, fatal, panic)")
	flag.Parse()

	os.Setenv("LOG_LEVEL", logLevel)

	if rootDir == "" {
		logging.Log.Fatal("Root directory path is required")
	}

	alDirs, err := scanner.TraverseDirectory(rootDir)
	if err != nil {
		logging.Log.Fatalf("Failed to traverse root directory: %v", err)
	}

	for _, alDir := range alDirs {
		configPaths, err := scanner.ScanDirectoryForConfigFiles(alDir)
		if err != nil {
			logging.Log.Errorf("Failed to scan AL directory for config files: %s, error: %v", alDir, err)
			continue
		}

		if err := parser.DiscoverAndProcessConfigs(alDir, configPaths); err != nil {
			logging.Log.Errorf("Error processing configs for AL directory %s: %v", alDir, err)
		}
	}
	logging.Log.Info("Completed processing all AL directories.")
}
