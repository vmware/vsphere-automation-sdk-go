// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: LdapIdentitySources
// Used by client-side stubs.

package aaa

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/core"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/lib"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/protocol/client"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt/model"
)

const _ = core.SupportedByRuntimeVersion1

type LdapIdentitySourcesClient interface {

	// Delete an LDAP identity source. Users defined in that source will no longer be able to access NSX.
	//
	// @param ldapIdentitySourceIdParam (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Delete(ldapIdentitySourceIdParam string) error

	// Attempt to connect to an LDAP server and retrieve the server certificate it presents.
	//
	// @param identitySourceLdapServerEndpointParam (required)
	// @return com.vmware.nsx_policy.model.PeerCertificateChain
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Fetchcertificate(identitySourceLdapServerEndpointParam model.IdentitySourceLdapServerEndpoint) (model.PeerCertificateChain, error)

	// Return details about one LDAP identity source
	//
	// @param ldapIdentitySourceIdParam (required)
	// @return com.vmware.nsx_policy.model.LdapIdentitySource
	// The return value will contain all the properties defined in model.LdapIdentitySource.
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(ldapIdentitySourceIdParam string) (*data.StructValue, error)

	// Return a list of all configured LDAP identity sources.
	//
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @return com.vmware.nsx_policy.model.LdapIdentitySourceListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.LdapIdentitySourceListResult, error)

	// Attempt to connect to an existing LDAP identity source and report any errors encountered.
	//
	// @param ldapIdentitySourceIdParam (required)
	// @return com.vmware.nsx_policy.model.LdapIdentitySourceProbeResults
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Probe(ldapIdentitySourceIdParam string) (model.LdapIdentitySourceProbeResults, error)

	// Verify that the configuration of an LDAP identity source is correct before actually creating the source.
	//
	// @param ldapIdentitySourceParam (required)
	// The parameter must contain all the properties defined in model.LdapIdentitySource.
	// @return com.vmware.nsx_policy.model.LdapIdentitySourceProbeResults
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Probeidentitysource(ldapIdentitySourceParam *data.StructValue) (model.LdapIdentitySourceProbeResults, error)

	// Attempt to connect to an LDAP server and ensure that the server can be contacted using the given URL and authentication credentials.
	//
	// @param identitySourceLdapServerParam (required)
	// @return com.vmware.nsx_policy.model.IdentitySourceLdapServerProbeResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Probeldapserver(identitySourceLdapServerParam model.IdentitySourceLdapServer) (model.IdentitySourceLdapServerProbeResult, error)

	// Update the configuration of an existing LDAP identity source. You may wish to verify the new configuration using the POST /aaa/ldap-identity-sources?action=probe API before changing the configuration.
	//
	// @param ldapIdentitySourceIdParam (required)
	// @param ldapIdentitySourceParam (required)
	// The parameter must contain all the properties defined in model.LdapIdentitySource.
	// @return com.vmware.nsx_policy.model.LdapIdentitySource
	// The return value will contain all the properties defined in model.LdapIdentitySource.
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(ldapIdentitySourceIdParam string, ldapIdentitySourceParam *data.StructValue) (*data.StructValue, error)
}

type ldapIdentitySourcesClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewLdapIdentitySourcesClient(connector client.Connector) *ldapIdentitySourcesClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.nsx_policy.aaa.ldap_identity_sources")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"delete":              core.NewMethodIdentifier(interfaceIdentifier, "delete"),
		"fetchcertificate":    core.NewMethodIdentifier(interfaceIdentifier, "fetchcertificate"),
		"get":                 core.NewMethodIdentifier(interfaceIdentifier, "get"),
		"list":                core.NewMethodIdentifier(interfaceIdentifier, "list"),
		"probe":               core.NewMethodIdentifier(interfaceIdentifier, "probe"),
		"probeidentitysource": core.NewMethodIdentifier(interfaceIdentifier, "probeidentitysource"),
		"probeldapserver":     core.NewMethodIdentifier(interfaceIdentifier, "probeldapserver"),
		"update":              core.NewMethodIdentifier(interfaceIdentifier, "update"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	lIface := ldapIdentitySourcesClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &lIface
}

