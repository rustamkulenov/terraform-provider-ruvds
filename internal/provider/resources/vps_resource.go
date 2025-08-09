// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resources

import (
	"context"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/rustamkulenov/terraform-provider-ruvds/internal/api"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &VpsResource{}
var _ resource.ResourceWithImportState = &VpsResource{}

func NewVpsResource() resource.Resource {
	return &VpsResource{}
}

// VpsResource defines the resource implementation.
type VpsResource struct {
	importing bool
	client    *api.Client
}

// VpsResourceModel describes the resource data model.
type VpsResourceModel struct {
	ID                      types.Int32   `tfsdk:"id"`
	Status                  types.String  `tfsdk:"status"`
	CreateProgress          types.Int32   `tfsdk:"create_progress"`
	DataCenterID            types.Int32   `tfsdk:"datacenter_id"`
	TariffID                types.Int32   `tfsdk:"tariff_id"`
	PaymentPeriod           types.Int32   `tfsdk:"payment_period"`
	OSID                    types.Int32   `tfsdk:"os_id"`
	CPU                     types.Int32   `tfsdk:"cpu"`
	RAM                     types.Float32 `tfsdk:"ram"`
	VRAM                    types.Int32   `tfsdk:"vram"`
	Drive                   types.Int32   `tfsdk:"drive"`
	DriveTariffID           types.Int32   `tfsdk:"drive_tariff_id"`
	IP                      types.Int32   `tfsdk:"ip"`
	DDOSProtection          types.Float32 `tfsdk:"ddos_protection"`
	PaidTill                types.String  `tfsdk:"paid_till"`
	TemplateID              types.String  `tfsdk:"template_id"`
	AdditionalDrive         types.Int32   `tfsdk:"additional_drive"`
	AdditionalDriveTariffID types.Int32   `tfsdk:"additional_drive_tariff_id"`
	UserComment             types.String  `tfsdk:"user_comment"`
	SShKeyId                types.String  `tfsdk:"ssh_key_id"`
	ComputerName            types.String  `tfsdk:"computer_name"`
	IPAddress               types.String  `tfsdk:"ip_address"`
	Netmask                 types.String  `tfsdk:"netmask"`
	Gateway                 types.String  `tfsdk:"gateway"`
}

func (r *VpsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vps"
}

func (r *VpsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "VPS resource for managing virtual servers",

		Attributes: map[string]schema.Attribute{
			"id": schema.Int32Attribute{
				MarkdownDescription: "The ID of the example resource.",
				Computed:            true,
			},
			"status": schema.StringAttribute{
				MarkdownDescription: "The status of the virtual server.",
				Computed:            true,
			},
			"create_progress": schema.Int32Attribute{
				MarkdownDescription: "The progress of the virtual server creation.",
				Computed:            true,
			},
			"datacenter_id": schema.Int32Attribute{
				MarkdownDescription: "The ID of the data center where the virtual server is located.",
				Required:            true,
			},
			"tariff_id": schema.Int32Attribute{
				MarkdownDescription: "The ID of the tariff plan for the virtual server.",
				Required:            true,
			},
			"payment_period": schema.Int32Attribute{
				MarkdownDescription: "The payment period for the virtual server.",
				Required:            true,
			},
			"os_id": schema.Int32Attribute{
				MarkdownDescription: "The ID of the operating system installed on the virtual server.",
				Required:            true,
			},
			"cpu": schema.Int32Attribute{
				MarkdownDescription: "The number of CPU cores allocated to the virtual server.",
				Required:            true,
			},
			"ram": schema.Float32Attribute{
				MarkdownDescription: "The amount of RAM allocated to the virtual server in GB.",
				Required:            true,
			},
			"vram": schema.Int32Attribute{
				MarkdownDescription: "The amount of VRAM allocated to the virtual server in MB.",
				Optional:            true,
			},
			"drive": schema.Int32Attribute{
				MarkdownDescription: "The size of the primary drive allocated to the virtual server in GB.",
				Required:            true,
			},
			"drive_tariff_id": schema.Int32Attribute{
				MarkdownDescription: "The ID of the tariff plan for the primary drive.",
				Required:            true,
			},
			"ip": schema.Int32Attribute{
				MarkdownDescription: "The number of IP address allocated to the virtual server.",
				Required:            true,
			},
			"ddos_protection": schema.Float32Attribute{
				MarkdownDescription: "The level of DDoS protection applied to the virtual server.",
				Optional:            true,
			},
			"paid_till": schema.StringAttribute{
				MarkdownDescription: "The date until which the virtual server is paid.",
				Computed:            true,
			},
			"template_id": schema.StringAttribute{
				MarkdownDescription: "The ID of the template used for the virtual server.",
				Optional:            true,
			},
			"additional_drive": schema.Int32Attribute{
				MarkdownDescription: "The size of the additional drive allocated to the virtual server in GB.",
				Optional:            true,
			},
			"additional_drive_tariff_id": schema.Int32Attribute{
				MarkdownDescription: "The ID of the tariff plan for the additional drive.",
				Optional:            true,
			},
			"user_comment": schema.StringAttribute{
				MarkdownDescription: "A user comment associated with the virtual server.",
				Optional:            true,
			},
			"ssh_key_id": schema.StringAttribute{
				MarkdownDescription: "The ID of the SSH key associated with the virtual server.",
				Optional:            true,
			},
			"computer_name": schema.StringAttribute{
				MarkdownDescription: "The name of the computer for the virtual server.",
				Optional:            true,
			},
			"ip_address": schema.StringAttribute{Computed: true},
			"netmask":    schema.StringAttribute{Computed: true},
			"gateway":    schema.StringAttribute{Computed: true},
		},
	}
}

