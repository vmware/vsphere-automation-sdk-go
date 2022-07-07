// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: UpgradeUnitGroups
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

type UpgradeUnitGroupsClient interface {

	// Add upgrade units to specified upgrade unit group. The upgrade units will be added at the end of the upgrade unit list.
	//
	// @param groupIdParam (required)
	// @param upgradeUnitListParam (required)
	// @return com.vmware.nsx.model.UpgradeUnitList
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Addupgradeunits(groupIdParam string, upgradeUnitListParam model.UpgradeUnitList) (model.UpgradeUnitList, error)

	// Create a group of upgrade units.
	//
	// @param upgradeUnitGroupParam (required)
	// @return com.vmware.nsx.model.UpgradeUnitGroup
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Create(upgradeUnitGroupParam model.UpgradeUnitGroup) (model.UpgradeUnitGroup, error)

	// Delete the specified group. NOTE - A group can be deleted only if it is empty. If user tries to delete a group containing one or more upgrade units, the operation will fail and an error will be returned.
	//
	// @param groupIdParam (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Delete(groupIdParam string) error

	// Returns information about a specific upgrade unit group in the upgrade plan. If request parameter summary is set to true, then only count of upgrade units will be returned, upgrade units list will be empty.
	//
	// @param groupIdParam (required)
	// @param summaryParam Flag indicating whether to return the summary (optional, default to false)
	// @return com.vmware.nsx.model.UpgradeUnitGroup
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(groupIdParam string, summaryParam *bool) (model.UpgradeUnitGroup, error)

	// Return information of all upgrade unit groups in the upgrade plan. If request parameter summary is set to true, then only count of upgrade units will be returned, upgrade units list will be empty. If request parameter component type is specified, then all upgrade unit groups for that component will be returned.
	//
	// @param componentTypeParam Component type based on which upgrade unit groups to be filtered (optional)
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @param summaryParam Flag indicating whether to return summary (optional, default to false)
	// @param syncParam Synchronize before returning upgrade unit groups (optional, default to false)
	// @return com.vmware.nsx.model.UpgradeUnitGroupListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(componentTypeParam *string, cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string, summaryParam *bool, syncParam *bool) (model.UpgradeUnitGroupListResult, error)

	// Reorder an upgrade unit group by placing it before/after the specified upgrade unit group.
	//
	// @param groupIdParam (required)
	// @param reorderRequestParam (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Reorder(groupIdParam string, reorderRequestParam model.ReorderRequest) error

	// Update the specified upgrade unit group. Removal of upgrade units from the group using this is not allowed. An error will be returned in that case.
	//
	// @param groupIdParam (required)
	// @param upgradeUnitGroupParam (required)
	// @return com.vmware.nsx.model.UpgradeUnitGroup
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(groupIdParam string, upgradeUnitGroupParam model.UpgradeUnitGroup) (model.UpgradeUnitGroup, error)
}

type upgradeUnitGroupsClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewUpgradeUnitGroupsClient(connector client.Connector) *upgradeUnitGroupsClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.nsx.upgrade.upgrade_unit_groups")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"addupgradeunits": core.NewMethodIdentifier(interfaceIdentifier, "addupgradeunits"),
		"create":          core.NewMethodIdentifier(interfaceIdentifier, "create"),
		"delete":          core.NewMethodIdentifier(interfaceIdentifier, "delete"),
		"get":             core.NewMethodIdentifier(interfaceIdentifier, "get"),
		"list":            core.NewMethodIdentifier(interfaceIdentifier, "list"),
		"reorder":         core.NewMethodIdentifier(interfaceIdentifier, "reorder"),
		"update":          core.NewMethodIdentifier(interfaceIdentifier, "update"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	uIface := upgradeUnitGroupsClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &uIface
}

