// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: UpgradeUnitGroups
// Used by client-side stubs.

package upgrade

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	nsxModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type UpgradeUnitGroupsClient interface {

	// Add upgrade units to specified upgrade unit group. The upgrade units will be added at the end of the upgrade unit list.
	//
	// @param groupIdParam (required)
	// @param upgradeUnitListParam (required)
	// @return com.vmware.nsx.model.UpgradeUnitList
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Addupgradeunits(groupIdParam string, upgradeUnitListParam nsxModel.UpgradeUnitList) (nsxModel.UpgradeUnitList, error)

	// Create a group of upgrade units.
	//
	// @param upgradeUnitGroupParam (required)
	// @return com.vmware.nsx.model.UpgradeUnitGroup
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Create(upgradeUnitGroupParam nsxModel.UpgradeUnitGroup) (nsxModel.UpgradeUnitGroup, error)

	// Delete the specified group. NOTE - A group can be deleted only if it is empty. If user tries to delete a group containing one or more upgrade units, the operation will fail and an error will be returned.
	//
	// @param groupIdParam (required)
	//
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
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(groupIdParam string, summaryParam *bool) (nsxModel.UpgradeUnitGroup, error)

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
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(componentTypeParam *string, cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string, summaryParam *bool, syncParam *bool) (nsxModel.UpgradeUnitGroupListResult, error)

	// Reorder an upgrade unit group by placing it before/after the specified upgrade unit group.
	//
	// @param groupIdParam (required)
	// @param reorderRequestParam (required)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Reorder(groupIdParam string, reorderRequestParam nsxModel.ReorderRequest) error

	// Update the specified upgrade unit group. Removal of upgrade units from the group using this is not allowed. An error will be returned in that case.
	//
	// @param groupIdParam (required)
	// @param upgradeUnitGroupParam (required)
	// @return com.vmware.nsx.model.UpgradeUnitGroup
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(groupIdParam string, upgradeUnitGroupParam nsxModel.UpgradeUnitGroup) (nsxModel.UpgradeUnitGroup, error)
}

type upgradeUnitGroupsClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewUpgradeUnitGroupsClient(connector vapiProtocolClient_.Connector) *upgradeUnitGroupsClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx.upgrade.upgrade_unit_groups")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"addupgradeunits": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "addupgradeunits"),
		"create":          vapiCore_.NewMethodIdentifier(interfaceIdentifier, "create"),
		"delete":          vapiCore_.NewMethodIdentifier(interfaceIdentifier, "delete"),
		"get":             vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
		"list":            vapiCore_.NewMethodIdentifier(interfaceIdentifier, "list"),
		"reorder":         vapiCore_.NewMethodIdentifier(interfaceIdentifier, "reorder"),
		"update":          vapiCore_.NewMethodIdentifier(interfaceIdentifier, "update"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	uIface := upgradeUnitGroupsClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &uIface
}

