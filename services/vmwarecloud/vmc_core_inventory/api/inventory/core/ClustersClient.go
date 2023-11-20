// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Clusters
// Used by client-side stubs.

package core

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	vmwarecloudVmc_core_inventoryModel "github.com/vmware/vsphere-automation-sdk-go/services/vmwarecloud/vmc_core_inventory/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type ClustersClient interface {

	// Get a collection of cluster scoped by organization
	//
	// @param orgIdParam organization identifier
	// @param deploymentIdParam parent deployment identifier
	// @param includeDeletedResourcesParam should the deleted resources be included in the results?
	// @param pageParam Zero-based page index (0..N)
	// @param sizeParam The size of the page to be returned
	// @param sortParam Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported.
	// @return OK
	//
	// @throws Unauthenticated Unauthorized
	// @throws Unauthorized Forbidden
	GetCoreClusters(orgIdParam string, deploymentIdParam *string, includeDeletedResourcesParam *bool, pageParam *int64, sizeParam *int64, sortParam []string) (vmwarecloudVmc_core_inventoryModel.ClusterResponse, error)

	// Get details for a specific core cluster
	//
	// @param orgIdParam organization identifier
	// @param clusterIdParam cluster identifier
	// @param includeDeletedResourcesParam should the deleted resources be included in the results?
	// @return OK
	//
	// @throws Unauthenticated Unauthorized
	// @throws Unauthorized Forbidden
	// @throws NotFound Not Found
	GetCoreClusterById(orgIdParam string, clusterIdParam string, includeDeletedResourcesParam *bool) (vmwarecloudVmc_core_inventoryModel.Cluster, error)
}

type clustersClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewClustersClient(connector vapiProtocolClient_.Connector) *clustersClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.vmwarecloud.vmc_core_inventory.api.inventory.core.clusters")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"get_core_clusters":      vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get_core_clusters"),
		"get_core_cluster_by_id": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get_core_cluster_by_id"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	cIface := clustersClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &cIface
}

func (cIface *clustersClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := cIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (cIface *clustersClient) GetCoreClusters(orgIdParam string, deploymentIdParam *string, includeDeletedResourcesParam *bool, pageParam *int64, sizeParam *int64, sortParam []string) (vmwarecloudVmc_core_inventoryModel.ClusterResponse, error) {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	operationRestMetaData := clustersGetCoreClustersRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(clustersGetCoreClustersInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("DeploymentId", deploymentIdParam)
	sv.AddStructField("IncludeDeletedResources", includeDeletedResourcesParam)
	sv.AddStructField("Page", pageParam)
	sv.AddStructField("Size", sizeParam)
	sv.AddStructField("Sort", sortParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput vmwarecloudVmc_core_inventoryModel.ClusterResponse
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.vmwarecloud.vmc_core_inventory.api.inventory.core.clusters", "get_core_clusters", inputDataValue, executionContext)
	var emptyOutput vmwarecloudVmc_core_inventoryModel.ClusterResponse
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), ClustersGetCoreClustersOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(vmwarecloudVmc_core_inventoryModel.ClusterResponse), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (cIface *clustersClient) GetCoreClusterById(orgIdParam string, clusterIdParam string, includeDeletedResourcesParam *bool) (vmwarecloudVmc_core_inventoryModel.Cluster, error) {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	operationRestMetaData := clustersGetCoreClusterByIdRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(clustersGetCoreClusterByIdInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("ClusterId", clusterIdParam)
	sv.AddStructField("IncludeDeletedResources", includeDeletedResourcesParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput vmwarecloudVmc_core_inventoryModel.Cluster
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.vmwarecloud.vmc_core_inventory.api.inventory.core.clusters", "get_core_cluster_by_id", inputDataValue, executionContext)
	var emptyOutput vmwarecloudVmc_core_inventoryModel.Cluster
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), ClustersGetCoreClusterByIdOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(vmwarecloudVmc_core_inventoryModel.Cluster), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
