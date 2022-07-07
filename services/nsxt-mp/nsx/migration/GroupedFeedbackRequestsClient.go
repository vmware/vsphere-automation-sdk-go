// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: GroupedFeedbackRequests
// Used by client-side stubs.

package migration

import (
	"github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
)

const _ = core.SupportedByRuntimeVersion1

type GroupedFeedbackRequestsClient interface {

	// Get feedback details of NSX-V to be migrated, grouped by feedback type.
	//
	// @param categoryParam Category on which feedback request should be filtered (optional)
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param federationSiteIdParam Id of the site in NSX-T Federation (optional)
	// @param hashParam Hash based on which feedback request should be filtered (optional)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param networkLayerParam Network layer for which feedback is generated (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @param stateParam Filter based on current state of the feedback request (optional, default to ALL)
	// @param subCategoryParam Sub category based on which feedback request should be filtered (optional)
	// @return com.vmware.nsx.model.GroupedMigrationFeedbackRequestListResult
	// @throws Unauthenticated  Unauthorized
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(categoryParam *string, cursorParam *string, federationSiteIdParam *string, hashParam *string, includedFieldsParam *string, networkLayerParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string, stateParam *string, subCategoryParam *string) (model.GroupedMigrationFeedbackRequestListResult, error)
}

type groupedFeedbackRequestsClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewGroupedFeedbackRequestsClient(connector client.Connector) *groupedFeedbackRequestsClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.nsx.migration.grouped_feedback_requests")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"list": core.NewMethodIdentifier(interfaceIdentifier, "list"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	gIface := groupedFeedbackRequestsClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &gIface
}

func (gIface *groupedFeedbackRequestsClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := gIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (gIface *groupedFeedbackRequestsClient) List(categoryParam *string, cursorParam *string, federationSiteIdParam *string, hashParam *string, includedFieldsParam *string, networkLayerParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string, stateParam *string, subCategoryParam *string) (model.GroupedMigrationFeedbackRequestListResult, error) {
	typeConverter := gIface.connector.TypeConverter()
	executionContext := gIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(groupedFeedbackRequestsListInputType(), typeConverter)
	sv.AddStructField("Category", categoryParam)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("FederationSiteId", federationSiteIdParam)
	sv.AddStructField("Hash", hashParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("NetworkLayer", networkLayerParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	sv.AddStructField("State", stateParam)
	sv.AddStructField("SubCategory", subCategoryParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.GroupedMigrationFeedbackRequestListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := groupedFeedbackRequestsListRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	gIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := gIface.connector.GetApiProvider().Invoke("com.vmware.nsx.migration.grouped_feedback_requests", "list", inputDataValue, executionContext)
	var emptyOutput model.GroupedMigrationFeedbackRequestListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), groupedFeedbackRequestsListOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.GroupedMigrationFeedbackRequestListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), gIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
