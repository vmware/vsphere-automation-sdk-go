// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: ConsolidatedEffectiveIpAddressMembers
// Used by client-side stubs.

package ns_groups

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	nsxModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type ConsolidatedEffectiveIpAddressMembersClient interface {

	// Returns consolidated effective ip address members of the specified NSGroup. Applicable in case of federated environment. The response contains site-wise list of consolidated effective IP address members. In the response, for the local-site, the list will contain static and dynamicaly translated IPs. For the remote sites, the list will contain only the dynamically translated IPs. The static IPs will not be seen in the response of this API. Hence, user can refer to the local-site Ip response in the API results or the group definition to see the static IP membership of the Group. This API is applicable only for NSGroups containing either VirtualMachine, VIF, LogicalSwitch, LogicalPort or IPSet member type. For NSGroups containing other member types,it returns an empty list. Use the cursor value in the response to fetch the next page. If there is no cursor value for a response, it implies the last page in the results for the query.
	//
	// @param nsGroupIdParam NSGroup Id (required)
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param ipFilterParam IP address, range, or subnet (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param siteIdParam UUID of the site from which the effective IP addresses are to be fetched (optional)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @return com.vmware.nsx.model.ConsolidatedEffectiveIPAddressMemberListResult
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(nsGroupIdParam string, cursorParam *string, includedFieldsParam *string, ipFilterParam *string, pageSizeParam *int64, siteIdParam *string, sortAscendingParam *bool, sortByParam *string) (nsxModel.ConsolidatedEffectiveIPAddressMemberListResult, error)
}

type consolidatedEffectiveIpAddressMembersClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewConsolidatedEffectiveIpAddressMembersClient(connector vapiProtocolClient_.Connector) *consolidatedEffectiveIpAddressMembersClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx.ns_groups.consolidated_effective_ip_address_members")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"list": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "list"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	cIface := consolidatedEffectiveIpAddressMembersClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &cIface
}

func (cIface *consolidatedEffectiveIpAddressMembersClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := cIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (cIface *consolidatedEffectiveIpAddressMembersClient) List(nsGroupIdParam string, cursorParam *string, includedFieldsParam *string, ipFilterParam *string, pageSizeParam *int64, siteIdParam *string, sortAscendingParam *bool, sortByParam *string) (nsxModel.ConsolidatedEffectiveIPAddressMemberListResult, error) {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	operationRestMetaData := consolidatedEffectiveIpAddressMembersListRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(consolidatedEffectiveIpAddressMembersListInputType(), typeConverter)
	sv.AddStructField("NsGroupId", nsGroupIdParam)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("IpFilter", ipFilterParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SiteId", siteIdParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.ConsolidatedEffectiveIPAddressMemberListResult
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.nsx.ns_groups.consolidated_effective_ip_address_members", "list", inputDataValue, executionContext)
	var emptyOutput nsxModel.ConsolidatedEffectiveIPAddressMemberListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), ConsolidatedEffectiveIpAddressMembersListOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.ConsolidatedEffectiveIPAddressMemberListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
