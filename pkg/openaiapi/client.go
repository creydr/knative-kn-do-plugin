package openaiapi

import (
	"github.com/creydr/knative-kn-do-plugin/pkg/config"
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

func NewClient(clientConfig config.Config) openai.Client {
	return openai.NewClient(
		option.WithBaseURL(clientConfig.BaseUrl),
		option.WithAPIKey(clientConfig.ApiKey),
	)
}
