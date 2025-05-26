package plugin

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/briandowns/spinner"
	"github.com/creydr/knative-kn-do-plugin/pkg/config"
	"github.com/creydr/knative-kn-do-plugin/pkg/k8s"
	"github.com/creydr/knative-kn-do-plugin/pkg/openaiapi"
	"github.com/creydr/knative-kn-do-plugin/pkg/openaiapi/function"
	"github.com/openai/openai-go"
)

func Run(message string) error {
	ctx := context.Background()

	config := config.NewFromEnv()
	client := openaiapi.NewClient(config)

	fmt.Printf("so you want to %s...\n", message)

	params := openai.ChatCompletionNewParams{
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(message),
		},
		Seed:  openai.Int(0),
		Model: config.Model,
	}

	mappings, err := Mappings()
	if err != nil {
		return fmt.Errorf("could not setup k8s-function to llm-function mappings: %w", err)
	}

	for _, mapping := range mappings {
		params.Tools = append(params.Tools, function.WrapIntoChatCompletionToolParam(mapping.FunctionDefinitionParam))
	}

	chatCompletionSpinner := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	chatCompletionSpinner.Suffix = " Trying to understand what you said...\n"
	chatCompletionSpinner.FinalMSG = "I think I know what you want me to do...\n"
	chatCompletionSpinner.Start()

	// Make initial chat completion request
	completion, err := client.Chat.Completions.New(ctx, params)
	if err != nil {
		chatCompletionSpinner.Start()
		return fmt.Errorf("failed to execute API request: %w", err)
	}
	chatCompletionSpinner.Stop()

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

		if err := mapping.K8sHandler.Handle(ctx, args); err != nil {
			return fmt.Errorf("failed to call mapping function for %s: %w", funcName, err)
		}
	}

	return nil
}
