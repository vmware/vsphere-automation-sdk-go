// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Aggregated_with_pagination
// Used by client-side stubs.

package drafts

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	nsx_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type Aggregated_with_paginationClient interface {

	// Get a paginated aggregated configuration of a given draft. This aggregated configuration is the differnece between the current published firewall configuration and a firewall configuration stored in a given draft. For an initial API call, if request_id is present in a response, then this is a paginated aggregated configuration of a given draft, containing all the security policies from the aggregated configuration. Using this request_id, more granular aggregated configuration, at security policy level, can be fetched from subsequent API calls. Absence of request_id suggests that whole aggregated configuration has been returned as a response to initial API call, as the size of aggregated configuration is not big enough to need pagination.
	//
	// @param draftIdParam (required)
	// @param requestIdParam Request identifier to track subsequent API calls (optional)
	// @param rootPathParam Path of the root object of subtree (optional)
	// @return com.vmware.nsx_policy.model.PolicyDraftPaginatedAggregatedConfigurationResult
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(draftIdParam string, requestIdParam *string, rootPathParam *string) (nsx_policyModel.PolicyDraftPaginatedAggregatedConfigurationResult, error)
}

type aggregated_with_paginationClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewAggregated_with_paginationClient(connector vapiProtocolClient_.Connector) *aggregated_with_paginationClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx_policy.global_infra.drafts.aggregated_with_pagination")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"get": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	aIface := aggregated_with_paginationClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &aIface
}

func (aIface *aggregated_with_paginationClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := aIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (aIface *aggregated_with_paginationClient) Get(draftIdParam string, requestIdParam *string, rootPathParam *string) (nsx_policyModel.PolicyDraftPaginatedAggregatedConfigurationResult, error) {
	typeConverter := aIface.connector.TypeConverter()
	executionContext := aIface.connector.NewExecutionContext()
	operationRestMetaData := aggregatedWithPaginationGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(aggregatedWithPaginationGetInputType(), typeConverter)
	sv.AddStructField("DraftId", draftIdParam)
	sv.AddStructField("RequestId", requestIdParam)
	sv.AddStructField("RootPath", rootPathParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.PolicyDraftPaginatedAggregatedConfigurationResult
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := aIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.global_infra.drafts.aggregated_with_pagination", "get", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.PolicyDraftPaginatedAggregatedConfigurationResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), AggregatedWithPaginationGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.PolicyDraftPaginatedAggregatedConfigurationResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), aIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
