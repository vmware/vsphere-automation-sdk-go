// Copyright (c) 2019-2026 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: EffectiveProfiles
// Used by client-side stubs.

package service_configs

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	nsxModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type EffectiveProfilesClient interface {

	// Returns the effective profiles applied to the specified Resource.
	//  Please use APIs in the Policy section corresponding to the profile type - IPFIX_SWITCH_PROFILE - Policy > Monitoring > IPFIX > Switch IPFIX Profiles IPFIX_COLLECTOR_PROFILE - Policy > Monitoring > IPFIX > Switch IPFIX Collectors DFW_CPU_MEM_THRESHOLDS_PROFILE - Policy > Security > Security Profiles > Cpu Memory Thresholds Profiles DFW_SESSIONTIMER_PROFILE - Policy > Security > Security Profiles > DNS Security Profiles GI_SERVICE_PROFILE - Policy > Security > Service Insertion > Service References DFW_FLOOD_PROTECTION_PROFILE - Policy > Security > Security Profiles > Flood Protection Profiles GENERAL_SECURITY_SETTINGS_PROFILE - Policy > Security > Security Profiles > General Security Profiles DFW_DNS_PROFILE - Policy > Security > Security Profiles > DNS Security Profiles PACE_HOST_CONFIG_PROFILE SHA_PROFILE - Policy > Monitoring > System Health Agent LATENCY_STAT_PROFILE - Policy > Monitoring > Latency IP_DISCOVERY_SWITCHING_UPM_PROFILE - Policy > Networking > Connectivity > Segments > Segment Profiles > IP Discovery Profiles SYSTEM_HEALTH_PLUGIN_PROFILE - Policy > Monitoring > System Health Agent COMMON_AGENT_HOST_CONFIG_PROFILE ODS_PROFILE - Policy > Monitoring > Online Diagnostic System ODS_DYNAMIC_INSTANCE - Policy > Monitoring > Online Diagnostic System
	//
	// Deprecated: This API element is deprecated.
	//
	// @param resourceIdParam The resource for which the effective profiles are to be fetched (required)
	// @param resourceTypeParam This enum defines the valid Resource types to be used in effective profiles API (required)
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includedFieldsParam Note - this parameter currently only works when used with the search APIs /policy/api/v1/search/query and /policy/api/v1/search/dsl. It is ignored for other list APIs. (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param profileTypeParam Fetch effectivw profiles of the given profile_type (optional)
	// @param sortAscendingParam If true, results are sorted in ascending order (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @return com.vmware.nsx.model.EffectiveProfileListResult
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(resourceIdParam string, resourceTypeParam string, cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, profileTypeParam *string, sortAscendingParam *bool, sortByParam *string) (nsxModel.EffectiveProfileListResult, error)
}

type effectiveProfilesClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewEffectiveProfilesClient(connector vapiProtocolClient_.Connector) *effectiveProfilesClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx.service_configs.effective_profiles")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"list": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "list"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	eIface := effectiveProfilesClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &eIface
}

func (eIface *effectiveProfilesClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := eIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (eIface *effectiveProfilesClient) List(resourceIdParam string, resourceTypeParam string, cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, profileTypeParam *string, sortAscendingParam *bool, sortByParam *string) (nsxModel.EffectiveProfileListResult, error) {
	typeConverter := eIface.connector.TypeConverter()
	executionContext := eIface.connector.NewExecutionContext()
	operationRestMetaData := effectiveProfilesListRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(effectiveProfilesListInputType(), typeConverter)
	sv.AddStructField("ResourceId", resourceIdParam)
	sv.AddStructField("ResourceType", resourceTypeParam)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("ProfileType", profileTypeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.EffectiveProfileListResult
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := eIface.connector.GetApiProvider().Invoke("com.vmware.nsx.service_configs.effective_profiles", "list", inputDataValue, executionContext)
	var emptyOutput nsxModel.EffectiveProfileListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), EffectiveProfilesListOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.EffectiveProfileListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), eIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
