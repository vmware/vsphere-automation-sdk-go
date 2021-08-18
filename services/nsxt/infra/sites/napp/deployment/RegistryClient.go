// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Registry
// Used by client-side stubs.

package deployment

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/core"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/lib"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/protocol/client"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt/model"
)

const _ = core.SupportedByRuntimeVersion1

type RegistryClient interface {

	// Get helm repo and docker registry url for chart deployment.
	//
	// @param siteIdParam (required)
	// @return com.vmware.nsx_policy.model.DeploymentRegistry
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(siteIdParam string) (model.DeploymentRegistry, error)

	// In case of helm chart present in private repository and containers are hosted in private hub. You can change docker and helm url to use these private URL instead of default public one.
	//
	// @param siteIdParam (required)
	// @param deploymentRegistryParam (required)
	// @return com.vmware.nsx_policy.model.DeploymentRegistry
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(siteIdParam string, deploymentRegistryParam model.DeploymentRegistry) (model.DeploymentRegistry, error)
}

type registryClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewRegistryClient(connector client.Connector) *registryClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.nsx_policy.infra.sites.napp.deployment.registry")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"get":    core.NewMethodIdentifier(interfaceIdentifier, "get"),
		"update": core.NewMethodIdentifier(interfaceIdentifier, "update"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	rIface := registryClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &rIface
}

func (rIface *registryClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := rIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (rIface *registryClient) Get(siteIdParam string) (model.DeploymentRegistry, error) {
	typeConverter := rIface.connector.TypeConverter()
	executionContext := rIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(registryGetInputType(), typeConverter)
	sv.AddStructField("SiteId", siteIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.DeploymentRegistry
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := registryGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	rIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := rIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.sites.napp.deployment.registry", "get", inputDataValue, executionContext)
	var emptyOutput model.DeploymentRegistry
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), registryGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.DeploymentRegistry), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), rIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (rIface *registryClient) Update(siteIdParam string, deploymentRegistryParam model.DeploymentRegistry) (model.DeploymentRegistry, error) {
	typeConverter := rIface.connector.TypeConverter()
	executionContext := rIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(registryUpdateInputType(), typeConverter)
	sv.AddStructField("SiteId", siteIdParam)
	sv.AddStructField("DeploymentRegistry", deploymentRegistryParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.DeploymentRegistry
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := registryUpdateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	rIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := rIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.sites.napp.deployment.registry", "update", inputDataValue, executionContext)
	var emptyOutput model.DeploymentRegistry
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), registryUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.DeploymentRegistry), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), rIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
