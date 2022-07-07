// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: CommuntyLists
// Used by client-side stubs.

package bgp

import (
	"github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
)

const _ = core.SupportedByRuntimeVersion1

type CommuntyListsClient interface {

	// Update a specific BGP community list from a Logical Router
	//
	// @param logicalRouterIdParam (required)
	// @param communityListIdParam (required)
	// @param bGPCommunityListParam (required)
	// @return com.vmware.nsx.model.BGPCommunityList
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(logicalRouterIdParam string, communityListIdParam string, bGPCommunityListParam model.BGPCommunityList) (model.BGPCommunityList, error)
}

type communtyListsClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewCommuntyListsClient(connector client.Connector) *communtyListsClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.nsx.logical_routers.routing.bgp.communty_lists")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"update": core.NewMethodIdentifier(interfaceIdentifier, "update"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	cIface := communtyListsClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &cIface
}

func (cIface *communtyListsClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := cIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (cIface *communtyListsClient) Update(logicalRouterIdParam string, communityListIdParam string, bGPCommunityListParam model.BGPCommunityList) (model.BGPCommunityList, error) {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(communtyListsUpdateInputType(), typeConverter)
	sv.AddStructField("LogicalRouterId", logicalRouterIdParam)
	sv.AddStructField("CommunityListId", communityListIdParam)
	sv.AddStructField("BGPCommunityList", bGPCommunityListParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.BGPCommunityList
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := communtyListsUpdateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	cIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.nsx.logical_routers.routing.bgp.communty_lists", "update", inputDataValue, executionContext)
	var emptyOutput model.BGPCommunityList
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), communtyListsUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.BGPCommunityList), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
