// Copyright (c) 2019-2026 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Infra
// Used by client-side stubs.

package projects

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	nsx_global_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-gm/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type InfraClient interface {

	// Read infra. Returns only the infra related properties. Inner object are not populated.
	//
	// @param orgIdParam The organization ID (required)
	// @param projectIdParam The project ID (required)
	// @param basePathParam Base path of the resource for which user wants to retrieve the hierarchy. This should be the fully qualified path for the resource. - Sample examples - base_path=/infra/domains/default/groups/Group1 base_path=/infra/domains/default/security-policies/SecurityPolicy1/rules/Rule1 (optional)
	// @param filterParam Filter string, can contain multiple or single java regular expressions separated by ';'. By default populates immediate child resources of the resource indicated by the URL. These child resources will be filtered by the type provided in the filter. It is recommended to use type_filter parameter instead of filter parameter. - Sample query string to prevent loading services and deployment zones: filter=Type-^(?!.\*?(?:Service|DeploymentZone)).\*$ - Sample query string to populate all the Group objects under Infra & Domain: filter=Type-Domain%7CGroup - Sample query string to load every policy object under Infra: filter=Type-.\* (optional)
	// @param typeFilterParam Advanced filter string in which user can directly specify the resourceTypes to be filtered. Can be used in conjunction with base_path. - Sample example of type_filter to load all groups - type_filter=Group - Sample example of multiple type_filter - type_filter=Group;SercurityPolicy;RedirectionPolicy - Sample example to load all groups in default domain using base_path in conjunction with type_filter - base_path=/infra/domains/default&type_filter=Group (optional)
	// @return com.vmware.nsx_global_policy.model.Infra
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(orgIdParam string, projectIdParam string, basePathParam *string, filterParam *string, typeFilterParam *string) (nsx_global_policyModel.Infra, error)
}

type infraClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewInfraClient(connector vapiProtocolClient_.Connector) *infraClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx_global_policy.orgs.projects.infra")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"get": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
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

func (iIface *infraClient) Get(orgIdParam string, projectIdParam string, basePathParam *string, filterParam *string, typeFilterParam *string) (nsx_global_policyModel.Infra, error) {
	typeConverter := iIface.connector.TypeConverter()
	executionContext := iIface.connector.NewExecutionContext()
	operationRestMetaData := infraGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(infraGetInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("ProjectId", projectIdParam)
	sv.AddStructField("BasePath", basePathParam)
	sv.AddStructField("Filter", filterParam)
	sv.AddStructField("TypeFilter", typeFilterParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_global_policyModel.Infra
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := iIface.connector.GetApiProvider().Invoke("com.vmware.nsx_global_policy.orgs.projects.infra", "get", inputDataValue, executionContext)
	var emptyOutput nsx_global_policyModel.Infra
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), InfraGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_global_policyModel.Infra), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), iIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
