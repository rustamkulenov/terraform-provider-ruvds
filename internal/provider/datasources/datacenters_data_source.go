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

// DatacentersDataSource is a data source for retrieving information about datacenters.
type DatacentersDataSource struct {
	client *api.Client
}

func NewDatacentersDataSource() datasource.DataSource {
	return &DatacentersDataSource{}
}

var _ datasource.DataSource = &DatacentersDataSource{}

// DatacentersDataSourceModel describes the data source data model.
type DatacentersDataSourceModel struct {
	Codes     types.List   `tfsdk:"codes"`
	InCountry types.String `tfsdk:"in_country"`
}

func (d *DatacentersDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_datacenters"
}

func (d *DatacentersDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Datacenters data source",

		Attributes: map[string]schema.Attribute{
			"in_country": schema.StringAttribute{
				MarkdownDescription: "Country code to filter datacenters",
				Optional:            true,
			},
			"codes": schema.ListAttribute{
				MarkdownDescription: "Data Center codes",
				Computed:            true,
				ElementType:         types.StringType,
			},
		},
	}
}

func (d *DatacentersDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *DatacentersDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data DatacentersDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	dcs, err := d.client.GetDataCenters()
	if err != nil {
		resp.Diagnostics.AddError(
			"Client Error",
			fmt.Sprintf("Unable to read datacenters, got error: %s", err),
		)
		return
	}

	codes := make([]string, 0, len(dcs.DataCenters))
	for _, dc := range dcs.DataCenters {
		if data.InCountry.IsNull() || dc.GetDatacenterCountryCode() == data.InCountry.ValueString() {
			codes = append(codes, dc.GetDatacenterCode())
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
