// Copyright (c) 2019-2026 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: HostVmkAllocations
// Used by client-side stubs.

package host_network_configs

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	nsx_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type HostVmkAllocationsClient interface {

	// Delete a Host Vmk Allocation by ID and release the IP addresses allocated to the vmks.
	//
	// @param siteIdParam (required)
	// @param enforcementpointIdParam (required)
	// @param hostNetworkConfigsIdParam (required)
	// @param allocationIdParam (required)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Delete(siteIdParam string, enforcementpointIdParam string, hostNetworkConfigsIdParam string, allocationIdParam string) error

	// Get a Host Vmk Allocation by ID.
	//
	// @param siteIdParam (required)
	// @param enforcementpointIdParam (required)
	// @param hostNetworkConfigsIdParam (required)
	// @param allocationIdParam (required)
	// @return com.vmware.nsx_policy.model.HostVmkAllocation
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(siteIdParam string, enforcementpointIdParam string, hostNetworkConfigsIdParam string, allocationIdParam string) (nsx_policyModel.HostVmkAllocation, error)

	// List all host vmknic allocations for a given host network config.
	//
	// @param siteIdParam (required)
	// @param enforcementpointIdParam (required)
	// @param hostNetworkConfigsIdParam (required)
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param deviceIdParam Device ID of the vmknic to get the network id for. e.g. vmk0, vmk1. (optional)
	// @param hostExternalIdParam The external ID of the specific host in NSX. This is in the format compute-manager-uuid:host-moid and can be retrieved from the discovered node API in NSX, GET https://&lt;policy-mgr&gt;/policy/api/v1/fabric/discovered-nodes/ e.g. 5a0b7b11-4475-4b72-8822-225e348f4b4e:host-10 (optional)
	// @param includeConflictsParam If true, resources currently in a sandboxed state due to federation conflicts will be included in the list results alongside active intent. Sandboxed resources will have has_conflict=true in the response. Applicable on LM only. On FNS, all resources including conflicting ones are always returned regardless of this parameter. Default is false. (optional, default to false)
	// @param includeMarkForDeleteObjectsParam If true, resources that are marked for deletion will be included in the results. By default, these resources are not included. (optional, default to false)
	// @param includedFieldsParam Note - this parameter currently only works when used with the search APIs /policy/api/v1/search/query and /policy/api/v1/search/dsl. It is ignored for other list APIs. (optional)
	// @param infraNetworkPathParam Policy path of the Infra Network to filter vmk allocations by. (optional)
	// @param ipPoolPathParam Policy path of the IPv4 or IPv6 pool (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam If true, results are sorted in ascending order (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @return com.vmware.nsx_policy.model.HostVmkAllocationListResult
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(siteIdParam string, enforcementpointIdParam string, hostNetworkConfigsIdParam string, cursorParam *string, deviceIdParam *string, hostExternalIdParam *string, includeConflictsParam *bool, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, infraNetworkPathParam *string, ipPoolPathParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (nsx_policyModel.HostVmkAllocationListResult, error)

	// Patch a Host Vmk Allocation which specifies how vmk on infraNetwork is allocated and deallocated ipv4 and ipv6 addresses. Note: The 'vmk_allocations' list should always include the complete list of desired allocations. Any missing allocations from the list will be considered as not desired and freed up.
	//
	// @param siteIdParam (required)
	// @param enforcementpointIdParam (required)
	// @param hostNetworkConfigsIdParam (required)
	// @param allocationIdParam (required)
	// @param hostVmkAllocationParam (required)
	// @return com.vmware.nsx_policy.model.HostVmkAllocation
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Patch(siteIdParam string, enforcementpointIdParam string, hostNetworkConfigsIdParam string, allocationIdParam string, hostVmkAllocationParam nsx_policyModel.HostVmkAllocation) (nsx_policyModel.HostVmkAllocation, error)

	// Create or update a Host Vmk Allocation which specifies how vmk on infraNetwork is allocated and deallocated ipv4 and ipv6 addresses. Note: The 'vmk_allocations' list should always include the complete list of desired allocations. Any missing allocations from the list will be considered as not desired and freed up.
	//
	// @param siteIdParam (required)
	// @param enforcementpointIdParam (required)
	// @param hostNetworkConfigsIdParam (required)
	// @param allocationIdParam (required)
	// @param hostVmkAllocationParam (required)
	// @return com.vmware.nsx_policy.model.HostVmkAllocation
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(siteIdParam string, enforcementpointIdParam string, hostNetworkConfigsIdParam string, allocationIdParam string, hostVmkAllocationParam nsx_policyModel.HostVmkAllocation) (nsx_policyModel.HostVmkAllocation, error)
}

type hostVmkAllocationsClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewHostVmkAllocationsClient(connector vapiProtocolClient_.Connector) *hostVmkAllocationsClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx_policy.infra.sites.enforcement_points.host_network_configs.host_vmk_allocations")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"delete": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "delete"),
		"get":    vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
		"list":   vapiCore_.NewMethodIdentifier(interfaceIdentifier, "list"),
		"patch":  vapiCore_.NewMethodIdentifier(interfaceIdentifier, "patch"),
		"update": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "update"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	hIface := hostVmkAllocationsClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &hIface
}

