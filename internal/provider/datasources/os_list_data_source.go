package datasources

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/rustamkulenov/terraform-provider-ruvds/internal/api"
)

// OSListDataSource is a data source for retrieving information about operating systems.
type OSListDataSource struct {
	client *api.Client
}

func NewOSListDataSource() datasource.DataSource {
	return &OSListDataSource{}
}

var _ datasource.DataSource = &OSListDataSource{}

// OSListDataSourceModel describes the data source data model.
type OSListDataSourceModel struct {
	// Names is a list of OS codes available in the provider.
	Codes    types.List   `tfsdk:"codes"`
	WithType types.String `tfsdk:"with_type"`
}

func (d *OSListDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_os_list"
}

func (d *OSListDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Available Operating Systems data source",

		Attributes: map[string]schema.Attribute{
			"with_type": schema.StringAttribute{
				MarkdownDescription: "Type of operating system to filter (e.g., 'linux', 'windows')",
				Optional:            true,
			},
			"codes": schema.ListAttribute{
				MarkdownDescription: "OS codes available in the provider",
				Computed:            true,
				ElementType:         types.StringType,
			},
		},
	}
}

func (d *OSListDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *OSListDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data OSListDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	oses, err := d.client.GetOSList()
	if err != nil {
		resp.Diagnostics.AddError(
			"Client Error",
			fmt.Sprintf("Unable to read OSes, got error: %s", err),
		)
		return
	}

	codes := make([]string, 0, len(oses.Items))
	for _, os := range oses.Items {
		if data.WithType.IsNull() || os.Type == data.WithType.ValueString() {
			codes = append(codes, os.GetCode())
		}
	}
	codesres, diags := types.ListValueFrom(ctx, types.StringType, codes)
	resp.Diagnostics.Append(diags...)
	data.Codes = codesres

	// Write logs using the tflog package
	// Documentation: https://terraform.io/plugin/log
	tflog.Trace(ctx, "read a data source")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
