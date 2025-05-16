package k8s

import (
	"context"
	"fmt"

	"k8s.io/client-go/dynamic"
)

type DeleteKindHandler struct {
	client dynamic.Interface
}

func (kh DeleteKindHandler) Handle(ctx context.Context, args Arguments) error {
	fmt.Printf("Deleting Kind %s with name %s in namespace %s\n", args.get("kind", ""), args.get("name", ""), args.get("namespace", "default"))

	return nil
}

func NewDeleteKindHandler(client dynamic.Interface) *DeleteKindHandler {
	return &DeleteKindHandler{client: client}
}
