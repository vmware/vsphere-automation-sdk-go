/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: Service.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package vmon

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/vapi/std"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)



// The ``StartupType`` enumeration class defines valid Startup Type for services managed by vMon.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type ServiceStartupType string

const (
    // Service Startup type is Manual, thus issuing an explicit start on the service will start it.
	ServiceStartupType_MANUAL ServiceStartupType = "MANUAL"
    // Service Startup type is Automatic, thus during starting all services or issuing explicit start on the service will start it.
	ServiceStartupType_AUTOMATIC ServiceStartupType = "AUTOMATIC"
    // Service Startup type is Disabled, thus it will not start unless the startup type changes to manual or automatic.
	ServiceStartupType_DISABLED ServiceStartupType = "DISABLED"
)

func (s ServiceStartupType) ServiceStartupType() bool {
	switch s {
	case ServiceStartupType_MANUAL:
		return true
	case ServiceStartupType_AUTOMATIC:
		return true
	case ServiceStartupType_DISABLED:
		return true
	default:
		return false
	}
}


// The ``State`` enumeration class defines valid Run State for services.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type ServiceState string

const (
    // Service Run State is Starting, it is still not functional
	ServiceState_STARTING ServiceState = "STARTING"
    // Service Run State is Stopping, it is not functional
	ServiceState_STOPPING ServiceState = "STOPPING"
    // Service Run State is Started, it is fully functional
	ServiceState_STARTED ServiceState = "STARTED"
    // Service Run State is Stopped
	ServiceState_STOPPED ServiceState = "STOPPED"
)

func (s ServiceState) ServiceState() bool {
	switch s {
	case ServiceState_STARTING:
		return true
	case ServiceState_STOPPING:
		return true
	case ServiceState_STARTED:
		return true
	case ServiceState_STOPPED:
		return true
	default:
		return false
	}
}


// The ``Health`` enumeration class defines the possible values for health of a service.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type ServiceHealth string

const (
    // Service is in degraded state, it is not functional.
	ServiceHealth_DEGRADED ServiceHealth = "DEGRADED"
    // Service is in a healthy state and is fully functional.
	ServiceHealth_HEALTHY ServiceHealth = "HEALTHY"
    // Service is healthy with warnings.
	ServiceHealth_HEALTHY_WITH_WARNINGS ServiceHealth = "HEALTHY_WITH_WARNINGS"
)

func (h ServiceHealth) ServiceHealth() bool {
	switch h {
	case ServiceHealth_DEGRADED:
		return true
	case ServiceHealth_HEALTHY:
		return true
	case ServiceHealth_HEALTHY_WITH_WARNINGS:
		return true
	default:
		return false
	}
}


// The ``Info`` class contains information about a service.
type ServiceInfo struct {
    // Service name key. Can be used to lookup resource bundle
	NameKey string
    // Service description key. Can be used to lookup resource bundle
	DescriptionKey string
    // Startup Type.
	StartupType ServiceStartupType
    // Running State.
	State ServiceState
    // Health of service.
	Health *ServiceHealth
    // Localizable messages associated with the health of the service
	HealthMessages []std.LocalizableMessage
}

// The ``UpdateSpec`` class describes the changes to be made to the configuration of the service.
type ServiceUpdateSpec struct {
    // Startup Type
	StartupType *ServiceStartupType
}



func serviceStartInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service"] = bindings.NewIdType([]string{"com.vmware.appliance.vmon.Service"}, "")
	fieldNameMap["service"] = "Service"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func serviceStartOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func serviceStartRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["service"] = bindings.NewIdType([]string{"com.vmware.appliance.vmon.Service"}, "")
	fieldNameMap["service"] = "Service"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"",
		"null",
		"",
		resultHeaders,
		0,
		errorHeaders,
		map[string]int{"NotFound": 404,"NotAllowedInCurrentState": 400,"TimedOut": 500,"Error": 500})
}

func serviceStopInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service"] = bindings.NewIdType([]string{"com.vmware.appliance.vmon.Service"}, "")
	fieldNameMap["service"] = "Service"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func serviceStopOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func serviceStopRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["service"] = bindings.NewIdType([]string{"com.vmware.appliance.vmon.Service"}, "")
	fieldNameMap["service"] = "Service"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"",
		"null",
		"",
		resultHeaders,
		0,
		errorHeaders,
		map[string]int{"NotFound": 404,"Error": 500,"NotAllowedInCurrentState": 400})
}

func serviceRestartInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service"] = bindings.NewIdType([]string{"com.vmware.appliance.vmon.Service"}, "")
	fieldNameMap["service"] = "Service"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func serviceRestartOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func serviceRestartRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["service"] = bindings.NewIdType([]string{"com.vmware.appliance.vmon.Service"}, "")
	fieldNameMap["service"] = "Service"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"",
		"null",
		"",
		resultHeaders,
		0,
		errorHeaders,
		map[string]int{"NotFound": 404,"TimedOut": 500,"NotAllowedInCurrentState": 400,"Error": 500})
}

func serviceGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service"] = bindings.NewIdType([]string{"com.vmware.appliance.vmon.Service"}, "")
	fieldNameMap["service"] = "Service"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func serviceGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(ServiceInfoBindingType)
}

func serviceGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["service"] = bindings.NewIdType([]string{"com.vmware.appliance.vmon.Service"}, "")
	fieldNameMap["service"] = "Service"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"",
		"null",
		"",
		resultHeaders,
		0,
		errorHeaders,
		map[string]int{"NotFound": 404,"Error": 500})
}

func serviceUpdateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service"] = bindings.NewIdType([]string{"com.vmware.appliance.vmon.Service"}, "")
	fields["spec"] = bindings.NewReferenceType(ServiceUpdateSpecBindingType)
	fieldNameMap["service"] = "Service"
	fieldNameMap["spec"] = "Spec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func serviceUpdateOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func serviceUpdateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["service"] = bindings.NewIdType([]string{"com.vmware.appliance.vmon.Service"}, "")
	fields["spec"] = bindings.NewReferenceType(ServiceUpdateSpecBindingType)
	fieldNameMap["service"] = "Service"
	fieldNameMap["spec"] = "Spec"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"",
		"null",
		"",
		resultHeaders,
		0,
		errorHeaders,
		map[string]int{"NotFound": 404,"Error": 500,"NotAllowedInCurrentState": 400})
}

func serviceListDetailsInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func serviceListDetailsOutputType() bindings.BindingType {
	return bindings.NewMapType(bindings.NewIdType([]string{"com.vmware.appliance.vmon.Service"}, ""), bindings.NewReferenceType(ServiceInfoBindingType),reflect.TypeOf(map[string]ServiceInfo{}))
}

func serviceListDetailsRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"",
		"null",
		"",
		resultHeaders,
		0,
		errorHeaders,
		map[string]int{"Error": 500})
}


func ServiceInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["name_key"] = bindings.NewStringType()
	fieldNameMap["name_key"] = "NameKey"
	fields["description_key"] = bindings.NewStringType()
	fieldNameMap["description_key"] = "DescriptionKey"
	fields["startup_type"] = bindings.NewEnumType("com.vmware.appliance.vmon.service.startup_type", reflect.TypeOf(ServiceStartupType(ServiceStartupType_MANUAL)))
	fieldNameMap["startup_type"] = "StartupType"
	fields["state"] = bindings.NewEnumType("com.vmware.appliance.vmon.service.state", reflect.TypeOf(ServiceState(ServiceState_STARTING)))
	fieldNameMap["state"] = "State"
	fields["health"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.appliance.vmon.service.health", reflect.TypeOf(ServiceHealth(ServiceHealth_DEGRADED))))
	fieldNameMap["health"] = "Health"
	fields["health_messages"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(std.LocalizableMessageBindingType), reflect.TypeOf([]std.LocalizableMessage{})))
	fieldNameMap["health_messages"] = "HealthMessages"
	var validators = []bindings.Validator{}
	uv1 := bindings.NewUnionValidator("state",
		map[string][]bindings.FieldData{
			"STARTED": []bindings.FieldData{
				bindings.NewFieldData("health", true),
				bindings.NewFieldData("health_messages", true),
			},
			"STARTING": []bindings.FieldData{},
			"STOPPING": []bindings.FieldData{},
			"STOPPED": []bindings.FieldData{},
		},
	)
	validators = append(validators, uv1)
	return bindings.NewStructType("com.vmware.appliance.vmon.service.info", fields, reflect.TypeOf(ServiceInfo{}), fieldNameMap, validators)
}

func ServiceUpdateSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["startup_type"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.appliance.vmon.service.startup_type", reflect.TypeOf(ServiceStartupType(ServiceStartupType_MANUAL))))
	fieldNameMap["startup_type"] = "StartupType"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.appliance.vmon.service.update_spec", fields, reflect.TypeOf(ServiceUpdateSpec{}), fieldNameMap, validators)
}

