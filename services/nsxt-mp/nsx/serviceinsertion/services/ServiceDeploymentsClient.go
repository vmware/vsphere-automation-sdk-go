// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: ServiceDeployments
// Used by client-side stubs.

package services

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	nsxModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type ServiceDeploymentsClient interface {

	// This will deploy a particular service on a given cluster / host. Internally multiple service instance can be created during the deployment. If there are no issues in the parameters, the call returns immediately, and the service VMs will be deployed asynchronously. To get the overall status of the deployment or to get the status of individual service vm, please use the deployment status APIs.
	//
	// @param serviceIdParam (required)
	// @param serviceDeploymentParam (required)
	// @return com.vmware.nsx.model.ServiceDeployment
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Create(serviceIdParam string, serviceDeploymentParam nsxModel.ServiceDeployment) (nsxModel.ServiceDeployment, error)

	// Remove the service deployment. Will remove all the Service VMs that were created as part of this deployment. User can send optional force delete option which will force remove the deployment, but should be used only when the regular delete is not working. Regular delete will ensure proper cleanup of Service VMs and related objects. Directly calling this API without trying regular undeploy will result in unexpected results, and orphan objects.
	//
	// @param serviceIdParam (required)
	// @param serviceDeploymentIdParam (required)
	// @param forceParam Force delete the resource even if it is being used somewhere (optional, default to false)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Delete(serviceIdParam string, serviceDeploymentIdParam string, forceParam *bool) error

	// Returns detail of service deployment.
	//
	// @param serviceIdParam (required)
	// @param serviceDeploymentIdParam (required)
	// @return com.vmware.nsx.model.ServiceDeployment
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(serviceIdParam string, serviceDeploymentIdParam string) (nsxModel.ServiceDeployment, error)

	// Returns the list of deployments for the given service
	//
	// @param serviceIdParam (required)
	// @return com.vmware.nsx.model.ServiceDeploymentListResult
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(serviceIdParam string) (nsxModel.ServiceDeploymentListResult, error)

	// Update an existing Service Deployment.
	//  This API is deprecated since only property we can change on service deployment is display name, which is used for the SVM name. Changing the name will cause the name of the deployment to go out of sync with the deployed VM.
	//
	// Deprecated: This API element is deprecated.
	//
	// @param serviceIdParam (required)
	// @param serviceDeploymentIdParam (required)
	// @param serviceDeploymentParam (required)
	// @return com.vmware.nsx.model.ServiceDeployment
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(serviceIdParam string, serviceDeploymentIdParam string, serviceDeploymentParam nsxModel.ServiceDeployment) (nsxModel.ServiceDeployment, error)

	// If new deployment spec is provided, the deployment will be moved to the provided spec provided that current deployment state is either UPGRADE_FAILED or DEPLOYMENT_SUCCESSFUL If same deployment spec is provided, upgrade will be done only if current deployment state is UPGRADE_FAILED If the exisiting data store is upgraded or changed, there is an optional parameter to supply the new storage Id during upgrade which would perform the deployment upgrade on the new data store
	//
	// @param serviceIdParam (required)
	// @param serviceDeploymentIdParam (required)
	// @param deploymentSpecNameParam (required)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Upgrade(serviceIdParam string, serviceDeploymentIdParam string, deploymentSpecNameParam nsxModel.DeploymentSpecName) error
}

type serviceDeploymentsClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewServiceDeploymentsClient(connector vapiProtocolClient_.Connector) *serviceDeploymentsClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx.serviceinsertion.services.service_deployments")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"create":  vapiCore_.NewMethodIdentifier(interfaceIdentifier, "create"),
		"delete":  vapiCore_.NewMethodIdentifier(interfaceIdentifier, "delete"),
		"get":     vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
		"list":    vapiCore_.NewMethodIdentifier(interfaceIdentifier, "list"),
		"update":  vapiCore_.NewMethodIdentifier(interfaceIdentifier, "update"),
		"upgrade": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "upgrade"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	sIface := serviceDeploymentsClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &sIface
}

