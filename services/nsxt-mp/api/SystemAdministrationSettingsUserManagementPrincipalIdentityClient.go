// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: SystemAdministrationSettingsUserManagementPrincipalIdentity
// Used by client-side stubs.

package api

import (
	"github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/model"
)

const _ = core.SupportedByRuntimeVersion1

type SystemAdministrationSettingsUserManagementPrincipalIdentityClient interface {

	// This request also fetches the issuer and jwks_uri meta-data from the OIDC end-point and stores it.
	//
	// @param oidcEndPointParam (required)
	// @return com.vmware.model.OidcEndPoint
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	AddOidcEndPoint(oidcEndPointParam model.OidcEndPoint) (model.OidcEndPoint, error)

	// Delete a principal identity. It does not delete the certificate.
	//
	// @param principalIdentityIdParam Unique id of the principal identity to delete (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	DeletePrincipalIdentity(principalIdentityIdParam string) error

	// Delete a token-based principal identity.
	//
	// @param principalIdentityIdParam Unique id of the token-based principal identity to delete (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	DeleteTokenBasedPrincipalIdentity(principalIdentityIdParam string) error

	// When ?refresh=true is added to the request, the meta-data is newly fetched from the OIDC end-point.
	//
	// @param idParam (required)
	// @param refreshParam Refresh meta-data (optional, default to false)
	// @return com.vmware.model.OidcEndPoint
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetOidcEndPoint(idParam string, refreshParam *bool) (model.OidcEndPoint, error)

	// Returns the list of principals registered with a certificate.
	// @return com.vmware.model.PrincipalIdentityList
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetPrincipalIdentities() (model.PrincipalIdentityList, error)

	// Get a stored principal identity
	//
	// @param principalIdentityIdParam ID of the principal identity to get (required)
	// @return com.vmware.model.PrincipalIdentity
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetPrincipalIdentity(principalIdentityIdParam string) (model.PrincipalIdentity, error)

	// Get a stored token-based principal identity
	//
	// @param principalIdentityIdParam ID of the principal identity to get (required)
	// @return com.vmware.model.TokenBasedPrincipalIdentity
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetTokenBasedPrincipalIdentity(principalIdentityIdParam string) (model.TokenBasedPrincipalIdentity, error)

	// Return the list of OpenID Connect end-points.
	// @return com.vmware.model.OidcEndPointListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ListOidcEndPoints() (model.OidcEndPointListResult, error)

	// Return the list of token-based principal identities. | These don't have certificate or role information.
	// @return com.vmware.model.TokenBasedPrincipalIdentityListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ListTokenBasedPrincipalIdentities() (model.TokenBasedPrincipalIdentityListResult, error)

	// Associates a principal's name with a certificate that is used to authenticate. The combination name and node_id needs to be unique across token-based and certificate-based principal identities. Deprecated, use POST /trust-management/principal-identities/with-certificate instead.
	//
	// @param principalIdentityParam (required)
	// @return com.vmware.model.PrincipalIdentity
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	RegisterPrincipalIdentity(principalIdentityParam model.PrincipalIdentity) (model.PrincipalIdentity, error)

	// Create a principal identity with a new, unused, certificate. The combination name and node_id needs to be unique across token-based and certificate-based principal identities.
	//
	// @param principalIdentityWithCertificateParam (required)
	// @return com.vmware.model.PrincipalIdentity
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	RegisterPrincipalIdentityWithCertificate(principalIdentityWithCertificateParam model.PrincipalIdentityWithCertificate) (model.PrincipalIdentity, error)

	// Register a principal identity that is going to be authenticated through a token. The combination name and node_id needs to be unique across token-based and certificate-based principal identities.
	//
	// @param tokenBasedPrincipalIdentityParam (required)
	// @return com.vmware.model.TokenBasedPrincipalIdentity
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	RegisterTokenBasedPrincipalIdentity(tokenBasedPrincipalIdentityParam model.TokenBasedPrincipalIdentity) (model.TokenBasedPrincipalIdentity, error)

	// Update a OpenID Connect end-point's thumbprint used to connect to the oidc_uri through SSL
	//
	// @param updateOidcEndPointThumbprintRequestParam (required)
	// @return com.vmware.model.OidcEndPoint
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	UpdateOidcEndPointThumbprintUpdateThumbprint(updateOidcEndPointThumbprintRequestParam model.UpdateOidcEndPointThumbprintRequest) (model.OidcEndPoint, error)

	// Update a principal identity's certificate
	//
	// @param updatePrincipalIdentityCertificateRequestParam (required)
	// @return com.vmware.model.PrincipalIdentity
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	UpdatePrincipalIdentityCertificateUpdateCertificate(updatePrincipalIdentityCertificateRequestParam model.UpdatePrincipalIdentityCertificateRequest) (model.PrincipalIdentity, error)
}

