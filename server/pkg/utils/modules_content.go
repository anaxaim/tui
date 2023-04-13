package utils

import (
	"errors"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-config-inspect/tfconfig"

	"github.com/anaxaim/tui/server/pkg/model"
)

var ErrNotFoundVariable = errors.New("variable not found in content")

func LoadTFModule(moduleDir string) (*model.ParsedContent, error) {
	tfModule, diags := tfconfig.LoadModule(moduleDir)
	if diags.HasErrors() {
		return nil, fmt.Errorf("failed to load Terraform module: %w", diags)
	}

	resources := make(map[string]*model.ParsedResource, len(tfModule.ManagedResources))

	for _, v := range tfModule.ManagedResources {
		fullName := fmt.Sprintf("%s.%s", v.Type, v.Name)
		res := &model.ParsedResource{
			Provider: v.Provider.Name,
		}
		resources[fullName] = res
	}

	dataSources := make(map[string]*model.ParsedResource, len(tfModule.DataResources))

	for _, v := range tfModule.DataResources {
		fullName := fmt.Sprintf("%s.%s", v.Type, v.Name)
		dataS := &model.ParsedResource{
			Provider: v.Provider.Name,
		}
		dataSources[fullName] = dataS
	}

	variables := make(map[string]*model.ParsedVariable, len(tfModule.Variables))

	for _, v := range tfModule.Variables {
		variable := &model.ParsedVariable{
			Name:        v.Name,
			Type:        v.Type,
			Description: v.Description,
			Default:     v.Default,
			Required:    v.Required,
			Sensitive:   v.Sensitive,
		}
		variables[v.Name] = variable
	}

	outputs := make(map[string]*model.ParsedOutput, len(tfModule.Outputs))

	for _, v := range tfModule.Outputs {
		output := &model.ParsedOutput{
			Name:        v.Name,
			Description: v.Description,
			Sensitive:   v.Sensitive,
		}
		outputs[v.Name] = output
	}

	parsedContent := &model.ParsedContent{
		Variables:   variables,
		Outputs:     outputs,
		Providers:   tfModule.RequiredProviders,
		Resources:   resources,
		DataSources: dataSources,
	}

	return parsedContent, nil
}

func UpdateContentVariables(contentVariables string, moduleVariables []model.Variable) (string, error) {
	newContent := ""

	for _, moduleVariable := range moduleVariables {
		varName := fmt.Sprintf("variable \"%s\"", moduleVariable.Name)
		varDefault := fmt.Sprintf("default = \"%s\"", moduleVariable.DefaultValue)

		varNameIndex := strings.Index(contentVariables, varName)
		if varNameIndex != -1 {
			curlyBraceIndex := strings.Index(contentVariables[varNameIndex:], "{")
			newLineIndex := strings.Index(contentVariables[varNameIndex+curlyBraceIndex:], "\n")

			newContent = contentVariables[:varNameIndex+curlyBraceIndex+newLineIndex+1] + "  " + varDefault + "\n" + contentVariables[varNameIndex+curlyBraceIndex+newLineIndex+1:]
		} else {
			return "", fmt.Errorf("%w: %s", ErrNotFoundVariable, moduleVariable.Name)
		}
	}

	return newContent, nil
}
