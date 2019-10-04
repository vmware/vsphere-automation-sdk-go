/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Staged.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package staged

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/apis/appliance/update"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/vapi/std"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
    "time"
)




// The ``Info`` class contains information about the staged update
 type Info struct {
    StagingComplete bool
    Version string
    Name *string
    Description std.LocalizableMessage
    Priority update.CommonInfo_Priority
    Severity update.CommonInfo_Severity
    UpdateType update.CommonInfo_Category
    ReleaseDate time.Time
    RebootRequired bool
    Size int64
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
      "null",
      "",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"Error": 500,"Unauthenticated": 401,"Unauthorized": 403,"NotAllowedInCurrentState": 400})
}


func deleteInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
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
       map[string]int{"Error": 500,"Unauthenticated": 401,"Unauthorized": 403})
}



func InfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["staging_complete"] = bindings.NewBooleanType()
    fieldNameMap["staging_complete"] = "StagingComplete"
    fields["version"] = bindings.NewIdType([]string {"com.vmware.appliance.update.pending"}, "")
    fieldNameMap["version"] = "Version"
    fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["name"] = "Name"
    fields["description"] = bindings.NewReferenceType(std.LocalizableMessageBindingType)
    fieldNameMap["description"] = "Description"
    fields["priority"] = bindings.NewEnumType("com.vmware.appliance.update.common_info.priority", reflect.TypeOf(update.CommonInfo_Priority(update.CommonInfo_Priority_HIGH)))
    fieldNameMap["priority"] = "Priority"
    fields["severity"] = bindings.NewEnumType("com.vmware.appliance.update.common_info.severity", reflect.TypeOf(update.CommonInfo_Severity(update.CommonInfo_Severity_CRITICAL)))
    fieldNameMap["severity"] = "Severity"
    fields["update_type"] = bindings.NewEnumType("com.vmware.appliance.update.common_info.category", reflect.TypeOf(update.CommonInfo_Category(update.CommonInfo_Category_SECURITY)))
    fieldNameMap["update_type"] = "UpdateType"
    fields["release_date"] = bindings.NewDateTimeType()
    fieldNameMap["release_date"] = "ReleaseDate"
    fields["reboot_required"] = bindings.NewBooleanType()
    fieldNameMap["reboot_required"] = "RebootRequired"
    fields["size"] = bindings.NewIntegerType()
    fieldNameMap["size"] = "Size"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.update.staged.info",fields, reflect.TypeOf(Info{}), fieldNameMap, validators)
}


