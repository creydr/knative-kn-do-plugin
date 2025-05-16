package k8s

import "context"

type Handler interface {
	Handle(context.Context, Arguments) error
}
