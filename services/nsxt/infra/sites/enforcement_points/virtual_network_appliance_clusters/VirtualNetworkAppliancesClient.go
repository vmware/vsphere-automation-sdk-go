// Copyright (c) 2019-2025 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: VirtualNetworkAppliances
// Used by client-side stubs.

package virtual_network_appliance_clusters

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	nsx_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type VirtualNetworkAppliancesClient interface {

	// This API is used to delete the virtual network appliance (VNA). When \"force\" is used as a query parameter, the system will attempt to delete the VNA from vCenter. This \"best effort\" approach ensures the VNA is not left in a stale state, even if failures occur during the deletion process. However, VNA cannot be deleted if it hosts any logical networking entities, even when using \"force\". After a successful force delete API execution, user must confirm that the VNA virtual machine has been removed from the vCenter Server.
	//
	// @param siteIdParam (required)
	// @param enforcementpointIdParam (required)
	// @param virtualNetworkApplianceClusterIdParam (required)
	// @param virtualNetworkApplianceIdParam (required)
	// @param forceParam If the delete API fails to delete the virtual network appliance, this flag is used to forcibly delete the stale virtual network appliance from the system. (optional, default to false)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Delete(siteIdParam string, enforcementpointIdParam string, virtualNetworkApplianceClusterIdParam string, virtualNetworkApplianceIdParam string, forceParam *bool) error

	// Read a virtual network appliance(VNA) under an Enforcement Point
	//
	// @param siteIdParam (required)
	// @param enforcementpointIdParam (required)
	// @param virtualNetworkApplianceClusterIdParam (required)
	// @param virtualNetworkApplianceIdParam (required)
	// @return com.vmware.nsx_policy.model.BaseNetworkAppliance
	// The return value will contain all the properties defined in nsx_policyModel.BaseNetworkAppliance.
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(siteIdParam string, enforcementpointIdParam string, virtualNetworkApplianceClusterIdParam string, virtualNetworkApplianceIdParam string) (*vapiData_.StructValue, error)

	// List virtual network appliance(VNA) under an Enforcement Point.
	//
	// @param siteIdParam (required)
	// @param enforcementpointIdParam (required)
	// @param virtualNetworkApplianceClusterIdParam (required)
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param inMaintenanceModeParam If the flag is true then virtual network appliances (VNA) with maintenance mode 'ENABLED' in desired state will be returned, otherwise virtual network appliances (VNA) in 'DISABLED' desired state will be returned. (optional)
	// @param includeMarkForDeleteObjectsParam If true, resources that are marked for deletion will be included in the results. By default, these resources are not included. (optional, default to false)
	// @param includedFieldsParam Note - this parameter currently only works when used with the search APIs /policy/api/v1/search/query and /policy/api/v1/search/dsl. It is ignored for other list APIs. (optional)
	// @param managementIpParam Virtual network appliance(VNA) with provided management IP address will be returned. This property can only be used alone. It can not be combined with other filtering properties. (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam If true, results are sorted in ascending order (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @param transportZonePathParam Virtual Network Appliance with provided transport zone path will be returned. (optional)
	// @return com.vmware.nsx_policy.model.VirtualNetworkApplianceListResult
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(siteIdParam string, enforcementpointIdParam string, virtualNetworkApplianceClusterIdParam string, cursorParam *string, inMaintenanceModeParam *bool, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, managementIpParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string, transportZonePathParam *string) (nsx_policyModel.VirtualNetworkApplianceListResult, error)

	// If the passed VirtualNetworkAppliance(VNA) does not already exist, create a new VirtualNetworkAppliance. If it already exists, patch it.
	//
	// @param siteIdParam (required)
	// @param enforcementpointIdParam (required)
	// @param virtualNetworkApplianceClusterIdParam (required)
	// @param virtualNetworkApplianceIdParam (required)
	// @param baseNetworkApplianceParam (required)
	// The parameter must contain all the properties defined in nsx_policyModel.BaseNetworkAppliance.
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Patch(siteIdParam string, enforcementpointIdParam string, virtualNetworkApplianceClusterIdParam string, virtualNetworkApplianceIdParam string, baseNetworkApplianceParam *vapiData_.StructValue) error

	// If the passed virtual network appliance(VNA) does not already exist, create a new virtual network appliance(VNA). If it already exists, patch it.
	//
	// @param siteIdParam (required)
	// @param enforcementpointIdParam (required)
	// @param virtualNetworkApplianceClusterIdParam (required)
	// @param virtualNetworkApplianceIdParam (required)
	// @param baseNetworkApplianceParam (required)
	// The parameter must contain all the properties defined in nsx_policyModel.BaseNetworkAppliance.
	// @return com.vmware.nsx_policy.model.BaseNetworkAppliance
	// The return value will contain all the properties defined in nsx_policyModel.BaseNetworkAppliance.
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(siteIdParam string, enforcementpointIdParam string, virtualNetworkApplianceClusterIdParam string, virtualNetworkApplianceIdParam string, baseNetworkApplianceParam *vapiData_.StructValue) (*vapiData_.StructValue, error)
}

type virtualNetworkAppliancesClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewVirtualNetworkAppliancesClient(connector vapiProtocolClient_.Connector) *virtualNetworkAppliancesClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx_policy.infra.sites.enforcement_points.virtual_network_appliance_clusters.virtual_network_appliances")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"delete": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "delete"),
		"get":    vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
		"list":   vapiCore_.NewMethodIdentifier(interfaceIdentifier, "list"),
		"patch":  vapiCore_.NewMethodIdentifier(interfaceIdentifier, "patch"),
		"update": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "update"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	vIface := virtualNetworkAppliancesClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &vIface
}

func (vIface *virtualNetworkAppliancesClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := vIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (vIface *virtualNetworkAppliancesClient) Delete(siteIdParam string, enforcementpointIdParam string, virtualNetworkApplianceClusterIdParam string, virtualNetworkApplianceIdParam string, forceParam *bool) error {
	typeConverter := vIface.connector.TypeConverter()
	executionContext := vIface.connector.NewExecutionContext()
	operationRestMetaData := virtualNetworkAppliancesDeleteRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(virtualNetworkAppliancesDeleteInputType(), typeConverter)
	sv.AddStructField("SiteId", siteIdParam)
	sv.AddStructField("EnforcementpointId", enforcementpointIdParam)
	sv.AddStructField("VirtualNetworkApplianceClusterId", virtualNetworkApplianceClusterIdParam)
	sv.AddStructField("VirtualNetworkApplianceId", virtualNetworkApplianceIdParam)
	sv.AddStructField("Force", forceParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := vIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.sites.enforcement_points.virtual_network_appliance_clusters.virtual_network_appliances", "delete", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), vIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (vIface *virtualNetworkAppliancesClient) Get(siteIdParam string, enforcementpointIdParam string, virtualNetworkApplianceClusterIdParam string, virtualNetworkApplianceIdParam string) (*vapiData_.StructValue, error) {
	typeConverter := vIface.connector.TypeConverter()
	executionContext := vIface.connector.NewExecutionContext()
	operationRestMetaData := virtualNetworkAppliancesGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(virtualNetworkAppliancesGetInputType(), typeConverter)
	sv.AddStructField("SiteId", siteIdParam)
	sv.AddStructField("EnforcementpointId", enforcementpointIdParam)
	sv.AddStructField("VirtualNetworkApplianceClusterId", virtualNetworkApplianceClusterIdParam)
	sv.AddStructField("VirtualNetworkApplianceId", virtualNetworkApplianceIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput *vapiData_.StructValue
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := vIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.sites.enforcement_points.virtual_network_appliance_clusters.virtual_network_appliances", "get", inputDataValue, executionContext)
	var emptyOutput *vapiData_.StructValue
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), VirtualNetworkAppliancesGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(*vapiData_.StructValue), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), vIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (vIface *virtualNetworkAppliancesClient) List(siteIdParam string, enforcementpointIdParam string, virtualNetworkApplianceClusterIdParam string, cursorParam *string, inMaintenanceModeParam *bool, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, managementIpParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string, transportZonePathParam *string) (nsx_policyModel.VirtualNetworkApplianceListResult, error) {
	typeConverter := vIface.connector.TypeConverter()
	executionContext := vIface.connector.NewExecutionContext()
	operationRestMetaData := virtualNetworkAppliancesListRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(virtualNetworkAppliancesListInputType(), typeConverter)
	sv.AddStructField("SiteId", siteIdParam)
	sv.AddStructField("EnforcementpointId", enforcementpointIdParam)
	sv.AddStructField("VirtualNetworkApplianceClusterId", virtualNetworkApplianceClusterIdParam)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("InMaintenanceMode", inMaintenanceModeParam)
	sv.AddStructField("IncludeMarkForDeleteObjects", includeMarkForDeleteObjectsParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("ManagementIp", managementIpParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	sv.AddStructField("TransportZonePath", transportZonePathParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.VirtualNetworkApplianceListResult
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := vIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.sites.enforcement_points.virtual_network_appliance_clusters.virtual_network_appliances", "list", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.VirtualNetworkApplianceListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), VirtualNetworkAppliancesListOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.VirtualNetworkApplianceListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), vIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (vIface *virtualNetworkAppliancesClient) Patch(siteIdParam string, enforcementpointIdParam string, virtualNetworkApplianceClusterIdParam string, virtualNetworkApplianceIdParam string, baseNetworkApplianceParam *vapiData_.StructValue) error {
	typeConverter := vIface.connector.TypeConverter()
	executionContext := vIface.connector.NewExecutionContext()
	operationRestMetaData := virtualNetworkAppliancesPatchRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(virtualNetworkAppliancesPatchInputType(), typeConverter)
	sv.AddStructField("SiteId", siteIdParam)
	sv.AddStructField("EnforcementpointId", enforcementpointIdParam)
	sv.AddStructField("VirtualNetworkApplianceClusterId", virtualNetworkApplianceClusterIdParam)
	sv.AddStructField("VirtualNetworkApplianceId", virtualNetworkApplianceIdParam)
	sv.AddStructField("BaseNetworkAppliance", baseNetworkApplianceParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := vIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.sites.enforcement_points.virtual_network_appliance_clusters.virtual_network_appliances", "patch", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), vIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (vIface *virtualNetworkAppliancesClient) Update(siteIdParam string, enforcementpointIdParam string, virtualNetworkApplianceClusterIdParam string, virtualNetworkApplianceIdParam string, baseNetworkApplianceParam *vapiData_.StructValue) (*vapiData_.StructValue, error) {
	typeConverter := vIface.connector.TypeConverter()
	executionContext := vIface.connector.NewExecutionContext()
	operationRestMetaData := virtualNetworkAppliancesUpdateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(virtualNetworkAppliancesUpdateInputType(), typeConverter)
	sv.AddStructField("SiteId", siteIdParam)
	sv.AddStructField("EnforcementpointId", enforcementpointIdParam)
	sv.AddStructField("VirtualNetworkApplianceClusterId", virtualNetworkApplianceClusterIdParam)
	sv.AddStructField("VirtualNetworkApplianceId", virtualNetworkApplianceIdParam)
	sv.AddStructField("BaseNetworkAppliance", baseNetworkApplianceParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput *vapiData_.StructValue
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := vIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.sites.enforcement_points.virtual_network_appliance_clusters.virtual_network_appliances", "update", inputDataValue, executionContext)
	var emptyOutput *vapiData_.StructValue
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), VirtualNetworkAppliancesUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(*vapiData_.StructValue), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), vIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
