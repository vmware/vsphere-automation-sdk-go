/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: Install.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package deployment

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)



// The ``VcsaEmbeddedSpec`` class contains information used to configure an embedded standalone or replicated vCenter Server. This class was added in vSphere API 6.7.
type InstallVcsaEmbeddedSpec struct {
    // Spec used to configure a standalone embedded vCenter Server. This field describes how the standalone vCenter Server appliance should be configured. This property was added in vSphere API 6.7.
	Standalone *StandaloneSpec
    // Spec used to configure a replicated embedded vCenter Server. This field describes how the replicated vCenter Server appliance should be configured. This property was added in vSphere API 6.7.
	Replicated *ReplicatedSpec
    // Whether CEIP should be enabled or disabled. This property was added in vSphere API 6.7.
	CeipEnabled bool
}

// The ``InstallSpec`` class contains information used to configure the appliance installation. This class was added in vSphere API 6.7.
type InstallInstallSpec struct {
    // Spec used to configure an embedded vCenter Server. This field describes how the embedded vCenter Server appliance should be configured. This property was added in vSphere API 6.7.
	VcsaEmbedded InstallVcsaEmbeddedSpec
    // Use the default option for any questions that may come up during appliance configuration. This property was added in vSphere API 6.7.
	AutoAnswer *bool
}



func installGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func installGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(InstallInstallSpecBindingType)
}

func installGetRestMetadata() protocol.OperationRestMetadata {
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
		"null",
		"",
		resultHeaders,
		0,
		errorHeaders,
		map[string]int{"Unauthenticated": 401,"NotAllowedInCurrentState": 400})
}

func installCheckInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["spec"] = bindings.NewReferenceType(InstallInstallSpecBindingType)
	fieldNameMap["spec"] = "Spec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func installCheckOutputType() bindings.BindingType {
	return bindings.NewReferenceType(CheckInfoBindingType)
}

func installCheckRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["spec"] = bindings.NewReferenceType(InstallInstallSpecBindingType)
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
		map[string]int{"Unauthenticated": 401,"InvalidArgument": 400,"NotAllowedInCurrentState": 400})
}

func installStartInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["spec"] = bindings.NewReferenceType(InstallInstallSpecBindingType)
	fieldNameMap["spec"] = "Spec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func installStartOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func installStartRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["spec"] = bindings.NewReferenceType(InstallInstallSpecBindingType)
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
		map[string]int{"Unauthenticated": 401,"InvalidArgument": 400,"NotAllowedInCurrentState": 400})
}

func installCancelInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func installCancelOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func installCancelRestMetadata() protocol.OperationRestMetadata {
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
		"null",
		"",
		resultHeaders,
		0,
		errorHeaders,
		map[string]int{"Unauthenticated": 401,"NotAllowedInCurrentState": 400})
}


func InstallVcsaEmbeddedSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["standalone"] = bindings.NewOptionalType(bindings.NewReferenceType(StandaloneSpecBindingType))
	fieldNameMap["standalone"] = "Standalone"
	fields["replicated"] = bindings.NewOptionalType(bindings.NewReferenceType(ReplicatedSpecBindingType))
	fieldNameMap["replicated"] = "Replicated"
	fields["ceip_enabled"] = bindings.NewBooleanType()
	fieldNameMap["ceip_enabled"] = "CeipEnabled"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.deployment.install.vcsa_embedded_spec", fields, reflect.TypeOf(InstallVcsaEmbeddedSpec{}), fieldNameMap, validators)
}

func InstallInstallSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vcsa_embedded"] = bindings.NewReferenceType(InstallVcsaEmbeddedSpecBindingType)
	fieldNameMap["vcsa_embedded"] = "VcsaEmbedded"
	fields["auto_answer"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["auto_answer"] = "AutoAnswer"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.deployment.install.install_spec", fields, reflect.TypeOf(InstallInstallSpec{}), fieldNameMap, validators)
}

