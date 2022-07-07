// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: ServiceDeployments
// Used by client-side stubs.

package services

import (
	"github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
)

const _ = core.SupportedByRuntimeVersion1

type ServiceDeploymentsClient interface {

	// This will deploy a particular service on a given cluster / host. Internally multiple service instance can be created during the deployment. If there are no issues in the parameters, the call returns immediately, and the service VMs will be deployed asynchronously. To get the overall status of the deployment or to get the status of individual service vm, please use the deployment status APIs.
	//
	// @param serviceIdParam (required)
	// @param serviceDeploymentParam (required)
	// @return com.vmware.nsx.model.ServiceDeployment
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Create(serviceIdParam string, serviceDeploymentParam model.ServiceDeployment) (model.ServiceDeployment, error)

	// Remove the service deployment. Will remove all the Service VMs that were created as part of this deployment. User can send optional force delete option which will force remove the deployment, but should be used only when the regular delete is not working. Regular delete will ensure proper cleanup of Service VMs and related objects. Directly calling this API without trying regular undeploy will result in unexpected results, and orphan objects.
	//
	// @param serviceIdParam (required)
	// @param serviceDeploymentIdParam (required)
	// @param forceParam Force delete the resource even if it is being used somewhere (optional, default to false)
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
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(serviceIdParam string, serviceDeploymentIdParam string) (model.ServiceDeployment, error)

	// Returns the list of deployments for the given service
	//
	// @param serviceIdParam (required)
	// @return com.vmware.nsx.model.ServiceDeploymentListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(serviceIdParam string) (model.ServiceDeploymentListResult, error)

	// This API is deprecated since only property we can change on service deployment is display name, which is used for the SVM name. Changing the name will cause the name of the deployment to go out of sync with the deployed VM.
	//
	// @param serviceIdParam (required)
	// @param serviceDeploymentIdParam (required)
	// @param serviceDeploymentParam (required)
	// @return com.vmware.nsx.model.ServiceDeployment
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(serviceIdParam string, serviceDeploymentIdParam string, serviceDeploymentParam model.ServiceDeployment) (model.ServiceDeployment, error)

	// If new deployment spec is provided, the deployment will be moved to the provided spec provided that current deployment state is either UPGRADE_FAILED or DEPLOYMENT_SUCCESSFUL If same deployment spec is provided, upgrade will be done only if current deployment state is UPGRADE_FAILED
	//
	// @param serviceIdParam (required)
	// @param serviceDeploymentIdParam (required)
	// @param deploymentSpecNameParam (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Upgrade(serviceIdParam string, serviceDeploymentIdParam string, deploymentSpecNameParam model.DeploymentSpecName) error
}

type serviceDeploymentsClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewServiceDeploymentsClient(connector client.Connector) *serviceDeploymentsClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.nsx.serviceinsertion.services.service_deployments")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"create":  core.NewMethodIdentifier(interfaceIdentifier, "create"),
		"delete":  core.NewMethodIdentifier(interfaceIdentifier, "delete"),
		"get":     core.NewMethodIdentifier(interfaceIdentifier, "get"),
		"list":    core.NewMethodIdentifier(interfaceIdentifier, "list"),
		"update":  core.NewMethodIdentifier(interfaceIdentifier, "update"),
		"upgrade": core.NewMethodIdentifier(interfaceIdentifier, "upgrade"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	sIface := serviceDeploymentsClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &sIface
}

func (sIface *serviceDeploymentsClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := sIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (sIface *serviceDeploymentsClient) Create(serviceIdParam string, serviceDeploymentParam model.ServiceDeployment) (model.ServiceDeployment, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(serviceDeploymentsCreateInputType(), typeConverter)
	sv.AddStructField("ServiceId", serviceIdParam)
	sv.AddStructField("ServiceDeployment", serviceDeploymentParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ServiceDeployment
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := serviceDeploymentsCreateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx.serviceinsertion.services.service_deployments", "create", inputDataValue, executionContext)
	var emptyOutput model.ServiceDeployment
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), serviceDeploymentsCreateOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ServiceDeployment), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *serviceDeploymentsClient) Delete(serviceIdParam string, serviceDeploymentIdParam string, forceParam *bool) error {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(serviceDeploymentsDeleteInputType(), typeConverter)
	sv.AddStructField("ServiceId", serviceIdParam)
	sv.AddStructField("ServiceDeploymentId", serviceDeploymentIdParam)
	sv.AddStructField("Force", forceParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := serviceDeploymentsDeleteRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx.serviceinsertion.services.service_deployments", "delete", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (sIface *serviceDeploymentsClient) Get(serviceIdParam string, serviceDeploymentIdParam string) (model.ServiceDeployment, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(serviceDeploymentsGetInputType(), typeConverter)
	sv.AddStructField("ServiceId", serviceIdParam)
	sv.AddStructField("ServiceDeploymentId", serviceDeploymentIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ServiceDeployment
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := serviceDeploymentsGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx.serviceinsertion.services.service_deployments", "get", inputDataValue, executionContext)
	var emptyOutput model.ServiceDeployment
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), serviceDeploymentsGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ServiceDeployment), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *serviceDeploymentsClient) List(serviceIdParam string) (model.ServiceDeploymentListResult, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(serviceDeploymentsListInputType(), typeConverter)
	sv.AddStructField("ServiceId", serviceIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ServiceDeploymentListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := serviceDeploymentsListRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx.serviceinsertion.services.service_deployments", "list", inputDataValue, executionContext)
	var emptyOutput model.ServiceDeploymentListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), serviceDeploymentsListOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ServiceDeploymentListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *serviceDeploymentsClient) Update(serviceIdParam string, serviceDeploymentIdParam string, serviceDeploymentParam model.ServiceDeployment) (model.ServiceDeployment, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(serviceDeploymentsUpdateInputType(), typeConverter)
	sv.AddStructField("ServiceId", serviceIdParam)
	sv.AddStructField("ServiceDeploymentId", serviceDeploymentIdParam)
	sv.AddStructField("ServiceDeployment", serviceDeploymentParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ServiceDeployment
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := serviceDeploymentsUpdateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx.serviceinsertion.services.service_deployments", "update", inputDataValue, executionContext)
	var emptyOutput model.ServiceDeployment
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), serviceDeploymentsUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ServiceDeployment), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *serviceDeploymentsClient) Upgrade(serviceIdParam string, serviceDeploymentIdParam string, deploymentSpecNameParam model.DeploymentSpecName) error {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(serviceDeploymentsUpgradeInputType(), typeConverter)
	sv.AddStructField("ServiceId", serviceIdParam)
	sv.AddStructField("ServiceDeploymentId", serviceDeploymentIdParam)
	sv.AddStructField("DeploymentSpecName", deploymentSpecNameParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := serviceDeploymentsUpgradeRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx.serviceinsertion.services.service_deployments", "upgrade", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}