func (r *VpsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Parse the import ID which contains Id of VDI to import
	importID := req.ID

	intValue, err := strconv.Atoi(importID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid import ID",
			fmt.Sprintf("Could not parse import ID as int32: %v", err),
		)
		return
	}

	// Convert to the appropriate type for state
	state := VpsResourceModel{
		ID: types.Int32Value(int32(intValue)),
	}

	// Set the state
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *VpsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*api.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *api.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *VpsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan VpsResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)

	if resp.Diagnostics.HasError() {
		return
	}

	if r.importing {
		// Set defaults for mandatory fields during import
		if plan.Drive.IsNull() {
			plan.Drive = types.Int32Value(1)
		}
	}

	// Call API to create the virtual server
	vps := api.CreateVpsRequest(
		plan.DataCenterID.ValueInt32(),
		plan.TariffID.ValueInt32(),
		plan.PaymentPeriod.ValueInt32(),
		plan.OSID.ValueInt32(),
		plan.CPU.ValueInt32(),
		plan.RAM.ValueFloat32(),
		plan.Drive.ValueInt32(),
		plan.DriveTariffID.ValueInt32(),
		plan.IP.ValueInt32(),
	)

	response, descr, err := r.client.CreateVps(&vps)
	if err != nil {
		resp.Diagnostics.AddError(
			"Client Error",
			fmt.Sprintf("Unable to create virtual server, got error: %s: %v", err, *descr),
		)
		return
	}

	// Save received values in state
	plan.ID = types.Int32Value(response.VirtualServerId)
	plan.Status = types.StringValue(*response.Status.Status)
	plan.CreateProgress = types.Int32Value(response.Status.CreateProgress)
	plan.PaidTill = types.StringValue(*response.Status.PaidTill)
	if len(response.Status.NetworkV4) > 0 {
		plan.IPAddress = types.StringValue(response.Status.NetworkV4[0].IPAddress)
		plan.Netmask = types.StringValue(response.Status.NetworkV4[0].Netmask)
		plan.Gateway = types.StringValue(response.Status.NetworkV4[0].Gateway)
	}

	// Write logs using the tflog package
	// Documentation: https://terraform.io/plugin/log
	tflog.Trace(ctx, "created a resource")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *VpsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data VpsResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Get virtual server by ID and update the model
	if data.ID.ValueInt32() == 0 {
		resp.Diagnostics.AddError(
			"Invalid ID",
			"Virtual server ID is not set. Cannot read.",
		)
		return
	}
	vps, err := r.client.GetVps(data.ID.ValueInt32())
	if err != nil {
		resp.Diagnostics.AddError(
			"Client Error",
			fmt.Sprintf("Unable to read virtual server with ID %d, got error: %s", data.ID.ValueInt32(), err),
		)
		return
	}
	data.CreateProgress = types.Int32Value(vps.CreateProgress)
	data.DataCenterID = types.Int32Value(vps.DataCenterId)
	data.TariffID = types.Int32Value(vps.TariffId)
	data.PaymentPeriod = types.Int32Value(vps.PaymentPeriod)
	data.OSID = types.Int32Value(vps.OSId)
	data.CPU = types.Int32Value(vps.CPU)
	data.RAM = types.Float32Value(vps.RAM)
	data.Drive = types.Int32Value(vps.Drive)
	data.DriveTariffID = types.Int32Value(vps.DriveTariffId)
	data.IP = types.Int32Value(vps.IP)
	data.DDOSProtection = types.Float32Value(vps.DDOSProtection)
	if len(vps.NetworkV4) > 0 {
		data.IPAddress = types.StringValue(vps.NetworkV4[0].IPAddress)
		data.Gateway = types.StringValue(vps.NetworkV4[0].Gateway)
		data.Netmask = types.StringValue(vps.NetworkV4[0].Netmask)
	}

	if vps.VRAM != nil {
		data.VRAM = types.Int32Value(*vps.VRAM)
	} else {
		data.VRAM = types.Int32Null()
	}
	if vps.PaidTill != nil {
		data.PaidTill = types.StringValue(*vps.PaidTill)
	} else {
		data.PaidTill = types.StringNull()
	}
	if vps.Status != nil {
		data.Status = types.StringValue(*vps.Status)
	} else {
		data.Status = types.StringNull()
	}
	if vps.TemplateId != nil {
		data.TemplateID = types.StringValue(*vps.TemplateId)
	} else {
		data.TemplateID = types.StringNull()
	}
	if vps.AdditionalDrive != nil {
		data.AdditionalDrive = types.Int32Value(*vps.AdditionalDrive)
	} else {
		data.AdditionalDrive = types.Int32Null()
	}
	if vps.AdditionalDriveTariffId != nil {
		data.AdditionalDriveTariffID = types.Int32Value(*vps.AdditionalDriveTariffId)
	} else {
		data.AdditionalDriveTariffID = types.Int32Null()
	}
	if vps.UserComment != nil {
		data.UserComment = types.StringValue(*vps.UserComment)
	} else {
		data.UserComment = types.StringNull()
	}
	if vps.SShKeyID != nil {
		data.SShKeyId = types.StringValue(*vps.SShKeyID)
	} else {
		data.SShKeyId = types.StringNull()
	}
	if vps.ComputerName != nil {
		data.ComputerName = types.StringValue(*vps.ComputerName)
	} else {
		data.ComputerName = types.StringNull()
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data VpsResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// If applicable, this is a great opportunity to initialize any necessary
	// provider client data and make a call using it.
	// httpResp, err := r.client.Do(httpReq)
	// if err != nil {
	//     resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update example, got error: %s", err))
	//     return
	// }

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data VpsResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	serverid := data.ID.ValueInt32()
	if serverid == 0 {
		resp.Diagnostics.AddError(
			"Invalid ID",
			"Virtual server ID is not set. Cannot delete.",
		)
		return
	}
	// Call API to delete the virtual server
	actionResult, descr, err := r.client.DeleteVps(serverid)
	if err != nil {
		resp.Diagnostics.AddError(
			"Client Error",
			fmt.Sprintf("Unable to delete virtual server with ID %d, got error: %s: %v", serverid, err, *descr),
		)
		return
	}
	// Log the action result
	tflog.Trace(ctx, fmt.Sprintf("Delete action result: ID=%d, Type=%s, Status=%s, Progress=%d, Started=%s, Finished=%s, ResourceID=%d, ResourceType=%s",
		actionResult.ID, actionResult.Type, actionResult.Status, actionResult.Progress,
		actionResult.Started, actionResult.Finished, actionResult.ResourceId, actionResult.ResourceType))
}
