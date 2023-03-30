// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: StorageSpec
// Used by client-side stubs.

package esxs

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	vmcModel "github.com/vmware/vsphere-automation-sdk-go/services/vmc/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type StorageSpecClient interface {

	// Get storage specs for a host.
	//
	// @param orgParam Organization identifier (required)
	// @param sddcParam Sddc identifier (required)
	// @param clusterParam cluster identifier (required)
	// @param esxParam esx identifier (required)
	//
	// @throws Unauthenticated  Unauthorized
	// @throws InvalidRequest  Invalid or missing parameters
	// @throws Unauthorized  Forbidden
	GetStorageSpecs(orgParam string, sddcParam string, clusterParam string, esxParam string) ([]vmcModel.VsanDiskgroupMapping, error)
}

type storageSpecClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewStorageSpecClient(connector vapiProtocolClient_.Connector) *storageSpecClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.vmc.orgs.sddcs.clusters.esxs.storage_spec")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"get_storage_specs": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get_storage_specs"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	sIface := storageSpecClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &sIface
}

func (sIface *storageSpecClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := sIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (sIface *storageSpecClient) GetStorageSpecs(orgParam string, sddcParam string, clusterParam string, esxParam string) ([]vmcModel.VsanDiskgroupMapping, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	operationRestMetaData := storageSpecGetStorageSpecsRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(storageSpecGetStorageSpecsInputType(), typeConverter)
	sv.AddStructField("Org", orgParam)
	sv.AddStructField("Sddc", sddcParam)
	sv.AddStructField("Cluster", clusterParam)
	sv.AddStructField("Esx", esxParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput []vmcModel.VsanDiskgroupMapping
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.vmc.orgs.sddcs.clusters.esxs.storage_spec", "get_storage_specs", inputDataValue, executionContext)
	var emptyOutput []vmcModel.VsanDiskgroupMapping
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), StorageSpecGetStorageSpecsOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.([]vmcModel.VsanDiskgroupMapping), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
