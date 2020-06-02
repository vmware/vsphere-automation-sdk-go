/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: Rules.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package config

import (
	"reflect"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/vmc/model"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/protocol"
)





func rulesAddInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org"] = bindings.NewStringType()
	fields["sddc"] = bindings.NewStringType()
	fields["edge_id"] = bindings.NewStringType()
	fields["nat_rules"] = bindings.NewReferenceType(model.NatRulesBindingType)
	fieldNameMap["org"] = "Org"
	fieldNameMap["sddc"] = "Sddc"
	fieldNameMap["edge_id"] = "EdgeId"
	fieldNameMap["nat_rules"] = "NatRules"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func rulesAddOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func rulesAddRestMetadata() protocol.OperationRestMetadata {
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
	fields["edge_id"] = bindings.NewStringType()
	fields["nat_rules"] = bindings.NewReferenceType(model.NatRulesBindingType)
	fieldNameMap["org"] = "Org"
	fieldNameMap["sddc"] = "Sddc"
	fieldNameMap["edge_id"] = "EdgeId"
	fieldNameMap["nat_rules"] = "NatRules"
	paramsTypeMap["org"] = bindings.NewStringType()
	paramsTypeMap["sddc"] = bindings.NewStringType()
	paramsTypeMap["nat_rules"] = bindings.NewReferenceType(model.NatRulesBindingType)
	paramsTypeMap["edge_id"] = bindings.NewStringType()
	paramsTypeMap["org"] = bindings.NewStringType()
	paramsTypeMap["sddc"] = bindings.NewStringType()
	paramsTypeMap["edgeId"] = bindings.NewStringType()
	pathParams["edge_id"] = "edgeId"
	pathParams["org"] = "org"
	pathParams["sddc"] = "sddc"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
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
		"nat_rules",
		"POST",
		"/vmc/api/orgs/{org}/sddcs/{sddc}/networks/4.0/edges/{edgeId}/nat/config/rules",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400,"com.vmware.vapi.std.errors.unauthorized": 403,"com.vmware.vapi.std.errors.not_found": 404})
}

func rulesDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org"] = bindings.NewStringType()
	fields["sddc"] = bindings.NewStringType()
	fields["edge_id"] = bindings.NewStringType()
	fields["rule_id"] = bindings.NewIntegerType()
	fieldNameMap["org"] = "Org"
	fieldNameMap["sddc"] = "Sddc"
	fieldNameMap["edge_id"] = "EdgeId"
	fieldNameMap["rule_id"] = "RuleId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func rulesDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func rulesDeleteRestMetadata() protocol.OperationRestMetadata {
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
	fields["edge_id"] = bindings.NewStringType()
	fields["rule_id"] = bindings.NewIntegerType()
	fieldNameMap["org"] = "Org"
	fieldNameMap["sddc"] = "Sddc"
	fieldNameMap["edge_id"] = "EdgeId"
	fieldNameMap["rule_id"] = "RuleId"
	paramsTypeMap["rule_id"] = bindings.NewIntegerType()
	paramsTypeMap["org"] = bindings.NewStringType()
	paramsTypeMap["sddc"] = bindings.NewStringType()
	paramsTypeMap["edge_id"] = bindings.NewStringType()
	paramsTypeMap["org"] = bindings.NewStringType()
	paramsTypeMap["sddc"] = bindings.NewStringType()
	paramsTypeMap["edgeId"] = bindings.NewStringType()
	paramsTypeMap["ruleId"] = bindings.NewIntegerType()
	pathParams["edge_id"] = "edgeId"
	pathParams["org"] = "org"
	pathParams["sddc"] = "sddc"
	pathParams["rule_id"] = "ruleId"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
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
		"DELETE",
		"/vmc/api/orgs/{org}/sddcs/{sddc}/networks/4.0/edges/{edgeId}/nat/config/rules/{ruleId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400,"com.vmware.vapi.std.errors.unauthorized": 403,"com.vmware.vapi.std.errors.not_found": 404})
}

func rulesUpdateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org"] = bindings.NewStringType()
	fields["sddc"] = bindings.NewStringType()
	fields["edge_id"] = bindings.NewStringType()
	fields["rule_id"] = bindings.NewIntegerType()
	fields["nsxnatrule"] = bindings.NewReferenceType(model.NsxnatruleBindingType)
	fieldNameMap["org"] = "Org"
	fieldNameMap["sddc"] = "Sddc"
	fieldNameMap["edge_id"] = "EdgeId"
	fieldNameMap["rule_id"] = "RuleId"
	fieldNameMap["nsxnatrule"] = "Nsxnatrule"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func rulesUpdateOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func rulesUpdateRestMetadata() protocol.OperationRestMetadata {
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
	fields["edge_id"] = bindings.NewStringType()
	fields["rule_id"] = bindings.NewIntegerType()
	fields["nsxnatrule"] = bindings.NewReferenceType(model.NsxnatruleBindingType)
	fieldNameMap["org"] = "Org"
	fieldNameMap["sddc"] = "Sddc"
	fieldNameMap["edge_id"] = "EdgeId"
	fieldNameMap["rule_id"] = "RuleId"
	fieldNameMap["nsxnatrule"] = "Nsxnatrule"
	paramsTypeMap["rule_id"] = bindings.NewIntegerType()
	paramsTypeMap["org"] = bindings.NewStringType()
	paramsTypeMap["sddc"] = bindings.NewStringType()
	paramsTypeMap["edge_id"] = bindings.NewStringType()
	paramsTypeMap["nsxnatrule"] = bindings.NewReferenceType(model.NsxnatruleBindingType)
	paramsTypeMap["org"] = bindings.NewStringType()
	paramsTypeMap["sddc"] = bindings.NewStringType()
	paramsTypeMap["edgeId"] = bindings.NewStringType()
	paramsTypeMap["ruleId"] = bindings.NewIntegerType()
	pathParams["edge_id"] = "edgeId"
	pathParams["org"] = "org"
	pathParams["sddc"] = "sddc"
	pathParams["rule_id"] = "ruleId"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
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
		"nsxnatrule",
		"PUT",
		"/vmc/api/orgs/{org}/sddcs/{sddc}/networks/4.0/edges/{edgeId}/nat/config/rules/{ruleId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400,"com.vmware.vapi.std.errors.unauthorized": 403,"com.vmware.vapi.std.errors.not_found": 404})
}


