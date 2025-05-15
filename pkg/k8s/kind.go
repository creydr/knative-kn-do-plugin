package k8s

import "fmt"

func DeleteKind(args Arguments) error {
	fmt.Printf("Deleting Kind %s with name %s in namespace %s\n", args.get("kind", ""), args.get("name", ""), args.get("namespace", "default"))

	return nil
}
