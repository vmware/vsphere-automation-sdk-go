/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: Ethernet.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package hardware

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)

// Resource type for the virtual Ethernet adapter.
const Ethernet_RESOURCE_TYPE = "com.vmware.vcenter.vm.hardware.Ethernet"


// The ``EmulationType`` enumeration class defines the valid emulation types for a virtual Ethernet adapter.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type EthernetEmulationType string

const (
    // E1000 ethernet adapter.
	EthernetEmulationType_E1000 EthernetEmulationType = "E1000"
    // E1000e ethernet adapter.
	EthernetEmulationType_E1000E EthernetEmulationType = "E1000E"
    // AMD Lance PCNet32 Ethernet adapter.
	EthernetEmulationType_PCNET32 EthernetEmulationType = "PCNET32"
    // VMware Vmxnet virtual Ethernet adapter.
	EthernetEmulationType_VMXNET EthernetEmulationType = "VMXNET"
    // VMware Vmxnet2 virtual Ethernet adapter.
	EthernetEmulationType_VMXNET2 EthernetEmulationType = "VMXNET2"
    // VMware Vmxnet3 virtual Ethernet adapter.
	EthernetEmulationType_VMXNET3 EthernetEmulationType = "VMXNET3"
)

func (e EthernetEmulationType) EthernetEmulationType() bool {
	switch e {
	case EthernetEmulationType_E1000:
		return true
	case EthernetEmulationType_E1000E:
		return true
	case EthernetEmulationType_PCNET32:
		return true
	case EthernetEmulationType_VMXNET:
		return true
	case EthernetEmulationType_VMXNET2:
		return true
	case EthernetEmulationType_VMXNET3:
		return true
	default:
		return false
	}
}


// The ``MacAddressType`` enumeration class defines the valid MAC address origins for a virtual Ethernet adapter.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type EthernetMacAddressType string

const (
    // MAC address is assigned statically.
	EthernetMacAddressType_MANUAL EthernetMacAddressType = "MANUAL"
    // MAC address is generated automatically.
	EthernetMacAddressType_GENERATED EthernetMacAddressType = "GENERATED"
    // MAC address is assigned by vCenter Server.
	EthernetMacAddressType_ASSIGNED EthernetMacAddressType = "ASSIGNED"
)

func (m EthernetMacAddressType) EthernetMacAddressType() bool {
	switch m {
	case EthernetMacAddressType_MANUAL:
		return true
	case EthernetMacAddressType_GENERATED:
		return true
	case EthernetMacAddressType_ASSIGNED:
		return true
	default:
		return false
	}
}


// The ``BackingType`` enumeration class defines the valid backing types for a virtual Ethernet adapter.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type EthernetBackingType string

const (
    // vSphere standard portgroup network backing.
	EthernetBackingType_STANDARD_PORTGROUP EthernetBackingType = "STANDARD_PORTGROUP"
    // Legacy host device network backing. Imported VMs may have virtual Ethernet adapters with this type of backing, but this type of backing cannot be used to create or to update a virtual Ethernet adapter.
	EthernetBackingType_HOST_DEVICE EthernetBackingType = "HOST_DEVICE"
    // Distributed virtual switch backing.
	EthernetBackingType_DISTRIBUTED_PORTGROUP EthernetBackingType = "DISTRIBUTED_PORTGROUP"
    // Opaque network backing.
	EthernetBackingType_OPAQUE_NETWORK EthernetBackingType = "OPAQUE_NETWORK"
)

func (b EthernetBackingType) EthernetBackingType() bool {
	switch b {
	case EthernetBackingType_STANDARD_PORTGROUP:
		return true
	case EthernetBackingType_HOST_DEVICE:
		return true
	case EthernetBackingType_DISTRIBUTED_PORTGROUP:
		return true
	case EthernetBackingType_OPAQUE_NETWORK:
		return true
	default:
		return false
	}
}


