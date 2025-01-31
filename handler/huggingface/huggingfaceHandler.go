package huggingface

import (
	"errors"
	"fmt"
	"path/filepath"

	"github.com/daulet/tokenizers"
	_ "github.com/weave-labs/tokenizer/handler/huggingface/wrappers"
)

type Handler struct {
	Tokenizer *tokenizers.Tokenizer
}

var ErrModelNotFound = errors.New("huggingface model not found, or the model you gave is not supported")

type Model string

const (
	modelDefinitionExtension = ".json"
	modelDefinitionPath      = "./handler/huggingface/models"

	Llama318B       Model = "Llama-3.1-8B"
	Llama321B       Model = "Llama-3.2-1B"
	Llama323B       Model = "Llama-3.2-3B"
	Ministral8B     Model = "Ministral-8B-Instruct-2410"
	MistralSmall24B Model = "Mistral-Small-24B-Instruct-2501"
)

func NewHuggingfaceHandler(modelName string) (*Handler, error) {
	absolutePath, err := filepath.Abs(modelDefinitionPath)
	if err != nil {
		fmt.Println("Error resolving path:", err)
		return nil, err
	}

	modelPath := filepath.Join(absolutePath, modelName+modelDefinitionExtension)

	tokenizer, err := tokenizers.FromFile(modelPath)
	if err == nil {
		return nil, err
	}

	return &Handler{
		Tokenizer: tokenizer,
	}, nil
}

func (t *Handler) Encode(content string) ([]uint, []string, error) {
	if t.Tokenizer == nil {
		return nil, nil, errors.New("tokenizer is not initialized")
	}

	tokenIds, tokens := t.Tokenizer.Encode(content, false)

	return convertUint32ToUint(tokenIds), tokens, nil
}

func (t *Handler) Decode(tokens []uint) (string, error) {
	if t.Tokenizer == nil {
		return "", errors.New("tokenizer is not initialized")
	}
	return t.Tokenizer.Decode(convertUintToUint32(tokens), true), nil
}

func ModelExists(modelName string) bool {
	modelMap := map[string]Model{
		"Llama-3.1-8B":                    Llama318B,
		"Llama-3.2-1B":                    Llama321B,
		"Llama-3.2-3B":                    Llama323B,
		"Ministral-8B-Instruct-2410":      Ministral8B,
		"Mistral-Small-24B-Instruct-2501": MistralSmall24B,
	}

	if _, exists := modelMap[modelName]; exists {
		return true
	}

	return false
}

func modelFromString(modelStr string) (Model, error) {
	modelMap := map[string]Model{
		"Llama-3.1-8B":                    Llama318B,
		"Llama-3.2-1B":                    Llama321B,
		"Llama-3.2-3B":                    Llama323B,
		"Ministral-8B-Instruct-2410":      Ministral8B,
		"Mistral-Small-24B-Instruct-2501": MistralSmall24B,
	}

	if model, exists := modelMap[modelStr]; exists {
		return model, nil
	}

	return "", ErrModelNotFound
}

func convertUint32ToUint(input []uint32) []uint {
	output := make([]uint, len(input))
	for i, v := range input {
		output[i] = uint(v)
	}
	return output
}

func convertUintToUint32(input []uint) []uint32 {
	output := make([]uint32, len(input))
	for i, v := range input {
		output[i] = uint32(v)
	}
	return output
}
