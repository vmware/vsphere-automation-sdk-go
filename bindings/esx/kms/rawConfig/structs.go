/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: RawConfig.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package rawConfig

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
    "time"
)



// The ``InfoType`` enumeration class describes the type of information that is populated in Info when returned by RawConfig#get.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type InfoType string

const (
    // All properties in Info are populated.
     InfoType_FULL InfoType = "FULL"
    // The Info#configuration property is not populated. 
    //
    //  The Info#configuration property may be a large document. Specifying a InfoType#InfoType_BRIEF request will reduce the size of the Info class on common operations.
     InfoType_BRIEF InfoType = "BRIEF"
)

func (i InfoType) InfoType() bool {
    switch i {
        case InfoType_FULL:
            return true
        case InfoType_BRIEF:
            return true
        default:
            return false
    }
}





// The ``Info`` class contains information that describes the persistent user configuration.
 type Info struct {
    Configuration []byte
    Fingerprint string
    LastUpdateTime time.Time
}






// The ``SetSpec`` class contains information that describes a new configuration to applied using RawConfig#set.
 type SetSpec struct {
    Configuration []byte
}









func getInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["type"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.esx.kms.raw_config.info_type", reflect.TypeOf(InfoType(InfoType_FULL))))
    fieldNameMap["type"] = "Type_"
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
    paramsTypeMap["type"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.esx.kms.raw_config.info_type", reflect.TypeOf(InfoType(InfoType_FULL))))
    queryParams["type"] = "type"
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
      "/esx/kms/raw-config",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"Error": 500,"InvalidArgument": 400,"Unauthenticated": 401,"Unauthorized": 403})
}


func setInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["spec"] = bindings.NewReferenceType(SetSpecBindingType)
    fieldNameMap["spec"] = "Spec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func setOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func setRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["spec"] = bindings.NewReferenceType(SetSpecBindingType)
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "spec",
      "PUT",
      "/esx/kms/raw-config",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"Error": 500,"InvalidArgument": 400,"Unauthenticated": 401,"Unauthorized": 403})
}



func InfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["configuration"] = bindings.NewOptionalType(bindings.NewBlobType())
    fieldNameMap["configuration"] = "Configuration"
    fields["fingerprint"] = bindings.NewStringType()
    fieldNameMap["fingerprint"] = "Fingerprint"
    fields["last_update_time"] = bindings.NewDateTimeType()
    fieldNameMap["last_update_time"] = "LastUpdateTime"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.esx.kms.raw_config.info",fields, reflect.TypeOf(Info{}), fieldNameMap, validators)
}

func SetSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["configuration"] = bindings.NewBlobType()
    fieldNameMap["configuration"] = "Configuration"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.esx.kms.raw_config.set_spec",fields, reflect.TypeOf(SetSpec{}), fieldNameMap, validators)
}


