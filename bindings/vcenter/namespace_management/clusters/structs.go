/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Clusters.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package clusters

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/vapi/std"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/bindings/vcenter/namespace_management"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)



// The ``ConfigStatus`` enumeration class describes the status of reaching the desired state configuration for the cluster. **Warning:** This enumeration is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type ConfigStatus string

const (
    // The Namespace configuration is being applied to the cluster. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     ConfigStatus_CONFIGURING ConfigStatus = "CONFIGURING"
    // The Namespace configuration is being removed from the cluster. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     ConfigStatus_REMOVING ConfigStatus = "REMOVING"
    // The cluster is configured correctly with the Namespace configuration. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     ConfigStatus_RUNNING ConfigStatus = "RUNNING"
    // Failed to apply the Namespace configuration to the cluster, user intervention needed. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     ConfigStatus_ERROR ConfigStatus = "ERROR"
)

func (c ConfigStatus) ConfigStatus() bool {
    switch c {
        case ConfigStatus_CONFIGURING:
            return true
        case ConfigStatus_REMOVING:
            return true
        case ConfigStatus_RUNNING:
            return true
        case ConfigStatus_ERROR:
            return true
        default:
            return false
    }
}




// The ``KubernetesStatus`` enumeration class describes the cluster's ability to deploy pods. **Warning:** This enumeration is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type KubernetesStatus string

const (
    // The cluster is able to accept pods. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     KubernetesStatus_READY KubernetesStatus = "READY"
    // The cluster may be able to accept pods, but has warning messages. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     KubernetesStatus_WARNING KubernetesStatus = "WARNING"
    // The cluster may not be able to accept pods and has error messages. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     KubernetesStatus_ERROR KubernetesStatus = "ERROR"
)

func (k KubernetesStatus) KubernetesStatus() bool {
    switch k {
        case KubernetesStatus_READY:
            return true
        case KubernetesStatus_WARNING:
            return true
        case KubernetesStatus_ERROR:
            return true
        default:
            return false
    }
}




// Identifies the network plugin that cluster networking functionalities for this vSphere Namespaces Cluster. **Warning:** This enumeration is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type NetworkProvider string

const (
    // NSX-T Container Plugin. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     NetworkProvider_NSXT_CONTAINER_PLUGIN NetworkProvider = "NSXT_CONTAINER_PLUGIN"
)

func (n NetworkProvider) NetworkProvider() bool {
    switch n {
        case NetworkProvider_NSXT_CONTAINER_PLUGIN:
            return true
        default:
            return false
    }
}





// The ``Message`` class contains the information about the object configuration. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type Message struct {
    Severity Message_Severity
    Details *std.LocalizableMessage
}




    
    // The ``Severity`` enumeration class represents the severity of the message. **Warning:** This enumeration is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
     
    type Message_Severity string

    const (
        // Informational message. This may be accompanied by vCenter event. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
         Message_Severity_INFO Message_Severity = "INFO"
        // Warning message. This may be accompanied by vCenter event. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
         Message_Severity_WARNING Message_Severity = "WARNING"
        // Error message. This is accompanied by vCenter event and/or alarm. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
         Message_Severity_ERROR Message_Severity = "ERROR"
    )

    func (s Message_Severity) Message_Severity() bool {
        switch s {
            case Message_Severity_INFO:
                return true
            case Message_Severity_WARNING:
                return true
            case Message_Severity_ERROR:
                return true
            default:
                return false
        }
    }



// The ``Stats`` class contains the basic runtime statistics about a vSphere Namespaces-enabled cluster. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type Stats struct {
    CpuUsed int64
    CpuCapacity int64
    MemoryUsed int64
    MemoryCapacity int64
    StorageUsed int64
    StorageCapacity int64
}






// The ``Summary`` class contains the basic information about the cluster statistics and status related to vSphere Namespaces. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type Summary struct {
    Cluster string
    ClusterName string
    Stats Stats
    ConfigStatus ConfigStatus
    KubernetesStatus KubernetesStatus
}






// The ``Info`` class contains detailed information about the cluster statistics and status related to vSphere Namespaces. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type Info struct {
    StatSummary Stats
    ConfigStatus ConfigStatus
    Messages []Message
    KubernetesStatus KubernetesStatus
    KubernetesStatusMessages []Message
    ApiServerManagementEndpoint string
    ApiServerClusterEndpoint string
    ApiServers map[string]bool
    TlsManagementEndpointCertificate *string
    TlsEndpointCertificate *string
}






