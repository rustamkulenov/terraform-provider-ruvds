package datasources

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/rustamkulenov/terraform-provider-ruvds/internal/api"
)

// VpsListDataSource is a data source for retrieving information about operating VPS.
type VpsListDataSource struct {
	client *api.Client
}

func NewVpsListDataSource() datasource.DataSource {
	return &VpsListDataSource{}
}

var _ datasource.DataSource = &VpsListDataSource{}

// VpsListDataSourceModel describes the data source data model.
type VpsListDataSourceModel struct {
	// List of servers.
	Servers types.List `tfsdk:"servers"`
}

type VpsModel struct {
	ID                      types.Int64   `tfsdk:"id"`
	Status                  types.String  `tfsdk:"status"`
	CreateProgress          types.Int64   `tfsdk:"create_progress"`
	DataCenterID            types.Int64   `tfsdk:"datacenter_id"`
	TariffID                types.Int64   `tfsdk:"tariff_id"`
	PaymentPeriod           types.Int64   `tfsdk:"payment_period"`
	OSID                    types.Int64   `tfsdk:"os_id"`
	CPU                     types.Int64   `tfsdk:"cpu"`
	RAM                     types.Float64 `tfsdk:"ram"`
	VRAM                    types.Int64   `tfsdk:"vram"`
	Drive                   types.Int64   `tfsdk:"drive"`
	DriveTariffID           types.Int64   `tfsdk:"drive_tariff_id"`
	IP                      types.Int64   `tfsdk:"ip"`
	DDOSProtection          types.Float32 `tfsdk:"ddos_protection"`
	PaidTill                types.String  `tfsdk:"paid_till"`
	TemplateID              types.String  `tfsdk:"template_id"`
	AdditionalDrive         types.Int64   `tfsdk:"additional_drive"`
	AdditionalDriveTariffID types.Int64   `tfsdk:"additional_drive_tariff_id"`
	UserComment             types.String  `tfsdk:"user_comment"`
}

func (d *VpsListDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vps_list"
}

func (d *VpsListDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "VPS list data source",

		Attributes: map[string]schema.Attribute{
			"servers": schema.ListNestedAttribute{
				MarkdownDescription: "List of virtual servers",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id":                         schema.Int64Attribute{Computed: true},
						"status":                     schema.StringAttribute{Computed: true},
						"create_progress":            schema.Int64Attribute{Computed: true},
						"datacenter_id":              schema.Int64Attribute{Computed: true},
						"tariff_id":                  schema.Int64Attribute{Computed: true},
						"payment_period":             schema.Int64Attribute{Computed: true},
						"os_id":                      schema.Int64Attribute{Computed: true},
						"cpu":                        schema.Int64Attribute{Computed: true},
						"ram":                        schema.Float64Attribute{Computed: true},
						"vram":                       schema.Int64Attribute{Computed: true},
						"drive":                      schema.Int64Attribute{Computed: true},
						"drive_tariff_id":            schema.Int64Attribute{Computed: true},
						"ip":                         schema.Int64Attribute{Computed: true},
						"ddos_protection":            schema.Float32Attribute{Computed: true},
						"paid_till":                  schema.StringAttribute{Computed: true},
						"template_id":                schema.StringAttribute{Computed: true},
						"additional_drive":           schema.Int64Attribute{Computed: true},
						"additional_drive_tariff_id": schema.Int64Attribute{Computed: true},
						"user_comment":               schema.StringAttribute{Computed: true},
					},
				},
			},
		},
	}
}

func (d *VpsListDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *VpsListDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data VpsListDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	vpses, err := d.client.GetVpsList()
	if err != nil {
		resp.Diagnostics.AddError(
			"Client Error",
			fmt.Sprintf("Unable to read VPS list, got error: %s", err),
		)
		return
	}

	servers := make([]VpsModel, 0, len(vpses.VirtualServers))
	for _, srv := range vpses.VirtualServers {
		vps := VpsModel{
			ID:             types.Int64Value(int64(srv.ID)),
			CreateProgress: types.Int64Value(int64(srv.CreateProgress)),
			DataCenterID:   types.Int64Value(int64(srv.DataCenterId)),
			TariffID:       types.Int64Value(int64(srv.TariffId)),
			PaymentPeriod:  types.Int64Value(int64(srv.PaymentPeriod)),
			OSID:           types.Int64Value(int64(srv.OSId)),
			CPU:            types.Int64Value(int64(srv.CPU)),
			RAM:            types.Float64Value(float64(srv.RAM)),
			VRAM:           types.Int64Value(int64(srv.VRAM)),
			Drive:          types.Int64Value(int64(srv.Drive)),
			DriveTariffID:  types.Int64Value(int64(srv.DriveTariffId)),
			IP:             types.Int64Value(int64(srv.IP)),
			DDOSProtection: types.Float32Value(srv.DDOSProtection),
			PaidTill:       types.StringValue(srv.PaidTill),
		}
		if srv.Status != nil {
			vps.Status = types.StringValue(*srv.Status)
		} else {
			vps.Status = types.StringNull()
		}
		if srv.TemplateId != nil {
			vps.TemplateID = types.StringValue(*srv.TemplateId)
		} else {
			vps.TemplateID = types.StringNull()
		}
		if srv.UserComment != nil {
			vps.UserComment = types.StringValue(*srv.UserComment)
		} else {
			vps.UserComment = types.StringNull()
		}
		if srv.AdditionalDrive == nil {
			vps.AdditionalDrive = types.Int64Null()
		} else {
			vps.AdditionalDrive = types.Int64Value(int64(*srv.AdditionalDrive))
		}
		if srv.AdditionalDriveTariffId == nil {
			vps.AdditionalDriveTariffID = types.Int64Null()
		} else {
			vps.AdditionalDriveTariffID = types.Int64Value(int64(*srv.AdditionalDriveTariffId))
		}
		servers = append(servers, vps)
	}

	vpsListType := types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id":                         types.Int64Type,
			"status":                     types.StringType,
			"create_progress":            types.Int64Type,
			"datacenter_id":              types.Int64Type,
			"tariff_id":                  types.Int64Type,
			"payment_period":             types.Int64Type,
			"os_id":                      types.Int64Type,
			"cpu":                        types.Int64Type,
			"ram":                        types.Float64Type,
			"vram":                       types.Int64Type,
			"drive":                      types.Int64Type,
			"drive_tariff_id":            types.Int64Type,
			"ip":                         types.Int64Type,
			"ddos_protection":            types.Float32Type,
			"paid_till":                  types.StringType,
			"template_id":                types.StringType,
			"additional_drive":           types.Int64Type,
			"additional_drive_tariff_id": types.Int64Type,
			"user_comment":               types.StringType,
		},
	}
	vpsres, diags := types.ListValueFrom(ctx, vpsListType, servers)
	resp.Diagnostics.Append(diags...)
	data.Servers = vpsres

	// Write logs using the tflog package
	// Documentation: https://terraform.io/plugin/log
	tflog.Trace(ctx, "read a data source")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
