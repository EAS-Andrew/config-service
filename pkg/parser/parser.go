package parser

import (
	"os"
	"path/filepath"

	"config-service/pkg/logging"
	"config-service/pkg/models"
	"config-service/pkg/processor"

	gopkg "gopkg.in/yaml.v2"
)

func ParseConfigFiles(configPaths []string) ([]interface{}, error) {
	var configs []interface{}
	configNames := make(map[string]bool)
	for _, path := range configPaths {
		bytes, err := os.ReadFile(filepath.Clean(path))
		if err != nil {
			logging.Log.Errorf("Error reading configuration file: %v", err)
			return nil, err
		}
		var rawConfig map[string]interface{}
		if err := gopkg.Unmarshal(bytes, &rawConfig); err != nil {
			logging.Log.Errorf("Error unmarshalling configuration file: %v", err)
			return nil, err
		}
		kind, ok := rawConfig["kind"].(string)
		if !ok {
			logging.Log.Errorf("Configuration file does not specify kind: %s", path)
			continue
		}
		metadata, ok := rawConfig["metadata"].(map[interface{}]interface{})
		if !ok {
			logging.Log.Errorf("Configuration file does not contain metadata: %s", path)
			continue
		}
		name, ok := metadata["name"].(string)
		if !ok {
			logging.Log.Errorf("Configuration metadata does not specify name: %s", path)
			continue
		}
		configKey := kind + "-" + name
		if _, exists := configNames[configKey]; exists {
			logging.Log.Errorf("Duplicate configuration name %s of kind %s found in file: %s. SKIPPING", name, kind, path)
			continue
		}
		switch kind {
		case "Environment":
			var env models.Environment
			if err := gopkg.Unmarshal(bytes, &env); err != nil {
				logging.Log.Errorf("Error unmarshalling Environment configuration: %v", err)
				return nil, err
			}
			configs = append(configs, env)
		case "DeploymentRepository":
			var repo models.DeploymentRepository
			if err := gopkg.Unmarshal(bytes, &repo); err != nil {
				logging.Log.Errorf("Error unmarshalling DeploymentRepository configuration: %v", err)
				return nil, err
			}
			configs = append(configs, repo)
		case "Component":
			var comp models.Component
			if err := gopkg.Unmarshal(bytes, &comp); err != nil {
				logging.Log.Errorf("Error unmarshalling Component configuration: %v", err)
				return nil, err
			}
			configs = append(configs, comp)
		case "ArtefactRepository":
			var artefact models.ArtefactRepository
			if err := gopkg.Unmarshal(bytes, &artefact); err != nil {
				logging.Log.Errorf("Error unmarshalling ArtefactRepository configuration: %v", err)
				return nil, err
			}
			configs = append(configs, artefact)
		default:
			logging.Log.Errorf("Unsupported configuration kind %s in file: %s", kind, path)
		}
		configNames[configKey] = true
	}
	return configs, nil
}

func DiscoverAndProcessConfigs(alCode string, configPaths []string) error {
	configs, err := ParseConfigFiles(configPaths)
	if err != nil {
		logging.Log.Errorf("Failed to parse config files for AL %s: %v", alCode, err)
		return err
	}

	processorFactory := processor.NewProcessorFactory()
	processors := processorFactory.GetProcessors()

	for _, proc := range processors {
		err := proc.ProcessAllConfigsForALFolder(alCode, filterConfigsForProcessor(proc, configs))
		if err != nil {
			logging.Log.Errorf("Failed to process configurations for AL %s with processor %T: %v", alCode, proc, err)
			return err
		}
	}

	logging.Log.Infof("Successfully discovered and processed configurations for AL %s", alCode)
	return nil
}

func filterConfigsForProcessor(proc processor.ALFolderProcessor, configs []interface{}) []interface{} {
	var filteredConfigs []interface{}
	for _, config := range configs {
		if proc.Accepts(config) {
			filteredConfigs = append(filteredConfigs, config)
		}
	}
	return filteredConfigs
}
