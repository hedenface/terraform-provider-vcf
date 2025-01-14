---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "vcf_domain Data Source - terraform-provider-vcf"
subcategory: ""
description: |-
  
---

# vcf_domain (Data Source)

A workload domain is a policy based resource container with specific availability and performance attributes that combines compute (vSphere),
storage (vSAN/NFS/VMFS on FC/VVOL) and networking (NSX) into a single consumable entity.

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `domain_id` (String) The ID of the Domain to be used as data source

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `cluster` (List of Object) Specification representing the clusters in the workload domain (see [below for nested schema](#nestedatt--cluster))
- `id` (String) The ID of this resource.
- `is_management_sso_domain` (Boolean) Shows whether the domain is joined to the management domain SSO
- `name` (String) Name of the domain
- `nsx_cluster_ref` (List of Object) Represents NSX Manager cluster references associated with the domain (see [below for nested schema](#nestedatt--nsx_cluster_ref))
- `sso_id` (String) ID of the SSO domain associated with the workload domain
- `sso_name` (String) Name of the SSO domain associated with the workload domain
- `status` (String) Status of the workload domain
- `type` (String) Type of the workload domain
- `vcenter_fqdn` (String) Fully qualified domain name of the vCenter Server instance
- `vcenter_id` (String) ID of the vCenter Server instance

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `read` (String) 


<a id="nestedatt--cluster"></a>
### Nested Schema for `cluster`

Read-Only:

- `cluster_image_id` (String) ID of the cluster image to be used with the cluster
- `evc_mode` (String) Cluster EVC mode
- `geneve_vlan_id` (Number) VLAN ID use for NSX Geneve in the workload domain
- `high_availability_enabled` (Boolean) vSphere High Availability settings for the cluster
- `host` (List of Object) List of ESXi host information in the workload domain (see [below for nested schema](#nestedobjatt--cluster--host))
- `id` (String) ID of the cluster
- `is_default` (Boolean) Status of the cluster if default or not
- `is_stretched` (Boolean) Status of the cluster if stretched or not
- `name` (String) Name of the cluster in the workload domain
- `nfs_datastores` (List of Object) Cluster storage configuration for NFS (see [below for nested schema](#nestedobjatt--cluster--nfs_datastores))
- `primary_datastore_name` (String) Name of the primary datastore
- `primary_datastore_type` (String) Storage type of the primary datastore
- `vds` (List of Object) vSphere Distributed Switches in the cluster (see [below for nested schema](#nestedobjatt--cluster--vds))
- `vmfs_datastore` (List of Object) Cluster storage configuration for VMFS (see [below for nested schema](#nestedobjatt--cluster--vmfs_datastore))
- `vsan_datastore` (List of Object) Cluster storage configuration for vSAN (see [below for nested schema](#nestedobjatt--cluster--vsan_datastore))
- `vsan_remote_datastore_cluster` Cluster storage configuration for vSAN Remote Datastore (List of Object) (see [below for nested schema](#nestedobjatt--cluster--vsan_remote_datastore_cluster))
- `vvol_datastores` (List of Object) Cluster storage configuration for VVOL (see [below for nested schema](#nestedobjatt--cluster--vvol_datastores))

<a id="nestedobjatt--cluster--host"></a>
### Nested Schema for `cluster.host`

Read-Only:

- `availability_zone_name` (String) Availability Zone Name
- `host_name` (String) Host name of the ESXi host
- `id` (String) ID of the host (UUID)
- `ip_address` (String) IPv4 address of the ESXi host
- `license_key` (String) License key for an ESXi host
- `password` (String) Password to authenticate to the ESXi host
- `serial_number` (String) Serial number of the ESXi host
- `ssh_thumbprint` (String) SSH thumbprint of the ESXi host
- `username` (String) Username to authenticate to the ESXi host
- `vmnic` (List of Object) vmnic configuration for the ESXi host (see [below for nested schema](#nestedobjatt--cluster--host--vmnic))

<a id="nestedobjatt--cluster--host--vmnic"></a>
### Nested Schema for `cluster.host.vmnic`

Read-Only:

- `id` (String) ESXI host vmnic ID associated with a VDS
- `uplink` (String) Uplink associated with vmnic
- `vds_name` (String) Name of the VDS associated with the ESXi host



<a id="nestedobjatt--cluster--nfs_datastores"></a>
### Nested Schema for `cluster.nfs_datastores`

Read-Only:

- `datastore_name` (String) NFS datastore name used for cluster creation
- `path` (String) Shared directory path used for NFS based cluster creation
- `read_only` (Boolean) Readonly is used to identify whether to mount the directory as readOnly or not
- `server_name` (String) Fully qualified domain name or IP address of the NFS endpoint
- `user_tag` (String) User tag used to annotate NFS share


<a id="nestedobjatt--cluster--vds"></a>
### Nested Schema for `cluster.vds`

Read-Only:

- `name` (String) vSphere Distributed Switch name
- `is_used_by_nsx` (Boolean) Identifies if the vSphere distributed switch is used by NSX
- `nioc_bandwidth_allocations` (List of Object) List of Network I/O Control Bandwidth Allocations for System Traffic based on shares, reservation, and limit (see [below for nested schema](#nestedobjatt--cluster--vds--nioc_bandwidth_allocations))
- `portgroup` (List of Object) List of portgroups associated with the vSphere Distributed Switch (see [below for nested schema](#nestedobjatt--cluster--vds--portgroup))

<a id="nestedobjatt--cluster--vds--nioc_bandwidth_allocations"></a>
### Nested Schema for `cluster.vds.nioc_bandwidth_allocations`

Read-Only:

- `limit` (Number) The maximum allowed usage for a traffic class belonging to this resource pool per host physical NIC
- `reservation` (Number) Amount of bandwidth resource that is guaranteed available to the host infrastructure traffic class.
- `shares` (Number) The number of shares allocated. Used to determine resource allocation in case of resource contention.
- `shares_level` (String) The allocation level. The level is a simplified view of shares. Levels map to a pre-determined set of numeric values for shares.
- `type` (String) Host infrastructure traffic type.


<a id="nestedobjatt--cluster--vds--portgroup"></a>
### Nested Schema for `cluster.vds.portgroup`

Read-Only:

- `name` (String) Port group name
- `active_uplinks` (List of String) List of active uplinks associated with portgroup.
- `transport_type` (String) Port group transport type



<a id="nestedobjatt--cluster--vmfs_datastore"></a>
### Nested Schema for `cluster.vmfs_datastore`

Read-Only:

- `datastore_names` (List of String) VMFS datastore names used for VMFS on FC for cluster creation


<a id="nestedobjatt--cluster--vsan_datastore"></a>
### Nested Schema for `cluster.vsan_datastore`

Read-Only:

- `datastore_name` (String) vSAN datastore name
- `dedup_and_compression_enabled` (Boolean) Signals if vSAN deduplication and compression is enabled
- `failures_to_tolerate` (Number) Number of ESXi host failures to tolerate in the vSAN cluster
- `license_key` (String) vSAN license key used


<a id="nestedobjatt--cluster--vsan_remote_datastore_cluster"></a>
### Nested Schema for `cluster.vsan_remote_datastore_cluster`

Read-Only:

- `datastore_uuids` (List of String) vSAN HCI Mesh remote datastore UUIDs


<a id="nestedobjatt--cluster--vvol_datastores"></a>
### Nested Schema for `cluster.vvol_datastores`

Read-Only:

- `datastore_name` (String) vVol datastore name used
- `storage_container_id` (String) UUID of the VASA storage container
- `storage_protocol_type` (String) Type of the VASA storage protocol.
- `user_id` (String) UUID of the VASA storage user
- `vasa_provider_id` (String) UUID of the VASA storage provider



<a id="nestedatt--nsx_cluster_ref"></a>
### Nested Schema for `nsx_cluster_ref`

Read-Only:

- `id` (String) NSX Manager cluster ID
- `vip` (String) Virtual IP (VIP) for the NSX Manager cluster
- `vip_fqdn` (String) Fully qualified domain name of the NSX Manager cluster VIP


