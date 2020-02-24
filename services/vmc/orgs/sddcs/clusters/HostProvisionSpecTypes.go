/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: HostProvisionSpec.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package clusters

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/services/vmc/model"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)





func hostProvisionSpecGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org"] = bindings.NewStringType()
	fields["sddc"] = bindings.NewStringType()
	fields["cluster"] = bindings.NewStringType()
	fields["action"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["org"] = "Org"
	fieldNameMap["sddc"] = "Sddc"
	fieldNameMap["cluster"] = "Cluster"
	fieldNameMap["action"] = "Action"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func hostProvisionSpecGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.HostProvisionSpecBindingType)
}

func hostProvisionSpecGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["org"] = bindings.NewStringType()
	fields["sddc"] = bindings.NewStringType()
	fields["cluster"] = bindings.NewStringType()
	fields["action"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["org"] = "Org"
	fieldNameMap["sddc"] = "Sddc"
	fieldNameMap["cluster"] = "Cluster"
	fieldNameMap["action"] = "Action"
	paramsTypeMap["cluster"] = bindings.NewStringType()
	paramsTypeMap["action"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["org"] = bindings.NewStringType()
	paramsTypeMap["sddc"] = bindings.NewStringType()
	paramsTypeMap["org"] = bindings.NewStringType()
	paramsTypeMap["sddc"] = bindings.NewStringType()
	paramsTypeMap["cluster"] = bindings.NewStringType()
	pathParams["cluster"] = "cluster"
	pathParams["org"] = "org"
	pathParams["sddc"] = "sddc"
	queryParams["action"] = "action"
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
		"GET",
		"/vmc/api/orgs/{org}/sddcs/{sddc}/clusters/{cluster}/host-provision-spec",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"Unauthenticated": 401,"Unauthorized": 403,"InternalServerError": 500})
}


