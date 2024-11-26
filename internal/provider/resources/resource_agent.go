/*
Copyright 2024 Hewlett Packard Enterprise Development LP.
*/

package resources

import (
	"context"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/aruba-uxi/terraform-provider-hpeuxi/internal/provider/util"
	config_api_client "github.com/aruba-uxi/terraform-provider-hpeuxi/pkg/config-api-client"
)

var (
	_ resource.Resource              = &agentResource{}
	_ resource.ResourceWithConfigure = &agentResource{}
)

type agentResourceModel struct {
	ID                 types.String `tfsdk:"id"`
	Serial             types.String `tfsdk:"serial"`
	Name               types.String `tfsdk:"name"`
	ModelNumber        types.String `tfsdk:"model_number"`
	WifiMacAddress     types.String `tfsdk:"wifi_mac_address"`
	EthernetMacAddress types.String `tfsdk:"ethernet_mac_address"`
	Notes              types.String `tfsdk:"notes"`
	PcapMode           types.String `tfsdk:"pcap_mode"`
}

func NewAgentResource() resource.Resource {
	return &agentResource{}
}

type agentResource struct {
	client *config_api_client.APIClient
}

func (r *agentResource) Metadata(
	ctx context.Context,
	req resource.MetadataRequest,
	resp *resource.MetadataResponse,
) {
	resp.TypeName = req.ProviderTypeName + "_agent"
}

func (r *agentResource) Schema(
	_ context.Context,
	_ resource.SchemaRequest,
	resp *resource.SchemaResponse,
) {
	resp.Schema = schema.Schema{
		Description: "Manages an agent.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "The identifier of the agent.",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"serial": schema.StringAttribute{
				Description: "The serial number of the agent.",
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: "The name of the agent.",
				Required:    true,
			},
			"model_number": schema.StringAttribute{
				Description: "The model number of the agent.",
				Computed:    true,
			},
			"wifi_mac_address": schema.StringAttribute{
				Description: "The wifi mac address of the agent.",
				Computed:    true,
			},
			"ethernet_mac_address": schema.StringAttribute{
				Description: "The ethernet mac address of the agent.",
				Computed:    true,
			},
			"notes": schema.StringAttribute{
				Description: "The notes of the agent.",
				Optional:    true,
				Computed:    true,
			},
			"pcap_mode": schema.StringAttribute{
				Description: "The packet capture mode of the agent.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func (r *agentResource) Configure(
	_ context.Context,
	req resource.ConfigureRequest,
	resp *resource.ConfigureResponse,
) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*config_api_client.APIClient)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			"Resource type: Group. Please report this issue to the provider developers.",
		)

		return
	}

	r.client = client
}

func (r *agentResource) Create(
	ctx context.Context,
	req resource.CreateRequest,
	resp *resource.CreateResponse,
) {
	var plan agentResourceModel
	diags := req.Plan.Get(ctx, &plan)
	diags.AddError(
		"operation not supported",
		"creating an agent is not supported; agents can only be imported",
	)
	resp.Diagnostics.Append(diags...)
}

func (r *agentResource) Read(
	ctx context.Context,
	req resource.ReadRequest,
	resp *resource.ReadResponse,
) {
	var state agentResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	request := r.client.ConfigurationAPI.
		AgentsGet(ctx).
		Id(state.ID.ValueString())
	agentResponse, response, err := util.RetryForTooManyRequests(request.Execute)
	// defer response.Body.Close()
	errorPresent, errorDetail := util.RaiseForStatus(response, err)

	errorSummary := util.GenerateErrorSummary("read", "uxi_agent")

	if errorPresent {
		resp.Diagnostics.AddError(errorSummary, errorDetail)

		return
	}

	if len(agentResponse.Items) != 1 {
		resp.State.RemoveResource(ctx)

		return
	}
	agent := agentResponse.Items[0]

	state.ID = types.StringValue(agent.Id)
	state.Name = types.StringValue(agent.Name)
	state.Serial = types.StringValue(agent.Serial)
	state.ModelNumber = types.StringPointerValue(agent.ModelNumber.Get())
	state.WifiMacAddress = types.StringPointerValue(agent.WifiMacAddress.Get())
	state.EthernetMacAddress = types.StringPointerValue(agent.EthernetMacAddress.Get())
	state.Notes = types.StringPointerValue(agent.Notes.Get())
	state.PcapMode = types.StringPointerValue(agent.PcapMode.Get())

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *agentResource) Update(
	ctx context.Context,
	req resource.UpdateRequest,
	resp *resource.UpdateResponse,
) {
	var plan agentResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	errorSummary := util.GenerateErrorSummary("update", "uxi_agent")
	patchRequest := config_api_client.NewAgentsPatchRequest()
	patchRequest.Name = plan.Name.ValueStringPointer()
	if !plan.Notes.IsUnknown() {
		patchRequest.Notes = plan.Notes.ValueStringPointer()
	}
	plannedPcapMode := plan.PcapMode.ValueStringPointer()
	if !plan.PcapMode.IsUnknown() && plannedPcapMode != nil {
		pcapMode, err := config_api_client.NewPcapModeFromValue(*plannedPcapMode)
		if err != nil {
			resp.Diagnostics.AddError(errorSummary, err.Error())

			return
		}
		patchRequest.PcapMode = pcapMode
	}
	request := r.client.ConfigurationAPI.
		AgentsPatch(ctx, plan.ID.ValueString()).
		AgentsPatchRequest(*patchRequest)
	agent, response, err := util.RetryForTooManyRequests(request.Execute)
	// defer response.Body.Close()

	errorPresent, errorDetail := util.RaiseForStatus(response, err)

	if errorPresent {
		resp.Diagnostics.AddError(errorSummary, errorDetail)

		return
	}

	plan.ID = types.StringValue(agent.Id)
	plan.Name = types.StringValue(agent.Name)
	plan.Serial = types.StringValue(agent.Serial)
	plan.ModelNumber = types.StringPointerValue(agent.ModelNumber.Get())
	plan.WifiMacAddress = types.StringPointerValue(agent.WifiMacAddress.Get())
	plan.EthernetMacAddress = types.StringPointerValue(agent.EthernetMacAddress.Get())
	plan.Notes = types.StringPointerValue(agent.Notes.Get())
	if agent.PcapMode.Get() != nil {
		plan.PcapMode = types.StringValue(string(*agent.PcapMode.Get()))
	}

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *agentResource) Delete(
	ctx context.Context,
	req resource.DeleteRequest,
	resp *resource.DeleteResponse,
) {
	var state agentResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	request := r.client.ConfigurationAPI.AgentsDelete(ctx, state.ID.ValueString())

	_, response, err := util.RetryForTooManyRequests(request.Execute)
	// defer response.Body.Close()
	errorPresent, errorDetail := util.RaiseForStatus(response, err)

	if errorPresent {
		if response != nil && response.StatusCode == http.StatusNotFound {
			resp.State.RemoveResource(ctx)

			return
		}
		resp.Diagnostics.AddError(util.GenerateErrorSummary("delete", "uxi_agent"), errorDetail)

		return
	}
}

func (r *agentResource) ImportState(
	ctx context.Context,
	req resource.ImportStateRequest,
	resp *resource.ImportStateResponse,
) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
