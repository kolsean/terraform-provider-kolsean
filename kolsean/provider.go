package kolsean

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"kolsean_object_v1": {
				Schema: map[string]*schema.Schema{
					"name": {
						Type:     schema.TypeString,
						Required: true,
						ForceNew: true,
					},
				},
				CreateContext: func(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
					name := d.Get("name").(string)
					d.Set("name", name)
					d.SetId(name)
					return nil
				},
				ReadContext: func(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
					name := d.Get("name")
					d.Set("name", name)
					return nil
				},
				DeleteContext: func(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
					d.SetId("")
					return nil
				},
			},
		},
	}
}
