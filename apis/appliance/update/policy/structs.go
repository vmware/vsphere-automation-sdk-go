/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Policy.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package policy

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)



// The ``DayOfWeek`` enumeration class defines the set of days
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





// The ``Time`` class defines weekday and time the automatic check for new updates will be run
 type Time struct {
    Day DayOfWeek
    Hour int64
    Minute int64
}






// The ``Info`` class defines automatic update checking and staging policy.
 type Info struct {
    CustomURL *string
    DefaultURL string
    Username *string
    CheckSchedule []Time
    AutoStage bool
    AutoUpdate bool
    ManualControl bool
}






// The ``Config`` class defines automatic update checking and staging policy.
 type Config struct {
    CustomURL *string
    Username *string
    Password *string
    CheckSchedule []Time
    AutoStage bool
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
       map[string]int{"Error": 500,"Unauthenticated": 401,"Unauthorized": 403})
}


func setInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["policy"] = bindings.NewReferenceType(ConfigBindingType)
    fieldNameMap["policy"] = "Policy"
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



func TimeBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["day"] = bindings.NewEnumType("com.vmware.appliance.update.policy.day_of_week", reflect.TypeOf(DayOfWeek(DayOfWeek_MONDAY)))
    fieldNameMap["day"] = "Day"
    fields["hour"] = bindings.NewIntegerType()
    fieldNameMap["hour"] = "Hour"
    fields["minute"] = bindings.NewIntegerType()
    fieldNameMap["minute"] = "Minute"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.update.policy.time",fields, reflect.TypeOf(Time{}), fieldNameMap, validators)
}

func InfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["custom_URL"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["custom_URL"] = "CustomURL"
    fields["default_URL"] = bindings.NewStringType()
    fieldNameMap["default_URL"] = "DefaultURL"
    fields["username"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["username"] = "Username"
    fields["check_schedule"] = bindings.NewListType(bindings.NewReferenceType(TimeBindingType), reflect.TypeOf([]Time{}))
    fieldNameMap["check_schedule"] = "CheckSchedule"
    fields["auto_stage"] = bindings.NewBooleanType()
    fieldNameMap["auto_stage"] = "AutoStage"
    fields["auto_update"] = bindings.NewBooleanType()
    fieldNameMap["auto_update"] = "AutoUpdate"
    fields["manual_control"] = bindings.NewBooleanType()
    fieldNameMap["manual_control"] = "ManualControl"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.update.policy.info",fields, reflect.TypeOf(Info{}), fieldNameMap, validators)
}

func ConfigBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["custom_URL"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["custom_URL"] = "CustomURL"
    fields["username"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["username"] = "Username"
    fields["password"] = bindings.NewOptionalType(bindings.NewSecretType())
    fieldNameMap["password"] = "Password"
    fields["check_schedule"] = bindings.NewListType(bindings.NewReferenceType(TimeBindingType), reflect.TypeOf([]Time{}))
    fieldNameMap["check_schedule"] = "CheckSchedule"
    fields["auto_stage"] = bindings.NewBooleanType()
    fieldNameMap["auto_stage"] = "AutoStage"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.update.policy.config",fields, reflect.TypeOf(Config{}), fieldNameMap, validators)
}


