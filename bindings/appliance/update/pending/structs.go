/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Pending.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package pending

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/bindings/appliance"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/bindings/appliance/update"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/vapi/std"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
    "net/url"
    "time"
)



// The ``SourceType`` enumeration class defines the supported types of sources of updates.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type SourceType string

const (
    // Do not perform a new check, return the previous result
     SourceType_LAST_CHECK SourceType = "LAST_CHECK"
    // Check the local sources, ISO devices, staged area
     SourceType_LOCAL SourceType = "LOCAL"
    // Check the local sources, ISO devices, staged area, then online repository as stated in update policy
     SourceType_LOCAL_AND_ONLINE SourceType = "LOCAL_AND_ONLINE"
)

func (s SourceType) SourceType() bool {
    switch s {
        case SourceType_LAST_CHECK:
            return true
        case SourceType_LOCAL:
            return true
        case SourceType_LOCAL_AND_ONLINE:
            return true
        default:
            return false
    }
}





// The ``Info`` class contains the extended information about the update
 type Info struct {
    Name *string
    Contents []std.LocalizableMessage
    ServicesWillBeStopped []update.ServiceInfo
    Eulas []std.LocalizableMessage
    Staged bool
    KnowledgeBase *url.URL
    Description std.LocalizableMessage
    Priority update.CommonInfo_Priority
    Severity update.CommonInfo_Severity
    UpdateType update.CommonInfo_Category
    ReleaseDate time.Time
    RebootRequired bool
    Size int64
}






// The ``Question`` class describes a item of information that must be provided by the user in order to install the update.
 type Question struct {
    DataItem string
    Text std.LocalizableMessage
    Description std.LocalizableMessage
    Type_ Question_InputType
    AllowedValues []string
    Regexp *string
    DefaultAnswer *string
}




    
    // The ``InputType`` enumeration class defines representation of field fields in GUI or CLI
    //
    // <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
     
    type Question_InputType string

    const (
        // plain text answer
         Question_InputType_PLAIN_TEXT Question_InputType = "PLAIN_TEXT"
        // Yes/No,On/Off,Checkbox answer
         Question_InputType_BOOLEAN Question_InputType = "BOOLEAN"
        // Password (masked) answer
         Question_InputType_PASSWORD Question_InputType = "PASSWORD"
    )

    func (i Question_InputType) Question_InputType() bool {
        switch i {
            case Question_InputType_PLAIN_TEXT:
                return true
            case Question_InputType_BOOLEAN:
                return true
            case Question_InputType_PASSWORD:
                return true
            default:
                return false
        }
    }



// The ``PrecheckResult`` class contains estimates of how long it will take install and rollback an update as well as a list of possible warnings and problems with installing the update.
 type PrecheckResult struct {
    CheckTime time.Time
    EstimatedTimeToInstall *int64
    EstimatedTimeToRollback *int64
    RebootRequired bool
    Issues *appliance.Notifications
    Questions []Question
}









func listInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["source_type"] = bindings.NewEnumType("com.vmware.appliance.update.pending.source_type", reflect.TypeOf(SourceType(SourceType_LAST_CHECK)))
    fields["url"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["source_type"] = "SourceType"
    fieldNameMap["url"] = "Url"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func listOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewReferenceType(update.SummaryBindingType), reflect.TypeOf([]update.Summary{}))
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
      "null",
      "",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"Error": 500,"NotFound": 404,"Unauthenticated": 401,"Unauthorized": 403})
}


func getInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["version"] = bindings.NewIdType([]string {"com.vmware.appliance.update.pending"}, "")
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
       map[string]int{"Error": 500,"Unauthenticated": 401,"Unauthorized": 403,"NotFound": 404,"AlreadyInDesiredState": 400})
}


func precheckInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["version"] = bindings.NewIdType([]string {"com.vmware.appliance.update.pending"}, "")
    fieldNameMap["version"] = "Version"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func precheckOutputType() bindings.BindingType {
    return bindings.NewReferenceType(PrecheckResultBindingType)
}

func precheckRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"Error": 500,"Unauthenticated": 401,"Unauthorized": 403,"NotFound": 404,"AlreadyInDesiredState": 400,"NotAllowedInCurrentState": 400})
}


func stageInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["version"] = bindings.NewIdType([]string {"com.vmware.appliance.update.pending"}, "")
    fieldNameMap["version"] = "Version"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func stageOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func stageRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"Error": 500,"Unauthenticated": 401,"Unauthorized": 403,"NotFound": 404,"AlreadyInDesiredState": 400,"AlreadyExists": 400,"NotAllowedInCurrentState": 400})
}


func validateInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["version"] = bindings.NewIdType([]string {"com.vmware.appliance.update.pending"}, "")
    fields["user_data"] = bindings.NewMapType(bindings.NewIdType([]string {"com.vmware.applicance.update.pending.dataitem"}, ""), bindings.NewStringType(),reflect.TypeOf(map[string]string{}))
    fieldNameMap["version"] = "Version"
    fieldNameMap["user_data"] = "UserData"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func validateOutputType() bindings.BindingType {
    return bindings.NewReferenceType(appliance.NotificationsBindingType)
}

func validateRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"Error": 500,"Unauthenticated": 401,"Unauthorized": 403,"NotFound": 404,"AlreadyInDesiredState": 400,"NotAllowedInCurrentState": 400})
}


func installInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["version"] = bindings.NewIdType([]string {"com.vmware.appliance.update.pending"}, "")
    fields["user_data"] = bindings.NewMapType(bindings.NewIdType([]string {"com.vmware.applicance.update.pending.dataitem"}, ""), bindings.NewStringType(),reflect.TypeOf(map[string]string{}))
    fieldNameMap["version"] = "Version"
    fieldNameMap["user_data"] = "UserData"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func installOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func installRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"Error": 500,"Unauthenticated": 401,"Unauthorized": 403,"NotFound": 404,"AlreadyInDesiredState": 400,"NotAllowedInCurrentState": 400})
}


func stageAndInstallInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["version"] = bindings.NewIdType([]string {"com.vmware.appliance.update.pending"}, "")
    fields["user_data"] = bindings.NewMapType(bindings.NewIdType([]string {"com.vmware.applicance.update.pending.dataitem"}, ""), bindings.NewStringType(),reflect.TypeOf(map[string]string{}))
    fieldNameMap["version"] = "Version"
    fieldNameMap["user_data"] = "UserData"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func stageAndInstallOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func stageAndInstallRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"Error": 500,"Unauthenticated": 401,"Unauthorized": 403,"NotFound": 404,"AlreadyInDesiredState": 400,"NotAllowedInCurrentState": 400})
}



func InfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["name"] = "Name"
    fields["contents"] = bindings.NewListType(bindings.NewReferenceType(std.LocalizableMessageBindingType), reflect.TypeOf([]std.LocalizableMessage{}))
    fieldNameMap["contents"] = "Contents"
    fields["services_will_be_stopped"] = bindings.NewListType(bindings.NewReferenceType(update.ServiceInfoBindingType), reflect.TypeOf([]update.ServiceInfo{}))
    fieldNameMap["services_will_be_stopped"] = "ServicesWillBeStopped"
    fields["eulas"] = bindings.NewListType(bindings.NewReferenceType(std.LocalizableMessageBindingType), reflect.TypeOf([]std.LocalizableMessage{}))
    fieldNameMap["eulas"] = "Eulas"
    fields["staged"] = bindings.NewBooleanType()
    fieldNameMap["staged"] = "Staged"
    fields["knowledge_base"] = bindings.NewOptionalType(bindings.NewUriType())
    fieldNameMap["knowledge_base"] = "KnowledgeBase"
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
    return bindings.NewStructType("com.vmware.appliance.update.pending.info",fields, reflect.TypeOf(Info{}), fieldNameMap, validators)
}

func QuestionBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["data_item"] = bindings.NewIdType([]string {"com.vmware.applicance.update.pending.dataitem"}, "")
    fieldNameMap["data_item"] = "DataItem"
    fields["text"] = bindings.NewReferenceType(std.LocalizableMessageBindingType)
    fieldNameMap["text"] = "Text"
    fields["description"] = bindings.NewReferenceType(std.LocalizableMessageBindingType)
    fieldNameMap["description"] = "Description"
    fields["type"] = bindings.NewEnumType("com.vmware.appliance.update.pending.question.input_type", reflect.TypeOf(Question_InputType(Question_InputType_PLAIN_TEXT)))
    fieldNameMap["type"] = "Type_"
    fields["allowed_values"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
    fieldNameMap["allowed_values"] = "AllowedValues"
    fields["regexp"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["regexp"] = "Regexp"
    fields["default_answer"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["default_answer"] = "DefaultAnswer"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("type",
        map[string][]bindings.FieldData {
            "PLAIN_TEXT": []bindings.FieldData {
                 bindings.NewFieldData("allowed_values", false),
                 bindings.NewFieldData("regexp", false),
                 bindings.NewFieldData("default_answer", false),
            },
            "BOOLEAN": []bindings.FieldData {},
            "PASSWORD": []bindings.FieldData {},
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.appliance.update.pending.question",fields, reflect.TypeOf(Question{}), fieldNameMap, validators)
}

func PrecheckResultBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["check_time"] = bindings.NewDateTimeType()
    fieldNameMap["check_time"] = "CheckTime"
    fields["estimated_time_to_install"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["estimated_time_to_install"] = "EstimatedTimeToInstall"
    fields["estimated_time_to_rollback"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["estimated_time_to_rollback"] = "EstimatedTimeToRollback"
    fields["reboot_required"] = bindings.NewBooleanType()
    fieldNameMap["reboot_required"] = "RebootRequired"
    fields["issues"] = bindings.NewOptionalType(bindings.NewReferenceType(appliance.NotificationsBindingType))
    fieldNameMap["issues"] = "Issues"
    fields["questions"] = bindings.NewListType(bindings.NewReferenceType(QuestionBindingType), reflect.TypeOf([]Question{}))
    fieldNameMap["questions"] = "Questions"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.update.pending.precheck_result",fields, reflect.TypeOf(PrecheckResult{}), fieldNameMap, validators)
}


