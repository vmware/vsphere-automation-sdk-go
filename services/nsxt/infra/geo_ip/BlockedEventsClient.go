// Copyright (c) 2019-2025 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: BlockedEvents
// Used by client-side stubs.

package geo_ip

import (
	nsx_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type BlockedEventsClient interface {

    // Get Geo IP Blocked Events. The system will return the latest 25k events stored. User can use optional request parameters like source_country_code or source_ip_address to filter the response.
    //
    // @param cursorParam Cursor for getting next page of records (optional)
    // @param destinationCountryCodeParam Comma Separated Destination Country Codes of Geo IP Blocked Event (optional)
    // @param destinationIpAddressParam Comma Separated Destination IP Addresses of Geo IP Blocked Event (optional)
    // @param directionParam Comma Separated Directions of Traffic (optional)
    // @param includeAllProjectsParam (optional, default to false)
    // @param pageSizeParam Maximum number of results to return in this page (optional, default to 500)
    // @param ruleIdParam Comma Separated Gateway Firewall Rule Ids of Geo IP Blocked Event (optional)
    // @param sourceCountryCodeParam Comma Separated Source Country Codes of Geo IP Blocked Event (optional)
    // @param sourceIpAddressParam Comma Separated Source IP Addresses of Geo IP Blocked Event (optional)
    // @return com.vmware.nsx_policy.model.GeoIpBlockedEventsList
    //
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get(cursorParam *string, destinationCountryCodeParam *string, destinationIpAddressParam *string, directionParam *string, includeAllProjectsParam *bool, pageSizeParam *int64, ruleIdParam *string, sourceCountryCodeParam *string, sourceIpAddressParam *string) (nsx_policyModel.GeoIpBlockedEventsList, error)
}


type blockedEventsClient struct {
	connector           	   vapiProtocolClient_.Connector
	interfaceDefinition 	   vapiCore_.InterfaceDefinition
	errorsBindingMap           map[string]vapiBindings_.BindingType
}

func NewBlockedEventsClient(connector vapiProtocolClient_.Connector) *blockedEventsClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx_policy.infra.geo_ip.blocked_events")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"get": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	bIface := blockedEventsClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &bIface
}

func (bIface *blockedEventsClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := bIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (bIface *blockedEventsClient) Get(cursorParam *string, destinationCountryCodeParam *string, destinationIpAddressParam *string, directionParam *string, includeAllProjectsParam *bool, pageSizeParam *int64, ruleIdParam *string, sourceCountryCodeParam *string, sourceIpAddressParam *string) (nsx_policyModel.GeoIpBlockedEventsList, error) {
	typeConverter := bIface.connector.TypeConverter()
	executionContext := bIface.connector.NewExecutionContext()
	operationRestMetaData := blockedEventsGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(blockedEventsGetInputType(), typeConverter)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("DestinationCountryCode", destinationCountryCodeParam)
	sv.AddStructField("DestinationIpAddress", destinationIpAddressParam)
	sv.AddStructField("Direction", directionParam)
	sv.AddStructField("IncludeAllProjects", includeAllProjectsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("RuleId", ruleIdParam)
	sv.AddStructField("SourceCountryCode", sourceCountryCodeParam)
	sv.AddStructField("SourceIpAddress", sourceIpAddressParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.GeoIpBlockedEventsList
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := bIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.geo_ip.blocked_events", "get", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.GeoIpBlockedEventsList
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), BlockedEventsGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.GeoIpBlockedEventsList), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), bIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

