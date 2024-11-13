// Copyright (c) 2019-2024 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: PortRealTimeEthProfileBindingMaps
// Used by client-side stubs.

package ports

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	nsx_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type PortRealTimeEthProfileBindingMapsClient interface {

	// API will delete Port Real Time Eth Profile Binding Profile.
	//
	// @param segmentIdParam Segment ID (required)
	// @param portIdParam Port ID (required)
	// @param portRealTimeEthProfileBindingMapIdParam Port Real Time Eth Profile Binding Map ID (required)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Delete(segmentIdParam string, portIdParam string, portRealTimeEthProfileBindingMapIdParam string) error

	// API will get Port Real Time Eth Profile Binding Map.
	//
	// @param segmentIdParam Segment ID (required)
	// @param portIdParam Port ID (required)
	// @param portRealTimeEthProfileBindingMapIdParam Port Real Time Eth Profile Binding Map ID (required)
	// @return com.vmware.nsx_policy.model.PortRealTimeEthProfileBindingMap
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(segmentIdParam string, portIdParam string, portRealTimeEthProfileBindingMapIdParam string) (nsx_policyModel.PortRealTimeEthProfileBindingMap, error)

	// API will list all Port Real Time Eth Profile Binding Maps in current port id.
	//
	// @param segmentIdParam (required)
	// @param portIdParam (required)
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @return com.vmware.nsx_policy.model.PortRealTimeEthProfileBindingMapListResult
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(segmentIdParam string, portIdParam string, cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (nsx_policyModel.PortRealTimeEthProfileBindingMapListResult, error)

	// API will create Port Real Time Eth Profile Binding Map. For objects with no binding maps, default profile is applied.
	//
	// @param segmentIdParam Segment ID (required)
	// @param portIdParam Port ID (required)
	// @param portRealTimeEthProfileBindingMapIdParam Port Real Time Eth Profile Binding Map ID (required)
	// @param portRealTimeEthProfileBindingMapParam (required)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Patch(segmentIdParam string, portIdParam string, portRealTimeEthProfileBindingMapIdParam string, portRealTimeEthProfileBindingMapParam nsx_policyModel.PortRealTimeEthProfileBindingMap) error

	// API will update Port Real Time Eth Profile Binding Map. For objects with no binding maps, default profile is applied.
	//
	// @param segmentIdParam Segment ID (required)
	// @param portIdParam Port ID (required)
	// @param portRealTimeEthProfileBindingMapIdParam Port Real Time Eth Profile Binding Map ID (required)
	// @param portRealTimeEthProfileBindingMapParam (required)
	// @return com.vmware.nsx_policy.model.PortRealTimeEthProfileBindingMap
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(segmentIdParam string, portIdParam string, portRealTimeEthProfileBindingMapIdParam string, portRealTimeEthProfileBindingMapParam nsx_policyModel.PortRealTimeEthProfileBindingMap) (nsx_policyModel.PortRealTimeEthProfileBindingMap, error)
}

type portRealTimeEthProfileBindingMapsClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewPortRealTimeEthProfileBindingMapsClient(connector vapiProtocolClient_.Connector) *portRealTimeEthProfileBindingMapsClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx_policy.infra.segments.ports.port_real_time_eth_profile_binding_maps")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"delete": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "delete"),
		"get":    vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
		"list":   vapiCore_.NewMethodIdentifier(interfaceIdentifier, "list"),
		"patch":  vapiCore_.NewMethodIdentifier(interfaceIdentifier, "patch"),
		"update": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "update"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	pIface := portRealTimeEthProfileBindingMapsClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &pIface
}

