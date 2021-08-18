// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: TransportNodes
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

type TransportNodesClient interface {

	// Delete intelligence transport node config.
	//
	// @param siteIdParam (required)
	// @param tnIdParam (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Delete(siteIdParam string, tnIdParam string) error

	// Get a transport node configuration for intelligence.
	//
	// @param siteIdParam (required)
	// @param tnIdParam (required)
	// @return com.vmware.nsx_policy.model.IntelligenceTransportNodeInfo
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(siteIdParam string, tnIdParam string) (model.IntelligenceTransportNodeInfo, error)

	// Get a list of transport nodes that have NSX Intelligence data collection enabled and/or disabled.
	//
	// @param siteIdParam (required)
	// @param transportNodeTypeParam NSX Intelligence transport node type (required)
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param enabledParam Data collection enabled (optional)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @param standaloneParam Standalone transport node (optional)
	// @return com.vmware.nsx_policy.model.IntelligenceTransportNodeInfoList
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(siteIdParam string, transportNodeTypeParam string, cursorParam *string, enabledParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string, standaloneParam *bool) (model.IntelligenceTransportNodeInfoList, error)

	// Enable / disable NSX Intelligence data collection for a list of standalone transport nodes. If a transport node is belonging to a transport node cluster, this API can't be used to adjust data collection configuration. Use transport node cluster API to adjust data collection configuration.
	//
	// @param siteIdParam (required)
	// @param intelligenceTransportNodeInfoListParam (required)
	// @return com.vmware.nsx_policy.model.IntelligenceTransportNodeInfoList
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Patch(siteIdParam string, intelligenceTransportNodeInfoListParam model.IntelligenceTransportNodeInfoList) (model.IntelligenceTransportNodeInfoList, error)

	// Enable / disable NSX Intelligence data collection on the selected standalone transport node. If a transport node is belonging to a transport node cluster, this API can't be used to adjust data collection configuration. Use transport node cluster API to adjust data collection configuration.
	//
	// @param siteIdParam (required)
	// @param tnIdParam (required)
	// @param intelligenceTransportNodeInfoParam (required)
	// @return com.vmware.nsx_policy.model.IntelligenceTransportNodeInfo
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Patch0(siteIdParam string, tnIdParam string, intelligenceTransportNodeInfoParam model.IntelligenceTransportNodeInfo) (model.IntelligenceTransportNodeInfo, error)

	// Enable / disable NSX Intelligence data collection on the selected standalone transport node. If a transport node is belonging to a transport node cluster, this API can't be used to adjust data collection configuration. Use transport node cluster API to adjust data collection configuration.
	//
	// @param siteIdParam (required)
	// @param tnIdParam (required)
	// @param intelligenceTransportNodeInfoParam (required)
	// @return com.vmware.nsx_policy.model.IntelligenceTransportNodeInfo
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(siteIdParam string, tnIdParam string, intelligenceTransportNodeInfoParam model.IntelligenceTransportNodeInfo) (model.IntelligenceTransportNodeInfo, error)
}

type transportNodesClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewTransportNodesClient(connector client.Connector) *transportNodesClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.nsx_policy.infra.sites.intelligence.transport_nodes")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"delete":  core.NewMethodIdentifier(interfaceIdentifier, "delete"),
		"get":     core.NewMethodIdentifier(interfaceIdentifier, "get"),
		"list":    core.NewMethodIdentifier(interfaceIdentifier, "list"),
		"patch":   core.NewMethodIdentifier(interfaceIdentifier, "patch"),
		"patch_0": core.NewMethodIdentifier(interfaceIdentifier, "patch_0"),
		"update":  core.NewMethodIdentifier(interfaceIdentifier, "update"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	tIface := transportNodesClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &tIface
}

