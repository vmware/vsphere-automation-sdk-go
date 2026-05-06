// Copyright (c) 2019-2026 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Batch
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

type BatchClient interface {

	// Creates/Updates new service configs sent in batch request. This API returns ALL the service configs that are created/updated.
	//  Please use APIs in the Policy section corresponding to the profile type - IPFIX_SWITCH_PROFILE - Policy > Monitoring > IPFIX > Switch IPFIX Profiles IPFIX_COLLECTOR_PROFILE - Policy > Monitoring > IPFIX > Switch IPFIX Collectors DFW_CPU_MEM_THRESHOLDS_PROFILE - Policy > Security > Security Profiles > Cpu Memory Thresholds Profiles DFW_SESSIONTIMER_PROFILE - Policy > Security > Security Profiles > DNS Security Profiles GI_SERVICE_PROFILE - Policy > Security > Service Insertion > Service References DFW_FLOOD_PROTECTION_PROFILE - Policy > Security > Security Profiles > Flood Protection Profiles GENERAL_SECURITY_SETTINGS_PROFILE - Policy > Security > Security Profiles > General Security Profiles DFW_DNS_PROFILE - Policy > Security > Security Profiles > DNS Security Profiles PACE_HOST_CONFIG_PROFILE SHA_PROFILE - Policy > Monitoring > System Health Agent LATENCY_STAT_PROFILE - Policy > Monitoring > Latency IP_DISCOVERY_SWITCHING_UPM_PROFILE - Policy > Networking > Connectivity > Segments > Segment Profiles > IP Discovery Profiles SYSTEM_HEALTH_PLUGIN_PROFILE - Policy > Monitoring > System Health Agent COMMON_AGENT_HOST_CONFIG_PROFILE ODS_PROFILE - Policy > Monitoring > Online Diagnostic System ODS_DYNAMIC_INSTANCE - Policy > Monitoring > Online Diagnostic System
	//
	// Deprecated: This API element is deprecated.
	//
	// @param serviceConfigListParam (required)
	// @return com.vmware.nsx.model.ServiceConfigListResult
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Create(serviceConfigListParam nsxModel.ServiceConfigList) (nsxModel.ServiceConfigListResult, error)
}

type batchClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewBatchClient(connector vapiProtocolClient_.Connector) *batchClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx.service_configs.batch")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"create": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "create"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	bIface := batchClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &bIface
}

func (bIface *batchClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := bIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (bIface *batchClient) Create(serviceConfigListParam nsxModel.ServiceConfigList) (nsxModel.ServiceConfigListResult, error) {
	typeConverter := bIface.connector.TypeConverter()
	executionContext := bIface.connector.NewExecutionContext()
	operationRestMetaData := batchCreateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(batchCreateInputType(), typeConverter)
	sv.AddStructField("ServiceConfigList", serviceConfigListParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.ServiceConfigListResult
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := bIface.connector.GetApiProvider().Invoke("com.vmware.nsx.service_configs.batch", "create", inputDataValue, executionContext)
	var emptyOutput nsxModel.ServiceConfigListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), BatchCreateOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.ServiceConfigListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), bIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
