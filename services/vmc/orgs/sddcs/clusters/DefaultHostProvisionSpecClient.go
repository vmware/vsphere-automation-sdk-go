
/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Client stubs for service: HostProvisionSpec
 * Functions that implement the generated HostProvisionSpecClient interface
 */


package clusters

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/services/vmc/model"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/core"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/lib"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/log"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol/client"
)

type DefaultHostProvisionSpecClient struct {
	interfaceName       string
	interfaceDefinition core.InterfaceDefinition
	methodIdentifiers   []core.MethodIdentifier
	methodNameToDefMap  map[string]*core.MethodDefinition
	errorBindingMap     map[string]bindings.BindingType
	interfaceIdentifier core.InterfaceIdentifier
	connector           client.Connector
}

func NewDefaultHostProvisionSpecClient(connector client.Connector) *DefaultHostProvisionSpecClient {
	interfaceName := "com.vmware.vmc.orgs.sddcs.clusters.host_provision_spec"
	interfaceIdentifier := core.NewInterfaceIdentifier(interfaceName)
	methodIdentifiers := []core.MethodIdentifier{
		core.NewMethodIdentifier(interfaceIdentifier, "get"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorBindingMap := make(map[string]bindings.BindingType)
	errorBindingMap[errors.InternalServerError{}.Error()] = errors.InternalServerErrorBindingType()
	errorBindingMap[errors.InvalidArgument{}.Error()] = errors.InvalidArgumentBindingType()
	errorBindingMap[errors.OperationNotFound{}.Error()] = errors.OperationNotFoundBindingType()
	errorBindingMap[errors.UnexpectedInput{}.Error()] = errors.UnexpectedInputBindingType()
	errorBindingMap[errors.ServiceUnavailable{}.Error()] = errors.ServiceUnavailableBindingType()
	errorBindingMap[errors.TimedOut{}.Error()] = errors.TimedOutBindingType()
	hIface := DefaultHostProvisionSpecClient{interfaceName: interfaceName, methodIdentifiers: methodIdentifiers, interfaceDefinition: interfaceDefinition, errorBindingMap: errorBindingMap, interfaceIdentifier: interfaceIdentifier, connector: connector}
	hIface.methodNameToDefMap = make(map[string]*core.MethodDefinition)
	hIface.methodNameToDefMap["get"] = hIface.getMethodDefinition()
	return &hIface
}

func (hIface *DefaultHostProvisionSpecClient) Get(orgParam string, sddcParam string, clusterParam string, actionParam *string) (model.HostProvisionSpec, error) {
	typeConverter := hIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(hIface.interfaceIdentifier, "get")
	sv := bindings.NewStructValueBuilder(hostProvisionSpecGetInputType(), typeConverter)
	sv.AddStructField("Org", orgParam)
	sv.AddStructField("Sddc", sddcParam)
	sv.AddStructField("Cluster", clusterParam)
	sv.AddStructField("Action", actionParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.HostProvisionSpec
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := hostProvisionSpecGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	hIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := hIface.Invoke(hIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput model.HostProvisionSpec
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), hostProvisionSpecGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.HostProvisionSpec), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), hIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}


func (hIface *DefaultHostProvisionSpecClient) Invoke(ctx *core.ExecutionContext, methodId core.MethodIdentifier, inputDataValue data.DataValue) core.MethodResult {
	methodResult := hIface.connector.GetApiProvider().Invoke(hIface.interfaceName, methodId.Name(), inputDataValue, ctx)
	return methodResult
}


func (hIface *DefaultHostProvisionSpecClient) getMethodDefinition() *core.MethodDefinition {
	interfaceIdentifier := core.NewInterfaceIdentifier(hIface.interfaceName)
	typeConverter := hIface.connector.TypeConverter()

	input, inputError := typeConverter.ConvertToDataDefinition(hostProvisionSpecGetInputType())
	output, outputError := typeConverter.ConvertToDataDefinition(hostProvisionSpecGetOutputType())
	if inputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultHostProvisionSpecClient.get method's input - %s",
			bindings.VAPIerrorsToError(inputError).Error())
		return nil
	}
	if outputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultHostProvisionSpecClient.get method's output - %s",
			bindings.VAPIerrorsToError(outputError).Error())
		return nil
	}
	methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "get")
	errorDefinitions := make([]data.ErrorDefinition, 0)
	hIface.errorBindingMap[errors.Unauthenticated{}.Error()] = errors.UnauthenticatedBindingType()
	errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.UnauthenticatedBindingType())
	if errError1 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultHostProvisionSpecClient.get method's errors.Unauthenticated error - %s",
			bindings.VAPIerrorsToError(errError1).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
	hIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	if errError2 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultHostProvisionSpecClient.get method's errors.Unauthorized error - %s",
			bindings.VAPIerrorsToError(errError2).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
	hIface.errorBindingMap[errors.InternalServerError{}.Error()] = errors.InternalServerErrorBindingType()
	errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.InternalServerErrorBindingType())
	if errError3 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultHostProvisionSpecClient.get method's errors.InternalServerError error - %s",
			bindings.VAPIerrorsToError(errError3).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))

	methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
	return &methodDefinition
}
