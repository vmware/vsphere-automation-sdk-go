/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Upgrade.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package upgrade

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/vapi/std"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
    "net/url"
    "time"
)



// The ``ExecutionMode`` enumeration class defines the Execution mode of NSX environment. **Warning:** This enumeration is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type ExecutionMode string

const (
    // The clusters with NSX enabled will be upgraded in parallel. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     ExecutionMode_BATCH ExecutionMode = "BATCH"
    // The clusters with NSX enabled will be upgraded in sequence one at a time. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     ExecutionMode_SERIAL ExecutionMode = "SERIAL"
)

func (e ExecutionMode) ExecutionMode() bool {
    switch e {
        case ExecutionMode_BATCH:
            return true
        case ExecutionMode_SERIAL:
            return true
        default:
            return false
    }
}




// The ``NSXNodeType`` enumeration class defines the type of the node in a cluster. **Warning:** This enumeration is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type NSXNodeType string

const (
    // The node type to identify NSX Manager. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     NSXNodeType_MANAGER NSXNodeType = "MANAGER"
    // The node type to identify NSX Edge VM. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     NSXNodeType_EDGE NSXNodeType = "EDGE"
    // The node type to identify a Host. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     NSXNodeType_HOST NSXNodeType = "HOST"
)

func (n NSXNodeType) NSXNodeType() bool {
    switch n {
        case NSXNodeType_MANAGER:
            return true
        case NSXNodeType_EDGE:
            return true
        case NSXNodeType_HOST:
            return true
        default:
            return false
    }
}




// The ``DeploymentStatus`` enumeration class defines the overall upgrade status of the deployment, including the clusters and the nodes. **Warning:** This enumeration is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type DeploymentStatus string

const (
    // The upgrade is not yet started for the deployment. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     DeploymentStatus_NOT_STARTED DeploymentStatus = "NOT_STARTED"
    // The upgrade is in progress for the deployment. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     DeploymentStatus_IN_PROGRESS DeploymentStatus = "IN_PROGRESS"
    // The upgrade will be paused after the current operation is completed. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     DeploymentStatus_PAUSE_PENDING DeploymentStatus = "PAUSE_PENDING"
    // The upgrade is paused for the deployment. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     DeploymentStatus_PAUSED DeploymentStatus = "PAUSED"
    // The upgrade is completed for the deployment. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     DeploymentStatus_COMPLETED DeploymentStatus = "COMPLETED"
)

func (d DeploymentStatus) DeploymentStatus() bool {
    switch d {
        case DeploymentStatus_NOT_STARTED:
            return true
        case DeploymentStatus_IN_PROGRESS:
            return true
        case DeploymentStatus_PAUSE_PENDING:
            return true
        case DeploymentStatus_PAUSED:
            return true
        case DeploymentStatus_COMPLETED:
            return true
        default:
            return false
    }
}




// The ``ClusterStatus`` enumeration class defines the upgrade status of the cluster. **Warning:** This enumeration is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type ClusterStatus string

const (
    // The upgrade is not yet started for the cluster. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     ClusterStatus_NOT_STARTED ClusterStatus = "NOT_STARTED"
    // The upgrade is in queue, and not yet started for the cluster. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     ClusterStatus_IN_QUEUE ClusterStatus = "IN_QUEUE"
    // The upgrade is in progress for the cluster. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     ClusterStatus_IN_PROGRESS ClusterStatus = "IN_PROGRESS"
    // The upgrade has failed for the cluster. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     ClusterStatus_FAILED ClusterStatus = "FAILED"
    // The upgrade is completed for the cluster. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     ClusterStatus_COMPLETED ClusterStatus = "COMPLETED"
)

func (c ClusterStatus) ClusterStatus() bool {
    switch c {
        case ClusterStatus_NOT_STARTED:
            return true
        case ClusterStatus_IN_QUEUE:
            return true
        case ClusterStatus_IN_PROGRESS:
            return true
        case ClusterStatus_FAILED:
            return true
        case ClusterStatus_COMPLETED:
            return true
        default:
            return false
    }
}




// The ``Status`` enumeration class defines the current status of any operation. **Warning:** This enumeration is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type Status string

