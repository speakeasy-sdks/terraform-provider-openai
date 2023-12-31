// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"openai/internal/sdk/pkg/models/shared"
)

func (r *ImageDataSourceModel) ToGetSDKType() *shared.CreateImageRequest {
	n := new(int64)
	if !r.N.IsUnknown() && !r.N.IsNull() {
		*n = r.N.ValueInt64()
	} else {
		n = nil
	}
	prompt := r.Prompt.ValueString()
	responseFormat := new(shared.CreateImageRequestResponseFormat)
	if !r.ResponseFormat.IsUnknown() && !r.ResponseFormat.IsNull() {
		*responseFormat = shared.CreateImageRequestResponseFormat(r.ResponseFormat.ValueString())
	} else {
		responseFormat = nil
	}
	size := new(shared.CreateImageRequestSize)
	if !r.Size.IsUnknown() && !r.Size.IsNull() {
		*size = shared.CreateImageRequestSize(r.Size.ValueString())
	} else {
		size = nil
	}
	user := new(string)
	if !r.User.IsUnknown() && !r.User.IsNull() {
		*user = r.User.ValueString()
	} else {
		user = nil
	}
	out := shared.CreateImageRequest{
		N:              n,
		Prompt:         prompt,
		ResponseFormat: responseFormat,
		Size:           size,
		User:           user,
	}
	return &out
}

func (r *ImageDataSourceModel) RefreshFromGetResponse(resp *shared.ImagesResponse) {
	r.Created = types.Int64Value(resp.Created)
	r.Data = nil
	for _, dataItem := range resp.Data {
		var data1 ImagesResponseData
		if dataItem.B64JSON != nil {
			data1.B64JSON = types.StringValue(*dataItem.B64JSON)
		} else {
			data1.B64JSON = types.StringNull()
		}
		if dataItem.URL != nil {
			data1.URL = types.StringValue(*dataItem.URL)
		} else {
			data1.URL = types.StringNull()
		}
		r.Data = append(r.Data, data1)
	}
}
