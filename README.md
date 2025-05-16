# kn-do-plugin

`kn-do-plugin` is a plugin of Knative Client to manage resources based on natural language.

## Description

With this plugin, you don't need to be aware of the `kn` commands and can manage your resources in natural language. For example like
```
$ kn do create a new broker named my-broker in the namespace my-namespace
```

## Install

```
go install github.com/creydr/knative-kn-do-plugin/cmd@latest
```

## Configuration

As the plugin uses a LLM for natural language processing, it requires an API endpoint for a model (and maybe a API token) via the following environment variables:

| env                             | description                                      | default                     |
|---------------------------------|--------------------------------------------------|-----------------------------|
| `KN_DO_OPENAI_API_ENDPOINT`     | Endpoint of an OpenAI API compatible LLM server  | `http://localhost:11434/v1` |
| `KN_DO_OPENAI_API_ENDPOINT_KEY` | API key for accessing the above API              |                             |
| `KN_DO_MODEL_NAME`              | Model name to use                                | `qwen3:1.7b`                |

### Server setup

Instead of connecting to a public endpoint, it's also possible to run a LLM locally and connect to it. For example by using [ollama](https://ollama.com):

```
ollama serve
```
and then with a model which supports function calls / tools, e.g.:
```
ollama run qwen3:1.7b
```