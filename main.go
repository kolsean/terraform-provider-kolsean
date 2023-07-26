package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/kolsean/terraform-provider-kolsean/kolsean"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: kolsean.Provider,
	})
}
