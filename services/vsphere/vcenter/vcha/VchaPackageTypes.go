/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for package: com.vmware.vcenter.vcha.
 * Includes binding types of a top level structures and enumerations.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package vcha

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
)


// The ``IpFamily`` enumeration class defines the Ip address family. This enumeration was added in vSphere API 6.7.1.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type IpFamily string

const (
    // IPV4 address family. This constant field was added in vSphere API 6.7.1.
	IpFamily_IPV4 IpFamily = "IPV4"
    // IPv6 address family. This constant field was added in vSphere API 6.7.1.
	IpFamily_IPV6 IpFamily = "IPV6"
)

func (i IpFamily) IpFamily() bool {
	switch i {
	case IpFamily_IPV4:
		return true
	case IpFamily_IPV6:
		return true
	default:
		return false
	}
}


// The ``NetworkType`` enumeration class defines the type of a vCenter Server network. This enumeration was added in vSphere API 6.7.1.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type NetworkType string

const (
    // vSphere standard portgroup network. This constant field was added in vSphere API 6.7.1.
	NetworkType_STANDARD_PORTGROUP NetworkType = "STANDARD_PORTGROUP"
    // Distributed virtual switch. This constant field was added in vSphere API 6.7.1.
	NetworkType_DISTRIBUTED_PORTGROUP NetworkType = "DISTRIBUTED_PORTGROUP"
)

func (n NetworkType) NetworkType() bool {
	switch n {
	case NetworkType_STANDARD_PORTGROUP:
		return true
	case NetworkType_DISTRIBUTED_PORTGROUP:
		return true
	default:
		return false
	}
}


// The ``CertificateInfo`` Class contains information about the SSL certificate for a management vCenter server. This class was added in vSphere API 6.7.1.
type CertificateInfo struct {
    // The SHA-256 thumbprint of the SSL certificate for a management vCenter server. This property was added in vSphere API 6.7.1.
	SslThumbprint string
}

// The ``ConnectionSpec`` class contains information required to connect to a vCenter server. The connection to the vCenter server always uses the HTTPS protocol. This class was added in vSphere API 6.7.1.
type ConnectionSpec struct {
    // IP Address or DNS of the vCenter. This property was added in vSphere API 6.7.1.
	Hostname string
    // Port number. This property was added in vSphere API 6.7.1.
	Port *int64
    // SHA1 hash of the server SSL certificate. This property was added in vSphere API 6.7.1.
	SslThumbprint *string
    // Username to access the server. This property was added in vSphere API 6.7.1.
	Username *string
    // Password for the specified user. This property was added in vSphere API 6.7.1.
	Password *string
}

// The ``CredentialsSpec`` class contains information to connect to the vCenter server managing the VCHA nodes. This class was added in vSphere API 6.7.1.
type CredentialsSpec struct {
    // Connection information for the management vCenter Server of the Active Node in a VCHA Cluster. This property was added in vSphere API 6.7.1.
	ActiveLocation ConnectionSpec
}

// The ``DiskInfo`` class contains information to describe the storage configuration of a vCenter virtual machine. This class was added in vSphere API 6.7.1.
type DiskInfo struct {
    // The identifier of the datastore to put all the virtual disks on. This property was added in vSphere API 6.7.1.
	Datastore string
    // The name of the datastore. This property was added in vSphere API 6.7.1.
	DatastoreName string
}

// The ``DiskSpec`` class contains information to describe the storage configuration of a vCenter virtual machine. This class was added in vSphere API 6.7.1.
type DiskSpec struct {
    // The identifier of the datastore to put all the virtual disks on. This property was added in vSphere API 6.7.1.
	Datastore *string
}

// The ``IpSpec`` class contains IP information used to configure a network interface. This class was added in vSphere API 6.7.1.
type IpSpec struct {
    // Family of the IP address to configure the interface. This property was added in vSphere API 6.7.1.
	IpFamily IpFamily
    // If the family of the ip is IPV4, then this will point to IPv4 address specification. This property was added in vSphere API 6.7.1.
	Ipv4 *Ipv4Spec
    // If the family of the ip is IPV6, then this will point to IPv6 address specification. This property was added in vSphere API 6.7.1.
	Ipv6 *Ipv6Spec
    // The IP address of the Gateway for this interface. This property was added in vSphere API 6.7.1.
	DefaultGateway *string
    // The list of IP addresses of the DNS servers for this interface. This list is a comma separated list. This property was added in vSphere API 6.7.1.
	DnsServers []string
}

