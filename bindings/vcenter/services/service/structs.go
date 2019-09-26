/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Service.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package service

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/vapi/std"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)



// The ``StartupType`` enumeration class defines valid Startup Type for vCenter Server services.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type StartupType string

const (
    // Service Startup type is Manual, thus issuing an explicit start on the service will start it.
     StartupType_MANUAL StartupType = "MANUAL"
    // Service Startup type is Automatic, thus during starting all services or issuing explicit start on the service will start it.
     StartupType_AUTOMATIC StartupType = "AUTOMATIC"
    // Service Startup type is Disabled, thus it will not start unless the startup type changes to manual or automatic.
     StartupType_DISABLED StartupType = "DISABLED"
)

func (s StartupType) StartupType() bool {
    switch s {
        case StartupType_MANUAL:
            return true
        case StartupType_AUTOMATIC:
            return true
        case StartupType_DISABLED:
            return true
        default:
            return false
    }
}




// The ``State`` enumeration class defines valid Run State for services.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type State string

const (
    // Service Run State is Starting, it is still not functional
     State_STARTING State = "STARTING"
    // Service Run State is Stopping, it is not functional
     State_STOPPING State = "STOPPING"
    // Service Run State is Started, it is fully functional
     State_STARTED State = "STARTED"
    // Service Run State is Stopped
     State_STOPPED State = "STOPPED"
)

func (s State) State() bool {
    switch s {
        case State_STARTING:
            return true
        case State_STOPPING:
            return true
        case State_STARTED:
            return true
        case State_STOPPED:
            return true
        default:
            return false
    }
}




// The ``Health`` enumeration class defines the possible values for health of a service.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type Health string

const (
    // Service is in degraded state, it is not functional.
     Health_DEGRADED Health = "DEGRADED"
    // Service is in a healthy state and is fully functional.
     Health_HEALTHY Health = "HEALTHY"
    // Service is healthy with warnings.
     Health_HEALTHY_WITH_WARNINGS Health = "HEALTHY_WITH_WARNINGS"
)

func (h Health) Health() bool {
    switch h {
        case Health_DEGRADED:
            return true
        case Health_HEALTHY:
            return true
        case Health_HEALTHY_WITH_WARNINGS:
            return true
        default:
            return false
    }
}





// The ``Info`` class contains information about a service.
 type Info struct {
    NameKey string
    DescriptionKey string
    StartupType StartupType
    State State
    Health *Health
    HealthMessages []std.LocalizableMessage
}






// The ``UpdateSpec`` class describes the changes to be made to the configuration of the service.
 type UpdateSpec struct {
    StartupType *StartupType
}









func startInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["service"] = bindings.NewIdType([]string {"com.vmware.vcenter.services.Service"}, "")
    fieldNameMap["service"] = "Service"
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
       map[string]int{"NotFound": 404,"NotAllowedInCurrentState": 400,"TimedOut": 500,"Error": 500})
}


func stopInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["service"] = bindings.NewIdType([]string {"com.vmware.vcenter.services.Service"}, "")
    fieldNameMap["service"] = "Service"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func stopOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func stopRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"NotFound": 404,"Error": 500,"NotAllowedInCurrentState": 400})
}


func restartInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["service"] = bindings.NewIdType([]string {"com.vmware.vcenter.services.Service"}, "")
    fieldNameMap["service"] = "Service"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func restartOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func restartRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"NotFound": 404,"TimedOut": 500,"NotAllowedInCurrentState": 400,"Error": 500})
}


func getInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["service"] = bindings.NewIdType([]string {"com.vmware.vcenter.services.Service"}, "")
    fieldNameMap["service"] = "Service"
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


func updateInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["service"] = bindings.NewIdType([]string {"com.vmware.vcenter.services.Service"}, "")
    fields["spec"] = bindings.NewReferenceType(UpdateSpecBindingType)
    fieldNameMap["service"] = "Service"
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
       map[string]int{"NotFound": 404,"Error": 500,"NotAllowedInCurrentState": 400})
}


func listDetailsInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func listDetailsOutputType() bindings.BindingType {
    return bindings.NewMapType(bindings.NewIdType([]string {"com.vmware.vcenter.services.Service"}, ""), bindings.NewReferenceType(InfoBindingType),reflect.TypeOf(map[string]Info{}))
}

func listDetailsRestMetadata() protocol.OperationRestMetadata {
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



func InfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["name_key"] = bindings.NewStringType()
    fieldNameMap["name_key"] = "NameKey"
    fields["description_key"] = bindings.NewStringType()
    fieldNameMap["description_key"] = "DescriptionKey"
    fields["startup_type"] = bindings.NewEnumType("com.vmware.vcenter.services.service.startup_type", reflect.TypeOf(StartupType(StartupType_MANUAL)))
    fieldNameMap["startup_type"] = "StartupType"
    fields["state"] = bindings.NewEnumType("com.vmware.vcenter.services.service.state", reflect.TypeOf(State(State_STARTING)))
    fieldNameMap["state"] = "State"
    fields["health"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.services.service.health", reflect.TypeOf(Health(Health_DEGRADED))))
    fieldNameMap["health"] = "Health"
    fields["health_messages"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(std.LocalizableMessageBindingType), reflect.TypeOf([]std.LocalizableMessage{})))
    fieldNameMap["health_messages"] = "HealthMessages"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("state",
        map[string][]bindings.FieldData {
            "STARTED": []bindings.FieldData {
                 bindings.NewFieldData("health", true),
                 bindings.NewFieldData("health_messages", true),
            },
            "STARTING": []bindings.FieldData {},
            "STOPPING": []bindings.FieldData {},
            "STOPPED": []bindings.FieldData {},
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.vcenter.services.service.info",fields, reflect.TypeOf(Info{}), fieldNameMap, validators)
}

func UpdateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["startup_type"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.services.service.startup_type", reflect.TypeOf(StartupType(StartupType_MANUAL))))
    fieldNameMap["startup_type"] = "StartupType"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.services.service.update_spec",fields, reflect.TypeOf(UpdateSpec{}), fieldNameMap, validators)
}


