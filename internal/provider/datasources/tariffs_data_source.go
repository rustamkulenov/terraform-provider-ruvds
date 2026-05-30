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

// TariffsDataSource is a data source for retrieving all tariffs from the RUVDS API.
type TariffsDataSource struct {
	client *api.Client
}

func NewTariffsDataSource() datasource.DataSource {
	return &TariffsDataSource{}
}

var _ datasource.DataSource = &TariffsDataSource{}

// TariffsDataSourceModel describes the data source data model.
type TariffsDataSourceModel struct {
	Vps                    types.List `tfsdk:"vps"`
	Drive                  types.List `tfsdk:"drive"`
	AdditionalDrive        types.List `tfsdk:"additional_drive"`
	AdditionalService      types.List `tfsdk:"additional_service"`
	PaymentPeriodDiscounts types.List `tfsdk:"payment_period_discounts"`
	// OnlyActive is an optional filter to return only active tariffs with is_active=true.
	OnlyActive types.Bool `tfsdk:"only_active"`
}

type TariffVpsModel struct {
	Id       types.Int32   `tfsdk:"id"`
	Name     types.String  `tfsdk:"name"`
	CPU      types.Float32 `tfsdk:"cpu"`
	RAM      types.Float32 `tfsdk:"ram"`
	VRAM     types.Float32 `tfsdk:"vram"`
	IP       types.Float32 `tfsdk:"ip"`
	IsActive types.Bool    `tfsdk:"is_active"`
}

type TariffDriveModel struct {
	Id       types.Int32   `tfsdk:"id"`
	Name     types.String  `tfsdk:"name"`
	Price    types.Float32 `tfsdk:"price"`
	IsActive types.Bool    `tfsdk:"is_active"`
}

type TariffAdditionalDriveModel struct {
	Id       types.Int32   `tfsdk:"id"`
	Name     types.String  `tfsdk:"name"`
	Price    types.Float32 `tfsdk:"price"`
	IsActive types.Bool    `tfsdk:"is_active"`
}

type TariffAdditionalServiceModel struct {
	Id       types.Int32   `tfsdk:"id"`
	Name     types.String  `tfsdk:"name"`
	Price    types.Float32 `tfsdk:"price"`
	IsActive types.Bool    `tfsdk:"is_active"`
}

type TariffPaymentPeriodDiscountModel struct {
	PaymentPeriod types.Int32   `tfsdk:"payment_period"`
	Discount      types.Float32 `tfsdk:"discount"`
}

func (d *TariffsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_tariffs"
}

func (d *TariffsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "All tariffs data source",
		Attributes: map[string]schema.Attribute{
			"only_active": schema.BoolAttribute{
				MarkdownDescription: "If true, return only active tariffs with is_active=true.",
				Optional:            true,
			},
			"vps": schema.ListNestedAttribute{
				MarkdownDescription: "List of VPS tariffs.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id":        schema.Int32Attribute{Computed: true, MarkdownDescription: "Tariff ID."},
						"name":      schema.StringAttribute{Computed: true, MarkdownDescription: "Tariff name."},
						"cpu":       schema.Float32Attribute{Computed: true, MarkdownDescription: "Price per virtual CPU core."},
						"ram":       schema.Float32Attribute{Computed: true, MarkdownDescription: "Price per 1 GB RAM."},
						"vram":      schema.Float32Attribute{Computed: true, MarkdownDescription: "Price per 1 GB VRAM."},
						"ip":        schema.Float32Attribute{Computed: true, MarkdownDescription: "Price per 1 additional IP address."},
						"is_active": schema.BoolAttribute{Computed: true, MarkdownDescription: "Indicates if the tariff is active."},
					},
				},
			},
			"drive": schema.ListNestedAttribute{
				MarkdownDescription: "List of drive tariffs.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id":        schema.Int32Attribute{Computed: true, MarkdownDescription: "Tariff ID."},
						"name":      schema.StringAttribute{Computed: true, MarkdownDescription: "Tariff name."},
						"price":     schema.Float32Attribute{Computed: true, MarkdownDescription: "Price per 1 GB."},
						"is_active": schema.BoolAttribute{Computed: true, MarkdownDescription: "Indicates if the tariff is active."},
					},
				},
			},
			"additional_drive": schema.ListNestedAttribute{
				MarkdownDescription: "List of additional drive tariffs.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id":        schema.Int32Attribute{Computed: true, MarkdownDescription: "Tariff ID."},
						"name":      schema.StringAttribute{Computed: true, MarkdownDescription: "Tariff name."},
						"price":     schema.Float32Attribute{Computed: true, MarkdownDescription: "Price per 1 GB."},
						"is_active": schema.BoolAttribute{Computed: true, MarkdownDescription: "Indicates if the tariff is active."},
					},
				},
			},
			"additional_service": schema.ListNestedAttribute{
				MarkdownDescription: "List of additional service tariffs.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id":        schema.Int32Attribute{Computed: true, MarkdownDescription: "Tariff ID."},
						"name":      schema.StringAttribute{Computed: true, MarkdownDescription: "Tariff name."},
						"price":     schema.Float32Attribute{Computed: true, MarkdownDescription: "Price per service, depending on the service type."},
						"is_active": schema.BoolAttribute{Computed: true, MarkdownDescription: "Indicates if the tariff is active."},
					},
				},
			},
			"payment_period_discounts": schema.ListNestedAttribute{
				MarkdownDescription: "List of payment period discounts.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"payment_period": schema.Int32Attribute{Computed: true, MarkdownDescription: "Payment period enum. see API description for reference."},
						"discount":       schema.Float32Attribute{Computed: true, MarkdownDescription: "Discount amount (0.0-1.0)."},
					},
				},
			},
		},
	}
}

