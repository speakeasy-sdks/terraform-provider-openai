// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"terraform/internal/sdk/pkg/models/shared"
)

type CreateImageResponse struct {
	ContentType string
	// OK
	ImagesResponse *shared.ImagesResponse
	StatusCode     int
	RawResponse    *http.Response
}
