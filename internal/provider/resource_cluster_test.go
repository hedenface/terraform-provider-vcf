/*
 *  Copyright 2023 VMware, Inc.
 *    SPDX-License-Identifier: MPL-2.0
 */

package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/vmware/terraform-provider-vcf/internal/constants"
	validationUtils "github.com/vmware/terraform-provider-vcf/internal/validation"
	"github.com/vmware/vcf-sdk-go/client/clusters"
	"log"
	"os"
	"strings"
	"testing"
)

func TestAccResourceVcfCluster(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testCheckVcfClusterDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVcfClusterResourceConfig(
					os.Getenv(constants.VcfTestDomainDataSourceId),
					os.Getenv(constants.VcfTestHost5Fqdn),
					os.Getenv(constants.VcfTestHost5Pass),
					os.Getenv(constants.VcfTestHost6Fqdn),
					os.Getenv(constants.VcfTestHost6Pass),
					os.Getenv(constants.VcfTestHost7Fqdn),
					os.Getenv(constants.VcfTestHost7Pass),
					os.Getenv(constants.VcfTestEsxiLicenseKey),
					os.Getenv(constants.VcfTestVsanLicenseKey),
					"",
					""),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("vcf_cluster.cluster1", "name"),
					resource.TestCheckResourceAttrSet("vcf_cluster.cluster1", "primary_datastore_name"),
					resource.TestCheckResourceAttrSet("vcf_cluster.cluster1", "primary_datastore_type"),
					resource.TestCheckResourceAttrSet("vcf_cluster.cluster1", "is_default"),
					resource.TestCheckResourceAttrSet("vcf_cluster.cluster1", "is_stretched"),
					resource.TestCheckResourceAttrSet("vcf_cluster.cluster1", "host.0.id"),
					resource.TestCheckResourceAttrSet("vcf_cluster.cluster1", "host.1.id"),
					resource.TestCheckResourceAttrSet("vcf_cluster.cluster1", "host.2.id"),
				),
			},
			{
				ResourceName:     "vcf_cluster.cluster1",
				ImportState:      true,
				ImportStateCheck: clusterImportStateCheck,
			},
			{
				// add another host to the cluster
				Config: testAccVcfClusterResourceConfig(
					os.Getenv(constants.VcfTestDomainDataSourceId),
					os.Getenv(constants.VcfTestHost5Fqdn),
					os.Getenv(constants.VcfTestHost5Pass),
					os.Getenv(constants.VcfTestHost6Fqdn),
					os.Getenv(constants.VcfTestHost6Pass),
					os.Getenv(constants.VcfTestHost7Fqdn),
					os.Getenv(constants.VcfTestHost7Pass),
					os.Getenv(constants.VcfTestEsxiLicenseKey),
					os.Getenv(constants.VcfTestVsanLicenseKey),
					testAccVcfHostCommissionConfig(
						"host4",
						os.Getenv(constants.VcfTestHost8Fqdn),
						os.Getenv(constants.VcfTestHost8Pass)),
					testAccVcfHostInClusterConfig("host4",
						os.Getenv(constants.VcfTestEsxiLicenseKey),
						"sfo-m01-cl01")),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("vcf_cluster.cluster1", "name"),
					resource.TestCheckResourceAttrSet("vcf_cluster.cluster1", "primary_datastore_name"),
					resource.TestCheckResourceAttrSet("vcf_cluster.cluster1", "primary_datastore_type"),
					resource.TestCheckResourceAttrSet("vcf_cluster.cluster1", "is_default"),
					resource.TestCheckResourceAttrSet("vcf_cluster.cluster1", "is_stretched"),
					resource.TestCheckResourceAttrSet("vcf_cluster.cluster1", "host.0.id"),
					resource.TestCheckResourceAttrSet("vcf_cluster.cluster1", "host.1.id"),
					resource.TestCheckResourceAttrSet("vcf_cluster.cluster1", "host.2.id"),
					resource.TestCheckResourceAttrSet("vcf_cluster.cluster1", "host.3.id"),
				),
			},
			{
				// remove the added host
				Config: testAccVcfClusterResourceConfig(
					os.Getenv(constants.VcfTestDomainDataSourceId),
					os.Getenv(constants.VcfTestHost5Fqdn),
					os.Getenv(constants.VcfTestHost5Pass),
					os.Getenv(constants.VcfTestHost6Fqdn),
					os.Getenv(constants.VcfTestHost6Pass),
					os.Getenv(constants.VcfTestHost7Fqdn),
					os.Getenv(constants.VcfTestHost7Pass),
					os.Getenv(constants.VcfTestEsxiLicenseKey),
					os.Getenv(constants.VcfTestVsanLicenseKey),
					testAccVcfHostCommissionConfig(
						"host4",
						os.Getenv(constants.VcfTestHost8Fqdn),
						os.Getenv(constants.VcfTestHost8Pass)),
					""),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("vcf_cluster.cluster1", "name"),
					resource.TestCheckResourceAttrSet("vcf_cluster.cluster1", "primary_datastore_name"),
					resource.TestCheckResourceAttrSet("vcf_cluster.cluster1", "primary_datastore_type"),
					resource.TestCheckResourceAttrSet("vcf_cluster.cluster1", "is_default"),
					resource.TestCheckResourceAttrSet("vcf_cluster.cluster1", "is_stretched"),
					resource.TestCheckResourceAttrSet("vcf_cluster.cluster1", "host.0.id"),
					resource.TestCheckResourceAttrSet("vcf_cluster.cluster1", "host.1.id"),
					resource.TestCheckResourceAttrSet("vcf_cluster.cluster1", "host.2.id"),
				),
			},
		},
	})
}

