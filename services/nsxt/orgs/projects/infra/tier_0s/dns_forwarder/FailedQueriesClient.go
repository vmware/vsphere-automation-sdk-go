// Copyright (c) 2019-2024 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: FailedQueries
// Used by client-side stubs.

package dns_forwarder

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	nsx_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type FailedQueriesClient interface {

	// Return the given count of recent failed DNS queries from DNS forwarder. Since the DNS forwarder is running in Active/Standby HA mode on transport nodes, the given count of queries will be returned from each nodes. Hence the total queries returned could be doubled. If no count is specified, 100 recent failed queries are returned. If the recent failures is less than the given count, all the failures will be returned. The maximum count is 1,000.
	//
	// @param orgIdParam The organization ID (required)
	// @param projectIdParam The project ID (required)
	// @param tier0IdParam (required)
	// @param countParam The count of the failed DNS queries (optional, default to 100)
	// @param enforcementPointPathParam An enforcement point path, on which the action is to be performed (optional, default to /infra/sites/default/enforcement-points/default)
	// @return com.vmware.nsx_policy.model.PolicyDnsFailedQueries
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(orgIdParam string, projectIdParam string, tier0IdParam string, countParam *int64, enforcementPointPathParam *string) (nsx_policyModel.PolicyDnsFailedQueries, error)
}

type failedQueriesClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewFailedQueriesClient(connector vapiProtocolClient_.Connector) *failedQueriesClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx_policy.orgs.projects.infra.tier_0s.dns_forwarder.failed_queries")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"get": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	fIface := failedQueriesClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &fIface
}

func (fIface *failedQueriesClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := fIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (fIface *failedQueriesClient) Get(orgIdParam string, projectIdParam string, tier0IdParam string, countParam *int64, enforcementPointPathParam *string) (nsx_policyModel.PolicyDnsFailedQueries, error) {
	typeConverter := fIface.connector.TypeConverter()
	executionContext := fIface.connector.NewExecutionContext()
	operationRestMetaData := failedQueriesGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(failedQueriesGetInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("ProjectId", projectIdParam)
	sv.AddStructField("Tier0Id", tier0IdParam)
	sv.AddStructField("Count", countParam)
	sv.AddStructField("EnforcementPointPath", enforcementPointPathParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.PolicyDnsFailedQueries
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := fIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.orgs.projects.infra.tier_0s.dns_forwarder.failed_queries", "get", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.PolicyDnsFailedQueries
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), FailedQueriesGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.PolicyDnsFailedQueries), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), fIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
