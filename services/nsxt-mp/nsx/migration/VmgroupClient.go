// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Vmgroup
// Used by client-side stubs.

package migration

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	nsxModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type VmgroupClient interface {

	// For each VM group, the following three high level steps are performed in sequence. 1. Call pre VM group migrate API. 2. Migrate (by vmotion,in place, etc.,) VMs in the VM group. This step will be done by user independent of MC. 3. Call post VM group migrate API with the same VM group id used in the pre VM group migrate API. This API specifically deals with post VM group migrate API. When post VM group migrate API is invoked for a VM group id, MC performs following actions. - Add security tags on the VMs migrated. For the VMs mentioned in the failed VM instance uuids, this operation is skipped.
	//
	// @param postVmGroupMigrationSpecParam (required)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Postmigrate(postVmGroupMigrationSpecParam nsxModel.PostVmGroupMigrationSpec) error

	// For each VM group, the following three high level steps are performed in sequence. 1. Call pre VM group migrate API. 2. Migrate (by vmotion,in place, etc.,) VMs in the VM group. This step will be done by user independent of MC. 3. Call post VM group migrate API with the same VM group id used in the pre VM group migrate API. This API specifically deals with pre VM group migrate API. When pre VM group migrate API is invoked for a VM group id, MC performs following actions. - Checks segmentation realization state. - Creates segment ports. - Creates temporary security groups.
	//
	// @param preVmGroupMigrationSpecParam (required)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Premigrate(preVmGroupMigrationSpecParam nsxModel.PreVmGroupMigrationSpec) error
}

type vmgroupClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewVmgroupClient(connector vapiProtocolClient_.Connector) *vmgroupClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx.migration.vmgroup")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"postmigrate": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "postmigrate"),
		"premigrate":  vapiCore_.NewMethodIdentifier(interfaceIdentifier, "premigrate"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	vIface := vmgroupClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &vIface
}

func (vIface *vmgroupClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := vIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (vIface *vmgroupClient) Postmigrate(postVmGroupMigrationSpecParam nsxModel.PostVmGroupMigrationSpec) error {
	typeConverter := vIface.connector.TypeConverter()
	executionContext := vIface.connector.NewExecutionContext()
	operationRestMetaData := vmgroupPostmigrateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(vmgroupPostmigrateInputType(), typeConverter)
	sv.AddStructField("PostVmGroupMigrationSpec", postVmGroupMigrationSpecParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := vIface.connector.GetApiProvider().Invoke("com.vmware.nsx.migration.vmgroup", "postmigrate", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), vIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (vIface *vmgroupClient) Premigrate(preVmGroupMigrationSpecParam nsxModel.PreVmGroupMigrationSpec) error {
	typeConverter := vIface.connector.TypeConverter()
	executionContext := vIface.connector.NewExecutionContext()
	operationRestMetaData := vmgroupPremigrateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(vmgroupPremigrateInputType(), typeConverter)
	sv.AddStructField("PreVmGroupMigrationSpec", preVmGroupMigrationSpecParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := vIface.connector.GetApiProvider().Invoke("com.vmware.nsx.migration.vmgroup", "premigrate", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), vIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}