func (d *TariffsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *TariffsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data TariffsDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tariffs, err := d.client.GetTariffs()
	if err != nil {
		resp.Diagnostics.AddError(
			"Client Error",
			fmt.Sprintf("Unable to read tariffs, got error: %s", err),
		)
		return
	}

	onlyActive := !data.OnlyActive.IsNull() && data.OnlyActive.ValueBool()

	vps := make([]TariffVpsModel, 0, len(tariffs.VpsTariffs))
	for _, tariff := range tariffs.VpsTariffs {
		if onlyActive && !tariff.IsActive {
			continue
		}
		vps = append(vps, TariffVpsModel{
			Id:       types.Int32Value(tariff.Id),
			Name:     types.StringValue(tariff.Name),
			CPU:      types.Float32Value(tariff.CPU),
			RAM:      types.Float32Value(tariff.RAM),
			VRAM:     types.Float32Value(tariff.VRAM),
			IP:       types.Float32Value(tariff.IP),
			IsActive: types.BoolValue(tariff.IsActive),
		})
	}

	drive := make([]TariffDriveModel, 0, len(tariffs.DriveTariffs))
	for _, tariff := range tariffs.DriveTariffs {
		if onlyActive && !tariff.IsActive {
			continue
		}
		drive = append(drive, TariffDriveModel{
			Id:       types.Int32Value(tariff.Id),
			Name:     types.StringValue(tariff.Name),
			Price:    types.Float32Value(tariff.Price),
			IsActive: types.BoolValue(tariff.IsActive),
		})
	}

	additionalDrive := make([]TariffAdditionalDriveModel, 0, len(tariffs.AdditionalDriveTariffs))
	for _, tariff := range tariffs.AdditionalDriveTariffs {
		if onlyActive && !tariff.IsActive {
			continue
		}
		additionalDrive = append(additionalDrive, TariffAdditionalDriveModel{
			Id:       types.Int32Value(tariff.Id),
			Name:     types.StringValue(tariff.Name),
			Price:    types.Float32Value(tariff.Price),
			IsActive: types.BoolValue(tariff.IsActive),
		})
	}

	additionalService := make([]TariffAdditionalServiceModel, 0, len(tariffs.AdditionalServiceTariffs))
	for _, tariff := range tariffs.AdditionalServiceTariffs {
		if onlyActive && !tariff.IsActive {
			continue
		}
		additionalService = append(additionalService, TariffAdditionalServiceModel{
			Id:       types.Int32Value(tariff.Id),
			Name:     types.StringValue(tariff.Name),
			Price:    types.Float32Value(tariff.Price),
			IsActive: types.BoolValue(tariff.IsActive),
		})
	}

	paymentPeriodDiscounts := make([]TariffPaymentPeriodDiscountModel, 0, len(tariffs.PaymentPeriodDiscounts))
	for _, discount := range tariffs.PaymentPeriodDiscounts {
		paymentPeriodDiscounts = append(paymentPeriodDiscounts, TariffPaymentPeriodDiscountModel{
			PaymentPeriod: types.Int32Value(discount.PaymentPeriod),
			Discount:      types.Float32Value(discount.Discount),
		})
	}

	vpsType := types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id":        types.Int32Type,
			"name":      types.StringType,
			"cpu":       types.Float32Type,
			"ram":       types.Float32Type,
			"vram":      types.Float32Type,
			"ip":        types.Float32Type,
			"is_active": types.BoolType,
		},
	}

	driveType := types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id":        types.Int32Type,
			"name":      types.StringType,
			"price":     types.Float32Type,
			"is_active": types.BoolType,
		},
	}

	additionalDriveType := driveType
	additionalServiceType := driveType

	paymentPeriodDiscountType := types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"payment_period": types.Int32Type,
			"discount":       types.Float32Type,
		},
	}

	vpsRes, diags := types.ListValueFrom(ctx, vpsType, vps)
	resp.Diagnostics.Append(diags...)
	data.Vps = vpsRes

	driveRes, diags := types.ListValueFrom(ctx, driveType, drive)
	resp.Diagnostics.Append(diags...)
	data.Drive = driveRes

	additionalDriveRes, diags := types.ListValueFrom(ctx, additionalDriveType, additionalDrive)
	resp.Diagnostics.Append(diags...)
	data.AdditionalDrive = additionalDriveRes

	additionalServiceRes, diags := types.ListValueFrom(ctx, additionalServiceType, additionalService)
	resp.Diagnostics.Append(diags...)
	data.AdditionalService = additionalServiceRes

	paymentPeriodDiscountRes, diags := types.ListValueFrom(ctx, paymentPeriodDiscountType, paymentPeriodDiscounts)
	resp.Diagnostics.Append(diags...)
	data.PaymentPeriodDiscounts = paymentPeriodDiscountRes

	tflog.Trace(ctx, "read a data source")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
