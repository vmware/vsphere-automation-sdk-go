// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: RemediationPolicy
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

type RemediationPolicyClient interface {

	// Get the current Remediation policy on a cluster.
	//
	// @param orgParam org identifier (required)
	// @param sddcParam Sddc identifier (required)
	// @param clusterParam cluster identifier (required)
	// @return com.vmware.model.RemediationClusterInfo
	//
	// @throws Unauthenticated  Unauthorized
	// @throws InvalidRequest  Invalid action or bad argument
	// @throws Unauthorized  Forbidden
	Get(orgParam string, sddcParam string, clusterParam string) (model.RemediationClusterInfo, error)

	// Change the current Remediation policy on a cluster.
	//
	// @param orgParam org identifier (required)
	// @param sddcParam Sddc identifier (required)
	// @param clusterParam cluster identifier (required)
	// @param remediationPolicyPatchParam Remediation policy patch (required)
	// @return com.vmware.model.Task
	//
	// @throws Unauthenticated  Unauthorized
	// @throws InvalidRequest  Invalid action or bad argument
	// @throws Unauthorized  Forbidden
	Patch(orgParam string, sddcParam string, clusterParam string, remediationPolicyPatchParam model.RemediationPolicyPatch) (model.Task, error)
}

type remediationPolicyClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewRemediationPolicyClient(connector vapiProtocolClient_.Connector) *remediationPolicyClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.api.orgs.sddcs.clusters.remediation_policy")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"get":   vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
		"patch": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "patch"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	rIface := remediationPolicyClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &rIface
}

func (rIface *remediationPolicyClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := rIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (rIface *remediationPolicyClient) Get(orgParam string, sddcParam string, clusterParam string) (model.RemediationClusterInfo, error) {
	typeConverter := rIface.connector.TypeConverter()
	executionContext := rIface.connector.NewExecutionContext()
	operationRestMetaData := remediationPolicyGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(remediationPolicyGetInputType(), typeConverter)
	sv.AddStructField("Org", orgParam)
	sv.AddStructField("Sddc", sddcParam)
	sv.AddStructField("Cluster", clusterParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.RemediationClusterInfo
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := rIface.connector.GetApiProvider().Invoke("com.vmware.api.orgs.sddcs.clusters.remediation_policy", "get", inputDataValue, executionContext)
	var emptyOutput model.RemediationClusterInfo
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), RemediationPolicyGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.RemediationClusterInfo), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), rIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (rIface *remediationPolicyClient) Patch(orgParam string, sddcParam string, clusterParam string, remediationPolicyPatchParam model.RemediationPolicyPatch) (model.Task, error) {
	typeConverter := rIface.connector.TypeConverter()
	executionContext := rIface.connector.NewExecutionContext()
	operationRestMetaData := remediationPolicyPatchRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(remediationPolicyPatchInputType(), typeConverter)
	sv.AddStructField("Org", orgParam)
	sv.AddStructField("Sddc", sddcParam)
	sv.AddStructField("Cluster", clusterParam)
	sv.AddStructField("RemediationPolicyPatch", remediationPolicyPatchParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.Task
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := rIface.connector.GetApiProvider().Invoke("com.vmware.api.orgs.sddcs.clusters.remediation_policy", "patch", inputDataValue, executionContext)
	var emptyOutput model.Task
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), RemediationPolicyPatchOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.Task), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), rIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
