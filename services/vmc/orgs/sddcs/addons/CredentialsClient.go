// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Credentials
// Used by client-side stubs.

package addons

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	vmcModel "github.com/vmware/vsphere-automation-sdk-go/services/vmc/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type CredentialsClient interface {

	// Associated a new add on credentials with the SDDC such as HCX
	//
	// @param orgParam Org id of the associated SDDC (required)
	// @param sddcIdParam Id of the SDDC (required)
	// @param addonTypeParam Add on type (required)
	// @param credentialsParam Credentials creation payload (required)
	// @return com.vmware.vmc.model.NewCredentials
	//
	// @throws ConcurrentChange  Credentials with same name exists with in the scope of addOnType
	// @throws InvalidRequest  Invalid input
	// @throws Unauthorized  Forbidden
	Create(orgParam string, sddcIdParam string, addonTypeParam string, credentialsParam vmcModel.NewCredentials) (vmcModel.NewCredentials, error)

	// Get credential details by name
	//
	// @param orgParam Org id of the associated SDDC (required)
	// @param sddcIdParam Id of the SDDC (required)
	// @param addonTypeParam Add on type (required)
	// @param nameParam name of the credentials (required)
	// @return com.vmware.vmc.model.NewCredentials
	//
	// @throws Unauthorized  Forbidden
	Get(orgParam string, sddcIdParam string, addonTypeParam string, nameParam string) (vmcModel.NewCredentials, error)

	// List all the credentials assoicated with an addon type with in a SDDC
	//
	// @param orgParam Org id of the associated SDDC (required)
	// @param sddcIdParam Id of the SDDC (required)
	// @param addonTypeParam Add on type (required)
	//
	// @throws Unauthorized  Forbidden
	List(orgParam string, sddcIdParam string, addonTypeParam string) ([]vmcModel.NewCredentials, error)

	// Update credential details
	//
	// @param orgParam Org id of the associated SDDC (required)
	// @param sddcIdParam Id of the SDDC (required)
	// @param addonTypeParam Add on type (required)
	// @param nameParam name of the credentials (required)
	// @param credentialsParam Credentials update payload (required)
	// @return com.vmware.vmc.model.NewCredentials
	//
	// @throws InvalidRequest  Bad request
	// @throws Unauthorized  Forbidden
	Update(orgParam string, sddcIdParam string, addonTypeParam string, nameParam string, credentialsParam vmcModel.UpdateCredentials) (vmcModel.NewCredentials, error)
}

type credentialsClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewCredentialsClient(connector vapiProtocolClient_.Connector) *credentialsClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.vmc.orgs.sddcs.addons.credentials")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"create": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "create"),
		"get":    vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
		"list":   vapiCore_.NewMethodIdentifier(interfaceIdentifier, "list"),
		"update": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "update"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	cIface := credentialsClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &cIface
}

func (cIface *credentialsClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := cIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (cIface *credentialsClient) Create(orgParam string, sddcIdParam string, addonTypeParam string, credentialsParam vmcModel.NewCredentials) (vmcModel.NewCredentials, error) {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	operationRestMetaData := credentialsCreateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(credentialsCreateInputType(), typeConverter)
	sv.AddStructField("Org", orgParam)
	sv.AddStructField("SddcId", sddcIdParam)
	sv.AddStructField("AddonType", addonTypeParam)
	sv.AddStructField("Credentials", credentialsParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput vmcModel.NewCredentials
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.vmc.orgs.sddcs.addons.credentials", "create", inputDataValue, executionContext)
	var emptyOutput vmcModel.NewCredentials
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), CredentialsCreateOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(vmcModel.NewCredentials), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (cIface *credentialsClient) Get(orgParam string, sddcIdParam string, addonTypeParam string, nameParam string) (vmcModel.NewCredentials, error) {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	operationRestMetaData := credentialsGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(credentialsGetInputType(), typeConverter)
	sv.AddStructField("Org", orgParam)
	sv.AddStructField("SddcId", sddcIdParam)
	sv.AddStructField("AddonType", addonTypeParam)
	sv.AddStructField("Name", nameParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput vmcModel.NewCredentials
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.vmc.orgs.sddcs.addons.credentials", "get", inputDataValue, executionContext)
	var emptyOutput vmcModel.NewCredentials
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), CredentialsGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(vmcModel.NewCredentials), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (cIface *credentialsClient) List(orgParam string, sddcIdParam string, addonTypeParam string) ([]vmcModel.NewCredentials, error) {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	operationRestMetaData := credentialsListRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(credentialsListInputType(), typeConverter)
	sv.AddStructField("Org", orgParam)
	sv.AddStructField("SddcId", sddcIdParam)
	sv.AddStructField("AddonType", addonTypeParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput []vmcModel.NewCredentials
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.vmc.orgs.sddcs.addons.credentials", "list", inputDataValue, executionContext)
	var emptyOutput []vmcModel.NewCredentials
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), CredentialsListOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.([]vmcModel.NewCredentials), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (cIface *credentialsClient) Update(orgParam string, sddcIdParam string, addonTypeParam string, nameParam string, credentialsParam vmcModel.UpdateCredentials) (vmcModel.NewCredentials, error) {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	operationRestMetaData := credentialsUpdateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(credentialsUpdateInputType(), typeConverter)
	sv.AddStructField("Org", orgParam)
	sv.AddStructField("SddcId", sddcIdParam)
	sv.AddStructField("AddonType", addonTypeParam)
	sv.AddStructField("Name", nameParam)
	sv.AddStructField("Credentials", credentialsParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput vmcModel.NewCredentials
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.vmc.orgs.sddcs.addons.credentials", "update", inputDataValue, executionContext)
	var emptyOutput vmcModel.NewCredentials
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), CredentialsUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(vmcModel.NewCredentials), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
