/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Recovery.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package recovery

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/vapi/std"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
    "time"
)



// The ``LossType`` enumeration class defines the type of loss to recover from. **Warning:** This enumeration is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type LossType string

const (
    // Indicates loss of storage and servers. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     LossType_ALL LossType = "ALL"
)

func (l LossType) LossType() bool {
    switch l {
        case LossType_ALL:
            return true
        default:
            return false
    }
}




// The ``RecoveryStage`` enumeration class defines the stage of the NSX infrastructure recovery operation. **Warning:** This enumeration is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type RecoveryStage string

const (
    // Recovery operation is not in progress. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     RecoveryStage_NONE RecoveryStage = "NONE"
    // Recovery operation is resetting the service account password. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     RecoveryStage_SERVICE_ACCOUNT_PASSWD_RESET RecoveryStage = "SERVICE_ACCOUNT_PASSWD_RESET"
    // Recovery operation is locating the NSX-I VM. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     RecoveryStage_NSX_LOCATE RecoveryStage = "NSX_LOCATE"
    // Recovery operation is creating NSX infrastructure for WCP clusters. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     RecoveryStage_NSX_INSTALL RecoveryStage = "NSX_INSTALL"
    // Recovery operation is complete. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     RecoveryStage_DONE RecoveryStage = "DONE"
)

func (r RecoveryStage) RecoveryStage() bool {
    switch r {
        case RecoveryStage_NONE:
            return true
        case RecoveryStage_SERVICE_ACCOUNT_PASSWD_RESET:
            return true
        case RecoveryStage_NSX_LOCATE:
            return true
        case RecoveryStage_NSX_INSTALL:
            return true
        case RecoveryStage_DONE:
            return true
        default:
            return false
    }
}





// The ``Info`` class contains information related to recovery needed. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type Info struct {
    Needed bool
    Loss *LossType
}






// The ``ExecuteSpec`` class contains information related to recovery of NSX infrastructure for WCP clusters. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type ExecuteSpec struct {
    RootPassword string
}






// The ``ExecutionStatus`` class contains information related to the status of recovery of NSX infrastructure for WCP clusters. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type ExecutionStatus struct {
    Stage RecoveryStage
    Errors []std.LocalizableMessage
    StartTime time.Time
    EndTime time.Time
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
      "POST",
      "/vcenter/nsx/recovery",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"Unauthorized": 403,"Unauthenticated": 401,"InternalServerError": 500,"Error": 500})
}


func executeInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["spec"] = bindings.NewReferenceType(ExecuteSpecBindingType)
    fieldNameMap["spec"] = "Spec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func executeOutputType() bindings.BindingType {
    return bindings.NewIdType([]string {"Task"}, "")
}

func executeRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["spec"] = bindings.NewReferenceType(ExecuteSpecBindingType)
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "spec",
      "POST",
      "/vcenter/nsx/recovery",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"Unauthenticated": 401,"Unauthorized": 403,"NotAllowedInCurrentState": 400,"AlreadyInDesiredState": 400,"Unsupported": 400,"InternalServerError": 500,"Error": 500})
}


func executeStatusInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func executeStatusOutputType() bindings.BindingType {
    return bindings.NewReferenceType(ExecutionStatusBindingType)
}

func executeStatusRestMetadata() protocol.OperationRestMetadata {
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
      "GET",
      "/vcenter/nsx/recovery/execute/status",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"Unauthenticated": 401,"Unauthorized": 403,"Error": 500})
}



func InfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["needed"] = bindings.NewBooleanType()
    fieldNameMap["needed"] = "Needed"
    fields["loss"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.nsx.recovery.loss_type", reflect.TypeOf(LossType(LossType_ALL))))
    fieldNameMap["loss"] = "Loss"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.nsx.recovery.info",fields, reflect.TypeOf(Info{}), fieldNameMap, validators)
}

func ExecuteSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["root_password"] = bindings.NewSecretType()
    fieldNameMap["root_password"] = "RootPassword"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.nsx.recovery.execute_spec",fields, reflect.TypeOf(ExecuteSpec{}), fieldNameMap, validators)
}

func ExecutionStatusBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["stage"] = bindings.NewEnumType("com.vmware.vcenter.nsx.recovery.recovery_stage", reflect.TypeOf(RecoveryStage(RecoveryStage_NONE)))
    fieldNameMap["stage"] = "Stage"
    fields["errors"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(std.LocalizableMessageBindingType), reflect.TypeOf([]std.LocalizableMessage{})))
    fieldNameMap["errors"] = "Errors"
    fields["start_time"] = bindings.NewDateTimeType()
    fieldNameMap["start_time"] = "StartTime"
    fields["end_time"] = bindings.NewDateTimeType()
    fieldNameMap["end_time"] = "EndTime"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.nsx.recovery.execution_status",fields, reflect.TypeOf(ExecutionStatus{}), fieldNameMap, validators)
}


