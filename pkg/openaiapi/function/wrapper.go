package function

import "github.com/openai/openai-go"

// WrapIntoChatCompletionToolParam simply wraps a FunctionDefinitionParam into a ChatCompletionToolParam,
// so not each function to create a ChatCompletionToolParam, needs to configure it fully and can return only the
// FunctionDefinitionParam instead.
func WrapIntoChatCompletionToolParam(fdp openai.FunctionDefinitionParam) openai.ChatCompletionToolParam {
	return openai.ChatCompletionToolParam{
		Type:     "function",
		Function: fdp,
	}
}
