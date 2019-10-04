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
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/apis/appliance"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
    "time"
)



// The ``State`` enumeration class defines the various states the appliance update can be in.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type State string

const (
    // The appliance is up to date.
     State_UP_TO_DATE State = "UP_TO_DATE"
    // A new update is available.
     State_UPDATES_PENDING State = "UPDATES_PENDING"
    // The appliance update is in progress of downloading an update.
     State_STAGE_IN_PROGRESS State = "STAGE_IN_PROGRESS"
    // The appliance update is in progress of installing an update.
     State_INSTALL_IN_PROGRESS State = "INSTALL_IN_PROGRESS"
    // The appliance update failed and cannot recover.
     State_INSTALL_FAILED State = "INSTALL_FAILED"
    // The appliance update failed and recovery is in progress.
     State_ROLLBACK_IN_PROGRESS State = "ROLLBACK_IN_PROGRESS"
)

func (s State) State() bool {
    switch s {
        case State_UP_TO_DATE:
            return true
        case State_UPDATES_PENDING:
            return true
        case State_STAGE_IN_PROGRESS:
            return true
        case State_INSTALL_IN_PROGRESS:
            return true
        case State_INSTALL_FAILED:
            return true
        case State_ROLLBACK_IN_PROGRESS:
            return true
        default:
            return false
    }
}





// The ``Info`` class describes the state of the appliance update.
 type Info struct {
    State State
    Task *appliance.TaskInfo
    Version string
    LatestQueryTime *time.Time
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


func cancelInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func cancelOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func cancelRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"Error": 500,"NotAllowedInCurrentState": 400,"Unauthenticated": 401,"Unauthorized": 403})
}



func InfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["state"] = bindings.NewEnumType("com.vmware.appliance.update.state", reflect.TypeOf(State(State_UP_TO_DATE)))
    fieldNameMap["state"] = "State"
    fields["task"] = bindings.NewOptionalType(bindings.NewReferenceType(appliance.TaskInfoBindingType))
    fieldNameMap["task"] = "Task"
    fields["version"] = bindings.NewStringType()
    fieldNameMap["version"] = "Version"
    fields["latest_query_time"] = bindings.NewOptionalType(bindings.NewDateTimeType())
    fieldNameMap["latest_query_time"] = "LatestQueryTime"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("state",
        map[string][]bindings.FieldData {
            "UP_TO_DATE": []bindings.FieldData {},
            "UPDATES_PENDING": []bindings.FieldData {},
            "STAGE_IN_PROGRESS": []bindings.FieldData {},
            "INSTALL_IN_PROGRESS": []bindings.FieldData {},
            "INSTALL_FAILED": []bindings.FieldData {},
            "ROLLBACK_IN_PROGRESS": []bindings.FieldData {},
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.appliance.update.info",fields, reflect.TypeOf(Info{}), fieldNameMap, validators)
}


