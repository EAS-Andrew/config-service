package processor

import (
	"config-service/pkg/logging"
	"config-service/pkg/models"
	"fmt"
	"strings"

	"sync"

	"github.com/sirupsen/logrus"
)

type NexusRMProcessor struct{}

func (n *NexusRMProcessor) ProcessAllConfigsForALFolder(alCode string, configs []interface{}) error {
	var wg sync.WaitGroup
	errorsChan := make(chan error, len(configs))

	for _, config := range configs {
		wg.Add(1)
		go func(config interface{}) {
			defer wg.Done()
			switch cfg := config.(type) {
			case models.ArtefactRepository:
				if err := n.processArtefactRepository(&cfg); err != nil {
					logging.Log.WithFields(logrus.Fields{"AL code": alCode, "config": cfg, "error": err.Error()}).Error("Failed to process ArtefactRepository config in NexusRMProcessor")
					errorsChan <- fmt.Errorf("NexusRMProcessor: Failed to process ArtefactRepository config for AL code %s: %v", alCode, err)
				}
			default:
			}
		}(config)
	}

	wg.Wait()
	close(errorsChan)

	var errors []string
	for err := range errorsChan {
		if err != nil {
			errors = append(errors, err.Error())
		}
	}

	if len(errors) > 0 {
		return fmt.Errorf("NexusRMProcessor: Encountered errors while processing configs for AL folder %s: %s", alCode, strings.Join(errors, "; "))
	}

	logging.Log.Infof("NexusRMProcessor: Successfully processed configs for AL folder: %s", alCode)
	return nil
}

func (n *NexusRMProcessor) processArtefactRepository(config *models.ArtefactRepository) error {
	logging.Log.Infof("NexusRMProcessor: Processing ArtefactRepository config: %s", config.Metadata.Name)

	if config.Spec.Type != "oci" {
		logging.Log.WithFields(logrus.Fields{"name": config.Metadata.Name, "type": config.Spec.Type}).Info("NexusRMProcessor: ArtefactRepository type is not OCI, skipping.")
		return nil
	}

	return nil
}

func (n *NexusRMProcessor) Accepts(config interface{}) bool {
	_, ok := config.(models.ArtefactRepository)
	return ok
}
