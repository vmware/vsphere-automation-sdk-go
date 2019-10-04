/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: KeyOperation.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package keyOperation

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)




// The ``GeneratedKey`` class contains properties that are returned by KeyOperation#generateKey.
 type GeneratedKey struct {
    Ciphertext string
    Plaintext *string
}






// The ``EncryptResult`` class contains properties that are returned by KeyOperation#encrypt.
 type EncryptResult struct {
    Ciphertext string
}






// The ``DecryptResult`` class contains properties that are returned by KeyOperation#decrypt.
 type DecryptResult struct {
    Plaintext string
}









func generateKeyInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["provider"] = bindings.NewIdType([]string {"com.vmware.esx.kms.providers"}, "")
    fields["num_of_bytes"] = bindings.NewIntegerType()
    fieldNameMap["provider"] = "Provider"
    fieldNameMap["num_of_bytes"] = "NumOfBytes"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func generateKeyOutputType() bindings.BindingType {
    return bindings.NewReferenceType(GeneratedKeyBindingType)
}

func generateKeyRestMetadata() protocol.OperationRestMetadata {
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
      "null",
      "",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"InvalidArgument": 400,"NotFound": 404,"Unauthenticated": 401,"Unauthorized": 403,"Error": 500})
}


func encryptInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["provider"] = bindings.NewIdType([]string {"com.vmware.esx.kms.providers"}, "")
    fields["plaintext"] = bindings.NewSecretType()
    fieldNameMap["provider"] = "Provider"
    fieldNameMap["plaintext"] = "Plaintext"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func encryptOutputType() bindings.BindingType {
    return bindings.NewReferenceType(EncryptResultBindingType)
}

func encryptRestMetadata() protocol.OperationRestMetadata {
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
      "null",
      "",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"InvalidArgument": 400,"NotFound": 404,"Unauthenticated": 401,"Unauthorized": 403,"Error": 500})
}


func decryptInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["provider"] = bindings.NewIdType([]string {"com.vmware.esx.kms.providers"}, "")
    fields["ciphertext"] = bindings.NewStringType()
    fieldNameMap["provider"] = "Provider"
    fieldNameMap["ciphertext"] = "Ciphertext"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func decryptOutputType() bindings.BindingType {
    return bindings.NewReferenceType(DecryptResultBindingType)
}

func decryptRestMetadata() protocol.OperationRestMetadata {
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
      "null",
      "",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"InvalidArgument": 400,"NotFound": 404,"Unauthenticated": 401,"Unauthorized": 403,"Error": 500})
}



func GeneratedKeyBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["ciphertext"] = bindings.NewStringType()
    fieldNameMap["ciphertext"] = "Ciphertext"
    fields["plaintext"] = bindings.NewOptionalType(bindings.NewSecretType())
    fieldNameMap["plaintext"] = "Plaintext"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.esx.kms.key_operation.generated_key",fields, reflect.TypeOf(GeneratedKey{}), fieldNameMap, validators)
}

func EncryptResultBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["ciphertext"] = bindings.NewStringType()
    fieldNameMap["ciphertext"] = "Ciphertext"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.esx.kms.key_operation.encrypt_result",fields, reflect.TypeOf(EncryptResult{}), fieldNameMap, validators)
}

func DecryptResultBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["plaintext"] = bindings.NewSecretType()
    fieldNameMap["plaintext"] = "Plaintext"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.esx.kms.key_operation.decrypt_result",fields, reflect.TypeOf(DecryptResult{}), fieldNameMap, validators)
}