type systemAdministrationSettingsUserManagementPrincipalIdentityClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewSystemAdministrationSettingsUserManagementPrincipalIdentityClient(connector client.Connector) *systemAdministrationSettingsUserManagementPrincipalIdentityClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.api.system_administration_settings_user_management_principal_identity")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"add_oidc_end_point":                                       core.NewMethodIdentifier(interfaceIdentifier, "add_oidc_end_point"),
		"delete_principal_identity":                                core.NewMethodIdentifier(interfaceIdentifier, "delete_principal_identity"),
		"delete_token_based_principal_identity":                    core.NewMethodIdentifier(interfaceIdentifier, "delete_token_based_principal_identity"),
		"get_oidc_end_point":                                       core.NewMethodIdentifier(interfaceIdentifier, "get_oidc_end_point"),
		"get_principal_identities":                                 core.NewMethodIdentifier(interfaceIdentifier, "get_principal_identities"),
		"get_principal_identity":                                   core.NewMethodIdentifier(interfaceIdentifier, "get_principal_identity"),
		"get_token_based_principal_identity":                       core.NewMethodIdentifier(interfaceIdentifier, "get_token_based_principal_identity"),
		"list_oidc_end_points":                                     core.NewMethodIdentifier(interfaceIdentifier, "list_oidc_end_points"),
		"list_token_based_principal_identities":                    core.NewMethodIdentifier(interfaceIdentifier, "list_token_based_principal_identities"),
		"register_principal_identity":                              core.NewMethodIdentifier(interfaceIdentifier, "register_principal_identity"),
		"register_principal_identity_with_certificate":             core.NewMethodIdentifier(interfaceIdentifier, "register_principal_identity_with_certificate"),
		"register_token_based_principal_identity":                  core.NewMethodIdentifier(interfaceIdentifier, "register_token_based_principal_identity"),
		"update_oidc_end_point_thumbprint_update_thumbprint":       core.NewMethodIdentifier(interfaceIdentifier, "update_oidc_end_point_thumbprint_update_thumbprint"),
		"update_principal_identity_certificate_update_certificate": core.NewMethodIdentifier(interfaceIdentifier, "update_principal_identity_certificate_update_certificate"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	sIface := systemAdministrationSettingsUserManagementPrincipalIdentityClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &sIface
}

