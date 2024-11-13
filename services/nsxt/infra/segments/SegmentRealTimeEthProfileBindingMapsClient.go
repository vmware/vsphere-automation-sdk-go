// Copyright (c) 2019-2024 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: SegmentRealTimeEthProfileBindingMaps
// Used by client-side stubs.

package segments

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	nsx_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type SegmentRealTimeEthProfileBindingMapsClient interface {

	// API will delete Segment Real Time Eth Profile Binding Profile.
	//
	// @param segmentIdParam Segment ID (required)
	// @param segmentRealTimeEthProfileBindingMapIdParam Segment Real Time Eth Profile Binding Map ID (required)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Delete(segmentIdParam string, segmentRealTimeEthProfileBindingMapIdParam string) error

	// API will get Segment Real Time Eth Profile Binding Map.
	//
	// @param segmentIdParam Segment ID (required)
	// @param segmentRealTimeEthProfileBindingMapIdParam Segment Real Time Eth Profile Binding Map ID (required)
	// @return com.vmware.nsx_policy.model.SegmentRealTimeEthProfileBindingMap
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(segmentIdParam string, segmentRealTimeEthProfileBindingMapIdParam string) (nsx_policyModel.SegmentRealTimeEthProfileBindingMap, error)

	// API will list all Segment Real Time Eth Profile Binding Maps in current segment id.
	//
	// @param segmentIdParam (required)
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @return com.vmware.nsx_policy.model.SegmentRealTimeEthProfileBindingMapListResult
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(segmentIdParam string, cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (nsx_policyModel.SegmentRealTimeEthProfileBindingMapListResult, error)

	// API will create Segment Real Time Eth profile binding map. For objects with no binding maps, default profile is applied.
	//
	// @param segmentIdParam Segment ID (required)
	// @param segmentRealTimeEthProfileBindingMapIdParam Segment Real Time Eth Profile Binding Map ID (required)
	// @param segmentRealTimeEthProfileBindingMapParam (required)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Patch(segmentIdParam string, segmentRealTimeEthProfileBindingMapIdParam string, segmentRealTimeEthProfileBindingMapParam nsx_policyModel.SegmentRealTimeEthProfileBindingMap) error

	// API will update Segment Real Time Eth Profile Binding Map. For objects with no binding maps, default profile is applied.
	//
	// @param segmentIdParam Segment ID (required)
	// @param segmentRealTimeEthProfileBindingMapIdParam Segment Real Time Eth Profile Binding Map ID (required)
	// @param segmentRealTimeEthProfileBindingMapParam (required)
	// @return com.vmware.nsx_policy.model.SegmentRealTimeEthProfileBindingMap
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(segmentIdParam string, segmentRealTimeEthProfileBindingMapIdParam string, segmentRealTimeEthProfileBindingMapParam nsx_policyModel.SegmentRealTimeEthProfileBindingMap) (nsx_policyModel.SegmentRealTimeEthProfileBindingMap, error)
}

type segmentRealTimeEthProfileBindingMapsClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewSegmentRealTimeEthProfileBindingMapsClient(connector vapiProtocolClient_.Connector) *segmentRealTimeEthProfileBindingMapsClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx_policy.infra.segments.segment_real_time_eth_profile_binding_maps")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"delete": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "delete"),
		"get":    vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
		"list":   vapiCore_.NewMethodIdentifier(interfaceIdentifier, "list"),
		"patch":  vapiCore_.NewMethodIdentifier(interfaceIdentifier, "patch"),
		"update": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "update"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	sIface := segmentRealTimeEthProfileBindingMapsClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &sIface
}

