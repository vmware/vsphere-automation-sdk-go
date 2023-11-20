// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: DnsSecurityProfileBindingMaps
// Used by client-side stubs.

package groups

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	nsx_global_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-gm/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type DnsSecurityProfileBindingMapsClient interface {

	// API will delete DNS security profile binding map
	//
	// @param domainIdParam Domain ID (required)
	// @param groupIdParam Group ID (required)
	// @param dnsSecurityProfileBindingMapIdParam DNS security profile binding map ID (required)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Delete(domainIdParam string, groupIdParam string, dnsSecurityProfileBindingMapIdParam string) error

	// API will get DNS security profile binding map
	//
	// @param domainIdParam Domain ID (required)
	// @param groupIdParam Group ID (required)
	// @param dnsSecurityProfileBindingMapIdParam DNS security profile binding map ID (required)
	// @return com.vmware.nsx_global_policy.model.DnsSecurityProfileBindingMap
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(domainIdParam string, groupIdParam string, dnsSecurityProfileBindingMapIdParam string) (nsx_global_policyModel.DnsSecurityProfileBindingMap, error)

	// API will get DNS security profile binding map
	//
	// @param domainIdParam (required)
	// @param groupIdParam (required)
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includeMarkForDeleteObjectsParam Include objects that are marked for deletion in results (optional, default to false)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @return com.vmware.nsx_global_policy.model.DnsSecurityProfileBindingMapListResult
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(domainIdParam string, groupIdParam string, cursorParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (nsx_global_policyModel.DnsSecurityProfileBindingMapListResult, error)

	// API will create or update DNS security profile binding map
	//
	// @param domainIdParam Domain ID (required)
	// @param groupIdParam Group ID (required)
	// @param dnsSecurityProfileBindingMapIdParam DNS security profile binding map ID (required)
	// @param dnsSecurityProfileBindingMapParam (required)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Patch(domainIdParam string, groupIdParam string, dnsSecurityProfileBindingMapIdParam string, dnsSecurityProfileBindingMapParam nsx_global_policyModel.DnsSecurityProfileBindingMap) error

	// API will update DNS security profile binding map
	//
	// @param domainIdParam Domain ID (required)
	// @param groupIdParam Group ID (required)
	// @param dnsSecurityProfileBindingMapIdParam DNS security profile binding map ID (required)
	// @param dnsSecurityProfileBindingMapParam (required)
	// @return com.vmware.nsx_global_policy.model.DnsSecurityProfileBindingMap
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(domainIdParam string, groupIdParam string, dnsSecurityProfileBindingMapIdParam string, dnsSecurityProfileBindingMapParam nsx_global_policyModel.DnsSecurityProfileBindingMap) (nsx_global_policyModel.DnsSecurityProfileBindingMap, error)
}

type dnsSecurityProfileBindingMapsClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewDnsSecurityProfileBindingMapsClient(connector vapiProtocolClient_.Connector) *dnsSecurityProfileBindingMapsClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx_global_policy.global_infra.domains.groups.dns_security_profile_binding_maps")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"delete": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "delete"),
		"get":    vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
		"list":   vapiCore_.NewMethodIdentifier(interfaceIdentifier, "list"),
		"patch":  vapiCore_.NewMethodIdentifier(interfaceIdentifier, "patch"),
		"update": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "update"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	dIface := dnsSecurityProfileBindingMapsClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &dIface
}

