// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: TransportNodes
// Used by client-side stubs.

package members

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

	// Get Transport Nodes that belong to this Group. This API is applicable for Groups containing Transport Node member type. For Groups containing other member types, an empty list is returned.
	//
	// @param orgIdParam (required)
	// @param deploymentGroupIdParam (required)
	// @param vhcIdParam (required)
	// @param domainIdParam Domain id (required)
	// @param groupIdParam Group Id (required)
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param enforcementPointPathParam String Path of the enforcement point (optional)
	// @param includeMarkForDeleteObjectsParam Include objects that are marked for deletion in results (optional, default to false)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @return com.vmware.nsx_policy.model.PolicyGroupMembersListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(orgIdParam string, deploymentGroupIdParam string, vhcIdParam string, domainIdParam string, groupIdParam string, cursorParam *string, enforcementPointPathParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.PolicyGroupMembersListResult, error)
}

type transportNodesClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewTransportNodesClient(connector client.Connector) *transportNodesClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.nsx_policy.org_root.orgs.deployment_groups.vhcs.global_infra.domains.groups.members.transport_nodes")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"list": core.NewMethodIdentifier(interfaceIdentifier, "list"),
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

func (tIface *transportNodesClient) List(orgIdParam string, deploymentGroupIdParam string, vhcIdParam string, domainIdParam string, groupIdParam string, cursorParam *string, enforcementPointPathParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.PolicyGroupMembersListResult, error) {
	typeConverter := tIface.connector.TypeConverter()
	executionContext := tIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(transportNodesListInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("DeploymentGroupId", deploymentGroupIdParam)
	sv.AddStructField("VhcId", vhcIdParam)
	sv.AddStructField("DomainId", domainIdParam)
	sv.AddStructField("GroupId", groupIdParam)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("EnforcementPointPath", enforcementPointPathParam)
	sv.AddStructField("IncludeMarkForDeleteObjects", includeMarkForDeleteObjectsParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.PolicyGroupMembersListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := transportNodesListRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	tIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := tIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.org_root.orgs.deployment_groups.vhcs.global_infra.domains.groups.members.transport_nodes", "list", inputDataValue, executionContext)
	var emptyOutput model.PolicyGroupMembersListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), transportNodesListOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.PolicyGroupMembersListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
