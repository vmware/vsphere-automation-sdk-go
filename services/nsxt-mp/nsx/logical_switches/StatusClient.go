// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Status
// Used by client-side stubs.

package logical_switches

import (
	"github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
)

const _ = core.SupportedByRuntimeVersion1

type StatusClient interface {

	// Returns Operational status of all logical switches. The query parameter \"source=realtime\" is not supported.
	//
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param diagnosticParam Flag to enable showing of transit logical switch. (optional, default to false)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @param sourceParam Data source type. (optional)
	// @param switchTypeParam Logical Switch type (optional)
	// @param switchingProfileIdParam Switching Profile identifier (optional)
	// @param transportTypeParam Mode of transport supported in the transport zone for this logical switch (optional)
	// @param transportZoneIdParam Transport zone identifier (optional)
	// @param uplinkTeamingPolicyNameParam The logical switch's uplink teaming policy name (optional)
	// @param vlanParam Virtual Local Area Network Identifier (optional)
	// @param vniParam VNI of the OVERLAY LogicalSwitch(es) to return. (optional)
	// @return com.vmware.nsx.model.LogicalSwitchStatusSummary
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(cursorParam *string, diagnosticParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string, sourceParam *string, switchTypeParam *string, switchingProfileIdParam *string, transportTypeParam *string, transportZoneIdParam *string, uplinkTeamingPolicyNameParam *string, vlanParam *int64, vniParam *int64) (model.LogicalSwitchStatusSummary, error)
}

type statusClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewStatusClient(connector client.Connector) *statusClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.nsx.logical_switches.status")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"get": core.NewMethodIdentifier(interfaceIdentifier, "get"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	sIface := statusClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &sIface
}

func (sIface *statusClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := sIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (sIface *statusClient) Get(cursorParam *string, diagnosticParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string, sourceParam *string, switchTypeParam *string, switchingProfileIdParam *string, transportTypeParam *string, transportZoneIdParam *string, uplinkTeamingPolicyNameParam *string, vlanParam *int64, vniParam *int64) (model.LogicalSwitchStatusSummary, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(statusGetInputType(), typeConverter)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("Diagnostic", diagnosticParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	sv.AddStructField("Source", sourceParam)
	sv.AddStructField("SwitchType", switchTypeParam)
	sv.AddStructField("SwitchingProfileId", switchingProfileIdParam)
	sv.AddStructField("TransportType", transportTypeParam)
	sv.AddStructField("TransportZoneId", transportZoneIdParam)
	sv.AddStructField("UplinkTeamingPolicyName", uplinkTeamingPolicyNameParam)
	sv.AddStructField("Vlan", vlanParam)
	sv.AddStructField("Vni", vniParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.LogicalSwitchStatusSummary
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := statusGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx.logical_switches.status", "get", inputDataValue, executionContext)
	var emptyOutput model.LogicalSwitchStatusSummary
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), statusGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.LogicalSwitchStatusSummary), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
