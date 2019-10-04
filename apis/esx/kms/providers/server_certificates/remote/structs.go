/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Remote.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package remote

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/vapi/std"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)




// The ``Summary`` class contains a summary of a remote key server certificate.
 type Summary struct {
    ServerName string
    Certificate *string
    ErrorMessages []std.LocalizableMessage
    Trusted bool
}






// The ``FilterSpec`` class contains properties used to filter the results when listing remote server certificates.
 type FilterSpec struct {
    ServerNames map[string]bool
    Trusted *bool
}









func listInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["provider"] = bindings.NewIdType([]string {"com.vmware.esx.kms.providers"}, "")
    fields["spec"] = bindings.NewOptionalType(bindings.NewReferenceType(FilterSpecBindingType))
    fieldNameMap["provider"] = "Provider"
    fieldNameMap["spec"] = "Spec"
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
    paramsTypeMap["provider"] = bindings.NewIdType([]string {"com.vmware.esx.kms.providers"}, "")
    paramsTypeMap["spec.server_names"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewStringType(), reflect.TypeOf(map[string]bool{})))
    paramsTypeMap["spec.trusted"] = bindings.NewOptionalType(bindings.NewBooleanType())
    paramsTypeMap["provider"] = bindings.NewIdType([]string {"com.vmware.esx.kms.providers"}, "")
    pathParams["provider"] = "provider"
    queryParams["spec.server_names"] = "server_names"
    queryParams["spec.trusted"] = "trusted"
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
      "/esx/kms/providers/{provider}/server-certificates/remote",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"InvalidArgument": 400,"NotFound": 404,"Unauthenticated": 401,"Unauthorized": 403,"Error": 500})
}



func SummaryBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["server_name"] = bindings.NewStringType()
    fieldNameMap["server_name"] = "ServerName"
    fields["certificate"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["certificate"] = "Certificate"
    fields["error_messages"] = bindings.NewListType(bindings.NewReferenceType(std.LocalizableMessageBindingType), reflect.TypeOf([]std.LocalizableMessage{}))
    fieldNameMap["error_messages"] = "ErrorMessages"
    fields["trusted"] = bindings.NewBooleanType()
    fieldNameMap["trusted"] = "Trusted"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.esx.kms.providers.server_certificates.remote.summary",fields, reflect.TypeOf(Summary{}), fieldNameMap, validators)
}

func FilterSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["server_names"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewStringType(), reflect.TypeOf(map[string]bool{})))
    fieldNameMap["server_names"] = "ServerNames"
    fields["trusted"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["trusted"] = "Trusted"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.esx.kms.providers.server_certificates.remote.filter_spec",fields, reflect.TypeOf(FilterSpec{}), fieldNameMap, validators)
}


