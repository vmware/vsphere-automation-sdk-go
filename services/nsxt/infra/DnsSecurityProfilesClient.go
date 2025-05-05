// Copyright (c) 2019-2025 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: DnsSecurityProfiles
// Used by client-side stubs.

package infra

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	nsx_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type DnsSecurityProfilesClient interface {

	// Delete DNS security profile
	//
	// @param profileIdParam (required)
	// @param overrideParam If true, the global resource can be over written locally. This means that there will be a local only resource in place of the global resource that can reflect local specific settings and values. The global object will continue to exist but will not be used for any configuration until this local object is removed. When the object is overridden the Global resource continues to exist unmodified, while the overridden object is created with all of the user specified values. The Global resource may be updated in the background, however, the overridden object may only be updated by the user. Once the user removes the overridden copy, the Global resource will then resume being used in the configuration. (optional, default to false)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Delete(profileIdParam string, overrideParam *bool) error

	// Read the DNS Forwarder for the given tier-0 instance
	//
	// @param profileIdParam (required)
	// @return com.vmware.nsx_policy.model.DnsSecurityProfile
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(profileIdParam string) (nsx_policyModel.DnsSecurityProfile, error)

	// List all DNS security profiles
	//
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includeMarkForDeleteObjectsParam If true, resources that are marked for deletion will be included in the results. By default, these resources are not included. (optional, default to false)
	// @param includedFieldsParam Note - this parameter currently only works when used with the search APIs /policy/api/v1/search/query and /policy/api/v1/search/dsl. It is ignored for other list APIs. (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam If true, results are sorted in ascending order (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @return com.vmware.nsx_policy.model.DnsSecurityProfileListResult
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(cursorParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (nsx_policyModel.DnsSecurityProfileListResult, error)

	// Create or update DNS security profile
	//
	// @param profileIdParam (required)
	// @param dnsSecurityProfileParam (required)
	// @param overrideParam If true, the global resource can be over written locally. This means that there will be a local only resource in place of the global resource that can reflect local specific settings and values. The global object will continue to exist but will not be used for any configuration until this local object is removed. When the object is overridden the Global resource continues to exist unmodified, while the overridden object is created with all of the user specified values. The Global resource may be updated in the background, however, the overridden object may only be updated by the user. Once the user removes the overridden copy, the Global resource will then resume being used in the configuration. (optional, default to false)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Patch(profileIdParam string, dnsSecurityProfileParam nsx_policyModel.DnsSecurityProfile, overrideParam *bool) error

	// Create or update DNS security profile
	//
	// @param profileIdParam (required)
	// @param dnsSecurityProfileParam (required)
	// @param overrideParam If true, the global resource can be over written locally. This means that there will be a local only resource in place of the global resource that can reflect local specific settings and values. The global object will continue to exist but will not be used for any configuration until this local object is removed. When the object is overridden the Global resource continues to exist unmodified, while the overridden object is created with all of the user specified values. The Global resource may be updated in the background, however, the overridden object may only be updated by the user. Once the user removes the overridden copy, the Global resource will then resume being used in the configuration. (optional, default to false)
	// @return com.vmware.nsx_policy.model.DnsSecurityProfile
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(profileIdParam string, dnsSecurityProfileParam nsx_policyModel.DnsSecurityProfile, overrideParam *bool) (nsx_policyModel.DnsSecurityProfile, error)
}

type dnsSecurityProfilesClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewDnsSecurityProfilesClient(connector vapiProtocolClient_.Connector) *dnsSecurityProfilesClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx_policy.infra.dns_security_profiles")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"delete": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "delete"),
		"get":    vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
		"list":   vapiCore_.NewMethodIdentifier(interfaceIdentifier, "list"),
		"patch":  vapiCore_.NewMethodIdentifier(interfaceIdentifier, "patch"),
		"update": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "update"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	dIface := dnsSecurityProfilesClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &dIface
}

func (dIface *dnsSecurityProfilesClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := dIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (dIface *dnsSecurityProfilesClient) Delete(profileIdParam string, overrideParam *bool) error {
	typeConverter := dIface.connector.TypeConverter()
	executionContext := dIface.connector.NewExecutionContext()
	operationRestMetaData := dnsSecurityProfilesDeleteRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(dnsSecurityProfilesDeleteInputType(), typeConverter)
	sv.AddStructField("ProfileId", profileIdParam)
	sv.AddStructField("Override", overrideParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := dIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.dns_security_profiles", "delete", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), dIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (dIface *dnsSecurityProfilesClient) Get(profileIdParam string) (nsx_policyModel.DnsSecurityProfile, error) {
	typeConverter := dIface.connector.TypeConverter()
	executionContext := dIface.connector.NewExecutionContext()
	operationRestMetaData := dnsSecurityProfilesGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(dnsSecurityProfilesGetInputType(), typeConverter)
	sv.AddStructField("ProfileId", profileIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.DnsSecurityProfile
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := dIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.dns_security_profiles", "get", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.DnsSecurityProfile
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), DnsSecurityProfilesGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.DnsSecurityProfile), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), dIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (dIface *dnsSecurityProfilesClient) List(cursorParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (nsx_policyModel.DnsSecurityProfileListResult, error) {
	typeConverter := dIface.connector.TypeConverter()
	executionContext := dIface.connector.NewExecutionContext()
	operationRestMetaData := dnsSecurityProfilesListRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(dnsSecurityProfilesListInputType(), typeConverter)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludeMarkForDeleteObjects", includeMarkForDeleteObjectsParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.DnsSecurityProfileListResult
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := dIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.dns_security_profiles", "list", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.DnsSecurityProfileListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), DnsSecurityProfilesListOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.DnsSecurityProfileListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), dIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (dIface *dnsSecurityProfilesClient) Patch(profileIdParam string, dnsSecurityProfileParam nsx_policyModel.DnsSecurityProfile, overrideParam *bool) error {
	typeConverter := dIface.connector.TypeConverter()
	executionContext := dIface.connector.NewExecutionContext()
	operationRestMetaData := dnsSecurityProfilesPatchRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(dnsSecurityProfilesPatchInputType(), typeConverter)
	sv.AddStructField("ProfileId", profileIdParam)
	sv.AddStructField("DnsSecurityProfile", dnsSecurityProfileParam)
	sv.AddStructField("Override", overrideParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := dIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.dns_security_profiles", "patch", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), dIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (dIface *dnsSecurityProfilesClient) Update(profileIdParam string, dnsSecurityProfileParam nsx_policyModel.DnsSecurityProfile, overrideParam *bool) (nsx_policyModel.DnsSecurityProfile, error) {
	typeConverter := dIface.connector.TypeConverter()
	executionContext := dIface.connector.NewExecutionContext()
	operationRestMetaData := dnsSecurityProfilesUpdateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(dnsSecurityProfilesUpdateInputType(), typeConverter)
	sv.AddStructField("ProfileId", profileIdParam)
	sv.AddStructField("DnsSecurityProfile", dnsSecurityProfileParam)
	sv.AddStructField("Override", overrideParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.DnsSecurityProfile
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := dIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.dns_security_profiles", "update", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.DnsSecurityProfile
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), DnsSecurityProfilesUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.DnsSecurityProfile), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), dIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
