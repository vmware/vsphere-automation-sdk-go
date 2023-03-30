// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: TlsInspectionConfigProfileBindings
// Used by client-side stubs.

package tier_1s

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	nsx_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type TlsInspectionConfigProfileBindingsClient interface {

	// API will get TLS Config Profile Binding Map for Tier-1 Logical Router.
	//
	// @param tier1IdParam (required)
	// @param tlsInspectionConfigProfileBindingIdParam (required)
	// @return com.vmware.nsx_policy.model.TlsConfigProfileBindingMap
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(tier1IdParam string, tlsInspectionConfigProfileBindingIdParam string) (nsx_policyModel.TlsConfigProfileBindingMap, error)
}

type tlsInspectionConfigProfileBindingsClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewTlsInspectionConfigProfileBindingsClient(connector vapiProtocolClient_.Connector) *tlsInspectionConfigProfileBindingsClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx_policy.global_infra.tier_1s.tls_inspection_config_profile_bindings")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"get": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	tIface := tlsInspectionConfigProfileBindingsClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &tIface
}

func (tIface *tlsInspectionConfigProfileBindingsClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := tIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (tIface *tlsInspectionConfigProfileBindingsClient) Get(tier1IdParam string, tlsInspectionConfigProfileBindingIdParam string) (nsx_policyModel.TlsConfigProfileBindingMap, error) {
	typeConverter := tIface.connector.TypeConverter()
	executionContext := tIface.connector.NewExecutionContext()
	operationRestMetaData := tlsInspectionConfigProfileBindingsGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(tlsInspectionConfigProfileBindingsGetInputType(), typeConverter)
	sv.AddStructField("Tier1Id", tier1IdParam)
	sv.AddStructField("TlsInspectionConfigProfileBindingId", tlsInspectionConfigProfileBindingIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.TlsConfigProfileBindingMap
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := tIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.global_infra.tier_1s.tls_inspection_config_profile_bindings", "get", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.TlsConfigProfileBindingMap
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), TlsInspectionConfigProfileBindingsGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.TlsConfigProfileBindingMap), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
