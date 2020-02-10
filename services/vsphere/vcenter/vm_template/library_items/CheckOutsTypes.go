/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: CheckOuts.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package library_items

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
	"time"
)



// The ``CheckOutSpec`` class defines the information required to check out a library item containing a virtual machine template. This class was added in vSphere API 6.9.1.
type CheckOutsCheckOutSpec struct {
    // Name of the virtual machine to check out of the library item. This property was added in vSphere API 6.9.1.
	Name *string
    // Information used to place the checked out virtual machine. This property was added in vSphere API 6.9.1.
	Placement *CheckOutsPlacementSpec
    // Specifies whether the virtual machine should be powered on after check out. This property was added in vSphere API 6.9.1.
	PoweredOn *bool
}

// The ``PlacementSpec`` class contains information used to place a checked out virtual machine onto resources within the vCenter inventory. The specified compute resource should have access to the storage associated with the checked out virtual machine. This class was added in vSphere API 6.9.1.
type CheckOutsPlacementSpec struct {
    // Virtual machine folder into which the virtual machine should be placed. This property was added in vSphere API 6.9.1.
	Folder *string
    // Resource pool into which the virtual machine should be placed. This property was added in vSphere API 6.9.1.
	ResourcePool *string
    // Host onto which the virtual machine should be placed. If ``host`` and ``resourcePool`` are both specified, ``resourcePool`` must belong to ``host``. If ``host`` and ``cluster`` are both specified, ``host`` must be a member of ``cluster``. This property was added in vSphere API 6.9.1.
	Host *string
    // Cluster onto which the virtual machine should be placed. If ``cluster`` and ``resourcePool`` are both specified, ``resourcePool`` must belong to ``cluster``. If ``cluster`` and ``host`` are both specified, ``host`` must be a member of ``cluster``. This property was added in vSphere API 6.9.1.
	Cluster *string
}

// The ``CheckInSpec`` class defines the information required to check in a virtual machine into a library item. This class was added in vSphere API 6.9.1.
type CheckOutsCheckInSpec struct {
    // Message describing the changes made to the virtual machine. This property was added in vSphere API 6.9.1.
	Message string
}

// The ``Summary`` class contains commonly used information about a checked out virtual machine. This class was added in vSphere API 6.9.1.
type CheckOutsSummary struct {
    // Identifier of the checked out virtual machine. This property was added in vSphere API 6.9.1.
	Vm string
}

// The ``Info`` class contains information about a checked out virtual machine. This class was added in vSphere API 6.9.1.
type CheckOutsInfo struct {
    // Date and time when the virtual machine was checked out. This property was added in vSphere API 6.9.1.
	Time time.Time
    // Name of the user who checked out the virtual machine. This property was added in vSphere API 6.9.1.
	User string
}



func checkOutsCheckOutInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["template_library_item"] = bindings.NewIdType([]string{"com.vmware.content.library.Item"}, "")
	fields["spec"] = bindings.NewOptionalType(bindings.NewReferenceType(CheckOutsCheckOutSpecBindingType))
	fieldNameMap["template_library_item"] = "TemplateLibraryItem"
	fieldNameMap["spec"] = "Spec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func checkOutsCheckOutOutputType() bindings.BindingType {
	return bindings.NewIdType([]string{"VirtualMachine"}, "")
}

func checkOutsCheckOutRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["template_library_item"] = bindings.NewIdType([]string{"com.vmware.content.library.Item"}, "")
	fields["spec"] = bindings.NewOptionalType(bindings.NewReferenceType(CheckOutsCheckOutSpecBindingType))
	fieldNameMap["template_library_item"] = "TemplateLibraryItem"
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
		map[string]int{"AlreadyExists": 400,"InvalidArgument": 400,"NotFound": 404,"ResourceInaccessible": 500,"UnableToAllocateResource": 500,"Unauthenticated": 401,"Unauthorized": 403,"Error": 500})
}

func checkOutsCheckInInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["template_library_item"] = bindings.NewIdType([]string{"com.vmware.content.library.Item"}, "")
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fields["spec"] = bindings.NewOptionalType(bindings.NewReferenceType(CheckOutsCheckInSpecBindingType))
	fieldNameMap["template_library_item"] = "TemplateLibraryItem"
	fieldNameMap["vm"] = "Vm"
	fieldNameMap["spec"] = "Spec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func checkOutsCheckInOutputType() bindings.BindingType {
	return bindings.NewIdType([]string{"com.vmware.content.library.item.Version"}, "")
}

func checkOutsCheckInRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["template_library_item"] = bindings.NewIdType([]string{"com.vmware.content.library.Item"}, "")
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fields["spec"] = bindings.NewOptionalType(bindings.NewReferenceType(CheckOutsCheckInSpecBindingType))
	fieldNameMap["template_library_item"] = "TemplateLibraryItem"
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
		map[string]int{"InvalidArgument": 400,"NotAllowedInCurrentState": 400,"NotFound": 404,"ResourceInaccessible": 500,"Unauthenticated": 401,"Unauthorized": 403,"Error": 500})
}

func checkOutsListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["template_library_item"] = bindings.NewIdType([]string{"com.vmware.content.library.Item"}, "")
	fieldNameMap["template_library_item"] = "TemplateLibraryItem"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func checkOutsListOutputType() bindings.BindingType {
	return bindings.NewListType(bindings.NewReferenceType(CheckOutsSummaryBindingType), reflect.TypeOf([]CheckOutsSummary{}))
}

func checkOutsListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["template_library_item"] = bindings.NewIdType([]string{"com.vmware.content.library.Item"}, "")
	fieldNameMap["template_library_item"] = "TemplateLibraryItem"
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
		map[string]int{"NotFound": 404,"InvalidArgument": 400,"Unauthenticated": 401,"Unauthorized": 403,"Error": 500})
}

func checkOutsGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["template_library_item"] = bindings.NewIdType([]string{"com.vmware.content.library.Item"}, "")
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fieldNameMap["template_library_item"] = "TemplateLibraryItem"
	fieldNameMap["vm"] = "Vm"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func checkOutsGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(CheckOutsInfoBindingType)
}

func checkOutsGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["template_library_item"] = bindings.NewIdType([]string{"com.vmware.content.library.Item"}, "")
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fieldNameMap["template_library_item"] = "TemplateLibraryItem"
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
		map[string]int{"NotFound": 404,"InvalidArgument": 400,"Unauthenticated": 401,"Unauthorized": 403,"Error": 500})
}

func checkOutsDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["template_library_item"] = bindings.NewIdType([]string{"com.vmware.content.library.Item"}, "")
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fieldNameMap["template_library_item"] = "TemplateLibraryItem"
	fieldNameMap["vm"] = "Vm"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func checkOutsDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func checkOutsDeleteRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["template_library_item"] = bindings.NewIdType([]string{"com.vmware.content.library.Item"}, "")
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fieldNameMap["template_library_item"] = "TemplateLibraryItem"
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
		map[string]int{"NotFound": 404,"InvalidArgument": 400,"NotAllowedInCurrentState": 400,"ResourceBusy": 500,"ResourceInaccessible": 500,"Unauthenticated": 401,"Unauthorized": 403,"Error": 500})
}


func CheckOutsCheckOutSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["name"] = "Name"
	fields["placement"] = bindings.NewOptionalType(bindings.NewReferenceType(CheckOutsPlacementSpecBindingType))
	fieldNameMap["placement"] = "Placement"
	fields["powered_on"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["powered_on"] = "PoweredOn"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.vm_template.library_items.check_outs.check_out_spec", fields, reflect.TypeOf(CheckOutsCheckOutSpec{}), fieldNameMap, validators)
}

func CheckOutsPlacementSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["folder"] = bindings.NewOptionalType(bindings.NewIdType([]string{"Folder"}, ""))
	fieldNameMap["folder"] = "Folder"
	fields["resource_pool"] = bindings.NewOptionalType(bindings.NewIdType([]string{"ResourcePool"}, ""))
	fieldNameMap["resource_pool"] = "ResourcePool"
	fields["host"] = bindings.NewOptionalType(bindings.NewIdType([]string{"HostSystem"}, ""))
	fieldNameMap["host"] = "Host"
	fields["cluster"] = bindings.NewOptionalType(bindings.NewIdType([]string{"ClusterComputeResource"}, ""))
	fieldNameMap["cluster"] = "Cluster"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.vm_template.library_items.check_outs.placement_spec", fields, reflect.TypeOf(CheckOutsPlacementSpec{}), fieldNameMap, validators)
}

func CheckOutsCheckInSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["message"] = bindings.NewStringType()
	fieldNameMap["message"] = "Message"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.vm_template.library_items.check_outs.check_in_spec", fields, reflect.TypeOf(CheckOutsCheckInSpec{}), fieldNameMap, validators)
}

func CheckOutsSummaryBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fieldNameMap["vm"] = "Vm"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.vm_template.library_items.check_outs.summary", fields, reflect.TypeOf(CheckOutsSummary{}), fieldNameMap, validators)
}

func CheckOutsInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["time"] = bindings.NewDateTimeType()
	fieldNameMap["time"] = "Time"
	fields["user"] = bindings.NewStringType()
	fieldNameMap["user"] = "User"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.vm_template.library_items.check_outs.info", fields, reflect.TypeOf(CheckOutsInfo{}), fieldNameMap, validators)
}

