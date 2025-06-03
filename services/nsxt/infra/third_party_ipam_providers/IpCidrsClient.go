// Copyright (c) 2019-2025 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: IpCidrs
// Used by client-side stubs.

package third_party_ipam_providers

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	nsx_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type IpCidrsClient interface {

	// Paginated list of third party IPAM Providers Cidr info.
	//
	// @param providerInstanceIdParam (required)
	// @param ipamThirdPartyProviderCommonParamsParam (required)
	// @return com.vmware.nsx_policy.model.CidrInfoListResult
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Create(providerInstanceIdParam string, ipamThirdPartyProviderCommonParamsParam nsx_policyModel.IpamThirdPartyProviderCommonParams) (nsx_policyModel.CidrInfoListResult, error)
}

type ipCidrsClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewIpCidrsClient(connector vapiProtocolClient_.Connector) *ipCidrsClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx_policy.infra.third_party_ipam_providers.ip_cidrs")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"create": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "create"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	iIface := ipCidrsClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &iIface
}

func (iIface *ipCidrsClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := iIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (iIface *ipCidrsClient) Create(providerInstanceIdParam string, ipamThirdPartyProviderCommonParamsParam nsx_policyModel.IpamThirdPartyProviderCommonParams) (nsx_policyModel.CidrInfoListResult, error) {
	typeConverter := iIface.connector.TypeConverter()
	executionContext := iIface.connector.NewExecutionContext()
	operationRestMetaData := ipCidrsCreateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(ipCidrsCreateInputType(), typeConverter)
	sv.AddStructField("ProviderInstanceId", providerInstanceIdParam)
	sv.AddStructField("IpamThirdPartyProviderCommonParams", ipamThirdPartyProviderCommonParamsParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.CidrInfoListResult
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := iIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.third_party_ipam_providers.ip_cidrs", "create", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.CidrInfoListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), IpCidrsCreateOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.CidrInfoListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), iIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
