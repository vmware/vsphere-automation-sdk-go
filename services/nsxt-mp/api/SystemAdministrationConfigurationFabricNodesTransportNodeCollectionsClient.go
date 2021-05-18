// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: SystemAdministrationConfigurationFabricNodesTransportNodeCollections
// Used by client-side stubs.

package api

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/core"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/lib"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/protocol/client"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt-mp/model"
)

const _ = core.SupportedByRuntimeVersion1

type SystemAdministrationConfigurationFabricNodesTransportNodeCollectionsClient interface {

	// When transport node collection is created the hosts which are part of compute collection will be prepared automatically i.e. NSX Manager attempts to install the NSX components on hosts. Transport nodes for these hosts are created using the configuration specified in transport node profile.
	//
	// @param transportNodeCollectionParam (required)
	// @param applyProfileParam Indicates if the Transport Node Profile (TNP) configuration should be applied during creation (optional, default to true)
	// @return com.vmware.model.TransportNodeCollection
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	CreateTransportNodeCollection(transportNodeCollectionParam model.TransportNodeCollection, applyProfileParam *bool) (model.TransportNodeCollection, error)

	// By deleting transport node collection, we are detaching the transport node profile(TNP) from the compute collection. It has no effect on existing transport nodes. However, new hosts added to the compute collection will no longer be automatically converted to NSX transport node. Detaching TNP from compute collection does not delete TNP.
	//
	// @param transportNodeCollectionIdParam (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	DeleteTransportNodeCollection(transportNodeCollectionIdParam string) error

	// Returns transport node collection by id
	//
	// @param transportNodeCollectionIdParam (required)
	// @return com.vmware.model.TransportNodeCollection
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetTransportNodeCollection(transportNodeCollectionIdParam string) (model.TransportNodeCollection, error)

	// Returns the state of transport node collection based on the states of transport nodes of the hosts which are part of compute collection.
	//
	// @param transportNodeCollectionIdParam (required)
	// @return com.vmware.model.TransportNodeCollectionState
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetTransportNodeCollectionState(transportNodeCollectionIdParam string) (model.TransportNodeCollectionState, error)

	// Returns all Transport Node collections
	//
	// @param clusterMoidParam Managed object ID of cluster in VC (optional)
	// @param computeCollectionIdParam Compute collection id (optional)
	// @param vcInstanceUuidParam UUID for VC deployment (optional)
	// @return com.vmware.model.TransportNodeCollectionListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ListTransportNodeCollections(clusterMoidParam *string, computeCollectionIdParam *string, vcInstanceUuidParam *string) (model.TransportNodeCollectionListResult, error)

	// This API is relevant for compute collection on which vLCM is enabled. This API shpuld be invoked to retry the realization of transport node profile on the compute collection. This is useful when profile realization had failed because of error in vLCM. This API has no effect if vLCM is not enabled on the computer collection.
	//
	// @param transportNodeCollectionIdParam (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	RetryTransportNodeCollectionRealizationRetryProfileRealization(transportNodeCollectionIdParam string) error

	// Attach different transport node profile to compute collection by updating transport node collection.
	//
	// @param transportNodeCollectionIdParam (required)
	// @param transportNodeCollectionParam (required)
	// @return com.vmware.model.TransportNodeCollection
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	UpdateTransportNodeCollection(transportNodeCollectionIdParam string, transportNodeCollectionParam model.TransportNodeCollection) (model.TransportNodeCollection, error)
}

type systemAdministrationConfigurationFabricNodesTransportNodeCollectionsClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewSystemAdministrationConfigurationFabricNodesTransportNodeCollectionsClient(connector client.Connector) *systemAdministrationConfigurationFabricNodesTransportNodeCollectionsClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.api.system_administration_configuration_fabric_nodes_transport_node_collections")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"create_transport_node_collection":                                      core.NewMethodIdentifier(interfaceIdentifier, "create_transport_node_collection"),
		"delete_transport_node_collection":                                      core.NewMethodIdentifier(interfaceIdentifier, "delete_transport_node_collection"),
		"get_transport_node_collection":                                         core.NewMethodIdentifier(interfaceIdentifier, "get_transport_node_collection"),
		"get_transport_node_collection_state":                                   core.NewMethodIdentifier(interfaceIdentifier, "get_transport_node_collection_state"),
		"list_transport_node_collections":                                       core.NewMethodIdentifier(interfaceIdentifier, "list_transport_node_collections"),
		"retry_transport_node_collection_realization_retry_profile_realization": core.NewMethodIdentifier(interfaceIdentifier, "retry_transport_node_collection_realization_retry_profile_realization"),
		"update_transport_node_collection":                                      core.NewMethodIdentifier(interfaceIdentifier, "update_transport_node_collection"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	sIface := systemAdministrationConfigurationFabricNodesTransportNodeCollectionsClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &sIface
}

