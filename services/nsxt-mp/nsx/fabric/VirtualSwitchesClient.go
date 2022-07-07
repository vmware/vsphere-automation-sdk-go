// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: VirtualSwitches
// Used by client-side stubs.

package fabric

import (
	"github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
)

const _ = core.SupportedByRuntimeVersion1

type VirtualSwitchesClient interface {

	// Returns information about all virtual switches based on the request parameters.
	//
	// @param cmLocalIdParam Local Id of the virtual switch (optional)
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param discoveredNodeIdParam Discovered node ID (optional)
	// @param displayNameParam Display name of the virtual switch (optional)
	// @param externalIdParam External id of the virtual switch (optional)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param originIdParam ID of the compute manager (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @param uuidParam UUID of the switch (optional)
	// @return com.vmware.nsx.model.VirtualSwitchListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(cmLocalIdParam *string, cursorParam *string, discoveredNodeIdParam *string, displayNameParam *string, externalIdParam *string, includedFieldsParam *string, originIdParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string, uuidParam *string) (model.VirtualSwitchListResult, error)
}

type virtualSwitchesClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewVirtualSwitchesClient(connector client.Connector) *virtualSwitchesClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.nsx.fabric.virtual_switches")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"list": core.NewMethodIdentifier(interfaceIdentifier, "list"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	vIface := virtualSwitchesClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &vIface
}

func (vIface *virtualSwitchesClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := vIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (vIface *virtualSwitchesClient) List(cmLocalIdParam *string, cursorParam *string, discoveredNodeIdParam *string, displayNameParam *string, externalIdParam *string, includedFieldsParam *string, originIdParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string, uuidParam *string) (model.VirtualSwitchListResult, error) {
	typeConverter := vIface.connector.TypeConverter()
	executionContext := vIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(virtualSwitchesListInputType(), typeConverter)
	sv.AddStructField("CmLocalId", cmLocalIdParam)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("DiscoveredNodeId", discoveredNodeIdParam)
	sv.AddStructField("DisplayName", displayNameParam)
	sv.AddStructField("ExternalId", externalIdParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("OriginId", originIdParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	sv.AddStructField("Uuid", uuidParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.VirtualSwitchListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := virtualSwitchesListRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	vIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := vIface.connector.GetApiProvider().Invoke("com.vmware.nsx.fabric.virtual_switches", "list", inputDataValue, executionContext)
	var emptyOutput model.VirtualSwitchListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), virtualSwitchesListOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.VirtualSwitchListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), vIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
