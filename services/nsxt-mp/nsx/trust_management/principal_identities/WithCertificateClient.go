// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: WithCertificate
// Used by client-side stubs.

package principal_identities

import (
	"github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
)

const _ = core.SupportedByRuntimeVersion1

type WithCertificateClient interface {

	// Create a principal identity with a new, unused, certificate. The combination name and node_id needs to be unique across token-based and certificate-based principal identities.
	//
	// @param principalIdentityWithCertificateParam (required)
	// @return com.vmware.nsx.model.PrincipalIdentity
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Create(principalIdentityWithCertificateParam model.PrincipalIdentityWithCertificate) (model.PrincipalIdentity, error)
}

type withCertificateClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewWithCertificateClient(connector client.Connector) *withCertificateClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.nsx.trust_management.principal_identities.with_certificate")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"create": core.NewMethodIdentifier(interfaceIdentifier, "create"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	wIface := withCertificateClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &wIface
}

func (wIface *withCertificateClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := wIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (wIface *withCertificateClient) Create(principalIdentityWithCertificateParam model.PrincipalIdentityWithCertificate) (model.PrincipalIdentity, error) {
	typeConverter := wIface.connector.TypeConverter()
	executionContext := wIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(withCertificateCreateInputType(), typeConverter)
	sv.AddStructField("PrincipalIdentityWithCertificate", principalIdentityWithCertificateParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.PrincipalIdentity
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := withCertificateCreateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	wIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := wIface.connector.GetApiProvider().Invoke("com.vmware.nsx.trust_management.principal_identities.with_certificate", "create", inputDataValue, executionContext)
	var emptyOutput model.PrincipalIdentity
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), withCertificateCreateOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.PrincipalIdentity), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), wIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
