// Copyright (c) 2019-2026 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: LbNodeUsageSummary
// Used by client-side stubs.

package infra

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	nsx_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type LbNodeUsageSummaryClient interface {

	// The API is used to retrieve the load balancer node usage summary of all nodes for every enforcement point. - If the parameter ?include_usages=true exists, the property node_usages are included in response. By default, the property node_usages is not included in response. - If parameter ?enforcement_point_path=<enforcement-point-path> exists, only node usage summary from specific enforcement point is included in response. If no enforcement point path is specified, information will be aggregated from each enforcement point. - If parameter ?node_type=VirtualNetworkAppliance exists, it will only return the LB node usage summary for virtual network appliances. If parameter ?node_type=EdgeNode exists, it will only return the LB node usage summary for edge nodes. If the parameter is not specified, it will return LB node usage summary for both edge nodes and virtual network appliances. VCF Load Balancer availability in terms of use-cases and editions is specified in the VMware Cloud Foundation Feature Comparison and Upgrade Paths Guide. Please review before consuming these APIs.
	//
	// @param enforcementPointPathParam Specify enforcement point path. (optional)
	// @param includeUsagesParam Specify whether to include usages in response. (optional)
	// @param nodeTypeParam Specify node type. (optional)
	// @return com.vmware.nsx_policy.model.AggregateLBNodeUsageSummary
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(enforcementPointPathParam *string, includeUsagesParam *bool, nodeTypeParam *string) (nsx_policyModel.AggregateLBNodeUsageSummary, error)
}

type lbNodeUsageSummaryClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewLbNodeUsageSummaryClient(connector vapiProtocolClient_.Connector) *lbNodeUsageSummaryClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx_policy.infra.lb_node_usage_summary")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"get": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	lIface := lbNodeUsageSummaryClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &lIface
}

func (lIface *lbNodeUsageSummaryClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := lIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (lIface *lbNodeUsageSummaryClient) Get(enforcementPointPathParam *string, includeUsagesParam *bool, nodeTypeParam *string) (nsx_policyModel.AggregateLBNodeUsageSummary, error) {
	typeConverter := lIface.connector.TypeConverter()
	executionContext := lIface.connector.NewExecutionContext()
	operationRestMetaData := lbNodeUsageSummaryGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(lbNodeUsageSummaryGetInputType(), typeConverter)
	sv.AddStructField("EnforcementPointPath", enforcementPointPathParam)
	sv.AddStructField("IncludeUsages", includeUsagesParam)
	sv.AddStructField("NodeType", nodeTypeParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.AggregateLBNodeUsageSummary
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := lIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.lb_node_usage_summary", "get", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.AggregateLBNodeUsageSummary
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), LbNodeUsageSummaryGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.AggregateLBNodeUsageSummary), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), lIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
