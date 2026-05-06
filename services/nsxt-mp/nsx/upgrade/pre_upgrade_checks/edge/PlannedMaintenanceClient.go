// Copyright (c) 2019-2026 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: PlannedMaintenance
// Used by client-side stubs.

package edge

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	nsxModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type PlannedMaintenanceClient interface {

	// This API will accept the edge identification information either edge uuid or edge path. Peer(s) of the given edge will be checked to verify their Edge health, LR status, BGP neighbor status. Healthy peer edge(s) ensure minimal downtime when target edge enters maintenance mode. The progress status OK, if checks have passed. Otherwise API will return NOT_OK. If status returns WAIT, In that case system expects client to wait and poll the API after every 1 minute till checks reach terminal state OK or NOT_OK within stipulated internal retry intervals.
	//
	// @param preUpgradeEdgeHealthCheckRequestParam (required)
	// @return com.vmware.nsx.model.PreUpgradeHostHealthCheckStatuses
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Create(preUpgradeEdgeHealthCheckRequestParam nsxModel.PreUpgradeEdgeHealthCheckRequest) (nsxModel.PreUpgradeHostHealthCheckStatuses, error)
}

type plannedMaintenanceClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewPlannedMaintenanceClient(connector vapiProtocolClient_.Connector) *plannedMaintenanceClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx.upgrade.pre_upgrade_checks.edge.planned_maintenance")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"create": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "create"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	pIface := plannedMaintenanceClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &pIface
}

func (pIface *plannedMaintenanceClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := pIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (pIface *plannedMaintenanceClient) Create(preUpgradeEdgeHealthCheckRequestParam nsxModel.PreUpgradeEdgeHealthCheckRequest) (nsxModel.PreUpgradeHostHealthCheckStatuses, error) {
	typeConverter := pIface.connector.TypeConverter()
	executionContext := pIface.connector.NewExecutionContext()
	operationRestMetaData := plannedMaintenanceCreateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(plannedMaintenanceCreateInputType(), typeConverter)
	sv.AddStructField("PreUpgradeEdgeHealthCheckRequest", preUpgradeEdgeHealthCheckRequestParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.PreUpgradeHostHealthCheckStatuses
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := pIface.connector.GetApiProvider().Invoke("com.vmware.nsx.upgrade.pre_upgrade_checks.edge.planned_maintenance", "create", inputDataValue, executionContext)
	var emptyOutput nsxModel.PreUpgradeHostHealthCheckStatuses
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), PlannedMaintenanceCreateOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.PreUpgradeHostHealthCheckStatuses), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), pIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
