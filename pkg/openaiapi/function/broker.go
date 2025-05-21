package function

import "github.com/creydr/knative-kn-do-plugin/pkg/k8s"

func CreateBroker() *FunctionData {
	return &FunctionData{
		Description: "Creates a Knative Broker in the cluster",
		Parameters: Parameters{
			Parameter{
				Name:        k8s.BrokerNameArgName,
				Type:        "string",
				Description: "The name of the Broker",
				Required:    true,
			},
			Parameter{
				Name:        k8s.BrokerNamespaceArgName,
				Type:        "string",
				Description: "The Kubernetes namespace where the Broker should be created in",
			},
			Parameter{
				Name:        k8s.BrokerClassArgName,
				Type:        "string",
				Description: "The class or type of the Broker",
			},
		},
	}
}
