// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: PortMonitoringProfileBindingMaps
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

type PortMonitoringProfileBindingMapsClient interface {

	// API will delete Infra Port Monitoring Profile Binding Profile.
	//
	// @param infraSegmentIdParam Infra Segment ID (required)
	// @param infraPortIdParam Infra Port ID (required)
	// @param portMonitoringProfileBindingMapIdParam Port Monitoring Profile Binding Map ID (required)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Delete(infraSegmentIdParam string, infraPortIdParam string, portMonitoringProfileBindingMapIdParam string) error

	// API will get Infra Port Monitoring Profile Binding Map.
	//
	// @param infraSegmentIdParam Infra Segment ID (required)
	// @param infraPortIdParam Infra Port ID (required)
	// @param portMonitoringProfileBindingMapIdParam Port Monitoring Profile Binding Map ID (required)
	// @return com.vmware.nsx_policy.model.PortMonitoringProfileBindingMap
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(infraSegmentIdParam string, infraPortIdParam string, portMonitoringProfileBindingMapIdParam string) (nsx_policyModel.PortMonitoringProfileBindingMap, error)

	// API will list all Infra Port Monitoring Profile Binding Maps in current port id.
	//
	// @param infraSegmentIdParam (required)
	// @param infraPortIdParam (required)
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includeMarkForDeleteObjectsParam Include objects that are marked for deletion in results (optional, default to false)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @return com.vmware.nsx_policy.model.PortMonitoringProfileBindingMapListResult
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(infraSegmentIdParam string, infraPortIdParam string, cursorParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (nsx_policyModel.PortMonitoringProfileBindingMapListResult, error)

	// API will create Infra Port Monitoring Profile Binding Map.
	//
	// @param infraSegmentIdParam Infra Segment ID (required)
	// @param infraPortIdParam Infra Port ID (required)
	// @param portMonitoringProfileBindingMapIdParam Port Monitoring Profile Binding Map ID (required)
	// @param portMonitoringProfileBindingMapParam (required)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Patch(infraSegmentIdParam string, infraPortIdParam string, portMonitoringProfileBindingMapIdParam string, portMonitoringProfileBindingMapParam nsx_policyModel.PortMonitoringProfileBindingMap) error

	// API will update Infra Port Monitoring Profile Binding Map.
	//
	// @param infraSegmentIdParam InfraSegment ID (required)
	// @param infraPortIdParam Infra Port ID (required)
	// @param portMonitoringProfileBindingMapIdParam Port Monitoring Profile Binding Map ID (required)
	// @param portMonitoringProfileBindingMapParam (required)
	// @return com.vmware.nsx_policy.model.PortMonitoringProfileBindingMap
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(infraSegmentIdParam string, infraPortIdParam string, portMonitoringProfileBindingMapIdParam string, portMonitoringProfileBindingMapParam nsx_policyModel.PortMonitoringProfileBindingMap) (nsx_policyModel.PortMonitoringProfileBindingMap, error)
}

type portMonitoringProfileBindingMapsClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewPortMonitoringProfileBindingMapsClient(connector vapiProtocolClient_.Connector) *portMonitoringProfileBindingMapsClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx_policy.infra.segments.ports.port_monitoring_profile_binding_maps")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"delete": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "delete"),
		"get":    vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
		"list":   vapiCore_.NewMethodIdentifier(interfaceIdentifier, "list"),
		"patch":  vapiCore_.NewMethodIdentifier(interfaceIdentifier, "patch"),
		"update": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "update"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	pIface := portMonitoringProfileBindingMapsClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &pIface
}

