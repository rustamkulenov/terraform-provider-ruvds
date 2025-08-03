package provider

import (
	"context"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/ephemeral"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/rustamkulenov/terraform-provider-ruvds/internal/api"
	"github.com/rustamkulenov/terraform-provider-ruvds/internal/provider/datasources"
)

// Ensure RuVdsProvider satisfies various provider interfaces.
var _ provider.Provider = &RuVdsProvider{}
var _ provider.ProviderWithFunctions = &RuVdsProvider{}
var _ provider.ProviderWithEphemeralResources = &RuVdsProvider{}

// RuVdsProvider defines the provider implementation.
type RuVdsProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// RuVdsProviderModel describes the provider data model.
type RuVdsProviderModel struct {
	Endpoint types.String `tfsdk:"endpoint"`
	Token    types.String `tfsdk:"token"`
}

func (p *RuVdsProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "ruvds"
	resp.Version = p.version
}

func (p *RuVdsProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"endpoint": schema.StringAttribute{
				MarkdownDescription: "RuVds provider attribute",
				Optional:            true,
			},
			"token": schema.StringAttribute{
				MarkdownDescription: "RuVds provider API token. Can also be set via the `RUVDS_API_TOKEN` environment variable.",
				Optional:            true,
				Sensitive:           true,
			},
		},
	}
}

func (p *RuVdsProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var data RuVdsProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	token := data.Token.ValueString()

	// If token not provided then try to get it from environment variable
	if token == "" {
		token = os.Getenv("RUVDS_API_TOKEN")
		if token == "" {
			resp.Diagnostics.AddError(
				"Missing RuVDS API Token",
				"Please set the RUVDS_API_TOKEN environment variable or provide a token in the provider configuration.",
			)
			return
		}
	}

	// Client configuration for data sources and resources
	client := api.NewClient(
		token,
		data.Endpoint.ValueString(),
	)
	resp.DataSourceData = client
	resp.ResourceData = client
}

func (p *RuVdsProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewExampleResource,
	}
}

func (p *RuVdsProvider) EphemeralResources(ctx context.Context) []func() ephemeral.EphemeralResource {
	return []func() ephemeral.EphemeralResource{
		NewExampleEphemeralResource,
	}
}

func (p *RuVdsProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		datasources.NewDatacenterDataSource,
		datasources.NewDatacentersDataSource,
		datasources.NewOSListDataSource,
		datasources.NewOSDataSource,
		datasources.NewVpsListDataSource,
	}
}

func (p *RuVdsProvider) Functions(ctx context.Context) []func() function.Function {
	return []func() function.Function{
		NewExampleFunction,
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &RuVdsProvider{
			version: version,
		}
	}
}
