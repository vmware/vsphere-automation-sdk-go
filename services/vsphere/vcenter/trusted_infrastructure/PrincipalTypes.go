/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: Principal.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package trusted_infrastructure

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)



// The ``Info`` class contains the information about the principal and certificates used by this vCenter to retrieve tokens. This class was added in vSphere API 7.0.0.
type PrincipalInfo struct {
    // The certificates used by the STS to sign tokens for this vCenter. This property was added in vSphere API 7.0.0.
	Certificates []X509CertChain
    // The service which created and signed the security token. This property was added in vSphere API 7.0.0.
	Issuer string
    // The principal used by this vCenter instance to retrieve tokens. Currently this is the vCenter solution user. This property was added in vSphere API 7.0.0.
	Principal StsPrincipal
    // The user-friednly name of the vCenter. This property was added in vSphere API 7.0.0.
	Name string
}



func principalGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func principalGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(PrincipalInfoBindingType)
}

func principalGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
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
		"/vcenter/trusted-infrastructure/principal",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"Error": 500,"Unauthenticated": 401})
}


func PrincipalInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["certificates"] = bindings.NewListType(bindings.NewReferenceType(X509CertChainBindingType), reflect.TypeOf([]X509CertChain{}))
	fieldNameMap["certificates"] = "Certificates"
	fields["issuer"] = bindings.NewStringType()
	fieldNameMap["issuer"] = "Issuer"
	fields["principal"] = bindings.NewReferenceType(StsPrincipalBindingType)
	fieldNameMap["principal"] = "Principal"
	fields["name"] = bindings.NewStringType()
	fieldNameMap["name"] = "Name"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.trusted_infrastructure.principal.info", fields, reflect.TypeOf(PrincipalInfo{}), fieldNameMap, validators)
}

