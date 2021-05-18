// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: CompatibleSubnets
// Used by client-side stubs.

package account_link

import (
	"github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	"github.com/vmware/vsphere-automation-sdk-go/services/vmc/model"
)

const _ = core.SupportedByRuntimeVersion1

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
	// @throws Unauthenticated  Unauthorized
	// @throws Unauthorized  Forbidden
	Get(orgParam string, linkedAccountIdParam string, regionParam *string, sddcParam *string, forceRefreshParam *bool, instanceTypeParam *string, sddcTypeParam *string, numOfHostsParam *int64) (model.AwsCompatibleSubnets, error)

	// Sets which subnet to use to link accounts and finishes the linking process
	//
	// @param orgParam Organization identifier (required)
	// @return com.vmware.vmc.model.AwsSubnet
	// @throws Unauthenticated  Unauthorized
	// @throws Unauthorized  Forbidden
	Post(orgParam string) (model.AwsSubnet, error)
}

type compatibleSubnetsClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewCompatibleSubnetsClient(connector client.Connector) *compatibleSubnetsClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.vmc.orgs.account_link.compatible_subnets")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"get":  core.NewMethodIdentifier(interfaceIdentifier, "get"),
		"post": core.NewMethodIdentifier(interfaceIdentifier, "post"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	cIface := compatibleSubnetsClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &cIface
}

func (cIface *compatibleSubnetsClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := cIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (cIface *compatibleSubnetsClient) Get(orgParam string, linkedAccountIdParam string, regionParam *string, sddcParam *string, forceRefreshParam *bool, instanceTypeParam *string, sddcTypeParam *string, numOfHostsParam *int64) (model.AwsCompatibleSubnets, error) {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(compatibleSubnetsGetInputType(), typeConverter)
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
		var emptyOutput model.AwsCompatibleSubnets
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := compatibleSubnetsGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	cIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.vmc.orgs.account_link.compatible_subnets", "get", inputDataValue, executionContext)
	var emptyOutput model.AwsCompatibleSubnets
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), compatibleSubnetsGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.AwsCompatibleSubnets), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (cIface *compatibleSubnetsClient) Post(orgParam string) (model.AwsSubnet, error) {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(compatibleSubnetsPostInputType(), typeConverter)
	sv.AddStructField("Org", orgParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.AwsSubnet
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := compatibleSubnetsPostRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	cIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.vmc.orgs.account_link.compatible_subnets", "post", inputDataValue, executionContext)
	var emptyOutput model.AwsSubnet
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), compatibleSubnetsPostOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.AwsSubnet), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
