// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: Publish.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package msft_licensing

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	vmcModel "github.com/vmware/vsphere-automation-sdk-go/services/vmc/model"
	"reflect"
)

func publishPostInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org"] = vapiBindings_.NewStringType()
	fields["sddc"] = vapiBindings_.NewStringType()
	fields["cluster"] = vapiBindings_.NewStringType()
	fields["config_change"] = vapiBindings_.NewReferenceType(vmcModel.MsftLicensingConfigBindingType)
	fieldNameMap["org"] = "Org"
	fieldNameMap["sddc"] = "Sddc"
	fieldNameMap["cluster"] = "Cluster"
	fieldNameMap["config_change"] = "ConfigChange"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func PublishPostOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(vmcModel.TaskBindingType)
}

func publishPostRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["org"] = vapiBindings_.NewStringType()
	fields["sddc"] = vapiBindings_.NewStringType()
	fields["cluster"] = vapiBindings_.NewStringType()
	fields["config_change"] = vapiBindings_.NewReferenceType(vmcModel.MsftLicensingConfigBindingType)
	fieldNameMap["org"] = "Org"
	fieldNameMap["sddc"] = "Sddc"
	fieldNameMap["cluster"] = "Cluster"
	fieldNameMap["config_change"] = "ConfigChange"
	paramsTypeMap["cluster"] = vapiBindings_.NewStringType()
	paramsTypeMap["org"] = vapiBindings_.NewStringType()
	paramsTypeMap["sddc"] = vapiBindings_.NewStringType()
	paramsTypeMap["config_change"] = vapiBindings_.NewReferenceType(vmcModel.MsftLicensingConfigBindingType)
	paramsTypeMap["org"] = vapiBindings_.NewStringType()
	paramsTypeMap["sddc"] = vapiBindings_.NewStringType()
	paramsTypeMap["cluster"] = vapiBindings_.NewStringType()
	pathParams["cluster"] = "cluster"
	pathParams["org"] = "org"
	pathParams["sddc"] = "sddc"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
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
		"config_change",
		"POST",
		"/vmc/api/orgs/{org}/sddcs/{sddc}/clusters/{cluster}/msft-licensing/publish",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.unauthorized": 403})
}