func (sIface *segmentRealTimeEthProfileBindingMapsClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := sIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (sIface *segmentRealTimeEthProfileBindingMapsClient) Delete(segmentIdParam string, segmentRealTimeEthProfileBindingMapIdParam string) error {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	operationRestMetaData := segmentRealTimeEthProfileBindingMapsDeleteRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(segmentRealTimeEthProfileBindingMapsDeleteInputType(), typeConverter)
	sv.AddStructField("SegmentId", segmentIdParam)
	sv.AddStructField("SegmentRealTimeEthProfileBindingMapId", segmentRealTimeEthProfileBindingMapIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.segments.segment_real_time_eth_profile_binding_maps", "delete", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (sIface *segmentRealTimeEthProfileBindingMapsClient) Get(segmentIdParam string, segmentRealTimeEthProfileBindingMapIdParam string) (nsx_policyModel.SegmentRealTimeEthProfileBindingMap, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	operationRestMetaData := segmentRealTimeEthProfileBindingMapsGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(segmentRealTimeEthProfileBindingMapsGetInputType(), typeConverter)
	sv.AddStructField("SegmentId", segmentIdParam)
	sv.AddStructField("SegmentRealTimeEthProfileBindingMapId", segmentRealTimeEthProfileBindingMapIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.SegmentRealTimeEthProfileBindingMap
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.segments.segment_real_time_eth_profile_binding_maps", "get", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.SegmentRealTimeEthProfileBindingMap
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), SegmentRealTimeEthProfileBindingMapsGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.SegmentRealTimeEthProfileBindingMap), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *segmentRealTimeEthProfileBindingMapsClient) List(segmentIdParam string, cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (nsx_policyModel.SegmentRealTimeEthProfileBindingMapListResult, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	operationRestMetaData := segmentRealTimeEthProfileBindingMapsListRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(segmentRealTimeEthProfileBindingMapsListInputType(), typeConverter)
	sv.AddStructField("SegmentId", segmentIdParam)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.SegmentRealTimeEthProfileBindingMapListResult
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.segments.segment_real_time_eth_profile_binding_maps", "list", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.SegmentRealTimeEthProfileBindingMapListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), SegmentRealTimeEthProfileBindingMapsListOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.SegmentRealTimeEthProfileBindingMapListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *segmentRealTimeEthProfileBindingMapsClient) Patch(segmentIdParam string, segmentRealTimeEthProfileBindingMapIdParam string, segmentRealTimeEthProfileBindingMapParam nsx_policyModel.SegmentRealTimeEthProfileBindingMap) error {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	operationRestMetaData := segmentRealTimeEthProfileBindingMapsPatchRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(segmentRealTimeEthProfileBindingMapsPatchInputType(), typeConverter)
	sv.AddStructField("SegmentId", segmentIdParam)
	sv.AddStructField("SegmentRealTimeEthProfileBindingMapId", segmentRealTimeEthProfileBindingMapIdParam)
	sv.AddStructField("SegmentRealTimeEthProfileBindingMap", segmentRealTimeEthProfileBindingMapParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.segments.segment_real_time_eth_profile_binding_maps", "patch", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (sIface *segmentRealTimeEthProfileBindingMapsClient) Update(segmentIdParam string, segmentRealTimeEthProfileBindingMapIdParam string, segmentRealTimeEthProfileBindingMapParam nsx_policyModel.SegmentRealTimeEthProfileBindingMap) (nsx_policyModel.SegmentRealTimeEthProfileBindingMap, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	operationRestMetaData := segmentRealTimeEthProfileBindingMapsUpdateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(segmentRealTimeEthProfileBindingMapsUpdateInputType(), typeConverter)
	sv.AddStructField("SegmentId", segmentIdParam)
	sv.AddStructField("SegmentRealTimeEthProfileBindingMapId", segmentRealTimeEthProfileBindingMapIdParam)
	sv.AddStructField("SegmentRealTimeEthProfileBindingMap", segmentRealTimeEthProfileBindingMapParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.SegmentRealTimeEthProfileBindingMap
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.segments.segment_real_time_eth_profile_binding_maps", "update", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.SegmentRealTimeEthProfileBindingMap
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), SegmentRealTimeEthProfileBindingMapsUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.SegmentRealTimeEthProfileBindingMap), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
