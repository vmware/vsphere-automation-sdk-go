// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: SystemAdministrationConfigurationFabricComputeCollections
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

type SystemAdministrationConfigurationFabricComputeCollectionsClient interface {

	// Fabric templates are fabric configurations applied at the compute collection level. This configurations is used to decide what automated operations should be a run when a host membership changes. This functionality is deprecated. Use Transport Node Profiles instead of this template.
	//
	// @param computeCollectionFabricTemplateParam (required)
	// @return com.vmware.model.ComputeCollectionFabricTemplate
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	CreateComputeCollectionFabricTemplate(computeCollectionFabricTemplateParam model.ComputeCollectionFabricTemplate) (model.ComputeCollectionFabricTemplate, error)

	// Deletes compute collection fabric template for the given id. This functionality is deprecated. Use Transport Node Profiles instead of this template.
	//
	// @param fabricTemplateIdParam (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	DeleteComputeCollectionFabricTemplate(fabricTemplateIdParam string) error

	// Get compute collection fabric template for the given id. This functionality is deprecated. Use Transport Node Profiles instead of this template.
	//
	// @param fabricTemplateIdParam (required)
	// @return com.vmware.model.ComputeCollectionFabricTemplate
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetComputeCollectionFabricTemplate(fabricTemplateIdParam string) (model.ComputeCollectionFabricTemplate, error)

	// Get status of member host nodes of the compute-collection. Only nsx prepared host nodes in the specified compute-collection are included in the response. cc-ext-id should be of type VC_Cluster.
	//
	// @param ccExtIdParam (required)
	// @return com.vmware.model.HostNodeStatusListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetHostNodeStatusOnComputeCollection(ccExtIdParam string) (model.HostNodeStatusListResult, error)

	// Returns compute collection fabric templates. This functionality is deprecated. Use Transport Node Profiles instead of this template.
	//
	// @param computeCollectionIdParam Compute collection id (optional)
	// @return com.vmware.model.ComputeCollectionFabricTemplateListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ListComputeCollectionFabricTemplates(computeCollectionIdParam *string) (model.ComputeCollectionFabricTemplateListResult, error)

	// Returns list of physical network interfaces for all discovered nodes in compute collection. Interface information includes PNIC name, hostswitch name it's attached to(if any) and MAC address.
	//
	// @param ccExtIdParam (required)
	// @return com.vmware.model.ComputeCollectionNetworkInterfacesListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ListComputeCollectionPhysicalNetworkInterfaces(ccExtIdParam string) (model.ComputeCollectionNetworkInterfacesListResult, error)

	// Returns information about all compute collections.
	//
	// @param cmLocalIdParam Local Id of the compute collection in the Compute Manager (optional)
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param discoveredNodeIdParam Id of the discovered node which belongs to this Compute Collection (optional)
	// @param displayNameParam Name of the ComputeCollection in source compute manager (optional)
	// @param externalIdParam External ID of the ComputeCollection in the source Compute manager, e.g. mo-ref in VC (optional)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param nodeIdParam Id of the fabric node created from a discovered node belonging to this Compute Collection (optional)
	// @param originIdParam Id of the compute manager from where this Compute Collection was discovered (optional)
	// @param originTypeParam ComputeCollection type like VC_Cluster. Here the Compute Manager type prefix would help in differentiating similar named Compute Collection types from different Compute Managers (optional)
	// @param ownerIdParam Id of the owner of compute collection in the Compute Manager (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @return com.vmware.model.ComputeCollectionListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ListComputeCollections(cmLocalIdParam *string, cursorParam *string, discoveredNodeIdParam *string, displayNameParam *string, externalIdParam *string, includedFieldsParam *string, nodeIdParam *string, originIdParam *string, originTypeParam *string, ownerIdParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.ComputeCollectionListResult, error)

	// Perform action specific to NSX on the compute-collection. cc-ext-id should be of type VC_Cluster.
	//
	// @param ccExtIdParam (required)
	// @param actionParam Supported actions on compute-collection (optional)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	PerformActionOnComputeCollection(ccExtIdParam string, actionParam *string) error

	// Returns information about a specific compute collection.
	//
	// @param ccExtIdParam (required)
	// @return com.vmware.model.ComputeCollection
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ReadComputeCollection(ccExtIdParam string) (model.ComputeCollection, error)

	// Updates compute collection fabric template for the given id. This functionality is deprecated. Use Transport Node Profiles instead of this template.
	//
	// @param fabricTemplateIdParam (required)
	// @param computeCollectionFabricTemplateParam (required)
	// @return com.vmware.model.ComputeCollectionFabricTemplate
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	UpdateComputeCollectionFabricTemplate(fabricTemplateIdParam string, computeCollectionFabricTemplateParam model.ComputeCollectionFabricTemplate) (model.ComputeCollectionFabricTemplate, error)
}

