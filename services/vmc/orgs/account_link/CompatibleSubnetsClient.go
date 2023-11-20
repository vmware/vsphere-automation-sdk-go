// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: CompatibleSubnets
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

type CompatibleSubnetsClient interface {

	// Gets a customer's compatible subnets for account linking
	//
	// @param orgParam Organization identifier (required)
	// @param linkedAccountIdParam The linked connected account identifier (required)
	// @param regionParam The region of the cloud resources to work in (optional)
	// @param sddcParam sddc (optional)
	// @param forceRefreshParam When true, forces the mappings for datacenters to be refreshed for the connected account. (optional)
	// @param instanceTypeParam The server instance type to be used. (optional)
	// @param sddcTypeParam The sddc type to be used. (1NODE, SingleAZ, MultiAZ) (optional)
	// @param numOfHostsParam The number of hosts (optional)
	// @return com.vmware.vmc.model.AwsCompatibleSubnets
	//
	// @throws Unauthenticated  Unauthorized
	// @throws Unauthorized  Forbidden
	Get(orgParam string, linkedAccountIdParam string, regionParam *string, sddcParam *string, forceRefreshParam *bool, instanceTypeParam *string, sddcTypeParam *string, numOfHostsParam *int64) (vmcModel.AwsCompatibleSubnets, error)

	// Sets which subnet to use to link accounts and finishes the linking process
	//
	// @param orgParam Organization identifier (required)
	// @return com.vmware.vmc.model.AwsSubnet
	//
	// @throws Unauthenticated  Unauthorized
	// @throws Unauthorized  Forbidden
	Post(orgParam string) (vmcModel.AwsSubnet, error)
}

type compatibleSubnetsClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewCompatibleSubnetsClient(connector vapiProtocolClient_.Connector) *compatibleSubnetsClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.vmc.orgs.account_link.compatible_subnets")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"get":  vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
		"post": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "post"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	cIface := compatibleSubnetsClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &cIface
}

func (cIface *compatibleSubnetsClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := cIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (cIface *compatibleSubnetsClient) Get(orgParam string, linkedAccountIdParam string, regionParam *string, sddcParam *string, forceRefreshParam *bool, instanceTypeParam *string, sddcTypeParam *string, numOfHostsParam *int64) (vmcModel.AwsCompatibleSubnets, error) {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	operationRestMetaData := compatibleSubnetsGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(compatibleSubnetsGetInputType(), typeConverter)
	sv.AddStructField("Org", orgParam)
	sv.AddStructField("LinkedAccountId", linkedAccountIdParam)
	sv.AddStructField("Region", regionParam)
	sv.AddStructField("Sddc", sddcParam)
	sv.AddStructField("ForceRefresh", forceRefreshParam)
	sv.AddStructField("InstanceType", instanceTypeParam)
	sv.AddStructField("SddcType", sddcTypeParam)
	sv.AddStructField("NumOfHosts", numOfHostsParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput vmcModel.AwsCompatibleSubnets
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.vmc.orgs.account_link.compatible_subnets", "get", inputDataValue, executionContext)
	var emptyOutput vmcModel.AwsCompatibleSubnets
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), CompatibleSubnetsGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(vmcModel.AwsCompatibleSubnets), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (cIface *compatibleSubnetsClient) Post(orgParam string) (vmcModel.AwsSubnet, error) {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	operationRestMetaData := compatibleSubnetsPostRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(compatibleSubnetsPostInputType(), typeConverter)
	sv.AddStructField("Org", orgParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput vmcModel.AwsSubnet
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.vmc.orgs.account_link.compatible_subnets", "post", inputDataValue, executionContext)
	var emptyOutput vmcModel.AwsSubnet
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), CompatibleSubnetsPostOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(vmcModel.AwsSubnet), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
