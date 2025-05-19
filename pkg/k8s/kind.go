package k8s

import (
	"context"
	"fmt"

	"k8s.io/client-go/dynamic"
)

const (
	KindArgName          ArgumentName = "kind"
	KindNameArgName      ArgumentName = "name"
	KindNamespaceArgName ArgumentName = "namespace"
)

type DeleteKindHandler struct {
	client dynamic.Interface
}

func (kh DeleteKindHandler) Handle(ctx context.Context, args Arguments) error {
	fmt.Printf("Deleting Kind %s with name %s in namespace %s\n", args.get(KindArgName, ""), args.get(KindNameArgName, ""), args.get(KindNamespaceArgName, "default"))

	return nil
}

func NewDeleteKindHandler(client dynamic.Interface) *DeleteKindHandler {
	return &DeleteKindHandler{client: client}
}
