// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: CloudServiceVMCOnAWSTrafficGroup
// Used by client-side stubs.

package api

import (
	"github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-vmc-aws-integration/model"
)

const _ = core.SupportedByRuntimeVersion1

type CloudServiceVMCOnAWSTrafficGroupClient interface {

	// Delete the Traffic Group, disassociate the prefix list with the traffic group.
	//
	// @param trafficGroupIdParam (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	DeleteTrafficGroup(trafficGroupIdParam string) error

	// Delete the specified association map for a traffic group.
	//
	// @param trafficGroupIdParam Traffic group id (required)
	// @param mapIdParam Association map id (required)
	// @param forceParam Force delete the resource even if it is being used somewhere (optional, default to false)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	DeleteTrafficGroupAssociationMap(trafficGroupIdParam string, mapIdParam string, forceParam *bool) error

	// Get the traffic group information by traffic group ID.
	//
	// @param trafficGroupIdParam (required)
	// @return com.vmware.model.TrafficGroup
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetTrafficGroup(trafficGroupIdParam string) (model.TrafficGroup, error)

	// @param trafficGroupIdParam Traffic group id (required)
	// @param mapIdParam Association map id (required)
	// @return com.vmware.model.TrafficGroupAssociationMap
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetTrafficGroupAssociationMap(trafficGroupIdParam string, mapIdParam string) (model.TrafficGroupAssociationMap, error)

	// Retrieve association maps for a traffic group given its ID.
	//
	// @param trafficGroupIdParam Traffic group id (required)
	// @return com.vmware.model.TrafficGroupAssociationMapsListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ListTrafficGroupAssociationMaps(trafficGroupIdParam string) (model.TrafficGroupAssociationMapsListResult, error)

	// A list of all traffic groups, which are associated with prefix lists.
	//
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @param stateParam Filter parameter to filter Traffic Groups by Consolidated State (optional)
	// @param verboseParam Verbose info requested (optional, default to false)
	// @return com.vmware.model.TrafficGroupsListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ListTrafficGroups(cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string, stateParam *string, verboseParam *bool) (model.TrafficGroupsListResult, error)

	// Perform an action on a Traffic Group
	//
	// @param trafficGroupIdParam (required)
	// @param actionMessageParam (required)
	// @param actionParam Action performed on the Traffic Group (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	TriggerTrafficGroupAction(trafficGroupIdParam string, actionMessageParam model.ActionMessage, actionParam string) error

	// This API is used to create or update a traffic group, a Prefix List is associated to the traffic group.
	//
	// @param trafficGroupIdParam traffic group id (required)
	// @param trafficGroupParam (required)
	// @return com.vmware.model.TrafficGroup
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	UpdateTrafficGroup(trafficGroupIdParam string, trafficGroupParam model.TrafficGroup) (model.TrafficGroup, error)

	// @param trafficGroupIdParam Traffic group id (required)
	// @param mapIdParam Association map id (required)
	// @param trafficGroupAssociationMapParam (required)
	// @return com.vmware.model.TrafficGroupAssociationMap
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	UpdateTrafficGroupAssociationMap(trafficGroupIdParam string, mapIdParam string, trafficGroupAssociationMapParam model.TrafficGroupAssociationMap) (model.TrafficGroupAssociationMap, error)
}

type cloudServiceVMCOnAWSTrafficGroupClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewCloudServiceVMCOnAWSTrafficGroupClient(connector client.Connector) *cloudServiceVMCOnAWSTrafficGroupClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.api.cloud_service_VMC_on_AWS_traffic_group")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"delete_traffic_group":                 core.NewMethodIdentifier(interfaceIdentifier, "delete_traffic_group"),
		"delete_traffic_group_association_map": core.NewMethodIdentifier(interfaceIdentifier, "delete_traffic_group_association_map"),
		"get_traffic_group":                    core.NewMethodIdentifier(interfaceIdentifier, "get_traffic_group"),
		"get_traffic_group_association_map":    core.NewMethodIdentifier(interfaceIdentifier, "get_traffic_group_association_map"),
		"list_traffic_group_association_maps":  core.NewMethodIdentifier(interfaceIdentifier, "list_traffic_group_association_maps"),
		"list_traffic_groups":                  core.NewMethodIdentifier(interfaceIdentifier, "list_traffic_groups"),
		"trigger_traffic_group_action":         core.NewMethodIdentifier(interfaceIdentifier, "trigger_traffic_group_action"),
		"update_traffic_group":                 core.NewMethodIdentifier(interfaceIdentifier, "update_traffic_group"),
		"update_traffic_group_association_map": core.NewMethodIdentifier(interfaceIdentifier, "update_traffic_group_association_map"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	cIface := cloudServiceVMCOnAWSTrafficGroupClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &cIface
}

