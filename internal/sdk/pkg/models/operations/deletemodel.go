// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"terraform/internal/sdk/pkg/models/shared"
)

type DeleteModelRequest struct {
	// The model to delete
	Model string `pathParam:"style=simple,explode=false,name=model"`
}

type DeleteModelResponse struct {
	ContentType string
	// OK
	DeleteModelResponse *shared.DeleteModelResponse
	StatusCode          int
	RawResponse         *http.Response
}