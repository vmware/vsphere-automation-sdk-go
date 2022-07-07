// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: UpgradeUnits
// Used by client-side stubs.

package upgrade

import (
	"github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
)

const _ = core.SupportedByRuntimeVersion1

type UpgradeUnitsClient interface {

	// Get a specific upgrade unit
	//
	// @param upgradeUnitIdParam (required)
	// @return com.vmware.nsx.model.UpgradeUnit
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(upgradeUnitIdParam string) (model.UpgradeUnit, error)

	// Get upgrade units
	//
	// @param componentTypeParam Component type based on which upgrade units to be filtered (optional)
	// @param currentVersionParam Current version of upgrade unit based on which upgrade units to be filtered (optional)
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param groupIdParam UUID of group based on which upgrade units to be filtered (optional)
	// @param hasWarningsParam Flag to indicate whether to return only upgrade units with warnings (optional, default to false)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param metadataParam Metadata about upgrade unit to filter on (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @param upgradeUnitTypeParam Upgrade unit type based on which upgrade units to be filtered (optional)
	// @return com.vmware.nsx.model.UpgradeUnitListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(componentTypeParam *string, currentVersionParam *string, cursorParam *string, groupIdParam *string, hasWarningsParam *bool, includedFieldsParam *string, metadataParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string, upgradeUnitTypeParam *string) (model.UpgradeUnitListResult, error)
}

type upgradeUnitsClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewUpgradeUnitsClient(connector client.Connector) *upgradeUnitsClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.nsx.upgrade.upgrade_units")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"get":  core.NewMethodIdentifier(interfaceIdentifier, "get"),
		"list": core.NewMethodIdentifier(interfaceIdentifier, "list"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	uIface := upgradeUnitsClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &uIface
}

func (uIface *upgradeUnitsClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := uIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (uIface *upgradeUnitsClient) Get(upgradeUnitIdParam string) (model.UpgradeUnit, error) {
	typeConverter := uIface.connector.TypeConverter()
	executionContext := uIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(upgradeUnitsGetInputType(), typeConverter)
	sv.AddStructField("UpgradeUnitId", upgradeUnitIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.UpgradeUnit
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := upgradeUnitsGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	uIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := uIface.connector.GetApiProvider().Invoke("com.vmware.nsx.upgrade.upgrade_units", "get", inputDataValue, executionContext)
	var emptyOutput model.UpgradeUnit
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), upgradeUnitsGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.UpgradeUnit), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), uIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (uIface *upgradeUnitsClient) List(componentTypeParam *string, currentVersionParam *string, cursorParam *string, groupIdParam *string, hasWarningsParam *bool, includedFieldsParam *string, metadataParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string, upgradeUnitTypeParam *string) (model.UpgradeUnitListResult, error) {
	typeConverter := uIface.connector.TypeConverter()
	executionContext := uIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(upgradeUnitsListInputType(), typeConverter)
	sv.AddStructField("ComponentType", componentTypeParam)
	sv.AddStructField("CurrentVersion", currentVersionParam)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("GroupId", groupIdParam)
	sv.AddStructField("HasWarnings", hasWarningsParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("Metadata", metadataParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	sv.AddStructField("UpgradeUnitType", upgradeUnitTypeParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.UpgradeUnitListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := upgradeUnitsListRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	uIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := uIface.connector.GetApiProvider().Invoke("com.vmware.nsx.upgrade.upgrade_units", "list", inputDataValue, executionContext)
	var emptyOutput model.UpgradeUnitListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), upgradeUnitsListOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.UpgradeUnitListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), uIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
