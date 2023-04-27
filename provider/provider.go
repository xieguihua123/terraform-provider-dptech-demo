package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure ScaffoldingProvider satisfies various provider interfaces.
var _ provider.Provider = &ScaffoldingProvider{}

// ScaffoldingProvider defines the provider implementation.
type ScaffoldingProvider struct {
	version string
}

// ScaffoldingProviderModel describes the provider data model.
type ScaffoldingProviderModel struct {
	Name    types.String `tfsdk:"name"`
	Address types.String `tfsdk:"address"`
}

func (p *ScaffoldingProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "dptech-demo"
	fmt.Println("111111111111111111111111111111111")
}

func (p *ScaffoldingProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	fmt.Println("111111111111111111111111111111111")
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				MarkdownDescription: "  provider attribute",
				Optional:            true,
			}, "address": schema.StringAttribute{
				MarkdownDescription: "  provider attribute",
				Optional:            true,
			},
		},
	}
}

func (p *ScaffoldingProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var data ScaffoldingProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	fmt.Println("111111111111111111111111111111111")
	if resp.Diagnostics.HasError() {
		return
	}
	address := ""
	if !data.Address.IsNull() {
		address = data.Address.ValueString()
	}

	if data.Name.IsNull() {
		tflog.Info(ctx, "Name is NULL")
		return
	}
	if data.Address.IsNull() {
		tflog.Info(ctx, "address is NULL")
		return
	}
	tflog.Info(ctx, address)
	client, err := NewClient(&address)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create HashiCups API Client",
			"An unexpected error occurred when creating the HashiCups API client. "+
				"If the error is not clear, please contact the provider developers.\n\n"+
				"HashiCups Client Error: "+err.Error(),
		)
		return
	}

	resp.DataSourceData = client
	resp.ResourceData = client
}

func (p *ScaffoldingProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewExampleResource,
	}
}

func (p *ScaffoldingProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewExampleDataSource,
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &ScaffoldingProvider{
			version: version,
		}
	}
}
