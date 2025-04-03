// Copyright (c) 2019-2025 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: ServiceNodes
// Used by client-side stubs.

package service_clusters

import (
	nsx_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type ServiceNodesClient interface {

    // Delete Service Node.
    //
    // @param siteIdParam (required)
    // @param enforcementpointIdParam (required)
    // @param serviceClusterIdParam (required)
    // @param serviceNodeIdParam (required)
    // @param forceParam Force delete the resource even if it is being used somewhere (optional, default to false)
    // @param relocateAndDeleteParam Relocate auto allocated service contexts from given service node (optional, default to false)
    //
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Delete(siteIdParam string, enforcementpointIdParam string, serviceClusterIdParam string, serviceNodeIdParam string, forceParam *bool, relocateAndDeleteParam *bool) error

    // Read a Service Node under an Enforcement Point
    //
    // @param siteIdParam (required)
    // @param enforcementpointIdParam (required)
    // @param serviceClusterIdParam (required)
    // @param serviceNodeIdParam (required)
    // @return com.vmware.nsx_policy.model.PolicyServiceNode
    //
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get(siteIdParam string, enforcementpointIdParam string, serviceClusterIdParam string, serviceNodeIdParam string) (nsx_policyModel.PolicyServiceNode, error)

    // List Service Nodes under an Enforcement Point
    //
    // @param siteIdParam (required)
    // @param enforcementpointIdParam (required)
    // @param serviceClusterIdParam (required)
    // @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
    // @param inMaintenanceModeParam Maintenance mode flag (optional)
    // @param includeMarkForDeleteObjectsParam Include objects that are marked for deletion in results (optional, default to false)
    // @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
    // @param managementIpParam Service node management IP address (optional)
    // @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
    // @param sortAscendingParam (optional)
    // @param sortByParam Field by which records are sorted (optional)
    // @param transportZonePathParam Transport zone path (optional)
    // @return com.vmware.nsx_policy.model.PolicyServiceNodeListResult
    //
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	List(siteIdParam string, enforcementpointIdParam string, serviceClusterIdParam string, cursorParam *string, inMaintenanceModeParam *bool, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, managementIpParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string, transportZonePathParam *string) (nsx_policyModel.PolicyServiceNodeListResult, error)

    // If the passed Service Node does not already exist, create a new Service Node. If it already exists, patch it.
    //
    // @param siteIdParam (required)
    // @param enforcementpointIdParam (required)
    // @param serviceClusterIdParam (required)
    // @param serviceNodeIdParam (required)
    // @param policyServiceNodeParam (required)
    //
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Patch(siteIdParam string, enforcementpointIdParam string, serviceClusterIdParam string, serviceNodeIdParam string, policyServiceNodeParam nsx_policyModel.PolicyServiceNode) error

    // If the passed Service Node does not already exist, create a new Service Node. If it already exists, patch it.
    //
    // @param siteIdParam (required)
    // @param enforcementpointIdParam (required)
    // @param serviceClusterIdParam (required)
    // @param serviceNodeIdParam (required)
    // @param policyServiceNodeParam (required)
    // @return com.vmware.nsx_policy.model.PolicyServiceNode
    //
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Update(siteIdParam string, enforcementpointIdParam string, serviceClusterIdParam string, serviceNodeIdParam string, policyServiceNodeParam nsx_policyModel.PolicyServiceNode) (nsx_policyModel.PolicyServiceNode, error)
}


type serviceNodesClient struct {
	connector           	   vapiProtocolClient_.Connector
	interfaceDefinition 	   vapiCore_.InterfaceDefinition
	errorsBindingMap           map[string]vapiBindings_.BindingType
}

func NewServiceNodesClient(connector vapiProtocolClient_.Connector) *serviceNodesClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx_policy.infra.sites.enforcement_points.service_clusters.service_nodes")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"delete": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "delete"),
		"get": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
		"list": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "list"),
		"patch": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "patch"),
		"update": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "update"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	sIface := serviceNodesClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &sIface
}