func (sIface *systemAdministrationConfigurationFabricNodesTransportNodeCollectionsClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := sIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (sIface *systemAdministrationConfigurationFabricNodesTransportNodeCollectionsClient) CreateTransportNodeCollection(transportNodeCollectionParam model.TransportNodeCollection, applyProfileParam *bool) (model.TransportNodeCollection, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationFabricNodesTransportNodeCollectionsCreateTransportNodeCollectionInputType(), typeConverter)
	sv.AddStructField("TransportNodeCollection", transportNodeCollectionParam)
	sv.AddStructField("ApplyProfile", applyProfileParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.TransportNodeCollection
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationFabricNodesTransportNodeCollectionsCreateTransportNodeCollectionRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_fabric_nodes_transport_node_collections", "create_transport_node_collection", inputDataValue, executionContext)
	var emptyOutput model.TransportNodeCollection
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationConfigurationFabricNodesTransportNodeCollectionsCreateTransportNodeCollectionOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.TransportNodeCollection), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationConfigurationFabricNodesTransportNodeCollectionsClient) DeleteTransportNodeCollection(transportNodeCollectionIdParam string) error {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationFabricNodesTransportNodeCollectionsDeleteTransportNodeCollectionInputType(), typeConverter)
	sv.AddStructField("TransportNodeCollectionId", transportNodeCollectionIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationFabricNodesTransportNodeCollectionsDeleteTransportNodeCollectionRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_fabric_nodes_transport_node_collections", "delete_transport_node_collection", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (sIface *systemAdministrationConfigurationFabricNodesTransportNodeCollectionsClient) GetTransportNodeCollection(transportNodeCollectionIdParam string) (model.TransportNodeCollection, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationFabricNodesTransportNodeCollectionsGetTransportNodeCollectionInputType(), typeConverter)
	sv.AddStructField("TransportNodeCollectionId", transportNodeCollectionIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.TransportNodeCollection
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationFabricNodesTransportNodeCollectionsGetTransportNodeCollectionRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_fabric_nodes_transport_node_collections", "get_transport_node_collection", inputDataValue, executionContext)
	var emptyOutput model.TransportNodeCollection
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationConfigurationFabricNodesTransportNodeCollectionsGetTransportNodeCollectionOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.TransportNodeCollection), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationConfigurationFabricNodesTransportNodeCollectionsClient) GetTransportNodeCollectionState(transportNodeCollectionIdParam string) (model.TransportNodeCollectionState, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationFabricNodesTransportNodeCollectionsGetTransportNodeCollectionStateInputType(), typeConverter)
	sv.AddStructField("TransportNodeCollectionId", transportNodeCollectionIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.TransportNodeCollectionState
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationFabricNodesTransportNodeCollectionsGetTransportNodeCollectionStateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_fabric_nodes_transport_node_collections", "get_transport_node_collection_state", inputDataValue, executionContext)
	var emptyOutput model.TransportNodeCollectionState
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationConfigurationFabricNodesTransportNodeCollectionsGetTransportNodeCollectionStateOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.TransportNodeCollectionState), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationConfigurationFabricNodesTransportNodeCollectionsClient) ListTransportNodeCollections(clusterMoidParam *string, computeCollectionIdParam *string, vcInstanceUuidParam *string) (model.TransportNodeCollectionListResult, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationFabricNodesTransportNodeCollectionsListTransportNodeCollectionsInputType(), typeConverter)
	sv.AddStructField("ClusterMoid", clusterMoidParam)
	sv.AddStructField("ComputeCollectionId", computeCollectionIdParam)
	sv.AddStructField("VcInstanceUuid", vcInstanceUuidParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.TransportNodeCollectionListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationFabricNodesTransportNodeCollectionsListTransportNodeCollectionsRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_fabric_nodes_transport_node_collections", "list_transport_node_collections", inputDataValue, executionContext)
	var emptyOutput model.TransportNodeCollectionListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationConfigurationFabricNodesTransportNodeCollectionsListTransportNodeCollectionsOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.TransportNodeCollectionListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationConfigurationFabricNodesTransportNodeCollectionsClient) RetryTransportNodeCollectionRealizationRetryProfileRealization(transportNodeCollectionIdParam string) error {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationFabricNodesTransportNodeCollectionsRetryTransportNodeCollectionRealizationRetryProfileRealizationInputType(), typeConverter)
	sv.AddStructField("TransportNodeCollectionId", transportNodeCollectionIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationFabricNodesTransportNodeCollectionsRetryTransportNodeCollectionRealizationRetryProfileRealizationRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_fabric_nodes_transport_node_collections", "retry_transport_node_collection_realization_retry_profile_realization", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (sIface *systemAdministrationConfigurationFabricNodesTransportNodeCollectionsClient) UpdateTransportNodeCollection(transportNodeCollectionIdParam string, transportNodeCollectionParam model.TransportNodeCollection) (model.TransportNodeCollection, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationFabricNodesTransportNodeCollectionsUpdateTransportNodeCollectionInputType(), typeConverter)
	sv.AddStructField("TransportNodeCollectionId", transportNodeCollectionIdParam)
	sv.AddStructField("TransportNodeCollection", transportNodeCollectionParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.TransportNodeCollection
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationFabricNodesTransportNodeCollectionsUpdateTransportNodeCollectionRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_fabric_nodes_transport_node_collections", "update_transport_node_collection", inputDataValue, executionContext)
	var emptyOutput model.TransportNodeCollection
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationConfigurationFabricNodesTransportNodeCollectionsUpdateTransportNodeCollectionOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.TransportNodeCollection), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
