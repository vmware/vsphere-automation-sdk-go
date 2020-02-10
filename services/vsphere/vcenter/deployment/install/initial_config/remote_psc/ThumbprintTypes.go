/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: Thumbprint.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package remote_psc

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)



// The ``RemoteSpec`` class contains the information used to connect to the remote PSC. This class was added in vSphere API 6.7.
type ThumbprintRemoteSpec struct {
    // The IP address or DNS resolvable name of the remote PSC. This property was added in vSphere API 6.7.
	Address string
    // The HTTPS port of the remote PSC. This property was added in vSphere API 6.7.
	HttpsPort *int64
}



func thumbprintGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["spec"] = bindings.NewReferenceType(ThumbprintRemoteSpecBindingType)
	fieldNameMap["spec"] = "Spec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func thumbprintGetOutputType() bindings.BindingType {
	return bindings.NewStringType()
}

func thumbprintGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["spec"] = bindings.NewReferenceType(ThumbprintRemoteSpecBindingType)
	fieldNameMap["spec"] = "Spec"
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
		"null",
		"",
		resultHeaders,
		0,
		errorHeaders,
		map[string]int{"Unauthenticated": 401,"Error": 500})
}


func ThumbprintRemoteSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["address"] = bindings.NewStringType()
	fieldNameMap["address"] = "Address"
	fields["https_port"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["https_port"] = "HttpsPort"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.deployment.install.initial_config.remote_psc.thumbprint.remote_spec", fields, reflect.TypeOf(ThumbprintRemoteSpec{}), fieldNameMap, validators)
}