func (sIface *serviceNodesClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := sIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (sIface *serviceNodesClient) Delete(siteIdParam string, enforcementpointIdParam string, serviceClusterIdParam string, serviceNodeIdParam string, forceParam *bool, relocateAndDeleteParam *bool) error {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	operationRestMetaData := serviceNodesDeleteRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(serviceNodesDeleteInputType(), typeConverter)
	sv.AddStructField("SiteId", siteIdParam)
	sv.AddStructField("EnforcementpointId", enforcementpointIdParam)
	sv.AddStructField("ServiceClusterId", serviceClusterIdParam)
	sv.AddStructField("ServiceNodeId", serviceNodeIdParam)
	sv.AddStructField("Force", forceParam)
	sv.AddStructField("RelocateAndDelete", relocateAndDeleteParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.sites.enforcement_points.service_clusters.service_nodes", "delete", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (sIface *serviceNodesClient) Get(siteIdParam string, enforcementpointIdParam string, serviceClusterIdParam string, serviceNodeIdParam string) (nsx_policyModel.PolicyServiceNode, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	operationRestMetaData := serviceNodesGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(serviceNodesGetInputType(), typeConverter)
	sv.AddStructField("SiteId", siteIdParam)
	sv.AddStructField("EnforcementpointId", enforcementpointIdParam)
	sv.AddStructField("ServiceClusterId", serviceClusterIdParam)
	sv.AddStructField("ServiceNodeId", serviceNodeIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.PolicyServiceNode
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.sites.enforcement_points.service_clusters.service_nodes", "get", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.PolicyServiceNode
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), ServiceNodesGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.PolicyServiceNode), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *serviceNodesClient) List(siteIdParam string, enforcementpointIdParam string, serviceClusterIdParam string, cursorParam *string, inMaintenanceModeParam *bool, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, managementIpParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string, transportZonePathParam *string) (nsx_policyModel.PolicyServiceNodeListResult, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	operationRestMetaData := serviceNodesListRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(serviceNodesListInputType(), typeConverter)
	sv.AddStructField("SiteId", siteIdParam)
	sv.AddStructField("EnforcementpointId", enforcementpointIdParam)
	sv.AddStructField("ServiceClusterId", serviceClusterIdParam)
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
		var emptyOutput nsx_policyModel.PolicyServiceNodeListResult
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.sites.enforcement_points.service_clusters.service_nodes", "list", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.PolicyServiceNodeListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), ServiceNodesListOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.PolicyServiceNodeListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *serviceNodesClient) Patch(siteIdParam string, enforcementpointIdParam string, serviceClusterIdParam string, serviceNodeIdParam string, policyServiceNodeParam nsx_policyModel.PolicyServiceNode) error {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	operationRestMetaData := serviceNodesPatchRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(serviceNodesPatchInputType(), typeConverter)
	sv.AddStructField("SiteId", siteIdParam)
	sv.AddStructField("EnforcementpointId", enforcementpointIdParam)
	sv.AddStructField("ServiceClusterId", serviceClusterIdParam)
	sv.AddStructField("ServiceNodeId", serviceNodeIdParam)
	sv.AddStructField("PolicyServiceNode", policyServiceNodeParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.sites.enforcement_points.service_clusters.service_nodes", "patch", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (sIface *serviceNodesClient) Update(siteIdParam string, enforcementpointIdParam string, serviceClusterIdParam string, serviceNodeIdParam string, policyServiceNodeParam nsx_policyModel.PolicyServiceNode) (nsx_policyModel.PolicyServiceNode, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	operationRestMetaData := serviceNodesUpdateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(serviceNodesUpdateInputType(), typeConverter)
	sv.AddStructField("SiteId", siteIdParam)
	sv.AddStructField("EnforcementpointId", enforcementpointIdParam)
	sv.AddStructField("ServiceClusterId", serviceClusterIdParam)
	sv.AddStructField("ServiceNodeId", serviceNodeIdParam)
	sv.AddStructField("PolicyServiceNode", policyServiceNodeParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.PolicyServiceNode
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.sites.enforcement_points.service_clusters.service_nodes", "update", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.PolicyServiceNode
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), ServiceNodesUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.PolicyServiceNode), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

