// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: PortSecurityProfileBindingMaps
// Used by client-side stubs.

package ports

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	nsx_global_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-gm/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type PortSecurityProfileBindingMapsClient interface {

	// API will return details of the port security profile binding map. If the security profile binding map does not exist, it will return 404.
	//
	// @param tier1IdParam tier-1 gateway id (required)
	// @param segmentIdParam segment id (required)
	// @param portIdParam port id (required)
	// @param portSecurityProfileBindingMapIdParam port security profile binding map id (required)
	// @return com.vmware.nsx_global_policy.model.PortSecurityProfileBindingMap
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(tier1IdParam string, segmentIdParam string, portIdParam string, portSecurityProfileBindingMapIdParam string) (nsx_global_policyModel.PortSecurityProfileBindingMap, error)

	// API will list all port security profile binding maps.
	//
	// @param tier1IdParam tier-1 gateway id (required)
	// @param segmentIdParam segment id (required)
	// @param portIdParam port id (required)
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @return com.vmware.nsx_global_policy.model.PortSecurityProfileBindingMapListResult
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(tier1IdParam string, segmentIdParam string, portIdParam string, cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (nsx_global_policyModel.PortSecurityProfileBindingMapListResult, error)
}

type portSecurityProfileBindingMapsClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewPortSecurityProfileBindingMapsClient(connector vapiProtocolClient_.Connector) *portSecurityProfileBindingMapsClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx_global_policy.global_infra.tier_1s.segments.ports.port_security_profile_binding_maps")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"get":  vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
		"list": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "list"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	pIface := portSecurityProfileBindingMapsClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &pIface
}

func (pIface *portSecurityProfileBindingMapsClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := pIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (pIface *portSecurityProfileBindingMapsClient) Get(tier1IdParam string, segmentIdParam string, portIdParam string, portSecurityProfileBindingMapIdParam string) (nsx_global_policyModel.PortSecurityProfileBindingMap, error) {
	typeConverter := pIface.connector.TypeConverter()
	executionContext := pIface.connector.NewExecutionContext()
	operationRestMetaData := portSecurityProfileBindingMapsGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(portSecurityProfileBindingMapsGetInputType(), typeConverter)
	sv.AddStructField("Tier1Id", tier1IdParam)
	sv.AddStructField("SegmentId", segmentIdParam)
	sv.AddStructField("PortId", portIdParam)
	sv.AddStructField("PortSecurityProfileBindingMapId", portSecurityProfileBindingMapIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_global_policyModel.PortSecurityProfileBindingMap
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := pIface.connector.GetApiProvider().Invoke("com.vmware.nsx_global_policy.global_infra.tier_1s.segments.ports.port_security_profile_binding_maps", "get", inputDataValue, executionContext)
	var emptyOutput nsx_global_policyModel.PortSecurityProfileBindingMap
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), PortSecurityProfileBindingMapsGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_global_policyModel.PortSecurityProfileBindingMap), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), pIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (pIface *portSecurityProfileBindingMapsClient) List(tier1IdParam string, segmentIdParam string, portIdParam string, cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (nsx_global_policyModel.PortSecurityProfileBindingMapListResult, error) {
	typeConverter := pIface.connector.TypeConverter()
	executionContext := pIface.connector.NewExecutionContext()
	operationRestMetaData := portSecurityProfileBindingMapsListRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(portSecurityProfileBindingMapsListInputType(), typeConverter)
	sv.AddStructField("Tier1Id", tier1IdParam)
	sv.AddStructField("SegmentId", segmentIdParam)
	sv.AddStructField("PortId", portIdParam)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_global_policyModel.PortSecurityProfileBindingMapListResult
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := pIface.connector.GetApiProvider().Invoke("com.vmware.nsx_global_policy.global_infra.tier_1s.segments.ports.port_security_profile_binding_maps", "list", inputDataValue, executionContext)
	var emptyOutput nsx_global_policyModel.PortSecurityProfileBindingMapListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), PortSecurityProfileBindingMapsListOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_global_policyModel.PortSecurityProfileBindingMapListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), pIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
