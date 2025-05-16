package plugin

import (
	"fmt"

	"github.com/creydr/knative-kn-do-plugin/pkg/k8s"
	"github.com/creydr/knative-kn-do-plugin/pkg/openaiapi/function"
	"github.com/openai/openai-go"
	"k8s.io/client-go/dynamic"
	eventingv1 "knative.dev/eventing/pkg/client/clientset/versioned/typed/eventing/v1"
)

type toolMapping struct {
	FunctionDefinitionParam openai.FunctionDefinitionParam
	K8sHandler              k8s.Handler
}

type ToolMappings map[string]toolMapping

func (m ToolMappings) add(funcName string, funcData *function.FunctionData, k8sHandler k8s.Handler) {
	m[funcName] = toolMapping{
		FunctionDefinitionParam: funcData.ToFunctionDefinitionParam(funcName),
		K8sHandler:              k8sHandler,
	}
}

// Mappings returns a map for the mapping between a tool function for the OpenAI tool request and the corresponding
// function to handle the request against the Kubernetes API
func Mappings() (ToolMappings, error) {
	m := make(ToolMappings)

	restConfig, err := k8s.GetRestConfig()
	if err != nil {
		return nil, fmt.Errorf("could not create RestConfig: %w", err)
	}

	eventingClient, err := eventingv1.NewForConfig(restConfig)
	if err != nil {
		return nil, fmt.Errorf("could not create EventingClient: %w", err)
	}
	dynamicClient, err := dynamic.NewForConfig(restConfig)
	if err != nil {
		return nil, fmt.Errorf("could not create DynamicClient: %w", err)
	}

	m.add("create_broker", function.CreateBroker(), k8s.NewCreateBrokerHandler(eventingClient))
	m.add("delete_kind", function.DeleteKind(), k8s.NewDeleteKindHandler(dynamicClient))

	return m, nil
}
