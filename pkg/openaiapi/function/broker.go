package function

func CreateBroker() *FunctionData {
	return &FunctionData{
		Description: "Creates a Knative Broker in the cluster",
		Parameters: Parameters{
			Parameter{
				Name:        "name",
				Type:        "string",
				Description: "The name of the Broker",
				Required:    true,
			},
			Parameter{
				Name:        "namespace",
				Type:        "string",
				Description: "The Kubernetes namespace where the Broker should be created in",
			},
		},
	}
}