// The ``BackingInfo`` class contains information about the physical resource backing a virtual Ethernet adapter.
type EthernetBackingInfo struct {
    // Backing type for the virtual Ethernet adapter.
	Type_ EthernetBackingType
    // Identifier of the network backing the virtual Ethernet adapter.
	Network *string
    // Name of the standard portgroup backing the virtual Ethernet adapter.
	NetworkName *string
    // Name of the device backing the virtual Ethernet adapter.
	HostDevice *string
    // UUID of the distributed virtual switch that backs the virtual Ethernet adapter.
	DistributedSwitchUuid *string
    // Key of the distributed virtual port that backs the virtual Ethernet adapter.
	DistributedPort *string
    // Server-generated cookie that identifies the connection to the port. This ookie may be used to verify that the virtual machine is the rightful owner of the port.
	ConnectionCookie *int64
    // Type of the opaque network that backs the virtual Ethernet adapter.
	OpaqueNetworkType *string
    // Identifier of the opaque network that backs the virtual Ethernet adapter.
	OpaqueNetworkId *string
}

// The ``BackingSpec`` class provides a specification of the physical resource that backs a virtual Ethernet adapter.
type EthernetBackingSpec struct {
    // Backing type for the virtual Ethernet adapter.
	Type_ EthernetBackingType
    // Identifier of the network that backs the virtual Ethernet adapter.
	Network *string
    // Key of the distributed virtual port that backs the virtual Ethernet adapter. Depending on the type of the Portgroup, the port may be specified using this field. If the portgroup type is early-binding (also known as static), a port is assigned when the Ethernet adapter is configured to use the port. The port may be either automatically or specifically assigned based on the value of this property. If the portgroup type is ephemeral, the port is created and assigned to a virtual machine when it is powered on and the Ethernet adapter is connected. This property cannot be specified as no free ports exist before use.
	DistributedPort *string
}

// The ``Info`` class contains information about a virtual Ethernet adapter.
type EthernetInfo struct {
    // Device label.
	Label string
    // Ethernet adapter emulation type.
	Type_ EthernetEmulationType
    // Flag indicating whether Universal Pass-Through (UPT) compatibility is enabled on this virtual Ethernet adapter.
	UptCompatibilityEnabled *bool
    // MAC address type.
	MacType EthernetMacAddressType
    // MAC address.
	MacAddress *string
    // Address of the virtual Ethernet adapter on the PCI bus. If the PCI address is invalid, the server will change it when the VM is started or as the device is hot added.
	PciSlotNumber *int64
    // Flag indicating whether wake-on-LAN is enabled on this virtual Ethernet adapter.
	WakeOnLanEnabled bool
    // Physical resource backing for the virtual Ethernet adapter.
	Backing EthernetBackingInfo
    // Connection status of the virtual device.
	State ConnectionState
    // Flag indicating whether the virtual device should be connected whenever the virtual machine is powered on.
	StartConnected bool
    // Flag indicating whether the guest can connect and disconnect the device.
	AllowGuestControl bool
}

// The ``CreateSpec`` class provides a specification for the configuration of a newly-created virtual Ethernet adapter.
type EthernetCreateSpec struct {
    // Ethernet adapter emulation type.
	Type_ *EthernetEmulationType
    // Flag indicating whether Universal Pass-Through (UPT) compatibility is enabled on this virtual Ethernet adapter.
	UptCompatibilityEnabled *bool
    // MAC address type.
	MacType *EthernetMacAddressType
    // MAC address.
	MacAddress *string
    // Address of the virtual Ethernet adapter on the PCI bus. If the PCI address is invalid, the server will change when it the VM is started or as the device is hot added.
	PciSlotNumber *int64
    // Flag indicating whether wake-on-LAN is enabled on this virtual Ethernet adapter.
	WakeOnLanEnabled *bool
    // Physical resource backing for the virtual Ethernet adapter.
	Backing *EthernetBackingSpec
    // Flag indicating whether the virtual device should be connected whenever the virtual machine is powered on.
	StartConnected *bool
    // Flag indicating whether the guest can connect and disconnect the device.
	AllowGuestControl *bool
}

