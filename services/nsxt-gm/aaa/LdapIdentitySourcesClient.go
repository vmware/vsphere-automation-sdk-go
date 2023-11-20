// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: LdapIdentitySources
// Used by client-side stubs.

package aaa

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	nsx_global_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-gm/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type LdapIdentitySourcesClient interface {

	// Delete an LDAP identity source. Users defined in that source will no longer be able to access NSX.
	//
	// @param ldapIdentitySourceIdParam (required)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Delete(ldapIdentitySourceIdParam string) error

	// Attempt to connect to an LDAP server and retrieve the server certificate it presents.
	//
	// @param identitySourceLdapServerEndpointParam (required)
	// @return com.vmware.nsx_global_policy.model.PeerCertificateChain
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Fetchcertificate(identitySourceLdapServerEndpointParam nsx_global_policyModel.IdentitySourceLdapServerEndpoint) (nsx_global_policyModel.PeerCertificateChain, error)

	// Return details about one LDAP identity source
	//
	// @param ldapIdentitySourceIdParam (required)
	// @return com.vmware.nsx_global_policy.model.LdapIdentitySource
	// The return value will contain all the properties defined in nsx_global_policyModel.LdapIdentitySource.
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(ldapIdentitySourceIdParam string) (*vapiData_.StructValue, error)

	// Return a list of all configured LDAP identity sources.
	//
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @return com.vmware.nsx_global_policy.model.LdapIdentitySourceListResult
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (nsx_global_policyModel.LdapIdentitySourceListResult, error)

	// Attempt to connect to an existing LDAP identity source and report any errors encountered.
	//
	// @param ldapIdentitySourceIdParam (required)
	// @return com.vmware.nsx_global_policy.model.LdapIdentitySourceProbeResults
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Probe(ldapIdentitySourceIdParam string) (nsx_global_policyModel.LdapIdentitySourceProbeResults, error)

	// Verify that the configuration of an LDAP identity source is correct before actually creating the source.
	//
	// @param ldapIdentitySourceParam (required)
	// The parameter must contain all the properties defined in nsx_global_policyModel.LdapIdentitySource.
	// @return com.vmware.nsx_global_policy.model.LdapIdentitySourceProbeResults
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Probeidentitysource(ldapIdentitySourceParam *vapiData_.StructValue) (nsx_global_policyModel.LdapIdentitySourceProbeResults, error)

	// Attempt to connect to an LDAP server and ensure that the server can be contacted using the given URL and authentication credentials.
	//
	// @param identitySourceLdapServerParam (required)
	// @return com.vmware.nsx_global_policy.model.IdentitySourceLdapServerProbeResult
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Probeldapserver(identitySourceLdapServerParam nsx_global_policyModel.IdentitySourceLdapServer) (nsx_global_policyModel.IdentitySourceLdapServerProbeResult, error)

	// Create a new LDAP identity source or update the configuration of an existing LDAP identity source. You may wish to verify the new configuration using the POST /aaa/ldap-identity-sources?action=probe API before creating or changing the configuration. Note that if you are using LDAP on an active and standby NSX-T Global Manager in a federated environment, you must use the same name for your LDAP identity sources on the active and standby Global Managers.
	//
	// @param ldapIdentitySourceIdParam (required)
	// @param ldapIdentitySourceParam (required)
	// The parameter must contain all the properties defined in nsx_global_policyModel.LdapIdentitySource.
	// @return com.vmware.nsx_global_policy.model.LdapIdentitySource
	// The return value will contain all the properties defined in nsx_global_policyModel.LdapIdentitySource.
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(ldapIdentitySourceIdParam string, ldapIdentitySourceParam *vapiData_.StructValue) (*vapiData_.StructValue, error)
}

type ldapIdentitySourcesClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewLdapIdentitySourcesClient(connector vapiProtocolClient_.Connector) *ldapIdentitySourcesClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx_global_policy.aaa.ldap_identity_sources")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"delete":              vapiCore_.NewMethodIdentifier(interfaceIdentifier, "delete"),
		"fetchcertificate":    vapiCore_.NewMethodIdentifier(interfaceIdentifier, "fetchcertificate"),
		"get":                 vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
		"list":                vapiCore_.NewMethodIdentifier(interfaceIdentifier, "list"),
		"probe":               vapiCore_.NewMethodIdentifier(interfaceIdentifier, "probe"),
		"probeidentitysource": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "probeidentitysource"),
		"probeldapserver":     vapiCore_.NewMethodIdentifier(interfaceIdentifier, "probeldapserver"),
		"update":              vapiCore_.NewMethodIdentifier(interfaceIdentifier, "update"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	lIface := ldapIdentitySourcesClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &lIface
}

func (lIface *ldapIdentitySourcesClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := lIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (lIface *ldapIdentitySourcesClient) Delete(ldapIdentitySourceIdParam string) error {
	typeConverter := lIface.connector.TypeConverter()
	executionContext := lIface.connector.NewExecutionContext()
	operationRestMetaData := ldapIdentitySourcesDeleteRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(ldapIdentitySourcesDeleteInputType(), typeConverter)
	sv.AddStructField("LdapIdentitySourceId", ldapIdentitySourceIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := lIface.connector.GetApiProvider().Invoke("com.vmware.nsx_global_policy.aaa.ldap_identity_sources", "delete", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), lIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (lIface *ldapIdentitySourcesClient) Fetchcertificate(identitySourceLdapServerEndpointParam nsx_global_policyModel.IdentitySourceLdapServerEndpoint) (nsx_global_policyModel.PeerCertificateChain, error) {
	typeConverter := lIface.connector.TypeConverter()
	executionContext := lIface.connector.NewExecutionContext()
	operationRestMetaData := ldapIdentitySourcesFetchcertificateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(ldapIdentitySourcesFetchcertificateInputType(), typeConverter)
	sv.AddStructField("IdentitySourceLdapServerEndpoint", identitySourceLdapServerEndpointParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_global_policyModel.PeerCertificateChain
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := lIface.connector.GetApiProvider().Invoke("com.vmware.nsx_global_policy.aaa.ldap_identity_sources", "fetchcertificate", inputDataValue, executionContext)
	var emptyOutput nsx_global_policyModel.PeerCertificateChain
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), LdapIdentitySourcesFetchcertificateOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_global_policyModel.PeerCertificateChain), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), lIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (lIface *ldapIdentitySourcesClient) Get(ldapIdentitySourceIdParam string) (*vapiData_.StructValue, error) {
	typeConverter := lIface.connector.TypeConverter()
	executionContext := lIface.connector.NewExecutionContext()
	operationRestMetaData := ldapIdentitySourcesGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(ldapIdentitySourcesGetInputType(), typeConverter)
	sv.AddStructField("LdapIdentitySourceId", ldapIdentitySourceIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput *vapiData_.StructValue
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := lIface.connector.GetApiProvider().Invoke("com.vmware.nsx_global_policy.aaa.ldap_identity_sources", "get", inputDataValue, executionContext)
	var emptyOutput *vapiData_.StructValue
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), LdapIdentitySourcesGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(*vapiData_.StructValue), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), lIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (lIface *ldapIdentitySourcesClient) List(cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (nsx_global_policyModel.LdapIdentitySourceListResult, error) {
	typeConverter := lIface.connector.TypeConverter()
	executionContext := lIface.connector.NewExecutionContext()
	operationRestMetaData := ldapIdentitySourcesListRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(ldapIdentitySourcesListInputType(), typeConverter)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_global_policyModel.LdapIdentitySourceListResult
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := lIface.connector.GetApiProvider().Invoke("com.vmware.nsx_global_policy.aaa.ldap_identity_sources", "list", inputDataValue, executionContext)
	var emptyOutput nsx_global_policyModel.LdapIdentitySourceListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), LdapIdentitySourcesListOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_global_policyModel.LdapIdentitySourceListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), lIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (lIface *ldapIdentitySourcesClient) Probe(ldapIdentitySourceIdParam string) (nsx_global_policyModel.LdapIdentitySourceProbeResults, error) {
	typeConverter := lIface.connector.TypeConverter()
	executionContext := lIface.connector.NewExecutionContext()
	operationRestMetaData := ldapIdentitySourcesProbeRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(ldapIdentitySourcesProbeInputType(), typeConverter)
	sv.AddStructField("LdapIdentitySourceId", ldapIdentitySourceIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_global_policyModel.LdapIdentitySourceProbeResults
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := lIface.connector.GetApiProvider().Invoke("com.vmware.nsx_global_policy.aaa.ldap_identity_sources", "probe", inputDataValue, executionContext)
	var emptyOutput nsx_global_policyModel.LdapIdentitySourceProbeResults
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), LdapIdentitySourcesProbeOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_global_policyModel.LdapIdentitySourceProbeResults), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), lIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (lIface *ldapIdentitySourcesClient) Probeidentitysource(ldapIdentitySourceParam *vapiData_.StructValue) (nsx_global_policyModel.LdapIdentitySourceProbeResults, error) {
	typeConverter := lIface.connector.TypeConverter()
	executionContext := lIface.connector.NewExecutionContext()
	operationRestMetaData := ldapIdentitySourcesProbeidentitysourceRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(ldapIdentitySourcesProbeidentitysourceInputType(), typeConverter)
	sv.AddStructField("LdapIdentitySource", ldapIdentitySourceParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_global_policyModel.LdapIdentitySourceProbeResults
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := lIface.connector.GetApiProvider().Invoke("com.vmware.nsx_global_policy.aaa.ldap_identity_sources", "probeidentitysource", inputDataValue, executionContext)
	var emptyOutput nsx_global_policyModel.LdapIdentitySourceProbeResults
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), LdapIdentitySourcesProbeidentitysourceOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_global_policyModel.LdapIdentitySourceProbeResults), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), lIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (lIface *ldapIdentitySourcesClient) Probeldapserver(identitySourceLdapServerParam nsx_global_policyModel.IdentitySourceLdapServer) (nsx_global_policyModel.IdentitySourceLdapServerProbeResult, error) {
	typeConverter := lIface.connector.TypeConverter()
	executionContext := lIface.connector.NewExecutionContext()
	operationRestMetaData := ldapIdentitySourcesProbeldapserverRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(ldapIdentitySourcesProbeldapserverInputType(), typeConverter)
	sv.AddStructField("IdentitySourceLdapServer", identitySourceLdapServerParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_global_policyModel.IdentitySourceLdapServerProbeResult
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := lIface.connector.GetApiProvider().Invoke("com.vmware.nsx_global_policy.aaa.ldap_identity_sources", "probeldapserver", inputDataValue, executionContext)
	var emptyOutput nsx_global_policyModel.IdentitySourceLdapServerProbeResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), LdapIdentitySourcesProbeldapserverOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_global_policyModel.IdentitySourceLdapServerProbeResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), lIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (lIface *ldapIdentitySourcesClient) Update(ldapIdentitySourceIdParam string, ldapIdentitySourceParam *vapiData_.StructValue) (*vapiData_.StructValue, error) {
	typeConverter := lIface.connector.TypeConverter()
	executionContext := lIface.connector.NewExecutionContext()
	operationRestMetaData := ldapIdentitySourcesUpdateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(ldapIdentitySourcesUpdateInputType(), typeConverter)
	sv.AddStructField("LdapIdentitySourceId", ldapIdentitySourceIdParam)
	sv.AddStructField("LdapIdentitySource", ldapIdentitySourceParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput *vapiData_.StructValue
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := lIface.connector.GetApiProvider().Invoke("com.vmware.nsx_global_policy.aaa.ldap_identity_sources", "update", inputDataValue, executionContext)
	var emptyOutput *vapiData_.StructValue
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), LdapIdentitySourcesUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(*vapiData_.StructValue), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), lIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