const (
    // The operation is not yet started. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     Status_NOT_STARTED Status = "NOT_STARTED"
    // The operation is in progress. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     Status_IN_PROGRESS Status = "IN_PROGRESS"
    // The operation is paused. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     Status_PAUSED Status = "PAUSED"
    // The operation has failed. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     Status_FAILED Status = "FAILED"
    // The operation is completed. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     Status_COMPLETED Status = "COMPLETED"
)

func (s Status) Status() bool {
    switch s {
        case Status_NOT_STARTED:
            return true
        case Status_IN_PROGRESS:
            return true
        case Status_PAUSED:
            return true
        case Status_FAILED:
            return true
        case Status_COMPLETED:
            return true
        default:
            return false
    }
}





// The ``Messages`` class contains information about the info, warnings and erros if any. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type Messages struct {
    Info []std.LocalizableMessage
    Warnings []std.LocalizableMessage
    Errors []std.LocalizableMessage
}






// The ``NSXNodeInfo`` class contains the information about the nodes in the cluster. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type NSXNodeInfo struct {
    Id string
    NodeType NSXNodeType
    PercentComplete *float64
    PrecheckMessages Messages
    UpgradeMessages Messages
}






// The ``ClusterInfo`` class contains information about the cluster. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type ClusterInfo struct {
    Id string
    DisplayName string
    CurrentVersion string
    LastUpdated *time.Time
    Status ClusterStatus
    PercentComplete *float64
    EdgeList []NSXNodeInfo
    HostList []NSXNodeInfo
    PrecheckMessages Messages
    UpgradeMessages Messages
}






// The ``DeploymentInfo`` class contains information about the complete NSX deployment. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type DeploymentInfo struct {
    Clusters []ClusterInfo
    Managers []NSXNodeInfo
    Status DeploymentStatus
    PrecheckMessages Messages
    UpgradeMessages Messages
}






// The ``Release`` class contains information about the NSX Release. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type Release struct {
    Version string
    Name string
    ReleaseNotes *url.URL
}






// The ``Job`` class contains information about the job to be performed. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type Job struct {
    Name string
    Status Status
    StartTime *time.Time
    EndTime *time.Time
    PercentComplete *float64
}






// The ``Operation`` class contains information about the operation to be performed. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type Operation struct {
    Name string
    Status Status
    StartTime *time.Time
    EndTime *time.Time
    PercentComplete *float64
    Jobs []Job
}






// The ``Info`` class contains information about NSX deployment on the environment. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type Info struct {
    CurrentVersion string
    AvailableRelease *Release
    CurrentOperation *Operation
    Deployment DeploymentInfo
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
      "GET",
      "/vcenter/nsx/deployment/upgrade",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"Unauthenticated": 401,"Unauthorized": 403,"NotAllowedInCurrentState": 400,"Error": 500})
}


func checkInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func checkOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func checkRestMetadata() protocol.OperationRestMetadata {
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
      "/vcenter/nsx/deployment/upgrade",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"Unauthenticated": 401,"Unauthorized": 403,"NotAllowedInCurrentState": 400,"Error": 500})
}


func startInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["mode"] = bindings.NewEnumType("com.vmware.vcenter.nsx.deployment.upgrade.execution_mode", reflect.TypeOf(ExecutionMode(ExecutionMode_BATCH)))
    fieldNameMap["mode"] = "Mode"
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
    paramsTypeMap["mode"] = bindings.NewEnumType("com.vmware.vcenter.nsx.deployment.upgrade.execution_mode", reflect.TypeOf(ExecutionMode(ExecutionMode_BATCH)))
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "mode",
      "POST",
      "/vcenter/nsx/deployment/upgrade",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"Unauthenticated": 401,"Unauthorized": 403,"NotAllowedInCurrentState": 400,"AlreadyInDesiredState": 400,"Error": 500})
}


func pauseInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func pauseOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func pauseRestMetadata() protocol.OperationRestMetadata {
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
      "/vcenter/nsx/deployment/upgrade",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"Unauthenticated": 401,"Unauthorized": 403,"NotAllowedInCurrentState": 400,"AlreadyInDesiredState": 400,"Error": 500})
}


func resumeInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func resumeOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func resumeRestMetadata() protocol.OperationRestMetadata {
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
      "/vcenter/nsx/deployment/upgrade",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"Unauthenticated": 401,"Unauthorized": 403,"NotAllowedInCurrentState": 400,"Error": 500})
}



func MessagesBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["info"] = bindings.NewListType(bindings.NewReferenceType(std.LocalizableMessageBindingType), reflect.TypeOf([]std.LocalizableMessage{}))
    fieldNameMap["info"] = "Info"
    fields["warnings"] = bindings.NewListType(bindings.NewReferenceType(std.LocalizableMessageBindingType), reflect.TypeOf([]std.LocalizableMessage{}))
    fieldNameMap["warnings"] = "Warnings"
    fields["errors"] = bindings.NewListType(bindings.NewReferenceType(std.LocalizableMessageBindingType), reflect.TypeOf([]std.LocalizableMessage{}))
    fieldNameMap["errors"] = "Errors"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.nsx.deployment.upgrade.messages",fields, reflect.TypeOf(Messages{}), fieldNameMap, validators)
}

func NSXNodeInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["id"] = bindings.NewIdType([]string {"com.vmware.vcenter.deployment.node"}, "")
    fieldNameMap["id"] = "Id"
    fields["node_type"] = bindings.NewEnumType("com.vmware.vcenter.nsx.deployment.upgrade.NSX_node_type", reflect.TypeOf(NSXNodeType(NSXNodeType_MANAGER)))
    fieldNameMap["node_type"] = "NodeType"
    fields["percent_complete"] = bindings.NewOptionalType(bindings.NewDoubleType())
    fieldNameMap["percent_complete"] = "PercentComplete"
    fields["precheck_messages"] = bindings.NewReferenceType(MessagesBindingType)
    fieldNameMap["precheck_messages"] = "PrecheckMessages"
    fields["upgrade_messages"] = bindings.NewReferenceType(MessagesBindingType)
    fieldNameMap["upgrade_messages"] = "UpgradeMessages"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.nsx.deployment.upgrade.NSX_node_info",fields, reflect.TypeOf(NSXNodeInfo{}), fieldNameMap, validators)
}

func ClusterInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["id"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    fieldNameMap["id"] = "Id"
    fields["display_name"] = bindings.NewStringType()
    fieldNameMap["display_name"] = "DisplayName"
    fields["current_version"] = bindings.NewStringType()
    fieldNameMap["current_version"] = "CurrentVersion"
    fields["last_updated"] = bindings.NewOptionalType(bindings.NewDateTimeType())
    fieldNameMap["last_updated"] = "LastUpdated"
    fields["status"] = bindings.NewEnumType("com.vmware.vcenter.nsx.deployment.upgrade.cluster_status", reflect.TypeOf(ClusterStatus(ClusterStatus_NOT_STARTED)))
    fieldNameMap["status"] = "Status"
    fields["percent_complete"] = bindings.NewOptionalType(bindings.NewDoubleType())
    fieldNameMap["percent_complete"] = "PercentComplete"
    fields["edge_list"] = bindings.NewListType(bindings.NewReferenceType(NSXNodeInfoBindingType), reflect.TypeOf([]NSXNodeInfo{}))
    fieldNameMap["edge_list"] = "EdgeList"
    fields["host_list"] = bindings.NewListType(bindings.NewReferenceType(NSXNodeInfoBindingType), reflect.TypeOf([]NSXNodeInfo{}))
    fieldNameMap["host_list"] = "HostList"
    fields["precheck_messages"] = bindings.NewReferenceType(MessagesBindingType)
    fieldNameMap["precheck_messages"] = "PrecheckMessages"
    fields["upgrade_messages"] = bindings.NewReferenceType(MessagesBindingType)
    fieldNameMap["upgrade_messages"] = "UpgradeMessages"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.nsx.deployment.upgrade.cluster_info",fields, reflect.TypeOf(ClusterInfo{}), fieldNameMap, validators)
}

func DeploymentInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["clusters"] = bindings.NewListType(bindings.NewReferenceType(ClusterInfoBindingType), reflect.TypeOf([]ClusterInfo{}))
    fieldNameMap["clusters"] = "Clusters"
    fields["managers"] = bindings.NewListType(bindings.NewReferenceType(NSXNodeInfoBindingType), reflect.TypeOf([]NSXNodeInfo{}))
    fieldNameMap["managers"] = "Managers"
    fields["status"] = bindings.NewEnumType("com.vmware.vcenter.nsx.deployment.upgrade.deployment_status", reflect.TypeOf(DeploymentStatus(DeploymentStatus_NOT_STARTED)))
    fieldNameMap["status"] = "Status"
    fields["precheck_messages"] = bindings.NewReferenceType(MessagesBindingType)
    fieldNameMap["precheck_messages"] = "PrecheckMessages"
    fields["upgrade_messages"] = bindings.NewReferenceType(MessagesBindingType)
    fieldNameMap["upgrade_messages"] = "UpgradeMessages"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.nsx.deployment.upgrade.deployment_info",fields, reflect.TypeOf(DeploymentInfo{}), fieldNameMap, validators)
}

func ReleaseBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["version"] = bindings.NewStringType()
    fieldNameMap["version"] = "Version"
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    fields["release_notes"] = bindings.NewOptionalType(bindings.NewUriType())
    fieldNameMap["release_notes"] = "ReleaseNotes"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.nsx.deployment.upgrade.release",fields, reflect.TypeOf(Release{}), fieldNameMap, validators)
}

func JobBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    fields["status"] = bindings.NewEnumType("com.vmware.vcenter.nsx.deployment.upgrade.status", reflect.TypeOf(Status(Status_NOT_STARTED)))
    fieldNameMap["status"] = "Status"
    fields["start_time"] = bindings.NewOptionalType(bindings.NewDateTimeType())
    fieldNameMap["start_time"] = "StartTime"
    fields["end_time"] = bindings.NewOptionalType(bindings.NewDateTimeType())
    fieldNameMap["end_time"] = "EndTime"
    fields["percent_complete"] = bindings.NewOptionalType(bindings.NewDoubleType())
    fieldNameMap["percent_complete"] = "PercentComplete"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.nsx.deployment.upgrade.job",fields, reflect.TypeOf(Job{}), fieldNameMap, validators)
}

func OperationBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    fields["status"] = bindings.NewEnumType("com.vmware.vcenter.nsx.deployment.upgrade.status", reflect.TypeOf(Status(Status_NOT_STARTED)))
    fieldNameMap["status"] = "Status"
    fields["start_time"] = bindings.NewOptionalType(bindings.NewDateTimeType())
    fieldNameMap["start_time"] = "StartTime"
    fields["end_time"] = bindings.NewOptionalType(bindings.NewDateTimeType())
    fieldNameMap["end_time"] = "EndTime"
    fields["percent_complete"] = bindings.NewOptionalType(bindings.NewDoubleType())
    fieldNameMap["percent_complete"] = "PercentComplete"
    fields["jobs"] = bindings.NewListType(bindings.NewReferenceType(JobBindingType), reflect.TypeOf([]Job{}))
    fieldNameMap["jobs"] = "Jobs"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.nsx.deployment.upgrade.operation",fields, reflect.TypeOf(Operation{}), fieldNameMap, validators)
}

func InfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["current_version"] = bindings.NewStringType()
    fieldNameMap["current_version"] = "CurrentVersion"
    fields["available_release"] = bindings.NewOptionalType(bindings.NewReferenceType(ReleaseBindingType))
    fieldNameMap["available_release"] = "AvailableRelease"
    fields["current_operation"] = bindings.NewOptionalType(bindings.NewReferenceType(OperationBindingType))
    fieldNameMap["current_operation"] = "CurrentOperation"
    fields["deployment"] = bindings.NewReferenceType(DeploymentInfoBindingType)
    fieldNameMap["deployment"] = "Deployment"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.nsx.deployment.upgrade.info",fields, reflect.TypeOf(Info{}), fieldNameMap, validators)
}


