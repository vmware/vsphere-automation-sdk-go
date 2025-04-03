// Copyright (c) 2019-2025 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: FirewallIdentityStores.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package global_infra

import (
	"reflect"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
)

// Possible value for ``action`` of method FirewallIdentityStores#create.
const FirewallIdentityStores_CREATE_ACTION_FULL_SYNC = "FULL_SYNC"
// Possible value for ``action`` of method FirewallIdentityStores#create.
const FirewallIdentityStores_CREATE_ACTION_DELTA_SYNC = "DELTA_SYNC"
// Possible value for ``action`` of method FirewallIdentityStores#create.
const FirewallIdentityStores_CREATE_ACTION_STOP_SYNC = "STOP_SYNC"




func firewallIdentityStoresCreateInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["firewall_identity_store_id"] = vapiBindings_.NewStringType()
	fields["action"] = vapiBindings_.NewStringType()
	fields["delay"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fields["enforcement_point_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["firewall_identity_store_id"] = "FirewallIdentityStoreId"
	fieldNameMap["action"] = "Action"
	fieldNameMap["delay"] = "Delay"
	fieldNameMap["enforcement_point_path"] = "EnforcementPointPath"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func FirewallIdentityStoresCreateOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func firewallIdentityStoresCreateRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["firewall_identity_store_id"] = vapiBindings_.NewStringType()
	fields["action"] = vapiBindings_.NewStringType()
	fields["delay"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fields["enforcement_point_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["firewall_identity_store_id"] = "FirewallIdentityStoreId"
	fieldNameMap["action"] = "Action"
	fieldNameMap["delay"] = "Delay"
	fieldNameMap["enforcement_point_path"] = "EnforcementPointPath"
	paramsTypeMap["firewall_identity_store_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["delay"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	paramsTypeMap["action"] = vapiBindings_.NewStringType()
	paramsTypeMap["enforcement_point_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["firewallIdentityStoreId"] = vapiBindings_.NewStringType()
	pathParams["firewall_identity_store_id"] = "firewallIdentityStoreId"
	queryParams["delay"] = "delay"
	queryParams["action"] = "action"
	queryParams["enforcement_point_path"] = "enforcement_point_path"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return vapiProtocol_.NewOperationRestMetadata(
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
		"POST",
		"/policy/api/v1/global-infra/firewall-identity-stores/{firewallIdentityStoreId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400,"com.vmware.vapi.std.errors.unauthorized": 403,"com.vmware.vapi.std.errors.service_unavailable": 503,"com.vmware.vapi.std.errors.internal_server_error": 500,"com.vmware.vapi.std.errors.not_found": 404})
}


