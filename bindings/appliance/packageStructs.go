/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for package: com.vmware.appliance.
 * Includes binding types of a top level structures and enumerations.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package appliance

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/cis/task"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/vapi/std"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "time"
)



// The ``Notification`` class describes a notification that can be reported by the appliance task
type Notification struct {
    Id string
    Time *time.Time
    Message std.LocalizableMessage
    Resolution *std.LocalizableMessage
}






// The ``Notifications`` class contains info/warning/error messages that can be reported be the appliance task.
type Notifications struct {
    Info []Notification
    Warnings []Notification
    Errors []Notification
}






// The ``SubtaskInfo`` class contains information about one of the subtasks that makes up an appliance task.
type SubtaskInfo struct {
    Progress *task.Progress
    Result *Notifications
    Description std.LocalizableMessage
    Service string
    Operation string
    Parent *string
    Target *std.DynamicID
    Status task.Status
    Cancelable bool
    Error *data.ErrorValue
    StartTime *time.Time
    EndTime *time.Time
    User *string
}






// The ``TaskInfo`` class contains information about an appliance task and the subtasks of which it consists.
type TaskInfo struct {
    Progress *task.Progress
    SubtaskOrder []string
    Subtasks map[string]SubtaskInfo
    Description std.LocalizableMessage
    Service string
    Operation string
    Parent *string
    Target *std.DynamicID
    Status task.Status
    Cancelable bool
    Error *data.ErrorValue
    StartTime *time.Time
    EndTime *time.Time
    User *string
}










func NotificationBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["id"] = bindings.NewStringType()
    fieldNameMap["id"] = "Id"
    fields["time"] = bindings.NewOptionalType(bindings.NewDateTimeType())
    fieldNameMap["time"] = "Time"
    fields["message"] = bindings.NewReferenceType(std.LocalizableMessageBindingType)
    fieldNameMap["message"] = "Message"
    fields["resolution"] = bindings.NewOptionalType(bindings.NewReferenceType(std.LocalizableMessageBindingType))
    fieldNameMap["resolution"] = "Resolution"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.notification",fields, reflect.TypeOf(Notification{}), fieldNameMap, validators)
}

func NotificationsBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["info"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(NotificationBindingType), reflect.TypeOf([]Notification{})))
    fieldNameMap["info"] = "Info"
    fields["warnings"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(NotificationBindingType), reflect.TypeOf([]Notification{})))
    fieldNameMap["warnings"] = "Warnings"
    fields["errors"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(NotificationBindingType), reflect.TypeOf([]Notification{})))
    fieldNameMap["errors"] = "Errors"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.notifications",fields, reflect.TypeOf(Notifications{}), fieldNameMap, validators)
}

func SubtaskInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["progress"] = bindings.NewOptionalType(bindings.NewReferenceType(task.ProgressBindingType))
    fieldNameMap["progress"] = "Progress"
    fields["result"] = bindings.NewOptionalType(bindings.NewReferenceType(NotificationsBindingType))
    fieldNameMap["result"] = "Result"
    fields["description"] = bindings.NewReferenceType(std.LocalizableMessageBindingType)
    fieldNameMap["description"] = "Description"
    fields["service"] = bindings.NewIdType([]string {"com.vmware.vapi.service"}, "")
    fieldNameMap["service"] = "Service"
    fields["operation"] = bindings.NewIdType([]string {"com.vmware.vapi.operation"}, "")
    fieldNameMap["operation"] = "Operation"
    fields["parent"] = bindings.NewOptionalType(bindings.NewIdType([]string {"com.vmware.cis.task"}, ""))
    fieldNameMap["parent"] = "Parent"
    fields["target"] = bindings.NewOptionalType(bindings.NewReferenceType(std.DynamicIDBindingType))
    fieldNameMap["target"] = "Target"
    fields["status"] = bindings.NewEnumType("com.vmware.cis.task.status", reflect.TypeOf(task.Status(task.Status_PENDING)))
    fieldNameMap["status"] = "Status"
    fields["cancelable"] = bindings.NewBooleanType()
    fieldNameMap["cancelable"] = "Cancelable"
    fields["error"] = bindings.NewOptionalType(bindings.NewAnyErrorType())
    fieldNameMap["error"] = "Error"
    fields["start_time"] = bindings.NewOptionalType(bindings.NewDateTimeType())
    fieldNameMap["start_time"] = "StartTime"
    fields["end_time"] = bindings.NewOptionalType(bindings.NewDateTimeType())
    fieldNameMap["end_time"] = "EndTime"
    fields["user"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["user"] = "User"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("status",
        map[string][]bindings.FieldData {
            "RUNNING": []bindings.FieldData {
                 bindings.NewFieldData("progress", true),
                 bindings.NewFieldData("result", false),
                 bindings.NewFieldData("start_time", true),
            },
            "BLOCKED": []bindings.FieldData {
                 bindings.NewFieldData("progress", true),
                 bindings.NewFieldData("result", false),
                 bindings.NewFieldData("start_time", true),
            },
            "SUCCEEDED": []bindings.FieldData {
                 bindings.NewFieldData("progress", true),
                 bindings.NewFieldData("result", false),
                 bindings.NewFieldData("start_time", true),
                 bindings.NewFieldData("end_time", true),
            },
            "FAILED": []bindings.FieldData {
                 bindings.NewFieldData("progress", true),
                 bindings.NewFieldData("result", false),
                 bindings.NewFieldData("error", false),
                 bindings.NewFieldData("start_time", true),
                 bindings.NewFieldData("end_time", true),
            },
            "PENDING": []bindings.FieldData {},
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.appliance.subtask_info",fields, reflect.TypeOf(SubtaskInfo{}), fieldNameMap, validators)
}

func TaskInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["progress"] = bindings.NewOptionalType(bindings.NewReferenceType(task.ProgressBindingType))
    fieldNameMap["progress"] = "Progress"
    fields["subtask_order"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
    fieldNameMap["subtask_order"] = "SubtaskOrder"
    fields["subtasks"] = bindings.NewMapType(bindings.NewStringType(), bindings.NewReferenceType(SubtaskInfoBindingType),reflect.TypeOf(map[string]SubtaskInfo{}))
    fieldNameMap["subtasks"] = "Subtasks"
    fields["description"] = bindings.NewReferenceType(std.LocalizableMessageBindingType)
    fieldNameMap["description"] = "Description"
    fields["service"] = bindings.NewIdType([]string {"com.vmware.vapi.service"}, "")
    fieldNameMap["service"] = "Service"
    fields["operation"] = bindings.NewIdType([]string {"com.vmware.vapi.operation"}, "")
    fieldNameMap["operation"] = "Operation"
    fields["parent"] = bindings.NewOptionalType(bindings.NewIdType([]string {"com.vmware.cis.task"}, ""))
    fieldNameMap["parent"] = "Parent"
    fields["target"] = bindings.NewOptionalType(bindings.NewReferenceType(std.DynamicIDBindingType))
    fieldNameMap["target"] = "Target"
    fields["status"] = bindings.NewEnumType("com.vmware.cis.task.status", reflect.TypeOf(task.Status(task.Status_PENDING)))
    fieldNameMap["status"] = "Status"
    fields["cancelable"] = bindings.NewBooleanType()
    fieldNameMap["cancelable"] = "Cancelable"
    fields["error"] = bindings.NewOptionalType(bindings.NewAnyErrorType())
    fieldNameMap["error"] = "Error"
    fields["start_time"] = bindings.NewOptionalType(bindings.NewDateTimeType())
    fieldNameMap["start_time"] = "StartTime"
    fields["end_time"] = bindings.NewOptionalType(bindings.NewDateTimeType())
    fieldNameMap["end_time"] = "EndTime"
    fields["user"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["user"] = "User"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("status",
        map[string][]bindings.FieldData {
            "RUNNING": []bindings.FieldData {
                 bindings.NewFieldData("progress", true),
                 bindings.NewFieldData("start_time", true),
            },
            "BLOCKED": []bindings.FieldData {
                 bindings.NewFieldData("progress", true),
                 bindings.NewFieldData("start_time", true),
            },
            "SUCCEEDED": []bindings.FieldData {
                 bindings.NewFieldData("progress", true),
                 bindings.NewFieldData("start_time", true),
                 bindings.NewFieldData("end_time", true),
            },
            "FAILED": []bindings.FieldData {
                 bindings.NewFieldData("progress", true),
                 bindings.NewFieldData("error", false),
                 bindings.NewFieldData("start_time", true),
                 bindings.NewFieldData("end_time", true),
            },
            "PENDING": []bindings.FieldData {},
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.appliance.task_info",fields, reflect.TypeOf(TaskInfo{}), fieldNameMap, validators)
}