func testAccVcfHostInClusterConfig(hostResourceId, esxLicenseKey, clusterName string) string {
	return fmt.Sprintf(
		`host {
		id = vcf_host.%s.id
		license_key = %q
		vmnic {
			id = "vmnic0"
			vds_name = "%s-vds01"
		}
		vmnic {
			id = "vmnic1"
			vds_name = "%s-vds01"
		}
	}	
	`, hostResourceId, esxLicenseKey, clusterName, clusterName)
}

func testAccVcfHostCommissionConfig(hostResourceId, hostFqdn, hostPass string) string {
	return fmt.Sprintf(`
	resource "vcf_host" %q {
		fqdn      = %q
		username  = "root"
		password  = %q
		network_pool_id = vcf_network_pool.domain_pool.id
		storage_type = "VSAN"
	}
	`, hostResourceId, hostFqdn, hostPass)
}

func testAccVcfClusterResourceConfig(domainId, host1Fqdn, host1Pass, host2Fqdn, host2Pass,
	host3Fqdn, host3Pass, esxLicenseKey, vsanLicenseKey,
	additionalCommissionHostConfig, additionalHostInClusterConfig string) string {
	return fmt.Sprintf(`
	resource "vcf_network_pool" "domain_pool" {
		name    = "cluster-pool"
		network {
			gateway   = "192.168.12.1"
			mask      = "255.255.255.0"
			mtu       = 9000
			subnet    = "192.168.12.0"
			type      = "VSAN"
			vlan_id   = 100
			ip_pools {
				start = "192.168.12.5"
				end   = "192.168.12.50"
			}
		}
		network {
			gateway   = "192.168.13.1"
			mask      = "255.255.255.0"
			mtu       = 9000
			subnet    = "192.168.13.0"
			type      = "vMotion"
			vlan_id   = 100
			ip_pools {
			  start = "192.168.13.5"
			  end   = "192.168.13.50"
			}
		  }
	}

	resource "vcf_host" "host1" {
		fqdn      = %q
		username  = "root"
		password  = %q
		network_pool_id = vcf_network_pool.domain_pool.id
		storage_type = "VSAN"
	}
	resource "vcf_host" "host2" {
		fqdn      = %q
		username  = "root"
		password  = %q
		network_pool_id = vcf_network_pool.domain_pool.id
		storage_type = "VSAN"
	}
	resource "vcf_host" "host3" {
		fqdn      = %q
		username  = "root"
		password  = %q
		network_pool_id = vcf_network_pool.domain_pool.id
		storage_type = "VSAN"
	}
	%s
	resource "vcf_cluster" "cluster1" {
		domain_id = %q
		name = "sfo-m01-cl01"
		host {
			id = vcf_host.host1.id
			license_key = %q
			vmnic {
				id = "vmnic0"
				vds_name = "sfo-m01-cl01-vds01"
			}
			vmnic {
				id = "vmnic1"
				vds_name = "sfo-m01-cl01-vds01"
			}
		}
		host {
			id = vcf_host.host2.id
			license_key = %q
			vmnic {
				id = "vmnic0"
				vds_name = "sfo-m01-cl01-vds01"
			}
			vmnic {
				id = "vmnic1"
				vds_name = "sfo-m01-cl01-vds01"
			}
		}
		host {
			id = vcf_host.host3.id
			license_key = %q
			vmnic {
				id = "vmnic0"
				vds_name = "sfo-m01-cl01-vds01"
			}
			vmnic {
				id = "vmnic1"
				vds_name = "sfo-m01-cl01-vds01"
			}
		}
		%s
		vds {
			name = "sfo-m01-cl01-vds01"
			portgroup {
				name = "sfo-m01-cl01-vds01-pg-mgmt"
				transport_type = "MANAGEMENT"
			}
			portgroup {
				name = "sfo-m01-cl01-vds01-pg-vsan"
				transport_type = "VSAN"
			}
			portgroup {
				name = "sfo-m01-cl01-vds01-pg-vmotion"
				transport_type = "VMOTION"
			}
		}
		vsan_datastore {
			datastore_name = "sfo-m01-cl01-ds-vsan01"
			failures_to_tolerate = 1
			license_key = %q
		}
		geneve_vlan_id = 3
	}
	`, host1Fqdn, host1Pass, host2Fqdn, host2Pass, host3Fqdn, host3Pass, additionalCommissionHostConfig, domainId,
		esxLicenseKey, esxLicenseKey, esxLicenseKey, additionalHostInClusterConfig, vsanLicenseKey)
}