// The ``Ipv4Spec`` class contains IPV4 information used to configure a network interface. This class was added in vSphere API 6.7.1.
type Ipv4Spec struct {
    // IPV4 address to be used to configure the interface. This property was added in vSphere API 6.7.1.
	Address string
    // The subnet mask for the interface. This property was added in vSphere API 6.7.1.
	SubnetMask *string
    // The CIDR prefix for the interface. This property was added in vSphere API 6.7.1.
	Prefix *int64
}

// The ``Ipv6Spec`` class contains IPV6 information used to configure a network interface. This class was added in vSphere API 6.7.1.
type Ipv6Spec struct {
    // IPv6 address to be used to configure the interface. This property was added in vSphere API 6.7.1.
	Address string
    // The CIDR prefix for the interface. This property was added in vSphere API 6.7.1.
	Prefix int64
}

// The ``PlacementInfo`` class contains information to describe the inventory placement of a single node of a VCHA cluster.
//  The active node's management vCenter server credentials are required to populate all properties except biosUuid. This class was added in vSphere API 6.7.1.
type PlacementInfo struct {
    // The hostname of the vCenter server that is managing the VCHA node. This property was added in vSphere API 6.7.1.
	ManagementVcenterName string
    // The virtual machine name of the VCHA node. This property was added in vSphere API 6.7.1.
	VmName string
    // The identifier of the datacenter of the VCHA node. This property was added in vSphere API 6.7.1.
	Datacenter string
    // The name of the datacenter of the VCHA node. This property was added in vSphere API 6.7.1.
	DatacenterName string
    // The identifier of the host of the VCHA node. This property was added in vSphere API 6.7.1.
	Host string
    // The name of the host of the VCHA node. This property was added in vSphere API 6.7.1.
	HostName string
    // The identifier of the cluster of which ``host`` is member. This property was added in vSphere API 6.7.1.
	Cluster *string
    // The name of the cluster of which ``host`` is member. This property was added in vSphere API 6.7.1.
	ClusterName *string
    // The identifier of the Network object used for the HA network. This property was added in vSphere API 6.7.1.
	HaNetwork *string
    // The name of the Network object used for the HA network. This property was added in vSphere API 6.7.1.
	HaNetworkName *string
    // The type of the Network object used for the HA network. This property was added in vSphere API 6.7.1.
	HaNetworkType *NetworkType
    // The identifier of the Network object used for the Management network. This property was added in vSphere API 6.7.1.
	ManagementNetwork string
    // The name of the Network object used for the Management network. This property was added in vSphere API 6.7.1.
	ManagementNetworkName string
    // The type of the Network object used for the Management network. This property was added in vSphere API 6.7.1.
	ManagementNetworkType NetworkType
    // The storage information of the VCHA node. This property was added in vSphere API 6.7.1.
	Storage DiskInfo
    // BIOS UUID for the node. This property was added in vSphere API 6.7.1.
	BiosUuid *string
}

// The ``PlacementSpec`` class contains information to describe the inventory placement of a single node of a VCHA cluster. This class was added in vSphere API 6.7.1.
type PlacementSpec struct {
    // The name of the VCHA node to be used for the virtual machine name. This property was added in vSphere API 6.7.1.
	Name string
    // The identifier of the folder to deploy the VCHA node to. This property was added in vSphere API 6.7.1.
	Folder string
    // The identifier of the host to deploy the VCHA node to. This property was added in vSphere API 6.7.1.
	Host *string
    // The identifier of the resource pool to deploy the VCHA node to. This property was added in vSphere API 6.7.1.
	ResourcePool *string
    // The type of the Network object used by the HA network.
    //  If the PlacementSpec#haNetwork property is set, then the PlacementSpec#haNetworkType field must be set.
    //  If the PlacementSpec#haNetwork property is null, then the PlacementSpec#haNetworkType property is ignored. This property was added in vSphere API 6.7.1.
	HaNetworkType *NetworkType
    // The identifier of the Network object used for the HA network.
    //  If the PlacementSpec#haNetwork property is set, then the {#link #haNetworkType} property must be set.
    //  If the PlacementSpec#haNetwork property is null, then the PlacementSpec#haNetworkType property is ignored. This property was added in vSphere API 6.7.1.
	HaNetwork *string
    // The type of the Network object used by the Management network.
    //  If the PlacementSpec#managementNetwork property is set, then the {#link #managementNetworkType} field must be set.
    //  If the PlacementSpec#managementNetwork property is null, then the PlacementSpec#managementNetworkType property is ignored. This property was added in vSphere API 6.7.1.
	ManagementNetworkType *NetworkType
    // The identifier of the Network object used for the Management network. If the PlacementSpec#managementNetwork property is set, then the PlacementSpec#managementNetworkType property must be set.
    //  If the PlacementSpec#managementNetwork property is null, then the PlacementSpec#managementNetworkType property is ignored. This property was added in vSphere API 6.7.1.
	ManagementNetwork *string
    // The storage specification to deploy the VCHA node to. This property was added in vSphere API 6.7.1.
	Storage *DiskSpec
}




func CertificateInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["ssl_thumbprint"] = bindings.NewStringType()
	fieldNameMap["ssl_thumbprint"] = "SslThumbprint"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.vcha.certificate_info", fields, reflect.TypeOf(CertificateInfo{}), fieldNameMap, validators)
}

func ConnectionSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["hostname"] = bindings.NewStringType()
	fieldNameMap["hostname"] = "Hostname"
	fields["port"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["port"] = "Port"
	fields["ssl_thumbprint"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["ssl_thumbprint"] = "SslThumbprint"
	fields["username"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["username"] = "Username"
	fields["password"] = bindings.NewOptionalType(bindings.NewSecretType())
	fieldNameMap["password"] = "Password"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.vcha.connection_spec", fields, reflect.TypeOf(ConnectionSpec{}), fieldNameMap, validators)
}

func CredentialsSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["active_location"] = bindings.NewReferenceType(ConnectionSpecBindingType)
	fieldNameMap["active_location"] = "ActiveLocation"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.vcha.credentials_spec", fields, reflect.TypeOf(CredentialsSpec{}), fieldNameMap, validators)
}

func DiskInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["datastore"] = bindings.NewIdType([]string{"Datastore:VCenter"}, "")
	fieldNameMap["datastore"] = "Datastore"
	fields["datastore_name"] = bindings.NewStringType()
	fieldNameMap["datastore_name"] = "DatastoreName"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.vcha.disk_info", fields, reflect.TypeOf(DiskInfo{}), fieldNameMap, validators)
}

func DiskSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["datastore"] = bindings.NewOptionalType(bindings.NewIdType([]string{"Datastore:VCenter"}, ""))
	fieldNameMap["datastore"] = "Datastore"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.vcha.disk_spec", fields, reflect.TypeOf(DiskSpec{}), fieldNameMap, validators)
}

func IpSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["ip_family"] = bindings.NewEnumType("com.vmware.vcenter.vcha.ip_family", reflect.TypeOf(IpFamily(IpFamily_IPV4)))
	fieldNameMap["ip_family"] = "IpFamily"
	fields["ipv4"] = bindings.NewOptionalType(bindings.NewReferenceType(Ipv4SpecBindingType))
	fieldNameMap["ipv4"] = "Ipv4"
	fields["ipv6"] = bindings.NewOptionalType(bindings.NewReferenceType(Ipv6SpecBindingType))
	fieldNameMap["ipv6"] = "Ipv6"
	fields["default_gateway"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["default_gateway"] = "DefaultGateway"
	fields["dns_servers"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["dns_servers"] = "DnsServers"
	var validators = []bindings.Validator{}
	uv1 := bindings.NewUnionValidator("ip_family",
		map[string][]bindings.FieldData{
			"IPV4": []bindings.FieldData{
				bindings.NewFieldData("ipv4", true),
			},
			"IPV6": []bindings.FieldData{
				bindings.NewFieldData("ipv6", true),
			},
		},
	)
	validators = append(validators, uv1)
	return bindings.NewStructType("com.vmware.vcenter.vcha.ip_spec", fields, reflect.TypeOf(IpSpec{}), fieldNameMap, validators)
}

func Ipv4SpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["address"] = bindings.NewStringType()
	fieldNameMap["address"] = "Address"
	fields["subnet_mask"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["subnet_mask"] = "SubnetMask"
	fields["prefix"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["prefix"] = "Prefix"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.vcha.ipv4_spec", fields, reflect.TypeOf(Ipv4Spec{}), fieldNameMap, validators)
}

func Ipv6SpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["address"] = bindings.NewStringType()
	fieldNameMap["address"] = "Address"
	fields["prefix"] = bindings.NewIntegerType()
	fieldNameMap["prefix"] = "Prefix"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.vcha.ipv6_spec", fields, reflect.TypeOf(Ipv6Spec{}), fieldNameMap, validators)
}

func PlacementInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["management_vcenter_name"] = bindings.NewStringType()
	fieldNameMap["management_vcenter_name"] = "ManagementVcenterName"
	fields["vm_name"] = bindings.NewStringType()
	fieldNameMap["vm_name"] = "VmName"
	fields["datacenter"] = bindings.NewIdType([]string{"Datacenter:VCenter"}, "")
	fieldNameMap["datacenter"] = "Datacenter"
	fields["datacenter_name"] = bindings.NewStringType()
	fieldNameMap["datacenter_name"] = "DatacenterName"
	fields["host"] = bindings.NewIdType([]string{"HostSystem:VCenter"}, "")
	fieldNameMap["host"] = "Host"
	fields["host_name"] = bindings.NewStringType()
	fieldNameMap["host_name"] = "HostName"
	fields["cluster"] = bindings.NewOptionalType(bindings.NewIdType([]string{"ClusterComputeResource:VCenter"}, ""))
	fieldNameMap["cluster"] = "Cluster"
	fields["cluster_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cluster_name"] = "ClusterName"
	fields["ha_network"] = bindings.NewOptionalType(bindings.NewIdType([]string{"Network:VCenter"}, ""))
	fieldNameMap["ha_network"] = "HaNetwork"
	fields["ha_network_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["ha_network_name"] = "HaNetworkName"
	fields["ha_network_type"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.vcha.network_type", reflect.TypeOf(NetworkType(NetworkType_STANDARD_PORTGROUP))))
	fieldNameMap["ha_network_type"] = "HaNetworkType"
	fields["management_network"] = bindings.NewIdType([]string{"Network:VCenter"}, "")
	fieldNameMap["management_network"] = "ManagementNetwork"
	fields["management_network_name"] = bindings.NewStringType()
	fieldNameMap["management_network_name"] = "ManagementNetworkName"
	fields["management_network_type"] = bindings.NewEnumType("com.vmware.vcenter.vcha.network_type", reflect.TypeOf(NetworkType(NetworkType_STANDARD_PORTGROUP)))
	fieldNameMap["management_network_type"] = "ManagementNetworkType"
	fields["storage"] = bindings.NewReferenceType(DiskInfoBindingType)
	fieldNameMap["storage"] = "Storage"
	fields["bios_uuid"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["bios_uuid"] = "BiosUuid"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.vcha.placement_info", fields, reflect.TypeOf(PlacementInfo{}), fieldNameMap, validators)
}

func PlacementSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["name"] = bindings.NewStringType()
	fieldNameMap["name"] = "Name"
	fields["folder"] = bindings.NewIdType([]string{"Folder:VCenter"}, "")
	fieldNameMap["folder"] = "Folder"
	fields["host"] = bindings.NewOptionalType(bindings.NewIdType([]string{"HostSystem:VCenter"}, ""))
	fieldNameMap["host"] = "Host"
	fields["resource_pool"] = bindings.NewOptionalType(bindings.NewIdType([]string{"ResourcePool:VCenter"}, ""))
	fieldNameMap["resource_pool"] = "ResourcePool"
	fields["ha_network_type"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.vcha.network_type", reflect.TypeOf(NetworkType(NetworkType_STANDARD_PORTGROUP))))
	fieldNameMap["ha_network_type"] = "HaNetworkType"
	fields["ha_network"] = bindings.NewOptionalType(bindings.NewIdType([]string{"Network:VCenter"}, ""))
	fieldNameMap["ha_network"] = "HaNetwork"
	fields["management_network_type"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.vcha.network_type", reflect.TypeOf(NetworkType(NetworkType_STANDARD_PORTGROUP))))
	fieldNameMap["management_network_type"] = "ManagementNetworkType"
	fields["management_network"] = bindings.NewOptionalType(bindings.NewIdType([]string{"Network:VCenter"}, ""))
	fieldNameMap["management_network"] = "ManagementNetwork"
	fields["storage"] = bindings.NewOptionalType(bindings.NewReferenceType(DiskSpecBindingType))
	fieldNameMap["storage"] = "Storage"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.vcha.placement_spec", fields, reflect.TypeOf(PlacementSpec{}), fieldNameMap, validators)
}


