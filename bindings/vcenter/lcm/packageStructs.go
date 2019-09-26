/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for package: com.vmware.vcenter.lcm.
 * Includes binding types of a top level structures and enumerations.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package lcm

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/cis/task"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/vapi/std"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "time"
)


// The size of appliance to be deployed. **Warning:** This enumeration is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.

type ApplianceSize string

const (
    // Appliance size of 'tiny', Default vCPUs: 2, Memory: 8GB, VM: 100, Hosts: 10. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     ApplianceSize_TINY ApplianceSize = "TINY"
    // Appliance size of 'small', Default vCPUs: 4, Memory: 16GB, VM: 1000, Hosts: 100. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     ApplianceSize_SMALL ApplianceSize = "SMALL"
    // Appliance size of 'medium', Default vCPUs: 8, Memory: 24GB, VM: 4000, Hosts: 400. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     ApplianceSize_MEDIUM ApplianceSize = "MEDIUM"
    // Appliance size of 'large', Default vCPUs: 16, Memory: 32GB, VM: 10000, Hosts: 1000. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     ApplianceSize_LARGE ApplianceSize = "LARGE"
    // Appliance size of 'extra large', Default vCPUs: 24, Memory: 48GB, VM: 35000, Hosts: 2000. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     ApplianceSize_XLARGE ApplianceSize = "XLARGE"
)

func (a ApplianceSize) ApplianceSize() bool {
    switch a {
        case ApplianceSize_TINY:
            return true
        case ApplianceSize_SMALL:
            return true
        case ApplianceSize_MEDIUM:
            return true
        case ApplianceSize_LARGE:
            return true
        case ApplianceSize_XLARGE:
            return true
        default:
            return false
    }
}




// The type of appliance to be deployed. **Warning:** This enumeration is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.

type ApplianceType string

const (
    // Management node. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     ApplianceType_VCSA_EXTERNAL ApplianceType = "VCSA_EXTERNAL"
    // Embedded node. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     ApplianceType_VCSA_EMBEDDED ApplianceType = "VCSA_EMBEDDED"
    // Infrastructure node. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     ApplianceType_PSC ApplianceType = "PSC"
    // VMC node. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     ApplianceType_VMC ApplianceType = "VMC"
)

func (a ApplianceType) ApplianceType() bool {
    switch a {
        case ApplianceType_VCSA_EXTERNAL:
            return true
        case ApplianceType_VCSA_EMBEDDED:
            return true
        case ApplianceType_PSC:
            return true
        case ApplianceType_VMC:
            return true
        default:
            return false
    }
}




// The storage size of the appliance to be deployed. **Warning:** This enumeration is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.

type StorageSize string

const (
    // Large storage. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     StorageSize_LARGE StorageSize = "LARGE"
    // Extra large storage. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     StorageSize_XLARGE StorageSize = "XLARGE"
    // Regular storage. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     StorageSize_REGULAR StorageSize = "REGULAR"
)

func (s StorageSize) StorageSize() bool {
    switch s {
        case StorageSize_LARGE:
            return true
        case StorageSize_XLARGE:
            return true
        case StorageSize_REGULAR:
            return true
        default:
            return false
    }
}





// Connection information for source/destination location. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type Connection struct {
    Hostname string
    Username string
    Password string
    HttpsPort *int64
    SslVerify *bool
    SslThumbprint *string
}






// Information about the appliance deployed. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type DeploymentInfo struct {
    ApplianceName string
    ApplianceFqdn *string
    ApplianceIps []string
}






// Container to control deployment. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type DeploymentOption struct {
    SkipOptions map[DeploymentOption_SkipOptions]bool
}




    
    // Skippable tasks. **Warning:** This enumeration is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
     
    type DeploymentOption_SkipOptions string

    const (
        // Skips the sso check. This should only be used when performing precheck for install/upgrade of management node before infrastructure node is deployed. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
         DeploymentOption_SkipOptions_SKIP_SSO_CHECK DeploymentOption_SkipOptions = "SKIP_SSO_CHECK"
    )

    func (s DeploymentOption_SkipOptions) DeploymentOption_SkipOptions() bool {
        switch s {
            case DeploymentOption_SkipOptions_SKIP_SSO_CHECK:
                return true
            default:
                return false
        }
    }



// Spec to describe the new appliance. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type DestinationAppliance struct {
    ApplianceName string
    ApplianceType ApplianceType
    ApplianceSize *ApplianceSize
    ApplianceDiskSize *StorageSize
    RootPassword string
    ThinDiskMode bool
    OvaLocation string
    OvaLocationSslVerify *bool
    OvaLocationSslThumbprint *string
    OvftoolLocation string
    OvftoolLocationSslVerify *bool
    OvftoolLocationSslThumbprint *string
    Services DestinationApplianceService
    Network Network
    Time Time
    OvftoolArguments map[string]string
    VcsaEmbedded *VcsaEmbedded
    Vmc *ExternalVcsa
}






