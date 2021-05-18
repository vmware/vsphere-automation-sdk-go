// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: ManagementPlaneAPINetworkingLogicalRoutingAndServicesNAT
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

type ManagementPlaneAPINetworkingLogicalRoutingAndServicesNATClient interface {

	// Add a NAT rule in a specific logical router.
	//
	// @param logicalRouterIdParam (required)
	// @param natRuleParam (required)
	// @return com.vmware.model.NatRule
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	AddNatRule(logicalRouterIdParam string, natRuleParam model.NatRule) (model.NatRule, error)

	// Create multiple NAT rules in a specific logical router. The API succeeds only when all rules are accepted and created successfully. Any one validation voilation will fail the API, no rule will be created. The ruleIds of each rules can be found from the responsed message.
	//
	// @param logicalRouterIdParam (required)
	// @param natRuleListParam (required)
	// @return com.vmware.model.NatRuleList
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	AddNatRulesCreateMultiple(logicalRouterIdParam string, natRuleListParam model.NatRuleList) (model.NatRuleList, error)

	// Delete a specific NAT rule from a logical router
	//
	// @param logicalRouterIdParam (required)
	// @param ruleIdParam (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	DeleteNatRule(logicalRouterIdParam string, ruleIdParam string) error

	// Get a specific NAT rule from a given logical router
	//
	// @param logicalRouterIdParam (required)
	// @param ruleIdParam (required)
	// @return com.vmware.model.NatRule
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetNatRule(logicalRouterIdParam string, ruleIdParam string) (model.NatRule, error)

	// Returns the summation of statistics for all rules from all nodes for the Specified Logical Router. Also gives the per transport node statistics for provided logical router. The query parameter \"source=realtime\" is not supported.
	//
	// @param logicalRouterIdParam (required)
	// @param sourceParam Data source type. (optional)
	// @return com.vmware.model.NatStatisticsPerLogicalRouter
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetNatStatisticsPerLogicalRouter(logicalRouterIdParam string, sourceParam *string) (model.NatStatisticsPerLogicalRouter, error)

	// Returns the summation of statistics from all nodes for the Specified Logical Router NAT Rule. Query parameter \"source=realtime\" is the only supported source.
	//
	// @param logicalRouterIdParam (required)
	// @param ruleIdParam (required)
	// @param sourceParam Data source type. (optional)
	// @return com.vmware.model.NatStatisticsPerRule
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetNatStatisticsPerRule(logicalRouterIdParam string, ruleIdParam string, sourceParam *string) (model.NatStatisticsPerRule, error)

	// Returns the summation of statistics for all rules from all logical routers which are present on given transport node. Only cached statistics are supported. The query parameter \"source=realtime\" is not supported.
	//
	// @param nodeIdParam (required)
	// @param sourceParam Data source type. (optional)
	// @return com.vmware.model.NatStatisticsPerTransportNode
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetNatStatisticsPerTransportNode(nodeIdParam string, sourceParam *string) (model.NatStatisticsPerTransportNode, error)

	// Returns paginated list of all user defined NAT rules of the specific logical router. If a rule_type is provided, only the given type of rules will be returned. If no rule_type is specified, the rule_type will be defaulted to NATv4, i.e. only the NATv4 rules will be listed.
	//
	// @param logicalRouterIdParam (required)
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param ruleTypeParam Action type for getting NAT rules (optional)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @return com.vmware.model.NatRuleListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ListNatRules(logicalRouterIdParam string, cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, ruleTypeParam *string, sortAscendingParam *bool, sortByParam *string) (model.NatRuleListResult, error)

	// Update a specific NAT rule from a given logical router.
	//
	// @param logicalRouterIdParam (required)
	// @param ruleIdParam (required)
	// @param natRuleParam (required)
	// @return com.vmware.model.NatRule
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	UpdateNatRule(logicalRouterIdParam string, ruleIdParam string, natRuleParam model.NatRule) (model.NatRule, error)
}

