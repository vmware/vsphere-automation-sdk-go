// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Certificates
// Used by client-side stubs.

package trust_management

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	nsxModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type CertificatesClient interface {

	// Look up the Certificate Profile matching the service-type and apply the certificate. When the Certificate Profile has cluster_certificate=false, the node_id parameter is required to designate the node where the certificate needs to be applied.
	//
	// @param certIdParam ID of certificate to apply (required)
	// @param serviceTypeParam Supported service types, that are using certificates. (required)
	// @param nodeIdParam Node Id (optional)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Applycertificate(certIdParam string, serviceTypeParam string, nodeIdParam *string) error

	// Removes the specified certificate. The private key associated with the certificate is also deleted.
	//
	// @param certIdParam ID of certificate to delete (required)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Delete(certIdParam string) error

	// Attempt to connect to an TLS service endpoint and retrieve the server certificate chain it presents.
	//
	// @param tlsServiceEndpointParam (required)
	// @return com.vmware.nsx.model.PeerCertificateChain
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Fetchpeercertificatechain(tlsServiceEndpointParam nsxModel.TlsServiceEndpoint) (nsxModel.PeerCertificateChain, error)

	// Returns information for the specified certificate ID, including the certificate's UUID; resource_type (for example, certificate_self_signed, certificate_ca, or certificate_signed); pem_encoded data; and history of the certificate (who created or modified it and when). For additional information, include the ?details=true modifier at the end of the request URI.
	//
	// @param certIdParam ID of certificate to read (required)
	// @param detailsParam whether to expand the pem data and show all its details (optional, default to false)
	// @return com.vmware.nsx.model.Certificate
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(certIdParam string, detailsParam *bool) (nsxModel.Certificate, error)

	// Adds a new private-public certificate or a chain of certificates (CAs) and, optionally, a private key that can be applied to one of the user-facing components (appliance management or edge). The certificate and the key should be stored in PEM format. If no private key is provided, the certificate is used as a client certificate in the trust store. A private key can be uploaded for a CA certificate only if the \"purpose\" parameter is set to \"signing-ca\".
	//
	// @param trustObjectDataParam (required)
	// @return com.vmware.nsx.model.CertificateList
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Importcertificate(trustObjectDataParam nsxModel.TrustObjectData) (nsxModel.CertificateList, error)

	// Add a CA certificate as a trust anchor
	//
	// @param aliasParam Alias under which to store the trusted CA in the trust-store (required)
	// @param trustObjectDataParam (required)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Importtrustedca(aliasParam string, trustObjectDataParam nsxModel.TrustObjectData) error

	// Returns all certificate information viewable by the user, including each certificate's UUID; resource_type (for example, certificate_self_signed, certificate_ca, or certificate_signed); pem_encoded data; and history of the certificate (who created or modified it and when). For additional information, include the ?details=true modifier at the end of the request URI.
	//
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param detailsParam whether to expand the pem data and show all its details (optional, default to false)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param nodeIdParam Node ID of certificate to return (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @param type_Param Type of certificate to return (optional)
	// @return com.vmware.nsx.model.CertificateList
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(cursorParam *string, detailsParam *bool, includedFieldsParam *string, nodeIdParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string, type_Param *string) (nsxModel.CertificateList, error)

	// Set a certificate that has been imported to be the Appliance Proxy certificate used for communicating with Appliance Proxies on other sites.
	//
	//  Please use below APIs.
	//  POST https://<nsx-mgr>/api/v1/trust-management/certificates/<cert-id>?action=apply_certificate&service_type=APH
	//
	// Deprecated: This API element is deprecated.
	//
	// @param setInterSiteAphCertificateRequestParam (required)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Setapplianceproxycertificateforintersitecommunication(setInterSiteAphCertificateRequestParam nsxModel.SetInterSiteAphCertificateRequest) error

	// Set a certificate that has been imported to be either the principal identity certificate for the local cluster with either GM or LM service type. Currently, the service type specified must match the current service type of the local cluster.
	//
	//  Please use below APIs.
	//  POST https://<nsx-mgr>/api/v1/trust-management/certificates/<cert-id>?action=apply_certificate&service_type=LOCAL_MANAGER
	//  POST https://<nsx-mgr>/api/v1/trust-management/certificates/<cert-id>?action=apply_certificate&service_type=GLOBAL_MANAGER
	//
	// Deprecated: This API element is deprecated.
	//
	// @param setPrincipalIdentityCertificateForFederationRequestParam (required)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Setpicertificateforfederation(setPrincipalIdentityCertificateForFederationRequestParam nsxModel.SetPrincipalIdentityCertificateForFederationRequest) error

	// Checks whether certificate is valid. When the certificate contains a chain, the full chain is validated. The usage parameter can be SERVER (default) or CLIENT. This indicates whether the certificate needs to be validated as a server-auth or a client-auth certificate.
	//
	// @param certIdParam ID of certificate to validate (required)
	// @param usageParam Usage Type of the Certificate, SERVER or CLIENT. Default is SERVER (optional)
	// @return com.vmware.nsx.model.CertificateCheckingStatus
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Validate(certIdParam string, usageParam *string) (nsxModel.CertificateCheckingStatus, error)
}

type certificatesClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewCertificatesClient(connector vapiProtocolClient_.Connector) *certificatesClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx.trust_management.certificates")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"applycertificate":          vapiCore_.NewMethodIdentifier(interfaceIdentifier, "applycertificate"),
		"delete":                    vapiCore_.NewMethodIdentifier(interfaceIdentifier, "delete"),
		"fetchpeercertificatechain": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "fetchpeercertificatechain"),
		"get":                       vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
		"importcertificate":         vapiCore_.NewMethodIdentifier(interfaceIdentifier, "importcertificate"),
		"importtrustedca":           vapiCore_.NewMethodIdentifier(interfaceIdentifier, "importtrustedca"),
		"list":                      vapiCore_.NewMethodIdentifier(interfaceIdentifier, "list"),
		"setapplianceproxycertificateforintersitecommunication": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "setapplianceproxycertificateforintersitecommunication"),
		"setpicertificateforfederation":                         vapiCore_.NewMethodIdentifier(interfaceIdentifier, "setpicertificateforfederation"),
		"validate":                                              vapiCore_.NewMethodIdentifier(interfaceIdentifier, "validate"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	cIface := certificatesClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &cIface
}

