package plugin

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/creydr/knative-kn-do-plugin/pkg/config"
	"github.com/creydr/knative-kn-do-plugin/pkg/k8s"
	"github.com/creydr/knative-kn-do-plugin/pkg/openaiapi"
	"github.com/creydr/knative-kn-do-plugin/pkg/openaiapi/function"
	"github.com/openai/openai-go"
)

func Run(message string) error {
	config := config.NewFromEnv()
	client := openaiapi.NewClient(config)

	ctx := context.Background()

	fmt.Printf("so you want to %s...\n", message)

	params := openai.ChatCompletionNewParams{
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(message),
		},
		Seed:  openai.Int(0),
		Model: config.Model,
	}

	mappings := Mappings()
	for _, mapping := range mappings {
		params.Tools = append(params.Tools, function.WrapIntoChatCompletionToolParam(mapping.FunctionDefinitionParam))
	}

	// Make initial chat completion request
	completion, err := client.Chat.Completions.New(ctx, params)
	if err != nil {
		return fmt.Errorf("failed to execute API request: %w", err)
	}

	toolCalls := completion.Choices[0].Message.ToolCalls

	// Return early if there are no tool calls
	if len(toolCalls) == 0 {
		fmt.Printf("No function call")
		return nil
	}

	for _, toolCall := range toolCalls {
		funcName := toolCall.Function.Name
		mapping, ok := mappings[funcName]
		if !ok {
			return fmt.Errorf("function %s is not defined in mapping", funcName)
		}

		var args k8s.Arguments
		if err := json.Unmarshal([]byte(toolCall.Function.Arguments), &args); err != nil {
			return fmt.Errorf("failed to extract function argument: %w", err)
		}

		if err := mapping.K8sFunc(args); err != nil {
			return fmt.Errorf("failed to call mapping function for %s: %w", funcName, err)
		}
	}

	return nil
}
