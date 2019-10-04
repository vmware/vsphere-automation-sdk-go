/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Subscriptions.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package subscriptions

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/apis/vmc/model"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)

// Possible value for ``offerType`` of method Subscriptions#get0.
const Subscriptions_GET_0_OFFER_TYPE_TERM = "TERM"
// Possible value for ``offerType`` of method Subscriptions#get0.
const Subscriptions_GET_0_OFFER_TYPE_ON_DEMAND = "ON_DEMAND"






func createInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["org"] = bindings.NewStringType()
    fields["subscription_request"] = bindings.NewReferenceType(model.SubscriptionRequestBindingType)
    fieldNameMap["org"] = "Org"
    fieldNameMap["subscription_request"] = "SubscriptionRequest"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func createOutputType() bindings.BindingType {
    return bindings.NewReferenceType(model.TaskBindingType)
}

func createRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["org"] = bindings.NewStringType()
    paramsTypeMap["subscription_request"] = bindings.NewReferenceType(model.SubscriptionRequestBindingType)
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
      "subscription_request",
      "POST",
      "/vmc/api/orgs/{org}/subscriptions",
       resultHeaders,
       201,
       errorHeaders,
       map[string]int{"Unauthenticated": 401,"Unauthorized": 403,"InternalServerError": 500})
}


func getInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["org"] = bindings.NewStringType()
    fields["subscription"] = bindings.NewStringType()
    fieldNameMap["org"] = "Org"
    fieldNameMap["subscription"] = "Subscription"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func getOutputType() bindings.BindingType {
    return bindings.NewReferenceType(model.SubscriptionDetailsBindingType)
}

func getRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["org"] = bindings.NewStringType()
    paramsTypeMap["subscription"] = bindings.NewStringType()
    paramsTypeMap["org"] = bindings.NewStringType()
    paramsTypeMap["subscription"] = bindings.NewStringType()
    pathParams["org"] = "org"
    pathParams["subscription"] = "subscription"
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "GET",
      "/vmc/api/orgs/{org}/subscriptions/{subscription}",
       resultHeaders,
       200,
       errorHeaders,
       map[string]int{"InternalServerError": 500,"NotFound": 404})
}


func get0InputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["org"] = bindings.NewStringType()
    fields["offer_type"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["org"] = "Org"
    fieldNameMap["offer_type"] = "OfferType"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func get0OutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewReferenceType(model.SubscriptionDetailsBindingType), reflect.TypeOf([]model.SubscriptionDetails{}))
}

func get0RestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["offer_type"] = bindings.NewOptionalType(bindings.NewStringType())
    paramsTypeMap["org"] = bindings.NewStringType()
    paramsTypeMap["org"] = bindings.NewStringType()
    pathParams["org"] = "org"
    queryParams["offer_type"] = "offer_type"
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
      "/vmc/api/orgs/{org}/subscriptions",
       resultHeaders,
       200,
       errorHeaders,
       map[string]int{"Unauthenticated": 401,"InternalServerError": 500,"NotFound": 404})
}




