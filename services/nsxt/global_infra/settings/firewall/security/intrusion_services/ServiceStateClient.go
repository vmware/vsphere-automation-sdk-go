// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: ServiceState
// Used by client-side stubs.

package intrusion_services

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/core"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/lib"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/protocol/client"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt/model"
)

const _ = core.SupportedByRuntimeVersion1

type ServiceStateClient interface {

	// Get the indexing progress of IDS Events like start time, current position and total records.
	// @return com.vmware.nsx_policy.model.PolicyIdsServiceState
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get() (model.PolicyIdsServiceState, error)
}

type serviceStateClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewServiceStateClient(connector client.Connector) *serviceStateClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.nsx_policy.global_infra.settings.firewall.security.intrusion_services.service_state")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"get": core.NewMethodIdentifier(interfaceIdentifier, "get"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	sIface := serviceStateClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &sIface
}

func (sIface *serviceStateClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := sIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (sIface *serviceStateClient) Get() (model.PolicyIdsServiceState, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(serviceStateGetInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.PolicyIdsServiceState
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := serviceStateGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.global_infra.settings.firewall.security.intrusion_services.service_state", "get", inputDataValue, executionContext)
	var emptyOutput model.PolicyIdsServiceState
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), serviceStateGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.PolicyIdsServiceState), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
