package main

import (
	"./client"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceBoxFolder() *schema.Resource {
	return &schema.Resource{
		Create: resourceBoxFolderCreate,
		Read:   resourceBoxFolderRead,
		Update: resourceBoxFolderUpdate,
		Delete: resourceBoxFolderDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceBoxFolderCreate(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	name := d.Get("name").(string)
	id := d.Get("id").(string)

	err := apiClient.CreateFolder(name, id)
	if err != nil {
		return err
	}

	d.SetId(name)

	return nil
}

func resourceBoxFolderRead(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	id := d.Get("id").(string)
	
	err := apiClient.
	return nil
}

func resourceBoxFolderUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceBoxFolderRead(d, m)
}

func resourceBoxFolderDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
