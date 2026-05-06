// Copyright (c) 2019-2026 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Zones
// Used by client-side stubs.

package fabric

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	nsxModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type ZonesClient interface {

	// Returns information about a specific zone, along with the list of its associated clusters discovered from the compute manager
	//
	// @param zoneExtIdParam (required)
	// @return com.vmware.nsx.model.Zone
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(zoneExtIdParam string) (nsxModel.Zone, error)

	// Returns information about all zones discovered from the compute manager
	//
	// @param cmLocalIdParam The zone id/ name as is received from the compute manager, e.g. \"california-zone\" (optional)
	// @param computeCollectionExternalIdParam External id of the compute collection (cluster), that is associated with a zone (optional)
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includedFieldsParam Note - this parameter currently only works when used with the search APIs /policy/api/v1/search/query and /policy/api/v1/search/dsl. It is ignored for other list APIs. (optional)
	// @param originIdParam Id of the compute manager from where the zone was discovered, e.g. \"ecd5d517-db8d-46cf-bf52-f7e3663b7232\" (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam If true, results are sorted in ascending order (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @return com.vmware.nsx.model.ZoneListResult
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(cmLocalIdParam *string, computeCollectionExternalIdParam *string, cursorParam *string, includedFieldsParam *string, originIdParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (nsxModel.ZoneListResult, error)
}

type zonesClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewZonesClient(connector vapiProtocolClient_.Connector) *zonesClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx.fabric.zones")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"get":  vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
		"list": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "list"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	zIface := zonesClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &zIface
}

func (zIface *zonesClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := zIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (zIface *zonesClient) Get(zoneExtIdParam string) (nsxModel.Zone, error) {
	typeConverter := zIface.connector.TypeConverter()
	executionContext := zIface.connector.NewExecutionContext()
	operationRestMetaData := zonesGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(zonesGetInputType(), typeConverter)
	sv.AddStructField("ZoneExtId", zoneExtIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.Zone
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := zIface.connector.GetApiProvider().Invoke("com.vmware.nsx.fabric.zones", "get", inputDataValue, executionContext)
	var emptyOutput nsxModel.Zone
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), ZonesGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.Zone), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), zIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (zIface *zonesClient) List(cmLocalIdParam *string, computeCollectionExternalIdParam *string, cursorParam *string, includedFieldsParam *string, originIdParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (nsxModel.ZoneListResult, error) {
	typeConverter := zIface.connector.TypeConverter()
	executionContext := zIface.connector.NewExecutionContext()
	operationRestMetaData := zonesListRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(zonesListInputType(), typeConverter)
	sv.AddStructField("CmLocalId", cmLocalIdParam)
	sv.AddStructField("ComputeCollectionExternalId", computeCollectionExternalIdParam)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("OriginId", originIdParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.ZoneListResult
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := zIface.connector.GetApiProvider().Invoke("com.vmware.nsx.fabric.zones", "list", inputDataValue, executionContext)
	var emptyOutput nsxModel.ZoneListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), ZonesListOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.ZoneListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), zIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
