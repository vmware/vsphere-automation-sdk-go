/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Update.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package update

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)



// ``AutoUpdateNotification`` enumeration class Defines state for automatic update notification.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type AutoUpdateNotification string

const (
    // Automatic update notification is disabled. Disable periodically query the configured url for updates.
     AutoUpdateNotification_disabled AutoUpdateNotification = "disabled"
    // Automatic update notification is enabled. Enable periodically query the configured url for updates.
     AutoUpdateNotification_enabled AutoUpdateNotification = "enabled"
)

func (a AutoUpdateNotification) AutoUpdateNotification() bool {
    switch a {
        case AutoUpdateNotification_disabled:
            return true
        case AutoUpdateNotification_enabled:
            return true
        default:
            return false
    }
}




// ``UpdateDay`` enumeration class Defines days to query for updates.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type UpdateDay string

const (
    // query for updates on Monday
     UpdateDay_Monday UpdateDay = "Monday"
    // query for updates on Tuesday
     UpdateDay_Tuesday UpdateDay = "Tuesday"
    // query for updates on Friday
     UpdateDay_Friday UpdateDay = "Friday"
    // query for updates on Wednesday
     UpdateDay_Wednesday UpdateDay = "Wednesday"
    // query for updates on Thursday
     UpdateDay_Thursday UpdateDay = "Thursday"
    // query for updates on Saturday
     UpdateDay_Saturday UpdateDay = "Saturday"
    // query for updates on Sunday
     UpdateDay_Sunday UpdateDay = "Sunday"
    // query for updates everyday
     UpdateDay_Everyday UpdateDay = "Everyday"
)

func (u UpdateDay) UpdateDay() bool {
    switch u {
        case UpdateDay_Monday:
            return true
        case UpdateDay_Tuesday:
            return true
        case UpdateDay_Friday:
            return true
        case UpdateDay_Wednesday:
            return true
        case UpdateDay_Thursday:
            return true
        case UpdateDay_Saturday:
            return true
        case UpdateDay_Sunday:
            return true
        case UpdateDay_Everyday:
            return true
        default:
            return false
    }
}





// ``UpdateStructSet`` class Structure to set url update repository.
 type UpdateStructSet struct {
    CurrentURL string
    CheckUpdates AutoUpdateNotification
    Time string
    Day UpdateDay
    Username string
    Password string
}






// ``UpdateStructGet`` class Structure to get url update repository.
 type UpdateStructGet struct {
    CurrentURL string
    DefaultURL string
    CheckUpdates AutoUpdateNotification
    Time string
    Day UpdateDay
    LatestUpdateInstallTime string
    LatestUpdateQueryTime string
    Username string
    Password string
}









func setInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["config"] = bindings.NewReferenceType(UpdateStructSetBindingType)
    fieldNameMap["config"] = "Config"
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


func getInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func getOutputType() bindings.BindingType {
    return bindings.NewReferenceType(UpdateStructGetBindingType)
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



func UpdateStructSetBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["current_URL"] = bindings.NewStringType()
    fieldNameMap["current_URL"] = "CurrentURL"
    fields["check_updates"] = bindings.NewEnumType("com.vmware.appliance.techpreview.system.update.auto_update_notification", reflect.TypeOf(AutoUpdateNotification(AutoUpdateNotification_disabled)))
    fieldNameMap["check_updates"] = "CheckUpdates"
    fields["time"] = bindings.NewStringType()
    fieldNameMap["time"] = "Time"
    fields["day"] = bindings.NewEnumType("com.vmware.appliance.techpreview.system.update.update_day", reflect.TypeOf(UpdateDay(UpdateDay_Monday)))
    fieldNameMap["day"] = "Day"
    fields["username"] = bindings.NewStringType()
    fieldNameMap["username"] = "Username"
    fields["password"] = bindings.NewSecretType()
    fieldNameMap["password"] = "Password"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.techpreview.system.update.update_struct_set",fields, reflect.TypeOf(UpdateStructSet{}), fieldNameMap, validators)
}

func UpdateStructGetBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["current_URL"] = bindings.NewStringType()
    fieldNameMap["current_URL"] = "CurrentURL"
    fields["default_URL"] = bindings.NewStringType()
    fieldNameMap["default_URL"] = "DefaultURL"
    fields["check_updates"] = bindings.NewEnumType("com.vmware.appliance.techpreview.system.update.auto_update_notification", reflect.TypeOf(AutoUpdateNotification(AutoUpdateNotification_disabled)))
    fieldNameMap["check_updates"] = "CheckUpdates"
    fields["time"] = bindings.NewStringType()
    fieldNameMap["time"] = "Time"
    fields["day"] = bindings.NewEnumType("com.vmware.appliance.techpreview.system.update.update_day", reflect.TypeOf(UpdateDay(UpdateDay_Monday)))
    fieldNameMap["day"] = "Day"
    fields["latest_update_install_time"] = bindings.NewStringType()
    fieldNameMap["latest_update_install_time"] = "LatestUpdateInstallTime"
    fields["latest_update_query_time"] = bindings.NewStringType()
    fieldNameMap["latest_update_query_time"] = "LatestUpdateQueryTime"
    fields["username"] = bindings.NewStringType()
    fieldNameMap["username"] = "Username"
    fields["password"] = bindings.NewStringType()
    fieldNameMap["password"] = "Password"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.techpreview.system.update.update_struct_get",fields, reflect.TypeOf(UpdateStructGet{}), fieldNameMap, validators)
}


