/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Job.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package job

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/vapi/std"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
    "time"
)



// The ``Status`` enumeration class defines the status values that can be reported for an operation.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type Status string

const (
    // The operation is not running.
     Status_NONE Status = "NONE"
    // The operation is in progress.
     Status_RUNNING Status = "RUNNING"
    // The operation completed successfully.
     Status_SUCCEEDED Status = "SUCCEEDED"
    // The operation failed.
     Status_FAILED Status = "FAILED"
)

func (s Status) Status() bool {
    switch s {
        case Status_NONE:
            return true
        case Status_RUNNING:
            return true
        case Status_SUCCEEDED:
            return true
        case Status_FAILED:
            return true
        default:
            return false
    }
}





// The ``CreateSpec`` class has the fields to request the start of reconciliation job.
 type CreateSpec struct {
    SsoAdminUserName *string
    SsoAdminUserPassword *string
    IgnoreWarnings *bool
}






// The ``Info`` class represents the reconciliation job information. It contains information related to current Status, any associated messages and progress as percentage.
 type Info struct {
    Description std.LocalizableMessage
    Service string
    Operation string
    Parent *string
    Target *std.DynamicID
    Status Status
    Cancelable *bool
    Error *data.ErrorValue
    StartTime *time.Time
    EndTime *time.Time
    Messages []std.LocalizableMessage
    Progress int64
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
    return bindings.NewReferenceType(InfoBindingType)
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
       map[string]int{"FeatureInUse": 400,"NotAllowedInCurrentState": 400,"InvalidArgument": 400,"Error": 500})
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



func CreateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["sso_admin_user_name"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["sso_admin_user_name"] = "SsoAdminUserName"
    fields["sso_admin_user_password"] = bindings.NewOptionalType(bindings.NewSecretType())
    fieldNameMap["sso_admin_user_password"] = "SsoAdminUserPassword"
    fields["ignore_warnings"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["ignore_warnings"] = "IgnoreWarnings"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.recovery.reconciliation.job.create_spec",fields, reflect.TypeOf(CreateSpec{}), fieldNameMap, validators)
}

func InfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["description"] = bindings.NewReferenceType(std.LocalizableMessageBindingType)
    fieldNameMap["description"] = "Description"
    fields["service"] = bindings.NewIdType([]string {"com.vmware.appliance.recovery.reconciliation.job"}, "")
    fieldNameMap["service"] = "Service"
    fields["operation"] = bindings.NewStringType()
    fieldNameMap["operation"] = "Operation"
    fields["parent"] = bindings.NewOptionalType(bindings.NewIdType([]string {"com.vmware.appliance.recovery.reconciliation.job"}, ""))
    fieldNameMap["parent"] = "Parent"
    fields["target"] = bindings.NewOptionalType(bindings.NewReferenceType(std.DynamicIDBindingType))
    fieldNameMap["target"] = "Target"
    fields["status"] = bindings.NewEnumType("com.vmware.appliance.recovery.reconciliation.job.status", reflect.TypeOf(Status(Status_NONE)))
    fieldNameMap["status"] = "Status"
    fields["cancelable"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["cancelable"] = "Cancelable"
    fields["error"] = bindings.NewOptionalType(bindings.NewAnyErrorType())
    fieldNameMap["error"] = "Error"
    fields["start_time"] = bindings.NewOptionalType(bindings.NewDateTimeType())
    fieldNameMap["start_time"] = "StartTime"
    fields["end_time"] = bindings.NewOptionalType(bindings.NewDateTimeType())
    fieldNameMap["end_time"] = "EndTime"
    fields["messages"] = bindings.NewListType(bindings.NewReferenceType(std.LocalizableMessageBindingType), reflect.TypeOf([]std.LocalizableMessage{}))
    fieldNameMap["messages"] = "Messages"
    fields["progress"] = bindings.NewIntegerType()
    fieldNameMap["progress"] = "Progress"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("status",
        map[string][]bindings.FieldData {
            "FAILED": []bindings.FieldData {
                 bindings.NewFieldData("error", false),
                 bindings.NewFieldData("start_time", true),
                 bindings.NewFieldData("end_time", true),
            },
            "RUNNING": []bindings.FieldData {
                 bindings.NewFieldData("start_time", true),
            },
            "SUCCEEDED": []bindings.FieldData {
                 bindings.NewFieldData("start_time", true),
                 bindings.NewFieldData("end_time", true),
            },
            "NONE": []bindings.FieldData {},
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.appliance.recovery.reconciliation.job.info",fields, reflect.TypeOf(Info{}), fieldNameMap, validators)
}