type managementPlaneAPINetworkingLogicalRoutingAndServicesNATClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewManagementPlaneAPINetworkingLogicalRoutingAndServicesNATClient(connector client.Connector) *managementPlaneAPINetworkingLogicalRoutingAndServicesNATClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.api.management_plane_API_networking_logical_routing_and_services_NAT")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"add_nat_rule":                          core.NewMethodIdentifier(interfaceIdentifier, "add_nat_rule"),
		"add_nat_rules_create_multiple":         core.NewMethodIdentifier(interfaceIdentifier, "add_nat_rules_create_multiple"),
		"delete_nat_rule":                       core.NewMethodIdentifier(interfaceIdentifier, "delete_nat_rule"),
		"get_nat_rule":                          core.NewMethodIdentifier(interfaceIdentifier, "get_nat_rule"),
		"get_nat_statistics_per_logical_router": core.NewMethodIdentifier(interfaceIdentifier, "get_nat_statistics_per_logical_router"),
		"get_nat_statistics_per_rule":           core.NewMethodIdentifier(interfaceIdentifier, "get_nat_statistics_per_rule"),
		"get_nat_statistics_per_transport_node": core.NewMethodIdentifier(interfaceIdentifier, "get_nat_statistics_per_transport_node"),
		"list_nat_rules":                        core.NewMethodIdentifier(interfaceIdentifier, "list_nat_rules"),
		"update_nat_rule":                       core.NewMethodIdentifier(interfaceIdentifier, "update_nat_rule"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	mIface := managementPlaneAPINetworkingLogicalRoutingAndServicesNATClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &mIface
}

