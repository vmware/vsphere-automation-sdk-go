// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Vmgroup
// Used by client-side stubs.

package migration

import (
	"github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
)

const _ = core.SupportedByRuntimeVersion1

type VmgroupClient interface {

	// For each VM group, the following three high level steps are performed in sequence. 1. Call pre VM group migrate API. 2. Migrate (by vmotion,in place, etc.,) VMs in the VM group. This step will be done by user independent of MC. 3. Call post VM group migrate API with the same VM group id used in the pre VM group migrate API. This API specifically deals with post VM group migrate API. When post VM group migrate API is invoked for a VM group id, MC performs following actions. - Add security tags on the VMs migrated. For the VMs mentioned in the failed VM instance uuids, this operation is skipped.
	//
	// @param postVmGroupMigrationSpecParam (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Postmigrate(postVmGroupMigrationSpecParam model.PostVmGroupMigrationSpec) error

	// For each VM group, the following three high level steps are performed in sequence. 1. Call pre VM group migrate API. 2. Migrate (by vmotion,in place, etc.,) VMs in the VM group. This step will be done by user independent of MC. 3. Call post VM group migrate API with the same VM group id used in the pre VM group migrate API. This API specifically deals with pre VM group migrate API. When pre VM group migrate API is invoked for a VM group id, MC performs following actions. - Checks segmentation realization state. - Creates segment ports. - Creates temporary security groups.
	//
	// @param preVmGroupMigrationSpecParam (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Premigrate(preVmGroupMigrationSpecParam model.PreVmGroupMigrationSpec) error
}

type vmgroupClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewVmgroupClient(connector client.Connector) *vmgroupClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.nsx.migration.vmgroup")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"postmigrate": core.NewMethodIdentifier(interfaceIdentifier, "postmigrate"),
		"premigrate":  core.NewMethodIdentifier(interfaceIdentifier, "premigrate"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	vIface := vmgroupClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &vIface
}

func (vIface *vmgroupClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := vIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (vIface *vmgroupClient) Postmigrate(postVmGroupMigrationSpecParam model.PostVmGroupMigrationSpec) error {
	typeConverter := vIface.connector.TypeConverter()
	executionContext := vIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(vmgroupPostmigrateInputType(), typeConverter)
	sv.AddStructField("PostVmGroupMigrationSpec", postVmGroupMigrationSpecParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := vmgroupPostmigrateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	vIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := vIface.connector.GetApiProvider().Invoke("com.vmware.nsx.migration.vmgroup", "postmigrate", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), vIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (vIface *vmgroupClient) Premigrate(preVmGroupMigrationSpecParam model.PreVmGroupMigrationSpec) error {
	typeConverter := vIface.connector.TypeConverter()
	executionContext := vIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(vmgroupPremigrateInputType(), typeConverter)
	sv.AddStructField("PreVmGroupMigrationSpec", preVmGroupMigrationSpecParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := vmgroupPremigrateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	vIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := vIface.connector.GetApiProvider().Invoke("com.vmware.nsx.migration.vmgroup", "premigrate", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), vIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}
