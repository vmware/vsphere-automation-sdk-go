/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: Host.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package vcenter

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
type HostConnectionState string

const (
    // Host is connected to the vCenter Server
	HostConnectionState_CONNECTED HostConnectionState = "CONNECTED"
    // Host is disconnected from the vCenter Server
	HostConnectionState_DISCONNECTED HostConnectionState = "DISCONNECTED"
    // VirtualCenter is not receiving heartbeats from the server. The state automatically changes to connected once heartbeats are received again.
	HostConnectionState_NOT_RESPONDING HostConnectionState = "NOT_RESPONDING"
)

func (c HostConnectionState) HostConnectionState() bool {
	switch c {
	case HostConnectionState_CONNECTED:
		return true
	case HostConnectionState_DISCONNECTED:
		return true
	case HostConnectionState_NOT_RESPONDING:
		return true
	default:
		return false
	}
}


// The ``PowerState`` enumeration class defines the power states of a host.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type HostPowerState string

const (
    // The host is powered on. A host that is entering standby mode is also in this state.
	HostPowerState_POWERED_ON HostPowerState = "POWERED_ON"
    // The host was specifically powered off by the user through vCenter server. This state is not a cetain state, because after vCenter server issues the command to power off the host, the host might crash, or kill all the processes but fail to power off.
	HostPowerState_POWERED_OFF HostPowerState = "POWERED_OFF"
    // The host was specifically put in standby mode, either explicitly by the user, or automatically by DPM. This state is not a cetain state, because after VirtualCenter issues the command to put the host in standby state, the host might crash, or kill all the processes but fail to enter standby mode. A host that is exiting standby mode is also in this state.
	HostPowerState_STANDBY HostPowerState = "STANDBY"
)

func (p HostPowerState) HostPowerState() bool {
	switch p {
	case HostPowerState_POWERED_ON:
		return true
	case HostPowerState_POWERED_OFF:
		return true
	case HostPowerState_STANDBY:
		return true
	default:
		return false
	}
}


// The ``CreateSpec`` class defines the information used to create a host.
type HostCreateSpec struct {
    // The IP address or DNS resolvable name of the host.
	Hostname string
    // The port of the host.
	Port *int64
    // The administrator account on the host.
	UserName string
    // The password for the administrator account on the host.
	Password string
    // Host and cluster folder in which the new standalone host should be created.
	Folder *string
    // Type of host's SSL certificate verification to be done.
	ThumbprintVerification HostCreateSpecThumbprintVerification
    // The thumbprint of the SSL certificate, which the host is expected to have. The thumbprint is always computed using the SHA1 hash and is the string representation of that hash in the format: xx:xx:xx:xx:xx:xx:xx:xx:xx:xx:xx:xx:xx:xx:xx:xx:xx:xx:xx:xx where, 'x' represents a hexadecimal digit.
	Thumbprint *string
    // Whether host should be added to the vCenter Server even if it is being managed by another vCenter Server. The original vCenterServer loses connection to the host.
	ForceAdd *bool
}

// The ``ThumbprintVerification`` enumeration class defines the thumbprint verification schemes for a host's SSL certificate.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type HostCreateSpecThumbprintVerification string

const (
    // Accept the host's thumbprint without verifying it.
	HostCreateSpecThumbprintVerification_NONE HostCreateSpecThumbprintVerification = "NONE"
    // Host's SSL certificate verified by checking its thumbprint against the specified thumbprint.
	HostCreateSpecThumbprintVerification_THUMBPRINT HostCreateSpecThumbprintVerification = "THUMBPRINT"
)

func (t HostCreateSpecThumbprintVerification) HostCreateSpecThumbprintVerification() bool {
	switch t {
	case HostCreateSpecThumbprintVerification_NONE:
		return true
	case HostCreateSpecThumbprintVerification_THUMBPRINT:
		return true
	default:
		return false
	}
}


// The ``FilterSpec`` class contains properties used to filter the results when listing hosts (see Host#list). If multiple properties are specified, only hosts matching all of the properties match the filter.
type HostFilterSpec struct {
    // Identifiers of hosts that can match the filter.
	Hosts map[string]bool
    // Names that hosts must have to match the filter (see HostSummary#name).
	Names map[string]bool
    // Folders that must contain the hosts for the hosts to match the filter.
	Folders map[string]bool
    // Datacenters that must contain the hosts for the hosts to match the filter.
	Datacenters map[string]bool
    // If true, only hosts that are not part of a cluster can match the filter, and if false, only hosts that are are part of a cluster can match the filter.
	Standalone *bool
    // Clusters that must contain the hosts for the hosts to match the filter.
	Clusters map[string]bool
    // Connection states that a host must be in to match the filter (see HostSummary#connectionState.
	ConnectionStates map[HostConnectionState]bool
}

// The ``Summary`` class contains commonly used information about a host in vCenter Server.
type HostSummary struct {
    // Identifier of the host.
	Host string
    // Name of the host.
	Name string
    // Connection status of the host
	ConnectionState HostConnectionState
    // Power state of the host
	PowerState *HostPowerState
}



func hostCreateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["spec"] = bindings.NewReferenceType(HostCreateSpecBindingType)
	fieldNameMap["spec"] = "Spec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func hostCreateOutputType() bindings.BindingType {
	return bindings.NewIdType([]string{"HostSystem"}, "")
}

func hostCreateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["spec"] = bindings.NewReferenceType(HostCreateSpecBindingType)
	fieldNameMap["spec"] = "Spec"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
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
		map[string]int{"AlreadyExists": 400,"Error": 500,"InvalidArgument": 400,"InvalidElementType": 400,"NotFound": 404,"ResourceInUse": 400,"UnableToAllocateResource": 500,"Unauthenticated": 401,"Unsupported": 400,"ServiceUnavailable": 503,"Unauthorized": 403})
}

func hostDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["host"] = bindings.NewIdType([]string{"HostSystem"}, "")
	fieldNameMap["host"] = "Host"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func hostDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func hostDeleteRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["host"] = bindings.NewIdType([]string{"HostSystem"}, "")
	fieldNameMap["host"] = "Host"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
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
		map[string]int{"Error": 500,"NotFound": 404,"ResourceInUse": 400,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}

func hostListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["filter"] = bindings.NewOptionalType(bindings.NewReferenceType(HostFilterSpecBindingType))
	fieldNameMap["filter"] = "Filter"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func hostListOutputType() bindings.BindingType {
	return bindings.NewListType(bindings.NewReferenceType(HostSummaryBindingType), reflect.TypeOf([]HostSummary{}))
}

func hostListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["filter"] = bindings.NewOptionalType(bindings.NewReferenceType(HostFilterSpecBindingType))
	fieldNameMap["filter"] = "Filter"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
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
		map[string]int{"InvalidArgument": 400,"UnableToAllocateResource": 500,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}

func hostConnectInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["host"] = bindings.NewIdType([]string{"HostSystem"}, "")
	fieldNameMap["host"] = "Host"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func hostConnectOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func hostConnectRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["host"] = bindings.NewIdType([]string{"HostSystem"}, "")
	fieldNameMap["host"] = "Host"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
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
		map[string]int{"AlreadyInDesiredState": 400,"Error": 500,"NotFound": 404,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}

func hostDisconnectInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["host"] = bindings.NewIdType([]string{"HostSystem"}, "")
	fieldNameMap["host"] = "Host"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func hostDisconnectOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func hostDisconnectRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["host"] = bindings.NewIdType([]string{"HostSystem"}, "")
	fieldNameMap["host"] = "Host"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
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
		map[string]int{"AlreadyInDesiredState": 400,"Error": 500,"NotFound": 404,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}


func HostCreateSpecBindingType() bindings.BindingType {
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
	fields["folder"] = bindings.NewOptionalType(bindings.NewIdType([]string{"Folder"}, ""))
	fieldNameMap["folder"] = "Folder"
	fields["thumbprint_verification"] = bindings.NewEnumType("com.vmware.vcenter.host.create_spec.thumbprint_verification", reflect.TypeOf(HostCreateSpecThumbprintVerification(HostCreateSpecThumbprintVerification_NONE)))
	fieldNameMap["thumbprint_verification"] = "ThumbprintVerification"
	fields["thumbprint"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["thumbprint"] = "Thumbprint"
	fields["force_add"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["force_add"] = "ForceAdd"
	var validators = []bindings.Validator{}
	uv1 := bindings.NewUnionValidator("thumbprint_verification",
		map[string][]bindings.FieldData{
			"THUMBPRINT": []bindings.FieldData{
				bindings.NewFieldData("thumbprint", true),
			},
			"NONE": []bindings.FieldData{},
		},
	)
	validators = append(validators, uv1)
	return bindings.NewStructType("com.vmware.vcenter.host.create_spec", fields, reflect.TypeOf(HostCreateSpec{}), fieldNameMap, validators)
}

func HostFilterSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["hosts"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string{"HostSystem"}, ""), reflect.TypeOf(map[string]bool{})))
	fieldNameMap["hosts"] = "Hosts"
	fields["names"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewStringType(), reflect.TypeOf(map[string]bool{})))
	fieldNameMap["names"] = "Names"
	fields["folders"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string{"Folder"}, ""), reflect.TypeOf(map[string]bool{})))
	fieldNameMap["folders"] = "Folders"
	fields["datacenters"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string{"Datacenter"}, ""), reflect.TypeOf(map[string]bool{})))
	fieldNameMap["datacenters"] = "Datacenters"
	fields["standalone"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["standalone"] = "Standalone"
	fields["clusters"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string{"ClusterComputeResource"}, ""), reflect.TypeOf(map[string]bool{})))
	fieldNameMap["clusters"] = "Clusters"
	fields["connection_states"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewEnumType("com.vmware.vcenter.host.connection_state", reflect.TypeOf(HostConnectionState(HostConnectionState_CONNECTED))), reflect.TypeOf(map[HostConnectionState]bool{})))
	fieldNameMap["connection_states"] = "ConnectionStates"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.host.filter_spec", fields, reflect.TypeOf(HostFilterSpec{}), fieldNameMap, validators)
}

func HostSummaryBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["host"] = bindings.NewIdType([]string{"HostSystem"}, "")
	fieldNameMap["host"] = "Host"
	fields["name"] = bindings.NewStringType()
	fieldNameMap["name"] = "Name"
	fields["connection_state"] = bindings.NewEnumType("com.vmware.vcenter.host.connection_state", reflect.TypeOf(HostConnectionState(HostConnectionState_CONNECTED)))
	fieldNameMap["connection_state"] = "ConnectionState"
	fields["power_state"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.host.power_state", reflect.TypeOf(HostPowerState(HostPowerState_POWERED_ON))))
	fieldNameMap["power_state"] = "PowerState"
	var validators = []bindings.Validator{}
	uv1 := bindings.NewUnionValidator("connection_state",
		map[string][]bindings.FieldData{
			"CONNECTED": []bindings.FieldData{
				bindings.NewFieldData("power_state", true),
			},
			"DISCONNECTED": []bindings.FieldData{},
			"NOT_RESPONDING": []bindings.FieldData{},
		},
	)
	validators = append(validators, uv1)
	return bindings.NewStructType("com.vmware.vcenter.host.summary", fields, reflect.TypeOf(HostSummary{}), fieldNameMap, validators)
}

