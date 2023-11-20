// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: WithCertificate
// Used by client-side stubs.

package principal_identities

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	nsxModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type WithCertificateClient interface {

	// Create a principal identity with a new, unused, certificate. The combination name and node_id needs to be unique across token-based and certificate-based principal identities.
	//
	// @param principalIdentityWithCertificateParam (required)
	// @return com.vmware.nsx.model.PrincipalIdentity
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Create(principalIdentityWithCertificateParam nsxModel.PrincipalIdentityWithCertificate) (nsxModel.PrincipalIdentity, error)
}

type withCertificateClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewWithCertificateClient(connector vapiProtocolClient_.Connector) *withCertificateClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx.trust_management.principal_identities.with_certificate")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"create": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "create"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	wIface := withCertificateClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &wIface
}

func (wIface *withCertificateClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := wIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (wIface *withCertificateClient) Create(principalIdentityWithCertificateParam nsxModel.PrincipalIdentityWithCertificate) (nsxModel.PrincipalIdentity, error) {
	typeConverter := wIface.connector.TypeConverter()
	executionContext := wIface.connector.NewExecutionContext()
	operationRestMetaData := withCertificateCreateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(withCertificateCreateInputType(), typeConverter)
	sv.AddStructField("PrincipalIdentityWithCertificate", principalIdentityWithCertificateParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.PrincipalIdentity
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := wIface.connector.GetApiProvider().Invoke("com.vmware.nsx.trust_management.principal_identities.with_certificate", "create", inputDataValue, executionContext)
	var emptyOutput nsxModel.PrincipalIdentity
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), WithCertificateCreateOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.PrincipalIdentity), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), wIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
