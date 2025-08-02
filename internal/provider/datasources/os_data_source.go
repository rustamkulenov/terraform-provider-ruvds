package datasources

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/rustamkulenov/terraform-provider-ruvds/internal/api"
)

// OSDataSource is a data source for retrieving information about OS by its code.
// Currently, it accepts name as a code.
type OSDataSource struct {
	client *api.Client
}

func NewOSDataSource() datasource.DataSource {
	return &OSDataSource{}
}

var _ datasource.DataSource = &OSDataSource{}

// OSDataSourceModel describes the data source data model.
type OSDataSourceModel struct {
	Id               types.Int32  `tfsdk:"id"`
	Name             types.String `tfsdk:"name"`
	Code             types.String `tfsdk:"code"`
	IsActive         types.Bool   `tfsdk:"is_active"`
	Type             types.String `tfsdk:"type"`
	SshKeysSupported types.Bool   `tfsdk:"ssh_keys_supported"`
	// WithCode is the input code (OS is and name) to find the OS.
	// It is not stored in the state.
	WithCode types.String `tfsdk:"with_code"`
}

func (d *OSDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_os"
}

func (d *OSDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "OS data source",

		Attributes: map[string]schema.Attribute{
			"id": schema.Int32Attribute{
				MarkdownDescription: "OS identifier",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "OS name",
				Computed:            true,
			},
			"code": schema.StringAttribute{
				MarkdownDescription: "OS code defined from its id and name",
				Computed:            true,
			},
			"is_active": schema.BoolAttribute{
				MarkdownDescription: "A value indicating if the OS is active",
				Computed:            true,
			},
			"type": schema.StringAttribute{
				MarkdownDescription: "Type of the OS (e.g., 'linux', 'windows')",
				Computed:            true,
			},
			"ssh_keys_supported": schema.BoolAttribute{
				MarkdownDescription: "A value indicating if the OS supports SSH keys",
				Computed:            true,
			},
			// WithCode is the input code (OS id and name) to find the OS.
			"with_code": schema.StringAttribute{
				MarkdownDescription: "Data Center code to find",
				Required:            true,
			},
		},
	}
}

func (d *OSDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*api.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *api.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	d.client = client
}

func (d *OSDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data OSDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	oses, err := d.client.GetOS()
	if err != nil {
		resp.Diagnostics.AddError(
			"Client Error",
			fmt.Sprintf("Unable to read OSes, got error: %s", err),
		)
		return
	}
	for _, os := range oses.Items {
		if os.GetCode() == strings.ToLower(data.WithCode.ValueString()) {
			data.Id = types.Int32Value(os.ID)
			data.Name = types.StringValue(os.Name)
			data.Code = types.StringValue(os.GetCode())
			data.IsActive = types.BoolValue(os.IsActive)
			data.Type = types.StringValue(os.Type)
			data.SshKeysSupported = types.BoolValue(os.SshKeysSupported)
			break
		}
	}

	if data.Id.IsNull() {
		resp.Diagnostics.AddError(
			"OS Not Found",
			fmt.Sprintf("OS with code '%s' not found", data.WithCode.ValueString()),
		)
		return
	}

	// Write logs using the tflog package
	// Documentation: https://terraform.io/plugin/log
	tflog.Trace(ctx, "read a data source")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
