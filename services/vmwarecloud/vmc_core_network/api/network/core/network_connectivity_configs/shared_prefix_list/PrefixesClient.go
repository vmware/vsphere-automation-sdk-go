// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Prefixes
// Used by client-side stubs.

package shared_prefix_list

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	vmwarecloudVmc_core_networkModel "github.com/vmware/vsphere-automation-sdk-go/services/vmwarecloud/vmc_core_network/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type PrefixesClient interface {

	// get all prefix list entry of given shared prefix list id
	//
	// @param orgIdParam organization identifier
	// @param idParam resource identifier
	// @param sharedPrefixListIdParam Shared Prefix List Id
	// @param pageParam Number of requested page (0-based)
	// @param sizeParam Size of each page
	// @param sortParam Sort order of paged requests
	// @param filterParam Filtering criteria of paged requests
	// @return OK
	//
	// @throws Unauthenticated Unauthorized
	// @throws Unauthorized Forbidden
	// @throws NotFound Cannot find the shared prefix list with given identifier
	GetSharedPrefixListById(orgIdParam string, idParam string, sharedPrefixListIdParam string, pageParam *int64, sizeParam *int64, sortParam []string, filterParam []string) (vmwarecloudVmc_core_networkModel.SharedPrefixListEntryResponse, error)
}

type prefixesClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewPrefixesClient(connector vapiProtocolClient_.Connector) *prefixesClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.vmwarecloud.vmc_core_network.api.network.core.network_connectivity_configs.shared_prefix_list.prefixes")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"get_shared_prefix_list_by_id": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get_shared_prefix_list_by_id"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	pIface := prefixesClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &pIface
}

func (pIface *prefixesClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := pIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (pIface *prefixesClient) GetSharedPrefixListById(orgIdParam string, idParam string, sharedPrefixListIdParam string, pageParam *int64, sizeParam *int64, sortParam []string, filterParam []string) (vmwarecloudVmc_core_networkModel.SharedPrefixListEntryResponse, error) {
	typeConverter := pIface.connector.TypeConverter()
	executionContext := pIface.connector.NewExecutionContext()
	operationRestMetaData := prefixesGetSharedPrefixListByIdRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(prefixesGetSharedPrefixListByIdInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("Id", idParam)
	sv.AddStructField("SharedPrefixListId", sharedPrefixListIdParam)
	sv.AddStructField("Page", pageParam)
	sv.AddStructField("Size", sizeParam)
	sv.AddStructField("Sort", sortParam)
	sv.AddStructField("Filter", filterParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput vmwarecloudVmc_core_networkModel.SharedPrefixListEntryResponse
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := pIface.connector.GetApiProvider().Invoke("com.vmware.vmwarecloud.vmc_core_network.api.network.core.network_connectivity_configs.shared_prefix_list.prefixes", "get_shared_prefix_list_by_id", inputDataValue, executionContext)
	var emptyOutput vmwarecloudVmc_core_networkModel.SharedPrefixListEntryResponse
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), PrefixesGetSharedPrefixListByIdOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(vmwarecloudVmc_core_networkModel.SharedPrefixListEntryResponse), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), pIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
