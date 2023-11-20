// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: WorkflowRouteDestination
// Used by client-side stubs.

package vmc_aws

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	vmwarecloudVmc_aws_cho_providerModel "github.com/vmware/vsphere-automation-sdk-go/services/vmwarecloud/vmc_aws_cho_provider/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type WorkflowRouteDestinationClient interface {

	// Get workflow destination.
	//
	// @param orgIdParam organization identifier
	// @param workflowRoutingPayloadParam Payload for getting the workflow destination
	// @return OK
	//
	// @throws Unauthenticated Unauthorized
	// @throws Unauthorized Forbidden
	GetWorkflowDestination(orgIdParam string, workflowRoutingPayloadParam vmwarecloudVmc_aws_cho_providerModel.WorkflowRoutingPayload) (vmwarecloudVmc_aws_cho_providerModel.WorkflowDestinationResponse, error)
}

type workflowRouteDestinationClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewWorkflowRouteDestinationClient(connector vapiProtocolClient_.Connector) *workflowRouteDestinationClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.vmwarecloud.vmc_aws_cho_provider.api.inventory.vmc_aws.workflow_route_destination")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"get_workflow_destination": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get_workflow_destination"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	wIface := workflowRouteDestinationClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &wIface
}

func (wIface *workflowRouteDestinationClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := wIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (wIface *workflowRouteDestinationClient) GetWorkflowDestination(orgIdParam string, workflowRoutingPayloadParam vmwarecloudVmc_aws_cho_providerModel.WorkflowRoutingPayload) (vmwarecloudVmc_aws_cho_providerModel.WorkflowDestinationResponse, error) {
	typeConverter := wIface.connector.TypeConverter()
	executionContext := wIface.connector.NewExecutionContext()
	operationRestMetaData := workflowRouteDestinationGetWorkflowDestinationRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(workflowRouteDestinationGetWorkflowDestinationInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("WorkflowRoutingPayload", workflowRoutingPayloadParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput vmwarecloudVmc_aws_cho_providerModel.WorkflowDestinationResponse
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := wIface.connector.GetApiProvider().Invoke("com.vmware.vmwarecloud.vmc_aws_cho_provider.api.inventory.vmc_aws.workflow_route_destination", "get_workflow_destination", inputDataValue, executionContext)
	var emptyOutput vmwarecloudVmc_aws_cho_providerModel.WorkflowDestinationResponse
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), WorkflowRouteDestinationGetWorkflowDestinationOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(vmwarecloudVmc_aws_cho_providerModel.WorkflowDestinationResponse), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), wIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
