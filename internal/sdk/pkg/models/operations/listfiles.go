// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"terraform/internal/sdk/pkg/models/shared"
)

type ListFilesResponse struct {
	ContentType string
	// OK
	ListFilesResponse *shared.ListFilesResponse
	StatusCode        int
	RawResponse       *http.Response
}