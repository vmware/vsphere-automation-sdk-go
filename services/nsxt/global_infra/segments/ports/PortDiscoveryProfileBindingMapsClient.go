// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: PortDiscoveryProfileBindingMaps
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

type PortDiscoveryProfileBindingMapsClient interface {

	// API will delete Infra Port Discovery Profile Binding Profile
	//
	// @param infraSegmentIdParam Infra Segment ID (required)
	// @param infraPortIdParam Infra Port ID (required)
	// @param portDiscoveryProfileBindingMapIdParam Port Discovery Profile Binding Map ID (required)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Delete(infraSegmentIdParam string, infraPortIdParam string, portDiscoveryProfileBindingMapIdParam string) error

	// API will get Infra Port Discovery Profile Binding Map
	//
	// @param infraSegmentIdParam Infra Segment ID (required)
	// @param infraPortIdParam Infra Port ID (required)
	// @param portDiscoveryProfileBindingMapIdParam Port Discovery Profile Binding Map ID (required)
	// @return com.vmware.nsx_policy.model.PortDiscoveryProfileBindingMap
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(infraSegmentIdParam string, infraPortIdParam string, portDiscoveryProfileBindingMapIdParam string) (nsx_policyModel.PortDiscoveryProfileBindingMap, error)

	// API will list all Infra Port Discovery Profile Binding Maps in current port id.
	//
	// @param infraSegmentIdParam (required)
	// @param infraPortIdParam (required)
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includeMarkForDeleteObjectsParam Include objects that are marked for deletion in results (optional, default to false)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @return com.vmware.nsx_policy.model.PortDiscoveryProfileBindingMapListResult
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(infraSegmentIdParam string, infraPortIdParam string, cursorParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (nsx_policyModel.PortDiscoveryProfileBindingMapListResult, error)

	// API will create Infra Port Discovery Profile Binding Map. For objects with no binding maps, default profile is applied.
	//
	// @param infraSegmentIdParam Infra Segment ID (required)
	// @param infraPortIdParam Port ID (required)
	// @param portDiscoveryProfileBindingMapIdParam Port Discovery Profile Binding Map ID (required)
	// @param portDiscoveryProfileBindingMapParam (required)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Patch(infraSegmentIdParam string, infraPortIdParam string, portDiscoveryProfileBindingMapIdParam string, portDiscoveryProfileBindingMapParam nsx_policyModel.PortDiscoveryProfileBindingMap) error

	// API will update Infra Port Discovery Profile Binding Map. For objects with no binding maps, default profile is applied.
	//
	// @param infraSegmentIdParam Infra Segment ID (required)
	// @param infraPortIdParam Infra Port ID (required)
	// @param portDiscoveryProfileBindingMapIdParam Port Discovery Profile Binding Map ID (required)
	// @param portDiscoveryProfileBindingMapParam (required)
	// @return com.vmware.nsx_policy.model.PortDiscoveryProfileBindingMap
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(infraSegmentIdParam string, infraPortIdParam string, portDiscoveryProfileBindingMapIdParam string, portDiscoveryProfileBindingMapParam nsx_policyModel.PortDiscoveryProfileBindingMap) (nsx_policyModel.PortDiscoveryProfileBindingMap, error)
}

type portDiscoveryProfileBindingMapsClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewPortDiscoveryProfileBindingMapsClient(connector vapiProtocolClient_.Connector) *portDiscoveryProfileBindingMapsClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx_policy.global_infra.segments.ports.port_discovery_profile_binding_maps")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"delete": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "delete"),
		"get":    vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
		"list":   vapiCore_.NewMethodIdentifier(interfaceIdentifier, "list"),
		"patch":  vapiCore_.NewMethodIdentifier(interfaceIdentifier, "patch"),
		"update": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "update"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	pIface := portDiscoveryProfileBindingMapsClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &pIface
}

func (pIface *portDiscoveryProfileBindingMapsClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := pIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (pIface *portDiscoveryProfileBindingMapsClient) Delete(infraSegmentIdParam string, infraPortIdParam string, portDiscoveryProfileBindingMapIdParam string) error {
	typeConverter := pIface.connector.TypeConverter()
	executionContext := pIface.connector.NewExecutionContext()
	operationRestMetaData := portDiscoveryProfileBindingMapsDeleteRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(portDiscoveryProfileBindingMapsDeleteInputType(), typeConverter)
	sv.AddStructField("InfraSegmentId", infraSegmentIdParam)
	sv.AddStructField("InfraPortId", infraPortIdParam)
	sv.AddStructField("PortDiscoveryProfileBindingMapId", portDiscoveryProfileBindingMapIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := pIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.global_infra.segments.ports.port_discovery_profile_binding_maps", "delete", inputDataValue, executionContext)
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

func (pIface *portDiscoveryProfileBindingMapsClient) Get(infraSegmentIdParam string, infraPortIdParam string, portDiscoveryProfileBindingMapIdParam string) (nsx_policyModel.PortDiscoveryProfileBindingMap, error) {
	typeConverter := pIface.connector.TypeConverter()
	executionContext := pIface.connector.NewExecutionContext()
	operationRestMetaData := portDiscoveryProfileBindingMapsGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(portDiscoveryProfileBindingMapsGetInputType(), typeConverter)
	sv.AddStructField("InfraSegmentId", infraSegmentIdParam)
	sv.AddStructField("InfraPortId", infraPortIdParam)
	sv.AddStructField("PortDiscoveryProfileBindingMapId", portDiscoveryProfileBindingMapIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.PortDiscoveryProfileBindingMap
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := pIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.global_infra.segments.ports.port_discovery_profile_binding_maps", "get", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.PortDiscoveryProfileBindingMap
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), PortDiscoveryProfileBindingMapsGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.PortDiscoveryProfileBindingMap), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), pIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (pIface *portDiscoveryProfileBindingMapsClient) List(infraSegmentIdParam string, infraPortIdParam string, cursorParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (nsx_policyModel.PortDiscoveryProfileBindingMapListResult, error) {
	typeConverter := pIface.connector.TypeConverter()
	executionContext := pIface.connector.NewExecutionContext()
	operationRestMetaData := portDiscoveryProfileBindingMapsListRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(portDiscoveryProfileBindingMapsListInputType(), typeConverter)
	sv.AddStructField("InfraSegmentId", infraSegmentIdParam)
	sv.AddStructField("InfraPortId", infraPortIdParam)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludeMarkForDeleteObjects", includeMarkForDeleteObjectsParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.PortDiscoveryProfileBindingMapListResult
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := pIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.global_infra.segments.ports.port_discovery_profile_binding_maps", "list", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.PortDiscoveryProfileBindingMapListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), PortDiscoveryProfileBindingMapsListOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.PortDiscoveryProfileBindingMapListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), pIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (pIface *portDiscoveryProfileBindingMapsClient) Patch(infraSegmentIdParam string, infraPortIdParam string, portDiscoveryProfileBindingMapIdParam string, portDiscoveryProfileBindingMapParam nsx_policyModel.PortDiscoveryProfileBindingMap) error {
	typeConverter := pIface.connector.TypeConverter()
	executionContext := pIface.connector.NewExecutionContext()
	operationRestMetaData := portDiscoveryProfileBindingMapsPatchRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(portDiscoveryProfileBindingMapsPatchInputType(), typeConverter)
	sv.AddStructField("InfraSegmentId", infraSegmentIdParam)
	sv.AddStructField("InfraPortId", infraPortIdParam)
	sv.AddStructField("PortDiscoveryProfileBindingMapId", portDiscoveryProfileBindingMapIdParam)
	sv.AddStructField("PortDiscoveryProfileBindingMap", portDiscoveryProfileBindingMapParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := pIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.global_infra.segments.ports.port_discovery_profile_binding_maps", "patch", inputDataValue, executionContext)
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

func (pIface *portDiscoveryProfileBindingMapsClient) Update(infraSegmentIdParam string, infraPortIdParam string, portDiscoveryProfileBindingMapIdParam string, portDiscoveryProfileBindingMapParam nsx_policyModel.PortDiscoveryProfileBindingMap) (nsx_policyModel.PortDiscoveryProfileBindingMap, error) {
	typeConverter := pIface.connector.TypeConverter()
	executionContext := pIface.connector.NewExecutionContext()
	operationRestMetaData := portDiscoveryProfileBindingMapsUpdateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(portDiscoveryProfileBindingMapsUpdateInputType(), typeConverter)
	sv.AddStructField("InfraSegmentId", infraSegmentIdParam)
	sv.AddStructField("InfraPortId", infraPortIdParam)
	sv.AddStructField("PortDiscoveryProfileBindingMapId", portDiscoveryProfileBindingMapIdParam)
	sv.AddStructField("PortDiscoveryProfileBindingMap", portDiscoveryProfileBindingMapParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.PortDiscoveryProfileBindingMap
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := pIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.global_infra.segments.ports.port_discovery_profile_binding_maps", "update", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.PortDiscoveryProfileBindingMap
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), PortDiscoveryProfileBindingMapsUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.PortDiscoveryProfileBindingMap), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), pIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