func (pIface *portMonitoringProfileBindingMapsClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := pIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (pIface *portMonitoringProfileBindingMapsClient) Delete(infraSegmentIdParam string, infraPortIdParam string, portMonitoringProfileBindingMapIdParam string) error {
	typeConverter := pIface.connector.TypeConverter()
	executionContext := pIface.connector.NewExecutionContext()
	operationRestMetaData := portMonitoringProfileBindingMapsDeleteRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(portMonitoringProfileBindingMapsDeleteInputType(), typeConverter)
	sv.AddStructField("InfraSegmentId", infraSegmentIdParam)
	sv.AddStructField("InfraPortId", infraPortIdParam)
	sv.AddStructField("PortMonitoringProfileBindingMapId", portMonitoringProfileBindingMapIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := pIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.segments.ports.port_monitoring_profile_binding_maps", "delete", inputDataValue, executionContext)
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

func (pIface *portMonitoringProfileBindingMapsClient) Get(infraSegmentIdParam string, infraPortIdParam string, portMonitoringProfileBindingMapIdParam string) (nsx_policyModel.PortMonitoringProfileBindingMap, error) {
	typeConverter := pIface.connector.TypeConverter()
	executionContext := pIface.connector.NewExecutionContext()
	operationRestMetaData := portMonitoringProfileBindingMapsGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(portMonitoringProfileBindingMapsGetInputType(), typeConverter)
	sv.AddStructField("InfraSegmentId", infraSegmentIdParam)
	sv.AddStructField("InfraPortId", infraPortIdParam)
	sv.AddStructField("PortMonitoringProfileBindingMapId", portMonitoringProfileBindingMapIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.PortMonitoringProfileBindingMap
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := pIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.segments.ports.port_monitoring_profile_binding_maps", "get", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.PortMonitoringProfileBindingMap
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), PortMonitoringProfileBindingMapsGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.PortMonitoringProfileBindingMap), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), pIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (pIface *portMonitoringProfileBindingMapsClient) List(infraSegmentIdParam string, infraPortIdParam string, cursorParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (nsx_policyModel.PortMonitoringProfileBindingMapListResult, error) {
	typeConverter := pIface.connector.TypeConverter()
	executionContext := pIface.connector.NewExecutionContext()
	operationRestMetaData := portMonitoringProfileBindingMapsListRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(portMonitoringProfileBindingMapsListInputType(), typeConverter)
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
		var emptyOutput nsx_policyModel.PortMonitoringProfileBindingMapListResult
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := pIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.segments.ports.port_monitoring_profile_binding_maps", "list", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.PortMonitoringProfileBindingMapListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), PortMonitoringProfileBindingMapsListOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.PortMonitoringProfileBindingMapListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), pIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (pIface *portMonitoringProfileBindingMapsClient) Patch(infraSegmentIdParam string, infraPortIdParam string, portMonitoringProfileBindingMapIdParam string, portMonitoringProfileBindingMapParam nsx_policyModel.PortMonitoringProfileBindingMap) error {
	typeConverter := pIface.connector.TypeConverter()
	executionContext := pIface.connector.NewExecutionContext()
	operationRestMetaData := portMonitoringProfileBindingMapsPatchRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(portMonitoringProfileBindingMapsPatchInputType(), typeConverter)
	sv.AddStructField("InfraSegmentId", infraSegmentIdParam)
	sv.AddStructField("InfraPortId", infraPortIdParam)
	sv.AddStructField("PortMonitoringProfileBindingMapId", portMonitoringProfileBindingMapIdParam)
	sv.AddStructField("PortMonitoringProfileBindingMap", portMonitoringProfileBindingMapParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := pIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.segments.ports.port_monitoring_profile_binding_maps", "patch", inputDataValue, executionContext)
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

func (pIface *portMonitoringProfileBindingMapsClient) Update(infraSegmentIdParam string, infraPortIdParam string, portMonitoringProfileBindingMapIdParam string, portMonitoringProfileBindingMapParam nsx_policyModel.PortMonitoringProfileBindingMap) (nsx_policyModel.PortMonitoringProfileBindingMap, error) {
	typeConverter := pIface.connector.TypeConverter()
	executionContext := pIface.connector.NewExecutionContext()
	operationRestMetaData := portMonitoringProfileBindingMapsUpdateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(portMonitoringProfileBindingMapsUpdateInputType(), typeConverter)
	sv.AddStructField("InfraSegmentId", infraSegmentIdParam)
	sv.AddStructField("InfraPortId", infraPortIdParam)
	sv.AddStructField("PortMonitoringProfileBindingMapId", portMonitoringProfileBindingMapIdParam)
	sv.AddStructField("PortMonitoringProfileBindingMap", portMonitoringProfileBindingMapParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.PortMonitoringProfileBindingMap
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := pIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.segments.ports.port_monitoring_profile_binding_maps", "update", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.PortMonitoringProfileBindingMap
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), PortMonitoringProfileBindingMapsUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.PortMonitoringProfileBindingMap), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), pIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
