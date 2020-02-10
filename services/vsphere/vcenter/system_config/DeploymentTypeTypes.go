/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: DeploymentType.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package system_config

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/services/vsphere/vcenter/deployment"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)



// The ``Info`` class contains the fields used to get the appliance type. This class was added in vSphere API 6.7.
type DeploymentTypeInfo struct {
    // The type of the appliance. This property was added in vSphere API 6.7.
	Type_ deployment.ApplianceType
}

// The ``ReconfigureSpec`` class contains the fields used to get and set the appliance type. This class was added in vSphere API 6.7.
type DeploymentTypeReconfigureSpec struct {
    // The type of the appliance. This property was added in vSphere API 6.7.
	Type_ deployment.ApplianceType
    // External PSC to register with when reconfiguring a VCSA_EMBEDDED appliance to a VCSA_EXTERNAL appliance. This property was added in vSphere API 6.7.
	RemotePsc *deployment.RemotePscSpec
}



func deploymentTypeGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func deploymentTypeGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(DeploymentTypeInfoBindingType)
}

func deploymentTypeGetRestMetadata() protocol.OperationRestMetadata {
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

func deploymentTypeReconfigureInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["spec"] = bindings.NewReferenceType(DeploymentTypeReconfigureSpecBindingType)
	fieldNameMap["spec"] = "Spec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func deploymentTypeReconfigureOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func deploymentTypeReconfigureRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["spec"] = bindings.NewReferenceType(DeploymentTypeReconfigureSpecBindingType)
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
		map[string]int{"Unsupported": 400,"InvalidArgument": 400,"Unauthenticated": 401,"NotAllowedInCurrentState": 400})
}


func DeploymentTypeInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["type"] = bindings.NewEnumType("com.vmware.vcenter.deployment.appliance_type", reflect.TypeOf(deployment.ApplianceType(deployment.ApplianceType_VCSA_EMBEDDED)))
	fieldNameMap["type"] = "Type_"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.system_config.deployment_type.info", fields, reflect.TypeOf(DeploymentTypeInfo{}), fieldNameMap, validators)
}

func DeploymentTypeReconfigureSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["type"] = bindings.NewEnumType("com.vmware.vcenter.deployment.appliance_type", reflect.TypeOf(deployment.ApplianceType(deployment.ApplianceType_VCSA_EMBEDDED)))
	fieldNameMap["type"] = "Type_"
	fields["remote_psc"] = bindings.NewOptionalType(bindings.NewReferenceType(deployment.RemotePscSpecBindingType))
	fieldNameMap["remote_psc"] = "RemotePsc"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.system_config.deployment_type.reconfigure_spec", fields, reflect.TypeOf(DeploymentTypeReconfigureSpec{}), fieldNameMap, validators)
}

