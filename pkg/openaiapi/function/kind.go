package function

import "github.com/creydr/knative-kn-do-plugin/pkg/k8s"

func DeleteKind() *FunctionData {
	return &FunctionData{
		Description: "Deletes a Knative resource (e.g. a Broker) from the cluster",
		Parameters: Parameters{
			Parameter{
				Name:        k8s.KindArgName,
				Type:        "string",
				Description: "The Kubernetes type of the resource which should be deleted",
			},
			Parameter{
				Name:        k8s.KindNameArgName,
				Type:        "string",
				Description: "The name of the resource",
				Required:    true,
			},
			Parameter{
				Name:        k8s.KindNamespaceArgName,
				Type:        "string",
				Description: "The Kubernetes namespace where the resource should be deleted from",
			},
		},
	}
}