func (sIface *systemAdministrationSettingsUserManagementPrincipalIdentityClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := sIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (sIface *systemAdministrationSettingsUserManagementPrincipalIdentityClient) AddOidcEndPoint(oidcEndPointParam model.OidcEndPoint) (model.OidcEndPoint, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationSettingsUserManagementPrincipalIdentityAddOidcEndPointInputType(), typeConverter)
	sv.AddStructField("OidcEndPoint", oidcEndPointParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.OidcEndPoint
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationSettingsUserManagementPrincipalIdentityAddOidcEndPointRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_settings_user_management_principal_identity", "add_oidc_end_point", inputDataValue, executionContext)
	var emptyOutput model.OidcEndPoint
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationSettingsUserManagementPrincipalIdentityAddOidcEndPointOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.OidcEndPoint), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationSettingsUserManagementPrincipalIdentityClient) DeletePrincipalIdentity(principalIdentityIdParam string) error {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationSettingsUserManagementPrincipalIdentityDeletePrincipalIdentityInputType(), typeConverter)
	sv.AddStructField("PrincipalIdentityId", principalIdentityIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationSettingsUserManagementPrincipalIdentityDeletePrincipalIdentityRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_settings_user_management_principal_identity", "delete_principal_identity", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (sIface *systemAdministrationSettingsUserManagementPrincipalIdentityClient) DeleteTokenBasedPrincipalIdentity(principalIdentityIdParam string) error {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationSettingsUserManagementPrincipalIdentityDeleteTokenBasedPrincipalIdentityInputType(), typeConverter)
	sv.AddStructField("PrincipalIdentityId", principalIdentityIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationSettingsUserManagementPrincipalIdentityDeleteTokenBasedPrincipalIdentityRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_settings_user_management_principal_identity", "delete_token_based_principal_identity", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (sIface *systemAdministrationSettingsUserManagementPrincipalIdentityClient) GetOidcEndPoint(idParam string, refreshParam *bool) (model.OidcEndPoint, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationSettingsUserManagementPrincipalIdentityGetOidcEndPointInputType(), typeConverter)
	sv.AddStructField("Id", idParam)
	sv.AddStructField("Refresh", refreshParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.OidcEndPoint
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationSettingsUserManagementPrincipalIdentityGetOidcEndPointRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_settings_user_management_principal_identity", "get_oidc_end_point", inputDataValue, executionContext)
	var emptyOutput model.OidcEndPoint
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationSettingsUserManagementPrincipalIdentityGetOidcEndPointOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.OidcEndPoint), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationSettingsUserManagementPrincipalIdentityClient) GetPrincipalIdentities() (model.PrincipalIdentityList, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationSettingsUserManagementPrincipalIdentityGetPrincipalIdentitiesInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.PrincipalIdentityList
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationSettingsUserManagementPrincipalIdentityGetPrincipalIdentitiesRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_settings_user_management_principal_identity", "get_principal_identities", inputDataValue, executionContext)
	var emptyOutput model.PrincipalIdentityList
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationSettingsUserManagementPrincipalIdentityGetPrincipalIdentitiesOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.PrincipalIdentityList), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationSettingsUserManagementPrincipalIdentityClient) GetPrincipalIdentity(principalIdentityIdParam string) (model.PrincipalIdentity, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationSettingsUserManagementPrincipalIdentityGetPrincipalIdentityInputType(), typeConverter)
	sv.AddStructField("PrincipalIdentityId", principalIdentityIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.PrincipalIdentity
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationSettingsUserManagementPrincipalIdentityGetPrincipalIdentityRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_settings_user_management_principal_identity", "get_principal_identity", inputDataValue, executionContext)
	var emptyOutput model.PrincipalIdentity
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationSettingsUserManagementPrincipalIdentityGetPrincipalIdentityOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.PrincipalIdentity), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationSettingsUserManagementPrincipalIdentityClient) GetTokenBasedPrincipalIdentity(principalIdentityIdParam string) (model.TokenBasedPrincipalIdentity, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationSettingsUserManagementPrincipalIdentityGetTokenBasedPrincipalIdentityInputType(), typeConverter)
	sv.AddStructField("PrincipalIdentityId", principalIdentityIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.TokenBasedPrincipalIdentity
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationSettingsUserManagementPrincipalIdentityGetTokenBasedPrincipalIdentityRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_settings_user_management_principal_identity", "get_token_based_principal_identity", inputDataValue, executionContext)
	var emptyOutput model.TokenBasedPrincipalIdentity
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationSettingsUserManagementPrincipalIdentityGetTokenBasedPrincipalIdentityOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.TokenBasedPrincipalIdentity), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationSettingsUserManagementPrincipalIdentityClient) ListOidcEndPoints() (model.OidcEndPointListResult, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationSettingsUserManagementPrincipalIdentityListOidcEndPointsInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.OidcEndPointListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationSettingsUserManagementPrincipalIdentityListOidcEndPointsRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_settings_user_management_principal_identity", "list_oidc_end_points", inputDataValue, executionContext)
	var emptyOutput model.OidcEndPointListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationSettingsUserManagementPrincipalIdentityListOidcEndPointsOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.OidcEndPointListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationSettingsUserManagementPrincipalIdentityClient) ListTokenBasedPrincipalIdentities() (model.TokenBasedPrincipalIdentityListResult, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationSettingsUserManagementPrincipalIdentityListTokenBasedPrincipalIdentitiesInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.TokenBasedPrincipalIdentityListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationSettingsUserManagementPrincipalIdentityListTokenBasedPrincipalIdentitiesRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_settings_user_management_principal_identity", "list_token_based_principal_identities", inputDataValue, executionContext)
	var emptyOutput model.TokenBasedPrincipalIdentityListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationSettingsUserManagementPrincipalIdentityListTokenBasedPrincipalIdentitiesOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.TokenBasedPrincipalIdentityListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationSettingsUserManagementPrincipalIdentityClient) RegisterPrincipalIdentity(principalIdentityParam model.PrincipalIdentity) (model.PrincipalIdentity, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationSettingsUserManagementPrincipalIdentityRegisterPrincipalIdentityInputType(), typeConverter)
	sv.AddStructField("PrincipalIdentity", principalIdentityParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.PrincipalIdentity
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationSettingsUserManagementPrincipalIdentityRegisterPrincipalIdentityRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_settings_user_management_principal_identity", "register_principal_identity", inputDataValue, executionContext)
	var emptyOutput model.PrincipalIdentity
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationSettingsUserManagementPrincipalIdentityRegisterPrincipalIdentityOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.PrincipalIdentity), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationSettingsUserManagementPrincipalIdentityClient) RegisterPrincipalIdentityWithCertificate(principalIdentityWithCertificateParam model.PrincipalIdentityWithCertificate) (model.PrincipalIdentity, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationSettingsUserManagementPrincipalIdentityRegisterPrincipalIdentityWithCertificateInputType(), typeConverter)
	sv.AddStructField("PrincipalIdentityWithCertificate", principalIdentityWithCertificateParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.PrincipalIdentity
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationSettingsUserManagementPrincipalIdentityRegisterPrincipalIdentityWithCertificateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_settings_user_management_principal_identity", "register_principal_identity_with_certificate", inputDataValue, executionContext)
	var emptyOutput model.PrincipalIdentity
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationSettingsUserManagementPrincipalIdentityRegisterPrincipalIdentityWithCertificateOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.PrincipalIdentity), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationSettingsUserManagementPrincipalIdentityClient) RegisterTokenBasedPrincipalIdentity(tokenBasedPrincipalIdentityParam model.TokenBasedPrincipalIdentity) (model.TokenBasedPrincipalIdentity, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationSettingsUserManagementPrincipalIdentityRegisterTokenBasedPrincipalIdentityInputType(), typeConverter)
	sv.AddStructField("TokenBasedPrincipalIdentity", tokenBasedPrincipalIdentityParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.TokenBasedPrincipalIdentity
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationSettingsUserManagementPrincipalIdentityRegisterTokenBasedPrincipalIdentityRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_settings_user_management_principal_identity", "register_token_based_principal_identity", inputDataValue, executionContext)
	var emptyOutput model.TokenBasedPrincipalIdentity
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationSettingsUserManagementPrincipalIdentityRegisterTokenBasedPrincipalIdentityOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.TokenBasedPrincipalIdentity), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationSettingsUserManagementPrincipalIdentityClient) UpdateOidcEndPointThumbprintUpdateThumbprint(updateOidcEndPointThumbprintRequestParam model.UpdateOidcEndPointThumbprintRequest) (model.OidcEndPoint, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationSettingsUserManagementPrincipalIdentityUpdateOidcEndPointThumbprintUpdateThumbprintInputType(), typeConverter)
	sv.AddStructField("UpdateOidcEndPointThumbprintRequest", updateOidcEndPointThumbprintRequestParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.OidcEndPoint
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationSettingsUserManagementPrincipalIdentityUpdateOidcEndPointThumbprintUpdateThumbprintRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_settings_user_management_principal_identity", "update_oidc_end_point_thumbprint_update_thumbprint", inputDataValue, executionContext)
	var emptyOutput model.OidcEndPoint
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationSettingsUserManagementPrincipalIdentityUpdateOidcEndPointThumbprintUpdateThumbprintOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.OidcEndPoint), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationSettingsUserManagementPrincipalIdentityClient) UpdatePrincipalIdentityCertificateUpdateCertificate(updatePrincipalIdentityCertificateRequestParam model.UpdatePrincipalIdentityCertificateRequest) (model.PrincipalIdentity, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationSettingsUserManagementPrincipalIdentityUpdatePrincipalIdentityCertificateUpdateCertificateInputType(), typeConverter)
	sv.AddStructField("UpdatePrincipalIdentityCertificateRequest", updatePrincipalIdentityCertificateRequestParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.PrincipalIdentity
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationSettingsUserManagementPrincipalIdentityUpdatePrincipalIdentityCertificateUpdateCertificateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_settings_user_management_principal_identity", "update_principal_identity_certificate_update_certificate", inputDataValue, executionContext)
	var emptyOutput model.PrincipalIdentity
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationSettingsUserManagementPrincipalIdentityUpdatePrincipalIdentityCertificateUpdateCertificateOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.PrincipalIdentity), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
