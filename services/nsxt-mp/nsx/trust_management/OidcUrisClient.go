// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: OidcUris
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

type OidcUrisClient interface {

	// This request also fetches the issuer and jwks_uri meta-data from the OIDC end-point and stores it.
	//
	// @param oidcEndPointParam (required)
	// @return com.vmware.nsx.model.OidcEndPoint
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Create(oidcEndPointParam nsxModel.OidcEndPoint) (nsxModel.OidcEndPoint, error)

	// When ?refresh=true is added to the request, the meta-data is newly fetched from the OIDC end-point.
	//
	// @param idParam (required)
	// @param refreshParam Refresh meta-data (optional, default to false)
	// @return com.vmware.nsx.model.OidcEndPoint
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(idParam string, refreshParam *bool) (nsxModel.OidcEndPoint, error)

	// Return the list of OpenID Connect end-points.
	//
	// @param oidcTypeParam Type of OIDC endpoint to return (optional)
	// @return com.vmware.nsx.model.OidcEndPointListResult
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(oidcTypeParam *string) (nsxModel.OidcEndPointListResult, error)

	// Refresh an OpenID Connect end-point by re-reading data from the OIDC URI.
	//
	// @param idParam (required)
	// @return com.vmware.nsx.model.OidcEndPoint
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Refresh(idParam string) (nsxModel.OidcEndPoint, error)

	// Update the properties of an OpenID Connect end-point. The oidc_uri property may not be changed. If you need to update the oidc_uri, you should delete the OIDC end-point and create a new one with the correct oidc_uri. This request also re-fetches the issuer, jwks_uri, and other meta-data from the OIDC end-point and stores it.
	//
	// @param idParam (required)
	// @param oidcEndPointParam (required)
	// @return com.vmware.nsx.model.OidcEndPoint
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(idParam string, oidcEndPointParam nsxModel.OidcEndPoint) (nsxModel.OidcEndPoint, error)

	// Update a OpenID Connect end-point's thumbprint used to connect to the oidc_uri through SSL
	//
	// @param updateOidcEndPointThumbprintRequestParam (required)
	// @return com.vmware.nsx.model.OidcEndPoint
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Updatethumbprint(updateOidcEndPointThumbprintRequestParam nsxModel.UpdateOidcEndPointThumbprintRequest) (nsxModel.OidcEndPoint, error)
}

type oidcUrisClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewOidcUrisClient(connector vapiProtocolClient_.Connector) *oidcUrisClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx.trust_management.oidc_uris")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"create":           vapiCore_.NewMethodIdentifier(interfaceIdentifier, "create"),
		"get":              vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
		"list":             vapiCore_.NewMethodIdentifier(interfaceIdentifier, "list"),
		"refresh":          vapiCore_.NewMethodIdentifier(interfaceIdentifier, "refresh"),
		"update":           vapiCore_.NewMethodIdentifier(interfaceIdentifier, "update"),
		"updatethumbprint": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "updatethumbprint"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	oIface := oidcUrisClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &oIface
}

func (oIface *oidcUrisClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := oIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (oIface *oidcUrisClient) Create(oidcEndPointParam nsxModel.OidcEndPoint) (nsxModel.OidcEndPoint, error) {
	typeConverter := oIface.connector.TypeConverter()
	executionContext := oIface.connector.NewExecutionContext()
	operationRestMetaData := oidcUrisCreateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(oidcUrisCreateInputType(), typeConverter)
	sv.AddStructField("OidcEndPoint", oidcEndPointParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.OidcEndPoint
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := oIface.connector.GetApiProvider().Invoke("com.vmware.nsx.trust_management.oidc_uris", "create", inputDataValue, executionContext)
	var emptyOutput nsxModel.OidcEndPoint
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), OidcUrisCreateOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.OidcEndPoint), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), oIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (oIface *oidcUrisClient) Get(idParam string, refreshParam *bool) (nsxModel.OidcEndPoint, error) {
	typeConverter := oIface.connector.TypeConverter()
	executionContext := oIface.connector.NewExecutionContext()
	operationRestMetaData := oidcUrisGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(oidcUrisGetInputType(), typeConverter)
	sv.AddStructField("Id", idParam)
	sv.AddStructField("Refresh", refreshParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.OidcEndPoint
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := oIface.connector.GetApiProvider().Invoke("com.vmware.nsx.trust_management.oidc_uris", "get", inputDataValue, executionContext)
	var emptyOutput nsxModel.OidcEndPoint
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), OidcUrisGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.OidcEndPoint), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), oIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (oIface *oidcUrisClient) List(oidcTypeParam *string) (nsxModel.OidcEndPointListResult, error) {
	typeConverter := oIface.connector.TypeConverter()
	executionContext := oIface.connector.NewExecutionContext()
	operationRestMetaData := oidcUrisListRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(oidcUrisListInputType(), typeConverter)
	sv.AddStructField("OidcType", oidcTypeParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.OidcEndPointListResult
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := oIface.connector.GetApiProvider().Invoke("com.vmware.nsx.trust_management.oidc_uris", "list", inputDataValue, executionContext)
	var emptyOutput nsxModel.OidcEndPointListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), OidcUrisListOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.OidcEndPointListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), oIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (oIface *oidcUrisClient) Refresh(idParam string) (nsxModel.OidcEndPoint, error) {
	typeConverter := oIface.connector.TypeConverter()
	executionContext := oIface.connector.NewExecutionContext()
	operationRestMetaData := oidcUrisRefreshRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(oidcUrisRefreshInputType(), typeConverter)
	sv.AddStructField("Id", idParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.OidcEndPoint
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := oIface.connector.GetApiProvider().Invoke("com.vmware.nsx.trust_management.oidc_uris", "refresh", inputDataValue, executionContext)
	var emptyOutput nsxModel.OidcEndPoint
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), OidcUrisRefreshOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.OidcEndPoint), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), oIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (oIface *oidcUrisClient) Update(idParam string, oidcEndPointParam nsxModel.OidcEndPoint) (nsxModel.OidcEndPoint, error) {
	typeConverter := oIface.connector.TypeConverter()
	executionContext := oIface.connector.NewExecutionContext()
	operationRestMetaData := oidcUrisUpdateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(oidcUrisUpdateInputType(), typeConverter)
	sv.AddStructField("Id", idParam)
	sv.AddStructField("OidcEndPoint", oidcEndPointParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.OidcEndPoint
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := oIface.connector.GetApiProvider().Invoke("com.vmware.nsx.trust_management.oidc_uris", "update", inputDataValue, executionContext)
	var emptyOutput nsxModel.OidcEndPoint
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), OidcUrisUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.OidcEndPoint), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), oIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (oIface *oidcUrisClient) Updatethumbprint(updateOidcEndPointThumbprintRequestParam nsxModel.UpdateOidcEndPointThumbprintRequest) (nsxModel.OidcEndPoint, error) {
	typeConverter := oIface.connector.TypeConverter()
	executionContext := oIface.connector.NewExecutionContext()
	operationRestMetaData := oidcUrisUpdatethumbprintRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(oidcUrisUpdatethumbprintInputType(), typeConverter)
	sv.AddStructField("UpdateOidcEndPointThumbprintRequest", updateOidcEndPointThumbprintRequestParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.OidcEndPoint
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := oIface.connector.GetApiProvider().Invoke("com.vmware.nsx.trust_management.oidc_uris", "updatethumbprint", inputDataValue, executionContext)
	var emptyOutput nsxModel.OidcEndPoint
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), OidcUrisUpdatethumbprintOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.OidcEndPoint), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), oIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
