// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Credentials
// Used by client-side stubs.

package addons

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/core"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/lib"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/protocol/client"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/vmc/model"
)

const _ = core.SupportedByRuntimeVersion1

type CredentialsClient interface {

	// Associated a new add on credentials with the SDDC such as HCX
	//
	// @param orgParam Org id of the associated SDDC (required)
	// @param sddcIdParam Id of the SDDC (required)
	// @param addonTypeParam Add on type (required)
	// @param credentialsParam Credentials creation payload (required)
	// @return com.vmware.vmc.model.NewCredentials
	// @throws ConcurrentChange  Credentials with same name exists with in the scope of addOnType
	// @throws InvalidRequest  Invalid input
	// @throws Unauthorized  Forbidden
	Create(orgParam string, sddcIdParam string, addonTypeParam string, credentialsParam model.NewCredentials) (model.NewCredentials, error)

	// Get credential details by name
	//
	// @param orgParam Org id of the associated SDDC (required)
	// @param sddcIdParam Id of the SDDC (required)
	// @param addonTypeParam Add on type (required)
	// @param nameParam name of the credentials (required)
	// @return com.vmware.vmc.model.NewCredentials
	// @throws Unauthorized  Forbidden
	Get(orgParam string, sddcIdParam string, addonTypeParam string, nameParam string) (model.NewCredentials, error)

	// List all the credentials assoicated with an addon type with in a SDDC
	//
	// @param orgParam Org id of the associated SDDC (required)
	// @param sddcIdParam Id of the SDDC (required)
	// @param addonTypeParam Add on type (required)
	// @throws Unauthorized  Forbidden
	List(orgParam string, sddcIdParam string, addonTypeParam string) ([]model.NewCredentials, error)

	// Update credential details
	//
	// @param orgParam Org id of the associated SDDC (required)
	// @param sddcIdParam Id of the SDDC (required)
	// @param addonTypeParam Add on type (required)
	// @param nameParam name of the credentials (required)
	// @param credentialsParam Credentials update payload (required)
	// @return com.vmware.vmc.model.NewCredentials
	// @throws InvalidRequest  Bad request
	// @throws Unauthorized  Forbidden
	Update(orgParam string, sddcIdParam string, addonTypeParam string, nameParam string, credentialsParam model.UpdateCredentials) (model.NewCredentials, error)
}

type credentialsClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewCredentialsClient(connector client.Connector) *credentialsClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.vmc.orgs.sddcs.addons.credentials")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"create": core.NewMethodIdentifier(interfaceIdentifier, "create"),
		"get":    core.NewMethodIdentifier(interfaceIdentifier, "get"),
		"list":   core.NewMethodIdentifier(interfaceIdentifier, "list"),
		"update": core.NewMethodIdentifier(interfaceIdentifier, "update"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	cIface := credentialsClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &cIface
}

func (cIface *credentialsClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := cIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (cIface *credentialsClient) Create(orgParam string, sddcIdParam string, addonTypeParam string, credentialsParam model.NewCredentials) (model.NewCredentials, error) {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(credentialsCreateInputType(), typeConverter)
	sv.AddStructField("Org", orgParam)
	sv.AddStructField("SddcId", sddcIdParam)
	sv.AddStructField("AddonType", addonTypeParam)
	sv.AddStructField("Credentials", credentialsParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.NewCredentials
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := credentialsCreateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	cIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.vmc.orgs.sddcs.addons.credentials", "create", inputDataValue, executionContext)
	var emptyOutput model.NewCredentials
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), credentialsCreateOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.NewCredentials), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (cIface *credentialsClient) Get(orgParam string, sddcIdParam string, addonTypeParam string, nameParam string) (model.NewCredentials, error) {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(credentialsGetInputType(), typeConverter)
	sv.AddStructField("Org", orgParam)
	sv.AddStructField("SddcId", sddcIdParam)
	sv.AddStructField("AddonType", addonTypeParam)
	sv.AddStructField("Name", nameParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.NewCredentials
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := credentialsGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	cIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.vmc.orgs.sddcs.addons.credentials", "get", inputDataValue, executionContext)
	var emptyOutput model.NewCredentials
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), credentialsGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.NewCredentials), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (cIface *credentialsClient) List(orgParam string, sddcIdParam string, addonTypeParam string) ([]model.NewCredentials, error) {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(credentialsListInputType(), typeConverter)
	sv.AddStructField("Org", orgParam)
	sv.AddStructField("SddcId", sddcIdParam)
	sv.AddStructField("AddonType", addonTypeParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput []model.NewCredentials
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := credentialsListRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	cIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.vmc.orgs.sddcs.addons.credentials", "list", inputDataValue, executionContext)
	var emptyOutput []model.NewCredentials
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), credentialsListOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.([]model.NewCredentials), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (cIface *credentialsClient) Update(orgParam string, sddcIdParam string, addonTypeParam string, nameParam string, credentialsParam model.UpdateCredentials) (model.NewCredentials, error) {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(credentialsUpdateInputType(), typeConverter)
	sv.AddStructField("Org", orgParam)
	sv.AddStructField("SddcId", sddcIdParam)
	sv.AddStructField("AddonType", addonTypeParam)
	sv.AddStructField("Name", nameParam)
	sv.AddStructField("Credentials", credentialsParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.NewCredentials
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := credentialsUpdateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	cIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.vmc.orgs.sddcs.addons.credentials", "update", inputDataValue, executionContext)
	var emptyOutput model.NewCredentials
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), credentialsUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.NewCredentials), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