func (tIface *transportNodesClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := tIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (tIface *transportNodesClient) Delete(siteIdParam string, tnIdParam string) error {
	typeConverter := tIface.connector.TypeConverter()
	executionContext := tIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(transportNodesDeleteInputType(), typeConverter)
	sv.AddStructField("SiteId", siteIdParam)
	sv.AddStructField("TnId", tnIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := transportNodesDeleteRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	tIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := tIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.sites.intelligence.transport_nodes", "delete", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (tIface *transportNodesClient) Get(siteIdParam string, tnIdParam string) (model.IntelligenceTransportNodeInfo, error) {
	typeConverter := tIface.connector.TypeConverter()
	executionContext := tIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(transportNodesGetInputType(), typeConverter)
	sv.AddStructField("SiteId", siteIdParam)
	sv.AddStructField("TnId", tnIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.IntelligenceTransportNodeInfo
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := transportNodesGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	tIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := tIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.sites.intelligence.transport_nodes", "get", inputDataValue, executionContext)
	var emptyOutput model.IntelligenceTransportNodeInfo
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), transportNodesGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.IntelligenceTransportNodeInfo), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (tIface *transportNodesClient) List(siteIdParam string, transportNodeTypeParam string, cursorParam *string, enabledParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string, standaloneParam *bool) (model.IntelligenceTransportNodeInfoList, error) {
	typeConverter := tIface.connector.TypeConverter()
	executionContext := tIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(transportNodesListInputType(), typeConverter)
	sv.AddStructField("SiteId", siteIdParam)
	sv.AddStructField("TransportNodeType", transportNodeTypeParam)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("Enabled", enabledParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	sv.AddStructField("Standalone", standaloneParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.IntelligenceTransportNodeInfoList
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := transportNodesListRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	tIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := tIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.sites.intelligence.transport_nodes", "list", inputDataValue, executionContext)
	var emptyOutput model.IntelligenceTransportNodeInfoList
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), transportNodesListOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.IntelligenceTransportNodeInfoList), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (tIface *transportNodesClient) Patch(siteIdParam string, intelligenceTransportNodeInfoListParam model.IntelligenceTransportNodeInfoList) (model.IntelligenceTransportNodeInfoList, error) {
	typeConverter := tIface.connector.TypeConverter()
	executionContext := tIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(transportNodesPatchInputType(), typeConverter)
	sv.AddStructField("SiteId", siteIdParam)
	sv.AddStructField("IntelligenceTransportNodeInfoList", intelligenceTransportNodeInfoListParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.IntelligenceTransportNodeInfoList
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := transportNodesPatchRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	tIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := tIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.sites.intelligence.transport_nodes", "patch", inputDataValue, executionContext)
	var emptyOutput model.IntelligenceTransportNodeInfoList
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), transportNodesPatchOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.IntelligenceTransportNodeInfoList), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (tIface *transportNodesClient) Patch0(siteIdParam string, tnIdParam string, intelligenceTransportNodeInfoParam model.IntelligenceTransportNodeInfo) (model.IntelligenceTransportNodeInfo, error) {
	typeConverter := tIface.connector.TypeConverter()
	executionContext := tIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(transportNodesPatch0InputType(), typeConverter)
	sv.AddStructField("SiteId", siteIdParam)
	sv.AddStructField("TnId", tnIdParam)
	sv.AddStructField("IntelligenceTransportNodeInfo", intelligenceTransportNodeInfoParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.IntelligenceTransportNodeInfo
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := transportNodesPatch0RestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	tIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := tIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.sites.intelligence.transport_nodes", "patch_0", inputDataValue, executionContext)
	var emptyOutput model.IntelligenceTransportNodeInfo
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), transportNodesPatch0OutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.IntelligenceTransportNodeInfo), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (tIface *transportNodesClient) Update(siteIdParam string, tnIdParam string, intelligenceTransportNodeInfoParam model.IntelligenceTransportNodeInfo) (model.IntelligenceTransportNodeInfo, error) {
	typeConverter := tIface.connector.TypeConverter()
	executionContext := tIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(transportNodesUpdateInputType(), typeConverter)
	sv.AddStructField("SiteId", siteIdParam)
	sv.AddStructField("TnId", tnIdParam)
	sv.AddStructField("IntelligenceTransportNodeInfo", intelligenceTransportNodeInfoParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.IntelligenceTransportNodeInfo
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := transportNodesUpdateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	tIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := tIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.sites.intelligence.transport_nodes", "update", inputDataValue, executionContext)
	var emptyOutput model.IntelligenceTransportNodeInfo
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), transportNodesUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.IntelligenceTransportNodeInfo), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
