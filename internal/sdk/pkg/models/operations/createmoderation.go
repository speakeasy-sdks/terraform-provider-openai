// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"terraform/internal/sdk/pkg/models/shared"
)

type CreateModerationResponse struct {
	ContentType string
	// OK
	CreateModerationResponse *shared.CreateModerationResponse
	StatusCode               int
	RawResponse              *http.Response
}
