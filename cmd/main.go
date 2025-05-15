package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/creydr/knative-kn-do-plugin/pkg/plugin"
)

func main() {
	message := ""
	if len(os.Args) >= 2 {
		message = strings.Join(os.Args[1:], " ")
	} else {
		fmt.Printf("Usage: %s [your command]\n", os.Args[0])
		os.Exit(1)
	}

	if err := plugin.Run(message); err != nil {
		panic(err)
	}
}
