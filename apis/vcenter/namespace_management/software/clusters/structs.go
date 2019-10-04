/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Clusters.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package clusters

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/vapi/std"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
    "time"
)



// The ``State`` enumeration class describes the state of the upgrade. **Warning:** This enumeration is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type State string

const (
    // Upgrade is in progress. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     State_PENDING State = "PENDING"
    // Cluster is ready when there is no upgrade or upgrade is completed. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     State_READY State = "READY"
    // Upgrade failed and need user intervention. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     State_ERROR State = "ERROR"
)

func (s State) State() bool {
    switch s {
        case State_PENDING:
            return true
        case State_READY:
            return true
        case State_ERROR:
            return true
        default:
            return false
    }
}





// The ``Result`` class contains the result of batch upgrade method. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type Result struct {
    Res Result_Res
    Exception *data.ErrorValue
}




    
    // The ``Res`` enumeration class represents the upgrade invocation result for each cluster. **Warning:** This enumeration is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
     
    type Result_Res string

    const (
        // Upgrade is started. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
         Result_Res_STARTED Result_Res = "STARTED"
        // Upgrade is rejected. This implies pre-check failed when invoking upgrade of the cluster. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
         Result_Res_REJECTED Result_Res = "REJECTED"
    )

    func (r Result_Res) Result_Res() bool {
        switch r {
            case Result_Res_STARTED:
                return true
            case Result_Res_REJECTED:
                return true
            default:
                return false
        }
    }



// The ``UpgradeSpec`` class contains the specification required to upgrade a cluster. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type UpgradeSpec struct {
    DesiredVersion string
}






// The ``Summary`` class contains basic information about the cluster upgrade related information. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type Summary struct {
    Cluster string
    ClusterName string
    CurrentVersion string
    AvailableVersions []string
    LastUpgradedDate *time.Time
    DesiredVersion *string
    State State
}






// The ``Info`` class contains detailed information about the cluster upgrade status and related information. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type Info struct {
    CurrentVersion string
    AvailableVersions []string
    LastUpgradedDate *time.Time
    Messages []Message
    State State
    UpgradeStatus *UpgradeStatus
}






// The ``UpgradeStatus`` class contains detailed information about the cluster when upgraded is in progress. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type UpgradeStatus struct {
    DesiredVersion *string
    Messages []Message
    Progress *UpgradeProgress
}






// The ``UpgradeProgress`` class contains detailed information about the cluster upgrade progess. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type UpgradeProgress struct {
    Total int64
    Completed int64
    Message std.LocalizableMessage
}






// The ``Message`` class contains the information about the object configuration. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type Message struct {
    Severity Message_Severity
    Details std.LocalizableMessage
}




    
    // The ``Severity`` enumeration class represents the severity of the message. **Warning:** This enumeration is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
     
    type Message_Severity string

    const (
        // Informational message. This may be accompanied by vCenter event. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
         Message_Severity_INFO Message_Severity = "INFO"
        // Warning message. This may be accompanied by vCenter event. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
         Message_Severity_WARNING Message_Severity = "WARNING"
        // Error message. This is accompanied by vCenter event and/or alarm. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
         Message_Severity_ERROR Message_Severity = "ERROR"
    )

    func (s Message_Severity) Message_Severity() bool {
        switch s {
            case Message_Severity_INFO:
                return true
            case Message_Severity_WARNING:
                return true
            case Message_Severity_ERROR:
                return true
            default:
                return false
        }
    }






func upgradeInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    fields["spec"] = bindings.NewReferenceType(UpgradeSpecBindingType)
    fieldNameMap["cluster"] = "Cluster"
    fieldNameMap["spec"] = "Spec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func upgradeOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func upgradeRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    paramsTypeMap["spec"] = bindings.NewReferenceType(UpgradeSpecBindingType)
    paramsTypeMap["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    pathParams["cluster"] = "cluster"
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
      "/vcenter/namespace-management/software/clusters/{cluster}",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"Error": 500,"Unsupported": 400,"NotFound": 404,"Unauthenticated": 401,"Unauthorized": 403})
}


func upgradeMultipleInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["specs"] = bindings.NewMapType(bindings.NewIdType([]string {"ClusterComputeResource"}, ""), bindings.NewReferenceType(UpgradeSpecBindingType),reflect.TypeOf(map[string]UpgradeSpec{}))
    fieldNameMap["specs"] = "Specs"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func upgradeMultipleOutputType() bindings.BindingType {
    return bindings.NewMapType(bindings.NewIdType([]string {"ClusterComputeResource"}, ""), bindings.NewReferenceType(ResultBindingType),reflect.TypeOf(map[string]Result{}))
}

func upgradeMultipleRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["specs"] = bindings.NewMapType(bindings.NewIdType([]string {"ClusterComputeResource"}, ""), bindings.NewReferenceType(UpgradeSpecBindingType),reflect.TypeOf(map[string]UpgradeSpec{}))
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "specs",
      "POST",
      "/vcenter/namespace-management/software/clusters",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"Error": 500,"Unauthenticated": 401,"Unauthorized": 403})
}


func getInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    fieldNameMap["cluster"] = "Cluster"
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
    paramsTypeMap["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    paramsTypeMap["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    pathParams["cluster"] = "cluster"
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
      "/vcenter/namespace-management/software/clusters/{cluster}",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"NotFound": 404,"Error": 500,"Unauthenticated": 401,"Unauthorized": 403})
}


func listInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
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
    errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "GET",
      "/vcenter/namespace-management/software/clusters",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"Error": 500,"Unauthenticated": 401,"Unauthorized": 403})
}



func ResultBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["res"] = bindings.NewEnumType("com.vmware.vcenter.namespace_management.software.clusters.result.res", reflect.TypeOf(Result_Res(Result_Res_STARTED)))
    fieldNameMap["res"] = "Res"
    fields["exception"] = bindings.NewOptionalType(bindings.NewAnyErrorType())
    fieldNameMap["exception"] = "Exception"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("res",
        map[string][]bindings.FieldData {
            "REJECTED": []bindings.FieldData {
                 bindings.NewFieldData("exception", true),
            },
            "STARTED": []bindings.FieldData {},
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.vcenter.namespace_management.software.clusters.result",fields, reflect.TypeOf(Result{}), fieldNameMap, validators)
}

func UpgradeSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["desired_version"] = bindings.NewStringType()
    fieldNameMap["desired_version"] = "DesiredVersion"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.namespace_management.software.clusters.upgrade_spec",fields, reflect.TypeOf(UpgradeSpec{}), fieldNameMap, validators)
}

func SummaryBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    fieldNameMap["cluster"] = "Cluster"
    fields["cluster_name"] = bindings.NewIdType([]string {"ClusterComputeResource.name"}, "")
    fieldNameMap["cluster_name"] = "ClusterName"
    fields["current_version"] = bindings.NewStringType()
    fieldNameMap["current_version"] = "CurrentVersion"
    fields["available_versions"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
    fieldNameMap["available_versions"] = "AvailableVersions"
    fields["last_upgraded_date"] = bindings.NewOptionalType(bindings.NewDateTimeType())
    fieldNameMap["last_upgraded_date"] = "LastUpgradedDate"
    fields["desired_version"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["desired_version"] = "DesiredVersion"
    fields["state"] = bindings.NewEnumType("com.vmware.vcenter.namespace_management.software.clusters.state", reflect.TypeOf(State(State_PENDING)))
    fieldNameMap["state"] = "State"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.namespace_management.software.clusters.summary",fields, reflect.TypeOf(Summary{}), fieldNameMap, validators)
}

func InfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["current_version"] = bindings.NewStringType()
    fieldNameMap["current_version"] = "CurrentVersion"
    fields["available_versions"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
    fieldNameMap["available_versions"] = "AvailableVersions"
    fields["last_upgraded_date"] = bindings.NewOptionalType(bindings.NewDateTimeType())
    fieldNameMap["last_upgraded_date"] = "LastUpgradedDate"
    fields["messages"] = bindings.NewListType(bindings.NewReferenceType(MessageBindingType), reflect.TypeOf([]Message{}))
    fieldNameMap["messages"] = "Messages"
    fields["state"] = bindings.NewEnumType("com.vmware.vcenter.namespace_management.software.clusters.state", reflect.TypeOf(State(State_PENDING)))
    fieldNameMap["state"] = "State"
    fields["upgrade_status"] = bindings.NewOptionalType(bindings.NewReferenceType(UpgradeStatusBindingType))
    fieldNameMap["upgrade_status"] = "UpgradeStatus"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.namespace_management.software.clusters.info",fields, reflect.TypeOf(Info{}), fieldNameMap, validators)
}

func UpgradeStatusBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["desired_version"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["desired_version"] = "DesiredVersion"
    fields["messages"] = bindings.NewListType(bindings.NewReferenceType(MessageBindingType), reflect.TypeOf([]Message{}))
    fieldNameMap["messages"] = "Messages"
    fields["progress"] = bindings.NewOptionalType(bindings.NewReferenceType(UpgradeProgressBindingType))
    fieldNameMap["progress"] = "Progress"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.namespace_management.software.clusters.upgrade_status",fields, reflect.TypeOf(UpgradeStatus{}), fieldNameMap, validators)
}

func UpgradeProgressBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["total"] = bindings.NewIntegerType()
    fieldNameMap["total"] = "Total"
    fields["completed"] = bindings.NewIntegerType()
    fieldNameMap["completed"] = "Completed"
    fields["message"] = bindings.NewReferenceType(std.LocalizableMessageBindingType)
    fieldNameMap["message"] = "Message"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.namespace_management.software.clusters.upgrade_progress",fields, reflect.TypeOf(UpgradeProgress{}), fieldNameMap, validators)
}

func MessageBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["severity"] = bindings.NewEnumType("com.vmware.vcenter.namespace_management.software.clusters.message.severity", reflect.TypeOf(Message_Severity(Message_Severity_INFO)))
    fieldNameMap["severity"] = "Severity"
    fields["details"] = bindings.NewReferenceType(std.LocalizableMessageBindingType)
    fieldNameMap["details"] = "Details"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.namespace_management.software.clusters.message",fields, reflect.TypeOf(Message{}), fieldNameMap, validators)
}


