/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: SiteRecoveryVersions.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package draas

import (
	"reflect"
	"github.com/vmware/vsphere-automation-sdk-go/services/vmc/draas/model"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
)

// Possible value for ``versionSource`` of method SiteRecoveryVersions#get.
const SiteRecoveryVersions_GET_VERSION_SOURCE_VAMICLI = "vamicli"
// Possible value for ``versionSource`` of method SiteRecoveryVersions#get.
const SiteRecoveryVersions_GET_VERSION_SOURCE_LS = "ls"




func siteRecoveryVersionsGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org"] = bindings.NewStringType()
	fields["sddc"] = bindings.NewStringType()
	fields["version_source"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["org"] = "Org"
	fieldNameMap["sddc"] = "Sddc"
	fieldNameMap["version_source"] = "VersionSource"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func siteRecoveryVersionsGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.SiteRecoveryVersionsBindingType)
}

func siteRecoveryVersionsGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["org"] = bindings.NewStringType()
	fields["sddc"] = bindings.NewStringType()
	fields["version_source"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["org"] = "Org"
	fieldNameMap["sddc"] = "Sddc"
	fieldNameMap["version_source"] = "VersionSource"
	paramsTypeMap["org"] = bindings.NewStringType()
	paramsTypeMap["sddc"] = bindings.NewStringType()
	paramsTypeMap["version_source"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["org"] = bindings.NewStringType()
	paramsTypeMap["sddc"] = bindings.NewStringType()
	pathParams["org"] = "org"
	pathParams["sddc"] = "sddc"
	queryParams["version_source"] = "version_source"
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
		"",
		"GET",
		"/vmc/draas/api/orgs/{org}/sddcs/{sddc}/site-recovery/versions",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.unauthenticated": 401,"com.vmware.vapi.std.errors.invalid_request": 400,"com.vmware.vapi.std.errors.unauthorized": 403,"com.vmware.vapi.std.errors.not_found": 404})
}