// The ``UpdateSpec`` class describes the updates to be made to the configuration of a virtual Ethernet adapter.
type EthernetUpdateSpec struct {
    // Flag indicating whether Universal Pass-Through (UPT) compatibility should be enabled on this virtual Ethernet adapter. 
    //
    //  This property may be modified at any time, and changes will be applied the next time the virtual machine is powered on.
	UptCompatibilityEnabled *bool
    // MAC address type. 
    //
    //  This property may be modified at any time, and changes will be applied the next time the virtual machine is powered on.
	MacType *EthernetMacAddressType
    // MAC address. 
    //
    //  This property may be modified at any time, and changes will be applied the next time the virtual machine is powered on.
	MacAddress *string
    // Flag indicating whether wake-on-LAN shoud be enabled on this virtual Ethernet adapter. 
    //
    //  This property may be modified at any time, and changes will be applied the next time the virtual machine is powered on.
	WakeOnLanEnabled *bool
    // Physical resource backing for the virtual Ethernet adapter. 
    //
    //  This property may be modified at any time, and changes will be applied the next time the virtual machine is powered on.
	Backing *EthernetBackingSpec
    // Flag indicating whether the virtual device should be connected whenever the virtual machine is powered on.
	StartConnected *bool
    // Flag indicating whether the guest can connect and disconnect the device.
	AllowGuestControl *bool
}

// The ``Summary`` class contains commonly used information about a virtual Ethernet adapter.
type EthernetSummary struct {
    // Identifier of the virtual Ethernet adapter.
	Nic string
}



func ethernetListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fieldNameMap["vm"] = "Vm"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func ethernetListOutputType() bindings.BindingType {
	return bindings.NewListType(bindings.NewReferenceType(EthernetSummaryBindingType), reflect.TypeOf([]EthernetSummary{}))
}

func ethernetListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fieldNameMap["vm"] = "Vm"
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
		map[string]int{"Error": 500,"NotFound": 404,"ResourceInaccessible": 500,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}

func ethernetGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fields["nic"] = bindings.NewIdType([]string{"com.vmware.vcenter.vm.hardware.Ethernet"}, "")
	fieldNameMap["vm"] = "Vm"
	fieldNameMap["nic"] = "Nic"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func ethernetGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(EthernetInfoBindingType)
}

func ethernetGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fields["nic"] = bindings.NewIdType([]string{"com.vmware.vcenter.vm.hardware.Ethernet"}, "")
	fieldNameMap["vm"] = "Vm"
	fieldNameMap["nic"] = "Nic"
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
		map[string]int{"Error": 500,"NotFound": 404,"ResourceInaccessible": 500,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}

func ethernetCreateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fields["spec"] = bindings.NewReferenceType(EthernetCreateSpecBindingType)
	fieldNameMap["vm"] = "Vm"
	fieldNameMap["spec"] = "Spec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func ethernetCreateOutputType() bindings.BindingType {
	return bindings.NewIdType([]string{"com.vmware.vcenter.vm.hardware.Ethernet"}, "")
}

func ethernetCreateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fields["spec"] = bindings.NewReferenceType(EthernetCreateSpecBindingType)
	fieldNameMap["vm"] = "Vm"
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
		map[string]int{"Error": 500,"NotFound": 404,"UnableToAllocateResource": 500,"InvalidArgument": 400,"ResourceBusy": 500,"ResourceInaccessible": 500,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403,"Unsupported": 400})
}

func ethernetUpdateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fields["nic"] = bindings.NewIdType([]string{"com.vmware.vcenter.vm.hardware.Ethernet"}, "")
	fields["spec"] = bindings.NewReferenceType(EthernetUpdateSpecBindingType)
	fieldNameMap["vm"] = "Vm"
	fieldNameMap["nic"] = "Nic"
	fieldNameMap["spec"] = "Spec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func ethernetUpdateOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func ethernetUpdateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fields["nic"] = bindings.NewIdType([]string{"com.vmware.vcenter.vm.hardware.Ethernet"}, "")
	fields["spec"] = bindings.NewReferenceType(EthernetUpdateSpecBindingType)
	fieldNameMap["vm"] = "Vm"
	fieldNameMap["nic"] = "Nic"
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
		map[string]int{"Error": 500,"InvalidArgument": 400,"NotFound": 404,"ResourceBusy": 500,"ResourceInaccessible": 500,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}

func ethernetDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fields["nic"] = bindings.NewIdType([]string{"com.vmware.vcenter.vm.hardware.Ethernet"}, "")
	fieldNameMap["vm"] = "Vm"
	fieldNameMap["nic"] = "Nic"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func ethernetDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func ethernetDeleteRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fields["nic"] = bindings.NewIdType([]string{"com.vmware.vcenter.vm.hardware.Ethernet"}, "")
	fieldNameMap["vm"] = "Vm"
	fieldNameMap["nic"] = "Nic"
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
		map[string]int{"Error": 500,"NotFound": 404,"ResourceBusy": 500,"ResourceInaccessible": 500,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}

func ethernetConnectInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fields["nic"] = bindings.NewIdType([]string{"com.vmware.vcenter.vm.hardware.Ethernet"}, "")
	fieldNameMap["vm"] = "Vm"
	fieldNameMap["nic"] = "Nic"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func ethernetConnectOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func ethernetConnectRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fields["nic"] = bindings.NewIdType([]string{"com.vmware.vcenter.vm.hardware.Ethernet"}, "")
	fieldNameMap["vm"] = "Vm"
	fieldNameMap["nic"] = "Nic"
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
		map[string]int{"Error": 500,"NotFound": 404,"AlreadyInDesiredState": 400,"NotAllowedInCurrentState": 400,"ResourceBusy": 500,"ResourceInaccessible": 500,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}

func ethernetDisconnectInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fields["nic"] = bindings.NewIdType([]string{"com.vmware.vcenter.vm.hardware.Ethernet"}, "")
	fieldNameMap["vm"] = "Vm"
	fieldNameMap["nic"] = "Nic"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func ethernetDisconnectOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func ethernetDisconnectRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fields["nic"] = bindings.NewIdType([]string{"com.vmware.vcenter.vm.hardware.Ethernet"}, "")
	fieldNameMap["vm"] = "Vm"
	fieldNameMap["nic"] = "Nic"
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
		map[string]int{"Error": 500,"NotFound": 404,"AlreadyInDesiredState": 400,"NotAllowedInCurrentState": 400,"ResourceBusy": 500,"ResourceInaccessible": 500,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}


func EthernetBackingInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["type"] = bindings.NewEnumType("com.vmware.vcenter.vm.hardware.ethernet.backing_type", reflect.TypeOf(EthernetBackingType(EthernetBackingType_STANDARD_PORTGROUP)))
	fieldNameMap["type"] = "Type_"
	fields["network"] = bindings.NewOptionalType(bindings.NewIdType([]string{"Network"}, ""))
	fieldNameMap["network"] = "Network"
	fields["network_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["network_name"] = "NetworkName"
	fields["host_device"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["host_device"] = "HostDevice"
	fields["distributed_switch_uuid"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["distributed_switch_uuid"] = "DistributedSwitchUuid"
	fields["distributed_port"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["distributed_port"] = "DistributedPort"
	fields["connection_cookie"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["connection_cookie"] = "ConnectionCookie"
	fields["opaque_network_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["opaque_network_type"] = "OpaqueNetworkType"
	fields["opaque_network_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["opaque_network_id"] = "OpaqueNetworkId"
	var validators = []bindings.Validator{}
	uv1 := bindings.NewUnionValidator("type",
		map[string][]bindings.FieldData{
			"STANDARD_PORTGROUP": []bindings.FieldData{
				bindings.NewFieldData("network", false),
				bindings.NewFieldData("network_name", true),
			},
			"DISTRIBUTED_PORTGROUP": []bindings.FieldData{
				bindings.NewFieldData("network", false),
				bindings.NewFieldData("distributed_switch_uuid", true),
				bindings.NewFieldData("distributed_port", false),
				bindings.NewFieldData("connection_cookie", false),
			},
			"OPAQUE_NETWORK": []bindings.FieldData{
				bindings.NewFieldData("network", false),
				bindings.NewFieldData("opaque_network_type", true),
				bindings.NewFieldData("opaque_network_id", true),
			},
			"HOST_DEVICE": []bindings.FieldData{
				bindings.NewFieldData("host_device", true),
			},
		},
	)
	validators = append(validators, uv1)
	return bindings.NewStructType("com.vmware.vcenter.vm.hardware.ethernet.backing_info", fields, reflect.TypeOf(EthernetBackingInfo{}), fieldNameMap, validators)
}

func EthernetBackingSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["type"] = bindings.NewEnumType("com.vmware.vcenter.vm.hardware.ethernet.backing_type", reflect.TypeOf(EthernetBackingType(EthernetBackingType_STANDARD_PORTGROUP)))
	fieldNameMap["type"] = "Type_"
	fields["network"] = bindings.NewOptionalType(bindings.NewIdType([]string{"Network"}, ""))
	fieldNameMap["network"] = "Network"
	fields["distributed_port"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["distributed_port"] = "DistributedPort"
	var validators = []bindings.Validator{}
	uv1 := bindings.NewUnionValidator("type",
		map[string][]bindings.FieldData{
			"STANDARD_PORTGROUP": []bindings.FieldData{
				bindings.NewFieldData("network", true),
			},
			"DISTRIBUTED_PORTGROUP": []bindings.FieldData{
				bindings.NewFieldData("network", true),
				bindings.NewFieldData("distributed_port", false),
			},
			"OPAQUE_NETWORK": []bindings.FieldData{
				bindings.NewFieldData("network", true),
			},
			"HOST_DEVICE": []bindings.FieldData{},
		},
	)
	validators = append(validators, uv1)
	return bindings.NewStructType("com.vmware.vcenter.vm.hardware.ethernet.backing_spec", fields, reflect.TypeOf(EthernetBackingSpec{}), fieldNameMap, validators)
}

func EthernetInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["label"] = bindings.NewStringType()
	fieldNameMap["label"] = "Label"
	fields["type"] = bindings.NewEnumType("com.vmware.vcenter.vm.hardware.ethernet.emulation_type", reflect.TypeOf(EthernetEmulationType(EthernetEmulationType_E1000)))
	fieldNameMap["type"] = "Type_"
	fields["upt_compatibility_enabled"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["upt_compatibility_enabled"] = "UptCompatibilityEnabled"
	fields["mac_type"] = bindings.NewEnumType("com.vmware.vcenter.vm.hardware.ethernet.mac_address_type", reflect.TypeOf(EthernetMacAddressType(EthernetMacAddressType_MANUAL)))
	fieldNameMap["mac_type"] = "MacType"
	fields["mac_address"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["mac_address"] = "MacAddress"
	fields["pci_slot_number"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["pci_slot_number"] = "PciSlotNumber"
	fields["wake_on_lan_enabled"] = bindings.NewBooleanType()
	fieldNameMap["wake_on_lan_enabled"] = "WakeOnLanEnabled"
	fields["backing"] = bindings.NewReferenceType(EthernetBackingInfoBindingType)
	fieldNameMap["backing"] = "Backing"
	fields["state"] = bindings.NewEnumType("com.vmware.vcenter.vm.hardware.connection_state", reflect.TypeOf(ConnectionState(ConnectionState_CONNECTED)))
	fieldNameMap["state"] = "State"
	fields["start_connected"] = bindings.NewBooleanType()
	fieldNameMap["start_connected"] = "StartConnected"
	fields["allow_guest_control"] = bindings.NewBooleanType()
	fieldNameMap["allow_guest_control"] = "AllowGuestControl"
	var validators = []bindings.Validator{}
	uv1 := bindings.NewUnionValidator("type",
		map[string][]bindings.FieldData{
			"VMXNET3": []bindings.FieldData{
				bindings.NewFieldData("upt_compatibility_enabled", true),
			},
			"E1000": []bindings.FieldData{},
			"E1000E": []bindings.FieldData{},
			"PCNET32": []bindings.FieldData{},
			"VMXNET": []bindings.FieldData{},
			"VMXNET2": []bindings.FieldData{},
		},
	)
	validators = append(validators, uv1)
	return bindings.NewStructType("com.vmware.vcenter.vm.hardware.ethernet.info", fields, reflect.TypeOf(EthernetInfo{}), fieldNameMap, validators)
}

func EthernetCreateSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["type"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.vm.hardware.ethernet.emulation_type", reflect.TypeOf(EthernetEmulationType(EthernetEmulationType_E1000))))
	fieldNameMap["type"] = "Type_"
	fields["upt_compatibility_enabled"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["upt_compatibility_enabled"] = "UptCompatibilityEnabled"
	fields["mac_type"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.vm.hardware.ethernet.mac_address_type", reflect.TypeOf(EthernetMacAddressType(EthernetMacAddressType_MANUAL))))
	fieldNameMap["mac_type"] = "MacType"
	fields["mac_address"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["mac_address"] = "MacAddress"
	fields["pci_slot_number"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["pci_slot_number"] = "PciSlotNumber"
	fields["wake_on_lan_enabled"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["wake_on_lan_enabled"] = "WakeOnLanEnabled"
	fields["backing"] = bindings.NewOptionalType(bindings.NewReferenceType(EthernetBackingSpecBindingType))
	fieldNameMap["backing"] = "Backing"
	fields["start_connected"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["start_connected"] = "StartConnected"
	fields["allow_guest_control"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["allow_guest_control"] = "AllowGuestControl"
	var validators = []bindings.Validator{}
	uv1 := bindings.NewUnionValidator("type",
		map[string][]bindings.FieldData{
			"VMXNET3": []bindings.FieldData{
				bindings.NewFieldData("upt_compatibility_enabled", false),
			},
			"E1000": []bindings.FieldData{},
			"E1000E": []bindings.FieldData{},
			"PCNET32": []bindings.FieldData{},
			"VMXNET": []bindings.FieldData{},
			"VMXNET2": []bindings.FieldData{},
		},
	)
	validators = append(validators, uv1)
	uv2 := bindings.NewUnionValidator("mac_type",
		map[string][]bindings.FieldData{
			"MANUAL": []bindings.FieldData{
				bindings.NewFieldData("mac_address", true),
			},
			"GENERATED": []bindings.FieldData{},
			"ASSIGNED": []bindings.FieldData{},
		},
	)
	validators = append(validators, uv2)
	return bindings.NewStructType("com.vmware.vcenter.vm.hardware.ethernet.create_spec", fields, reflect.TypeOf(EthernetCreateSpec{}), fieldNameMap, validators)
}

func EthernetUpdateSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["upt_compatibility_enabled"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["upt_compatibility_enabled"] = "UptCompatibilityEnabled"
	fields["mac_type"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.vm.hardware.ethernet.mac_address_type", reflect.TypeOf(EthernetMacAddressType(EthernetMacAddressType_MANUAL))))
	fieldNameMap["mac_type"] = "MacType"
	fields["mac_address"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["mac_address"] = "MacAddress"
	fields["wake_on_lan_enabled"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["wake_on_lan_enabled"] = "WakeOnLanEnabled"
	fields["backing"] = bindings.NewOptionalType(bindings.NewReferenceType(EthernetBackingSpecBindingType))
	fieldNameMap["backing"] = "Backing"
	fields["start_connected"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["start_connected"] = "StartConnected"
	fields["allow_guest_control"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["allow_guest_control"] = "AllowGuestControl"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.vm.hardware.ethernet.update_spec", fields, reflect.TypeOf(EthernetUpdateSpec{}), fieldNameMap, validators)
}

func EthernetSummaryBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["nic"] = bindings.NewIdType([]string{"com.vmware.vcenter.vm.hardware.Ethernet"}, "")
	fieldNameMap["nic"] = "Nic"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.vm.hardware.ethernet.summary", fields, reflect.TypeOf(EthernetSummary{}), fieldNameMap, validators)
}