func (pIface *portRealTimeEthProfileBindingMapsClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := pIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (pIface *portRealTimeEthProfileBindingMapsClient) Delete(segmentIdParam string, portIdParam string, portRealTimeEthProfileBindingMapIdParam string) error {
	typeConverter := pIface.connector.TypeConverter()
	executionContext := pIface.connector.NewExecutionContext()
	operationRestMetaData := portRealTimeEthProfileBindingMapsDeleteRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(portRealTimeEthProfileBindingMapsDeleteInputType(), typeConverter)
	sv.AddStructField("SegmentId", segmentIdParam)
	sv.AddStructField("PortId", portIdParam)
	sv.AddStructField("PortRealTimeEthProfileBindingMapId", portRealTimeEthProfileBindingMapIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := pIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.segments.ports.port_real_time_eth_profile_binding_maps", "delete", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), pIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (pIface *portRealTimeEthProfileBindingMapsClient) Get(segmentIdParam string, portIdParam string, portRealTimeEthProfileBindingMapIdParam string) (nsx_policyModel.PortRealTimeEthProfileBindingMap, error) {
	typeConverter := pIface.connector.TypeConverter()
	executionContext := pIface.connector.NewExecutionContext()
	operationRestMetaData := portRealTimeEthProfileBindingMapsGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(portRealTimeEthProfileBindingMapsGetInputType(), typeConverter)
	sv.AddStructField("SegmentId", segmentIdParam)
	sv.AddStructField("PortId", portIdParam)
	sv.AddStructField("PortRealTimeEthProfileBindingMapId", portRealTimeEthProfileBindingMapIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.PortRealTimeEthProfileBindingMap
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := pIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.segments.ports.port_real_time_eth_profile_binding_maps", "get", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.PortRealTimeEthProfileBindingMap
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), PortRealTimeEthProfileBindingMapsGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.PortRealTimeEthProfileBindingMap), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), pIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (pIface *portRealTimeEthProfileBindingMapsClient) List(segmentIdParam string, portIdParam string, cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (nsx_policyModel.PortRealTimeEthProfileBindingMapListResult, error) {
	typeConverter := pIface.connector.TypeConverter()
	executionContext := pIface.connector.NewExecutionContext()
	operationRestMetaData := portRealTimeEthProfileBindingMapsListRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(portRealTimeEthProfileBindingMapsListInputType(), typeConverter)
	sv.AddStructField("SegmentId", segmentIdParam)
	sv.AddStructField("PortId", portIdParam)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.PortRealTimeEthProfileBindingMapListResult
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := pIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.segments.ports.port_real_time_eth_profile_binding_maps", "list", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.PortRealTimeEthProfileBindingMapListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), PortRealTimeEthProfileBindingMapsListOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.PortRealTimeEthProfileBindingMapListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), pIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (pIface *portRealTimeEthProfileBindingMapsClient) Patch(segmentIdParam string, portIdParam string, portRealTimeEthProfileBindingMapIdParam string, portRealTimeEthProfileBindingMapParam nsx_policyModel.PortRealTimeEthProfileBindingMap) error {
	typeConverter := pIface.connector.TypeConverter()
	executionContext := pIface.connector.NewExecutionContext()
	operationRestMetaData := portRealTimeEthProfileBindingMapsPatchRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(portRealTimeEthProfileBindingMapsPatchInputType(), typeConverter)
	sv.AddStructField("SegmentId", segmentIdParam)
	sv.AddStructField("PortId", portIdParam)
	sv.AddStructField("PortRealTimeEthProfileBindingMapId", portRealTimeEthProfileBindingMapIdParam)
	sv.AddStructField("PortRealTimeEthProfileBindingMap", portRealTimeEthProfileBindingMapParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := pIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.segments.ports.port_real_time_eth_profile_binding_maps", "patch", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), pIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (pIface *portRealTimeEthProfileBindingMapsClient) Update(segmentIdParam string, portIdParam string, portRealTimeEthProfileBindingMapIdParam string, portRealTimeEthProfileBindingMapParam nsx_policyModel.PortRealTimeEthProfileBindingMap) (nsx_policyModel.PortRealTimeEthProfileBindingMap, error) {
	typeConverter := pIface.connector.TypeConverter()
	executionContext := pIface.connector.NewExecutionContext()
	operationRestMetaData := portRealTimeEthProfileBindingMapsUpdateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(portRealTimeEthProfileBindingMapsUpdateInputType(), typeConverter)
	sv.AddStructField("SegmentId", segmentIdParam)
	sv.AddStructField("PortId", portIdParam)
	sv.AddStructField("PortRealTimeEthProfileBindingMapId", portRealTimeEthProfileBindingMapIdParam)
	sv.AddStructField("PortRealTimeEthProfileBindingMap", portRealTimeEthProfileBindingMapParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.PortRealTimeEthProfileBindingMap
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := pIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.segments.ports.port_real_time_eth_profile_binding_maps", "update", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.PortRealTimeEthProfileBindingMap
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), PortRealTimeEthProfileBindingMapsUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.PortRealTimeEthProfileBindingMap), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), pIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
