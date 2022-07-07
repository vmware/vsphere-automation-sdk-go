// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: AuditLogs.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package administration

import (
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
	"reflect"
)

func auditLogsCreateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["audit_log_request"] = bindings.NewReferenceType(model.AuditLogRequestBindingType)
	fields["cursor"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["audit_log_request"] = "AuditLogRequest"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["fields"] = "Fields"
	fieldNameMap["page_size"] = "PageSize"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func auditLogsCreateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.AuditLogListResultBindingType)
}

func auditLogsCreateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["audit_log_request"] = bindings.NewReferenceType(model.AuditLogRequestBindingType)
	fields["cursor"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["audit_log_request"] = "AuditLogRequest"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["fields"] = "Fields"
	fieldNameMap["page_size"] = "PageSize"
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["audit_log_request"] = bindings.NewReferenceType(model.AuditLogRequestBindingType)
	queryParams["cursor"] = "cursor"
	queryParams["fields"] = "fields"
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
		"audit_log_request",
		"POST",
		"/api/v1/administration/audit-logs",
		"",
		resultHeaders,
		201,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.concurrent_change": 409, "com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
