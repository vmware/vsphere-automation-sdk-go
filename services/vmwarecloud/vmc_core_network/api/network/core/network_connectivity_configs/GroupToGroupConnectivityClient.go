// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: GroupToGroupConnectivity
// Used by client-side stubs.

package network_connectivity_configs

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	vmwarecloudVmc_core_networkModel "github.com/vmware/vsphere-automation-sdk-go/services/vmwarecloud/vmc_core_network/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type GroupToGroupConnectivityClient interface {

	// Get group-to-group-connectivity details for a specific group
	//
	// @param orgIdParam organization identifier
	// @param idParam resource identifier
	// @return OK
	//
	// @throws Unauthenticated Unauthorized
	// @throws Unauthorized Forbidden
	// @throws NotFound Cannot find the group with given identifier
	GetGroupToGroupConnectivity(orgIdParam string, idParam string) (vmwarecloudVmc_core_networkModel.GroupToGroupConnectivity, error)

	// Get group-to-group-connectivity details for the specified id
	//
	// @param orgIdParam organization identifier
	// @param idParam resource identifier
	// @param groupToGroupConnectivityIdParam Group To Group Connectivity Id
	// @return OK
	//
	// @throws Unauthenticated Unauthorized
	// @throws Unauthorized Forbidden
	// @throws NotFound Cannot find the group to group connectivity details with given identifier
	GetGroupToGroupConnectivityById(orgIdParam string, idParam string, groupToGroupConnectivityIdParam string) (vmwarecloudVmc_core_networkModel.GroupToGroupConnectivity, error)
}

type groupToGroupConnectivityClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewGroupToGroupConnectivityClient(connector vapiProtocolClient_.Connector) *groupToGroupConnectivityClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.vmwarecloud.vmc_core_network.api.network.core.network_connectivity_configs.group_to_group_connectivity")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"get_group_to_group_connectivity":       vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get_group_to_group_connectivity"),
		"get_group_to_group_connectivity_by_id": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get_group_to_group_connectivity_by_id"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	gIface := groupToGroupConnectivityClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &gIface
}

func (gIface *groupToGroupConnectivityClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := gIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (gIface *groupToGroupConnectivityClient) GetGroupToGroupConnectivity(orgIdParam string, idParam string) (vmwarecloudVmc_core_networkModel.GroupToGroupConnectivity, error) {
	typeConverter := gIface.connector.TypeConverter()
	executionContext := gIface.connector.NewExecutionContext()
	operationRestMetaData := groupToGroupConnectivityGetGroupToGroupConnectivityRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(groupToGroupConnectivityGetGroupToGroupConnectivityInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("Id", idParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput vmwarecloudVmc_core_networkModel.GroupToGroupConnectivity
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := gIface.connector.GetApiProvider().Invoke("com.vmware.vmwarecloud.vmc_core_network.api.network.core.network_connectivity_configs.group_to_group_connectivity", "get_group_to_group_connectivity", inputDataValue, executionContext)
	var emptyOutput vmwarecloudVmc_core_networkModel.GroupToGroupConnectivity
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), GroupToGroupConnectivityGetGroupToGroupConnectivityOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(vmwarecloudVmc_core_networkModel.GroupToGroupConnectivity), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), gIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (gIface *groupToGroupConnectivityClient) GetGroupToGroupConnectivityById(orgIdParam string, idParam string, groupToGroupConnectivityIdParam string) (vmwarecloudVmc_core_networkModel.GroupToGroupConnectivity, error) {
	typeConverter := gIface.connector.TypeConverter()
	executionContext := gIface.connector.NewExecutionContext()
	operationRestMetaData := groupToGroupConnectivityGetGroupToGroupConnectivityByIdRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(groupToGroupConnectivityGetGroupToGroupConnectivityByIdInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("Id", idParam)
	sv.AddStructField("GroupToGroupConnectivityId", groupToGroupConnectivityIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput vmwarecloudVmc_core_networkModel.GroupToGroupConnectivity
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := gIface.connector.GetApiProvider().Invoke("com.vmware.vmwarecloud.vmc_core_network.api.network.core.network_connectivity_configs.group_to_group_connectivity", "get_group_to_group_connectivity_by_id", inputDataValue, executionContext)
	var emptyOutput vmwarecloudVmc_core_networkModel.GroupToGroupConnectivity
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), GroupToGroupConnectivityGetGroupToGroupConnectivityByIdOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(vmwarecloudVmc_core_networkModel.GroupToGroupConnectivity), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), gIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
