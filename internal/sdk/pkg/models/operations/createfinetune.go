// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"openai/internal/sdk/pkg/models/shared"
)

type CreateFineTuneResponse struct {
	ContentType string
	// OK
	FineTune    *shared.FineTune
	StatusCode  int
	RawResponse *http.Response
}
