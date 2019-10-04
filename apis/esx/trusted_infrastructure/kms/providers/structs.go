/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Providers.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package providers

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/vapi/std"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)

// Resource type for a Key Provider.
const Providers_RESOURCE_TYPE = "com.vmware.esx.trusted_infrastructure.kms.provider"



// The ``Summary`` class contains a summary of a key provider.
 type Summary struct {
    Provider string
    Available bool
    Details []std.LocalizableMessage
}






// The ``FilterSpec`` class contains properties used to filter the results when listing the providers.
 type FilterSpec struct {
    Providers map[string]bool
    Available *bool
}









func listInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["filter_spec"] = bindings.NewOptionalType(bindings.NewReferenceType(FilterSpecBindingType))
    fieldNameMap["filter_spec"] = "FilterSpec"
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
    paramsTypeMap["filter_spec.available"] = bindings.NewOptionalType(bindings.NewBooleanType())
    paramsTypeMap["filter_spec.providers"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string {"com.vmware.esx.trusted_infrastructure.kms.provider"}, ""), reflect.TypeOf(map[string]bool{})))
    queryParams["filter_spec.available"] = "available"
    queryParams["filter_spec.providers"] = "providers"
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
      "/esx/trusted-infrastructure/kms/providers",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"Unauthenticated": 401,"Unauthorized": 403,"Error": 500})
}



func SummaryBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["provider"] = bindings.NewIdType([]string {"com.vmware.esx.trusted_infrastructure.kms.provider"}, "")
    fieldNameMap["provider"] = "Provider"
    fields["available"] = bindings.NewBooleanType()
    fieldNameMap["available"] = "Available"
    fields["details"] = bindings.NewListType(bindings.NewReferenceType(std.LocalizableMessageBindingType), reflect.TypeOf([]std.LocalizableMessage{}))
    fieldNameMap["details"] = "Details"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.esx.trusted_infrastructure.kms.providers.summary",fields, reflect.TypeOf(Summary{}), fieldNameMap, validators)
}

func FilterSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["providers"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string {"com.vmware.esx.trusted_infrastructure.kms.provider"}, ""), reflect.TypeOf(map[string]bool{})))
    fieldNameMap["providers"] = "Providers"
    fields["available"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["available"] = "Available"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.esx.trusted_infrastructure.kms.providers.filter_spec",fields, reflect.TypeOf(FilterSpec{}), fieldNameMap, validators)
}


