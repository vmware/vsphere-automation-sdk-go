// Copyright (c) 2019-2026 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: State
// Used by client-side stubs.

package distributed_vlan_connections

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	nsx_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type StateClient interface {

	// Get validation state for the specified Distributed VLAN Connection. Use network_span_path or zone_external_id to filter nodes by span or availability zone. Use status to filter nodes by validation status. Use host_transport_node_path and include_node_details=true to query uplink-level validation details on a specific host.
	//
	// @param distributedVlanConnectionIdParam (required)
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param hostTransportNodePathParam Policy path of a specific host transport node. If specified, only the state for this host transport node will be returned. (optional)
	// @param includeNodeDetailsParam Flag indicating whether to query live uplink and gateway validation details from the host transport node. If set to true, node_detail_status will be populated. (optional, default to false)
	// @param includedFieldsParam Note - this parameter currently only works when used with the search APIs /policy/api/v1/search/query and /policy/api/v1/search/dsl. It is ignored for other list APIs. (optional)
	// @param networkSpanPathParam Filter nodes by the policy path of a cluster-based network span. Cannot be combined with zone_external_id. (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam If true, results are sorted in ascending order (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @param statusParam The status of the validation on a transport node. VALIDATION_UNKNOWN - Validation has not been performed yet, or status is unknown. VALIDATION_IN_PROGRESS - The validation is running on this transport node. VALIDATION_SUCCESS - The validation is successful on this transport node. All gateway addresses of the Distributed VLAN Connection on this transport node are validated successfully. VALIDATION_TIMED_OUT - The validation timed out on this transport node. VALIDATION_FAILED - There is a validation failure on this transport node. Any validation failure on any gateway address of the Distributed VLAN Connection on this transport node is considered a validation failure. VALIDATION_REALIZATION_FAILURE_HOST_SWITCH_NOT_FOUND - Validation realization failed on this transport node because the host switch of the default overlay transport zone was not found. VALIDATION_REALIZATION_FAILURE_NO_CONNECTED_UPLINK - Validation realization failed on this transport node because there are no connected uplinks on the host switch. VALIDATION_REALIZATION_FAILURE_INTERNAL_ERROR - Validation realization failed on this transport node because of an internal error. (optional)
	// @param zoneExternalIdParam Filter nodes by the external ID of an availability zone. Cannot be combined with network_span_path. (optional)
	// @return com.vmware.nsx_policy.model.DistributedVlanConnectionState
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(distributedVlanConnectionIdParam string, cursorParam *string, hostTransportNodePathParam *string, includeNodeDetailsParam *bool, includedFieldsParam *string, networkSpanPathParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string, statusParam *string, zoneExternalIdParam *string) (nsx_policyModel.DistributedVlanConnectionState, error)
}

type stateClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewStateClient(connector vapiProtocolClient_.Connector) *stateClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx_policy.infra.distributed_vlan_connections.state")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"list": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "list"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	sIface := stateClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &sIface
}

func (sIface *stateClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := sIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (sIface *stateClient) List(distributedVlanConnectionIdParam string, cursorParam *string, hostTransportNodePathParam *string, includeNodeDetailsParam *bool, includedFieldsParam *string, networkSpanPathParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string, statusParam *string, zoneExternalIdParam *string) (nsx_policyModel.DistributedVlanConnectionState, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	operationRestMetaData := stateListRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(stateListInputType(), typeConverter)
	sv.AddStructField("DistributedVlanConnectionId", distributedVlanConnectionIdParam)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("HostTransportNodePath", hostTransportNodePathParam)
	sv.AddStructField("IncludeNodeDetails", includeNodeDetailsParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("NetworkSpanPath", networkSpanPathParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	sv.AddStructField("Status", statusParam)
	sv.AddStructField("ZoneExternalId", zoneExternalIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.DistributedVlanConnectionState
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.distributed_vlan_connections.state", "list", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.DistributedVlanConnectionState
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), StateListOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.DistributedVlanConnectionState), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
