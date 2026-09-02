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

	// Delete a DnsZone. The zone cannot be deleted if any DnsRecord resources reference it via zone_path, or if it is currently shared via a Share resource. Remove all DNS records referencing this zone and all Share resources before deleting.
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

	// Read a DnsZone by ID within the specified DnsService.
	//
	// @param orgIdParam (required)
	// @param projectIdParam (required)
	// @param dnsServiceIdParam (required)
	// @param zoneIdParam (required)
	// @return com.vmware.nsx_policy.model.DnsZone
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(orgIdParam string, projectIdParam string, dnsServiceIdParam string, zoneIdParam string) (nsx_policyModel.DnsZone, error)

	// List all DnsZone resources under the specified DnsService.
	//
	// @param orgIdParam (required)
	// @param projectIdParam (required)
	// @param dnsServiceIdParam (required)
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includeConflictsParam If true, resources currently in a sandboxed state due to federation conflicts will be included in the list results alongside active intent. Sandboxed resources will have has_conflict=true in the response. Applicable on LM only. On FNS, all resources including conflicting ones are always returned regardless of this parameter. Default is false. (optional, default to false)
	// @param includeMarkForDeleteObjectsParam If true, resources that are marked for deletion will be included in the results. By default, these resources are not included. (optional, default to false)
	// @param includedFieldsParam Note - this parameter currently only works when used with the search APIs /policy/api/v1/search/query and /policy/api/v1/search/dsl. It is ignored for other list APIs. (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam If true, results are sorted in ascending order (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @return com.vmware.nsx_policy.model.DnsZoneListResult
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(orgIdParam string, projectIdParam string, dnsServiceIdParam string, cursorParam *string, includeConflictsParam *bool, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (nsx_policyModel.DnsZoneListResult, error)

	// Patch a DnsZone. Only provided fields are updated. The dns_domain_name field is immutable and cannot be changed after creation.
	//
	// @param orgIdParam (required)
	// @param projectIdParam (required)
	// @param dnsServiceIdParam (required)
	// @param zoneIdParam (required)
	// @param dnsZoneParam (required)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Patch(orgIdParam string, projectIdParam string, dnsServiceIdParam string, zoneIdParam string, dnsZoneParam nsx_policyModel.DnsZone) error

	// Create or update a DnsZone under the specified DnsService. The dns_domain_name is immutable after creation. Domain names must be unique within the parent DNS service. Supports forward zones (e.g. \"example.com\") and reverse zones (e.g. \"12.168.192.in-addr.arpa\").
	//
	// @param orgIdParam (required)
	// @param projectIdParam (required)
	// @param dnsServiceIdParam (required)
	// @param zoneIdParam (required)
	// @param dnsZoneParam (required)
	// @return com.vmware.nsx_policy.model.DnsZone
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(orgIdParam string, projectIdParam string, dnsServiceIdParam string, zoneIdParam string, dnsZoneParam nsx_policyModel.DnsZone) (nsx_policyModel.DnsZone, error)
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

func (zIface *zonesClient) Get(orgIdParam string, projectIdParam string, dnsServiceIdParam string, zoneIdParam string) (nsx_policyModel.DnsZone, error) {
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
		var emptyOutput nsx_policyModel.DnsZone
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := zIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.orgs.projects.dns_services.zones", "get", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.DnsZone
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), ZonesGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.DnsZone), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), zIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (zIface *zonesClient) List(orgIdParam string, projectIdParam string, dnsServiceIdParam string, cursorParam *string, includeConflictsParam *bool, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (nsx_policyModel.DnsZoneListResult, error) {
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
	sv.AddStructField("IncludeConflicts", includeConflictsParam)
	sv.AddStructField("IncludeMarkForDeleteObjects", includeMarkForDeleteObjectsParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.DnsZoneListResult
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := zIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.orgs.projects.dns_services.zones", "list", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.DnsZoneListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), ZonesListOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.DnsZoneListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), zIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (zIface *zonesClient) Patch(orgIdParam string, projectIdParam string, dnsServiceIdParam string, zoneIdParam string, dnsZoneParam nsx_policyModel.DnsZone) error {
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
	sv.AddStructField("DnsZone", dnsZoneParam)
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

func (zIface *zonesClient) Update(orgIdParam string, projectIdParam string, dnsServiceIdParam string, zoneIdParam string, dnsZoneParam nsx_policyModel.DnsZone) (nsx_policyModel.DnsZone, error) {
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
	sv.AddStructField("DnsZone", dnsZoneParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.DnsZone
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := zIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.orgs.projects.dns_services.zones", "update", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.DnsZone
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), ZonesUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.DnsZone), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), zIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
