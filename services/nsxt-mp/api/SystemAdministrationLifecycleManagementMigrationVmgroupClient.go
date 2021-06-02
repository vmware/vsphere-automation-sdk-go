// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: SystemAdministrationLifecycleManagementMigrationVmgroup
// Used by client-side stubs.

package api

import (
	"github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/model"
)

const _ = core.SupportedByRuntimeVersion1

type SystemAdministrationLifecycleManagementMigrationVmgroupClient interface {

	// For each VM group, the following three high level steps are performed in sequence. 1. Call pre VM group migrate API. 2. Migrate (by vmotion,in place, etc.,) VMs in the VM group. This step will be done by user independent of MC. 3. Call post VM group migrate API with the same VM group id used in the pre VM group migrate API. This API specifically deals with post VM group migrate API. When post VM group migrate API is invoked for a VM group id, MC performs following actions. - Add security tags on the VMs migrated. For the VMs mentioned in the failed VM instance uuids, this operation is skipped.
	//
	// @param postVmGroupMigrationSpecParam (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	PostVmGroupMigratePostMigrate(postVmGroupMigrationSpecParam model.PostVmGroupMigrationSpec) error

	// For each VM group, the following three high level steps are performed in sequence. 1. Call pre VM group migrate API. 2. Migrate (by vmotion,in place, etc.,) VMs in the VM group. This step will be done by user independent of MC. 3. Call post VM group migrate API with the same VM group id used in the pre VM group migrate API. This API specifically deals with pre VM group migrate API. When pre VM group migrate API is invoked for a VM group id, MC performs following actions. - Checks segmentation realization state. - Creates segment ports. - Creates temporary security groups.
	//
	// @param preVmGroupMigrationSpecParam (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	PreVmGroupMigratePreMigrate(preVmGroupMigrationSpecParam model.PreVmGroupMigrationSpec) error
}

type systemAdministrationLifecycleManagementMigrationVmgroupClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewSystemAdministrationLifecycleManagementMigrationVmgroupClient(connector client.Connector) *systemAdministrationLifecycleManagementMigrationVmgroupClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.api.system_administration_lifecycle_management_migration_vmgroup")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"post_vm_group_migrate_post_migrate": core.NewMethodIdentifier(interfaceIdentifier, "post_vm_group_migrate_post_migrate"),
		"pre_vm_group_migrate_pre_migrate":   core.NewMethodIdentifier(interfaceIdentifier, "pre_vm_group_migrate_pre_migrate"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	sIface := systemAdministrationLifecycleManagementMigrationVmgroupClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &sIface
}

func (sIface *systemAdministrationLifecycleManagementMigrationVmgroupClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := sIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (sIface *systemAdministrationLifecycleManagementMigrationVmgroupClient) PostVmGroupMigratePostMigrate(postVmGroupMigrationSpecParam model.PostVmGroupMigrationSpec) error {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationLifecycleManagementMigrationVmgroupPostVmGroupMigratePostMigrateInputType(), typeConverter)
	sv.AddStructField("PostVmGroupMigrationSpec", postVmGroupMigrationSpecParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationLifecycleManagementMigrationVmgroupPostVmGroupMigratePostMigrateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_lifecycle_management_migration_vmgroup", "post_vm_group_migrate_post_migrate", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (sIface *systemAdministrationLifecycleManagementMigrationVmgroupClient) PreVmGroupMigratePreMigrate(preVmGroupMigrationSpecParam model.PreVmGroupMigrationSpec) error {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationLifecycleManagementMigrationVmgroupPreVmGroupMigratePreMigrateInputType(), typeConverter)
	sv.AddStructField("PreVmGroupMigrationSpec", preVmGroupMigrationSpecParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationLifecycleManagementMigrationVmgroupPreVmGroupMigratePreMigrateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_lifecycle_management_migration_vmgroup", "pre_vm_group_migrate_pre_migrate", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}
