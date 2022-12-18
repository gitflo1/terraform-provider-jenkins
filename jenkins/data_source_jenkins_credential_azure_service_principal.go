package jenkins

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceJenkinsCredentialAzureServicePrincipal() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceJenkinsCredentialAzureServicePrincipalRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Description: "The credential id of the Azure serivce principal credential.",
				Required:    true,
			},
			"folder": {
				Type:        schema.TypeString,
				Description: "The Jenkins folder that contains the Service principal credentials.",
				Required:    true,
			},
			"scope": {
				Type:             schema.TypeString,
				Description:      "The Jenkins scope assigned to the Service Principal credentials.",
				Optional:         true,
				Default:          "GLOBAL",
				ValidateDiagFunc: validateCredentialScope,
			},
			"domain": {
				Type:        schema.TypeString,
				Description: "The Jenkins domain that the credentials will be added to.",
				Optional:    true,
				Default:     "_",
			},
			"subscription_id": {
				Type:        schema.TypeString,
				Description: "The Azure subscription id.",
				Computed:    true,
			},
			"client_id": {
				Type:        schema.TypeString,
				Description: "The client id (application id) of the Azure Service Principal.",
				Computed:    true,
			},
			"client_secret": {
				Type:        schema.TypeString,
				Description: "The client secret of the Azure Service Principal. Cannot be used with certificate_id.",
				Computed:    true,
			},
			"certificate_id": {
				Type:        schema.TypeString,
				Description: "The certificate reference of the Azure Service Principal, pointing to a Jenkins certificate credential. Cannot be used with client_secret.",
				Computed:    true,
			},
			"tenant": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceJenkinsCredentialAzureServicePrincipalRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// TODO: Get and populate object correctly.
	name := d.Get("name").(string)
	folderName := d.Get("folder").(string)
	subid := d.Get("subscription_id").(string)

	d.SetId(formatFolderName(folderName + "/" + name))
	d.Set("folder", folderName)
	d.Set("subscription_id", subid)
	d.Set("client_id", "")
	d.Set("client_secret", "")
	d.Set("certificate_id", "")
	d.Set("tenant", "")
	return resourceJenkinsCredentialAzureServicePrincipalRead(ctx, d, meta)
}