func (sIface *serviceDeploymentsClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := sIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (sIface *serviceDeploymentsClient) Create(serviceIdParam string, serviceDeploymentParam nsxModel.ServiceDeployment) (nsxModel.ServiceDeployment, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	operationRestMetaData := serviceDeploymentsCreateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(serviceDeploymentsCreateInputType(), typeConverter)
	sv.AddStructField("ServiceId", serviceIdParam)
	sv.AddStructField("ServiceDeployment", serviceDeploymentParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.ServiceDeployment
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx.serviceinsertion.services.service_deployments", "create", inputDataValue, executionContext)
	var emptyOutput nsxModel.ServiceDeployment
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), ServiceDeploymentsCreateOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.ServiceDeployment), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *serviceDeploymentsClient) Delete(serviceIdParam string, serviceDeploymentIdParam string, forceParam *bool) error {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	operationRestMetaData := serviceDeploymentsDeleteRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(serviceDeploymentsDeleteInputType(), typeConverter)
	sv.AddStructField("ServiceId", serviceIdParam)
	sv.AddStructField("ServiceDeploymentId", serviceDeploymentIdParam)
	sv.AddStructField("Force", forceParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx.serviceinsertion.services.service_deployments", "delete", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (sIface *serviceDeploymentsClient) Get(serviceIdParam string, serviceDeploymentIdParam string) (nsxModel.ServiceDeployment, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	operationRestMetaData := serviceDeploymentsGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(serviceDeploymentsGetInputType(), typeConverter)
	sv.AddStructField("ServiceId", serviceIdParam)
	sv.AddStructField("ServiceDeploymentId", serviceDeploymentIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.ServiceDeployment
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx.serviceinsertion.services.service_deployments", "get", inputDataValue, executionContext)
	var emptyOutput nsxModel.ServiceDeployment
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), ServiceDeploymentsGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.ServiceDeployment), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *serviceDeploymentsClient) List(serviceIdParam string) (nsxModel.ServiceDeploymentListResult, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	operationRestMetaData := serviceDeploymentsListRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(serviceDeploymentsListInputType(), typeConverter)
	sv.AddStructField("ServiceId", serviceIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.ServiceDeploymentListResult
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx.serviceinsertion.services.service_deployments", "list", inputDataValue, executionContext)
	var emptyOutput nsxModel.ServiceDeploymentListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), ServiceDeploymentsListOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.ServiceDeploymentListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *serviceDeploymentsClient) Update(serviceIdParam string, serviceDeploymentIdParam string, serviceDeploymentParam nsxModel.ServiceDeployment) (nsxModel.ServiceDeployment, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	operationRestMetaData := serviceDeploymentsUpdateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(serviceDeploymentsUpdateInputType(), typeConverter)
	sv.AddStructField("ServiceId", serviceIdParam)
	sv.AddStructField("ServiceDeploymentId", serviceDeploymentIdParam)
	sv.AddStructField("ServiceDeployment", serviceDeploymentParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.ServiceDeployment
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx.serviceinsertion.services.service_deployments", "update", inputDataValue, executionContext)
	var emptyOutput nsxModel.ServiceDeployment
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), ServiceDeploymentsUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.ServiceDeployment), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *serviceDeploymentsClient) Upgrade(serviceIdParam string, serviceDeploymentIdParam string, deploymentSpecNameParam nsxModel.DeploymentSpecName) error {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	operationRestMetaData := serviceDeploymentsUpgradeRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(serviceDeploymentsUpgradeInputType(), typeConverter)
	sv.AddStructField("ServiceId", serviceIdParam)
	sv.AddStructField("ServiceDeploymentId", serviceDeploymentIdParam)
	sv.AddStructField("DeploymentSpecName", deploymentSpecNameParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx.serviceinsertion.services.service_deployments", "upgrade", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}
