// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"terraform/internal/sdk"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &ImageDataSource{}
var _ datasource.DataSourceWithConfigure = &ImageDataSource{}

func NewImageDataSource() datasource.DataSource {
	return &ImageDataSource{}
}

// ImageDataSource is the data source implementation.
type ImageDataSource struct {
	client *sdk.SDK
}

// ImageDataSourceModel describes the data model.
type ImageDataSourceModel struct {
	Created        types.Int64          `tfsdk:"created"`
	Data           []ImagesResponseData `tfsdk:"data"`
	N              types.Int64          `tfsdk:"n"`
	Prompt         types.String         `tfsdk:"prompt"`
	ResponseFormat types.String         `tfsdk:"response_format"`
	Size           types.String         `tfsdk:"size"`
	User           types.String         `tfsdk:"user"`
}

// Metadata returns the data source type name.
func (r *ImageDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_image"
}

// Schema defines the schema for the data source.
func (r *ImageDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Image DataSource",

		Attributes: map[string]schema.Attribute{
			"created": schema.Int64Attribute{
				Computed: true,
			},
			"data": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"b64_json": schema.StringAttribute{
							Computed: true,
						},
						"url": schema.StringAttribute{
							Computed: true,
						},
					},
				},
			},
			"n": schema.Int64Attribute{
				Optional:    true,
				Description: `The number of images to generate. Must be between 1 and 10.`,
			},
			"prompt": schema.StringAttribute{
				Required:    true,
				Description: `A text description of the desired image(s). The maximum length is 1000 characters.`,
			},
			"response_format": schema.StringAttribute{
				Optional: true,
				Validators: []validator.String{
					stringvalidator.OneOf(
						"url",
						"b64_json",
					),
				},
				MarkdownDescription: `must be one of [url, b64_json]` + "\n" +
					`The format in which the generated images are returned. Must be one of ` + "`" + `url` + "`" + ` or ` + "`" + `b64_json` + "`" + `.`,
			},
			"size": schema.StringAttribute{
				Optional: true,
				Validators: []validator.String{
					stringvalidator.OneOf(
						"256x256",
						"512x512",
						"1024x1024",
					),
				},
				MarkdownDescription: `must be one of [256x256, 512x512, 1024x1024]` + "\n" +
					`The size of the generated images. Must be one of ` + "`" + `256x256` + "`" + `, ` + "`" + `512x512` + "`" + `, or ` + "`" + `1024x1024` + "`" + `.`,
			},
			"user": schema.StringAttribute{
				Optional: true,
				MarkdownDescription: `A unique identifier representing your end-user, which can help OpenAI to monitor and detect abuse. [Learn more](/docs/guides/safety-best-practices/end-user-ids).` + "\n" +
					``,
			},
		},
	}
}

func (r *ImageDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.SDK)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected DataSource Configure Type",
			fmt.Sprintf("Expected *sdk.SDK, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *ImageDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *ImageDataSourceModel
	var item types.Object

	resp.Diagnostics.Append(req.Config.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	request := *data.ToGetSDKType()
	res, err := r.client.OpenAI.CreateImage(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if res.ImagesResponse == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromGetResponse(res.ImagesResponse)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