// The configuration of vCenter services. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type DestinationApplianceService struct {
    Rhttpproxy *ReverseProxy
    Ssh Ssh
}






// Configuration of destination location. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type DestinationLocation struct {
    Esx *Esx
    Vcenter *Vc
}






// Configuration of the replicated Single Sign-On for Embedded type deployment. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type EmbeddedReplicatedVcsa struct {
    SsoAdminPassword string
    SsoDomainName string
    PartnerHostname string
    SslVerify *bool
    SslThumbprint *string
    HttpsPort *int64
}






// Configuration of the standalone Single Sign-On for Embedded type deployment. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type EmbeddedStandaloneVcsa struct {
    SsoAdminPassword string
    SsoDomainName string
}






// Configuration of ESX. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type Esx struct {
    Connection Connection
    Inventory EsxInventory
}






// Configuration of ESX's inventory. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type EsxInventory struct {
    DatastoreName string
    NetworkName *string
    ResourcePoolPath *string
}






// Configuration of the external tools used. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type ExternalTool struct {
    Name string
    Hostname *string
    Location string
}






// Configuration of the Single Sign-On for Management type deployment. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type ExternalVcsa struct {
    SsoAdminPassword string
    SsoDomainName string
    PscHostname string
    SslVerify *bool
    SslThumbprint *string
    HttpsPort *int64
}






// The specification that represents a install operation. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type InstallSpec struct {
    DestinationLocation DestinationLocation
    DestinationAppliance DestinationAppliance
    ScheduledStart *time.Time
    DelayTolerance *int64
}






// Network configuration of the appliance to be deployed. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type Network struct {
    Hostname *string
    IpFamily *TemporaryNetwork_IpType
    Mode TemporaryNetwork_NetworkMode
    Ip *string
    DnsServers []string
    Prefix *int64
    Gateway *string
}






// The ``Notification`` class describes a notification that can be reported by the appliance task, which can be of type info, warning or errors.
type Notification struct {
    Id string
    Time *time.Time
    Message std.LocalizableMessage
    Resolution *std.LocalizableMessage
}






// The ``Notifications`` class contains info/warning/error messages that can be reported be the appliance task.
type Notifications struct {
    Info []Notification
    Warnings []Notification
    Errors []Notification
}






// Spec used to configure a Platform Services Controller. This section describes how the Platform Services Controller appliance should be configured. If unset, either ``#vcsaEmbedded`` or ``#vcsaExternal`` must be provided. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type Psc struct {
    Standalone *PscStandalone
    Replicated *PscReplicated
    CeipEnabled bool
}






// Configuration of the replicated Single Sign-On for PSC type deployment. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type PscReplicated struct {
    SsoAdminPassword string
    SsoDomainName string
    PartnerHostname string
    SslVerify *bool
    SslThumbprint *string
    HttpsPort *int64
    SsoSiteName *string
}






// Configuration of the standalone Single Sign-On for Embedded type deployment. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type PscStandalone struct {
    SsoAdminPassword string
    SsoDomainName string
    SsoSiteName *string
}






// Container of info, warning and error messages associated with a single task. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type Result struct {
    Info []Notification
    Warnings []Notification
    Errors []Notification
}






// Port numbers on which the vCenter Server Appliance communicates with the other vSphere components. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type ReverseProxy struct {
    HttpPort *int64
    HttpsPort *int64
}






// Setting to enable SSH on the deployed appliance. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type Ssh struct {
    Enabled *bool
}






// Container that contains the status information about a single task. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type SubTaskInfo struct {
    Progress task.Progress
    LastUpdatedTime time.Time
    Result *Result
    ExternalTools []ExternalTool
    Description std.LocalizableMessage
    Service string
    Operation string
    Parent *string
    Target *std.DynamicID
    Status task.Status
    Cancelable bool
    Error *data.ErrorValue
    StartTime *time.Time
    EndTime *time.Time
    User *string
}






// The container that contains the status information of a deployment. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type TaskInfo struct {
    MetadataFile string
    State *string
    Progress *task.Progress
    LastUpdatedTime time.Time
    SubtaskOrder [][]string
    Subtasks map[string]SubTaskInfo
    ApplianceInfo *DeploymentInfo
    Result data.DataValue
    AdditionalInfo *string
    Description std.LocalizableMessage
    Service string
    Operation string
    Parent *string
    Target *std.DynamicID
    Status task.Status
    Cancelable bool
    Error *data.ErrorValue
    StartTime *time.Time
    EndTime *time.Time
    User *string
}






