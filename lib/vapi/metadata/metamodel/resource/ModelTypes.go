// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: Model.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package resource

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/protocol"
	"reflect"
)

func modelListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["resource_id"] = bindings.NewIdType([]string{"com.vmware.vapi.resource"}, "")
	fieldNameMap["resource_id"] = "ResourceId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func modelListOutputType() bindings.BindingType {
	return bindings.NewSetType(bindings.NewIdType([]string{"com.vmware.vapi.structure"}, ""), reflect.TypeOf(map[string]bool{}))
}

func modelListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["resource_id"] = bindings.NewIdType([]string{"com.vmware.vapi.resource"}, "")
	fieldNameMap["resource_id"] = "ResourceId"
	paramsTypeMap["resource_id"] = bindings.NewIdType([]string{"com.vmware.vapi.resource"}, "")
	paramsTypeMap["resourceId"] = bindings.NewIdType([]string{"com.vmware.vapi.resource"}, "")
	pathParams["resource_id"] = "resourceId"
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
		"/vapi/metadata/metamodel/resource/{resourceId}/model",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.not_found": 404})
}
