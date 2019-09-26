/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Configuration.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package configuration

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)



// The amount of information captured in the log files varies, depending on the level setting.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type LogLevel string

const (
    // Disables logging.
     LogLevel_NONE LogLevel = "NONE"
    // Logging limited to error messages.
     LogLevel_ERROR LogLevel = "ERROR"
    // Error messages plus warning messages are logged.
     LogLevel_WARNING LogLevel = "WARNING"
    // Default setting. Errors, warnings, plus informational messages about normal operations are logged. Acceptable for production environments.
     LogLevel_INFO LogLevel = "INFO"
    // Can facilitate troubleshooting and debugging. Not recommended for production environments.
     LogLevel_VERBOSE LogLevel = "VERBOSE"
    // Extended verbose logging. Provides complete detail. Use for debugging and to facilitate client application development only. Not recommended for production environments.
     LogLevel_TRIVIA LogLevel = "TRIVIA"
)

func (l LogLevel) LogLevel() bool {
    switch l {
        case LogLevel_NONE:
            return true
        case LogLevel_ERROR:
            return true
        case LogLevel_WARNING:
            return true
        case LogLevel_INFO:
            return true
        case LogLevel_VERBOSE:
            return true
        case LogLevel_TRIVIA:
            return true
        default:
            return false
    }
}





// The ``Info`` class contains vStats configuration.
 type Info struct {
    LogLevel LogLevel
}






// The ``UpdateSpec`` class contains modifiable properties from vStats configuration.
 type UpdateSpec struct {
    LogLevel *LogLevel
}









func updateInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["update_spec"] = bindings.NewReferenceType(UpdateSpecBindingType)
    fieldNameMap["update_spec"] = "UpdateSpec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func updateOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func updateRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["update_spec"] = bindings.NewReferenceType(UpdateSpecBindingType)
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "update_spec",
      "PATCH",
      "/stats/configuration",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"Error": 500,"InvalidArgument": 400,"Unauthenticated": 401,"Unauthorized": 403})
}


func getInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
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
      "/stats/configuration",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"Error": 500,"Unauthenticated": 401,"Unauthorized": 403})
}



func InfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["log_level"] = bindings.NewEnumType("com.vmware.vstats.configuration.log_level", reflect.TypeOf(LogLevel(LogLevel_NONE)))
    fieldNameMap["log_level"] = "LogLevel"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vstats.configuration.info",fields, reflect.TypeOf(Info{}), fieldNameMap, validators)
}

func UpdateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["log_level"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vstats.configuration.log_level", reflect.TypeOf(LogLevel(LogLevel_NONE))))
    fieldNameMap["log_level"] = "LogLevel"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vstats.configuration.update_spec",fields, reflect.TypeOf(UpdateSpec{}), fieldNameMap, validators)
}


