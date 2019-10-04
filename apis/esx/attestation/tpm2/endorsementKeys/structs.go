/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: EndorsementKeys.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package endorsementKeys

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)

// Resource type for TPM 2.0 endorsement key.
const EndorsementKeys_RESOURCE_TYPE = "com.vmware.esx.attestation.tpm2.endorsement_keys"



// The ``Summary`` class contains information that summarizes a TPM endorsement key.
 type Summary struct {
    Name string
}






// The ``Info`` class contains information that describes a TPM endorsement key.
 type Info struct {
    PublicKey string
}






// The ``CreateSpec`` class contains information that describes a TPM endorsement key. 
//
//  Only one of CreateSpec#publicKey or CreateSpec#certificate must be specified.
 type CreateSpec struct {
    Name string
    PublicKey *string
    Certificate *string
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
      "/esx/attestation/tpm2/endorsement-keys",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"Error": 500,"Unauthenticated": 401,"Unauthorized": 403})
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
    return bindings.NewVoidType()
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
      "/esx/attestation/tpm2/endorsement-keys",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"AlreadyExists": 400,"Error": 500,"InvalidArgument": 400,"Unauthenticated": 401,"Unauthorized": 403})
}


func deleteInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["name"] = bindings.NewIdType([]string {"com.vmware.esx.attestation.tpm2.endorsement_keys"}, "")
    fieldNameMap["name"] = "Name"
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
    paramsTypeMap["name"] = bindings.NewIdType([]string {"com.vmware.esx.attestation.tpm2.endorsement_keys"}, "")
    paramsTypeMap["name"] = bindings.NewIdType([]string {"com.vmware.esx.attestation.tpm2.endorsement_keys"}, "")
    pathParams["name"] = "name"
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
      "/esx/attestation/tpm2/endorsement-keys/{name}",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"Error": 500,"InvalidArgument": 400,"NotFound": 404,"Unauthenticated": 401,"Unauthorized": 403})
}


func getInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["name"] = bindings.NewIdType([]string {"com.vmware.esx.attestation.tpm2.endorsement_keys"}, "")
    fieldNameMap["name"] = "Name"
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
    paramsTypeMap["name"] = bindings.NewIdType([]string {"com.vmware.esx.attestation.tpm2.endorsement_keys"}, "")
    paramsTypeMap["name"] = bindings.NewIdType([]string {"com.vmware.esx.attestation.tpm2.endorsement_keys"}, "")
    pathParams["name"] = "name"
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
      "/esx/attestation/tpm2/endorsement-keys/{name}",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"Error": 500,"InvalidArgument": 400,"NotFound": 404,"Unauthenticated": 401,"Unauthorized": 403})
}



func SummaryBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["name"] = bindings.NewIdType([]string {"com.vmware.esx.attestation.tpm2.endorsement_keys"}, "")
    fieldNameMap["name"] = "Name"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.esx.attestation.tpm2.endorsement_keys.summary",fields, reflect.TypeOf(Summary{}), fieldNameMap, validators)
}

func InfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["public_key"] = bindings.NewStringType()
    fieldNameMap["public_key"] = "PublicKey"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.esx.attestation.tpm2.endorsement_keys.info",fields, reflect.TypeOf(Info{}), fieldNameMap, validators)
}

func CreateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["name"] = bindings.NewIdType([]string {"com.vmware.esx.attestation.tpm2.endorsement_keys"}, "")
    fieldNameMap["name"] = "Name"
    fields["public_key"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["public_key"] = "PublicKey"
    fields["certificate"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["certificate"] = "Certificate"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.esx.attestation.tpm2.endorsement_keys.create_spec",fields, reflect.TypeOf(CreateSpec{}), fieldNameMap, validators)
}


