// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: Versionpromote.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package rules

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	vmwarecloudVmc_aws_cho_providerModel "github.com/vmware/vsphere-automation-sdk-go/services/vmwarecloud/vmc_aws_cho_provider/model"
	"reflect"
)

func versionpromotePromoteRuleVersionInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["rule_version_promotion_payload"] = vapiBindings_.NewReferenceType(vmwarecloudVmc_aws_cho_providerModel.RuleVersionPromotionPayloadBindingType)
	fieldNameMap["rule_version_promotion_payload"] = "RuleVersionPromotionPayload"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func VersionpromotePromoteRuleVersionOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func versionpromotePromoteRuleVersionRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["rule_version_promotion_payload"] = vapiBindings_.NewReferenceType(vmwarecloudVmc_aws_cho_providerModel.RuleVersionPromotionPayloadBindingType)
	fieldNameMap["rule_version_promotion_payload"] = "RuleVersionPromotionPayload"
	paramsTypeMap["rule_version_promotion_payload"] = vapiBindings_.NewReferenceType(vmwarecloudVmc_aws_cho_providerModel.RuleVersionPromotionPayloadBindingType)
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
		"rule_version_promotion_payload",
		"POST",
		"/api/inventory/vmc-aws/workflow/rules/version:promote",
		"application/json",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.unauthenticated": 401, "com.vmware.vapi.std.errors.unauthorized": 403})
}
