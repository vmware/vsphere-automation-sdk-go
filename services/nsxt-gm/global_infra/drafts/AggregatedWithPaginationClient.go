// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Aggregated_with_pagination
// Used by client-side stubs.

package drafts

import (
	"github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-gm/model"
)

const _ = core.SupportedByRuntimeVersion1

type Aggregated_with_paginationClient interface {

	// Get a paginated aggregated configuration of a given draft. This aggregated configuration is the differnece between the current published firewall configuration and a firewall configuration stored in a given draft. For an initial API call, if request_id is present in a response, then this is a paginated aggregated configuration of a given draft, containing all the security policies from the aggregated configuration. Using this request_id, more granular aggregated configuration, at security policy level, can be fetched from subsequent API calls. Absence of request_id suggests that whole aggregated configuration has been returned as a response to initial API call, as the size of aggregated configuration is not big enough to need pagination.
	//
	// @param draftIdParam (required)
	// @param requestIdParam Request identifier to track subsequent API calls (optional)
	// @param rootPathParam Path of the root object of subtree (optional)
	// @return com.vmware.nsx_global_policy.model.PolicyDraftPaginatedAggregatedConfigurationResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(draftIdParam string, requestIdParam *string, rootPathParam *string) (model.PolicyDraftPaginatedAggregatedConfigurationResult, error)
}

type aggregated_with_paginationClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewAggregated_with_paginationClient(connector client.Connector) *aggregated_with_paginationClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.nsx_global_policy.global_infra.drafts.aggregated_with_pagination")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"get": core.NewMethodIdentifier(interfaceIdentifier, "get"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	aIface := aggregated_with_paginationClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &aIface
}

func (aIface *aggregated_with_paginationClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := aIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (aIface *aggregated_with_paginationClient) Get(draftIdParam string, requestIdParam *string, rootPathParam *string) (model.PolicyDraftPaginatedAggregatedConfigurationResult, error) {
	typeConverter := aIface.connector.TypeConverter()
	executionContext := aIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(aggregatedWithPaginationGetInputType(), typeConverter)
	sv.AddStructField("DraftId", draftIdParam)
	sv.AddStructField("RequestId", requestIdParam)
	sv.AddStructField("RootPath", rootPathParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.PolicyDraftPaginatedAggregatedConfigurationResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := aggregatedWithPaginationGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	aIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := aIface.connector.GetApiProvider().Invoke("com.vmware.nsx_global_policy.global_infra.drafts.aggregated_with_pagination", "get", inputDataValue, executionContext)
	var emptyOutput model.PolicyDraftPaginatedAggregatedConfigurationResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), aggregatedWithPaginationGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.PolicyDraftPaginatedAggregatedConfigurationResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), aIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
