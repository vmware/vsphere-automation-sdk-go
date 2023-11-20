// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Setup
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

type SetupClient interface {

	// Add ALB endpoint details for non cross VC migration modes.
	//
	// @param aviEndPointParam (required)
	//
	// @throws Unauthenticated  Unauthorized
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Addalbinfo(aviEndPointParam nsxModel.AviEndPoint) error

	// Add NSX-V to NSX-T site mapping.
	//
	// @param v2tSiteMappingSpecParam (required)
	//
	// @throws Unauthenticated  Unauthorized
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Addv2tsitemapping(v2tSiteMappingSpecParam nsxModel.V2tSiteMappingSpec) error

	// Get setup details of NSX-V to be migrated.
	// @return com.vmware.nsx.model.MigrationSetupInfo
	//
	// @throws Unauthenticated  Unauthorized
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get() (nsxModel.MigrationSetupInfo, error)

	// This API is to support the add Host workflow on NSX-T once the V2T migration process has started. There are two high level cases. For each case, please refer to the detailed steps. Case - 1 : For migration mode - CONFIG_AND_EDGE_MIGRATION_WITH_BYOT, CONFIG_AND_EDGE_MIGRATION_WITH_BYOT_ON_FEDERATION 1. Verify that overall migration status is in SUCCESS state in migration co-ordinator. This means all migration stages are completed and we are at a point where we can move workloads using the pre-migrate, post-migrate API's. 2. Follow steps mentioned in product documentation to add a new host to cluster. 3. Invoke this API (Accept Host Transport Node) to perform the V2T migration steps on newly added host transport node. Once API is successful, make sure that newly added HOST can take worklods (Ex: Traffic checks, etc.,). For Case-1, this API should be invoked only when following conditions are met. - The overall migration status is in SUCCESS status. - All the migration stages are completed and we are at a point where we can move workloads. - The migration mode selected has to be among the ones designated for case-1 - The new host has been added to a cluster by following all the steps mentioned in product documentation. Case - 2 : For migration modes that perform HOST migration. 1. Verify that overall migration status is in PAUSED state in migration co-ordinator. 2. Follow steps mentioned in product documentation to add a new host to cluster. 3. Invoke the migration co-ordinator sync HOST groups API. This will result in migration co-ordinator updating its inventory and accouting for the newly added host transport node. At this point, the list host upgrade units API call to MC will also show the newly added host transport node. But do NOT resume the migration. Since this host transport node is added after the V2T migration has started, we need to perform migration steps on this host through this special API (Accept Host Transport Node) 4. Invoke this API (Accept Host Transport Node) to perform the V2T migration steps on newly added host transport node. Once the API is successful, we should be seeing that migration status has been marked as SUCCESS for the newly added host transport node. Very Importnant Note : Make sure that migration status is PAUSED in migration co-ordinator when performing steps 2 through 4. Also ensure that all the steps 2 through 4 pass without failures. For Case-2, this API should be invoked only when following conditions are met. - The overall migration status is in PAUSED status. - The migration mode selected has HOST component in it. - All the migration stages before HOST stage are completed. - The new host has been added to a cluster by following all the steps mentioned in product documentation.
	//
	// @param newHostTransportNodeSpecParam (required)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Migratenewlyaddedhosttransportnode(newHostTransportNodeSpecParam nsxModel.NewHostTransportNodeSpec) error

	// Set the NSX-V ESG to NSX-T Router mapping option.
	//
	// @param mappingOptionParam Mapping option (required)
	//
	// @throws Unauthenticated  Unauthorized
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Setesgtoroutermappingoption(mappingOptionParam string) error

	// Provide setup details of NSX-V to be migrated.
	//
	// @param migrationSetupInfoParam (required)
	// @return com.vmware.nsx.model.MigrationSetupInfo
	//
	// @throws Unauthenticated  Unauthorized
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(migrationSetupInfoParam nsxModel.MigrationSetupInfo) (nsxModel.MigrationSetupInfo, error)
}

type setupClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewSetupClient(connector vapiProtocolClient_.Connector) *setupClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx.migration.setup")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"addalbinfo":                         vapiCore_.NewMethodIdentifier(interfaceIdentifier, "addalbinfo"),
		"addv2tsitemapping":                  vapiCore_.NewMethodIdentifier(interfaceIdentifier, "addv2tsitemapping"),
		"get":                                vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
		"migratenewlyaddedhosttransportnode": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "migratenewlyaddedhosttransportnode"),
		"setesgtoroutermappingoption":        vapiCore_.NewMethodIdentifier(interfaceIdentifier, "setesgtoroutermappingoption"),
		"update":                             vapiCore_.NewMethodIdentifier(interfaceIdentifier, "update"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	sIface := setupClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &sIface
}

func (sIface *setupClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := sIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (sIface *setupClient) Addalbinfo(aviEndPointParam nsxModel.AviEndPoint) error {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	operationRestMetaData := setupAddalbinfoRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(setupAddalbinfoInputType(), typeConverter)
	sv.AddStructField("AviEndPoint", aviEndPointParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx.migration.setup", "addalbinfo", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (sIface *setupClient) Addv2tsitemapping(v2tSiteMappingSpecParam nsxModel.V2tSiteMappingSpec) error {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	operationRestMetaData := setupAddv2tsitemappingRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(setupAddv2tsitemappingInputType(), typeConverter)
	sv.AddStructField("V2tSiteMappingSpec", v2tSiteMappingSpecParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx.migration.setup", "addv2tsitemapping", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (sIface *setupClient) Get() (nsxModel.MigrationSetupInfo, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	operationRestMetaData := setupGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(setupGetInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.MigrationSetupInfo
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx.migration.setup", "get", inputDataValue, executionContext)
	var emptyOutput nsxModel.MigrationSetupInfo
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), SetupGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.MigrationSetupInfo), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *setupClient) Migratenewlyaddedhosttransportnode(newHostTransportNodeSpecParam nsxModel.NewHostTransportNodeSpec) error {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	operationRestMetaData := setupMigratenewlyaddedhosttransportnodeRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(setupMigratenewlyaddedhosttransportnodeInputType(), typeConverter)
	sv.AddStructField("NewHostTransportNodeSpec", newHostTransportNodeSpecParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx.migration.setup", "migratenewlyaddedhosttransportnode", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (sIface *setupClient) Setesgtoroutermappingoption(mappingOptionParam string) error {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	operationRestMetaData := setupSetesgtoroutermappingoptionRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(setupSetesgtoroutermappingoptionInputType(), typeConverter)
	sv.AddStructField("MappingOption", mappingOptionParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx.migration.setup", "setesgtoroutermappingoption", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (sIface *setupClient) Update(migrationSetupInfoParam nsxModel.MigrationSetupInfo) (nsxModel.MigrationSetupInfo, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	operationRestMetaData := setupUpdateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(setupUpdateInputType(), typeConverter)
	sv.AddStructField("MigrationSetupInfo", migrationSetupInfoParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.MigrationSetupInfo
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx.migration.setup", "update", inputDataValue, executionContext)
	var emptyOutput nsxModel.MigrationSetupInfo
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), SetupUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.MigrationSetupInfo), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