func (cIface *certificatesClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := cIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (cIface *certificatesClient) Applycertificate(certIdParam string, serviceTypeParam string, nodeIdParam *string) error {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	operationRestMetaData := certificatesApplycertificateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(certificatesApplycertificateInputType(), typeConverter)
	sv.AddStructField("CertId", certIdParam)
	sv.AddStructField("ServiceType", serviceTypeParam)
	sv.AddStructField("NodeId", nodeIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.nsx.trust_management.certificates", "applycertificate", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (cIface *certificatesClient) Delete(certIdParam string) error {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	operationRestMetaData := certificatesDeleteRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(certificatesDeleteInputType(), typeConverter)
	sv.AddStructField("CertId", certIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.nsx.trust_management.certificates", "delete", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (cIface *certificatesClient) Fetchpeercertificatechain(tlsServiceEndpointParam nsxModel.TlsServiceEndpoint) (nsxModel.PeerCertificateChain, error) {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	operationRestMetaData := certificatesFetchpeercertificatechainRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(certificatesFetchpeercertificatechainInputType(), typeConverter)
	sv.AddStructField("TlsServiceEndpoint", tlsServiceEndpointParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.PeerCertificateChain
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.nsx.trust_management.certificates", "fetchpeercertificatechain", inputDataValue, executionContext)
	var emptyOutput nsxModel.PeerCertificateChain
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), CertificatesFetchpeercertificatechainOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.PeerCertificateChain), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (cIface *certificatesClient) Get(certIdParam string, detailsParam *bool) (nsxModel.Certificate, error) {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	operationRestMetaData := certificatesGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(certificatesGetInputType(), typeConverter)
	sv.AddStructField("CertId", certIdParam)
	sv.AddStructField("Details", detailsParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.Certificate
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.nsx.trust_management.certificates", "get", inputDataValue, executionContext)
	var emptyOutput nsxModel.Certificate
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), CertificatesGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.Certificate), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (cIface *certificatesClient) Importcertificate(trustObjectDataParam nsxModel.TrustObjectData) (nsxModel.CertificateList, error) {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	operationRestMetaData := certificatesImportcertificateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(certificatesImportcertificateInputType(), typeConverter)
	sv.AddStructField("TrustObjectData", trustObjectDataParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.CertificateList
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.nsx.trust_management.certificates", "importcertificate", inputDataValue, executionContext)
	var emptyOutput nsxModel.CertificateList
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), CertificatesImportcertificateOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.CertificateList), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (cIface *certificatesClient) Importtrustedca(aliasParam string, trustObjectDataParam nsxModel.TrustObjectData) error {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	operationRestMetaData := certificatesImporttrustedcaRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(certificatesImporttrustedcaInputType(), typeConverter)
	sv.AddStructField("Alias", aliasParam)
	sv.AddStructField("TrustObjectData", trustObjectDataParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.nsx.trust_management.certificates", "importtrustedca", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (cIface *certificatesClient) List(cursorParam *string, detailsParam *bool, includedFieldsParam *string, nodeIdParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string, type_Param *string) (nsxModel.CertificateList, error) {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	operationRestMetaData := certificatesListRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(certificatesListInputType(), typeConverter)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("Details", detailsParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("NodeId", nodeIdParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	sv.AddStructField("Type_", type_Param)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.CertificateList
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.nsx.trust_management.certificates", "list", inputDataValue, executionContext)
	var emptyOutput nsxModel.CertificateList
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), CertificatesListOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.CertificateList), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (cIface *certificatesClient) Setapplianceproxycertificateforintersitecommunication(setInterSiteAphCertificateRequestParam nsxModel.SetInterSiteAphCertificateRequest) error {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	operationRestMetaData := certificatesSetapplianceproxycertificateforintersitecommunicationRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(certificatesSetapplianceproxycertificateforintersitecommunicationInputType(), typeConverter)
	sv.AddStructField("SetInterSiteAphCertificateRequest", setInterSiteAphCertificateRequestParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.nsx.trust_management.certificates", "setapplianceproxycertificateforintersitecommunication", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (cIface *certificatesClient) Setpicertificateforfederation(setPrincipalIdentityCertificateForFederationRequestParam nsxModel.SetPrincipalIdentityCertificateForFederationRequest) error {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	operationRestMetaData := certificatesSetpicertificateforfederationRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(certificatesSetpicertificateforfederationInputType(), typeConverter)
	sv.AddStructField("SetPrincipalIdentityCertificateForFederationRequest", setPrincipalIdentityCertificateForFederationRequestParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.nsx.trust_management.certificates", "setpicertificateforfederation", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (cIface *certificatesClient) Validate(certIdParam string, usageParam *string) (nsxModel.CertificateCheckingStatus, error) {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	operationRestMetaData := certificatesValidateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(certificatesValidateInputType(), typeConverter)
	sv.AddStructField("CertId", certIdParam)
	sv.AddStructField("Usage", usageParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.CertificateCheckingStatus
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.nsx.trust_management.certificates", "validate", inputDataValue, executionContext)
	var emptyOutput nsxModel.CertificateCheckingStatus
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), CertificatesValidateOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.CertificateCheckingStatus), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