func testCheckVcfClusterDestroy(state *terraform.State) error {
	vcfClient := testAccProvider.Meta().(*SddcManagerClient)
	apiClient := vcfClient.ApiClient

	for _, rs := range state.RootModule().Resources {
		if rs.Type != "vcf_cluster" {
			continue
		}

		clusterId := rs.Primary.Attributes["id"]
		getClusterParams := clusters.NewGetClusterParams().
			WithTimeout(constants.DefaultVcfApiCallTimeout).
			WithContext(context.TODO())
		getClusterParams.ID = clusterId

		clusterResult, err := apiClient.Clusters.GetCluster(getClusterParams)
		if err != nil && strings.Contains(err.Error(), "CLUSTER_NOT_FOUND") {
			log.Println("error = ", err)
			return nil
		}
		if clusterResult != nil && clusterResult.Payload != nil {
			return fmt.Errorf("domain with id %q not destroyed", clusterId)
		}

	}

	// Did not find the cluster
	return nil
}

func clusterImportStateCheck(states []*terraform.InstanceState) error {
	for _, state := range states {
		if state.Ephemeral.Type != "vcf_cluster" {
			continue
		}
		if state.Attributes["domain_id"] != os.Getenv(constants.VcfTestDomainDataSourceId) {
			return fmt.Errorf("cluster has wrong domain_id attribute set")
		}
		if validationUtils.IsEmpty(state.Attributes["id"]) {
			return fmt.Errorf("cluster has no id attribute set")
		}
		if state.Attributes["name"] != "sfo-m01-cl01" {
			return fmt.Errorf("cluster has wrong name attribute set")
		}
		if state.Attributes["primary_datastore_name"] != "sfo-m01-cl01-ds-vsan01" {
			return fmt.Errorf("cluster has wrong primary_datastore_name attribute set")
		}
		if state.Attributes["primary_datastore_type"] != "VSAN" {
			return fmt.Errorf("cluster has wrong primary_datastore_type attribute set")
		}
		if validationUtils.IsEmpty(state.Attributes["is_default"]) {
			return fmt.Errorf("cluster has no is_default attribute set")
		}
		if validationUtils.IsEmpty(state.Attributes["is_stretched"]) {
			return fmt.Errorf("cluster has no is_stretched attribute set")
		}
		if validationUtils.IsEmpty(state.Attributes["host.0.id"]) {
			return fmt.Errorf("cluster has no host.0.id attribute set")
		}
		if validationUtils.IsEmpty(state.Attributes["host.1.id"]) {
			return fmt.Errorf("cluster has no host.1.id attribute set")
		}
		if validationUtils.IsEmpty(state.Attributes["host.2.id"]) {
			return fmt.Errorf("cluster has no host.2.id attribute set")
		}
		if validationUtils.IsEmpty(state.Attributes["host.0.ip_address"]) {
			return fmt.Errorf("cluster has no host.0.ip_address attribute set")
		}
		if validationUtils.IsEmpty(state.Attributes["host.1.ip_address"]) {
			return fmt.Errorf("cluster has no host.1.ip_address attribute set")
		}
		if validationUtils.IsEmpty(state.Attributes["host.2.ip_address"]) {
			return fmt.Errorf("cluster has no host.2.ip_address attribute set")
		}
		if validationUtils.IsEmpty(state.Attributes["host.0.host_name"]) {
			return fmt.Errorf("cluster has no host.0.host_name attribute set")
		}
		if validationUtils.IsEmpty(state.Attributes["host.1.host_name"]) {
			return fmt.Errorf("cluster has no host.1.host_name attribute set")
		}
		if validationUtils.IsEmpty(state.Attributes["host.2.host_name"]) {
			return fmt.Errorf("cluster has no host.2.host_name attribute set")
		}
		return nil
	}
	return fmt.Errorf("cluster InstanceState not found! Import failed")
}