func (dIface *dnsSecurityProfileBindingMapsClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := dIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (dIface *dnsSecurityProfileBindingMapsClient) Delete(domainIdParam string, groupIdParam string, dnsSecurityProfileBindingMapIdParam string) error {
	typeConverter := dIface.connector.TypeConverter()
	executionContext := dIface.connector.NewExecutionContext()
	operationRestMetaData := dnsSecurityProfileBindingMapsDeleteRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(dnsSecurityProfileBindingMapsDeleteInputType(), typeConverter)
	sv.AddStructField("DomainId", domainIdParam)
	sv.AddStructField("GroupId", groupIdParam)
	sv.AddStructField("DnsSecurityProfileBindingMapId", dnsSecurityProfileBindingMapIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := dIface.connector.GetApiProvider().Invoke("com.vmware.nsx_global_policy.global_infra.domains.groups.dns_security_profile_binding_maps", "delete", inputDataValue, executionContext)
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

func (dIface *dnsSecurityProfileBindingMapsClient) Get(domainIdParam string, groupIdParam string, dnsSecurityProfileBindingMapIdParam string) (nsx_global_policyModel.DnsSecurityProfileBindingMap, error) {
	typeConverter := dIface.connector.TypeConverter()
	executionContext := dIface.connector.NewExecutionContext()
	operationRestMetaData := dnsSecurityProfileBindingMapsGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(dnsSecurityProfileBindingMapsGetInputType(), typeConverter)
	sv.AddStructField("DomainId", domainIdParam)
	sv.AddStructField("GroupId", groupIdParam)
	sv.AddStructField("DnsSecurityProfileBindingMapId", dnsSecurityProfileBindingMapIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_global_policyModel.DnsSecurityProfileBindingMap
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := dIface.connector.GetApiProvider().Invoke("com.vmware.nsx_global_policy.global_infra.domains.groups.dns_security_profile_binding_maps", "get", inputDataValue, executionContext)
	var emptyOutput nsx_global_policyModel.DnsSecurityProfileBindingMap
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), DnsSecurityProfileBindingMapsGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_global_policyModel.DnsSecurityProfileBindingMap), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), dIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (dIface *dnsSecurityProfileBindingMapsClient) List(domainIdParam string, groupIdParam string, cursorParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (nsx_global_policyModel.DnsSecurityProfileBindingMapListResult, error) {
	typeConverter := dIface.connector.TypeConverter()
	executionContext := dIface.connector.NewExecutionContext()
	operationRestMetaData := dnsSecurityProfileBindingMapsListRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(dnsSecurityProfileBindingMapsListInputType(), typeConverter)
	sv.AddStructField("DomainId", domainIdParam)
	sv.AddStructField("GroupId", groupIdParam)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludeMarkForDeleteObjects", includeMarkForDeleteObjectsParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_global_policyModel.DnsSecurityProfileBindingMapListResult
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := dIface.connector.GetApiProvider().Invoke("com.vmware.nsx_global_policy.global_infra.domains.groups.dns_security_profile_binding_maps", "list", inputDataValue, executionContext)
	var emptyOutput nsx_global_policyModel.DnsSecurityProfileBindingMapListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), DnsSecurityProfileBindingMapsListOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_global_policyModel.DnsSecurityProfileBindingMapListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), dIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (dIface *dnsSecurityProfileBindingMapsClient) Patch(domainIdParam string, groupIdParam string, dnsSecurityProfileBindingMapIdParam string, dnsSecurityProfileBindingMapParam nsx_global_policyModel.DnsSecurityProfileBindingMap) error {
	typeConverter := dIface.connector.TypeConverter()
	executionContext := dIface.connector.NewExecutionContext()
	operationRestMetaData := dnsSecurityProfileBindingMapsPatchRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(dnsSecurityProfileBindingMapsPatchInputType(), typeConverter)
	sv.AddStructField("DomainId", domainIdParam)
	sv.AddStructField("GroupId", groupIdParam)
	sv.AddStructField("DnsSecurityProfileBindingMapId", dnsSecurityProfileBindingMapIdParam)
	sv.AddStructField("DnsSecurityProfileBindingMap", dnsSecurityProfileBindingMapParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := dIface.connector.GetApiProvider().Invoke("com.vmware.nsx_global_policy.global_infra.domains.groups.dns_security_profile_binding_maps", "patch", inputDataValue, executionContext)
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

func (dIface *dnsSecurityProfileBindingMapsClient) Update(domainIdParam string, groupIdParam string, dnsSecurityProfileBindingMapIdParam string, dnsSecurityProfileBindingMapParam nsx_global_policyModel.DnsSecurityProfileBindingMap) (nsx_global_policyModel.DnsSecurityProfileBindingMap, error) {
	typeConverter := dIface.connector.TypeConverter()
	executionContext := dIface.connector.NewExecutionContext()
	operationRestMetaData := dnsSecurityProfileBindingMapsUpdateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(dnsSecurityProfileBindingMapsUpdateInputType(), typeConverter)
	sv.AddStructField("DomainId", domainIdParam)
	sv.AddStructField("GroupId", groupIdParam)
	sv.AddStructField("DnsSecurityProfileBindingMapId", dnsSecurityProfileBindingMapIdParam)
	sv.AddStructField("DnsSecurityProfileBindingMap", dnsSecurityProfileBindingMapParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_global_policyModel.DnsSecurityProfileBindingMap
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := dIface.connector.GetApiProvider().Invoke("com.vmware.nsx_global_policy.global_infra.domains.groups.dns_security_profile_binding_maps", "update", inputDataValue, executionContext)
	var emptyOutput nsx_global_policyModel.DnsSecurityProfileBindingMap
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), DnsSecurityProfileBindingMapsUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_global_policyModel.DnsSecurityProfileBindingMap), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), dIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
