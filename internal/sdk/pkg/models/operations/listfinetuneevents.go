// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"terraform/internal/sdk/pkg/models/shared"
)

type ListFineTuneEventsRequest struct {
	// The ID of the fine-tune job to get events for.
	//
	FineTuneID string `pathParam:"style=simple,explode=false,name=fine_tune_id"`
	// Whether to stream events for the fine-tune job. If set to true,
	// events will be sent as data-only
	// [server-sent events](https://developer.mozilla.org/en-US/docs/Web/API/Server-sent_events/Using_server-sent_events#Event_stream_format)
	// as they become available. The stream will terminate with a
	// `data: [DONE]` message when the job is finished (succeeded, cancelled,
	// or failed).
	//
	// If set to false, only events generated so far will be returned.
	//
	Stream *bool `queryParam:"style=form,explode=true,name=stream"`
}

type ListFineTuneEventsResponse struct {
	ContentType string
	// OK
	ListFineTuneEventsResponse *shared.ListFineTuneEventsResponse
	StatusCode                 int
	RawResponse                *http.Response
}
