// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: Reconcile.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package search

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	"reflect"
)

// Possible value for ``action`` of method Reconcile#create.
const Reconcile_CREATE_ACTION_COMPLETE_REINDEXING = "COMPLETE_REINDEXING"

// Possible value for ``action`` of method Reconcile#create.
const Reconcile_CREATE_ACTION_OPENSEARCH_RESTORE = "OPENSEARCH_RESTORE"

// Possible value for ``action`` of method Reconcile#create.
const Reconcile_CREATE_ACTION_PRODUCT_RESTORE = "PRODUCT_RESTORE"

// Possible value for ``action`` of method Reconcile#create.
const Reconcile_CREATE_ACTION_PRODUCT_UPGRADE = "PRODUCT_UPGRADE"

// Possible value for ``action`` of method Reconcile#create.
const Reconcile_CREATE_ACTION_REBALANCE = "REBALANCE"

func reconcileCreateInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["action"] = vapiBindings_.NewStringType()
	fields["override"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["scope"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["action"] = "Action"
	fieldNameMap["override"] = "Override"
	fieldNameMap["scope"] = "Scope"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func ReconcileCreateOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func reconcileCreateRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["action"] = vapiBindings_.NewStringType()
	fields["override"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["scope"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["action"] = "Action"
	fieldNameMap["override"] = "Override"
	fieldNameMap["scope"] = "Scope"
	paramsTypeMap["scope"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["action"] = vapiBindings_.NewStringType()
	paramsTypeMap["override"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	queryParams["scope"] = "scope"
	queryParams["action"] = "action"
	queryParams["override"] = "override"
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
		"/api/v1/search/reconcile",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
