// Copyright (c) 2019-2026 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Dvpg_validation_results
// Used by client-side stubs.

package workflows

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	nsx_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type Dvpg_validation_resultsClient interface {

	// Paginated list of per-Distributed Virtual Port Group validation results from the last VALIDATION phase. Supports page_size, cursor, sort_by, sort_ascending, and filters: name, vlan_id, derived, has_vm, compatibility.
	//
	// @param workflowIdParam Workflow ID (required)
	// @param compatibilityParam Multi-value exact match. Multiple values can be provided as a comma-separated list. A Distributed Virtual Port Group result is included if its status matches any element of this list. Allowed values: SUPPORTED, PARTIALLY_SUPPORTED, UNSUPPORTED. (optional)
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param derivedParam Exact match on derived status. If true, only Distributed Virtual Port Groups that have been migrated are returned. If false, only Distributed Virtual Port Groups that have not been migrated are returned. (optional)
	// @param hasVmParam If true, only Distributed Virtual Port Groups that have VMs attached are returned. (optional)
	// @param includedFieldsParam Note - this parameter currently only works when used with the search APIs /policy/api/v1/search/query and /policy/api/v1/search/dsl. It is ignored for other list APIs. (optional)
	// @param nameParam Case-insensitive substring filter on the Distributed Virtual Port Group display name. (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam If true, results are sorted in ascending order (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @param vlanIdParam Exact match filter on the Distributed Virtual Port Group VLAN configuration. (optional)
	// @return com.vmware.nsx_policy.model.DvpgValidationResultsList
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(workflowIdParam string, compatibilityParam *string, cursorParam *string, derivedParam *bool, hasVmParam *bool, includedFieldsParam *string, nameParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string, vlanIdParam *string) (nsx_policyModel.DvpgValidationResultsList, error)
}

type dvpg_validation_resultsClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewDvpg_validation_resultsClient(connector vapiProtocolClient_.Connector) *dvpg_validation_resultsClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx_policy.workflows.dvpg_validation_results")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"list": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "list"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	dIface := dvpg_validation_resultsClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &dIface
}

func (dIface *dvpg_validation_resultsClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := dIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (dIface *dvpg_validation_resultsClient) List(workflowIdParam string, compatibilityParam *string, cursorParam *string, derivedParam *bool, hasVmParam *bool, includedFieldsParam *string, nameParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string, vlanIdParam *string) (nsx_policyModel.DvpgValidationResultsList, error) {
	typeConverter := dIface.connector.TypeConverter()
	executionContext := dIface.connector.NewExecutionContext()
	operationRestMetaData := dvpgValidationResultsListRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(dvpgValidationResultsListInputType(), typeConverter)
	sv.AddStructField("WorkflowId", workflowIdParam)
	sv.AddStructField("Compatibility", compatibilityParam)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("Derived", derivedParam)
	sv.AddStructField("HasVm", hasVmParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("Name", nameParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	sv.AddStructField("VlanId", vlanIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.DvpgValidationResultsList
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := dIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.workflows.dvpg_validation_results", "list", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.DvpgValidationResultsList
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), DvpgValidationResultsListOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.DvpgValidationResultsList), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), dIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
