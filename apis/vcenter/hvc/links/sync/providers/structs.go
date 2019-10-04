/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Providers.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package providers

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/vapi/std"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
    "time"
)

// Resource type for Sync Providers. **Warning:** This constant field is available as technical preview. It may be changed in a future release.
const Providers_RESOURCE_TYPE = "com.vmware.vcenter.hvc.links.sync.Providers"


// The ``Status`` enumeration class defines valid sync status. **Warning:** This enumeration is available as technical preview. It may be changed in a future release.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type Status string

const (
    // If Sync was successful. **Warning:** This constant field is available as technical preview. It may be changed in a future release.
     Status_SUCCEEDED Status = "SUCCEEDED"
    // If Sync failed. **Warning:** This constant field is available as technical preview. It may be changed in a future release.
     Status_FAILED Status = "FAILED"
    // If Sync has not been triggered. **Warning:** This constant field is available as technical preview. It may be changed in a future release.
     Status_NO_SYNC_FOUND Status = "NO_SYNC_FOUND"
)

func (s Status) Status() bool {
    switch s {
        case Status_SUCCEEDED:
            return true
        case Status_FAILED:
            return true
        case Status_NO_SYNC_FOUND:
            return true
        default:
            return false
    }
}





// The ``Info`` class contains information about sync for a provider. **Warning:** This class is available as technical preview. It may be changed in a future release.
 type Info struct {
    LastSyncTime *time.Time
    Status Status
    PollingIntervalInSeconds int64
    CurrentSessionInfo *SessionInfo
    StatusMessage *std.LocalizableMessage
}






// The ``SessionInfo`` class contains sync session information. **Warning:** This class is available as technical preview. It may be changed in a future release.
 type SessionInfo struct {
    Stage SessionInfo_Stage
    CompletedWork int64
    TotalWork int64
    CompletionTime *time.Time
    StartTime time.Time
    Exception *std.LocalizableMessage
}




    
    // The ``Stage`` class defines the different stages of Sync. **Warning:** This enumeration is available as technical preview. It may be changed in a future release.
    //
    // <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
     
    type SessionInfo_Stage string

    const (
        // Changes are being detected on the source replica. **Warning:** This constant field is available as technical preview. It may be changed in a future release.
         SessionInfo_Stage_CHANGE_DETECTION SessionInfo_Stage = "CHANGE_DETECTION"
        // Changes from the source replica are being enumerated. **Warning:** This constant field is available as technical preview. It may be changed in a future release.
         SessionInfo_Stage_CHANGE_ENUMERATION SessionInfo_Stage = "CHANGE_ENUMERATION"
        // Changes are being applied to the destination replica. **Warning:** This constant field is available as technical preview. It may be changed in a future release.
         SessionInfo_Stage_CHANGE_APPLICATION SessionInfo_Stage = "CHANGE_APPLICATION"
        // Sync has completed. **Warning:** This constant field is available as technical preview. It may be changed in a future release.
         SessionInfo_Stage_COMPLETED SessionInfo_Stage = "COMPLETED"
        // Sync failed. **Warning:** This constant field is available as technical preview. It may be changed in a future release.
         SessionInfo_Stage_FAILED SessionInfo_Stage = "FAILED"
        // Session is waiting for progress to be set. **Warning:** This constant field is available as technical preview. It may be changed in a future release.
         SessionInfo_Stage_WAITING SessionInfo_Stage = "WAITING"
    )

    func (s SessionInfo_Stage) SessionInfo_Stage() bool {
        switch s {
            case SessionInfo_Stage_CHANGE_DETECTION:
                return true
            case SessionInfo_Stage_CHANGE_ENUMERATION:
                return true
            case SessionInfo_Stage_CHANGE_APPLICATION:
                return true
            case SessionInfo_Stage_COMPLETED:
                return true
            case SessionInfo_Stage_FAILED:
                return true
            case SessionInfo_Stage_WAITING:
                return true
            default:
                return false
        }
    }



// The ``Summary`` class contains information about a provider. **Warning:** This class is available as technical preview. It may be changed in a future release.
 type Summary struct {
    Provider string
}






// The ``Credentials`` interface specifies user credentials to make a successful connection to remote endpoint. **Warning:** This class is available as technical preview. It may be changed in a future release.
 type Credentials struct {
    UserName string
    Password string
}









func listInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["link"] = bindings.NewIdType([]string {"com.vmware.vcenter.hvc.Links"}, "")
    fieldNameMap["link"] = "Link"
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
       map[string]int{"Error": 500,"Unauthorized": 403})
}


func getInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["link"] = bindings.NewIdType([]string {"com.vmware.vcenter.hvc.Links"}, "")
    fields["provider"] = bindings.NewIdType([]string {"com.vmware.vcenter.hvc.sync.Providers"}, "")
    fieldNameMap["link"] = "Link"
    fieldNameMap["provider"] = "Provider"
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
       map[string]int{"Error": 500,"NotFound": 404,"Unauthorized": 403})
}


func startInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["link"] = bindings.NewIdType([]string {"com.vmware.vcenter.hvc.Links"}, "")
    fields["provider"] = bindings.NewIdType([]string {"com.vmware.vcenter.hvc.sync.Providers"}, "")
    fields["credentials"] = bindings.NewOptionalType(bindings.NewReferenceType(CredentialsBindingType))
    fieldNameMap["link"] = "Link"
    fieldNameMap["provider"] = "Provider"
    fieldNameMap["credentials"] = "Credentials"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func startOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func startRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"Error": 500,"NotFound": 404,"Unauthorized": 403,"ResourceBusy": 500})
}



func InfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["last_sync_time"] = bindings.NewOptionalType(bindings.NewDateTimeType())
    fieldNameMap["last_sync_time"] = "LastSyncTime"
    fields["status"] = bindings.NewEnumType("com.vmware.vcenter.hvc.links.sync.providers.status", reflect.TypeOf(Status(Status_SUCCEEDED)))
    fieldNameMap["status"] = "Status"
    fields["polling_interval_in_seconds"] = bindings.NewIntegerType()
    fieldNameMap["polling_interval_in_seconds"] = "PollingIntervalInSeconds"
    fields["current_session_info"] = bindings.NewOptionalType(bindings.NewReferenceType(SessionInfoBindingType))
    fieldNameMap["current_session_info"] = "CurrentSessionInfo"
    fields["status_message"] = bindings.NewOptionalType(bindings.NewReferenceType(std.LocalizableMessageBindingType))
    fieldNameMap["status_message"] = "StatusMessage"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("status",
        map[string][]bindings.FieldData {
            "FAILED": []bindings.FieldData {
                 bindings.NewFieldData("status_message", true),
            },
            "SUCCEEDED": []bindings.FieldData {},
            "NO_SYNC_FOUND": []bindings.FieldData {},
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.vcenter.hvc.links.sync.providers.info",fields, reflect.TypeOf(Info{}), fieldNameMap, validators)
}

func SessionInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["stage"] = bindings.NewEnumType("com.vmware.vcenter.hvc.links.sync.providers.session_info.stage", reflect.TypeOf(SessionInfo_Stage(SessionInfo_Stage_CHANGE_DETECTION)))
    fieldNameMap["stage"] = "Stage"
    fields["completed_work"] = bindings.NewIntegerType()
    fieldNameMap["completed_work"] = "CompletedWork"
    fields["total_work"] = bindings.NewIntegerType()
    fieldNameMap["total_work"] = "TotalWork"
    fields["completion_time"] = bindings.NewOptionalType(bindings.NewDateTimeType())
    fieldNameMap["completion_time"] = "CompletionTime"
    fields["start_time"] = bindings.NewDateTimeType()
    fieldNameMap["start_time"] = "StartTime"
    fields["exception"] = bindings.NewOptionalType(bindings.NewReferenceType(std.LocalizableMessageBindingType))
    fieldNameMap["exception"] = "Exception"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("stage",
        map[string][]bindings.FieldData {
            "FAILED": []bindings.FieldData {
                 bindings.NewFieldData("completion_time", true),
                 bindings.NewFieldData("exception", true),
            },
            "COMPLETED": []bindings.FieldData {
                 bindings.NewFieldData("completion_time", true),
            },
            "CHANGE_DETECTION": []bindings.FieldData {},
            "CHANGE_ENUMERATION": []bindings.FieldData {},
            "CHANGE_APPLICATION": []bindings.FieldData {},
            "WAITING": []bindings.FieldData {},
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.vcenter.hvc.links.sync.providers.session_info",fields, reflect.TypeOf(SessionInfo{}), fieldNameMap, validators)
}

func SummaryBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["provider"] = bindings.NewIdType([]string {"com.vmware.vcenter.hvc.links.sync.Providers"}, "")
    fieldNameMap["provider"] = "Provider"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.hvc.links.sync.providers.summary",fields, reflect.TypeOf(Summary{}), fieldNameMap, validators)
}

func CredentialsBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["user_name"] = bindings.NewStringType()
    fieldNameMap["user_name"] = "UserName"
    fields["password"] = bindings.NewSecretType()
    fieldNameMap["password"] = "Password"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.hvc.links.sync.providers.credentials",fields, reflect.TypeOf(Credentials{}), fieldNameMap, validators)
}


