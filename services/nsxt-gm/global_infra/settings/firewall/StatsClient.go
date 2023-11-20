// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Stats
// Used by client-side stubs.

package firewall

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type StatsClient interface {

	// Sets firewall rule statistics counter to zero. This operation is supported for given category, for example: DFW i.e. for all layer3 firewall (transport nodes only) rules or EDGE i.e. for all layer3 edge firewall (edge nodes only) rules. - no enforcement point path specified: On global manager, it is mandatory to give an enforcement point path. On local manager, reset of stats will be executed for each enforcement point. - {enforcement_point_path}: Reset of stats will be executed only for the given enforcement point.
	//
	// @param categoryParam Aggregation statistic category (required)
	// @param containerClusterPathParam String Path of the Container Cluster entity (optional)
	// @param enforcementPointPathParam String Path of the enforcement point (optional)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Reset(categoryParam string, containerClusterPathParam *string, enforcementPointPathParam *string) error
}

type statsClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewStatsClient(connector vapiProtocolClient_.Connector) *statsClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx_global_policy.global_infra.settings.firewall.stats")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"reset": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "reset"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	sIface := statsClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &sIface
}

func (sIface *statsClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := sIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (sIface *statsClient) Reset(categoryParam string, containerClusterPathParam *string, enforcementPointPathParam *string) error {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	operationRestMetaData := statsResetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(statsResetInputType(), typeConverter)
	sv.AddStructField("Category", categoryParam)
	sv.AddStructField("ContainerClusterPath", containerClusterPathParam)
	sv.AddStructField("EnforcementPointPath", enforcementPointPathParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx_global_policy.global_infra.settings.firewall.stats", "reset", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}
