package main

import (
	"context"
	"flag"
	"log"

	"dptech/edu/hashicups-pf/provider"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
)

var (
	version string
)

func main() {
	var debug bool

	flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	opts := providerserver.ServeOpts{
		Address: "dptech/edu/hashicups-pf",
		Debug:   debug,
	}

	err := providerserver.Serve(context.Background(), provider.New(version), opts)

	if err != nil {
		log.Fatal(err.Error())
	}
}
