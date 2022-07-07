// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: MigrationUnit
// Used by client-side stubs.

package migration_unit_groups

import (
	"github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
)

const _ = core.SupportedByRuntimeVersion1

type MigrationUnitClient interface {

	// Reorder an migration unit within the migration unit group by placing it before/after the specified migration unit
	//
	// @param groupIdParam (required)
	// @param migrationUnitIdParam (required)
	// @param reorderMigrationRequestParam (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Reorder(groupIdParam string, migrationUnitIdParam string, reorderMigrationRequestParam model.ReorderMigrationRequest) error
}

type migrationUnitClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewMigrationUnitClient(connector client.Connector) *migrationUnitClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.nsx.migration.migration_unit_groups.migration_unit")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"reorder": core.NewMethodIdentifier(interfaceIdentifier, "reorder"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	mIface := migrationUnitClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &mIface
}

func (mIface *migrationUnitClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := mIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (mIface *migrationUnitClient) Reorder(groupIdParam string, migrationUnitIdParam string, reorderMigrationRequestParam model.ReorderMigrationRequest) error {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(migrationUnitReorderInputType(), typeConverter)
	sv.AddStructField("GroupId", groupIdParam)
	sv.AddStructField("MigrationUnitId", migrationUnitIdParam)
	sv.AddStructField("ReorderMigrationRequest", reorderMigrationRequestParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := migrationUnitReorderRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.nsx.migration.migration_unit_groups.migration_unit", "reorder", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}
