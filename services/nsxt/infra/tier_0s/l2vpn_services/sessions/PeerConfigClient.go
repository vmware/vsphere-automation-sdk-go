// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: PeerConfig
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

type PeerConfigClient interface {

	// Get peer config for the Tier-0 L2VPN session to configure the remote side of the tunnel. - no enforcement point path specified: L2VPN Session Peer Codes will be evaluated on each enforcement point. - enforcement point paths specified: L2VPN Session Peer Codes are evaluated only on the given enforcement points. API supported only when L2VPN Service is in Server Mode.
	//
	// @param tier0IdParam (required)
	// @param serviceIdParam (required)
	// @param sessionIdParam (required)
	// @param enforcementPointPathParam String Path of the enforcement point (optional)
	// @return com.vmware.nsx_policy.model.AggregateL2VPNSessionPeerConfig
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(tier0IdParam string, serviceIdParam string, sessionIdParam string, enforcementPointPathParam *string) (model.AggregateL2VPNSessionPeerConfig, error)
}

type peerConfigClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewPeerConfigClient(connector client.Connector) *peerConfigClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.nsx_policy.infra.tier_0s.l2vpn_services.sessions.peer_config")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"get": core.NewMethodIdentifier(interfaceIdentifier, "get"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	pIface := peerConfigClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &pIface
}

func (pIface *peerConfigClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := pIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (pIface *peerConfigClient) Get(tier0IdParam string, serviceIdParam string, sessionIdParam string, enforcementPointPathParam *string) (model.AggregateL2VPNSessionPeerConfig, error) {
	typeConverter := pIface.connector.TypeConverter()
	executionContext := pIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(peerConfigGetInputType(), typeConverter)
	sv.AddStructField("Tier0Id", tier0IdParam)
	sv.AddStructField("ServiceId", serviceIdParam)
	sv.AddStructField("SessionId", sessionIdParam)
	sv.AddStructField("EnforcementPointPath", enforcementPointPathParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.AggregateL2VPNSessionPeerConfig
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := peerConfigGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	pIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := pIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.tier_0s.l2vpn_services.sessions.peer_config", "get", inputDataValue, executionContext)
	var emptyOutput model.AggregateL2VPNSessionPeerConfig
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), peerConfigGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.AggregateL2VPNSessionPeerConfig), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), pIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
