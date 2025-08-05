// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resources

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/rustamkulenov/terraform-provider-ruvds/internal/api"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &SshResource{}
var _ resource.ResourceWithImportState = &SshResource{}

func NewSshResource() resource.Resource {
	return &SshResource{}
}

// SshResource defines the resource implementation.
type SshResource struct {
	client *api.Client
}

// SshResourceModel describes the resource data model.
type SshResourceModel struct {
	SshKeyId          types.String `tfsdk:"ssh_key_id"`
	Name              types.String `tfsdk:"name"`
	PublicKey         types.String `tfsdk:"public_key"`
	Md5Fingerprint    types.String `tfsdk:"md5_fingerprint"`
	Sha256Fingerprint types.String `tfsdk:"sha256_fingerprint"`
}

func (r *SshResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ssh"
}

func (r *SshResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "SSH resource for managing SSH keys in RuVDS.",

		Attributes: map[string]schema.Attribute{
			"ssh_key_id": schema.StringAttribute{
				MarkdownDescription: "SSH key ID",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Name of the SSH key",
				Required:            true,
			},
			"public_key": schema.StringAttribute{
				MarkdownDescription: "Public key content. Optional because can be imported",
				Optional:            true,
			},
			"md5_fingerprint": schema.StringAttribute{
				MarkdownDescription: "MD5 fingerprint of the SSH key",
				Computed:            true,
			},
			"sha256_fingerprint": schema.StringAttribute{
				MarkdownDescription: "SHA256 fingerprint of the SSH key",
				Computed:            true,
			},
		},
	}
}

func (r *SshResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Parse the import ID (simple passthrough in this case)
	resource.ImportStatePassthroughID(ctx, path.Root("name"), req, resp)
}

func (r *SshResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *SshResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan SshResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Call API to create the SSH key
	key, descr, err := r.client.CreateSshKey(api.CreateSSHKeyRequest{
		Name:      plan.Name.ValueString(),
		PublicKey: plan.PublicKey.ValueString(),
	})

	if err != nil {
		resp.Diagnostics.AddError(
			"Client Error",
			fmt.Sprintf("Unable to create SSH key, got error: %s: %v", err, *descr),
		)
		return
	}

	// Save received values in state
	plan.SshKeyId = types.StringValue(key.SshKeyId)
	plan.Md5Fingerprint = types.StringValue(key.Md5Fingerprint)
	plan.Sha256Fingerprint = types.StringValue(key.Sha256Fingerprint)

	// Write logs using the tflog package
	// Documentation: https://terraform.io/plugin/log
	tflog.Trace(ctx, "created a resource")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *SshResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SshResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// If name is set then get all Ssh keys and find the key
	if data.Name.ValueString() != "" {
		keys, err := r.client.GetSshKeyList()
		if err != nil {
			resp.Diagnostics.AddError(
				"Client Error",
				fmt.Sprintf("Unable to read SSH key by name %s, got error: %s", data.Name.ValueString(), err),
			)
			return
		}
		for _, ssh := range keys.SshKeys {
			if ssh.Name == data.Name.ValueString() {
				data.Name = types.StringValue(ssh.Name)
				data.PublicKey = types.StringValue(ssh.PublicKey)
				data.Md5Fingerprint = types.StringValue(ssh.Md5Fingerprint)
				data.Sha256Fingerprint = types.StringValue(ssh.Sha256Fingerprint)
				data.SshKeyId = types.StringValue(ssh.SshKeyId)

				// Save updated data into Terraform state
				resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
				return
			}
		}
	}

	// Get SSH key by ID and update the model
	if data.SshKeyId.ValueString() == "" {
		resp.Diagnostics.AddError(
			"Invalid SSH Key ID",
			"SSH Key ID is not set. Cannot read.",
		)
		return
	}
	sshKey, err := r.client.GetSshKey(data.SshKeyId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Client Error",
			fmt.Sprintf("Unable to read SSH key with ID %s, got error: %s", data.SshKeyId.ValueString(), err),
		)
		return
	}
	data.Name = types.StringValue(sshKey.Name)
	data.PublicKey = types.StringValue(sshKey.PublicKey)
	data.Md5Fingerprint = types.StringValue(sshKey.Md5Fingerprint)
	data.Sha256Fingerprint = types.StringValue(sshKey.Sha256Fingerprint)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SshResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data SshResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SshResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SshResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	keyid := data.SshKeyId.ValueString()
	if keyid == "" {
		resp.Diagnostics.AddError(
			"Invalid ID",
			"SSH key ID is not set. Cannot delete.",
		)
		return
	}
	// Call API to delete the ssh key (it does not return HTTP 200)
	_, _ = r.client.DeleteSshKey(keyid)
}
