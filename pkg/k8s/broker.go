package k8s

import (
	"context"
	"fmt"

	v1 "knative.dev/eventing/pkg/client/clientset/versioned/typed/eventing/v1"
)

const (
	BrokerNameArgName      ArgumentName = "name"
	BrokerNamespaceArgName ArgumentName = "namespace"
)

type CreateBrokerHandler struct {
	client *v1.EventingV1Client
}

func (bh CreateBrokerHandler) Handle(ctx context.Context, args Arguments) error {
	fmt.Printf("Creating Broker %s in namespace %s\n", args.get("name", ""), args.get("namespace", "default"))
	fmt.Printf("Creating Broker %s in namespace %s\n", args.get(BrokerNameArgName, ""), args.get(BrokerNamespaceArgName, "default"))
	fmt.Printf("Creating Broker %s\n", args.get(BrokerNameArgName))

	return nil
}

func NewCreateBrokerHandler(client *v1.EventingV1Client) *CreateBrokerHandler {
	return &CreateBrokerHandler{client: client}
}
