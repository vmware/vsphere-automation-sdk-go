/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Archive.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package archive

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/apis/appliance/recovery/backup"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
    "net/url"
    "time"
)




// The ``Info`` class represents backup archive information.
 type Info struct {
    Timestamp time.Time
    Location url.URL
    Parts []string
    Version string
    SystemName string
    Comment string
}






// The ``Summary`` class contains commonly used information about a backup archive.
 type Summary struct {
    Archive string
    Timestamp time.Time
    Version string
    Comment string
}






// The ``FilterSpec`` class contains properties used to filter the results when listing backup archives (see Archive#list). If multiple properties are specified, only backup archives matching all of the properties match the filter.
 type FilterSpec struct {
    StartTimestamp *time.Time
    EndTimestamp *time.Time
    CommentSubstring *string
    MaxResults *int64
}









func getInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["spec"] = bindings.NewReferenceType(backup.LocationSpecBindingType)
    fields["system_name"] = bindings.NewIdType([]string {"com.vmware.appliance.recovery.backup.system_name"}, "")
    fields["archive"] = bindings.NewIdType([]string {"com.vmware.appliance.recovery.backup.system_name.archive"}, "")
    fieldNameMap["spec"] = "Spec"
    fieldNameMap["system_name"] = "SystemName"
    fieldNameMap["archive"] = "Archive"
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


func listInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["loc_spec"] = bindings.NewReferenceType(backup.LocationSpecBindingType)
    fields["system_name"] = bindings.NewIdType([]string {"com.vmware.appliance.recovery.backup.system_name"}, "")
    fields["filter_spec"] = bindings.NewReferenceType(FilterSpecBindingType)
    fieldNameMap["loc_spec"] = "LocSpec"
    fieldNameMap["system_name"] = "SystemName"
    fieldNameMap["filter_spec"] = "FilterSpec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func listOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewReferenceType(SummaryBindingType), reflect.TypeOf([]Summary{}))
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
       map[string]int{"NotFound": 404,"Error": 500})
}



func InfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["timestamp"] = bindings.NewDateTimeType()
    fieldNameMap["timestamp"] = "Timestamp"
    fields["location"] = bindings.NewUriType()
    fieldNameMap["location"] = "Location"
    fields["parts"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
    fieldNameMap["parts"] = "Parts"
    fields["version"] = bindings.NewStringType()
    fieldNameMap["version"] = "Version"
    fields["system_name"] = bindings.NewIdType([]string {"com.vmware.appliance.recovery.backup.system_name"}, "")
    fieldNameMap["system_name"] = "SystemName"
    fields["comment"] = bindings.NewStringType()
    fieldNameMap["comment"] = "Comment"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.recovery.backup.system_name.archive.info",fields, reflect.TypeOf(Info{}), fieldNameMap, validators)
}

func SummaryBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["archive"] = bindings.NewIdType([]string {"com.vmware.appliance.recovery.backup.system_name.archive"}, "")
    fieldNameMap["archive"] = "Archive"
    fields["timestamp"] = bindings.NewDateTimeType()
    fieldNameMap["timestamp"] = "Timestamp"
    fields["version"] = bindings.NewStringType()
    fieldNameMap["version"] = "Version"
    fields["comment"] = bindings.NewStringType()
    fieldNameMap["comment"] = "Comment"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.recovery.backup.system_name.archive.summary",fields, reflect.TypeOf(Summary{}), fieldNameMap, validators)
}

func FilterSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["start_timestamp"] = bindings.NewOptionalType(bindings.NewDateTimeType())
    fieldNameMap["start_timestamp"] = "StartTimestamp"
    fields["end_timestamp"] = bindings.NewOptionalType(bindings.NewDateTimeType())
    fieldNameMap["end_timestamp"] = "EndTimestamp"
    fields["comment_substring"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["comment_substring"] = "CommentSubstring"
    fields["max_results"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["max_results"] = "MaxResults"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.recovery.backup.system_name.archive.filter_spec",fields, reflect.TypeOf(FilterSpec{}), fieldNameMap, validators)
}


