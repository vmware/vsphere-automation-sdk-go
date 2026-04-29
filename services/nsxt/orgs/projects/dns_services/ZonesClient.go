// Copyright (c) 2019-2026 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Zones
// Used by client-side stubs.

package dns_services

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	nsx_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type ZonesClient interface {

	// Delete a ProjectDnsZone. The zone cannot be deleted if any DnsRecord resources reference it via zone_path, or if it is currently shared via a Share resource. Remove all DNS records referencing this zone and all Share resources before deleting.
	//
	// @param orgIdParam (required)
	// @param projectIdParam (required)
	// @param dnsServiceIdParam (required)
	// @param zoneIdParam (required)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Delete(orgIdParam string, projectIdParam string, dnsServiceIdParam string, zoneIdParam string) error

	// Read a ProjectDnsZone by ID within the specified PolicyDnsService.
	//
	// @param orgIdParam (required)
	// @param projectIdParam (required)
	// @param dnsServiceIdParam (required)
	// @param zoneIdParam (required)
	// @return com.vmware.nsx_policy.model.ProjectDnsZone
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(orgIdParam string, projectIdParam string, dnsServiceIdParam string, zoneIdParam string) (nsx_policyModel.ProjectDnsZone, error)

	// List all ProjectDnsZone resources under the specified PolicyDnsService.
	//
	// @param orgIdParam (required)
	// @param projectIdParam (required)
	// @param dnsServiceIdParam (required)
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includeMarkForDeleteObjectsParam If true, resources that are marked for deletion will be included in the results. By default, these resources are not included. (optional, default to false)
	// @param includedFieldsParam Note - this parameter currently only works when used with the search APIs /policy/api/v1/search/query and /policy/api/v1/search/dsl. It is ignored for other list APIs. (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam If true, results are sorted in ascending order (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @return com.vmware.nsx_policy.model.ProjectDnsZoneListResult
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(orgIdParam string, projectIdParam string, dnsServiceIdParam string, cursorParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (nsx_policyModel.ProjectDnsZoneListResult, error)

	// Patch a ProjectDnsZone. Only provided fields are updated. The dns_domain_name field is immutable and cannot be changed after creation.
	//
	// @param orgIdParam (required)
	// @param projectIdParam (required)
	// @param dnsServiceIdParam (required)
	// @param zoneIdParam (required)
	// @param projectDnsZoneParam (required)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Patch(orgIdParam string, projectIdParam string, dnsServiceIdParam string, zoneIdParam string, projectDnsZoneParam nsx_policyModel.ProjectDnsZone) error

	// Create or update a ProjectDnsZone under the specified PolicyDnsService. The dns_domain_name is immutable after creation. Domain names must be unique within the parent DNS service. Supports forward zones (e.g. \"example.com\") and reverse zones (e.g. \"12.168.192.in-addr.arpa\").
	//
	// @param orgIdParam (required)
	// @param projectIdParam (required)
	// @param dnsServiceIdParam (required)
	// @param zoneIdParam (required)
	// @param projectDnsZoneParam (required)
	// @return com.vmware.nsx_policy.model.ProjectDnsZone
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(orgIdParam string, projectIdParam string, dnsServiceIdParam string, zoneIdParam string, projectDnsZoneParam nsx_policyModel.ProjectDnsZone) (nsx_policyModel.ProjectDnsZone, error)
}

type zonesClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewZonesClient(connector vapiProtocolClient_.Connector) *zonesClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx_policy.orgs.projects.dns_services.zones")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"delete": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "delete"),
		"get":    vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
		"list":   vapiCore_.NewMethodIdentifier(interfaceIdentifier, "list"),
		"patch":  vapiCore_.NewMethodIdentifier(interfaceIdentifier, "patch"),
		"update": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "update"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	zIface := zonesClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &zIface
}

func (zIface *zonesClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := zIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (zIface *zonesClient) Delete(orgIdParam string, projectIdParam string, dnsServiceIdParam string, zoneIdParam string) error {
	typeConverter := zIface.connector.TypeConverter()
	executionContext := zIface.connector.NewExecutionContext()
	operationRestMetaData := zonesDeleteRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(zonesDeleteInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("ProjectId", projectIdParam)
	sv.AddStructField("DnsServiceId", dnsServiceIdParam)
	sv.AddStructField("ZoneId", zoneIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := zIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.orgs.projects.dns_services.zones", "delete", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), zIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (zIface *zonesClient) Get(orgIdParam string, projectIdParam string, dnsServiceIdParam string, zoneIdParam string) (nsx_policyModel.ProjectDnsZone, error) {
	typeConverter := zIface.connector.TypeConverter()
	executionContext := zIface.connector.NewExecutionContext()
	operationRestMetaData := zonesGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(zonesGetInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("ProjectId", projectIdParam)
	sv.AddStructField("DnsServiceId", dnsServiceIdParam)
	sv.AddStructField("ZoneId", zoneIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.ProjectDnsZone
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := zIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.orgs.projects.dns_services.zones", "get", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.ProjectDnsZone
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), ZonesGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.ProjectDnsZone), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), zIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (zIface *zonesClient) List(orgIdParam string, projectIdParam string, dnsServiceIdParam string, cursorParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (nsx_policyModel.ProjectDnsZoneListResult, error) {
	typeConverter := zIface.connector.TypeConverter()
	executionContext := zIface.connector.NewExecutionContext()
	operationRestMetaData := zonesListRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(zonesListInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("ProjectId", projectIdParam)
	sv.AddStructField("DnsServiceId", dnsServiceIdParam)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludeMarkForDeleteObjects", includeMarkForDeleteObjectsParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.ProjectDnsZoneListResult
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := zIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.orgs.projects.dns_services.zones", "list", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.ProjectDnsZoneListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), ZonesListOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.ProjectDnsZoneListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), zIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (zIface *zonesClient) Patch(orgIdParam string, projectIdParam string, dnsServiceIdParam string, zoneIdParam string, projectDnsZoneParam nsx_policyModel.ProjectDnsZone) error {
	typeConverter := zIface.connector.TypeConverter()
	executionContext := zIface.connector.NewExecutionContext()
	operationRestMetaData := zonesPatchRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(zonesPatchInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("ProjectId", projectIdParam)
	sv.AddStructField("DnsServiceId", dnsServiceIdParam)
	sv.AddStructField("ZoneId", zoneIdParam)
	sv.AddStructField("ProjectDnsZone", projectDnsZoneParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := zIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.orgs.projects.dns_services.zones", "patch", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), zIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (zIface *zonesClient) Update(orgIdParam string, projectIdParam string, dnsServiceIdParam string, zoneIdParam string, projectDnsZoneParam nsx_policyModel.ProjectDnsZone) (nsx_policyModel.ProjectDnsZone, error) {
	typeConverter := zIface.connector.TypeConverter()
	executionContext := zIface.connector.NewExecutionContext()
	operationRestMetaData := zonesUpdateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(zonesUpdateInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("ProjectId", projectIdParam)
	sv.AddStructField("DnsServiceId", dnsServiceIdParam)
	sv.AddStructField("ZoneId", zoneIdParam)
	sv.AddStructField("ProjectDnsZone", projectDnsZoneParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.ProjectDnsZone
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := zIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.orgs.projects.dns_services.zones", "update", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.ProjectDnsZone
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), ZonesUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.ProjectDnsZone), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), zIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
