// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: EnableVrScaleOut
// Used by client-side stubs.

package draas

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	vmcDraasModel "github.com/vmware/vsphere-automation-sdk-go/services/vmc/draas/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type EnableVrScaleOutClient interface {

	// Enable VR scale out for Activated with DR SDDC
	//
	// @param orgParam Organization identifier (required)
	// @param sddcParam sddc identifier (required)
	// @param scaleOutParam If = 'true', will set scale-out property in hms config to true. (required)
	// @param dryRunParam Run the HMS scale-out API with dryRun=true,false or skip running it. (optional)
	// @param isRetryParam If = 'true', will not send start notifications (optional)
	// @return com.vmware.vmc.draas.model.Task
	//
	// @throws Unauthenticated  Unauthorized
	// @throws InvalidRequest  Bad Request
	// @throws Unauthorized  Forbidden
	Post(orgParam string, sddcParam string, scaleOutParam bool, dryRunParam *string, isRetryParam *bool) (vmcDraasModel.Task, error)

	// Set VR scale out property
	//
	// @param orgParam Organization identifier (required)
	// @param sddcParam sddc identifier (required)
	// @param scaleOutParam If = 'true', will set scale-out property in hms config to true. (required)
	// @return com.vmware.vmc.draas.model.Task
	//
	// @throws Unauthenticated  Unauthorized
	// @throws InvalidRequest  Bad Request
	// @throws Unauthorized  Forbidden
	Post0(orgParam string, sddcParam string, scaleOutParam bool) (vmcDraasModel.Task, error)
}

type enableVrScaleOutClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewEnableVrScaleOutClient(connector vapiProtocolClient_.Connector) *enableVrScaleOutClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.vmc.draas.enable_vr_scale_out")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"post":   vapiCore_.NewMethodIdentifier(interfaceIdentifier, "post"),
		"post_0": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "post_0"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	eIface := enableVrScaleOutClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &eIface
}

func (eIface *enableVrScaleOutClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := eIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (eIface *enableVrScaleOutClient) Post(orgParam string, sddcParam string, scaleOutParam bool, dryRunParam *string, isRetryParam *bool) (vmcDraasModel.Task, error) {
	typeConverter := eIface.connector.TypeConverter()
	executionContext := eIface.connector.NewExecutionContext()
	operationRestMetaData := enableVrScaleOutPostRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(enableVrScaleOutPostInputType(), typeConverter)
	sv.AddStructField("Org", orgParam)
	sv.AddStructField("Sddc", sddcParam)
	sv.AddStructField("ScaleOut", scaleOutParam)
	sv.AddStructField("DryRun", dryRunParam)
	sv.AddStructField("IsRetry", isRetryParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput vmcDraasModel.Task
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := eIface.connector.GetApiProvider().Invoke("com.vmware.vmc.draas.enable_vr_scale_out", "post", inputDataValue, executionContext)
	var emptyOutput vmcDraasModel.Task
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), EnableVrScaleOutPostOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(vmcDraasModel.Task), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), eIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (eIface *enableVrScaleOutClient) Post0(orgParam string, sddcParam string, scaleOutParam bool) (vmcDraasModel.Task, error) {
	typeConverter := eIface.connector.TypeConverter()
	executionContext := eIface.connector.NewExecutionContext()
	operationRestMetaData := enableVrScaleOutPost0RestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(enableVrScaleOutPost0InputType(), typeConverter)
	sv.AddStructField("Org", orgParam)
	sv.AddStructField("Sddc", sddcParam)
	sv.AddStructField("ScaleOut", scaleOutParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput vmcDraasModel.Task
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := eIface.connector.GetApiProvider().Invoke("com.vmware.vmc.draas.enable_vr_scale_out", "post_0", inputDataValue, executionContext)
	var emptyOutput vmcDraasModel.Task
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), EnableVrScaleOutPost0OutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(vmcDraasModel.Task), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), eIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
