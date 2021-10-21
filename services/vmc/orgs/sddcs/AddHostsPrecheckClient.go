// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: AddHostsPrecheck
// Used by client-side stubs.

package sddcs

import (
	"github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	"github.com/vmware/vsphere-automation-sdk-go/services/vmc/model"
)

const _ = core.SupportedByRuntimeVersion1

type AddHostsPrecheckClient interface {

	// Triggers an add host precheck task
	//
	// @param orgParam Organization identifier (required)
	// @param sddcParam Sddc identifier (required)
	// @param esxConfigParam esxConfig (required)
	// @return com.vmware.vmc.model.Task
	// @throws Unauthenticated  Unauthorized
	// @throws InvalidRequest  Bad Request.
	// @throws Unauthorized  Access not allowed to the operation for the current user.
	// @throws NotFound  Cannot find the SDDC with the given id.
	AddHostPrecheckTask(orgParam string, sddcParam string, esxConfigParam model.EsxConfig) (model.Task, error)
}

type addHostsPrecheckClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewAddHostsPrecheckClient(connector client.Connector) *addHostsPrecheckClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.vmc.orgs.sddcs.add_hosts_precheck")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"add_host_precheck_task": core.NewMethodIdentifier(interfaceIdentifier, "add_host_precheck_task"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	aIface := addHostsPrecheckClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &aIface
}

func (aIface *addHostsPrecheckClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := aIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (aIface *addHostsPrecheckClient) AddHostPrecheckTask(orgParam string, sddcParam string, esxConfigParam model.EsxConfig) (model.Task, error) {
	typeConverter := aIface.connector.TypeConverter()
	executionContext := aIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(addHostsPrecheckAddHostPrecheckTaskInputType(), typeConverter)
	sv.AddStructField("Org", orgParam)
	sv.AddStructField("Sddc", sddcParam)
	sv.AddStructField("EsxConfig", esxConfigParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.Task
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := addHostsPrecheckAddHostPrecheckTaskRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	aIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := aIface.connector.GetApiProvider().Invoke("com.vmware.vmc.orgs.sddcs.add_hosts_precheck", "add_host_precheck_task", inputDataValue, executionContext)
	var emptyOutput model.Task
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), addHostsPrecheckAddHostPrecheckTaskOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.Task), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), aIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