func (lIface *ldapIdentitySourcesClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := lIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (lIface *ldapIdentitySourcesClient) Delete(ldapIdentitySourceIdParam string) error {
	typeConverter := lIface.connector.TypeConverter()
	executionContext := lIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(ldapIdentitySourcesDeleteInputType(), typeConverter)
	sv.AddStructField("LdapIdentitySourceId", ldapIdentitySourceIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := ldapIdentitySourcesDeleteRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	lIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := lIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.aaa.ldap_identity_sources", "delete", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), lIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (lIface *ldapIdentitySourcesClient) Fetchcertificate(identitySourceLdapServerEndpointParam model.IdentitySourceLdapServerEndpoint) (model.PeerCertificateChain, error) {
	typeConverter := lIface.connector.TypeConverter()
	executionContext := lIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(ldapIdentitySourcesFetchcertificateInputType(), typeConverter)
	sv.AddStructField("IdentitySourceLdapServerEndpoint", identitySourceLdapServerEndpointParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.PeerCertificateChain
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := ldapIdentitySourcesFetchcertificateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	lIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := lIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.aaa.ldap_identity_sources", "fetchcertificate", inputDataValue, executionContext)
	var emptyOutput model.PeerCertificateChain
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), ldapIdentitySourcesFetchcertificateOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.PeerCertificateChain), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), lIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (lIface *ldapIdentitySourcesClient) Get(ldapIdentitySourceIdParam string) (*data.StructValue, error) {
	typeConverter := lIface.connector.TypeConverter()
	executionContext := lIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(ldapIdentitySourcesGetInputType(), typeConverter)
	sv.AddStructField("LdapIdentitySourceId", ldapIdentitySourceIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput *data.StructValue
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := ldapIdentitySourcesGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	lIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := lIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.aaa.ldap_identity_sources", "get", inputDataValue, executionContext)
	var emptyOutput *data.StructValue
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), ldapIdentitySourcesGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(*data.StructValue), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), lIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (lIface *ldapIdentitySourcesClient) List(cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.LdapIdentitySourceListResult, error) {
	typeConverter := lIface.connector.TypeConverter()
	executionContext := lIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(ldapIdentitySourcesListInputType(), typeConverter)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.LdapIdentitySourceListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := ldapIdentitySourcesListRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	lIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := lIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.aaa.ldap_identity_sources", "list", inputDataValue, executionContext)
	var emptyOutput model.LdapIdentitySourceListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), ldapIdentitySourcesListOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.LdapIdentitySourceListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), lIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (lIface *ldapIdentitySourcesClient) Probe(ldapIdentitySourceIdParam string) (model.LdapIdentitySourceProbeResults, error) {
	typeConverter := lIface.connector.TypeConverter()
	executionContext := lIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(ldapIdentitySourcesProbeInputType(), typeConverter)
	sv.AddStructField("LdapIdentitySourceId", ldapIdentitySourceIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.LdapIdentitySourceProbeResults
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := ldapIdentitySourcesProbeRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	lIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := lIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.aaa.ldap_identity_sources", "probe", inputDataValue, executionContext)
	var emptyOutput model.LdapIdentitySourceProbeResults
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), ldapIdentitySourcesProbeOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.LdapIdentitySourceProbeResults), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), lIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (lIface *ldapIdentitySourcesClient) Probeidentitysource(ldapIdentitySourceParam *data.StructValue) (model.LdapIdentitySourceProbeResults, error) {
	typeConverter := lIface.connector.TypeConverter()
	executionContext := lIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(ldapIdentitySourcesProbeidentitysourceInputType(), typeConverter)
	sv.AddStructField("LdapIdentitySource", ldapIdentitySourceParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.LdapIdentitySourceProbeResults
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := ldapIdentitySourcesProbeidentitysourceRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	lIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := lIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.aaa.ldap_identity_sources", "probeidentitysource", inputDataValue, executionContext)
	var emptyOutput model.LdapIdentitySourceProbeResults
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), ldapIdentitySourcesProbeidentitysourceOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.LdapIdentitySourceProbeResults), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), lIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (lIface *ldapIdentitySourcesClient) Probeldapserver(identitySourceLdapServerParam model.IdentitySourceLdapServer) (model.IdentitySourceLdapServerProbeResult, error) {
	typeConverter := lIface.connector.TypeConverter()
	executionContext := lIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(ldapIdentitySourcesProbeldapserverInputType(), typeConverter)
	sv.AddStructField("IdentitySourceLdapServer", identitySourceLdapServerParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.IdentitySourceLdapServerProbeResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := ldapIdentitySourcesProbeldapserverRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	lIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := lIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.aaa.ldap_identity_sources", "probeldapserver", inputDataValue, executionContext)
	var emptyOutput model.IdentitySourceLdapServerProbeResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), ldapIdentitySourcesProbeldapserverOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.IdentitySourceLdapServerProbeResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), lIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (lIface *ldapIdentitySourcesClient) Update(ldapIdentitySourceIdParam string, ldapIdentitySourceParam *data.StructValue) (*data.StructValue, error) {
	typeConverter := lIface.connector.TypeConverter()
	executionContext := lIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(ldapIdentitySourcesUpdateInputType(), typeConverter)
	sv.AddStructField("LdapIdentitySourceId", ldapIdentitySourceIdParam)
	sv.AddStructField("LdapIdentitySource", ldapIdentitySourceParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput *data.StructValue
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := ldapIdentitySourcesUpdateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	lIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := lIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.aaa.ldap_identity_sources", "update", inputDataValue, executionContext)
	var emptyOutput *data.StructValue
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), ldapIdentitySourcesUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(*data.StructValue), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), lIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
