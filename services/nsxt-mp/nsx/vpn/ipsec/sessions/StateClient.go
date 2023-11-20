// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: State
// Used by client-side stubs.

package sessions

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	nsxModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type StateClient interface {

	// Return realized state information of a ipsec vpn session. Any configuration update that affects the ipsec vpn session can use this API to get its realized state by passing a request_id returned by the configuration change operation. e.g. Update configuration of ipsec vpn session, service, endpoints, profiles, etc. It will return a service disabled error, if the ipsec vpn service associated with the session is disabled.
	//
	//  Please use below Policy APIs.
	//  GET /policy/api/v1/infra/realized-state/realized-entity?realized_path=/policy/api/v1/infra/tier-0s/<tier-0-id>/ipsec-vpn-services/<service-id>/sessions/<session-id>
	//  GET /policy/api/v1/infra/realized-state/realized-entity?realized_path=/policy/api/v1/infra/tier-1s/<tier-1-id>/ipsec-vpn-services/<service-id>/sessions/<session-id>
	//
	// Deprecated: This API element is deprecated.
	//
	// @param ipsecVpnSessionIdParam (required)
	// @param barrierIdParam (optional)
	// @param requestIdParam Realization request ID (optional)
	// @return com.vmware.nsx.model.IPSecVPNSessionState
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(ipsecVpnSessionIdParam string, barrierIdParam *int64, requestIdParam *string) (nsxModel.IPSecVPNSessionState, error)
}

type stateClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewStateClient(connector vapiProtocolClient_.Connector) *stateClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx.vpn.ipsec.sessions.state")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"get": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	sIface := stateClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &sIface
}

func (sIface *stateClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := sIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (sIface *stateClient) Get(ipsecVpnSessionIdParam string, barrierIdParam *int64, requestIdParam *string) (nsxModel.IPSecVPNSessionState, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	operationRestMetaData := stateGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(stateGetInputType(), typeConverter)
	sv.AddStructField("IpsecVpnSessionId", ipsecVpnSessionIdParam)
	sv.AddStructField("BarrierId", barrierIdParam)
	sv.AddStructField("RequestId", requestIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.IPSecVPNSessionState
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx.vpn.ipsec.sessions.state", "get", inputDataValue, executionContext)
	var emptyOutput nsxModel.IPSecVPNSessionState
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), StateGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.IPSecVPNSessionState), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
