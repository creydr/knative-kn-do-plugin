package function

import (
	"github.com/creydr/knative-kn-do-plugin/pkg/k8s"
	"github.com/openai/openai-go"
)

type FunctionData struct {
	Description string
	Parameters  Parameters
}

func (funcData *FunctionData) ToFunctionDefinitionParam(funcName string) openai.FunctionDefinitionParam {
	return openai.FunctionDefinitionParam{
		Name:        funcName,
		Description: openai.String(funcData.Description),
		Parameters:  funcData.Parameters.ToFunctionParameters(),
	}
}

type Parameters []Parameter

type Parameter struct {
	Name        k8s.ArgumentName
	Type        string
	Description string
	Required    bool
}

func (f Parameters) ToFunctionParameters() openai.FunctionParameters {
	fp := openai.FunctionParameters{
		"type": "object",
	}

	properties := map[string]interface{}{}
	var required []string
	for _, f := range f {
		properties[string(f.Name)] = map[string]string{
			"type":        f.Type,
			"description": f.Description,
		}

		if f.Required {
			required = append(required, string(f.Name))
		}
	}

	fp["properties"] = properties
	fp["required"] = required

	return fp
}
