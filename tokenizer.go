package tokenizer

import (
	"errors"

	"github.com/weave-labs/tokenizer/handler/huggingface"
	"github.com/weave-labs/tokenizer/handler/openAI"
)

type ITokenizerHandler interface {
	Encode(content string) ([]uint, []string, error)
	Decode(tokens []uint) (string, error)
}

type TokenizerService struct {
	Tokenizer ITokenizerHandler
}

var ErrModelNotFound = errors.New("model not found")

func NewTokenizerService(modelName string) (*TokenizerService, error) {

	if openAI.ModelExists(modelName) {
		tokenizer, err := openAI.NewOpenAIHandler(modelName)
		if err != nil {
			return nil, err
		}
		return &TokenizerService{
			Tokenizer: tokenizer,
		}, nil
	}

	if huggingface.ModelExists(modelName) {
		tokenizer, err := huggingface.NewHuggingfaceHandler(modelName)
		if err != nil {
			return nil, err
		}
		return &TokenizerService{
			Tokenizer: tokenizer,
		}, nil
	}

	return nil, ErrModelNotFound
}
