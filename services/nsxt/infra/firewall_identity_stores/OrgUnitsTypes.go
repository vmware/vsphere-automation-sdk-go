// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: OrgUnits.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package firewall_identity_stores

import (
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
	"reflect"
)

func orgUnitsListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["firewall_identity_store_id"] = bindings.NewStringType()
	fields["enforcement_point_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["firewall_identity_store_id"] = "FirewallIdentityStoreId"
	fieldNameMap["enforcement_point_path"] = "EnforcementPointPath"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func orgUnitsListOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.DirectoryOrgUnitListResultsBindingType)
}

func orgUnitsListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["firewall_identity_store_id"] = bindings.NewStringType()
	fields["enforcement_point_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["firewall_identity_store_id"] = "FirewallIdentityStoreId"
	fieldNameMap["enforcement_point_path"] = "EnforcementPointPath"
	paramsTypeMap["firewall_identity_store_id"] = bindings.NewStringType()
	paramsTypeMap["enforcement_point_path"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["firewallIdentityStoreId"] = bindings.NewStringType()
	pathParams["firewall_identity_store_id"] = "firewallIdentityStoreId"
	queryParams["enforcement_point_path"] = "enforcement_point_path"
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
		"/policy/api/v1/infra/firewall-identity-stores/{firewallIdentityStoreId}/org-units",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
