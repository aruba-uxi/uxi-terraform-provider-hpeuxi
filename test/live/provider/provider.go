/*
Copyright 2024 Hewlett Packard Enterprise Development LP.
*/

package provider

import (
	"github.com/aruba-uxi/terraform-provider-hpeuxi/internal/provider"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

var (
	ProviderConfig                  = `provider "uxi" {}`
	TestAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
		"uxi": providerserver.NewProtocol6WithError(provider.New("test")()),
	}
)
