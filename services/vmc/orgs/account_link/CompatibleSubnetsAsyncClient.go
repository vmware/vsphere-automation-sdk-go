// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: CompatibleSubnetsAsync
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
	// @throws Unauthenticated  Unauthorized
	// @throws Unauthorized  Forbidden
	Get(orgParam string, linkedAccountIdParam string, regionParam *string, sddcParam *string, instanceTypeParam *string, sddcTypeParam *string, numOfHostsParam *int64) (model.Task, error)

	// Sets which subnet to use to link accounts and finishes the linking process via a task
	//
	// @param awsSubnetParam The subnet chosen by the customer (required)
	// @param orgParam Organization identifier (required)
	// @return com.vmware.vmc.model.Task
	// @throws Unauthenticated  Unauthorized
	// @throws Unauthorized  Forbidden
	Post(awsSubnetParam model.AwsSubnet, orgParam string) (model.Task, error)
}

type compatibleSubnetsAsyncClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewCompatibleSubnetsAsyncClient(connector client.Connector) *compatibleSubnetsAsyncClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.vmc.orgs.account_link.compatible_subnets_async")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"get":  core.NewMethodIdentifier(interfaceIdentifier, "get"),
		"post": core.NewMethodIdentifier(interfaceIdentifier, "post"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	cIface := compatibleSubnetsAsyncClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &cIface
}

func (cIface *compatibleSubnetsAsyncClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := cIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (cIface *compatibleSubnetsAsyncClient) Get(orgParam string, linkedAccountIdParam string, regionParam *string, sddcParam *string, instanceTypeParam *string, sddcTypeParam *string, numOfHostsParam *int64) (model.Task, error) {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(compatibleSubnetsAsyncGetInputType(), typeConverter)
	sv.AddStructField("Org", orgParam)
	sv.AddStructField("LinkedAccountId", linkedAccountIdParam)
	sv.AddStructField("Region", regionParam)
	sv.AddStructField("Sddc", sddcParam)
	sv.AddStructField("InstanceType", instanceTypeParam)
	sv.AddStructField("SddcType", sddcTypeParam)
	sv.AddStructField("NumOfHosts", numOfHostsParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.Task
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := compatibleSubnetsAsyncGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	cIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.vmc.orgs.account_link.compatible_subnets_async", "get", inputDataValue, executionContext)
	var emptyOutput model.Task
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), compatibleSubnetsAsyncGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.Task), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (cIface *compatibleSubnetsAsyncClient) Post(awsSubnetParam model.AwsSubnet, orgParam string) (model.Task, error) {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(compatibleSubnetsAsyncPostInputType(), typeConverter)
	sv.AddStructField("AwsSubnet", awsSubnetParam)
	sv.AddStructField("Org", orgParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.Task
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := compatibleSubnetsAsyncPostRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	cIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.vmc.orgs.account_link.compatible_subnets_async", "post", inputDataValue, executionContext)
	var emptyOutput model.Task
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), compatibleSubnetsAsyncPostOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.Task), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
