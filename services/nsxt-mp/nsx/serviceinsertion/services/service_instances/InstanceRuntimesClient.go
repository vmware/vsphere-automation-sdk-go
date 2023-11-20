// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: InstanceRuntimes
// Used by client-side stubs.

package service_instances

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	nsxModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type InstanceRuntimesClient interface {

	// Set service VM either in or out of maintenance mode for maintenance mode, or in service or out of service for runtime state. Only one value can be set at one time.
	//
	// @param serviceIdParam (required)
	// @param serviceInstanceIdParam (required)
	// @param instanceRuntimeIdParam (required)
	// @param actionParam (optional)
	// @param unhealthyReasonParam Reason for the unhealthy state (optional)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Create(serviceIdParam string, serviceInstanceIdParam string, instanceRuntimeIdParam string, actionParam *string, unhealthyReasonParam *string) error

	// Undeploy one service VM as standalone or two service VMs as HA. Associated deployment information and instance runtime will also be deleted once service VMs have been un-deployed successfully.
	//  This API has been deprecated, please use below Policy API
	//  DELETE /policy/api/v1/infra/tier-0s/<tier-0-id>/locale-services/<locale-service-id>/service-instances/<service-instance-id> DELETE /policy/api/v1/infra/tier-1s/<tier-1-id>/locale-services/<locale-service-id>/service-instances/<service-instance-id>
	//
	// Deprecated: This API element is deprecated.
	//
	// @param serviceIdParam (required)
	// @param serviceInstanceIdParam (required)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Delete(serviceIdParam string, serviceInstanceIdParam string) error

	// Deploys one service VM as standalone, or two service VMs as HA where one VM is active and another one is standby. During the deployment of service VMs, service will be set up based on deployment events using callbacks.
	//  This API has been deprecated, please use below Policy API
	//  PUT /policy/api/v1/infra/tier-0s/<tier-0-id>/locale-services/<locale-service-id>/service-instances/<service-instance-id> PUT /policy/api/v1/infra/tier-1s/<tier-1-id>/locale-services/<locale-service-id>/service-instances/<service-instance-id>
	//
	// Deprecated: This API element is deprecated.
	//
	// @param serviceIdParam (required)
	// @param serviceInstanceIdParam (required)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Deploy(serviceIdParam string, serviceInstanceIdParam string) error

	// Returns list of instance runtimes of service VMs being deployed for a given service instance id
	//
	// @param serviceIdParam (required)
	// @param serviceInstanceIdParam (required)
	// @return com.vmware.nsx.model.InstanceRuntimeListResult
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(serviceIdParam string, serviceInstanceIdParam string) (nsxModel.InstanceRuntimeListResult, error)

	// Upgrade service VMs using newer version of OVF. Upgrade is a 2 step process. Update the 'deployment_spec_name' in the ServiceInstance to the new DeploymentSpec to which the service VMs will be upgraded, folowed by this 'upgrade' api. In case of HA, the stand-by service VM will be upgrade first. Once it has been upgraded, it switches to be the Active one and then the other VM will be upgrade.
	//  This API has been deprecated, please use below Policy API
	//  PATCH /policy/api/v1/infra/tier-0s/<tier-0-id>/locale-services/<locale-service-id>/service-instances/<service-instance-id> PATCH /policy/api/v1/infra/tier-1s/<tier-1-id>/locale-services/<locale-service-id>/service-instances/<service-instance-id>
	//
	// Deprecated: This API element is deprecated.
	//
	// @param serviceIdParam (required)
	// @param serviceInstanceIdParam (required)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Upgrade(serviceIdParam string, serviceInstanceIdParam string) error
}

type instanceRuntimesClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewInstanceRuntimesClient(connector vapiProtocolClient_.Connector) *instanceRuntimesClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx.serviceinsertion.services.service_instances.instance_runtimes")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"create":  vapiCore_.NewMethodIdentifier(interfaceIdentifier, "create"),
		"delete":  vapiCore_.NewMethodIdentifier(interfaceIdentifier, "delete"),
		"deploy":  vapiCore_.NewMethodIdentifier(interfaceIdentifier, "deploy"),
		"list":    vapiCore_.NewMethodIdentifier(interfaceIdentifier, "list"),
		"upgrade": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "upgrade"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	iIface := instanceRuntimesClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &iIface
}

