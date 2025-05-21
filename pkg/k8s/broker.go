package k8s

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	clientv1eventing "knative.dev/client/pkg/eventing/v1"
	v1 "knative.dev/eventing/pkg/client/clientset/versioned/typed/eventing/v1"
)

const (
	BrokerNameArgName      ArgumentName = "name"
	BrokerNamespaceArgName ArgumentName = "namespace"
	BrokerClassArgName     ArgumentName = "brokerclass"
)

type CreateBrokerHandler struct {
	client *v1.EventingV1Client
}

func (bh CreateBrokerHandler) Handle(ctx context.Context, args Arguments) error {
	fmt.Printf("Creating Broker %s\n", args.get(BrokerNameArgName))

	name := args.get(BrokerNameArgName)
	brokerBuilder := clientv1eventing.NewBrokerBuilder(name.(string))

	if namespace := args.get(BrokerNamespaceArgName); namespace != nil {
		brokerBuilder.Namespace(namespace.(string))
	} else {
		brokerBuilder.Namespace(GetDefaultNamespace())
	}
	if brokerClass := args.get(BrokerClassArgName); brokerClass != nil {
		brokerBuilder.Class(brokerClass.(string))
	}

	broker := brokerBuilder.Build()

	_, err := bh.client.Brokers(broker.Namespace).Create(ctx, broker, metav1.CreateOptions{})
	if err != nil {
		return fmt.Errorf("failed to create Broker %s: %w", broker.Name, err)
	}

	return nil
}

func NewCreateBrokerHandler(client *v1.EventingV1Client) *CreateBrokerHandler {
	return &CreateBrokerHandler{client: client}
}
