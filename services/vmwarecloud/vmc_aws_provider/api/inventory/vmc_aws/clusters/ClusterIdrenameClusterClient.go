// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: ClusterIdrenameCluster
// Used by client-side stubs.

package clusters

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	vmwarecloudVmc_aws_providerModel "github.com/vmware/vsphere-automation-sdk-go/services/vmwarecloud/vmc_aws_provider/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type ClusterIdrenameClusterClient interface {

	// Allows users to rename the cluster.
	//
	// @param orgIdParam organization identifier
	// @param clusterIdParam cluster identifier
	// @param clusterRenameIntentParam Payload for renaming the cluster request
	// @return Activity accepted for processing
	//
	// @throws Unauthenticated Unauthorized
	// @throws Unauthorized Forbidden
	RenameCluster(orgIdParam string, clusterIdParam string, clusterRenameIntentParam vmwarecloudVmc_aws_providerModel.ClusterRenameIntent) (vmwarecloudVmc_aws_providerModel.ActivityResponse, error)
}

type clusterIdrenameClusterClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewClusterIdrenameClusterClient(connector vapiProtocolClient_.Connector) *clusterIdrenameClusterClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.vmwarecloud.vmc_aws_provider.api.inventory.vmc_aws.clusters.cluster_idrename_cluster")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"rename_cluster": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "rename_cluster"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	cIface := clusterIdrenameClusterClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &cIface
}

func (cIface *clusterIdrenameClusterClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := cIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (cIface *clusterIdrenameClusterClient) RenameCluster(orgIdParam string, clusterIdParam string, clusterRenameIntentParam vmwarecloudVmc_aws_providerModel.ClusterRenameIntent) (vmwarecloudVmc_aws_providerModel.ActivityResponse, error) {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	operationRestMetaData := clusterIdrenameClusterRenameClusterRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(clusterIdrenameClusterRenameClusterInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("ClusterId", clusterIdParam)
	sv.AddStructField("ClusterRenameIntent", clusterRenameIntentParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput vmwarecloudVmc_aws_providerModel.ActivityResponse
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.vmwarecloud.vmc_aws_provider.api.inventory.vmc_aws.clusters.cluster_idrename_cluster", "rename_cluster", inputDataValue, executionContext)
	var emptyOutput vmwarecloudVmc_aws_providerModel.ActivityResponse
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), ClusterIdrenameClusterRenameClusterOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(vmwarecloudVmc_aws_providerModel.ActivityResponse), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
