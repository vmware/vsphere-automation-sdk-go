// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: ProfileStatus
// Used by client-side stubs.

package orgs

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	vmwarecloudVmc_core_orgModel "github.com/vmware/vsphere-automation-sdk-go/services/vmwarecloud/vmc_core_org/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type ProfileStatusClient interface {

	// Validate the profile of the org
	//
	// @param providerParam provider identifier
	// @param orgIdParam Unique identifier (GUID) of the organization.
	// @return OK
	//
	// @throws InvalidRequest Bad Request
	// @throws Unauthenticated Unauthorized
	// @throws Unauthorized Forbidden
	// @throws NotFound Not Found
	ValidateOrgProfile(providerParam string, orgIdParam string) (vmwarecloudVmc_core_orgModel.OrgProfileValidationResponse, error)
}

type profileStatusClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewProfileStatusClient(connector vapiProtocolClient_.Connector) *profileStatusClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.vmwarecloud.vmc_core_org.api.organization.orgs.profile_status")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"validate_org_profile": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "validate_org_profile"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	pIface := profileStatusClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &pIface
}

func (pIface *profileStatusClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := pIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (pIface *profileStatusClient) ValidateOrgProfile(providerParam string, orgIdParam string) (vmwarecloudVmc_core_orgModel.OrgProfileValidationResponse, error) {
	typeConverter := pIface.connector.TypeConverter()
	executionContext := pIface.connector.NewExecutionContext()
	operationRestMetaData := profileStatusValidateOrgProfileRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(profileStatusValidateOrgProfileInputType(), typeConverter)
	sv.AddStructField("Provider", providerParam)
	sv.AddStructField("OrgId", orgIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput vmwarecloudVmc_core_orgModel.OrgProfileValidationResponse
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := pIface.connector.GetApiProvider().Invoke("com.vmware.vmwarecloud.vmc_core_org.api.organization.orgs.profile_status", "validate_org_profile", inputDataValue, executionContext)
	var emptyOutput vmwarecloudVmc_core_orgModel.OrgProfileValidationResponse
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), ProfileStatusValidateOrgProfileOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(vmwarecloudVmc_core_orgModel.OrgProfileValidationResponse), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), pIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
