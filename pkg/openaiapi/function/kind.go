package function

func DeleteKind() *FunctionData {
	return &FunctionData{
		Description: "Deletes a Knative resource (e.g. a Broker) from the cluster",
		Parameters: Parameters{
			Parameter{
				Name:        "kind",
				Type:        "string",
				Description: "The Kubernetes type of the resource which should be deleted",
			},
			Parameter{
				Name:        "name",
				Type:        "string",
				Description: "The name of the resource",
				Required:    true,
			},
			Parameter{
				Name:        "namespace",
				Type:        "string",
				Description: "The Kubernetes namespace where the resource should be deleted from",
			},
		},
	}
}
