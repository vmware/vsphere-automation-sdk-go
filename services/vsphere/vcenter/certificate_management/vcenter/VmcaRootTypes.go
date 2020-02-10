/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: VmcaRoot.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package vcenter

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)



// The ``CreateSpec`` contains information. to generate a Private Key and CSR. This class was added in vSphere API 6.9.1.
type VmcaRootCreateSpec struct {
    // The size of the key to be used for public and private key generation. This property was added in vSphere API 6.9.1.
	KeySize *int64
    // The common name of the host for which certificate is generated. This property was added in vSphere API 6.9.1.
	CommonName *string
    // Organization field in certificate subject. This property was added in vSphere API 6.9.1.
	Organization *string
    // Organization unit field in certificate subject. This property was added in vSphere API 6.9.1.
	OrganizationUnit *string
    // Locality field in certificate subject. This property was added in vSphere API 6.9.1.
	Locality *string
    // State field in certificate subject. This property was added in vSphere API 6.9.1.
	StateOrProvince *string
    // Country field in certificate subject. This property was added in vSphere API 6.9.1.
	Country *string
    // Email field in Certificate extensions. This property was added in vSphere API 6.9.1.
	EmailAddress *string
    // SubjectAltName is list of Dns Names and Ip addresses. This property was added in vSphere API 6.9.1.
	SubjectAltName []string
}



func vmcaRootCreateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["spec"] = bindings.NewOptionalType(bindings.NewReferenceType(VmcaRootCreateSpecBindingType))
	fieldNameMap["spec"] = "Spec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func vmcaRootCreateOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func vmcaRootCreateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["spec"] = bindings.NewOptionalType(bindings.NewReferenceType(VmcaRootCreateSpecBindingType))
	fieldNameMap["spec"] = "Spec"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
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
		map[string]int{"Error": 500})
}


func VmcaRootCreateSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["key_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["key_size"] = "KeySize"
	fields["common_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["common_name"] = "CommonName"
	fields["organization"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["organization"] = "Organization"
	fields["organization_unit"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["organization_unit"] = "OrganizationUnit"
	fields["locality"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["locality"] = "Locality"
	fields["state_or_province"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["state_or_province"] = "StateOrProvince"
	fields["country"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["country"] = "Country"
	fields["email_address"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["email_address"] = "EmailAddress"
	fields["subject_alt_name"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["subject_alt_name"] = "SubjectAltName"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.certificate_management.vcenter.vmca_root.create_spec", fields, reflect.TypeOf(VmcaRootCreateSpec{}), fieldNameMap, validators)
}

