/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
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


// The ``IpFamily`` enumeration class defines the Ip address family.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.

type IpFamily string

const (
    // IPV4 address family
     IpFamily_IPV4 IpFamily = "IPV4"
    // IPv6 address family
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




// The ``NetworkType`` enumeration class defines the type of a vCenter Server network.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.

type NetworkType string

const (
    // vSphere standard portgroup network.
     NetworkType_STANDARD_PORTGROUP NetworkType = "STANDARD_PORTGROUP"
    // Distributed virtual switch.
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





// The ``CertificateInfo`` Class contains information about the SSL certificate for a management vCenter server.
type CertificateInfo struct {
    SslThumbprint string
}






// The ``ConnectionSpec`` class contains information required to connect to a vCenter server. The connection to the vCenter server always uses the HTTPS protocol.
type ConnectionSpec struct {
    Hostname string
    Port *int64
    SslThumbprint *string
    Username *string
    Password *string
}






// The ``CredentialsSpec`` class contains information to connect to the vCenter server managing the VCHA nodes.
type CredentialsSpec struct {
    ActiveLocation ConnectionSpec
}






// The ``DiskInfo`` class contains information to describe the storage configuration of a vCenter virtual machine.
type DiskInfo struct {
    Datastore string
    DatastoreName string
}






// The ``DiskSpec`` class contains information to describe the storage configuration of a vCenter virtual machine.
type DiskSpec struct {
    Datastore *string
}






// The ``IpSpec`` class contains IP information used to configure a network interface.
type IpSpec struct {
    IpFamily IpFamily
    Ipv4 *Ipv4Spec
    Ipv6 *Ipv6Spec
    DefaultGateway *string
    DnsServers []string
}






// The ``Ipv4Spec`` class contains IPV4 information used to configure a network interface.
type Ipv4Spec struct {
    Address string
    SubnetMask *string
    Prefix *int64
}






// The ``Ipv6Spec`` class contains IPV6 information used to configure a network interface.
type Ipv6Spec struct {
    Address string
    Prefix int64
}






// The ``PlacementInfo`` class contains information to describe the inventory placement of a single node of a VCHA cluster.
//  The active node's management vCenter server credentials are required to populate all properties except biosUuid.
type PlacementInfo struct {
    ManagementVcenterName string
    ManagementVcenterServerGuid *string
    VmName string
    Datacenter string
    DatacenterName string
    Host string
    HostName string
    Cluster *string
    ClusterName *string
    HaNetwork *string
    HaNetworkName *string
    HaNetworkType *NetworkType
    ManagementNetwork string
    ManagementNetworkName string
    ManagementNetworkType NetworkType
    Storage DiskInfo
    BiosUuid *string
}






// The ``PlacementSpec`` class contains information to describe the inventory placement of a single node of a VCHA cluster.
type PlacementSpec struct {
    Name string
    Folder string
    Host *string
    ResourcePool *string
    HaNetworkType *NetworkType
    HaNetwork *string
    ManagementNetworkType *NetworkType
    ManagementNetwork *string
    Storage *DiskSpec
}










func CertificateInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["ssl_thumbprint"] = bindings.NewStringType()
    fieldNameMap["ssl_thumbprint"] = "SslThumbprint"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vcha.certificate_info",fields, reflect.TypeOf(CertificateInfo{}), fieldNameMap, validators)
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
    return bindings.NewStructType("com.vmware.vcenter.vcha.connection_spec",fields, reflect.TypeOf(ConnectionSpec{}), fieldNameMap, validators)
}

func CredentialsSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["active_location"] = bindings.NewReferenceType(ConnectionSpecBindingType)
    fieldNameMap["active_location"] = "ActiveLocation"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vcha.credentials_spec",fields, reflect.TypeOf(CredentialsSpec{}), fieldNameMap, validators)
}

func DiskInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["datastore"] = bindings.NewIdType([]string {"Datastore:VCenter"}, "")
    fieldNameMap["datastore"] = "Datastore"
    fields["datastore_name"] = bindings.NewStringType()
    fieldNameMap["datastore_name"] = "DatastoreName"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vcha.disk_info",fields, reflect.TypeOf(DiskInfo{}), fieldNameMap, validators)
}

func DiskSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["datastore"] = bindings.NewOptionalType(bindings.NewIdType([]string {"Datastore:VCenter"}, ""))
    fieldNameMap["datastore"] = "Datastore"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vcha.disk_spec",fields, reflect.TypeOf(DiskSpec{}), fieldNameMap, validators)
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
        map[string][]bindings.FieldData {
            "IPV4": []bindings.FieldData {
                 bindings.NewFieldData("ipv4", true),
            },
            "IPV6": []bindings.FieldData {
                 bindings.NewFieldData("ipv6", true),
            },
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.vcenter.vcha.ip_spec",fields, reflect.TypeOf(IpSpec{}), fieldNameMap, validators)
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
    return bindings.NewStructType("com.vmware.vcenter.vcha.ipv4_spec",fields, reflect.TypeOf(Ipv4Spec{}), fieldNameMap, validators)
}

func Ipv6SpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["address"] = bindings.NewStringType()
    fieldNameMap["address"] = "Address"
    fields["prefix"] = bindings.NewIntegerType()
    fieldNameMap["prefix"] = "Prefix"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vcha.ipv6_spec",fields, reflect.TypeOf(Ipv6Spec{}), fieldNameMap, validators)
}

func PlacementInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["management_vcenter_name"] = bindings.NewStringType()
    fieldNameMap["management_vcenter_name"] = "ManagementVcenterName"
    fields["management_vcenter_server_guid"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["management_vcenter_server_guid"] = "ManagementVcenterServerGuid"
    fields["vm_name"] = bindings.NewStringType()
    fieldNameMap["vm_name"] = "VmName"
    fields["datacenter"] = bindings.NewIdType([]string {"Datacenter:VCenter"}, "")
    fieldNameMap["datacenter"] = "Datacenter"
    fields["datacenter_name"] = bindings.NewStringType()
    fieldNameMap["datacenter_name"] = "DatacenterName"
    fields["host"] = bindings.NewIdType([]string {"HostSystem:VCenter"}, "")
    fieldNameMap["host"] = "Host"
    fields["host_name"] = bindings.NewStringType()
    fieldNameMap["host_name"] = "HostName"
    fields["cluster"] = bindings.NewOptionalType(bindings.NewIdType([]string {"ClusterComputeResource:VCenter"}, ""))
    fieldNameMap["cluster"] = "Cluster"
    fields["cluster_name"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["cluster_name"] = "ClusterName"
    fields["ha_network"] = bindings.NewOptionalType(bindings.NewIdType([]string {"Network:VCenter"}, ""))
    fieldNameMap["ha_network"] = "HaNetwork"
    fields["ha_network_name"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["ha_network_name"] = "HaNetworkName"
    fields["ha_network_type"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.vcha.network_type", reflect.TypeOf(NetworkType(NetworkType_STANDARD_PORTGROUP))))
    fieldNameMap["ha_network_type"] = "HaNetworkType"
    fields["management_network"] = bindings.NewIdType([]string {"Network:VCenter"}, "")
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
    return bindings.NewStructType("com.vmware.vcenter.vcha.placement_info",fields, reflect.TypeOf(PlacementInfo{}), fieldNameMap, validators)
}

func PlacementSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    fields["folder"] = bindings.NewIdType([]string {"Folder:VCenter"}, "")
    fieldNameMap["folder"] = "Folder"
    fields["host"] = bindings.NewOptionalType(bindings.NewIdType([]string {"HostSystem:VCenter"}, ""))
    fieldNameMap["host"] = "Host"
    fields["resource_pool"] = bindings.NewOptionalType(bindings.NewIdType([]string {"ResourcePool:VCenter"}, ""))
    fieldNameMap["resource_pool"] = "ResourcePool"
    fields["ha_network_type"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.vcha.network_type", reflect.TypeOf(NetworkType(NetworkType_STANDARD_PORTGROUP))))
    fieldNameMap["ha_network_type"] = "HaNetworkType"
    fields["ha_network"] = bindings.NewOptionalType(bindings.NewIdType([]string {"Network:VCenter"}, ""))
    fieldNameMap["ha_network"] = "HaNetwork"
    fields["management_network_type"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.vcha.network_type", reflect.TypeOf(NetworkType(NetworkType_STANDARD_PORTGROUP))))
    fieldNameMap["management_network_type"] = "ManagementNetworkType"
    fields["management_network"] = bindings.NewOptionalType(bindings.NewIdType([]string {"Network:VCenter"}, ""))
    fieldNameMap["management_network"] = "ManagementNetwork"
    fields["storage"] = bindings.NewOptionalType(bindings.NewReferenceType(DiskSpecBindingType))
    fieldNameMap["storage"] = "Storage"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vcha.placement_spec",fields, reflect.TypeOf(PlacementSpec{}), fieldNameMap, validators)
}


