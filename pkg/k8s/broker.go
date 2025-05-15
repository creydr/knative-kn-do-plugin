package k8s

import (
	"fmt"
)

func CreateBroker(args Arguments) error {
	fmt.Printf("Creating Broker %s in namespace %s\n", args.get("name", ""), args.get("namespace", "default"))

	return nil
}
