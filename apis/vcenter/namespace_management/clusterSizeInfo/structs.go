/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: ClusterSizeInfo.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package clusterSizeInfo

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/apis/vcenter/namespace_management"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)




// The ``VmInfo`` class contains the information about the configuration of the virtual machines which would be created for Namespaces. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type VmInfo struct {
    Count int64
    CoresPerSocket int64
    Memory int64
    Capacity int64
}






// The ``Info`` class contains the information about limits associated with a ``SizingHint``. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type Info struct {
    NumSupportedPods int64
    NumSupportedServices int64
    DefaultServiceCidr namespace_management.Ipv4Cidr
    DefaultPodCidr namespace_management.Ipv4Cidr
    MasterVmInfo VmInfo
    WorkerVmInfo *VmInfo
}









func getInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func getOutputType() bindings.BindingType {
    return bindings.NewMapType(bindings.NewEnumType("com.vmware.vcenter.namespace_management.sizing_hint", reflect.TypeOf(namespace_management.SizingHint(namespace_management.SizingHint_TINY))), bindings.NewReferenceType(InfoBindingType),reflect.TypeOf(map[namespace_management.SizingHint]Info{}))
}

func getRestMetadata() protocol.OperationRestMetadata {
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
      "/vcenter/namespace-management/cluster-size-info",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"Error": 500,"Unauthenticated": 401})
}



func VmInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["count"] = bindings.NewIntegerType()
    fieldNameMap["count"] = "Count"
    fields["cores_per_socket"] = bindings.NewIntegerType()
    fieldNameMap["cores_per_socket"] = "CoresPerSocket"
    fields["memory"] = bindings.NewIntegerType()
    fieldNameMap["memory"] = "Memory"
    fields["capacity"] = bindings.NewIntegerType()
    fieldNameMap["capacity"] = "Capacity"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.namespace_management.cluster_size_info.vm_info",fields, reflect.TypeOf(VmInfo{}), fieldNameMap, validators)
}

func InfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["num_supported_pods"] = bindings.NewIntegerType()
    fieldNameMap["num_supported_pods"] = "NumSupportedPods"
    fields["num_supported_services"] = bindings.NewIntegerType()
    fieldNameMap["num_supported_services"] = "NumSupportedServices"
    fields["default_service_cidr"] = bindings.NewReferenceType(namespace_management.Ipv4CidrBindingType)
    fieldNameMap["default_service_cidr"] = "DefaultServiceCidr"
    fields["default_pod_cidr"] = bindings.NewReferenceType(namespace_management.Ipv4CidrBindingType)
    fieldNameMap["default_pod_cidr"] = "DefaultPodCidr"
    fields["master_vm_info"] = bindings.NewReferenceType(VmInfoBindingType)
    fieldNameMap["master_vm_info"] = "MasterVmInfo"
    fields["worker_vm_info"] = bindings.NewOptionalType(bindings.NewReferenceType(VmInfoBindingType))
    fieldNameMap["worker_vm_info"] = "WorkerVmInfo"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.namespace_management.cluster_size_info.info",fields, reflect.TypeOf(Info{}), fieldNameMap, validators)
}


