// Copyright (c) 2019-2026 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: HypervisorExcludeList
// Used by client-side stubs.

package hypervisor_security

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	nsx_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type HypervisorExcludeListClient interface {

	// Read the singleton Hypervisor Security Exclude List containing NSX Group paths that are excluded from Hypervisor Security enforcement.
	// @return com.vmware.nsx_policy.model.HypervisorSecurityExcludeList
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get() (nsx_policyModel.HypervisorSecurityExcludeList, error)

	// Partially update the singleton Hypervisor Security Exclude List. Only the fields present in the request body are updated; omitted fields retain their current values.
	//
	// @param hypervisorSecurityExcludeListParam (required)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Patch(hypervisorSecurityExcludeListParam nsx_policyModel.HypervisorSecurityExcludeList) error

	// Create or replace the singleton Hypervisor Security Exclude List. The request body must contain the complete desired set of member Group paths; any existing members not present in the request body will be removed.
	//
	// @param hypervisorSecurityExcludeListParam (required)
	// @return com.vmware.nsx_policy.model.HypervisorSecurityExcludeList
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(hypervisorSecurityExcludeListParam nsx_policyModel.HypervisorSecurityExcludeList) (nsx_policyModel.HypervisorSecurityExcludeList, error)
}

type hypervisorExcludeListClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewHypervisorExcludeListClient(connector vapiProtocolClient_.Connector) *hypervisorExcludeListClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx_policy.infra.settings.firewall.hypervisor_security.hypervisor_exclude_list")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"get":    vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
		"patch":  vapiCore_.NewMethodIdentifier(interfaceIdentifier, "patch"),
		"update": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "update"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	hIface := hypervisorExcludeListClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &hIface
}

func (hIface *hypervisorExcludeListClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := hIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (hIface *hypervisorExcludeListClient) Get() (nsx_policyModel.HypervisorSecurityExcludeList, error) {
	typeConverter := hIface.connector.TypeConverter()
	executionContext := hIface.connector.NewExecutionContext()
	operationRestMetaData := hypervisorExcludeListGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(hypervisorExcludeListGetInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.HypervisorSecurityExcludeList
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := hIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.settings.firewall.hypervisor_security.hypervisor_exclude_list", "get", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.HypervisorSecurityExcludeList
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), HypervisorExcludeListGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.HypervisorSecurityExcludeList), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), hIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (hIface *hypervisorExcludeListClient) Patch(hypervisorSecurityExcludeListParam nsx_policyModel.HypervisorSecurityExcludeList) error {
	typeConverter := hIface.connector.TypeConverter()
	executionContext := hIface.connector.NewExecutionContext()
	operationRestMetaData := hypervisorExcludeListPatchRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(hypervisorExcludeListPatchInputType(), typeConverter)
	sv.AddStructField("HypervisorSecurityExcludeList", hypervisorSecurityExcludeListParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := hIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.settings.firewall.hypervisor_security.hypervisor_exclude_list", "patch", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), hIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (hIface *hypervisorExcludeListClient) Update(hypervisorSecurityExcludeListParam nsx_policyModel.HypervisorSecurityExcludeList) (nsx_policyModel.HypervisorSecurityExcludeList, error) {
	typeConverter := hIface.connector.TypeConverter()
	executionContext := hIface.connector.NewExecutionContext()
	operationRestMetaData := hypervisorExcludeListUpdateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(hypervisorExcludeListUpdateInputType(), typeConverter)
	sv.AddStructField("HypervisorSecurityExcludeList", hypervisorSecurityExcludeListParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.HypervisorSecurityExcludeList
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := hIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.settings.firewall.hypervisor_security.hypervisor_exclude_list", "update", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.HypervisorSecurityExcludeList
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), HypervisorExcludeListUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.HypervisorSecurityExcludeList), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), hIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
