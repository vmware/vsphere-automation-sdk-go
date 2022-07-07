// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: TokenPrincipalIdentities
// Used by client-side stubs.

package trust_management

import (
	"github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
)

const _ = core.SupportedByRuntimeVersion1

type TokenPrincipalIdentitiesClient interface {

	// Register a principal identity that is going to be authenticated through a token. The combination name and node_id needs to be unique across token-based and certificate-based principal identities.
	//
	// @param tokenBasedPrincipalIdentityParam (required)
	// @return com.vmware.nsx.model.TokenBasedPrincipalIdentity
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Create(tokenBasedPrincipalIdentityParam model.TokenBasedPrincipalIdentity) (model.TokenBasedPrincipalIdentity, error)

	// Delete a token-based principal identity.
	//
	// @param principalIdentityIdParam Unique id of the token-based principal identity to delete (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Delete(principalIdentityIdParam string) error

	// Get a stored token-based principal identity
	//
	// @param principalIdentityIdParam ID of the principal identity to get (required)
	// @return com.vmware.nsx.model.TokenBasedPrincipalIdentity
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(principalIdentityIdParam string) (model.TokenBasedPrincipalIdentity, error)

	// Return the list of token-based principal identities. | These don't have certificate or role information.
	// @return com.vmware.nsx.model.TokenBasedPrincipalIdentityListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List() (model.TokenBasedPrincipalIdentityListResult, error)
}

type tokenPrincipalIdentitiesClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewTokenPrincipalIdentitiesClient(connector client.Connector) *tokenPrincipalIdentitiesClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.nsx.trust_management.token_principal_identities")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"create": core.NewMethodIdentifier(interfaceIdentifier, "create"),
		"delete": core.NewMethodIdentifier(interfaceIdentifier, "delete"),
		"get":    core.NewMethodIdentifier(interfaceIdentifier, "get"),
		"list":   core.NewMethodIdentifier(interfaceIdentifier, "list"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	tIface := tokenPrincipalIdentitiesClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &tIface
}

func (tIface *tokenPrincipalIdentitiesClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := tIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (tIface *tokenPrincipalIdentitiesClient) Create(tokenBasedPrincipalIdentityParam model.TokenBasedPrincipalIdentity) (model.TokenBasedPrincipalIdentity, error) {
	typeConverter := tIface.connector.TypeConverter()
	executionContext := tIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(tokenPrincipalIdentitiesCreateInputType(), typeConverter)
	sv.AddStructField("TokenBasedPrincipalIdentity", tokenBasedPrincipalIdentityParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.TokenBasedPrincipalIdentity
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := tokenPrincipalIdentitiesCreateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	tIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := tIface.connector.GetApiProvider().Invoke("com.vmware.nsx.trust_management.token_principal_identities", "create", inputDataValue, executionContext)
	var emptyOutput model.TokenBasedPrincipalIdentity
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), tokenPrincipalIdentitiesCreateOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.TokenBasedPrincipalIdentity), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (tIface *tokenPrincipalIdentitiesClient) Delete(principalIdentityIdParam string) error {
	typeConverter := tIface.connector.TypeConverter()
	executionContext := tIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(tokenPrincipalIdentitiesDeleteInputType(), typeConverter)
	sv.AddStructField("PrincipalIdentityId", principalIdentityIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := tokenPrincipalIdentitiesDeleteRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	tIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := tIface.connector.GetApiProvider().Invoke("com.vmware.nsx.trust_management.token_principal_identities", "delete", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (tIface *tokenPrincipalIdentitiesClient) Get(principalIdentityIdParam string) (model.TokenBasedPrincipalIdentity, error) {
	typeConverter := tIface.connector.TypeConverter()
	executionContext := tIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(tokenPrincipalIdentitiesGetInputType(), typeConverter)
	sv.AddStructField("PrincipalIdentityId", principalIdentityIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.TokenBasedPrincipalIdentity
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := tokenPrincipalIdentitiesGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	tIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := tIface.connector.GetApiProvider().Invoke("com.vmware.nsx.trust_management.token_principal_identities", "get", inputDataValue, executionContext)
	var emptyOutput model.TokenBasedPrincipalIdentity
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), tokenPrincipalIdentitiesGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.TokenBasedPrincipalIdentity), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (tIface *tokenPrincipalIdentitiesClient) List() (model.TokenBasedPrincipalIdentityListResult, error) {
	typeConverter := tIface.connector.TypeConverter()
	executionContext := tIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(tokenPrincipalIdentitiesListInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.TokenBasedPrincipalIdentityListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := tokenPrincipalIdentitiesListRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	tIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := tIface.connector.GetApiProvider().Invoke("com.vmware.nsx.trust_management.token_principal_identities", "list", inputDataValue, executionContext)
	var emptyOutput model.TokenBasedPrincipalIdentityListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), tokenPrincipalIdentitiesListOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.TokenBasedPrincipalIdentityListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
