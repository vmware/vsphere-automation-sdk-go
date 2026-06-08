// Copyright (c) 2019-2026 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Status
// Used by client-side stubs.

package realized_state

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	nsx_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type StatusClient interface {

	// Get Consolidated Status of an intent object (with or without enforcement specific status details). The request is evaluated as follows: - <intent_path>: the request is evaluated on all enforcement points for the given intent without enforcement point specific details. - <intent_path, include_enforced_status=true>: the request is evaluated on all enforcement points for the given intent with enforcement point specific details. This is only supported when intent_path refers to an object of one of the following types: SECURITY_POLICY, GATEWAY_POLICY, IDS_SECURITY_POLICY, IDS_GATEWAY_POLICY, IDS_COMPRESSED_SIGNATURE, IDS_COMPRESSED_CUSTOM_SIGNATURE, TLS_POLICY, VPC_SECURITY_POLICY, VPC_GATEWAY_POLICY, PUBLIC_CLOUD_VPC_SECURITY_POLICY, Passing the path of an unsupported type (e.g. a rule) with include_enforced_status=true will result in an error. In such cases, pass the path of the parent policy object instead (e.g., the SecurityPolicy or GatewayPolicy that contains the rule).
	//
	// @param intentPathParam Policy Path referencing an intent object. (required)
	// @param includeEnforcedStatusParam Flag conveying whether to include detailed per-transport-node enforcement status in the response. When set to true, each item in consolidated_status_per_enforcement_point is returned as ConsolidatedStatusNsxT (with the enforced_status field populated). When set to false (default), each item is returned as ConsolidatedStatusPerEnforcementPoint (enforced_status absent). Use the resource_type field in each item to distinguish between the two types. This flag is only supported when intent_path refers to an object of one of the following types: SECURITY_POLICY, GATEWAY_POLICY, IDS_SECURITY_POLICY, IDS_GATEWAY_POLICY, IDS_COMPRESSED_SIGNATURE, IDS_COMPRESSED_CUSTOM_SIGNATURE, TLS_POLICY, VPC_SECURITY_POLICY, VPC_GATEWAY_POLICY, PUBLIC_CLOUD_VPC_SECURITY_POLICY, COMMUNICATION_MAP (deprecated alias for SECURITY_POLICY), EDGE_COMMUNICATION_MAP (deprecated alias for GATEWAY_POLICY). Setting this flag for any other type (e.g. a rule path) will result in an error. In such cases, pass the path of the parent policy object instead (e.g., the SecurityPolicy or GatewayPolicy that contains the rule). (optional, default to false)
	// @param sitePathParam Policy Path referencing a site. This is applicable only on a GlobalManager. If no site_path is specified, then based on the span of the intent the response will be fetched from the respective sites (optional)
	// @return com.vmware.nsx_policy.model.ConsolidatedRealizedStatus
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(intentPathParam string, includeEnforcedStatusParam *bool, sitePathParam *string) (nsx_policyModel.ConsolidatedRealizedStatus, error)
}

type statusClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewStatusClient(connector vapiProtocolClient_.Connector) *statusClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx_policy.global_infra.realized_state.status")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"get": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	sIface := statusClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &sIface
}

func (sIface *statusClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := sIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (sIface *statusClient) Get(intentPathParam string, includeEnforcedStatusParam *bool, sitePathParam *string) (nsx_policyModel.ConsolidatedRealizedStatus, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	operationRestMetaData := statusGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(statusGetInputType(), typeConverter)
	sv.AddStructField("IntentPath", intentPathParam)
	sv.AddStructField("IncludeEnforcedStatus", includeEnforcedStatusParam)
	sv.AddStructField("SitePath", sitePathParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.ConsolidatedRealizedStatus
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.global_infra.realized_state.status", "get", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.ConsolidatedRealizedStatus
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), StatusGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.ConsolidatedRealizedStatus), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
