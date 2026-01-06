// Copyright (c) 2019-2026 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Infra
// Used by client-side stubs.

package nsx_policy

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	nsx_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type InfraClient interface {

	// Read infra. Returns only the infra related properties. Inner object are not populated.
	//
	// @param basePathParam Base path of the resource for which user wants to retrieve the hierarchy. This should be the fully qualified path for the resource. - Sample examples - base_path=/infra/domains/default/groups/Group1 base_path=/infra/domains/default/security-policies/SecurityPolicy1/rules/Rule1 (optional)
	// @param filterParam Filter string, can contain multiple or single java regular expressions separated by ';'. By default populates immediate child resources of the resource indicated by the URL. These child resources will be filtered by the type provided in the filter. It is recommended to use type_filter parameter instead of filter parameter. - Sample query string to prevent loading services and deployment zones: filter=Type-^(?!.\*?(?:Service|DeploymentZone)).\*$ - Sample query string to populate all the Group objects under Infra & Domain: filter=Type-Domain%7CGroup - Sample query string to load every policy object under Infra: filter=Type-.\* (optional)
	// @param typeFilterParam Advanced filter string in which user can directly specify the resourceTypes to be filtered. Can be used in conjunction with base_path. - Sample example of type_filter to load all groups - type_filter=Group - Sample example of multiple type_filter - type_filter=Group;SercurityPolicy;RedirectionPolicy - Sample example to load all groups in default domain using base_path in conjunction with type_filter - base_path=/infra/domains/default&type_filter=Group (optional)
	// @return com.vmware.nsx_policy.model.Infra
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(basePathParam *string, filterParam *string, typeFilterParam *string) (nsx_policyModel.Infra, error)

	// Patch API at infra level can be used in two flavours 1. Like a regular API to update Infra object 2. Hierarchical API: To create/update/delete entire or part of intent hierarchy Hierarchical API: Provides users a way to create entire or part of intent in single API invocation. Input is expressed in a tree format. Each node in tree can have multiple children of different types. System will resolve the dependencies of nodes within the intent tree and will create the model. Children for any node can be specified using ChildResourceReference or ChildPolicyConfigResource. If a resource is specified using ChildResourceReference then it will not be updated only its children will be updated. If Object is specified using ChildPolicyConfigResource, object along with its children will be updated. Hierarchical API can also be used to delete any sub-branch of entire tree. Hierarchical API supports up to 5000 intent creation on LM and 1000 on GM.
	//
	// @param infraParam (required)
	// @param enforceRevisionCheckParam If this is set to true, each child object in the request needs to have _revision property set correctly. System will honor the revision numbers while updating the resources. (optional, default to false)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Patch(infraParam nsx_policyModel.Infra, enforceRevisionCheckParam *bool) error

	// Updates only the single infra object. This does not allow hierarchical updates of entities.
	//
	// @param infraParam (required)
	// @return com.vmware.nsx_policy.model.Infra
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(infraParam nsx_policyModel.Infra) (nsx_policyModel.Infra, error)
}

type infraClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewInfraClient(connector vapiProtocolClient_.Connector) *infraClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx_policy.infra")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"get":    vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
		"patch":  vapiCore_.NewMethodIdentifier(interfaceIdentifier, "patch"),
		"update": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "update"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	iIface := infraClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &iIface
}

func (iIface *infraClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := iIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (iIface *infraClient) Get(basePathParam *string, filterParam *string, typeFilterParam *string) (nsx_policyModel.Infra, error) {
	typeConverter := iIface.connector.TypeConverter()
	executionContext := iIface.connector.NewExecutionContext()
	operationRestMetaData := infraGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(infraGetInputType(), typeConverter)
	sv.AddStructField("BasePath", basePathParam)
	sv.AddStructField("Filter", filterParam)
	sv.AddStructField("TypeFilter", typeFilterParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.Infra
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := iIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra", "get", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.Infra
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), InfraGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.Infra), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), iIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (iIface *infraClient) Patch(infraParam nsx_policyModel.Infra, enforceRevisionCheckParam *bool) error {
	typeConverter := iIface.connector.TypeConverter()
	executionContext := iIface.connector.NewExecutionContext()
	operationRestMetaData := infraPatchRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(infraPatchInputType(), typeConverter)
	sv.AddStructField("Infra", infraParam)
	sv.AddStructField("EnforceRevisionCheck", enforceRevisionCheckParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := iIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra", "patch", inputDataValue, executionContext)
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

func (iIface *infraClient) Update(infraParam nsx_policyModel.Infra) (nsx_policyModel.Infra, error) {
	typeConverter := iIface.connector.TypeConverter()
	executionContext := iIface.connector.NewExecutionContext()
	operationRestMetaData := infraUpdateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(infraUpdateInputType(), typeConverter)
	sv.AddStructField("Infra", infraParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.Infra
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := iIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra", "update", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.Infra
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), InfraUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.Infra), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), iIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
