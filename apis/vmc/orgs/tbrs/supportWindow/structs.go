/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: SupportWindow.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package supportWindow

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
    fields["minimum_seats_available"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fields["created_by"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["org"] = "Org"
    fieldNameMap["minimum_seats_available"] = "MinimumSeatsAvailable"
    fieldNameMap["created_by"] = "CreatedBy"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func getOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewReferenceType(model.SupportWindowBindingType), reflect.TypeOf([]model.SupportWindow{}))
}

func getRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["created_by"] = bindings.NewOptionalType(bindings.NewStringType())
    paramsTypeMap["org"] = bindings.NewStringType()
    paramsTypeMap["minimum_seats_available"] = bindings.NewOptionalType(bindings.NewIntegerType())
    paramsTypeMap["org"] = bindings.NewStringType()
    pathParams["org"] = "org"
    queryParams["created_by"] = "createdBy"
    queryParams["minimum_seats_available"] = "minimumSeatsAvailable"
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
      "/vmc/api/orgs/{org}/tbrs/support-window",
       resultHeaders,
       200,
       errorHeaders,
       map[string]int{"Unauthenticated": 401,"InvalidRequest": 400,"Unauthorized": 403,"NotFound": 404})
}


func putInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["org"] = bindings.NewStringType()
    fields["id"] = bindings.NewStringType()
    fields["sddc_id"] = bindings.NewReferenceType(model.SddcIdBindingType)
    fieldNameMap["org"] = "Org"
    fieldNameMap["id"] = "Id"
    fieldNameMap["sddc_id"] = "SddcId"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func putOutputType() bindings.BindingType {
    return bindings.NewReferenceType(model.SupportWindowIdBindingType)
}

func putRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["org"] = bindings.NewStringType()
    paramsTypeMap["id"] = bindings.NewStringType()
    paramsTypeMap["sddc_id"] = bindings.NewReferenceType(model.SddcIdBindingType)
    paramsTypeMap["org"] = bindings.NewStringType()
    paramsTypeMap["id"] = bindings.NewStringType()
    pathParams["org"] = "org"
    pathParams["id"] = "id"
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "sddc_id",
      "PUT",
      "/vmc/api/orgs/{org}/tbrs/support-window/{id}",
       resultHeaders,
       200,
       errorHeaders,
       map[string]int{"Unauthenticated": 401,"InvalidRequest": 400,"Unauthorized": 403,"NotFound": 404})
}




