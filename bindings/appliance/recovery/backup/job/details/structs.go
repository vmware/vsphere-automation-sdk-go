/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Details.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package details

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/cis/task"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/vapi/std"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
    "net/url"
    "time"
)



// The ``Type`` enumeration class defines the type of backup job.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type Type string

const (
    // Job type is Scheduled.
     Type_SCHEDULED Type = "SCHEDULED"
    // Job type is Manual.
     Type_MANUAL Type = "MANUAL"
)

func (t Type) Type() bool {
    switch t {
        case Type_SCHEDULED:
            return true
        case Type_MANUAL:
            return true
        default:
            return false
    }
}





// The ``Info`` class contains information about a backup job.
 type Info struct {
    Location url.URL
    Duration *int64
    Size *int64
    Progress *task.Progress
    LocationUser string
    Type_ Type
    Messages []std.LocalizableMessage
    Build *BuildInfo
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






// The ``FilterSpec`` class contains properties used to filter the results when listing backup jobs details (see Details#list).
 type FilterSpec struct {
    Jobs map[string]bool
}






// The ``BuildInfo`` class contains information about the build of the appliance.
 type BuildInfo struct {
    VersionName string
    Version string
    BuildNumber string
}









func listInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["filter"] = bindings.NewOptionalType(bindings.NewReferenceType(FilterSpecBindingType))
    fieldNameMap["filter"] = "Filter"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func listOutputType() bindings.BindingType {
    return bindings.NewMapType(bindings.NewIdType([]string {"com.vmware.appliance.recovery.backup.job"}, ""), bindings.NewReferenceType(InfoBindingType),reflect.TypeOf(map[string]Info{}))
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



func InfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["location"] = bindings.NewUriType()
    fieldNameMap["location"] = "Location"
    fields["duration"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["duration"] = "Duration"
    fields["size"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["size"] = "Size"
    fields["progress"] = bindings.NewOptionalType(bindings.NewReferenceType(task.ProgressBindingType))
    fieldNameMap["progress"] = "Progress"
    fields["location_user"] = bindings.NewStringType()
    fieldNameMap["location_user"] = "LocationUser"
    fields["type"] = bindings.NewEnumType("com.vmware.appliance.recovery.backup.job.details.type", reflect.TypeOf(Type(Type_SCHEDULED)))
    fieldNameMap["type"] = "Type_"
    fields["messages"] = bindings.NewListType(bindings.NewReferenceType(std.LocalizableMessageBindingType), reflect.TypeOf([]std.LocalizableMessage{}))
    fieldNameMap["messages"] = "Messages"
    fields["build"] = bindings.NewOptionalType(bindings.NewReferenceType(BuildInfoBindingType))
    fieldNameMap["build"] = "Build"
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
            "SUCCEEDED": []bindings.FieldData {
                 bindings.NewFieldData("duration", true),
                 bindings.NewFieldData("size", true),
                 bindings.NewFieldData("progress", true),
                 bindings.NewFieldData("start_time", true),
                 bindings.NewFieldData("end_time", true),
            },
            "FAILED": []bindings.FieldData {
                 bindings.NewFieldData("duration", true),
                 bindings.NewFieldData("size", true),
                 bindings.NewFieldData("progress", true),
                 bindings.NewFieldData("error", false),
                 bindings.NewFieldData("start_time", true),
                 bindings.NewFieldData("end_time", true),
            },
            "RUNNING": []bindings.FieldData {
                 bindings.NewFieldData("duration", true),
                 bindings.NewFieldData("size", true),
                 bindings.NewFieldData("progress", true),
                 bindings.NewFieldData("start_time", true),
            },
            "BLOCKED": []bindings.FieldData {
                 bindings.NewFieldData("start_time", true),
            },
            "PENDING": []bindings.FieldData {},
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.appliance.recovery.backup.job.details.info",fields, reflect.TypeOf(Info{}), fieldNameMap, validators)
}

func FilterSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["jobs"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string {"com.vmware.appliance.recovery.backup.job"}, ""), reflect.TypeOf(map[string]bool{})))
    fieldNameMap["jobs"] = "Jobs"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.recovery.backup.job.details.filter_spec",fields, reflect.TypeOf(FilterSpec{}), fieldNameMap, validators)
}

func BuildInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["version_name"] = bindings.NewStringType()
    fieldNameMap["version_name"] = "VersionName"
    fields["version"] = bindings.NewStringType()
    fieldNameMap["version"] = "Version"
    fields["build_number"] = bindings.NewStringType()
    fieldNameMap["build_number"] = "BuildNumber"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.recovery.backup.job.details.build_info",fields, reflect.TypeOf(BuildInfo{}), fieldNameMap, validators)
}