func (mIface *managementPlaneAPINetworkingLogicalRoutingAndServicesNATClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := mIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (mIface *managementPlaneAPINetworkingLogicalRoutingAndServicesNATClient) AddNatRule(logicalRouterIdParam string, natRuleParam model.NatRule) (model.NatRule, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingLogicalRoutingAndServicesNATAddNatRuleInputType(), typeConverter)
	sv.AddStructField("LogicalRouterId", logicalRouterIdParam)
	sv.AddStructField("NatRule", natRuleParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.NatRule
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingLogicalRoutingAndServicesNATAddNatRuleRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_logical_routing_and_services_NAT", "add_nat_rule", inputDataValue, executionContext)
	var emptyOutput model.NatRule
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingLogicalRoutingAndServicesNATAddNatRuleOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.NatRule), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingLogicalRoutingAndServicesNATClient) AddNatRulesCreateMultiple(logicalRouterIdParam string, natRuleListParam model.NatRuleList) (model.NatRuleList, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingLogicalRoutingAndServicesNATAddNatRulesCreateMultipleInputType(), typeConverter)
	sv.AddStructField("LogicalRouterId", logicalRouterIdParam)
	sv.AddStructField("NatRuleList", natRuleListParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.NatRuleList
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingLogicalRoutingAndServicesNATAddNatRulesCreateMultipleRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_logical_routing_and_services_NAT", "add_nat_rules_create_multiple", inputDataValue, executionContext)
	var emptyOutput model.NatRuleList
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingLogicalRoutingAndServicesNATAddNatRulesCreateMultipleOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.NatRuleList), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingLogicalRoutingAndServicesNATClient) DeleteNatRule(logicalRouterIdParam string, ruleIdParam string) error {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingLogicalRoutingAndServicesNATDeleteNatRuleInputType(), typeConverter)
	sv.AddStructField("LogicalRouterId", logicalRouterIdParam)
	sv.AddStructField("RuleId", ruleIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingLogicalRoutingAndServicesNATDeleteNatRuleRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_logical_routing_and_services_NAT", "delete_nat_rule", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingLogicalRoutingAndServicesNATClient) GetNatRule(logicalRouterIdParam string, ruleIdParam string) (model.NatRule, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingLogicalRoutingAndServicesNATGetNatRuleInputType(), typeConverter)
	sv.AddStructField("LogicalRouterId", logicalRouterIdParam)
	sv.AddStructField("RuleId", ruleIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.NatRule
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingLogicalRoutingAndServicesNATGetNatRuleRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_logical_routing_and_services_NAT", "get_nat_rule", inputDataValue, executionContext)
	var emptyOutput model.NatRule
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingLogicalRoutingAndServicesNATGetNatRuleOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.NatRule), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingLogicalRoutingAndServicesNATClient) GetNatStatisticsPerLogicalRouter(logicalRouterIdParam string, sourceParam *string) (model.NatStatisticsPerLogicalRouter, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingLogicalRoutingAndServicesNATGetNatStatisticsPerLogicalRouterInputType(), typeConverter)
	sv.AddStructField("LogicalRouterId", logicalRouterIdParam)
	sv.AddStructField("Source", sourceParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.NatStatisticsPerLogicalRouter
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingLogicalRoutingAndServicesNATGetNatStatisticsPerLogicalRouterRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_logical_routing_and_services_NAT", "get_nat_statistics_per_logical_router", inputDataValue, executionContext)
	var emptyOutput model.NatStatisticsPerLogicalRouter
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingLogicalRoutingAndServicesNATGetNatStatisticsPerLogicalRouterOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.NatStatisticsPerLogicalRouter), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingLogicalRoutingAndServicesNATClient) GetNatStatisticsPerRule(logicalRouterIdParam string, ruleIdParam string, sourceParam *string) (model.NatStatisticsPerRule, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingLogicalRoutingAndServicesNATGetNatStatisticsPerRuleInputType(), typeConverter)
	sv.AddStructField("LogicalRouterId", logicalRouterIdParam)
	sv.AddStructField("RuleId", ruleIdParam)
	sv.AddStructField("Source", sourceParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.NatStatisticsPerRule
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingLogicalRoutingAndServicesNATGetNatStatisticsPerRuleRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_logical_routing_and_services_NAT", "get_nat_statistics_per_rule", inputDataValue, executionContext)
	var emptyOutput model.NatStatisticsPerRule
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingLogicalRoutingAndServicesNATGetNatStatisticsPerRuleOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.NatStatisticsPerRule), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingLogicalRoutingAndServicesNATClient) GetNatStatisticsPerTransportNode(nodeIdParam string, sourceParam *string) (model.NatStatisticsPerTransportNode, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingLogicalRoutingAndServicesNATGetNatStatisticsPerTransportNodeInputType(), typeConverter)
	sv.AddStructField("NodeId", nodeIdParam)
	sv.AddStructField("Source", sourceParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.NatStatisticsPerTransportNode
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingLogicalRoutingAndServicesNATGetNatStatisticsPerTransportNodeRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_logical_routing_and_services_NAT", "get_nat_statistics_per_transport_node", inputDataValue, executionContext)
	var emptyOutput model.NatStatisticsPerTransportNode
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingLogicalRoutingAndServicesNATGetNatStatisticsPerTransportNodeOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.NatStatisticsPerTransportNode), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingLogicalRoutingAndServicesNATClient) ListNatRules(logicalRouterIdParam string, cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, ruleTypeParam *string, sortAscendingParam *bool, sortByParam *string) (model.NatRuleListResult, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingLogicalRoutingAndServicesNATListNatRulesInputType(), typeConverter)
	sv.AddStructField("LogicalRouterId", logicalRouterIdParam)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("RuleType", ruleTypeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.NatRuleListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingLogicalRoutingAndServicesNATListNatRulesRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_logical_routing_and_services_NAT", "list_nat_rules", inputDataValue, executionContext)
	var emptyOutput model.NatRuleListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingLogicalRoutingAndServicesNATListNatRulesOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.NatRuleListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingLogicalRoutingAndServicesNATClient) UpdateNatRule(logicalRouterIdParam string, ruleIdParam string, natRuleParam model.NatRule) (model.NatRule, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingLogicalRoutingAndServicesNATUpdateNatRuleInputType(), typeConverter)
	sv.AddStructField("LogicalRouterId", logicalRouterIdParam)
	sv.AddStructField("RuleId", ruleIdParam)
	sv.AddStructField("NatRule", natRuleParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.NatRule
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingLogicalRoutingAndServicesNATUpdateNatRuleRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_logical_routing_and_services_NAT", "update_nat_rule", inputDataValue, executionContext)
	var emptyOutput model.NatRule
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingLogicalRoutingAndServicesNATUpdateNatRuleOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.NatRule), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
