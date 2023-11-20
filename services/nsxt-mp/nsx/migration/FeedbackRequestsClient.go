// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: FeedbackRequests
// Used by client-side stubs.

package migration

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	nsxModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type FeedbackRequestsClient interface {

	// Get feedback details of NSX-V to be migrated.
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
	// @return com.vmware.nsx.model.MigrationFeedbackRequestListResult
	//
	// @throws Unauthenticated  Unauthorized
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(categoryParam *string, cursorParam *string, federationSiteIdParam *string, hashParam *string, includedFieldsParam *string, networkLayerParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string, stateParam *string, subCategoryParam *string) (nsxModel.MigrationFeedbackRequestListResult, error)
}

type feedbackRequestsClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewFeedbackRequestsClient(connector vapiProtocolClient_.Connector) *feedbackRequestsClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx.migration.feedback_requests")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"list": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "list"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	fIface := feedbackRequestsClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &fIface
}

func (fIface *feedbackRequestsClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := fIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (fIface *feedbackRequestsClient) List(categoryParam *string, cursorParam *string, federationSiteIdParam *string, hashParam *string, includedFieldsParam *string, networkLayerParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string, stateParam *string, subCategoryParam *string) (nsxModel.MigrationFeedbackRequestListResult, error) {
	typeConverter := fIface.connector.TypeConverter()
	executionContext := fIface.connector.NewExecutionContext()
	operationRestMetaData := feedbackRequestsListRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(feedbackRequestsListInputType(), typeConverter)
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
		var emptyOutput nsxModel.MigrationFeedbackRequestListResult
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := fIface.connector.GetApiProvider().Invoke("com.vmware.nsx.migration.feedback_requests", "list", inputDataValue, executionContext)
	var emptyOutput nsxModel.MigrationFeedbackRequestListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), FeedbackRequestsListOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.MigrationFeedbackRequestListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), fIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
