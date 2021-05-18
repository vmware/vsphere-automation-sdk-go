// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: ConsolidatedEffectiveIpAddresses
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

type ConsolidatedEffectiveIpAddressesClient interface {

	// Returns consolidated effective ip address members of the specified NSGroup. Applicable in case of federated environment. The response contains site-wise list of consolidated effective IP address members. In the response, for the local-site, the list will contain static and dynamicaly translated IPs. For the remote sites, the list will contain only the dynamically translated IPs. The static IPs will not be seen in the response of this API. Hence, user can refer to the local-site Ip response in the API results or the group definition to see the static IP membership of the Group. This API is applicable only for Global Groups containing (directly or via nesting) either VirtualMachine, VIF, Segment, SegmentPort or IPSet member type. Use the cursor value in the response to fetch the next page. If there is no cursor value for a response, it implies the last page in the results for the query.
	//
	// @param domainIdParam Domain id (required)
	// @param groupIdParam Group Id (required)
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param enforcementPointPathParam String Path of the enforcement point (optional)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param ipFilterParam IP address, range, or subnet (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param siteIdParam UUID of the site from which the effective IP addresses are to be fetched (optional)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @return com.vmware.nsx_policy.model.ConsolidatedEffectiveIPAddressMemberListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(domainIdParam string, groupIdParam string, cursorParam *string, enforcementPointPathParam *string, includedFieldsParam *string, ipFilterParam *string, pageSizeParam *int64, siteIdParam *string, sortAscendingParam *bool, sortByParam *string) (model.ConsolidatedEffectiveIPAddressMemberListResult, error)
}

type consolidatedEffectiveIpAddressesClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewConsolidatedEffectiveIpAddressesClient(connector client.Connector) *consolidatedEffectiveIpAddressesClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.nsx_policy.infra.domains.groups.members.consolidated_effective_ip_addresses")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"list": core.NewMethodIdentifier(interfaceIdentifier, "list"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	cIface := consolidatedEffectiveIpAddressesClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &cIface
}

func (cIface *consolidatedEffectiveIpAddressesClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := cIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (cIface *consolidatedEffectiveIpAddressesClient) List(domainIdParam string, groupIdParam string, cursorParam *string, enforcementPointPathParam *string, includedFieldsParam *string, ipFilterParam *string, pageSizeParam *int64, siteIdParam *string, sortAscendingParam *bool, sortByParam *string) (model.ConsolidatedEffectiveIPAddressMemberListResult, error) {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(consolidatedEffectiveIpAddressesListInputType(), typeConverter)
	sv.AddStructField("DomainId", domainIdParam)
	sv.AddStructField("GroupId", groupIdParam)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("EnforcementPointPath", enforcementPointPathParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("IpFilter", ipFilterParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SiteId", siteIdParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ConsolidatedEffectiveIPAddressMemberListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := consolidatedEffectiveIpAddressesListRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	cIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.domains.groups.members.consolidated_effective_ip_addresses", "list", inputDataValue, executionContext)
	var emptyOutput model.ConsolidatedEffectiveIPAddressMemberListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), consolidatedEffectiveIpAddressesListOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ConsolidatedEffectiveIPAddressMemberListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
