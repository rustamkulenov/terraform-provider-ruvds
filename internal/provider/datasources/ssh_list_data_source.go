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

// SshListDataSource is a data source for retrieving information about SSH keys.
type SshListDataSource struct {
	client *api.Client
}

func NewSshListDataSource() datasource.DataSource {
	return &SshListDataSource{}
}

var _ datasource.DataSource = &SshListDataSource{}

// SshListDataSourceModel describes the data source data model.
type SshListDataSourceModel struct {
	// List of ssh keys.
	Servers types.List `tfsdk:"ssh_keys"`
}

type SshModel struct {
	SshKeyId          types.String `tfsdk:"ssh_key_id"`
	Name              types.String `tfsdk:"name"`
	PublicKey         types.String `tfsdk:"public_key"`
	Md5Fingerprint    types.String `tfsdk:"md5_fingerprint"`
	Sha256Fingerprint types.String `tfsdk:"sha256_fingerprint"`
}

func (d *SshListDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ssh_list"
}

func (d *SshListDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "SSH list data source",

		Attributes: map[string]schema.Attribute{
			"ssh_keys": schema.ListNestedAttribute{
				MarkdownDescription: "List of deployed SSH keys",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"ssh_key_id":         schema.StringAttribute{Computed: true, Description: "SSH key ID"},
						"name":               schema.StringAttribute{Computed: true, Description: "SSH key name"},
						"public_key":         schema.StringAttribute{Computed: true, Description: "SSH public key"},
						"md5_fingerprint":    schema.StringAttribute{Computed: true, Description: "MD5 fingerprint of the SSH key"},
						"sha256_fingerprint": schema.StringAttribute{Computed: true, Description: "SHA256 fingerprint of the SSH key"},
					},
				},
			},
		},
	}
}

func (d *SshListDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *SshListDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data SshListDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	sshKeysDto, err := d.client.GetSshKeyList()
	if err != nil {
		resp.Diagnostics.AddError(
			"Client Error",
			fmt.Sprintf("Unable to read SSH list, got error: %s", err),
		)
		return
	}

	sshKeys := make([]SshModel, 0, len(sshKeysDto.SshKeys))
	for _, ssh := range sshKeysDto.SshKeys {
		vps := SshModel{
			SshKeyId:          types.StringValue(ssh.SshKeyId),
			Name:              types.StringValue(ssh.Name),
			PublicKey:         types.StringValue(ssh.PublicKey),
			Md5Fingerprint:    types.StringValue(ssh.Md5Fingerprint),
			Sha256Fingerprint: types.StringValue(ssh.Sha256Fingerprint),
		}
		sshKeys = append(sshKeys, vps)
	}

	sshKeysType := types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ssh_key_id":         types.StringType,
			"name":               types.StringType,
			"public_key":         types.StringType,
			"md5_fingerprint":    types.StringType,
			"sha256_fingerprint": types.StringType,
		},
	}
	sshres, diags := types.ListValueFrom(ctx, sshKeysType, sshKeys)
	resp.Diagnostics.Append(diags...)
	data.Servers = sshres

	// Write logs using the tflog package
	// Documentation: https://terraform.io/plugin/log
	tflog.Trace(ctx, "read a data source")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