func (uIface *upgradeUnitGroupsClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := uIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (uIface *upgradeUnitGroupsClient) Addupgradeunits(groupIdParam string, upgradeUnitListParam nsxModel.UpgradeUnitList) (nsxModel.UpgradeUnitList, error) {
	typeConverter := uIface.connector.TypeConverter()
	executionContext := uIface.connector.NewExecutionContext()
	operationRestMetaData := upgradeUnitGroupsAddupgradeunitsRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(upgradeUnitGroupsAddupgradeunitsInputType(), typeConverter)
	sv.AddStructField("GroupId", groupIdParam)
	sv.AddStructField("UpgradeUnitList", upgradeUnitListParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.UpgradeUnitList
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := uIface.connector.GetApiProvider().Invoke("com.vmware.nsx.upgrade.upgrade_unit_groups", "addupgradeunits", inputDataValue, executionContext)
	var emptyOutput nsxModel.UpgradeUnitList
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), UpgradeUnitGroupsAddupgradeunitsOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.UpgradeUnitList), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), uIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (uIface *upgradeUnitGroupsClient) Create(upgradeUnitGroupParam nsxModel.UpgradeUnitGroup) (nsxModel.UpgradeUnitGroup, error) {
	typeConverter := uIface.connector.TypeConverter()
	executionContext := uIface.connector.NewExecutionContext()
	operationRestMetaData := upgradeUnitGroupsCreateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(upgradeUnitGroupsCreateInputType(), typeConverter)
	sv.AddStructField("UpgradeUnitGroup", upgradeUnitGroupParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.UpgradeUnitGroup
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := uIface.connector.GetApiProvider().Invoke("com.vmware.nsx.upgrade.upgrade_unit_groups", "create", inputDataValue, executionContext)
	var emptyOutput nsxModel.UpgradeUnitGroup
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), UpgradeUnitGroupsCreateOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.UpgradeUnitGroup), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), uIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (uIface *upgradeUnitGroupsClient) Delete(groupIdParam string) error {
	typeConverter := uIface.connector.TypeConverter()
	executionContext := uIface.connector.NewExecutionContext()
	operationRestMetaData := upgradeUnitGroupsDeleteRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(upgradeUnitGroupsDeleteInputType(), typeConverter)
	sv.AddStructField("GroupId", groupIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := uIface.connector.GetApiProvider().Invoke("com.vmware.nsx.upgrade.upgrade_unit_groups", "delete", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), uIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (uIface *upgradeUnitGroupsClient) Get(groupIdParam string, summaryParam *bool) (nsxModel.UpgradeUnitGroup, error) {
	typeConverter := uIface.connector.TypeConverter()
	executionContext := uIface.connector.NewExecutionContext()
	operationRestMetaData := upgradeUnitGroupsGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(upgradeUnitGroupsGetInputType(), typeConverter)
	sv.AddStructField("GroupId", groupIdParam)
	sv.AddStructField("Summary", summaryParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.UpgradeUnitGroup
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := uIface.connector.GetApiProvider().Invoke("com.vmware.nsx.upgrade.upgrade_unit_groups", "get", inputDataValue, executionContext)
	var emptyOutput nsxModel.UpgradeUnitGroup
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), UpgradeUnitGroupsGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.UpgradeUnitGroup), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), uIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (uIface *upgradeUnitGroupsClient) List(componentTypeParam *string, cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string, summaryParam *bool, syncParam *bool) (nsxModel.UpgradeUnitGroupListResult, error) {
	typeConverter := uIface.connector.TypeConverter()
	executionContext := uIface.connector.NewExecutionContext()
	operationRestMetaData := upgradeUnitGroupsListRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(upgradeUnitGroupsListInputType(), typeConverter)
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
		var emptyOutput nsxModel.UpgradeUnitGroupListResult
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := uIface.connector.GetApiProvider().Invoke("com.vmware.nsx.upgrade.upgrade_unit_groups", "list", inputDataValue, executionContext)
	var emptyOutput nsxModel.UpgradeUnitGroupListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), UpgradeUnitGroupsListOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.UpgradeUnitGroupListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), uIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (uIface *upgradeUnitGroupsClient) Reorder(groupIdParam string, reorderRequestParam nsxModel.ReorderRequest) error {
	typeConverter := uIface.connector.TypeConverter()
	executionContext := uIface.connector.NewExecutionContext()
	operationRestMetaData := upgradeUnitGroupsReorderRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(upgradeUnitGroupsReorderInputType(), typeConverter)
	sv.AddStructField("GroupId", groupIdParam)
	sv.AddStructField("ReorderRequest", reorderRequestParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := uIface.connector.GetApiProvider().Invoke("com.vmware.nsx.upgrade.upgrade_unit_groups", "reorder", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), uIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (uIface *upgradeUnitGroupsClient) Update(groupIdParam string, upgradeUnitGroupParam nsxModel.UpgradeUnitGroup) (nsxModel.UpgradeUnitGroup, error) {
	typeConverter := uIface.connector.TypeConverter()
	executionContext := uIface.connector.NewExecutionContext()
	operationRestMetaData := upgradeUnitGroupsUpdateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(upgradeUnitGroupsUpdateInputType(), typeConverter)
	sv.AddStructField("GroupId", groupIdParam)
	sv.AddStructField("UpgradeUnitGroup", upgradeUnitGroupParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.UpgradeUnitGroup
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := uIface.connector.GetApiProvider().Invoke("com.vmware.nsx.upgrade.upgrade_unit_groups", "update", inputDataValue, executionContext)
	var emptyOutput nsxModel.UpgradeUnitGroup
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), UpgradeUnitGroupsUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.UpgradeUnitGroup), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), uIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