type systemAdministrationConfigurationFabricComputeCollectionsClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewSystemAdministrationConfigurationFabricComputeCollectionsClient(connector client.Connector) *systemAdministrationConfigurationFabricComputeCollectionsClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.api.system_administration_configuration_fabric_compute_collections")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"create_compute_collection_fabric_template":           core.NewMethodIdentifier(interfaceIdentifier, "create_compute_collection_fabric_template"),
		"delete_compute_collection_fabric_template":           core.NewMethodIdentifier(interfaceIdentifier, "delete_compute_collection_fabric_template"),
		"get_compute_collection_fabric_template":              core.NewMethodIdentifier(interfaceIdentifier, "get_compute_collection_fabric_template"),
		"get_host_node_status_on_compute_collection":          core.NewMethodIdentifier(interfaceIdentifier, "get_host_node_status_on_compute_collection"),
		"list_compute_collection_fabric_templates":            core.NewMethodIdentifier(interfaceIdentifier, "list_compute_collection_fabric_templates"),
		"list_compute_collection_physical_network_interfaces": core.NewMethodIdentifier(interfaceIdentifier, "list_compute_collection_physical_network_interfaces"),
		"list_compute_collections":                            core.NewMethodIdentifier(interfaceIdentifier, "list_compute_collections"),
		"perform_action_on_compute_collection":                core.NewMethodIdentifier(interfaceIdentifier, "perform_action_on_compute_collection"),
		"read_compute_collection":                             core.NewMethodIdentifier(interfaceIdentifier, "read_compute_collection"),
		"update_compute_collection_fabric_template":           core.NewMethodIdentifier(interfaceIdentifier, "update_compute_collection_fabric_template"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	sIface := systemAdministrationConfigurationFabricComputeCollectionsClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &sIface
}

