package plugin

import (
	k8s2 "github.com/creydr/knative-kn-do-plugin/pkg/k8s"
	"github.com/creydr/knative-kn-do-plugin/pkg/openaiapi/function"
	"github.com/openai/openai-go"
)

type toolMapping struct {
	FunctionDefinitionParam openai.FunctionDefinitionParam
	K8sFunc                 k8s2.Func
}

type ToolMappings map[string]toolMapping

func (m ToolMappings) add(funcName string, funcData *function.FunctionData, k8sFunc k8s2.Func) {
	m[funcName] = toolMapping{
		FunctionDefinitionParam: funcData.ToFunctionDefinitionParam(funcName),
		K8sFunc:                 k8sFunc,
	}
}

// Mappings returns a map for the mapping between a tool function for the OpenAI tool request and the corresponding
// function to handle the request against the Kubernetes API
func Mappings() ToolMappings {
	m := make(ToolMappings)

	m.add("create_broker", function.CreateBroker(), k8s2.CreateBroker)
	m.add("delete_kind", function.DeleteKind(), k8s2.DeleteKind)

	return m
}