func (cIface *cloudServiceVMCOnAWSTrafficGroupClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := cIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (cIface *cloudServiceVMCOnAWSTrafficGroupClient) DeleteTrafficGroup(trafficGroupIdParam string) error {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(cloudServiceVMCOnAWSTrafficGroupDeleteTrafficGroupInputType(), typeConverter)
	sv.AddStructField("TrafficGroupId", trafficGroupIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := cloudServiceVMCOnAWSTrafficGroupDeleteTrafficGroupRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	cIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.api.cloud_service_VMC_on_AWS_traffic_group", "delete_traffic_group", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (cIface *cloudServiceVMCOnAWSTrafficGroupClient) DeleteTrafficGroupAssociationMap(trafficGroupIdParam string, mapIdParam string, forceParam *bool) error {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(cloudServiceVMCOnAWSTrafficGroupDeleteTrafficGroupAssociationMapInputType(), typeConverter)
	sv.AddStructField("TrafficGroupId", trafficGroupIdParam)
	sv.AddStructField("MapId", mapIdParam)
	sv.AddStructField("Force", forceParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := cloudServiceVMCOnAWSTrafficGroupDeleteTrafficGroupAssociationMapRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	cIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.api.cloud_service_VMC_on_AWS_traffic_group", "delete_traffic_group_association_map", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (cIface *cloudServiceVMCOnAWSTrafficGroupClient) GetTrafficGroup(trafficGroupIdParam string) (model.TrafficGroup, error) {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(cloudServiceVMCOnAWSTrafficGroupGetTrafficGroupInputType(), typeConverter)
	sv.AddStructField("TrafficGroupId", trafficGroupIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.TrafficGroup
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := cloudServiceVMCOnAWSTrafficGroupGetTrafficGroupRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	cIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.api.cloud_service_VMC_on_AWS_traffic_group", "get_traffic_group", inputDataValue, executionContext)
	var emptyOutput model.TrafficGroup
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), cloudServiceVMCOnAWSTrafficGroupGetTrafficGroupOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.TrafficGroup), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (cIface *cloudServiceVMCOnAWSTrafficGroupClient) GetTrafficGroupAssociationMap(trafficGroupIdParam string, mapIdParam string) (model.TrafficGroupAssociationMap, error) {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(cloudServiceVMCOnAWSTrafficGroupGetTrafficGroupAssociationMapInputType(), typeConverter)
	sv.AddStructField("TrafficGroupId", trafficGroupIdParam)
	sv.AddStructField("MapId", mapIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.TrafficGroupAssociationMap
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := cloudServiceVMCOnAWSTrafficGroupGetTrafficGroupAssociationMapRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	cIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.api.cloud_service_VMC_on_AWS_traffic_group", "get_traffic_group_association_map", inputDataValue, executionContext)
	var emptyOutput model.TrafficGroupAssociationMap
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), cloudServiceVMCOnAWSTrafficGroupGetTrafficGroupAssociationMapOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.TrafficGroupAssociationMap), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (cIface *cloudServiceVMCOnAWSTrafficGroupClient) ListTrafficGroupAssociationMaps(trafficGroupIdParam string) (model.TrafficGroupAssociationMapsListResult, error) {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(cloudServiceVMCOnAWSTrafficGroupListTrafficGroupAssociationMapsInputType(), typeConverter)
	sv.AddStructField("TrafficGroupId", trafficGroupIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.TrafficGroupAssociationMapsListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := cloudServiceVMCOnAWSTrafficGroupListTrafficGroupAssociationMapsRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	cIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.api.cloud_service_VMC_on_AWS_traffic_group", "list_traffic_group_association_maps", inputDataValue, executionContext)
	var emptyOutput model.TrafficGroupAssociationMapsListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), cloudServiceVMCOnAWSTrafficGroupListTrafficGroupAssociationMapsOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.TrafficGroupAssociationMapsListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (cIface *cloudServiceVMCOnAWSTrafficGroupClient) ListTrafficGroups(cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string, stateParam *string, verboseParam *bool) (model.TrafficGroupsListResult, error) {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(cloudServiceVMCOnAWSTrafficGroupListTrafficGroupsInputType(), typeConverter)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	sv.AddStructField("State", stateParam)
	sv.AddStructField("Verbose", verboseParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.TrafficGroupsListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := cloudServiceVMCOnAWSTrafficGroupListTrafficGroupsRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	cIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.api.cloud_service_VMC_on_AWS_traffic_group", "list_traffic_groups", inputDataValue, executionContext)
	var emptyOutput model.TrafficGroupsListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), cloudServiceVMCOnAWSTrafficGroupListTrafficGroupsOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.TrafficGroupsListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (cIface *cloudServiceVMCOnAWSTrafficGroupClient) TriggerTrafficGroupAction(trafficGroupIdParam string, actionMessageParam model.ActionMessage, actionParam string) error {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(cloudServiceVMCOnAWSTrafficGroupTriggerTrafficGroupActionInputType(), typeConverter)
	sv.AddStructField("TrafficGroupId", trafficGroupIdParam)
	sv.AddStructField("ActionMessage", actionMessageParam)
	sv.AddStructField("Action", actionParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := cloudServiceVMCOnAWSTrafficGroupTriggerTrafficGroupActionRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	cIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.api.cloud_service_VMC_on_AWS_traffic_group", "trigger_traffic_group_action", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (cIface *cloudServiceVMCOnAWSTrafficGroupClient) UpdateTrafficGroup(trafficGroupIdParam string, trafficGroupParam model.TrafficGroup) (model.TrafficGroup, error) {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(cloudServiceVMCOnAWSTrafficGroupUpdateTrafficGroupInputType(), typeConverter)
	sv.AddStructField("TrafficGroupId", trafficGroupIdParam)
	sv.AddStructField("TrafficGroup", trafficGroupParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.TrafficGroup
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := cloudServiceVMCOnAWSTrafficGroupUpdateTrafficGroupRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	cIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.api.cloud_service_VMC_on_AWS_traffic_group", "update_traffic_group", inputDataValue, executionContext)
	var emptyOutput model.TrafficGroup
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), cloudServiceVMCOnAWSTrafficGroupUpdateTrafficGroupOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.TrafficGroup), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (cIface *cloudServiceVMCOnAWSTrafficGroupClient) UpdateTrafficGroupAssociationMap(trafficGroupIdParam string, mapIdParam string, trafficGroupAssociationMapParam model.TrafficGroupAssociationMap) (model.TrafficGroupAssociationMap, error) {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(cloudServiceVMCOnAWSTrafficGroupUpdateTrafficGroupAssociationMapInputType(), typeConverter)
	sv.AddStructField("TrafficGroupId", trafficGroupIdParam)
	sv.AddStructField("MapId", mapIdParam)
	sv.AddStructField("TrafficGroupAssociationMap", trafficGroupAssociationMapParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.TrafficGroupAssociationMap
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := cloudServiceVMCOnAWSTrafficGroupUpdateTrafficGroupAssociationMapRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	cIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.api.cloud_service_VMC_on_AWS_traffic_group", "update_traffic_group_association_map", inputDataValue, executionContext)
	var emptyOutput model.TrafficGroupAssociationMap
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), cloudServiceVMCOnAWSTrafficGroupUpdateTrafficGroupAssociationMapOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.TrafficGroupAssociationMap), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