func (uIface *upgradeUnitGroupsClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := uIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (uIface *upgradeUnitGroupsClient) Addupgradeunits(groupIdParam string, upgradeUnitListParam model.UpgradeUnitList) (model.UpgradeUnitList, error) {
	typeConverter := uIface.connector.TypeConverter()
	executionContext := uIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(upgradeUnitGroupsAddupgradeunitsInputType(), typeConverter)
	sv.AddStructField("GroupId", groupIdParam)
	sv.AddStructField("UpgradeUnitList", upgradeUnitListParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.UpgradeUnitList
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := upgradeUnitGroupsAddupgradeunitsRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	uIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := uIface.connector.GetApiProvider().Invoke("com.vmware.nsx.upgrade.upgrade_unit_groups", "addupgradeunits", inputDataValue, executionContext)
	var emptyOutput model.UpgradeUnitList
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), upgradeUnitGroupsAddupgradeunitsOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.UpgradeUnitList), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), uIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (uIface *upgradeUnitGroupsClient) Create(upgradeUnitGroupParam model.UpgradeUnitGroup) (model.UpgradeUnitGroup, error) {
	typeConverter := uIface.connector.TypeConverter()
	executionContext := uIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(upgradeUnitGroupsCreateInputType(), typeConverter)
	sv.AddStructField("UpgradeUnitGroup", upgradeUnitGroupParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.UpgradeUnitGroup
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := upgradeUnitGroupsCreateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	uIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := uIface.connector.GetApiProvider().Invoke("com.vmware.nsx.upgrade.upgrade_unit_groups", "create", inputDataValue, executionContext)
	var emptyOutput model.UpgradeUnitGroup
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), upgradeUnitGroupsCreateOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.UpgradeUnitGroup), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), uIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (uIface *upgradeUnitGroupsClient) Delete(groupIdParam string) error {
	typeConverter := uIface.connector.TypeConverter()
	executionContext := uIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(upgradeUnitGroupsDeleteInputType(), typeConverter)
	sv.AddStructField("GroupId", groupIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := upgradeUnitGroupsDeleteRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	uIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := uIface.connector.GetApiProvider().Invoke("com.vmware.nsx.upgrade.upgrade_unit_groups", "delete", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), uIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (uIface *upgradeUnitGroupsClient) Get(groupIdParam string, summaryParam *bool) (model.UpgradeUnitGroup, error) {
	typeConverter := uIface.connector.TypeConverter()
	executionContext := uIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(upgradeUnitGroupsGetInputType(), typeConverter)
	sv.AddStructField("GroupId", groupIdParam)
	sv.AddStructField("Summary", summaryParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.UpgradeUnitGroup
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := upgradeUnitGroupsGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	uIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := uIface.connector.GetApiProvider().Invoke("com.vmware.nsx.upgrade.upgrade_unit_groups", "get", inputDataValue, executionContext)
	var emptyOutput model.UpgradeUnitGroup
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), upgradeUnitGroupsGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.UpgradeUnitGroup), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), uIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (uIface *upgradeUnitGroupsClient) List(componentTypeParam *string, cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string, summaryParam *bool, syncParam *bool) (model.UpgradeUnitGroupListResult, error) {
	typeConverter := uIface.connector.TypeConverter()
	executionContext := uIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(upgradeUnitGroupsListInputType(), typeConverter)
	sv.AddStructField("ComponentType", componentTypeParam)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	sv.AddStructField("Summary", summaryParam)
	sv.AddStructField("Sync", syncParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.UpgradeUnitGroupListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := upgradeUnitGroupsListRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	uIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := uIface.connector.GetApiProvider().Invoke("com.vmware.nsx.upgrade.upgrade_unit_groups", "list", inputDataValue, executionContext)
	var emptyOutput model.UpgradeUnitGroupListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), upgradeUnitGroupsListOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.UpgradeUnitGroupListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), uIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (uIface *upgradeUnitGroupsClient) Reorder(groupIdParam string, reorderRequestParam model.ReorderRequest) error {
	typeConverter := uIface.connector.TypeConverter()
	executionContext := uIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(upgradeUnitGroupsReorderInputType(), typeConverter)
	sv.AddStructField("GroupId", groupIdParam)
	sv.AddStructField("ReorderRequest", reorderRequestParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := upgradeUnitGroupsReorderRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	uIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := uIface.connector.GetApiProvider().Invoke("com.vmware.nsx.upgrade.upgrade_unit_groups", "reorder", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), uIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (uIface *upgradeUnitGroupsClient) Update(groupIdParam string, upgradeUnitGroupParam model.UpgradeUnitGroup) (model.UpgradeUnitGroup, error) {
	typeConverter := uIface.connector.TypeConverter()
	executionContext := uIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(upgradeUnitGroupsUpdateInputType(), typeConverter)
	sv.AddStructField("GroupId", groupIdParam)
	sv.AddStructField("UpgradeUnitGroup", upgradeUnitGroupParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.UpgradeUnitGroup
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := upgradeUnitGroupsUpdateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	uIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := uIface.connector.GetApiProvider().Invoke("com.vmware.nsx.upgrade.upgrade_unit_groups", "update", inputDataValue, executionContext)
	var emptyOutput model.UpgradeUnitGroup
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), upgradeUnitGroupsUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.UpgradeUnitGroup), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), uIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
