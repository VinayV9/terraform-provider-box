package main

import (
	"./client"
	"github.com/hashicorp/terraform/helper/schema"
)

// Provider for box
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: providerSchema(),
		ResourcesMap: map[string]*schema.Resource{
			"box_folder": resourceBoxFolder(),
		},
		ConfigureFunc: providerConfigure,
	}
}

// BoxClient schemma
type BoxClient struct {
	AccessToken string
}

func providerSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"access_token": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "access_token used to authenticate with the service provider",
		},
		"base_url": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "base url used to authenticate with the service provider",
		},
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	accessToken := d.Get("access_token").(string)
	baseURL := d.Get("base_url").(string)
	return client.NewClient(accessToken, baseURL), nil
}

// client_id = ezxiv7hc7ninreix0m4b1vantpzdlflh
// primary_access_token = NvocCaINeboZfFLGXzzJGC5A3IURu8S5
