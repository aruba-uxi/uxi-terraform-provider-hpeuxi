package util

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/aruba-uxi/terraform-provider-hpeuxi/pkg/config-api-client"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/live/config"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/stretchr/testify/assert"
)

func GetAgentProperties(id string) config_api_client.AgentItem {
	result, _, err := Client.ConfigurationAPI.
		AgentsGet(context.Background()).
		Id(id).
		Execute()
	if err != nil {
		panic(err)
	}
	if len(result.Items) != 1 {
		panic("agent with id `" + id + "` could not be found")
	}
	return result.Items[0]
}

func CheckStateAgainstAgent(
	t *testing.T,
	agent config_api_client.AgentItem,
) resource.TestCheckFunc {
	return resource.ComposeAggregateTestCheckFunc(
		resource.TestCheckResourceAttr("data.uxi_agent.my_agent", "id", config.AgentPermanentId),
		resource.TestCheckResourceAttr("data.uxi_agent.my_agent", "serial", agent.Serial),
		TestOptionalValue(t, "data.uxi_agent.my_agent", "model_number", agent.ModelNumber.Get()),
		resource.TestCheckResourceAttrWith(
			"data.uxi_agent.my_agent",
			"name",
			func(value string) error {
				assert.Equal(t, value, agent.Name)
				return nil
			},
		),
		TestOptionalValue(
			t,
			"data.uxi_agent.my_agent",
			"wifi_mac_address",
			agent.WifiMacAddress.Get(),
		),
		TestOptionalValue(
			t,
			"data.uxi_agent.my_agent",
			"ethernet_mac_address",
			agent.EthernetMacAddress.Get(),
		),
		TestOptionalValue(t, "data.uxi_agent.my_agent", "notes", agent.Notes.Get()),
		TestOptionalValue(t, "data.uxi_agent.my_agent", "pcap_mode", agent.PcapMode.Get()),
	)
}

type ProvisionAgent struct {
	CustomerId        string
	ProvisionToken    string
	DeviceSerial      string
	DeviceGatewayHost string
}

type provisionAgentRequest struct {
	Uid            string `json:"uid"`
	CustomerUid    string `json:"customer_uid"`
	ProvisionToken string `json:"provision_token"`
	PlatformName   string `json:"platform_name"`
	DeviceSerial   string `json:"device_serial"`
}

func (p ProvisionAgent) Provision() (string, error) {
	url := p.DeviceGatewayHost + "/provision-zebra-device"
	id, err := p.generateId()
	if err != nil {
		return id, err
	}

	request := provisionAgentRequest{
		Uid:            id,
		CustomerUid:    p.CustomerId,
		ProvisionToken: p.ProvisionToken,
		PlatformName:   "zebra",
		DeviceSerial:   p.DeviceSerial,
	}
	jsonData, err := json.Marshal(request)
	if err != nil {
		return id, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return id, err
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return id, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusCreated {
		return id, fmt.Errorf(
			"unexpected status code returned: %d\nresponse: %s",
			resp.StatusCode,
			string(body),
		)
	}

	return id, nil
}

func (p ProvisionAgent) generateId() (string, error) {
	// Create an MD5 hash of the serial string
	hasher := md5.New()
	hasher.Write([]byte(p.DeviceSerial))
	digest := hasher.Sum(nil)

	// Use the first 16 bytes of the digest to create a UUID v3
	uuid, err := uuid.FromBytes(digest[:16])
	if err != nil {
		return "", err
	}
	uuid[6] = (uuid[6] & 0x0f) | 0x30 // Set the version to 3 (MD5-based UUID)
	uuid[8] = (uuid[8] & 0x3f) | 0x80 // Set the variant to RFC 4122

	return uuid.String(), nil
}

func GetAgent(id string) *config_api_client.AgentItem {
	result, _, err := Client.ConfigurationAPI.
		AgentsGet(context.Background()).
		Id(id).
		Execute()
	if err != nil {
		panic(err)
	}
	if len(result.Items) != 1 {
		return nil
	}
	return &result.Items[0]
}
