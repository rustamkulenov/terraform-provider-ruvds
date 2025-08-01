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

// DatacenterDataSource is a data source for retrieving information about datacenter by its code.
type DatacenterDataSource struct {
	client *api.Client
}

func NewDatacenterDataSource() datasource.DataSource {
	return &DatacenterDataSource{}
}

var _ datasource.DataSource = &DatacenterDataSource{}

// DatacenterDataSourceModel describes the data source data model.
type DatacenterDataSourceModel struct {
	Id      types.Int32  `tfsdk:"id"`
	Name    types.String `tfsdk:"name"`
	Code    types.String `tfsdk:"code"`
	Country types.String `tfsdk:"country"`
	// WithCode is the input code to find the datacenter.
	// It is not stored in the state.
	WithCode types.String `tfsdk:"with_code"`
}

func (d *DatacenterDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_datacenter"
}

func (d *DatacenterDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Datacenter data source",

		Attributes: map[string]schema.Attribute{
			"id": schema.Int32Attribute{
				MarkdownDescription: "Data Center identifier",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Data Center name",
				Computed:            true,
			},
			"code": schema.StringAttribute{
				MarkdownDescription: "Data Center code defined from its name",
				Computed:            true,
			},
			"country": schema.StringAttribute{
				MarkdownDescription: "Data Center country code defined from its name",
				Computed:            true,
			},
			"with_code": schema.StringAttribute{
				MarkdownDescription: "Data Center code to find",
				Required:            true,
			},
		},
	}
}

func (d *DatacenterDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *DatacenterDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data DatacenterDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	dcs, err := d.client.GetDataCenters()
	if err != nil {
		resp.Diagnostics.AddError(
			"Client Error",
			fmt.Sprintf("Unable to read datacenter, got error: %s", err),
		)
		return
	}
	for _, dc := range dcs.DataCenters {
		if dc.GetDatacenterCode() == data.WithCode.ValueString() {
			data.Id = types.Int32Value(dc.ID)
			data.Name = types.StringValue(dc.Name)
			data.Code = types.StringValue(dc.GetDatacenterCode())
			data.Country = types.StringValue(dc.GetDatacenterCountryCode())
			break
		}
	}

	if data.Id.IsNull() {
		resp.Diagnostics.AddError(
			"Datacenter Not Found",
			fmt.Sprintf("Datacenter with code '%s' not found", data.WithCode.ValueString()),
		)
		return
	}

	// Write logs using the tflog package
	// Documentation: https://terraform.io/plugin/log
	tflog.Trace(ctx, "read a data source")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
