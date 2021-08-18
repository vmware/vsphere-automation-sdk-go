// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Clusters
// Used by client-side stubs.

package intelligence

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/core"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/lib"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/protocol/client"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt/model"
)

const _ = core.SupportedByRuntimeVersion1

type ClustersClient interface {

	// Enable / disable NSX Intelligence data collection on the selected transport node cluster.
	//
	// @param siteIdParam (required)
	// @param clusterIdParam (required)
	// @return com.vmware.nsx_policy.model.IntelligenceTransportNodeClusterInfo
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(siteIdParam string, clusterIdParam string) (model.IntelligenceTransportNodeClusterInfo, error)

	// Get a list of transport node clusters that have NSX Intelligence data collection enabled and/or disabled.
	//
	// @param siteIdParam (required)
	// @param clusterTypeParam NSX Intelligence transport node cluster type (required)
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param enabledParam Data collection enabled (optional)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @return com.vmware.nsx_policy.model.IntelligenceTransportNodeClusterInfoList
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(siteIdParam string, clusterTypeParam string, cursorParam *string, enabledParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.IntelligenceTransportNodeClusterInfoList, error)

	// Enable / disable NSX Intelligence data collection for a list of transport node clusters.
	//
	// @param siteIdParam (required)
	// @param intelligenceTransportNodeClusterInfoListParam (required)
	// @return com.vmware.nsx_policy.model.IntelligenceTransportNodeClusterInfoList
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Patch(siteIdParam string, intelligenceTransportNodeClusterInfoListParam model.IntelligenceTransportNodeClusterInfoList) (model.IntelligenceTransportNodeClusterInfoList, error)

	// Enable / disable NSX Intelligence data collection on the selected transport node cluster.
	//
	// @param siteIdParam (required)
	// @param clusterIdParam (required)
	// @param intelligenceTransportNodeClusterInfoParam (required)
	// @return com.vmware.nsx_policy.model.IntelligenceTransportNodeClusterInfo
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Patch0(siteIdParam string, clusterIdParam string, intelligenceTransportNodeClusterInfoParam model.IntelligenceTransportNodeClusterInfo) (model.IntelligenceTransportNodeClusterInfo, error)

	// Enable / disable NSX Intelligence data collection on the selected transport node cluster.
	//
	// @param siteIdParam (required)
	// @param clusterIdParam (required)
	// @param intelligenceTransportNodeClusterInfoParam (required)
	// @return com.vmware.nsx_policy.model.IntelligenceTransportNodeClusterInfo
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(siteIdParam string, clusterIdParam string, intelligenceTransportNodeClusterInfoParam model.IntelligenceTransportNodeClusterInfo) (model.IntelligenceTransportNodeClusterInfo, error)
}

type clustersClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewClustersClient(connector client.Connector) *clustersClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.nsx_policy.infra.sites.intelligence.clusters")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"get":     core.NewMethodIdentifier(interfaceIdentifier, "get"),
		"list":    core.NewMethodIdentifier(interfaceIdentifier, "list"),
		"patch":   core.NewMethodIdentifier(interfaceIdentifier, "patch"),
		"patch_0": core.NewMethodIdentifier(interfaceIdentifier, "patch_0"),
		"update":  core.NewMethodIdentifier(interfaceIdentifier, "update"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	cIface := clustersClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &cIface
}

func (cIface *clustersClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := cIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (cIface *clustersClient) Get(siteIdParam string, clusterIdParam string) (model.IntelligenceTransportNodeClusterInfo, error) {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(clustersGetInputType(), typeConverter)
	sv.AddStructField("SiteId", siteIdParam)
	sv.AddStructField("ClusterId", clusterIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.IntelligenceTransportNodeClusterInfo
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := clustersGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	cIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.sites.intelligence.clusters", "get", inputDataValue, executionContext)
	var emptyOutput model.IntelligenceTransportNodeClusterInfo
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), clustersGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.IntelligenceTransportNodeClusterInfo), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (cIface *clustersClient) List(siteIdParam string, clusterTypeParam string, cursorParam *string, enabledParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.IntelligenceTransportNodeClusterInfoList, error) {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(clustersListInputType(), typeConverter)
	sv.AddStructField("SiteId", siteIdParam)
	sv.AddStructField("ClusterType", clusterTypeParam)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("Enabled", enabledParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.IntelligenceTransportNodeClusterInfoList
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := clustersListRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	cIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.sites.intelligence.clusters", "list", inputDataValue, executionContext)
	var emptyOutput model.IntelligenceTransportNodeClusterInfoList
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), clustersListOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.IntelligenceTransportNodeClusterInfoList), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (cIface *clustersClient) Patch(siteIdParam string, intelligenceTransportNodeClusterInfoListParam model.IntelligenceTransportNodeClusterInfoList) (model.IntelligenceTransportNodeClusterInfoList, error) {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(clustersPatchInputType(), typeConverter)
	sv.AddStructField("SiteId", siteIdParam)
	sv.AddStructField("IntelligenceTransportNodeClusterInfoList", intelligenceTransportNodeClusterInfoListParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.IntelligenceTransportNodeClusterInfoList
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := clustersPatchRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	cIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.sites.intelligence.clusters", "patch", inputDataValue, executionContext)
	var emptyOutput model.IntelligenceTransportNodeClusterInfoList
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), clustersPatchOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.IntelligenceTransportNodeClusterInfoList), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (cIface *clustersClient) Patch0(siteIdParam string, clusterIdParam string, intelligenceTransportNodeClusterInfoParam model.IntelligenceTransportNodeClusterInfo) (model.IntelligenceTransportNodeClusterInfo, error) {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(clustersPatch0InputType(), typeConverter)
	sv.AddStructField("SiteId", siteIdParam)
	sv.AddStructField("ClusterId", clusterIdParam)
	sv.AddStructField("IntelligenceTransportNodeClusterInfo", intelligenceTransportNodeClusterInfoParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.IntelligenceTransportNodeClusterInfo
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := clustersPatch0RestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	cIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.sites.intelligence.clusters", "patch_0", inputDataValue, executionContext)
	var emptyOutput model.IntelligenceTransportNodeClusterInfo
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), clustersPatch0OutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.IntelligenceTransportNodeClusterInfo), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (cIface *clustersClient) Update(siteIdParam string, clusterIdParam string, intelligenceTransportNodeClusterInfoParam model.IntelligenceTransportNodeClusterInfo) (model.IntelligenceTransportNodeClusterInfo, error) {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(clustersUpdateInputType(), typeConverter)
	sv.AddStructField("SiteId", siteIdParam)
	sv.AddStructField("ClusterId", clusterIdParam)
	sv.AddStructField("IntelligenceTransportNodeClusterInfo", intelligenceTransportNodeClusterInfoParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.IntelligenceTransportNodeClusterInfo
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := clustersUpdateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	cIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.sites.intelligence.clusters", "update", inputDataValue, executionContext)
	var emptyOutput model.IntelligenceTransportNodeClusterInfo
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), clustersUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.IntelligenceTransportNodeClusterInfo), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
