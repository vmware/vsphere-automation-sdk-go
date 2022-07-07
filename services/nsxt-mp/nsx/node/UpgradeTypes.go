// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: Upgrade.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package node

import (
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
	"reflect"
)

func upgradeGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["bundle_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["upgrade_task_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["bundle_name"] = "BundleName"
	fieldNameMap["upgrade_task_id"] = "UpgradeTaskId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func upgradeGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.UpgradeTaskPropertiesBindingType)
}

func upgradeGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["bundle_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["upgrade_task_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["bundle_name"] = "BundleName"
	fieldNameMap["upgrade_task_id"] = "UpgradeTaskId"
	paramsTypeMap["bundle_name"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["upgrade_task_id"] = bindings.NewOptionalType(bindings.NewStringType())
	queryParams["bundle_name"] = "bundle_name"
	queryParams["upgrade_task_id"] = "upgrade_task_id"
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
		"/api/v1/node/upgrade",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