// Configuration of the temporary network which is used during upgrade/migrate. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type TemporaryNetwork struct {
    IpFamily *TemporaryNetwork_IpType
    Mode TemporaryNetwork_NetworkMode
    Ip *string
    DnsServers []string
    Prefix *int64
    Gateway *string
}




    
    // Network IP address family. **Warning:** This enumeration is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
     
    type TemporaryNetwork_IpType string

    const (
        // IPv4 Type of IP address. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
         TemporaryNetwork_IpType_IPV4 TemporaryNetwork_IpType = "IPV4"
        // IPv6 Type of IP address. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
         TemporaryNetwork_IpType_IPV6 TemporaryNetwork_IpType = "IPV6"
    )

    func (i TemporaryNetwork_IpType) TemporaryNetwork_IpType() bool {
        switch i {
            case TemporaryNetwork_IpType_IPV4:
                return true
            case TemporaryNetwork_IpType_IPV6:
                return true
            default:
                return false
        }
    }

    
    // Network mode. **Warning:** This enumeration is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
     
    type TemporaryNetwork_NetworkMode string

    const (
        // DHCP mode. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
         TemporaryNetwork_NetworkMode_DHCP TemporaryNetwork_NetworkMode = "DHCP"
        // Static IP mode. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
         TemporaryNetwork_NetworkMode_STATIC TemporaryNetwork_NetworkMode = "STATIC"
    )

    func (n TemporaryNetwork_NetworkMode) TemporaryNetwork_NetworkMode() bool {
        switch n {
            case TemporaryNetwork_NetworkMode_DHCP:
                return true
            case TemporaryNetwork_NetworkMode_STATIC:
                return true
            default:
                return false
        }
    }



// NTP setting of the appliance to be deployed. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type Time struct {
    NtpServers []string
}






// Configuration of the VC that hosts/will host an appliance. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type Vc struct {
    Connection Connection
    Inventory VcInventory
}






// Inventory information about a VCenter. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type VcInventory struct {
    VmFolderPath *string
    ResourcePoolPath *string
    ClusterPath *string
    HostPath *string
    DatastoreName *string
    DatastoreClusterName *string
    NetworkName string
}






// Spec used to configure an embedded vCenter Server. This field describes how the embedded vCenter Server appliance should be configured. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type VcsaEmbedded struct {
    Standalone *EmbeddedStandaloneVcsa
    Replicated *EmbeddedReplicatedVcsa
    CeipEnabled bool
}










func ConnectionBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["hostname"] = bindings.NewStringType()
    fieldNameMap["hostname"] = "Hostname"
    fields["username"] = bindings.NewStringType()
    fieldNameMap["username"] = "Username"
    fields["password"] = bindings.NewSecretType()
    fieldNameMap["password"] = "Password"
    fields["https_port"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["https_port"] = "HttpsPort"
    fields["ssl_verify"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["ssl_verify"] = "SslVerify"
    fields["ssl_thumbprint"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["ssl_thumbprint"] = "SslThumbprint"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.lcm.connection",fields, reflect.TypeOf(Connection{}), fieldNameMap, validators)
}

func DeploymentInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["appliance_name"] = bindings.NewStringType()
    fieldNameMap["appliance_name"] = "ApplianceName"
    fields["appliance_fqdn"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["appliance_fqdn"] = "ApplianceFqdn"
    fields["appliance_ips"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
    fieldNameMap["appliance_ips"] = "ApplianceIps"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.lcm.deployment_info",fields, reflect.TypeOf(DeploymentInfo{}), fieldNameMap, validators)
}

func DeploymentOptionBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["skip_options"] = bindings.NewMapType(bindings.NewEnumType("com.vmware.vcenter.lcm.deployment_option.skip_options", reflect.TypeOf(DeploymentOption_SkipOptions(DeploymentOption_SkipOptions_SKIP_SSO_CHECK))), bindings.NewBooleanType(),reflect.TypeOf(map[DeploymentOption_SkipOptions]bool{}))
    fieldNameMap["skip_options"] = "SkipOptions"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.lcm.deployment_option",fields, reflect.TypeOf(DeploymentOption{}), fieldNameMap, validators)
}

func DestinationApplianceBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["appliance_name"] = bindings.NewStringType()
    fieldNameMap["appliance_name"] = "ApplianceName"
    fields["appliance_type"] = bindings.NewEnumType("com.vmware.vcenter.lcm.appliance_type", reflect.TypeOf(ApplianceType(ApplianceType_VCSA_EXTERNAL)))
    fieldNameMap["appliance_type"] = "ApplianceType"
    fields["appliance_size"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.lcm.appliance_size", reflect.TypeOf(ApplianceSize(ApplianceSize_TINY))))
    fieldNameMap["appliance_size"] = "ApplianceSize"
    fields["appliance_disk_size"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.lcm.storage_size", reflect.TypeOf(StorageSize(StorageSize_LARGE))))
    fieldNameMap["appliance_disk_size"] = "ApplianceDiskSize"
    fields["root_password"] = bindings.NewSecretType()
    fieldNameMap["root_password"] = "RootPassword"
    fields["thin_disk_mode"] = bindings.NewBooleanType()
    fieldNameMap["thin_disk_mode"] = "ThinDiskMode"
    fields["ova_location"] = bindings.NewStringType()
    fieldNameMap["ova_location"] = "OvaLocation"
    fields["ova_location_ssl_verify"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["ova_location_ssl_verify"] = "OvaLocationSslVerify"
    fields["ova_location_ssl_thumbprint"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["ova_location_ssl_thumbprint"] = "OvaLocationSslThumbprint"
    fields["ovftool_location"] = bindings.NewStringType()
    fieldNameMap["ovftool_location"] = "OvftoolLocation"
    fields["ovftool_location_ssl_verify"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["ovftool_location_ssl_verify"] = "OvftoolLocationSslVerify"
    fields["ovftool_location_ssl_thumbprint"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["ovftool_location_ssl_thumbprint"] = "OvftoolLocationSslThumbprint"
    fields["services"] = bindings.NewReferenceType(DestinationApplianceServiceBindingType)
    fieldNameMap["services"] = "Services"
    fields["network"] = bindings.NewReferenceType(NetworkBindingType)
    fieldNameMap["network"] = "Network"
    fields["time"] = bindings.NewReferenceType(TimeBindingType)
    fieldNameMap["time"] = "Time"
    fields["ovftool_arguments"] = bindings.NewOptionalType(bindings.NewMapType(bindings.NewStringType(), bindings.NewStringType(),reflect.TypeOf(map[string]string{})))
    fieldNameMap["ovftool_arguments"] = "OvftoolArguments"
    fields["vcsa_embedded"] = bindings.NewOptionalType(bindings.NewReferenceType(VcsaEmbeddedBindingType))
    fieldNameMap["vcsa_embedded"] = "VcsaEmbedded"
    fields["vmc"] = bindings.NewOptionalType(bindings.NewReferenceType(ExternalVcsaBindingType))
    fieldNameMap["vmc"] = "Vmc"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("appliance_type",
        map[string][]bindings.FieldData {
            "VCSA_EMBEDDED": []bindings.FieldData {
                 bindings.NewFieldData("appliance_size", false),
                 bindings.NewFieldData("appliance_disk_size", false),
                 bindings.NewFieldData("vcsa_embedded", true),
            },
            "VMC": []bindings.FieldData {
                 bindings.NewFieldData("vmc", true),
            },
            "VCSA_EXTERNAL": []bindings.FieldData {},
            "PSC": []bindings.FieldData {},
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.vcenter.lcm.destination_appliance",fields, reflect.TypeOf(DestinationAppliance{}), fieldNameMap, validators)
}

func DestinationApplianceServiceBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["rhttpproxy"] = bindings.NewOptionalType(bindings.NewReferenceType(ReverseProxyBindingType))
    fieldNameMap["rhttpproxy"] = "Rhttpproxy"
    fields["ssh"] = bindings.NewReferenceType(SshBindingType)
    fieldNameMap["ssh"] = "Ssh"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.lcm.destination_appliance_service",fields, reflect.TypeOf(DestinationApplianceService{}), fieldNameMap, validators)
}

func DestinationLocationBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["esx"] = bindings.NewOptionalType(bindings.NewReferenceType(EsxBindingType))
    fieldNameMap["esx"] = "Esx"
    fields["vcenter"] = bindings.NewOptionalType(bindings.NewReferenceType(VcBindingType))
    fieldNameMap["vcenter"] = "Vcenter"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.lcm.destination_location",fields, reflect.TypeOf(DestinationLocation{}), fieldNameMap, validators)
}

func EmbeddedReplicatedVcsaBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["sso_admin_password"] = bindings.NewSecretType()
    fieldNameMap["sso_admin_password"] = "SsoAdminPassword"
    fields["sso_domain_name"] = bindings.NewStringType()
    fieldNameMap["sso_domain_name"] = "SsoDomainName"
    fields["partner_hostname"] = bindings.NewStringType()
    fieldNameMap["partner_hostname"] = "PartnerHostname"
    fields["ssl_verify"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["ssl_verify"] = "SslVerify"
    fields["ssl_thumbprint"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["ssl_thumbprint"] = "SslThumbprint"
    fields["https_port"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["https_port"] = "HttpsPort"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.lcm.embedded_replicated_vcsa",fields, reflect.TypeOf(EmbeddedReplicatedVcsa{}), fieldNameMap, validators)
}

func EmbeddedStandaloneVcsaBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["sso_admin_password"] = bindings.NewSecretType()
    fieldNameMap["sso_admin_password"] = "SsoAdminPassword"
    fields["sso_domain_name"] = bindings.NewStringType()
    fieldNameMap["sso_domain_name"] = "SsoDomainName"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.lcm.embedded_standalone_vcsa",fields, reflect.TypeOf(EmbeddedStandaloneVcsa{}), fieldNameMap, validators)
}

func EsxBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["connection"] = bindings.NewReferenceType(ConnectionBindingType)
    fieldNameMap["connection"] = "Connection"
    fields["inventory"] = bindings.NewReferenceType(EsxInventoryBindingType)
    fieldNameMap["inventory"] = "Inventory"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.lcm.esx",fields, reflect.TypeOf(Esx{}), fieldNameMap, validators)
}

func EsxInventoryBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["datastore_name"] = bindings.NewStringType()
    fieldNameMap["datastore_name"] = "DatastoreName"
    fields["network_name"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["network_name"] = "NetworkName"
    fields["resource_pool_path"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["resource_pool_path"] = "ResourcePoolPath"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.lcm.esx_inventory",fields, reflect.TypeOf(EsxInventory{}), fieldNameMap, validators)
}

func ExternalToolBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    fields["hostname"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["hostname"] = "Hostname"
    fields["location"] = bindings.NewStringType()
    fieldNameMap["location"] = "Location"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.lcm.external_tool",fields, reflect.TypeOf(ExternalTool{}), fieldNameMap, validators)
}

func ExternalVcsaBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["sso_admin_password"] = bindings.NewSecretType()
    fieldNameMap["sso_admin_password"] = "SsoAdminPassword"
    fields["sso_domain_name"] = bindings.NewStringType()
    fieldNameMap["sso_domain_name"] = "SsoDomainName"
    fields["psc_hostname"] = bindings.NewStringType()
    fieldNameMap["psc_hostname"] = "PscHostname"
    fields["ssl_verify"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["ssl_verify"] = "SslVerify"
    fields["ssl_thumbprint"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["ssl_thumbprint"] = "SslThumbprint"
    fields["https_port"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["https_port"] = "HttpsPort"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.lcm.external_vcsa",fields, reflect.TypeOf(ExternalVcsa{}), fieldNameMap, validators)
}

func InstallSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["destination_location"] = bindings.NewReferenceType(DestinationLocationBindingType)
    fieldNameMap["destination_location"] = "DestinationLocation"
    fields["destination_appliance"] = bindings.NewReferenceType(DestinationApplianceBindingType)
    fieldNameMap["destination_appliance"] = "DestinationAppliance"
    fields["scheduled_start"] = bindings.NewOptionalType(bindings.NewDateTimeType())
    fieldNameMap["scheduled_start"] = "ScheduledStart"
    fields["delay_tolerance"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["delay_tolerance"] = "DelayTolerance"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.lcm.install_spec",fields, reflect.TypeOf(InstallSpec{}), fieldNameMap, validators)
}

func NetworkBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["hostname"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["hostname"] = "Hostname"
    fields["ip_family"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.lcm.temporary_network.ip_type", reflect.TypeOf(TemporaryNetwork_IpType(TemporaryNetwork_IpType_IPV4))))
    fieldNameMap["ip_family"] = "IpFamily"
    fields["mode"] = bindings.NewEnumType("com.vmware.vcenter.lcm.temporary_network.network_mode", reflect.TypeOf(TemporaryNetwork_NetworkMode(TemporaryNetwork_NetworkMode_DHCP)))
    fieldNameMap["mode"] = "Mode"
    fields["ip"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["ip"] = "Ip"
    fields["dns_servers"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
    fieldNameMap["dns_servers"] = "DnsServers"
    fields["prefix"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["prefix"] = "Prefix"
    fields["gateway"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["gateway"] = "Gateway"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("mode",
        map[string][]bindings.FieldData {
            "STATIC": []bindings.FieldData {
                 bindings.NewFieldData("hostname", false),
                 bindings.NewFieldData("ip", true),
                 bindings.NewFieldData("dns_servers", true),
                 bindings.NewFieldData("prefix", true),
                 bindings.NewFieldData("gateway", true),
            },
            "DHCP": []bindings.FieldData {},
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.vcenter.lcm.network",fields, reflect.TypeOf(Network{}), fieldNameMap, validators)
}

func NotificationBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["id"] = bindings.NewStringType()
    fieldNameMap["id"] = "Id"
    fields["time"] = bindings.NewOptionalType(bindings.NewDateTimeType())
    fieldNameMap["time"] = "Time"
    fields["message"] = bindings.NewReferenceType(std.LocalizableMessageBindingType)
    fieldNameMap["message"] = "Message"
    fields["resolution"] = bindings.NewOptionalType(bindings.NewReferenceType(std.LocalizableMessageBindingType))
    fieldNameMap["resolution"] = "Resolution"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.lcm.notification",fields, reflect.TypeOf(Notification{}), fieldNameMap, validators)
}

func NotificationsBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["info"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(NotificationBindingType), reflect.TypeOf([]Notification{})))
    fieldNameMap["info"] = "Info"
    fields["warnings"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(NotificationBindingType), reflect.TypeOf([]Notification{})))
    fieldNameMap["warnings"] = "Warnings"
    fields["errors"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(NotificationBindingType), reflect.TypeOf([]Notification{})))
    fieldNameMap["errors"] = "Errors"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.lcm.notifications",fields, reflect.TypeOf(Notifications{}), fieldNameMap, validators)
}

func PscBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["standalone"] = bindings.NewOptionalType(bindings.NewReferenceType(PscStandaloneBindingType))
    fieldNameMap["standalone"] = "Standalone"
    fields["replicated"] = bindings.NewOptionalType(bindings.NewReferenceType(PscReplicatedBindingType))
    fieldNameMap["replicated"] = "Replicated"
    fields["ceip_enabled"] = bindings.NewBooleanType()
    fieldNameMap["ceip_enabled"] = "CeipEnabled"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.lcm.psc",fields, reflect.TypeOf(Psc{}), fieldNameMap, validators)
}

func PscReplicatedBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["sso_admin_password"] = bindings.NewSecretType()
    fieldNameMap["sso_admin_password"] = "SsoAdminPassword"
    fields["sso_domain_name"] = bindings.NewStringType()
    fieldNameMap["sso_domain_name"] = "SsoDomainName"
    fields["partner_hostname"] = bindings.NewStringType()
    fieldNameMap["partner_hostname"] = "PartnerHostname"
    fields["ssl_verify"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["ssl_verify"] = "SslVerify"
    fields["ssl_thumbprint"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["ssl_thumbprint"] = "SslThumbprint"
    fields["https_port"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["https_port"] = "HttpsPort"
    fields["sso_site_name"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["sso_site_name"] = "SsoSiteName"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.lcm.psc_replicated",fields, reflect.TypeOf(PscReplicated{}), fieldNameMap, validators)
}

func PscStandaloneBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["sso_admin_password"] = bindings.NewSecretType()
    fieldNameMap["sso_admin_password"] = "SsoAdminPassword"
    fields["sso_domain_name"] = bindings.NewStringType()
    fieldNameMap["sso_domain_name"] = "SsoDomainName"
    fields["sso_site_name"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["sso_site_name"] = "SsoSiteName"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.lcm.psc_standalone",fields, reflect.TypeOf(PscStandalone{}), fieldNameMap, validators)
}

func ResultBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["info"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(NotificationBindingType), reflect.TypeOf([]Notification{})))
    fieldNameMap["info"] = "Info"
    fields["warnings"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(NotificationBindingType), reflect.TypeOf([]Notification{})))
    fieldNameMap["warnings"] = "Warnings"
    fields["errors"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(NotificationBindingType), reflect.TypeOf([]Notification{})))
    fieldNameMap["errors"] = "Errors"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.lcm.result",fields, reflect.TypeOf(Result{}), fieldNameMap, validators)
}

func ReverseProxyBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["http_port"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["http_port"] = "HttpPort"
    fields["https_port"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["https_port"] = "HttpsPort"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.lcm.reverse_proxy",fields, reflect.TypeOf(ReverseProxy{}), fieldNameMap, validators)
}

func SshBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["enabled"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["enabled"] = "Enabled"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.lcm.ssh",fields, reflect.TypeOf(Ssh{}), fieldNameMap, validators)
}

func SubTaskInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["progress"] = bindings.NewReferenceType(task.ProgressBindingType)
    fieldNameMap["progress"] = "Progress"
    fields["last_updated_time"] = bindings.NewDateTimeType()
    fieldNameMap["last_updated_time"] = "LastUpdatedTime"
    fields["result"] = bindings.NewOptionalType(bindings.NewReferenceType(ResultBindingType))
    fieldNameMap["result"] = "Result"
    fields["external_tools"] = bindings.NewListType(bindings.NewReferenceType(ExternalToolBindingType), reflect.TypeOf([]ExternalTool{}))
    fieldNameMap["external_tools"] = "ExternalTools"
    fields["description"] = bindings.NewReferenceType(std.LocalizableMessageBindingType)
    fieldNameMap["description"] = "Description"
    fields["service"] = bindings.NewIdType([]string {"com.vmware.vapi.service"}, "")
    fieldNameMap["service"] = "Service"
    fields["operation"] = bindings.NewIdType([]string {"com.vmware.vapi.operation"}, "")
    fieldNameMap["operation"] = "Operation"
    fields["parent"] = bindings.NewOptionalType(bindings.NewIdType([]string {"com.vmware.cis.task"}, ""))
    fieldNameMap["parent"] = "Parent"
    fields["target"] = bindings.NewOptionalType(bindings.NewReferenceType(std.DynamicIDBindingType))
    fieldNameMap["target"] = "Target"
    fields["status"] = bindings.NewEnumType("com.vmware.cis.task.status", reflect.TypeOf(task.Status(task.Status_PENDING)))
    fieldNameMap["status"] = "Status"
    fields["cancelable"] = bindings.NewBooleanType()
    fieldNameMap["cancelable"] = "Cancelable"
    fields["error"] = bindings.NewOptionalType(bindings.NewAnyErrorType())
    fieldNameMap["error"] = "Error"
    fields["start_time"] = bindings.NewOptionalType(bindings.NewDateTimeType())
    fieldNameMap["start_time"] = "StartTime"
    fields["end_time"] = bindings.NewOptionalType(bindings.NewDateTimeType())
    fieldNameMap["end_time"] = "EndTime"
    fields["user"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["user"] = "User"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("status",
        map[string][]bindings.FieldData {
            "FAILED": []bindings.FieldData {
                 bindings.NewFieldData("error", false),
                 bindings.NewFieldData("start_time", true),
                 bindings.NewFieldData("end_time", true),
            },
            "RUNNING": []bindings.FieldData {
                 bindings.NewFieldData("start_time", true),
            },
            "BLOCKED": []bindings.FieldData {
                 bindings.NewFieldData("start_time", true),
            },
            "SUCCEEDED": []bindings.FieldData {
                 bindings.NewFieldData("start_time", true),
                 bindings.NewFieldData("end_time", true),
            },
            "PENDING": []bindings.FieldData {},
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.vcenter.lcm.sub_task_info",fields, reflect.TypeOf(SubTaskInfo{}), fieldNameMap, validators)
}

func TaskInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["metadata_file"] = bindings.NewStringType()
    fieldNameMap["metadata_file"] = "MetadataFile"
    fields["state"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["state"] = "State"
    fields["progress"] = bindings.NewOptionalType(bindings.NewReferenceType(task.ProgressBindingType))
    fieldNameMap["progress"] = "Progress"
    fields["last_updated_time"] = bindings.NewDateTimeType()
    fieldNameMap["last_updated_time"] = "LastUpdatedTime"
    fields["subtask_order"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})), reflect.TypeOf([][]string{})))
    fieldNameMap["subtask_order"] = "SubtaskOrder"
    fields["subtasks"] = bindings.NewOptionalType(bindings.NewMapType(bindings.NewStringType(), bindings.NewReferenceType(SubTaskInfoBindingType),reflect.TypeOf(map[string]SubTaskInfo{})))
    fieldNameMap["subtasks"] = "Subtasks"
    fields["appliance_info"] = bindings.NewOptionalType(bindings.NewReferenceType(DeploymentInfoBindingType))
    fieldNameMap["appliance_info"] = "ApplianceInfo"
    fields["result"] = bindings.NewOptionalType(bindings.NewOpaqueType())
    fieldNameMap["result"] = "Result"
    fields["additional_info"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["additional_info"] = "AdditionalInfo"
    fields["description"] = bindings.NewReferenceType(std.LocalizableMessageBindingType)
    fieldNameMap["description"] = "Description"
    fields["service"] = bindings.NewIdType([]string {"com.vmware.vapi.service"}, "")
    fieldNameMap["service"] = "Service"
    fields["operation"] = bindings.NewIdType([]string {"com.vmware.vapi.operation"}, "")
    fieldNameMap["operation"] = "Operation"
    fields["parent"] = bindings.NewOptionalType(bindings.NewIdType([]string {"com.vmware.cis.task"}, ""))
    fieldNameMap["parent"] = "Parent"
    fields["target"] = bindings.NewOptionalType(bindings.NewReferenceType(std.DynamicIDBindingType))
    fieldNameMap["target"] = "Target"
    fields["status"] = bindings.NewEnumType("com.vmware.cis.task.status", reflect.TypeOf(task.Status(task.Status_PENDING)))
    fieldNameMap["status"] = "Status"
    fields["cancelable"] = bindings.NewBooleanType()
    fieldNameMap["cancelable"] = "Cancelable"
    fields["error"] = bindings.NewOptionalType(bindings.NewAnyErrorType())
    fieldNameMap["error"] = "Error"
    fields["start_time"] = bindings.NewOptionalType(bindings.NewDateTimeType())
    fieldNameMap["start_time"] = "StartTime"
    fields["end_time"] = bindings.NewOptionalType(bindings.NewDateTimeType())
    fieldNameMap["end_time"] = "EndTime"
    fields["user"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["user"] = "User"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("status",
        map[string][]bindings.FieldData {
            "RUNNING": []bindings.FieldData {
                 bindings.NewFieldData("progress", true),
                 bindings.NewFieldData("start_time", true),
            },
            "FAILED": []bindings.FieldData {
                 bindings.NewFieldData("progress", true),
                 bindings.NewFieldData("error", false),
                 bindings.NewFieldData("start_time", true),
                 bindings.NewFieldData("end_time", true),
            },
            "SUCCEEDED": []bindings.FieldData {
                 bindings.NewFieldData("progress", true),
                 bindings.NewFieldData("start_time", true),
                 bindings.NewFieldData("end_time", true),
            },
            "BLOCKED": []bindings.FieldData {
                 bindings.NewFieldData("progress", true),
                 bindings.NewFieldData("start_time", true),
            },
            "PENDING": []bindings.FieldData {},
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.vcenter.lcm.task_info",fields, reflect.TypeOf(TaskInfo{}), fieldNameMap, validators)
}

func TemporaryNetworkBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["ip_family"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.lcm.temporary_network.ip_type", reflect.TypeOf(TemporaryNetwork_IpType(TemporaryNetwork_IpType_IPV4))))
    fieldNameMap["ip_family"] = "IpFamily"
    fields["mode"] = bindings.NewEnumType("com.vmware.vcenter.lcm.temporary_network.network_mode", reflect.TypeOf(TemporaryNetwork_NetworkMode(TemporaryNetwork_NetworkMode_DHCP)))
    fieldNameMap["mode"] = "Mode"
    fields["ip"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["ip"] = "Ip"
    fields["dns_servers"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
    fieldNameMap["dns_servers"] = "DnsServers"
    fields["prefix"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["prefix"] = "Prefix"
    fields["gateway"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["gateway"] = "Gateway"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("mode",
        map[string][]bindings.FieldData {
            "STATIC": []bindings.FieldData {
                 bindings.NewFieldData("ip", true),
                 bindings.NewFieldData("dns_servers", true),
                 bindings.NewFieldData("prefix", true),
                 bindings.NewFieldData("gateway", true),
            },
            "DHCP": []bindings.FieldData {},
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.vcenter.lcm.temporary_network",fields, reflect.TypeOf(TemporaryNetwork{}), fieldNameMap, validators)
}

func TimeBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["ntp_servers"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
    fieldNameMap["ntp_servers"] = "NtpServers"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.lcm.time",fields, reflect.TypeOf(Time{}), fieldNameMap, validators)
}

func VcBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["connection"] = bindings.NewReferenceType(ConnectionBindingType)
    fieldNameMap["connection"] = "Connection"
    fields["inventory"] = bindings.NewReferenceType(VcInventoryBindingType)
    fieldNameMap["inventory"] = "Inventory"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.lcm.vc",fields, reflect.TypeOf(Vc{}), fieldNameMap, validators)
}

func VcInventoryBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vm_folder_path"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["vm_folder_path"] = "VmFolderPath"
    fields["resource_pool_path"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["resource_pool_path"] = "ResourcePoolPath"
    fields["cluster_path"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["cluster_path"] = "ClusterPath"
    fields["host_path"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["host_path"] = "HostPath"
    fields["datastore_name"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["datastore_name"] = "DatastoreName"
    fields["datastore_cluster_name"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["datastore_cluster_name"] = "DatastoreClusterName"
    fields["network_name"] = bindings.NewStringType()
    fieldNameMap["network_name"] = "NetworkName"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.lcm.vc_inventory",fields, reflect.TypeOf(VcInventory{}), fieldNameMap, validators)
}

func VcsaEmbeddedBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["standalone"] = bindings.NewOptionalType(bindings.NewReferenceType(EmbeddedStandaloneVcsaBindingType))
    fieldNameMap["standalone"] = "Standalone"
    fields["replicated"] = bindings.NewOptionalType(bindings.NewReferenceType(EmbeddedReplicatedVcsaBindingType))
    fieldNameMap["replicated"] = "Replicated"
    fields["ceip_enabled"] = bindings.NewBooleanType()
    fieldNameMap["ceip_enabled"] = "CeipEnabled"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.lcm.vcsa_embedded",fields, reflect.TypeOf(VcsaEmbedded{}), fieldNameMap, validators)
}