// The ``Ipv4Range`` contains specification to configure multiple interfaces in IPv4. The range of IPv4 addresses is derived by incrementing the startingAddress to the specified addressCount. To use the object for a single IPv4 address specification, set addressCount to 1. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type Ipv4Range struct {
    StartingAddress string
    AddressCount int64
    SubnetMask string
    Gateway string
}






// The ``NetworkSpec`` contains information related to network configuration for one or more interfaces. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type NetworkSpec struct {
    FloatingIP *string
    Network string
    Mode NetworkSpec_Ipv4Mode
    AddressRange *Ipv4Range
}




    
    // The ``Ipv4Mode`` enumeration class defines various IPv4 address assignment modes. **Warning:** This enumeration is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
     
    type NetworkSpec_Ipv4Mode string

    const (
        // The address is automatically assigned by a DHCP server. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
         NetworkSpec_Ipv4Mode_DHCP NetworkSpec_Ipv4Mode = "DHCP"
        // The address is static. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
         NetworkSpec_Ipv4Mode_STATICRANGE NetworkSpec_Ipv4Mode = "STATICRANGE"
    )

    func (i NetworkSpec_Ipv4Mode) NetworkSpec_Ipv4Mode() bool {
        switch i {
            case NetworkSpec_Ipv4Mode_DHCP:
                return true
            case NetworkSpec_Ipv4Mode_STATICRANGE:
                return true
            default:
                return false
        }
    }



// The ``ImageRegistry`` class contains the specification required to configure container image registry endpoint. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type ImageRegistry struct {
    Hostname string
    Port *int64
}






// The ``ImageStorageSpec`` class contains the specification required to configure storage used for container images. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type ImageStorageSpec struct {
    StoragePolicy string
}






// The ``NCPClusterNetworkEnableSpec`` class encapsulates the NSX Container Plugin-specific cluster networking configuration parameters for the vSphere Namespaces Cluster Enable operation. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type NCPClusterNetworkEnableSpec struct {
    PodCidrs []namespace_management.Ipv4Cidr
    IngressCidrs []namespace_management.Ipv4Cidr
    EgressCidrs []namespace_management.Ipv4Cidr
    ClusterDistributedSwitch string
    NsxEdgeCluster string
    SubnetSize *int64
}






// The ``NCPClusterNetworkUpdateSpec`` class encapsulates the NSX Container Plugin-specific cluster networking configuration parameters for the vSphere Namespaces Cluster Update operation. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type NCPClusterNetworkUpdateSpec struct {
    PodCidrs []namespace_management.Ipv4Cidr
    IngressCidrs []namespace_management.Ipv4Cidr
    EgressCidrs []namespace_management.Ipv4Cidr
    DefaultIngressTlsCertificate *string
}






// The ``NCPClusterNetworkSetSpec`` class encapsulates the NSX Container Plugin-specific cluster networking configuration parameters for the vSphere Namespaces Cluster Set operation. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type NCPClusterNetworkSetSpec struct {
    PodCidrs []namespace_management.Ipv4Cidr
    IngressCidrs []namespace_management.Ipv4Cidr
    EgressCidrs []namespace_management.Ipv4Cidr
    DefaultIngressTlsCertificate string
}






// The ``EnableSpec`` class contains the specification required to enable vSphere Namespaces on a cluster. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type EnableSpec struct {
    SizeHint namespace_management.SizingHint
    ServiceCidr namespace_management.Ipv4Cidr
    PodCidr *namespace_management.Ipv4Cidr
    NetworkProvider NetworkProvider
    NcpClusterNetworkSpec *NCPClusterNetworkEnableSpec
    MasterManagementNetwork NetworkSpec
    MasterClusterNetwork NetworkSpec
    WorkerClusterNetwork NetworkSpec
    MasterAdditionalNetworks []NetworkSpec
    WorkerAdditionalNetworks []NetworkSpec
    MasterDNS []string
    WorkerDNS []string
    MasterDNSSearchDomains []string
    WorkerDNSSearchDomains []string
    MasterNTPServers []string
    MasterStoragePolicy string
    EphemeralStoragePolicy string
    LoginBanner *string
    MasterDNSNames []string
    ImageStorage ImageStorageSpec
    DefaultImageRegistry *ImageRegistry
    DefaultImageRepository *string
}






