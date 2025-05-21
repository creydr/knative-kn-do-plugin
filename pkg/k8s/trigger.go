package k8s

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	clientv1eventing "knative.dev/client/pkg/eventing/v1"
	v1 "knative.dev/eventing/pkg/client/clientset/versioned/typed/eventing/v1"
	"knative.dev/pkg/apis"
	duckv1 "knative.dev/pkg/apis/duck/v1"
)

const (
	TriggerNameArgName        ArgumentName = "name"
	TriggerNamespaceArgName   ArgumentName = "namespace"
	TriggerBrokerArgName      ArgumentName = "broker"
	TriggerDestinationArgName ArgumentName = "destination"
)

type CreateTriggerHandler struct {
	client *v1.EventingV1Client
}

func (bh CreateTriggerHandler) Handle(ctx context.Context, args Arguments) error {
	fmt.Printf("Creating Trigger %s\n", args.get(TriggerNameArgName))

	name := args.get(TriggerNameArgName).(string)
	broker := args.get(TriggerBrokerArgName).(string)
	sink := args.get(TriggerDestinationArgName).(string)
	sinkURI, err := apis.ParseURL(sink)
	if err != nil {
		return fmt.Errorf("failed to parse destination URI: %v", err)
	}
	dest := duckv1.Destination{
		URI: sinkURI,
	}

	triggerBuilder := clientv1eventing.
		NewTriggerBuilder(name).
		Broker(broker).
		Subscriber(&dest)

	if namespace := args.get(TriggerNamespaceArgName); namespace != nil {
		triggerBuilder.Namespace(namespace.(string))
	} else {
		triggerBuilder.Namespace(GetDefaultNamespace())
	}

	trigger := triggerBuilder.Build()

	_, err = bh.client.Triggers(trigger.Namespace).Create(ctx, trigger, metav1.CreateOptions{})
	if err != nil {
		return fmt.Errorf("failed to create Trigger %s: %w", name, err)
	}

	return nil
}

func NewCreateTriggerHandler(client *v1.EventingV1Client) *CreateTriggerHandler {
	return &CreateTriggerHandler{client: client}
}
