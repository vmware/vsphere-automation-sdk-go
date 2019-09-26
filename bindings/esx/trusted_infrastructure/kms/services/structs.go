/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Services.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package services

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/bindings/esx/trusted_infrastructure"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)

// The resource type for the Services interface.
const Services_RESOURCE_TYPE = "com.vmware.esx.trusted_infrastructure.kms.service"



// The ``Summary`` class contains a summary of a KMS service instance.
 type Summary struct {
    Service string
    Address trusted_infrastructure.NetworkAddress
    Group string
}






// The ``Info`` class contains all the stored information about a KMS service instance.
 type Info struct {
    Address trusted_infrastructure.NetworkAddress
    TrustedCA trusted_infrastructure.X509CertChain
    Group string
}






// The ``CreateSpec`` class contains the data necessary for adding a KMS service instance to the environment.
 type CreateSpec struct {
    Address trusted_infrastructure.NetworkAddress
    TrustedCA trusted_infrastructure.X509CertChain
    Group string
}






// The ``FilterSpec`` class contains the data necessary for identifying a KMS service instance.
 type FilterSpec struct {
    Services map[string]bool
    Address []trusted_infrastructure.NetworkAddress
    Group map[string]bool
}









func listInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["filter"] = bindings.NewOptionalType(bindings.NewReferenceType(FilterSpecBindingType))
    fieldNameMap["filter"] = "Filter"
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
    paramsTypeMap["filter"] = bindings.NewOptionalType(bindings.NewReferenceType(FilterSpecBindingType))
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "filter",
      "POST",
      "/esx/trusted-infrastructure/kms/services",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"Error": 500,"Unauthenticated": 401})
}


func getInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["service"] = bindings.NewIdType([]string {"com.vmware.esx.trusted_infrastructure.kms.service"}, "")
    fieldNameMap["service"] = "Service"
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
    paramsTypeMap["service"] = bindings.NewIdType([]string {"com.vmware.esx.trusted_infrastructure.kms.service"}, "")
    paramsTypeMap["service"] = bindings.NewIdType([]string {"com.vmware.esx.trusted_infrastructure.kms.service"}, "")
    pathParams["service"] = "service"
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
      "/esx/trusted-infrastructure/kms/services/{service}",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"NotFound": 404,"Error": 500,"Unauthenticated": 401})
}


func createInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["spec"] = bindings.NewReferenceType(CreateSpecBindingType)
    fieldNameMap["spec"] = "Spec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func createOutputType() bindings.BindingType {
    return bindings.NewIdType([]string {"com.vmware.esx.trusted_infrastructure.kms.service"}, "")
}

func createRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["spec"] = bindings.NewReferenceType(CreateSpecBindingType)
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
      "/esx/trusted-infrastructure/kms/services",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"AlreadyExists": 400,"InvalidArgument": 400,"UnableToAllocateResource": 500,"Error": 500,"Unauthenticated": 401})
}


func deleteInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["service"] = bindings.NewIdType([]string {"com.vmware.esx.trusted_infrastructure.kms.service"}, "")
    fieldNameMap["service"] = "Service"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func deleteOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func deleteRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["service"] = bindings.NewIdType([]string {"com.vmware.esx.trusted_infrastructure.kms.service"}, "")
    paramsTypeMap["service"] = bindings.NewIdType([]string {"com.vmware.esx.trusted_infrastructure.kms.service"}, "")
    pathParams["service"] = "service"
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "DELETE",
      "/esx/trusted-infrastructure/kms/services/{service}",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"NotFound": 404,"Error": 500,"Unauthenticated": 401})
}



func SummaryBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["service"] = bindings.NewIdType([]string {"com.vmware.esx.trusted_infrastructure.kms.service"}, "")
    fieldNameMap["service"] = "Service"
    fields["address"] = bindings.NewReferenceType(trusted_infrastructure.NetworkAddressBindingType)
    fieldNameMap["address"] = "Address"
    fields["group"] = bindings.NewStringType()
    fieldNameMap["group"] = "Group"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.esx.trusted_infrastructure.kms.services.summary",fields, reflect.TypeOf(Summary{}), fieldNameMap, validators)
}

func InfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["address"] = bindings.NewReferenceType(trusted_infrastructure.NetworkAddressBindingType)
    fieldNameMap["address"] = "Address"
    fields["trusted_CA"] = bindings.NewReferenceType(trusted_infrastructure.X509CertChainBindingType)
    fieldNameMap["trusted_CA"] = "TrustedCA"
    fields["group"] = bindings.NewStringType()
    fieldNameMap["group"] = "Group"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.esx.trusted_infrastructure.kms.services.info",fields, reflect.TypeOf(Info{}), fieldNameMap, validators)
}

func CreateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["address"] = bindings.NewReferenceType(trusted_infrastructure.NetworkAddressBindingType)
    fieldNameMap["address"] = "Address"
    fields["trusted_CA"] = bindings.NewReferenceType(trusted_infrastructure.X509CertChainBindingType)
    fieldNameMap["trusted_CA"] = "TrustedCA"
    fields["group"] = bindings.NewStringType()
    fieldNameMap["group"] = "Group"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.esx.trusted_infrastructure.kms.services.create_spec",fields, reflect.TypeOf(CreateSpec{}), fieldNameMap, validators)
}

func FilterSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["services"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string {"com.vmware.esx.trusted_infrastructure.kms.service"}, ""), reflect.TypeOf(map[string]bool{})))
    fieldNameMap["services"] = "Services"
    fields["address"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(trusted_infrastructure.NetworkAddressBindingType), reflect.TypeOf([]trusted_infrastructure.NetworkAddress{})))
    fieldNameMap["address"] = "Address"
    fields["group"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewStringType(), reflect.TypeOf(map[string]bool{})))
    fieldNameMap["group"] = "Group"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.esx.trusted_infrastructure.kms.services.filter_spec",fields, reflect.TypeOf(FilterSpec{}), fieldNameMap, validators)
}