func (hIface *hostVmkAllocationsClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := hIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (hIface *hostVmkAllocationsClient) Delete(siteIdParam string, enforcementpointIdParam string, hostNetworkConfigsIdParam string, allocationIdParam string) error {
	typeConverter := hIface.connector.TypeConverter()
	executionContext := hIface.connector.NewExecutionContext()
	operationRestMetaData := hostVmkAllocationsDeleteRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(hostVmkAllocationsDeleteInputType(), typeConverter)
	sv.AddStructField("SiteId", siteIdParam)
	sv.AddStructField("EnforcementpointId", enforcementpointIdParam)
	sv.AddStructField("HostNetworkConfigsId", hostNetworkConfigsIdParam)
	sv.AddStructField("AllocationId", allocationIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := hIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.sites.enforcement_points.host_network_configs.host_vmk_allocations", "delete", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), hIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (hIface *hostVmkAllocationsClient) Get(siteIdParam string, enforcementpointIdParam string, hostNetworkConfigsIdParam string, allocationIdParam string) (nsx_policyModel.HostVmkAllocation, error) {
	typeConverter := hIface.connector.TypeConverter()
	executionContext := hIface.connector.NewExecutionContext()
	operationRestMetaData := hostVmkAllocationsGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(hostVmkAllocationsGetInputType(), typeConverter)
	sv.AddStructField("SiteId", siteIdParam)
	sv.AddStructField("EnforcementpointId", enforcementpointIdParam)
	sv.AddStructField("HostNetworkConfigsId", hostNetworkConfigsIdParam)
	sv.AddStructField("AllocationId", allocationIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.HostVmkAllocation
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := hIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.sites.enforcement_points.host_network_configs.host_vmk_allocations", "get", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.HostVmkAllocation
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), HostVmkAllocationsGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.HostVmkAllocation), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), hIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (hIface *hostVmkAllocationsClient) List(siteIdParam string, enforcementpointIdParam string, hostNetworkConfigsIdParam string, cursorParam *string, deviceIdParam *string, hostExternalIdParam *string, includeConflictsParam *bool, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, infraNetworkPathParam *string, ipPoolPathParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (nsx_policyModel.HostVmkAllocationListResult, error) {
	typeConverter := hIface.connector.TypeConverter()
	executionContext := hIface.connector.NewExecutionContext()
	operationRestMetaData := hostVmkAllocationsListRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(hostVmkAllocationsListInputType(), typeConverter)
	sv.AddStructField("SiteId", siteIdParam)
	sv.AddStructField("EnforcementpointId", enforcementpointIdParam)
	sv.AddStructField("HostNetworkConfigsId", hostNetworkConfigsIdParam)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("DeviceId", deviceIdParam)
	sv.AddStructField("HostExternalId", hostExternalIdParam)
	sv.AddStructField("IncludeConflicts", includeConflictsParam)
	sv.AddStructField("IncludeMarkForDeleteObjects", includeMarkForDeleteObjectsParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("InfraNetworkPath", infraNetworkPathParam)
	sv.AddStructField("IpPoolPath", ipPoolPathParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.HostVmkAllocationListResult
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := hIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.sites.enforcement_points.host_network_configs.host_vmk_allocations", "list", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.HostVmkAllocationListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), HostVmkAllocationsListOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.HostVmkAllocationListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), hIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (hIface *hostVmkAllocationsClient) Patch(siteIdParam string, enforcementpointIdParam string, hostNetworkConfigsIdParam string, allocationIdParam string, hostVmkAllocationParam nsx_policyModel.HostVmkAllocation) (nsx_policyModel.HostVmkAllocation, error) {
	typeConverter := hIface.connector.TypeConverter()
	executionContext := hIface.connector.NewExecutionContext()
	operationRestMetaData := hostVmkAllocationsPatchRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(hostVmkAllocationsPatchInputType(), typeConverter)
	sv.AddStructField("SiteId", siteIdParam)
	sv.AddStructField("EnforcementpointId", enforcementpointIdParam)
	sv.AddStructField("HostNetworkConfigsId", hostNetworkConfigsIdParam)
	sv.AddStructField("AllocationId", allocationIdParam)
	sv.AddStructField("HostVmkAllocation", hostVmkAllocationParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.HostVmkAllocation
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := hIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.sites.enforcement_points.host_network_configs.host_vmk_allocations", "patch", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.HostVmkAllocation
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), HostVmkAllocationsPatchOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.HostVmkAllocation), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), hIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (hIface *hostVmkAllocationsClient) Update(siteIdParam string, enforcementpointIdParam string, hostNetworkConfigsIdParam string, allocationIdParam string, hostVmkAllocationParam nsx_policyModel.HostVmkAllocation) (nsx_policyModel.HostVmkAllocation, error) {
	typeConverter := hIface.connector.TypeConverter()
	executionContext := hIface.connector.NewExecutionContext()
	operationRestMetaData := hostVmkAllocationsUpdateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(hostVmkAllocationsUpdateInputType(), typeConverter)
	sv.AddStructField("SiteId", siteIdParam)
	sv.AddStructField("EnforcementpointId", enforcementpointIdParam)
	sv.AddStructField("HostNetworkConfigsId", hostNetworkConfigsIdParam)
	sv.AddStructField("AllocationId", allocationIdParam)
	sv.AddStructField("HostVmkAllocation", hostVmkAllocationParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.HostVmkAllocation
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := hIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.sites.enforcement_points.host_network_configs.host_vmk_allocations", "update", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.HostVmkAllocation
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), HostVmkAllocationsUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.HostVmkAllocation), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), hIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
