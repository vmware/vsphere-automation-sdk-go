// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: SystemAdministrationLifecycleManagementMigrationSetup
// Used by client-side stubs.

package api

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/core"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/lib"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/protocol/client"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt-mp/model"
)

const _ = core.SupportedByRuntimeVersion1

type SystemAdministrationLifecycleManagementMigrationSetupClient interface {

	// Get setup details of NSX-V to be migrated.
	// @return com.vmware.model.MigrationSetupInfo
	// @throws Unauthenticated  Unauthorized
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetNsxvSetupDetails() (model.MigrationSetupInfo, error)

	// Provide setup details of NSX-V to be migrated.
	//
	// @param migrationSetupInfoParam (required)
	// @return com.vmware.model.MigrationSetupInfo
	// @throws Unauthenticated  Unauthorized
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	UpdateNsxvSetupDetails(migrationSetupInfoParam model.MigrationSetupInfo) (model.MigrationSetupInfo, error)
}

type systemAdministrationLifecycleManagementMigrationSetupClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewSystemAdministrationLifecycleManagementMigrationSetupClient(connector client.Connector) *systemAdministrationLifecycleManagementMigrationSetupClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.api.system_administration_lifecycle_management_migration_setup")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"get_nsxv_setup_details":    core.NewMethodIdentifier(interfaceIdentifier, "get_nsxv_setup_details"),
		"update_nsxv_setup_details": core.NewMethodIdentifier(interfaceIdentifier, "update_nsxv_setup_details"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	sIface := systemAdministrationLifecycleManagementMigrationSetupClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &sIface
}

func (sIface *systemAdministrationLifecycleManagementMigrationSetupClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := sIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (sIface *systemAdministrationLifecycleManagementMigrationSetupClient) GetNsxvSetupDetails() (model.MigrationSetupInfo, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationLifecycleManagementMigrationSetupGetNsxvSetupDetailsInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.MigrationSetupInfo
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationLifecycleManagementMigrationSetupGetNsxvSetupDetailsRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_lifecycle_management_migration_setup", "get_nsxv_setup_details", inputDataValue, executionContext)
	var emptyOutput model.MigrationSetupInfo
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationLifecycleManagementMigrationSetupGetNsxvSetupDetailsOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.MigrationSetupInfo), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationLifecycleManagementMigrationSetupClient) UpdateNsxvSetupDetails(migrationSetupInfoParam model.MigrationSetupInfo) (model.MigrationSetupInfo, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationLifecycleManagementMigrationSetupUpdateNsxvSetupDetailsInputType(), typeConverter)
	sv.AddStructField("MigrationSetupInfo", migrationSetupInfoParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.MigrationSetupInfo
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationLifecycleManagementMigrationSetupUpdateNsxvSetupDetailsRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_lifecycle_management_migration_setup", "update_nsxv_setup_details", inputDataValue, executionContext)
	var emptyOutput model.MigrationSetupInfo
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationLifecycleManagementMigrationSetupUpdateNsxvSetupDetailsOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.MigrationSetupInfo), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
