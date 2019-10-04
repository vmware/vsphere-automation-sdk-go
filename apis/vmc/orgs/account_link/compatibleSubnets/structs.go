/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: CompatibleSubnets.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package compatibleSubnets

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
    fields["linked_account_id"] = bindings.NewOptionalType(bindings.NewStringType())
    fields["region"] = bindings.NewOptionalType(bindings.NewStringType())
    fields["sddc"] = bindings.NewOptionalType(bindings.NewStringType())
    fields["force_refresh"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fields["instance_type"] = bindings.NewOptionalType(bindings.NewStringType())
    fields["sddc_type"] = bindings.NewOptionalType(bindings.NewStringType())
    fields["num_of_hosts"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["org"] = "Org"
    fieldNameMap["linked_account_id"] = "LinkedAccountId"
    fieldNameMap["region"] = "Region"
    fieldNameMap["sddc"] = "Sddc"
    fieldNameMap["force_refresh"] = "ForceRefresh"
    fieldNameMap["instance_type"] = "InstanceType"
    fieldNameMap["sddc_type"] = "SddcType"
    fieldNameMap["num_of_hosts"] = "NumOfHosts"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func getOutputType() bindings.BindingType {
    return bindings.NewReferenceType(model.AwsCompatibleSubnetsBindingType)
}

func getRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["sddc"] = bindings.NewOptionalType(bindings.NewStringType())
    paramsTypeMap["instance_type"] = bindings.NewOptionalType(bindings.NewStringType())
    paramsTypeMap["region"] = bindings.NewOptionalType(bindings.NewStringType())
    paramsTypeMap["linked_account_id"] = bindings.NewOptionalType(bindings.NewStringType())
    paramsTypeMap["org"] = bindings.NewStringType()
    paramsTypeMap["force_refresh"] = bindings.NewOptionalType(bindings.NewBooleanType())
    paramsTypeMap["num_of_hosts"] = bindings.NewOptionalType(bindings.NewIntegerType())
    paramsTypeMap["sddc_type"] = bindings.NewOptionalType(bindings.NewStringType())
    paramsTypeMap["org"] = bindings.NewStringType()
    pathParams["org"] = "org"
    queryParams["linked_account_id"] = "linkedAccountId"
    queryParams["num_of_hosts"] = "numOfHosts"
    queryParams["sddc"] = "sddc"
    queryParams["instance_type"] = "instanceType"
    queryParams["region"] = "region"
    queryParams["sddc_type"] = "sddcType"
    queryParams["force_refresh"] = "forceRefresh"
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
      "/vmc/api/orgs/{org}/account-link/compatible-subnets",
       resultHeaders,
       200,
       errorHeaders,
       map[string]int{"Unauthenticated": 401,"Unauthorized": 403})
}


func postInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["org"] = bindings.NewStringType()
    fieldNameMap["org"] = "Org"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func postOutputType() bindings.BindingType {
    return bindings.NewReferenceType(model.AwsSubnetBindingType)
}

func postRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["org"] = bindings.NewStringType()
    paramsTypeMap["org"] = bindings.NewStringType()
    pathParams["org"] = "org"
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
      "/vmc/api/orgs/{org}/account-link/compatible-subnets",
       resultHeaders,
       200,
       errorHeaders,
       map[string]int{"Unauthenticated": 401,"Unauthorized": 403})
}




