// Copyright (c) 2019-2026 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: VdsScopeMembers
// Used by client-side stubs.

package subnets

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	nsx_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type VdsScopeMembersClient interface {

	// Retrieves all Virtual Distributed Switches (VDSes) associated with the specified VPC subnet with pagination and filtering support. Only includes VDSes with vCenter version 9.2 or higher to ensure compatibility. The VDS details are fetched from the VdsFolderSpan table, which contains precomputed mappings of network IDs to VDS collections. This provides efficient access to VDS information without real-time computation. Supports filtering by compute manager name, datacenter name, VDS name, and VDS external ID. Results are paginated using cursor-based pagination with configurable page size. Returns empty array if no compatible VDSes are found.
	//
	// @param orgIdParam (required)
	// @param projectIdParam (required)
	// @param vpcIdParam (required)
	// @param subnetIdParam (required)
	// @param cmNameParam Filter VDS scope members by compute manager name. Supports partial, case-insensitive matching. (optional)
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param dcNameParam Filter VDS scope members by datacenter name. Supports partial, case-insensitive matching. (optional)
	// @param includedFieldsParam Note - this parameter currently only works when used with the search APIs /policy/api/v1/search/query and /policy/api/v1/search/dsl. It is ignored for other list APIs. (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam If true, results are sorted in ascending order (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @param vdsExternalIdParam Filter VDS scope members by VDS external identifier. Supports partial, case-insensitive matching. (optional)
	// @param vdsNameParam Filter VDS scope members by VDS name. Supports partial, case-insensitive matching. (optional)
	// @return com.vmware.nsx_policy.model.VdsScopeMembersListResult
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(orgIdParam string, projectIdParam string, vpcIdParam string, subnetIdParam string, cmNameParam *string, cursorParam *string, dcNameParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string, vdsExternalIdParam *string, vdsNameParam *string) (nsx_policyModel.VdsScopeMembersListResult, error)
}

type vdsScopeMembersClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewVdsScopeMembersClient(connector vapiProtocolClient_.Connector) *vdsScopeMembersClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx_policy.orgs.projects.vpcs.subnets.vds_scope_members")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"list": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "list"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	vIface := vdsScopeMembersClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &vIface
}

func (vIface *vdsScopeMembersClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := vIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (vIface *vdsScopeMembersClient) List(orgIdParam string, projectIdParam string, vpcIdParam string, subnetIdParam string, cmNameParam *string, cursorParam *string, dcNameParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string, vdsExternalIdParam *string, vdsNameParam *string) (nsx_policyModel.VdsScopeMembersListResult, error) {
	typeConverter := vIface.connector.TypeConverter()
	executionContext := vIface.connector.NewExecutionContext()
	operationRestMetaData := vdsScopeMembersListRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(vdsScopeMembersListInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("ProjectId", projectIdParam)
	sv.AddStructField("VpcId", vpcIdParam)
	sv.AddStructField("SubnetId", subnetIdParam)
	sv.AddStructField("CmName", cmNameParam)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("DcName", dcNameParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	sv.AddStructField("VdsExternalId", vdsExternalIdParam)
	sv.AddStructField("VdsName", vdsNameParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.VdsScopeMembersListResult
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := vIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.orgs.projects.vpcs.subnets.vds_scope_members", "list", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.VdsScopeMembersListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), VdsScopeMembersListOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.VdsScopeMembersListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), vIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
