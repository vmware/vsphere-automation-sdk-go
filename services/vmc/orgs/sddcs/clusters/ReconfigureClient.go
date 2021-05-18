// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Reconfigure
// Used by client-side stubs.

package clusters

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/core"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/lib"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/protocol/client"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/vmc/model"
)

const _ = core.SupportedByRuntimeVersion1

type ReconfigureClient interface {

	// This reconfiguration will handle changing both the number of hosts and the cluster storage capacity.
	//
	// @param orgParam Organization identifier (required)
	// @param sddcParam Sddc identifier (required)
	// @param clusterParam cluster identifier (required)
	// @param clusterReconfigureParamsParam clusterReconfigureParams (required)
	// @return com.vmware.vmc.model.Task
	// @throws Unauthenticated  Unauthorized
	// @throws InvalidRequest  Bad Request. The SDDC is not in a valid state. , Method not allowed
	// @throws Unauthorized  Access not allowed to the operation for the current user
	// @throws NotFound  Cannot find the cluster with the given id
	ClusterReconfig(orgParam string, sddcParam string, clusterParam string, clusterReconfigureParamsParam model.ClusterReconfigureParams) (model.Task, error)
}

type reconfigureClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewReconfigureClient(connector client.Connector) *reconfigureClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.vmc.orgs.sddcs.clusters.reconfigure")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"cluster_reconfig": core.NewMethodIdentifier(interfaceIdentifier, "cluster_reconfig"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	rIface := reconfigureClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &rIface
}

func (rIface *reconfigureClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := rIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (rIface *reconfigureClient) ClusterReconfig(orgParam string, sddcParam string, clusterParam string, clusterReconfigureParamsParam model.ClusterReconfigureParams) (model.Task, error) {
	typeConverter := rIface.connector.TypeConverter()
	executionContext := rIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(reconfigureClusterReconfigInputType(), typeConverter)
	sv.AddStructField("Org", orgParam)
	sv.AddStructField("Sddc", sddcParam)
	sv.AddStructField("Cluster", clusterParam)
	sv.AddStructField("ClusterReconfigureParams", clusterReconfigureParamsParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.Task
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := reconfigureClusterReconfigRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	rIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := rIface.connector.GetApiProvider().Invoke("com.vmware.vmc.orgs.sddcs.clusters.reconfigure", "cluster_reconfig", inputDataValue, executionContext)
	var emptyOutput model.Task
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), reconfigureClusterReconfigOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.Task), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), rIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
