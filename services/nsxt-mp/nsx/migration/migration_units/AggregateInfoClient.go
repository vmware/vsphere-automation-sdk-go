// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: AggregateInfo
// Used by client-side stubs.

package migration_units

import (
	"github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
)

const _ = core.SupportedByRuntimeVersion1

type AggregateInfoClient interface {

	// Get migration units aggregate-info
	//
	// @param componentTypeParam Component type based on which migration units to be filtered (optional)
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param groupIdParam Identifier of group based on which migration units to be filtered (optional)
	// @param hasErrorsParam Flag to indicate whether to return only migration units with errors (optional, default to false)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param metadataParam Metadata about migration unit to filter on (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param selectionStatusParam Flag to indicate whether to return only selected, only deselected or both type of migration units (optional, default to ALL)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @return com.vmware.nsx.model.MigrationUnitAggregateInfoListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(componentTypeParam *string, cursorParam *string, groupIdParam *string, hasErrorsParam *bool, includedFieldsParam *string, metadataParam *string, pageSizeParam *int64, selectionStatusParam *string, sortAscendingParam *bool, sortByParam *string) (model.MigrationUnitAggregateInfoListResult, error)
}

type aggregateInfoClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewAggregateInfoClient(connector client.Connector) *aggregateInfoClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.nsx.migration.migration_units.aggregate_info")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"list": core.NewMethodIdentifier(interfaceIdentifier, "list"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	aIface := aggregateInfoClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &aIface
}

func (aIface *aggregateInfoClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := aIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (aIface *aggregateInfoClient) List(componentTypeParam *string, cursorParam *string, groupIdParam *string, hasErrorsParam *bool, includedFieldsParam *string, metadataParam *string, pageSizeParam *int64, selectionStatusParam *string, sortAscendingParam *bool, sortByParam *string) (model.MigrationUnitAggregateInfoListResult, error) {
	typeConverter := aIface.connector.TypeConverter()
	executionContext := aIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(aggregateInfoListInputType(), typeConverter)
	sv.AddStructField("ComponentType", componentTypeParam)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("GroupId", groupIdParam)
	sv.AddStructField("HasErrors", hasErrorsParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("Metadata", metadataParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SelectionStatus", selectionStatusParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.MigrationUnitAggregateInfoListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := aggregateInfoListRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	aIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := aIface.connector.GetApiProvider().Invoke("com.vmware.nsx.migration.migration_units.aggregate_info", "list", inputDataValue, executionContext)
	var emptyOutput model.MigrationUnitAggregateInfoListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), aggregateInfoListOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.MigrationUnitAggregateInfoListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), aIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
