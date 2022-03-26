


# MetaKube API.
This describes possible operations which can be made against the MetaKube API.
  

## Informations

### Version

2.18

### Terms Of Service

https://www.syseleven.de/agb-sla-metakube

## Content negotiation

### URI Schemes
  * https

### Consumes
  * application/json

### Produces
  * application/octet-stream
  * application/json
  * application/yaml

## Access control

### Security Schemes

#### api_key (header: Authorization)



> **Type**: apikey

### Security Requirements
  * api_key

## All endpoints

###  addon

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| POST | /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/addons | [create addon](#create-addon) |  |
| POST | /api/v2/projects/{project_id}/clusters/{cluster_id}/addons | [create addon v2](#create-addon-v2) |  |
| DELETE | /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/addons/{addon_id} | [delete addon](#delete-addon) | Deletes the given addon that belongs to the cluster. |
| DELETE | /api/v2/projects/{project_id}/clusters/{cluster_id}/addons/{addon_id} | [delete addon v2](#delete-addon-v2) | Deletes the given addon that belongs to the cluster. |
| GET | /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/addons/{addon_id} | [get addon](#get-addon) | Gets an addon that is assigned to the given cluster. |
| GET | /api/v2/projects/{project_id}/clusters/{cluster_id}/addons/{addon_id} | [get addon v2](#get-addon-v2) | Gets an addon that is assigned to the given cluster. |
| POST | /api/v1/addons | [list accessible addons](#list-accessible-addons) |  |
| GET | /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/addons | [list addons](#list-addons) |  |
| GET | /api/v2/projects/{project_id}/clusters/{cluster_id}/addons | [list addons v2](#list-addons-v2) |  |
| GET | /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/installableaddons | [list installable addons](#list-installable-addons) |  |
| GET | /api/v2/projects/{project_id}/clusters/{cluster_id}/installableaddons | [list installable addons v2](#list-installable-addons-v2) |  |
| PATCH | /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/addons/{addon_id} | [patch addon](#patch-addon) | Patches an addon that is assigned to the given cluster. |
| PATCH | /api/v2/projects/{project_id}/clusters/{cluster_id}/addons/{addon_id} | [patch addon v2](#patch-addon-v2) | Patches an addon that is assigned to the given cluster. |
  


###  admin

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| PUT | /api/v1/admin/metering/configurations | [create or update metering configurations](#create-or-update-metering-configurations) |  |
| PUT | /api/v1/admin/metering/credentials | [create or update metering credentials](#create-or-update-metering-credentials) |  |
| DELETE | /api/v1/admin/admission/plugins/{name} | [delete admission plugin](#delete-admission-plugin) | Deletes the admission plugin. |
| DELETE | /api/v1/admin/seeds/{seed_name} | [delete seed](#delete-seed) | Deletes the seed CRD object from the Kubermatic. |
| GET | /api/v1/admin | [get admins](#get-admins) | Returns list of admin users. |
| GET | /api/v1/admin/admission/plugins/{name} | [get admission plugin](#get-admission-plugin) | Gets the admission plugin. |
| GET | /api/v1/admin/settings/customlinks | [get kubermatic custom links](#get-kubermatic-custom-links) | Gets the custom links. |
| GET | /api/v1/admin/settings | [get kubermatic settings](#get-kubermatic-settings) | Gets the global settings. |
| GET | /api/v1/admin/seeds/{seed_name} | [get seed](#get-seed) | Returns the seed object. |
| GET | /api/v1/admin/admission/plugins | [list admission plugins](#list-admission-plugins) | Returns all admission plugins from the CRDs. |
| GET | /api/v1/admin/seeds | [list seeds](#list-seeds) | Returns all seeds from the CRDs. |
| PATCH | /api/v1/admin/settings | [patch kubermatic settings](#patch-kubermatic-settings) | Patches the global settings. |
| PUT | /api/v1/admin | [set admin](#set-admin) | Allows setting and clearing admin role for users. |
| PATCH | /api/v1/admin/admission/plugins/{name} | [update admission plugin](#update-admission-plugin) | Updates the admission plugin. |
| PATCH | /api/v1/admin/seeds/{seed_name} | [update seed](#update-seed) | Updates the seed. |
  


###  alibaba

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| GET | /api/v1/providers/alibaba/instancetypes | [list alibaba instance types](#list-alibaba-instance-types) | Lists available Alibaba instance types. |
| GET | /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/providers/alibaba/instancetypes | [list alibaba instance types no credentials](#list-alibaba-instance-types-no-credentials) |  |
| GET | /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/alibaba/instancetypes | [list alibaba instance types no credentials v2](#list-alibaba-instance-types-no-credentials-v2) |  |
| GET | /api/v1/providers/alibaba/vswitches | [list alibaba v switches](#list-alibaba-v-switches) | Lists available Alibaba vSwitches. |
| GET | /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/alibaba/vswitches | [list alibaba v switches no credentials v2](#list-alibaba-v-switches-no-credentials-v2) |  |
| GET | /api/v1/providers/alibaba/zones | [list alibaba zones](#list-alibaba-zones) | Lists available Alibaba zones. |
| GET | /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/providers/alibaba/zones | [list alibaba zones no credentials](#list-alibaba-zones-no-credentials) |  |
| GET | /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/alibaba/zones | [list alibaba zones no credentials v2](#list-alibaba-zones-no-credentials-v2) |  |
  


###  allowedregistries

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| DELETE | /api/v2/allowedregistries/{allowed_registry} | [delete allowed registry](#delete-allowed-registry) | Deletes the given allowed registry. |
| GET | /api/v2/allowedregistries/{allowed_registry} | [get allowed registry](#get-allowed-registry) |  |
| PATCH | /api/v2/allowedregistries/{allowed_registry} | [patch allowed registry](#patch-allowed-registry) |  |
  


###  allowedregistry

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| POST | /api/v2/allowedregistries | [create allowed registry](#create-allowed-registry) |  |
| GET | /api/v2/allowedregistries | [list allowed registries](#list-allowed-registries) | List allowed registries. |
  


###  anexia

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| GET | /api/v1/providers/anexia/templates | [list anexia templates](#list-anexia-templates) |  |
| GET | /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/anexia/templates | [list anexia templates no credentials v2](#list-anexia-templates-no-credentials-v2) |  |
| GET | /api/v1/providers/anexia/vlans | [list anexia vlans](#list-anexia-vlans) |  |
| GET | /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/anexia/vlans | [list anexia vlans no credentials v2](#list-anexia-vlans-no-credentials-v2) |  |
  


###  aws

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| GET | /api/v1/providers/aws/{dc}/securitygroups | [list a w s security groups](#list-a-w-s-security-groups) |  |
| GET | /api/v1/providers/aws/sizes | [list a w s sizes](#list-a-w-s-sizes) | Lists available AWS sizes. |
| GET | /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/providers/aws/sizes | [list a w s sizes no credentials](#list-a-w-s-sizes-no-credentials) |  |
| GET | /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/aws/sizes | [list a w s sizes no credentials v2](#list-a-w-s-sizes-no-credentials-v2) |  |
| GET | /api/v1/providers/aws/{dc}/subnets | [list a w s subnets](#list-a-w-s-subnets) |  |
| GET | /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/providers/aws/subnets | [list a w s subnets no credentials](#list-a-w-s-subnets-no-credentials) |  |
| GET | /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/aws/subnets | [list a w s subnets no credentials v2](#list-a-w-s-subnets-no-credentials-v2) |  |
| GET | /api/v1/providers/aws/{dc}/vpcs | [list a w s v p c s](#list-a-w-s-v-p-c-s) |  |
  


###  azure

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| GET | /api/v1/providers/azure/availabilityzones | [list azure availability zones](#list-azure-availability-zones) |  |
| GET | /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/providers/azure/availabilityzones | [list azure availability zones no credentials](#list-azure-availability-zones-no-credentials) |  |
| GET | /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/azure/availabilityzones | [list azure availability zones no credentials v2](#list-azure-availability-zones-no-credentials-v2) |  |
| GET | /api/v2/providers/azure/resourcegroups | [list azure resource groups](#list-azure-resource-groups) |  |
| GET | /api/v2/providers/azure/routetables | [list azure route tables](#list-azure-route-tables) |  |
| GET | /api/v2/providers/azure/securitygroups | [list azure security groups](#list-azure-security-groups) |  |
| GET | /api/v1/providers/azure/sizes | [list azure sizes](#list-azure-sizes) |  |
| GET | /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/providers/azure/sizes | [list azure sizes no credentials](#list-azure-sizes-no-credentials) |  |
| GET | /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/azure/sizes | [list azure sizes no credentials v2](#list-azure-sizes-no-credentials-v2) |  |
| GET | /api/v2/providers/azure/subnets | [list azure subnets](#list-azure-subnets) |  |
| GET | /api/v2/providers/azure/vnets | [list azure vnets](#list-azure-vnets) |  |
  


###  backupcredentials

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| PUT | /api/v2/seeds/{seed_name}/backupcredentials | [create or update backup credentials](#create-or-update-backup-credentials) |  |
  


###  constraint

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| POST | /api/v2/constraints | [create default constraint](#create-default-constraint) |  |
| GET | /api/v2/constraints/{constraint_name} | [get default constraint](#get-default-constraint) |  |
| GET | /api/v2/constraints | [list default constraint](#list-default-constraint) | List default constraint. |
| PATCH | /api/v2/constraints/{constraint_name} | [patch default constraint](#patch-default-constraint) |  |
  


###  constraints

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| DELETE | /api/v2/constraints/{constraint_name} | [delete default constraint](#delete-default-constraint) | Deletes a specified default constraint. |
  


###  constrainttemplates

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| POST | /api/v2/constrainttemplates | [create constraint template](#create-constraint-template) |  |
| DELETE | /api/v2/constrainttemplates/{ct_name} | [delete constraint template](#delete-constraint-template) |  |
| GET | /api/v2/constrainttemplates/{ct_name} | [get constraint template](#get-constraint-template) |  |
| GET | /api/v2/constrainttemplates | [list constraint templates](#list-constraint-templates) | List constraint templates. |
| PATCH | /api/v2/constrainttemplates/{ct_name} | [patch constraint template](#patch-constraint-template) |  |
  


###  credentials

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| GET | /api/v1/providers/{provider_name}/presets/credentials | [list credentials](#list-credentials) |  |
  


###  datacenter

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| POST | /api/v1/seed/{seed_name}/dc | [create d c](#create-d-c) | Create the datacenter for a specified seed. |
| DELETE | /api/v1/seed/{seed_name}/dc/{dc} | [delete d c](#delete-d-c) | Delete the datacenter. |
| GET | /api/v1/providers/{provider_name}/dc/{dc} | [get d c for provider](#get-d-c-for-provider) | Get the datacenter for the specified provider. |
| GET | /api/v1/seed/{seed_name}/dc/{dc} | [get d c for seed](#get-d-c-for-seed) | Returns the specified datacenter for the specified seed. |
| GET | /api/v1/dc/{dc} | [get datacenter](#get-datacenter) |  |
| GET | /api/v1/providers/{provider_name}/dc | [list d c for provider](#list-d-c-for-provider) | Returns all datacenters for the specified provider. |
| GET | /api/v1/seed/{seed_name}/dc | [list d c for seed](#list-d-c-for-seed) | Returns all datacenters for the specified seed. |
| GET | /api/v1/dc | [list datacenters](#list-datacenters) |  |
| PATCH | /api/v1/seed/{seed_name}/dc/{dc} | [patch d c](#patch-d-c) | Patch the datacenter. |
| PUT | /api/v1/seed/{seed_name}/dc/{dc} | [update d c](#update-d-c) | Update the datacenter. The datacenter spec will be overwritten with the one provided in the request. |
  


###  digitalocean

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| GET | /api/v1/providers/digitalocean/sizes | [list digitalocean sizes](#list-digitalocean-sizes) |  |
| GET | /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/providers/digitalocean/sizes | [list digitalocean sizes no credentials](#list-digitalocean-sizes-no-credentials) |  |
| GET | /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/digitalocean/sizes | [list digitalocean sizes no credentials v2](#list-digitalocean-sizes-no-credentials-v2) |  |
  


###  etcdbackupconfig

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| POST | /api/v2/projects/{project_id}/clusters/{cluster_id}/etcdbackupconfigs | [create etcd backup config](#create-etcd-backup-config) |  |
| DELETE | /api/v2/projects/{project_id}/clusters/{cluster_id}/etcdbackupconfigs/{ebc_id} | [delete etcd backup config](#delete-etcd-backup-config) |  |
| GET | /api/v2/projects/{project_id}/clusters/{cluster_id}/etcdbackupconfigs/{ebc_id} | [get etcd backup config](#get-etcd-backup-config) |  |
| GET | /api/v2/projects/{project_id}/clusters/{cluster_id}/etcdbackupconfigs | [list etcd backup config](#list-etcd-backup-config) |  |
| GET | /api/v2/projects/{project_id}/etcdbackupconfigs | [list project etcd backup config](#list-project-etcd-backup-config) |  |
| PATCH | /api/v2/projects/{project_id}/clusters/{cluster_id}/etcdbackupconfigs/{ebc_id} | [patch etcd backup config](#patch-etcd-backup-config) |  |
  


###  etcdrestore

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| POST | /api/v2/projects/{project_id}/clusters/{cluster_id}/etcdrestores | [create etcd restore](#create-etcd-restore) |  |
| DELETE | /api/v2/projects/{project_id}/clusters/{cluster_id}/etcdrestores/{er_name} | [delete etcd restore](#delete-etcd-restore) |  |
| GET | /api/v2/projects/{project_id}/clusters/{cluster_id}/etcdrestores/{er_name} | [get etcd restore](#get-etcd-restore) |  |
| GET | /api/v2/projects/{project_id}/clusters/{cluster_id}/etcdrestores | [list etcd restore](#list-etcd-restore) |  |
| GET | /api/v2/projects/{project_id}/etcdrestores | [list project etcd restore](#list-project-etcd-restore) |  |
  


###  gcp

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| GET | /api/v1/providers/gcp/disktypes | [list g c p disk types](#list-g-c-p-disk-types) |  |
| GET | /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/providers/gcp/disktypes | [list g c p disk types no credentials](#list-g-c-p-disk-types-no-credentials) |  |
| GET | /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/gcp/disktypes | [list g c p disk types no credentials v2](#list-g-c-p-disk-types-no-credentials-v2) |  |
| GET | /api/v1/providers/gcp/networks | [list g c p networks](#list-g-c-p-networks) |  |
| GET | /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/providers/gcp/networks | [list g c p networks no credentials](#list-g-c-p-networks-no-credentials) |  |
| GET | /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/gcp/networks | [list g c p networks no credentials v2](#list-g-c-p-networks-no-credentials-v2) |  |
| GET | /api/v1/providers/gcp/sizes | [list g c p sizes](#list-g-c-p-sizes) |  |
| GET | /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/providers/gcp/sizes | [list g c p sizes no credentials](#list-g-c-p-sizes-no-credentials) |  |
| GET | /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/gcp/sizes | [list g c p sizes no credentials v2](#list-g-c-p-sizes-no-credentials-v2) |  |
| GET | /api/v1/providers/gcp/{dc}/subnetworks | [list g c p subnetworks](#list-g-c-p-subnetworks) |  |
| GET | /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/providers/gcp/subnetworks | [list g c p subnetworks no credentials](#list-g-c-p-subnetworks-no-credentials) |  |
| GET | /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/gcp/subnetworks | [list g c p subnetworks no credentials v2](#list-g-c-p-subnetworks-no-credentials-v2) |  |
| GET | /api/v1/providers/gcp/{dc}/zones | [list g c p zones](#list-g-c-p-zones) |  |
| GET | /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/providers/gcp/zones | [list g c p zones no credentials](#list-g-c-p-zones-no-credentials) |  |
| GET | /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/gcp/zones | [list g c p zones no credentials v2](#list-g-c-p-zones-no-credentials-v2) |  |
  


###  hetzner

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| GET | /api/v1/providers/hetzner/sizes | [list hetzner sizes](#list-hetzner-sizes) |  |
| GET | /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/providers/hetzner/sizes | [list hetzner sizes no credentials](#list-hetzner-sizes-no-credentials) |  |
| GET | /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/hetzner/sizes | [list hetzner sizes no credentials v2](#list-hetzner-sizes-no-credentials-v2) |  |
  


###  metering

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| GET | /api/v1/admin/metering/reports/{report_name} | [get metering report](#get-metering-report) |  |
| GET | /api/v1/admin/metering/reports | [list metering reports](#list-metering-reports) |  |
  


###  metric

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| GET | /api/v2/projects/{project_id}/clusters/{cluster_id}/machinedeployments/{machinedeployment_id}/nodes/metrics | [list machine deployment metrics](#list-machine-deployment-metrics) | Lists metrics that belong to the given machine deployment. |
| GET | /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/nodedeployments/{nodedeployment_id}/nodes/metrics | [list node deployment metrics](#list-node-deployment-metrics) | Lists metrics that belong to the given node deployment. |
  


###  mlaadminsetting

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| POST | /api/v2/projects/{project_id}/clusters/{cluster_id}/mlaadminsetting | [create m l a admin setting](#create-m-l-a-admin-setting) |  |
| DELETE | /api/v2/projects/{project_id}/clusters/{cluster_id}/mlaadminsetting | [delete m l a admin setting](#delete-m-l-a-admin-setting) | Deletes the MLA admin setting that belongs to the cluster. |
| GET | /api/v2/projects/{project_id}/clusters/{cluster_id}/mlaadminsetting | [get m l a admin setting](#get-m-l-a-admin-setting) | Gets MLA Admin settings for the given cluster. |
| PUT | /api/v2/projects/{project_id}/clusters/{cluster_id}/mlaadminsetting | [update m l a admin setting](#update-m-l-a-admin-setting) | Updates the MLA admin setting for the given cluster. |
  


###  openstack

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| GET | /api/v1/providers/openstack/availabilityzones | [list openstack availability zones](#list-openstack-availability-zones) |  |
| GET | /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/providers/openstack/availabilityzones | [list openstack availability zones no credentials](#list-openstack-availability-zones-no-credentials) |  |
| GET | /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/openstack/availabilityzones | [list openstack availability zones no credentials v2](#list-openstack-availability-zones-no-credentials-v2) |  |
| GET | /api/v1/providers/openstack/images | [list openstack images](#list-openstack-images) |  |
| GET | /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/providers/openstack/images | [list openstack images no credentials](#list-openstack-images-no-credentials) |  |
| GET | /api/v1/providers/openstack/networks | [list openstack networks](#list-openstack-networks) |  |
| GET | /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/providers/openstack/networks | [list openstack networks no credentials](#list-openstack-networks-no-credentials) |  |
| GET | /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/openstack/networks | [list openstack networks no credentials v2](#list-openstack-networks-no-credentials-v2) |  |
| GET | /api/v1/providers/openstack/quotalimits | [list openstack quota limits](#list-openstack-quota-limits) |  |
| GET | /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/providers/openstack/quotalimits | [list openstack quota limits no credentials](#list-openstack-quota-limits-no-credentials) |  |
| GET | /api/v1/providers/openstack/securitygroups | [list openstack security groups](#list-openstack-security-groups) |  |
| GET | /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/providers/openstack/securitygroups | [list openstack security groups no credentials](#list-openstack-security-groups-no-credentials) |  |
| GET | /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/openstack/securitygroups | [list openstack security groups no credentials v2](#list-openstack-security-groups-no-credentials-v2) |  |
| GET | /api/v1/providers/openstack/sizes | [list openstack sizes](#list-openstack-sizes) |  |
| GET | /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/providers/openstack/sizes | [list openstack sizes no credentials](#list-openstack-sizes-no-credentials) |  |
| GET | /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/openstack/sizes | [list openstack sizes no credentials v2](#list-openstack-sizes-no-credentials-v2) |  |
| GET | /api/v1/providers/openstack/subnets | [list openstack subnets](#list-openstack-subnets) |  |
| GET | /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/providers/openstack/subnets | [list openstack subnets no credentials](#list-openstack-subnets-no-credentials) |  |
| GET | /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/openstack/subnets | [list openstack subnets no credentials v2](#list-openstack-subnets-no-credentials-v2) |  |
| GET | /api/v1/providers/openstack/tenants | [list openstack tenants](#list-openstack-tenants) |  |
| GET | /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/providers/openstack/tenants | [list openstack tenants no credentials](#list-openstack-tenants-no-credentials) |  |
| GET | /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/openstack/tenants | [list openstack tenants no credentials v2](#list-openstack-tenants-no-credentials-v2) |  |
  


###  operations

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| GET | /api/v1/kubeconfig | [create o ID c kubeconfig](#create-o-id-c-kubeconfig) |  |
| GET | /api/v1/addonconfigs/{addon_id} | [get addon config](#get-addon-config) | Returns specified addon config. |
| GET | /api/v1/admission/plugins/{version} | [get admission plugins](#get-admission-plugins) | Returns specified addon config. |
| GET | /api/v1/addonconfigs | [list addon configs](#list-addon-configs) | Returns all available addon configs. |
| GET | /api/v1/labels/system | [list system labels](#list-system-labels) |  |
| POST | /api/v2/projects/{project_id}/clusters/{cluster_id}/externalccmmigration | [migrate cluster to external c c m](#migrate-cluster-to-external-c-c-m) |  |
  


###  packet

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| GET | /api/v1/providers/packet/sizes | [list packet sizes](#list-packet-sizes) |  |
| GET | /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/providers/packet/sizes | [list packet sizes no credentials](#list-packet-sizes-no-credentials) |  |
| GET | /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/packet/sizes | [list packet sizes no credentials v2](#list-packet-sizes-no-credentials-v2) |  |
  


###  preset

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| POST | /api/v2/providers/{provider_name}/presets | [create preset](#create-preset) |  |
| GET | /api/v2/presets | [list presets](#list-presets) |  |
| GET | /api/v2/providers/{provider_name}/presets | [list provider presets](#list-provider-presets) |  |
| PUT | /api/v2/providers/{provider_name}/presets | [update preset](#update-preset) |  |
| PUT | /api/v2/presets/{preset_name}/status | [update preset status](#update-preset-status) | Updates the status of a preset. It can enable or disable it, so that it won't be listed by the list endpoints. |
  


###  project

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| PUT | /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/sshkeys/{key_id} | [assign SSH key to cluster](#assign-ssh-key-to-cluster) |  |
| PUT | /api/v2/projects/{project_id}/clusters/{cluster_id}/sshkeys/{key_id} | [assign SSH key to cluster v2](#assign-ssh-key-to-cluster-v2) |  |
| POST | /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/clusterroles/{role_id}/clusterbindings | [bind user to cluster role](#bind-user-to-cluster-role) |  |
| POST | /api/v2/projects/{project_id}/clusters/{cluster_id}/clusterroles/{role_id}/clusterbindings | [bind user to cluster role v2](#bind-user-to-cluster-role-v2) |  |
| POST | /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/roles/{namespace}/{role_id}/bindings | [bind user to role](#bind-user-to-role) |  |
| POST | /api/v2/projects/{project_id}/clusters/{cluster_id}/roles/{namespace}/{role_id}/bindings | [bind user to role v2](#bind-user-to-role-v2) |  |
| POST | /api/v1/projects/{project_id}/dc/{dc}/clusters | [create cluster](#create-cluster) | Creates a cluster for the given project. |
| POST | /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/clusterroles | [create cluster role](#create-cluster-role) |  |
| POST | /api/v2/projects/{project_id}/clustertemplates | [create cluster template](#create-cluster-template) | Creates a cluster templates for the given project. |
| POST | /api/v2/projects/{project_id}/clustertemplates/{template_id}/instances | [create cluster template instance](#create-cluster-template-instance) | Create cluster template instance. |
| POST | /api/v2/projects/{project_id}/clusters | [create cluster v2](#create-cluster-v2) | Creates a cluster for the given project. |
| POST | /api/v2/projects/{project_id}/clusters/{cluster_id}/constraints | [create constraint](#create-constraint) | Creates a given constraint for the specified cluster. |
| POST | /api/v2/projects/{project_id}/kubernetes/clusters | [create external cluster](#create-external-cluster) | Creates an external cluster for the given project. |
| POST | /api/v2/projects/{project_id}/clusters/{cluster_id}/gatekeeper/config | [create gatekeeper config](#create-gatekeeper-config) |  |
| POST | /api/v2/projects/{project_id}/clusters/{cluster_id}/machinedeployments | [create machine deployment](#create-machine-deployment) |  |
| POST | /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/nodedeployments | [create node deployment](#create-node-deployment) |  |
| POST | /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/ndrequests | [create node deployment request](#create-node-deployment-request) |  |
| POST | /api/v1/projects | [create project](#create-project) | Creates a brand new project. |
| POST | /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/roles | [create role](#create-role) |  |
| POST | /api/v1/projects/{project_id}/sshkeys | [create SSH key](#create-ssh-key) | Adds the given SSH key to the specified project. |
| DELETE | /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id} | [delete cluster](#delete-cluster) |  |
| DELETE | /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/clusterroles/{role_id} | [delete cluster role](#delete-cluster-role) |  |
| DELETE | /api/v2/projects/{project_id}/clustertemplates/{template_id} | [delete cluster template](#delete-cluster-template) | Delete cluster template. |
| DELETE | /api/v2/projects/{project_id}/clusters/{cluster_id} | [delete cluster v2](#delete-cluster-v2) |  |
| DELETE | /api/v2/projects/{project_id}/clusters/{cluster_id}/constraints/{constraint_name} | [delete constraint](#delete-constraint) | Deletes a specified constraint for the given cluster. |
| DELETE | /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id} | [delete external cluster](#delete-external-cluster) |  |
| DELETE | /api/v2/projects/{project_id}/clusters/{cluster_id}/gatekeeper/config | [delete gatekeeper config](#delete-gatekeeper-config) | Deletes the gatekeeper sync config for the specified cluster. |
| DELETE | /api/v2/projects/{project_id}/clusters/{cluster_id}/machinedeployments/{machinedeployment_id} | [delete machine deployment](#delete-machine-deployment) | Deletes the given machine deployment that belongs to the cluster. |
| DELETE | /api/v2/projects/{project_id}/clusters/{cluster_id}/machinedeployments/nodes/{node_id} | [delete machine deployment node](#delete-machine-deployment-node) | Deletes the given node that belongs to the machine deployment. |
| DELETE | /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/nodedeployments/{nodedeployment_id} | [delete node deployment](#delete-node-deployment) | Deletes the given node deployment that belongs to the cluster. |
| DELETE | /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/ndrequests/{ndrequest_id} | [delete node deployment request](#delete-node-deployment-request) | Deletes the given NodeDeploymentRequest that belongs to the cluster. |
| DELETE | /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/nodes/{node_id} | [delete node for cluster legacy](#delete-node-for-cluster-legacy) | Deprecated:
Deletes the given node that belongs to the cluster. |
| DELETE | /api/v1/projects/{project_id} | [delete project](#delete-project) | Deletes the project with the given ID. |
| DELETE | /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/roles/{namespace}/{role_id} | [delete role](#delete-role) |  |
| DELETE | /api/v1/projects/{project_id}/sshkeys/{key_id} | [delete SSH key](#delete-ssh-key) | Removes the given SSH Key from the system. |
| DELETE | /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/sshkeys/{key_id} | [detach SSH key from cluster](#detach-ssh-key-from-cluster) |  |
| DELETE | /api/projects/{project_id}/clusters/{cluster_id}/sshkeys/{key_id} | [detach SSH key from cluster v2](#detach-ssh-key-from-cluster-v2) |  |
| GET | /api/v2/projects/{project_id}/clusters/{cluster_id}/alertmanager/config | [get alertmanager](#get-alertmanager) | Gets the alertmanager configuration for the specified cluster. |
| GET | /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id} | [get cluster](#get-cluster) |  |
| GET | /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/events | [get cluster events](#get-cluster-events) | Gets the events related to the specified cluster. |
| GET | /api/v2/projects/{project_id}/clusters/{cluster_id}/events | [get cluster events v2](#get-cluster-events-v2) | Gets the events related to the specified cluster. |
| GET | /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/health | [get cluster health](#get-cluster-health) |  |
| GET | /api/v2/projects/{project_id}/clusters/{cluster_id}/health | [get cluster health v2](#get-cluster-health-v2) |  |
| GET | /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/kubeconfig | [get cluster kubeconfig](#get-cluster-kubeconfig) | Gets the kubeconfig for the specified cluster. |
| GET | /api/v2/projects/{project_id}/clusters/{cluster_id}/kubeconfig | [get cluster kubeconfig v2](#get-cluster-kubeconfig-v2) | Gets the kubeconfig for the specified cluster. |
| GET | /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/metrics | [get cluster metrics](#get-cluster-metrics) |  |
| GET | /api/v2/projects/{project_id}/clusters/{cluster_id}/metrics | [get cluster metrics v2](#get-cluster-metrics-v2) |  |
| GET | /api/v2/projects/{project_id}/clusters/{cluster_id}/oidc | [get cluster oidc](#get-cluster-oidc) | Gets the OIDC params for the specified cluster with OIDC authentication. |
| GET | /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/roles/{role_id} | [get cluster role](#get-cluster-role) |  |
| GET | /api/v2/projects/{project_id}/clustertemplates/{template_id} | [get cluster template](#get-cluster-template) | Get cluster template. |
| GET | /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/upgrades | [get cluster upgrades](#get-cluster-upgrades) |  |
| GET | /api/v2/projects/{project_id}/clusters/{cluster_id}/upgrades | [get cluster upgrades v2](#get-cluster-upgrades-v2) |  |
| GET | /api/v2/projects/{project_id}/clusters/{cluster_id} | [get cluster v2](#get-cluster-v2) |  |
| GET | /api/v2/projects/{project_id}/clusters/{cluster_id}/constraints/{constraint_name} | [get constraint](#get-constraint) | Gets an specified constraint for the given cluster. |
| GET | /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id} | [get external cluster](#get-external-cluster) | Gets an external cluster for the given project. |
| GET | /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/metrics | [get external cluster metrics](#get-external-cluster-metrics) |  |
| GET | /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/nodes/{node_id} | [get external cluster node](#get-external-cluster-node) | Gets an external cluster node. |
| GET | /api/v2/projects/{project_id}/clusters/{cluster_id}/gatekeeper/config | [get gatekeeper config](#get-gatekeeper-config) | Gets the gatekeeper sync config for the specified cluster. |
| GET | /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/kubeloginkubeconfig | [get kube login cluster kubeconfig](#get-kube-login-cluster-kubeconfig) | Gets the kubeconfig for the specified cluster with oidc authentication that works nicely with kube-login. |
| GET | /api/v2/projects/{project_id}/clusters/{cluster_id}/machinedeployments/{machinedeployment_id} | [get machine deployment](#get-machine-deployment) | Gets a machine deployment that is assigned to the given cluster. |
| GET | /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/nodedeployments/{nodedeployment_id} | [get node deployment](#get-node-deployment) | Gets a node deployment that is assigned to the given cluster. |
| GET | /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/ndrequests/{ndrequest_id} | [get node deployment request](#get-node-deployment-request) | Gets a NodeDeploymentRequest that is assigned to the given cluster. |
| GET | /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/oidckubeconfig | [get oidc cluster kubeconfig](#get-oidc-cluster-kubeconfig) | Gets the kubeconfig for the specified cluster with oidc authentication. |
| GET | /api/v2/projects/{project_id}/clusters/{cluster_id}/oidckubeconfig | [get oidc cluster kubeconfig v2](#get-oidc-cluster-kubeconfig-v2) | Gets the kubeconfig for the specified cluster with oidc authentication. |
| GET | /api/v1/projects/{project_id} | [get project](#get-project) |  |
| GET | /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/roles/{namespace}/{role_id} | [get role](#get-role) |  |
| GET | /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/clusterroles | [list cluster role](#list-cluster-role) |  |
| GET | /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/clusterbindings | [list cluster role binding](#list-cluster-role-binding) |  |
| GET | /api/v2/projects/{project_id}/clusters/{cluster_id}/clusterbindings | [list cluster role binding v2](#list-cluster-role-binding-v2) |  |
| GET | /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/clusterrolenames | [list cluster role names](#list-cluster-role-names) |  |
| GET | /api/v2/projects/{project_id}/clusters/{cluster_id}/clusterrolenames | [list cluster role names v2](#list-cluster-role-names-v2) |  |
| GET | /api/v2/projects/{project_id}/clusters/{cluster_id}/clusterroles | [list cluster role v2](#list-cluster-role-v2) |  |
| GET | /api/v2/projects/{project_id}/clustertemplates | [list cluster templates](#list-cluster-templates) | List cluster templates for the given project. |
| GET | /api/v1/projects/{project_id}/dc/{dc}/clusters | [list clusters](#list-clusters) | Lists clusters for the specified project and data center. |
| GET | /api/v1/projects/{project_id}/clusters | [list clusters for project](#list-clusters-for-project) | Lists clusters for the specified project. |
| GET | /api/v2/projects/{project_id}/clusters | [list clusters v2](#list-clusters-v2) | Lists clusters for the specified project. |
| GET | /api/v2/projects/{project_id}/clusters/{cluster_id}/constraints | [list constraints](#list-constraints) | Lists constraints for the specified cluster. |
| GET | /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/events | [list external cluster events](#list-external-cluster-events) | Gets an external cluster events. |
| GET | /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/nodes | [list external cluster nodes](#list-external-cluster-nodes) | Gets an external cluster nodes. |
| GET | /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/nodesmetrics | [list external cluster nodes metrics](#list-external-cluster-nodes-metrics) | Gets an external cluster nodes metrics. |
| GET | /api/v2/projects/{project_id}/kubernetes/clusters | [list external clusters](#list-external-clusters) | Lists external clusters for the specified project. |
| GET | /api/v2/projects/{project_id}/clusters/{cluster_id}/machinedeployments/{machinedeployment_id}/nodes | [list machine deployment nodes](#list-machine-deployment-nodes) | Lists nodes that belong to the given machine deployment. |
| GET | /api/v2/projects/{project_id}/clusters/{cluster_id}/machinedeployments/{machinedeployment_id}/nodes/events | [list machine deployment nodes events](#list-machine-deployment-nodes-events) | Lists machine deployment events. If query parameter `type` is set to `warning` then only warning events are retrieved. |
| GET | /api/v2/projects/{project_id}/clusters/{cluster_id}/machinedeployments | [list machine deployments](#list-machine-deployments) |  |
| GET | /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/namespaces | [list namespace](#list-namespace) |  |
| GET | /api/v2/projects/{project_id}/clusters/{cluster_id}/namespaces | [list namespace v2](#list-namespace-v2) |  |
| GET | /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/nodedeployments/{nodedeployment_id}/nodes | [list node deployment nodes](#list-node-deployment-nodes) | Lists nodes that belong to the given node deployment. |
| GET | /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/nodedeployments/{nodedeployment_id}/nodes/events | [list node deployment nodes events](#list-node-deployment-nodes-events) | Lists node deployment events. If query parameter `type` is set to `warning` then only warning events are retrieved. |
| GET | /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/ndrequests | [list node deployment requests](#list-node-deployment-requests) |  |
| GET | /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/nodedeployments | [list node deployments](#list-node-deployments) |  |
| GET | /api/v2/projects/{project_id}/clusters/{cluster_id}/nodes | [list nodes for cluster](#list-nodes-for-cluster) | This endpoint is used for kubeadm cluster. |
| GET | /api/v1/projects | [list projects](#list-projects) | Lists projects that an authenticated user is a member of. |
| GET | /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/roles | [list role](#list-role) |  |
| GET | /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/bindings | [list role binding](#list-role-binding) |  |
| GET | /api/v2/projects/{project_id}/clusters/{cluster_id}/bindings | [list role binding v2](#list-role-binding-v2) |  |
| GET | /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/rolenames | [list role names](#list-role-names) |  |
| GET | /api/v2/projects/{project_id}/clusters/{cluster_id}/rolenames | [list role names v2](#list-role-names-v2) |  |
| GET | /api/v2/projects/{project_id}/clusters/{cluster_id}/roles | [list role v2](#list-role-v2) |  |
| GET | /api/v1/projects/{project_id}/sshkeys | [list SSH keys](#list-ssh-keys) | Lists SSH Keys that belong to the given project. |
| GET | /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/sshkeys | [list SSH keys assigned to cluster](#list-ssh-keys-assigned-to-cluster) |  |
| GET | /api/v2/projects/{project_id}/clusters/{cluster_id}/sshkeys | [list SSH keys assigned to cluster v2](#list-ssh-keys-assigned-to-cluster-v2) |  |
| PATCH | /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id} | [patch cluster](#patch-cluster) | Patches the given cluster using JSON Merge Patch method (https://tools.ietf.org/html/rfc7396). |
| PATCH | /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/clusterroles/{role_id} | [patch cluster role](#patch-cluster-role) |  |
| PATCH | /api/v2/projects/{project_id}/clusters/{cluster_id} | [patch cluster v2](#patch-cluster-v2) | Patches the given cluster using JSON Merge Patch method (https://tools.ietf.org/html/rfc7396). |
| PATCH | /api/v2/projects/{project_id}/clusters/{cluster_id}/constraints/{constraint_name} | [patch constraint](#patch-constraint) | Patches a given constraint for the specified cluster. |
| PATCH | /api/v2/projects/{project_id}/clusters/{cluster_id}/gatekeeper/config | [patch gatekeeper config](#patch-gatekeeper-config) | Patches the gatekeeper config for the specified cluster. |
| PATCH | /api/v2/projects/{project_id}/clusters/{cluster_id}/machinedeployments/{machinedeployment_id} | [patch machine deployment](#patch-machine-deployment) |  |
| PATCH | /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/nodedeployments/{nodedeployment_id} | [patch node deployment](#patch-node-deployment) |  |
| PATCH | /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/ndrequests/{ndrequest_id} | [patch node deployment request](#patch-node-deployment-request) | Patches a NodeDeploymentRequest that is assigned to the given cluster. |
| PATCH | /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/roles/{namespace}/{role_id} | [patch role](#patch-role) |  |
| DELETE | /api/v2/projects/{project_id}/clusters/{cluster_id}/alertmanager/config | [reset alertmanager](#reset-alertmanager) | Resets the alertmanager configuration to default for the specified cluster. |
| POST | /api/v2/projects/{project_id}/clusters/{cluster_id}/machinedeployments/{machinedeployment_id} | [restart machine deployment](#restart-machine-deployment) | Schedules rolling restart of a machine deployment that is assigned to the given cluster. |
| PUT | /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/token | [revoke cluster admin token](#revoke-cluster-admin-token) |  |
| PUT | /api/v2/projects/{project_id}/clusters/{cluster_id}/token | [revoke cluster admin token v2](#revoke-cluster-admin-token-v2) |  |
| PUT | /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/viewertoken | [revoke cluster viewer token](#revoke-cluster-viewer-token) |  |
| PUT | /api/v2/projects/{project_id}/clusters/{cluster_id}/viewertoken | [revoke cluster viewer token v2](#revoke-cluster-viewer-token-v2) |  |
| DELETE | /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/clusterroles/{role_id}/clusterbindings | [unbind user from cluster role binding](#unbind-user-from-cluster-role-binding) |  |
| DELETE | /api/v2/projects/{project_id}/clusters/{cluster_id}/clusterroles/{role_id}/clusterbindings | [unbind user from cluster role binding v2](#unbind-user-from-cluster-role-binding-v2) |  |
| DELETE | /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/roles/{namespace}/{role_id}/bindings | [unbind user from role binding](#unbind-user-from-role-binding) |  |
| DELETE | /api/v2/projects/{project_id}/clusters/{cluster_id}/roles/{namespace}/{role_id}/bindings | [unbind user from role binding v2](#unbind-user-from-role-binding-v2) |  |
| PUT | /api/v2/projects/{project_id}/clusters/{cluster_id}/alertmanager/config | [update alertmanager](#update-alertmanager) |  |
| PUT | /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id} | [update external cluster](#update-external-cluster) | Updates an external cluster for the given project. |
| PUT | /api/v1/projects/{project_id} | [update project](#update-project) |  |
| PUT | /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/nodes/upgrades | [upgrade cluster node deployments](#upgrade-cluster-node-deployments) |  |
| PUT | /api/v2/projects/{project_id}/clusters/{cluster_id}/nodes/upgrades | [upgrade cluster node deployments v2](#upgrade-cluster-node-deployments-v2) |  |
  


###  rulegroup

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| POST | /api/v2/projects/{project_id}/clusters/{cluster_id}/rulegroups | [create rule group](#create-rule-group) |  |
| DELETE | /api/v2/projects/{project_id}/clusters/{cluster_id}/rulegroups/{rulegroup_id} | [delete rule group](#delete-rule-group) | Deletes the given rule group that belongs to the cluster. |
| GET | /api/v2/projects/{project_id}/clusters/{cluster_id}/rulegroups/{rulegroup_id} | [get rule group](#get-rule-group) | Gets a specified rule group for the given cluster. |
| GET | /api/v2/projects/{project_id}/clusters/{cluster_id}/rulegroups | [list rule groups](#list-rule-groups) |  |
| PUT | /api/v2/projects/{project_id}/clusters/{cluster_id}/rulegroups/{rulegroup_id} | [update rule group](#update-rule-group) | Updates the specified rule group for the given cluster. |
  


###  seed

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| GET | /api/v2/seeds/{seed_name}/settings | [get seed settings](#get-seed-settings) | Gets the seed settings. |
| GET | /api/v1/seed | [list seed names](#list-seed-names) |  |
  


###  serviceaccounts

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| POST | /api/v1/projects/{project_id}/serviceaccounts | [add service account to project](#add-service-account-to-project) |  |
| DELETE | /api/v1/projects/{project_id}/serviceaccounts/{serviceaccount_id} | [delete service account](#delete-service-account) |  |
| GET | /api/v1/projects/{project_id}/serviceaccounts | [list service accounts](#list-service-accounts) |  |
| PUT | /api/v1/projects/{project_id}/serviceaccounts/{serviceaccount_id} | [update service account](#update-service-account) |  |
  


###  settings

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| GET | /api/v1/me/settings | [get current user settings](#get-current-user-settings) | Returns settings of the current user. |
| PATCH | /api/v1/me/settings | [patch current user settings](#patch-current-user-settings) | Updates settings of the current user. |
  


###  tokens

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| POST | /api/v1/projects/{project_id}/serviceaccounts/{serviceaccount_id}/tokens | [add token to service account](#add-token-to-service-account) |  |
| DELETE | /api/v1/projects/{project_id}/serviceaccounts/{serviceaccount_id}/tokens/{token_id} | [delete service account token](#delete-service-account-token) |  |
| GET | /api/v1/projects/{project_id}/serviceaccounts/{serviceaccount_id}/tokens | [list service account tokens](#list-service-account-tokens) |  |
| PATCH | /api/v1/projects/{project_id}/serviceaccounts/{serviceaccount_id}/tokens/{token_id} | [patch service account token](#patch-service-account-token) |  |
| PUT | /api/v1/projects/{project_id}/serviceaccounts/{serviceaccount_id}/tokens/{token_id} | [update service account token](#update-service-account-token) |  |
  


###  users

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| POST | /api/v1/projects/{project_id}/users | [add user to project](#add-user-to-project) |  |
| DELETE | /api/v1/projects/{project_id}/users/{user_id} | [delete user from project](#delete-user-from-project) |  |
| PUT | /api/v1/projects/{project_id}/users/{user_id} | [edit user in project](#edit-user-in-project) |  |
| GET | /api/v1/me | [get current user](#get-current-user) | Returns information about the current user. |
| GET | /api/v1/projects/{project_id}/users | [get users for project](#get-users-for-project) |  |
| POST | /api/v1/me/logout | [logout current user](#logout-current-user) | Adds current authorization bearer token to the blacklist. |
  


###  version

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| GET | /api/v2/providers/{provider_name}/versions | [list versions by provider](#list-versions-by-provider) |  |
  


###  versions

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| GET | /api/v1/upgrades/cluster | [get master versions](#get-master-versions) |  |
| GET | /api/v1/version | [get meta kube versions](#get-meta-kube-versions) | Get versions of running MetaKube components. |
| GET | /api/v1/upgrades/node | [get node upgrades](#get-node-upgrades) |  |
  


###  vsphere

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| GET | /api/v2/providers/vsphere/datastores | [list v sphere datastores](#list-v-sphere-datastores) |  |
| GET | /api/v1/providers/vsphere/folders | [list v sphere folders](#list-v-sphere-folders) |  |
| GET | /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/providers/vsphere/folders | [list v sphere folders no credentials](#list-v-sphere-folders-no-credentials) |  |
| GET | /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/vsphere/folders | [list v sphere folders no credentials v2](#list-v-sphere-folders-no-credentials-v2) |  |
| GET | /api/v1/providers/vsphere/networks | [list v sphere networks](#list-v-sphere-networks) |  |
| GET | /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/providers/vsphere/networks | [list v sphere networks no credentials](#list-v-sphere-networks-no-credentials) |  |
| GET | /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/vsphere/networks | [list v sphere networks no credentials v2](#list-v-sphere-networks-no-credentials-v2) |  |
  


## Paths

### <span id="add-service-account-to-project"></span> add service account to project (*addServiceAccountToProject*)

```
POST /api/v1/projects/{project_id}/serviceaccounts
```

Adds the given service account to the given project

#### Consumes
  * application/json

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| project_id | `path` | string | `string` |  | ✓ |  |  |
| Body | `body` | [ServiceAccount](#service-account) | `models.ServiceAccount` | |  | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [201](#add-service-account-to-project-201) | Created | ServiceAccount |  | [schema](#add-service-account-to-project-201-schema) |
| [401](#add-service-account-to-project-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#add-service-account-to-project-401-schema) |
| [403](#add-service-account-to-project-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#add-service-account-to-project-403-schema) |
| [default](#add-service-account-to-project-default) | | errorResponse |  | [schema](#add-service-account-to-project-default-schema) |

#### Responses


##### <span id="add-service-account-to-project-201"></span> 201 - ServiceAccount
Status: Created

###### <span id="add-service-account-to-project-201-schema"></span> Schema
   
  

[ServiceAccount](#service-account)

##### <span id="add-service-account-to-project-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="add-service-account-to-project-401-schema"></span> Schema

##### <span id="add-service-account-to-project-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="add-service-account-to-project-403-schema"></span> Schema

##### <span id="add-service-account-to-project-default"></span> Default Response
errorResponse

###### <span id="add-service-account-to-project-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="add-token-to-service-account"></span> add token to service account (*addTokenToServiceAccount*)

```
POST /api/v1/projects/{project_id}/serviceaccounts/{serviceaccount_id}/tokens
```

Generates a token for the given service account

#### Consumes
  * application/json

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| project_id | `path` | string | `string` |  | ✓ |  |  |
| serviceaccount_id | `path` | string | `string` |  | ✓ |  |  |
| Body | `body` | [ServiceAccountToken](#service-account-token) | `models.ServiceAccountToken` | |  | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [201](#add-token-to-service-account-201) | Created | ServiceAccountToken |  | [schema](#add-token-to-service-account-201-schema) |
| [401](#add-token-to-service-account-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#add-token-to-service-account-401-schema) |
| [403](#add-token-to-service-account-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#add-token-to-service-account-403-schema) |
| [default](#add-token-to-service-account-default) | | errorResponse |  | [schema](#add-token-to-service-account-default-schema) |

#### Responses


##### <span id="add-token-to-service-account-201"></span> 201 - ServiceAccountToken
Status: Created

###### <span id="add-token-to-service-account-201-schema"></span> Schema
   
  

[ServiceAccountToken](#service-account-token)

##### <span id="add-token-to-service-account-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="add-token-to-service-account-401-schema"></span> Schema

##### <span id="add-token-to-service-account-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="add-token-to-service-account-403-schema"></span> Schema

##### <span id="add-token-to-service-account-default"></span> Default Response
errorResponse

###### <span id="add-token-to-service-account-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="add-user-to-project"></span> add user to project (*addUserToProject*)

```
POST /api/v1/projects/{project_id}/users
```

Adds the given user to the given project

#### Consumes
  * application/json

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| project_id | `path` | string | `string` |  | ✓ |  |  |
| Body | `body` | [User](#user) | `models.User` | |  | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [201](#add-user-to-project-201) | Created | User |  | [schema](#add-user-to-project-201-schema) |
| [401](#add-user-to-project-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#add-user-to-project-401-schema) |
| [403](#add-user-to-project-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#add-user-to-project-403-schema) |
| [default](#add-user-to-project-default) | | errorResponse |  | [schema](#add-user-to-project-default-schema) |

#### Responses


##### <span id="add-user-to-project-201"></span> 201 - User
Status: Created

###### <span id="add-user-to-project-201-schema"></span> Schema
   
  

[User](#user)

##### <span id="add-user-to-project-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="add-user-to-project-401-schema"></span> Schema

##### <span id="add-user-to-project-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="add-user-to-project-403-schema"></span> Schema

##### <span id="add-user-to-project-default"></span> Default Response
errorResponse

###### <span id="add-user-to-project-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="assign-ssh-key-to-cluster"></span> assign SSH key to cluster (*assignSSHKeyToCluster*)

```
PUT /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/sshkeys/{key_id}
```

Assigns an existing ssh key to the given cluster

#### Consumes
  * application/json

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| dc | `path` | string | `string` |  | ✓ |  |  |
| key_id | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [201](#assign-ssh-key-to-cluster-201) | Created | SSHKey |  | [schema](#assign-ssh-key-to-cluster-201-schema) |
| [401](#assign-ssh-key-to-cluster-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#assign-ssh-key-to-cluster-401-schema) |
| [403](#assign-ssh-key-to-cluster-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#assign-ssh-key-to-cluster-403-schema) |
| [default](#assign-ssh-key-to-cluster-default) | | errorResponse |  | [schema](#assign-ssh-key-to-cluster-default-schema) |

#### Responses


##### <span id="assign-ssh-key-to-cluster-201"></span> 201 - SSHKey
Status: Created

###### <span id="assign-ssh-key-to-cluster-201-schema"></span> Schema
   
  

[SSHKey](#ssh-key)

##### <span id="assign-ssh-key-to-cluster-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="assign-ssh-key-to-cluster-401-schema"></span> Schema

##### <span id="assign-ssh-key-to-cluster-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="assign-ssh-key-to-cluster-403-schema"></span> Schema

##### <span id="assign-ssh-key-to-cluster-default"></span> Default Response
errorResponse

###### <span id="assign-ssh-key-to-cluster-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="assign-ssh-key-to-cluster-v2"></span> assign SSH key to cluster v2 (*assignSSHKeyToClusterV2*)

```
PUT /api/v2/projects/{project_id}/clusters/{cluster_id}/sshkeys/{key_id}
```

Assigns an existing ssh key to the given cluster

#### Consumes
  * application/json

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| key_id | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [201](#assign-ssh-key-to-cluster-v2-201) | Created | SSHKey |  | [schema](#assign-ssh-key-to-cluster-v2-201-schema) |
| [401](#assign-ssh-key-to-cluster-v2-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#assign-ssh-key-to-cluster-v2-401-schema) |
| [403](#assign-ssh-key-to-cluster-v2-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#assign-ssh-key-to-cluster-v2-403-schema) |
| [default](#assign-ssh-key-to-cluster-v2-default) | | errorResponse |  | [schema](#assign-ssh-key-to-cluster-v2-default-schema) |

#### Responses


##### <span id="assign-ssh-key-to-cluster-v2-201"></span> 201 - SSHKey
Status: Created

###### <span id="assign-ssh-key-to-cluster-v2-201-schema"></span> Schema
   
  

[SSHKey](#ssh-key)

##### <span id="assign-ssh-key-to-cluster-v2-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="assign-ssh-key-to-cluster-v2-401-schema"></span> Schema

##### <span id="assign-ssh-key-to-cluster-v2-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="assign-ssh-key-to-cluster-v2-403-schema"></span> Schema

##### <span id="assign-ssh-key-to-cluster-v2-default"></span> Default Response
errorResponse

###### <span id="assign-ssh-key-to-cluster-v2-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="bind-user-to-cluster-role"></span> bind user to cluster role (*bindUserToClusterRole*)

```
POST /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/clusterroles/{role_id}/clusterbindings
```

Binds user to cluster role

#### Consumes
  * application/json

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| dc | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |
| role_id | `path` | string | `string` |  | ✓ |  |  |
| Body | `body` | [ClusterRoleUser](#cluster-role-user) | `models.ClusterRoleUser` | |  | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#bind-user-to-cluster-role-200) | OK | ClusterRoleBinding |  | [schema](#bind-user-to-cluster-role-200-schema) |
| [401](#bind-user-to-cluster-role-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#bind-user-to-cluster-role-401-schema) |
| [403](#bind-user-to-cluster-role-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#bind-user-to-cluster-role-403-schema) |
| [default](#bind-user-to-cluster-role-default) | | errorResponse |  | [schema](#bind-user-to-cluster-role-default-schema) |

#### Responses


##### <span id="bind-user-to-cluster-role-200"></span> 200 - ClusterRoleBinding
Status: OK

###### <span id="bind-user-to-cluster-role-200-schema"></span> Schema
   
  

[ClusterRoleBinding](#cluster-role-binding)

##### <span id="bind-user-to-cluster-role-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="bind-user-to-cluster-role-401-schema"></span> Schema

##### <span id="bind-user-to-cluster-role-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="bind-user-to-cluster-role-403-schema"></span> Schema

##### <span id="bind-user-to-cluster-role-default"></span> Default Response
errorResponse

###### <span id="bind-user-to-cluster-role-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="bind-user-to-cluster-role-v2"></span> bind user to cluster role v2 (*bindUserToClusterRoleV2*)

```
POST /api/v2/projects/{project_id}/clusters/{cluster_id}/clusterroles/{role_id}/clusterbindings
```

Binds user to cluster role

#### Consumes
  * application/json

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |
| role_id | `path` | string | `string` |  | ✓ |  |  |
| Body | `body` | [ClusterRoleUser](#cluster-role-user) | `models.ClusterRoleUser` | |  | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#bind-user-to-cluster-role-v2-200) | OK | ClusterRoleBinding |  | [schema](#bind-user-to-cluster-role-v2-200-schema) |
| [401](#bind-user-to-cluster-role-v2-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#bind-user-to-cluster-role-v2-401-schema) |
| [403](#bind-user-to-cluster-role-v2-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#bind-user-to-cluster-role-v2-403-schema) |
| [default](#bind-user-to-cluster-role-v2-default) | | errorResponse |  | [schema](#bind-user-to-cluster-role-v2-default-schema) |

#### Responses


##### <span id="bind-user-to-cluster-role-v2-200"></span> 200 - ClusterRoleBinding
Status: OK

###### <span id="bind-user-to-cluster-role-v2-200-schema"></span> Schema
   
  

[ClusterRoleBinding](#cluster-role-binding)

##### <span id="bind-user-to-cluster-role-v2-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="bind-user-to-cluster-role-v2-401-schema"></span> Schema

##### <span id="bind-user-to-cluster-role-v2-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="bind-user-to-cluster-role-v2-403-schema"></span> Schema

##### <span id="bind-user-to-cluster-role-v2-default"></span> Default Response
errorResponse

###### <span id="bind-user-to-cluster-role-v2-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="bind-user-to-role"></span> bind user to role (*bindUserToRole*)

```
POST /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/roles/{namespace}/{role_id}/bindings
```

Binds user to the role

#### Consumes
  * application/json

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| dc | `path` | string | `string` |  | ✓ |  |  |
| namespace | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |
| role_id | `path` | string | `string` |  | ✓ |  |  |
| Body | `body` | [RoleUser](#role-user) | `models.RoleUser` | |  | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#bind-user-to-role-200) | OK | RoleBinding |  | [schema](#bind-user-to-role-200-schema) |
| [401](#bind-user-to-role-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#bind-user-to-role-401-schema) |
| [403](#bind-user-to-role-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#bind-user-to-role-403-schema) |
| [default](#bind-user-to-role-default) | | errorResponse |  | [schema](#bind-user-to-role-default-schema) |

#### Responses


##### <span id="bind-user-to-role-200"></span> 200 - RoleBinding
Status: OK

###### <span id="bind-user-to-role-200-schema"></span> Schema
   
  

[RoleBinding](#role-binding)

##### <span id="bind-user-to-role-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="bind-user-to-role-401-schema"></span> Schema

##### <span id="bind-user-to-role-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="bind-user-to-role-403-schema"></span> Schema

##### <span id="bind-user-to-role-default"></span> Default Response
errorResponse

###### <span id="bind-user-to-role-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="bind-user-to-role-v2"></span> bind user to role v2 (*bindUserToRoleV2*)

```
POST /api/v2/projects/{project_id}/clusters/{cluster_id}/roles/{namespace}/{role_id}/bindings
```

Binds user to the role

#### Consumes
  * application/json

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| namespace | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |
| role_id | `path` | string | `string` |  | ✓ |  |  |
| Body | `body` | [RoleUser](#role-user) | `models.RoleUser` | |  | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#bind-user-to-role-v2-200) | OK | RoleBinding |  | [schema](#bind-user-to-role-v2-200-schema) |
| [401](#bind-user-to-role-v2-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#bind-user-to-role-v2-401-schema) |
| [403](#bind-user-to-role-v2-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#bind-user-to-role-v2-403-schema) |
| [default](#bind-user-to-role-v2-default) | | errorResponse |  | [schema](#bind-user-to-role-v2-default-schema) |

#### Responses


##### <span id="bind-user-to-role-v2-200"></span> 200 - RoleBinding
Status: OK

###### <span id="bind-user-to-role-v2-200-schema"></span> Schema
   
  

[RoleBinding](#role-binding)

##### <span id="bind-user-to-role-v2-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="bind-user-to-role-v2-401-schema"></span> Schema

##### <span id="bind-user-to-role-v2-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="bind-user-to-role-v2-403-schema"></span> Schema

##### <span id="bind-user-to-role-v2-default"></span> Default Response
errorResponse

###### <span id="bind-user-to-role-v2-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="create-addon"></span> create addon (*createAddon*)

```
POST /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/addons
```

Creates an addon that will belong to the given cluster

#### Consumes
  * application/json

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| dc | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |
| Body | `body` | [Addon](#addon) | `models.Addon` | |  | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [201](#create-addon-201) | Created | Addon |  | [schema](#create-addon-201-schema) |
| [401](#create-addon-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#create-addon-401-schema) |
| [403](#create-addon-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#create-addon-403-schema) |
| [default](#create-addon-default) | | errorResponse |  | [schema](#create-addon-default-schema) |

#### Responses


##### <span id="create-addon-201"></span> 201 - Addon
Status: Created

###### <span id="create-addon-201-schema"></span> Schema
   
  

[Addon](#addon)

##### <span id="create-addon-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="create-addon-401-schema"></span> Schema

##### <span id="create-addon-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="create-addon-403-schema"></span> Schema

##### <span id="create-addon-default"></span> Default Response
errorResponse

###### <span id="create-addon-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="create-addon-v2"></span> create addon v2 (*createAddonV2*)

```
POST /api/v2/projects/{project_id}/clusters/{cluster_id}/addons
```

Creates an addon that will belong to the given cluster

#### Consumes
  * application/json

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |
| Body | `body` | [Addon](#addon) | `models.Addon` | |  | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [201](#create-addon-v2-201) | Created | Addon |  | [schema](#create-addon-v2-201-schema) |
| [401](#create-addon-v2-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#create-addon-v2-401-schema) |
| [403](#create-addon-v2-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#create-addon-v2-403-schema) |
| [default](#create-addon-v2-default) | | errorResponse |  | [schema](#create-addon-v2-default-schema) |

#### Responses


##### <span id="create-addon-v2-201"></span> 201 - Addon
Status: Created

###### <span id="create-addon-v2-201-schema"></span> Schema
   
  

[Addon](#addon)

##### <span id="create-addon-v2-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="create-addon-v2-401-schema"></span> Schema

##### <span id="create-addon-v2-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="create-addon-v2-403-schema"></span> Schema

##### <span id="create-addon-v2-default"></span> Default Response
errorResponse

###### <span id="create-addon-v2-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="create-allowed-registry"></span> create allowed registry (*createAllowedRegistry*)

```
POST /api/v2/allowedregistries
```

Creates a allowed registry

#### Consumes
  * application/json

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| Body | `body` | [WrBody](#wr-body) | `models.WrBody` | |  | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [201](#create-allowed-registry-201) | Created | AllowedRegistry |  | [schema](#create-allowed-registry-201-schema) |
| [401](#create-allowed-registry-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#create-allowed-registry-401-schema) |
| [403](#create-allowed-registry-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#create-allowed-registry-403-schema) |
| [default](#create-allowed-registry-default) | | errorResponse |  | [schema](#create-allowed-registry-default-schema) |

#### Responses


##### <span id="create-allowed-registry-201"></span> 201 - AllowedRegistry
Status: Created

###### <span id="create-allowed-registry-201-schema"></span> Schema
   
  

[AllowedRegistry](#allowed-registry)

##### <span id="create-allowed-registry-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="create-allowed-registry-401-schema"></span> Schema

##### <span id="create-allowed-registry-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="create-allowed-registry-403-schema"></span> Schema

##### <span id="create-allowed-registry-default"></span> Default Response
errorResponse

###### <span id="create-allowed-registry-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="create-cluster"></span> Creates a cluster for the given project. (*createCluster*)

```
POST /api/v1/projects/{project_id}/dc/{dc}/clusters
```

#### Consumes
  * application/json

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| dc | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |
| Body | `body` | [CreateClusterSpec](#create-cluster-spec) | `models.CreateClusterSpec` | |  | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [201](#create-cluster-201) | Created | Cluster |  | [schema](#create-cluster-201-schema) |
| [401](#create-cluster-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#create-cluster-401-schema) |
| [403](#create-cluster-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#create-cluster-403-schema) |
| [default](#create-cluster-default) | | errorResponse |  | [schema](#create-cluster-default-schema) |

#### Responses


##### <span id="create-cluster-201"></span> 201 - Cluster
Status: Created

###### <span id="create-cluster-201-schema"></span> Schema
   
  

[Cluster](#cluster)

##### <span id="create-cluster-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="create-cluster-401-schema"></span> Schema

##### <span id="create-cluster-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="create-cluster-403-schema"></span> Schema

##### <span id="create-cluster-default"></span> Default Response
errorResponse

###### <span id="create-cluster-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="create-cluster-role"></span> create cluster role (*createClusterRole*)

```
POST /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/clusterroles
```

Creates cluster role

#### Consumes
  * application/json

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| dc | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |
| Body | `body` | [ClusterRole](#cluster-role) | `models.ClusterRole` | |  | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [201](#create-cluster-role-201) | Created | ClusterRole |  | [schema](#create-cluster-role-201-schema) |
| [401](#create-cluster-role-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#create-cluster-role-401-schema) |
| [403](#create-cluster-role-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#create-cluster-role-403-schema) |
| [default](#create-cluster-role-default) | | errorResponse |  | [schema](#create-cluster-role-default-schema) |

#### Responses


##### <span id="create-cluster-role-201"></span> 201 - ClusterRole
Status: Created

###### <span id="create-cluster-role-201-schema"></span> Schema
   
  

[ClusterRole](#cluster-role)

##### <span id="create-cluster-role-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="create-cluster-role-401-schema"></span> Schema

##### <span id="create-cluster-role-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="create-cluster-role-403-schema"></span> Schema

##### <span id="create-cluster-role-default"></span> Default Response
errorResponse

###### <span id="create-cluster-role-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="create-cluster-template"></span> Creates a cluster templates for the given project. (*createClusterTemplate*)

```
POST /api/v2/projects/{project_id}/clustertemplates
```

#### Consumes
  * application/json

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| project_id | `path` | string | `string` |  | ✓ |  |  |
| Body | `body` | [CreateClusterTemplateBody](#create-cluster-template-body) | `CreateClusterTemplateBody` | |  | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [201](#create-cluster-template-201) | Created | ClusterTemplate |  | [schema](#create-cluster-template-201-schema) |
| [401](#create-cluster-template-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#create-cluster-template-401-schema) |
| [403](#create-cluster-template-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#create-cluster-template-403-schema) |
| [default](#create-cluster-template-default) | | errorResponse |  | [schema](#create-cluster-template-default-schema) |

#### Responses


##### <span id="create-cluster-template-201"></span> 201 - ClusterTemplate
Status: Created

###### <span id="create-cluster-template-201-schema"></span> Schema
   
  

[ClusterTemplate](#cluster-template)

##### <span id="create-cluster-template-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="create-cluster-template-401-schema"></span> Schema

##### <span id="create-cluster-template-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="create-cluster-template-403-schema"></span> Schema

##### <span id="create-cluster-template-default"></span> Default Response
errorResponse

###### <span id="create-cluster-template-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

###### Inlined models

**<span id="create-cluster-template-body"></span> CreateClusterTemplateBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| DNSDomain | string| `string` |  | |  |  |
| Name | string| `string` |  | |  |  |
| PodsCIDR | string| `string` |  | |  |  |
| Scope | string| `string` |  | |  |  |
| ServicesCIDR | string| `string` |  | |  |  |
| UserSSHKeys | [][ClusterTemplateSSHKey](#cluster-template-ssh-key)| `[]*models.ClusterTemplateSSHKey` |  | |  |  |
| cluster | [Cluster](#cluster)| `models.Cluster` |  | |  |  |
| nodeDeployment | [NodeDeployment](#node-deployment)| `models.NodeDeployment` |  | |  |  |



### <span id="create-cluster-template-instance"></span> Create cluster template instance. (*createClusterTemplateInstance*)

```
POST /api/v2/projects/{project_id}/clustertemplates/{template_id}/instances
```

#### Consumes
  * application/json

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| project_id | `path` | string | `string` |  | ✓ |  |  |
| template_id | `path` | string | `string` |  | ✓ |  |  |
| Body | `body` | [CreateClusterTemplateInstanceBody](#create-cluster-template-instance-body) | `CreateClusterTemplateInstanceBody` | |  | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [201](#create-cluster-template-instance-201) | Created | ClusterTemplateInstance |  | [schema](#create-cluster-template-instance-201-schema) |
| [401](#create-cluster-template-instance-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#create-cluster-template-instance-401-schema) |
| [403](#create-cluster-template-instance-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#create-cluster-template-instance-403-schema) |
| [default](#create-cluster-template-instance-default) | | errorResponse |  | [schema](#create-cluster-template-instance-default-schema) |

#### Responses


##### <span id="create-cluster-template-instance-201"></span> 201 - ClusterTemplateInstance
Status: Created

###### <span id="create-cluster-template-instance-201-schema"></span> Schema
   
  

[ClusterTemplateInstance](#cluster-template-instance)

##### <span id="create-cluster-template-instance-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="create-cluster-template-instance-401-schema"></span> Schema

##### <span id="create-cluster-template-instance-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="create-cluster-template-instance-403-schema"></span> Schema

##### <span id="create-cluster-template-instance-default"></span> Default Response
errorResponse

###### <span id="create-cluster-template-instance-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

###### Inlined models

**<span id="create-cluster-template-instance-body"></span> CreateClusterTemplateInstanceBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Replicas | int64 (formatted integer)| `int64` |  | |  |  |



### <span id="create-cluster-v2"></span> Creates a cluster for the given project. (*createClusterV2*)

```
POST /api/v2/projects/{project_id}/clusters
```

#### Consumes
  * application/json

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| project_id | `path` | string | `string` |  | ✓ |  |  |
| Body | `body` | [CreateClusterSpec](#create-cluster-spec) | `models.CreateClusterSpec` | |  | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [201](#create-cluster-v2-201) | Created | Cluster |  | [schema](#create-cluster-v2-201-schema) |
| [401](#create-cluster-v2-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#create-cluster-v2-401-schema) |
| [403](#create-cluster-v2-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#create-cluster-v2-403-schema) |
| [default](#create-cluster-v2-default) | | errorResponse |  | [schema](#create-cluster-v2-default-schema) |

#### Responses


##### <span id="create-cluster-v2-201"></span> 201 - Cluster
Status: Created

###### <span id="create-cluster-v2-201-schema"></span> Schema
   
  

[Cluster](#cluster)

##### <span id="create-cluster-v2-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="create-cluster-v2-401-schema"></span> Schema

##### <span id="create-cluster-v2-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="create-cluster-v2-403-schema"></span> Schema

##### <span id="create-cluster-v2-default"></span> Default Response
errorResponse

###### <span id="create-cluster-v2-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="create-constraint"></span> Creates a given constraint for the specified cluster. (*createConstraint*)

```
POST /api/v2/projects/{project_id}/clusters/{cluster_id}/constraints
```

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |
| Body | `body` | [ConstraintBody](#constraint-body) | `models.ConstraintBody` | | ✓ | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#create-constraint-200) | OK | Constraint |  | [schema](#create-constraint-200-schema) |
| [401](#create-constraint-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#create-constraint-401-schema) |
| [403](#create-constraint-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#create-constraint-403-schema) |
| [default](#create-constraint-default) | | errorResponse |  | [schema](#create-constraint-default-schema) |

#### Responses


##### <span id="create-constraint-200"></span> 200 - Constraint
Status: OK

###### <span id="create-constraint-200-schema"></span> Schema
   
  

[Constraint](#constraint)

##### <span id="create-constraint-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="create-constraint-401-schema"></span> Schema

##### <span id="create-constraint-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="create-constraint-403-schema"></span> Schema

##### <span id="create-constraint-default"></span> Default Response
errorResponse

###### <span id="create-constraint-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="create-constraint-template"></span> create constraint template (*createConstraintTemplate*)

```
POST /api/v2/constrainttemplates
```

Create constraint template

#### Consumes
  * application/json

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| Body | `body` | [CtBody](#ct-body) | `models.CtBody` | |  | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#create-constraint-template-200) | OK | ConstraintTemplate |  | [schema](#create-constraint-template-200-schema) |
| [401](#create-constraint-template-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#create-constraint-template-401-schema) |
| [403](#create-constraint-template-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#create-constraint-template-403-schema) |
| [default](#create-constraint-template-default) | | errorResponse |  | [schema](#create-constraint-template-default-schema) |

#### Responses


##### <span id="create-constraint-template-200"></span> 200 - ConstraintTemplate
Status: OK

###### <span id="create-constraint-template-200-schema"></span> Schema
   
  

[ConstraintTemplate](#constraint-template)

##### <span id="create-constraint-template-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="create-constraint-template-401-schema"></span> Schema

##### <span id="create-constraint-template-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="create-constraint-template-403-schema"></span> Schema

##### <span id="create-constraint-template-default"></span> Default Response
errorResponse

###### <span id="create-constraint-template-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="create-d-c"></span> Create the datacenter for a specified seed. (*createDC*)

```
POST /api/v1/seed/{seed_name}/dc
```

#### Consumes
  * application/json

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| seed_name | `path` | string | `string` |  | ✓ |  |  |
| Body | `body` | [CreateDCBody](#create-d-c-body) | `CreateDCBody` | |  | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [201](#create-d-c-201) | Created | Datacenter |  | [schema](#create-d-c-201-schema) |
| [401](#create-d-c-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#create-d-c-401-schema) |
| [403](#create-d-c-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#create-d-c-403-schema) |
| [default](#create-d-c-default) | | errorResponse |  | [schema](#create-d-c-default-schema) |

#### Responses


##### <span id="create-d-c-201"></span> 201 - Datacenter
Status: Created

###### <span id="create-d-c-201-schema"></span> Schema
   
  

[Datacenter](#datacenter)

##### <span id="create-d-c-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="create-d-c-401-schema"></span> Schema

##### <span id="create-d-c-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="create-d-c-403-schema"></span> Schema

##### <span id="create-d-c-default"></span> Default Response
errorResponse

###### <span id="create-d-c-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

###### Inlined models

**<span id="create-d-c-body"></span> CreateDCBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Name | string| `string` |  | |  |  |
| spec | [DatacenterSpec](#datacenter-spec)| `models.DatacenterSpec` |  | |  |  |



### <span id="create-default-constraint"></span> create default constraint (*createDefaultConstraint*)

```
POST /api/v2/constraints
```

Creates default constraint

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| Body | `body` | [ConstraintBody](#constraint-body) | `models.ConstraintBody` | | ✓ | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#create-default-constraint-200) | OK | Constraint |  | [schema](#create-default-constraint-200-schema) |
| [401](#create-default-constraint-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#create-default-constraint-401-schema) |
| [403](#create-default-constraint-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#create-default-constraint-403-schema) |
| [default](#create-default-constraint-default) | | errorResponse |  | [schema](#create-default-constraint-default-schema) |

#### Responses


##### <span id="create-default-constraint-200"></span> 200 - Constraint
Status: OK

###### <span id="create-default-constraint-200-schema"></span> Schema
   
  

[Constraint](#constraint)

##### <span id="create-default-constraint-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="create-default-constraint-401-schema"></span> Schema

##### <span id="create-default-constraint-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="create-default-constraint-403-schema"></span> Schema

##### <span id="create-default-constraint-default"></span> Default Response
errorResponse

###### <span id="create-default-constraint-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="create-etcd-backup-config"></span> create etcd backup config (*createEtcdBackupConfig*)

```
POST /api/v2/projects/{project_id}/clusters/{cluster_id}/etcdbackupconfigs
```

Creates a etcd backup config that will belong to the given cluster

#### Consumes
  * application/json

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |
| Body | `body` | [EbcBody](#ebc-body) | `models.EbcBody` | |  | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [201](#create-etcd-backup-config-201) | Created | EtcdBackupConfig |  | [schema](#create-etcd-backup-config-201-schema) |
| [401](#create-etcd-backup-config-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#create-etcd-backup-config-401-schema) |
| [403](#create-etcd-backup-config-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#create-etcd-backup-config-403-schema) |
| [default](#create-etcd-backup-config-default) | | errorResponse |  | [schema](#create-etcd-backup-config-default-schema) |

#### Responses


##### <span id="create-etcd-backup-config-201"></span> 201 - EtcdBackupConfig
Status: Created

###### <span id="create-etcd-backup-config-201-schema"></span> Schema
   
  

[EtcdBackupConfig](#etcd-backup-config)

##### <span id="create-etcd-backup-config-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="create-etcd-backup-config-401-schema"></span> Schema

##### <span id="create-etcd-backup-config-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="create-etcd-backup-config-403-schema"></span> Schema

##### <span id="create-etcd-backup-config-default"></span> Default Response
errorResponse

###### <span id="create-etcd-backup-config-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="create-etcd-restore"></span> create etcd restore (*createEtcdRestore*)

```
POST /api/v2/projects/{project_id}/clusters/{cluster_id}/etcdrestores
```

Creates a etcd backup restore for a given cluster

#### Consumes
  * application/json

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |
| Body | `body` | [ErBody](#er-body) | `models.ErBody` | |  | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [201](#create-etcd-restore-201) | Created | EtcdBackupConfig |  | [schema](#create-etcd-restore-201-schema) |
| [401](#create-etcd-restore-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#create-etcd-restore-401-schema) |
| [403](#create-etcd-restore-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#create-etcd-restore-403-schema) |
| [default](#create-etcd-restore-default) | | errorResponse |  | [schema](#create-etcd-restore-default-schema) |

#### Responses


##### <span id="create-etcd-restore-201"></span> 201 - EtcdBackupConfig
Status: Created

###### <span id="create-etcd-restore-201-schema"></span> Schema
   
  

[EtcdBackupConfig](#etcd-backup-config)

##### <span id="create-etcd-restore-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="create-etcd-restore-401-schema"></span> Schema

##### <span id="create-etcd-restore-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="create-etcd-restore-403-schema"></span> Schema

##### <span id="create-etcd-restore-default"></span> Default Response
errorResponse

###### <span id="create-etcd-restore-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="create-external-cluster"></span> Creates an external cluster for the given project. (*createExternalCluster*)

```
POST /api/v2/projects/{project_id}/kubernetes/clusters
```

#### Consumes
  * application/json

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| project_id | `path` | string | `string` |  | ✓ |  |  |
| Body | `body` | [Body](#body) | `models.Body` | |  | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [201](#create-external-cluster-201) | Created | Cluster |  | [schema](#create-external-cluster-201-schema) |
| [401](#create-external-cluster-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#create-external-cluster-401-schema) |
| [403](#create-external-cluster-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#create-external-cluster-403-schema) |
| [default](#create-external-cluster-default) | | errorResponse |  | [schema](#create-external-cluster-default-schema) |

#### Responses


##### <span id="create-external-cluster-201"></span> 201 - Cluster
Status: Created

###### <span id="create-external-cluster-201-schema"></span> Schema
   
  

[Cluster](#cluster)

##### <span id="create-external-cluster-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="create-external-cluster-401-schema"></span> Schema

##### <span id="create-external-cluster-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="create-external-cluster-403-schema"></span> Schema

##### <span id="create-external-cluster-default"></span> Default Response
errorResponse

###### <span id="create-external-cluster-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="create-gatekeeper-config"></span> create gatekeeper config (*createGatekeeperConfig*)

```
POST /api/v2/projects/{project_id}/clusters/{cluster_id}/gatekeeper/config
```

Creates a gatekeeper config for the given cluster

#### Consumes
  * application/json

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |
| Body | `body` | [GatekeeperConfig](#gatekeeper-config) | `models.GatekeeperConfig` | | ✓ | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [201](#create-gatekeeper-config-201) | Created | GatekeeperConfig |  | [schema](#create-gatekeeper-config-201-schema) |
| [401](#create-gatekeeper-config-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#create-gatekeeper-config-401-schema) |
| [403](#create-gatekeeper-config-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#create-gatekeeper-config-403-schema) |
| [default](#create-gatekeeper-config-default) | | errorResponse |  | [schema](#create-gatekeeper-config-default-schema) |

#### Responses


##### <span id="create-gatekeeper-config-201"></span> 201 - GatekeeperConfig
Status: Created

###### <span id="create-gatekeeper-config-201-schema"></span> Schema
   
  

[GatekeeperConfig](#gatekeeper-config)

##### <span id="create-gatekeeper-config-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="create-gatekeeper-config-401-schema"></span> Schema

##### <span id="create-gatekeeper-config-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="create-gatekeeper-config-403-schema"></span> Schema

##### <span id="create-gatekeeper-config-default"></span> Default Response
errorResponse

###### <span id="create-gatekeeper-config-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="create-m-l-a-admin-setting"></span> create m l a admin setting (*createMLAAdminSetting*)

```
POST /api/v2/projects/{project_id}/clusters/{cluster_id}/mlaadminsetting
```

Creates MLA admin setting that will belong to the given cluster

#### Consumes
  * application/json

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |
| Body | `body` | [MLAAdminSetting](#m-l-a-admin-setting) | `models.MLAAdminSetting` | | ✓ | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [201](#create-m-l-a-admin-setting-201) | Created | MLAAdminSetting |  | [schema](#create-m-l-a-admin-setting-201-schema) |
| [401](#create-m-l-a-admin-setting-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#create-m-l-a-admin-setting-401-schema) |
| [403](#create-m-l-a-admin-setting-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#create-m-l-a-admin-setting-403-schema) |
| [default](#create-m-l-a-admin-setting-default) | | errorResponse |  | [schema](#create-m-l-a-admin-setting-default-schema) |

#### Responses


##### <span id="create-m-l-a-admin-setting-201"></span> 201 - MLAAdminSetting
Status: Created

###### <span id="create-m-l-a-admin-setting-201-schema"></span> Schema
   
  

[MLAAdminSetting](#m-l-a-admin-setting)

##### <span id="create-m-l-a-admin-setting-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="create-m-l-a-admin-setting-401-schema"></span> Schema

##### <span id="create-m-l-a-admin-setting-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="create-m-l-a-admin-setting-403-schema"></span> Schema

##### <span id="create-m-l-a-admin-setting-default"></span> Default Response
errorResponse

###### <span id="create-m-l-a-admin-setting-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="create-machine-deployment"></span> create machine deployment (*createMachineDeployment*)

```
POST /api/v2/projects/{project_id}/clusters/{cluster_id}/machinedeployments
```

Creates a machine deployment that will belong to the given cluster

#### Consumes
  * application/json

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |
| Body | `body` | [NodeDeployment](#node-deployment) | `models.NodeDeployment` | |  | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [201](#create-machine-deployment-201) | Created | NodeDeployment |  | [schema](#create-machine-deployment-201-schema) |
| [401](#create-machine-deployment-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#create-machine-deployment-401-schema) |
| [403](#create-machine-deployment-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#create-machine-deployment-403-schema) |
| [default](#create-machine-deployment-default) | | errorResponse |  | [schema](#create-machine-deployment-default-schema) |

#### Responses


##### <span id="create-machine-deployment-201"></span> 201 - NodeDeployment
Status: Created

###### <span id="create-machine-deployment-201-schema"></span> Schema
   
  

[NodeDeployment](#node-deployment)

##### <span id="create-machine-deployment-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="create-machine-deployment-401-schema"></span> Schema

##### <span id="create-machine-deployment-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="create-machine-deployment-403-schema"></span> Schema

##### <span id="create-machine-deployment-default"></span> Default Response
errorResponse

###### <span id="create-machine-deployment-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="create-node-deployment"></span> create node deployment (*createNodeDeployment*)

```
POST /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/nodedeployments
```

Creates a node deployment that will belong to the given cluster

#### Consumes
  * application/json

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| dc | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |
| Body | `body` | [NodeDeployment](#node-deployment) | `models.NodeDeployment` | |  | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [201](#create-node-deployment-201) | Created | NodeDeployment |  | [schema](#create-node-deployment-201-schema) |
| [401](#create-node-deployment-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#create-node-deployment-401-schema) |
| [403](#create-node-deployment-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#create-node-deployment-403-schema) |
| [default](#create-node-deployment-default) | | errorResponse |  | [schema](#create-node-deployment-default-schema) |

#### Responses


##### <span id="create-node-deployment-201"></span> 201 - NodeDeployment
Status: Created

###### <span id="create-node-deployment-201-schema"></span> Schema
   
  

[NodeDeployment](#node-deployment)

##### <span id="create-node-deployment-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="create-node-deployment-401-schema"></span> Schema

##### <span id="create-node-deployment-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="create-node-deployment-403-schema"></span> Schema

##### <span id="create-node-deployment-default"></span> Default Response
errorResponse

###### <span id="create-node-deployment-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="create-node-deployment-request"></span> create node deployment request (*createNodeDeploymentRequest*)

```
POST /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/ndrequests
```

Creates a NodeDeploymentRequest that will belong to the given cluster

#### Consumes
  * application/json

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| dc | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |
| Body | `body` | [NodeDeploymentRequest](#node-deployment-request) | `models.NodeDeploymentRequest` | |  | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [201](#create-node-deployment-request-201) | Created | NodeDeploymentRequest |  | [schema](#create-node-deployment-request-201-schema) |
| [401](#create-node-deployment-request-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#create-node-deployment-request-401-schema) |
| [403](#create-node-deployment-request-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#create-node-deployment-request-403-schema) |
| [default](#create-node-deployment-request-default) | | errorResponse |  | [schema](#create-node-deployment-request-default-schema) |

#### Responses


##### <span id="create-node-deployment-request-201"></span> 201 - NodeDeploymentRequest
Status: Created

###### <span id="create-node-deployment-request-201-schema"></span> Schema
   
  

[NodeDeploymentRequest](#node-deployment-request)

##### <span id="create-node-deployment-request-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="create-node-deployment-request-401-schema"></span> Schema

##### <span id="create-node-deployment-request-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="create-node-deployment-request-403-schema"></span> Schema

##### <span id="create-node-deployment-request-default"></span> Default Response
errorResponse

###### <span id="create-node-deployment-request-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="create-o-id-c-kubeconfig"></span> create o ID c kubeconfig (*createOIDCKubeconfig*)

```
GET /api/v1/kubeconfig
```

Starts OIDC flow and generates kubeconfig, the generated config
contains OIDC provider authentication info

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `query` | string | `string` |  |  |  |  |
| project_id | `query` | string | `string` |  |  |  |  |
| user_id | `query` | string | `string` |  |  |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#create-o-id-c-kubeconfig-200) | OK | Kubeconfig is a clusters kubeconfig |  | [schema](#create-o-id-c-kubeconfig-200-schema) |
| [default](#create-o-id-c-kubeconfig-default) | | errorResponse |  | [schema](#create-o-id-c-kubeconfig-default-schema) |

#### Responses


##### <span id="create-o-id-c-kubeconfig-200"></span> 200 - Kubeconfig is a clusters kubeconfig
Status: OK

###### <span id="create-o-id-c-kubeconfig-200-schema"></span> Schema
   
  

[]uint8 (formatted integer)

##### <span id="create-o-id-c-kubeconfig-default"></span> Default Response
errorResponse

###### <span id="create-o-id-c-kubeconfig-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="create-or-update-backup-credentials"></span> create or update backup credentials (*createOrUpdateBackupCredentials*)

```
PUT /api/v2/seeds/{seed_name}/backupcredentials
```

Creates or updates backup credentials for a given seed

#### Consumes
  * application/json

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| seed_name | `path` | string | `string` |  | ✓ |  |  |
| Body | `body` | [BcBody](#bc-body) | `models.BcBody` | |  | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#create-or-update-backup-credentials-200) | OK | EmptyResponse is a empty response |  | [schema](#create-or-update-backup-credentials-200-schema) |
| [401](#create-or-update-backup-credentials-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#create-or-update-backup-credentials-401-schema) |
| [403](#create-or-update-backup-credentials-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#create-or-update-backup-credentials-403-schema) |
| [default](#create-or-update-backup-credentials-default) | | errorResponse |  | [schema](#create-or-update-backup-credentials-default-schema) |

#### Responses


##### <span id="create-or-update-backup-credentials-200"></span> 200 - EmptyResponse is a empty response
Status: OK

###### <span id="create-or-update-backup-credentials-200-schema"></span> Schema

##### <span id="create-or-update-backup-credentials-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="create-or-update-backup-credentials-401-schema"></span> Schema

##### <span id="create-or-update-backup-credentials-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="create-or-update-backup-credentials-403-schema"></span> Schema

##### <span id="create-or-update-backup-credentials-default"></span> Default Response
errorResponse

###### <span id="create-or-update-backup-credentials-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="create-or-update-metering-configurations"></span> create or update metering configurations (*createOrUpdateMeteringConfigurations*)

```
PUT /api/v1/admin/metering/configurations
```

Configures KKP metering tool. Only available in Kubermatic Enterprise Edition

#### Produces
  * application/json

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#create-or-update-metering-configurations-200) | OK | EmptyResponse is a empty response |  | [schema](#create-or-update-metering-configurations-200-schema) |
| [401](#create-or-update-metering-configurations-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#create-or-update-metering-configurations-401-schema) |
| [403](#create-or-update-metering-configurations-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#create-or-update-metering-configurations-403-schema) |
| [default](#create-or-update-metering-configurations-default) | | errorResponse |  | [schema](#create-or-update-metering-configurations-default-schema) |

#### Responses


##### <span id="create-or-update-metering-configurations-200"></span> 200 - EmptyResponse is a empty response
Status: OK

###### <span id="create-or-update-metering-configurations-200-schema"></span> Schema

##### <span id="create-or-update-metering-configurations-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="create-or-update-metering-configurations-401-schema"></span> Schema

##### <span id="create-or-update-metering-configurations-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="create-or-update-metering-configurations-403-schema"></span> Schema

##### <span id="create-or-update-metering-configurations-default"></span> Default Response
errorResponse

###### <span id="create-or-update-metering-configurations-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="create-or-update-metering-credentials"></span> create or update metering credentials (*createOrUpdateMeteringCredentials*)

```
PUT /api/v1/admin/metering/credentials
```

Creates or updates the metering tool credentials. Only available in Kubermatic Enterprise Edition

#### Produces
  * application/json

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#create-or-update-metering-credentials-200) | OK | EmptyResponse is a empty response |  | [schema](#create-or-update-metering-credentials-200-schema) |
| [401](#create-or-update-metering-credentials-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#create-or-update-metering-credentials-401-schema) |
| [403](#create-or-update-metering-credentials-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#create-or-update-metering-credentials-403-schema) |
| [default](#create-or-update-metering-credentials-default) | | errorResponse |  | [schema](#create-or-update-metering-credentials-default-schema) |

#### Responses


##### <span id="create-or-update-metering-credentials-200"></span> 200 - EmptyResponse is a empty response
Status: OK

###### <span id="create-or-update-metering-credentials-200-schema"></span> Schema

##### <span id="create-or-update-metering-credentials-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="create-or-update-metering-credentials-401-schema"></span> Schema

##### <span id="create-or-update-metering-credentials-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="create-or-update-metering-credentials-403-schema"></span> Schema

##### <span id="create-or-update-metering-credentials-default"></span> Default Response
errorResponse

###### <span id="create-or-update-metering-credentials-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="create-preset"></span> create preset (*createPreset*)

```
POST /api/v2/providers/{provider_name}/presets
```

Creates the preset

#### Consumes
  * application/json

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| provider_name | `path` | string | `string` |  | ✓ |  |  |
| Body | `body` | [Preset](#preset) | `models.Preset` | | ✓ | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#create-preset-200) | OK | Preset |  | [schema](#create-preset-200-schema) |
| [401](#create-preset-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#create-preset-401-schema) |
| [403](#create-preset-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#create-preset-403-schema) |
| [default](#create-preset-default) | | errorResponse |  | [schema](#create-preset-default-schema) |

#### Responses


##### <span id="create-preset-200"></span> 200 - Preset
Status: OK

###### <span id="create-preset-200-schema"></span> Schema
   
  

[Preset](#preset)

##### <span id="create-preset-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="create-preset-401-schema"></span> Schema

##### <span id="create-preset-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="create-preset-403-schema"></span> Schema

##### <span id="create-preset-default"></span> Default Response
errorResponse

###### <span id="create-preset-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="create-project"></span> Creates a brand new project. (*createProject*)

```
POST /api/v1/projects
```

Note that this endpoint can be consumed by every authenticated user.

#### Consumes
  * application/json

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| Body | `body` | [CreateProjectBody](#create-project-body) | `CreateProjectBody` | |  | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [201](#create-project-201) | Created | Project |  | [schema](#create-project-201-schema) |
| [401](#create-project-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#create-project-401-schema) |
| [409](#create-project-409) | Conflict | EmptyResponse is a empty response |  | [schema](#create-project-409-schema) |
| [default](#create-project-default) | | errorResponse |  | [schema](#create-project-default-schema) |

#### Responses


##### <span id="create-project-201"></span> 201 - Project
Status: Created

###### <span id="create-project-201-schema"></span> Schema
   
  

[Project](#project)

##### <span id="create-project-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="create-project-401-schema"></span> Schema

##### <span id="create-project-409"></span> 409 - EmptyResponse is a empty response
Status: Conflict

###### <span id="create-project-409-schema"></span> Schema

##### <span id="create-project-default"></span> Default Response
errorResponse

###### <span id="create-project-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

###### Inlined models

**<span id="create-project-body"></span> CreateProjectBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Labels | map of string| `map[string]string` |  | |  |  |
| Name | string| `string` |  | |  |  |
| Users | []string| `[]string` |  | | human user email list for the service account in projectmanagers group |  |



### <span id="create-role"></span> create role (*createRole*)

```
POST /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/roles
```

Creates cluster role

#### Consumes
  * application/json

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| dc | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |
| Body | `body` | [Role](#role) | `models.Role` | |  | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [201](#create-role-201) | Created | Role |  | [schema](#create-role-201-schema) |
| [401](#create-role-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#create-role-401-schema) |
| [403](#create-role-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#create-role-403-schema) |
| [default](#create-role-default) | | errorResponse |  | [schema](#create-role-default-schema) |

#### Responses


##### <span id="create-role-201"></span> 201 - Role
Status: Created

###### <span id="create-role-201-schema"></span> Schema
   
  

[Role](#role)

##### <span id="create-role-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="create-role-401-schema"></span> Schema

##### <span id="create-role-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="create-role-403-schema"></span> Schema

##### <span id="create-role-default"></span> Default Response
errorResponse

###### <span id="create-role-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="create-rule-group"></span> create rule group (*createRuleGroup*)

```
POST /api/v2/projects/{project_id}/clusters/{cluster_id}/rulegroups
```

Creates a rule group that will belong to the given cluster

#### Consumes
  * application/json

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |
| Body | `body` | [RuleGroup](#rule-group) | `models.RuleGroup` | | ✓ | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [201](#create-rule-group-201) | Created | RuleGroup |  | [schema](#create-rule-group-201-schema) |
| [401](#create-rule-group-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#create-rule-group-401-schema) |
| [403](#create-rule-group-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#create-rule-group-403-schema) |
| [default](#create-rule-group-default) | | errorResponse |  | [schema](#create-rule-group-default-schema) |

#### Responses


##### <span id="create-rule-group-201"></span> 201 - RuleGroup
Status: Created

###### <span id="create-rule-group-201-schema"></span> Schema
   
  

[RuleGroup](#rule-group)

##### <span id="create-rule-group-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="create-rule-group-401-schema"></span> Schema

##### <span id="create-rule-group-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="create-rule-group-403-schema"></span> Schema

##### <span id="create-rule-group-default"></span> Default Response
errorResponse

###### <span id="create-rule-group-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="create-ssh-key"></span> Adds the given SSH key to the specified project. (*createSSHKey*)

```
POST /api/v1/projects/{project_id}/sshkeys
```

#### Consumes
  * application/json

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| project_id | `path` | string | `string` |  | ✓ |  |  |
| Key | `body` | [SSHKey](#ssh-key) | `models.SSHKey` | |  | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [201](#create-ssh-key-201) | Created | SSHKey |  | [schema](#create-ssh-key-201-schema) |
| [401](#create-ssh-key-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#create-ssh-key-401-schema) |
| [403](#create-ssh-key-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#create-ssh-key-403-schema) |
| [default](#create-ssh-key-default) | | errorResponse |  | [schema](#create-ssh-key-default-schema) |

#### Responses


##### <span id="create-ssh-key-201"></span> 201 - SSHKey
Status: Created

###### <span id="create-ssh-key-201-schema"></span> Schema
   
  

[SSHKey](#ssh-key)

##### <span id="create-ssh-key-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="create-ssh-key-401-schema"></span> Schema

##### <span id="create-ssh-key-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="create-ssh-key-403-schema"></span> Schema

##### <span id="create-ssh-key-default"></span> Default Response
errorResponse

###### <span id="create-ssh-key-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="delete-addon"></span> Deletes the given addon that belongs to the cluster. (*deleteAddon*)

```
DELETE /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/addons/{addon_id}
```

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| addon_id | `path` | string | `string` |  | ✓ |  |  |
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| dc | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#delete-addon-200) | OK | EmptyResponse is a empty response |  | [schema](#delete-addon-200-schema) |
| [401](#delete-addon-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#delete-addon-401-schema) |
| [403](#delete-addon-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#delete-addon-403-schema) |
| [default](#delete-addon-default) | | errorResponse |  | [schema](#delete-addon-default-schema) |

#### Responses


##### <span id="delete-addon-200"></span> 200 - EmptyResponse is a empty response
Status: OK

###### <span id="delete-addon-200-schema"></span> Schema

##### <span id="delete-addon-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="delete-addon-401-schema"></span> Schema

##### <span id="delete-addon-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="delete-addon-403-schema"></span> Schema

##### <span id="delete-addon-default"></span> Default Response
errorResponse

###### <span id="delete-addon-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="delete-addon-v2"></span> Deletes the given addon that belongs to the cluster. (*deleteAddonV2*)

```
DELETE /api/v2/projects/{project_id}/clusters/{cluster_id}/addons/{addon_id}
```

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| addon_id | `path` | string | `string` |  | ✓ |  |  |
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#delete-addon-v2-200) | OK | EmptyResponse is a empty response |  | [schema](#delete-addon-v2-200-schema) |
| [401](#delete-addon-v2-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#delete-addon-v2-401-schema) |
| [403](#delete-addon-v2-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#delete-addon-v2-403-schema) |
| [default](#delete-addon-v2-default) | | errorResponse |  | [schema](#delete-addon-v2-default-schema) |

#### Responses


##### <span id="delete-addon-v2-200"></span> 200 - EmptyResponse is a empty response
Status: OK

###### <span id="delete-addon-v2-200-schema"></span> Schema

##### <span id="delete-addon-v2-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="delete-addon-v2-401-schema"></span> Schema

##### <span id="delete-addon-v2-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="delete-addon-v2-403-schema"></span> Schema

##### <span id="delete-addon-v2-default"></span> Default Response
errorResponse

###### <span id="delete-addon-v2-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="delete-admission-plugin"></span> Deletes the admission plugin. (*deleteAdmissionPlugin*)

```
DELETE /api/v1/admin/admission/plugins/{name}
```

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| name | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#delete-admission-plugin-200) | OK | EmptyResponse is a empty response |  | [schema](#delete-admission-plugin-200-schema) |
| [401](#delete-admission-plugin-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#delete-admission-plugin-401-schema) |
| [403](#delete-admission-plugin-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#delete-admission-plugin-403-schema) |
| [default](#delete-admission-plugin-default) | | errorResponse |  | [schema](#delete-admission-plugin-default-schema) |

#### Responses


##### <span id="delete-admission-plugin-200"></span> 200 - EmptyResponse is a empty response
Status: OK

###### <span id="delete-admission-plugin-200-schema"></span> Schema

##### <span id="delete-admission-plugin-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="delete-admission-plugin-401-schema"></span> Schema

##### <span id="delete-admission-plugin-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="delete-admission-plugin-403-schema"></span> Schema

##### <span id="delete-admission-plugin-default"></span> Default Response
errorResponse

###### <span id="delete-admission-plugin-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="delete-allowed-registry"></span> Deletes the given allowed registry. (*deleteAllowedRegistry*)

```
DELETE /api/v2/allowedregistries/{allowed_registry}
```

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| allowed_registry | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#delete-allowed-registry-200) | OK | EmptyResponse is a empty response |  | [schema](#delete-allowed-registry-200-schema) |
| [401](#delete-allowed-registry-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#delete-allowed-registry-401-schema) |
| [403](#delete-allowed-registry-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#delete-allowed-registry-403-schema) |
| [default](#delete-allowed-registry-default) | | errorResponse |  | [schema](#delete-allowed-registry-default-schema) |

#### Responses


##### <span id="delete-allowed-registry-200"></span> 200 - EmptyResponse is a empty response
Status: OK

###### <span id="delete-allowed-registry-200-schema"></span> Schema

##### <span id="delete-allowed-registry-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="delete-allowed-registry-401-schema"></span> Schema

##### <span id="delete-allowed-registry-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="delete-allowed-registry-403-schema"></span> Schema

##### <span id="delete-allowed-registry-default"></span> Default Response
errorResponse

###### <span id="delete-allowed-registry-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="delete-cluster"></span> delete cluster (*deleteCluster*)

```
DELETE /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}
```

Deletes the specified cluster

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| dc | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |
| DeleteLoadBalancers | `header` | boolean | `bool` |  |  |  |  |
| DeleteVolumes | `header` | boolean | `bool` |  |  |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#delete-cluster-200) | OK | EmptyResponse is a empty response |  | [schema](#delete-cluster-200-schema) |
| [401](#delete-cluster-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#delete-cluster-401-schema) |
| [403](#delete-cluster-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#delete-cluster-403-schema) |
| [default](#delete-cluster-default) | | errorResponse |  | [schema](#delete-cluster-default-schema) |

#### Responses


##### <span id="delete-cluster-200"></span> 200 - EmptyResponse is a empty response
Status: OK

###### <span id="delete-cluster-200-schema"></span> Schema

##### <span id="delete-cluster-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="delete-cluster-401-schema"></span> Schema

##### <span id="delete-cluster-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="delete-cluster-403-schema"></span> Schema

##### <span id="delete-cluster-default"></span> Default Response
errorResponse

###### <span id="delete-cluster-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="delete-cluster-role"></span> delete cluster role (*deleteClusterRole*)

```
DELETE /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/clusterroles/{role_id}
```

Delete the cluster role with the given name

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| dc | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |
| role_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#delete-cluster-role-200) | OK | EmptyResponse is a empty response |  | [schema](#delete-cluster-role-200-schema) |
| [401](#delete-cluster-role-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#delete-cluster-role-401-schema) |
| [403](#delete-cluster-role-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#delete-cluster-role-403-schema) |
| [default](#delete-cluster-role-default) | | errorResponse |  | [schema](#delete-cluster-role-default-schema) |

#### Responses


##### <span id="delete-cluster-role-200"></span> 200 - EmptyResponse is a empty response
Status: OK

###### <span id="delete-cluster-role-200-schema"></span> Schema

##### <span id="delete-cluster-role-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="delete-cluster-role-401-schema"></span> Schema

##### <span id="delete-cluster-role-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="delete-cluster-role-403-schema"></span> Schema

##### <span id="delete-cluster-role-default"></span> Default Response
errorResponse

###### <span id="delete-cluster-role-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="delete-cluster-template"></span> Delete cluster template. (*deleteClusterTemplate*)

```
DELETE /api/v2/projects/{project_id}/clustertemplates/{template_id}
```

#### Consumes
  * application/json

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| project_id | `path` | string | `string` |  | ✓ |  |  |
| template_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#delete-cluster-template-200) | OK | EmptyResponse is a empty response |  | [schema](#delete-cluster-template-200-schema) |
| [401](#delete-cluster-template-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#delete-cluster-template-401-schema) |
| [403](#delete-cluster-template-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#delete-cluster-template-403-schema) |
| [default](#delete-cluster-template-default) | | errorResponse |  | [schema](#delete-cluster-template-default-schema) |

#### Responses


##### <span id="delete-cluster-template-200"></span> 200 - EmptyResponse is a empty response
Status: OK

###### <span id="delete-cluster-template-200-schema"></span> Schema

##### <span id="delete-cluster-template-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="delete-cluster-template-401-schema"></span> Schema

##### <span id="delete-cluster-template-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="delete-cluster-template-403-schema"></span> Schema

##### <span id="delete-cluster-template-default"></span> Default Response
errorResponse

###### <span id="delete-cluster-template-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="delete-cluster-v2"></span> delete cluster v2 (*deleteClusterV2*)

```
DELETE /api/v2/projects/{project_id}/clusters/{cluster_id}
```

Deletes the specified cluster

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |
| DeleteLoadBalancers | `header` | boolean | `bool` |  |  |  |  |
| DeleteVolumes | `header` | boolean | `bool` |  |  |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#delete-cluster-v2-200) | OK | EmptyResponse is a empty response |  | [schema](#delete-cluster-v2-200-schema) |
| [401](#delete-cluster-v2-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#delete-cluster-v2-401-schema) |
| [403](#delete-cluster-v2-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#delete-cluster-v2-403-schema) |
| [default](#delete-cluster-v2-default) | | errorResponse |  | [schema](#delete-cluster-v2-default-schema) |

#### Responses


##### <span id="delete-cluster-v2-200"></span> 200 - EmptyResponse is a empty response
Status: OK

###### <span id="delete-cluster-v2-200-schema"></span> Schema

##### <span id="delete-cluster-v2-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="delete-cluster-v2-401-schema"></span> Schema

##### <span id="delete-cluster-v2-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="delete-cluster-v2-403-schema"></span> Schema

##### <span id="delete-cluster-v2-default"></span> Default Response
errorResponse

###### <span id="delete-cluster-v2-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="delete-constraint"></span> Deletes a specified constraint for the given cluster. (*deleteConstraint*)

```
DELETE /api/v2/projects/{project_id}/clusters/{cluster_id}/constraints/{constraint_name}
```

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| constraint_name | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#delete-constraint-200) | OK | EmptyResponse is a empty response |  | [schema](#delete-constraint-200-schema) |
| [401](#delete-constraint-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#delete-constraint-401-schema) |
| [403](#delete-constraint-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#delete-constraint-403-schema) |
| [default](#delete-constraint-default) | | errorResponse |  | [schema](#delete-constraint-default-schema) |

#### Responses


##### <span id="delete-constraint-200"></span> 200 - EmptyResponse is a empty response
Status: OK

###### <span id="delete-constraint-200-schema"></span> Schema

##### <span id="delete-constraint-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="delete-constraint-401-schema"></span> Schema

##### <span id="delete-constraint-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="delete-constraint-403-schema"></span> Schema

##### <span id="delete-constraint-default"></span> Default Response
errorResponse

###### <span id="delete-constraint-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="delete-constraint-template"></span> delete constraint template (*deleteConstraintTemplate*)

```
DELETE /api/v2/constrainttemplates/{ct_name}
```

Deletes the specified cluster

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| ct_name | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#delete-constraint-template-200) | OK | EmptyResponse is a empty response |  | [schema](#delete-constraint-template-200-schema) |
| [401](#delete-constraint-template-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#delete-constraint-template-401-schema) |
| [403](#delete-constraint-template-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#delete-constraint-template-403-schema) |
| [default](#delete-constraint-template-default) | | errorResponse |  | [schema](#delete-constraint-template-default-schema) |

#### Responses


##### <span id="delete-constraint-template-200"></span> 200 - EmptyResponse is a empty response
Status: OK

###### <span id="delete-constraint-template-200-schema"></span> Schema

##### <span id="delete-constraint-template-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="delete-constraint-template-401-schema"></span> Schema

##### <span id="delete-constraint-template-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="delete-constraint-template-403-schema"></span> Schema

##### <span id="delete-constraint-template-default"></span> Default Response
errorResponse

###### <span id="delete-constraint-template-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="delete-d-c"></span> Delete the datacenter. (*deleteDC*)

```
DELETE /api/v1/seed/{seed_name}/dc/{dc}
```

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| dc | `path` | string | `string` |  | ✓ |  |  |
| seed_name | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#delete-d-c-200) | OK | EmptyResponse is a empty response |  | [schema](#delete-d-c-200-schema) |
| [401](#delete-d-c-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#delete-d-c-401-schema) |
| [403](#delete-d-c-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#delete-d-c-403-schema) |
| [default](#delete-d-c-default) | | errorResponse |  | [schema](#delete-d-c-default-schema) |

#### Responses


##### <span id="delete-d-c-200"></span> 200 - EmptyResponse is a empty response
Status: OK

###### <span id="delete-d-c-200-schema"></span> Schema

##### <span id="delete-d-c-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="delete-d-c-401-schema"></span> Schema

##### <span id="delete-d-c-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="delete-d-c-403-schema"></span> Schema

##### <span id="delete-d-c-default"></span> Default Response
errorResponse

###### <span id="delete-d-c-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="delete-default-constraint"></span> Deletes a specified default constraint. (*deleteDefaultConstraint*)

```
DELETE /api/v2/constraints/{constraint_name}
```

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| constraint_name | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#delete-default-constraint-200) | OK | EmptyResponse is a empty response |  | [schema](#delete-default-constraint-200-schema) |
| [401](#delete-default-constraint-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#delete-default-constraint-401-schema) |
| [403](#delete-default-constraint-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#delete-default-constraint-403-schema) |
| [default](#delete-default-constraint-default) | | errorResponse |  | [schema](#delete-default-constraint-default-schema) |

#### Responses


##### <span id="delete-default-constraint-200"></span> 200 - EmptyResponse is a empty response
Status: OK

###### <span id="delete-default-constraint-200-schema"></span> Schema

##### <span id="delete-default-constraint-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="delete-default-constraint-401-schema"></span> Schema

##### <span id="delete-default-constraint-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="delete-default-constraint-403-schema"></span> Schema

##### <span id="delete-default-constraint-default"></span> Default Response
errorResponse

###### <span id="delete-default-constraint-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="delete-etcd-backup-config"></span> delete etcd backup config (*deleteEtcdBackupConfig*)

```
DELETE /api/v2/projects/{project_id}/clusters/{cluster_id}/etcdbackupconfigs/{ebc_id}
```

Deletes a etcd backup config for a given cluster based on its id

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| ebc_id | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#delete-etcd-backup-config-200) | OK | EmptyResponse is a empty response |  | [schema](#delete-etcd-backup-config-200-schema) |
| [401](#delete-etcd-backup-config-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#delete-etcd-backup-config-401-schema) |
| [403](#delete-etcd-backup-config-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#delete-etcd-backup-config-403-schema) |
| [default](#delete-etcd-backup-config-default) | | errorResponse |  | [schema](#delete-etcd-backup-config-default-schema) |

#### Responses


##### <span id="delete-etcd-backup-config-200"></span> 200 - EmptyResponse is a empty response
Status: OK

###### <span id="delete-etcd-backup-config-200-schema"></span> Schema

##### <span id="delete-etcd-backup-config-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="delete-etcd-backup-config-401-schema"></span> Schema

##### <span id="delete-etcd-backup-config-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="delete-etcd-backup-config-403-schema"></span> Schema

##### <span id="delete-etcd-backup-config-default"></span> Default Response
errorResponse

###### <span id="delete-etcd-backup-config-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="delete-etcd-restore"></span> delete etcd restore (*deleteEtcdRestore*)

```
DELETE /api/v2/projects/{project_id}/clusters/{cluster_id}/etcdrestores/{er_name}
```

Deletes a etcd restore config for a given cluster based on its name

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| er_name | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#delete-etcd-restore-200) | OK | EmptyResponse is a empty response |  | [schema](#delete-etcd-restore-200-schema) |
| [401](#delete-etcd-restore-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#delete-etcd-restore-401-schema) |
| [403](#delete-etcd-restore-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#delete-etcd-restore-403-schema) |
| [409](#delete-etcd-restore-409) | Conflict | errorResponse |  | [schema](#delete-etcd-restore-409-schema) |
| [default](#delete-etcd-restore-default) | | errorResponse |  | [schema](#delete-etcd-restore-default-schema) |

#### Responses


##### <span id="delete-etcd-restore-200"></span> 200 - EmptyResponse is a empty response
Status: OK

###### <span id="delete-etcd-restore-200-schema"></span> Schema

##### <span id="delete-etcd-restore-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="delete-etcd-restore-401-schema"></span> Schema

##### <span id="delete-etcd-restore-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="delete-etcd-restore-403-schema"></span> Schema

##### <span id="delete-etcd-restore-409"></span> 409 - errorResponse
Status: Conflict

###### <span id="delete-etcd-restore-409-schema"></span> Schema
   
  

[ErrorResponse](#error-response)

##### <span id="delete-etcd-restore-default"></span> Default Response
errorResponse

###### <span id="delete-etcd-restore-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="delete-external-cluster"></span> delete external cluster (*deleteExternalCluster*)

```
DELETE /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}
```

Deletes the specified external cluster

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#delete-external-cluster-200) | OK | EmptyResponse is a empty response |  | [schema](#delete-external-cluster-200-schema) |
| [401](#delete-external-cluster-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#delete-external-cluster-401-schema) |
| [403](#delete-external-cluster-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#delete-external-cluster-403-schema) |
| [default](#delete-external-cluster-default) | | errorResponse |  | [schema](#delete-external-cluster-default-schema) |

#### Responses


##### <span id="delete-external-cluster-200"></span> 200 - EmptyResponse is a empty response
Status: OK

###### <span id="delete-external-cluster-200-schema"></span> Schema

##### <span id="delete-external-cluster-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="delete-external-cluster-401-schema"></span> Schema

##### <span id="delete-external-cluster-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="delete-external-cluster-403-schema"></span> Schema

##### <span id="delete-external-cluster-default"></span> Default Response
errorResponse

###### <span id="delete-external-cluster-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="delete-gatekeeper-config"></span> Deletes the gatekeeper sync config for the specified cluster. (*deleteGatekeeperConfig*)

```
DELETE /api/v2/projects/{project_id}/clusters/{cluster_id}/gatekeeper/config
```

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#delete-gatekeeper-config-200) | OK | EmptyResponse is a empty response |  | [schema](#delete-gatekeeper-config-200-schema) |
| [401](#delete-gatekeeper-config-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#delete-gatekeeper-config-401-schema) |
| [403](#delete-gatekeeper-config-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#delete-gatekeeper-config-403-schema) |
| [default](#delete-gatekeeper-config-default) | | errorResponse |  | [schema](#delete-gatekeeper-config-default-schema) |

#### Responses


##### <span id="delete-gatekeeper-config-200"></span> 200 - EmptyResponse is a empty response
Status: OK

###### <span id="delete-gatekeeper-config-200-schema"></span> Schema

##### <span id="delete-gatekeeper-config-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="delete-gatekeeper-config-401-schema"></span> Schema

##### <span id="delete-gatekeeper-config-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="delete-gatekeeper-config-403-schema"></span> Schema

##### <span id="delete-gatekeeper-config-default"></span> Default Response
errorResponse

###### <span id="delete-gatekeeper-config-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="delete-m-l-a-admin-setting"></span> Deletes the MLA admin setting that belongs to the cluster. (*deleteMLAAdminSetting*)

```
DELETE /api/v2/projects/{project_id}/clusters/{cluster_id}/mlaadminsetting
```

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#delete-m-l-a-admin-setting-200) | OK | EmptyResponse is a empty response |  | [schema](#delete-m-l-a-admin-setting-200-schema) |
| [401](#delete-m-l-a-admin-setting-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#delete-m-l-a-admin-setting-401-schema) |
| [403](#delete-m-l-a-admin-setting-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#delete-m-l-a-admin-setting-403-schema) |
| [default](#delete-m-l-a-admin-setting-default) | | errorResponse |  | [schema](#delete-m-l-a-admin-setting-default-schema) |

#### Responses


##### <span id="delete-m-l-a-admin-setting-200"></span> 200 - EmptyResponse is a empty response
Status: OK

###### <span id="delete-m-l-a-admin-setting-200-schema"></span> Schema

##### <span id="delete-m-l-a-admin-setting-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="delete-m-l-a-admin-setting-401-schema"></span> Schema

##### <span id="delete-m-l-a-admin-setting-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="delete-m-l-a-admin-setting-403-schema"></span> Schema

##### <span id="delete-m-l-a-admin-setting-default"></span> Default Response
errorResponse

###### <span id="delete-m-l-a-admin-setting-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="delete-machine-deployment"></span> Deletes the given machine deployment that belongs to the cluster. (*deleteMachineDeployment*)

```
DELETE /api/v2/projects/{project_id}/clusters/{cluster_id}/machinedeployments/{machinedeployment_id}
```

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| machinedeployment_id | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#delete-machine-deployment-200) | OK | EmptyResponse is a empty response |  | [schema](#delete-machine-deployment-200-schema) |
| [401](#delete-machine-deployment-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#delete-machine-deployment-401-schema) |
| [403](#delete-machine-deployment-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#delete-machine-deployment-403-schema) |
| [default](#delete-machine-deployment-default) | | errorResponse |  | [schema](#delete-machine-deployment-default-schema) |

#### Responses


##### <span id="delete-machine-deployment-200"></span> 200 - EmptyResponse is a empty response
Status: OK

###### <span id="delete-machine-deployment-200-schema"></span> Schema

##### <span id="delete-machine-deployment-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="delete-machine-deployment-401-schema"></span> Schema

##### <span id="delete-machine-deployment-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="delete-machine-deployment-403-schema"></span> Schema

##### <span id="delete-machine-deployment-default"></span> Default Response
errorResponse

###### <span id="delete-machine-deployment-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="delete-machine-deployment-node"></span> Deletes the given node that belongs to the machine deployment. (*deleteMachineDeploymentNode*)

```
DELETE /api/v2/projects/{project_id}/clusters/{cluster_id}/machinedeployments/nodes/{node_id}
```

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| node_id | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#delete-machine-deployment-node-200) | OK | EmptyResponse is a empty response |  | [schema](#delete-machine-deployment-node-200-schema) |
| [401](#delete-machine-deployment-node-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#delete-machine-deployment-node-401-schema) |
| [403](#delete-machine-deployment-node-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#delete-machine-deployment-node-403-schema) |
| [default](#delete-machine-deployment-node-default) | | errorResponse |  | [schema](#delete-machine-deployment-node-default-schema) |

#### Responses


##### <span id="delete-machine-deployment-node-200"></span> 200 - EmptyResponse is a empty response
Status: OK

###### <span id="delete-machine-deployment-node-200-schema"></span> Schema

##### <span id="delete-machine-deployment-node-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="delete-machine-deployment-node-401-schema"></span> Schema

##### <span id="delete-machine-deployment-node-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="delete-machine-deployment-node-403-schema"></span> Schema

##### <span id="delete-machine-deployment-node-default"></span> Default Response
errorResponse

###### <span id="delete-machine-deployment-node-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="delete-node-deployment"></span> Deletes the given node deployment that belongs to the cluster. (*deleteNodeDeployment*)

```
DELETE /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/nodedeployments/{nodedeployment_id}
```

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| dc | `path` | string | `string` |  | ✓ |  |  |
| nodedeployment_id | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#delete-node-deployment-200) | OK | EmptyResponse is a empty response |  | [schema](#delete-node-deployment-200-schema) |
| [401](#delete-node-deployment-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#delete-node-deployment-401-schema) |
| [403](#delete-node-deployment-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#delete-node-deployment-403-schema) |
| [default](#delete-node-deployment-default) | | errorResponse |  | [schema](#delete-node-deployment-default-schema) |

#### Responses


##### <span id="delete-node-deployment-200"></span> 200 - EmptyResponse is a empty response
Status: OK

###### <span id="delete-node-deployment-200-schema"></span> Schema

##### <span id="delete-node-deployment-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="delete-node-deployment-401-schema"></span> Schema

##### <span id="delete-node-deployment-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="delete-node-deployment-403-schema"></span> Schema

##### <span id="delete-node-deployment-default"></span> Default Response
errorResponse

###### <span id="delete-node-deployment-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="delete-node-deployment-request"></span> Deletes the given NodeDeploymentRequest that belongs to the cluster. (*deleteNodeDeploymentRequest*)

```
DELETE /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/ndrequests/{ndrequest_id}
```

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| dc | `path` | string | `string` |  | ✓ |  |  |
| ndrequest_id | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#delete-node-deployment-request-200) | OK | EmptyResponse is a empty response |  | [schema](#delete-node-deployment-request-200-schema) |
| [401](#delete-node-deployment-request-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#delete-node-deployment-request-401-schema) |
| [403](#delete-node-deployment-request-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#delete-node-deployment-request-403-schema) |
| [default](#delete-node-deployment-request-default) | | errorResponse |  | [schema](#delete-node-deployment-request-default-schema) |

#### Responses


##### <span id="delete-node-deployment-request-200"></span> 200 - EmptyResponse is a empty response
Status: OK

###### <span id="delete-node-deployment-request-200-schema"></span> Schema

##### <span id="delete-node-deployment-request-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="delete-node-deployment-request-401-schema"></span> Schema

##### <span id="delete-node-deployment-request-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="delete-node-deployment-request-403-schema"></span> Schema

##### <span id="delete-node-deployment-request-default"></span> Default Response
errorResponse

###### <span id="delete-node-deployment-request-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="delete-node-for-cluster-legacy"></span> Deprecated:
Deletes the given node that belongs to the cluster. (*deleteNodeForClusterLegacy*)

```
DELETE /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/nodes/{node_id}
```

This endpoint is deprecated, please create a Node Deployment instead.

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| dc | `path` | string | `string` |  | ✓ |  |  |
| node_id | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#delete-node-for-cluster-legacy-200) | OK | EmptyResponse is a empty response |  | [schema](#delete-node-for-cluster-legacy-200-schema) |
| [401](#delete-node-for-cluster-legacy-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#delete-node-for-cluster-legacy-401-schema) |
| [403](#delete-node-for-cluster-legacy-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#delete-node-for-cluster-legacy-403-schema) |
| [default](#delete-node-for-cluster-legacy-default) | | errorResponse |  | [schema](#delete-node-for-cluster-legacy-default-schema) |

#### Responses


##### <span id="delete-node-for-cluster-legacy-200"></span> 200 - EmptyResponse is a empty response
Status: OK

###### <span id="delete-node-for-cluster-legacy-200-schema"></span> Schema

##### <span id="delete-node-for-cluster-legacy-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="delete-node-for-cluster-legacy-401-schema"></span> Schema

##### <span id="delete-node-for-cluster-legacy-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="delete-node-for-cluster-legacy-403-schema"></span> Schema

##### <span id="delete-node-for-cluster-legacy-default"></span> Default Response
errorResponse

###### <span id="delete-node-for-cluster-legacy-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="delete-project"></span> Deletes the project with the given ID. (*deleteProject*)

```
DELETE /api/v1/projects/{project_id}
```

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#delete-project-200) | OK | EmptyResponse is a empty response |  | [schema](#delete-project-200-schema) |
| [401](#delete-project-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#delete-project-401-schema) |
| [403](#delete-project-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#delete-project-403-schema) |
| [default](#delete-project-default) | | errorResponse |  | [schema](#delete-project-default-schema) |

#### Responses


##### <span id="delete-project-200"></span> 200 - EmptyResponse is a empty response
Status: OK

###### <span id="delete-project-200-schema"></span> Schema

##### <span id="delete-project-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="delete-project-401-schema"></span> Schema

##### <span id="delete-project-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="delete-project-403-schema"></span> Schema

##### <span id="delete-project-default"></span> Default Response
errorResponse

###### <span id="delete-project-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="delete-role"></span> delete role (*deleteRole*)

```
DELETE /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/roles/{namespace}/{role_id}
```

Delete the cluster role with the given name

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| dc | `path` | string | `string` |  | ✓ |  |  |
| namespace | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |
| role_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#delete-role-200) | OK | EmptyResponse is a empty response |  | [schema](#delete-role-200-schema) |
| [401](#delete-role-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#delete-role-401-schema) |
| [403](#delete-role-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#delete-role-403-schema) |
| [default](#delete-role-default) | | errorResponse |  | [schema](#delete-role-default-schema) |

#### Responses


##### <span id="delete-role-200"></span> 200 - EmptyResponse is a empty response
Status: OK

###### <span id="delete-role-200-schema"></span> Schema

##### <span id="delete-role-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="delete-role-401-schema"></span> Schema

##### <span id="delete-role-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="delete-role-403-schema"></span> Schema

##### <span id="delete-role-default"></span> Default Response
errorResponse

###### <span id="delete-role-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="delete-rule-group"></span> Deletes the given rule group that belongs to the cluster. (*deleteRuleGroup*)

```
DELETE /api/v2/projects/{project_id}/clusters/{cluster_id}/rulegroups/{rulegroup_id}
```

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |
| rulegroup_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#delete-rule-group-200) | OK | EmptyResponse is a empty response |  | [schema](#delete-rule-group-200-schema) |
| [401](#delete-rule-group-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#delete-rule-group-401-schema) |
| [403](#delete-rule-group-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#delete-rule-group-403-schema) |
| [default](#delete-rule-group-default) | | errorResponse |  | [schema](#delete-rule-group-default-schema) |

#### Responses


##### <span id="delete-rule-group-200"></span> 200 - EmptyResponse is a empty response
Status: OK

###### <span id="delete-rule-group-200-schema"></span> Schema

##### <span id="delete-rule-group-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="delete-rule-group-401-schema"></span> Schema

##### <span id="delete-rule-group-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="delete-rule-group-403-schema"></span> Schema

##### <span id="delete-rule-group-default"></span> Default Response
errorResponse

###### <span id="delete-rule-group-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="delete-ssh-key"></span> Removes the given SSH Key from the system. (*deleteSSHKey*)

```
DELETE /api/v1/projects/{project_id}/sshkeys/{key_id}
```

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| key_id | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#delete-ssh-key-200) | OK | EmptyResponse is a empty response |  | [schema](#delete-ssh-key-200-schema) |
| [401](#delete-ssh-key-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#delete-ssh-key-401-schema) |
| [403](#delete-ssh-key-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#delete-ssh-key-403-schema) |
| [default](#delete-ssh-key-default) | | errorResponse |  | [schema](#delete-ssh-key-default-schema) |

#### Responses


##### <span id="delete-ssh-key-200"></span> 200 - EmptyResponse is a empty response
Status: OK

###### <span id="delete-ssh-key-200-schema"></span> Schema

##### <span id="delete-ssh-key-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="delete-ssh-key-401-schema"></span> Schema

##### <span id="delete-ssh-key-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="delete-ssh-key-403-schema"></span> Schema

##### <span id="delete-ssh-key-default"></span> Default Response
errorResponse

###### <span id="delete-ssh-key-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="delete-seed"></span> Deletes the seed CRD object from the Kubermatic. (*deleteSeed*)

```
DELETE /api/v1/admin/seeds/{seed_name}
```

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| seed_name | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#delete-seed-200) | OK | EmptyResponse is a empty response |  | [schema](#delete-seed-200-schema) |
| [401](#delete-seed-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#delete-seed-401-schema) |
| [403](#delete-seed-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#delete-seed-403-schema) |
| [default](#delete-seed-default) | | errorResponse |  | [schema](#delete-seed-default-schema) |

#### Responses


##### <span id="delete-seed-200"></span> 200 - EmptyResponse is a empty response
Status: OK

###### <span id="delete-seed-200-schema"></span> Schema

##### <span id="delete-seed-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="delete-seed-401-schema"></span> Schema

##### <span id="delete-seed-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="delete-seed-403-schema"></span> Schema

##### <span id="delete-seed-default"></span> Default Response
errorResponse

###### <span id="delete-seed-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="delete-service-account"></span> delete service account (*deleteServiceAccount*)

```
DELETE /api/v1/projects/{project_id}/serviceaccounts/{serviceaccount_id}
```

Deletes service account for the given project

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| project_id | `path` | string | `string` |  | ✓ |  |  |
| serviceaccount_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#delete-service-account-200) | OK | EmptyResponse is a empty response |  | [schema](#delete-service-account-200-schema) |
| [401](#delete-service-account-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#delete-service-account-401-schema) |
| [403](#delete-service-account-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#delete-service-account-403-schema) |
| [default](#delete-service-account-default) | | errorResponse |  | [schema](#delete-service-account-default-schema) |

#### Responses


##### <span id="delete-service-account-200"></span> 200 - EmptyResponse is a empty response
Status: OK

###### <span id="delete-service-account-200-schema"></span> Schema

##### <span id="delete-service-account-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="delete-service-account-401-schema"></span> Schema

##### <span id="delete-service-account-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="delete-service-account-403-schema"></span> Schema

##### <span id="delete-service-account-default"></span> Default Response
errorResponse

###### <span id="delete-service-account-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="delete-service-account-token"></span> delete service account token (*deleteServiceAccountToken*)

```
DELETE /api/v1/projects/{project_id}/serviceaccounts/{serviceaccount_id}/tokens/{token_id}
```

Deletes the token

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| project_id | `path` | string | `string` |  | ✓ |  |  |
| serviceaccount_id | `path` | string | `string` |  | ✓ |  |  |
| token_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#delete-service-account-token-200) | OK | EmptyResponse is a empty response |  | [schema](#delete-service-account-token-200-schema) |
| [401](#delete-service-account-token-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#delete-service-account-token-401-schema) |
| [403](#delete-service-account-token-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#delete-service-account-token-403-schema) |
| [default](#delete-service-account-token-default) | | errorResponse |  | [schema](#delete-service-account-token-default-schema) |

#### Responses


##### <span id="delete-service-account-token-200"></span> 200 - EmptyResponse is a empty response
Status: OK

###### <span id="delete-service-account-token-200-schema"></span> Schema

##### <span id="delete-service-account-token-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="delete-service-account-token-401-schema"></span> Schema

##### <span id="delete-service-account-token-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="delete-service-account-token-403-schema"></span> Schema

##### <span id="delete-service-account-token-default"></span> Default Response
errorResponse

###### <span id="delete-service-account-token-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="delete-user-from-project"></span> delete user from project (*deleteUserFromProject*)

```
DELETE /api/v1/projects/{project_id}/users/{user_id}
```

Removes the given member from the project

#### Consumes
  * application/json

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| project_id | `path` | string | `string` |  | ✓ |  |  |
| user_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#delete-user-from-project-200) | OK | User |  | [schema](#delete-user-from-project-200-schema) |
| [401](#delete-user-from-project-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#delete-user-from-project-401-schema) |
| [403](#delete-user-from-project-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#delete-user-from-project-403-schema) |
| [default](#delete-user-from-project-default) | | errorResponse |  | [schema](#delete-user-from-project-default-schema) |

#### Responses


##### <span id="delete-user-from-project-200"></span> 200 - User
Status: OK

###### <span id="delete-user-from-project-200-schema"></span> Schema
   
  

[User](#user)

##### <span id="delete-user-from-project-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="delete-user-from-project-401-schema"></span> Schema

##### <span id="delete-user-from-project-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="delete-user-from-project-403-schema"></span> Schema

##### <span id="delete-user-from-project-default"></span> Default Response
errorResponse

###### <span id="delete-user-from-project-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="detach-ssh-key-from-cluster"></span> detach SSH key from cluster (*detachSSHKeyFromCluster*)

```
DELETE /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/sshkeys/{key_id}
```

Unassignes an ssh key from the given cluster

#### Consumes
  * application/json

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| dc | `path` | string | `string` |  | ✓ |  |  |
| key_id | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#detach-ssh-key-from-cluster-200) | OK | EmptyResponse is a empty response |  | [schema](#detach-ssh-key-from-cluster-200-schema) |
| [401](#detach-ssh-key-from-cluster-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#detach-ssh-key-from-cluster-401-schema) |
| [403](#detach-ssh-key-from-cluster-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#detach-ssh-key-from-cluster-403-schema) |
| [default](#detach-ssh-key-from-cluster-default) | | errorResponse |  | [schema](#detach-ssh-key-from-cluster-default-schema) |

#### Responses


##### <span id="detach-ssh-key-from-cluster-200"></span> 200 - EmptyResponse is a empty response
Status: OK

###### <span id="detach-ssh-key-from-cluster-200-schema"></span> Schema

##### <span id="detach-ssh-key-from-cluster-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="detach-ssh-key-from-cluster-401-schema"></span> Schema

##### <span id="detach-ssh-key-from-cluster-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="detach-ssh-key-from-cluster-403-schema"></span> Schema

##### <span id="detach-ssh-key-from-cluster-default"></span> Default Response
errorResponse

###### <span id="detach-ssh-key-from-cluster-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="detach-ssh-key-from-cluster-v2"></span> detach SSH key from cluster v2 (*detachSSHKeyFromClusterV2*)

```
DELETE /api/projects/{project_id}/clusters/{cluster_id}/sshkeys/{key_id}
```

Unassignes an ssh key from the given cluster

#### Consumes
  * application/json

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| key_id | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#detach-ssh-key-from-cluster-v2-200) | OK | EmptyResponse is a empty response |  | [schema](#detach-ssh-key-from-cluster-v2-200-schema) |
| [401](#detach-ssh-key-from-cluster-v2-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#detach-ssh-key-from-cluster-v2-401-schema) |
| [403](#detach-ssh-key-from-cluster-v2-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#detach-ssh-key-from-cluster-v2-403-schema) |
| [default](#detach-ssh-key-from-cluster-v2-default) | | errorResponse |  | [schema](#detach-ssh-key-from-cluster-v2-default-schema) |

#### Responses


##### <span id="detach-ssh-key-from-cluster-v2-200"></span> 200 - EmptyResponse is a empty response
Status: OK

###### <span id="detach-ssh-key-from-cluster-v2-200-schema"></span> Schema

##### <span id="detach-ssh-key-from-cluster-v2-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="detach-ssh-key-from-cluster-v2-401-schema"></span> Schema

##### <span id="detach-ssh-key-from-cluster-v2-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="detach-ssh-key-from-cluster-v2-403-schema"></span> Schema

##### <span id="detach-ssh-key-from-cluster-v2-default"></span> Default Response
errorResponse

###### <span id="detach-ssh-key-from-cluster-v2-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="edit-user-in-project"></span> edit user in project (*editUserInProject*)

```
PUT /api/v1/projects/{project_id}/users/{user_id}
```

Changes membership of the given user for the given project

#### Consumes
  * application/json

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| project_id | `path` | string | `string` |  | ✓ |  |  |
| user_id | `path` | string | `string` |  | ✓ |  |  |
| Body | `body` | [User](#user) | `models.User` | |  | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#edit-user-in-project-200) | OK | User |  | [schema](#edit-user-in-project-200-schema) |
| [401](#edit-user-in-project-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#edit-user-in-project-401-schema) |
| [403](#edit-user-in-project-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#edit-user-in-project-403-schema) |
| [default](#edit-user-in-project-default) | | errorResponse |  | [schema](#edit-user-in-project-default-schema) |

#### Responses


##### <span id="edit-user-in-project-200"></span> 200 - User
Status: OK

###### <span id="edit-user-in-project-200-schema"></span> Schema
   
  

[User](#user)

##### <span id="edit-user-in-project-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="edit-user-in-project-401-schema"></span> Schema

##### <span id="edit-user-in-project-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="edit-user-in-project-403-schema"></span> Schema

##### <span id="edit-user-in-project-default"></span> Default Response
errorResponse

###### <span id="edit-user-in-project-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="get-addon"></span> Gets an addon that is assigned to the given cluster. (*getAddon*)

```
GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/addons/{addon_id}
```

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| addon_id | `path` | string | `string` |  | ✓ |  |  |
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| dc | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-addon-200) | OK | Addon |  | [schema](#get-addon-200-schema) |
| [401](#get-addon-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#get-addon-401-schema) |
| [403](#get-addon-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#get-addon-403-schema) |
| [default](#get-addon-default) | | errorResponse |  | [schema](#get-addon-default-schema) |

#### Responses


##### <span id="get-addon-200"></span> 200 - Addon
Status: OK

###### <span id="get-addon-200-schema"></span> Schema
   
  

[Addon](#addon)

##### <span id="get-addon-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="get-addon-401-schema"></span> Schema

##### <span id="get-addon-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="get-addon-403-schema"></span> Schema

##### <span id="get-addon-default"></span> Default Response
errorResponse

###### <span id="get-addon-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="get-addon-config"></span> Returns specified addon config. (*getAddonConfig*)

```
GET /api/v1/addonconfigs/{addon_id}
```

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| addon_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-addon-config-200) | OK | AddonConfig |  | [schema](#get-addon-config-200-schema) |
| [401](#get-addon-config-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#get-addon-config-401-schema) |
| [default](#get-addon-config-default) | | errorResponse |  | [schema](#get-addon-config-default-schema) |

#### Responses


##### <span id="get-addon-config-200"></span> 200 - AddonConfig
Status: OK

###### <span id="get-addon-config-200-schema"></span> Schema
   
  

[AddonConfig](#addon-config)

##### <span id="get-addon-config-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="get-addon-config-401-schema"></span> Schema

##### <span id="get-addon-config-default"></span> Default Response
errorResponse

###### <span id="get-addon-config-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="get-addon-v2"></span> Gets an addon that is assigned to the given cluster. (*getAddonV2*)

```
GET /api/v2/projects/{project_id}/clusters/{cluster_id}/addons/{addon_id}
```

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| addon_id | `path` | string | `string` |  | ✓ |  |  |
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-addon-v2-200) | OK | Addon |  | [schema](#get-addon-v2-200-schema) |
| [401](#get-addon-v2-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#get-addon-v2-401-schema) |
| [403](#get-addon-v2-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#get-addon-v2-403-schema) |
| [default](#get-addon-v2-default) | | errorResponse |  | [schema](#get-addon-v2-default-schema) |

#### Responses


##### <span id="get-addon-v2-200"></span> 200 - Addon
Status: OK

###### <span id="get-addon-v2-200-schema"></span> Schema
   
  

[Addon](#addon)

##### <span id="get-addon-v2-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="get-addon-v2-401-schema"></span> Schema

##### <span id="get-addon-v2-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="get-addon-v2-403-schema"></span> Schema

##### <span id="get-addon-v2-default"></span> Default Response
errorResponse

###### <span id="get-addon-v2-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="get-admins"></span> Returns list of admin users. (*getAdmins*)

```
GET /api/v1/admin
```

#### Produces
  * application/json

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-admins-200) | OK | Admin |  | [schema](#get-admins-200-schema) |
| [401](#get-admins-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#get-admins-401-schema) |
| [403](#get-admins-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#get-admins-403-schema) |
| [default](#get-admins-default) | | errorResponse |  | [schema](#get-admins-default-schema) |

#### Responses


##### <span id="get-admins-200"></span> 200 - Admin
Status: OK

###### <span id="get-admins-200-schema"></span> Schema
   
  

[][Admin](#admin)

##### <span id="get-admins-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="get-admins-401-schema"></span> Schema

##### <span id="get-admins-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="get-admins-403-schema"></span> Schema

##### <span id="get-admins-default"></span> Default Response
errorResponse

###### <span id="get-admins-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="get-admission-plugin"></span> Gets the admission plugin. (*getAdmissionPlugin*)

```
GET /api/v1/admin/admission/plugins/{name}
```

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| name | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-admission-plugin-200) | OK | AdmissionPlugin |  | [schema](#get-admission-plugin-200-schema) |
| [401](#get-admission-plugin-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#get-admission-plugin-401-schema) |
| [403](#get-admission-plugin-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#get-admission-plugin-403-schema) |
| [default](#get-admission-plugin-default) | | errorResponse |  | [schema](#get-admission-plugin-default-schema) |

#### Responses


##### <span id="get-admission-plugin-200"></span> 200 - AdmissionPlugin
Status: OK

###### <span id="get-admission-plugin-200-schema"></span> Schema
   
  

[AdmissionPlugin](#admission-plugin)

##### <span id="get-admission-plugin-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="get-admission-plugin-401-schema"></span> Schema

##### <span id="get-admission-plugin-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="get-admission-plugin-403-schema"></span> Schema

##### <span id="get-admission-plugin-default"></span> Default Response
errorResponse

###### <span id="get-admission-plugin-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="get-admission-plugins"></span> Returns specified addon config. (*getAdmissionPlugins*)

```
GET /api/v1/admission/plugins/{version}
```

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| version | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-admission-plugins-200) | OK | AdmissionPluginList |  | [schema](#get-admission-plugins-200-schema) |
| [401](#get-admission-plugins-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#get-admission-plugins-401-schema) |
| [default](#get-admission-plugins-default) | | errorResponse |  | [schema](#get-admission-plugins-default-schema) |

#### Responses


##### <span id="get-admission-plugins-200"></span> 200 - AdmissionPluginList
Status: OK

###### <span id="get-admission-plugins-200-schema"></span> Schema
   
  


 [AdmissionPluginList](#admission-plugin-list)

##### <span id="get-admission-plugins-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="get-admission-plugins-401-schema"></span> Schema

##### <span id="get-admission-plugins-default"></span> Default Response
errorResponse

###### <span id="get-admission-plugins-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="get-alertmanager"></span> Gets the alertmanager configuration for the specified cluster. (*getAlertmanager*)

```
GET /api/v2/projects/{project_id}/clusters/{cluster_id}/alertmanager/config
```

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-alertmanager-200) | OK | Alertmanager |  | [schema](#get-alertmanager-200-schema) |
| [401](#get-alertmanager-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#get-alertmanager-401-schema) |
| [403](#get-alertmanager-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#get-alertmanager-403-schema) |
| [default](#get-alertmanager-default) | | errorResponse |  | [schema](#get-alertmanager-default-schema) |

#### Responses


##### <span id="get-alertmanager-200"></span> 200 - Alertmanager
Status: OK

###### <span id="get-alertmanager-200-schema"></span> Schema
   
  

[Alertmanager](#alertmanager)

##### <span id="get-alertmanager-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="get-alertmanager-401-schema"></span> Schema

##### <span id="get-alertmanager-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="get-alertmanager-403-schema"></span> Schema

##### <span id="get-alertmanager-default"></span> Default Response
errorResponse

###### <span id="get-alertmanager-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="get-allowed-registry"></span> get allowed registry (*getAllowedRegistry*)

```
GET /api/v2/allowedregistries/{allowed_registry}
```

Get allowed registries specified by name

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| allowed_registry | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-allowed-registry-200) | OK | AllowedRegistry |  | [schema](#get-allowed-registry-200-schema) |
| [401](#get-allowed-registry-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#get-allowed-registry-401-schema) |
| [403](#get-allowed-registry-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#get-allowed-registry-403-schema) |
| [default](#get-allowed-registry-default) | | errorResponse |  | [schema](#get-allowed-registry-default-schema) |

#### Responses


##### <span id="get-allowed-registry-200"></span> 200 - AllowedRegistry
Status: OK

###### <span id="get-allowed-registry-200-schema"></span> Schema
   
  

[AllowedRegistry](#allowed-registry)

##### <span id="get-allowed-registry-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="get-allowed-registry-401-schema"></span> Schema

##### <span id="get-allowed-registry-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="get-allowed-registry-403-schema"></span> Schema

##### <span id="get-allowed-registry-default"></span> Default Response
errorResponse

###### <span id="get-allowed-registry-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="get-cluster"></span> get cluster (*getCluster*)

```
GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}
```

Gets the cluster with the given name

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| dc | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-cluster-200) | OK | Cluster |  | [schema](#get-cluster-200-schema) |
| [401](#get-cluster-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#get-cluster-401-schema) |
| [403](#get-cluster-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#get-cluster-403-schema) |
| [default](#get-cluster-default) | | errorResponse |  | [schema](#get-cluster-default-schema) |

#### Responses


##### <span id="get-cluster-200"></span> 200 - Cluster
Status: OK

###### <span id="get-cluster-200-schema"></span> Schema
   
  

[Cluster](#cluster)

##### <span id="get-cluster-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="get-cluster-401-schema"></span> Schema

##### <span id="get-cluster-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="get-cluster-403-schema"></span> Schema

##### <span id="get-cluster-default"></span> Default Response
errorResponse

###### <span id="get-cluster-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="get-cluster-events"></span> Gets the events related to the specified cluster. (*getClusterEvents*)

```
GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/events
```

#### Produces
  * application/yaml

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| dc | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |
| type | `query` | string | `string` |  |  |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-cluster-events-200) | OK | Event |  | [schema](#get-cluster-events-200-schema) |
| [401](#get-cluster-events-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#get-cluster-events-401-schema) |
| [403](#get-cluster-events-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#get-cluster-events-403-schema) |
| [default](#get-cluster-events-default) | | errorResponse |  | [schema](#get-cluster-events-default-schema) |

#### Responses


##### <span id="get-cluster-events-200"></span> 200 - Event
Status: OK

###### <span id="get-cluster-events-200-schema"></span> Schema
   
  

[][Event](#event)

##### <span id="get-cluster-events-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="get-cluster-events-401-schema"></span> Schema

##### <span id="get-cluster-events-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="get-cluster-events-403-schema"></span> Schema

##### <span id="get-cluster-events-default"></span> Default Response
errorResponse

###### <span id="get-cluster-events-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="get-cluster-events-v2"></span> Gets the events related to the specified cluster. (*getClusterEventsV2*)

```
GET /api/v2/projects/{project_id}/clusters/{cluster_id}/events
```

#### Produces
  * application/yaml

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |
| type | `query` | string | `string` |  |  |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-cluster-events-v2-200) | OK | Event |  | [schema](#get-cluster-events-v2-200-schema) |
| [401](#get-cluster-events-v2-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#get-cluster-events-v2-401-schema) |
| [403](#get-cluster-events-v2-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#get-cluster-events-v2-403-schema) |
| [default](#get-cluster-events-v2-default) | | errorResponse |  | [schema](#get-cluster-events-v2-default-schema) |

#### Responses


##### <span id="get-cluster-events-v2-200"></span> 200 - Event
Status: OK

###### <span id="get-cluster-events-v2-200-schema"></span> Schema
   
  

[][Event](#event)

##### <span id="get-cluster-events-v2-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="get-cluster-events-v2-401-schema"></span> Schema

##### <span id="get-cluster-events-v2-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="get-cluster-events-v2-403-schema"></span> Schema

##### <span id="get-cluster-events-v2-default"></span> Default Response
errorResponse

###### <span id="get-cluster-events-v2-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="get-cluster-health"></span> get cluster health (*getClusterHealth*)

```
GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/health
```

Returns the cluster's component health status

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| dc | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-cluster-health-200) | OK | ClusterHealth |  | [schema](#get-cluster-health-200-schema) |
| [401](#get-cluster-health-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#get-cluster-health-401-schema) |
| [403](#get-cluster-health-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#get-cluster-health-403-schema) |
| [default](#get-cluster-health-default) | | errorResponse |  | [schema](#get-cluster-health-default-schema) |

#### Responses


##### <span id="get-cluster-health-200"></span> 200 - ClusterHealth
Status: OK

###### <span id="get-cluster-health-200-schema"></span> Schema
   
  

[ClusterHealth](#cluster-health)

##### <span id="get-cluster-health-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="get-cluster-health-401-schema"></span> Schema

##### <span id="get-cluster-health-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="get-cluster-health-403-schema"></span> Schema

##### <span id="get-cluster-health-default"></span> Default Response
errorResponse

###### <span id="get-cluster-health-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="get-cluster-health-v2"></span> get cluster health v2 (*getClusterHealthV2*)

```
GET /api/v2/projects/{project_id}/clusters/{cluster_id}/health
```

Returns the cluster's component health status

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-cluster-health-v2-200) | OK | ClusterHealth |  | [schema](#get-cluster-health-v2-200-schema) |
| [401](#get-cluster-health-v2-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#get-cluster-health-v2-401-schema) |
| [403](#get-cluster-health-v2-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#get-cluster-health-v2-403-schema) |
| [default](#get-cluster-health-v2-default) | | errorResponse |  | [schema](#get-cluster-health-v2-default-schema) |

#### Responses


##### <span id="get-cluster-health-v2-200"></span> 200 - ClusterHealth
Status: OK

###### <span id="get-cluster-health-v2-200-schema"></span> Schema
   
  

[ClusterHealth](#cluster-health)

##### <span id="get-cluster-health-v2-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="get-cluster-health-v2-401-schema"></span> Schema

##### <span id="get-cluster-health-v2-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="get-cluster-health-v2-403-schema"></span> Schema

##### <span id="get-cluster-health-v2-default"></span> Default Response
errorResponse

###### <span id="get-cluster-health-v2-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="get-cluster-kubeconfig"></span> Gets the kubeconfig for the specified cluster. (*getClusterKubeconfig*)

```
GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/kubeconfig
```

#### Produces
  * application/octet-stream

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| dc | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-cluster-kubeconfig-200) | OK | Kubeconfig is a clusters kubeconfig |  | [schema](#get-cluster-kubeconfig-200-schema) |
| [401](#get-cluster-kubeconfig-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#get-cluster-kubeconfig-401-schema) |
| [403](#get-cluster-kubeconfig-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#get-cluster-kubeconfig-403-schema) |
| [default](#get-cluster-kubeconfig-default) | | errorResponse |  | [schema](#get-cluster-kubeconfig-default-schema) |

#### Responses


##### <span id="get-cluster-kubeconfig-200"></span> 200 - Kubeconfig is a clusters kubeconfig
Status: OK

###### <span id="get-cluster-kubeconfig-200-schema"></span> Schema
   
  

[]uint8 (formatted integer)

##### <span id="get-cluster-kubeconfig-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="get-cluster-kubeconfig-401-schema"></span> Schema

##### <span id="get-cluster-kubeconfig-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="get-cluster-kubeconfig-403-schema"></span> Schema

##### <span id="get-cluster-kubeconfig-default"></span> Default Response
errorResponse

###### <span id="get-cluster-kubeconfig-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="get-cluster-kubeconfig-v2"></span> Gets the kubeconfig for the specified cluster. (*getClusterKubeconfigV2*)

```
GET /api/v2/projects/{project_id}/clusters/{cluster_id}/kubeconfig
```

#### Produces
  * application/octet-stream

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-cluster-kubeconfig-v2-200) | OK | Kubeconfig is a clusters kubeconfig |  | [schema](#get-cluster-kubeconfig-v2-200-schema) |
| [401](#get-cluster-kubeconfig-v2-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#get-cluster-kubeconfig-v2-401-schema) |
| [403](#get-cluster-kubeconfig-v2-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#get-cluster-kubeconfig-v2-403-schema) |
| [default](#get-cluster-kubeconfig-v2-default) | | errorResponse |  | [schema](#get-cluster-kubeconfig-v2-default-schema) |

#### Responses


##### <span id="get-cluster-kubeconfig-v2-200"></span> 200 - Kubeconfig is a clusters kubeconfig
Status: OK

###### <span id="get-cluster-kubeconfig-v2-200-schema"></span> Schema
   
  

[]uint8 (formatted integer)

##### <span id="get-cluster-kubeconfig-v2-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="get-cluster-kubeconfig-v2-401-schema"></span> Schema

##### <span id="get-cluster-kubeconfig-v2-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="get-cluster-kubeconfig-v2-403-schema"></span> Schema

##### <span id="get-cluster-kubeconfig-v2-default"></span> Default Response
errorResponse

###### <span id="get-cluster-kubeconfig-v2-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="get-cluster-metrics"></span> get cluster metrics (*getClusterMetrics*)

```
GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/metrics
```

Gets cluster metrics

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| dc | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-cluster-metrics-200) | OK | ClusterMetrics |  | [schema](#get-cluster-metrics-200-schema) |
| [401](#get-cluster-metrics-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#get-cluster-metrics-401-schema) |
| [403](#get-cluster-metrics-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#get-cluster-metrics-403-schema) |
| [default](#get-cluster-metrics-default) | | errorResponse |  | [schema](#get-cluster-metrics-default-schema) |

#### Responses


##### <span id="get-cluster-metrics-200"></span> 200 - ClusterMetrics
Status: OK

###### <span id="get-cluster-metrics-200-schema"></span> Schema
   
  

[ClusterMetrics](#cluster-metrics)

##### <span id="get-cluster-metrics-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="get-cluster-metrics-401-schema"></span> Schema

##### <span id="get-cluster-metrics-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="get-cluster-metrics-403-schema"></span> Schema

##### <span id="get-cluster-metrics-default"></span> Default Response
errorResponse

###### <span id="get-cluster-metrics-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="get-cluster-metrics-v2"></span> get cluster metrics v2 (*getClusterMetricsV2*)

```
GET /api/v2/projects/{project_id}/clusters/{cluster_id}/metrics
```

Gets cluster metrics

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-cluster-metrics-v2-200) | OK | ClusterMetrics |  | [schema](#get-cluster-metrics-v2-200-schema) |
| [401](#get-cluster-metrics-v2-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#get-cluster-metrics-v2-401-schema) |
| [403](#get-cluster-metrics-v2-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#get-cluster-metrics-v2-403-schema) |
| [default](#get-cluster-metrics-v2-default) | | errorResponse |  | [schema](#get-cluster-metrics-v2-default-schema) |

#### Responses


##### <span id="get-cluster-metrics-v2-200"></span> 200 - ClusterMetrics
Status: OK

###### <span id="get-cluster-metrics-v2-200-schema"></span> Schema
   
  

[ClusterMetrics](#cluster-metrics)

##### <span id="get-cluster-metrics-v2-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="get-cluster-metrics-v2-401-schema"></span> Schema

##### <span id="get-cluster-metrics-v2-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="get-cluster-metrics-v2-403-schema"></span> Schema

##### <span id="get-cluster-metrics-v2-default"></span> Default Response
errorResponse

###### <span id="get-cluster-metrics-v2-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="get-cluster-oidc"></span> Gets the OIDC params for the specified cluster with OIDC authentication. (*getClusterOidc*)

```
GET /api/v2/projects/{project_id}/clusters/{cluster_id}/oidc
```

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-cluster-oidc-200) | OK | OIDCSpec |  | [schema](#get-cluster-oidc-200-schema) |
| [401](#get-cluster-oidc-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#get-cluster-oidc-401-schema) |
| [403](#get-cluster-oidc-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#get-cluster-oidc-403-schema) |
| [default](#get-cluster-oidc-default) | | errorResponse |  | [schema](#get-cluster-oidc-default-schema) |

#### Responses


##### <span id="get-cluster-oidc-200"></span> 200 - OIDCSpec
Status: OK

###### <span id="get-cluster-oidc-200-schema"></span> Schema
   
  

[OIDCSpec](#o-id-c-spec)

##### <span id="get-cluster-oidc-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="get-cluster-oidc-401-schema"></span> Schema

##### <span id="get-cluster-oidc-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="get-cluster-oidc-403-schema"></span> Schema

##### <span id="get-cluster-oidc-default"></span> Default Response
errorResponse

###### <span id="get-cluster-oidc-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="get-cluster-role"></span> get cluster role (*getClusterRole*)

```
GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/roles/{role_id}
```

Gets the cluster role with the given name

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| dc | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |
| role_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-cluster-role-200) | OK | ClusterRole |  | [schema](#get-cluster-role-200-schema) |
| [401](#get-cluster-role-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#get-cluster-role-401-schema) |
| [403](#get-cluster-role-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#get-cluster-role-403-schema) |
| [default](#get-cluster-role-default) | | errorResponse |  | [schema](#get-cluster-role-default-schema) |

#### Responses


##### <span id="get-cluster-role-200"></span> 200 - ClusterRole
Status: OK

###### <span id="get-cluster-role-200-schema"></span> Schema
   
  

[ClusterRole](#cluster-role)

##### <span id="get-cluster-role-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="get-cluster-role-401-schema"></span> Schema

##### <span id="get-cluster-role-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="get-cluster-role-403-schema"></span> Schema

##### <span id="get-cluster-role-default"></span> Default Response
errorResponse

###### <span id="get-cluster-role-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="get-cluster-template"></span> Get cluster template. (*getClusterTemplate*)

```
GET /api/v2/projects/{project_id}/clustertemplates/{template_id}
```

#### Consumes
  * application/json

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| project_id | `path` | string | `string` |  | ✓ |  |  |
| template_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-cluster-template-200) | OK | ClusterTemplate |  | [schema](#get-cluster-template-200-schema) |
| [401](#get-cluster-template-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#get-cluster-template-401-schema) |
| [403](#get-cluster-template-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#get-cluster-template-403-schema) |
| [default](#get-cluster-template-default) | | errorResponse |  | [schema](#get-cluster-template-default-schema) |

#### Responses


##### <span id="get-cluster-template-200"></span> 200 - ClusterTemplate
Status: OK

###### <span id="get-cluster-template-200-schema"></span> Schema
   
  

[ClusterTemplate](#cluster-template)

##### <span id="get-cluster-template-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="get-cluster-template-401-schema"></span> Schema

##### <span id="get-cluster-template-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="get-cluster-template-403-schema"></span> Schema

##### <span id="get-cluster-template-default"></span> Default Response
errorResponse

###### <span id="get-cluster-template-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="get-cluster-upgrades"></span> get cluster upgrades (*getClusterUpgrades*)

```
GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/upgrades
```

Gets possible cluster upgrades

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| dc | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-cluster-upgrades-200) | OK | MasterVersion |  | [schema](#get-cluster-upgrades-200-schema) |
| [401](#get-cluster-upgrades-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#get-cluster-upgrades-401-schema) |
| [403](#get-cluster-upgrades-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#get-cluster-upgrades-403-schema) |
| [default](#get-cluster-upgrades-default) | | errorResponse |  | [schema](#get-cluster-upgrades-default-schema) |

#### Responses


##### <span id="get-cluster-upgrades-200"></span> 200 - MasterVersion
Status: OK

###### <span id="get-cluster-upgrades-200-schema"></span> Schema
   
  

[][MasterVersion](#master-version)

##### <span id="get-cluster-upgrades-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="get-cluster-upgrades-401-schema"></span> Schema

##### <span id="get-cluster-upgrades-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="get-cluster-upgrades-403-schema"></span> Schema

##### <span id="get-cluster-upgrades-default"></span> Default Response
errorResponse

###### <span id="get-cluster-upgrades-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="get-cluster-upgrades-v2"></span> get cluster upgrades v2 (*getClusterUpgradesV2*)

```
GET /api/v2/projects/{project_id}/clusters/{cluster_id}/upgrades
```

Gets possible cluster upgrades

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-cluster-upgrades-v2-200) | OK | MasterVersion |  | [schema](#get-cluster-upgrades-v2-200-schema) |
| [401](#get-cluster-upgrades-v2-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#get-cluster-upgrades-v2-401-schema) |
| [403](#get-cluster-upgrades-v2-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#get-cluster-upgrades-v2-403-schema) |
| [default](#get-cluster-upgrades-v2-default) | | errorResponse |  | [schema](#get-cluster-upgrades-v2-default-schema) |

#### Responses


##### <span id="get-cluster-upgrades-v2-200"></span> 200 - MasterVersion
Status: OK

###### <span id="get-cluster-upgrades-v2-200-schema"></span> Schema
   
  

[][MasterVersion](#master-version)

##### <span id="get-cluster-upgrades-v2-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="get-cluster-upgrades-v2-401-schema"></span> Schema

##### <span id="get-cluster-upgrades-v2-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="get-cluster-upgrades-v2-403-schema"></span> Schema

##### <span id="get-cluster-upgrades-v2-default"></span> Default Response
errorResponse

###### <span id="get-cluster-upgrades-v2-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="get-cluster-v2"></span> get cluster v2 (*getClusterV2*)

```
GET /api/v2/projects/{project_id}/clusters/{cluster_id}
```

Gets the cluster with the given name

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-cluster-v2-200) | OK | Cluster |  | [schema](#get-cluster-v2-200-schema) |
| [401](#get-cluster-v2-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#get-cluster-v2-401-schema) |
| [403](#get-cluster-v2-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#get-cluster-v2-403-schema) |
| [default](#get-cluster-v2-default) | | errorResponse |  | [schema](#get-cluster-v2-default-schema) |

#### Responses


##### <span id="get-cluster-v2-200"></span> 200 - Cluster
Status: OK

###### <span id="get-cluster-v2-200-schema"></span> Schema
   
  

[Cluster](#cluster)

##### <span id="get-cluster-v2-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="get-cluster-v2-401-schema"></span> Schema

##### <span id="get-cluster-v2-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="get-cluster-v2-403-schema"></span> Schema

##### <span id="get-cluster-v2-default"></span> Default Response
errorResponse

###### <span id="get-cluster-v2-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="get-constraint"></span> Gets an specified constraint for the given cluster. (*getConstraint*)

```
GET /api/v2/projects/{project_id}/clusters/{cluster_id}/constraints/{constraint_name}
```

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| constraint_name | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-constraint-200) | OK | Constraint |  | [schema](#get-constraint-200-schema) |
| [401](#get-constraint-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#get-constraint-401-schema) |
| [403](#get-constraint-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#get-constraint-403-schema) |
| [default](#get-constraint-default) | | errorResponse |  | [schema](#get-constraint-default-schema) |

#### Responses


##### <span id="get-constraint-200"></span> 200 - Constraint
Status: OK

###### <span id="get-constraint-200-schema"></span> Schema
   
  

[Constraint](#constraint)

##### <span id="get-constraint-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="get-constraint-401-schema"></span> Schema

##### <span id="get-constraint-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="get-constraint-403-schema"></span> Schema

##### <span id="get-constraint-default"></span> Default Response
errorResponse

###### <span id="get-constraint-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="get-constraint-template"></span> get constraint template (*getConstraintTemplate*)

```
GET /api/v2/constrainttemplates/{ct_name}
```

Get constraint templates specified by name

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| ct_name | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-constraint-template-200) | OK | ConstraintTemplate |  | [schema](#get-constraint-template-200-schema) |
| [401](#get-constraint-template-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#get-constraint-template-401-schema) |
| [403](#get-constraint-template-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#get-constraint-template-403-schema) |
| [default](#get-constraint-template-default) | | errorResponse |  | [schema](#get-constraint-template-default-schema) |

#### Responses


##### <span id="get-constraint-template-200"></span> 200 - ConstraintTemplate
Status: OK

###### <span id="get-constraint-template-200-schema"></span> Schema
   
  

[ConstraintTemplate](#constraint-template)

##### <span id="get-constraint-template-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="get-constraint-template-401-schema"></span> Schema

##### <span id="get-constraint-template-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="get-constraint-template-403-schema"></span> Schema

##### <span id="get-constraint-template-default"></span> Default Response
errorResponse

###### <span id="get-constraint-template-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="get-current-user"></span> Returns information about the current user. (*getCurrentUser*)

```
GET /api/v1/me
```

#### Produces
  * application/json

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-current-user-200) | OK | User |  | [schema](#get-current-user-200-schema) |
| [401](#get-current-user-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#get-current-user-401-schema) |
| [default](#get-current-user-default) | | errorResponse |  | [schema](#get-current-user-default-schema) |

#### Responses


##### <span id="get-current-user-200"></span> 200 - User
Status: OK

###### <span id="get-current-user-200-schema"></span> Schema
   
  

[User](#user)

##### <span id="get-current-user-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="get-current-user-401-schema"></span> Schema

##### <span id="get-current-user-default"></span> Default Response
errorResponse

###### <span id="get-current-user-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="get-current-user-settings"></span> Returns settings of the current user. (*getCurrentUserSettings*)

```
GET /api/v1/me/settings
```

#### Produces
  * application/json

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-current-user-settings-200) | OK | UserSettings |  | [schema](#get-current-user-settings-200-schema) |
| [401](#get-current-user-settings-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#get-current-user-settings-401-schema) |
| [default](#get-current-user-settings-default) | | errorResponse |  | [schema](#get-current-user-settings-default-schema) |

#### Responses


##### <span id="get-current-user-settings-200"></span> 200 - UserSettings
Status: OK

###### <span id="get-current-user-settings-200-schema"></span> Schema
   
  

[UserSettings](#user-settings)

##### <span id="get-current-user-settings-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="get-current-user-settings-401-schema"></span> Schema

##### <span id="get-current-user-settings-default"></span> Default Response
errorResponse

###### <span id="get-current-user-settings-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="get-d-c-for-provider"></span> Get the datacenter for the specified provider. (*getDCForProvider*)

```
GET /api/v1/providers/{provider_name}/dc/{dc}
```

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| dc | `path` | string | `string` |  | ✓ |  |  |
| provider_name | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-d-c-for-provider-200) | OK | Datacenter |  | [schema](#get-d-c-for-provider-200-schema) |
| [401](#get-d-c-for-provider-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#get-d-c-for-provider-401-schema) |
| [403](#get-d-c-for-provider-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#get-d-c-for-provider-403-schema) |
| [default](#get-d-c-for-provider-default) | | errorResponse |  | [schema](#get-d-c-for-provider-default-schema) |

#### Responses


##### <span id="get-d-c-for-provider-200"></span> 200 - Datacenter
Status: OK

###### <span id="get-d-c-for-provider-200-schema"></span> Schema
   
  

[Datacenter](#datacenter)

##### <span id="get-d-c-for-provider-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="get-d-c-for-provider-401-schema"></span> Schema

##### <span id="get-d-c-for-provider-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="get-d-c-for-provider-403-schema"></span> Schema

##### <span id="get-d-c-for-provider-default"></span> Default Response
errorResponse

###### <span id="get-d-c-for-provider-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="get-d-c-for-seed"></span> Returns the specified datacenter for the specified seed. (*getDCForSeed*)

```
GET /api/v1/seed/{seed_name}/dc/{dc}
```

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| dc | `path` | string | `string` |  | ✓ |  |  |
| seed_name | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-d-c-for-seed-200) | OK | Datacenter |  | [schema](#get-d-c-for-seed-200-schema) |
| [401](#get-d-c-for-seed-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#get-d-c-for-seed-401-schema) |
| [403](#get-d-c-for-seed-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#get-d-c-for-seed-403-schema) |
| [default](#get-d-c-for-seed-default) | | errorResponse |  | [schema](#get-d-c-for-seed-default-schema) |

#### Responses


##### <span id="get-d-c-for-seed-200"></span> 200 - Datacenter
Status: OK

###### <span id="get-d-c-for-seed-200-schema"></span> Schema
   
  

[Datacenter](#datacenter)

##### <span id="get-d-c-for-seed-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="get-d-c-for-seed-401-schema"></span> Schema

##### <span id="get-d-c-for-seed-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="get-d-c-for-seed-403-schema"></span> Schema

##### <span id="get-d-c-for-seed-default"></span> Default Response
errorResponse

###### <span id="get-d-c-for-seed-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="get-datacenter"></span> get datacenter (*getDatacenter*)

```
GET /api/v1/dc/{dc}
```

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| dc | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-datacenter-200) | OK | Datacenter |  | [schema](#get-datacenter-200-schema) |
| [default](#get-datacenter-default) | | errorResponse |  | [schema](#get-datacenter-default-schema) |

#### Responses


##### <span id="get-datacenter-200"></span> 200 - Datacenter
Status: OK

###### <span id="get-datacenter-200-schema"></span> Schema
   
  

[Datacenter](#datacenter)

##### <span id="get-datacenter-default"></span> Default Response
errorResponse

###### <span id="get-datacenter-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="get-default-constraint"></span> get default constraint (*getDefaultConstraint*)

```
GET /api/v2/constraints/{constraint_name}
```

Gets an specified default constraint

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| constraint_name | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-default-constraint-200) | OK | Constraint |  | [schema](#get-default-constraint-200-schema) |
| [401](#get-default-constraint-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#get-default-constraint-401-schema) |
| [403](#get-default-constraint-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#get-default-constraint-403-schema) |
| [default](#get-default-constraint-default) | | errorResponse |  | [schema](#get-default-constraint-default-schema) |

#### Responses


##### <span id="get-default-constraint-200"></span> 200 - Constraint
Status: OK

###### <span id="get-default-constraint-200-schema"></span> Schema
   
  

[Constraint](#constraint)

##### <span id="get-default-constraint-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="get-default-constraint-401-schema"></span> Schema

##### <span id="get-default-constraint-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="get-default-constraint-403-schema"></span> Schema

##### <span id="get-default-constraint-default"></span> Default Response
errorResponse

###### <span id="get-default-constraint-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="get-etcd-backup-config"></span> get etcd backup config (*getEtcdBackupConfig*)

```
GET /api/v2/projects/{project_id}/clusters/{cluster_id}/etcdbackupconfigs/{ebc_id}
```

Gets a etcd backup config for a given cluster based on its id

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| ebc_id | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-etcd-backup-config-200) | OK | EtcdBackupConfig |  | [schema](#get-etcd-backup-config-200-schema) |
| [401](#get-etcd-backup-config-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#get-etcd-backup-config-401-schema) |
| [403](#get-etcd-backup-config-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#get-etcd-backup-config-403-schema) |
| [default](#get-etcd-backup-config-default) | | errorResponse |  | [schema](#get-etcd-backup-config-default-schema) |

#### Responses


##### <span id="get-etcd-backup-config-200"></span> 200 - EtcdBackupConfig
Status: OK

###### <span id="get-etcd-backup-config-200-schema"></span> Schema
   
  

[EtcdBackupConfig](#etcd-backup-config)

##### <span id="get-etcd-backup-config-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="get-etcd-backup-config-401-schema"></span> Schema

##### <span id="get-etcd-backup-config-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="get-etcd-backup-config-403-schema"></span> Schema

##### <span id="get-etcd-backup-config-default"></span> Default Response
errorResponse

###### <span id="get-etcd-backup-config-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="get-etcd-restore"></span> get etcd restore (*getEtcdRestore*)

```
GET /api/v2/projects/{project_id}/clusters/{cluster_id}/etcdrestores/{er_name}
```

Gets a etcd backup restore for a given cluster based on its name

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| er_name | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-etcd-restore-200) | OK | EtcdRestore |  | [schema](#get-etcd-restore-200-schema) |
| [401](#get-etcd-restore-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#get-etcd-restore-401-schema) |
| [403](#get-etcd-restore-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#get-etcd-restore-403-schema) |
| [default](#get-etcd-restore-default) | | errorResponse |  | [schema](#get-etcd-restore-default-schema) |

#### Responses


##### <span id="get-etcd-restore-200"></span> 200 - EtcdRestore
Status: OK

###### <span id="get-etcd-restore-200-schema"></span> Schema
   
  

[EtcdRestore](#etcd-restore)

##### <span id="get-etcd-restore-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="get-etcd-restore-401-schema"></span> Schema

##### <span id="get-etcd-restore-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="get-etcd-restore-403-schema"></span> Schema

##### <span id="get-etcd-restore-default"></span> Default Response
errorResponse

###### <span id="get-etcd-restore-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="get-external-cluster"></span> Gets an external cluster for the given project. (*getExternalCluster*)

```
GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}
```

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-external-cluster-200) | OK | Cluster |  | [schema](#get-external-cluster-200-schema) |
| [401](#get-external-cluster-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#get-external-cluster-401-schema) |
| [403](#get-external-cluster-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#get-external-cluster-403-schema) |
| [default](#get-external-cluster-default) | | errorResponse |  | [schema](#get-external-cluster-default-schema) |

#### Responses


##### <span id="get-external-cluster-200"></span> 200 - Cluster
Status: OK

###### <span id="get-external-cluster-200-schema"></span> Schema
   
  

[Cluster](#cluster)

##### <span id="get-external-cluster-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="get-external-cluster-401-schema"></span> Schema

##### <span id="get-external-cluster-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="get-external-cluster-403-schema"></span> Schema

##### <span id="get-external-cluster-default"></span> Default Response
errorResponse

###### <span id="get-external-cluster-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="get-external-cluster-metrics"></span> get external cluster metrics (*getExternalClusterMetrics*)

```
GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/metrics
```

Gets cluster metrics

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-external-cluster-metrics-200) | OK | ClusterMetrics |  | [schema](#get-external-cluster-metrics-200-schema) |
| [401](#get-external-cluster-metrics-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#get-external-cluster-metrics-401-schema) |
| [403](#get-external-cluster-metrics-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#get-external-cluster-metrics-403-schema) |
| [default](#get-external-cluster-metrics-default) | | errorResponse |  | [schema](#get-external-cluster-metrics-default-schema) |

#### Responses


##### <span id="get-external-cluster-metrics-200"></span> 200 - ClusterMetrics
Status: OK

###### <span id="get-external-cluster-metrics-200-schema"></span> Schema
   
  

[ClusterMetrics](#cluster-metrics)

##### <span id="get-external-cluster-metrics-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="get-external-cluster-metrics-401-schema"></span> Schema

##### <span id="get-external-cluster-metrics-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="get-external-cluster-metrics-403-schema"></span> Schema

##### <span id="get-external-cluster-metrics-default"></span> Default Response
errorResponse

###### <span id="get-external-cluster-metrics-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="get-external-cluster-node"></span> Gets an external cluster node. (*getExternalClusterNode*)

```
GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/nodes/{node_id}
```

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| node_id | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-external-cluster-node-200) | OK | Node |  | [schema](#get-external-cluster-node-200-schema) |
| [401](#get-external-cluster-node-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#get-external-cluster-node-401-schema) |
| [403](#get-external-cluster-node-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#get-external-cluster-node-403-schema) |
| [default](#get-external-cluster-node-default) | | errorResponse |  | [schema](#get-external-cluster-node-default-schema) |

#### Responses


##### <span id="get-external-cluster-node-200"></span> 200 - Node
Status: OK

###### <span id="get-external-cluster-node-200-schema"></span> Schema
   
  

[Node](#node)

##### <span id="get-external-cluster-node-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="get-external-cluster-node-401-schema"></span> Schema

##### <span id="get-external-cluster-node-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="get-external-cluster-node-403-schema"></span> Schema

##### <span id="get-external-cluster-node-default"></span> Default Response
errorResponse

###### <span id="get-external-cluster-node-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="get-gatekeeper-config"></span> Gets the gatekeeper sync config for the specified cluster. (*getGatekeeperConfig*)

```
GET /api/v2/projects/{project_id}/clusters/{cluster_id}/gatekeeper/config
```

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-gatekeeper-config-200) | OK | GatekeeperConfig |  | [schema](#get-gatekeeper-config-200-schema) |
| [401](#get-gatekeeper-config-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#get-gatekeeper-config-401-schema) |
| [403](#get-gatekeeper-config-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#get-gatekeeper-config-403-schema) |
| [default](#get-gatekeeper-config-default) | | errorResponse |  | [schema](#get-gatekeeper-config-default-schema) |

#### Responses


##### <span id="get-gatekeeper-config-200"></span> 200 - GatekeeperConfig
Status: OK

###### <span id="get-gatekeeper-config-200-schema"></span> Schema
   
  

[GatekeeperConfig](#gatekeeper-config)

##### <span id="get-gatekeeper-config-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="get-gatekeeper-config-401-schema"></span> Schema

##### <span id="get-gatekeeper-config-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="get-gatekeeper-config-403-schema"></span> Schema

##### <span id="get-gatekeeper-config-default"></span> Default Response
errorResponse

###### <span id="get-gatekeeper-config-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="get-kube-login-cluster-kubeconfig"></span> Gets the kubeconfig for the specified cluster with oidc authentication that works nicely with kube-login. (*getKubeLoginClusterKubeconfig*)

```
GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/kubeloginkubeconfig
```

#### Produces
  * application/yaml

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| dc | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-kube-login-cluster-kubeconfig-200) | OK | Kubeconfig is a clusters kubeconfig |  | [schema](#get-kube-login-cluster-kubeconfig-200-schema) |
| [401](#get-kube-login-cluster-kubeconfig-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#get-kube-login-cluster-kubeconfig-401-schema) |
| [403](#get-kube-login-cluster-kubeconfig-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#get-kube-login-cluster-kubeconfig-403-schema) |
| [default](#get-kube-login-cluster-kubeconfig-default) | | errorResponse |  | [schema](#get-kube-login-cluster-kubeconfig-default-schema) |

#### Responses


##### <span id="get-kube-login-cluster-kubeconfig-200"></span> 200 - Kubeconfig is a clusters kubeconfig
Status: OK

###### <span id="get-kube-login-cluster-kubeconfig-200-schema"></span> Schema
   
  

[]uint8 (formatted integer)

##### <span id="get-kube-login-cluster-kubeconfig-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="get-kube-login-cluster-kubeconfig-401-schema"></span> Schema

##### <span id="get-kube-login-cluster-kubeconfig-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="get-kube-login-cluster-kubeconfig-403-schema"></span> Schema

##### <span id="get-kube-login-cluster-kubeconfig-default"></span> Default Response
errorResponse

###### <span id="get-kube-login-cluster-kubeconfig-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="get-kubermatic-custom-links"></span> Gets the custom links. (*getKubermaticCustomLinks*)

```
GET /api/v1/admin/settings/customlinks
```

#### Produces
  * application/json

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-kubermatic-custom-links-200) | OK | GlobalCustomLinks |  | [schema](#get-kubermatic-custom-links-200-schema) |
| [401](#get-kubermatic-custom-links-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#get-kubermatic-custom-links-401-schema) |
| [403](#get-kubermatic-custom-links-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#get-kubermatic-custom-links-403-schema) |
| [default](#get-kubermatic-custom-links-default) | | errorResponse |  | [schema](#get-kubermatic-custom-links-default-schema) |

#### Responses


##### <span id="get-kubermatic-custom-links-200"></span> 200 - GlobalCustomLinks
Status: OK

###### <span id="get-kubermatic-custom-links-200-schema"></span> Schema
   
  


 [GlobalCustomLinks](#global-custom-links)

##### <span id="get-kubermatic-custom-links-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="get-kubermatic-custom-links-401-schema"></span> Schema

##### <span id="get-kubermatic-custom-links-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="get-kubermatic-custom-links-403-schema"></span> Schema

##### <span id="get-kubermatic-custom-links-default"></span> Default Response
errorResponse

###### <span id="get-kubermatic-custom-links-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="get-kubermatic-settings"></span> Gets the global settings. (*getKubermaticSettings*)

```
GET /api/v1/admin/settings
```

#### Produces
  * application/json

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-kubermatic-settings-200) | OK | GlobalSettings |  | [schema](#get-kubermatic-settings-200-schema) |
| [401](#get-kubermatic-settings-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#get-kubermatic-settings-401-schema) |
| [403](#get-kubermatic-settings-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#get-kubermatic-settings-403-schema) |
| [default](#get-kubermatic-settings-default) | | errorResponse |  | [schema](#get-kubermatic-settings-default-schema) |

#### Responses


##### <span id="get-kubermatic-settings-200"></span> 200 - GlobalSettings
Status: OK

###### <span id="get-kubermatic-settings-200-schema"></span> Schema
   
  

[GlobalSettings](#global-settings)

##### <span id="get-kubermatic-settings-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="get-kubermatic-settings-401-schema"></span> Schema

##### <span id="get-kubermatic-settings-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="get-kubermatic-settings-403-schema"></span> Schema

##### <span id="get-kubermatic-settings-default"></span> Default Response
errorResponse

###### <span id="get-kubermatic-settings-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="get-m-l-a-admin-setting"></span> Gets MLA Admin settings for the given cluster. (*getMLAAdminSetting*)

```
GET /api/v2/projects/{project_id}/clusters/{cluster_id}/mlaadminsetting
```

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-m-l-a-admin-setting-200) | OK | MLAAdminSetting |  | [schema](#get-m-l-a-admin-setting-200-schema) |
| [401](#get-m-l-a-admin-setting-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#get-m-l-a-admin-setting-401-schema) |
| [403](#get-m-l-a-admin-setting-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#get-m-l-a-admin-setting-403-schema) |
| [default](#get-m-l-a-admin-setting-default) | | errorResponse |  | [schema](#get-m-l-a-admin-setting-default-schema) |

#### Responses


##### <span id="get-m-l-a-admin-setting-200"></span> 200 - MLAAdminSetting
Status: OK

###### <span id="get-m-l-a-admin-setting-200-schema"></span> Schema
   
  

[MLAAdminSetting](#m-l-a-admin-setting)

##### <span id="get-m-l-a-admin-setting-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="get-m-l-a-admin-setting-401-schema"></span> Schema

##### <span id="get-m-l-a-admin-setting-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="get-m-l-a-admin-setting-403-schema"></span> Schema

##### <span id="get-m-l-a-admin-setting-default"></span> Default Response
errorResponse

###### <span id="get-m-l-a-admin-setting-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="get-machine-deployment"></span> Gets a machine deployment that is assigned to the given cluster. (*getMachineDeployment*)

```
GET /api/v2/projects/{project_id}/clusters/{cluster_id}/machinedeployments/{machinedeployment_id}
```

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| machinedeployment_id | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-machine-deployment-200) | OK | NodeDeployment |  | [schema](#get-machine-deployment-200-schema) |
| [401](#get-machine-deployment-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#get-machine-deployment-401-schema) |
| [403](#get-machine-deployment-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#get-machine-deployment-403-schema) |
| [default](#get-machine-deployment-default) | | errorResponse |  | [schema](#get-machine-deployment-default-schema) |

#### Responses


##### <span id="get-machine-deployment-200"></span> 200 - NodeDeployment
Status: OK

###### <span id="get-machine-deployment-200-schema"></span> Schema
   
  

[NodeDeployment](#node-deployment)

##### <span id="get-machine-deployment-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="get-machine-deployment-401-schema"></span> Schema

##### <span id="get-machine-deployment-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="get-machine-deployment-403-schema"></span> Schema

##### <span id="get-machine-deployment-default"></span> Default Response
errorResponse

###### <span id="get-machine-deployment-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="get-master-versions"></span> get master versions (*getMasterVersions*)

```
GET /api/v1/upgrades/cluster
```

Lists all versions which don't result in automatic updates

#### Produces
  * application/json

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-master-versions-200) | OK | MasterVersion |  | [schema](#get-master-versions-200-schema) |
| [default](#get-master-versions-default) | | errorResponse |  | [schema](#get-master-versions-default-schema) |

#### Responses


##### <span id="get-master-versions-200"></span> 200 - MasterVersion
Status: OK

###### <span id="get-master-versions-200-schema"></span> Schema
   
  

[][MasterVersion](#master-version)

##### <span id="get-master-versions-default"></span> Default Response
errorResponse

###### <span id="get-master-versions-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="get-meta-kube-versions"></span> Get versions of running MetaKube components. (*getMetaKubeVersions*)

```
GET /api/v1/version
```

#### Produces
  * application/json

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-meta-kube-versions-200) | OK | MetaKubeVersions |  | [schema](#get-meta-kube-versions-200-schema) |
| [default](#get-meta-kube-versions-default) | | errorResponse |  | [schema](#get-meta-kube-versions-default-schema) |

#### Responses


##### <span id="get-meta-kube-versions-200"></span> 200 - MetaKubeVersions
Status: OK

###### <span id="get-meta-kube-versions-200-schema"></span> Schema
   
  

[KubermaticVersions](#kubermatic-versions)

##### <span id="get-meta-kube-versions-default"></span> Default Response
errorResponse

###### <span id="get-meta-kube-versions-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="get-metering-report"></span> get metering report (*getMeteringReport*)

```
GET /api/v1/admin/metering/reports/{report_name}
```

Download a specific metering report. Provides an S3 pre signed URL valid for 1 hour. Only available in Kubermatic Enterprise Edition

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| report_name | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-metering-report-200) | OK | MeteringReportURL |  | [schema](#get-metering-report-200-schema) |
| [401](#get-metering-report-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#get-metering-report-401-schema) |
| [403](#get-metering-report-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#get-metering-report-403-schema) |
| [default](#get-metering-report-default) | | errorResponse |  | [schema](#get-metering-report-default-schema) |

#### Responses


##### <span id="get-metering-report-200"></span> 200 - MeteringReportURL
Status: OK

###### <span id="get-metering-report-200-schema"></span> Schema
   
  

| Name | Type | Go type | Default | Description | Example |
|------|------|---------| ------- |-------------|---------|
| getMeteringReportOKBody | string|  | |  |  |



##### <span id="get-metering-report-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="get-metering-report-401-schema"></span> Schema

##### <span id="get-metering-report-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="get-metering-report-403-schema"></span> Schema

##### <span id="get-metering-report-default"></span> Default Response
errorResponse

###### <span id="get-metering-report-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="get-node-deployment"></span> Gets a node deployment that is assigned to the given cluster. (*getNodeDeployment*)

```
GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/nodedeployments/{nodedeployment_id}
```

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| dc | `path` | string | `string` |  | ✓ |  |  |
| nodedeployment_id | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-node-deployment-200) | OK | NodeDeployment |  | [schema](#get-node-deployment-200-schema) |
| [401](#get-node-deployment-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#get-node-deployment-401-schema) |
| [403](#get-node-deployment-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#get-node-deployment-403-schema) |
| [default](#get-node-deployment-default) | | errorResponse |  | [schema](#get-node-deployment-default-schema) |

#### Responses


##### <span id="get-node-deployment-200"></span> 200 - NodeDeployment
Status: OK

###### <span id="get-node-deployment-200-schema"></span> Schema
   
  

[NodeDeployment](#node-deployment)

##### <span id="get-node-deployment-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="get-node-deployment-401-schema"></span> Schema

##### <span id="get-node-deployment-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="get-node-deployment-403-schema"></span> Schema

##### <span id="get-node-deployment-default"></span> Default Response
errorResponse

###### <span id="get-node-deployment-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="get-node-deployment-request"></span> Gets a NodeDeploymentRequest that is assigned to the given cluster. (*getNodeDeploymentRequest*)

```
GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/ndrequests/{ndrequest_id}
```

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| dc | `path` | string | `string` |  | ✓ |  |  |
| ndrequest_id | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-node-deployment-request-200) | OK | NodeDeploymentRequest |  | [schema](#get-node-deployment-request-200-schema) |
| [401](#get-node-deployment-request-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#get-node-deployment-request-401-schema) |
| [403](#get-node-deployment-request-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#get-node-deployment-request-403-schema) |
| [default](#get-node-deployment-request-default) | | errorResponse |  | [schema](#get-node-deployment-request-default-schema) |

#### Responses


##### <span id="get-node-deployment-request-200"></span> 200 - NodeDeploymentRequest
Status: OK

###### <span id="get-node-deployment-request-200-schema"></span> Schema
   
  

[NodeDeploymentRequest](#node-deployment-request)

##### <span id="get-node-deployment-request-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="get-node-deployment-request-401-schema"></span> Schema

##### <span id="get-node-deployment-request-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="get-node-deployment-request-403-schema"></span> Schema

##### <span id="get-node-deployment-request-default"></span> Default Response
errorResponse

###### <span id="get-node-deployment-request-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="get-node-upgrades"></span> get node upgrades (*getNodeUpgrades*)

```
GET /api/v1/upgrades/node
```

Gets possible node upgrades for a specific control plane version

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| control_plane_version | `query` | string | `string` |  |  |  |  |
| type | `query` | string | `string` |  |  |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-node-upgrades-200) | OK | MasterVersion |  | [schema](#get-node-upgrades-200-schema) |
| [401](#get-node-upgrades-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#get-node-upgrades-401-schema) |
| [403](#get-node-upgrades-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#get-node-upgrades-403-schema) |
| [default](#get-node-upgrades-default) | | errorResponse |  | [schema](#get-node-upgrades-default-schema) |

#### Responses


##### <span id="get-node-upgrades-200"></span> 200 - MasterVersion
Status: OK

###### <span id="get-node-upgrades-200-schema"></span> Schema
   
  

[][MasterVersion](#master-version)

##### <span id="get-node-upgrades-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="get-node-upgrades-401-schema"></span> Schema

##### <span id="get-node-upgrades-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="get-node-upgrades-403-schema"></span> Schema

##### <span id="get-node-upgrades-default"></span> Default Response
errorResponse

###### <span id="get-node-upgrades-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="get-oidc-cluster-kubeconfig"></span> Gets the kubeconfig for the specified cluster with oidc authentication. (*getOidcClusterKubeconfig*)

```
GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/oidckubeconfig
```

#### Produces
  * application/yaml

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| dc | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-oidc-cluster-kubeconfig-200) | OK | Kubeconfig is a clusters kubeconfig |  | [schema](#get-oidc-cluster-kubeconfig-200-schema) |
| [401](#get-oidc-cluster-kubeconfig-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#get-oidc-cluster-kubeconfig-401-schema) |
| [403](#get-oidc-cluster-kubeconfig-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#get-oidc-cluster-kubeconfig-403-schema) |
| [default](#get-oidc-cluster-kubeconfig-default) | | errorResponse |  | [schema](#get-oidc-cluster-kubeconfig-default-schema) |

#### Responses


##### <span id="get-oidc-cluster-kubeconfig-200"></span> 200 - Kubeconfig is a clusters kubeconfig
Status: OK

###### <span id="get-oidc-cluster-kubeconfig-200-schema"></span> Schema
   
  

[]uint8 (formatted integer)

##### <span id="get-oidc-cluster-kubeconfig-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="get-oidc-cluster-kubeconfig-401-schema"></span> Schema

##### <span id="get-oidc-cluster-kubeconfig-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="get-oidc-cluster-kubeconfig-403-schema"></span> Schema

##### <span id="get-oidc-cluster-kubeconfig-default"></span> Default Response
errorResponse

###### <span id="get-oidc-cluster-kubeconfig-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="get-oidc-cluster-kubeconfig-v2"></span> Gets the kubeconfig for the specified cluster with oidc authentication. (*getOidcClusterKubeconfigV2*)

```
GET /api/v2/projects/{project_id}/clusters/{cluster_id}/oidckubeconfig
```

#### Produces
  * application/octet-stream

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-oidc-cluster-kubeconfig-v2-200) | OK | Kubeconfig is a clusters kubeconfig |  | [schema](#get-oidc-cluster-kubeconfig-v2-200-schema) |
| [401](#get-oidc-cluster-kubeconfig-v2-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#get-oidc-cluster-kubeconfig-v2-401-schema) |
| [403](#get-oidc-cluster-kubeconfig-v2-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#get-oidc-cluster-kubeconfig-v2-403-schema) |
| [default](#get-oidc-cluster-kubeconfig-v2-default) | | errorResponse |  | [schema](#get-oidc-cluster-kubeconfig-v2-default-schema) |

#### Responses


##### <span id="get-oidc-cluster-kubeconfig-v2-200"></span> 200 - Kubeconfig is a clusters kubeconfig
Status: OK

###### <span id="get-oidc-cluster-kubeconfig-v2-200-schema"></span> Schema
   
  

[]uint8 (formatted integer)

##### <span id="get-oidc-cluster-kubeconfig-v2-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="get-oidc-cluster-kubeconfig-v2-401-schema"></span> Schema

##### <span id="get-oidc-cluster-kubeconfig-v2-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="get-oidc-cluster-kubeconfig-v2-403-schema"></span> Schema

##### <span id="get-oidc-cluster-kubeconfig-v2-default"></span> Default Response
errorResponse

###### <span id="get-oidc-cluster-kubeconfig-v2-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="get-project"></span> get project (*getProject*)

```
GET /api/v1/projects/{project_id}
```

Gets the project with the given ID

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-project-200) | OK | Project |  | [schema](#get-project-200-schema) |
| [401](#get-project-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#get-project-401-schema) |
| [409](#get-project-409) | Conflict | EmptyResponse is a empty response |  | [schema](#get-project-409-schema) |
| [default](#get-project-default) | | errorResponse |  | [schema](#get-project-default-schema) |

#### Responses


##### <span id="get-project-200"></span> 200 - Project
Status: OK

###### <span id="get-project-200-schema"></span> Schema
   
  

[Project](#project)

##### <span id="get-project-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="get-project-401-schema"></span> Schema

##### <span id="get-project-409"></span> 409 - EmptyResponse is a empty response
Status: Conflict

###### <span id="get-project-409-schema"></span> Schema

##### <span id="get-project-default"></span> Default Response
errorResponse

###### <span id="get-project-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="get-role"></span> get role (*getRole*)

```
GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/roles/{namespace}/{role_id}
```

Gets the role with the given name

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| dc | `path` | string | `string` |  | ✓ |  |  |
| namespace | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |
| role_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-role-200) | OK | Role |  | [schema](#get-role-200-schema) |
| [401](#get-role-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#get-role-401-schema) |
| [403](#get-role-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#get-role-403-schema) |
| [default](#get-role-default) | | errorResponse |  | [schema](#get-role-default-schema) |

#### Responses


##### <span id="get-role-200"></span> 200 - Role
Status: OK

###### <span id="get-role-200-schema"></span> Schema
   
  

[Role](#role)

##### <span id="get-role-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="get-role-401-schema"></span> Schema

##### <span id="get-role-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="get-role-403-schema"></span> Schema

##### <span id="get-role-default"></span> Default Response
errorResponse

###### <span id="get-role-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="get-rule-group"></span> Gets a specified rule group for the given cluster. (*getRuleGroup*)

```
GET /api/v2/projects/{project_id}/clusters/{cluster_id}/rulegroups/{rulegroup_id}
```

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |
| rulegroup_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-rule-group-200) | OK | RuleGroup |  | [schema](#get-rule-group-200-schema) |
| [401](#get-rule-group-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#get-rule-group-401-schema) |
| [403](#get-rule-group-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#get-rule-group-403-schema) |
| [default](#get-rule-group-default) | | errorResponse |  | [schema](#get-rule-group-default-schema) |

#### Responses


##### <span id="get-rule-group-200"></span> 200 - RuleGroup
Status: OK

###### <span id="get-rule-group-200-schema"></span> Schema
   
  

[RuleGroup](#rule-group)

##### <span id="get-rule-group-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="get-rule-group-401-schema"></span> Schema

##### <span id="get-rule-group-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="get-rule-group-403-schema"></span> Schema

##### <span id="get-rule-group-default"></span> Default Response
errorResponse

###### <span id="get-rule-group-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="get-seed"></span> Returns the seed object. (*getSeed*)

```
GET /api/v1/admin/seeds/{seed_name}
```

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| seed_name | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-seed-200) | OK | Seed |  | [schema](#get-seed-200-schema) |
| [401](#get-seed-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#get-seed-401-schema) |
| [403](#get-seed-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#get-seed-403-schema) |
| [default](#get-seed-default) | | errorResponse |  | [schema](#get-seed-default-schema) |

#### Responses


##### <span id="get-seed-200"></span> 200 - Seed
Status: OK

###### <span id="get-seed-200-schema"></span> Schema
   
  

[Seed](#seed)

##### <span id="get-seed-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="get-seed-401-schema"></span> Schema

##### <span id="get-seed-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="get-seed-403-schema"></span> Schema

##### <span id="get-seed-default"></span> Default Response
errorResponse

###### <span id="get-seed-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="get-seed-settings"></span> Gets the seed settings. (*getSeedSettings*)

```
GET /api/v2/seeds/{seed_name}/settings
```

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| seed_name | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-seed-settings-200) | OK | SeedSettings |  | [schema](#get-seed-settings-200-schema) |
| [401](#get-seed-settings-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#get-seed-settings-401-schema) |
| [403](#get-seed-settings-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#get-seed-settings-403-schema) |
| [default](#get-seed-settings-default) | | errorResponse |  | [schema](#get-seed-settings-default-schema) |

#### Responses


##### <span id="get-seed-settings-200"></span> 200 - SeedSettings
Status: OK

###### <span id="get-seed-settings-200-schema"></span> Schema
   
  

[SeedSettings](#seed-settings)

##### <span id="get-seed-settings-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="get-seed-settings-401-schema"></span> Schema

##### <span id="get-seed-settings-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="get-seed-settings-403-schema"></span> Schema

##### <span id="get-seed-settings-default"></span> Default Response
errorResponse

###### <span id="get-seed-settings-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="get-users-for-project"></span> get users for project (*getUsersForProject*)

```
GET /api/v1/projects/{project_id}/users
```

Get list of users for the given project

#### Consumes
  * application/json

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-users-for-project-200) | OK | User |  | [schema](#get-users-for-project-200-schema) |
| [401](#get-users-for-project-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#get-users-for-project-401-schema) |
| [403](#get-users-for-project-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#get-users-for-project-403-schema) |
| [default](#get-users-for-project-default) | | errorResponse |  | [schema](#get-users-for-project-default-schema) |

#### Responses


##### <span id="get-users-for-project-200"></span> 200 - User
Status: OK

###### <span id="get-users-for-project-200-schema"></span> Schema
   
  

[][User](#user)

##### <span id="get-users-for-project-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="get-users-for-project-401-schema"></span> Schema

##### <span id="get-users-for-project-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="get-users-for-project-403-schema"></span> Schema

##### <span id="get-users-for-project-default"></span> Default Response
errorResponse

###### <span id="get-users-for-project-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-a-w-s-security-groups"></span> list a w s security groups (*listAWSSecurityGroups*)

```
GET /api/v1/providers/aws/{dc}/securitygroups
```

Lists available AWS Security Groups

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| dc | `path` | string | `string` |  | ✓ |  |  |
| AccessKeyID | `header` | string | `string` |  |  |  |  |
| Credential | `header` | string | `string` |  |  |  |  |
| SecretAccessKey | `header` | string | `string` |  |  |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-a-w-s-security-groups-200) | OK | AWSSecurityGroupList |  | [schema](#list-a-w-s-security-groups-200-schema) |
| [default](#list-a-w-s-security-groups-default) | | errorResponse |  | [schema](#list-a-w-s-security-groups-default-schema) |

#### Responses


##### <span id="list-a-w-s-security-groups-200"></span> 200 - AWSSecurityGroupList
Status: OK

###### <span id="list-a-w-s-security-groups-200-schema"></span> Schema
   
  

[AWSSecurityGroupList](#a-w-s-security-group-list)

##### <span id="list-a-w-s-security-groups-default"></span> Default Response
errorResponse

###### <span id="list-a-w-s-security-groups-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-a-w-s-sizes"></span> Lists available AWS sizes. (*listAWSSizes*)

```
GET /api/v1/providers/aws/sizes
```

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| Region | `header` | string | `string` |  |  |  |  |
| architecture | `query` | string | `string` |  |  |  | architecture query parameter. Supports: arm64 and x64 types. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-a-w-s-sizes-200) | OK | AWSSizeList |  | [schema](#list-a-w-s-sizes-200-schema) |
| [default](#list-a-w-s-sizes-default) | | errorResponse |  | [schema](#list-a-w-s-sizes-default-schema) |

#### Responses


##### <span id="list-a-w-s-sizes-200"></span> 200 - AWSSizeList
Status: OK

###### <span id="list-a-w-s-sizes-200-schema"></span> Schema
   
  


 [AWSSizeList](#a-w-s-size-list)

##### <span id="list-a-w-s-sizes-default"></span> Default Response
errorResponse

###### <span id="list-a-w-s-sizes-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-a-w-s-sizes-no-credentials"></span> list a w s sizes no credentials (*listAWSSizesNoCredentials*)

```
GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/providers/aws/sizes
```

Lists available AWS sizes

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| dc | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-a-w-s-sizes-no-credentials-200) | OK | AWSSizeList |  | [schema](#list-a-w-s-sizes-no-credentials-200-schema) |
| [default](#list-a-w-s-sizes-no-credentials-default) | | errorResponse |  | [schema](#list-a-w-s-sizes-no-credentials-default-schema) |

#### Responses


##### <span id="list-a-w-s-sizes-no-credentials-200"></span> 200 - AWSSizeList
Status: OK

###### <span id="list-a-w-s-sizes-no-credentials-200-schema"></span> Schema
   
  


 [AWSSizeList](#a-w-s-size-list)

##### <span id="list-a-w-s-sizes-no-credentials-default"></span> Default Response
errorResponse

###### <span id="list-a-w-s-sizes-no-credentials-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-a-w-s-sizes-no-credentials-v2"></span> list a w s sizes no credentials v2 (*listAWSSizesNoCredentialsV2*)

```
GET /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/aws/sizes
```

Lists available AWS sizes

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |
| architecture | `query` | string | `string` |  |  |  | architecture query parameter. Supports: arm64 and x64 types. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-a-w-s-sizes-no-credentials-v2-200) | OK | AWSSizeList |  | [schema](#list-a-w-s-sizes-no-credentials-v2-200-schema) |
| [default](#list-a-w-s-sizes-no-credentials-v2-default) | | errorResponse |  | [schema](#list-a-w-s-sizes-no-credentials-v2-default-schema) |

#### Responses


##### <span id="list-a-w-s-sizes-no-credentials-v2-200"></span> 200 - AWSSizeList
Status: OK

###### <span id="list-a-w-s-sizes-no-credentials-v2-200-schema"></span> Schema
   
  


 [AWSSizeList](#a-w-s-size-list)

##### <span id="list-a-w-s-sizes-no-credentials-v2-default"></span> Default Response
errorResponse

###### <span id="list-a-w-s-sizes-no-credentials-v2-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-a-w-s-subnets"></span> list a w s subnets (*listAWSSubnets*)

```
GET /api/v1/providers/aws/{dc}/subnets
```

Lists available AWS subnets

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| dc | `path` | string | `string` |  | ✓ |  |  |
| AccessKeyID | `header` | string | `string` |  |  |  |  |
| Credential | `header` | string | `string` |  |  |  |  |
| SecretAccessKey | `header` | string | `string` |  |  |  |  |
| vpc | `header` | string | `string` |  |  |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-a-w-s-subnets-200) | OK | AWSSubnetList |  | [schema](#list-a-w-s-subnets-200-schema) |
| [default](#list-a-w-s-subnets-default) | | errorResponse |  | [schema](#list-a-w-s-subnets-default-schema) |

#### Responses


##### <span id="list-a-w-s-subnets-200"></span> 200 - AWSSubnetList
Status: OK

###### <span id="list-a-w-s-subnets-200-schema"></span> Schema
   
  


 [AWSSubnetList](#a-w-s-subnet-list)

##### <span id="list-a-w-s-subnets-default"></span> Default Response
errorResponse

###### <span id="list-a-w-s-subnets-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-a-w-s-subnets-no-credentials"></span> list a w s subnets no credentials (*listAWSSubnetsNoCredentials*)

```
GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/providers/aws/subnets
```

Lists available AWS subnets

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| dc | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-a-w-s-subnets-no-credentials-200) | OK | AWSSubnetList |  | [schema](#list-a-w-s-subnets-no-credentials-200-schema) |
| [default](#list-a-w-s-subnets-no-credentials-default) | | errorResponse |  | [schema](#list-a-w-s-subnets-no-credentials-default-schema) |

#### Responses


##### <span id="list-a-w-s-subnets-no-credentials-200"></span> 200 - AWSSubnetList
Status: OK

###### <span id="list-a-w-s-subnets-no-credentials-200-schema"></span> Schema
   
  


 [AWSSubnetList](#a-w-s-subnet-list)

##### <span id="list-a-w-s-subnets-no-credentials-default"></span> Default Response
errorResponse

###### <span id="list-a-w-s-subnets-no-credentials-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-a-w-s-subnets-no-credentials-v2"></span> list a w s subnets no credentials v2 (*listAWSSubnetsNoCredentialsV2*)

```
GET /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/aws/subnets
```

Lists available AWS subnets

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-a-w-s-subnets-no-credentials-v2-200) | OK | AWSSubnetList |  | [schema](#list-a-w-s-subnets-no-credentials-v2-200-schema) |
| [default](#list-a-w-s-subnets-no-credentials-v2-default) | | errorResponse |  | [schema](#list-a-w-s-subnets-no-credentials-v2-default-schema) |

#### Responses


##### <span id="list-a-w-s-subnets-no-credentials-v2-200"></span> 200 - AWSSubnetList
Status: OK

###### <span id="list-a-w-s-subnets-no-credentials-v2-200-schema"></span> Schema
   
  


 [AWSSubnetList](#a-w-s-subnet-list)

##### <span id="list-a-w-s-subnets-no-credentials-v2-default"></span> Default Response
errorResponse

###### <span id="list-a-w-s-subnets-no-credentials-v2-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-a-w-s-v-p-c-s"></span> list a w s v p c s (*listAWSVPCS*)

```
GET /api/v1/providers/aws/{dc}/vpcs
```

Lists available AWS vpc's

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| dc | `path` | string | `string` |  | ✓ |  |  |
| AccessKeyID | `header` | string | `string` |  |  |  |  |
| Credential | `header` | string | `string` |  |  |  |  |
| SecretAccessKey | `header` | string | `string` |  |  |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-a-w-s-v-p-c-s-200) | OK | AWSVPCList |  | [schema](#list-a-w-s-v-p-c-s-200-schema) |
| [default](#list-a-w-s-v-p-c-s-default) | | errorResponse |  | [schema](#list-a-w-s-v-p-c-s-default-schema) |

#### Responses


##### <span id="list-a-w-s-v-p-c-s-200"></span> 200 - AWSVPCList
Status: OK

###### <span id="list-a-w-s-v-p-c-s-200-schema"></span> Schema
   
  


 [AWSVPCList](#a-w-s-v-p-c-list)

##### <span id="list-a-w-s-v-p-c-s-default"></span> Default Response
errorResponse

###### <span id="list-a-w-s-v-p-c-s-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-accessible-addons"></span> list accessible addons (*listAccessibleAddons*)

```
POST /api/v1/addons
```

Lists names of addons that can be configured inside the user clusters

#### Consumes
  * application/json

#### Produces
  * application/json

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-accessible-addons-200) | OK | AccessibleAddons |  | [schema](#list-accessible-addons-200-schema) |
| [401](#list-accessible-addons-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#list-accessible-addons-401-schema) |
| [403](#list-accessible-addons-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#list-accessible-addons-403-schema) |
| [default](#list-accessible-addons-default) | | errorResponse |  | [schema](#list-accessible-addons-default-schema) |

#### Responses


##### <span id="list-accessible-addons-200"></span> 200 - AccessibleAddons
Status: OK

###### <span id="list-accessible-addons-200-schema"></span> Schema
   
  


 [AccessibleAddons](#accessible-addons)

##### <span id="list-accessible-addons-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="list-accessible-addons-401-schema"></span> Schema

##### <span id="list-accessible-addons-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="list-accessible-addons-403-schema"></span> Schema

##### <span id="list-accessible-addons-default"></span> Default Response
errorResponse

###### <span id="list-accessible-addons-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-addon-configs"></span> Returns all available addon configs. (*listAddonConfigs*)

```
GET /api/v1/addonconfigs
```

#### Produces
  * application/json

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-addon-configs-200) | OK | AddonConfig |  | [schema](#list-addon-configs-200-schema) |
| [401](#list-addon-configs-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#list-addon-configs-401-schema) |
| [default](#list-addon-configs-default) | | errorResponse |  | [schema](#list-addon-configs-default-schema) |

#### Responses


##### <span id="list-addon-configs-200"></span> 200 - AddonConfig
Status: OK

###### <span id="list-addon-configs-200-schema"></span> Schema
   
  

[][AddonConfig](#addon-config)

##### <span id="list-addon-configs-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="list-addon-configs-401-schema"></span> Schema

##### <span id="list-addon-configs-default"></span> Default Response
errorResponse

###### <span id="list-addon-configs-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-addons"></span> list addons (*listAddons*)

```
GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/addons
```

Lists addons that belong to the given cluster

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| dc | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-addons-200) | OK | Addon |  | [schema](#list-addons-200-schema) |
| [401](#list-addons-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#list-addons-401-schema) |
| [403](#list-addons-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#list-addons-403-schema) |
| [default](#list-addons-default) | | errorResponse |  | [schema](#list-addons-default-schema) |

#### Responses


##### <span id="list-addons-200"></span> 200 - Addon
Status: OK

###### <span id="list-addons-200-schema"></span> Schema
   
  

[][Addon](#addon)

##### <span id="list-addons-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="list-addons-401-schema"></span> Schema

##### <span id="list-addons-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="list-addons-403-schema"></span> Schema

##### <span id="list-addons-default"></span> Default Response
errorResponse

###### <span id="list-addons-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-addons-v2"></span> list addons v2 (*listAddonsV2*)

```
GET /api/v2/projects/{project_id}/clusters/{cluster_id}/addons
```

Lists addons that belong to the given cluster

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-addons-v2-200) | OK | Addon |  | [schema](#list-addons-v2-200-schema) |
| [401](#list-addons-v2-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#list-addons-v2-401-schema) |
| [403](#list-addons-v2-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#list-addons-v2-403-schema) |
| [default](#list-addons-v2-default) | | errorResponse |  | [schema](#list-addons-v2-default-schema) |

#### Responses


##### <span id="list-addons-v2-200"></span> 200 - Addon
Status: OK

###### <span id="list-addons-v2-200-schema"></span> Schema
   
  

[][Addon](#addon)

##### <span id="list-addons-v2-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="list-addons-v2-401-schema"></span> Schema

##### <span id="list-addons-v2-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="list-addons-v2-403-schema"></span> Schema

##### <span id="list-addons-v2-default"></span> Default Response
errorResponse

###### <span id="list-addons-v2-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-admission-plugins"></span> Returns all admission plugins from the CRDs. (*listAdmissionPlugins*)

```
GET /api/v1/admin/admission/plugins
```

#### Produces
  * application/json

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-admission-plugins-200) | OK | AdmissionPlugin |  | [schema](#list-admission-plugins-200-schema) |
| [401](#list-admission-plugins-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#list-admission-plugins-401-schema) |
| [403](#list-admission-plugins-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#list-admission-plugins-403-schema) |
| [default](#list-admission-plugins-default) | | errorResponse |  | [schema](#list-admission-plugins-default-schema) |

#### Responses


##### <span id="list-admission-plugins-200"></span> 200 - AdmissionPlugin
Status: OK

###### <span id="list-admission-plugins-200-schema"></span> Schema
   
  

[][AdmissionPlugin](#admission-plugin)

##### <span id="list-admission-plugins-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="list-admission-plugins-401-schema"></span> Schema

##### <span id="list-admission-plugins-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="list-admission-plugins-403-schema"></span> Schema

##### <span id="list-admission-plugins-default"></span> Default Response
errorResponse

###### <span id="list-admission-plugins-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-alibaba-instance-types"></span> Lists available Alibaba instance types. (*listAlibabaInstanceTypes*)

```
GET /api/v1/providers/alibaba/instancetypes
```

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| AccessKeyID | `header` | string | `string` |  |  |  |  |
| AccessKeySecret | `header` | string | `string` |  |  |  |  |
| Credential | `header` | string | `string` |  |  |  |  |
| Region | `header` | string | `string` |  |  |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-alibaba-instance-types-200) | OK | AlibabaInstanceTypeList |  | [schema](#list-alibaba-instance-types-200-schema) |
| [default](#list-alibaba-instance-types-default) | | errorResponse |  | [schema](#list-alibaba-instance-types-default-schema) |

#### Responses


##### <span id="list-alibaba-instance-types-200"></span> 200 - AlibabaInstanceTypeList
Status: OK

###### <span id="list-alibaba-instance-types-200-schema"></span> Schema
   
  


 [AlibabaInstanceTypeList](#alibaba-instance-type-list)

##### <span id="list-alibaba-instance-types-default"></span> Default Response
errorResponse

###### <span id="list-alibaba-instance-types-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-alibaba-instance-types-no-credentials"></span> list alibaba instance types no credentials (*listAlibabaInstanceTypesNoCredentials*)

```
GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/providers/alibaba/instancetypes
```

Lists available Alibaba Instance Types

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| dc | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |
| Region | `header` | string | `string` |  |  |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-alibaba-instance-types-no-credentials-200) | OK | AlibabaInstanceTypeList |  | [schema](#list-alibaba-instance-types-no-credentials-200-schema) |
| [default](#list-alibaba-instance-types-no-credentials-default) | | errorResponse |  | [schema](#list-alibaba-instance-types-no-credentials-default-schema) |

#### Responses


##### <span id="list-alibaba-instance-types-no-credentials-200"></span> 200 - AlibabaInstanceTypeList
Status: OK

###### <span id="list-alibaba-instance-types-no-credentials-200-schema"></span> Schema
   
  


 [AlibabaInstanceTypeList](#alibaba-instance-type-list)

##### <span id="list-alibaba-instance-types-no-credentials-default"></span> Default Response
errorResponse

###### <span id="list-alibaba-instance-types-no-credentials-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-alibaba-instance-types-no-credentials-v2"></span> list alibaba instance types no credentials v2 (*listAlibabaInstanceTypesNoCredentialsV2*)

```
GET /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/alibaba/instancetypes
```

Lists available Alibaba Instance Types

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |
| Region | `header` | string | `string` |  |  |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-alibaba-instance-types-no-credentials-v2-200) | OK | AlibabaInstanceTypeList |  | [schema](#list-alibaba-instance-types-no-credentials-v2-200-schema) |
| [default](#list-alibaba-instance-types-no-credentials-v2-default) | | errorResponse |  | [schema](#list-alibaba-instance-types-no-credentials-v2-default-schema) |

#### Responses


##### <span id="list-alibaba-instance-types-no-credentials-v2-200"></span> 200 - AlibabaInstanceTypeList
Status: OK

###### <span id="list-alibaba-instance-types-no-credentials-v2-200-schema"></span> Schema
   
  


 [AlibabaInstanceTypeList](#alibaba-instance-type-list)

##### <span id="list-alibaba-instance-types-no-credentials-v2-default"></span> Default Response
errorResponse

###### <span id="list-alibaba-instance-types-no-credentials-v2-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-alibaba-v-switches"></span> Lists available Alibaba vSwitches. (*listAlibabaVSwitches*)

```
GET /api/v1/providers/alibaba/vswitches
```

#### Produces
  * application/json

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-alibaba-v-switches-200) | OK | AlibabaVSwitchList |  | [schema](#list-alibaba-v-switches-200-schema) |
| [default](#list-alibaba-v-switches-default) | | errorResponse |  | [schema](#list-alibaba-v-switches-default-schema) |

#### Responses


##### <span id="list-alibaba-v-switches-200"></span> 200 - AlibabaVSwitchList
Status: OK

###### <span id="list-alibaba-v-switches-200-schema"></span> Schema
   
  


 [AlibabaVSwitchList](#alibaba-v-switch-list)

##### <span id="list-alibaba-v-switches-default"></span> Default Response
errorResponse

###### <span id="list-alibaba-v-switches-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-alibaba-v-switches-no-credentials-v2"></span> list alibaba v switches no credentials v2 (*listAlibabaVSwitchesNoCredentialsV2*)

```
GET /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/alibaba/vswitches
```

Lists available Alibaba vSwitches

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |
| Region | `header` | string | `string` |  |  |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-alibaba-v-switches-no-credentials-v2-200) | OK | AlibabaVSwitchList |  | [schema](#list-alibaba-v-switches-no-credentials-v2-200-schema) |
| [default](#list-alibaba-v-switches-no-credentials-v2-default) | | errorResponse |  | [schema](#list-alibaba-v-switches-no-credentials-v2-default-schema) |

#### Responses


##### <span id="list-alibaba-v-switches-no-credentials-v2-200"></span> 200 - AlibabaVSwitchList
Status: OK

###### <span id="list-alibaba-v-switches-no-credentials-v2-200-schema"></span> Schema
   
  


 [AlibabaVSwitchList](#alibaba-v-switch-list)

##### <span id="list-alibaba-v-switches-no-credentials-v2-default"></span> Default Response
errorResponse

###### <span id="list-alibaba-v-switches-no-credentials-v2-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-alibaba-zones"></span> Lists available Alibaba zones. (*listAlibabaZones*)

```
GET /api/v1/providers/alibaba/zones
```

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| AccessKeyID | `header` | string | `string` |  |  |  |  |
| AccessKeySecret | `header` | string | `string` |  |  |  |  |
| Credential | `header` | string | `string` |  |  |  |  |
| Region | `header` | string | `string` |  |  |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-alibaba-zones-200) | OK | AlibabaZoneList |  | [schema](#list-alibaba-zones-200-schema) |
| [default](#list-alibaba-zones-default) | | errorResponse |  | [schema](#list-alibaba-zones-default-schema) |

#### Responses


##### <span id="list-alibaba-zones-200"></span> 200 - AlibabaZoneList
Status: OK

###### <span id="list-alibaba-zones-200-schema"></span> Schema
   
  


 [AlibabaZoneList](#alibaba-zone-list)

##### <span id="list-alibaba-zones-default"></span> Default Response
errorResponse

###### <span id="list-alibaba-zones-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-alibaba-zones-no-credentials"></span> list alibaba zones no credentials (*listAlibabaZonesNoCredentials*)

```
GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/providers/alibaba/zones
```

Lists available Alibaba Instance Types

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| dc | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |
| Region | `header` | string | `string` |  |  |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-alibaba-zones-no-credentials-200) | OK | AlibabaZoneList |  | [schema](#list-alibaba-zones-no-credentials-200-schema) |
| [default](#list-alibaba-zones-no-credentials-default) | | errorResponse |  | [schema](#list-alibaba-zones-no-credentials-default-schema) |

#### Responses


##### <span id="list-alibaba-zones-no-credentials-200"></span> 200 - AlibabaZoneList
Status: OK

###### <span id="list-alibaba-zones-no-credentials-200-schema"></span> Schema
   
  


 [AlibabaZoneList](#alibaba-zone-list)

##### <span id="list-alibaba-zones-no-credentials-default"></span> Default Response
errorResponse

###### <span id="list-alibaba-zones-no-credentials-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-alibaba-zones-no-credentials-v2"></span> list alibaba zones no credentials v2 (*listAlibabaZonesNoCredentialsV2*)

```
GET /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/alibaba/zones
```

Lists available Alibaba Instance Types

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |
| Region | `header` | string | `string` |  |  |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-alibaba-zones-no-credentials-v2-200) | OK | AlibabaZoneList |  | [schema](#list-alibaba-zones-no-credentials-v2-200-schema) |
| [default](#list-alibaba-zones-no-credentials-v2-default) | | errorResponse |  | [schema](#list-alibaba-zones-no-credentials-v2-default-schema) |

#### Responses


##### <span id="list-alibaba-zones-no-credentials-v2-200"></span> 200 - AlibabaZoneList
Status: OK

###### <span id="list-alibaba-zones-no-credentials-v2-200-schema"></span> Schema
   
  


 [AlibabaZoneList](#alibaba-zone-list)

##### <span id="list-alibaba-zones-no-credentials-v2-default"></span> Default Response
errorResponse

###### <span id="list-alibaba-zones-no-credentials-v2-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-allowed-registries"></span> List allowed registries. (*listAllowedRegistries*)

```
GET /api/v2/allowedregistries
```

#### Produces
  * application/json

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-allowed-registries-200) | OK | AllowedRegistry |  | [schema](#list-allowed-registries-200-schema) |
| [401](#list-allowed-registries-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#list-allowed-registries-401-schema) |
| [403](#list-allowed-registries-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#list-allowed-registries-403-schema) |
| [default](#list-allowed-registries-default) | | errorResponse |  | [schema](#list-allowed-registries-default-schema) |

#### Responses


##### <span id="list-allowed-registries-200"></span> 200 - AllowedRegistry
Status: OK

###### <span id="list-allowed-registries-200-schema"></span> Schema
   
  

[][AllowedRegistry](#allowed-registry)

##### <span id="list-allowed-registries-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="list-allowed-registries-401-schema"></span> Schema

##### <span id="list-allowed-registries-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="list-allowed-registries-403-schema"></span> Schema

##### <span id="list-allowed-registries-default"></span> Default Response
errorResponse

###### <span id="list-allowed-registries-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-anexia-templates"></span> list anexia templates (*listAnexiaTemplates*)

```
GET /api/v1/providers/anexia/templates
```

Lists templates from anexia

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| Credential | `header` | string | `string` |  |  |  |  |
| Location | `header` | string | `string` |  |  |  |  |
| Token | `header` | string | `string` |  |  |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-anexia-templates-200) | OK | AnexiaTemplateList |  | [schema](#list-anexia-templates-200-schema) |
| [default](#list-anexia-templates-default) | | errorResponse |  | [schema](#list-anexia-templates-default-schema) |

#### Responses


##### <span id="list-anexia-templates-200"></span> 200 - AnexiaTemplateList
Status: OK

###### <span id="list-anexia-templates-200-schema"></span> Schema
   
  


 [AnexiaTemplateList](#anexia-template-list)

##### <span id="list-anexia-templates-default"></span> Default Response
errorResponse

###### <span id="list-anexia-templates-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-anexia-templates-no-credentials-v2"></span> list anexia templates no credentials v2 (*listAnexiaTemplatesNoCredentialsV2*)

```
GET /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/anexia/templates
```

Lists templates from Anexia

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-anexia-templates-no-credentials-v2-200) | OK | AnexiaTemplateList |  | [schema](#list-anexia-templates-no-credentials-v2-200-schema) |
| [default](#list-anexia-templates-no-credentials-v2-default) | | errorResponse |  | [schema](#list-anexia-templates-no-credentials-v2-default-schema) |

#### Responses


##### <span id="list-anexia-templates-no-credentials-v2-200"></span> 200 - AnexiaTemplateList
Status: OK

###### <span id="list-anexia-templates-no-credentials-v2-200-schema"></span> Schema
   
  


 [AnexiaTemplateList](#anexia-template-list)

##### <span id="list-anexia-templates-no-credentials-v2-default"></span> Default Response
errorResponse

###### <span id="list-anexia-templates-no-credentials-v2-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-anexia-vlans"></span> list anexia vlans (*listAnexiaVlans*)

```
GET /api/v1/providers/anexia/vlans
```

Lists vlans from anexia

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| Credential | `header` | string | `string` |  |  |  |  |
| Token | `header` | string | `string` |  |  |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-anexia-vlans-200) | OK | AnexiaVlanList |  | [schema](#list-anexia-vlans-200-schema) |
| [default](#list-anexia-vlans-default) | | errorResponse |  | [schema](#list-anexia-vlans-default-schema) |

#### Responses


##### <span id="list-anexia-vlans-200"></span> 200 - AnexiaVlanList
Status: OK

###### <span id="list-anexia-vlans-200-schema"></span> Schema
   
  


 [AnexiaVlanList](#anexia-vlan-list)

##### <span id="list-anexia-vlans-default"></span> Default Response
errorResponse

###### <span id="list-anexia-vlans-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-anexia-vlans-no-credentials-v2"></span> list anexia vlans no credentials v2 (*listAnexiaVlansNoCredentialsV2*)

```
GET /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/anexia/vlans
```

Lists vlans from Anexia

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-anexia-vlans-no-credentials-v2-200) | OK | AnexiaVlanList |  | [schema](#list-anexia-vlans-no-credentials-v2-200-schema) |
| [default](#list-anexia-vlans-no-credentials-v2-default) | | errorResponse |  | [schema](#list-anexia-vlans-no-credentials-v2-default-schema) |

#### Responses


##### <span id="list-anexia-vlans-no-credentials-v2-200"></span> 200 - AnexiaVlanList
Status: OK

###### <span id="list-anexia-vlans-no-credentials-v2-200-schema"></span> Schema
   
  


 [AnexiaVlanList](#anexia-vlan-list)

##### <span id="list-anexia-vlans-no-credentials-v2-default"></span> Default Response
errorResponse

###### <span id="list-anexia-vlans-no-credentials-v2-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-azure-availability-zones"></span> list azure availability zones (*listAzureAvailabilityZones*)

```
GET /api/v1/providers/azure/availabilityzones
```

Lists VM availability zones in an Azure region

#### Produces
  * application/json

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-azure-availability-zones-200) | OK | AzureAvailabilityZonesList |  | [schema](#list-azure-availability-zones-200-schema) |
| [default](#list-azure-availability-zones-default) | | errorResponse |  | [schema](#list-azure-availability-zones-default-schema) |

#### Responses


##### <span id="list-azure-availability-zones-200"></span> 200 - AzureAvailabilityZonesList
Status: OK

###### <span id="list-azure-availability-zones-200-schema"></span> Schema
   
  

[AzureAvailabilityZonesList](#azure-availability-zones-list)

##### <span id="list-azure-availability-zones-default"></span> Default Response
errorResponse

###### <span id="list-azure-availability-zones-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-azure-availability-zones-no-credentials"></span> list azure availability zones no credentials (*listAzureAvailabilityZonesNoCredentials*)

```
GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/providers/azure/availabilityzones
```

Lists available VM availability zones in an Azure region

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| dc | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |
| SKUName | `header` | string | `string` |  |  |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-azure-availability-zones-no-credentials-200) | OK | AzureAvailabilityZonesList |  | [schema](#list-azure-availability-zones-no-credentials-200-schema) |
| [default](#list-azure-availability-zones-no-credentials-default) | | errorResponse |  | [schema](#list-azure-availability-zones-no-credentials-default-schema) |

#### Responses


##### <span id="list-azure-availability-zones-no-credentials-200"></span> 200 - AzureAvailabilityZonesList
Status: OK

###### <span id="list-azure-availability-zones-no-credentials-200-schema"></span> Schema
   
  

[AzureAvailabilityZonesList](#azure-availability-zones-list)

##### <span id="list-azure-availability-zones-no-credentials-default"></span> Default Response
errorResponse

###### <span id="list-azure-availability-zones-no-credentials-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-azure-availability-zones-no-credentials-v2"></span> list azure availability zones no credentials v2 (*listAzureAvailabilityZonesNoCredentialsV2*)

```
GET /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/azure/availabilityzones
```

Lists available VM availability zones in an Azure region

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |
| SKUName | `header` | string | `string` |  |  |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-azure-availability-zones-no-credentials-v2-200) | OK | AzureAvailabilityZonesList |  | [schema](#list-azure-availability-zones-no-credentials-v2-200-schema) |
| [default](#list-azure-availability-zones-no-credentials-v2-default) | | errorResponse |  | [schema](#list-azure-availability-zones-no-credentials-v2-default-schema) |

#### Responses


##### <span id="list-azure-availability-zones-no-credentials-v2-200"></span> 200 - AzureAvailabilityZonesList
Status: OK

###### <span id="list-azure-availability-zones-no-credentials-v2-200-schema"></span> Schema
   
  

[AzureAvailabilityZonesList](#azure-availability-zones-list)

##### <span id="list-azure-availability-zones-no-credentials-v2-default"></span> Default Response
errorResponse

###### <span id="list-azure-availability-zones-no-credentials-v2-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-azure-resource-groups"></span> list azure resource groups (*listAzureResourceGroups*)

```
GET /api/v2/providers/azure/resourcegroups
```

Lists available VM resource groups

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| ClientID | `header` | string | `string` |  |  |  |  |
| ClientSecret | `header` | string | `string` |  |  |  |  |
| Credential | `header` | string | `string` |  |  |  |  |
| Location | `header` | string | `string` |  |  |  |  |
| SubscriptionID | `header` | string | `string` |  |  |  |  |
| TenantID | `header` | string | `string` |  |  |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-azure-resource-groups-200) | OK | AzureResourceGroupsList |  | [schema](#list-azure-resource-groups-200-schema) |
| [default](#list-azure-resource-groups-default) | | errorResponse |  | [schema](#list-azure-resource-groups-default-schema) |

#### Responses


##### <span id="list-azure-resource-groups-200"></span> 200 - AzureResourceGroupsList
Status: OK

###### <span id="list-azure-resource-groups-200-schema"></span> Schema
   
  

[AzureResourceGroupsList](#azure-resource-groups-list)

##### <span id="list-azure-resource-groups-default"></span> Default Response
errorResponse

###### <span id="list-azure-resource-groups-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-azure-route-tables"></span> list azure route tables (*listAzureRouteTables*)

```
GET /api/v2/providers/azure/routetables
```

Lists available VM route tables

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| ClientID | `header` | string | `string` |  |  |  |  |
| ClientSecret | `header` | string | `string` |  |  |  |  |
| Credential | `header` | string | `string` |  |  |  |  |
| Location | `header` | string | `string` |  |  |  |  |
| ResourceGroup | `header` | string | `string` |  |  |  |  |
| SubscriptionID | `header` | string | `string` |  |  |  |  |
| TenantID | `header` | string | `string` |  |  |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-azure-route-tables-200) | OK | AzureRouteTablesList |  | [schema](#list-azure-route-tables-200-schema) |
| [default](#list-azure-route-tables-default) | | errorResponse |  | [schema](#list-azure-route-tables-default-schema) |

#### Responses


##### <span id="list-azure-route-tables-200"></span> 200 - AzureRouteTablesList
Status: OK

###### <span id="list-azure-route-tables-200-schema"></span> Schema
   
  

[AzureRouteTablesList](#azure-route-tables-list)

##### <span id="list-azure-route-tables-default"></span> Default Response
errorResponse

###### <span id="list-azure-route-tables-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-azure-security-groups"></span> list azure security groups (*listAzureSecurityGroups*)

```
GET /api/v2/providers/azure/securitygroups
```

Lists available VM security groups

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| ClientID | `header` | string | `string` |  |  |  |  |
| ClientSecret | `header` | string | `string` |  |  |  |  |
| Credential | `header` | string | `string` |  |  |  |  |
| Location | `header` | string | `string` |  |  |  |  |
| ResourceGroup | `header` | string | `string` |  |  |  |  |
| SubscriptionID | `header` | string | `string` |  |  |  |  |
| TenantID | `header` | string | `string` |  |  |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-azure-security-groups-200) | OK | AzureSecurityGroupsList |  | [schema](#list-azure-security-groups-200-schema) |
| [default](#list-azure-security-groups-default) | | errorResponse |  | [schema](#list-azure-security-groups-default-schema) |

#### Responses


##### <span id="list-azure-security-groups-200"></span> 200 - AzureSecurityGroupsList
Status: OK

###### <span id="list-azure-security-groups-200-schema"></span> Schema
   
  

[AzureSecurityGroupsList](#azure-security-groups-list)

##### <span id="list-azure-security-groups-default"></span> Default Response
errorResponse

###### <span id="list-azure-security-groups-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-azure-sizes"></span> list azure sizes (*listAzureSizes*)

```
GET /api/v1/providers/azure/sizes
```

Lists available VM sizes in an Azure region

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| ClientID | `header` | string | `string` |  |  |  |  |
| ClientSecret | `header` | string | `string` |  |  |  |  |
| Credential | `header` | string | `string` |  |  |  |  |
| Location | `header` | string | `string` |  |  |  |  |
| SubscriptionID | `header` | string | `string` |  |  |  |  |
| TenantID | `header` | string | `string` |  |  |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-azure-sizes-200) | OK | AzureSizeList |  | [schema](#list-azure-sizes-200-schema) |
| [default](#list-azure-sizes-default) | | errorResponse |  | [schema](#list-azure-sizes-default-schema) |

#### Responses


##### <span id="list-azure-sizes-200"></span> 200 - AzureSizeList
Status: OK

###### <span id="list-azure-sizes-200-schema"></span> Schema
   
  


 [AzureSizeList](#azure-size-list)

##### <span id="list-azure-sizes-default"></span> Default Response
errorResponse

###### <span id="list-azure-sizes-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-azure-sizes-no-credentials"></span> list azure sizes no credentials (*listAzureSizesNoCredentials*)

```
GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/providers/azure/sizes
```

Lists available VM sizes in an Azure region

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| dc | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-azure-sizes-no-credentials-200) | OK | AzureSizeList |  | [schema](#list-azure-sizes-no-credentials-200-schema) |
| [default](#list-azure-sizes-no-credentials-default) | | errorResponse |  | [schema](#list-azure-sizes-no-credentials-default-schema) |

#### Responses


##### <span id="list-azure-sizes-no-credentials-200"></span> 200 - AzureSizeList
Status: OK

###### <span id="list-azure-sizes-no-credentials-200-schema"></span> Schema
   
  


 [AzureSizeList](#azure-size-list)

##### <span id="list-azure-sizes-no-credentials-default"></span> Default Response
errorResponse

###### <span id="list-azure-sizes-no-credentials-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-azure-sizes-no-credentials-v2"></span> list azure sizes no credentials v2 (*listAzureSizesNoCredentialsV2*)

```
GET /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/azure/sizes
```

Lists available VM sizes in an Azure region

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-azure-sizes-no-credentials-v2-200) | OK | AzureSizeList |  | [schema](#list-azure-sizes-no-credentials-v2-200-schema) |
| [default](#list-azure-sizes-no-credentials-v2-default) | | errorResponse |  | [schema](#list-azure-sizes-no-credentials-v2-default-schema) |

#### Responses


##### <span id="list-azure-sizes-no-credentials-v2-200"></span> 200 - AzureSizeList
Status: OK

###### <span id="list-azure-sizes-no-credentials-v2-200-schema"></span> Schema
   
  


 [AzureSizeList](#azure-size-list)

##### <span id="list-azure-sizes-no-credentials-v2-default"></span> Default Response
errorResponse

###### <span id="list-azure-sizes-no-credentials-v2-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-azure-subnets"></span> list azure subnets (*listAzureSubnets*)

```
GET /api/v2/providers/azure/subnets
```

Lists available VM subnets

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| ClientID | `header` | string | `string` |  |  |  |  |
| ClientSecret | `header` | string | `string` |  |  |  |  |
| Credential | `header` | string | `string` |  |  |  |  |
| ResourceGroup | `header` | string | `string` |  |  |  |  |
| SubscriptionID | `header` | string | `string` |  |  |  |  |
| TenantID | `header` | string | `string` |  |  |  |  |
| VirtualNetwork | `header` | string | `string` |  |  |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-azure-subnets-200) | OK | AzureSubnetsList |  | [schema](#list-azure-subnets-200-schema) |
| [default](#list-azure-subnets-default) | | errorResponse |  | [schema](#list-azure-subnets-default-schema) |

#### Responses


##### <span id="list-azure-subnets-200"></span> 200 - AzureSubnetsList
Status: OK

###### <span id="list-azure-subnets-200-schema"></span> Schema
   
  

[AzureSubnetsList](#azure-subnets-list)

##### <span id="list-azure-subnets-default"></span> Default Response
errorResponse

###### <span id="list-azure-subnets-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-azure-vnets"></span> list azure vnets (*listAzureVnets*)

```
GET /api/v2/providers/azure/vnets
```

Lists available VM virtual networks

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| ClientID | `header` | string | `string` |  |  |  |  |
| ClientSecret | `header` | string | `string` |  |  |  |  |
| Credential | `header` | string | `string` |  |  |  |  |
| Location | `header` | string | `string` |  |  |  |  |
| ResourceGroup | `header` | string | `string` |  |  |  |  |
| SubscriptionID | `header` | string | `string` |  |  |  |  |
| TenantID | `header` | string | `string` |  |  |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-azure-vnets-200) | OK | AzureVirtualNetworksList |  | [schema](#list-azure-vnets-200-schema) |
| [default](#list-azure-vnets-default) | | errorResponse |  | [schema](#list-azure-vnets-default-schema) |

#### Responses


##### <span id="list-azure-vnets-200"></span> 200 - AzureVirtualNetworksList
Status: OK

###### <span id="list-azure-vnets-200-schema"></span> Schema
   
  

[AzureVirtualNetworksList](#azure-virtual-networks-list)

##### <span id="list-azure-vnets-default"></span> Default Response
errorResponse

###### <span id="list-azure-vnets-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-cluster-role"></span> list cluster role (*listClusterRole*)

```
GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/clusterroles
```

Lists all ClusterRoles

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| dc | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-cluster-role-200) | OK | ClusterRole |  | [schema](#list-cluster-role-200-schema) |
| [401](#list-cluster-role-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#list-cluster-role-401-schema) |
| [403](#list-cluster-role-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#list-cluster-role-403-schema) |
| [default](#list-cluster-role-default) | | errorResponse |  | [schema](#list-cluster-role-default-schema) |

#### Responses


##### <span id="list-cluster-role-200"></span> 200 - ClusterRole
Status: OK

###### <span id="list-cluster-role-200-schema"></span> Schema
   
  

[][ClusterRole](#cluster-role)

##### <span id="list-cluster-role-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="list-cluster-role-401-schema"></span> Schema

##### <span id="list-cluster-role-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="list-cluster-role-403-schema"></span> Schema

##### <span id="list-cluster-role-default"></span> Default Response
errorResponse

###### <span id="list-cluster-role-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-cluster-role-binding"></span> list cluster role binding (*listClusterRoleBinding*)

```
GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/clusterbindings
```

List cluster role binding

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| dc | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-cluster-role-binding-200) | OK | ClusterRoleBinding |  | [schema](#list-cluster-role-binding-200-schema) |
| [401](#list-cluster-role-binding-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#list-cluster-role-binding-401-schema) |
| [403](#list-cluster-role-binding-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#list-cluster-role-binding-403-schema) |
| [default](#list-cluster-role-binding-default) | | errorResponse |  | [schema](#list-cluster-role-binding-default-schema) |

#### Responses


##### <span id="list-cluster-role-binding-200"></span> 200 - ClusterRoleBinding
Status: OK

###### <span id="list-cluster-role-binding-200-schema"></span> Schema
   
  

[][ClusterRoleBinding](#cluster-role-binding)

##### <span id="list-cluster-role-binding-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="list-cluster-role-binding-401-schema"></span> Schema

##### <span id="list-cluster-role-binding-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="list-cluster-role-binding-403-schema"></span> Schema

##### <span id="list-cluster-role-binding-default"></span> Default Response
errorResponse

###### <span id="list-cluster-role-binding-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-cluster-role-binding-v2"></span> list cluster role binding v2 (*listClusterRoleBindingV2*)

```
GET /api/v2/projects/{project_id}/clusters/{cluster_id}/clusterbindings
```

List cluster role binding

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-cluster-role-binding-v2-200) | OK | ClusterRoleBinding |  | [schema](#list-cluster-role-binding-v2-200-schema) |
| [401](#list-cluster-role-binding-v2-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#list-cluster-role-binding-v2-401-schema) |
| [403](#list-cluster-role-binding-v2-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#list-cluster-role-binding-v2-403-schema) |
| [default](#list-cluster-role-binding-v2-default) | | errorResponse |  | [schema](#list-cluster-role-binding-v2-default-schema) |

#### Responses


##### <span id="list-cluster-role-binding-v2-200"></span> 200 - ClusterRoleBinding
Status: OK

###### <span id="list-cluster-role-binding-v2-200-schema"></span> Schema
   
  

[][ClusterRoleBinding](#cluster-role-binding)

##### <span id="list-cluster-role-binding-v2-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="list-cluster-role-binding-v2-401-schema"></span> Schema

##### <span id="list-cluster-role-binding-v2-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="list-cluster-role-binding-v2-403-schema"></span> Schema

##### <span id="list-cluster-role-binding-v2-default"></span> Default Response
errorResponse

###### <span id="list-cluster-role-binding-v2-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-cluster-role-names"></span> list cluster role names (*listClusterRoleNames*)

```
GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/clusterrolenames
```

Lists all ClusterRoles

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| dc | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-cluster-role-names-200) | OK | ClusterRoleName |  | [schema](#list-cluster-role-names-200-schema) |
| [401](#list-cluster-role-names-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#list-cluster-role-names-401-schema) |
| [403](#list-cluster-role-names-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#list-cluster-role-names-403-schema) |
| [default](#list-cluster-role-names-default) | | errorResponse |  | [schema](#list-cluster-role-names-default-schema) |

#### Responses


##### <span id="list-cluster-role-names-200"></span> 200 - ClusterRoleName
Status: OK

###### <span id="list-cluster-role-names-200-schema"></span> Schema
   
  

[][ClusterRoleName](#cluster-role-name)

##### <span id="list-cluster-role-names-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="list-cluster-role-names-401-schema"></span> Schema

##### <span id="list-cluster-role-names-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="list-cluster-role-names-403-schema"></span> Schema

##### <span id="list-cluster-role-names-default"></span> Default Response
errorResponse

###### <span id="list-cluster-role-names-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-cluster-role-names-v2"></span> list cluster role names v2 (*listClusterRoleNamesV2*)

```
GET /api/v2/projects/{project_id}/clusters/{cluster_id}/clusterrolenames
```

Lists all ClusterRoles

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-cluster-role-names-v2-200) | OK | ClusterRoleName |  | [schema](#list-cluster-role-names-v2-200-schema) |
| [401](#list-cluster-role-names-v2-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#list-cluster-role-names-v2-401-schema) |
| [403](#list-cluster-role-names-v2-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#list-cluster-role-names-v2-403-schema) |
| [default](#list-cluster-role-names-v2-default) | | errorResponse |  | [schema](#list-cluster-role-names-v2-default-schema) |

#### Responses


##### <span id="list-cluster-role-names-v2-200"></span> 200 - ClusterRoleName
Status: OK

###### <span id="list-cluster-role-names-v2-200-schema"></span> Schema
   
  

[][ClusterRoleName](#cluster-role-name)

##### <span id="list-cluster-role-names-v2-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="list-cluster-role-names-v2-401-schema"></span> Schema

##### <span id="list-cluster-role-names-v2-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="list-cluster-role-names-v2-403-schema"></span> Schema

##### <span id="list-cluster-role-names-v2-default"></span> Default Response
errorResponse

###### <span id="list-cluster-role-names-v2-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-cluster-role-v2"></span> list cluster role v2 (*listClusterRoleV2*)

```
GET /api/v2/projects/{project_id}/clusters/{cluster_id}/clusterroles
```

Lists all ClusterRoles

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-cluster-role-v2-200) | OK | ClusterRole |  | [schema](#list-cluster-role-v2-200-schema) |
| [401](#list-cluster-role-v2-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#list-cluster-role-v2-401-schema) |
| [403](#list-cluster-role-v2-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#list-cluster-role-v2-403-schema) |
| [default](#list-cluster-role-v2-default) | | errorResponse |  | [schema](#list-cluster-role-v2-default-schema) |

#### Responses


##### <span id="list-cluster-role-v2-200"></span> 200 - ClusterRole
Status: OK

###### <span id="list-cluster-role-v2-200-schema"></span> Schema
   
  

[][ClusterRole](#cluster-role)

##### <span id="list-cluster-role-v2-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="list-cluster-role-v2-401-schema"></span> Schema

##### <span id="list-cluster-role-v2-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="list-cluster-role-v2-403-schema"></span> Schema

##### <span id="list-cluster-role-v2-default"></span> Default Response
errorResponse

###### <span id="list-cluster-role-v2-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-cluster-templates"></span> List cluster templates for the given project. (*listClusterTemplates*)

```
GET /api/v2/projects/{project_id}/clustertemplates
```

#### Consumes
  * application/json

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-cluster-templates-200) | OK | ClusterTemplateList |  | [schema](#list-cluster-templates-200-schema) |
| [401](#list-cluster-templates-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#list-cluster-templates-401-schema) |
| [403](#list-cluster-templates-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#list-cluster-templates-403-schema) |
| [default](#list-cluster-templates-default) | | errorResponse |  | [schema](#list-cluster-templates-default-schema) |

#### Responses


##### <span id="list-cluster-templates-200"></span> 200 - ClusterTemplateList
Status: OK

###### <span id="list-cluster-templates-200-schema"></span> Schema
   
  


 [ClusterTemplateList](#cluster-template-list)

##### <span id="list-cluster-templates-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="list-cluster-templates-401-schema"></span> Schema

##### <span id="list-cluster-templates-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="list-cluster-templates-403-schema"></span> Schema

##### <span id="list-cluster-templates-default"></span> Default Response
errorResponse

###### <span id="list-cluster-templates-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-clusters"></span> Lists clusters for the specified project and data center. (*listClusters*)

```
GET /api/v1/projects/{project_id}/dc/{dc}/clusters
```

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| dc | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-clusters-200) | OK | ClusterList |  | [schema](#list-clusters-200-schema) |
| [401](#list-clusters-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#list-clusters-401-schema) |
| [403](#list-clusters-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#list-clusters-403-schema) |
| [default](#list-clusters-default) | | errorResponse |  | [schema](#list-clusters-default-schema) |

#### Responses


##### <span id="list-clusters-200"></span> 200 - ClusterList
Status: OK

###### <span id="list-clusters-200-schema"></span> Schema
   
  


 [ClusterList](#cluster-list)

##### <span id="list-clusters-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="list-clusters-401-schema"></span> Schema

##### <span id="list-clusters-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="list-clusters-403-schema"></span> Schema

##### <span id="list-clusters-default"></span> Default Response
errorResponse

###### <span id="list-clusters-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-clusters-for-project"></span> Lists clusters for the specified project. (*listClustersForProject*)

```
GET /api/v1/projects/{project_id}/clusters
```

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-clusters-for-project-200) | OK | ClusterList |  | [schema](#list-clusters-for-project-200-schema) |
| [401](#list-clusters-for-project-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#list-clusters-for-project-401-schema) |
| [403](#list-clusters-for-project-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#list-clusters-for-project-403-schema) |
| [default](#list-clusters-for-project-default) | | errorResponse |  | [schema](#list-clusters-for-project-default-schema) |

#### Responses


##### <span id="list-clusters-for-project-200"></span> 200 - ClusterList
Status: OK

###### <span id="list-clusters-for-project-200-schema"></span> Schema
   
  


 [ClusterList](#cluster-list)

##### <span id="list-clusters-for-project-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="list-clusters-for-project-401-schema"></span> Schema

##### <span id="list-clusters-for-project-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="list-clusters-for-project-403-schema"></span> Schema

##### <span id="list-clusters-for-project-default"></span> Default Response
errorResponse

###### <span id="list-clusters-for-project-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-clusters-v2"></span> Lists clusters for the specified project. (*listClustersV2*)

```
GET /api/v2/projects/{project_id}/clusters
```

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-clusters-v2-200) | OK | ClusterList |  | [schema](#list-clusters-v2-200-schema) |
| [401](#list-clusters-v2-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#list-clusters-v2-401-schema) |
| [403](#list-clusters-v2-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#list-clusters-v2-403-schema) |
| [default](#list-clusters-v2-default) | | errorResponse |  | [schema](#list-clusters-v2-default-schema) |

#### Responses


##### <span id="list-clusters-v2-200"></span> 200 - ClusterList
Status: OK

###### <span id="list-clusters-v2-200-schema"></span> Schema
   
  


 [ClusterList](#cluster-list)

##### <span id="list-clusters-v2-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="list-clusters-v2-401-schema"></span> Schema

##### <span id="list-clusters-v2-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="list-clusters-v2-403-schema"></span> Schema

##### <span id="list-clusters-v2-default"></span> Default Response
errorResponse

###### <span id="list-clusters-v2-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-constraint-templates"></span> List constraint templates. (*listConstraintTemplates*)

```
GET /api/v2/constrainttemplates
```

#### Produces
  * application/json

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-constraint-templates-200) | OK | ConstraintTemplate |  | [schema](#list-constraint-templates-200-schema) |
| [401](#list-constraint-templates-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#list-constraint-templates-401-schema) |
| [403](#list-constraint-templates-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#list-constraint-templates-403-schema) |
| [default](#list-constraint-templates-default) | | errorResponse |  | [schema](#list-constraint-templates-default-schema) |

#### Responses


##### <span id="list-constraint-templates-200"></span> 200 - ConstraintTemplate
Status: OK

###### <span id="list-constraint-templates-200-schema"></span> Schema
   
  

[][ConstraintTemplate](#constraint-template)

##### <span id="list-constraint-templates-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="list-constraint-templates-401-schema"></span> Schema

##### <span id="list-constraint-templates-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="list-constraint-templates-403-schema"></span> Schema

##### <span id="list-constraint-templates-default"></span> Default Response
errorResponse

###### <span id="list-constraint-templates-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-constraints"></span> Lists constraints for the specified cluster. (*listConstraints*)

```
GET /api/v2/projects/{project_id}/clusters/{cluster_id}/constraints
```

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-constraints-200) | OK | Constraint |  | [schema](#list-constraints-200-schema) |
| [401](#list-constraints-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#list-constraints-401-schema) |
| [403](#list-constraints-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#list-constraints-403-schema) |
| [default](#list-constraints-default) | | errorResponse |  | [schema](#list-constraints-default-schema) |

#### Responses


##### <span id="list-constraints-200"></span> 200 - Constraint
Status: OK

###### <span id="list-constraints-200-schema"></span> Schema
   
  

[][Constraint](#constraint)

##### <span id="list-constraints-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="list-constraints-401-schema"></span> Schema

##### <span id="list-constraints-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="list-constraints-403-schema"></span> Schema

##### <span id="list-constraints-default"></span> Default Response
errorResponse

###### <span id="list-constraints-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-credentials"></span> list credentials (*listCredentials*)

```
GET /api/v1/providers/{provider_name}/presets/credentials
```

Lists credential names for the provider

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| provider_name | `path` | string | `string` |  | ✓ |  |  |
| datacenter | `query` | string | `string` |  |  |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-credentials-200) | OK | CredentialList |  | [schema](#list-credentials-200-schema) |
| [default](#list-credentials-default) | | errorResponse |  | [schema](#list-credentials-default-schema) |

#### Responses


##### <span id="list-credentials-200"></span> 200 - CredentialList
Status: OK

###### <span id="list-credentials-200-schema"></span> Schema
   
  

[CredentialList](#credential-list)

##### <span id="list-credentials-default"></span> Default Response
errorResponse

###### <span id="list-credentials-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-d-c-for-provider"></span> Returns all datacenters for the specified provider. (*listDCForProvider*)

```
GET /api/v1/providers/{provider_name}/dc
```

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| provider_name | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-d-c-for-provider-200) | OK | Datacenter |  | [schema](#list-d-c-for-provider-200-schema) |
| [401](#list-d-c-for-provider-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#list-d-c-for-provider-401-schema) |
| [403](#list-d-c-for-provider-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#list-d-c-for-provider-403-schema) |
| [default](#list-d-c-for-provider-default) | | errorResponse |  | [schema](#list-d-c-for-provider-default-schema) |

#### Responses


##### <span id="list-d-c-for-provider-200"></span> 200 - Datacenter
Status: OK

###### <span id="list-d-c-for-provider-200-schema"></span> Schema
   
  

[][Datacenter](#datacenter)

##### <span id="list-d-c-for-provider-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="list-d-c-for-provider-401-schema"></span> Schema

##### <span id="list-d-c-for-provider-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="list-d-c-for-provider-403-schema"></span> Schema

##### <span id="list-d-c-for-provider-default"></span> Default Response
errorResponse

###### <span id="list-d-c-for-provider-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-d-c-for-seed"></span> Returns all datacenters for the specified seed. (*listDCForSeed*)

```
GET /api/v1/seed/{seed_name}/dc
```

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| seed_name | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-d-c-for-seed-200) | OK | Datacenter |  | [schema](#list-d-c-for-seed-200-schema) |
| [401](#list-d-c-for-seed-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#list-d-c-for-seed-401-schema) |
| [403](#list-d-c-for-seed-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#list-d-c-for-seed-403-schema) |
| [default](#list-d-c-for-seed-default) | | errorResponse |  | [schema](#list-d-c-for-seed-default-schema) |

#### Responses


##### <span id="list-d-c-for-seed-200"></span> 200 - Datacenter
Status: OK

###### <span id="list-d-c-for-seed-200-schema"></span> Schema
   
  

[][Datacenter](#datacenter)

##### <span id="list-d-c-for-seed-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="list-d-c-for-seed-401-schema"></span> Schema

##### <span id="list-d-c-for-seed-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="list-d-c-for-seed-403-schema"></span> Schema

##### <span id="list-d-c-for-seed-default"></span> Default Response
errorResponse

###### <span id="list-d-c-for-seed-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-datacenters"></span> list datacenters (*listDatacenters*)

```
GET /api/v1/dc
```

#### Produces
  * application/json

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-datacenters-200) | OK | DatacenterList |  | [schema](#list-datacenters-200-schema) |
| [default](#list-datacenters-default) | | errorResponse |  | [schema](#list-datacenters-default-schema) |

#### Responses


##### <span id="list-datacenters-200"></span> 200 - DatacenterList
Status: OK

###### <span id="list-datacenters-200-schema"></span> Schema
   
  


 [DatacenterList](#datacenter-list)

##### <span id="list-datacenters-default"></span> Default Response
errorResponse

###### <span id="list-datacenters-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-default-constraint"></span> List default constraint. (*listDefaultConstraint*)

```
GET /api/v2/constraints
```

#### Produces
  * application/json

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-default-constraint-200) | OK | Constraint |  | [schema](#list-default-constraint-200-schema) |
| [401](#list-default-constraint-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#list-default-constraint-401-schema) |
| [403](#list-default-constraint-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#list-default-constraint-403-schema) |
| [default](#list-default-constraint-default) | | errorResponse |  | [schema](#list-default-constraint-default-schema) |

#### Responses


##### <span id="list-default-constraint-200"></span> 200 - Constraint
Status: OK

###### <span id="list-default-constraint-200-schema"></span> Schema
   
  

[][Constraint](#constraint)

##### <span id="list-default-constraint-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="list-default-constraint-401-schema"></span> Schema

##### <span id="list-default-constraint-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="list-default-constraint-403-schema"></span> Schema

##### <span id="list-default-constraint-default"></span> Default Response
errorResponse

###### <span id="list-default-constraint-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-digitalocean-sizes"></span> list digitalocean sizes (*listDigitaloceanSizes*)

```
GET /api/v1/providers/digitalocean/sizes
```

Lists sizes from digitalocean

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| Credential | `header` | string | `string` |  |  |  |  |
| DoToken | `header` | string | `string` |  |  |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-digitalocean-sizes-200) | OK | DigitaloceanSizeList |  | [schema](#list-digitalocean-sizes-200-schema) |
| [default](#list-digitalocean-sizes-default) | | errorResponse |  | [schema](#list-digitalocean-sizes-default-schema) |

#### Responses


##### <span id="list-digitalocean-sizes-200"></span> 200 - DigitaloceanSizeList
Status: OK

###### <span id="list-digitalocean-sizes-200-schema"></span> Schema
   
  

[DigitaloceanSizeList](#digitalocean-size-list)

##### <span id="list-digitalocean-sizes-default"></span> Default Response
errorResponse

###### <span id="list-digitalocean-sizes-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-digitalocean-sizes-no-credentials"></span> list digitalocean sizes no credentials (*listDigitaloceanSizesNoCredentials*)

```
GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/providers/digitalocean/sizes
```

Lists sizes from digitalocean

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| dc | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-digitalocean-sizes-no-credentials-200) | OK | DigitaloceanSizeList |  | [schema](#list-digitalocean-sizes-no-credentials-200-schema) |
| [default](#list-digitalocean-sizes-no-credentials-default) | | errorResponse |  | [schema](#list-digitalocean-sizes-no-credentials-default-schema) |

#### Responses


##### <span id="list-digitalocean-sizes-no-credentials-200"></span> 200 - DigitaloceanSizeList
Status: OK

###### <span id="list-digitalocean-sizes-no-credentials-200-schema"></span> Schema
   
  

[DigitaloceanSizeList](#digitalocean-size-list)

##### <span id="list-digitalocean-sizes-no-credentials-default"></span> Default Response
errorResponse

###### <span id="list-digitalocean-sizes-no-credentials-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-digitalocean-sizes-no-credentials-v2"></span> list digitalocean sizes no credentials v2 (*listDigitaloceanSizesNoCredentialsV2*)

```
GET /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/digitalocean/sizes
```

Lists sizes from digitalocean

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-digitalocean-sizes-no-credentials-v2-200) | OK | DigitaloceanSizeList |  | [schema](#list-digitalocean-sizes-no-credentials-v2-200-schema) |
| [default](#list-digitalocean-sizes-no-credentials-v2-default) | | errorResponse |  | [schema](#list-digitalocean-sizes-no-credentials-v2-default-schema) |

#### Responses


##### <span id="list-digitalocean-sizes-no-credentials-v2-200"></span> 200 - DigitaloceanSizeList
Status: OK

###### <span id="list-digitalocean-sizes-no-credentials-v2-200-schema"></span> Schema
   
  

[DigitaloceanSizeList](#digitalocean-size-list)

##### <span id="list-digitalocean-sizes-no-credentials-v2-default"></span> Default Response
errorResponse

###### <span id="list-digitalocean-sizes-no-credentials-v2-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-etcd-backup-config"></span> list etcd backup config (*listEtcdBackupConfig*)

```
GET /api/v2/projects/{project_id}/clusters/{cluster_id}/etcdbackupconfigs
```

List etcd backup configs for a given cluster

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-etcd-backup-config-200) | OK | EtcdBackupConfig |  | [schema](#list-etcd-backup-config-200-schema) |
| [401](#list-etcd-backup-config-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#list-etcd-backup-config-401-schema) |
| [403](#list-etcd-backup-config-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#list-etcd-backup-config-403-schema) |
| [default](#list-etcd-backup-config-default) | | errorResponse |  | [schema](#list-etcd-backup-config-default-schema) |

#### Responses


##### <span id="list-etcd-backup-config-200"></span> 200 - EtcdBackupConfig
Status: OK

###### <span id="list-etcd-backup-config-200-schema"></span> Schema
   
  

[][EtcdBackupConfig](#etcd-backup-config)

##### <span id="list-etcd-backup-config-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="list-etcd-backup-config-401-schema"></span> Schema

##### <span id="list-etcd-backup-config-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="list-etcd-backup-config-403-schema"></span> Schema

##### <span id="list-etcd-backup-config-default"></span> Default Response
errorResponse

###### <span id="list-etcd-backup-config-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-etcd-restore"></span> list etcd restore (*listEtcdRestore*)

```
GET /api/v2/projects/{project_id}/clusters/{cluster_id}/etcdrestores
```

List etcd backup restores for a given cluster

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-etcd-restore-200) | OK | EtcdRestore |  | [schema](#list-etcd-restore-200-schema) |
| [401](#list-etcd-restore-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#list-etcd-restore-401-schema) |
| [403](#list-etcd-restore-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#list-etcd-restore-403-schema) |
| [default](#list-etcd-restore-default) | | errorResponse |  | [schema](#list-etcd-restore-default-schema) |

#### Responses


##### <span id="list-etcd-restore-200"></span> 200 - EtcdRestore
Status: OK

###### <span id="list-etcd-restore-200-schema"></span> Schema
   
  

[][EtcdRestore](#etcd-restore)

##### <span id="list-etcd-restore-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="list-etcd-restore-401-schema"></span> Schema

##### <span id="list-etcd-restore-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="list-etcd-restore-403-schema"></span> Schema

##### <span id="list-etcd-restore-default"></span> Default Response
errorResponse

###### <span id="list-etcd-restore-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-external-cluster-events"></span> Gets an external cluster events. (*listExternalClusterEvents*)

```
GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/events
```

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |
| type | `query` | string | `string` |  |  |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-external-cluster-events-200) | OK | Event |  | [schema](#list-external-cluster-events-200-schema) |
| [401](#list-external-cluster-events-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#list-external-cluster-events-401-schema) |
| [403](#list-external-cluster-events-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#list-external-cluster-events-403-schema) |
| [default](#list-external-cluster-events-default) | | errorResponse |  | [schema](#list-external-cluster-events-default-schema) |

#### Responses


##### <span id="list-external-cluster-events-200"></span> 200 - Event
Status: OK

###### <span id="list-external-cluster-events-200-schema"></span> Schema
   
  

[][Event](#event)

##### <span id="list-external-cluster-events-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="list-external-cluster-events-401-schema"></span> Schema

##### <span id="list-external-cluster-events-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="list-external-cluster-events-403-schema"></span> Schema

##### <span id="list-external-cluster-events-default"></span> Default Response
errorResponse

###### <span id="list-external-cluster-events-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-external-cluster-nodes"></span> Gets an external cluster nodes. (*listExternalClusterNodes*)

```
GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/nodes
```

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-external-cluster-nodes-200) | OK | Node |  | [schema](#list-external-cluster-nodes-200-schema) |
| [401](#list-external-cluster-nodes-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#list-external-cluster-nodes-401-schema) |
| [403](#list-external-cluster-nodes-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#list-external-cluster-nodes-403-schema) |
| [default](#list-external-cluster-nodes-default) | | errorResponse |  | [schema](#list-external-cluster-nodes-default-schema) |

#### Responses


##### <span id="list-external-cluster-nodes-200"></span> 200 - Node
Status: OK

###### <span id="list-external-cluster-nodes-200-schema"></span> Schema
   
  

[][Node](#node)

##### <span id="list-external-cluster-nodes-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="list-external-cluster-nodes-401-schema"></span> Schema

##### <span id="list-external-cluster-nodes-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="list-external-cluster-nodes-403-schema"></span> Schema

##### <span id="list-external-cluster-nodes-default"></span> Default Response
errorResponse

###### <span id="list-external-cluster-nodes-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-external-cluster-nodes-metrics"></span> Gets an external cluster nodes metrics. (*listExternalClusterNodesMetrics*)

```
GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/nodesmetrics
```

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-external-cluster-nodes-metrics-200) | OK | NodeMetric |  | [schema](#list-external-cluster-nodes-metrics-200-schema) |
| [401](#list-external-cluster-nodes-metrics-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#list-external-cluster-nodes-metrics-401-schema) |
| [403](#list-external-cluster-nodes-metrics-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#list-external-cluster-nodes-metrics-403-schema) |
| [default](#list-external-cluster-nodes-metrics-default) | | errorResponse |  | [schema](#list-external-cluster-nodes-metrics-default-schema) |

#### Responses


##### <span id="list-external-cluster-nodes-metrics-200"></span> 200 - NodeMetric
Status: OK

###### <span id="list-external-cluster-nodes-metrics-200-schema"></span> Schema
   
  

[][NodeMetric](#node-metric)

##### <span id="list-external-cluster-nodes-metrics-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="list-external-cluster-nodes-metrics-401-schema"></span> Schema

##### <span id="list-external-cluster-nodes-metrics-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="list-external-cluster-nodes-metrics-403-schema"></span> Schema

##### <span id="list-external-cluster-nodes-metrics-default"></span> Default Response
errorResponse

###### <span id="list-external-cluster-nodes-metrics-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-external-clusters"></span> Lists external clusters for the specified project. (*listExternalClusters*)

```
GET /api/v2/projects/{project_id}/kubernetes/clusters
```

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-external-clusters-200) | OK | ClusterList |  | [schema](#list-external-clusters-200-schema) |
| [401](#list-external-clusters-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#list-external-clusters-401-schema) |
| [403](#list-external-clusters-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#list-external-clusters-403-schema) |
| [default](#list-external-clusters-default) | | errorResponse |  | [schema](#list-external-clusters-default-schema) |

#### Responses


##### <span id="list-external-clusters-200"></span> 200 - ClusterList
Status: OK

###### <span id="list-external-clusters-200-schema"></span> Schema
   
  


 [ClusterList](#cluster-list)

##### <span id="list-external-clusters-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="list-external-clusters-401-schema"></span> Schema

##### <span id="list-external-clusters-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="list-external-clusters-403-schema"></span> Schema

##### <span id="list-external-clusters-default"></span> Default Response
errorResponse

###### <span id="list-external-clusters-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-g-c-p-disk-types"></span> list g c p disk types (*listGCPDiskTypes*)

```
GET /api/v1/providers/gcp/disktypes
```

Lists disk types from GCP

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| Credential | `header` | string | `string` |  |  |  |  |
| ServiceAccount | `header` | string | `string` |  |  |  |  |
| Zone | `header` | string | `string` |  |  |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-g-c-p-disk-types-200) | OK | GCPDiskTypeList |  | [schema](#list-g-c-p-disk-types-200-schema) |
| [default](#list-g-c-p-disk-types-default) | | errorResponse |  | [schema](#list-g-c-p-disk-types-default-schema) |

#### Responses


##### <span id="list-g-c-p-disk-types-200"></span> 200 - GCPDiskTypeList
Status: OK

###### <span id="list-g-c-p-disk-types-200-schema"></span> Schema
   
  


 [GCPDiskTypeList](#g-c-p-disk-type-list)

##### <span id="list-g-c-p-disk-types-default"></span> Default Response
errorResponse

###### <span id="list-g-c-p-disk-types-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-g-c-p-disk-types-no-credentials"></span> list g c p disk types no credentials (*listGCPDiskTypesNoCredentials*)

```
GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/providers/gcp/disktypes
```

Lists disk types from GCP

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| dc | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |
| Zone | `header` | string | `string` |  |  |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-g-c-p-disk-types-no-credentials-200) | OK | GCPDiskTypeList |  | [schema](#list-g-c-p-disk-types-no-credentials-200-schema) |
| [default](#list-g-c-p-disk-types-no-credentials-default) | | errorResponse |  | [schema](#list-g-c-p-disk-types-no-credentials-default-schema) |

#### Responses


##### <span id="list-g-c-p-disk-types-no-credentials-200"></span> 200 - GCPDiskTypeList
Status: OK

###### <span id="list-g-c-p-disk-types-no-credentials-200-schema"></span> Schema
   
  


 [GCPDiskTypeList](#g-c-p-disk-type-list)

##### <span id="list-g-c-p-disk-types-no-credentials-default"></span> Default Response
errorResponse

###### <span id="list-g-c-p-disk-types-no-credentials-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-g-c-p-disk-types-no-credentials-v2"></span> list g c p disk types no credentials v2 (*listGCPDiskTypesNoCredentialsV2*)

```
GET /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/gcp/disktypes
```

Lists disk types from GCP

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |
| Zone | `header` | string | `string` |  |  |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-g-c-p-disk-types-no-credentials-v2-200) | OK | GCPDiskTypeList |  | [schema](#list-g-c-p-disk-types-no-credentials-v2-200-schema) |
| [default](#list-g-c-p-disk-types-no-credentials-v2-default) | | errorResponse |  | [schema](#list-g-c-p-disk-types-no-credentials-v2-default-schema) |

#### Responses


##### <span id="list-g-c-p-disk-types-no-credentials-v2-200"></span> 200 - GCPDiskTypeList
Status: OK

###### <span id="list-g-c-p-disk-types-no-credentials-v2-200-schema"></span> Schema
   
  


 [GCPDiskTypeList](#g-c-p-disk-type-list)

##### <span id="list-g-c-p-disk-types-no-credentials-v2-default"></span> Default Response
errorResponse

###### <span id="list-g-c-p-disk-types-no-credentials-v2-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-g-c-p-networks"></span> list g c p networks (*listGCPNetworks*)

```
GET /api/v1/providers/gcp/networks
```

Lists networks from GCP

#### Produces
  * application/json

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-g-c-p-networks-200) | OK | GCPNetworkList |  | [schema](#list-g-c-p-networks-200-schema) |
| [default](#list-g-c-p-networks-default) | | errorResponse |  | [schema](#list-g-c-p-networks-default-schema) |

#### Responses


##### <span id="list-g-c-p-networks-200"></span> 200 - GCPNetworkList
Status: OK

###### <span id="list-g-c-p-networks-200-schema"></span> Schema
   
  


 [GCPNetworkList](#g-c-p-network-list)

##### <span id="list-g-c-p-networks-default"></span> Default Response
errorResponse

###### <span id="list-g-c-p-networks-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-g-c-p-networks-no-credentials"></span> list g c p networks no credentials (*listGCPNetworksNoCredentials*)

```
GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/providers/gcp/networks
```

Lists available GCP networks

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| dc | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-g-c-p-networks-no-credentials-200) | OK | GCPNetworkList |  | [schema](#list-g-c-p-networks-no-credentials-200-schema) |
| [default](#list-g-c-p-networks-no-credentials-default) | | errorResponse |  | [schema](#list-g-c-p-networks-no-credentials-default-schema) |

#### Responses


##### <span id="list-g-c-p-networks-no-credentials-200"></span> 200 - GCPNetworkList
Status: OK

###### <span id="list-g-c-p-networks-no-credentials-200-schema"></span> Schema
   
  


 [GCPNetworkList](#g-c-p-network-list)

##### <span id="list-g-c-p-networks-no-credentials-default"></span> Default Response
errorResponse

###### <span id="list-g-c-p-networks-no-credentials-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-g-c-p-networks-no-credentials-v2"></span> list g c p networks no credentials v2 (*listGCPNetworksNoCredentialsV2*)

```
GET /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/gcp/networks
```

Lists available GCP networks

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-g-c-p-networks-no-credentials-v2-200) | OK | GCPNetworkList |  | [schema](#list-g-c-p-networks-no-credentials-v2-200-schema) |
| [default](#list-g-c-p-networks-no-credentials-v2-default) | | errorResponse |  | [schema](#list-g-c-p-networks-no-credentials-v2-default-schema) |

#### Responses


##### <span id="list-g-c-p-networks-no-credentials-v2-200"></span> 200 - GCPNetworkList
Status: OK

###### <span id="list-g-c-p-networks-no-credentials-v2-200-schema"></span> Schema
   
  


 [GCPNetworkList](#g-c-p-network-list)

##### <span id="list-g-c-p-networks-no-credentials-v2-default"></span> Default Response
errorResponse

###### <span id="list-g-c-p-networks-no-credentials-v2-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-g-c-p-sizes"></span> list g c p sizes (*listGCPSizes*)

```
GET /api/v1/providers/gcp/sizes
```

Lists machine sizes from GCP

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| Credential | `header` | string | `string` |  |  |  |  |
| ServiceAccount | `header` | string | `string` |  |  |  |  |
| Zone | `header` | string | `string` |  |  |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-g-c-p-sizes-200) | OK | GCPMachineSizeList |  | [schema](#list-g-c-p-sizes-200-schema) |
| [default](#list-g-c-p-sizes-default) | | errorResponse |  | [schema](#list-g-c-p-sizes-default-schema) |

#### Responses


##### <span id="list-g-c-p-sizes-200"></span> 200 - GCPMachineSizeList
Status: OK

###### <span id="list-g-c-p-sizes-200-schema"></span> Schema
   
  


 [GCPMachineSizeList](#g-c-p-machine-size-list)

##### <span id="list-g-c-p-sizes-default"></span> Default Response
errorResponse

###### <span id="list-g-c-p-sizes-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-g-c-p-sizes-no-credentials"></span> list g c p sizes no credentials (*listGCPSizesNoCredentials*)

```
GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/providers/gcp/sizes
```

Lists machine sizes from GCP

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| dc | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |
| Zone | `header` | string | `string` |  |  |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-g-c-p-sizes-no-credentials-200) | OK | GCPMachineSizeList |  | [schema](#list-g-c-p-sizes-no-credentials-200-schema) |
| [default](#list-g-c-p-sizes-no-credentials-default) | | errorResponse |  | [schema](#list-g-c-p-sizes-no-credentials-default-schema) |

#### Responses


##### <span id="list-g-c-p-sizes-no-credentials-200"></span> 200 - GCPMachineSizeList
Status: OK

###### <span id="list-g-c-p-sizes-no-credentials-200-schema"></span> Schema
   
  


 [GCPMachineSizeList](#g-c-p-machine-size-list)

##### <span id="list-g-c-p-sizes-no-credentials-default"></span> Default Response
errorResponse

###### <span id="list-g-c-p-sizes-no-credentials-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-g-c-p-sizes-no-credentials-v2"></span> list g c p sizes no credentials v2 (*listGCPSizesNoCredentialsV2*)

```
GET /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/gcp/sizes
```

Lists machine sizes from GCP

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |
| Zone | `header` | string | `string` |  |  |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-g-c-p-sizes-no-credentials-v2-200) | OK | GCPMachineSizeList |  | [schema](#list-g-c-p-sizes-no-credentials-v2-200-schema) |
| [default](#list-g-c-p-sizes-no-credentials-v2-default) | | errorResponse |  | [schema](#list-g-c-p-sizes-no-credentials-v2-default-schema) |

#### Responses


##### <span id="list-g-c-p-sizes-no-credentials-v2-200"></span> 200 - GCPMachineSizeList
Status: OK

###### <span id="list-g-c-p-sizes-no-credentials-v2-200-schema"></span> Schema
   
  


 [GCPMachineSizeList](#g-c-p-machine-size-list)

##### <span id="list-g-c-p-sizes-no-credentials-v2-default"></span> Default Response
errorResponse

###### <span id="list-g-c-p-sizes-no-credentials-v2-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-g-c-p-subnetworks"></span> list g c p subnetworks (*listGCPSubnetworks*)

```
GET /api/v1/providers/gcp/{dc}/subnetworks
```

Lists subnetworks from GCP

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| dc | `path` | string | `string` |  | ✓ |  |  |
| Credential | `header` | string | `string` |  |  |  |  |
| Network | `header` | string | `string` |  |  |  |  |
| ServiceAccount | `header` | string | `string` |  |  |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-g-c-p-subnetworks-200) | OK | GCPSubnetworkList |  | [schema](#list-g-c-p-subnetworks-200-schema) |
| [default](#list-g-c-p-subnetworks-default) | | errorResponse |  | [schema](#list-g-c-p-subnetworks-default-schema) |

#### Responses


##### <span id="list-g-c-p-subnetworks-200"></span> 200 - GCPSubnetworkList
Status: OK

###### <span id="list-g-c-p-subnetworks-200-schema"></span> Schema
   
  


 [GCPSubnetworkList](#g-c-p-subnetwork-list)

##### <span id="list-g-c-p-subnetworks-default"></span> Default Response
errorResponse

###### <span id="list-g-c-p-subnetworks-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-g-c-p-subnetworks-no-credentials"></span> list g c p subnetworks no credentials (*listGCPSubnetworksNoCredentials*)

```
GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/providers/gcp/subnetworks
```

Lists available GCP subnetworks

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| dc | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |
| Network | `header` | string | `string` |  |  |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-g-c-p-subnetworks-no-credentials-200) | OK | GCPSubnetworkList |  | [schema](#list-g-c-p-subnetworks-no-credentials-200-schema) |
| [default](#list-g-c-p-subnetworks-no-credentials-default) | | errorResponse |  | [schema](#list-g-c-p-subnetworks-no-credentials-default-schema) |

#### Responses


##### <span id="list-g-c-p-subnetworks-no-credentials-200"></span> 200 - GCPSubnetworkList
Status: OK

###### <span id="list-g-c-p-subnetworks-no-credentials-200-schema"></span> Schema
   
  


 [GCPSubnetworkList](#g-c-p-subnetwork-list)

##### <span id="list-g-c-p-subnetworks-no-credentials-default"></span> Default Response
errorResponse

###### <span id="list-g-c-p-subnetworks-no-credentials-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-g-c-p-subnetworks-no-credentials-v2"></span> list g c p subnetworks no credentials v2 (*listGCPSubnetworksNoCredentialsV2*)

```
GET /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/gcp/subnetworks
```

Lists available GCP subnetworks

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |
| Network | `header` | string | `string` |  |  |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-g-c-p-subnetworks-no-credentials-v2-200) | OK | GCPSubnetworkList |  | [schema](#list-g-c-p-subnetworks-no-credentials-v2-200-schema) |
| [default](#list-g-c-p-subnetworks-no-credentials-v2-default) | | errorResponse |  | [schema](#list-g-c-p-subnetworks-no-credentials-v2-default-schema) |

#### Responses


##### <span id="list-g-c-p-subnetworks-no-credentials-v2-200"></span> 200 - GCPSubnetworkList
Status: OK

###### <span id="list-g-c-p-subnetworks-no-credentials-v2-200-schema"></span> Schema
   
  


 [GCPSubnetworkList](#g-c-p-subnetwork-list)

##### <span id="list-g-c-p-subnetworks-no-credentials-v2-default"></span> Default Response
errorResponse

###### <span id="list-g-c-p-subnetworks-no-credentials-v2-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-g-c-p-zones"></span> list g c p zones (*listGCPZones*)

```
GET /api/v1/providers/gcp/{dc}/zones
```

Lists available GCP zones

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| dc | `path` | string | `string` |  | ✓ |  |  |
| Credential | `header` | string | `string` |  |  |  |  |
| ServiceAccount | `header` | string | `string` |  |  |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-g-c-p-zones-200) | OK | GCPZoneList |  | [schema](#list-g-c-p-zones-200-schema) |
| [default](#list-g-c-p-zones-default) | | errorResponse |  | [schema](#list-g-c-p-zones-default-schema) |

#### Responses


##### <span id="list-g-c-p-zones-200"></span> 200 - GCPZoneList
Status: OK

###### <span id="list-g-c-p-zones-200-schema"></span> Schema
   
  


 [GCPZoneList](#g-c-p-zone-list)

##### <span id="list-g-c-p-zones-default"></span> Default Response
errorResponse

###### <span id="list-g-c-p-zones-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-g-c-p-zones-no-credentials"></span> list g c p zones no credentials (*listGCPZonesNoCredentials*)

```
GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/providers/gcp/zones
```

Lists available GCP zones

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| dc | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-g-c-p-zones-no-credentials-200) | OK | GCPZoneList |  | [schema](#list-g-c-p-zones-no-credentials-200-schema) |
| [default](#list-g-c-p-zones-no-credentials-default) | | errorResponse |  | [schema](#list-g-c-p-zones-no-credentials-default-schema) |

#### Responses


##### <span id="list-g-c-p-zones-no-credentials-200"></span> 200 - GCPZoneList
Status: OK

###### <span id="list-g-c-p-zones-no-credentials-200-schema"></span> Schema
   
  


 [GCPZoneList](#g-c-p-zone-list)

##### <span id="list-g-c-p-zones-no-credentials-default"></span> Default Response
errorResponse

###### <span id="list-g-c-p-zones-no-credentials-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-g-c-p-zones-no-credentials-v2"></span> list g c p zones no credentials v2 (*listGCPZonesNoCredentialsV2*)

```
GET /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/gcp/zones
```

Lists available GCP zones

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-g-c-p-zones-no-credentials-v2-200) | OK | GCPZoneList |  | [schema](#list-g-c-p-zones-no-credentials-v2-200-schema) |
| [default](#list-g-c-p-zones-no-credentials-v2-default) | | errorResponse |  | [schema](#list-g-c-p-zones-no-credentials-v2-default-schema) |

#### Responses


##### <span id="list-g-c-p-zones-no-credentials-v2-200"></span> 200 - GCPZoneList
Status: OK

###### <span id="list-g-c-p-zones-no-credentials-v2-200-schema"></span> Schema
   
  


 [GCPZoneList](#g-c-p-zone-list)

##### <span id="list-g-c-p-zones-no-credentials-v2-default"></span> Default Response
errorResponse

###### <span id="list-g-c-p-zones-no-credentials-v2-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-hetzner-sizes"></span> list hetzner sizes (*listHetznerSizes*)

```
GET /api/v1/providers/hetzner/sizes
```

Lists sizes from hetzner

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| Credential | `header` | string | `string` |  |  |  |  |
| HetznerToken | `header` | string | `string` |  |  |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-hetzner-sizes-200) | OK | HetznerSizeList |  | [schema](#list-hetzner-sizes-200-schema) |
| [default](#list-hetzner-sizes-default) | | errorResponse |  | [schema](#list-hetzner-sizes-default-schema) |

#### Responses


##### <span id="list-hetzner-sizes-200"></span> 200 - HetznerSizeList
Status: OK

###### <span id="list-hetzner-sizes-200-schema"></span> Schema
   
  

[HetznerSizeList](#hetzner-size-list)

##### <span id="list-hetzner-sizes-default"></span> Default Response
errorResponse

###### <span id="list-hetzner-sizes-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-hetzner-sizes-no-credentials"></span> list hetzner sizes no credentials (*listHetznerSizesNoCredentials*)

```
GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/providers/hetzner/sizes
```

Lists sizes from hetzner

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| dc | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-hetzner-sizes-no-credentials-200) | OK | HetznerSizeList |  | [schema](#list-hetzner-sizes-no-credentials-200-schema) |
| [default](#list-hetzner-sizes-no-credentials-default) | | errorResponse |  | [schema](#list-hetzner-sizes-no-credentials-default-schema) |

#### Responses


##### <span id="list-hetzner-sizes-no-credentials-200"></span> 200 - HetznerSizeList
Status: OK

###### <span id="list-hetzner-sizes-no-credentials-200-schema"></span> Schema
   
  

[HetznerSizeList](#hetzner-size-list)

##### <span id="list-hetzner-sizes-no-credentials-default"></span> Default Response
errorResponse

###### <span id="list-hetzner-sizes-no-credentials-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-hetzner-sizes-no-credentials-v2"></span> list hetzner sizes no credentials v2 (*listHetznerSizesNoCredentialsV2*)

```
GET /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/hetzner/sizes
```

Lists sizes from hetzner

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-hetzner-sizes-no-credentials-v2-200) | OK | HetznerSizeList |  | [schema](#list-hetzner-sizes-no-credentials-v2-200-schema) |
| [default](#list-hetzner-sizes-no-credentials-v2-default) | | errorResponse |  | [schema](#list-hetzner-sizes-no-credentials-v2-default-schema) |

#### Responses


##### <span id="list-hetzner-sizes-no-credentials-v2-200"></span> 200 - HetznerSizeList
Status: OK

###### <span id="list-hetzner-sizes-no-credentials-v2-200-schema"></span> Schema
   
  

[HetznerSizeList](#hetzner-size-list)

##### <span id="list-hetzner-sizes-no-credentials-v2-default"></span> Default Response
errorResponse

###### <span id="list-hetzner-sizes-no-credentials-v2-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-installable-addons"></span> list installable addons (*listInstallableAddons*)

```
GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/installableaddons
```

Lists names of addons that can be installed inside the user cluster

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| dc | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-installable-addons-200) | OK | AccessibleAddons |  | [schema](#list-installable-addons-200-schema) |
| [401](#list-installable-addons-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#list-installable-addons-401-schema) |
| [403](#list-installable-addons-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#list-installable-addons-403-schema) |
| [default](#list-installable-addons-default) | | errorResponse |  | [schema](#list-installable-addons-default-schema) |

#### Responses


##### <span id="list-installable-addons-200"></span> 200 - AccessibleAddons
Status: OK

###### <span id="list-installable-addons-200-schema"></span> Schema
   
  


 [AccessibleAddons](#accessible-addons)

##### <span id="list-installable-addons-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="list-installable-addons-401-schema"></span> Schema

##### <span id="list-installable-addons-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="list-installable-addons-403-schema"></span> Schema

##### <span id="list-installable-addons-default"></span> Default Response
errorResponse

###### <span id="list-installable-addons-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-installable-addons-v2"></span> list installable addons v2 (*listInstallableAddonsV2*)

```
GET /api/v2/projects/{project_id}/clusters/{cluster_id}/installableaddons
```

Lists names of addons that can be installed inside the user cluster

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-installable-addons-v2-200) | OK | AccessibleAddons |  | [schema](#list-installable-addons-v2-200-schema) |
| [401](#list-installable-addons-v2-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#list-installable-addons-v2-401-schema) |
| [403](#list-installable-addons-v2-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#list-installable-addons-v2-403-schema) |
| [default](#list-installable-addons-v2-default) | | errorResponse |  | [schema](#list-installable-addons-v2-default-schema) |

#### Responses


##### <span id="list-installable-addons-v2-200"></span> 200 - AccessibleAddons
Status: OK

###### <span id="list-installable-addons-v2-200-schema"></span> Schema
   
  


 [AccessibleAddons](#accessible-addons)

##### <span id="list-installable-addons-v2-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="list-installable-addons-v2-401-schema"></span> Schema

##### <span id="list-installable-addons-v2-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="list-installable-addons-v2-403-schema"></span> Schema

##### <span id="list-installable-addons-v2-default"></span> Default Response
errorResponse

###### <span id="list-installable-addons-v2-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-machine-deployment-metrics"></span> Lists metrics that belong to the given machine deployment. (*listMachineDeploymentMetrics*)

```
GET /api/v2/projects/{project_id}/clusters/{cluster_id}/machinedeployments/{machinedeployment_id}/nodes/metrics
```

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| machinedeployment_id | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-machine-deployment-metrics-200) | OK | NodeMetric |  | [schema](#list-machine-deployment-metrics-200-schema) |
| [401](#list-machine-deployment-metrics-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#list-machine-deployment-metrics-401-schema) |
| [403](#list-machine-deployment-metrics-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#list-machine-deployment-metrics-403-schema) |
| [default](#list-machine-deployment-metrics-default) | | errorResponse |  | [schema](#list-machine-deployment-metrics-default-schema) |

#### Responses


##### <span id="list-machine-deployment-metrics-200"></span> 200 - NodeMetric
Status: OK

###### <span id="list-machine-deployment-metrics-200-schema"></span> Schema
   
  

[][NodeMetric](#node-metric)

##### <span id="list-machine-deployment-metrics-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="list-machine-deployment-metrics-401-schema"></span> Schema

##### <span id="list-machine-deployment-metrics-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="list-machine-deployment-metrics-403-schema"></span> Schema

##### <span id="list-machine-deployment-metrics-default"></span> Default Response
errorResponse

###### <span id="list-machine-deployment-metrics-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-machine-deployment-nodes"></span> Lists nodes that belong to the given machine deployment. (*listMachineDeploymentNodes*)

```
GET /api/v2/projects/{project_id}/clusters/{cluster_id}/machinedeployments/{machinedeployment_id}/nodes
```

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| machinedeployment_id | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |
| hideInitialConditions | `query` | boolean | `bool` |  |  |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-machine-deployment-nodes-200) | OK | Node |  | [schema](#list-machine-deployment-nodes-200-schema) |
| [401](#list-machine-deployment-nodes-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#list-machine-deployment-nodes-401-schema) |
| [403](#list-machine-deployment-nodes-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#list-machine-deployment-nodes-403-schema) |
| [default](#list-machine-deployment-nodes-default) | | errorResponse |  | [schema](#list-machine-deployment-nodes-default-schema) |

#### Responses


##### <span id="list-machine-deployment-nodes-200"></span> 200 - Node
Status: OK

###### <span id="list-machine-deployment-nodes-200-schema"></span> Schema
   
  

[][Node](#node)

##### <span id="list-machine-deployment-nodes-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="list-machine-deployment-nodes-401-schema"></span> Schema

##### <span id="list-machine-deployment-nodes-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="list-machine-deployment-nodes-403-schema"></span> Schema

##### <span id="list-machine-deployment-nodes-default"></span> Default Response
errorResponse

###### <span id="list-machine-deployment-nodes-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-machine-deployment-nodes-events"></span> Lists machine deployment events. If query parameter `type` is set to `warning` then only warning events are retrieved. (*listMachineDeploymentNodesEvents*)

```
GET /api/v2/projects/{project_id}/clusters/{cluster_id}/machinedeployments/{machinedeployment_id}/nodes/events
```

If the value is 'normal' then normal events are returned. If the query parameter is missing method returns all events.

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| machinedeployment_id | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |
| type | `query` | string | `string` |  |  |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-machine-deployment-nodes-events-200) | OK | Event |  | [schema](#list-machine-deployment-nodes-events-200-schema) |
| [401](#list-machine-deployment-nodes-events-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#list-machine-deployment-nodes-events-401-schema) |
| [403](#list-machine-deployment-nodes-events-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#list-machine-deployment-nodes-events-403-schema) |
| [default](#list-machine-deployment-nodes-events-default) | | errorResponse |  | [schema](#list-machine-deployment-nodes-events-default-schema) |

#### Responses


##### <span id="list-machine-deployment-nodes-events-200"></span> 200 - Event
Status: OK

###### <span id="list-machine-deployment-nodes-events-200-schema"></span> Schema
   
  

[][Event](#event)

##### <span id="list-machine-deployment-nodes-events-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="list-machine-deployment-nodes-events-401-schema"></span> Schema

##### <span id="list-machine-deployment-nodes-events-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="list-machine-deployment-nodes-events-403-schema"></span> Schema

##### <span id="list-machine-deployment-nodes-events-default"></span> Default Response
errorResponse

###### <span id="list-machine-deployment-nodes-events-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-machine-deployments"></span> list machine deployments (*listMachineDeployments*)

```
GET /api/v2/projects/{project_id}/clusters/{cluster_id}/machinedeployments
```

Lists machine deployments that belong to the given cluster

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-machine-deployments-200) | OK | NodeDeployment |  | [schema](#list-machine-deployments-200-schema) |
| [401](#list-machine-deployments-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#list-machine-deployments-401-schema) |
| [403](#list-machine-deployments-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#list-machine-deployments-403-schema) |
| [default](#list-machine-deployments-default) | | errorResponse |  | [schema](#list-machine-deployments-default-schema) |

#### Responses


##### <span id="list-machine-deployments-200"></span> 200 - NodeDeployment
Status: OK

###### <span id="list-machine-deployments-200-schema"></span> Schema
   
  

[][NodeDeployment](#node-deployment)

##### <span id="list-machine-deployments-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="list-machine-deployments-401-schema"></span> Schema

##### <span id="list-machine-deployments-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="list-machine-deployments-403-schema"></span> Schema

##### <span id="list-machine-deployments-default"></span> Default Response
errorResponse

###### <span id="list-machine-deployments-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-metering-reports"></span> list metering reports (*listMeteringReports*)

```
GET /api/v1/admin/metering/reports
```

List metering reports. Only available in Kubermatic Enterprise Edition

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| max_keys | `query` | int64 (formatted integer) | `int64` |  |  |  |  |
| start_after | `query` | string | `string` |  |  |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-metering-reports-200) | OK | MeteringReport |  | [schema](#list-metering-reports-200-schema) |
| [401](#list-metering-reports-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#list-metering-reports-401-schema) |
| [403](#list-metering-reports-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#list-metering-reports-403-schema) |
| [default](#list-metering-reports-default) | | errorResponse |  | [schema](#list-metering-reports-default-schema) |

#### Responses


##### <span id="list-metering-reports-200"></span> 200 - MeteringReport
Status: OK

###### <span id="list-metering-reports-200-schema"></span> Schema
   
  

[][MeteringReport](#metering-report)

##### <span id="list-metering-reports-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="list-metering-reports-401-schema"></span> Schema

##### <span id="list-metering-reports-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="list-metering-reports-403-schema"></span> Schema

##### <span id="list-metering-reports-default"></span> Default Response
errorResponse

###### <span id="list-metering-reports-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-namespace"></span> list namespace (*listNamespace*)

```
GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/namespaces
```

Lists all namespaces in the cluster

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| dc | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-namespace-200) | OK | Namespace |  | [schema](#list-namespace-200-schema) |
| [401](#list-namespace-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#list-namespace-401-schema) |
| [403](#list-namespace-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#list-namespace-403-schema) |
| [default](#list-namespace-default) | | errorResponse |  | [schema](#list-namespace-default-schema) |

#### Responses


##### <span id="list-namespace-200"></span> 200 - Namespace
Status: OK

###### <span id="list-namespace-200-schema"></span> Schema
   
  

[][Namespace](#namespace)

##### <span id="list-namespace-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="list-namespace-401-schema"></span> Schema

##### <span id="list-namespace-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="list-namespace-403-schema"></span> Schema

##### <span id="list-namespace-default"></span> Default Response
errorResponse

###### <span id="list-namespace-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-namespace-v2"></span> list namespace v2 (*listNamespaceV2*)

```
GET /api/v2/projects/{project_id}/clusters/{cluster_id}/namespaces
```

Lists all namespaces in the cluster

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-namespace-v2-200) | OK | Namespace |  | [schema](#list-namespace-v2-200-schema) |
| [401](#list-namespace-v2-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#list-namespace-v2-401-schema) |
| [403](#list-namespace-v2-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#list-namespace-v2-403-schema) |
| [default](#list-namespace-v2-default) | | errorResponse |  | [schema](#list-namespace-v2-default-schema) |

#### Responses


##### <span id="list-namespace-v2-200"></span> 200 - Namespace
Status: OK

###### <span id="list-namespace-v2-200-schema"></span> Schema
   
  

[][Namespace](#namespace)

##### <span id="list-namespace-v2-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="list-namespace-v2-401-schema"></span> Schema

##### <span id="list-namespace-v2-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="list-namespace-v2-403-schema"></span> Schema

##### <span id="list-namespace-v2-default"></span> Default Response
errorResponse

###### <span id="list-namespace-v2-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-node-deployment-metrics"></span> Lists metrics that belong to the given node deployment. (*listNodeDeploymentMetrics*)

```
GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/nodedeployments/{nodedeployment_id}/nodes/metrics
```

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| dc | `path` | string | `string` |  | ✓ |  |  |
| nodedeployment_id | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-node-deployment-metrics-200) | OK | NodeMetric |  | [schema](#list-node-deployment-metrics-200-schema) |
| [401](#list-node-deployment-metrics-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#list-node-deployment-metrics-401-schema) |
| [403](#list-node-deployment-metrics-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#list-node-deployment-metrics-403-schema) |
| [default](#list-node-deployment-metrics-default) | | errorResponse |  | [schema](#list-node-deployment-metrics-default-schema) |

#### Responses


##### <span id="list-node-deployment-metrics-200"></span> 200 - NodeMetric
Status: OK

###### <span id="list-node-deployment-metrics-200-schema"></span> Schema
   
  

[][NodeMetric](#node-metric)

##### <span id="list-node-deployment-metrics-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="list-node-deployment-metrics-401-schema"></span> Schema

##### <span id="list-node-deployment-metrics-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="list-node-deployment-metrics-403-schema"></span> Schema

##### <span id="list-node-deployment-metrics-default"></span> Default Response
errorResponse

###### <span id="list-node-deployment-metrics-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-node-deployment-nodes"></span> Lists nodes that belong to the given node deployment. (*listNodeDeploymentNodes*)

```
GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/nodedeployments/{nodedeployment_id}/nodes
```

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| dc | `path` | string | `string` |  | ✓ |  |  |
| nodedeployment_id | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |
| hideInitialConditions | `query` | boolean | `bool` |  |  |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-node-deployment-nodes-200) | OK | Node |  | [schema](#list-node-deployment-nodes-200-schema) |
| [401](#list-node-deployment-nodes-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#list-node-deployment-nodes-401-schema) |
| [403](#list-node-deployment-nodes-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#list-node-deployment-nodes-403-schema) |
| [default](#list-node-deployment-nodes-default) | | errorResponse |  | [schema](#list-node-deployment-nodes-default-schema) |

#### Responses


##### <span id="list-node-deployment-nodes-200"></span> 200 - Node
Status: OK

###### <span id="list-node-deployment-nodes-200-schema"></span> Schema
   
  

[][Node](#node)

##### <span id="list-node-deployment-nodes-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="list-node-deployment-nodes-401-schema"></span> Schema

##### <span id="list-node-deployment-nodes-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="list-node-deployment-nodes-403-schema"></span> Schema

##### <span id="list-node-deployment-nodes-default"></span> Default Response
errorResponse

###### <span id="list-node-deployment-nodes-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-node-deployment-nodes-events"></span> Lists node deployment events. If query parameter `type` is set to `warning` then only warning events are retrieved. (*listNodeDeploymentNodesEvents*)

```
GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/nodedeployments/{nodedeployment_id}/nodes/events
```

If the value is 'normal' then normal events are returned. If the query parameter is missing method returns all events.

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| dc | `path` | string | `string` |  | ✓ |  |  |
| nodedeployment_id | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |
| type | `query` | string | `string` |  |  |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-node-deployment-nodes-events-200) | OK | Event |  | [schema](#list-node-deployment-nodes-events-200-schema) |
| [401](#list-node-deployment-nodes-events-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#list-node-deployment-nodes-events-401-schema) |
| [403](#list-node-deployment-nodes-events-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#list-node-deployment-nodes-events-403-schema) |
| [default](#list-node-deployment-nodes-events-default) | | errorResponse |  | [schema](#list-node-deployment-nodes-events-default-schema) |

#### Responses


##### <span id="list-node-deployment-nodes-events-200"></span> 200 - Event
Status: OK

###### <span id="list-node-deployment-nodes-events-200-schema"></span> Schema
   
  

[][Event](#event)

##### <span id="list-node-deployment-nodes-events-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="list-node-deployment-nodes-events-401-schema"></span> Schema

##### <span id="list-node-deployment-nodes-events-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="list-node-deployment-nodes-events-403-schema"></span> Schema

##### <span id="list-node-deployment-nodes-events-default"></span> Default Response
errorResponse

###### <span id="list-node-deployment-nodes-events-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-node-deployment-requests"></span> list node deployment requests (*listNodeDeploymentRequests*)

```
GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/ndrequests
```

Lists NodeDeploymentRequests that belong to the given cluster

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| dc | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-node-deployment-requests-200) | OK | NodeDeploymentRequest |  | [schema](#list-node-deployment-requests-200-schema) |
| [401](#list-node-deployment-requests-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#list-node-deployment-requests-401-schema) |
| [403](#list-node-deployment-requests-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#list-node-deployment-requests-403-schema) |
| [default](#list-node-deployment-requests-default) | | errorResponse |  | [schema](#list-node-deployment-requests-default-schema) |

#### Responses


##### <span id="list-node-deployment-requests-200"></span> 200 - NodeDeploymentRequest
Status: OK

###### <span id="list-node-deployment-requests-200-schema"></span> Schema
   
  

[][NodeDeploymentRequest](#node-deployment-request)

##### <span id="list-node-deployment-requests-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="list-node-deployment-requests-401-schema"></span> Schema

##### <span id="list-node-deployment-requests-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="list-node-deployment-requests-403-schema"></span> Schema

##### <span id="list-node-deployment-requests-default"></span> Default Response
errorResponse

###### <span id="list-node-deployment-requests-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-node-deployments"></span> list node deployments (*listNodeDeployments*)

```
GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/nodedeployments
```

Lists node deployments that belong to the given cluster

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| dc | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-node-deployments-200) | OK | NodeDeployment |  | [schema](#list-node-deployments-200-schema) |
| [401](#list-node-deployments-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#list-node-deployments-401-schema) |
| [403](#list-node-deployments-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#list-node-deployments-403-schema) |
| [default](#list-node-deployments-default) | | errorResponse |  | [schema](#list-node-deployments-default-schema) |

#### Responses


##### <span id="list-node-deployments-200"></span> 200 - NodeDeployment
Status: OK

###### <span id="list-node-deployments-200-schema"></span> Schema
   
  

[][NodeDeployment](#node-deployment)

##### <span id="list-node-deployments-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="list-node-deployments-401-schema"></span> Schema

##### <span id="list-node-deployments-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="list-node-deployments-403-schema"></span> Schema

##### <span id="list-node-deployments-default"></span> Default Response
errorResponse

###### <span id="list-node-deployments-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-nodes-for-cluster"></span> This endpoint is used for kubeadm cluster. (*listNodesForCluster*)

```
GET /api/v2/projects/{project_id}/clusters/{cluster_id}/nodes
```

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |
| hideInitialConditions | `query` | boolean | `bool` |  |  |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-nodes-for-cluster-200) | OK | Node |  | [schema](#list-nodes-for-cluster-200-schema) |
| [401](#list-nodes-for-cluster-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#list-nodes-for-cluster-401-schema) |
| [403](#list-nodes-for-cluster-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#list-nodes-for-cluster-403-schema) |
| [default](#list-nodes-for-cluster-default) | | errorResponse |  | [schema](#list-nodes-for-cluster-default-schema) |

#### Responses


##### <span id="list-nodes-for-cluster-200"></span> 200 - Node
Status: OK

###### <span id="list-nodes-for-cluster-200-schema"></span> Schema
   
  

[][Node](#node)

##### <span id="list-nodes-for-cluster-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="list-nodes-for-cluster-401-schema"></span> Schema

##### <span id="list-nodes-for-cluster-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="list-nodes-for-cluster-403-schema"></span> Schema

##### <span id="list-nodes-for-cluster-default"></span> Default Response
errorResponse

###### <span id="list-nodes-for-cluster-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-openstack-availability-zones"></span> list openstack availability zones (*listOpenstackAvailabilityZones*)

```
GET /api/v1/providers/openstack/availabilityzones
```

Lists availability zones from openstack

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| ApplicationCredentialID | `header` | string | `string` |  |  |  |  |
| ApplicationCredentialSecret | `header` | string | `string` |  |  |  |  |
| Credential | `header` | string | `string` |  |  |  |  |
| DatacenterName | `header` | string | `string` |  |  |  |  |
| Domain | `header` | string | `string` |  |  |  |  |
| OIDCAuthentication | `header` | boolean | `bool` |  |  |  |  |
| Password | `header` | string | `string` |  |  |  |  |
| Tenant | `header` | string | `string` |  |  |  |  |
| TenantID | `header` | string | `string` |  |  |  |  |
| Username | `header` | string | `string` |  |  |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-openstack-availability-zones-200) | OK | OpenstackAvailabilityZone |  | [schema](#list-openstack-availability-zones-200-schema) |
| [default](#list-openstack-availability-zones-default) | | errorResponse |  | [schema](#list-openstack-availability-zones-default-schema) |

#### Responses


##### <span id="list-openstack-availability-zones-200"></span> 200 - OpenstackAvailabilityZone
Status: OK

###### <span id="list-openstack-availability-zones-200-schema"></span> Schema
   
  

[][OpenstackAvailabilityZone](#openstack-availability-zone)

##### <span id="list-openstack-availability-zones-default"></span> Default Response
errorResponse

###### <span id="list-openstack-availability-zones-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-openstack-availability-zones-no-credentials"></span> list openstack availability zones no credentials (*listOpenstackAvailabilityZonesNoCredentials*)

```
GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/providers/openstack/availabilityzones
```

Lists availability zones from openstack

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| dc | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-openstack-availability-zones-no-credentials-200) | OK | OpenstackAvailabilityZone |  | [schema](#list-openstack-availability-zones-no-credentials-200-schema) |
| [default](#list-openstack-availability-zones-no-credentials-default) | | errorResponse |  | [schema](#list-openstack-availability-zones-no-credentials-default-schema) |

#### Responses


##### <span id="list-openstack-availability-zones-no-credentials-200"></span> 200 - OpenstackAvailabilityZone
Status: OK

###### <span id="list-openstack-availability-zones-no-credentials-200-schema"></span> Schema
   
  

[][OpenstackAvailabilityZone](#openstack-availability-zone)

##### <span id="list-openstack-availability-zones-no-credentials-default"></span> Default Response
errorResponse

###### <span id="list-openstack-availability-zones-no-credentials-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-openstack-availability-zones-no-credentials-v2"></span> list openstack availability zones no credentials v2 (*listOpenstackAvailabilityZonesNoCredentialsV2*)

```
GET /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/openstack/availabilityzones
```

Lists availability zones from openstack

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-openstack-availability-zones-no-credentials-v2-200) | OK | OpenstackAvailabilityZone |  | [schema](#list-openstack-availability-zones-no-credentials-v2-200-schema) |
| [default](#list-openstack-availability-zones-no-credentials-v2-default) | | errorResponse |  | [schema](#list-openstack-availability-zones-no-credentials-v2-default-schema) |

#### Responses


##### <span id="list-openstack-availability-zones-no-credentials-v2-200"></span> 200 - OpenstackAvailabilityZone
Status: OK

###### <span id="list-openstack-availability-zones-no-credentials-v2-200-schema"></span> Schema
   
  

[][OpenstackAvailabilityZone](#openstack-availability-zone)

##### <span id="list-openstack-availability-zones-no-credentials-v2-default"></span> Default Response
errorResponse

###### <span id="list-openstack-availability-zones-no-credentials-v2-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-openstack-images"></span> list openstack images (*listOpenstackImages*)

```
GET /api/v1/providers/openstack/images
```

Lists images from openstack

#### Produces
  * application/json

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-openstack-images-200) | OK | Image |  | [schema](#list-openstack-images-200-schema) |
| [default](#list-openstack-images-default) | | errorResponse |  | [schema](#list-openstack-images-default-schema) |

#### Responses


##### <span id="list-openstack-images-200"></span> 200 - Image
Status: OK

###### <span id="list-openstack-images-200-schema"></span> Schema
   
  

[][Image](#image)

##### <span id="list-openstack-images-default"></span> Default Response
errorResponse

###### <span id="list-openstack-images-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-openstack-images-no-credentials"></span> list openstack images no credentials (*listOpenstackImagesNoCredentials*)

```
GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/providers/openstack/images
```

Lists images from openstack

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| dc | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-openstack-images-no-credentials-200) | OK | Image |  | [schema](#list-openstack-images-no-credentials-200-schema) |
| [default](#list-openstack-images-no-credentials-default) | | errorResponse |  | [schema](#list-openstack-images-no-credentials-default-schema) |

#### Responses


##### <span id="list-openstack-images-no-credentials-200"></span> 200 - Image
Status: OK

###### <span id="list-openstack-images-no-credentials-200-schema"></span> Schema
   
  

[][Image](#image)

##### <span id="list-openstack-images-no-credentials-default"></span> Default Response
errorResponse

###### <span id="list-openstack-images-no-credentials-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-openstack-networks"></span> list openstack networks (*listOpenstackNetworks*)

```
GET /api/v1/providers/openstack/networks
```

Lists networks from openstack

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| ApplicationCredentialID | `header` | string | `string` |  |  |  |  |
| ApplicationCredentialSecret | `header` | string | `string` |  |  |  |  |
| Credential | `header` | string | `string` |  |  |  |  |
| DatacenterName | `header` | string | `string` |  |  |  |  |
| Domain | `header` | string | `string` |  |  |  |  |
| OIDCAuthentication | `header` | boolean | `bool` |  |  |  |  |
| Password | `header` | string | `string` |  |  |  |  |
| Tenant | `header` | string | `string` |  |  |  |  |
| TenantID | `header` | string | `string` |  |  |  |  |
| Username | `header` | string | `string` |  |  |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-openstack-networks-200) | OK | OpenstackNetwork |  | [schema](#list-openstack-networks-200-schema) |
| [default](#list-openstack-networks-default) | | errorResponse |  | [schema](#list-openstack-networks-default-schema) |

#### Responses


##### <span id="list-openstack-networks-200"></span> 200 - OpenstackNetwork
Status: OK

###### <span id="list-openstack-networks-200-schema"></span> Schema
   
  

[][OpenstackNetwork](#openstack-network)

##### <span id="list-openstack-networks-default"></span> Default Response
errorResponse

###### <span id="list-openstack-networks-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-openstack-networks-no-credentials"></span> list openstack networks no credentials (*listOpenstackNetworksNoCredentials*)

```
GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/providers/openstack/networks
```

Lists networks from openstack

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| dc | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-openstack-networks-no-credentials-200) | OK | OpenstackNetwork |  | [schema](#list-openstack-networks-no-credentials-200-schema) |
| [default](#list-openstack-networks-no-credentials-default) | | errorResponse |  | [schema](#list-openstack-networks-no-credentials-default-schema) |

#### Responses


##### <span id="list-openstack-networks-no-credentials-200"></span> 200 - OpenstackNetwork
Status: OK

###### <span id="list-openstack-networks-no-credentials-200-schema"></span> Schema
   
  

[][OpenstackNetwork](#openstack-network)

##### <span id="list-openstack-networks-no-credentials-default"></span> Default Response
errorResponse

###### <span id="list-openstack-networks-no-credentials-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-openstack-networks-no-credentials-v2"></span> list openstack networks no credentials v2 (*listOpenstackNetworksNoCredentialsV2*)

```
GET /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/openstack/networks
```

Lists networks from openstack

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-openstack-networks-no-credentials-v2-200) | OK | OpenstackNetwork |  | [schema](#list-openstack-networks-no-credentials-v2-200-schema) |
| [default](#list-openstack-networks-no-credentials-v2-default) | | errorResponse |  | [schema](#list-openstack-networks-no-credentials-v2-default-schema) |

#### Responses


##### <span id="list-openstack-networks-no-credentials-v2-200"></span> 200 - OpenstackNetwork
Status: OK

###### <span id="list-openstack-networks-no-credentials-v2-200-schema"></span> Schema
   
  

[][OpenstackNetwork](#openstack-network)

##### <span id="list-openstack-networks-no-credentials-v2-default"></span> Default Response
errorResponse

###### <span id="list-openstack-networks-no-credentials-v2-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-openstack-quota-limits"></span> list openstack quota limits (*listOpenstackQuotaLimits*)

```
GET /api/v1/providers/openstack/quotalimits
```

Lists quotalimits for tenant from openstack

#### Produces
  * application/json

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-openstack-quota-limits-200) | OK | Quotas |  | [schema](#list-openstack-quota-limits-200-schema) |
| [default](#list-openstack-quota-limits-default) | | errorResponse |  | [schema](#list-openstack-quota-limits-default-schema) |

#### Responses


##### <span id="list-openstack-quota-limits-200"></span> 200 - Quotas
Status: OK

###### <span id="list-openstack-quota-limits-200-schema"></span> Schema
   
  

[Quotas](#quotas)

##### <span id="list-openstack-quota-limits-default"></span> Default Response
errorResponse

###### <span id="list-openstack-quota-limits-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-openstack-quota-limits-no-credentials"></span> list openstack quota limits no credentials (*listOpenstackQuotaLimitsNoCredentials*)

```
GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/providers/openstack/quotalimits
```

Lists quotalimits for tenant from openstack

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| dc | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-openstack-quota-limits-no-credentials-200) | OK | Quotas |  | [schema](#list-openstack-quota-limits-no-credentials-200-schema) |
| [default](#list-openstack-quota-limits-no-credentials-default) | | errorResponse |  | [schema](#list-openstack-quota-limits-no-credentials-default-schema) |

#### Responses


##### <span id="list-openstack-quota-limits-no-credentials-200"></span> 200 - Quotas
Status: OK

###### <span id="list-openstack-quota-limits-no-credentials-200-schema"></span> Schema
   
  

[Quotas](#quotas)

##### <span id="list-openstack-quota-limits-no-credentials-default"></span> Default Response
errorResponse

###### <span id="list-openstack-quota-limits-no-credentials-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-openstack-security-groups"></span> list openstack security groups (*listOpenstackSecurityGroups*)

```
GET /api/v1/providers/openstack/securitygroups
```

Lists security groups from openstack

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| ApplicationCredentialID | `header` | string | `string` |  |  |  |  |
| ApplicationCredentialSecret | `header` | string | `string` |  |  |  |  |
| Credential | `header` | string | `string` |  |  |  |  |
| DatacenterName | `header` | string | `string` |  |  |  |  |
| Domain | `header` | string | `string` |  |  |  |  |
| OIDCAuthentication | `header` | boolean | `bool` |  |  |  |  |
| Password | `header` | string | `string` |  |  |  |  |
| Tenant | `header` | string | `string` |  |  |  |  |
| TenantID | `header` | string | `string` |  |  |  |  |
| Username | `header` | string | `string` |  |  |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-openstack-security-groups-200) | OK | OpenstackSecurityGroup |  | [schema](#list-openstack-security-groups-200-schema) |
| [default](#list-openstack-security-groups-default) | | errorResponse |  | [schema](#list-openstack-security-groups-default-schema) |

#### Responses


##### <span id="list-openstack-security-groups-200"></span> 200 - OpenstackSecurityGroup
Status: OK

###### <span id="list-openstack-security-groups-200-schema"></span> Schema
   
  

[][OpenstackSecurityGroup](#openstack-security-group)

##### <span id="list-openstack-security-groups-default"></span> Default Response
errorResponse

###### <span id="list-openstack-security-groups-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-openstack-security-groups-no-credentials"></span> list openstack security groups no credentials (*listOpenstackSecurityGroupsNoCredentials*)

```
GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/providers/openstack/securitygroups
```

Lists security groups from openstack

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| dc | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-openstack-security-groups-no-credentials-200) | OK | OpenstackSecurityGroup |  | [schema](#list-openstack-security-groups-no-credentials-200-schema) |
| [default](#list-openstack-security-groups-no-credentials-default) | | errorResponse |  | [schema](#list-openstack-security-groups-no-credentials-default-schema) |

#### Responses


##### <span id="list-openstack-security-groups-no-credentials-200"></span> 200 - OpenstackSecurityGroup
Status: OK

###### <span id="list-openstack-security-groups-no-credentials-200-schema"></span> Schema
   
  

[][OpenstackSecurityGroup](#openstack-security-group)

##### <span id="list-openstack-security-groups-no-credentials-default"></span> Default Response
errorResponse

###### <span id="list-openstack-security-groups-no-credentials-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-openstack-security-groups-no-credentials-v2"></span> list openstack security groups no credentials v2 (*listOpenstackSecurityGroupsNoCredentialsV2*)

```
GET /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/openstack/securitygroups
```

Lists security groups from openstack

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-openstack-security-groups-no-credentials-v2-200) | OK | OpenstackSecurityGroup |  | [schema](#list-openstack-security-groups-no-credentials-v2-200-schema) |
| [default](#list-openstack-security-groups-no-credentials-v2-default) | | errorResponse |  | [schema](#list-openstack-security-groups-no-credentials-v2-default-schema) |

#### Responses


##### <span id="list-openstack-security-groups-no-credentials-v2-200"></span> 200 - OpenstackSecurityGroup
Status: OK

###### <span id="list-openstack-security-groups-no-credentials-v2-200-schema"></span> Schema
   
  

[][OpenstackSecurityGroup](#openstack-security-group)

##### <span id="list-openstack-security-groups-no-credentials-v2-default"></span> Default Response
errorResponse

###### <span id="list-openstack-security-groups-no-credentials-v2-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-openstack-sizes"></span> list openstack sizes (*listOpenstackSizes*)

```
GET /api/v1/providers/openstack/sizes
```

Lists sizes from openstack

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| ApplicationCredentialID | `header` | string | `string` |  |  |  |  |
| ApplicationCredentialSecret | `header` | string | `string` |  |  |  |  |
| Credential | `header` | string | `string` |  |  |  |  |
| DatacenterName | `header` | string | `string` |  |  |  |  |
| Domain | `header` | string | `string` |  |  |  |  |
| OIDCAuthentication | `header` | boolean | `bool` |  |  |  |  |
| Password | `header` | string | `string` |  |  |  |  |
| Tenant | `header` | string | `string` |  |  |  |  |
| TenantID | `header` | string | `string` |  |  |  |  |
| Username | `header` | string | `string` |  |  |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-openstack-sizes-200) | OK | OpenstackSize |  | [schema](#list-openstack-sizes-200-schema) |
| [default](#list-openstack-sizes-default) | | errorResponse |  | [schema](#list-openstack-sizes-default-schema) |

#### Responses


##### <span id="list-openstack-sizes-200"></span> 200 - OpenstackSize
Status: OK

###### <span id="list-openstack-sizes-200-schema"></span> Schema
   
  

[][OpenstackSize](#openstack-size)

##### <span id="list-openstack-sizes-default"></span> Default Response
errorResponse

###### <span id="list-openstack-sizes-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-openstack-sizes-no-credentials"></span> list openstack sizes no credentials (*listOpenstackSizesNoCredentials*)

```
GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/providers/openstack/sizes
```

Lists sizes from openstack

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| dc | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-openstack-sizes-no-credentials-200) | OK | OpenstackSize |  | [schema](#list-openstack-sizes-no-credentials-200-schema) |
| [default](#list-openstack-sizes-no-credentials-default) | | errorResponse |  | [schema](#list-openstack-sizes-no-credentials-default-schema) |

#### Responses


##### <span id="list-openstack-sizes-no-credentials-200"></span> 200 - OpenstackSize
Status: OK

###### <span id="list-openstack-sizes-no-credentials-200-schema"></span> Schema
   
  

[][OpenstackSize](#openstack-size)

##### <span id="list-openstack-sizes-no-credentials-default"></span> Default Response
errorResponse

###### <span id="list-openstack-sizes-no-credentials-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-openstack-sizes-no-credentials-v2"></span> list openstack sizes no credentials v2 (*listOpenstackSizesNoCredentialsV2*)

```
GET /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/openstack/sizes
```

Lists sizes from openstack

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-openstack-sizes-no-credentials-v2-200) | OK | OpenstackSize |  | [schema](#list-openstack-sizes-no-credentials-v2-200-schema) |
| [default](#list-openstack-sizes-no-credentials-v2-default) | | errorResponse |  | [schema](#list-openstack-sizes-no-credentials-v2-default-schema) |

#### Responses


##### <span id="list-openstack-sizes-no-credentials-v2-200"></span> 200 - OpenstackSize
Status: OK

###### <span id="list-openstack-sizes-no-credentials-v2-200-schema"></span> Schema
   
  

[][OpenstackSize](#openstack-size)

##### <span id="list-openstack-sizes-no-credentials-v2-default"></span> Default Response
errorResponse

###### <span id="list-openstack-sizes-no-credentials-v2-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-openstack-subnets"></span> list openstack subnets (*listOpenstackSubnets*)

```
GET /api/v1/providers/openstack/subnets
```

Lists subnets from openstack

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| ApplicationCredentialID | `header` | string | `string` |  |  |  |  |
| ApplicationCredentialSecret | `header` | string | `string` |  |  |  |  |
| Credential | `header` | string | `string` |  |  |  |  |
| DatacenterName | `header` | string | `string` |  |  |  |  |
| Domain | `header` | string | `string` |  |  |  |  |
| OIDCAuthentication | `header` | boolean | `bool` |  |  |  |  |
| Password | `header` | string | `string` |  |  |  |  |
| Tenant | `header` | string | `string` |  |  |  |  |
| TenantID | `header` | string | `string` |  |  |  |  |
| Username | `header` | string | `string` |  |  |  |  |
| network_id | `query` | string | `string` |  |  |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-openstack-subnets-200) | OK | OpenstackSubnet |  | [schema](#list-openstack-subnets-200-schema) |
| [default](#list-openstack-subnets-default) | | errorResponse |  | [schema](#list-openstack-subnets-default-schema) |

#### Responses


##### <span id="list-openstack-subnets-200"></span> 200 - OpenstackSubnet
Status: OK

###### <span id="list-openstack-subnets-200-schema"></span> Schema
   
  

[][OpenstackSubnet](#openstack-subnet)

##### <span id="list-openstack-subnets-default"></span> Default Response
errorResponse

###### <span id="list-openstack-subnets-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-openstack-subnets-no-credentials"></span> list openstack subnets no credentials (*listOpenstackSubnetsNoCredentials*)

```
GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/providers/openstack/subnets
```

Lists subnets from openstack

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| dc | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |
| network_id | `query` | string | `string` |  |  |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-openstack-subnets-no-credentials-200) | OK | OpenstackSubnet |  | [schema](#list-openstack-subnets-no-credentials-200-schema) |
| [default](#list-openstack-subnets-no-credentials-default) | | errorResponse |  | [schema](#list-openstack-subnets-no-credentials-default-schema) |

#### Responses


##### <span id="list-openstack-subnets-no-credentials-200"></span> 200 - OpenstackSubnet
Status: OK

###### <span id="list-openstack-subnets-no-credentials-200-schema"></span> Schema
   
  

[][OpenstackSubnet](#openstack-subnet)

##### <span id="list-openstack-subnets-no-credentials-default"></span> Default Response
errorResponse

###### <span id="list-openstack-subnets-no-credentials-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-openstack-subnets-no-credentials-v2"></span> list openstack subnets no credentials v2 (*listOpenstackSubnetsNoCredentialsV2*)

```
GET /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/openstack/subnets
```

Lists subnets from openstack

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |
| network_id | `query` | string | `string` |  |  |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-openstack-subnets-no-credentials-v2-200) | OK | OpenstackSubnet |  | [schema](#list-openstack-subnets-no-credentials-v2-200-schema) |
| [default](#list-openstack-subnets-no-credentials-v2-default) | | errorResponse |  | [schema](#list-openstack-subnets-no-credentials-v2-default-schema) |

#### Responses


##### <span id="list-openstack-subnets-no-credentials-v2-200"></span> 200 - OpenstackSubnet
Status: OK

###### <span id="list-openstack-subnets-no-credentials-v2-200-schema"></span> Schema
   
  

[][OpenstackSubnet](#openstack-subnet)

##### <span id="list-openstack-subnets-no-credentials-v2-default"></span> Default Response
errorResponse

###### <span id="list-openstack-subnets-no-credentials-v2-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-openstack-tenants"></span> list openstack tenants (*listOpenstackTenants*)

```
GET /api/v1/providers/openstack/tenants
```

Lists tenants from openstack

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| ApplicationCredentialID | `header` | string | `string` |  |  |  |  |
| ApplicationCredentialSecret | `header` | string | `string` |  |  |  |  |
| Credential | `header` | string | `string` |  |  |  |  |
| DatacenterName | `header` | string | `string` |  |  |  |  |
| Domain | `header` | string | `string` |  |  |  |  |
| OIDCAuthentication | `header` | boolean | `bool` |  |  |  |  |
| Password | `header` | string | `string` |  |  |  |  |
| Username | `header` | string | `string` |  |  |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-openstack-tenants-200) | OK | OpenstackTenant |  | [schema](#list-openstack-tenants-200-schema) |
| [default](#list-openstack-tenants-default) | | errorResponse |  | [schema](#list-openstack-tenants-default-schema) |

#### Responses


##### <span id="list-openstack-tenants-200"></span> 200 - OpenstackTenant
Status: OK

###### <span id="list-openstack-tenants-200-schema"></span> Schema
   
  

[][OpenstackTenant](#openstack-tenant)

##### <span id="list-openstack-tenants-default"></span> Default Response
errorResponse

###### <span id="list-openstack-tenants-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-openstack-tenants-no-credentials"></span> list openstack tenants no credentials (*listOpenstackTenantsNoCredentials*)

```
GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/providers/openstack/tenants
```

Lists tenants from openstack

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| dc | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-openstack-tenants-no-credentials-200) | OK | OpenstackTenant |  | [schema](#list-openstack-tenants-no-credentials-200-schema) |
| [default](#list-openstack-tenants-no-credentials-default) | | errorResponse |  | [schema](#list-openstack-tenants-no-credentials-default-schema) |

#### Responses


##### <span id="list-openstack-tenants-no-credentials-200"></span> 200 - OpenstackTenant
Status: OK

###### <span id="list-openstack-tenants-no-credentials-200-schema"></span> Schema
   
  

[][OpenstackTenant](#openstack-tenant)

##### <span id="list-openstack-tenants-no-credentials-default"></span> Default Response
errorResponse

###### <span id="list-openstack-tenants-no-credentials-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-openstack-tenants-no-credentials-v2"></span> list openstack tenants no credentials v2 (*listOpenstackTenantsNoCredentialsV2*)

```
GET /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/openstack/tenants
```

Lists tenants from openstack

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-openstack-tenants-no-credentials-v2-200) | OK | OpenstackTenant |  | [schema](#list-openstack-tenants-no-credentials-v2-200-schema) |
| [default](#list-openstack-tenants-no-credentials-v2-default) | | errorResponse |  | [schema](#list-openstack-tenants-no-credentials-v2-default-schema) |

#### Responses


##### <span id="list-openstack-tenants-no-credentials-v2-200"></span> 200 - OpenstackTenant
Status: OK

###### <span id="list-openstack-tenants-no-credentials-v2-200-schema"></span> Schema
   
  

[][OpenstackTenant](#openstack-tenant)

##### <span id="list-openstack-tenants-no-credentials-v2-default"></span> Default Response
errorResponse

###### <span id="list-openstack-tenants-no-credentials-v2-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-packet-sizes"></span> list packet sizes (*listPacketSizes*)

```
GET /api/v1/providers/packet/sizes
```

Lists sizes from packet

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| apiKey | `header` | string | `string` |  |  |  |  |
| credential | `header` | string | `string` |  |  |  |  |
| projectID | `header` | string | `string` |  |  |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-packet-sizes-200) | OK | PacketSizeList |  | [schema](#list-packet-sizes-200-schema) |
| [default](#list-packet-sizes-default) | | errorResponse |  | [schema](#list-packet-sizes-default-schema) |

#### Responses


##### <span id="list-packet-sizes-200"></span> 200 - PacketSizeList
Status: OK

###### <span id="list-packet-sizes-200-schema"></span> Schema
   
  

[][PacketSizeList](#packet-size-list)

##### <span id="list-packet-sizes-default"></span> Default Response
errorResponse

###### <span id="list-packet-sizes-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-packet-sizes-no-credentials"></span> list packet sizes no credentials (*listPacketSizesNoCredentials*)

```
GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/providers/packet/sizes
```

Lists sizes from packet

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| dc | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-packet-sizes-no-credentials-200) | OK | PacketSizeList |  | [schema](#list-packet-sizes-no-credentials-200-schema) |
| [default](#list-packet-sizes-no-credentials-default) | | errorResponse |  | [schema](#list-packet-sizes-no-credentials-default-schema) |

#### Responses


##### <span id="list-packet-sizes-no-credentials-200"></span> 200 - PacketSizeList
Status: OK

###### <span id="list-packet-sizes-no-credentials-200-schema"></span> Schema
   
  

[][PacketSizeList](#packet-size-list)

##### <span id="list-packet-sizes-no-credentials-default"></span> Default Response
errorResponse

###### <span id="list-packet-sizes-no-credentials-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-packet-sizes-no-credentials-v2"></span> list packet sizes no credentials v2 (*listPacketSizesNoCredentialsV2*)

```
GET /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/packet/sizes
```

Lists sizes from packet

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-packet-sizes-no-credentials-v2-200) | OK | PacketSizeList |  | [schema](#list-packet-sizes-no-credentials-v2-200-schema) |
| [default](#list-packet-sizes-no-credentials-v2-default) | | errorResponse |  | [schema](#list-packet-sizes-no-credentials-v2-default-schema) |

#### Responses


##### <span id="list-packet-sizes-no-credentials-v2-200"></span> 200 - PacketSizeList
Status: OK

###### <span id="list-packet-sizes-no-credentials-v2-200-schema"></span> Schema
   
  

[][PacketSizeList](#packet-size-list)

##### <span id="list-packet-sizes-no-credentials-v2-default"></span> Default Response
errorResponse

###### <span id="list-packet-sizes-no-credentials-v2-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-presets"></span> list presets (*listPresets*)

```
GET /api/v2/presets
```

Lists presets

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| disabled | `query` | boolean | `bool` |  |  |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-presets-200) | OK | PresetList |  | [schema](#list-presets-200-schema) |
| [401](#list-presets-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#list-presets-401-schema) |
| [403](#list-presets-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#list-presets-403-schema) |
| [default](#list-presets-default) | | errorResponse |  | [schema](#list-presets-default-schema) |

#### Responses


##### <span id="list-presets-200"></span> 200 - PresetList
Status: OK

###### <span id="list-presets-200-schema"></span> Schema
   
  

[PresetList](#preset-list)

##### <span id="list-presets-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="list-presets-401-schema"></span> Schema

##### <span id="list-presets-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="list-presets-403-schema"></span> Schema

##### <span id="list-presets-default"></span> Default Response
errorResponse

###### <span id="list-presets-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-project-etcd-backup-config"></span> list project etcd backup config (*listProjectEtcdBackupConfig*)

```
GET /api/v2/projects/{project_id}/etcdbackupconfigs
```

List etcd backup configs for a given project

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| project_id | `path` | string | `string` |  | ✓ |  |  |
| type | `query` | string | `string` |  |  |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-project-etcd-backup-config-200) | OK | EtcdBackupConfig |  | [schema](#list-project-etcd-backup-config-200-schema) |
| [401](#list-project-etcd-backup-config-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#list-project-etcd-backup-config-401-schema) |
| [403](#list-project-etcd-backup-config-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#list-project-etcd-backup-config-403-schema) |
| [default](#list-project-etcd-backup-config-default) | | errorResponse |  | [schema](#list-project-etcd-backup-config-default-schema) |

#### Responses


##### <span id="list-project-etcd-backup-config-200"></span> 200 - EtcdBackupConfig
Status: OK

###### <span id="list-project-etcd-backup-config-200-schema"></span> Schema
   
  

[][EtcdBackupConfig](#etcd-backup-config)

##### <span id="list-project-etcd-backup-config-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="list-project-etcd-backup-config-401-schema"></span> Schema

##### <span id="list-project-etcd-backup-config-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="list-project-etcd-backup-config-403-schema"></span> Schema

##### <span id="list-project-etcd-backup-config-default"></span> Default Response
errorResponse

###### <span id="list-project-etcd-backup-config-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-project-etcd-restore"></span> list project etcd restore (*listProjectEtcdRestore*)

```
GET /api/v2/projects/{project_id}/etcdrestores
```

List etcd backup restores for a given project

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-project-etcd-restore-200) | OK | EtcdRestore |  | [schema](#list-project-etcd-restore-200-schema) |
| [401](#list-project-etcd-restore-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#list-project-etcd-restore-401-schema) |
| [403](#list-project-etcd-restore-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#list-project-etcd-restore-403-schema) |
| [default](#list-project-etcd-restore-default) | | errorResponse |  | [schema](#list-project-etcd-restore-default-schema) |

#### Responses


##### <span id="list-project-etcd-restore-200"></span> 200 - EtcdRestore
Status: OK

###### <span id="list-project-etcd-restore-200-schema"></span> Schema
   
  

[][EtcdRestore](#etcd-restore)

##### <span id="list-project-etcd-restore-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="list-project-etcd-restore-401-schema"></span> Schema

##### <span id="list-project-etcd-restore-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="list-project-etcd-restore-403-schema"></span> Schema

##### <span id="list-project-etcd-restore-default"></span> Default Response
errorResponse

###### <span id="list-project-etcd-restore-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-projects"></span> Lists projects that an authenticated user is a member of. (*listProjects*)

```
GET /api/v1/projects
```

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| displayAll | `query` | boolean | `bool` |  |  |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-projects-200) | OK | Project |  | [schema](#list-projects-200-schema) |
| [401](#list-projects-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#list-projects-401-schema) |
| [409](#list-projects-409) | Conflict | EmptyResponse is a empty response |  | [schema](#list-projects-409-schema) |
| [default](#list-projects-default) | | errorResponse |  | [schema](#list-projects-default-schema) |

#### Responses


##### <span id="list-projects-200"></span> 200 - Project
Status: OK

###### <span id="list-projects-200-schema"></span> Schema
   
  

[][Project](#project)

##### <span id="list-projects-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="list-projects-401-schema"></span> Schema

##### <span id="list-projects-409"></span> 409 - EmptyResponse is a empty response
Status: Conflict

###### <span id="list-projects-409-schema"></span> Schema

##### <span id="list-projects-default"></span> Default Response
errorResponse

###### <span id="list-projects-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-provider-presets"></span> list provider presets (*listProviderPresets*)

```
GET /api/v2/providers/{provider_name}/presets
```

Lists presets for the provider

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| provider_name | `path` | string | `string` |  | ✓ |  |  |
| datacenter | `query` | string | `string` |  |  |  |  |
| disabled | `query` | boolean | `bool` |  |  |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-provider-presets-200) | OK | PresetList |  | [schema](#list-provider-presets-200-schema) |
| [401](#list-provider-presets-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#list-provider-presets-401-schema) |
| [403](#list-provider-presets-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#list-provider-presets-403-schema) |
| [default](#list-provider-presets-default) | | errorResponse |  | [schema](#list-provider-presets-default-schema) |

#### Responses


##### <span id="list-provider-presets-200"></span> 200 - PresetList
Status: OK

###### <span id="list-provider-presets-200-schema"></span> Schema
   
  

[PresetList](#preset-list)

##### <span id="list-provider-presets-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="list-provider-presets-401-schema"></span> Schema

##### <span id="list-provider-presets-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="list-provider-presets-403-schema"></span> Schema

##### <span id="list-provider-presets-default"></span> Default Response
errorResponse

###### <span id="list-provider-presets-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-role"></span> list role (*listRole*)

```
GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/roles
```

Lists all Roles

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| dc | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-role-200) | OK | Role |  | [schema](#list-role-200-schema) |
| [401](#list-role-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#list-role-401-schema) |
| [403](#list-role-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#list-role-403-schema) |
| [default](#list-role-default) | | errorResponse |  | [schema](#list-role-default-schema) |

#### Responses


##### <span id="list-role-200"></span> 200 - Role
Status: OK

###### <span id="list-role-200-schema"></span> Schema
   
  

[][Role](#role)

##### <span id="list-role-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="list-role-401-schema"></span> Schema

##### <span id="list-role-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="list-role-403-schema"></span> Schema

##### <span id="list-role-default"></span> Default Response
errorResponse

###### <span id="list-role-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-role-binding"></span> list role binding (*listRoleBinding*)

```
GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/bindings
```

List role binding

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| dc | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-role-binding-200) | OK | RoleBinding |  | [schema](#list-role-binding-200-schema) |
| [401](#list-role-binding-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#list-role-binding-401-schema) |
| [403](#list-role-binding-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#list-role-binding-403-schema) |
| [default](#list-role-binding-default) | | errorResponse |  | [schema](#list-role-binding-default-schema) |

#### Responses


##### <span id="list-role-binding-200"></span> 200 - RoleBinding
Status: OK

###### <span id="list-role-binding-200-schema"></span> Schema
   
  

[][RoleBinding](#role-binding)

##### <span id="list-role-binding-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="list-role-binding-401-schema"></span> Schema

##### <span id="list-role-binding-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="list-role-binding-403-schema"></span> Schema

##### <span id="list-role-binding-default"></span> Default Response
errorResponse

###### <span id="list-role-binding-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-role-binding-v2"></span> list role binding v2 (*listRoleBindingV2*)

```
GET /api/v2/projects/{project_id}/clusters/{cluster_id}/bindings
```

List role binding

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-role-binding-v2-200) | OK | RoleBinding |  | [schema](#list-role-binding-v2-200-schema) |
| [401](#list-role-binding-v2-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#list-role-binding-v2-401-schema) |
| [403](#list-role-binding-v2-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#list-role-binding-v2-403-schema) |
| [default](#list-role-binding-v2-default) | | errorResponse |  | [schema](#list-role-binding-v2-default-schema) |

#### Responses


##### <span id="list-role-binding-v2-200"></span> 200 - RoleBinding
Status: OK

###### <span id="list-role-binding-v2-200-schema"></span> Schema
   
  

[][RoleBinding](#role-binding)

##### <span id="list-role-binding-v2-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="list-role-binding-v2-401-schema"></span> Schema

##### <span id="list-role-binding-v2-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="list-role-binding-v2-403-schema"></span> Schema

##### <span id="list-role-binding-v2-default"></span> Default Response
errorResponse

###### <span id="list-role-binding-v2-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-role-names"></span> list role names (*listRoleNames*)

```
GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/rolenames
```

Lists all Role names with namespaces

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| dc | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-role-names-200) | OK | RoleName |  | [schema](#list-role-names-200-schema) |
| [401](#list-role-names-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#list-role-names-401-schema) |
| [403](#list-role-names-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#list-role-names-403-schema) |
| [default](#list-role-names-default) | | errorResponse |  | [schema](#list-role-names-default-schema) |

#### Responses


##### <span id="list-role-names-200"></span> 200 - RoleName
Status: OK

###### <span id="list-role-names-200-schema"></span> Schema
   
  

[][RoleName](#role-name)

##### <span id="list-role-names-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="list-role-names-401-schema"></span> Schema

##### <span id="list-role-names-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="list-role-names-403-schema"></span> Schema

##### <span id="list-role-names-default"></span> Default Response
errorResponse

###### <span id="list-role-names-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-role-names-v2"></span> list role names v2 (*listRoleNamesV2*)

```
GET /api/v2/projects/{project_id}/clusters/{cluster_id}/rolenames
```

Lists all Role names with namespaces

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-role-names-v2-200) | OK | RoleName |  | [schema](#list-role-names-v2-200-schema) |
| [401](#list-role-names-v2-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#list-role-names-v2-401-schema) |
| [403](#list-role-names-v2-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#list-role-names-v2-403-schema) |
| [default](#list-role-names-v2-default) | | errorResponse |  | [schema](#list-role-names-v2-default-schema) |

#### Responses


##### <span id="list-role-names-v2-200"></span> 200 - RoleName
Status: OK

###### <span id="list-role-names-v2-200-schema"></span> Schema
   
  

[][RoleName](#role-name)

##### <span id="list-role-names-v2-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="list-role-names-v2-401-schema"></span> Schema

##### <span id="list-role-names-v2-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="list-role-names-v2-403-schema"></span> Schema

##### <span id="list-role-names-v2-default"></span> Default Response
errorResponse

###### <span id="list-role-names-v2-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-role-v2"></span> list role v2 (*listRoleV2*)

```
GET /api/v2/projects/{project_id}/clusters/{cluster_id}/roles
```

Lists all Roles

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-role-v2-200) | OK | Role |  | [schema](#list-role-v2-200-schema) |
| [401](#list-role-v2-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#list-role-v2-401-schema) |
| [403](#list-role-v2-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#list-role-v2-403-schema) |
| [default](#list-role-v2-default) | | errorResponse |  | [schema](#list-role-v2-default-schema) |

#### Responses


##### <span id="list-role-v2-200"></span> 200 - Role
Status: OK

###### <span id="list-role-v2-200-schema"></span> Schema
   
  

[][Role](#role)

##### <span id="list-role-v2-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="list-role-v2-401-schema"></span> Schema

##### <span id="list-role-v2-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="list-role-v2-403-schema"></span> Schema

##### <span id="list-role-v2-default"></span> Default Response
errorResponse

###### <span id="list-role-v2-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-rule-groups"></span> list rule groups (*listRuleGroups*)

```
GET /api/v2/projects/{project_id}/clusters/{cluster_id}/rulegroups
```

Lists rule groups that belong to the given cluster

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |
| type | `query` | string | `string` |  |  |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-rule-groups-200) | OK | RuleGroup |  | [schema](#list-rule-groups-200-schema) |
| [401](#list-rule-groups-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#list-rule-groups-401-schema) |
| [403](#list-rule-groups-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#list-rule-groups-403-schema) |
| [default](#list-rule-groups-default) | | errorResponse |  | [schema](#list-rule-groups-default-schema) |

#### Responses


##### <span id="list-rule-groups-200"></span> 200 - RuleGroup
Status: OK

###### <span id="list-rule-groups-200-schema"></span> Schema
   
  

[][RuleGroup](#rule-group)

##### <span id="list-rule-groups-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="list-rule-groups-401-schema"></span> Schema

##### <span id="list-rule-groups-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="list-rule-groups-403-schema"></span> Schema

##### <span id="list-rule-groups-default"></span> Default Response
errorResponse

###### <span id="list-rule-groups-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-ssh-keys"></span> Lists SSH Keys that belong to the given project. (*listSSHKeys*)

```
GET /api/v1/projects/{project_id}/sshkeys
```

The returned collection is sorted by creation timestamp.

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-ssh-keys-200) | OK | SSHKey |  | [schema](#list-ssh-keys-200-schema) |
| [401](#list-ssh-keys-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#list-ssh-keys-401-schema) |
| [403](#list-ssh-keys-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#list-ssh-keys-403-schema) |
| [default](#list-ssh-keys-default) | | errorResponse |  | [schema](#list-ssh-keys-default-schema) |

#### Responses


##### <span id="list-ssh-keys-200"></span> 200 - SSHKey
Status: OK

###### <span id="list-ssh-keys-200-schema"></span> Schema
   
  

[][SSHKey](#ssh-key)

##### <span id="list-ssh-keys-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="list-ssh-keys-401-schema"></span> Schema

##### <span id="list-ssh-keys-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="list-ssh-keys-403-schema"></span> Schema

##### <span id="list-ssh-keys-default"></span> Default Response
errorResponse

###### <span id="list-ssh-keys-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-ssh-keys-assigned-to-cluster"></span> list SSH keys assigned to cluster (*listSSHKeysAssignedToCluster*)

```
GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/sshkeys
```

Lists ssh keys that are assigned to the cluster
The returned collection is sorted by creation timestamp.

#### Consumes
  * application/json

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| dc | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-ssh-keys-assigned-to-cluster-200) | OK | SSHKey |  | [schema](#list-ssh-keys-assigned-to-cluster-200-schema) |
| [401](#list-ssh-keys-assigned-to-cluster-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#list-ssh-keys-assigned-to-cluster-401-schema) |
| [403](#list-ssh-keys-assigned-to-cluster-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#list-ssh-keys-assigned-to-cluster-403-schema) |
| [default](#list-ssh-keys-assigned-to-cluster-default) | | errorResponse |  | [schema](#list-ssh-keys-assigned-to-cluster-default-schema) |

#### Responses


##### <span id="list-ssh-keys-assigned-to-cluster-200"></span> 200 - SSHKey
Status: OK

###### <span id="list-ssh-keys-assigned-to-cluster-200-schema"></span> Schema
   
  

[][SSHKey](#ssh-key)

##### <span id="list-ssh-keys-assigned-to-cluster-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="list-ssh-keys-assigned-to-cluster-401-schema"></span> Schema

##### <span id="list-ssh-keys-assigned-to-cluster-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="list-ssh-keys-assigned-to-cluster-403-schema"></span> Schema

##### <span id="list-ssh-keys-assigned-to-cluster-default"></span> Default Response
errorResponse

###### <span id="list-ssh-keys-assigned-to-cluster-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-ssh-keys-assigned-to-cluster-v2"></span> list SSH keys assigned to cluster v2 (*listSSHKeysAssignedToClusterV2*)

```
GET /api/v2/projects/{project_id}/clusters/{cluster_id}/sshkeys
```

Lists ssh keys that are assigned to the cluster
The returned collection is sorted by creation timestamp.

#### Consumes
  * application/json

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-ssh-keys-assigned-to-cluster-v2-200) | OK | SSHKey |  | [schema](#list-ssh-keys-assigned-to-cluster-v2-200-schema) |
| [401](#list-ssh-keys-assigned-to-cluster-v2-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#list-ssh-keys-assigned-to-cluster-v2-401-schema) |
| [403](#list-ssh-keys-assigned-to-cluster-v2-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#list-ssh-keys-assigned-to-cluster-v2-403-schema) |
| [default](#list-ssh-keys-assigned-to-cluster-v2-default) | | errorResponse |  | [schema](#list-ssh-keys-assigned-to-cluster-v2-default-schema) |

#### Responses


##### <span id="list-ssh-keys-assigned-to-cluster-v2-200"></span> 200 - SSHKey
Status: OK

###### <span id="list-ssh-keys-assigned-to-cluster-v2-200-schema"></span> Schema
   
  

[][SSHKey](#ssh-key)

##### <span id="list-ssh-keys-assigned-to-cluster-v2-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="list-ssh-keys-assigned-to-cluster-v2-401-schema"></span> Schema

##### <span id="list-ssh-keys-assigned-to-cluster-v2-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="list-ssh-keys-assigned-to-cluster-v2-403-schema"></span> Schema

##### <span id="list-ssh-keys-assigned-to-cluster-v2-default"></span> Default Response
errorResponse

###### <span id="list-ssh-keys-assigned-to-cluster-v2-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-seed-names"></span> list seed names (*listSeedNames*)

```
GET /api/v1/seed
```

#### Produces
  * application/json

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-seed-names-200) | OK | SeedNamesList |  | [schema](#list-seed-names-200-schema) |
| [default](#list-seed-names-default) | | errorResponse |  | [schema](#list-seed-names-default-schema) |

#### Responses


##### <span id="list-seed-names-200"></span> 200 - SeedNamesList
Status: OK

###### <span id="list-seed-names-200-schema"></span> Schema
   
  


 [SeedNamesList](#seed-names-list)

##### <span id="list-seed-names-default"></span> Default Response
errorResponse

###### <span id="list-seed-names-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-seeds"></span> Returns all seeds from the CRDs. (*listSeeds*)

```
GET /api/v1/admin/seeds
```

#### Produces
  * application/json

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-seeds-200) | OK | Seed |  | [schema](#list-seeds-200-schema) |
| [401](#list-seeds-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#list-seeds-401-schema) |
| [403](#list-seeds-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#list-seeds-403-schema) |
| [default](#list-seeds-default) | | errorResponse |  | [schema](#list-seeds-default-schema) |

#### Responses


##### <span id="list-seeds-200"></span> 200 - Seed
Status: OK

###### <span id="list-seeds-200-schema"></span> Schema
   
  

[][Seed](#seed)

##### <span id="list-seeds-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="list-seeds-401-schema"></span> Schema

##### <span id="list-seeds-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="list-seeds-403-schema"></span> Schema

##### <span id="list-seeds-default"></span> Default Response
errorResponse

###### <span id="list-seeds-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-service-account-tokens"></span> list service account tokens (*listServiceAccountTokens*)

```
GET /api/v1/projects/{project_id}/serviceaccounts/{serviceaccount_id}/tokens
```

List tokens for the given service account

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| project_id | `path` | string | `string` |  | ✓ |  |  |
| serviceaccount_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-service-account-tokens-200) | OK | PublicServiceAccountToken |  | [schema](#list-service-account-tokens-200-schema) |
| [401](#list-service-account-tokens-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#list-service-account-tokens-401-schema) |
| [403](#list-service-account-tokens-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#list-service-account-tokens-403-schema) |
| [default](#list-service-account-tokens-default) | | errorResponse |  | [schema](#list-service-account-tokens-default-schema) |

#### Responses


##### <span id="list-service-account-tokens-200"></span> 200 - PublicServiceAccountToken
Status: OK

###### <span id="list-service-account-tokens-200-schema"></span> Schema
   
  

[][PublicServiceAccountToken](#public-service-account-token)

##### <span id="list-service-account-tokens-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="list-service-account-tokens-401-schema"></span> Schema

##### <span id="list-service-account-tokens-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="list-service-account-tokens-403-schema"></span> Schema

##### <span id="list-service-account-tokens-default"></span> Default Response
errorResponse

###### <span id="list-service-account-tokens-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-service-accounts"></span> list service accounts (*listServiceAccounts*)

```
GET /api/v1/projects/{project_id}/serviceaccounts
```

List Service Accounts for the given project

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-service-accounts-200) | OK | ServiceAccount |  | [schema](#list-service-accounts-200-schema) |
| [401](#list-service-accounts-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#list-service-accounts-401-schema) |
| [403](#list-service-accounts-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#list-service-accounts-403-schema) |
| [default](#list-service-accounts-default) | | errorResponse |  | [schema](#list-service-accounts-default-schema) |

#### Responses


##### <span id="list-service-accounts-200"></span> 200 - ServiceAccount
Status: OK

###### <span id="list-service-accounts-200-schema"></span> Schema
   
  

[][ServiceAccount](#service-account)

##### <span id="list-service-accounts-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="list-service-accounts-401-schema"></span> Schema

##### <span id="list-service-accounts-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="list-service-accounts-403-schema"></span> Schema

##### <span id="list-service-accounts-default"></span> Default Response
errorResponse

###### <span id="list-service-accounts-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-system-labels"></span> list system labels (*listSystemLabels*)

```
GET /api/v1/labels/system
```

List restricted system labels

#### Produces
  * application/json

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-system-labels-200) | OK | ResourceLabelMap |  | [schema](#list-system-labels-200-schema) |
| [401](#list-system-labels-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#list-system-labels-401-schema) |
| [403](#list-system-labels-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#list-system-labels-403-schema) |
| [default](#list-system-labels-default) | | errorResponse |  | [schema](#list-system-labels-default-schema) |

#### Responses


##### <span id="list-system-labels-200"></span> 200 - ResourceLabelMap
Status: OK

###### <span id="list-system-labels-200-schema"></span> Schema
   
  

[ResourceLabelMap](#resource-label-map)

##### <span id="list-system-labels-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="list-system-labels-401-schema"></span> Schema

##### <span id="list-system-labels-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="list-system-labels-403-schema"></span> Schema

##### <span id="list-system-labels-default"></span> Default Response
errorResponse

###### <span id="list-system-labels-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-v-sphere-datastores"></span> list v sphere datastores (*listVSphereDatastores*)

```
GET /api/v2/providers/vsphere/datastores
```

Lists datastores from vsphere datacenter

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| Credential | `header` | string | `string` |  |  |  |  |
| DatacenterName | `header` | string | `string` |  |  |  |  |
| Password | `header` | string | `string` |  |  |  |  |
| Username | `header` | string | `string` |  |  |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-v-sphere-datastores-200) | OK | VSphereDatastoreList |  | [schema](#list-v-sphere-datastores-200-schema) |
| [default](#list-v-sphere-datastores-default) | | errorResponse |  | [schema](#list-v-sphere-datastores-default-schema) |

#### Responses


##### <span id="list-v-sphere-datastores-200"></span> 200 - VSphereDatastoreList
Status: OK

###### <span id="list-v-sphere-datastores-200-schema"></span> Schema
   
  

[][VSphereDatastoreList](#v-sphere-datastore-list)

##### <span id="list-v-sphere-datastores-default"></span> Default Response
errorResponse

###### <span id="list-v-sphere-datastores-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-v-sphere-folders"></span> list v sphere folders (*listVSphereFolders*)

```
GET /api/v1/providers/vsphere/folders
```

Lists folders from vsphere datacenter

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| Credential | `header` | string | `string` |  |  |  |  |
| DatacenterName | `header` | string | `string` |  |  |  |  |
| Password | `header` | string | `string` |  |  |  |  |
| Username | `header` | string | `string` |  |  |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-v-sphere-folders-200) | OK | VSphereFolder |  | [schema](#list-v-sphere-folders-200-schema) |
| [default](#list-v-sphere-folders-default) | | errorResponse |  | [schema](#list-v-sphere-folders-default-schema) |

#### Responses


##### <span id="list-v-sphere-folders-200"></span> 200 - VSphereFolder
Status: OK

###### <span id="list-v-sphere-folders-200-schema"></span> Schema
   
  

[][VSphereFolder](#v-sphere-folder)

##### <span id="list-v-sphere-folders-default"></span> Default Response
errorResponse

###### <span id="list-v-sphere-folders-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-v-sphere-folders-no-credentials"></span> list v sphere folders no credentials (*listVSphereFoldersNoCredentials*)

```
GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/providers/vsphere/folders
```

Lists folders from vsphere datacenter

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| dc | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-v-sphere-folders-no-credentials-200) | OK | VSphereFolder |  | [schema](#list-v-sphere-folders-no-credentials-200-schema) |
| [default](#list-v-sphere-folders-no-credentials-default) | | errorResponse |  | [schema](#list-v-sphere-folders-no-credentials-default-schema) |

#### Responses


##### <span id="list-v-sphere-folders-no-credentials-200"></span> 200 - VSphereFolder
Status: OK

###### <span id="list-v-sphere-folders-no-credentials-200-schema"></span> Schema
   
  

[][VSphereFolder](#v-sphere-folder)

##### <span id="list-v-sphere-folders-no-credentials-default"></span> Default Response
errorResponse

###### <span id="list-v-sphere-folders-no-credentials-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-v-sphere-folders-no-credentials-v2"></span> list v sphere folders no credentials v2 (*listVSphereFoldersNoCredentialsV2*)

```
GET /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/vsphere/folders
```

Lists folders from vsphere datacenter

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-v-sphere-folders-no-credentials-v2-200) | OK | VSphereFolder |  | [schema](#list-v-sphere-folders-no-credentials-v2-200-schema) |
| [default](#list-v-sphere-folders-no-credentials-v2-default) | | errorResponse |  | [schema](#list-v-sphere-folders-no-credentials-v2-default-schema) |

#### Responses


##### <span id="list-v-sphere-folders-no-credentials-v2-200"></span> 200 - VSphereFolder
Status: OK

###### <span id="list-v-sphere-folders-no-credentials-v2-200-schema"></span> Schema
   
  

[][VSphereFolder](#v-sphere-folder)

##### <span id="list-v-sphere-folders-no-credentials-v2-default"></span> Default Response
errorResponse

###### <span id="list-v-sphere-folders-no-credentials-v2-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-v-sphere-networks"></span> list v sphere networks (*listVSphereNetworks*)

```
GET /api/v1/providers/vsphere/networks
```

Lists networks from vsphere datacenter

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| Credential | `header` | string | `string` |  |  |  |  |
| DatacenterName | `header` | string | `string` |  |  |  |  |
| Password | `header` | string | `string` |  |  |  |  |
| Username | `header` | string | `string` |  |  |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-v-sphere-networks-200) | OK | VSphereNetwork |  | [schema](#list-v-sphere-networks-200-schema) |
| [default](#list-v-sphere-networks-default) | | errorResponse |  | [schema](#list-v-sphere-networks-default-schema) |

#### Responses


##### <span id="list-v-sphere-networks-200"></span> 200 - VSphereNetwork
Status: OK

###### <span id="list-v-sphere-networks-200-schema"></span> Schema
   
  

[][VSphereNetwork](#v-sphere-network)

##### <span id="list-v-sphere-networks-default"></span> Default Response
errorResponse

###### <span id="list-v-sphere-networks-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-v-sphere-networks-no-credentials"></span> list v sphere networks no credentials (*listVSphereNetworksNoCredentials*)

```
GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/providers/vsphere/networks
```

Lists networks from vsphere datacenter

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| dc | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-v-sphere-networks-no-credentials-200) | OK | VSphereNetwork |  | [schema](#list-v-sphere-networks-no-credentials-200-schema) |
| [default](#list-v-sphere-networks-no-credentials-default) | | errorResponse |  | [schema](#list-v-sphere-networks-no-credentials-default-schema) |

#### Responses


##### <span id="list-v-sphere-networks-no-credentials-200"></span> 200 - VSphereNetwork
Status: OK

###### <span id="list-v-sphere-networks-no-credentials-200-schema"></span> Schema
   
  

[][VSphereNetwork](#v-sphere-network)

##### <span id="list-v-sphere-networks-no-credentials-default"></span> Default Response
errorResponse

###### <span id="list-v-sphere-networks-no-credentials-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-v-sphere-networks-no-credentials-v2"></span> list v sphere networks no credentials v2 (*listVSphereNetworksNoCredentialsV2*)

```
GET /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/vsphere/networks
```

Lists networks from vsphere datacenter

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-v-sphere-networks-no-credentials-v2-200) | OK | VSphereNetwork |  | [schema](#list-v-sphere-networks-no-credentials-v2-200-schema) |
| [default](#list-v-sphere-networks-no-credentials-v2-default) | | errorResponse |  | [schema](#list-v-sphere-networks-no-credentials-v2-default-schema) |

#### Responses


##### <span id="list-v-sphere-networks-no-credentials-v2-200"></span> 200 - VSphereNetwork
Status: OK

###### <span id="list-v-sphere-networks-no-credentials-v2-200-schema"></span> Schema
   
  

[][VSphereNetwork](#v-sphere-network)

##### <span id="list-v-sphere-networks-no-credentials-v2-default"></span> Default Response
errorResponse

###### <span id="list-v-sphere-networks-no-credentials-v2-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="list-versions-by-provider"></span> list versions by provider (*listVersionsByProvider*)

```
GET /api/v2/providers/{provider_name}/versions
```

Lists all versions which don't result in automatic updates for a given provider

#### Consumes
  * application/json

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| provider_name | `path` | string | `string` |  | ✓ |  |  |
| type | `query` | string | `string` |  |  |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-versions-by-provider-200) | OK | VersionList |  | [schema](#list-versions-by-provider-200-schema) |
| [401](#list-versions-by-provider-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#list-versions-by-provider-401-schema) |
| [403](#list-versions-by-provider-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#list-versions-by-provider-403-schema) |
| [default](#list-versions-by-provider-default) | | errorResponse |  | [schema](#list-versions-by-provider-default-schema) |

#### Responses


##### <span id="list-versions-by-provider-200"></span> 200 - VersionList
Status: OK

###### <span id="list-versions-by-provider-200-schema"></span> Schema
   
  


 [VersionList](#version-list)

##### <span id="list-versions-by-provider-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="list-versions-by-provider-401-schema"></span> Schema

##### <span id="list-versions-by-provider-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="list-versions-by-provider-403-schema"></span> Schema

##### <span id="list-versions-by-provider-default"></span> Default Response
errorResponse

###### <span id="list-versions-by-provider-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="logout-current-user"></span> Adds current authorization bearer token to the blacklist. (*logoutCurrentUser*)

```
POST /api/v1/me/logout
```

Enforces user to login again with the new token.

#### Produces
  * application/json

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#logout-current-user-200) | OK | EmptyResponse is a empty response |  | [schema](#logout-current-user-200-schema) |
| [401](#logout-current-user-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#logout-current-user-401-schema) |
| [default](#logout-current-user-default) | | errorResponse |  | [schema](#logout-current-user-default-schema) |

#### Responses


##### <span id="logout-current-user-200"></span> 200 - EmptyResponse is a empty response
Status: OK

###### <span id="logout-current-user-200-schema"></span> Schema

##### <span id="logout-current-user-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="logout-current-user-401-schema"></span> Schema

##### <span id="logout-current-user-default"></span> Default Response
errorResponse

###### <span id="logout-current-user-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="migrate-cluster-to-external-c-c-m"></span> migrate cluster to external c c m (*migrateClusterToExternalCCM*)

```
POST /api/v2/projects/{project_id}/clusters/{cluster_id}/externalccmmigration
```

Enable the migration to the external CCM for the given cluster

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#migrate-cluster-to-external-c-c-m-200) | OK | EmptyResponse is a empty response |  | [schema](#migrate-cluster-to-external-c-c-m-200-schema) |
| [401](#migrate-cluster-to-external-c-c-m-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#migrate-cluster-to-external-c-c-m-401-schema) |
| [403](#migrate-cluster-to-external-c-c-m-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#migrate-cluster-to-external-c-c-m-403-schema) |
| [default](#migrate-cluster-to-external-c-c-m-default) | | errorResponse |  | [schema](#migrate-cluster-to-external-c-c-m-default-schema) |

#### Responses


##### <span id="migrate-cluster-to-external-c-c-m-200"></span> 200 - EmptyResponse is a empty response
Status: OK

###### <span id="migrate-cluster-to-external-c-c-m-200-schema"></span> Schema

##### <span id="migrate-cluster-to-external-c-c-m-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="migrate-cluster-to-external-c-c-m-401-schema"></span> Schema

##### <span id="migrate-cluster-to-external-c-c-m-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="migrate-cluster-to-external-c-c-m-403-schema"></span> Schema

##### <span id="migrate-cluster-to-external-c-c-m-default"></span> Default Response
errorResponse

###### <span id="migrate-cluster-to-external-c-c-m-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="patch-addon"></span> Patches an addon that is assigned to the given cluster. (*patchAddon*)

```
PATCH /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/addons/{addon_id}
```

#### Consumes
  * application/json

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| addon_id | `path` | string | `string` |  | ✓ |  |  |
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| dc | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |
| Body | `body` | [Addon](#addon) | `models.Addon` | |  | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#patch-addon-200) | OK | Addon |  | [schema](#patch-addon-200-schema) |
| [401](#patch-addon-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#patch-addon-401-schema) |
| [403](#patch-addon-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#patch-addon-403-schema) |
| [default](#patch-addon-default) | | errorResponse |  | [schema](#patch-addon-default-schema) |

#### Responses


##### <span id="patch-addon-200"></span> 200 - Addon
Status: OK

###### <span id="patch-addon-200-schema"></span> Schema
   
  

[Addon](#addon)

##### <span id="patch-addon-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="patch-addon-401-schema"></span> Schema

##### <span id="patch-addon-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="patch-addon-403-schema"></span> Schema

##### <span id="patch-addon-default"></span> Default Response
errorResponse

###### <span id="patch-addon-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="patch-addon-v2"></span> Patches an addon that is assigned to the given cluster. (*patchAddonV2*)

```
PATCH /api/v2/projects/{project_id}/clusters/{cluster_id}/addons/{addon_id}
```

#### Consumes
  * application/json

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| addon_id | `path` | string | `string` |  | ✓ |  |  |
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |
| Body | `body` | [Addon](#addon) | `models.Addon` | |  | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#patch-addon-v2-200) | OK | Addon |  | [schema](#patch-addon-v2-200-schema) |
| [401](#patch-addon-v2-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#patch-addon-v2-401-schema) |
| [403](#patch-addon-v2-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#patch-addon-v2-403-schema) |
| [default](#patch-addon-v2-default) | | errorResponse |  | [schema](#patch-addon-v2-default-schema) |

#### Responses


##### <span id="patch-addon-v2-200"></span> 200 - Addon
Status: OK

###### <span id="patch-addon-v2-200-schema"></span> Schema
   
  

[Addon](#addon)

##### <span id="patch-addon-v2-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="patch-addon-v2-401-schema"></span> Schema

##### <span id="patch-addon-v2-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="patch-addon-v2-403-schema"></span> Schema

##### <span id="patch-addon-v2-default"></span> Default Response
errorResponse

###### <span id="patch-addon-v2-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="patch-allowed-registry"></span> patch allowed registry (*patchAllowedRegistry*)

```
PATCH /api/v2/allowedregistries/{allowed_registry}
```

Patch a specified allowed registry

#### Consumes
  * application/json

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| allowed_registry | `path` | string | `string` |  | ✓ |  |  |
| Patch | `body` | [interface{}](#interface) | `interface{}` | |  | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#patch-allowed-registry-200) | OK | ConstraintTemplate |  | [schema](#patch-allowed-registry-200-schema) |
| [401](#patch-allowed-registry-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#patch-allowed-registry-401-schema) |
| [403](#patch-allowed-registry-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#patch-allowed-registry-403-schema) |
| [default](#patch-allowed-registry-default) | | errorResponse |  | [schema](#patch-allowed-registry-default-schema) |

#### Responses


##### <span id="patch-allowed-registry-200"></span> 200 - ConstraintTemplate
Status: OK

###### <span id="patch-allowed-registry-200-schema"></span> Schema
   
  

[ConstraintTemplate](#constraint-template)

##### <span id="patch-allowed-registry-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="patch-allowed-registry-401-schema"></span> Schema

##### <span id="patch-allowed-registry-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="patch-allowed-registry-403-schema"></span> Schema

##### <span id="patch-allowed-registry-default"></span> Default Response
errorResponse

###### <span id="patch-allowed-registry-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="patch-cluster"></span> Patches the given cluster using JSON Merge Patch method (https://tools.ietf.org/html/rfc7396). (*patchCluster*)

```
PATCH /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}
```

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| dc | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |
| Patch | `body` | [interface{}](#interface) | `interface{}` | |  | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#patch-cluster-200) | OK | Cluster |  | [schema](#patch-cluster-200-schema) |
| [401](#patch-cluster-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#patch-cluster-401-schema) |
| [403](#patch-cluster-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#patch-cluster-403-schema) |
| [default](#patch-cluster-default) | | errorResponse |  | [schema](#patch-cluster-default-schema) |

#### Responses


##### <span id="patch-cluster-200"></span> 200 - Cluster
Status: OK

###### <span id="patch-cluster-200-schema"></span> Schema
   
  

[Cluster](#cluster)

##### <span id="patch-cluster-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="patch-cluster-401-schema"></span> Schema

##### <span id="patch-cluster-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="patch-cluster-403-schema"></span> Schema

##### <span id="patch-cluster-default"></span> Default Response
errorResponse

###### <span id="patch-cluster-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="patch-cluster-role"></span> patch cluster role (*patchClusterRole*)

```
PATCH /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/clusterroles/{role_id}
```

Patch the cluster role with the given name

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| dc | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |
| role_id | `path` | string | `string` |  | ✓ |  |  |
| Patch | `body` | [interface{}](#interface) | `interface{}` | |  | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#patch-cluster-role-200) | OK | ClusterRole |  | [schema](#patch-cluster-role-200-schema) |
| [401](#patch-cluster-role-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#patch-cluster-role-401-schema) |
| [403](#patch-cluster-role-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#patch-cluster-role-403-schema) |
| [default](#patch-cluster-role-default) | | errorResponse |  | [schema](#patch-cluster-role-default-schema) |

#### Responses


##### <span id="patch-cluster-role-200"></span> 200 - ClusterRole
Status: OK

###### <span id="patch-cluster-role-200-schema"></span> Schema
   
  

[ClusterRole](#cluster-role)

##### <span id="patch-cluster-role-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="patch-cluster-role-401-schema"></span> Schema

##### <span id="patch-cluster-role-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="patch-cluster-role-403-schema"></span> Schema

##### <span id="patch-cluster-role-default"></span> Default Response
errorResponse

###### <span id="patch-cluster-role-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="patch-cluster-v2"></span> Patches the given cluster using JSON Merge Patch method (https://tools.ietf.org/html/rfc7396). (*patchClusterV2*)

```
PATCH /api/v2/projects/{project_id}/clusters/{cluster_id}
```

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |
| Patch | `body` | [interface{}](#interface) | `interface{}` | |  | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#patch-cluster-v2-200) | OK | Cluster |  | [schema](#patch-cluster-v2-200-schema) |
| [401](#patch-cluster-v2-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#patch-cluster-v2-401-schema) |
| [403](#patch-cluster-v2-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#patch-cluster-v2-403-schema) |
| [default](#patch-cluster-v2-default) | | errorResponse |  | [schema](#patch-cluster-v2-default-schema) |

#### Responses


##### <span id="patch-cluster-v2-200"></span> 200 - Cluster
Status: OK

###### <span id="patch-cluster-v2-200-schema"></span> Schema
   
  

[Cluster](#cluster)

##### <span id="patch-cluster-v2-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="patch-cluster-v2-401-schema"></span> Schema

##### <span id="patch-cluster-v2-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="patch-cluster-v2-403-schema"></span> Schema

##### <span id="patch-cluster-v2-default"></span> Default Response
errorResponse

###### <span id="patch-cluster-v2-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="patch-constraint"></span> Patches a given constraint for the specified cluster. (*patchConstraint*)

```
PATCH /api/v2/projects/{project_id}/clusters/{cluster_id}/constraints/{constraint_name}
```

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| constraint_name | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |
| Patch | `body` | [interface{}](#interface) | `interface{}` | |  | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#patch-constraint-200) | OK | Constraint |  | [schema](#patch-constraint-200-schema) |
| [401](#patch-constraint-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#patch-constraint-401-schema) |
| [403](#patch-constraint-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#patch-constraint-403-schema) |
| [default](#patch-constraint-default) | | errorResponse |  | [schema](#patch-constraint-default-schema) |

#### Responses


##### <span id="patch-constraint-200"></span> 200 - Constraint
Status: OK

###### <span id="patch-constraint-200-schema"></span> Schema
   
  

[Constraint](#constraint)

##### <span id="patch-constraint-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="patch-constraint-401-schema"></span> Schema

##### <span id="patch-constraint-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="patch-constraint-403-schema"></span> Schema

##### <span id="patch-constraint-default"></span> Default Response
errorResponse

###### <span id="patch-constraint-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="patch-constraint-template"></span> patch constraint template (*patchConstraintTemplate*)

```
PATCH /api/v2/constrainttemplates/{ct_name}
```

Patch a specified constraint template

#### Consumes
  * application/json

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| ct_name | `path` | string | `string` |  | ✓ |  |  |
| Patch | `body` | [interface{}](#interface) | `interface{}` | |  | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#patch-constraint-template-200) | OK | ConstraintTemplate |  | [schema](#patch-constraint-template-200-schema) |
| [401](#patch-constraint-template-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#patch-constraint-template-401-schema) |
| [403](#patch-constraint-template-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#patch-constraint-template-403-schema) |
| [default](#patch-constraint-template-default) | | errorResponse |  | [schema](#patch-constraint-template-default-schema) |

#### Responses


##### <span id="patch-constraint-template-200"></span> 200 - ConstraintTemplate
Status: OK

###### <span id="patch-constraint-template-200-schema"></span> Schema
   
  

[ConstraintTemplate](#constraint-template)

##### <span id="patch-constraint-template-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="patch-constraint-template-401-schema"></span> Schema

##### <span id="patch-constraint-template-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="patch-constraint-template-403-schema"></span> Schema

##### <span id="patch-constraint-template-default"></span> Default Response
errorResponse

###### <span id="patch-constraint-template-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="patch-current-user-settings"></span> Updates settings of the current user. (*patchCurrentUserSettings*)

```
PATCH /api/v1/me/settings
```

#### Consumes
  * application/json

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| Patch | `body` | [interface{}](#interface) | `interface{}` | |  | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#patch-current-user-settings-200) | OK | UserSettings |  | [schema](#patch-current-user-settings-200-schema) |
| [401](#patch-current-user-settings-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#patch-current-user-settings-401-schema) |
| [default](#patch-current-user-settings-default) | | errorResponse |  | [schema](#patch-current-user-settings-default-schema) |

#### Responses


##### <span id="patch-current-user-settings-200"></span> 200 - UserSettings
Status: OK

###### <span id="patch-current-user-settings-200-schema"></span> Schema
   
  

[UserSettings](#user-settings)

##### <span id="patch-current-user-settings-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="patch-current-user-settings-401-schema"></span> Schema

##### <span id="patch-current-user-settings-default"></span> Default Response
errorResponse

###### <span id="patch-current-user-settings-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="patch-d-c"></span> Patch the datacenter. (*patchDC*)

```
PATCH /api/v1/seed/{seed_name}/dc/{dc}
```

#### Consumes
  * application/json

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| dc | `path` | string | `string` |  | ✓ |  |  |
| seed_name | `path` | string | `string` |  | ✓ |  |  |
| Patch | `body` | [interface{}](#interface) | `interface{}` | |  | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#patch-d-c-200) | OK | Datacenter |  | [schema](#patch-d-c-200-schema) |
| [401](#patch-d-c-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#patch-d-c-401-schema) |
| [403](#patch-d-c-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#patch-d-c-403-schema) |
| [default](#patch-d-c-default) | | errorResponse |  | [schema](#patch-d-c-default-schema) |

#### Responses


##### <span id="patch-d-c-200"></span> 200 - Datacenter
Status: OK

###### <span id="patch-d-c-200-schema"></span> Schema
   
  

[Datacenter](#datacenter)

##### <span id="patch-d-c-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="patch-d-c-401-schema"></span> Schema

##### <span id="patch-d-c-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="patch-d-c-403-schema"></span> Schema

##### <span id="patch-d-c-default"></span> Default Response
errorResponse

###### <span id="patch-d-c-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="patch-default-constraint"></span> patch default constraint (*patchDefaultConstraint*)

```
PATCH /api/v2/constraints/{constraint_name}
```

Patch a specified default constraint

#### Consumes
  * application/json

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| constraint_name | `path` | string | `string` |  | ✓ |  |  |
| Patch | `body` | [interface{}](#interface) | `interface{}` | |  | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#patch-default-constraint-200) | OK | Constraint |  | [schema](#patch-default-constraint-200-schema) |
| [401](#patch-default-constraint-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#patch-default-constraint-401-schema) |
| [403](#patch-default-constraint-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#patch-default-constraint-403-schema) |
| [default](#patch-default-constraint-default) | | errorResponse |  | [schema](#patch-default-constraint-default-schema) |

#### Responses


##### <span id="patch-default-constraint-200"></span> 200 - Constraint
Status: OK

###### <span id="patch-default-constraint-200-schema"></span> Schema
   
  

[Constraint](#constraint)

##### <span id="patch-default-constraint-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="patch-default-constraint-401-schema"></span> Schema

##### <span id="patch-default-constraint-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="patch-default-constraint-403-schema"></span> Schema

##### <span id="patch-default-constraint-default"></span> Default Response
errorResponse

###### <span id="patch-default-constraint-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="patch-etcd-backup-config"></span> patch etcd backup config (*patchEtcdBackupConfig*)

```
PATCH /api/v2/projects/{project_id}/clusters/{cluster_id}/etcdbackupconfigs/{ebc_id}
```

Patches a etcd backup config for a given cluster based on its id

#### Consumes
  * application/json

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| ebc_id | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |
| Body | `body` | [EtcdBackupConfigSpec](#etcd-backup-config-spec) | `models.EtcdBackupConfigSpec` | | ✓ | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#patch-etcd-backup-config-200) | OK | EtcdBackupConfig |  | [schema](#patch-etcd-backup-config-200-schema) |
| [401](#patch-etcd-backup-config-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#patch-etcd-backup-config-401-schema) |
| [403](#patch-etcd-backup-config-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#patch-etcd-backup-config-403-schema) |
| [default](#patch-etcd-backup-config-default) | | errorResponse |  | [schema](#patch-etcd-backup-config-default-schema) |

#### Responses


##### <span id="patch-etcd-backup-config-200"></span> 200 - EtcdBackupConfig
Status: OK

###### <span id="patch-etcd-backup-config-200-schema"></span> Schema
   
  

[EtcdBackupConfig](#etcd-backup-config)

##### <span id="patch-etcd-backup-config-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="patch-etcd-backup-config-401-schema"></span> Schema

##### <span id="patch-etcd-backup-config-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="patch-etcd-backup-config-403-schema"></span> Schema

##### <span id="patch-etcd-backup-config-default"></span> Default Response
errorResponse

###### <span id="patch-etcd-backup-config-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="patch-gatekeeper-config"></span> Patches the gatekeeper config for the specified cluster. (*patchGatekeeperConfig*)

```
PATCH /api/v2/projects/{project_id}/clusters/{cluster_id}/gatekeeper/config
```

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |
| Patch | `body` | [interface{}](#interface) | `interface{}` | |  | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#patch-gatekeeper-config-200) | OK | GatekeeperConfig |  | [schema](#patch-gatekeeper-config-200-schema) |
| [401](#patch-gatekeeper-config-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#patch-gatekeeper-config-401-schema) |
| [403](#patch-gatekeeper-config-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#patch-gatekeeper-config-403-schema) |
| [default](#patch-gatekeeper-config-default) | | errorResponse |  | [schema](#patch-gatekeeper-config-default-schema) |

#### Responses


##### <span id="patch-gatekeeper-config-200"></span> 200 - GatekeeperConfig
Status: OK

###### <span id="patch-gatekeeper-config-200-schema"></span> Schema
   
  

[GatekeeperConfig](#gatekeeper-config)

##### <span id="patch-gatekeeper-config-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="patch-gatekeeper-config-401-schema"></span> Schema

##### <span id="patch-gatekeeper-config-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="patch-gatekeeper-config-403-schema"></span> Schema

##### <span id="patch-gatekeeper-config-default"></span> Default Response
errorResponse

###### <span id="patch-gatekeeper-config-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="patch-kubermatic-settings"></span> Patches the global settings. (*patchKubermaticSettings*)

```
PATCH /api/v1/admin/settings
```

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| Patch | `body` | [interface{}](#interface) | `interface{}` | |  | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#patch-kubermatic-settings-200) | OK | GlobalSettings |  | [schema](#patch-kubermatic-settings-200-schema) |
| [401](#patch-kubermatic-settings-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#patch-kubermatic-settings-401-schema) |
| [403](#patch-kubermatic-settings-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#patch-kubermatic-settings-403-schema) |
| [default](#patch-kubermatic-settings-default) | | errorResponse |  | [schema](#patch-kubermatic-settings-default-schema) |

#### Responses


##### <span id="patch-kubermatic-settings-200"></span> 200 - GlobalSettings
Status: OK

###### <span id="patch-kubermatic-settings-200-schema"></span> Schema
   
  

[GlobalSettings](#global-settings)

##### <span id="patch-kubermatic-settings-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="patch-kubermatic-settings-401-schema"></span> Schema

##### <span id="patch-kubermatic-settings-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="patch-kubermatic-settings-403-schema"></span> Schema

##### <span id="patch-kubermatic-settings-default"></span> Default Response
errorResponse

###### <span id="patch-kubermatic-settings-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="patch-machine-deployment"></span> patch machine deployment (*patchMachineDeployment*)

```
PATCH /api/v2/projects/{project_id}/clusters/{cluster_id}/machinedeployments/{machinedeployment_id}
```

Patches a machine deployment that is assigned to the given cluster. Please note that at the moment only
node deployment's spec can be updated by a patch, no other fields can be changed using this endpoint.

#### Consumes
  * application/json

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| machinedeployment_id | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |
| Patch | `body` | [interface{}](#interface) | `interface{}` | |  | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#patch-machine-deployment-200) | OK | NodeDeployment |  | [schema](#patch-machine-deployment-200-schema) |
| [401](#patch-machine-deployment-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#patch-machine-deployment-401-schema) |
| [403](#patch-machine-deployment-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#patch-machine-deployment-403-schema) |
| [default](#patch-machine-deployment-default) | | errorResponse |  | [schema](#patch-machine-deployment-default-schema) |

#### Responses


##### <span id="patch-machine-deployment-200"></span> 200 - NodeDeployment
Status: OK

###### <span id="patch-machine-deployment-200-schema"></span> Schema
   
  

[NodeDeployment](#node-deployment)

##### <span id="patch-machine-deployment-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="patch-machine-deployment-401-schema"></span> Schema

##### <span id="patch-machine-deployment-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="patch-machine-deployment-403-schema"></span> Schema

##### <span id="patch-machine-deployment-default"></span> Default Response
errorResponse

###### <span id="patch-machine-deployment-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="patch-node-deployment"></span> patch node deployment (*patchNodeDeployment*)

```
PATCH /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/nodedeployments/{nodedeployment_id}
```

Patches a node deployment that is assigned to the given cluster. Please note that at the moment only
node deployment's spec can be updated by a patch, no other fields can be changed using this endpoint.

#### Consumes
  * application/json

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| dc | `path` | string | `string` |  | ✓ |  |  |
| nodedeployment_id | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |
| Patch | `body` | [interface{}](#interface) | `interface{}` | |  | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#patch-node-deployment-200) | OK | NodeDeployment |  | [schema](#patch-node-deployment-200-schema) |
| [401](#patch-node-deployment-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#patch-node-deployment-401-schema) |
| [403](#patch-node-deployment-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#patch-node-deployment-403-schema) |
| [default](#patch-node-deployment-default) | | errorResponse |  | [schema](#patch-node-deployment-default-schema) |

#### Responses


##### <span id="patch-node-deployment-200"></span> 200 - NodeDeployment
Status: OK

###### <span id="patch-node-deployment-200-schema"></span> Schema
   
  

[NodeDeployment](#node-deployment)

##### <span id="patch-node-deployment-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="patch-node-deployment-401-schema"></span> Schema

##### <span id="patch-node-deployment-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="patch-node-deployment-403-schema"></span> Schema

##### <span id="patch-node-deployment-default"></span> Default Response
errorResponse

###### <span id="patch-node-deployment-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="patch-node-deployment-request"></span> Patches a NodeDeploymentRequest that is assigned to the given cluster. (*patchNodeDeploymentRequest*)

```
PATCH /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/ndrequests/{ndrequest_id}
```

#### Consumes
  * application/json

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| dc | `path` | string | `string` |  | ✓ |  |  |
| ndrequest_id | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |
| Patch | `body` | []uint8 (formatted integer) | `[]uint8` | |  | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#patch-node-deployment-request-200) | OK | NodeDeploymentRequest |  | [schema](#patch-node-deployment-request-200-schema) |
| [401](#patch-node-deployment-request-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#patch-node-deployment-request-401-schema) |
| [403](#patch-node-deployment-request-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#patch-node-deployment-request-403-schema) |
| [default](#patch-node-deployment-request-default) | | errorResponse |  | [schema](#patch-node-deployment-request-default-schema) |

#### Responses


##### <span id="patch-node-deployment-request-200"></span> 200 - NodeDeploymentRequest
Status: OK

###### <span id="patch-node-deployment-request-200-schema"></span> Schema
   
  

[NodeDeploymentRequest](#node-deployment-request)

##### <span id="patch-node-deployment-request-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="patch-node-deployment-request-401-schema"></span> Schema

##### <span id="patch-node-deployment-request-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="patch-node-deployment-request-403-schema"></span> Schema

##### <span id="patch-node-deployment-request-default"></span> Default Response
errorResponse

###### <span id="patch-node-deployment-request-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="patch-role"></span> patch role (*patchRole*)

```
PATCH /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/roles/{namespace}/{role_id}
```

Patch the role with the given name

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| dc | `path` | string | `string` |  | ✓ |  |  |
| namespace | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |
| role_id | `path` | string | `string` |  | ✓ |  |  |
| Patch | `body` | [interface{}](#interface) | `interface{}` | |  | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#patch-role-200) | OK | Role |  | [schema](#patch-role-200-schema) |
| [401](#patch-role-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#patch-role-401-schema) |
| [403](#patch-role-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#patch-role-403-schema) |
| [default](#patch-role-default) | | errorResponse |  | [schema](#patch-role-default-schema) |

#### Responses


##### <span id="patch-role-200"></span> 200 - Role
Status: OK

###### <span id="patch-role-200-schema"></span> Schema
   
  

[Role](#role)

##### <span id="patch-role-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="patch-role-401-schema"></span> Schema

##### <span id="patch-role-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="patch-role-403-schema"></span> Schema

##### <span id="patch-role-default"></span> Default Response
errorResponse

###### <span id="patch-role-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="patch-service-account-token"></span> patch service account token (*patchServiceAccountToken*)

```
PATCH /api/v1/projects/{project_id}/serviceaccounts/{serviceaccount_id}/tokens/{token_id}
```

Patches the token name

#### Consumes
  * application/json

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| project_id | `path` | string | `string` |  | ✓ |  |  |
| serviceaccount_id | `path` | string | `string` |  | ✓ |  |  |
| token_id | `path` | string | `string` |  | ✓ |  |  |
| Body | `body` | []uint8 (formatted integer) | `[]uint8` | |  | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#patch-service-account-token-200) | OK | PublicServiceAccountToken |  | [schema](#patch-service-account-token-200-schema) |
| [401](#patch-service-account-token-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#patch-service-account-token-401-schema) |
| [403](#patch-service-account-token-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#patch-service-account-token-403-schema) |
| [default](#patch-service-account-token-default) | | errorResponse |  | [schema](#patch-service-account-token-default-schema) |

#### Responses


##### <span id="patch-service-account-token-200"></span> 200 - PublicServiceAccountToken
Status: OK

###### <span id="patch-service-account-token-200-schema"></span> Schema
   
  

[PublicServiceAccountToken](#public-service-account-token)

##### <span id="patch-service-account-token-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="patch-service-account-token-401-schema"></span> Schema

##### <span id="patch-service-account-token-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="patch-service-account-token-403-schema"></span> Schema

##### <span id="patch-service-account-token-default"></span> Default Response
errorResponse

###### <span id="patch-service-account-token-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="reset-alertmanager"></span> Resets the alertmanager configuration to default for the specified cluster. (*resetAlertmanager*)

```
DELETE /api/v2/projects/{project_id}/clusters/{cluster_id}/alertmanager/config
```

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#reset-alertmanager-200) | OK | EmptyResponse is a empty response |  | [schema](#reset-alertmanager-200-schema) |
| [401](#reset-alertmanager-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#reset-alertmanager-401-schema) |
| [403](#reset-alertmanager-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#reset-alertmanager-403-schema) |
| [default](#reset-alertmanager-default) | | errorResponse |  | [schema](#reset-alertmanager-default-schema) |

#### Responses


##### <span id="reset-alertmanager-200"></span> 200 - EmptyResponse is a empty response
Status: OK

###### <span id="reset-alertmanager-200-schema"></span> Schema

##### <span id="reset-alertmanager-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="reset-alertmanager-401-schema"></span> Schema

##### <span id="reset-alertmanager-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="reset-alertmanager-403-schema"></span> Schema

##### <span id="reset-alertmanager-default"></span> Default Response
errorResponse

###### <span id="reset-alertmanager-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="restart-machine-deployment"></span> Schedules rolling restart of a machine deployment that is assigned to the given cluster. (*restartMachineDeployment*)

```
POST /api/v2/projects/{project_id}/clusters/{cluster_id}/machinedeployments/{machinedeployment_id}
```

#### Consumes
  * application/json

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| machinedeployment_id | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#restart-machine-deployment-200) | OK | NodeDeployment |  | [schema](#restart-machine-deployment-200-schema) |
| [401](#restart-machine-deployment-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#restart-machine-deployment-401-schema) |
| [403](#restart-machine-deployment-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#restart-machine-deployment-403-schema) |
| [default](#restart-machine-deployment-default) | | errorResponse |  | [schema](#restart-machine-deployment-default-schema) |

#### Responses


##### <span id="restart-machine-deployment-200"></span> 200 - NodeDeployment
Status: OK

###### <span id="restart-machine-deployment-200-schema"></span> Schema
   
  

[NodeDeployment](#node-deployment)

##### <span id="restart-machine-deployment-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="restart-machine-deployment-401-schema"></span> Schema

##### <span id="restart-machine-deployment-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="restart-machine-deployment-403-schema"></span> Schema

##### <span id="restart-machine-deployment-default"></span> Default Response
errorResponse

###### <span id="restart-machine-deployment-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="revoke-cluster-admin-token"></span> revoke cluster admin token (*revokeClusterAdminToken*)

```
PUT /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/token
```

Revokes the current admin token

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| dc | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#revoke-cluster-admin-token-200) | OK | EmptyResponse is a empty response |  | [schema](#revoke-cluster-admin-token-200-schema) |
| [401](#revoke-cluster-admin-token-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#revoke-cluster-admin-token-401-schema) |
| [403](#revoke-cluster-admin-token-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#revoke-cluster-admin-token-403-schema) |
| [default](#revoke-cluster-admin-token-default) | | errorResponse |  | [schema](#revoke-cluster-admin-token-default-schema) |

#### Responses


##### <span id="revoke-cluster-admin-token-200"></span> 200 - EmptyResponse is a empty response
Status: OK

###### <span id="revoke-cluster-admin-token-200-schema"></span> Schema

##### <span id="revoke-cluster-admin-token-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="revoke-cluster-admin-token-401-schema"></span> Schema

##### <span id="revoke-cluster-admin-token-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="revoke-cluster-admin-token-403-schema"></span> Schema

##### <span id="revoke-cluster-admin-token-default"></span> Default Response
errorResponse

###### <span id="revoke-cluster-admin-token-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="revoke-cluster-admin-token-v2"></span> revoke cluster admin token v2 (*revokeClusterAdminTokenV2*)

```
PUT /api/v2/projects/{project_id}/clusters/{cluster_id}/token
```

Revokes the current admin token

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#revoke-cluster-admin-token-v2-200) | OK | EmptyResponse is a empty response |  | [schema](#revoke-cluster-admin-token-v2-200-schema) |
| [401](#revoke-cluster-admin-token-v2-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#revoke-cluster-admin-token-v2-401-schema) |
| [403](#revoke-cluster-admin-token-v2-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#revoke-cluster-admin-token-v2-403-schema) |
| [default](#revoke-cluster-admin-token-v2-default) | | errorResponse |  | [schema](#revoke-cluster-admin-token-v2-default-schema) |

#### Responses


##### <span id="revoke-cluster-admin-token-v2-200"></span> 200 - EmptyResponse is a empty response
Status: OK

###### <span id="revoke-cluster-admin-token-v2-200-schema"></span> Schema

##### <span id="revoke-cluster-admin-token-v2-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="revoke-cluster-admin-token-v2-401-schema"></span> Schema

##### <span id="revoke-cluster-admin-token-v2-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="revoke-cluster-admin-token-v2-403-schema"></span> Schema

##### <span id="revoke-cluster-admin-token-v2-default"></span> Default Response
errorResponse

###### <span id="revoke-cluster-admin-token-v2-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="revoke-cluster-viewer-token"></span> revoke cluster viewer token (*revokeClusterViewerToken*)

```
PUT /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/viewertoken
```

Revokes the current viewer token

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| dc | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#revoke-cluster-viewer-token-200) | OK | EmptyResponse is a empty response |  | [schema](#revoke-cluster-viewer-token-200-schema) |
| [401](#revoke-cluster-viewer-token-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#revoke-cluster-viewer-token-401-schema) |
| [403](#revoke-cluster-viewer-token-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#revoke-cluster-viewer-token-403-schema) |
| [default](#revoke-cluster-viewer-token-default) | | errorResponse |  | [schema](#revoke-cluster-viewer-token-default-schema) |

#### Responses


##### <span id="revoke-cluster-viewer-token-200"></span> 200 - EmptyResponse is a empty response
Status: OK

###### <span id="revoke-cluster-viewer-token-200-schema"></span> Schema

##### <span id="revoke-cluster-viewer-token-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="revoke-cluster-viewer-token-401-schema"></span> Schema

##### <span id="revoke-cluster-viewer-token-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="revoke-cluster-viewer-token-403-schema"></span> Schema

##### <span id="revoke-cluster-viewer-token-default"></span> Default Response
errorResponse

###### <span id="revoke-cluster-viewer-token-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="revoke-cluster-viewer-token-v2"></span> revoke cluster viewer token v2 (*revokeClusterViewerTokenV2*)

```
PUT /api/v2/projects/{project_id}/clusters/{cluster_id}/viewertoken
```

Revokes the current viewer token

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#revoke-cluster-viewer-token-v2-200) | OK | EmptyResponse is a empty response |  | [schema](#revoke-cluster-viewer-token-v2-200-schema) |
| [401](#revoke-cluster-viewer-token-v2-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#revoke-cluster-viewer-token-v2-401-schema) |
| [403](#revoke-cluster-viewer-token-v2-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#revoke-cluster-viewer-token-v2-403-schema) |
| [default](#revoke-cluster-viewer-token-v2-default) | | errorResponse |  | [schema](#revoke-cluster-viewer-token-v2-default-schema) |

#### Responses


##### <span id="revoke-cluster-viewer-token-v2-200"></span> 200 - EmptyResponse is a empty response
Status: OK

###### <span id="revoke-cluster-viewer-token-v2-200-schema"></span> Schema

##### <span id="revoke-cluster-viewer-token-v2-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="revoke-cluster-viewer-token-v2-401-schema"></span> Schema

##### <span id="revoke-cluster-viewer-token-v2-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="revoke-cluster-viewer-token-v2-403-schema"></span> Schema

##### <span id="revoke-cluster-viewer-token-v2-default"></span> Default Response
errorResponse

###### <span id="revoke-cluster-viewer-token-v2-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="set-admin"></span> Allows setting and clearing admin role for users. (*setAdmin*)

```
PUT /api/v1/admin
```

#### Consumes
  * application/json

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| Body | `body` | [Admin](#admin) | `models.Admin` | |  | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#set-admin-200) | OK | Admin |  | [schema](#set-admin-200-schema) |
| [401](#set-admin-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#set-admin-401-schema) |
| [403](#set-admin-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#set-admin-403-schema) |
| [default](#set-admin-default) | | errorResponse |  | [schema](#set-admin-default-schema) |

#### Responses


##### <span id="set-admin-200"></span> 200 - Admin
Status: OK

###### <span id="set-admin-200-schema"></span> Schema
   
  

[Admin](#admin)

##### <span id="set-admin-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="set-admin-401-schema"></span> Schema

##### <span id="set-admin-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="set-admin-403-schema"></span> Schema

##### <span id="set-admin-default"></span> Default Response
errorResponse

###### <span id="set-admin-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="unbind-user-from-cluster-role-binding"></span> unbind user from cluster role binding (*unbindUserFromClusterRoleBinding*)

```
DELETE /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/clusterroles/{role_id}/clusterbindings
```

Unbinds user from cluster role binding

#### Consumes
  * application/json

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| dc | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |
| role_id | `path` | string | `string` |  | ✓ |  |  |
| Body | `body` | [ClusterRoleUser](#cluster-role-user) | `models.ClusterRoleUser` | |  | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#unbind-user-from-cluster-role-binding-200) | OK | ClusterRoleBinding |  | [schema](#unbind-user-from-cluster-role-binding-200-schema) |
| [401](#unbind-user-from-cluster-role-binding-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#unbind-user-from-cluster-role-binding-401-schema) |
| [403](#unbind-user-from-cluster-role-binding-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#unbind-user-from-cluster-role-binding-403-schema) |
| [default](#unbind-user-from-cluster-role-binding-default) | | errorResponse |  | [schema](#unbind-user-from-cluster-role-binding-default-schema) |

#### Responses


##### <span id="unbind-user-from-cluster-role-binding-200"></span> 200 - ClusterRoleBinding
Status: OK

###### <span id="unbind-user-from-cluster-role-binding-200-schema"></span> Schema
   
  

[ClusterRoleBinding](#cluster-role-binding)

##### <span id="unbind-user-from-cluster-role-binding-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="unbind-user-from-cluster-role-binding-401-schema"></span> Schema

##### <span id="unbind-user-from-cluster-role-binding-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="unbind-user-from-cluster-role-binding-403-schema"></span> Schema

##### <span id="unbind-user-from-cluster-role-binding-default"></span> Default Response
errorResponse

###### <span id="unbind-user-from-cluster-role-binding-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="unbind-user-from-cluster-role-binding-v2"></span> unbind user from cluster role binding v2 (*unbindUserFromClusterRoleBindingV2*)

```
DELETE /api/v2/projects/{project_id}/clusters/{cluster_id}/clusterroles/{role_id}/clusterbindings
```

Unbinds user from cluster role binding

#### Consumes
  * application/json

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |
| role_id | `path` | string | `string` |  | ✓ |  |  |
| Body | `body` | [ClusterRoleUser](#cluster-role-user) | `models.ClusterRoleUser` | |  | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#unbind-user-from-cluster-role-binding-v2-200) | OK | ClusterRoleBinding |  | [schema](#unbind-user-from-cluster-role-binding-v2-200-schema) |
| [401](#unbind-user-from-cluster-role-binding-v2-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#unbind-user-from-cluster-role-binding-v2-401-schema) |
| [403](#unbind-user-from-cluster-role-binding-v2-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#unbind-user-from-cluster-role-binding-v2-403-schema) |
| [default](#unbind-user-from-cluster-role-binding-v2-default) | | errorResponse |  | [schema](#unbind-user-from-cluster-role-binding-v2-default-schema) |

#### Responses


##### <span id="unbind-user-from-cluster-role-binding-v2-200"></span> 200 - ClusterRoleBinding
Status: OK

###### <span id="unbind-user-from-cluster-role-binding-v2-200-schema"></span> Schema
   
  

[ClusterRoleBinding](#cluster-role-binding)

##### <span id="unbind-user-from-cluster-role-binding-v2-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="unbind-user-from-cluster-role-binding-v2-401-schema"></span> Schema

##### <span id="unbind-user-from-cluster-role-binding-v2-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="unbind-user-from-cluster-role-binding-v2-403-schema"></span> Schema

##### <span id="unbind-user-from-cluster-role-binding-v2-default"></span> Default Response
errorResponse

###### <span id="unbind-user-from-cluster-role-binding-v2-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="unbind-user-from-role-binding"></span> unbind user from role binding (*unbindUserFromRoleBinding*)

```
DELETE /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/roles/{namespace}/{role_id}/bindings
```

Unbinds user from the role binding

#### Consumes
  * application/json

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| dc | `path` | string | `string` |  | ✓ |  |  |
| namespace | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |
| role_id | `path` | string | `string` |  | ✓ |  |  |
| Body | `body` | [RoleUser](#role-user) | `models.RoleUser` | |  | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#unbind-user-from-role-binding-200) | OK | RoleBinding |  | [schema](#unbind-user-from-role-binding-200-schema) |
| [401](#unbind-user-from-role-binding-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#unbind-user-from-role-binding-401-schema) |
| [403](#unbind-user-from-role-binding-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#unbind-user-from-role-binding-403-schema) |
| [default](#unbind-user-from-role-binding-default) | | errorResponse |  | [schema](#unbind-user-from-role-binding-default-schema) |

#### Responses


##### <span id="unbind-user-from-role-binding-200"></span> 200 - RoleBinding
Status: OK

###### <span id="unbind-user-from-role-binding-200-schema"></span> Schema
   
  

[RoleBinding](#role-binding)

##### <span id="unbind-user-from-role-binding-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="unbind-user-from-role-binding-401-schema"></span> Schema

##### <span id="unbind-user-from-role-binding-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="unbind-user-from-role-binding-403-schema"></span> Schema

##### <span id="unbind-user-from-role-binding-default"></span> Default Response
errorResponse

###### <span id="unbind-user-from-role-binding-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="unbind-user-from-role-binding-v2"></span> unbind user from role binding v2 (*unbindUserFromRoleBindingV2*)

```
DELETE /api/v2/projects/{project_id}/clusters/{cluster_id}/roles/{namespace}/{role_id}/bindings
```

Unbinds user from the role binding

#### Consumes
  * application/json

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| namespace | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |
| role_id | `path` | string | `string` |  | ✓ |  |  |
| Body | `body` | [RoleUser](#role-user) | `models.RoleUser` | |  | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#unbind-user-from-role-binding-v2-200) | OK | RoleBinding |  | [schema](#unbind-user-from-role-binding-v2-200-schema) |
| [401](#unbind-user-from-role-binding-v2-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#unbind-user-from-role-binding-v2-401-schema) |
| [403](#unbind-user-from-role-binding-v2-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#unbind-user-from-role-binding-v2-403-schema) |
| [default](#unbind-user-from-role-binding-v2-default) | | errorResponse |  | [schema](#unbind-user-from-role-binding-v2-default-schema) |

#### Responses


##### <span id="unbind-user-from-role-binding-v2-200"></span> 200 - RoleBinding
Status: OK

###### <span id="unbind-user-from-role-binding-v2-200-schema"></span> Schema
   
  

[RoleBinding](#role-binding)

##### <span id="unbind-user-from-role-binding-v2-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="unbind-user-from-role-binding-v2-401-schema"></span> Schema

##### <span id="unbind-user-from-role-binding-v2-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="unbind-user-from-role-binding-v2-403-schema"></span> Schema

##### <span id="unbind-user-from-role-binding-v2-default"></span> Default Response
errorResponse

###### <span id="unbind-user-from-role-binding-v2-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="update-admission-plugin"></span> Updates the admission plugin. (*updateAdmissionPlugin*)

```
PATCH /api/v1/admin/admission/plugins/{name}
```

#### Consumes
  * application/json

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| name | `path` | string | `string` |  | ✓ |  |  |
| Body | `body` | [AdmissionPlugin](#admission-plugin) | `models.AdmissionPlugin` | |  | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#update-admission-plugin-200) | OK | AdmissionPlugin |  | [schema](#update-admission-plugin-200-schema) |
| [401](#update-admission-plugin-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#update-admission-plugin-401-schema) |
| [403](#update-admission-plugin-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#update-admission-plugin-403-schema) |
| [default](#update-admission-plugin-default) | | errorResponse |  | [schema](#update-admission-plugin-default-schema) |

#### Responses


##### <span id="update-admission-plugin-200"></span> 200 - AdmissionPlugin
Status: OK

###### <span id="update-admission-plugin-200-schema"></span> Schema
   
  

[AdmissionPlugin](#admission-plugin)

##### <span id="update-admission-plugin-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="update-admission-plugin-401-schema"></span> Schema

##### <span id="update-admission-plugin-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="update-admission-plugin-403-schema"></span> Schema

##### <span id="update-admission-plugin-default"></span> Default Response
errorResponse

###### <span id="update-admission-plugin-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="update-alertmanager"></span> update alertmanager (*updateAlertmanager*)

```
PUT /api/v2/projects/{project_id}/clusters/{cluster_id}/alertmanager/config
```

Updates an alertmanager configuration for the given cluster

#### Consumes
  * application/json

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |
| Body | `body` | [Alertmanager](#alertmanager) | `models.Alertmanager` | | ✓ | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#update-alertmanager-200) | OK | Alertmanager |  | [schema](#update-alertmanager-200-schema) |
| [401](#update-alertmanager-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#update-alertmanager-401-schema) |
| [403](#update-alertmanager-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#update-alertmanager-403-schema) |
| [default](#update-alertmanager-default) | | errorResponse |  | [schema](#update-alertmanager-default-schema) |

#### Responses


##### <span id="update-alertmanager-200"></span> 200 - Alertmanager
Status: OK

###### <span id="update-alertmanager-200-schema"></span> Schema
   
  

[Alertmanager](#alertmanager)

##### <span id="update-alertmanager-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="update-alertmanager-401-schema"></span> Schema

##### <span id="update-alertmanager-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="update-alertmanager-403-schema"></span> Schema

##### <span id="update-alertmanager-default"></span> Default Response
errorResponse

###### <span id="update-alertmanager-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="update-d-c"></span> Update the datacenter. The datacenter spec will be overwritten with the one provided in the request. (*updateDC*)

```
PUT /api/v1/seed/{seed_name}/dc/{dc}
```

#### Consumes
  * application/json

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| dc | `path` | string | `string` |  | ✓ |  |  |
| seed_name | `path` | string | `string` |  | ✓ |  |  |
| Body | `body` | [UpdateDCBody](#update-d-c-body) | `UpdateDCBody` | |  | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#update-d-c-200) | OK | Datacenter |  | [schema](#update-d-c-200-schema) |
| [401](#update-d-c-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#update-d-c-401-schema) |
| [403](#update-d-c-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#update-d-c-403-schema) |
| [default](#update-d-c-default) | | errorResponse |  | [schema](#update-d-c-default-schema) |

#### Responses


##### <span id="update-d-c-200"></span> 200 - Datacenter
Status: OK

###### <span id="update-d-c-200-schema"></span> Schema
   
  

[Datacenter](#datacenter)

##### <span id="update-d-c-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="update-d-c-401-schema"></span> Schema

##### <span id="update-d-c-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="update-d-c-403-schema"></span> Schema

##### <span id="update-d-c-default"></span> Default Response
errorResponse

###### <span id="update-d-c-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

###### Inlined models

**<span id="update-d-c-body"></span> UpdateDCBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Name | string| `string` |  | |  |  |
| spec | [DatacenterSpec](#datacenter-spec)| `models.DatacenterSpec` |  | |  |  |



### <span id="update-external-cluster"></span> Updates an external cluster for the given project. (*updateExternalCluster*)

```
PUT /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}
```

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |
| Body | `body` | [Body](#body) | `models.Body` | |  | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#update-external-cluster-200) | OK | Cluster |  | [schema](#update-external-cluster-200-schema) |
| [401](#update-external-cluster-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#update-external-cluster-401-schema) |
| [403](#update-external-cluster-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#update-external-cluster-403-schema) |
| [default](#update-external-cluster-default) | | errorResponse |  | [schema](#update-external-cluster-default-schema) |

#### Responses


##### <span id="update-external-cluster-200"></span> 200 - Cluster
Status: OK

###### <span id="update-external-cluster-200-schema"></span> Schema
   
  

[Cluster](#cluster)

##### <span id="update-external-cluster-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="update-external-cluster-401-schema"></span> Schema

##### <span id="update-external-cluster-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="update-external-cluster-403-schema"></span> Schema

##### <span id="update-external-cluster-default"></span> Default Response
errorResponse

###### <span id="update-external-cluster-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="update-m-l-a-admin-setting"></span> Updates the MLA admin setting for the given cluster. (*updateMLAAdminSetting*)

```
PUT /api/v2/projects/{project_id}/clusters/{cluster_id}/mlaadminsetting
```

#### Consumes
  * application/json

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |
| Body | `body` | [MLAAdminSetting](#m-l-a-admin-setting) | `models.MLAAdminSetting` | | ✓ | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#update-m-l-a-admin-setting-200) | OK | MLAAdminSetting |  | [schema](#update-m-l-a-admin-setting-200-schema) |
| [401](#update-m-l-a-admin-setting-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#update-m-l-a-admin-setting-401-schema) |
| [403](#update-m-l-a-admin-setting-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#update-m-l-a-admin-setting-403-schema) |
| [default](#update-m-l-a-admin-setting-default) | | errorResponse |  | [schema](#update-m-l-a-admin-setting-default-schema) |

#### Responses


##### <span id="update-m-l-a-admin-setting-200"></span> 200 - MLAAdminSetting
Status: OK

###### <span id="update-m-l-a-admin-setting-200-schema"></span> Schema
   
  

[MLAAdminSetting](#m-l-a-admin-setting)

##### <span id="update-m-l-a-admin-setting-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="update-m-l-a-admin-setting-401-schema"></span> Schema

##### <span id="update-m-l-a-admin-setting-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="update-m-l-a-admin-setting-403-schema"></span> Schema

##### <span id="update-m-l-a-admin-setting-default"></span> Default Response
errorResponse

###### <span id="update-m-l-a-admin-setting-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="update-preset"></span> update preset (*updatePreset*)

```
PUT /api/v2/providers/{provider_name}/presets
```

Updates provider preset

#### Consumes
  * application/json

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| provider_name | `path` | string | `string` |  | ✓ |  |  |
| Body | `body` | [Preset](#preset) | `models.Preset` | | ✓ | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#update-preset-200) | OK | Preset |  | [schema](#update-preset-200-schema) |
| [401](#update-preset-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#update-preset-401-schema) |
| [403](#update-preset-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#update-preset-403-schema) |
| [default](#update-preset-default) | | errorResponse |  | [schema](#update-preset-default-schema) |

#### Responses


##### <span id="update-preset-200"></span> 200 - Preset
Status: OK

###### <span id="update-preset-200-schema"></span> Schema
   
  

[Preset](#preset)

##### <span id="update-preset-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="update-preset-401-schema"></span> Schema

##### <span id="update-preset-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="update-preset-403-schema"></span> Schema

##### <span id="update-preset-default"></span> Default Response
errorResponse

###### <span id="update-preset-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="update-preset-status"></span> Updates the status of a preset. It can enable or disable it, so that it won't be listed by the list endpoints. (*updatePresetStatus*)

```
PUT /api/v2/presets/{preset_name}/status
```

#### Consumes
  * application/json

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| preset_name | `path` | string | `string` |  | ✓ |  |  |
| provider | `query` | string | `string` |  |  |  |  |
| Body | `body` | [UpdatePresetStatusBody](#update-preset-status-body) | `UpdatePresetStatusBody` | | ✓ | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#update-preset-status-200) | OK | EmptyResponse is a empty response |  | [schema](#update-preset-status-200-schema) |
| [401](#update-preset-status-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#update-preset-status-401-schema) |
| [403](#update-preset-status-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#update-preset-status-403-schema) |
| [default](#update-preset-status-default) | | errorResponse |  | [schema](#update-preset-status-default-schema) |

#### Responses


##### <span id="update-preset-status-200"></span> 200 - EmptyResponse is a empty response
Status: OK

###### <span id="update-preset-status-200-schema"></span> Schema

##### <span id="update-preset-status-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="update-preset-status-401-schema"></span> Schema

##### <span id="update-preset-status-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="update-preset-status-403-schema"></span> Schema

##### <span id="update-preset-status-default"></span> Default Response
errorResponse

###### <span id="update-preset-status-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

###### Inlined models

**<span id="update-preset-status-body"></span> UpdatePresetStatusBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Enabled | boolean| `bool` |  | |  |  |



### <span id="update-project"></span> update project (*updateProject*)

```
PUT /api/v1/projects/{project_id}
```

Updates the given project

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| project_id | `path` | string | `string` |  | ✓ |  |  |
| Body | `body` | [Project](#project) | `models.Project` | |  | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#update-project-200) | OK | Project |  | [schema](#update-project-200-schema) |
| [400](#update-project-400) | Bad Request | EmptyResponse is a empty response |  | [schema](#update-project-400-schema) |
| [404](#update-project-404) | Not Found | EmptyResponse is a empty response |  | [schema](#update-project-404-schema) |
| [500](#update-project-500) | Internal Server Error | EmptyResponse is a empty response |  | [schema](#update-project-500-schema) |
| [501](#update-project-501) | Not Implemented | EmptyResponse is a empty response |  | [schema](#update-project-501-schema) |
| [default](#update-project-default) | | errorResponse |  | [schema](#update-project-default-schema) |

#### Responses


##### <span id="update-project-200"></span> 200 - Project
Status: OK

###### <span id="update-project-200-schema"></span> Schema
   
  

[Project](#project)

##### <span id="update-project-400"></span> 400 - EmptyResponse is a empty response
Status: Bad Request

###### <span id="update-project-400-schema"></span> Schema

##### <span id="update-project-404"></span> 404 - EmptyResponse is a empty response
Status: Not Found

###### <span id="update-project-404-schema"></span> Schema

##### <span id="update-project-500"></span> 500 - EmptyResponse is a empty response
Status: Internal Server Error

###### <span id="update-project-500-schema"></span> Schema

##### <span id="update-project-501"></span> 501 - EmptyResponse is a empty response
Status: Not Implemented

###### <span id="update-project-501-schema"></span> Schema

##### <span id="update-project-default"></span> Default Response
errorResponse

###### <span id="update-project-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="update-rule-group"></span> Updates the specified rule group for the given cluster. (*updateRuleGroup*)

```
PUT /api/v2/projects/{project_id}/clusters/{cluster_id}/rulegroups/{rulegroup_id}
```

#### Consumes
  * application/json

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |
| rulegroup_id | `path` | string | `string` |  | ✓ |  |  |
| Body | `body` | [RuleGroup](#rule-group) | `models.RuleGroup` | | ✓ | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#update-rule-group-200) | OK | RuleGroup |  | [schema](#update-rule-group-200-schema) |
| [401](#update-rule-group-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#update-rule-group-401-schema) |
| [403](#update-rule-group-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#update-rule-group-403-schema) |
| [default](#update-rule-group-default) | | errorResponse |  | [schema](#update-rule-group-default-schema) |

#### Responses


##### <span id="update-rule-group-200"></span> 200 - RuleGroup
Status: OK

###### <span id="update-rule-group-200-schema"></span> Schema
   
  

[RuleGroup](#rule-group)

##### <span id="update-rule-group-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="update-rule-group-401-schema"></span> Schema

##### <span id="update-rule-group-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="update-rule-group-403-schema"></span> Schema

##### <span id="update-rule-group-default"></span> Default Response
errorResponse

###### <span id="update-rule-group-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="update-seed"></span> Updates the seed. (*updateSeed*)

```
PATCH /api/v1/admin/seeds/{seed_name}
```

#### Consumes
  * application/json

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| seed_name | `path` | string | `string` |  | ✓ |  |  |
| Body | `body` | [UpdateSeedBody](#update-seed-body) | `UpdateSeedBody` | |  | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#update-seed-200) | OK | Seed |  | [schema](#update-seed-200-schema) |
| [401](#update-seed-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#update-seed-401-schema) |
| [403](#update-seed-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#update-seed-403-schema) |
| [default](#update-seed-default) | | errorResponse |  | [schema](#update-seed-default-schema) |

#### Responses


##### <span id="update-seed-200"></span> 200 - Seed
Status: OK

###### <span id="update-seed-200-schema"></span> Schema
   
  

[Seed](#seed)

##### <span id="update-seed-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="update-seed-401-schema"></span> Schema

##### <span id="update-seed-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="update-seed-403-schema"></span> Schema

##### <span id="update-seed-default"></span> Default Response
errorResponse

###### <span id="update-seed-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

###### Inlined models

**<span id="update-seed-body"></span> UpdateSeedBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Name | string| `string` |  | |  |  |
| spec | [SeedSpec](#seed-spec)| `models.SeedSpec` |  | |  |  |



### <span id="update-service-account"></span> update service account (*updateServiceAccount*)

```
PUT /api/v1/projects/{project_id}/serviceaccounts/{serviceaccount_id}
```

Updates service account for the given project

#### Consumes
  * application/json

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| project_id | `path` | string | `string` |  | ✓ |  |  |
| serviceaccount_id | `path` | string | `string` |  | ✓ |  |  |
| Body | `body` | [ServiceAccount](#service-account) | `models.ServiceAccount` | |  | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#update-service-account-200) | OK | ServiceAccount |  | [schema](#update-service-account-200-schema) |
| [401](#update-service-account-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#update-service-account-401-schema) |
| [403](#update-service-account-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#update-service-account-403-schema) |
| [default](#update-service-account-default) | | errorResponse |  | [schema](#update-service-account-default-schema) |

#### Responses


##### <span id="update-service-account-200"></span> 200 - ServiceAccount
Status: OK

###### <span id="update-service-account-200-schema"></span> Schema
   
  

[ServiceAccount](#service-account)

##### <span id="update-service-account-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="update-service-account-401-schema"></span> Schema

##### <span id="update-service-account-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="update-service-account-403-schema"></span> Schema

##### <span id="update-service-account-default"></span> Default Response
errorResponse

###### <span id="update-service-account-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="update-service-account-token"></span> update service account token (*updateServiceAccountToken*)

```
PUT /api/v1/projects/{project_id}/serviceaccounts/{serviceaccount_id}/tokens/{token_id}
```

Updates and regenerates the token

#### Consumes
  * application/json

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| project_id | `path` | string | `string` |  | ✓ |  |  |
| serviceaccount_id | `path` | string | `string` |  | ✓ |  |  |
| token_id | `path` | string | `string` |  | ✓ |  |  |
| Body | `body` | [PublicServiceAccountToken](#public-service-account-token) | `models.PublicServiceAccountToken` | |  | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#update-service-account-token-200) | OK | ServiceAccountToken |  | [schema](#update-service-account-token-200-schema) |
| [401](#update-service-account-token-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#update-service-account-token-401-schema) |
| [403](#update-service-account-token-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#update-service-account-token-403-schema) |
| [default](#update-service-account-token-default) | | errorResponse |  | [schema](#update-service-account-token-default-schema) |

#### Responses


##### <span id="update-service-account-token-200"></span> 200 - ServiceAccountToken
Status: OK

###### <span id="update-service-account-token-200-schema"></span> Schema
   
  

[ServiceAccountToken](#service-account-token)

##### <span id="update-service-account-token-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="update-service-account-token-401-schema"></span> Schema

##### <span id="update-service-account-token-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="update-service-account-token-403-schema"></span> Schema

##### <span id="update-service-account-token-default"></span> Default Response
errorResponse

###### <span id="update-service-account-token-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="upgrade-cluster-node-deployments"></span> upgrade cluster node deployments (*upgradeClusterNodeDeployments*)

```
PUT /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/nodes/upgrades
```

Upgrades node deployments in a cluster

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| dc | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |
| Body | `body` | [MasterVersion](#master-version) | `models.MasterVersion` | |  | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#upgrade-cluster-node-deployments-200) | OK | EmptyResponse is a empty response |  | [schema](#upgrade-cluster-node-deployments-200-schema) |
| [401](#upgrade-cluster-node-deployments-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#upgrade-cluster-node-deployments-401-schema) |
| [403](#upgrade-cluster-node-deployments-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#upgrade-cluster-node-deployments-403-schema) |
| [default](#upgrade-cluster-node-deployments-default) | | errorResponse |  | [schema](#upgrade-cluster-node-deployments-default-schema) |

#### Responses


##### <span id="upgrade-cluster-node-deployments-200"></span> 200 - EmptyResponse is a empty response
Status: OK

###### <span id="upgrade-cluster-node-deployments-200-schema"></span> Schema

##### <span id="upgrade-cluster-node-deployments-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="upgrade-cluster-node-deployments-401-schema"></span> Schema

##### <span id="upgrade-cluster-node-deployments-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="upgrade-cluster-node-deployments-403-schema"></span> Schema

##### <span id="upgrade-cluster-node-deployments-default"></span> Default Response
errorResponse

###### <span id="upgrade-cluster-node-deployments-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

### <span id="upgrade-cluster-node-deployments-v2"></span> upgrade cluster node deployments v2 (*upgradeClusterNodeDeploymentsV2*)

```
PUT /api/v2/projects/{project_id}/clusters/{cluster_id}/nodes/upgrades
```

Upgrades node deployments in a cluster

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| cluster_id | `path` | string | `string` |  | ✓ |  |  |
| project_id | `path` | string | `string` |  | ✓ |  |  |
| Body | `body` | [MasterVersion](#master-version) | `models.MasterVersion` | |  | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#upgrade-cluster-node-deployments-v2-200) | OK | EmptyResponse is a empty response |  | [schema](#upgrade-cluster-node-deployments-v2-200-schema) |
| [401](#upgrade-cluster-node-deployments-v2-401) | Unauthorized | EmptyResponse is a empty response |  | [schema](#upgrade-cluster-node-deployments-v2-401-schema) |
| [403](#upgrade-cluster-node-deployments-v2-403) | Forbidden | EmptyResponse is a empty response |  | [schema](#upgrade-cluster-node-deployments-v2-403-schema) |
| [default](#upgrade-cluster-node-deployments-v2-default) | | errorResponse |  | [schema](#upgrade-cluster-node-deployments-v2-default-schema) |

#### Responses


##### <span id="upgrade-cluster-node-deployments-v2-200"></span> 200 - EmptyResponse is a empty response
Status: OK

###### <span id="upgrade-cluster-node-deployments-v2-200-schema"></span> Schema

##### <span id="upgrade-cluster-node-deployments-v2-401"></span> 401 - EmptyResponse is a empty response
Status: Unauthorized

###### <span id="upgrade-cluster-node-deployments-v2-401-schema"></span> Schema

##### <span id="upgrade-cluster-node-deployments-v2-403"></span> 403 - EmptyResponse is a empty response
Status: Forbidden

###### <span id="upgrade-cluster-node-deployments-v2-403-schema"></span> Schema

##### <span id="upgrade-cluster-node-deployments-v2-default"></span> Default Response
errorResponse

###### <span id="upgrade-cluster-node-deployments-v2-default-schema"></span> Schema

  

[ErrorResponse](#error-response)

## Models

### <span id="a-w-s-cloud-spec"></span> AWSCloudSpec


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| AccessKeyID | string| `string` |  | |  |  |
| ControlPlaneRoleARN | string| `string` |  | | The IAM role, the control plane will use. The control plane will perform an assume-role |  |
| InstanceProfileName | string| `string` |  | |  |  |
| OpenstackBillingTenant | string| `string` |  | |  |  |
| RoleName | string| `string` |  | | DEPRECATED. Don't care for the role name. We only require the ControlPlaneRoleARN to be set so the control plane
can perform the assume-role.
We keep it for backwards compatibility (We use this name for cleanup purpose). |  |
| RouteTableID | string| `string` |  | |  |  |
| SecretAccessKey | string| `string` |  | |  |  |
| SecurityGroupID | string| `string` |  | |  |  |
| VPCID | string| `string` |  | |  |  |
| credentialsReference | [GlobalSecretKeySelector](#global-secret-key-selector)| `GlobalSecretKeySelector` |  | |  |  |



### <span id="a-w-s-node-spec"></span> AWSNodeSpec


> AWSNodeSpec aws specific node settings
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| AMI | string| `string` |  | | ami to use. Will be defaulted to a ami for your selected operating system and region. Only set this when you know what you do. |  |
| AssignPublicIP | boolean| `bool` |  | | This flag controls a property of the AWS instance. When set the AWS instance will get a public IP address
assigned during launch overriding a possible setting in the used AWS subnet. |  |
| AvailabilityZone | string| `string` |  | | Availability zone in which to place the node. It is coupled with the subnet to which the node will belong. |  |
| InstanceType | string| `string` | ✓ | |  | `t2.micro` |
| IsSpotInstance | boolean| `bool` |  | | IsSpotInstance indicates whether the created machine is an aws ec2 spot instance or on-demand ec2 instance. |  |
| SpotInstanceInterruptionBehavior | string| `string` |  | | SpotInstanceInterruptionBehavior sets the interruption behavior for the spot instance when capacity is no longer
available at the price you specified, if there is no capacity, or if a constraint cannot be met. Charges for EBS
volume storage apply when an instance is stopped. |  |
| SpotInstanceMaxPrice | string| `string` |  | | SpotInstanceMaxPrice is the maximum price you are willing to pay per instance hour. Your instance runs when
your maximum price is greater than the Spot Price. |  |
| SpotInstancePersistentRequest | boolean| `bool` |  | | SpotInstancePersistentRequest ensures that your request will be submitted every time your Spot Instance is terminated. |  |
| SubnetID | string| `string` |  | | The VPC subnet to which the node shall be connected. |  |
| Tags | map of string| `map[string]string` |  | | additional instance tags |  |
| VolumeSize | int64 (formatted integer)| `int64` | ✓ | | size of the volume in gb. Only one volume will be created |  |
| VolumeType | string| `string` | ✓ | |  | `gp2, io1, st1, sc1, standard` |



### <span id="a-w-s-security-group-list"></span> AWSSecurityGroupList


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| IDs | []string| `[]string` |  | |  |  |



### <span id="a-w-s-size"></span> AWSSize


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Architecture | string| `string` |  | |  |  |
| GPUs | int64 (formatted integer)| `int64` |  | |  |  |
| Memory | float (formatted number)| `float32` |  | |  |  |
| Name | string| `string` |  | |  |  |
| PrettyName | string| `string` |  | |  |  |
| Price | double (formatted number)| `float64` |  | |  |  |
| VCPUs | int64 (formatted integer)| `int64` |  | |  |  |



### <span id="a-w-s-size-list"></span> AWSSizeList


  

[][AWSSize](#a-w-s-size)

### <span id="a-w-s-subnet"></span> AWSSubnet


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| AvailabilityZone | string| `string` |  | |  |  |
| AvailabilityZoneID | string| `string` |  | |  |  |
| AvailableIPAddressCount | int64 (formatted integer)| `int64` |  | |  |  |
| DefaultForAz | boolean| `bool` |  | |  |  |
| ID | string| `string` |  | |  |  |
| IPv4CIDR | string| `string` |  | |  |  |
| IPv6CIDR | string| `string` |  | |  |  |
| IsDefaultSubnet | boolean| `bool` |  | |  |  |
| Name | string| `string` |  | |  |  |
| State | string| `string` |  | |  |  |
| Tags | [][AWSTag](#a-w-s-tag)| `[]*AWSTag` |  | |  |  |



### <span id="a-w-s-subnet-list"></span> AWSSubnetList


  

[][AWSSubnet](#a-w-s-subnet)

### <span id="a-w-s-tag"></span> AWSTag


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Key | string| `string` |  | |  |  |
| Value | string| `string` |  | |  |  |



### <span id="a-w-s-v-p-c"></span> AWSVPC


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| CidrBlock | string| `string` |  | | The primary IPv4 CIDR block for the VPC. |  |
| CidrBlockAssociationSet | [][AWSVpcCidrBlockAssociation](#a-w-s-vpc-cidr-block-association)| `[]*AWSVpcCidrBlockAssociation` |  | | Information about the IPv4 CIDR blocks associated with the VPC. |  |
| DhcpOptionsID | string| `string` |  | | The ID of the set of DHCP options you've associated with the VPC (or default
if the default options are associated with the VPC). |  |
| InstanceTenancy | string| `string` |  | | The allowed tenancy of instances launched into the VPC. |  |
| Ipv6CidrBlockAssociationSet | [][AWSVpcIPV6CidrBlockAssociation](#a-w-s-vpc-ip-v6-cidr-block-association)| `[]*AWSVpcIPV6CidrBlockAssociation` |  | | Information about the IPv6 CIDR blocks associated with the VPC. |  |
| IsDefault | boolean| `bool` |  | | Indicates whether the VPC is the default VPC. |  |
| Name | string| `string` |  | |  |  |
| OwnerID | string| `string` |  | | The ID of the AWS account that owns the VPC. |  |
| State | string| `string` |  | | The current state of the VPC. |  |
| Tags | [][AWSTag](#a-w-s-tag)| `[]*AWSTag` |  | | Any tags assigned to the VPC. |  |
| VpcID | string| `string` |  | | The ID of the VPC. |  |



### <span id="a-w-s-v-p-c-list"></span> AWSVPCList


  

[][AWSVPC](#a-w-s-v-p-c)

### <span id="a-w-s-vpc-cidr-block-association"></span> AWSVpcCidrBlockAssociation


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| AssociationID | string| `string` |  | | The association ID for the IPv4 CIDR block. |  |
| CidrBlock | string| `string` |  | | The IPv4 CIDR block. |  |
| State | string| `string` |  | | The state of the CIDR block. |  |
| StatusMessage | string| `string` |  | | A message about the status of the CIDR block, if applicable. |  |



### <span id="a-w-s-vpc-ipv6-cidr-block-association"></span> AWSVpcIpv6CidrBlockAssociation


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| AssociationID | string| `string` |  | | The association ID for the IPv4 CIDR block. |  |
| CidrBlock | string| `string` |  | | The IPv4 CIDR block. |  |
| State | string| `string` |  | | The state of the CIDR block. |  |
| StatusMessage | string| `string` |  | | A message about the status of the CIDR block, if applicable. |  |



### <span id="absolute"></span> Absolute


> Usage is a struct that contains the current resource usage and limits
of a tenant.
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| MaxImageMeta | int64 (formatted integer)| `int64` |  | | MaxImageMeta is the amount of image metadata available to a tenant. |  |
| MaxPersonality | int64 (formatted integer)| `int64` |  | | MaxPersonality is the amount of personality/files available to a tenant. |  |
| MaxPersonalitySize | int64 (formatted integer)| `int64` |  | | MaxPersonalitySize is the personality file size available to a tenant. |  |
| MaxSecurityGroupRules | int64 (formatted integer)| `int64` |  | | MaxSecurityGroupRules is the number of security group rules available to
a tenant. |  |
| MaxSecurityGroups | int64 (formatted integer)| `int64` |  | | MaxSecurityGroups is the number of security groups available to a tenant. |  |
| MaxServerGroupMembers | int64 (formatted integer)| `int64` |  | | MaxServerGroupMembers is the number of server group members available
to a tenant. |  |
| MaxServerGroups | int64 (formatted integer)| `int64` |  | | MaxServerGroups is the number of server groups available to a tenant. |  |
| MaxServerMeta | int64 (formatted integer)| `int64` |  | | MaxServerMeta is the amount of server metadata available to a tenant. |  |
| MaxTotalCores | int64 (formatted integer)| `int64` |  | | MaxTotalCores is the number of cores available to a tenant. |  |
| MaxTotalFloatingIps | int64 (formatted integer)| `int64` |  | | MaxTotalFloatingIps is the number of floating IPs available to a tenant. |  |
| MaxTotalInstances | int64 (formatted integer)| `int64` |  | | MaxTotalInstances is the number of instances/servers available to a tenant. |  |
| MaxTotalKeypairs | int64 (formatted integer)| `int64` |  | | MaxTotalKeypairs is the total keypairs available to a tenant. |  |
| MaxTotalRAMSize | int64 (formatted integer)| `int64` |  | | MaxTotalRAMSize is the total amount of RAM available to a tenant measured
in megabytes (MB). |  |
| TotalCoresUsed | int64 (formatted integer)| `int64` |  | | TotalCoresUsed is the number of cores currently in use. |  |
| TotalFloatingIpsUsed | int64 (formatted integer)| `int64` |  | | TotalFloatingIpsUsed is the number of floating IPs in use. |  |
| TotalInstancesUsed | int64 (formatted integer)| `int64` |  | | TotalInstancesUsed is the number of instances/servers in use. |  |
| TotalRAMUsed | int64 (formatted integer)| `int64` |  | | TotalRAMUsed is the total RAM/memory in use measured in megabytes (MB). |  |
| TotalSecurityGroupsUsed | int64 (formatted integer)| `int64` |  | | TotalSecurityGroupsUsed is the total number of security groups in use. |  |
| TotalServerGroupsUsed | int64 (formatted integer)| `int64` |  | | TotalServerGroupsUsed is the total number of server groups in use. |  |



### <span id="accessible-addons"></span> AccessibleAddons


  

[]string

### <span id="addon"></span> Addon


> Addon represents a predefined addon that users may install into their cluster
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Annotations | map of string| `map[string]string` |  | | Annotations that can be added to the resource |  |
| CreationTimestamp | date-time (formatted string)| `strfmt.DateTime` |  | | CreationTimestamp is a timestamp representing the server time when this object was created. |  |
| DeletionTimestamp | date-time (formatted string)| `strfmt.DateTime` |  | | DeletionTimestamp is a timestamp representing the server time when this object was deleted. |  |
| ID | string| `string` |  | | ID unique value that identifies the resource generated by the server. Read-Only. |  |
| Name | string| `string` |  | | Name represents human readable name for the resource |  |
| spec | [AddonSpec](#addon-spec)| `AddonSpec` |  | |  |  |



### <span id="addon-config"></span> AddonConfig


> AddonConfig represents a addon configuration
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Annotations | map of string| `map[string]string` |  | | Annotations that can be added to the resource |  |
| CreationTimestamp | date-time (formatted string)| `strfmt.DateTime` |  | | CreationTimestamp is a timestamp representing the server time when this object was created. |  |
| DeletionTimestamp | date-time (formatted string)| `strfmt.DateTime` |  | | DeletionTimestamp is a timestamp representing the server time when this object was deleted. |  |
| ID | string| `string` |  | | ID unique value that identifies the resource generated by the server. Read-Only. |  |
| Name | string| `string` |  | | Name represents human readable name for the resource |  |
| spec | [AddonConfigSpec](#addon-config-spec)| `AddonConfigSpec` |  | |  |  |



### <span id="addon-config-spec"></span> AddonConfigSpec


> AddonConfigSpec specifies configuration of addon
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Controls | [][AddonFormControl](#addon-form-control)| `[]*AddonFormControl` |  | | Controls that can be set for configured addon |  |
| Description | string| `string` |  | | Description of the configured addon, it will be displayed in the addon overview in the UI |  |
| Logo | string| `string` |  | | Logo of the configured addon, encoded in base64 |  |
| LogoFormat | string| `string` |  | | LogoFormat contains logo format of the configured addon, i.e. svg+xml |  |
| ShortDescription | string| `string` |  | | ShortDescription of the configured addon that contains more detailed information about the addon,
it will be displayed in the addon details view in the UI |  |



### <span id="addon-form-control"></span> AddonFormControl


> AddonFormControl specifies addon form control
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| DisplayName | string| `string` |  | | DisplayName is visible in the UI |  |
| HelpText | string| `string` |  | | HelpText is visible in the UI next to the control |  |
| InternalName | string| `string` |  | | InternalName is used internally to save in the addon object |  |
| Required | boolean| `bool` |  | | Required indicates if the control has to be set |  |
| Type | string| `string` |  | | Type of displayed control |  |



### <span id="addon-spec"></span> AddonSpec


> AddonSpec addon specification
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| ContinuouslyReconcile | boolean| `bool` |  | | ContinuouslyReconcile indicates that the addon cannot be deleted or modified outside of the UI after installation |  |
| IsDefault | boolean| `bool` |  | | IsDefault indicates whether the addon is default |  |
| Variables | map of any | `map[string]interface{}` |  | | Variables is free form data to use for parsing the manifest templates |  |



### <span id="admin"></span> Admin


> Admin represents admin user
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Email | string| `string` |  | | Email address of the admin user |  |
| IsAdmin | boolean| `bool` |  | | IsAdmin indicates admin role |  |
| Name | string| `string` |  | | Name of the admin user |  |



### <span id="admission-plugin"></span> AdmissionPlugin


> AdmissionPlugin represents an admission plugin
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Name | string| `string` |  | |  |  |
| Plugin | string| `string` |  | |  |  |
| fromVersion | [Semver](#semver)| `Semver` |  | |  |  |



### <span id="admission-plugin-list"></span> AdmissionPluginList


> AdmissionPluginList represents a list of admission plugins
  



[]string

### <span id="alertmanager"></span> Alertmanager


> Alertmanager represents an Alertmanager Configuration
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| spec | [AlertmanagerSpec](#alertmanager-spec)| `AlertmanagerSpec` |  | |  |  |



### <span id="alertmanager-spec"></span> AlertmanagerSpec


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Config | []uint8 (formatted integer)| `[]uint8` |  | | Config contains the alertmanager configuration in YAML |  |



### <span id="alibaba-cloud-spec"></span> AlibabaCloudSpec


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| AccessKeyID | string| `string` |  | |  |  |
| AccessKeySecret | string| `string` |  | |  |  |
| credentialsReference | [GlobalSecretKeySelector](#global-secret-key-selector)| `GlobalSecretKeySelector` |  | |  |  |



### <span id="alibaba-instance-type"></span> AlibabaInstanceType


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| CPUCoreCount | int64 (formatted integer)| `int64` |  | |  |  |
| GPUCoreCount | int64 (formatted integer)| `int64` |  | |  |  |
| ID | string| `string` |  | |  |  |
| MemorySize | double (formatted number)| `float64` |  | |  |  |



### <span id="alibaba-instance-type-list"></span> AlibabaInstanceTypeList


  

[][AlibabaInstanceType](#alibaba-instance-type)

### <span id="alibaba-node-spec"></span> AlibabaNodeSpec


> AlibabaNodeSpec alibaba specific node settings
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| DiskSize | string| `string` |  | |  |  |
| DiskType | string| `string` |  | |  |  |
| InstanceType | string| `string` |  | |  |  |
| InternetMaxBandwidthOut | string| `string` |  | |  |  |
| Labels | map of string| `map[string]string` |  | |  |  |
| VSwitchID | string| `string` |  | |  |  |
| ZoneID | string| `string` |  | |  |  |



### <span id="alibaba-v-switch"></span> AlibabaVSwitch


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| ID | string| `string` |  | |  |  |



### <span id="alibaba-v-switch-list"></span> AlibabaVSwitchList


  

[][AlibabaVSwitch](#alibaba-v-switch)

### <span id="alibaba-zone"></span> AlibabaZone


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| ID | string| `string` |  | |  |  |



### <span id="alibaba-zone-list"></span> AlibabaZoneList


  

[][AlibabaZone](#alibaba-zone)

### <span id="allowed-registry"></span> AllowedRegistry


> AllowedRegistry represents a object containing a allowed image registry prefix
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Name | string| `string` |  | |  |  |
| spec | [AllowedRegistrySpec](#allowed-registry-spec)| `AllowedRegistrySpec` |  | |  |  |



### <span id="allowed-registry-spec"></span> AllowedRegistrySpec


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| RegistryPrefix | string| `string` |  | | RegistryPrefix contains the prefix of the registry which will be allowed. User clusters will be able to deploy
only images which are prefixed with one of the allowed image registry prefixes. |  |



### <span id="anexia-cloud-spec"></span> AnexiaCloudSpec


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Token | string| `string` |  | |  |  |
| credentialsReference | [GlobalSecretKeySelector](#global-secret-key-selector)| `GlobalSecretKeySelector` |  | |  |  |



### <span id="anexia-node-spec"></span> AnexiaNodeSpec


> AnexiaNodeSpec anexia specific node settings
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| CPUs | int64 (formatted integer)| `int64` | ✓ | | CPUs states how many cpus the node will have. |  |
| DiskSize | int64 (formatted integer)| `int64` | ✓ | | DiskSize states the disk size that node will have. |  |
| Memory | int64 (formatted integer)| `int64` | ✓ | | Memory states the memory that node will have. |  |
| TemplateID | string| `string` | ✓ | | TemplateID instance template |  |
| VlanID | string| `string` | ✓ | | VlanID Instance vlanID |  |



### <span id="anexia-template"></span> AnexiaTemplate


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| ID | string| `string` |  | |  |  |



### <span id="anexia-template-list"></span> AnexiaTemplateList


  

[][AnexiaTemplate](#anexia-template)

### <span id="anexia-vlan"></span> AnexiaVlan


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| ID | string| `string` |  | |  |  |



### <span id="anexia-vlan-list"></span> AnexiaVlanList


  

[][AnexiaVlan](#anexia-vlan)

### <span id="audit-logging-settings"></span> AuditLoggingSettings


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Enabled | boolean| `bool` |  | |  |  |



### <span id="azure-availability-zones-list"></span> AzureAvailabilityZonesList


> AzureAvailabilityZonesList is the object representing the availability zones for vms in azure cloud provider
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Zones | []string| `[]string` |  | |  |  |



### <span id="azure-cloud-spec"></span> AzureCloudSpec


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| AssignAvailabilitySet | boolean| `bool` |  | |  |  |
| AvailabilitySet | string| `string` |  | |  |  |
| ClientID | string| `string` |  | |  |  |
| ClientSecret | string| `string` |  | |  |  |
| OpenstackBillingTenant | string| `string` |  | |  |  |
| ResourceGroup | string| `string` |  | |  |  |
| RouteTableName | string| `string` |  | |  |  |
| SecurityGroup | string| `string` |  | |  |  |
| SubnetName | string| `string` |  | |  |  |
| SubscriptionID | string| `string` |  | |  |  |
| TenantID | string| `string` |  | |  |  |
| VNetName | string| `string` |  | |  |  |
| VNetResourceGroup | string| `string` |  | |  |  |
| credentialsReference | [GlobalSecretKeySelector](#global-secret-key-selector)| `GlobalSecretKeySelector` |  | |  |  |
| loadBalancerSKU | [LBSKU](#l-b-s-k-u)| `LBSKU` |  | |  |  |



### <span id="azure-node-spec"></span> AzureNodeSpec


> AzureNodeSpec describes settings for an Azure node
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| AssignAvailabilitySet | boolean| `bool` |  | | AssignAvailabilitySet is used to check if an availability set should be created and assigned to the cluster. |  |
| AssignPublicIP | boolean| `bool` |  | | should the machine have a publicly accessible IP address |  |
| DataDiskSize | int32 (formatted integer)| `int32` |  | | Data disk size in GB |  |
| ImageID | string| `string` |  | | ImageID represents the ID of the image that should be used to run the node |  |
| OSDiskSize | int32 (formatted integer)| `int32` |  | | OS disk size in GB |  |
| Size | string| `string` | ✓ | | VM size |  |
| Tags | map of string| `map[string]string` |  | | Additional metadata to set |  |
| Zones | []string| `[]string` |  | | Zones represents the availability zones for azure vms |  |



### <span id="azure-resource-groups-list"></span> AzureResourceGroupsList


> AzureResourceGroupsList is the object representing the resource groups for vms in azure cloud provider
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| ResourceGroups | []string| `[]string` |  | |  |  |



### <span id="azure-route-tables-list"></span> AzureRouteTablesList


> AzureRouteTablesList is the object representing the route tables for vms in azure cloud provider
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| RouteTables | []string| `[]string` |  | |  |  |



### <span id="azure-security-groups-list"></span> AzureSecurityGroupsList


> AzureSecurityGroupsList is the object representing the security groups for vms in azure cloud provider
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| SecurityGroups | []string| `[]string` |  | |  |  |



### <span id="azure-size"></span> AzureSize


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| MaxDataDiskCount | int32 (formatted integer)| `int32` |  | |  |  |
| MemoryInMB | int32 (formatted integer)| `int32` |  | |  |  |
| Name | string| `string` |  | |  |  |
| NumberOfCores | int32 (formatted integer)| `int32` |  | |  |  |
| NumberOfGPUs | int32 (formatted integer)| `int32` |  | |  |  |
| OsDiskSizeInMB | int32 (formatted integer)| `int32` |  | |  |  |
| ResourceDiskSizeInMB | int32 (formatted integer)| `int32` |  | |  |  |



### <span id="azure-size-list"></span> AzureSizeList


  

[][AzureSize](#azure-size)

### <span id="azure-subnets-list"></span> AzureSubnetsList


> AzureSubnetsList is the object representing the subnets for vms in azure cloud provider
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Subnets | []string| `[]string` |  | |  |  |



### <span id="azure-virtual-networks-list"></span> AzureVirtualNetworksList


> AzureVirtualNetworksList is the object representing the virtual network for vms in azure cloud provider
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| VirtualNetworks | []string| `[]string` |  | |  |  |



### <span id="backup-credentials"></span> BackupCredentials


> BackupCredentials contains credentials for etcd backups
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| s3 | [S3BackupCredentials](#s3-backup-credentials)| `S3BackupCredentials` |  | |  |  |



### <span id="backup-status"></span> BackupStatus


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| BackupMessage | string| `string` |  | |  |  |
| BackupName | string| `string` |  | |  |  |
| DeleteJobName | string| `string` |  | |  |  |
| DeleteMessage | string| `string` |  | |  |  |
| JobName | string| `string` |  | |  |  |
| backupFinishedTime | [Time](#time)| `Time` |  | |  |  |
| backupPhase | [BackupStatusPhase](#backup-status-phase)| `BackupStatusPhase` |  | |  |  |
| backupStartTime | [Time](#time)| `Time` |  | |  |  |
| deleteFinishedTime | [Time](#time)| `Time` |  | |  |  |
| deletePhase | [BackupStatusPhase](#backup-status-phase)| `BackupStatusPhase` |  | |  |  |
| deleteStartTime | [Time](#time)| `Time` |  | |  |  |
| scheduledTime | [Time](#time)| `Time` |  | |  |  |



### <span id="backup-status-phase"></span> BackupStatusPhase


  

| Name | Type | Go type | Default | Description | Example |
|------|------|---------| ------- |-------------|---------|
| BackupStatusPhase | string| string | |  |  |



### <span id="bring-your-own-cloud-spec"></span> BringYourOwnCloudSpec


  

[interface{}](#interface)

### <span id="by-pod-status"></span> ByPodStatus


> ByPodStatus defines the observed state of ConstraintTemplate as seen by
an individual controller
+kubebuilder:pruning:PreserveUnknownFields
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Errors | [][CreateCRDError](#create-c-r-d-error)| `[]*CreateCRDError` |  | |  |  |
| ID | string| `string` |  | | a unique identifier for the pod that wrote the status |  |
| ObservedGeneration | int64 (formatted integer)| `int64` |  | |  |  |



### <span id="c-r-d"></span> CRD


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| spec | [CRDSpec](#c-r-d-spec)| `CRDSpec` |  | |  |  |



### <span id="c-r-d-spec"></span> CRDSpec


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| names | [Names](#names)| `Names` |  | |  |  |
| validation | [Validation](#validation)| `Validation` |  | |  |  |



### <span id="cent-o-s-spec"></span> CentOSSpec


> CentOSSpec contains CentOS specific settings
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| DistUpgradeOnBoot | boolean| `bool` |  | | do a dist-upgrade on boot and reboot it required afterwards |  |



### <span id="cleanup-options"></span> CleanupOptions


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Enabled | boolean| `bool` |  | |  |  |
| Enforced | boolean| `bool` |  | |  |  |



### <span id="cloud-spec"></span> CloudSpec


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| DatacenterName | string| `string` |  | | DatacenterName where the users 'cloud' lives in. |  |
| alibaba | [AlibabaCloudSpec](#alibaba-cloud-spec)| `AlibabaCloudSpec` |  | |  |  |
| anexia | [AnexiaCloudSpec](#anexia-cloud-spec)| `AnexiaCloudSpec` |  | |  |  |
| aws | [AWSCloudSpec](#a-w-s-cloud-spec)| `AWSCloudSpec` |  | |  |  |
| azure | [AzureCloudSpec](#azure-cloud-spec)| `AzureCloudSpec` |  | |  |  |
| bringyourown | [BringYourOwnCloudSpec](#bring-your-own-cloud-spec)| `BringYourOwnCloudSpec` |  | |  |  |
| digitalocean | [DigitaloceanCloudSpec](#digitalocean-cloud-spec)| `DigitaloceanCloudSpec` |  | |  |  |
| fake | [FakeCloudSpec](#fake-cloud-spec)| `FakeCloudSpec` |  | |  |  |
| gcp | [GCPCloudSpec](#g-c-p-cloud-spec)| `GCPCloudSpec` |  | |  |  |
| hetzner | [HetznerCloudSpec](#hetzner-cloud-spec)| `HetznerCloudSpec` |  | |  |  |
| kubevirt | [KubevirtCloudSpec](#kubevirt-cloud-spec)| `KubevirtCloudSpec` |  | |  |  |
| openstack | [OpenstackCloudSpec](#openstack-cloud-spec)| `OpenstackCloudSpec` |  | |  |  |
| packet | [PacketCloudSpec](#packet-cloud-spec)| `PacketCloudSpec` |  | |  |  |
| vsphere | [VSphereCloudSpec](#v-sphere-cloud-spec)| `VSphereCloudSpec` |  | |  |  |



### <span id="cluster"></span> Cluster


> Note:
Cluster has a custom MarshalJSON method defined
and thus the output may vary
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Annotations | map of string| `map[string]string` |  | | Annotations that can be added to the resource |  |
| CreationTimestamp | date-time (formatted string)| `strfmt.DateTime` |  | | CreationTimestamp is a timestamp representing the server time when this object was created. |  |
| Credential | string| `string` |  | |  |  |
| DeletionTimestamp | date-time (formatted string)| `strfmt.DateTime` |  | | DeletionTimestamp is a timestamp representing the server time when this object was deleted. |  |
| ID | string| `string` |  | | ID unique value that identifies the resource generated by the server. Read-Only. |  |
| InheritedLabels | map of string| `map[string]string` |  | |  |  |
| Labels | map of string| `map[string]string` |  | |  |  |
| Name | string| `string` |  | | Name represents human readable name for the resource |  |
| Type | string| `string` |  | |  |  |
| spec | [ClusterSpec](#cluster-spec)| `ClusterSpec` |  | |  |  |
| status | [ClusterStatus](#cluster-status)| `ClusterStatus` |  | |  |  |



### <span id="cluster-health"></span> ClusterHealth


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| apiserver | [HealthStatus](#health-status)| `HealthStatus` |  | |  |  |
| cloudProviderInfrastructure | [HealthStatus](#health-status)| `HealthStatus` |  | |  |  |
| controller | [HealthStatus](#health-status)| `HealthStatus` |  | |  |  |
| etcd | [HealthStatus](#health-status)| `HealthStatus` |  | |  |  |
| gatekeeperAudit | [HealthStatus](#health-status)| `HealthStatus` |  | |  |  |
| gatekeeperController | [HealthStatus](#health-status)| `HealthStatus` |  | |  |  |
| machineController | [HealthStatus](#health-status)| `HealthStatus` |  | |  |  |
| scheduler | [HealthStatus](#health-status)| `HealthStatus` |  | |  |  |
| userClusterControllerManager | [HealthStatus](#health-status)| `HealthStatus` |  | |  |  |



### <span id="cluster-list"></span> ClusterList


> ClusterList represents a list of clusters
  



[][Cluster](#cluster)

### <span id="cluster-metrics"></span> ClusterMetrics


> ClusterMetrics defines a metric for the given cluster
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Name | string| `string` |  | |  |  |
| controlPlane | [ControlPlaneMetrics](#control-plane-metrics)| `ControlPlaneMetrics` |  | |  |  |
| nodes | [NodesMetric](#nodes-metric)| `NodesMetric` |  | |  |  |



### <span id="cluster-networking-config"></span> ClusterNetworkingConfig


> ClusterNetworkingConfig specifies the different networking
parameters for a cluster.
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| DNSDomain | string| `string` |  | | Domain name for services. |  |
| KonnectivityEnabled | boolean| `bool` |  | | KonnectivityEnabled enables konnectivity for controlplane to node network communication. |  |
| NodeLocalDNSCacheEnabled | boolean| `bool` |  | | NodeLocalDNSCacheEnabled controls whether the NodeLocal DNS Cache feature is enabled.
Defaults to true. |  |
| ProxyMode | string| `string` |  | | ProxyMode defines the kube-proxy mode (ipvs/iptables).
Defaults to ipvs. |  |
| ipvs | [IPVSConfiguration](#ip-v-s-configuration)| `IPVSConfiguration` |  | |  |  |
| pods | [NetworkRanges](#network-ranges)| `NetworkRanges` |  | |  |  |
| services | [NetworkRanges](#network-ranges)| `NetworkRanges` |  | |  |  |



### <span id="cluster-role"></span> ClusterRole


> ClusterRole defines cluster RBAC role for the user cluster
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Annotations | map of string| `map[string]string` |  | | Annotations that can be added to the resource |  |
| CreationTimestamp | date-time (formatted string)| `strfmt.DateTime` |  | | CreationTimestamp is a timestamp representing the server time when this object was created. |  |
| DeletionTimestamp | date-time (formatted string)| `strfmt.DateTime` |  | | DeletionTimestamp is a timestamp representing the server time when this object was deleted. |  |
| ID | string| `string` |  | | ID unique value that identifies the resource generated by the server. Read-Only. |  |
| Name | string| `string` |  | | Name represents human readable name for the resource |  |
| Rules | [][PolicyRule](#policy-rule)| `[]*PolicyRule` |  | | Rules holds all the PolicyRules for this ClusterRole |  |



### <span id="cluster-role-binding"></span> ClusterRoleBinding


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| RoleRefName | string| `string` |  | |  |  |
| Subjects | [][Subject](#subject)| `[]*Subject` |  | | Subjects holds references to the objects the role applies to. |  |



### <span id="cluster-role-name"></span> ClusterRoleName


> ClusterRoleName defines RBAC cluster role name object for the user cluster
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Name | string| `string` |  | | Name of the cluster role. |  |



### <span id="cluster-role-user"></span> ClusterRoleUser


> ClusterRoleUser defines associated user with cluster role
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Group | string| `string` |  | |  |  |
| UserEmail | string| `string` |  | |  |  |



### <span id="cluster-spec"></span> ClusterSpec


> ClusterSpec defines the cluster specification
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| AdmissionPlugins | []string| `[]string` |  | | Additional Admission Controller plugins |  |
| ContainerRuntime | string| `string` |  | | ContainerRuntime to use, i.e. Docker or containerd. By default containerd will be used. |  |
| EnableUserSSHKeyAgent | boolean| `bool` |  | | EnableUserSSHKeyAgent control whether the UserSSHKeyAgent will be deployed in the user cluster or not.
If it was enabled, the agent will be deployed and used to sync the user ssh keys, that the user attach
to the created cluster. If the agent was disabled, it won't be deployed in the user cluster, thus after
the cluster creation any attached ssh keys won't be synced to the worker nodes. Once the agent is enabled/disabled
it cannot be changed after the cluster is being created. |  |
| MachineNetworks | [][MachineNetworkingConfig](#machine-networking-config)| `[]*MachineNetworkingConfig` |  | | MachineNetworks optionally specifies the parameters for IPAM. |  |
| PodNodeSelectorAdmissionPluginConfig | map of string| `map[string]string` |  | | PodNodeSelectorAdmissionPluginConfig provides the configuration for the PodNodeSelector.
It's used by the backend to create a configuration file for this plugin.
The key:value from the map is converted to the namespace:<node-selectors-labels> in the file.
The format in a file:
podNodeSelectorPluginConfig:
clusterDefaultNodeSelector: <node-selectors-labels>
namespace1: <node-selectors-labels>
namespace2: <node-selectors-labels> |  |
| UsePodNodeSelectorAdmissionPlugin | boolean| `bool` |  | | If active the PodNodeSelector admission plugin is configured at the apiserver |  |
| UsePodSecurityPolicyAdmissionPlugin | boolean| `bool` |  | | If active the PodSecurityPolicy admission plugin is configured at the apiserver |  |
| auditLogging | [AuditLoggingSettings](#audit-logging-settings)| `AuditLoggingSettings` |  | |  |  |
| cloud | [CloudSpec](#cloud-spec)| `CloudSpec` |  | |  |  |
| clusterNetwork | [ClusterNetworkingConfig](#cluster-networking-config)| `ClusterNetworkingConfig` |  | |  |  |
| mla | [MLASettings](#m-l-a-settings)| `MLASettings` |  | |  |  |
| oidc | [OIDCSettings](#o-id-c-settings)| `OIDCSettings` |  | |  |  |
| opaIntegration | [OPAIntegrationSettings](#o-p-a-integration-settings)| `OPAIntegrationSettings` |  | |  |  |
| serviceAccount | [ServiceAccountSettings](#service-account-settings)| `ServiceAccountSettings` |  | |  |  |
| sys11auth | [Sys11AuthSettings](#sys11-auth-settings)| `Sys11AuthSettings` |  | |  |  |
| updateWindow | [UpdateWindow](#update-window)| `UpdateWindow` |  | |  |  |
| version | [Semver](#semver)| `Semver` |  | |  |  |



### <span id="cluster-status"></span> ClusterStatus


> ClusterStatus defines the cluster status
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| URL | string| `string` |  | | URL specifies the address at which the cluster is available |  |
| externalCCMMigration | [ExternalCCMMigrationStatus](#external-c-c-m-migration-status)| `ExternalCCMMigrationStatus` |  | |  |  |
| version | [Semver](#semver)| `Semver` |  | |  |  |



### <span id="cluster-template"></span> ClusterTemplate


> ClusterTemplate represents a ClusterTemplate object
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| ID | string| `string` |  | |  |  |
| Name | string| `string` |  | |  |  |
| ProjectID | string| `string` |  | |  |  |
| Scope | string| `string` |  | |  |  |
| User | string| `string` |  | |  |  |
| UserSSHKeys | [][ClusterTemplateSSHKey](#cluster-template-ssh-key)| `[]*ClusterTemplateSSHKey` |  | |  |  |
| cluster | [Cluster](#cluster)| `Cluster` |  | |  |  |
| nodeDeployment | [NodeDeployment](#node-deployment)| `NodeDeployment` |  | |  |  |



### <span id="cluster-template-instance"></span> ClusterTemplateInstance


> ClusterTemplateInstance represents a ClusterTemplateInstance object
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Name | string| `string` |  | |  |  |
| spec | [ClusterTemplateInstanceSpec](#cluster-template-instance-spec)| `ClusterTemplateInstanceSpec` |  | |  |  |



### <span id="cluster-template-instance-spec"></span> ClusterTemplateInstanceSpec


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| ClusterTemplateID | string| `string` |  | |  |  |
| ClusterTemplateName | string| `string` |  | |  |  |
| ProjectID | string| `string` |  | |  |  |
| Replicas | int64 (formatted integer)| `int64` |  | |  |  |



### <span id="cluster-template-list"></span> ClusterTemplateList


> ClusterTemplateList represents a ClusterTemplate list
  



[][ClusterTemplate](#cluster-template)

### <span id="cluster-template-ssh-key"></span> ClusterTemplateSSHKey


> ClusterTemplateSSHKey represents SSH Key object for Cluster Template
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| ID | string| `string` |  | |  |  |
| Name | string| `string` |  | |  |  |



### <span id="cluster-type"></span> ClusterType


  

| Name | Type | Go type | Default | Description | Example |
|------|------|---------| ------- |-------------|---------|
| ClusterType | int8 (formatted integer)| int8 | |  |  |



### <span id="condition-status"></span> ConditionStatus


  

| Name | Type | Go type | Default | Description | Example |
|------|------|---------| ------- |-------------|---------|
| ConditionStatus | string| string | |  |  |



### <span id="constraint"></span> Constraint


> Constraint represents a gatekeeper Constraint
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Labels | map of string| `map[string]string` |  | |  |  |
| Name | string| `string` |  | |  |  |
| spec | [ConstraintSpec](#constraint-spec)| `ConstraintSpec` |  | |  |  |
| status | [ConstraintStatus](#constraint-status)| `ConstraintStatus` |  | |  |  |



### <span id="constraint-selector"></span> ConstraintSelector


> ConstraintSelector is the object holding the cluster selection filters
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Providers | []string| `[]string` |  | | Providers is a list of cloud providers to which the Constraint applies to. Empty means all providers are selected. |  |
| labelSelector | [LabelSelector](#label-selector)| `LabelSelector` |  | |  |  |



### <span id="constraint-spec"></span> ConstraintSpec


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| ConstraintType | string| `string` |  | | ConstraintType specifies the type of gatekeeper constraint that the constraint applies to |  |
| Disabled | boolean| `bool` |  | | Disabled  is the flag for disabling OPA constraints |  |
| match | [Match](#match)| `Match` |  | |  |  |
| parameters | [Parameters](#parameters)| `Parameters` |  | |  |  |
| selector | [ConstraintSelector](#constraint-selector)| `ConstraintSelector` |  | |  |  |



### <span id="constraint-status"></span> ConstraintStatus


> ConstraintStatus represents a constraint status which holds audit info
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| AuditTimestamp | string| `string` |  | |  |  |
| Enforcement | string| `string` |  | |  |  |
| Synced | boolean| `bool` |  | |  |  |
| Violations | [][Violation](#violation)| `[]*Violation` |  | |  |  |



### <span id="constraint-template"></span> ConstraintTemplate


> ConstraintTemplate represents a gatekeeper ConstraintTemplate
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Name | string| `string` |  | |  |  |
| spec | [ConstraintTemplateSpec](#constraint-template-spec)| `ConstraintTemplateSpec` |  | |  |  |
| status | [ConstraintTemplateStatus](#constraint-template-status)| `ConstraintTemplateStatus` |  | |  |  |



### <span id="constraint-template-selector"></span> ConstraintTemplateSelector


> ConstraintTemplateSelector is the object holding the cluster selection filters
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Providers | []string| `[]string` |  | | Providers is a list of cloud providers to which the Constraint Template applies to. Empty means all providers are selected. |  |
| labelSelector | [LabelSelector](#label-selector)| `LabelSelector` |  | |  |  |



### <span id="constraint-template-spec"></span> ConstraintTemplateSpec


> ConstraintTemplateSpec is the object representing the gatekeeper constraint template spec and kubermatic related spec
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Targets | [][Target](#target)| `[]*Target` |  | |  |  |
| crd | [CRD](#c-r-d)| `CRD` |  | |  |  |
| selector | [ConstraintTemplateSelector](#constraint-template-selector)| `ConstraintTemplateSelector` |  | |  |  |



### <span id="constraint-template-status"></span> ConstraintTemplateStatus


> ConstraintTemplateStatus defines the observed state of ConstraintTemplate
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| ByPod | [][ByPodStatus](#by-pod-status)| `[]*ByPodStatus` |  | |  |  |
| Created | boolean| `bool` |  | |  |  |



### <span id="control-plane-metrics"></span> ControlPlaneMetrics


> ControlPlaneMetrics defines a metric for the user cluster control plane resources
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| CPUTotalMillicores | int64 (formatted integer)| `int64` |  | | CPUTotalMillicores in m cores |  |
| MemoryTotalBytes | int64 (formatted integer)| `int64` |  | | MemoryTotalBytes in bytes |  |



### <span id="create-c-r-d-error"></span> CreateCRDError


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Code | string| `string` |  | |  |  |
| Location | string| `string` |  | |  |  |
| Message | string| `string` |  | |  |  |



### <span id="create-cluster-spec"></span> CreateClusterSpec


> CreateClusterSpec is the structure that is used to create cluster with its initial node deployment
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| DNSDomain | string| `string` |  | |  |  |
| PodsCIDR | string| `string` |  | |  |  |
| ServicesCIDR | string| `string` |  | |  |  |
| cluster | [Cluster](#cluster)| `Cluster` |  | |  |  |
| nodeDeployment | [NodeDeployment](#node-deployment)| `NodeDeployment` |  | |  |  |



### <span id="credential-list"></span> CredentialList


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Names | []string| `[]string` |  | |  |  |



### <span id="custom-link"></span> CustomLink


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Icon | string| `string` |  | |  |  |
| Label | string| `string` |  | |  |  |
| Location | string| `string` |  | |  |  |
| URL | string| `string` |  | |  |  |



### <span id="custom-links"></span> CustomLinks


  

[][CustomLink](#custom-link)

### <span id="datacenter"></span> Datacenter


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| metadata | [DatacenterMeta](#datacenter-meta)| `DatacenterMeta` |  | |  |  |
| spec | [DatacenterSpec](#datacenter-spec)| `DatacenterSpec` |  | |  |  |



### <span id="datacenter-list"></span> DatacenterList


> DatacenterList represents a list of datacenters
  



[][Datacenter](#datacenter)

### <span id="datacenter-meta"></span> DatacenterMeta


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Name | string| `string` |  | |  |  |



### <span id="datacenter-spec"></span> DatacenterSpec


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Country | string| `string` |  | | Optional: Country of the seed as ISO-3166 two-letter code, e.g. DE or UK.
It is used for informational purposes. |  |
| EnforceAuditLogging | boolean| `bool` |  | | EnforceAuditLogging enforces audit logging on every cluster within the DC,
ignoring cluster-specific settings. |  |
| EnforcePodSecurityPolicy | boolean| `bool` |  | | EnforcePodSecurityPolicy enforces pod security policy plugin on every clusters within the DC,
ignoring cluster-specific settings |  |
| Location | string| `string` |  | | Optional: Detailed location of the cluster, like "Hamburg" or "Datacenter 7".
It is used for informational purposes. |  |
| Provider | string| `string` |  | | Name of the datacenter provider. Extracted based on which provider is defined in the spec.
It is used for informational purposes. |  |
| RequiredEmailDomain | string| `string` |  | | Deprecated. Automatically migrated to the RequiredEmailDomains field. |  |
| RequiredEmailDomains | []string| `[]string` |  | |  |  |
| Seed | string| `string` |  | | Name of the seed this datacenter belongs to. |  |
| alibaba | [DatacenterSpecAlibaba](#datacenter-spec-alibaba)| `DatacenterSpecAlibaba` |  | |  |  |
| anexia | [DatacenterSpecAnexia](#datacenter-spec-anexia)| `DatacenterSpecAnexia` |  | |  |  |
| aws | [DatacenterSpecAWS](#datacenter-spec-a-w-s)| `DatacenterSpecAWS` |  | |  |  |
| azure | [DatacenterSpecAzure](#datacenter-spec-azure)| `DatacenterSpecAzure` |  | |  |  |
| bringyourown | [DatacenterSpecBringYourOwn](#datacenter-spec-bring-your-own)| `DatacenterSpecBringYourOwn` |  | |  |  |
| digitalocean | [DatacenterSpecDigitalocean](#datacenter-spec-digitalocean)| `DatacenterSpecDigitalocean` |  | |  |  |
| fake | [DatacenterSpecFake](#datacenter-spec-fake)| `DatacenterSpecFake` |  | |  |  |
| gcp | [DatacenterSpecGCP](#datacenter-spec-g-c-p)| `DatacenterSpecGCP` |  | |  |  |
| hetzner | [DatacenterSpecHetzner](#datacenter-spec-hetzner)| `DatacenterSpecHetzner` |  | |  |  |
| kubevirt | [DatacenterSpecKubevirt](#datacenter-spec-kubevirt)| `DatacenterSpecKubevirt` |  | |  |  |
| node | [NodeSettings](#node-settings)| `NodeSettings` |  | |  |  |
| openstack | [DatacenterSpecOpenstack](#datacenter-spec-openstack)| `DatacenterSpecOpenstack` |  | |  |  |
| packet | [DatacenterSpecPacket](#datacenter-spec-packet)| `DatacenterSpecPacket` |  | |  |  |
| vsphere | [DatacenterSpecVSphere](#datacenter-spec-v-sphere)| `DatacenterSpecVSphere` |  | |  |  |



### <span id="datacenter-spec-a-w-s"></span> DatacenterSpecAWS


> DatacenterSpecAWS describes an AWS datacenter
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Region | string| `string` |  | | The AWS region to use, e.g. "us-east-1". For a list of available regions, see
https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-regions-availability-zones.html |  |
| images | [ImageList](#image-list)| `ImageList` |  | |  |  |



### <span id="datacenter-spec-alibaba"></span> DatacenterSpecAlibaba


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Region | string| `string` |  | | Region to use, for a full list of regions see
https://www.alibabacloud.com/help/doc-detail/40654.htm |  |



### <span id="datacenter-spec-anexia"></span> DatacenterSpecAnexia


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| LocationID | string| `string` |  | | LocationID the location of the region |  |



### <span id="datacenter-spec-azure"></span> DatacenterSpecAzure


> DatacenterSpecAzure describes an Azure cloud datacenter
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Location | string| `string` |  | | Region to use, for example "westeurope". A list of available regions can be
found at https://azure.microsoft.com/en-us/global-infrastructure/locations/ |  |



### <span id="datacenter-spec-bring-your-own"></span> DatacenterSpecBringYourOwn


> DatacenterSpecBringYourOwn describes a datacenter our of bring your own nodes
  



[interface{}](#interface)

### <span id="datacenter-spec-digitalocean"></span> DatacenterSpecDigitalocean


> DatacenterSpecDigitalocean describes a DigitalOcean datacenter
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Region | string| `string` |  | | Datacenter location, e.g. "ams3". A list of existing datacenters can be found
at https://www.digitalocean.com/docs/platform/availability-matrix/ |  |



### <span id="datacenter-spec-fake"></span> DatacenterSpecFake


> DatacenterSpecFake describes a fake datacenter
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| FakeProperty | string| `string` |  | |  |  |



### <span id="datacenter-spec-g-c-p"></span> DatacenterSpecGCP


> DatacenterSpecGCP describes a GCP datacenter
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Region | string| `string` |  | | Region to use, for example "europe-west3", for a full list of regions see
https://cloud.google.com/compute/docs/regions-zones/ |  |
| Regional | boolean| `bool` |  | | Optional: Regional clusters spread their resources across multiple availability zones.
Refer to the official documentation for more details on this:
https://cloud.google.com/kubernetes-engine/docs/concepts/regional-clusters |  |
| ZoneSuffixes | []string| `[]string` |  | | List of enabled zones, for example [a, c]. See the link above for the available
zones in your chosen region. |  |



### <span id="datacenter-spec-hetzner"></span> DatacenterSpecHetzner


> DatacenterSpecHetzner describes a Hetzner cloud datacenter
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Datacenter | string| `string` |  | | Datacenter location, e.g. "nbg1-dc3". A list of existing datacenters can be found
at https://wiki.hetzner.de/index.php/Rechenzentren_und_Anbindung/en |  |
| Location | string| `string` |  | | Optional: Detailed location of the datacenter, like "Hamburg" or "Datacenter 7".
For informational purposes only. |  |
| Network | string| `string` |  | | Network is the pre-existing Hetzner network in which the machines are running.
While machines can be in multiple networks, a single one must be chosen for the
HCloud CCM to work. |  |



### <span id="datacenter-spec-kubevirt"></span> DatacenterSpecKubevirt


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| DNSPolicy | string| `string` |  | | DNSPolicy represents the dns policy for the pod. Valid values are 'ClusterFirstWithHostNet', 'ClusterFirst',
'Default' or 'None'. Defaults to "ClusterFirst". DNS parameters given in DNSConfig will be merged with the
policy selected with DNSPolicy. |  |
| dns_config | [PodDNSConfig](#pod-dns-config)| `PodDNSConfig` |  | |  |  |



### <span id="datacenter-spec-openstack"></span> DatacenterSpecOpenstack


> DatacenterSpecOpenstack describes an OpenStack datacenter
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| AuthURL | string| `string` |  | |  |  |
| AvailabilityZone | string| `string` |  | |  |  |
| DNSServers | []string| `[]string` |  | | Used for automatic network creation |  |
| EnabledFlavors | []string| `[]string` |  | | Optional: List of enabled flavors for the given datacenter |  |
| EnforceFloatingIP | boolean| `bool` |  | | Optional |  |
| IgnoreVolumeAZ | boolean| `bool` |  | | Optional |  |
| ManageSecurityGroups | boolean| `bool` |  | | Optional: Gets mapped to the "manage-security-groups" setting in the cloud config.
See https://kubernetes.io/docs/concepts/cluster-administration/cloud-providers/#load-balancer
This setting defaults to true. |  |
| NodeVolumeAttachLimit | uint64 (formatted integer)| `uint64` |  | |  |  |
| Region | string| `string` |  | |  |  |
| TrustDevicePath | boolean| `bool` |  | | Optional: Gets mapped to the "trust-device-path" setting in the cloud config.
See https://kubernetes.io/docs/concepts/cluster-administration/cloud-providers/#block-storage
This setting defaults to false. |  |
| UseOctavia | boolean| `bool` |  | | Optional: Gets mapped to the "use-octavia" setting in the cloud config.
use-octavia is enabled by default in CCM since v1.17.0, and disabled by
default with the in-tree cloud provider. |  |
| images | [ImageList](#image-list)| `ImageList` |  | |  |  |
| node_size_requirements | [OpenstackNodeSizeRequirements](#openstack-node-size-requirements)| `OpenstackNodeSizeRequirements` |  | |  |  |



### <span id="datacenter-spec-packet"></span> DatacenterSpecPacket


> DatacenterSpecPacket describes a Packet datacenter
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Facilities | []string| `[]string` |  | | The list of enabled facilities, for example "ams1", for a full list of available
facilities see https://support.packet.com/kb/articles/data-centers |  |



### <span id="datacenter-spec-v-sphere"></span> DatacenterSpecVSphere


> DatacenterSpecVSphere describes a vSphere datacenter
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| AllowInsecure | boolean| `bool` |  | | If set to true, disables the TLS certificate check against the endpoint. |  |
| Cluster | string| `string` |  | | Optional: The name of the vSphere cluster to use.
Cluster is deprecated and may be removed in future releases as it is
currently ignored.
The cluster hosting the VMs will be the same VM used as a template is
located. |  |
| Datacenter | string| `string` |  | | The name of the datacenter to use. |  |
| DefaultDatastore | string| `string` |  | | The default Datastore to be used for provisioning volumes using storage
classes/dynamic provisioning and for storing virtual machine files in
case no `Datastore` or `DatastoreCluster` is provided at Cluster level. |  |
| DefaultStoragePolicy | string| `string` |  | | The name of the storage policy to use for the storage class created in the user cluster. |  |
| Endpoint | string| `string` |  | | Endpoint URL to use, including protocol, for example "https://vcenter.example.com". |  |
| RootPath | string| `string` |  | | Optional: The root path for cluster specific VM folders. Each cluster gets its own
folder below the root folder. Must be the FQDN (for example
"/datacenter-1/vm/all-kubermatic-vms-in-here") and defaults to the root VM
folder: "/datacenter-1/vm" |  |
| infra_management_user | [VSphereCredentials](#v-sphere-credentials)| `VSphereCredentials` |  | |  |  |
| templates | [ImageList](#image-list)| `ImageList` |  | |  |  |



### <span id="digitalocean-cloud-spec"></span> DigitaloceanCloudSpec


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Token | string| `string` |  | |  |  |
| credentialsReference | [GlobalSecretKeySelector](#global-secret-key-selector)| `GlobalSecretKeySelector` |  | |  |  |



### <span id="digitalocean-node-spec"></span> DigitaloceanNodeSpec


> DigitaloceanNodeSpec digitalocean node settings
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Backups | boolean| `bool` |  | | enable backups for the droplet |  |
| IPv6 | boolean| `bool` |  | | enable ipv6 for the droplet |  |
| Monitoring | boolean| `bool` |  | | enable monitoring for the droplet |  |
| Size | string| `string` | ✓ | | droplet size slug |  |
| Tags | []string| `[]string` |  | | additional droplet tags |  |



### <span id="digitalocean-size"></span> DigitaloceanSize


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Available | boolean| `bool` |  | |  |  |
| Disk | int64 (formatted integer)| `int64` |  | |  |  |
| Memory | int64 (formatted integer)| `int64` |  | |  |  |
| PriceHourly | double (formatted number)| `float64` |  | |  |  |
| PriceMonthly | double (formatted number)| `float64` |  | |  |  |
| Regions | []string| `[]string` |  | |  |  |
| Slug | string| `string` |  | |  |  |
| Transfer | double (formatted number)| `float64` |  | |  |  |
| VCPUs | int64 (formatted integer)| `int64` |  | |  |  |



### <span id="digitalocean-size-list"></span> DigitaloceanSizeList


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Optimized | [][DigitaloceanSize](#digitalocean-size)| `[]*DigitaloceanSize` |  | |  |  |
| Standard | [][DigitaloceanSize](#digitalocean-size)| `[]*DigitaloceanSize` |  | |  |  |



### <span id="error-details"></span> ErrorDetails


> ErrorDetails contains details about the error
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Additional | []string| `[]string` |  | | Additional error messages |  |
| Code | int64 (formatted integer)| `int64` | ✓ | | The error code |  |
| Message | string| `string` | ✓ | | The error message |  |



### <span id="error-response"></span> ErrorResponse


> ErrorResponse is the default representation of an error
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| error | [ErrorDetails](#error-details)| `ErrorDetails` |  | |  |  |



### <span id="etcd-backup-config"></span> EtcdBackupConfig


> EtcdBackupConfig represents an object holding the configuration for etcd backups
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Annotations | map of string| `map[string]string` |  | | Annotations that can be added to the resource |  |
| CreationTimestamp | date-time (formatted string)| `strfmt.DateTime` |  | | CreationTimestamp is a timestamp representing the server time when this object was created. |  |
| DeletionTimestamp | date-time (formatted string)| `strfmt.DateTime` |  | | DeletionTimestamp is a timestamp representing the server time when this object was deleted. |  |
| ID | string| `string` |  | | ID unique value that identifies the resource generated by the server. Read-Only. |  |
| Name | string| `string` |  | | Name represents human readable name for the resource |  |
| spec | [EtcdBackupConfigSpec](#etcd-backup-config-spec)| `EtcdBackupConfigSpec` |  | |  |  |
| status | [EtcdBackupConfigStatus](#etcd-backup-config-status)| `EtcdBackupConfigStatus` |  | |  |  |



### <span id="etcd-backup-config-condition"></span> EtcdBackupConfigCondition


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Message | string| `string` |  | | Human readable message indicating details about last transition.
+optional |  |
| Reason | string| `string` |  | | (brief) reason for the condition's last transition.
+optional |  |
| lastHeartbeatTime | [Time](#time)| `Time` |  | |  |  |
| lastTransitionTime | [Time](#time)| `Time` |  | |  |  |
| status | [ConditionStatus](#condition-status)| `ConditionStatus` |  | |  |  |
| type | [EtcdBackupConfigConditionType](#etcd-backup-config-condition-type)| `EtcdBackupConfigConditionType` |  | |  |  |



### <span id="etcd-backup-config-condition-type"></span> EtcdBackupConfigConditionType


> EtcdBackupConfigConditionType is used to indicate the type of a EtcdBackupConfig condition. For all condition
types, the `true` value must indicate success. All condition types must be registered within
the `AllClusterConditionTypes` variable.
  



| Name | Type | Go type | Default | Description | Example |
|------|------|---------| ------- |-------------|---------|
| EtcdBackupConfigConditionType | string| string | | EtcdBackupConfigConditionType is used to indicate the type of a EtcdBackupConfig condition. For all condition
types, the `true` value must indicate success. All condition types must be registered within
the `AllClusterConditionTypes` variable. |  |



### <span id="etcd-backup-config-spec"></span> EtcdBackupConfigSpec


> EtcdBackupConfigSpec represents an object holding the etcd backup configuration specification
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| ClusterID | string| `string` |  | | ClusterID is the id of the cluster which will be backed up |  |
| Keep | int64 (formatted integer)| `int64` |  | | Keep is the number of backups to keep around before deleting the oldest one
If not set, defaults to DefaultKeptBackupsCount. Only used if Schedule is set. |  |
| Schedule | string| `string` |  | | Schedule is a cron expression defining when to perform
the backup. If not set, the backup is performed exactly
once, immediately. |  |



### <span id="etcd-backup-config-status"></span> EtcdBackupConfigStatus


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| CleanupRunning | boolean| `bool` |  | | If the controller was configured with a cleanupContainer, CleanupRunning keeps track of the corresponding job |  |
| Conditions | [][EtcdBackupConfigCondition](#etcd-backup-config-condition)| `[]*EtcdBackupConfigCondition` |  | | Conditions contains conditions of the EtcdBackupConfig |  |
| CurrentBackups | [][BackupStatus](#backup-status)| `[]*BackupStatus` |  | | CurrentBackups tracks the creation and deletion progress if all backups managed by the EtcdBackupConfig |  |



### <span id="etcd-restore"></span> EtcdRestore


> EtcdRestore represents an object holding the configuration for etcd backup restore
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Name | string| `string` |  | |  |  |
| spec | [EtcdRestoreSpec](#etcd-restore-spec)| `EtcdRestoreSpec` |  | |  |  |
| status | [EtcdRestoreStatus](#etcd-restore-status)| `EtcdRestoreStatus` |  | |  |  |



### <span id="etcd-restore-phase"></span> EtcdRestorePhase


  

| Name | Type | Go type | Default | Description | Example |
|------|------|---------| ------- |-------------|---------|
| EtcdRestorePhase | string| string | |  |  |



### <span id="etcd-restore-spec"></span> EtcdRestoreSpec


> EtcdRestoreSpec represents an object holding the etcd backup restore configuration specification
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| BackupDownloadCredentialsSecret | string| `string` |  | | BackupDownloadCredentialsSecret is the name of a secret in the cluster-xxx namespace containing
credentials needed to download the backup |  |
| BackupName | string| `string` |  | | BackupName is the name of the backup to restore from |  |
| ClusterID | string| `string` |  | | ClusterID is the id of the cluster which will be restored from the backup |  |



### <span id="etcd-restore-status"></span> EtcdRestoreStatus


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| phase | [EtcdRestorePhase](#etcd-restore-phase)| `EtcdRestorePhase` |  | |  |  |
| restoreTime | [Time](#time)| `Time` |  | |  |  |



### <span id="event"></span> Event


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Annotations | map of string| `map[string]string` |  | | Annotations that can be added to the resource |  |
| Count | int32 (formatted integer)| `int32` |  | | The number of times this event has occurred. |  |
| CreationTimestamp | date-time (formatted string)| `strfmt.DateTime` |  | | CreationTimestamp is a timestamp representing the server time when this object was created. |  |
| DeletionTimestamp | date-time (formatted string)| `strfmt.DateTime` |  | | DeletionTimestamp is a timestamp representing the server time when this object was deleted. |  |
| ID | string| `string` |  | | ID unique value that identifies the resource generated by the server. Read-Only. |  |
| LastTimestamp | date-time (formatted string)| `strfmt.DateTime` |  | | The time at which the most recent occurrence of this event was recorded. |  |
| Message | string| `string` |  | | A human-readable description of the status of this operation. |  |
| Name | string| `string` |  | | Name represents human readable name for the resource |  |
| Type | string| `string` |  | | Type of this event (i.e. normal or warning). New types could be added in the future. |  |
| involvedObject | [ObjectReferenceResource](#object-reference-resource)| `ObjectReferenceResource` |  | |  |  |



### <span id="expose-strategy"></span> ExposeStrategy


  

| Name | Type | Go type | Default | Description | Example |
|------|------|---------| ------- |-------------|---------|
| ExposeStrategy | string| string | |  |  |



### <span id="external-c-c-m-migration-status"></span> ExternalCCMMigrationStatus


  

| Name | Type | Go type | Default | Description | Example |
|------|------|---------| ------- |-------------|---------|
| ExternalCCMMigrationStatus | string| string | |  |  |



### <span id="external-documentation"></span> ExternalDocumentation


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Description | string| `string` |  | |  |  |
| URL | string| `string` |  | |  |  |



### <span id="fake-cloud-spec"></span> FakeCloudSpec


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Token | string| `string` |  | |  |  |



### <span id="flatcar-spec"></span> FlatcarSpec


> FlatcarSpec contains Flatcar Linux specific settings
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| DisableAutoUpdate | boolean| `bool` |  | | disable flatcar linux auto-update feature |  |



### <span id="g-c-p-cloud-spec"></span> GCPCloudSpec


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Network | string| `string` |  | |  |  |
| ServiceAccount | string| `string` |  | |  |  |
| Subnetwork | string| `string` |  | |  |  |
| credentialsReference | [GlobalSecretKeySelector](#global-secret-key-selector)| `GlobalSecretKeySelector` |  | |  |  |



### <span id="g-c-p-disk-type"></span> GCPDiskType


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Description | string| `string` |  | |  |  |
| Name | string| `string` |  | |  |  |



### <span id="g-c-p-disk-type-list"></span> GCPDiskTypeList


  

[][GCPDiskType](#g-c-p-disk-type)

### <span id="g-c-p-machine-size"></span> GCPMachineSize


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Description | string| `string` |  | |  |  |
| Memory | int64 (formatted integer)| `int64` |  | |  |  |
| Name | string| `string` |  | |  |  |
| VCPUs | int64 (formatted integer)| `int64` |  | |  |  |



### <span id="g-c-p-machine-size-list"></span> GCPMachineSizeList


  

[][GCPMachineSize](#g-c-p-machine-size)

### <span id="g-c-p-network"></span> GCPNetwork


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| AutoCreateSubnetworks | boolean| `bool` |  | |  |  |
| ID | uint64 (formatted integer)| `uint64` |  | |  |  |
| Kind | string| `string` |  | |  |  |
| Name | string| `string` |  | |  |  |
| Path | string| `string` |  | |  |  |
| Subnetworks | []string| `[]string` |  | |  |  |



### <span id="g-c-p-network-list"></span> GCPNetworkList


  

[][GCPNetwork](#g-c-p-network)

### <span id="g-c-p-node-spec"></span> GCPNodeSpec


> GCPNodeSpec gcp specific node settings
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| CustomImage | string| `string` |  | |  |  |
| DiskSize | int64 (formatted integer)| `int64` |  | |  |  |
| DiskType | string| `string` |  | |  |  |
| Labels | map of string| `map[string]string` |  | |  |  |
| MachineType | string| `string` |  | |  |  |
| Preemptible | boolean| `bool` |  | |  |  |
| Tags | []string| `[]string` |  | |  |  |
| Zone | string| `string` |  | |  |  |



### <span id="g-c-p-subnetwork"></span> GCPSubnetwork


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| GatewayAddress | string| `string` |  | |  |  |
| ID | uint64 (formatted integer)| `uint64` |  | |  |  |
| IPCidrRange | string| `string` |  | |  |  |
| Kind | string| `string` |  | |  |  |
| Name | string| `string` |  | |  |  |
| Network | string| `string` |  | |  |  |
| Path | string| `string` |  | |  |  |
| PrivateIPGoogleAccess | boolean| `bool` |  | |  |  |
| Region | string| `string` |  | |  |  |
| SelfLink | string| `string` |  | |  |  |



### <span id="g-c-p-subnetwork-list"></span> GCPSubnetworkList


  

[][GCPSubnetwork](#g-c-p-subnetwork)

### <span id="g-c-p-zone"></span> GCPZone


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Name | string| `string` |  | |  |  |



### <span id="g-c-p-zone-list"></span> GCPZoneList


  

[][GCPZone](#g-c-p-zone)

### <span id="g-v-k"></span> GVK


> GVK group version kind of a resource
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Group | string| `string` |  | |  |  |
| Kind | string| `string` |  | |  |  |
| Version | string| `string` |  | |  |  |



### <span id="gatekeeper-config"></span> GatekeeperConfig


> GatekeeperConfig represents a gatekeeper config
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| spec | [GatekeeperConfigSpec](#gatekeeper-config-spec)| `GatekeeperConfigSpec` |  | |  |  |



### <span id="gatekeeper-config-spec"></span> GatekeeperConfigSpec


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Match | [][MatchEntry](#match-entry)| `[]*MatchEntry` |  | | Configuration for namespace exclusion |  |
| readiness | [ReadinessSpec](#readiness-spec)| `ReadinessSpec` |  | |  |  |
| sync | [Sync](#sync)| `Sync` |  | |  |  |
| validation | [Validation](#validation)| `Validation` |  | |  |  |



### <span id="global-custom-links"></span> GlobalCustomLinks


> GlobalCustomLinks defines custom links for global settings
  



[][CustomLink](#custom-link)

### <span id="global-object-key-selector"></span> GlobalObjectKeySelector


> GlobalObjectKeySelector is needed as we can not use v1.SecretKeySelector
because it is not cross namespace
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| APIVersion | string| `string` |  | | API version of the referent.
+optional |  |
| FieldPath | string| `string` |  | | If referring to a piece of an object instead of an entire object, this string
should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2].
For example, if the object reference is to a container within a pod, this would take on a value like:
"spec.containers{name}" (where "name" refers to the name of the container that triggered
the event) or if no container name is specified "spec.containers[2]" (container with
index 2 in this pod). This syntax is chosen only to have some well-defined way of
referencing a part of an object.
TODO: this design is not final and this field is subject to change in the future.
+optional |  |
| Key | string| `string` |  | |  |  |
| Kind | string| `string` |  | | Kind of the referent.
More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
+optional |  |
| Name | string| `string` |  | | Name of the referent.
More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
+optional |  |
| Namespace | string| `string` |  | | Namespace of the referent.
More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/
+optional |  |
| ResourceVersion | string| `string` |  | | Specific resourceVersion to which this reference is made, if any.
More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency
+optional |  |
| uid | [UID](#uid)| `UID` |  | |  |  |



### <span id="global-secret-key-selector"></span> GlobalSecretKeySelector


  


* composed type [GlobalObjectKeySelector](#global-object-key-selector)

### <span id="global-settings"></span> GlobalSettings


> GlobalSettings defines global settings
  




* composed type [SettingSpec](#setting-spec)

### <span id="health-status"></span> HealthStatus


  

| Name | Type | Go type | Default | Description | Example |
|------|------|---------| ------- |-------------|---------|
| HealthStatus | int64 (formatted integer)| int64 | |  |  |



### <span id="hetzner-cloud-spec"></span> HetznerCloudSpec


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Network | string| `string` |  | | Network is the pre-existing Hetzner network in which the machines are running.
While machines can be in multiple networks, a single one must be chosen for the
HCloud CCM to work.
If this is empty, the network configured on the datacenter will be used. |  |
| Token | string| `string` |  | | Token is used to authenticate with the Hetzner cloud API. |  |
| credentialsReference | [GlobalSecretKeySelector](#global-secret-key-selector)| `GlobalSecretKeySelector` |  | |  |  |



### <span id="hetzner-node-spec"></span> HetznerNodeSpec


> HetznerNodeSpec Hetzner node settings
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Network | string| `string` |  | | network name |  |
| Type | string| `string` | ✓ | | server type |  |



### <span id="hetzner-size"></span> HetznerSize


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Cores | int64 (formatted integer)| `int64` |  | |  |  |
| Description | string| `string` |  | |  |  |
| Disk | int64 (formatted integer)| `int64` |  | |  |  |
| ID | int64 (formatted integer)| `int64` |  | |  |  |
| Memory | float (formatted number)| `float32` |  | |  |  |
| Name | string| `string` |  | |  |  |



### <span id="hetzner-size-list"></span> HetznerSizeList


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Dedicated | [][HetznerSize](#hetzner-size)| `[]*HetznerSize` |  | |  |  |
| Standard | [][HetznerSize](#hetzner-size)| `[]*HetznerSize` |  | |  |  |



### <span id="ip-v-s-configuration"></span> IPVSConfiguration


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| StrictArp | boolean| `bool` |  | | StrictArp configure arp_ignore and arp_announce to avoid answering ARP queries from kube-ipvs0 interface.
defaults to true. |  |



### <span id="image"></span> Image


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Created | string| `string` |  | | Created is the date when the image was created. |  |
| ID | string| `string` |  | | ID is the unique ID of an image. |  |
| Metadata | map of any | `map[string]interface{}` |  | | Metadata provides free-form key/value pairs that further describe the
image. |  |
| MinDisk | int64 (formatted integer)| `int64` |  | | MinDisk is the minimum amount of disk a flavor must have to be able
to create a server based on the image, measured in GB. |  |
| MinRAM | int64 (formatted integer)| `int64` |  | | MinRAM is the minimum amount of RAM a flavor must have to be able
to create a server based on the image, measured in MB. |  |
| Name | string| `string` |  | | Name provides a human-readable moniker for the OS image. |  |
| Progress | int64 (formatted integer)| `int64` |  | | The Progress and Status fields indicate image-creation status. |  |
| Status | string| `string` |  | | Status is the current status of the image. |  |
| Updated | string| `string` |  | | Update is the date when the image was updated. |  |



### <span id="image-list"></span> ImageList


> ImageList defines a map of operating system and the image to use
  



[ImageList](#image-list)

### <span id="json"></span> JSON


> These types are supported: bool, int64, float64, string, []interface{}, map[string]interface{} and nil.
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Raw | []uint8 (formatted integer)| `[]uint8` |  | |  |  |



### <span id="json-schema-definitions"></span> JSONSchemaDefinitions


  

[JSONSchemaDefinitions](#json-schema-definitions)

### <span id="json-schema-dependencies"></span> JSONSchemaDependencies


  

[JSONSchemaDependencies](#json-schema-dependencies)

### <span id="json-schema-props"></span> JSONSchemaProps


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| $schema | [JSONSchemaURL](#json-schema-url)| `JSONSchemaURL` |  | |  |  |
| AllOf | [][JSONSchemaProps](#json-schema-props)| `[]*JSONSchemaProps` |  | |  |  |
| AnyOf | [][JSONSchemaProps](#json-schema-props)| `[]*JSONSchemaProps` |  | |  |  |
| Description | string| `string` |  | |  |  |
| Enum | [][JSON](#json)| `[]*JSON` |  | |  |  |
| ExclusiveMaximum | boolean| `bool` |  | |  |  |
| ExclusiveMinimum | boolean| `bool` |  | |  |  |
| Format | string| `string` |  | | format is an OpenAPI v3 format string. Unknown formats are ignored. The following formats are validated:

bsonobjectid: a bson object ID, i.e. a 24 characters hex string
uri: an URI as parsed by Golang net/url.ParseRequestURI
email: an email address as parsed by Golang net/mail.ParseAddress
hostname: a valid representation for an Internet host name, as defined by RFC 1034, section 3.1 [RFC1034].
ipv4: an IPv4 IP as parsed by Golang net.ParseIP
ipv6: an IPv6 IP as parsed by Golang net.ParseIP
cidr: a CIDR as parsed by Golang net.ParseCIDR
mac: a MAC address as parsed by Golang net.ParseMAC
uuid: an UUID that allows uppercase defined by the regex (?i)^[0-9a-f]{8}-?[0-9a-f]{4}-?[0-9a-f]{4}-?[0-9a-f]{4}-?[0-9a-f]{12}$
uuid3: an UUID3 that allows uppercase defined by the regex (?i)^[0-9a-f]{8}-?[0-9a-f]{4}-?3[0-9a-f]{3}-?[0-9a-f]{4}-?[0-9a-f]{12}$
uuid4: an UUID4 that allows uppercase defined by the regex (?i)^[0-9a-f]{8}-?[0-9a-f]{4}-?4[0-9a-f]{3}-?[89ab][0-9a-f]{3}-?[0-9a-f]{12}$
uuid5: an UUID5 that allows uppercase defined by the regex (?i)^[0-9a-f]{8}-?[0-9a-f]{4}-?5[0-9a-f]{3}-?[89ab][0-9a-f]{3}-?[0-9a-f]{12}$
isbn: an ISBN10 or ISBN13 number string like "0321751043" or "978-0321751041"
isbn10: an ISBN10 number string like "0321751043"
isbn13: an ISBN13 number string like "978-0321751041"
creditcard: a credit card number defined by the regex ^(?:4[0-9]{12}(?:[0-9]{3})?|5[1-5][0-9]{14}|6(?:011|5[0-9][0-9])[0-9]{12}|3[47][0-9]{13}|3(?:0[0-5]|[68][0-9])[0-9]{11}|(?:2131|1800|35\\d{3})\\d{11})$ with any non digit characters mixed in
ssn: a U.S. social security number following the regex ^\\d{3}[- ]?\\d{2}[- ]?\\d{4}$
hexcolor: an hexadecimal color code like "#FFFFFF: following the regex ^#?([0-9a-fA-F]{3}|[0-9a-fA-F]{6})$
rgbcolor: an RGB color code like rgb like "rgb(255,255,2559"
byte: base64 encoded binary data
password: any kind of string
date: a date string like "2006-01-02" as defined by full-date in RFC3339
duration: a duration string like "22 ns" as parsed by Golang time.ParseDuration or compatible with Scala duration format
datetime: a date time string like "2014-12-15T19:30:20.000Z" as defined by date-time in RFC3339. |  |
| ID | string| `string` |  | |  |  |
| MaxItems | int64 (formatted integer)| `int64` |  | |  |  |
| MaxLength | int64 (formatted integer)| `int64` |  | |  |  |
| MaxProperties | int64 (formatted integer)| `int64` |  | |  |  |
| Maximum | double (formatted number)| `float64` |  | |  |  |
| MinItems | int64 (formatted integer)| `int64` |  | |  |  |
| MinLength | int64 (formatted integer)| `int64` |  | |  |  |
| MinProperties | int64 (formatted integer)| `int64` |  | |  |  |
| Minimum | double (formatted number)| `float64` |  | |  |  |
| MultipleOf | double (formatted number)| `float64` |  | |  |  |
| Nullable | boolean| `bool` |  | |  |  |
| OneOf | [][JSONSchemaProps](#json-schema-props)| `[]*JSONSchemaProps` |  | |  |  |
| Pattern | string| `string` |  | |  |  |
| PatternProperties | map of [JSONSchemaProps](#json-schema-props)| `map[string]JSONSchemaProps` |  | |  |  |
| Properties | map of [JSONSchemaProps](#json-schema-props)| `map[string]JSONSchemaProps` |  | |  |  |
| Ref | string| `string` |  | |  |  |
| Required | []string| `[]string` |  | |  |  |
| Title | string| `string` |  | |  |  |
| Type | string| `string` |  | |  |  |
| UniqueItems | boolean| `bool` |  | |  |  |
| XEmbeddedResource | boolean| `bool` |  | | x-kubernetes-embedded-resource defines that the value is an
embedded Kubernetes runtime.Object, with TypeMeta and
ObjectMeta. The type must be object. It is allowed to further
restrict the embedded object. kind, apiVersion and metadata
are validated automatically. x-kubernetes-preserve-unknown-fields
is allowed to be true, but does not have to be if the object
is fully specified (up to kind, apiVersion, metadata). |  |
| XIntOrString | boolean| `bool` |  | | x-kubernetes-int-or-string specifies that this value is
either an integer or a string. If this is true, an empty
type is allowed and type as child of anyOf is permitted
if following one of the following patterns:

1) anyOf:
type: integer
type: string
2) allOf:
anyOf:
type: integer
type: string
... zero or more |  |
| XListMapKeys | []string| `[]string` |  | | x-kubernetes-list-map-keys annotates an array with the x-kubernetes-list-type `map` by specifying the keys used
as the index of the map.

This tag MUST only be used on lists that have the "x-kubernetes-list-type"
extension set to "map". Also, the values specified for this attribute must
be a scalar typed field of the child structure (no nesting is supported).

The properties specified must either be required or have a default value,
to ensure those properties are present for all list items.

+optional |  |
| XListType | string| `string` |  | | x-kubernetes-list-type annotates an array to further describe its topology.
This extension must only be used on lists and may have 3 possible values:

1) `atomic`: the list is treated as a single entity, like a scalar.
Atomic lists will be entirely replaced when updated. This extension
may be used on any type of list (struct, scalar, ...).
2) `set`:
Sets are lists that must not have multiple items with the same value. Each
value must be a scalar, an object with x-kubernetes-map-type `atomic` or an
array with x-kubernetes-list-type `atomic`.
3) `map`:
These lists are like maps in that their elements have a non-index key
used to identify them. Order is preserved upon merge. The map tag
must only be used on a list with elements of type object.
Defaults to atomic for arrays.
+optional |  |
| XMapType | string| `string` |  | | x-kubernetes-map-type annotates an object to further describe its topology.
This extension must only be used when type is object and may have 2 possible values:

1) `granular`:
These maps are actual maps (key-value pairs) and each fields are independent
from each other (they can each be manipulated by separate actors). This is
the default behaviour for all maps.
2) `atomic`: the list is treated as a single entity, like a scalar.
Atomic maps will be entirely replaced when updated.
+optional |  |
| XPreserveUnknownFields | boolean| `bool` |  | | x-kubernetes-preserve-unknown-fields stops the API server
decoding step from pruning fields which are not specified
in the validation schema. This affects fields recursively,
but switches back to normal pruning behaviour if nested
properties or additionalProperties are specified in the schema.
This can either be true or undefined. False is forbidden. |  |
| additionalItems | [JSONSchemaPropsOrBool](#json-schema-props-or-bool)| `JSONSchemaPropsOrBool` |  | |  |  |
| additionalProperties | [JSONSchemaPropsOrBool](#json-schema-props-or-bool)| `JSONSchemaPropsOrBool` |  | |  |  |
| default | [JSON](#json)| `JSON` |  | |  |  |
| definitions | [JSONSchemaDefinitions](#json-schema-definitions)| `JSONSchemaDefinitions` |  | |  |  |
| dependencies | [JSONSchemaDependencies](#json-schema-dependencies)| `JSONSchemaDependencies` |  | |  |  |
| example | [JSON](#json)| `JSON` |  | |  |  |
| externalDocs | [ExternalDocumentation](#external-documentation)| `ExternalDocumentation` |  | |  |  |
| items | [JSONSchemaPropsOrArray](#json-schema-props-or-array)| `JSONSchemaPropsOrArray` |  | |  |  |
| not | [JSONSchemaProps](#json-schema-props)| `JSONSchemaProps` |  | |  |  |



### <span id="json-schema-props-or-array"></span> JSONSchemaPropsOrArray


> JSONSchemaPropsOrArray represents a value that can either be a JSONSchemaProps
or an array of JSONSchemaProps. Mainly here for serialization purposes.
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| JSONSchemas | [][JSONSchemaProps](#json-schema-props)| `[]*JSONSchemaProps` |  | |  |  |
| Schema | [JSONSchemaProps](#json-schema-props)| `JSONSchemaProps` |  | |  |  |



### <span id="json-schema-props-or-bool"></span> JSONSchemaPropsOrBool


> Defaults to true for the boolean property.
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Allows | boolean| `bool` |  | |  |  |
| Schema | [JSONSchemaProps](#json-schema-props)| `JSONSchemaProps` |  | |  |  |



### <span id="json-schema-props-or-string-array"></span> JSONSchemaPropsOrStringArray


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Property | []string| `[]string` |  | |  |  |
| Schema | [JSONSchemaProps](#json-schema-props)| `JSONSchemaProps` |  | |  |  |



### <span id="json-schema-url"></span> JSONSchemaURL


  

| Name | Type | Go type | Default | Description | Example |
|------|------|---------| ------- |-------------|---------|
| JSONSchemaURL | string| string | |  |  |



### <span id="kind"></span> Kind


> Kind specifies the resource Kind and APIGroup
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| APIGroups | []string| `[]string` |  | | APIGroups specifies the APIGroups of the resources |  |
| Kinds | []string| `[]string` |  | | Kinds specifies the kinds of the resources |  |



### <span id="kubermatic-versions"></span> KubermaticVersions


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| API | string| `string` |  | | Version of the Kubermatic API server. |  |



### <span id="kubevirt-cloud-spec"></span> KubevirtCloudSpec


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Kubeconfig | string| `string` |  | |  |  |
| credentialsReference | [GlobalSecretKeySelector](#global-secret-key-selector)| `GlobalSecretKeySelector` |  | |  |  |



### <span id="kubevirt-node-spec"></span> KubevirtNodeSpec


> KubevirtNodeSpec kubevirt specific node settings
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| CPUs | string| `string` | ✓ | | CPUs states how many cpus the kubevirt node will have. |  |
| Memory | string| `string` | ✓ | | Memory states the memory that kubevirt node will have. |  |
| Namespace | string| `string` | ✓ | | Namespace states in which namespace kubevirt node will be provisioned. |  |
| PVCSize | string| `string` | ✓ | | PVCSize states the size of the provisioned pvc per node. |  |
| SourceURL | string| `string` | ✓ | | SourceURL states the url from which the imported image will be downloaded. |  |
| StorageClassName | string| `string` | ✓ | | StorageClassName states the storage class name for the provisioned PVCs. |  |



### <span id="l-b-s-k-u"></span> LBSKU


  

| Name | Type | Go type | Default | Description | Example |
|------|------|---------| ------- |-------------|---------|
| LBSKU | string| string | |  |  |



### <span id="label-key-list"></span> LabelKeyList


  

[]string

### <span id="label-selector"></span> LabelSelector


> A label selector is a label query over a set of resources. The result of matchLabels and
matchExpressions are ANDed. An empty label selector matches all objects. A null
label selector matches no objects.
+structType=atomic
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| MatchExpressions | [][LabelSelectorRequirement](#label-selector-requirement)| `[]*LabelSelectorRequirement` |  | | matchExpressions is a list of label selector requirements. The requirements are ANDed.
+optional |  |
| MatchLabels | map of string| `map[string]string` |  | | matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels
map is equivalent to an element of matchExpressions, whose key field is "key", the
operator is "In", and the values array contains only "value". The requirements are ANDed.
+optional |  |



### <span id="label-selector-operator"></span> LabelSelectorOperator


  

| Name | Type | Go type | Default | Description | Example |
|------|------|---------| ------- |-------------|---------|
| LabelSelectorOperator | string| string | |  |  |



### <span id="label-selector-requirement"></span> LabelSelectorRequirement


> A label selector requirement is a selector that contains values, a key, and an operator that
relates the key and values.
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Key | string| `string` |  | | key is the label key that the selector applies to.
+patchMergeKey=key
+patchStrategy=merge |  |
| Values | []string| `[]string` |  | | values is an array of string values. If the operator is In or NotIn,
the values array must be non-empty. If the operator is Exists or DoesNotExist,
the values array must be empty. This array is replaced during a strategic
merge patch.
+optional |  |
| operator | [LabelSelectorOperator](#label-selector-operator)| `LabelSelectorOperator` |  | |  |  |



### <span id="limits"></span> Limits


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| absolute | [Absolute](#absolute)| `Absolute` |  | |  |  |



### <span id="logging-rate-limit-settings"></span> LoggingRateLimitSettings


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| IngestionBurstSize | int32 (formatted integer)| `int32` |  | | IngestionBurstSize represents ingestion burst size in number of requests (nginx `burst`). |  |
| IngestionRate | int32 (formatted integer)| `int32` |  | | IngestionRate represents ingestion rate limit in requests per second (nginx `rate` in `r/s`). |  |
| QueryBurstSize | int32 (formatted integer)| `int32` |  | | QueryBurstSize represents query burst size in number of requests (nginx `burst`). |  |
| QueryRate | int32 (formatted integer)| `int32` |  | | QueryRate represents query request rate limit per second (nginx `rate` in `r/s`). |  |



### <span id="m-l-a"></span> MLA


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| UserClusterMLAEnabled | boolean| `bool` |  | | whether the user cluster MLA (Monitoring, Logging & Alerting) stack is enabled in the seed |  |



### <span id="m-l-a-admin-setting"></span> MLAAdminSetting


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| loggingRateLimits | [LoggingRateLimitSettings](#logging-rate-limit-settings)| `LoggingRateLimitSettings` |  | |  |  |
| monitoringRateLimits | [MonitoringRateLimitSettings](#monitoring-rate-limit-settings)| `MonitoringRateLimitSettings` |  | |  |  |



### <span id="m-l-a-settings"></span> MLASettings


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| LoggingEnabled | boolean| `bool` |  | | LoggingEnabled is the flag for enabling logging in user cluster. |  |
| MonitoringEnabled | boolean| `bool` |  | | MonitoringEnabled is the flag for enabling monitoring in user cluster. |  |
| loggingResources | [ResourceRequirements](#resource-requirements)| `ResourceRequirements` |  | |  |  |
| monitoringResources | [ResourceRequirements](#resource-requirements)| `ResourceRequirements` |  | |  |  |



### <span id="machine-deployment-status"></span> MachineDeploymentStatus


> [MachineDeploymentStatus]
MachineDeploymentStatus defines the observed state of MachineDeployment
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| AvailableReplicas | int32 (formatted integer)| `int32` |  | | Total number of available machines (ready for at least minReadySeconds)
targeted by this deployment.
+optional |  |
| ObservedGeneration | int64 (formatted integer)| `int64` |  | | The generation observed by the deployment controller.
+optional |  |
| ReadyReplicas | int32 (formatted integer)| `int32` |  | | Total number of ready machines targeted by this deployment.
+optional |  |
| Replicas | int32 (formatted integer)| `int32` |  | | Total number of non-terminated machines targeted by this deployment
(their labels match the selector).
+optional |  |
| UnavailableReplicas | int32 (formatted integer)| `int32` |  | | Total number of unavailable machines targeted by this deployment.
This is the total number of machines that are still required for
the deployment to have 100% available capacity. They may either
be machines that are running but not yet available or machines
that still have not been created.
+optional |  |
| UpdatedReplicas | int32 (formatted integer)| `int32` |  | | Total number of non-terminated machines targeted by this deployment
that have the desired template spec.
+optional |  |



### <span id="machine-deployment-vm-resource-quota"></span> MachineDeploymentVMResourceQuota


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| EnableGPU | boolean| `bool` |  | |  |  |
| MaxCPU | int64 (formatted integer)| `int64` |  | | Maximal number of vCPU |  |
| MaxRAM | int64 (formatted integer)| `int64` |  | | Maximum RAM size in GB |  |
| MinCPU | int64 (formatted integer)| `int64` |  | | Minimal number of vCPU |  |
| MinRAM | int64 (formatted integer)| `int64` |  | | Minimal RAM size in GB |  |



### <span id="machine-networking-config"></span> MachineNetworkingConfig


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| CIDR | string| `string` |  | |  |  |
| DNSServers | []string| `[]string` |  | |  |  |
| Gateway | string| `string` |  | |  |  |



### <span id="master-version"></span> MasterVersion


> MasterVersion describes a version of the master components
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Default | boolean| `bool` |  | |  |  |
| RestrictedByKubeletVersion | boolean| `bool` |  | | If true, then given version control plane version is not compatible
with one of the kubelets inside cluster and shouldn't be used. |  |
| version | [Version](#version)| `Version` |  | |  |  |



### <span id="match"></span> Match


> Match contains the constraint to resource matching data
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| ExcludedNamespaces | []string| `[]string` |  | | ExcludedNamespaces is a list of namespace names. If defined, a constraint will only apply to resources not in a listed namespace. |  |
| Kinds | [][Kind](#kind)| `[]*Kind` |  | | Kinds accepts a list of objects with apiGroups and kinds fields that list the groups/kinds of objects to which
the constraint will apply. If multiple groups/kinds objects are specified, only one match is needed for the resource to be in scope |  |
| Namespaces | []string| `[]string` |  | | Namespaces is a list of namespace names. If defined, a constraint will only apply to resources in a listed namespace. |  |
| Scope | string| `string` |  | | Scope accepts *, Cluster, or Namespaced which determines if cluster-scoped and/or namesapced-scoped resources are selected. (defaults to *) |  |
| labelSelector | [LabelSelector](#label-selector)| `LabelSelector` |  | |  |  |
| namespaceSelector | [LabelSelector](#label-selector)| `LabelSelector` |  | |  |  |



### <span id="match-entry"></span> MatchEntry


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| ExcludedNamespaces | []string| `[]string` |  | | Namespaces which will be excluded |  |
| Processes | []string| `[]string` |  | | Processes which will be excluded in the given namespaces (sync, webhook, audit, *) |  |



### <span id="metering-configurations"></span> MeteringConfigurations


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Enabled | boolean| `bool` |  | |  |  |
| StorageClassName | string| `string` |  | | StorageClassName is the name of the storage class that the metering tool uses to save processed files before
exporting it to s3 bucket. Default value is kubermatic-fast. |  |
| StorageSize | string| `string` |  | | StorageSize is the size of the storage class. Default value is 100Gi. |  |



### <span id="metering-report"></span> MeteringReport


> MeteringReport holds objects names and metadata for available reports
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| LastModified | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| Name | string| `string` |  | |  |  |
| Size | int64 (formatted integer)| `int64` |  | |  |  |



### <span id="mla-options"></span> MlaOptions


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| LoggingEnabled | boolean| `bool` |  | |  |  |
| LoggingEnforced | boolean| `bool` |  | |  |  |
| MonitoringEnabled | boolean| `bool` |  | |  |  |
| MonitoringEnforced | boolean| `bool` |  | |  |  |



### <span id="monitoring-rate-limit-settings"></span> MonitoringRateLimitSettings


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| IngestionBurstSize | int32 (formatted integer)| `int32` |  | | IngestionBurstSize represents ingestion burst size in samples per second (Cortex `ingestion_burst_size`). |  |
| IngestionRate | int32 (formatted integer)| `int32` |  | | IngestionRate represents the ingestion rate limit in samples per second (Cortex `ingestion_rate`). |  |
| MaxSamplesPerQuery | int32 (formatted integer)| `int32` |  | | MaxSamplesPerQuery represents maximum number of samples during a query (Cortex `max_samples_per_query`). |  |
| MaxSeriesPerMetric | int32 (formatted integer)| `int32` |  | | MaxSeriesPerMetric represents maximum number of series per metric (Cortex `max_series_per_metric`). |  |
| MaxSeriesPerQuery | int32 (formatted integer)| `int32` |  | | MaxSeriesPerQuery represents maximum number of timeseries during a query (Cortex `max_series_per_query`). |  |
| MaxSeriesTotal | int32 (formatted integer)| `int32` |  | | MaxSeriesTotal represents maximum number of series per this user cluster (Cortex `max_series_per_user`). |  |
| QueryBurstSize | int32 (formatted integer)| `int32` |  | | QueryBurstSize represents query burst size in number of requests (nginx `burst`). |  |
| QueryRate | int32 (formatted integer)| `int32` |  | | QueryRate represents  query request rate limit per second (nginx `rate` in `r/s`). |  |



### <span id="names"></span> Names


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Kind | string| `string` |  | |  |  |
| ShortNames | []string| `[]string` |  | |  |  |



### <span id="namespace"></span> Namespace


> Namespace defines namespace
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Name | string| `string` |  | |  |  |



### <span id="network-ranges"></span> NetworkRanges


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| CIDRBlocks | []string| `[]string` |  | |  |  |



### <span id="node"></span> Node


> Node represents a worker node that is part of a cluster
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Annotations | map of string| `map[string]string` |  | | Annotations that can be added to the resource |  |
| CreationTimestamp | date-time (formatted string)| `strfmt.DateTime` |  | | CreationTimestamp is a timestamp representing the server time when this object was created. |  |
| DeletionTimestamp | date-time (formatted string)| `strfmt.DateTime` |  | | DeletionTimestamp is a timestamp representing the server time when this object was deleted. |  |
| ID | string| `string` |  | | ID unique value that identifies the resource generated by the server. Read-Only. |  |
| Name | string| `string` |  | | Name represents human readable name for the resource |  |
| spec | [NodeSpec](#node-spec)| `NodeSpec` |  | |  |  |
| status | [NodeStatus](#node-status)| `NodeStatus` |  | |  |  |



### <span id="node-address"></span> NodeAddress


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Address | string| `string` |  | |  | `192.168.1.1, node1.my.dns` |
| Type | string| `string` |  | |  | `ExternalIP, InternalIP, InternalDNS, ExternalDNS` |



### <span id="node-cloud-spec"></span> NodeCloudSpec


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| alibaba | [AlibabaNodeSpec](#alibaba-node-spec)| `AlibabaNodeSpec` |  | |  |  |
| anexia | [AnexiaNodeSpec](#anexia-node-spec)| `AnexiaNodeSpec` |  | |  |  |
| aws | [AWSNodeSpec](#a-w-s-node-spec)| `AWSNodeSpec` |  | |  |  |
| azure | [AzureNodeSpec](#azure-node-spec)| `AzureNodeSpec` |  | |  |  |
| digitalocean | [DigitaloceanNodeSpec](#digitalocean-node-spec)| `DigitaloceanNodeSpec` |  | |  |  |
| gcp | [GCPNodeSpec](#g-c-p-node-spec)| `GCPNodeSpec` |  | |  |  |
| hetzner | [HetznerNodeSpec](#hetzner-node-spec)| `HetznerNodeSpec` |  | |  |  |
| kubevirt | [KubevirtNodeSpec](#kubevirt-node-spec)| `KubevirtNodeSpec` |  | |  |  |
| openstack | [OpenstackNodeSpec](#openstack-node-spec)| `OpenstackNodeSpec` |  | |  |  |
| packet | [PacketNodeSpec](#packet-node-spec)| `PacketNodeSpec` |  | |  |  |
| vsphere | [VSphereNodeSpec](#v-sphere-node-spec)| `VSphereNodeSpec` |  | |  |  |



### <span id="node-deployment"></span> NodeDeployment


> NodeDeployment represents a set of worker nodes that is part of a cluster
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Annotations | map of string| `map[string]string` |  | | Annotations that can be added to the resource |  |
| CreationTimestamp | date-time (formatted string)| `strfmt.DateTime` |  | | CreationTimestamp is a timestamp representing the server time when this object was created. |  |
| DeletionTimestamp | date-time (formatted string)| `strfmt.DateTime` |  | | DeletionTimestamp is a timestamp representing the server time when this object was deleted. |  |
| ID | string| `string` |  | | ID unique value that identifies the resource generated by the server. Read-Only. |  |
| Name | string| `string` |  | | Name represents human readable name for the resource |  |
| spec | [NodeDeploymentSpec](#node-deployment-spec)| `NodeDeploymentSpec` |  | |  |  |
| status | [MachineDeploymentStatus](#machine-deployment-status)| `MachineDeploymentStatus` |  | |  |  |



### <span id="node-deployment-request"></span> NodeDeploymentRequest


> NodeDeploymentRequest represents an asynchronous request to create a NodeDeployment in a user cluster
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Annotations | map of string| `map[string]string` |  | | Annotations that can be added to the resource |  |
| CreationTimestamp | date-time (formatted string)| `strfmt.DateTime` |  | | CreationTimestamp is a timestamp representing the server time when this object was created. |  |
| DeletionTimestamp | date-time (formatted string)| `strfmt.DateTime` |  | | DeletionTimestamp is a timestamp representing the server time when this object was deleted. |  |
| ID | string| `string` |  | | ID unique value that identifies the resource generated by the server. Read-Only. |  |
| Name | string| `string` |  | | Name represents human readable name for the resource |  |
| spec | [NodeDeploymentRequestSpec](#node-deployment-request-spec)| `NodeDeploymentRequestSpec` |  | |  |  |



### <span id="node-deployment-request-spec"></span> NodeDeploymentRequestSpec


> NodeDeploymentRequestSpec node deployment request specification
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| nd | [NodeDeployment](#node-deployment)| `NodeDeployment` |  | |  |  |



### <span id="node-deployment-spec"></span> NodeDeploymentSpec


> NodeDeploymentSpec node deployment specification
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| DynamicConfig | boolean| `bool` |  | |  |  |
| MaxReplicas | int32 (formatted integer)| `int32` |  | |  |  |
| MinReplicas | int32 (formatted integer)| `int32` |  | |  |  |
| Paused | boolean| `bool` |  | |  |  |
| Replicas | int32 (formatted integer)| `int32` | ✓ | |  |  |
| template | [NodeSpec](#node-spec)| `NodeSpec` | ✓ | |  |  |



### <span id="node-metric"></span> NodeMetric


> NodeMetric defines a metric for the given node
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| CPUAvailableMillicores | int64 (formatted integer)| `int64` |  | |  |  |
| CPUTotalMillicores | int64 (formatted integer)| `int64` |  | | CPUTotalMillicores in m cores |  |
| CPUUsedPercentage | int64 (formatted integer)| `int64` |  | | CPUUsedPercentage in percentage |  |
| MemoryAvailableBytes | int64 (formatted integer)| `int64` |  | | MemoryAvailableBytes available memory for node |  |
| MemoryTotalBytes | int64 (formatted integer)| `int64` |  | | MemoryTotalBytes current memory usage in bytes |  |
| MemoryUsedPercentage | int64 (formatted integer)| `int64` |  | | MemoryUsedPercentage in percentage |  |
| Name | string| `string` |  | |  |  |



### <span id="node-resources"></span> NodeResources


> NodeResources cpu and memory of a node
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| CPU | string| `string` |  | |  |  |
| Memory | string| `string` |  | |  |  |



### <span id="node-settings"></span> NodeSettings


> NodeSettings are node specific flags which can be configured on datacenter level
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| HyperkubeImage | string| `string` |  | | Optional: The hyperkube image to use. Currently only Flatcar
makes use of this option. |  |
| InsecureRegistries | []string| `[]string` |  | | Optional: These image registries will be configured as insecure
on the container runtime. |  |
| PauseImage | string| `string` |  | | Optional: Translates to --pod-infra-container-image on the kubelet.
If not set, the kubelet will default it. |  |
| RegistryMirrors | []string| `[]string` |  | | Optional: These image registries will be configured as registry mirrors
on the container runtime. |  |
| http_proxy | [ProxyValue](#proxy-value)| `ProxyValue` |  | |  |  |
| no_proxy | [ProxyValue](#proxy-value)| `ProxyValue` |  | |  |  |



### <span id="node-spec"></span> NodeSpec


> NodeSpec node specification
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Labels | map of string| `map[string]string` |  | | Map of string keys and values that can be used to organize and categorize (scope and select) objects.
It will be applied to Nodes allowing users run their apps on specific Node using labelSelector. |  |
| SSHUserName | string| `string` |  | |  |  |
| Taints | [][TaintSpec](#taint-spec)| `[]*TaintSpec` |  | | List of taints to set on new nodes |  |
| cloud | [NodeCloudSpec](#node-cloud-spec)| `NodeCloudSpec` | ✓ | |  |  |
| operatingSystem | [OperatingSystemSpec](#operating-system-spec)| `OperatingSystemSpec` | ✓ | |  |  |
| versions | [NodeVersionInfo](#node-version-info)| `NodeVersionInfo` | ✓ | |  |  |



### <span id="node-status"></span> NodeStatus


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Addresses | [][NodeAddress](#node-address)| `[]*NodeAddress` |  | | different addresses of a node |  |
| ErrorMessage | string| `string` |  | | in case of a error this will contain a detailed error explanation |  |
| ErrorReason | string| `string` |  | | in case of a error this will contain a short error message |  |
| MachineName | string| `string` |  | | name of the actual Machine object |  |
| allocatable | [NodeResources](#node-resources)| `NodeResources` |  | |  |  |
| capacity | [NodeResources](#node-resources)| `NodeResources` |  | |  |  |
| nodeInfo | [NodeSystemInfo](#node-system-info)| `NodeSystemInfo` |  | |  |  |



### <span id="node-system-info"></span> NodeSystemInfo


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Architecture | string| `string` |  | |  |  |
| ContainerRuntime | string| `string` |  | |  |  |
| ContainerRuntimeVersion | string| `string` |  | |  |  |
| KernelVersion | string| `string` |  | |  |  |
| KubeletVersion | string| `string` |  | |  |  |
| OperatingSystem | string| `string` |  | |  |  |



### <span id="node-version-info"></span> NodeVersionInfo


> NodeVersionInfo node version information
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Kubelet | string| `string` |  | |  |  |



### <span id="nodes-metric"></span> NodesMetric


> NodesMetric defines a metric for a group of nodes
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| CPUAvailableMillicores | int64 (formatted integer)| `int64` |  | |  |  |
| CPUTotalMillicores | int64 (formatted integer)| `int64` |  | | CPUTotalMillicores in m cores |  |
| CPUUsedPercentage | int64 (formatted integer)| `int64` |  | | CPUUsedPercentage in percentage |  |
| MemoryAvailableBytes | int64 (formatted integer)| `int64` |  | | MemoryAvailableBytes available memory for node |  |
| MemoryTotalBytes | int64 (formatted integer)| `int64` |  | | MemoryTotalBytes current memory usage in bytes |  |
| MemoryUsedPercentage | int64 (formatted integer)| `int64` |  | | MemoryUsedPercentage in percentage |  |



### <span id="o-id-c-settings"></span> OIDCSettings


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| ClientID | string| `string` |  | |  |  |
| ClientSecret | string| `string` |  | |  |  |
| ExtraScopes | string| `string` |  | |  |  |
| GroupsClaim | string| `string` |  | |  |  |
| IssuerURL | string| `string` |  | |  |  |
| RequiredClaim | string| `string` |  | |  |  |
| UsernameClaim | string| `string` |  | |  |  |



### <span id="o-id-c-spec"></span> OIDCSpec


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| ClientID | string| `string` |  | |  |  |
| ClientSecret | string| `string` |  | |  |  |
| IssuerURL | string| `string` |  | |  |  |



### <span id="o-p-a-integration-settings"></span> OPAIntegrationSettings


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Enabled | boolean| `bool` |  | | Enabled is the flag for enabling OPA integration |  |
| ExperimentalEnableMutation | boolean| `bool` |  | | Enable mutation |  |
| WebhookTimeoutSeconds | int32 (formatted integer)| `int32` |  | | WebhookTimeout is the timeout that is set for the gatekeeper validating webhook admission review calls.
By default 10 seconds. |  |



### <span id="object-meta"></span> ObjectMeta


> ObjectMeta defines the set of fields that objects returned from the API have
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Annotations | map of string| `map[string]string` |  | | Annotations that can be added to the resource |  |
| CreationTimestamp | date-time (formatted string)| `strfmt.DateTime` |  | | CreationTimestamp is a timestamp representing the server time when this object was created. |  |
| DeletionTimestamp | date-time (formatted string)| `strfmt.DateTime` |  | | DeletionTimestamp is a timestamp representing the server time when this object was deleted. |  |
| ID | string| `string` |  | | ID unique value that identifies the resource generated by the server. Read-Only. |  |
| Name | string| `string` |  | | Name represents human readable name for the resource |  |



### <span id="object-reference"></span> ObjectReference


> New uses of this type are discouraged because of difficulty describing its usage when embedded in APIs.
1. Ignored fields.  It includes many fields which are not generally honored.  For instance, ResourceVersion and FieldPath are both very rarely valid in actual usage.
2. Invalid usage help.  It is impossible to add specific help for individual usage.  In most embedded usages, there are particular
restrictions like, "must refer only to types A and B" or "UID not honored" or "name must be restricted".
Those cannot be well described when embedded.
3. Inconsistent validation.  Because the usages are different, the validation rules are different by usage, which makes it hard for users to predict what will happen.
4. The fields are both imprecise and overly precise.  Kind is not a precise mapping to a URL. This can produce ambiguity
during interpretation and require a REST mapping.  In most cases, the dependency is on the group,resource tuple
and the version of the actual struct is irrelevant.
5. We cannot easily change it.  Because this type is embedded in many locations, updates to this type
will affect numerous schemas.  Don't make new APIs embed an underspecified API type they do not control.
Instead of using this type, create a locally provided and used type that is well-focused on your reference.
For example, ServiceReferences for admission registration: https://github.com/kubernetes/api/blob/release-1.17/admissionregistration/v1/types.go#L533 .
+k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
+structType=atomic
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| APIVersion | string| `string` |  | | API version of the referent.
+optional |  |
| FieldPath | string| `string` |  | | If referring to a piece of an object instead of an entire object, this string
should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2].
For example, if the object reference is to a container within a pod, this would take on a value like:
"spec.containers{name}" (where "name" refers to the name of the container that triggered
the event) or if no container name is specified "spec.containers[2]" (container with
index 2 in this pod). This syntax is chosen only to have some well-defined way of
referencing a part of an object.
TODO: this design is not final and this field is subject to change in the future.
+optional |  |
| Kind | string| `string` |  | | Kind of the referent.
More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
+optional |  |
| Name | string| `string` |  | | Name of the referent.
More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
+optional |  |
| Namespace | string| `string` |  | | Namespace of the referent.
More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/
+optional |  |
| ResourceVersion | string| `string` |  | | Specific resourceVersion to which this reference is made, if any.
More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency
+optional |  |
| uid | [UID](#uid)| `UID` |  | |  |  |



### <span id="object-reference-resource"></span> ObjectReferenceResource


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Name | string| `string` |  | | Name of the referent.
More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
+optional |  |
| Namespace | string| `string` |  | | Namespace of the referent.
More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/
+optional |  |
| Type | string| `string` |  | | Type of the referent. |  |



### <span id="opa-options"></span> OpaOptions


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Enabled | boolean| `bool` |  | |  |  |
| Enforced | boolean| `bool` |  | |  |  |



### <span id="openstack-availability-zone"></span> OpenstackAvailabilityZone


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Name | string| `string` |  | | Name is the name of the availability zone |  |



### <span id="openstack-cloud-spec"></span> OpenstackCloudSpec


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| ApplicationCredentialID | string| `string` |  | |  |  |
| ApplicationCredentialSecret | string| `string` |  | |  |  |
| Domain | string| `string` |  | |  |  |
| FloatingIPPool | string| `string` |  | | FloatingIPPool holds the name of the public network
The public network is reachable from the outside world
and should provide the pool of IP addresses to choose from.

When specified, all worker nodes will receive a public ip from this floating ip pool

Note that the network is external if the "External" field is set to true |  |
| Network | string| `string` |  | | Network holds the name of the internal network
When specified, all worker nodes will be attached to this network. If not specified, a network, subnet & router will be created

Note that the network is internal if the "External" field is set to false |  |
| Password | string| `string` |  | |  |  |
| RouterID | string| `string` |  | |  |  |
| SecurityGroups | string| `string` |  | |  |  |
| ServerGroupID | string| `string` |  | | ServerGroupID used as schedule hint shared between all machines in the cluster,
When not specified, soft-anti-affinity server group will be automatically created |  |
| SubnetCIDR | string| `string` |  | |  |  |
| SubnetID | string| `string` |  | |  |  |
| Tenant | string| `string` |  | |  |  |
| TenantID | string| `string` |  | |  |  |
| Token | string| `string` |  | | Used internally during cluster creation |  |
| UseOctavia | boolean| `bool` |  | | Whether or not to use Octavia for LoadBalancer type of Service
implementation instead of using Neutron-LBaaS.
Attention:Openstack CCM use Octavia as default load balancer
implementation since v1.17.0

Takes precedence over the 'use_octavia' flag provided at datacenter
level if both are specified.
+optional |  |
| UseToken | boolean| `bool` |  | |  |  |
| Username | string| `string` |  | |  |  |
| credentialsReference | [GlobalSecretKeySelector](#global-secret-key-selector)| `GlobalSecretKeySelector` |  | |  |  |



### <span id="openstack-network"></span> OpenstackNetwork


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| External | boolean| `bool` |  | | External set if network is the external network |  |
| ID | string| `string` |  | | Id uniquely identifies the current network |  |
| Name | string| `string` |  | | Name is the name of the network |  |



### <span id="openstack-node-size-requirements"></span> OpenstackNodeSizeRequirements


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| MinimumMemory | int64 (formatted integer)| `int64` |  | | MinimumMemory is the minimum required amount of memory, measured in MB |  |
| MinimumVCPUs | int64 (formatted integer)| `int64` |  | | VCPUs is the minimum required amount of (virtual) CPUs |  |



### <span id="openstack-node-spec"></span> OpenstackNodeSpec


> OpenstackNodeSpec openstack node settings
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| AvailabilityZone | string| `string` |  | | if not set, the default AZ from the Datacenter spec will be used |  |
| Flavor | string| `string` | ✓ | | instance flavor |  |
| Image | string| `string` | ✓ | | image to use |  |
| InstanceReadyCheckPeriod | string| `string` |  | | Period of time to check for instance ready status, i.e. 10s/1m |  |
| InstanceReadyCheckTimeout | string| `string` |  | | Max time to wait for the instance to be ready, i.e. 10s/1m |  |
| RootDiskSizeGB | int64 (formatted integer)| `int64` |  | | if set, the rootDisk will be a volume. If not, the rootDisk will be on ephemeral storage and its size will be derived from the flavor |  |
| Tags | map of string| `map[string]string` |  | | Additional metadata to set |  |
| UseFloatingIP | boolean| `bool` |  | | Defines whether floating ip should be used |  |



### <span id="openstack-security-group"></span> OpenstackSecurityGroup


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| ID | string| `string` |  | | Id uniquely identifies the current security group |  |
| Name | string| `string` |  | | Name is the name of the security group |  |



### <span id="openstack-size"></span> OpenstackSize


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Disk | int64 (formatted integer)| `int64` |  | | Disk is the amount of root disk, measured in GB |  |
| IsPublic | boolean| `bool` |  | | IsPublic indicates whether the size is public (available to all projects) or scoped to a set of projects |  |
| Memory | int64 (formatted integer)| `int64` |  | | MemoryTotalBytes is the amount of memory, measured in MB |  |
| Region | string| `string` |  | | Region specifies the geographic region in which the size resides |  |
| Slug | string| `string` |  | | Slug holds  the name of the size |  |
| Swap | int64 (formatted integer)| `int64` |  | | Swap is the amount of swap space, measured in MB |  |
| VCPUs | int64 (formatted integer)| `int64` |  | | VCPUs indicates how many (virtual) CPUs are available for this flavor |  |



### <span id="openstack-subnet"></span> OpenstackSubnet


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| ID | string| `string` |  | | Id uniquely identifies the subnet |  |
| Name | string| `string` |  | | Name is human-readable name for the subnet |  |



### <span id="openstack-tenant"></span> OpenstackTenant


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| ID | string| `string` |  | | Id uniquely identifies the current tenant |  |
| Name | string| `string` |  | | Name is the name of the tenant |  |



### <span id="operating-system-spec"></span> OperatingSystemSpec


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| centos | [CentOSSpec](#cent-o-s-spec)| `CentOSSpec` |  | |  |  |
| flatcar | [FlatcarSpec](#flatcar-spec)| `FlatcarSpec` |  | |  |  |
| rhel | [RHELSpec](#r-h-e-l-spec)| `RHELSpec` |  | |  |  |
| sles | [SLESSpec](#s-l-e-s-spec)| `SLESSpec` |  | |  |  |
| ubuntu | [UbuntuSpec](#ubuntu-spec)| `UbuntuSpec` |  | |  |  |



### <span id="packet-cpu"></span> PacketCPU


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Count | int64 (formatted integer)| `int64` |  | |  |  |
| Type | string| `string` |  | |  |  |



### <span id="packet-cloud-spec"></span> PacketCloudSpec


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| APIKey | string| `string` |  | |  |  |
| BillingCycle | string| `string` |  | |  |  |
| ProjectID | string| `string` |  | |  |  |
| credentialsReference | [GlobalSecretKeySelector](#global-secret-key-selector)| `GlobalSecretKeySelector` |  | |  |  |



### <span id="packet-drive"></span> PacketDrive


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Count | int64 (formatted integer)| `int64` |  | |  |  |
| Size | string| `string` |  | |  |  |
| Type | string| `string` |  | |  |  |



### <span id="packet-node-spec"></span> PacketNodeSpec


> PacketNodeSpec specifies packet specific node settings
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| InstanceType | string| `string` | ✓ | | InstanceType denotes the plan to which the device will be provisioned. |  |
| Tags | []string| `[]string` |  | | additional instance tags |  |



### <span id="packet-size"></span> PacketSize


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| CPUs | [][PacketCPU](#packet-cpu)| `[]*PacketCPU` |  | |  |  |
| Drives | [][PacketDrive](#packet-drive)| `[]*PacketDrive` |  | |  |  |
| Memory | string| `string` |  | |  |  |
| Name | string| `string` |  | |  |  |



### <span id="packet-size-list"></span> PacketSizeList


  

[][PacketSize](#packet-size)

### <span id="parameters"></span> Parameters


  

[Parameters](#parameters)

### <span id="pod-dns-config"></span> PodDNSConfig


> PodDNSConfig defines the DNS parameters of a pod in addition to
those generated from DNSPolicy.
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Nameservers | []string| `[]string` |  | | A list of DNS name server IP addresses.
This will be appended to the base nameservers generated from DNSPolicy.
Duplicated nameservers will be removed.
+optional |  |
| Options | [][PodDNSConfigOption](#pod-dns-config-option)| `[]*PodDNSConfigOption` |  | | A list of DNS resolver options.
This will be merged with the base options generated from DNSPolicy.
Duplicated entries will be removed. Resolution options given in Options
will override those that appear in the base DNSPolicy.
+optional |  |
| Searches | []string| `[]string` |  | | A list of DNS search domains for host-name lookup.
This will be appended to the base search paths generated from DNSPolicy.
Duplicated search paths will be removed.
+optional |  |



### <span id="pod-dns-config-option"></span> PodDNSConfigOption


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Name | string| `string` |  | | Required. |  |
| Value | string| `string` |  | | +optional |  |



### <span id="policy-rule"></span> PolicyRule


> PolicyRule holds information that describes a policy rule, but does not contain information
about who the rule applies to or which namespace the rule applies to.
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| APIGroups | []string| `[]string` |  | | APIGroups is the name of the APIGroup that contains the resources.  If multiple API groups are specified, any action requested against one of
the enumerated resources in any API group will be allowed.
+optional |  |
| NonResourceURLs | []string| `[]string` |  | | NonResourceURLs is a set of partial urls that a user should have access to.  *s are allowed, but only as the full, final step in the path
Since non-resource URLs are not namespaced, this field is only applicable for ClusterRoles referenced from a ClusterRoleBinding.
Rules can either apply to API resources (such as "pods" or "secrets") or non-resource URL paths (such as "/api"),  but not both.
+optional |  |
| ResourceNames | []string| `[]string` |  | | ResourceNames is an optional white list of names that the rule applies to.  An empty set means that everything is allowed.
+optional |  |
| Resources | []string| `[]string` |  | | Resources is a list of resources this rule applies to. '*' represents all resources.
+optional |  |
| Verbs | []string| `[]string` |  | | Verbs is a list of Verbs that apply to ALL the ResourceKinds and AttributeRestrictions contained in this rule. '*' represents all verbs. |  |



### <span id="preset"></span> Preset


> Preset represents a preset
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Enabled | boolean| `bool` |  | |  |  |
| Name | string| `string` |  | |  |  |
| Providers | [][PresetProvider](#preset-provider)| `[]*PresetProvider` |  | |  |  |



### <span id="preset-list"></span> PresetList


> PresetList represents a list of presets
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Items | [][Preset](#preset)| `[]*Preset` |  | |  |  |



### <span id="preset-provider"></span> PresetProvider


> PresetProvider represents a preset provider
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Enabled | boolean| `bool` |  | |  |  |
| name | [ProviderType](#provider-type)| `ProviderType` |  | |  |  |



### <span id="project"></span> Project


> Project is a top-level container for a set of resources
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Annotations | map of string| `map[string]string` |  | | Annotations that can be added to the resource |  |
| ClustersNumber | int64 (formatted integer)| `int64` |  | |  |  |
| CreationTimestamp | date-time (formatted string)| `strfmt.DateTime` |  | | CreationTimestamp is a timestamp representing the server time when this object was created. |  |
| DeletionTimestamp | date-time (formatted string)| `strfmt.DateTime` |  | | DeletionTimestamp is a timestamp representing the server time when this object was deleted. |  |
| ID | string| `string` |  | | ID unique value that identifies the resource generated by the server. Read-Only. |  |
| Labels | map of string| `map[string]string` |  | |  |  |
| Name | string| `string` |  | | Name represents human readable name for the resource |  |
| Owners | [][User](#user)| `[]*User` |  | | Owners an optional owners list for the given project |  |
| Status | string| `string` |  | |  |  |



### <span id="project-group"></span> ProjectGroup


> ProjectGroup is a helper data structure that
stores the information about a project and a group prefix that a user belongs to
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| GroupPrefix | string| `string` |  | |  |  |
| ID | string| `string` |  | |  |  |



### <span id="provider-type"></span> ProviderType


  

| Name | Type | Go type | Default | Description | Example |
|------|------|---------| ------- |-------------|---------|
| ProviderType | string| string | |  |  |



### <span id="proxy-settings"></span> ProxySettings


> ProxySettings allow configuring a HTTP proxy for the controlplanes
and nodes
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| http_proxy | [ProxyValue](#proxy-value)| `ProxyValue` |  | |  |  |
| no_proxy | [ProxyValue](#proxy-value)| `ProxyValue` |  | |  |  |



### <span id="proxy-value"></span> ProxyValue


  

| Name | Type | Go type | Default | Description | Example |
|------|------|---------| ------- |-------------|---------|
| ProxyValue | string| string | |  |  |



### <span id="public-a-w-s-cloud-spec"></span> PublicAWSCloudSpec


  

[interface{}](#interface)

### <span id="public-alibaba-cloud-spec"></span> PublicAlibabaCloudSpec


  

[interface{}](#interface)

### <span id="public-anexia-cloud-spec"></span> PublicAnexiaCloudSpec


  

[interface{}](#interface)

### <span id="public-azure-cloud-spec"></span> PublicAzureCloudSpec


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| AssignAvailabilitySet | boolean| `bool` |  | |  |  |



### <span id="public-bring-your-own-cloud-spec"></span> PublicBringYourOwnCloudSpec


  

[interface{}](#interface)

### <span id="public-cloud-spec"></span> PublicCloudSpec


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| DatacenterName | string| `string` |  | |  |  |
| alibaba | [PublicAlibabaCloudSpec](#public-alibaba-cloud-spec)| `PublicAlibabaCloudSpec` |  | |  |  |
| anexia | [PublicAnexiaCloudSpec](#public-anexia-cloud-spec)| `PublicAnexiaCloudSpec` |  | |  |  |
| aws | [PublicAWSCloudSpec](#public-a-w-s-cloud-spec)| `PublicAWSCloudSpec` |  | |  |  |
| azure | [PublicAzureCloudSpec](#public-azure-cloud-spec)| `PublicAzureCloudSpec` |  | |  |  |
| bringyourown | [PublicBringYourOwnCloudSpec](#public-bring-your-own-cloud-spec)| `PublicBringYourOwnCloudSpec` |  | |  |  |
| digitalocean | [PublicDigitaloceanCloudSpec](#public-digitalocean-cloud-spec)| `PublicDigitaloceanCloudSpec` |  | |  |  |
| fake | [PublicFakeCloudSpec](#public-fake-cloud-spec)| `PublicFakeCloudSpec` |  | |  |  |
| gcp | [PublicGCPCloudSpec](#public-g-c-p-cloud-spec)| `PublicGCPCloudSpec` |  | |  |  |
| hetzner | [PublicHetznerCloudSpec](#public-hetzner-cloud-spec)| `PublicHetznerCloudSpec` |  | |  |  |
| kubevirt | [PublicKubevirtCloudSpec](#public-kubevirt-cloud-spec)| `PublicKubevirtCloudSpec` |  | |  |  |
| openstack | [PublicOpenstackCloudSpec](#public-openstack-cloud-spec)| `PublicOpenstackCloudSpec` |  | |  |  |
| packet | [PublicPacketCloudSpec](#public-packet-cloud-spec)| `PublicPacketCloudSpec` |  | |  |  |
| vsphere | [PublicVSphereCloudSpec](#public-v-sphere-cloud-spec)| `PublicVSphereCloudSpec` |  | |  |  |



### <span id="public-digitalocean-cloud-spec"></span> PublicDigitaloceanCloudSpec


  

[interface{}](#interface)

### <span id="public-fake-cloud-spec"></span> PublicFakeCloudSpec


  

[interface{}](#interface)

### <span id="public-g-c-p-cloud-spec"></span> PublicGCPCloudSpec


  

[interface{}](#interface)

### <span id="public-hetzner-cloud-spec"></span> PublicHetznerCloudSpec


  

[interface{}](#interface)

### <span id="public-kubevirt-cloud-spec"></span> PublicKubevirtCloudSpec


  

[interface{}](#interface)

### <span id="public-openstack-cloud-spec"></span> PublicOpenstackCloudSpec


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Domain | string| `string` |  | |  |  |
| FloatingIPPool | string| `string` |  | |  |  |
| Network | string| `string` |  | |  |  |
| RouterID | string| `string` |  | |  |  |
| SecurityGroups | string| `string` |  | |  |  |
| SubnetCIDR | string| `string` |  | |  |  |
| SubnetID | string| `string` |  | |  |  |
| Tenant | string| `string` |  | |  |  |
| TenantID | string| `string` |  | |  |  |



### <span id="public-packet-cloud-spec"></span> PublicPacketCloudSpec


  

[interface{}](#interface)

### <span id="public-service-account-token"></span> PublicServiceAccountToken


> PublicServiceAccountToken represent an API service account token without secret fields
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Annotations | map of string| `map[string]string` |  | | Annotations that can be added to the resource |  |
| CreationTimestamp | date-time (formatted string)| `strfmt.DateTime` |  | | CreationTimestamp is a timestamp representing the server time when this object was created. |  |
| DeletionTimestamp | date-time (formatted string)| `strfmt.DateTime` |  | | DeletionTimestamp is a timestamp representing the server time when this object was deleted. |  |
| Expiry | date-time (formatted string)| `strfmt.DateTime` |  | | Expiry is a timestamp representing the time when this token will expire. |  |
| ID | string| `string` |  | | ID unique value that identifies the resource generated by the server. Read-Only. |  |
| Name | string| `string` |  | | Name represents human readable name for the resource |  |



### <span id="public-v-sphere-cloud-spec"></span> PublicVSphereCloudSpec


  

[interface{}](#interface)

### <span id="quantity"></span> Quantity


> The serialization format is:

<quantity>        ::= <signedNumber><suffix>
(Note that <suffix> may be empty, from the "" case in <decimalSI>.)
<digit>           ::= 0 | 1 | ... | 9
<digits>          ::= <digit> | <digit><digits>
<number>          ::= <digits> | <digits>.<digits> | <digits>. | .<digits>
<sign>            ::= "+" | "-"
<signedNumber>    ::= <number> | <sign><number>
<suffix>          ::= <binarySI> | <decimalExponent> | <decimalSI>
<binarySI>        ::= Ki | Mi | Gi | Ti | Pi | Ei
(International System of units; See: http://physics.nist.gov/cuu/Units/binary.html)
<decimalSI>       ::= m | "" | k | M | G | T | P | E
(Note that 1024 = 1Ki but 1000 = 1k; I didn't choose the capitalization.)
<decimalExponent> ::= "e" <signedNumber> | "E" <signedNumber>

No matter which of the three exponent forms is used, no quantity may represent
a number greater than 2^63-1 in magnitude, nor may it have more than 3 decimal
places. Numbers larger or more precise will be capped or rounded up.
(E.g.: 0.1m will rounded up to 1m.)
This may be extended in the future if we require larger or smaller quantities.

When a Quantity is parsed from a string, it will remember the type of suffix
it had, and will use the same type again when it is serialized.

Before serializing, Quantity will be put in "canonical form".
This means that Exponent/suffix will be adjusted up or down (with a
corresponding increase or decrease in Mantissa) such that:
a. No precision is lost
b. No fractional digits will be emitted
c. The exponent (or suffix) is as large as possible.
The sign will be omitted unless the number is negative.

Examples:
1.5 will be serialized as "1500m"
1.5Gi will be serialized as "1536Mi"

Note that the quantity will NEVER be internally represented by a
floating point number. That is the whole point of this exercise.

Non-canonical values will still parse as long as they are well formed,
but will be re-emitted in their canonical form. (So always use canonical
form, or don't diff.)

This format is intended to make it difficult to use these numbers without
writing some sort of special handling code in the hopes that that will
cause implementors to also use a fixed point implementation.

+protobuf=true
+protobuf.embed=string
+protobuf.options.marshal=false
+protobuf.options.(gogoproto.goproto_stringer)=false
+k8s:deepcopy-gen=true
+k8s:openapi-gen=true
  



[interface{}](#interface)

### <span id="quotas"></span> Quotas


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| FloatingIPQuota | int64 (formatted integer)| `int64` |  | | FloatingIpQuota Sys11 addition with the amount of used and attached floating ips |  |
| UsedFloatingIPCount | int64 (formatted integer)| `int64` |  | | UsedFloatingIpCount is the floating IP quota |  |
| limits | [Limits](#limits)| `Limits` |  | |  |  |



### <span id="r-h-e-l-spec"></span> RHELSpec


> RHELSpec contains rhel specific settings
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| DistUpgradeOnBoot | boolean| `bool` |  | | do a dist-upgrade on boot and reboot it required afterwards |  |
| RHELSubscriptionManagerPassword | string| `string` |  | |  |  |
| RHELSubscriptionManagerUser | string| `string` |  | |  |  |
| RHSMOfflineToken | string| `string` |  | |  |  |



### <span id="readiness-spec"></span> ReadinessSpec


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| StatsEnabled | boolean| `bool` |  | | enables stats for gatekeeper audit |  |



### <span id="report-url"></span> ReportURL


> ReportURL represent an S3 pre signed URL to download a report
  



| Name | Type | Go type | Default | Description | Example |
|------|------|---------| ------- |-------------|---------|
| ReportURL | string| string | | ReportURL represent an S3 pre signed URL to download a report |  |



### <span id="resource-label-map"></span> ResourceLabelMap


  

[ResourceLabelMap](#resource-label-map)

### <span id="resource-list"></span> ResourceList


  

[ResourceList](#resource-list)

### <span id="resource-requirements"></span> ResourceRequirements


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| limits | [ResourceList](#resource-list)| `ResourceList` |  | |  |  |
| requests | [ResourceList](#resource-list)| `ResourceList` |  | |  |  |



### <span id="resource-type"></span> ResourceType


  

| Name | Type | Go type | Default | Description | Example |
|------|------|---------| ------- |-------------|---------|
| ResourceType | string| string | |  |  |



### <span id="role"></span> Role


> Role defines RBAC role for the user cluster
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Annotations | map of string| `map[string]string` |  | | Annotations that can be added to the resource |  |
| CreationTimestamp | date-time (formatted string)| `strfmt.DateTime` |  | | CreationTimestamp is a timestamp representing the server time when this object was created. |  |
| DeletionTimestamp | date-time (formatted string)| `strfmt.DateTime` |  | | DeletionTimestamp is a timestamp representing the server time when this object was deleted. |  |
| ID | string| `string` |  | | ID unique value that identifies the resource generated by the server. Read-Only. |  |
| Name | string| `string` |  | | Name represents human readable name for the resource |  |
| Namespace | string| `string` |  | | Indicates the scope of this role. |  |
| Rules | [][PolicyRule](#policy-rule)| `[]*PolicyRule` |  | | Rules holds all the PolicyRules for this Role |  |



### <span id="role-binding"></span> RoleBinding


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Namespace | string| `string` |  | | Indicates the scope of this binding. |  |
| RoleRefName | string| `string` |  | |  |  |
| Subjects | [][Subject](#subject)| `[]*Subject` |  | | Subjects holds references to the objects the role applies to. |  |



### <span id="role-name"></span> RoleName


> RoleName defines RBAC role name object for the user cluster
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Name | string| `string` |  | | Name of the role. |  |
| Namespace | []string| `[]string` |  | | Indicates the scopes of this role. |  |



### <span id="role-user"></span> RoleUser


> RoleUser defines associated user with role
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Group | string| `string` |  | |  |  |
| UserEmail | string| `string` |  | |  |  |



### <span id="rule-group"></span> RuleGroup


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Data | []uint8 (formatted integer)| `[]uint8` |  | | contains the RuleGroup data. Ref: https://prometheus.io/docs/prometheus/latest/configuration/recording_rules/#rule_group |  |
| type | [RuleGroupType](#rule-group-type)| `RuleGroupType` |  | |  |  |



### <span id="rule-group-type"></span> RuleGroupType


  

| Name | Type | Go type | Default | Description | Example |
|------|------|---------| ------- |-------------|---------|
| RuleGroupType | string| string | |  |  |



### <span id="s3-backup-credentials"></span> S3BackupCredentials


> S3BackupCredentials contains credentials for S3 etcd backups
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| AccessKeyID | string| `string` |  | |  |  |
| SecretAccessKey | string| `string` |  | |  |  |



### <span id="s-l-e-s-spec"></span> SLESSpec


> SLESSpec contains SLES specific settings
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| DistUpgradeOnBoot | boolean| `bool` |  | | do a dist-upgrade on boot and reboot it required afterwards |  |



### <span id="ssh-key"></span> SSHKey


> SSHKey represents a ssh key
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Annotations | map of string| `map[string]string` |  | | Annotations that can be added to the resource |  |
| CreationTimestamp | date-time (formatted string)| `strfmt.DateTime` |  | | CreationTimestamp is a timestamp representing the server time when this object was created. |  |
| DeletionTimestamp | date-time (formatted string)| `strfmt.DateTime` |  | | DeletionTimestamp is a timestamp representing the server time when this object was deleted. |  |
| ID | string| `string` |  | | ID unique value that identifies the resource generated by the server. Read-Only. |  |
| Name | string| `string` |  | | Name represents human readable name for the resource |  |
| spec | [SSHKeySpec](#ssh-key-spec)| `SSHKeySpec` |  | |  |  |



### <span id="ssh-key-spec"></span> SSHKeySpec


> SSHKeySpec represents the details of a ssh key
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Fingerprint | string| `string` |  | |  |  |
| PublicKey | string| `string` |  | |  |  |



### <span id="seed"></span> Seed


> Seed represents a seed object
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Country | string| `string` |  | | Optional: Country of the seed as ISO-3166 two-letter code, e.g. DE or UK.
For informational purposes in the Kubermatic dashboard only. |  |
| Location | string| `string` |  | | Optional: Detailed location of the cluster, like "Hamburg" or "Datacenter 7".
For informational purposes in the Kubermatic dashboard only. |  |
| Name | string| `string` |  | | Name represents human readable name for the resource |  |
| SeedDNSOverwrite | string| `string` |  | | Optional: This can be used to override the DNS name used for this seed.
By default the seed name is used. |  |
| SeedDatacenters | map of [Datacenter](#datacenter)| `map[string]Datacenter` |  | | Datacenters contains a map of the possible datacenters (DCs) in this seed.
Each DC must have a globally unique identifier (i.e. names must be unique
across all seeds). |  |
| backupRestore | [SeedBackupRestoreConfiguration](#seed-backup-restore-configuration)| `SeedBackupRestoreConfiguration` |  | |  |  |
| expose_strategy | [ExposeStrategy](#expose-strategy)| `ExposeStrategy` |  | |  |  |
| kubeconfig | [ObjectReference](#object-reference)| `ObjectReference` |  | |  |  |
| mla | [SeedMLASettings](#seed-m-l-a-settings)| `SeedMLASettings` |  | |  |  |
| proxy_settings | [ProxySettings](#proxy-settings)| `ProxySettings` |  | |  |  |



### <span id="seed-backup-restore-configuration"></span> SeedBackupRestoreConfiguration


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| S3BucketName | string| `string` |  | | S3BucketName is the S3 bucket name to use for backup and restore. |  |
| S3Endpoint | string| `string` |  | | S3Endpoint is the S3 API endpoint to use for backup and restore. Defaults to s3.amazonaws.com. |  |



### <span id="seed-m-l-a-settings"></span> SeedMLASettings


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| UserClusterMLAEnabled | boolean| `bool` |  | | Optional: UserClusterMLAEnabled controls whether the user cluster MLA (Monitoring, Logging & Alerting) stack is enabled in the seed. |  |



### <span id="seed-names-list"></span> SeedNamesList


  

[]string

### <span id="seed-settings"></span> SeedSettings


> SeedSettings represents settings for a Seed cluster
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| SeedDNSOverwrite | string| `string` |  | | the Seed level seed dns overwrite |  |
| metering | [MeteringConfigurations](#metering-configurations)| `MeteringConfigurations` |  | |  |  |
| mla | [MLA](#m-l-a)| `MLA` |  | |  |  |



### <span id="seed-spec"></span> SeedSpec


> The spec for a seed data
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Country | string| `string` |  | | Optional: Country of the seed as ISO-3166 two-letter code, e.g. DE or UK.
For informational purposes in the Kubermatic dashboard only. |  |
| Location | string| `string` |  | | Optional: Detailed location of the cluster, like "Hamburg" or "Datacenter 7".
For informational purposes in the Kubermatic dashboard only. |  |
| SeedDNSOverwrite | string| `string` |  | | Optional: This can be used to override the DNS name used for this seed.
By default the seed name is used. |  |
| SeedDatacenters | map of [Datacenter](#datacenter)| `map[string]Datacenter` |  | | Datacenters contains a map of the possible datacenters (DCs) in this seed.
Each DC must have a globally unique identifier (i.e. names must be unique
across all seeds). |  |
| backupRestore | [SeedBackupRestoreConfiguration](#seed-backup-restore-configuration)| `SeedBackupRestoreConfiguration` |  | |  |  |
| expose_strategy | [ExposeStrategy](#expose-strategy)| `ExposeStrategy` |  | |  |  |
| kubeconfig | [ObjectReference](#object-reference)| `ObjectReference` |  | |  |  |
| mla | [SeedMLASettings](#seed-m-l-a-settings)| `SeedMLASettings` |  | |  |  |
| proxy_settings | [ProxySettings](#proxy-settings)| `ProxySettings` |  | |  |  |



### <span id="semver"></span> Semver


> Semver is struct that encapsulates semver.Semver struct so we can use it in API
+k8s:deepcopy-gen=true
  



[interface{}](#interface)

### <span id="service-account"></span> ServiceAccount


> ServiceAccount represent an API service account
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Annotations | map of string| `map[string]string` |  | | Annotations that can be added to the resource |  |
| CreationTimestamp | date-time (formatted string)| `strfmt.DateTime` |  | | CreationTimestamp is a timestamp representing the server time when this object was created. |  |
| DeletionTimestamp | date-time (formatted string)| `strfmt.DateTime` |  | | DeletionTimestamp is a timestamp representing the server time when this object was deleted. |  |
| Group | string| `string` |  | | Group that a service account belongs to |  |
| ID | string| `string` |  | | ID unique value that identifies the resource generated by the server. Read-Only. |  |
| Name | string| `string` |  | | Name represents human readable name for the resource |  |
| Status | string| `string` |  | | Status describes three stages of ServiceAccount life including Active, Inactive and Terminating |  |



### <span id="service-account-settings"></span> ServiceAccountSettings


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| APIAudiences | []string| `[]string` |  | | APIAudiences are the Identifiers of the API
If this is not specified, it will be set to a single element list containing the issuer URL |  |
| Issuer | string| `string` |  | | Issuer is the identifier of the service account token issuer
If this is not specified, it will be set to the URL of apiserver by default |  |
| TokenVolumeProjectionEnabled | boolean| `bool` |  | |  |  |



### <span id="service-account-token"></span> ServiceAccountToken


> ServiceAccountToken represent an API service account token
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Annotations | map of string| `map[string]string` |  | | Annotations that can be added to the resource |  |
| CreationTimestamp | date-time (formatted string)| `strfmt.DateTime` |  | | CreationTimestamp is a timestamp representing the server time when this object was created. |  |
| DeletionTimestamp | date-time (formatted string)| `strfmt.DateTime` |  | | DeletionTimestamp is a timestamp representing the server time when this object was deleted. |  |
| Expiry | date-time (formatted string)| `strfmt.DateTime` |  | | Expiry is a timestamp representing the time when this token will expire. |  |
| ID | string| `string` |  | | ID unique value that identifies the resource generated by the server. Read-Only. |  |
| Name | string| `string` |  | | Name represents human readable name for the resource |  |
| Token | string| `string` |  | | Token the JWT token |  |



### <span id="setting-spec"></span> SettingSpec


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| DefaultNodeCount | int8 (formatted integer)| `int8` |  | |  |  |
| DisplayAPIDocs | boolean| `bool` |  | |  |  |
| DisplayDemoInfo | boolean| `bool` |  | |  |  |
| DisplayTermsOfService | boolean| `bool` |  | |  |  |
| EnableDashboard | boolean| `bool` |  | |  |  |
| EnableExternalClusterImport | boolean| `bool` |  | |  |  |
| EnableOIDCKubeconfig | boolean| `bool` |  | |  |  |
| MlaAlertmanagerPrefix | string| `string` |  | |  |  |
| MlaGrafanaPrefix | string| `string` |  | |  |  |
| RestrictProjectCreation | boolean| `bool` |  | |  |  |
| UserProjectsLimit | int64 (formatted integer)| `int64` |  | |  |  |
| cleanupOptions | [CleanupOptions](#cleanup-options)| `CleanupOptions` |  | |  |  |
| clusterTypeOptions | [ClusterType](#cluster-type)| `ClusterType` |  | |  |  |
| customLinks | [CustomLinks](#custom-links)| `CustomLinks` |  | |  |  |
| machineDeploymentVMResourceQuota | [MachineDeploymentVMResourceQuota](#machine-deployment-vm-resource-quota)| `MachineDeploymentVMResourceQuota` |  | |  |  |
| mlaOptions | [MlaOptions](#mla-options)| `MlaOptions` |  | |  |  |
| opaOptions | [OpaOptions](#opa-options)| `OpaOptions` |  | |  |  |



### <span id="subject"></span> Subject


> or a value for non-objects such as user and group names.
+structType=atomic
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| APIGroup | string| `string` |  | | APIGroup holds the API group of the referenced subject.
Defaults to "" for ServiceAccount subjects.
Defaults to "rbac.authorization.k8s.io" for User and Group subjects.
+optional |  |
| Kind | string| `string` |  | | Kind of object being referenced. Values defined by this API group are "User", "Group", and "ServiceAccount".
If the Authorizer does not recognized the kind value, the Authorizer should report an error. |  |
| Name | string| `string` |  | | Name of the object being referenced. |  |
| Namespace | string| `string` |  | | Namespace of the referenced object.  If the object kind is non-namespace, such as "User" or "Group", and this value is not empty
the Authorizer should report an error.
+optional |  |



### <span id="sync"></span> Sync


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| SyncOnly | [][GVK](#g-v-k)| `[]*GVK` |  | | If non-empty, entries on this list will be replicated into OPA |  |



### <span id="sys11-auth-settings"></span> Sys11AuthSettings


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Realm | string| `string` |  | |  |  |



### <span id="taint-spec"></span> TaintSpec


> TaintSpec defines a node taint
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Effect | string| `string` |  | |  |  |
| Key | string| `string` |  | |  |  |
| Value | string| `string` |  | |  |  |



### <span id="target"></span> Target


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Libs | []string| `[]string` |  | |  |  |
| Rego | string| `string` |  | |  |  |
| Target | string| `string` |  | |  |  |



### <span id="time"></span> Time


> Programs using times should typically store and pass them as values,
not pointers. That is, time variables and struct fields should be of
type time.Time, not *time.Time.

A Time value can be used by multiple goroutines simultaneously except
that the methods GobDecode, UnmarshalBinary, UnmarshalJSON and
UnmarshalText are not concurrency-safe.

Time instants can be compared using the Before, After, and Equal methods.
The Sub method subtracts two instants, producing a Duration.
The Add method adds a Time and a Duration, producing a Time.

The zero value of type Time is January 1, year 1, 00:00:00.000000000 UTC.
As this time is unlikely to come up in practice, the IsZero method gives
a simple way of detecting a time that has not been initialized explicitly.

Each Time has associated with it a Location, consulted when computing the
presentation form of the time, such as in the Format, Hour, and Year methods.
The methods Local, UTC, and In return a Time with a specific location.
Changing the location in this way changes only the presentation; it does not
change the instant in time being denoted and therefore does not affect the
computations described in earlier paragraphs.

Representations of a Time value saved by the GobEncode, MarshalBinary,
MarshalJSON, and MarshalText methods store the Time.Location's offset, but not
the location name. They therefore lose information about Daylight Saving Time.

In addition to the required “wall clock” reading, a Time may contain an optional
reading of the current process's monotonic clock, to provide additional precision
for comparison or subtraction.
See the “Monotonic Clocks” section in the package documentation for details.

Note that the Go == operator compares not just the time instant but also the
Location and the monotonic clock reading. Therefore, Time values should not
be used as map or database keys without first guaranteeing that the
identical Location has been set for all values, which can be achieved
through use of the UTC or Local method, and that the monotonic clock reading
has been stripped by setting t = t.Round(0). In general, prefer t.Equal(u)
to t == u, since t.Equal uses the most accurate comparison available and
correctly handles the case when only one of its arguments has a monotonic
clock reading.
  



| Name | Type | Go type | Default | Description | Example |
|------|------|---------| ------- |-------------|---------|
| Time | date-time (formatted string)| strfmt.DateTime | | Programs using times should typically store and pass them as values,
not pointers. That is, time variables and struct fields should be of
type time.Time, not *time.Time.

A Time value can be used by multiple goroutines simultaneously except
that the methods GobDecode, UnmarshalBinary, UnmarshalJSON and
UnmarshalText are not concurrency-safe.

Time instants can be compared using the Before, After, and Equal methods.
The Sub method subtracts two instants, producing a Duration.
The Add method adds a Time and a Duration, producing a Time.

The zero value of type Time is January 1, year 1, 00:00:00.000000000 UTC.
As this time is unlikely to come up in practice, the IsZero method gives
a simple way of detecting a time that has not been initialized explicitly.

Each Time has associated with it a Location, consulted when computing the
presentation form of the time, such as in the Format, Hour, and Year methods.
The methods Local, UTC, and In return a Time with a specific location.
Changing the location in this way changes only the presentation; it does not
change the instant in time being denoted and therefore does not affect the
computations described in earlier paragraphs.

Representations of a Time value saved by the GobEncode, MarshalBinary,
MarshalJSON, and MarshalText methods store the Time.Location's offset, but not
the location name. They therefore lose information about Daylight Saving Time.

In addition to the required “wall clock” reading, a Time may contain an optional
reading of the current process's monotonic clock, to provide additional precision
for comparison or subtraction.
See the “Monotonic Clocks” section in the package documentation for details.

Note that the Go == operator compares not just the time instant but also the
Location and the monotonic clock reading. Therefore, Time values should not
be used as map or database keys without first guaranteeing that the
identical Location has been set for all values, which can be achieved
through use of the UTC or Local method, and that the monotonic clock reading
has been stripped by setting t = t.Round(0). In general, prefer t.Equal(u)
to t == u, since t.Equal uses the most accurate comparison available and
correctly handles the case when only one of its arguments has a monotonic
clock reading. |  |



### <span id="trace"></span> Trace


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Dump | string| `string` |  | | Also dump the state of OPA with the trace. Set to `All` to dump everything. |  |
| User | string| `string` |  | | Only trace requests from the specified user |  |
| kind | [GVK](#g-v-k)| `GVK` |  | |  |  |



### <span id="uid"></span> UID


> UID is a type that holds unique ID values, including UUIDs.  Because we
don't ONLY use UUIDs, this is an alias to string.  Being a type captures
intent and helps make sure that UIDs and names do not get conflated.
  



| Name | Type | Go type | Default | Description | Example |
|------|------|---------| ------- |-------------|---------|
| UID | string| string | | UID is a type that holds unique ID values, including UUIDs.  Because we
don't ONLY use UUIDs, this is an alias to string.  Being a type captures
intent and helps make sure that UIDs and names do not get conflated. |  |



### <span id="ubuntu-spec"></span> UbuntuSpec


> UbuntuSpec ubuntu specific settings
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| DistUpgradeOnBoot | boolean| `bool` |  | | do a dist-upgrade on boot and reboot it required afterwards |  |



### <span id="update-window"></span> UpdateWindow


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Length | string| `string` |  | |  |  |
| Start | string| `string` |  | |  |  |



### <span id="user"></span> User


> User represent an API user
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Annotations | map of string| `map[string]string` |  | | Annotations that can be added to the resource |  |
| CreationTimestamp | date-time (formatted string)| `strfmt.DateTime` |  | | CreationTimestamp is a timestamp representing the server time when this object was created. |  |
| DeletionTimestamp | date-time (formatted string)| `strfmt.DateTime` |  | | DeletionTimestamp is a timestamp representing the server time when this object was deleted. |  |
| Email | string| `string` |  | | Email an email address of the user |  |
| ID | string| `string` |  | | ID unique value that identifies the resource generated by the server. Read-Only. |  |
| IsAdmin | boolean| `bool` |  | | IsAdmin indicates admin role |  |
| Name | string| `string` |  | | Name represents human readable name for the resource |  |
| Projects | [][ProjectGroup](#project-group)| `[]*ProjectGroup` |  | | Projects holds the list of project the user belongs to
along with the group names |  |
| userSettings | [UserSettings](#user-settings)| `UserSettings` |  | |  |  |



### <span id="user-settings"></span> UserSettings


> UserSettings represent an user settings
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| CollapseSidenav | boolean| `bool` |  | |  |  |
| DisplayAllProjectsForAdmin | boolean| `bool` |  | |  |  |
| ItemsPerPage | int8 (formatted integer)| `int8` |  | |  |  |
| LastSeenChangelogVersion | string| `string` |  | |  |  |
| SelectProjectTableView | boolean| `bool` |  | |  |  |
| SelectedProjectID | string| `string` |  | |  |  |
| SelectedTheme | string| `string` |  | |  |  |



### <span id="v-sphere-cloud-spec"></span> VSphereCloudSpec


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Datastore | string| `string` |  | | Datastore to be used for storing virtual machines and as a default for
dynamic volume provisioning, it is mutually exclusive with
DatastoreCluster.
+optional |  |
| DatastoreCluster | string| `string` |  | | DatastoreCluster to be used for storing virtual machines, it is mutually
exclusive with Datastore.
+optional |  |
| Folder | string| `string` |  | | Folder is the folder to be used to group the provisioned virtual
machines.
+optional |  |
| Password | string| `string` |  | | Password is the vSphere user password.
+optional |  |
| ResourcePool | string| `string` |  | | ResourcePool is used to manage resources such as cpu and memory for vSphere virtual machines. The resource pool
should be defined on vSphere cluster level.
+optional |  |
| StoragePolicy | string| `string` |  | | StoragePolicy to be used for storage provisioning |  |
| Username | string| `string` |  | | Username is the vSphere user name.
+optional |  |
| VMNetName | string| `string` |  | | VMNetName is the name of the vSphere network. |  |
| credentialsReference | [GlobalSecretKeySelector](#global-secret-key-selector)| `GlobalSecretKeySelector` |  | |  |  |
| infraManagementUser | [VSphereCredentials](#v-sphere-credentials)| `VSphereCredentials` |  | |  |  |



### <span id="v-sphere-credentials"></span> VSphereCredentials


> VSphereCredentials credentials represents a credential for accessing vSphere
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Password | string| `string` |  | |  |  |
| Username | string| `string` |  | |  |  |



### <span id="v-sphere-datastore-list"></span> VSphereDatastoreList


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Datastores | []string| `[]string` |  | |  |  |



### <span id="v-sphere-folder"></span> VSphereFolder


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Path | string| `string` |  | | Path is the path of the folder |  |



### <span id="v-sphere-network"></span> VSphereNetwork


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| AbsolutePath | string| `string` |  | | AbsolutePath is the absolute path inside vCenter |  |
| Name | string| `string` |  | | Name is the name of the network |  |
| RelativePath | string| `string` |  | | RelativePath is the relative path inside the datacenter |  |
| Type | string| `string` |  | | Type defines the type of network |  |



### <span id="v-sphere-node-spec"></span> VSphereNodeSpec


> VSphereNodeSpec VSphere node settings
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| CPUs | int64 (formatted integer)| `int64` |  | |  |  |
| DiskSizeGB | int64 (formatted integer)| `int64` |  | |  |  |
| Memory | int64 (formatted integer)| `int64` |  | |  |  |
| Template | string| `string` |  | |  |  |



### <span id="validation"></span> Validation


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| LegacySchema | boolean| `bool` |  | | +kubebuilder:default=true |  |
| openAPIV3Schema | [JSONSchemaProps](#json-schema-props)| `JSONSchemaProps` |  | |  |  |



### <span id="version"></span> Version


  

[interface{}](#interface)

### <span id="version-list"></span> VersionList


> VersionList represents a list of versions
  



[][MasterVersion](#master-version)

### <span id="violation"></span> Violation


> Violation represents a gatekeeper constraint violation
  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| EnforcementAction | string| `string` |  | |  |  |
| Kind | string| `string` |  | |  |  |
| Message | string| `string` |  | |  |  |
| Name | string| `string` |  | |  |  |
| Namespace | string| `string` |  | |  |  |



### <span id="bc-body"></span> bcBody


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| backup_credentials | [BackupCredentials](#backup-credentials)| `BackupCredentials` |  | |  |  |



### <span id="body"></span> body


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Kubeconfig | string| `string` |  | | Kubeconfig Base64 encoded kubeconfig |  |
| Name | string| `string` |  | | Name is human readable name for the external cluster |  |



### <span id="constraint-body"></span> constraintBody


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Name | string| `string` |  | | Name is the name for the constraint |  |
| Spec | [ConstraintSpec](#constraint-spec)| `ConstraintSpec` |  | |  |  |



### <span id="ct-body"></span> ctBody


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Name | string| `string` |  | | Name of the constraint template |  |
| spec | [ConstraintTemplateSpec](#constraint-template-spec)| `ConstraintTemplateSpec` |  | |  |  |



### <span id="ebc-body"></span> ebcBody


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Name | string| `string` |  | | Name of the etcd backup config |  |
| spec | [EtcdBackupConfigSpec](#etcd-backup-config-spec)| `EtcdBackupConfigSpec` |  | |  |  |



### <span id="er-body"></span> erBody


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Name | string| `string` |  | | Name of the etcd backup restore. If not set, it will be generated |  |
| spec | [EtcdRestoreSpec](#etcd-restore-spec)| `EtcdRestoreSpec` |  | |  |  |



### <span id="unseal-keys"></span> unsealKeys


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Keys | []string| `[]string` |  | |  |  |



### <span id="wr-body"></span> wrBody


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Name | string| `string` |  | | Name of the allowed registry |  |
| spec | [AllowedRegistrySpec](#allowed-registry-spec)| `AllowedRegistrySpec` |  | |  |  |


