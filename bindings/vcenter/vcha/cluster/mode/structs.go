/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Mode.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package mode

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)



// The ``ClusterMode`` enumeration class defines the possible modes for a VCHA Cluster.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type ClusterMode string

const (
    // VCHA Cluster is enabled. State replication between the Active and Passive node is enabled and automatic failover is allowed.
     ClusterMode_ENABLED ClusterMode = "ENABLED"
    // VCHA Cluster is disabled. State replication between the Active and Passive node is disabled and automatic failover is not allowed.
     ClusterMode_DISABLED ClusterMode = "DISABLED"
    // VCHA Cluster is in maintenance mode. State replication between the and Passive node is enabled but automatic failover is not allowed.
     ClusterMode_MAINTENANCE ClusterMode = "MAINTENANCE"
)

func (c ClusterMode) ClusterMode() bool {
    switch c {
        case ClusterMode_ENABLED:
            return true
        case ClusterMode_DISABLED:
            return true
        case ClusterMode_MAINTENANCE:
            return true
        default:
            return false
    }
}





// The ``Info`` class contains the mode of the VCHA Cluster.
 type Info struct {
    Mode ClusterMode
}









func getInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
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
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "null",
      "",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"NotAllowedInCurrentState": 400,"Unauthorized": 403,"Error": 500})
}


func setInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["mode"] = bindings.NewEnumType("com.vmware.vcenter.vcha.cluster.mode.cluster_mode", reflect.TypeOf(ClusterMode(ClusterMode_ENABLED)))
    fieldNameMap["mode"] = "Mode"
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
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "null",
      "",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"Unauthorized": 403,"Error": 500})
}



func InfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["mode"] = bindings.NewEnumType("com.vmware.vcenter.vcha.cluster.mode.cluster_mode", reflect.TypeOf(ClusterMode(ClusterMode_ENABLED)))
    fieldNameMap["mode"] = "Mode"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vcha.cluster.mode.info",fields, reflect.TypeOf(Info{}), fieldNameMap, validators)
}


