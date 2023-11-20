// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Onboarding
// Used by client-side stubs.

package sites

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	nsx_global_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-gm/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type OnboardingClient interface {

	// Verifies and reports conflicting onboarding feature for a site. The response will contain first conflicting feature for the site configuration compared to corresponding global manager configuration.
	//
	// @param siteIdParam (required)
	// @param configOnboardingConflictRequestParam (required)
	// @return com.vmware.nsx_global_policy.model.ConfigOnboardingConflictStatus
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Checkconflict(siteIdParam string, configOnboardingConflictRequestParam nsx_global_policyModel.ConfigOnboardingConflictRequest) (nsx_global_policyModel.ConfigOnboardingConflictStatus, error)

	// Initiate config on-boarding of a Site. The entire on-boarding is async workflow controlled by API.
	//
	// @param siteIdParam (required)
	// @param configOnboardingRequestParam (required)
	// @return com.vmware.nsx_global_policy.model.ConfigOnboardingStatus
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Startonboarding(siteIdParam string, configOnboardingRequestParam nsx_global_policyModel.ConfigOnboardingRequest) (nsx_global_policyModel.ConfigOnboardingStatus, error)
}

type onboardingClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewOnboardingClient(connector vapiProtocolClient_.Connector) *onboardingClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx_global_policy.global_infra.sites.onboarding")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"checkconflict":   vapiCore_.NewMethodIdentifier(interfaceIdentifier, "checkconflict"),
		"startonboarding": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "startonboarding"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	oIface := onboardingClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &oIface
}

func (oIface *onboardingClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := oIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (oIface *onboardingClient) Checkconflict(siteIdParam string, configOnboardingConflictRequestParam nsx_global_policyModel.ConfigOnboardingConflictRequest) (nsx_global_policyModel.ConfigOnboardingConflictStatus, error) {
	typeConverter := oIface.connector.TypeConverter()
	executionContext := oIface.connector.NewExecutionContext()
	operationRestMetaData := onboardingCheckconflictRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(onboardingCheckconflictInputType(), typeConverter)
	sv.AddStructField("SiteId", siteIdParam)
	sv.AddStructField("ConfigOnboardingConflictRequest", configOnboardingConflictRequestParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_global_policyModel.ConfigOnboardingConflictStatus
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := oIface.connector.GetApiProvider().Invoke("com.vmware.nsx_global_policy.global_infra.sites.onboarding", "checkconflict", inputDataValue, executionContext)
	var emptyOutput nsx_global_policyModel.ConfigOnboardingConflictStatus
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), OnboardingCheckconflictOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_global_policyModel.ConfigOnboardingConflictStatus), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), oIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (oIface *onboardingClient) Startonboarding(siteIdParam string, configOnboardingRequestParam nsx_global_policyModel.ConfigOnboardingRequest) (nsx_global_policyModel.ConfigOnboardingStatus, error) {
	typeConverter := oIface.connector.TypeConverter()
	executionContext := oIface.connector.NewExecutionContext()
	operationRestMetaData := onboardingStartonboardingRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(onboardingStartonboardingInputType(), typeConverter)
	sv.AddStructField("SiteId", siteIdParam)
	sv.AddStructField("ConfigOnboardingRequest", configOnboardingRequestParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_global_policyModel.ConfigOnboardingStatus
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := oIface.connector.GetApiProvider().Invoke("com.vmware.nsx_global_policy.global_infra.sites.onboarding", "startonboarding", inputDataValue, executionContext)
	var emptyOutput nsx_global_policyModel.ConfigOnboardingStatus
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), OnboardingStartonboardingOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_global_policyModel.ConfigOnboardingStatus), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), oIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
