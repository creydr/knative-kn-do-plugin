package function

import "github.com/creydr/knative-kn-do-plugin/pkg/k8s"

func CreateTrigger() *FunctionData {
	return &FunctionData{
		Description: "Creates a Knative Trigger in the cluster to send events from a Broker to a Destination",
		Parameters: Parameters{
			Parameter{
				Name:        k8s.TriggerNameArgName,
				Type:        "string",
				Description: "The name of the Trigger",
				Required:    true,
			},
			Parameter{
				Name:        k8s.TriggerNamespaceArgName,
				Type:        "string",
				Description: "The Kubernetes namespace where the Trigger should be created in",
			},
			Parameter{
				Name:        k8s.TriggerBrokerArgName,
				Type:        "string",
				Description: "The Broker to which the Trigger should be connected to",
				Required:    true,
			},
			Parameter{
				Name:        k8s.TriggerDestinationArgName,
				Type:        "string",
				Description: "The Destination or Sink where the events should be sent to",
				Required:    true,
			},
		},
	}
}
