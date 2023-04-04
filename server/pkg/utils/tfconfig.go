package utils

import (
	"fmt"

	"github.com/hashicorp/terraform-config-inspect/tfconfig"

	"github.com/anaxaim/tui/server/pkg/model"
)

func LoadTFModule(moduleDir string) (*tfconfig.Module, error) {
	tfModule, diags := tfconfig.LoadModule(moduleDir)
	if diags.HasErrors() {
		return nil, fmt.Errorf("failed to load Terraform module: %w", diags)
	}

	return tfModule, nil
}

func MergeTfVariables(moduleVariables model.Variables, tfVariables map[string]*tfconfig.Variable) map[string]*tfconfig.Variable {
	mergedVars := make(map[string]*tfconfig.Variable)

	for _, v := range tfVariables {
		mergedVars[v.Name] = v
	}

	for _, v := range moduleVariables {
		if variable, ok := mergedVars[v.Name]; ok {
			if v.Type != "" {
				variable.Type = v.Type
			}

			if v.Description != "" {
				variable.Description = v.Description
			}

			if v.DefaultValue != "" {
				variable.Default = v.DefaultValue
			}
		} else {
			mergedVars[v.Name] = &tfconfig.Variable{
				Name:        v.Name,
				Type:        v.Type,
				Description: v.Description,
				Default:     v.DefaultValue,
			}
		}
	}

	return mergedVars
}
