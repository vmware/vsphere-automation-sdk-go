// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: RemoteMac
// Used by client-side stubs.

package sessions

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/core"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/lib"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/protocol/client"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt/model"
)

const _ = core.SupportedByRuntimeVersion1

type RemoteMacClient interface {

	// Returns L2Vpn session remote macs for a logical switch. Data is fetched from enforcement point. This API is deprecated. Please use GET /infra/tier-0s/<tier-0-id>/l2vpn-services/<service-id>/ sessions/<session-id>/remote-mac instead.
	//
	// @param tier0IdParam (required)
	// @param localeServiceIdParam (required)
	// @param serviceIdParam (required)
	// @param sessionIdParam (required)
	// @param enforcementPointPathParam String Path of the enforcement point (optional)
	// @param segmentPathParam Segment Path (optional)
	// @return com.vmware.nsx_policy.model.AggregateL2VpnSessionRemoteMac
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(tier0IdParam string, localeServiceIdParam string, serviceIdParam string, sessionIdParam string, enforcementPointPathParam *string, segmentPathParam *string) (model.AggregateL2VpnSessionRemoteMac, error)
}

type remoteMacClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewRemoteMacClient(connector client.Connector) *remoteMacClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.nsx_policy.infra.tier_0s.locale_services.l2vpn_services.sessions.remote_mac")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"get": core.NewMethodIdentifier(interfaceIdentifier, "get"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	rIface := remoteMacClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &rIface
}

func (rIface *remoteMacClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := rIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (rIface *remoteMacClient) Get(tier0IdParam string, localeServiceIdParam string, serviceIdParam string, sessionIdParam string, enforcementPointPathParam *string, segmentPathParam *string) (model.AggregateL2VpnSessionRemoteMac, error) {
	typeConverter := rIface.connector.TypeConverter()
	executionContext := rIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(remoteMacGetInputType(), typeConverter)
	sv.AddStructField("Tier0Id", tier0IdParam)
	sv.AddStructField("LocaleServiceId", localeServiceIdParam)
	sv.AddStructField("ServiceId", serviceIdParam)
	sv.AddStructField("SessionId", sessionIdParam)
	sv.AddStructField("EnforcementPointPath", enforcementPointPathParam)
	sv.AddStructField("SegmentPath", segmentPathParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.AggregateL2VpnSessionRemoteMac
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := remoteMacGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	rIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := rIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.tier_0s.locale_services.l2vpn_services.sessions.remote_mac", "get", inputDataValue, executionContext)
	var emptyOutput model.AggregateL2VpnSessionRemoteMac
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), remoteMacGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.AggregateL2VpnSessionRemoteMac), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), rIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
