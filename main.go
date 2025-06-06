/*
Copyright 2024 Hewlett Packard Enterprise Development LP.
*/

package main

import (
	"context"
	"flag"
	"log"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"

	"github.com/aruba-uxi/terraform-provider-hpeuxi/internal/provider"
)

var version string = "dev"

func main() {
	var debug bool

	flag.BoolVar(
		&debug,
		"debug",
		false,
		"set to true to run the provider with support for debuggers like delve",
	)
	flag.Parse()

	opts := providerserver.ServeOpts{
		Address: "registry.terraform.io/HewlettPackard/hpeuxi",
		Debug:   debug,
	}

	if err := providerserver.Serve(context.Background(), provider.New(version), opts); err != nil {
		log.Fatal(err.Error())
	}
}
