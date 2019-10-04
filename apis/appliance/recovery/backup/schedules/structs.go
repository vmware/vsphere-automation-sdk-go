/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Schedules.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package schedules

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/apis/appliance/recovery/backup/job"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
    "net/url"
)



// The ``DayOfWeek`` enumeration class defines the set of days when backup can be scheduled. The days can be specified as a list of individual days. You specify the days when you set the recurrence for a schedule. See RecurrenceInfo#days.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type DayOfWeek string

const (
    // Monday
     DayOfWeek_MONDAY DayOfWeek = "MONDAY"
    // Tuesday
     DayOfWeek_TUESDAY DayOfWeek = "TUESDAY"
    // Wednesday
     DayOfWeek_WEDNESDAY DayOfWeek = "WEDNESDAY"
    // Thursday
     DayOfWeek_THURSDAY DayOfWeek = "THURSDAY"
    // Friday
     DayOfWeek_FRIDAY DayOfWeek = "FRIDAY"
    // Saturday
     DayOfWeek_SATURDAY DayOfWeek = "SATURDAY"
    // Sunday
     DayOfWeek_SUNDAY DayOfWeek = "SUNDAY"
)

func (d DayOfWeek) DayOfWeek() bool {
    switch d {
        case DayOfWeek_MONDAY:
            return true
        case DayOfWeek_TUESDAY:
            return true
        case DayOfWeek_WEDNESDAY:
            return true
        case DayOfWeek_THURSDAY:
            return true
        case DayOfWeek_FRIDAY:
            return true
        case DayOfWeek_SATURDAY:
            return true
        case DayOfWeek_SUNDAY:
            return true
        default:
            return false
    }
}





// The ``RetentionInfo`` class contains retention information associated with a schedule.
 type RetentionInfo struct {
    MaxCount int64
}






// The ``RecurrenceInfo`` class contains the recurrence information associated with a schedule.
 type RecurrenceInfo struct {
    Minute int64
    Hour int64
    Days []DayOfWeek
}






// The ``CreateSpec`` class contains fields to be specified for creating a new schedule. The structure includes parts, location information, encryption password and enable flag.
 type CreateSpec struct {
    Parts []string
    BackupPassword *string
    Location url.URL
    LocationUser *string
    LocationPassword *string
    Enable *bool
    RecurrenceInfo *RecurrenceInfo
    RetentionInfo *RetentionInfo
}






// The ``Info`` class contains information about an existing schedule. The structure includes Schedule ID, parts, location information, encryption password, enable flag, recurrence and retention information.
 type Info struct {
    Parts []string
    Location url.URL
    LocationUser *string
    Enable bool
    RecurrenceInfo *RecurrenceInfo
    RetentionInfo *RetentionInfo
}






// The ``UpdateSpec`` class contains the fields of the existing schedule which can be updated.
 type UpdateSpec struct {
    Parts []string
    BackupPassword *string
    Location *url.URL
    LocationUser *string
    LocationPassword *string
    Enable *bool
    RecurrenceInfo *RecurrenceInfo
    RetentionInfo *RetentionInfo
}









func listInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func listOutputType() bindings.BindingType {
    return bindings.NewMapType(bindings.NewIdType([]string {"com.vmware.appliance.recovery.backup.schedule"}, ""), bindings.NewReferenceType(InfoBindingType),reflect.TypeOf(map[string]Info{}))
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


func runInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["schedule"] = bindings.NewIdType([]string {"com.vmware.appliance.recovery.backup.schedule"}, "")
    fields["comment"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["schedule"] = "Schedule"
    fieldNameMap["comment"] = "Comment"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func runOutputType() bindings.BindingType {
    return bindings.NewReferenceType(job.BackupJobStatusBindingType)
}

func runRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"FeatureInUse": 400,"NotFound": 404,"Error": 500})
}


func getInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["schedule"] = bindings.NewIdType([]string {"com.vmware.appliance.recovery.backup.schedule"}, "")
    fieldNameMap["schedule"] = "Schedule"
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


func createInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["schedule"] = bindings.NewIdType([]string {"com.vmware.appliance.recovery.backup.schedule"}, "")
    fields["spec"] = bindings.NewReferenceType(CreateSpecBindingType)
    fieldNameMap["schedule"] = "Schedule"
    fieldNameMap["spec"] = "Spec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func createOutputType() bindings.BindingType {
    return bindings.NewVoidType()
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
       map[string]int{"InvalidArgument": 400,"AlreadyExists": 400,"Error": 500})
}


func updateInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["schedule"] = bindings.NewIdType([]string {"com.vmware.appliance.recovery.backup.schedule"}, "")
    fields["spec"] = bindings.NewReferenceType(UpdateSpecBindingType)
    fieldNameMap["schedule"] = "Schedule"
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
       map[string]int{"InvalidArgument": 400,"NotFound": 404,"Error": 500})
}


func deleteInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["schedule"] = bindings.NewIdType([]string {"com.vmware.appliance.recovery.backup.schedule"}, "")
    fieldNameMap["schedule"] = "Schedule"
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



func RetentionInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["max_count"] = bindings.NewIntegerType()
    fieldNameMap["max_count"] = "MaxCount"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.recovery.backup.schedules.retention_info",fields, reflect.TypeOf(RetentionInfo{}), fieldNameMap, validators)
}

func RecurrenceInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["minute"] = bindings.NewIntegerType()
    fieldNameMap["minute"] = "Minute"
    fields["hour"] = bindings.NewIntegerType()
    fieldNameMap["hour"] = "Hour"
    fields["days"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewEnumType("com.vmware.appliance.recovery.backup.schedules.day_of_week", reflect.TypeOf(DayOfWeek(DayOfWeek_MONDAY))), reflect.TypeOf([]DayOfWeek{})))
    fieldNameMap["days"] = "Days"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.recovery.backup.schedules.recurrence_info",fields, reflect.TypeOf(RecurrenceInfo{}), fieldNameMap, validators)
}

func CreateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["parts"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
    fieldNameMap["parts"] = "Parts"
    fields["backup_password"] = bindings.NewOptionalType(bindings.NewSecretType())
    fieldNameMap["backup_password"] = "BackupPassword"
    fields["location"] = bindings.NewUriType()
    fieldNameMap["location"] = "Location"
    fields["location_user"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["location_user"] = "LocationUser"
    fields["location_password"] = bindings.NewOptionalType(bindings.NewSecretType())
    fieldNameMap["location_password"] = "LocationPassword"
    fields["enable"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["enable"] = "Enable"
    fields["recurrence_info"] = bindings.NewOptionalType(bindings.NewReferenceType(RecurrenceInfoBindingType))
    fieldNameMap["recurrence_info"] = "RecurrenceInfo"
    fields["retention_info"] = bindings.NewOptionalType(bindings.NewReferenceType(RetentionInfoBindingType))
    fieldNameMap["retention_info"] = "RetentionInfo"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.recovery.backup.schedules.create_spec",fields, reflect.TypeOf(CreateSpec{}), fieldNameMap, validators)
}

func InfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["parts"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
    fieldNameMap["parts"] = "Parts"
    fields["location"] = bindings.NewUriType()
    fieldNameMap["location"] = "Location"
    fields["location_user"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["location_user"] = "LocationUser"
    fields["enable"] = bindings.NewBooleanType()
    fieldNameMap["enable"] = "Enable"
    fields["recurrence_info"] = bindings.NewOptionalType(bindings.NewReferenceType(RecurrenceInfoBindingType))
    fieldNameMap["recurrence_info"] = "RecurrenceInfo"
    fields["retention_info"] = bindings.NewOptionalType(bindings.NewReferenceType(RetentionInfoBindingType))
    fieldNameMap["retention_info"] = "RetentionInfo"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.recovery.backup.schedules.info",fields, reflect.TypeOf(Info{}), fieldNameMap, validators)
}

func UpdateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["parts"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
    fieldNameMap["parts"] = "Parts"
    fields["backup_password"] = bindings.NewOptionalType(bindings.NewSecretType())
    fieldNameMap["backup_password"] = "BackupPassword"
    fields["location"] = bindings.NewOptionalType(bindings.NewUriType())
    fieldNameMap["location"] = "Location"
    fields["location_user"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["location_user"] = "LocationUser"
    fields["location_password"] = bindings.NewOptionalType(bindings.NewSecretType())
    fieldNameMap["location_password"] = "LocationPassword"
    fields["enable"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["enable"] = "Enable"
    fields["recurrence_info"] = bindings.NewOptionalType(bindings.NewReferenceType(RecurrenceInfoBindingType))
    fieldNameMap["recurrence_info"] = "RecurrenceInfo"
    fields["retention_info"] = bindings.NewOptionalType(bindings.NewReferenceType(RetentionInfoBindingType))
    fieldNameMap["retention_info"] = "RetentionInfo"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.recovery.backup.schedules.update_spec",fields, reflect.TypeOf(UpdateSpec{}), fieldNameMap, validators)
}


