// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: CompatibleSubnetsAsync
// Used by client-side stubs.

package account_link

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	vmcModel "github.com/vmware/vsphere-automation-sdk-go/services/vmc/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type CompatibleSubnetsAsyncClient interface {

	// Gets a customer's compatible subnets for account linking via a task. The information is returned as a member of the task (found in task.params['subnet_list_result'] when you are notified it is complete), and it's documented under ref /definitions/AwsCompatibleSubnets
	//
	// @param orgParam Organization identifier (required)
	// @param linkedAccountIdParam The linked connected account identifier (required)
	// @param regionParam The region of the cloud resources to work in (optional)
	// @param sddcParam sddc (optional)
	// @param instanceTypeParam The server instance type to be used. (optional)
	// @param sddcTypeParam The sddc type to be used. (1NODE, SingleAZ, MultiAZ) (optional)
	// @param numOfHostsParam The number of hosts (optional)
	// @return com.vmware.vmc.model.Task
	//
	// @throws Unauthenticated  Unauthorized
	// @throws Unauthorized  Forbidden
	Get(orgParam string, linkedAccountIdParam string, regionParam *string, sddcParam *string, instanceTypeParam *string, sddcTypeParam *string, numOfHostsParam *int64) (vmcModel.Task, error)

	// Sets which subnet to use to link accounts and finishes the linking process via a task
	//
	// @param awsSubnetParam The subnet chosen by the customer (required)
	// @param orgParam Organization identifier (required)
	// @return com.vmware.vmc.model.Task
	//
	// @throws Unauthenticated  Unauthorized
	// @throws Unauthorized  Forbidden
	Post(awsSubnetParam vmcModel.AwsSubnet, orgParam string) (vmcModel.Task, error)
}

type compatibleSubnetsAsyncClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewCompatibleSubnetsAsyncClient(connector vapiProtocolClient_.Connector) *compatibleSubnetsAsyncClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.vmc.orgs.account_link.compatible_subnets_async")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"get":  vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
		"post": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "post"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	cIface := compatibleSubnetsAsyncClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &cIface
}

func (cIface *compatibleSubnetsAsyncClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := cIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (cIface *compatibleSubnetsAsyncClient) Get(orgParam string, linkedAccountIdParam string, regionParam *string, sddcParam *string, instanceTypeParam *string, sddcTypeParam *string, numOfHostsParam *int64) (vmcModel.Task, error) {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	operationRestMetaData := compatibleSubnetsAsyncGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(compatibleSubnetsAsyncGetInputType(), typeConverter)
	sv.AddStructField("Org", orgParam)
	sv.AddStructField("LinkedAccountId", linkedAccountIdParam)
	sv.AddStructField("Region", regionParam)
	sv.AddStructField("Sddc", sddcParam)
	sv.AddStructField("InstanceType", instanceTypeParam)
	sv.AddStructField("SddcType", sddcTypeParam)
	sv.AddStructField("NumOfHosts", numOfHostsParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput vmcModel.Task
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.vmc.orgs.account_link.compatible_subnets_async", "get", inputDataValue, executionContext)
	var emptyOutput vmcModel.Task
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), CompatibleSubnetsAsyncGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(vmcModel.Task), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (cIface *compatibleSubnetsAsyncClient) Post(awsSubnetParam vmcModel.AwsSubnet, orgParam string) (vmcModel.Task, error) {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	operationRestMetaData := compatibleSubnetsAsyncPostRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(compatibleSubnetsAsyncPostInputType(), typeConverter)
	sv.AddStructField("AwsSubnet", awsSubnetParam)
	sv.AddStructField("Org", orgParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput vmcModel.Task
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.vmc.orgs.account_link.compatible_subnets_async", "post", inputDataValue, executionContext)
	var emptyOutput vmcModel.Task
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), CompatibleSubnetsAsyncPostOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(vmcModel.Task), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
