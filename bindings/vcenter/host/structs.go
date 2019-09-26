/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Host.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package host

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)

// The resource type for the vCenter Host.
const Host_RESOURCE_TYPE = "HostSystem"


// The ``ConnectionState`` enumeration class defines the connection status of a host.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type ConnectionState string

const (
    // Host is connected to the vCenter Server
     ConnectionState_CONNECTED ConnectionState = "CONNECTED"
    // Host is disconnected from the vCenter Server
     ConnectionState_DISCONNECTED ConnectionState = "DISCONNECTED"
    // VirtualCenter is not receiving heartbeats from the server. The state automatically changes to connected once heartbeats are received again.
     ConnectionState_NOT_RESPONDING ConnectionState = "NOT_RESPONDING"
)

func (c ConnectionState) ConnectionState() bool {
    switch c {
        case ConnectionState_CONNECTED:
            return true
        case ConnectionState_DISCONNECTED:
            return true
        case ConnectionState_NOT_RESPONDING:
            return true
        default:
            return false
    }
}




// The ``PowerState`` enumeration class defines the power states of a host.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type PowerState string

const (
    // The host is powered on. A host that is entering standby mode is also in this state.
     PowerState_POWERED_ON PowerState = "POWERED_ON"
    // The host was specifically powered off by the user through vCenter server. This state is not a cetain state, because after vCenter server issues the command to power off the host, the host might crash, or kill all the processes but fail to power off.
     PowerState_POWERED_OFF PowerState = "POWERED_OFF"
    // The host was specifically put in standby mode, either explicitly by the user, or automatically by DPM. This state is not a cetain state, because after VirtualCenter issues the command to put the host in standby state, the host might crash, or kill all the processes but fail to enter standby mode. A host that is exiting standby mode is also in this state.
     PowerState_STANDBY PowerState = "STANDBY"
)

func (p PowerState) PowerState() bool {
    switch p {
        case PowerState_POWERED_ON:
            return true
        case PowerState_POWERED_OFF:
            return true
        case PowerState_STANDBY:
            return true
        default:
            return false
    }
}





// The ``CreateSpec`` class defines the information used to create a host.
 type CreateSpec struct {
    Hostname string
    Port *int64
    UserName string
    Password string
    Folder *string
    ThumbprintVerification CreateSpec_ThumbprintVerification
    Thumbprint *string
    ForceAdd *bool
}




    
    // The ``ThumbprintVerification`` enumeration class defines the thumbprint verification schemes for a host's SSL certificate.
    //
    // <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
     
    type CreateSpec_ThumbprintVerification string

    const (
        // Accept the host's thumbprint without verifying it.
         CreateSpec_ThumbprintVerification_NONE CreateSpec_ThumbprintVerification = "NONE"
        // Host's SSL certificate verified by checking its thumbprint against the specified thumbprint.
         CreateSpec_ThumbprintVerification_THUMBPRINT CreateSpec_ThumbprintVerification = "THUMBPRINT"
    )

    func (t CreateSpec_ThumbprintVerification) CreateSpec_ThumbprintVerification() bool {
        switch t {
            case CreateSpec_ThumbprintVerification_NONE:
                return true
            case CreateSpec_ThumbprintVerification_THUMBPRINT:
                return true
            default:
                return false
        }
    }



// The ``FilterSpec`` class contains properties used to filter the results when listing hosts (see Host#list). If multiple properties are specified, only hosts matching all of the properties match the filter.
 type FilterSpec struct {
    Hosts map[string]bool
    Names map[string]bool
    Folders map[string]bool
    Datacenters map[string]bool
    Standalone *bool
    Clusters map[string]bool
    ConnectionStates map[ConnectionState]bool
}






// The ``Summary`` class contains commonly used information about a host in vCenter Server.
 type Summary struct {
    Host string
    Name string
    ConnectionState ConnectionState
    PowerState *PowerState
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
    return bindings.NewIdType([]string {"HostSystem"}, "")
}

func createRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"AlreadyExists": 400,"Error": 500,"InvalidArgument": 400,"InvalidElementType": 400,"NotFound": 404,"ResourceInUse": 400,"UnableToAllocateResource": 500,"Unauthenticated": 401,"Unsupported": 400,"ServiceUnavailable": 503,"Unauthorized": 403})
}


func deleteInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["host"] = bindings.NewIdType([]string {"HostSystem"}, "")
    fieldNameMap["host"] = "Host"
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
       map[string]int{"Error": 500,"NotFound": 404,"ResourceInUse": 400,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}


func listInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["filter"] = bindings.NewOptionalType(bindings.NewReferenceType(FilterSpecBindingType))
    fieldNameMap["filter"] = "Filter"
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
      "null",
      "",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"InvalidArgument": 400,"UnableToAllocateResource": 500,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}


func connectInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["host"] = bindings.NewIdType([]string {"HostSystem"}, "")
    fieldNameMap["host"] = "Host"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func connectOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func connectRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"AlreadyInDesiredState": 400,"Error": 500,"NotFound": 404,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}


func disconnectInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["host"] = bindings.NewIdType([]string {"HostSystem"}, "")
    fieldNameMap["host"] = "Host"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func disconnectOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func disconnectRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"AlreadyInDesiredState": 400,"Error": 500,"NotFound": 404,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}



func CreateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["hostname"] = bindings.NewStringType()
    fieldNameMap["hostname"] = "Hostname"
    fields["port"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["port"] = "Port"
    fields["user_name"] = bindings.NewStringType()
    fieldNameMap["user_name"] = "UserName"
    fields["password"] = bindings.NewSecretType()
    fieldNameMap["password"] = "Password"
    fields["folder"] = bindings.NewOptionalType(bindings.NewIdType([]string {"Folder"}, ""))
    fieldNameMap["folder"] = "Folder"
    fields["thumbprint_verification"] = bindings.NewEnumType("com.vmware.vcenter.host.create_spec.thumbprint_verification", reflect.TypeOf(CreateSpec_ThumbprintVerification(CreateSpec_ThumbprintVerification_NONE)))
    fieldNameMap["thumbprint_verification"] = "ThumbprintVerification"
    fields["thumbprint"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["thumbprint"] = "Thumbprint"
    fields["force_add"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["force_add"] = "ForceAdd"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("thumbprint_verification",
        map[string][]bindings.FieldData {
            "THUMBPRINT": []bindings.FieldData {
                 bindings.NewFieldData("thumbprint", true),
            },
            "NONE": []bindings.FieldData {},
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.vcenter.host.create_spec",fields, reflect.TypeOf(CreateSpec{}), fieldNameMap, validators)
}

func FilterSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["hosts"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string {"HostSystem"}, ""), reflect.TypeOf(map[string]bool{})))
    fieldNameMap["hosts"] = "Hosts"
    fields["names"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewStringType(), reflect.TypeOf(map[string]bool{})))
    fieldNameMap["names"] = "Names"
    fields["folders"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string {"Folder"}, ""), reflect.TypeOf(map[string]bool{})))
    fieldNameMap["folders"] = "Folders"
    fields["datacenters"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string {"Datacenter"}, ""), reflect.TypeOf(map[string]bool{})))
    fieldNameMap["datacenters"] = "Datacenters"
    fields["standalone"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["standalone"] = "Standalone"
    fields["clusters"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string {"ClusterComputeResource"}, ""), reflect.TypeOf(map[string]bool{})))
    fieldNameMap["clusters"] = "Clusters"
    fields["connection_states"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewEnumType("com.vmware.vcenter.host.connection_state", reflect.TypeOf(ConnectionState(ConnectionState_CONNECTED))), reflect.TypeOf(map[ConnectionState]bool{})))
    fieldNameMap["connection_states"] = "ConnectionStates"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.host.filter_spec",fields, reflect.TypeOf(FilterSpec{}), fieldNameMap, validators)
}

func SummaryBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["host"] = bindings.NewIdType([]string {"HostSystem"}, "")
    fieldNameMap["host"] = "Host"
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    fields["connection_state"] = bindings.NewEnumType("com.vmware.vcenter.host.connection_state", reflect.TypeOf(ConnectionState(ConnectionState_CONNECTED)))
    fieldNameMap["connection_state"] = "ConnectionState"
    fields["power_state"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.host.power_state", reflect.TypeOf(PowerState(PowerState_POWERED_ON))))
    fieldNameMap["power_state"] = "PowerState"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("connection_state",
        map[string][]bindings.FieldData {
            "CONNECTED": []bindings.FieldData {
                 bindings.NewFieldData("power_state", true),
            },
            "DISCONNECTED": []bindings.FieldData {},
            "NOT_RESPONDING": []bindings.FieldData {},
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.vcenter.host.summary",fields, reflect.TypeOf(Summary{}), fieldNameMap, validators)
}


