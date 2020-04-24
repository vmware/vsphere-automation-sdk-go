/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: Locale.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package vmc

import (
	"reflect"
	"github.com/vmware/vsphere-automation-sdk-go/services/vmc/model"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
)





func localeSetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vmc_locale"] = bindings.NewReferenceType(model.VmcLocaleBindingType)
	fieldNameMap["vmc_locale"] = "VmcLocale"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func localeSetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.VmcLocaleBindingType)
}

func localeSetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["vmc_locale"] = bindings.NewReferenceType(model.VmcLocaleBindingType)
	fieldNameMap["vmc_locale"] = "VmcLocale"
	paramsTypeMap["vmc_locale"] = bindings.NewReferenceType(model.VmcLocaleBindingType)
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
		dispatchHeaderParams,
		bodyFieldsMap,
		"",
		"vmc_locale",
		"POST",
		"/vmc/api/locale",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.unauthenticated": 401,"com.vmware.vapi.std.errors.unauthorized": 403})
}


