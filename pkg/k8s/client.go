package k8s

import (
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func GetRestConfig() (*rest.Config, error) {
	// This checks $KUBECONFIG first, then falls back to ~/.kube/config
	return clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		clientcmd.NewDefaultClientConfigLoadingRules(),
		&clientcmd.ConfigOverrides{},
	).ClientConfig()
}

func GetDefaultNamespace() string {
	cfg, err := clientcmd.NewDefaultClientConfigLoadingRules().Load()
	if err != nil {
		panic("could not load default config: " + err.Error())
	}

	ns := cfg.Contexts[cfg.CurrentContext].Namespace
	if ns == "" {
		ns = "default"
	}
	return ns
}
