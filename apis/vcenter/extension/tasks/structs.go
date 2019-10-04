/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Tasks.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package tasks

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/cis/task"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/vapi/std"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
    "time"
)

// Resource type for task. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
const Tasks_RESOURCE_TYPE = "com.vmware.cis.task"



// The ``CreateSpec`` class contains properties used to create a task. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type CreateSpec struct {
    Target *std.DynamicID
    OperationId string
    PercentDone *task.Progress
    Description *std.LocalizableMessage
    Status *task.Status
    User *string
    Cancelable bool
    ParentTask *string
    Error *data.ErrorValue
    Result data.DataValue
    StartTime *time.Time
    EndTime *time.Time
}






// The ``UpdateSpec`` class contains properties used to update the information of the operation associated with a task. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type UpdateSpec struct {
    PercentDone *task.Progress
    Description *std.LocalizableMessage
    Status *task.Status
    Error *data.ErrorValue
    Result data.DataValue
    StartTime *time.Time
    EndTime *time.Time
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
    return bindings.NewIdType([]string {"com.vmware.cis.task"}, "")
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
       map[string]int{"Error": 500,"NotAllowedInCurrentState": 400,"InvalidArgument": 400})
}


func updateInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["task_id"] = bindings.NewIdType([]string {"com.vmware.cis.task"}, "")
    fields["spec"] = bindings.NewReferenceType(UpdateSpecBindingType)
    fieldNameMap["task_id"] = "TaskId"
    fieldNameMap["spec"] = "Spec"
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
       map[string]int{"Error": 500,"NotAllowedInCurrentState": 400,"InvalidArgument": 400})
}



func CreateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["target"] = bindings.NewOptionalType(bindings.NewReferenceType(std.DynamicIDBindingType))
    fieldNameMap["target"] = "Target"
    fields["operation_id"] = bindings.NewIdType([]string {"com.vmware.cis.task"}, "")
    fieldNameMap["operation_id"] = "OperationId"
    fields["percent_done"] = bindings.NewOptionalType(bindings.NewReferenceType(task.ProgressBindingType))
    fieldNameMap["percent_done"] = "PercentDone"
    fields["description"] = bindings.NewOptionalType(bindings.NewReferenceType(std.LocalizableMessageBindingType))
    fieldNameMap["description"] = "Description"
    fields["status"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.cis.task.status", reflect.TypeOf(task.Status(task.Status_PENDING))))
    fieldNameMap["status"] = "Status"
    fields["user"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["user"] = "User"
    fields["cancelable"] = bindings.NewBooleanType()
    fieldNameMap["cancelable"] = "Cancelable"
    fields["parent_task"] = bindings.NewOptionalType(bindings.NewIdType([]string {"com.vmware.cis.task"}, ""))
    fieldNameMap["parent_task"] = "ParentTask"
    fields["error"] = bindings.NewOptionalType(bindings.NewAnyErrorType())
    fieldNameMap["error"] = "Error"
    fields["result"] = bindings.NewOptionalType(bindings.NewOpaqueType())
    fieldNameMap["result"] = "Result"
    fields["start_time"] = bindings.NewOptionalType(bindings.NewDateTimeType())
    fieldNameMap["start_time"] = "StartTime"
    fields["end_time"] = bindings.NewOptionalType(bindings.NewDateTimeType())
    fieldNameMap["end_time"] = "EndTime"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("status",
        map[string][]bindings.FieldData {
            "FAILED": []bindings.FieldData {
                 bindings.NewFieldData("error", false),
                 bindings.NewFieldData("result", false),
                 bindings.NewFieldData("start_time", false),
                 bindings.NewFieldData("end_time", false),
            },
            "RUNNING": []bindings.FieldData {
                 bindings.NewFieldData("result", false),
                 bindings.NewFieldData("start_time", false),
            },
            "BLOCKED": []bindings.FieldData {
                 bindings.NewFieldData("result", false),
                 bindings.NewFieldData("start_time", false),
            },
            "SUCCEEDED": []bindings.FieldData {
                 bindings.NewFieldData("result", false),
                 bindings.NewFieldData("start_time", false),
                 bindings.NewFieldData("end_time", false),
            },
            "PENDING": []bindings.FieldData {},
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.vcenter.extension.tasks.create_spec",fields, reflect.TypeOf(CreateSpec{}), fieldNameMap, validators)
}

func UpdateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["percent_done"] = bindings.NewOptionalType(bindings.NewReferenceType(task.ProgressBindingType))
    fieldNameMap["percent_done"] = "PercentDone"
    fields["description"] = bindings.NewOptionalType(bindings.NewReferenceType(std.LocalizableMessageBindingType))
    fieldNameMap["description"] = "Description"
    fields["status"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.cis.task.status", reflect.TypeOf(task.Status(task.Status_PENDING))))
    fieldNameMap["status"] = "Status"
    fields["error"] = bindings.NewOptionalType(bindings.NewAnyErrorType())
    fieldNameMap["error"] = "Error"
    fields["result"] = bindings.NewOptionalType(bindings.NewOpaqueType())
    fieldNameMap["result"] = "Result"
    fields["start_time"] = bindings.NewOptionalType(bindings.NewDateTimeType())
    fieldNameMap["start_time"] = "StartTime"
    fields["end_time"] = bindings.NewOptionalType(bindings.NewDateTimeType())
    fieldNameMap["end_time"] = "EndTime"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("status",
        map[string][]bindings.FieldData {
            "FAILED": []bindings.FieldData {
                 bindings.NewFieldData("error", false),
                 bindings.NewFieldData("result", false),
                 bindings.NewFieldData("start_time", false),
                 bindings.NewFieldData("end_time", false),
            },
            "RUNNING": []bindings.FieldData {
                 bindings.NewFieldData("result", false),
                 bindings.NewFieldData("start_time", false),
            },
            "BLOCKED": []bindings.FieldData {
                 bindings.NewFieldData("result", false),
                 bindings.NewFieldData("start_time", false),
            },
            "SUCCEEDED": []bindings.FieldData {
                 bindings.NewFieldData("result", false),
                 bindings.NewFieldData("start_time", false),
                 bindings.NewFieldData("end_time", false),
            },
            "PENDING": []bindings.FieldData {},
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.vcenter.extension.tasks.update_spec",fields, reflect.TypeOf(UpdateSpec{}), fieldNameMap, validators)
}


