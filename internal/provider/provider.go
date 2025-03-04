/* Copyright 2023 VMware, Inc.
   SPDX-License-Identifier: MPL-2.0 */

package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vmware/terraform-provider-vcf/internal/constants"
)

// Provider returns the resource configuration of the VCF provider.
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"sddc_manager_username": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Username to authenticate to SDDC Manager",
				DefaultFunc: schema.EnvDefaultFunc(constants.VcfTestUsername, nil),
			},
			"sddc_manager_password": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Password to authenticate to SDDC Manager",
				DefaultFunc: schema.EnvDefaultFunc(constants.VcfTestPassword, nil),
			},
			"sddc_manager_host": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Fully qualified domain name or IP address of the SDDC Manager",
				DefaultFunc: schema.EnvDefaultFunc(constants.VcfTestUrl, nil),
			},
			"allow_unverified_tls": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "If set, VMware VCF client will permit unverifiable TLS certificates.",
				DefaultFunc: schema.EnvDefaultFunc(constants.VcfTestAllowUnverifiedTls, false),
			},
		},

		DataSourcesMap: map[string]*schema.Resource{
			"vcf_domain":  DataSourceDomain(),
			"vcf_cluster": DataSourceCluster(),
		},

		ResourcesMap: map[string]*schema.Resource{
			"vcf_user":         ResourceUser(),
			"vcf_network_pool": ResourceNetworkPool(),
			"vcf_ceip":         ResourceCeip(),
			"vcf_host":         ResourceHost(),
			"vcf_domain":       ResourceDomain(),
			"vcf_cluster":      ResourceCluster(),
		},

		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(_ context.Context, data *schema.ResourceData) (interface{}, diag.Diagnostics) {
	username, isSetUsername := data.GetOk("sddc_manager_username")
	password, isSetPassword := data.GetOk("sddc_manager_password")
	hostName, isSetHost := data.GetOk("sddc_manager_host")
	if !isSetUsername || !isSetPassword || !isSetHost {
		return nil, diag.Errorf("SDDC Manager username, password, and host must be provided")
	}
	allowUnverifiedTls := data.Get("allow_unverified_tls")
	var newClient = NewSddcManagerClient(username.(string), password.(string), hostName.(string), allowUnverifiedTls.(bool))
	err := newClient.Connect()
	if err != nil {
		return nil, diag.FromErr(err)
	}
	return newClient, nil
}
