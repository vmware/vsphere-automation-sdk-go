// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: EdrsProvisioningSpec
// Used by client-side stubs.

package clusters

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	"github.com/vmware/vsphere-automation-sdk-go/services/vmc/autoscaler/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type EdrsProvisioningSpecClient interface {

	// Get the current EDRS provisioning spec on a cluster.
	//
	// @param orgParam org identifier (required)
	// @param sddcParam Sddc identifier (required)
	// @param clusterParam cluster identifier (required)
	// @return com.vmware.model.EdrsProvisioningSpec
	//
	// @throws Unauthenticated  Unauthorized
	// @throws InvalidRequest  Invalid action or bad argument
	// @throws Unauthorized  Forbidden
	Get(orgParam string, sddcParam string, clusterParam string) (model.EdrsProvisioningSpec, error)
}

type edrsProvisioningSpecClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewEdrsProvisioningSpecClient(connector vapiProtocolClient_.Connector) *edrsProvisioningSpecClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.api.orgs.sddcs.clusters.edrs_provisioning_spec")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"get": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	eIface := edrsProvisioningSpecClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &eIface
}

func (eIface *edrsProvisioningSpecClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := eIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (eIface *edrsProvisioningSpecClient) Get(orgParam string, sddcParam string, clusterParam string) (model.EdrsProvisioningSpec, error) {
	typeConverter := eIface.connector.TypeConverter()
	executionContext := eIface.connector.NewExecutionContext()
	operationRestMetaData := edrsProvisioningSpecGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(edrsProvisioningSpecGetInputType(), typeConverter)
	sv.AddStructField("Org", orgParam)
	sv.AddStructField("Sddc", sddcParam)
	sv.AddStructField("Cluster", clusterParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.EdrsProvisioningSpec
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := eIface.connector.GetApiProvider().Invoke("com.vmware.api.orgs.sddcs.clusters.edrs_provisioning_spec", "get", inputDataValue, executionContext)
	var emptyOutput model.EdrsProvisioningSpec
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), EdrsProvisioningSpecGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.EdrsProvisioningSpec), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), eIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
