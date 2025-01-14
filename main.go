package main

import (
	"github.com/ImCCTech/terraform-provider-todostore/todostore"
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() terraform.ResourceProvider {
			return todostore.Provider()
		},
	})
}
