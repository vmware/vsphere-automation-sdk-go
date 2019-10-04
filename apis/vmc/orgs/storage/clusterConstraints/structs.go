/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: ClusterConstraints.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package clusterConstraints

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/apis/vmc/model"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)







func getInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["org"] = bindings.NewStringType()
    fields["provider"] = bindings.NewStringType()
    fields["num_hosts"] = bindings.NewIntegerType()
    fields["one_node_reduced_capacity"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["org"] = "Org"
    fieldNameMap["provider"] = "Provider"
    fieldNameMap["num_hosts"] = "NumHosts"
    fieldNameMap["one_node_reduced_capacity"] = "OneNodeReducedCapacity"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func getOutputType() bindings.BindingType {
    return bindings.NewReferenceType(model.VsanConfigConstraintsBindingType)
}

func getRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["provider"] = bindings.NewStringType()
    paramsTypeMap["org"] = bindings.NewStringType()
    paramsTypeMap["num_hosts"] = bindings.NewIntegerType()
    paramsTypeMap["one_node_reduced_capacity"] = bindings.NewOptionalType(bindings.NewBooleanType())
    paramsTypeMap["org"] = bindings.NewStringType()
    pathParams["org"] = "org"
    queryParams["num_hosts"] = "num_hosts"
    queryParams["one_node_reduced_capacity"] = "one_node_reduced_capacity"
    queryParams["provider"] = "provider"
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
      "/vmc/api/orgs/{org}/storage/cluster-constraints",
       resultHeaders,
       200,
       errorHeaders,
       map[string]int{"Unauthenticated": 401,"InvalidRequest": 400,"Unauthorized": 403})
}




