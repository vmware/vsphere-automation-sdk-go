// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: PrincipalIdentities
// Used by client-side stubs.

package trust_management

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/core"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/lib"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/protocol/client"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt-gm/model"
)

const _ = core.SupportedByRuntimeVersion1

type PrincipalIdentitiesClient interface {

	// Associates a principal's name with a certificate that is used to authenticate. The combination name and node_id needs to be unique across token-based and certificate-based principal identities. Deprecated, use POST /trust-management/principal-identities/with-certificate instead.
	//
	// @param principalIdentityParam (required)
	// @return com.vmware.nsx_global_policy.model.PrincipalIdentity
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Create(principalIdentityParam model.PrincipalIdentity) (model.PrincipalIdentity, error)

	// Delete a principal identity. It does not delete the certificate.
	//
	// @param principalIdentityIdParam Unique id of the principal identity to delete (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Delete(principalIdentityIdParam string) error

	// Get a stored principal identity
	//
	// @param principalIdentityIdParam ID of the principal identity to get (required)
	// @return com.vmware.nsx_global_policy.model.PrincipalIdentity
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(principalIdentityIdParam string) (model.PrincipalIdentity, error)

	// Returns the list of principals registered with a certificate.
	// @return com.vmware.nsx_global_policy.model.PrincipalIdentityList
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List() (model.PrincipalIdentityList, error)

	// Update a principal identity's certificate
	//
	// @param updatePrincipalIdentityCertificateRequestParam (required)
	// @return com.vmware.nsx_global_policy.model.PrincipalIdentity
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Updatecertificate(updatePrincipalIdentityCertificateRequestParam model.UpdatePrincipalIdentityCertificateRequest) (model.PrincipalIdentity, error)
}

type principalIdentitiesClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewPrincipalIdentitiesClient(connector client.Connector) *principalIdentitiesClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.nsx_global_policy.trust_management.principal_identities")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"create":            core.NewMethodIdentifier(interfaceIdentifier, "create"),
		"delete":            core.NewMethodIdentifier(interfaceIdentifier, "delete"),
		"get":               core.NewMethodIdentifier(interfaceIdentifier, "get"),
		"list":              core.NewMethodIdentifier(interfaceIdentifier, "list"),
		"updatecertificate": core.NewMethodIdentifier(interfaceIdentifier, "updatecertificate"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	pIface := principalIdentitiesClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &pIface
}

func (pIface *principalIdentitiesClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := pIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (pIface *principalIdentitiesClient) Create(principalIdentityParam model.PrincipalIdentity) (model.PrincipalIdentity, error) {
	typeConverter := pIface.connector.TypeConverter()
	executionContext := pIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(principalIdentitiesCreateInputType(), typeConverter)
	sv.AddStructField("PrincipalIdentity", principalIdentityParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.PrincipalIdentity
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := principalIdentitiesCreateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	pIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := pIface.connector.GetApiProvider().Invoke("com.vmware.nsx_global_policy.trust_management.principal_identities", "create", inputDataValue, executionContext)
	var emptyOutput model.PrincipalIdentity
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), principalIdentitiesCreateOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.PrincipalIdentity), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), pIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (pIface *principalIdentitiesClient) Delete(principalIdentityIdParam string) error {
	typeConverter := pIface.connector.TypeConverter()
	executionContext := pIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(principalIdentitiesDeleteInputType(), typeConverter)
	sv.AddStructField("PrincipalIdentityId", principalIdentityIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := principalIdentitiesDeleteRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	pIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := pIface.connector.GetApiProvider().Invoke("com.vmware.nsx_global_policy.trust_management.principal_identities", "delete", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), pIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (pIface *principalIdentitiesClient) Get(principalIdentityIdParam string) (model.PrincipalIdentity, error) {
	typeConverter := pIface.connector.TypeConverter()
	executionContext := pIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(principalIdentitiesGetInputType(), typeConverter)
	sv.AddStructField("PrincipalIdentityId", principalIdentityIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.PrincipalIdentity
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := principalIdentitiesGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	pIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := pIface.connector.GetApiProvider().Invoke("com.vmware.nsx_global_policy.trust_management.principal_identities", "get", inputDataValue, executionContext)
	var emptyOutput model.PrincipalIdentity
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), principalIdentitiesGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.PrincipalIdentity), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), pIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (pIface *principalIdentitiesClient) List() (model.PrincipalIdentityList, error) {
	typeConverter := pIface.connector.TypeConverter()
	executionContext := pIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(principalIdentitiesListInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.PrincipalIdentityList
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := principalIdentitiesListRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	pIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := pIface.connector.GetApiProvider().Invoke("com.vmware.nsx_global_policy.trust_management.principal_identities", "list", inputDataValue, executionContext)
	var emptyOutput model.PrincipalIdentityList
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), principalIdentitiesListOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.PrincipalIdentityList), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), pIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (pIface *principalIdentitiesClient) Updatecertificate(updatePrincipalIdentityCertificateRequestParam model.UpdatePrincipalIdentityCertificateRequest) (model.PrincipalIdentity, error) {
	typeConverter := pIface.connector.TypeConverter()
	executionContext := pIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(principalIdentitiesUpdatecertificateInputType(), typeConverter)
	sv.AddStructField("UpdatePrincipalIdentityCertificateRequest", updatePrincipalIdentityCertificateRequestParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.PrincipalIdentity
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := principalIdentitiesUpdatecertificateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	pIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := pIface.connector.GetApiProvider().Invoke("com.vmware.nsx_global_policy.trust_management.principal_identities", "updatecertificate", inputDataValue, executionContext)
	var emptyOutput model.PrincipalIdentity
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), principalIdentitiesUpdatecertificateOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.PrincipalIdentity), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), pIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
