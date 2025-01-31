package openAI

import (
	"errors"

	"github.com/tiktoken-go/tokenizer"
)

type OpenAIHandler struct {
	Codec tokenizer.Codec
}

var ErrModelNotFound = errors.New("OpenAI model not found, or the model you gave is not supported")

func NewOpenAIHandler(modelName string) (*OpenAIHandler, error) {
	model, err := modelFromString(modelName)
	if err != nil {
		return nil, err
	}

	codec, err := tokenizer.ForModel(model)
	if err != nil {
		return nil, err
	}

	return &OpenAIHandler{
		Codec: codec,
	}, nil
}

func (t *OpenAIHandler) Encode(content string) ([]uint, []string, error) {
	return t.Codec.Encode(content)
}

func (t *OpenAIHandler) Decode(tokens []uint) (string, error) {
	return t.Codec.Decode(tokens)
}

func ModelExists(modelName string) bool {
	if _, err := modelFromString(modelName); err == nil {
		return true
	}
	return false
}

func modelFromString(modelStr string) (tokenizer.Model, error) {
	modelMap := map[string]tokenizer.Model{
		"o1-preview":    tokenizer.O1Preview,
		"o1-mini":       tokenizer.O1Mini,
		"gpt-4o":        tokenizer.GPT4o,
		"gpt-4":         tokenizer.GPT4,
		"gpt-3.5-turbo": tokenizer.GPT35Turbo,
		"gpt-3.5":       tokenizer.GPT35,
	}

	if model, exists := modelMap[modelStr]; exists {
		return model, nil
	}

	return "", ErrModelNotFound
}
