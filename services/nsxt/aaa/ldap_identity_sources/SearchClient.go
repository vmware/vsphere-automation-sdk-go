// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Search
// Used by client-side stubs.

package ldap_identity_sources

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/core"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/lib"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/protocol/client"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt/model"
)

const _ = core.SupportedByRuntimeVersion1

type SearchClient interface {

	// Search the LDAP identity source for users and groups that match the given filter_value. In most cases, the LDAP source performs a case-insensitive search.
	//
	// @param ldapIdentitySourceIdParam (required)
	// @param filterValueParam Search filter value (required)
	// @return com.vmware.nsx_policy.model.LdapIdentitySourceSearchResultList
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Create(ldapIdentitySourceIdParam string, filterValueParam string) (model.LdapIdentitySourceSearchResultList, error)
}

type searchClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewSearchClient(connector client.Connector) *searchClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.nsx_policy.aaa.ldap_identity_sources.search")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"create": core.NewMethodIdentifier(interfaceIdentifier, "create"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	sIface := searchClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &sIface
}

func (sIface *searchClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := sIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (sIface *searchClient) Create(ldapIdentitySourceIdParam string, filterValueParam string) (model.LdapIdentitySourceSearchResultList, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(searchCreateInputType(), typeConverter)
	sv.AddStructField("LdapIdentitySourceId", ldapIdentitySourceIdParam)
	sv.AddStructField("FilterValue", filterValueParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.LdapIdentitySourceSearchResultList
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := searchCreateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.aaa.ldap_identity_sources.search", "create", inputDataValue, executionContext)
	var emptyOutput model.LdapIdentitySourceSearchResultList
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), searchCreateOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.LdapIdentitySourceSearchResultList), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