func (sIface *systemAdministrationConfigurationFabricComputeCollectionsClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := sIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (sIface *systemAdministrationConfigurationFabricComputeCollectionsClient) CreateComputeCollectionFabricTemplate(computeCollectionFabricTemplateParam model.ComputeCollectionFabricTemplate) (model.ComputeCollectionFabricTemplate, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationFabricComputeCollectionsCreateComputeCollectionFabricTemplateInputType(), typeConverter)
	sv.AddStructField("ComputeCollectionFabricTemplate", computeCollectionFabricTemplateParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ComputeCollectionFabricTemplate
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationFabricComputeCollectionsCreateComputeCollectionFabricTemplateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_fabric_compute_collections", "create_compute_collection_fabric_template", inputDataValue, executionContext)
	var emptyOutput model.ComputeCollectionFabricTemplate
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationConfigurationFabricComputeCollectionsCreateComputeCollectionFabricTemplateOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ComputeCollectionFabricTemplate), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationConfigurationFabricComputeCollectionsClient) DeleteComputeCollectionFabricTemplate(fabricTemplateIdParam string) error {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationFabricComputeCollectionsDeleteComputeCollectionFabricTemplateInputType(), typeConverter)
	sv.AddStructField("FabricTemplateId", fabricTemplateIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationFabricComputeCollectionsDeleteComputeCollectionFabricTemplateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_fabric_compute_collections", "delete_compute_collection_fabric_template", inputDataValue, executionContext)
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

func (sIface *systemAdministrationConfigurationFabricComputeCollectionsClient) GetComputeCollectionFabricTemplate(fabricTemplateIdParam string) (model.ComputeCollectionFabricTemplate, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationFabricComputeCollectionsGetComputeCollectionFabricTemplateInputType(), typeConverter)
	sv.AddStructField("FabricTemplateId", fabricTemplateIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ComputeCollectionFabricTemplate
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationFabricComputeCollectionsGetComputeCollectionFabricTemplateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_fabric_compute_collections", "get_compute_collection_fabric_template", inputDataValue, executionContext)
	var emptyOutput model.ComputeCollectionFabricTemplate
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationConfigurationFabricComputeCollectionsGetComputeCollectionFabricTemplateOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ComputeCollectionFabricTemplate), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationConfigurationFabricComputeCollectionsClient) GetHostNodeStatusOnComputeCollection(ccExtIdParam string) (model.HostNodeStatusListResult, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationFabricComputeCollectionsGetHostNodeStatusOnComputeCollectionInputType(), typeConverter)
	sv.AddStructField("CcExtId", ccExtIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.HostNodeStatusListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationFabricComputeCollectionsGetHostNodeStatusOnComputeCollectionRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_fabric_compute_collections", "get_host_node_status_on_compute_collection", inputDataValue, executionContext)
	var emptyOutput model.HostNodeStatusListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationConfigurationFabricComputeCollectionsGetHostNodeStatusOnComputeCollectionOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.HostNodeStatusListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationConfigurationFabricComputeCollectionsClient) ListComputeCollectionFabricTemplates(computeCollectionIdParam *string) (model.ComputeCollectionFabricTemplateListResult, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationFabricComputeCollectionsListComputeCollectionFabricTemplatesInputType(), typeConverter)
	sv.AddStructField("ComputeCollectionId", computeCollectionIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ComputeCollectionFabricTemplateListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationFabricComputeCollectionsListComputeCollectionFabricTemplatesRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_fabric_compute_collections", "list_compute_collection_fabric_templates", inputDataValue, executionContext)
	var emptyOutput model.ComputeCollectionFabricTemplateListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationConfigurationFabricComputeCollectionsListComputeCollectionFabricTemplatesOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ComputeCollectionFabricTemplateListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationConfigurationFabricComputeCollectionsClient) ListComputeCollectionPhysicalNetworkInterfaces(ccExtIdParam string) (model.ComputeCollectionNetworkInterfacesListResult, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationFabricComputeCollectionsListComputeCollectionPhysicalNetworkInterfacesInputType(), typeConverter)
	sv.AddStructField("CcExtId", ccExtIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ComputeCollectionNetworkInterfacesListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationFabricComputeCollectionsListComputeCollectionPhysicalNetworkInterfacesRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_fabric_compute_collections", "list_compute_collection_physical_network_interfaces", inputDataValue, executionContext)
	var emptyOutput model.ComputeCollectionNetworkInterfacesListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationConfigurationFabricComputeCollectionsListComputeCollectionPhysicalNetworkInterfacesOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ComputeCollectionNetworkInterfacesListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationConfigurationFabricComputeCollectionsClient) ListComputeCollections(cmLocalIdParam *string, cursorParam *string, discoveredNodeIdParam *string, displayNameParam *string, externalIdParam *string, includedFieldsParam *string, nodeIdParam *string, originIdParam *string, originTypeParam *string, ownerIdParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.ComputeCollectionListResult, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationFabricComputeCollectionsListComputeCollectionsInputType(), typeConverter)
	sv.AddStructField("CmLocalId", cmLocalIdParam)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("DiscoveredNodeId", discoveredNodeIdParam)
	sv.AddStructField("DisplayName", displayNameParam)
	sv.AddStructField("ExternalId", externalIdParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("NodeId", nodeIdParam)
	sv.AddStructField("OriginId", originIdParam)
	sv.AddStructField("OriginType", originTypeParam)
	sv.AddStructField("OwnerId", ownerIdParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ComputeCollectionListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationFabricComputeCollectionsListComputeCollectionsRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_fabric_compute_collections", "list_compute_collections", inputDataValue, executionContext)
	var emptyOutput model.ComputeCollectionListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationConfigurationFabricComputeCollectionsListComputeCollectionsOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ComputeCollectionListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationConfigurationFabricComputeCollectionsClient) PerformActionOnComputeCollection(ccExtIdParam string, actionParam *string) error {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationFabricComputeCollectionsPerformActionOnComputeCollectionInputType(), typeConverter)
	sv.AddStructField("CcExtId", ccExtIdParam)
	sv.AddStructField("Action", actionParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationFabricComputeCollectionsPerformActionOnComputeCollectionRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_fabric_compute_collections", "perform_action_on_compute_collection", inputDataValue, executionContext)
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

func (sIface *systemAdministrationConfigurationFabricComputeCollectionsClient) ReadComputeCollection(ccExtIdParam string) (model.ComputeCollection, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationFabricComputeCollectionsReadComputeCollectionInputType(), typeConverter)
	sv.AddStructField("CcExtId", ccExtIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ComputeCollection
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationFabricComputeCollectionsReadComputeCollectionRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_fabric_compute_collections", "read_compute_collection", inputDataValue, executionContext)
	var emptyOutput model.ComputeCollection
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationConfigurationFabricComputeCollectionsReadComputeCollectionOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ComputeCollection), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationConfigurationFabricComputeCollectionsClient) UpdateComputeCollectionFabricTemplate(fabricTemplateIdParam string, computeCollectionFabricTemplateParam model.ComputeCollectionFabricTemplate) (model.ComputeCollectionFabricTemplate, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationFabricComputeCollectionsUpdateComputeCollectionFabricTemplateInputType(), typeConverter)
	sv.AddStructField("FabricTemplateId", fabricTemplateIdParam)
	sv.AddStructField("ComputeCollectionFabricTemplate", computeCollectionFabricTemplateParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ComputeCollectionFabricTemplate
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationFabricComputeCollectionsUpdateComputeCollectionFabricTemplateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_fabric_compute_collections", "update_compute_collection_fabric_template", inputDataValue, executionContext)
	var emptyOutput model.ComputeCollectionFabricTemplate
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationConfigurationFabricComputeCollectionsUpdateComputeCollectionFabricTemplateOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ComputeCollectionFabricTemplate), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