func (iIface *instanceRuntimesClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := iIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (iIface *instanceRuntimesClient) Create(serviceIdParam string, serviceInstanceIdParam string, instanceRuntimeIdParam string, actionParam *string, unhealthyReasonParam *string) error {
	typeConverter := iIface.connector.TypeConverter()
	executionContext := iIface.connector.NewExecutionContext()
	operationRestMetaData := instanceRuntimesCreateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(instanceRuntimesCreateInputType(), typeConverter)
	sv.AddStructField("ServiceId", serviceIdParam)
	sv.AddStructField("ServiceInstanceId", serviceInstanceIdParam)
	sv.AddStructField("InstanceRuntimeId", instanceRuntimeIdParam)
	sv.AddStructField("Action", actionParam)
	sv.AddStructField("UnhealthyReason", unhealthyReasonParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := iIface.connector.GetApiProvider().Invoke("com.vmware.nsx.serviceinsertion.services.service_instances.instance_runtimes", "create", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), iIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (iIface *instanceRuntimesClient) Delete(serviceIdParam string, serviceInstanceIdParam string) error {
	typeConverter := iIface.connector.TypeConverter()
	executionContext := iIface.connector.NewExecutionContext()
	operationRestMetaData := instanceRuntimesDeleteRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(instanceRuntimesDeleteInputType(), typeConverter)
	sv.AddStructField("ServiceId", serviceIdParam)
	sv.AddStructField("ServiceInstanceId", serviceInstanceIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := iIface.connector.GetApiProvider().Invoke("com.vmware.nsx.serviceinsertion.services.service_instances.instance_runtimes", "delete", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), iIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (iIface *instanceRuntimesClient) Deploy(serviceIdParam string, serviceInstanceIdParam string) error {
	typeConverter := iIface.connector.TypeConverter()
	executionContext := iIface.connector.NewExecutionContext()
	operationRestMetaData := instanceRuntimesDeployRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(instanceRuntimesDeployInputType(), typeConverter)
	sv.AddStructField("ServiceId", serviceIdParam)
	sv.AddStructField("ServiceInstanceId", serviceInstanceIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := iIface.connector.GetApiProvider().Invoke("com.vmware.nsx.serviceinsertion.services.service_instances.instance_runtimes", "deploy", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), iIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (iIface *instanceRuntimesClient) List(serviceIdParam string, serviceInstanceIdParam string) (nsxModel.InstanceRuntimeListResult, error) {
	typeConverter := iIface.connector.TypeConverter()
	executionContext := iIface.connector.NewExecutionContext()
	operationRestMetaData := instanceRuntimesListRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(instanceRuntimesListInputType(), typeConverter)
	sv.AddStructField("ServiceId", serviceIdParam)
	sv.AddStructField("ServiceInstanceId", serviceInstanceIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.InstanceRuntimeListResult
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := iIface.connector.GetApiProvider().Invoke("com.vmware.nsx.serviceinsertion.services.service_instances.instance_runtimes", "list", inputDataValue, executionContext)
	var emptyOutput nsxModel.InstanceRuntimeListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), InstanceRuntimesListOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.InstanceRuntimeListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), iIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (iIface *instanceRuntimesClient) Upgrade(serviceIdParam string, serviceInstanceIdParam string) error {
	typeConverter := iIface.connector.TypeConverter()
	executionContext := iIface.connector.NewExecutionContext()
	operationRestMetaData := instanceRuntimesUpgradeRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(instanceRuntimesUpgradeInputType(), typeConverter)
	sv.AddStructField("ServiceId", serviceIdParam)
	sv.AddStructField("ServiceInstanceId", serviceInstanceIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := iIface.connector.GetApiProvider().Invoke("com.vmware.nsx.serviceinsertion.services.service_instances.instance_runtimes", "upgrade", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), iIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}
