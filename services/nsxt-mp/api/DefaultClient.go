// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Default
// Used by client-side stubs.

package api

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/core"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/lib"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/protocol/client"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt-mp/model"
)

const _ = core.SupportedByRuntimeVersion1

type DefaultClient interface {

	// Look up the Certificate Profile matching the service-type and apply the certificate. When the Certificate Profile has cluster_certificate=false, the node_id parameter is required to designate the node where the certificate needs to be applied.
	//
	// @param certIdParam ID of certificate to apply (required)
	// @param serviceTypeParam Supported service types, that are using certificates. (required)
	// @param nodeIdParam Node Id (optional)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ApplyCertificateApplyCertificate(certIdParam string, serviceTypeParam string, nodeIdParam *string) error

	// Set a certificate that has been imported to be either the principal identity certificate for the local cluster with either GM or LM service type. Currently, the service type specified must match the current service type of the local cluster.
	//
	// @param setPrincipalIdentityCertificateForFederationRequestParam (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	SetPrincipalIdentityCertificateForFederationSetPiCertificateForFederation(setPrincipalIdentityCertificateForFederationRequestParam model.SetPrincipalIdentityCertificateForFederationRequest) error

	// Checks whether certificate is valid. When the certificate contains a chain, the full chain is validated. The usage parameter can be SERVER (default) or CLIENT. This indicates whether the certificate needs to be validated as a server-auth or a client-auth certificate.
	//
	// @param certIdParam ID of certificate to validate (required)
	// @param usageParam Usage Type of the Certificate, SERVER or CLIENT. Default is SERVER (optional)
	// @return com.vmware.model.CertificateCheckingStatus
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ValidateCertificateValidate(certIdParam string, usageParam *string) (model.CertificateCheckingStatus, error)
}

type defaultClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewDefaultClient(connector client.Connector) *defaultClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.api.default")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"apply_certificate_apply_certificate":                                                 core.NewMethodIdentifier(interfaceIdentifier, "apply_certificate_apply_certificate"),
		"set_principal_identity_certificate_for_federation_set_pi_certificate_for_federation": core.NewMethodIdentifier(interfaceIdentifier, "set_principal_identity_certificate_for_federation_set_pi_certificate_for_federation"),
		"validate_certificate_validate":                                                       core.NewMethodIdentifier(interfaceIdentifier, "validate_certificate_validate"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	dIface := defaultClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &dIface
}

func (dIface *defaultClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := dIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (dIface *defaultClient) ApplyCertificateApplyCertificate(certIdParam string, serviceTypeParam string, nodeIdParam *string) error {
	typeConverter := dIface.connector.TypeConverter()
	executionContext := dIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(defaultApplyCertificateApplyCertificateInputType(), typeConverter)
	sv.AddStructField("CertId", certIdParam)
	sv.AddStructField("ServiceType", serviceTypeParam)
	sv.AddStructField("NodeId", nodeIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := defaultApplyCertificateApplyCertificateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	dIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := dIface.connector.GetApiProvider().Invoke("com.vmware.api.default", "apply_certificate_apply_certificate", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), dIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (dIface *defaultClient) SetPrincipalIdentityCertificateForFederationSetPiCertificateForFederation(setPrincipalIdentityCertificateForFederationRequestParam model.SetPrincipalIdentityCertificateForFederationRequest) error {
	typeConverter := dIface.connector.TypeConverter()
	executionContext := dIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(defaultSetPrincipalIdentityCertificateForFederationSetPiCertificateForFederationInputType(), typeConverter)
	sv.AddStructField("SetPrincipalIdentityCertificateForFederationRequest", setPrincipalIdentityCertificateForFederationRequestParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := defaultSetPrincipalIdentityCertificateForFederationSetPiCertificateForFederationRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	dIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := dIface.connector.GetApiProvider().Invoke("com.vmware.api.default", "set_principal_identity_certificate_for_federation_set_pi_certificate_for_federation", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), dIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (dIface *defaultClient) ValidateCertificateValidate(certIdParam string, usageParam *string) (model.CertificateCheckingStatus, error) {
	typeConverter := dIface.connector.TypeConverter()
	executionContext := dIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(defaultValidateCertificateValidateInputType(), typeConverter)
	sv.AddStructField("CertId", certIdParam)
	sv.AddStructField("Usage", usageParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.CertificateCheckingStatus
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := defaultValidateCertificateValidateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	dIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := dIface.connector.GetApiProvider().Invoke("com.vmware.api.default", "validate_certificate_validate", inputDataValue, executionContext)
	var emptyOutput model.CertificateCheckingStatus
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), defaultValidateCertificateValidateOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.CertificateCheckingStatus), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), dIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
