// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: DeploymentIddeleteHost.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package deployments

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	vmwarecloudVmc_aws_cho_providerModel "github.com/vmware/vsphere-automation-sdk-go/services/vmwarecloud/vmc_aws_cho_provider/model"
	"reflect"
)

func deploymentIddeleteHostPostDeploymentDeleteHostInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org_id"] = vapiBindings_.NewStringType()
	fields["deployment_id"] = vapiBindings_.NewStringType()
	fields["delete_host_intent"] = vapiBindings_.NewReferenceType(vmwarecloudVmc_aws_cho_providerModel.DeleteHostIntentBindingType)
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["deployment_id"] = "DeploymentId"
	fieldNameMap["delete_host_intent"] = "DeleteHostIntent"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func DeploymentIddeleteHostPostDeploymentDeleteHostOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(vmwarecloudVmc_aws_cho_providerModel.ActivityResponseBindingType)
}

func deploymentIddeleteHostPostDeploymentDeleteHostRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["org_id"] = vapiBindings_.NewStringType()
	fields["deployment_id"] = vapiBindings_.NewStringType()
	fields["delete_host_intent"] = vapiBindings_.NewReferenceType(vmwarecloudVmc_aws_cho_providerModel.DeleteHostIntentBindingType)
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["deployment_id"] = "DeploymentId"
	fieldNameMap["delete_host_intent"] = "DeleteHostIntent"
	paramsTypeMap["org_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["deployment_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["delete_host_intent"] = vapiBindings_.NewReferenceType(vmwarecloudVmc_aws_cho_providerModel.DeleteHostIntentBindingType)
	paramsTypeMap["orgId"] = vapiBindings_.NewStringType()
	paramsTypeMap["deploymentId"] = vapiBindings_.NewStringType()
	pathParams["deployment_id"] = "deploymentId"
	pathParams["org_id"] = "orgId"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	errorHeaders["com.vmware.vapi.std.errors.unauthenticated"] = make(map[string]string)
	errorHeaders["com.vmware.vapi.std.errors.unauthenticated"]["challenge"] = "WWW-Authenticate"
	return vapiProtocol_.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		dispatchHeaderParams,
		bodyFieldsMap,
		"",
		"delete_host_intent",
		"POST",
		"/api/inventory/{orgId}/vmc-aws/deployments/{deploymentId}:delete-host",
		"application/json",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthenticated": 401, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.not_found": 404})
}