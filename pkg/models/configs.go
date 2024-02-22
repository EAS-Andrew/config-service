package models

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type Environment struct {
	APIVersion string `yaml:"apiVersion" validate:"required"`
	Kind       string `yaml:"kind" validate:"required,eq=Environment"`
	Metadata   struct {
		Name string `yaml:"name" validate:"required"`
	} `yaml:"metadata"`
	Spec struct {
		Namespace string `yaml:"namespace" validate:"required"`
		Cluster   string `yaml:"cluster" validate:"required"`
	} `yaml:"spec"`
}

type DeploymentRepository struct {
	APIVersion string `yaml:"apiVersion" validate:"required"`
	Kind       string `yaml:"kind" validate:"required,eq=DeploymentRepository"`
	Metadata   struct {
		Name string `yaml:"name" validate:"required"`
	} `yaml:"metadata"`
	Spec struct {
		Deployments []struct {
			Path        string `yaml:"path" validate:"required"`
			Environment string `yaml:"environment" validate:"required"`
		} `yaml:"deployments" validate:"required,dive,required"`
	} `yaml:"spec"`
}

type Component struct {
	APIVersion string `yaml:"apiVersion" validate:"required"`
	Kind       string `yaml:"kind" validate:"required,eq=Component"`
	Metadata   struct {
		Name string `yaml:"name" validate:"required"`
	} `yaml:"metadata"`
	Spec struct {
		Source struct {
			Git struct {
				Org      string `yaml:"org" validate:"required"`
				Repo     string `yaml:"repo" validate:"required"`
				Instance string `yaml:"instance" validate:"required"`
			} `yaml:"git" validate:"required"`
			Path string `yaml:"path" validate:"required"`
		} `yaml:"source" validate:"required"`
		Pipeline struct {
			Type   string `yaml:"type" validate:"required"`
			Params []struct {
				Name  string `yaml:"name" validate:"required"`
				Value string `yaml:"value" validate:"required"`
			} `yaml:"params" validate:"required,dive,required"`
		} `yaml:"pipeline" validate:"required"`
	} `yaml:"spec"`
}

type ArtefactRepository struct {
	APIVersion string `yaml:"apiVersion" validate:"required"`
	Kind       string `yaml:"kind" validate:"required,eq=ArtefactRepository"`
	Metadata   struct {
		Name string `yaml:"name" validate:"required"`
	} `yaml:"metadata"`
	Spec struct {
		Type   string `yaml:"type" validate:"required"`
		Access []struct {
			Role   string   `yaml:"role" validate:"required"`
			Users  []string `yaml:"users"`
			Groups []string `yaml:"groups"`
		} `yaml:"access"`
	} `yaml:"spec"`
}

func (e *Environment) Validate() error {
	validate := validator.New()
	return validate.Struct(e)
}

func (dr *DeploymentRepository) Validate() error {
	validate := validator.New()
	return validate.Struct(dr)
}

func (c *Component) Validate() error {
	validate := validator.New()
	return validate.Struct(c)
}

func (ar *ArtefactRepository) Validate() error {
	validate := validator.New()
	err := validate.Struct(ar)
	if err != nil {
		return err
	}
	for _, access := range ar.Spec.Access {
		if len(access.Users) == 0 && len(access.Groups) == 0 {
			return fmt.Errorf("at least one user or group must be provided in ArtefactRepository: %s", ar.Metadata.Name)
		}
	}
	return nil
}
