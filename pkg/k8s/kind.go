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
	fmt.Printf("Deleting Kind %s with name %s\n", args.get(KindArgName), args.get(KindNameArgName))

	return nil
}

func NewDeleteKindHandler(client dynamic.Interface) *DeleteKindHandler {
	return &DeleteKindHandler{client: client}
}
