// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Stats
// Used by client-side stubs.

package rules

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	nsxModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type StatsClient interface {

	// Get aggregated statistics for all rules for a given firewall section. The API only supports access to cached (source=cached) statistical data collected offline in the system. Data includes total number of packets, bytes, sessions counters and popularity index for a firewall rule and overall session count, max session count and max popularity index for all firewall rules on transport nodes or edge nodes. Aggregated statistics like maximum popularity index, maximum session count and total session count are computed with lower frequency compared to individual generic rule statistics, hence they may have a computation delay up to 15 minutes to reflect in response to this API.
	//
	//  Use the following Policy API -
	//  GET /policy/api/v1/infra/domains/<domain-id>/security-policies/<security-policy-id>/statistics
	//
	// Deprecated: This API element is deprecated.
	//
	// @param sectionIdParam (required)
	// @param sourceParam Data source type. (optional)
	// @return com.vmware.nsx.model.FirewallStatsList
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(sectionIdParam string, sourceParam *string) (nsxModel.FirewallStatsList, error)

	// Get aggregated statistics for a rule for given firewall section. The API only supports access to cached (source=cached) statistical data collected offline in the system. Data includes total number of packets, bytes, sessions counters and popularity index for a firewall rule and overall session count, max session count and max popularity index for all firewall rules on transport nodes or edge nodes. Aggregated statistics like maximum popularity index, maximum session count and total session count are computed with lower frequency compared to individual generic rule statistics, hence they may have a computation delay up to 15 minutes to reflect in response to this API.
	//
	//  Use the following Policy API -
	//  GET /policy/api/v1/infra/domains/<domain-id>/security-policies/<security-policy-id>/rules/<rule-id>/statistics
	//
	// Deprecated: This API element is deprecated.
	//
	// @param sectionIdParam (required)
	// @param ruleIdParam (required)
	// @param sourceParam Data source type. (optional)
	// @return com.vmware.nsx.model.FirewallStats
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get0(sectionIdParam string, ruleIdParam string, sourceParam *string) (nsxModel.FirewallStats, error)
}

type statsClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewStatsClient(connector vapiProtocolClient_.Connector) *statsClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx.firewall.sections.rules.stats")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"get":   vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
		"get_0": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get_0"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	sIface := statsClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &sIface
}

func (sIface *statsClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := sIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (sIface *statsClient) Get(sectionIdParam string, sourceParam *string) (nsxModel.FirewallStatsList, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	operationRestMetaData := statsGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(statsGetInputType(), typeConverter)
	sv.AddStructField("SectionId", sectionIdParam)
	sv.AddStructField("Source", sourceParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.FirewallStatsList
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx.firewall.sections.rules.stats", "get", inputDataValue, executionContext)
	var emptyOutput nsxModel.FirewallStatsList
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), StatsGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.FirewallStatsList), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *statsClient) Get0(sectionIdParam string, ruleIdParam string, sourceParam *string) (nsxModel.FirewallStats, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	operationRestMetaData := statsGet0RestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(statsGet0InputType(), typeConverter)
	sv.AddStructField("SectionId", sectionIdParam)
	sv.AddStructField("RuleId", ruleIdParam)
	sv.AddStructField("Source", sourceParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.FirewallStats
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx.firewall.sections.rules.stats", "get_0", inputDataValue, executionContext)
	var emptyOutput nsxModel.FirewallStats
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), StatsGet0OutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.FirewallStats), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
