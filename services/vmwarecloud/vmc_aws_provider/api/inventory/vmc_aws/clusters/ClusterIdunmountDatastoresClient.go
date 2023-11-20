// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: ClusterIdunmountDatastores
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

type ClusterIdunmountDatastoresClient interface {

	// Unmount nfs datastore to a cluster
	//
	// @param orgIdParam organization identifier
	// @param clusterIdParam cluster identifier
	// @param unmountNfsDatastoreIntentParam Payload for the Unmount NFS datastores
	// @return Activity accepted for processing
	//
	// @throws InvalidRequest Bad Request
	// @throws Unauthenticated Unauthorized
	// @throws Unauthorized Forbidden
	// @throws NotFound Not Found
	UnmountDataStoresOnCluster(orgIdParam string, clusterIdParam string, unmountNfsDatastoreIntentParam vmwarecloudVmc_aws_providerModel.UnmountNfsDatastoreIntent) (vmwarecloudVmc_aws_providerModel.ActivityResponse, error)
}

type clusterIdunmountDatastoresClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewClusterIdunmountDatastoresClient(connector vapiProtocolClient_.Connector) *clusterIdunmountDatastoresClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.vmwarecloud.vmc_aws_provider.api.inventory.vmc_aws.clusters.cluster_idunmount_datastores")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"unmount_data_stores_on_cluster": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "unmount_data_stores_on_cluster"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	cIface := clusterIdunmountDatastoresClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &cIface
}

func (cIface *clusterIdunmountDatastoresClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := cIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (cIface *clusterIdunmountDatastoresClient) UnmountDataStoresOnCluster(orgIdParam string, clusterIdParam string, unmountNfsDatastoreIntentParam vmwarecloudVmc_aws_providerModel.UnmountNfsDatastoreIntent) (vmwarecloudVmc_aws_providerModel.ActivityResponse, error) {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	operationRestMetaData := clusterIdunmountDatastoresUnmountDataStoresOnClusterRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(clusterIdunmountDatastoresUnmountDataStoresOnClusterInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("ClusterId", clusterIdParam)
	sv.AddStructField("UnmountNfsDatastoreIntent", unmountNfsDatastoreIntentParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput vmwarecloudVmc_aws_providerModel.ActivityResponse
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.vmwarecloud.vmc_aws_provider.api.inventory.vmc_aws.clusters.cluster_idunmount_datastores", "unmount_data_stores_on_cluster", inputDataValue, executionContext)
	var emptyOutput vmwarecloudVmc_aws_providerModel.ActivityResponse
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), ClusterIdunmountDatastoresUnmountDataStoresOnClusterOutputType())
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
