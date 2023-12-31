// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type CreateTranslationRequestFile struct {
	Content []byte `multipartForm:"content"`
	File    string `multipartForm:"name=file"`
}

// CreateTranslationRequestModel2 - ID of the model to use. Only `whisper-1` is currently available.
type CreateTranslationRequestModel2 string

const (
	CreateTranslationRequestModel2Whisper1 CreateTranslationRequestModel2 = "whisper-1"
)

func (e CreateTranslationRequestModel2) ToPointer() *CreateTranslationRequestModel2 {
	return &e
}

func (e *CreateTranslationRequestModel2) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "whisper-1":
		*e = CreateTranslationRequestModel2(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateTranslationRequestModel2: %v", v)
	}
}

type CreateTranslationRequest struct {
	// The audio file object (not file name) translate, in one of these formats: mp3, mp4, mpeg, mpga, m4a, wav, or webm.
	//
	File CreateTranslationRequestFile `multipartForm:"file"`
	// ID of the model to use. Only `whisper-1` is currently available.
	//
	Model interface{} `multipartForm:"name=model,json"`
	// An optional text to guide the model's style or continue a previous audio segment. The [prompt](/docs/guides/speech-to-text/prompting) should be in English.
	//
	Prompt *string `multipartForm:"name=prompt"`
	// The format of the transcript output, in one of these options: json, text, srt, verbose_json, or vtt.
	//
	ResponseFormat *string `multipartForm:"name=response_format"`
	// The sampling temperature, between 0 and 1. Higher values like 0.8 will make the output more random, while lower values like 0.2 will make it more focused and deterministic. If set to 0, the model will use [log probability](https://en.wikipedia.org/wiki/Log_probability) to automatically increase the temperature until certain thresholds are hit.
	//
	Temperature *float64 `multipartForm:"name=temperature"`
}
