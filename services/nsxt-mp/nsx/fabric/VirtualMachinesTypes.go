// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: VirtualMachines.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package fabric

import (
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
	"reflect"
)

func virtualMachinesAddtagsInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["virtual_machine_tag_update"] = bindings.NewReferenceType(model.VirtualMachineTagUpdateBindingType)
	fieldNameMap["virtual_machine_tag_update"] = "VirtualMachineTagUpdate"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func virtualMachinesAddtagsOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func virtualMachinesAddtagsRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["virtual_machine_tag_update"] = bindings.NewReferenceType(model.VirtualMachineTagUpdateBindingType)
	fieldNameMap["virtual_machine_tag_update"] = "VirtualMachineTagUpdate"
	paramsTypeMap["virtual_machine_tag_update"] = bindings.NewReferenceType(model.VirtualMachineTagUpdateBindingType)
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		dispatchHeaderParams,
		bodyFieldsMap,
		"action=add_tags",
		"virtual_machine_tag_update",
		"POST",
		"/api/v1/fabric/virtual-machines",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func virtualMachinesListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["display_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["exclude_vm_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["external_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["host_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["display_name"] = "DisplayName"
	fieldNameMap["exclude_vm_type"] = "ExcludeVmType"
	fieldNameMap["external_id"] = "ExternalId"
	fieldNameMap["host_id"] = "HostId"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func virtualMachinesListOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.VirtualMachineListResultBindingType)
}

func virtualMachinesListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["display_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["exclude_vm_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["external_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["host_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["display_name"] = "DisplayName"
	fieldNameMap["exclude_vm_type"] = "ExcludeVmType"
	fieldNameMap["external_id"] = "ExternalId"
	fieldNameMap["host_id"] = "HostId"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	paramsTypeMap["external_id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["exclude_vm_type"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["display_name"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["host_id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	queryParams["cursor"] = "cursor"
	queryParams["sort_ascending"] = "sort_ascending"
	queryParams["included_fields"] = "included_fields"
	queryParams["exclude_vm_type"] = "exclude_vm_type"
	queryParams["external_id"] = "external_id"
	queryParams["sort_by"] = "sort_by"
	queryParams["display_name"] = "display_name"
	queryParams["host_id"] = "host_id"
	queryParams["page_size"] = "page_size"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		dispatchHeaderParams,
		bodyFieldsMap,
		"",
		"",
		"GET",
		"/api/v1/fabric/virtual-machines",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func virtualMachinesRemovetagsInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["virtual_machine_tag_update"] = bindings.NewReferenceType(model.VirtualMachineTagUpdateBindingType)
	fieldNameMap["virtual_machine_tag_update"] = "VirtualMachineTagUpdate"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func virtualMachinesRemovetagsOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func virtualMachinesRemovetagsRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["virtual_machine_tag_update"] = bindings.NewReferenceType(model.VirtualMachineTagUpdateBindingType)
	fieldNameMap["virtual_machine_tag_update"] = "VirtualMachineTagUpdate"
	paramsTypeMap["virtual_machine_tag_update"] = bindings.NewReferenceType(model.VirtualMachineTagUpdateBindingType)
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		dispatchHeaderParams,
		bodyFieldsMap,
		"action=remove_tags",
		"virtual_machine_tag_update",
		"POST",
		"/api/v1/fabric/virtual-machines",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func virtualMachinesUpdatetagsInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["virtual_machine_tag_update"] = bindings.NewReferenceType(model.VirtualMachineTagUpdateBindingType)
	fieldNameMap["virtual_machine_tag_update"] = "VirtualMachineTagUpdate"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func virtualMachinesUpdatetagsOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func virtualMachinesUpdatetagsRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["virtual_machine_tag_update"] = bindings.NewReferenceType(model.VirtualMachineTagUpdateBindingType)
	fieldNameMap["virtual_machine_tag_update"] = "VirtualMachineTagUpdate"
	paramsTypeMap["virtual_machine_tag_update"] = bindings.NewReferenceType(model.VirtualMachineTagUpdateBindingType)
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		dispatchHeaderParams,
		bodyFieldsMap,
		"action=update_tags",
		"virtual_machine_tag_update",
		"POST",
		"/api/v1/fabric/virtual-machines",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