// The ``UpdateSpec`` class contains the specification required to update the configuration on the Cluster. This class is applied partially, and only the specified fields will replace or modify their existing counterparts. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type UpdateSpec struct {
    SizeHint *namespace_management.SizingHint
    ServiceCidr *namespace_management.Ipv4Cidr
    PodCidr *namespace_management.Ipv4Cidr
    NetworkProvider *NetworkProvider
    NcpClusterNetworkSpec *NCPClusterNetworkUpdateSpec
    MasterDNS []string
    WorkerDNS []string
    MasterDNSSearchDomains []string
    WorkerDNSSearchDomains []string
    MasterNTPServers []string
    MasterStoragePolicy *string
    EphemeralStoragePolicy *string
    LoginBanner *string
    ImageStorage *ImageStorageSpec
    DefaultImageRegistry *ImageRegistry
    DefaultImageRepository *string
    TlsEndpointCertificate *string
}






// The ``SetSpec`` class contains the specification required to set a new configuration on the Cluster. This class is applied in entirety, replacing the current specification fully. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type SetSpec struct {
    SizeHint namespace_management.SizingHint
    ServiceCidr namespace_management.Ipv4Cidr
    PodCidr *namespace_management.Ipv4Cidr
    NetworkProvider NetworkProvider
    NcpClusterNetworkSpec *NCPClusterNetworkSetSpec
    MasterDNS []string
    WorkerDNS []string
    MasterDNSSearchDomains []string
    WorkerDNSSearchDomains []string
    MasterNTPServers []string
    MasterStoragePolicy string
    EphemeralStoragePolicy string
    LoginBanner *string
    ImageStorage ImageStorageSpec
    DefaultImageRegistry *ImageRegistry
    DefaultImageRepository *string
}









func enableInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    fields["spec"] = bindings.NewReferenceType(EnableSpecBindingType)
    fieldNameMap["cluster"] = "Cluster"
    fieldNameMap["spec"] = "Spec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func enableOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func enableRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    paramsTypeMap["spec"] = bindings.NewReferenceType(EnableSpecBindingType)
    paramsTypeMap["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    pathParams["cluster"] = "cluster"
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "spec",
      "POST",
      "/vcenter/namespace-management/clusters/{cluster}",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"AlreadyExists": 400,"Error": 500,"NotFound": 404,"InvalidArgument": 400,"Unauthenticated": 401,"Unauthorized": 403})
}


func disableInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    fieldNameMap["cluster"] = "Cluster"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func disableOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func disableRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    paramsTypeMap["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    pathParams["cluster"] = "cluster"
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "POST",
      "/vcenter/namespace-management/clusters/{cluster}",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"Error": 500,"NotFound": 404,"Unauthenticated": 401,"Unauthorized": 403})
}


func getInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    fieldNameMap["cluster"] = "Cluster"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func getOutputType() bindings.BindingType {
    return bindings.NewReferenceType(InfoBindingType)
}

func getRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    paramsTypeMap["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    pathParams["cluster"] = "cluster"
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "GET",
      "/vcenter/namespace-management/clusters/{cluster}",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"NotFound": 404,"Unsupported": 400,"Error": 500,"Unauthenticated": 401,"Unauthorized": 403})
}


func listInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func listOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewReferenceType(SummaryBindingType), reflect.TypeOf([]Summary{}))
}

func listRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "GET",
      "/vcenter/namespace-management/clusters",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"Error": 500,"Unauthenticated": 401,"Unauthorized": 403})
}


func setInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    fields["spec"] = bindings.NewReferenceType(SetSpecBindingType)
    fieldNameMap["cluster"] = "Cluster"
    fieldNameMap["spec"] = "Spec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func setOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func setRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    paramsTypeMap["spec"] = bindings.NewReferenceType(SetSpecBindingType)
    paramsTypeMap["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    pathParams["cluster"] = "cluster"
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "spec",
      "PUT",
      "/vcenter/namespace-management/clusters/{cluster}",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"Error": 500,"InvalidArgument": 400,"NotAllowedInCurrentState": 400,"NotFound": 404,"Unauthenticated": 401,"Unauthorized": 403})
}


func updateInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    fields["spec"] = bindings.NewReferenceType(UpdateSpecBindingType)
    fieldNameMap["cluster"] = "Cluster"
    fieldNameMap["spec"] = "Spec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func updateOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func updateRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    paramsTypeMap["spec"] = bindings.NewReferenceType(UpdateSpecBindingType)
    paramsTypeMap["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    pathParams["cluster"] = "cluster"
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "spec",
      "PATCH",
      "/vcenter/namespace-management/clusters/{cluster}",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"Error": 500,"InvalidArgument": 400,"NotAllowedInCurrentState": 400,"NotFound": 404,"Unauthenticated": 401,"Unauthorized": 403})
}



func MessageBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["severity"] = bindings.NewEnumType("com.vmware.vcenter.namespace_management.clusters.message.severity", reflect.TypeOf(Message_Severity(Message_Severity_INFO)))
    fieldNameMap["severity"] = "Severity"
    fields["details"] = bindings.NewOptionalType(bindings.NewReferenceType(std.LocalizableMessageBindingType))
    fieldNameMap["details"] = "Details"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.namespace_management.clusters.message",fields, reflect.TypeOf(Message{}), fieldNameMap, validators)
}

func StatsBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["cpu_used"] = bindings.NewIntegerType()
    fieldNameMap["cpu_used"] = "CpuUsed"
    fields["cpu_capacity"] = bindings.NewIntegerType()
    fieldNameMap["cpu_capacity"] = "CpuCapacity"
    fields["memory_used"] = bindings.NewIntegerType()
    fieldNameMap["memory_used"] = "MemoryUsed"
    fields["memory_capacity"] = bindings.NewIntegerType()
    fieldNameMap["memory_capacity"] = "MemoryCapacity"
    fields["storage_used"] = bindings.NewIntegerType()
    fieldNameMap["storage_used"] = "StorageUsed"
    fields["storage_capacity"] = bindings.NewIntegerType()
    fieldNameMap["storage_capacity"] = "StorageCapacity"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.namespace_management.clusters.stats",fields, reflect.TypeOf(Stats{}), fieldNameMap, validators)
}

func SummaryBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    fieldNameMap["cluster"] = "Cluster"
    fields["cluster_name"] = bindings.NewIdType([]string {"ClusterComputeResource.name"}, "")
    fieldNameMap["cluster_name"] = "ClusterName"
    fields["stats"] = bindings.NewReferenceType(StatsBindingType)
    fieldNameMap["stats"] = "Stats"
    fields["config_status"] = bindings.NewEnumType("com.vmware.vcenter.namespace_management.clusters.config_status", reflect.TypeOf(ConfigStatus(ConfigStatus_CONFIGURING)))
    fieldNameMap["config_status"] = "ConfigStatus"
    fields["kubernetes_status"] = bindings.NewEnumType("com.vmware.vcenter.namespace_management.clusters.kubernetes_status", reflect.TypeOf(KubernetesStatus(KubernetesStatus_READY)))
    fieldNameMap["kubernetes_status"] = "KubernetesStatus"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.namespace_management.clusters.summary",fields, reflect.TypeOf(Summary{}), fieldNameMap, validators)
}

func InfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["stat_summary"] = bindings.NewReferenceType(StatsBindingType)
    fieldNameMap["stat_summary"] = "StatSummary"
    fields["config_status"] = bindings.NewEnumType("com.vmware.vcenter.namespace_management.clusters.config_status", reflect.TypeOf(ConfigStatus(ConfigStatus_CONFIGURING)))
    fieldNameMap["config_status"] = "ConfigStatus"
    fields["messages"] = bindings.NewListType(bindings.NewReferenceType(MessageBindingType), reflect.TypeOf([]Message{}))
    fieldNameMap["messages"] = "Messages"
    fields["kubernetes_status"] = bindings.NewEnumType("com.vmware.vcenter.namespace_management.clusters.kubernetes_status", reflect.TypeOf(KubernetesStatus(KubernetesStatus_READY)))
    fieldNameMap["kubernetes_status"] = "KubernetesStatus"
    fields["kubernetes_status_messages"] = bindings.NewListType(bindings.NewReferenceType(MessageBindingType), reflect.TypeOf([]Message{}))
    fieldNameMap["kubernetes_status_messages"] = "KubernetesStatusMessages"
    fields["api_server_management_endpoint"] = bindings.NewStringType()
    fieldNameMap["api_server_management_endpoint"] = "ApiServerManagementEndpoint"
    fields["api_server_cluster_endpoint"] = bindings.NewStringType()
    fieldNameMap["api_server_cluster_endpoint"] = "ApiServerClusterEndpoint"
    fields["api_servers"] = bindings.NewSetType(bindings.NewStringType(), reflect.TypeOf(map[string]bool{}))
    fieldNameMap["api_servers"] = "ApiServers"
    fields["tls_management_endpoint_certificate"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["tls_management_endpoint_certificate"] = "TlsManagementEndpointCertificate"
    fields["tls_endpoint_certificate"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["tls_endpoint_certificate"] = "TlsEndpointCertificate"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.namespace_management.clusters.info",fields, reflect.TypeOf(Info{}), fieldNameMap, validators)
}

func Ipv4RangeBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["starting_address"] = bindings.NewStringType()
    fieldNameMap["starting_address"] = "StartingAddress"
    fields["address_count"] = bindings.NewIntegerType()
    fieldNameMap["address_count"] = "AddressCount"
    fields["subnet_mask"] = bindings.NewStringType()
    fieldNameMap["subnet_mask"] = "SubnetMask"
    fields["gateway"] = bindings.NewStringType()
    fieldNameMap["gateway"] = "Gateway"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.namespace_management.clusters.ipv4_range",fields, reflect.TypeOf(Ipv4Range{}), fieldNameMap, validators)
}

func NetworkSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["floating_IP"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["floating_IP"] = "FloatingIP"
    fields["network"] = bindings.NewIdType([]string {"Network"}, "")
    fieldNameMap["network"] = "Network"
    fields["mode"] = bindings.NewEnumType("com.vmware.vcenter.namespace_management.clusters.network_spec.ipv4_mode", reflect.TypeOf(NetworkSpec_Ipv4Mode(NetworkSpec_Ipv4Mode_DHCP)))
    fieldNameMap["mode"] = "Mode"
    fields["address_range"] = bindings.NewOptionalType(bindings.NewReferenceType(Ipv4RangeBindingType))
    fieldNameMap["address_range"] = "AddressRange"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("mode",
        map[string][]bindings.FieldData {
            "DHCP": []bindings.FieldData {
                 bindings.NewFieldData("floating_IP", false),
            },
            "STATICRANGE": []bindings.FieldData {
                 bindings.NewFieldData("address_range", true),
            },
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.vcenter.namespace_management.clusters.network_spec",fields, reflect.TypeOf(NetworkSpec{}), fieldNameMap, validators)
}

func ImageRegistryBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["hostname"] = bindings.NewStringType()
    fieldNameMap["hostname"] = "Hostname"
    fields["port"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["port"] = "Port"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.namespace_management.clusters.image_registry",fields, reflect.TypeOf(ImageRegistry{}), fieldNameMap, validators)
}

func ImageStorageSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["storage_policy"] = bindings.NewIdType([]string {"SpsStorageProfile"}, "")
    fieldNameMap["storage_policy"] = "StoragePolicy"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.namespace_management.clusters.image_storage_spec",fields, reflect.TypeOf(ImageStorageSpec{}), fieldNameMap, validators)
}

func NCPClusterNetworkEnableSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["pod_cidrs"] = bindings.NewListType(bindings.NewReferenceType(namespace_management.Ipv4CidrBindingType), reflect.TypeOf([]namespace_management.Ipv4Cidr{}))
    fieldNameMap["pod_cidrs"] = "PodCidrs"
    fields["ingress_cidrs"] = bindings.NewListType(bindings.NewReferenceType(namespace_management.Ipv4CidrBindingType), reflect.TypeOf([]namespace_management.Ipv4Cidr{}))
    fieldNameMap["ingress_cidrs"] = "IngressCidrs"
    fields["egress_cidrs"] = bindings.NewListType(bindings.NewReferenceType(namespace_management.Ipv4CidrBindingType), reflect.TypeOf([]namespace_management.Ipv4Cidr{}))
    fieldNameMap["egress_cidrs"] = "EgressCidrs"
    fields["cluster_distributed_switch"] = bindings.NewIdType([]string {"vSphereDistributedSwitch"}, "")
    fieldNameMap["cluster_distributed_switch"] = "ClusterDistributedSwitch"
    fields["nsx_edge_cluster"] = bindings.NewIdType([]string {"NSXEdgeCluster"}, "")
    fieldNameMap["nsx_edge_cluster"] = "NsxEdgeCluster"
    fields["subnet_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["subnet_size"] = "SubnetSize"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.namespace_management.clusters.NCP_cluster_network_enable_spec",fields, reflect.TypeOf(NCPClusterNetworkEnableSpec{}), fieldNameMap, validators)
}

func NCPClusterNetworkUpdateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["pod_cidrs"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(namespace_management.Ipv4CidrBindingType), reflect.TypeOf([]namespace_management.Ipv4Cidr{})))
    fieldNameMap["pod_cidrs"] = "PodCidrs"
    fields["ingress_cidrs"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(namespace_management.Ipv4CidrBindingType), reflect.TypeOf([]namespace_management.Ipv4Cidr{})))
    fieldNameMap["ingress_cidrs"] = "IngressCidrs"
    fields["egress_cidrs"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(namespace_management.Ipv4CidrBindingType), reflect.TypeOf([]namespace_management.Ipv4Cidr{})))
    fieldNameMap["egress_cidrs"] = "EgressCidrs"
    fields["default_ingress_tls_certificate"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["default_ingress_tls_certificate"] = "DefaultIngressTlsCertificate"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.namespace_management.clusters.NCP_cluster_network_update_spec",fields, reflect.TypeOf(NCPClusterNetworkUpdateSpec{}), fieldNameMap, validators)
}

func NCPClusterNetworkSetSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["pod_cidrs"] = bindings.NewListType(bindings.NewReferenceType(namespace_management.Ipv4CidrBindingType), reflect.TypeOf([]namespace_management.Ipv4Cidr{}))
    fieldNameMap["pod_cidrs"] = "PodCidrs"
    fields["ingress_cidrs"] = bindings.NewListType(bindings.NewReferenceType(namespace_management.Ipv4CidrBindingType), reflect.TypeOf([]namespace_management.Ipv4Cidr{}))
    fieldNameMap["ingress_cidrs"] = "IngressCidrs"
    fields["egress_cidrs"] = bindings.NewListType(bindings.NewReferenceType(namespace_management.Ipv4CidrBindingType), reflect.TypeOf([]namespace_management.Ipv4Cidr{}))
    fieldNameMap["egress_cidrs"] = "EgressCidrs"
    fields["default_ingress_tls_certificate"] = bindings.NewStringType()
    fieldNameMap["default_ingress_tls_certificate"] = "DefaultIngressTlsCertificate"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.namespace_management.clusters.NCP_cluster_network_set_spec",fields, reflect.TypeOf(NCPClusterNetworkSetSpec{}), fieldNameMap, validators)
}

func EnableSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["size_hint"] = bindings.NewEnumType("com.vmware.vcenter.namespace_management.sizing_hint", reflect.TypeOf(namespace_management.SizingHint(namespace_management.SizingHint_TINY)))
    fieldNameMap["size_hint"] = "SizeHint"
    fields["service_cidr"] = bindings.NewReferenceType(namespace_management.Ipv4CidrBindingType)
    fieldNameMap["service_cidr"] = "ServiceCidr"
    fields["pod_cidr"] = bindings.NewOptionalType(bindings.NewReferenceType(namespace_management.Ipv4CidrBindingType))
    fieldNameMap["pod_cidr"] = "PodCidr"
    fields["network_provider"] = bindings.NewEnumType("com.vmware.vcenter.namespace_management.clusters.network_provider", reflect.TypeOf(NetworkProvider(NetworkProvider_NSXT_CONTAINER_PLUGIN)))
    fieldNameMap["network_provider"] = "NetworkProvider"
    fields["ncp_cluster_network_spec"] = bindings.NewOptionalType(bindings.NewReferenceType(NCPClusterNetworkEnableSpecBindingType))
    fieldNameMap["ncp_cluster_network_spec"] = "NcpClusterNetworkSpec"
    fields["master_management_network"] = bindings.NewReferenceType(NetworkSpecBindingType)
    fieldNameMap["master_management_network"] = "MasterManagementNetwork"
    fields["master_cluster_network"] = bindings.NewReferenceType(NetworkSpecBindingType)
    fieldNameMap["master_cluster_network"] = "MasterClusterNetwork"
    fields["worker_cluster_network"] = bindings.NewReferenceType(NetworkSpecBindingType)
    fieldNameMap["worker_cluster_network"] = "WorkerClusterNetwork"
    fields["master_additional_networks"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(NetworkSpecBindingType), reflect.TypeOf([]NetworkSpec{})))
    fieldNameMap["master_additional_networks"] = "MasterAdditionalNetworks"
    fields["worker_additional_networks"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(NetworkSpecBindingType), reflect.TypeOf([]NetworkSpec{})))
    fieldNameMap["worker_additional_networks"] = "WorkerAdditionalNetworks"
    fields["master_DNS"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
    fieldNameMap["master_DNS"] = "MasterDNS"
    fields["worker_DNS"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
    fieldNameMap["worker_DNS"] = "WorkerDNS"
    fields["master_DNS_search_domains"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
    fieldNameMap["master_DNS_search_domains"] = "MasterDNSSearchDomains"
    fields["worker_DNS_search_domains"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
    fieldNameMap["worker_DNS_search_domains"] = "WorkerDNSSearchDomains"
    fields["master_NTP_servers"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
    fieldNameMap["master_NTP_servers"] = "MasterNTPServers"
    fields["master_storage_policy"] = bindings.NewIdType([]string {"SpsStorageProfile"}, "")
    fieldNameMap["master_storage_policy"] = "MasterStoragePolicy"
    fields["ephemeral_storage_policy"] = bindings.NewIdType([]string {"SpsStorageProfile"}, "")
    fieldNameMap["ephemeral_storage_policy"] = "EphemeralStoragePolicy"
    fields["login_banner"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["login_banner"] = "LoginBanner"
    fields["Master_DNS_names"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
    fieldNameMap["Master_DNS_names"] = "MasterDNSNames"
    fields["image_storage"] = bindings.NewReferenceType(ImageStorageSpecBindingType)
    fieldNameMap["image_storage"] = "ImageStorage"
    fields["default_image_registry"] = bindings.NewOptionalType(bindings.NewReferenceType(ImageRegistryBindingType))
    fieldNameMap["default_image_registry"] = "DefaultImageRegistry"
    fields["default_image_repository"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["default_image_repository"] = "DefaultImageRepository"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("network_provider",
        map[string][]bindings.FieldData {
            "NSXT_CONTAINER_PLUGIN": []bindings.FieldData {
                 bindings.NewFieldData("ncp_cluster_network_spec", true),
            },
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.vcenter.namespace_management.clusters.enable_spec",fields, reflect.TypeOf(EnableSpec{}), fieldNameMap, validators)
}

func UpdateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["size_hint"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.namespace_management.sizing_hint", reflect.TypeOf(namespace_management.SizingHint(namespace_management.SizingHint_TINY))))
    fieldNameMap["size_hint"] = "SizeHint"
    fields["service_cidr"] = bindings.NewOptionalType(bindings.NewReferenceType(namespace_management.Ipv4CidrBindingType))
    fieldNameMap["service_cidr"] = "ServiceCidr"
    fields["pod_cidr"] = bindings.NewOptionalType(bindings.NewReferenceType(namespace_management.Ipv4CidrBindingType))
    fieldNameMap["pod_cidr"] = "PodCidr"
    fields["network_provider"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.namespace_management.clusters.network_provider", reflect.TypeOf(NetworkProvider(NetworkProvider_NSXT_CONTAINER_PLUGIN))))
    fieldNameMap["network_provider"] = "NetworkProvider"
    fields["ncp_cluster_network_spec"] = bindings.NewOptionalType(bindings.NewReferenceType(NCPClusterNetworkUpdateSpecBindingType))
    fieldNameMap["ncp_cluster_network_spec"] = "NcpClusterNetworkSpec"
    fields["master_DNS"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
    fieldNameMap["master_DNS"] = "MasterDNS"
    fields["worker_DNS"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
    fieldNameMap["worker_DNS"] = "WorkerDNS"
    fields["master_DNS_search_domains"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
    fieldNameMap["master_DNS_search_domains"] = "MasterDNSSearchDomains"
    fields["worker_DNS_search_domains"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
    fieldNameMap["worker_DNS_search_domains"] = "WorkerDNSSearchDomains"
    fields["master_NTP_servers"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
    fieldNameMap["master_NTP_servers"] = "MasterNTPServers"
    fields["master_storage_policy"] = bindings.NewOptionalType(bindings.NewIdType([]string {"SpsStorageProfile"}, ""))
    fieldNameMap["master_storage_policy"] = "MasterStoragePolicy"
    fields["ephemeral_storage_policy"] = bindings.NewOptionalType(bindings.NewIdType([]string {"SpsStorageProfile"}, ""))
    fieldNameMap["ephemeral_storage_policy"] = "EphemeralStoragePolicy"
    fields["login_banner"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["login_banner"] = "LoginBanner"
    fields["image_storage"] = bindings.NewOptionalType(bindings.NewReferenceType(ImageStorageSpecBindingType))
    fieldNameMap["image_storage"] = "ImageStorage"
    fields["default_image_registry"] = bindings.NewOptionalType(bindings.NewReferenceType(ImageRegistryBindingType))
    fieldNameMap["default_image_registry"] = "DefaultImageRegistry"
    fields["default_image_repository"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["default_image_repository"] = "DefaultImageRepository"
    fields["tls_endpoint_certificate"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["tls_endpoint_certificate"] = "TlsEndpointCertificate"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("network_provider",
        map[string][]bindings.FieldData {
            "NSXT_CONTAINER_PLUGIN": []bindings.FieldData {
                 bindings.NewFieldData("ncp_cluster_network_spec", false),
            },
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.vcenter.namespace_management.clusters.update_spec",fields, reflect.TypeOf(UpdateSpec{}), fieldNameMap, validators)
}

func SetSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["size_hint"] = bindings.NewEnumType("com.vmware.vcenter.namespace_management.sizing_hint", reflect.TypeOf(namespace_management.SizingHint(namespace_management.SizingHint_TINY)))
    fieldNameMap["size_hint"] = "SizeHint"
    fields["service_cidr"] = bindings.NewReferenceType(namespace_management.Ipv4CidrBindingType)
    fieldNameMap["service_cidr"] = "ServiceCidr"
    fields["pod_cidr"] = bindings.NewOptionalType(bindings.NewReferenceType(namespace_management.Ipv4CidrBindingType))
    fieldNameMap["pod_cidr"] = "PodCidr"
    fields["network_provider"] = bindings.NewEnumType("com.vmware.vcenter.namespace_management.clusters.network_provider", reflect.TypeOf(NetworkProvider(NetworkProvider_NSXT_CONTAINER_PLUGIN)))
    fieldNameMap["network_provider"] = "NetworkProvider"
    fields["ncp_cluster_network_spec"] = bindings.NewOptionalType(bindings.NewReferenceType(NCPClusterNetworkSetSpecBindingType))
    fieldNameMap["ncp_cluster_network_spec"] = "NcpClusterNetworkSpec"
    fields["master_DNS"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
    fieldNameMap["master_DNS"] = "MasterDNS"
    fields["worker_DNS"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
    fieldNameMap["worker_DNS"] = "WorkerDNS"
    fields["master_DNS_search_domains"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
    fieldNameMap["master_DNS_search_domains"] = "MasterDNSSearchDomains"
    fields["worker_DNS_search_domains"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
    fieldNameMap["worker_DNS_search_domains"] = "WorkerDNSSearchDomains"
    fields["master_NTP_servers"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
    fieldNameMap["master_NTP_servers"] = "MasterNTPServers"
    fields["master_storage_policy"] = bindings.NewIdType([]string {"SpsStorageProfile"}, "")
    fieldNameMap["master_storage_policy"] = "MasterStoragePolicy"
    fields["ephemeral_storage_policy"] = bindings.NewIdType([]string {"SpsStorageProfile"}, "")
    fieldNameMap["ephemeral_storage_policy"] = "EphemeralStoragePolicy"
    fields["login_banner"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["login_banner"] = "LoginBanner"
    fields["image_storage"] = bindings.NewReferenceType(ImageStorageSpecBindingType)
    fieldNameMap["image_storage"] = "ImageStorage"
    fields["default_image_registry"] = bindings.NewOptionalType(bindings.NewReferenceType(ImageRegistryBindingType))
    fieldNameMap["default_image_registry"] = "DefaultImageRegistry"
    fields["default_image_repository"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["default_image_repository"] = "DefaultImageRepository"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("network_provider",
        map[string][]bindings.FieldData {
            "NSXT_CONTAINER_PLUGIN": []bindings.FieldData {
                 bindings.NewFieldData("ncp_cluster_network_spec", true),
            },
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.vcenter.namespace_management.clusters.set_spec",fields, reflect.TypeOf(SetSpec{}), fieldNameMap, validators)
}


