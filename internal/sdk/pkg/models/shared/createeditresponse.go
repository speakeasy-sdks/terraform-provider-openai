// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type CreateEditResponseChoicesFinishReason string

const (
	CreateEditResponseChoicesFinishReasonStop   CreateEditResponseChoicesFinishReason = "stop"
	CreateEditResponseChoicesFinishReasonLength CreateEditResponseChoicesFinishReason = "length"
)

func (e CreateEditResponseChoicesFinishReason) ToPointer() *CreateEditResponseChoicesFinishReason {
	return &e
}

func (e *CreateEditResponseChoicesFinishReason) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "stop":
		fallthrough
	case "length":
		*e = CreateEditResponseChoicesFinishReason(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateEditResponseChoicesFinishReason: %v", v)
	}
}

type CreateEditResponseChoices struct {
	FinishReason CreateEditResponseChoicesFinishReason `json:"finish_reason"`
	Index        int64                                 `json:"index"`
	Text         string                                `json:"text"`
}

type CreateEditResponseUsage struct {
	CompletionTokens int64 `json:"completion_tokens"`
	PromptTokens     int64 `json:"prompt_tokens"`
	TotalTokens      int64 `json:"total_tokens"`
}

// CreateEditResponse - OK
type CreateEditResponse struct {
	Choices []CreateEditResponseChoices `json:"choices"`
	Created int64                       `json:"created"`
	Object  string                      `json:"object"`
	Usage   CreateEditResponseUsage     `json:"usage"`
}
