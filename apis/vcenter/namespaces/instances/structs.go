/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Instances.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package instances

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/vapi/std"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/apis/vcenter/namespaces/access"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)

// The resource type for namespace. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
const Instances_RESOURCE_TYPE = "com.vmware.vcenter.namespaces.Instance"


// The ``ConfigStatus`` enumeration class describes the status of reaching the desired state configuration for the namespace. **Warning:** This enumeration is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type ConfigStatus string

const (
    // The configuration is being applied to the namespace. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     ConfigStatus_CONFIGURING ConfigStatus = "CONFIGURING"
    // The configuration is being removed and namespace is being deleted. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     ConfigStatus_REMOVING ConfigStatus = "REMOVING"
    // The namespace is configured correctly. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     ConfigStatus_RUNNING ConfigStatus = "RUNNING"
    // Failed to apply the configuration to the namespace, user intervention needed. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     ConfigStatus_ERROR ConfigStatus = "ERROR"
)

func (c ConfigStatus) ConfigStatus() bool {
    switch c {
        case ConfigStatus_CONFIGURING:
            return true
        case ConfigStatus_REMOVING:
            return true
        case ConfigStatus_RUNNING:
            return true
        case ConfigStatus_ERROR:
            return true
        default:
            return false
    }
}





// The ``Access`` class contains the access control information for a subject on a namespace. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type Access struct {
    SubjectType access.SubjectType
    Subject string
    Domain string
    Role access.Role
}






// The ``StorageSpec`` class contains the specification required to configure storage associated with a namespace. Information in this class will result in storage quotas on the Kubernetes namespace. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type StorageSpec struct {
    Policy string
    Limit *int64
}






// The ``Message`` class contains the information about the object configuration. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type Message struct {
    Severity Message_MessageSeverity
    Details *std.LocalizableMessage
}




    
    // The ``MessageSeverity`` enumeration class represents the severity of the message. **Warning:** This enumeration is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
     
    type Message_MessageSeverity string

    const (
        // Informational message. This may be accompanied by vCenter event. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
         Message_MessageSeverity_INFO Message_MessageSeverity = "INFO"
        // Warning message. This may be accompanied by vCenter event. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
         Message_MessageSeverity_WARNING Message_MessageSeverity = "WARNING"
        // Error message. This is accompanied by vCenter event and/or alarm. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
         Message_MessageSeverity_ERROR Message_MessageSeverity = "ERROR"
    )

    func (m Message_MessageSeverity) Message_MessageSeverity() bool {
        switch m {
            case Message_MessageSeverity_INFO:
                return true
            case Message_MessageSeverity_WARNING:
                return true
            case Message_MessageSeverity_ERROR:
                return true
            default:
                return false
        }
    }



// The ``Stats`` class contains the basic runtime statistics about the namespace. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type Stats struct {
    CpuUsed int64
    MemoryUsed int64
    StorageUsed int64
}






// The ``Summary`` class contains basic information about the namespace. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type Summary struct {
    Namespace string
    Description string
    Cluster string
    ConfigStatus ConfigStatus
    Stats Stats
}






// The ``Info`` class contains detailed information about the namespace. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type Info struct {
    Cluster string
    ConfigStatus ConfigStatus
    Messages []Message
    Stats Stats
    Description string
    ResourceSpec *data.StructValue
    AccessList []Access
    StorageSpecs []StorageSpec
}






// The ``UpdateSpec`` class contains the specification required to update the configuration on the namespace. This class is applied partially, and only the specified fields will replace or modify their existing counterparts. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type UpdateSpec struct {
    Description *string
    ResourceSpec *data.StructValue
    AccessList []Access
    StorageSpecs []StorageSpec
}






// The ``SetSpec`` class contains the specification required to set a new configuration on the namespace. This class is applied in entirety, replacing the current specification fully. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type SetSpec struct {
    Description *string
    ResourceSpec *data.StructValue
    AccessList []Access
    StorageSpecs []StorageSpec
}






// The ``CreateSpec`` class contains the specification required to set up a namespace. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type CreateSpec struct {
    Namespace string
    Cluster string
    Description *string
    ResourceSpec *data.StructValue
    AccessList []Access
    StorageSpecs []StorageSpec
}









func createInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["spec"] = bindings.NewReferenceType(CreateSpecBindingType)
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
    paramsTypeMap["spec"] = bindings.NewReferenceType(CreateSpecBindingType)
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
      "/vcenter/namespaces/instances",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"AlreadyExists": 400,"Error": 500,"InvalidArgument": 400,"NotAllowedInCurrentState": 400,"NotFound": 404,"Unsupported": 400,"Unauthenticated": 401,"Unauthorized": 403})
}


func deleteInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["namespace"] = bindings.NewIdType([]string {"com.vmware.vcenter.namespaces.Instance"}, "")
    fieldNameMap["namespace"] = "Namespace"
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
    paramsTypeMap["namespace"] = bindings.NewIdType([]string {"com.vmware.vcenter.namespaces.Instance"}, "")
    paramsTypeMap["namespace"] = bindings.NewIdType([]string {"com.vmware.vcenter.namespaces.Instance"}, "")
    pathParams["namespace"] = "namespace"
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "DELETE",
      "/vcenter/namespaces/instances/{namespace}",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"Error": 500,"NotFound": 404,"Unauthenticated": 401,"Unauthorized": 403})
}


func getInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["namespace"] = bindings.NewIdType([]string {"com.vmware.vcenter.namespaces.Instance"}, "")
    fieldNameMap["namespace"] = "Namespace"
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
    paramsTypeMap["namespace"] = bindings.NewIdType([]string {"com.vmware.vcenter.namespaces.Instance"}, "")
    paramsTypeMap["namespace"] = bindings.NewIdType([]string {"com.vmware.vcenter.namespaces.Instance"}, "")
    pathParams["namespace"] = "namespace"
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
      "/vcenter/namespaces/instances/{namespace}",
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
      "/vcenter/namespaces/instances",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"Error": 500,"Unauthenticated": 401,"Unauthorized": 403})
}


func setInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["namespace"] = bindings.NewIdType([]string {"com.vmware.vcenter.namespaces.Instance"}, "")
    fields["spec"] = bindings.NewReferenceType(SetSpecBindingType)
    fieldNameMap["namespace"] = "Namespace"
    fieldNameMap["spec"] = "Spec"
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
    paramsTypeMap["spec"] = bindings.NewReferenceType(SetSpecBindingType)
    paramsTypeMap["namespace"] = bindings.NewIdType([]string {"com.vmware.vcenter.namespaces.Instance"}, "")
    paramsTypeMap["namespace"] = bindings.NewIdType([]string {"com.vmware.vcenter.namespaces.Instance"}, "")
    pathParams["namespace"] = "namespace"
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "spec",
      "PUT",
      "/vcenter/namespaces/instances/{namespace}",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"Error": 500,"InvalidArgument": 400,"NotAllowedInCurrentState": 400,"NotFound": 404,"Unauthenticated": 401,"Unauthorized": 403})
}


func updateInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["namespace"] = bindings.NewIdType([]string {"com.vmware.vcenter.namespaces.Instance"}, "")
    fields["spec"] = bindings.NewReferenceType(UpdateSpecBindingType)
    fieldNameMap["namespace"] = "Namespace"
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
    paramsTypeMap["spec"] = bindings.NewReferenceType(UpdateSpecBindingType)
    paramsTypeMap["namespace"] = bindings.NewIdType([]string {"com.vmware.vcenter.namespaces.Instance"}, "")
    paramsTypeMap["namespace"] = bindings.NewIdType([]string {"com.vmware.vcenter.namespaces.Instance"}, "")
    pathParams["namespace"] = "namespace"
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "spec",
      "PATCH",
      "/vcenter/namespaces/instances/{namespace}",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"Error": 500,"InvalidArgument": 400,"NotAllowedInCurrentState": 400,"NotFound": 404,"Unauthenticated": 401,"Unauthorized": 403})
}



func AccessBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["subject_type"] = bindings.NewEnumType("com.vmware.vcenter.namespaces.access.subject_type", reflect.TypeOf(access.SubjectType(access.SubjectType_USER)))
    fieldNameMap["subject_type"] = "SubjectType"
    fields["subject"] = bindings.NewStringType()
    fieldNameMap["subject"] = "Subject"
    fields["domain"] = bindings.NewStringType()
    fieldNameMap["domain"] = "Domain"
    fields["role"] = bindings.NewEnumType("com.vmware.vcenter.namespaces.access.role", reflect.TypeOf(access.Role(access.Role_EDIT)))
    fieldNameMap["role"] = "Role"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.namespaces.instances.access",fields, reflect.TypeOf(Access{}), fieldNameMap, validators)
}

func StorageSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["policy"] = bindings.NewIdType([]string {"SpsStorageProfile"}, "")
    fieldNameMap["policy"] = "Policy"
    fields["limit"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["limit"] = "Limit"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.namespaces.instances.storage_spec",fields, reflect.TypeOf(StorageSpec{}), fieldNameMap, validators)
}

func MessageBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["severity"] = bindings.NewEnumType("com.vmware.vcenter.namespaces.instances.message.message_severity", reflect.TypeOf(Message_MessageSeverity(Message_MessageSeverity_INFO)))
    fieldNameMap["severity"] = "Severity"
    fields["details"] = bindings.NewOptionalType(bindings.NewReferenceType(std.LocalizableMessageBindingType))
    fieldNameMap["details"] = "Details"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.namespaces.instances.message",fields, reflect.TypeOf(Message{}), fieldNameMap, validators)
}

func StatsBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["cpu_used"] = bindings.NewIntegerType()
    fieldNameMap["cpu_used"] = "CpuUsed"
    fields["memory_used"] = bindings.NewIntegerType()
    fieldNameMap["memory_used"] = "MemoryUsed"
    fields["storage_used"] = bindings.NewIntegerType()
    fieldNameMap["storage_used"] = "StorageUsed"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.namespaces.instances.stats",fields, reflect.TypeOf(Stats{}), fieldNameMap, validators)
}

func SummaryBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["namespace"] = bindings.NewIdType([]string {"com.vmware.vcenter.namespaces.Instance"}, "")
    fieldNameMap["namespace"] = "Namespace"
    fields["description"] = bindings.NewStringType()
    fieldNameMap["description"] = "Description"
    fields["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    fieldNameMap["cluster"] = "Cluster"
    fields["config_status"] = bindings.NewEnumType("com.vmware.vcenter.namespaces.instances.config_status", reflect.TypeOf(ConfigStatus(ConfigStatus_CONFIGURING)))
    fieldNameMap["config_status"] = "ConfigStatus"
    fields["stats"] = bindings.NewReferenceType(StatsBindingType)
    fieldNameMap["stats"] = "Stats"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.namespaces.instances.summary",fields, reflect.TypeOf(Summary{}), fieldNameMap, validators)
}

func InfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    fieldNameMap["cluster"] = "Cluster"
    fields["config_status"] = bindings.NewEnumType("com.vmware.vcenter.namespaces.instances.config_status", reflect.TypeOf(ConfigStatus(ConfigStatus_CONFIGURING)))
    fieldNameMap["config_status"] = "ConfigStatus"
    fields["messages"] = bindings.NewListType(bindings.NewReferenceType(MessageBindingType), reflect.TypeOf([]Message{}))
    fieldNameMap["messages"] = "Messages"
    fields["stats"] = bindings.NewReferenceType(StatsBindingType)
    fieldNameMap["stats"] = "Stats"
    fields["description"] = bindings.NewStringType()
    fieldNameMap["description"] = "Description"
    fields["resource_spec"] = bindings.NewOptionalType(bindings.NewDynamicStructType(nil, bindings.JSONRPC))
    fieldNameMap["resource_spec"] = "ResourceSpec"
    fields["access_list"] = bindings.NewListType(bindings.NewReferenceType(AccessBindingType), reflect.TypeOf([]Access{}))
    fieldNameMap["access_list"] = "AccessList"
    fields["storage_specs"] = bindings.NewListType(bindings.NewReferenceType(StorageSpecBindingType), reflect.TypeOf([]StorageSpec{}))
    fieldNameMap["storage_specs"] = "StorageSpecs"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.namespaces.instances.info",fields, reflect.TypeOf(Info{}), fieldNameMap, validators)
}

func UpdateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["description"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["description"] = "Description"
    fields["resource_spec"] = bindings.NewOptionalType(bindings.NewDynamicStructType(nil, bindings.JSONRPC))
    fieldNameMap["resource_spec"] = "ResourceSpec"
    fields["access_list"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(AccessBindingType), reflect.TypeOf([]Access{})))
    fieldNameMap["access_list"] = "AccessList"
    fields["storage_specs"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(StorageSpecBindingType), reflect.TypeOf([]StorageSpec{})))
    fieldNameMap["storage_specs"] = "StorageSpecs"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.namespaces.instances.update_spec",fields, reflect.TypeOf(UpdateSpec{}), fieldNameMap, validators)
}

func SetSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["description"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["description"] = "Description"
    fields["resource_spec"] = bindings.NewOptionalType(bindings.NewDynamicStructType(nil, bindings.JSONRPC))
    fieldNameMap["resource_spec"] = "ResourceSpec"
    fields["access_list"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(AccessBindingType), reflect.TypeOf([]Access{})))
    fieldNameMap["access_list"] = "AccessList"
    fields["storage_specs"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(StorageSpecBindingType), reflect.TypeOf([]StorageSpec{})))
    fieldNameMap["storage_specs"] = "StorageSpecs"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.namespaces.instances.set_spec",fields, reflect.TypeOf(SetSpec{}), fieldNameMap, validators)
}

func CreateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["namespace"] = bindings.NewIdType([]string {"com.vmware.vcenter.namespaces.Instance"}, "")
    fieldNameMap["namespace"] = "Namespace"
    fields["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    fieldNameMap["cluster"] = "Cluster"
    fields["description"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["description"] = "Description"
    fields["resource_spec"] = bindings.NewOptionalType(bindings.NewDynamicStructType(nil, bindings.JSONRPC))
    fieldNameMap["resource_spec"] = "ResourceSpec"
    fields["access_list"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(AccessBindingType), reflect.TypeOf([]Access{})))
    fieldNameMap["access_list"] = "AccessList"
    fields["storage_specs"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(StorageSpecBindingType), reflect.TypeOf([]StorageSpec{})))
    fieldNameMap["storage_specs"] = "StorageSpecs"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.namespaces.instances.create_spec",fields, reflect.TypeOf(CreateSpec{}), fieldNameMap, validators)
}


