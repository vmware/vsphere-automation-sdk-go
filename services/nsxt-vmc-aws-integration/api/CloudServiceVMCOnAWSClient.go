// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: CloudServiceVMCOnAWS
// Used by client-side stubs.

package api

import (
	"github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-vmc-aws-integration/model"
)

const _ = core.SupportedByRuntimeVersion1

type CloudServiceVMCOnAWSClient interface {

	// Retrieve the shadow account and linked VPC account information from VMC provider. This API is a live query to VMC provider.
	// @return com.vmware.model.VMCAccounts
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ListAccounts() (model.VMCAccounts, error)
}

type cloudServiceVMCOnAWSClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewCloudServiceVMCOnAWSClient(connector client.Connector) *cloudServiceVMCOnAWSClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.api.cloud_service_VMC_on_AWS")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"list_accounts": core.NewMethodIdentifier(interfaceIdentifier, "list_accounts"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	cIface := cloudServiceVMCOnAWSClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &cIface
}

func (cIface *cloudServiceVMCOnAWSClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := cIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (cIface *cloudServiceVMCOnAWSClient) ListAccounts() (model.VMCAccounts, error) {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(cloudServiceVMCOnAWSListAccountsInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.VMCAccounts
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := cloudServiceVMCOnAWSListAccountsRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	cIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.api.cloud_service_VMC_on_AWS", "list_accounts", inputDataValue, executionContext)
	var emptyOutput model.VMCAccounts
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), cloudServiceVMCOnAWSListAccountsOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.VMCAccounts), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
