/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Announcements.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package announcements

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/vapi/std"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
    "time"
)

// Resource type for the announcements. **Warning:** This constant field is available as technical preview. It may be changed in a future release.
const Announcements_RESOURCE_TYPE = "com.vmware.vcenter.system.announcement"


// The ``Severity`` enumeration class defines the severity of the announcements. **Warning:** This enumeration is available as technical preview. It may be changed in a future release.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type Severity string

const (
    // Critical problem. **Warning:** This constant field is available as technical preview. It may be changed in a future release.
     Severity_CRITICAL Severity = "CRITICAL"
    // Warning. **Warning:** This constant field is available as technical preview. It may be changed in a future release.
     Severity_WARNING Severity = "WARNING"
    // Information. **Warning:** This constant field is available as technical preview. It may be changed in a future release.
     Severity_INFO Severity = "INFO"
)

func (s Severity) Severity() bool {
    switch s {
        case Severity_CRITICAL:
            return true
        case Severity_WARNING:
            return true
        case Severity_INFO:
            return true
        default:
            return false
    }
}





// The ``Info`` class defines the announcement properties. **Warning:** This class is available as technical preview. It may be changed in a future release.
 type Info struct {
    Message std.LocalizableMessage
    Severity Severity
    ExpiresAt time.Time
}






// The ``Spec`` class defines the announcement properties for set and create operations. **Warning:** This class is available as technical preview. It may be changed in a future release.
 type Spec struct {
    Message std.LocalizableMessage
    Severity Severity
    ExpiresAt time.Time
}









func listInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func listOutputType() bindings.BindingType {
    return bindings.NewMapType(bindings.NewIdType([]string {"com.vmware.vcenter.system.announcements"}, ""), bindings.NewReferenceType(InfoBindingType),reflect.TypeOf(map[string]Info{}))
}

func listRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
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
       map[string]int{"Error": 500})
}


func getInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["announcement"] = bindings.NewIdType([]string {"com.vmware.vcenter.system.announcements"}, "")
    fieldNameMap["announcement"] = "Announcement"
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
       map[string]int{"Error": 500})
}


func createInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["spec"] = bindings.NewReferenceType(SpecBindingType)
    fieldNameMap["spec"] = "Spec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func createOutputType() bindings.BindingType {
    return bindings.NewIdType([]string {"com.vmware.vcenter.system.announcements"}, "")
}

func createRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
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
       map[string]int{"Error": 500})
}


func setInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["announcement"] = bindings.NewIdType([]string {"com.vmware.vcenter.system.announcements"}, "")
    fields["spec"] = bindings.NewReferenceType(SpecBindingType)
    fieldNameMap["announcement"] = "Announcement"
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
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
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
       map[string]int{"Error": 500})
}


func deleteInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["announcement"] = bindings.NewIdType([]string {"com.vmware.vcenter.system.announcements"}, "")
    fieldNameMap["announcement"] = "Announcement"
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
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
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
       map[string]int{"NotFound": 404,"Error": 500})
}



func InfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["message"] = bindings.NewReferenceType(std.LocalizableMessageBindingType)
    fieldNameMap["message"] = "Message"
    fields["severity"] = bindings.NewEnumType("com.vmware.vcenter.system.announcements.severity", reflect.TypeOf(Severity(Severity_CRITICAL)))
    fieldNameMap["severity"] = "Severity"
    fields["expires_at"] = bindings.NewDateTimeType()
    fieldNameMap["expires_at"] = "ExpiresAt"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.system.announcements.info",fields, reflect.TypeOf(Info{}), fieldNameMap, validators)
}

func SpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["message"] = bindings.NewReferenceType(std.LocalizableMessageBindingType)
    fieldNameMap["message"] = "Message"
    fields["severity"] = bindings.NewEnumType("com.vmware.vcenter.system.announcements.severity", reflect.TypeOf(Severity(Severity_CRITICAL)))
    fieldNameMap["severity"] = "Severity"
    fields["expires_at"] = bindings.NewDateTimeType()
    fieldNameMap["expires_at"] = "ExpiresAt"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.system.announcements.spec",fields, reflect.TypeOf(Spec{}), fieldNameMap, validators)
}


