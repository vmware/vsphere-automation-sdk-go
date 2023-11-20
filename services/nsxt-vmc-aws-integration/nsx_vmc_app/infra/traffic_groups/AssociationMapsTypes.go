// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: AssociationMaps.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package traffic_groups

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	nsx_vmc_appModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-vmc-aws-integration/nsx_vmc_app/model"
	"reflect"
)

func associationMapsDeleteInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["traffic_group_id"] = vapiBindings_.NewStringType()
	fields["map_id"] = vapiBindings_.NewStringType()
	fields["force"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["traffic_group_id"] = "TrafficGroupId"
	fieldNameMap["map_id"] = "MapId"
	fieldNameMap["force"] = "Force"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func AssociationMapsDeleteOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func associationMapsDeleteRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["traffic_group_id"] = vapiBindings_.NewStringType()
	fields["map_id"] = vapiBindings_.NewStringType()
	fields["force"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["traffic_group_id"] = "TrafficGroupId"
	fieldNameMap["map_id"] = "MapId"
	fieldNameMap["force"] = "Force"
	paramsTypeMap["map_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["traffic_group_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["force"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	paramsTypeMap["trafficGroupId"] = vapiBindings_.NewStringType()
	paramsTypeMap["mapId"] = vapiBindings_.NewStringType()
	pathParams["map_id"] = "mapId"
	pathParams["traffic_group_id"] = "trafficGroupId"
	queryParams["force"] = "force"
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
		"",
		"DELETE",
		"/cloud-service/api/v1/infra/traffic-groups/{trafficGroupId}/association-maps/{mapId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func associationMapsGetInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["traffic_group_id"] = vapiBindings_.NewStringType()
	fields["map_id"] = vapiBindings_.NewStringType()
	fieldNameMap["traffic_group_id"] = "TrafficGroupId"
	fieldNameMap["map_id"] = "MapId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func AssociationMapsGetOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_vmc_appModel.TrafficGroupAssociationMapBindingType)
}

func associationMapsGetRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["traffic_group_id"] = vapiBindings_.NewStringType()
	fields["map_id"] = vapiBindings_.NewStringType()
	fieldNameMap["traffic_group_id"] = "TrafficGroupId"
	fieldNameMap["map_id"] = "MapId"
	paramsTypeMap["map_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["traffic_group_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["trafficGroupId"] = vapiBindings_.NewStringType()
	paramsTypeMap["mapId"] = vapiBindings_.NewStringType()
	pathParams["map_id"] = "mapId"
	pathParams["traffic_group_id"] = "trafficGroupId"
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
		"",
		"GET",
		"/cloud-service/api/v1/infra/traffic-groups/{trafficGroupId}/association-maps/{mapId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func associationMapsListInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["traffic_group_id"] = vapiBindings_.NewStringType()
	fieldNameMap["traffic_group_id"] = "TrafficGroupId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func AssociationMapsListOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_vmc_appModel.TrafficGroupAssociationMapsListResultBindingType)
}

func associationMapsListRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["traffic_group_id"] = vapiBindings_.NewStringType()
	fieldNameMap["traffic_group_id"] = "TrafficGroupId"
	paramsTypeMap["traffic_group_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["trafficGroupId"] = vapiBindings_.NewStringType()
	pathParams["traffic_group_id"] = "trafficGroupId"
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
		"",
		"GET",
		"/cloud-service/api/v1/infra/traffic-groups/{trafficGroupId}/association-maps",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func associationMapsUpdateInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["traffic_group_id"] = vapiBindings_.NewStringType()
	fields["map_id"] = vapiBindings_.NewStringType()
	fields["traffic_group_association_map"] = vapiBindings_.NewReferenceType(nsx_vmc_appModel.TrafficGroupAssociationMapBindingType)
	fieldNameMap["traffic_group_id"] = "TrafficGroupId"
	fieldNameMap["map_id"] = "MapId"
	fieldNameMap["traffic_group_association_map"] = "TrafficGroupAssociationMap"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func AssociationMapsUpdateOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_vmc_appModel.TrafficGroupAssociationMapBindingType)
}

func associationMapsUpdateRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["traffic_group_id"] = vapiBindings_.NewStringType()
	fields["map_id"] = vapiBindings_.NewStringType()
	fields["traffic_group_association_map"] = vapiBindings_.NewReferenceType(nsx_vmc_appModel.TrafficGroupAssociationMapBindingType)
	fieldNameMap["traffic_group_id"] = "TrafficGroupId"
	fieldNameMap["map_id"] = "MapId"
	fieldNameMap["traffic_group_association_map"] = "TrafficGroupAssociationMap"
	paramsTypeMap["map_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["traffic_group_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["traffic_group_association_map"] = vapiBindings_.NewReferenceType(nsx_vmc_appModel.TrafficGroupAssociationMapBindingType)
	paramsTypeMap["trafficGroupId"] = vapiBindings_.NewStringType()
	paramsTypeMap["mapId"] = vapiBindings_.NewStringType()
	pathParams["map_id"] = "mapId"
	pathParams["traffic_group_id"] = "trafficGroupId"
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
		"traffic_group_association_map",
		"PUT",
		"/cloud-service/api/v1/infra/traffic-groups/{trafficGroupId}/association-maps/{mapId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
