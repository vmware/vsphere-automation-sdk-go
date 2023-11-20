// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: FeatureSummary
// Used by client-side stubs.

package onboarding

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	nsx_global_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-gm/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type FeatureSummaryClient interface {

	// Get consolidated list of conflicting entities summary for each supported feature for a site with an example.
	//
	// @param siteIdParam (required)
	// @return com.vmware.nsx_global_policy.model.ConflictingEntityListResponse
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(siteIdParam string) (nsx_global_policyModel.ConflictingEntityListResponse, error)
}

type featureSummaryClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewFeatureSummaryClient(connector vapiProtocolClient_.Connector) *featureSummaryClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx_global_policy.global_infra.sites.onboarding.feature_summary")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"get": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	fIface := featureSummaryClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &fIface
}

func (fIface *featureSummaryClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := fIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (fIface *featureSummaryClient) Get(siteIdParam string) (nsx_global_policyModel.ConflictingEntityListResponse, error) {
	typeConverter := fIface.connector.TypeConverter()
	executionContext := fIface.connector.NewExecutionContext()
	operationRestMetaData := featureSummaryGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(featureSummaryGetInputType(), typeConverter)
	sv.AddStructField("SiteId", siteIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_global_policyModel.ConflictingEntityListResponse
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := fIface.connector.GetApiProvider().Invoke("com.vmware.nsx_global_policy.global_infra.sites.onboarding.feature_summary", "get", inputDataValue, executionContext)
	var emptyOutput nsx_global_policyModel.ConflictingEntityListResponse
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), FeatureSummaryGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_global_policyModel.ConflictingEntityListResponse), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), fIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
