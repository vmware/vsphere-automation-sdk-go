/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: BaseImages.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package baseImages

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)

// Resource type for ESX base image.
const BaseImages_RESOURCE_TYPE = "com.vmware.esx.attestation.os.esx.base_images"



// The ``Summary`` class contains information that summarizes an ESX base image.
 type Summary struct {
    Version string
    DisplayName string
}






// The ``Info`` class contains information that describes an ESX base image.
 type Info struct {
    DisplayName string
}









func importFromImgdbInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["imgdb"] = bindings.NewBlobType()
    fieldNameMap["imgdb"] = "Imgdb"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func importFromImgdbOutputType() bindings.BindingType {
    return bindings.NewIdType([]string {"com.vmware.esx.attestation.os.esx.base_images"}, "")
}

func importFromImgdbRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["imgdb"] = bindings.NewBlobType()
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "imgdb",
      "POST",
      "/esx/attestation/os/esx/base-images",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"AlreadyExists": 400,"Error": 500,"InvalidArgument": 400,"Unauthenticated": 401,"Unauthorized": 403})
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
      "/esx/attestation/os/esx/base-images",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"Error": 500,"Unauthenticated": 401,"Unauthorized": 403})
}


func deleteInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["version"] = bindings.NewIdType([]string {"com.vmware.esx.attestation.os.esx.base_images"}, "")
    fieldNameMap["version"] = "Version"
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
    paramsTypeMap["version"] = bindings.NewIdType([]string {"com.vmware.esx.attestation.os.esx.base_images"}, "")
    paramsTypeMap["version"] = bindings.NewIdType([]string {"com.vmware.esx.attestation.os.esx.base_images"}, "")
    pathParams["version"] = "version"
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
      "/esx/attestation/os/esx/base-images/{version}",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"Error": 500,"InvalidArgument": 400,"NotFound": 404,"Unauthenticated": 401,"Unauthorized": 403})
}


func getInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["version"] = bindings.NewIdType([]string {"com.vmware.esx.attestation.os.esx.base_images"}, "")
    fieldNameMap["version"] = "Version"
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
    paramsTypeMap["version"] = bindings.NewIdType([]string {"com.vmware.esx.attestation.os.esx.base_images"}, "")
    paramsTypeMap["version"] = bindings.NewIdType([]string {"com.vmware.esx.attestation.os.esx.base_images"}, "")
    pathParams["version"] = "version"
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
      "/esx/attestation/os/esx/base-images/{version}",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"Error": 500,"InvalidArgument": 400,"NotFound": 404,"Unauthenticated": 401,"Unauthorized": 403})
}



func SummaryBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["version"] = bindings.NewIdType([]string {"com.vmware.esx.attestation.os.esx.base_images"}, "")
    fieldNameMap["version"] = "Version"
    fields["display_name"] = bindings.NewStringType()
    fieldNameMap["display_name"] = "DisplayName"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.esx.attestation.os.esx.base_images.summary",fields, reflect.TypeOf(Summary{}), fieldNameMap, validators)
}

func InfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["display_name"] = bindings.NewStringType()
    fieldNameMap["display_name"] = "DisplayName"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.esx.attestation.os.esx.base_images.info",fields, reflect.TypeOf(Info{}), fieldNameMap, validators)
}


