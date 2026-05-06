// Copyright (c) 2019-2026 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Upgrade
// Used by client-side stubs.

package nsx

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type UpgradeClient interface {

	// Aborts execution of pre-upgrade checks if already in progress. Halts the execution of checks awaiting execution at this point and makes best-effort attempts to stop checks already in execution. Returns without action if execution of pre-upgrade checks is not in progress.
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Abortpreupgradechecks() error

	// Run pre-defined checks to identify issues after upgrade of a component. The results of the checks are added to the respective upgrade units aggregate-info. The progress and status of post-upgrade checks is part of aggregate-info of individual upgrade unit groups. Returns HTTP status 500 with error code 30953 if execution of post-upgrade checks is already in progress.
	//
	// @param componentTypeParam (required)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Executepostupgradechecks(componentTypeParam string) error

	// Run pre-defined checks to identify potential issues which can be encountered during an upgrade or can cause an upgrade to fail. The results of the checks are added to the respective upgrade units aggregate-info. The progress and status of operation is part of upgrade status summary of individual components. Returns HTTP status 500 with error code 30953 if execution of pre-upgrade checks is already in progress.
	//
	// @param componentTypeParam Component type on which the pre-check will be executed. (optional)
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includedFieldsParam Note - this parameter currently only works when used with the search APIs /policy/api/v1/search/query and /policy/api/v1/search/dsl. It is ignored for other list APIs. (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param runOnlyImplicitPreChecksParam A boolean flag to determine the depth or type of pre-checks to execute. 'implicit' pre-checks are typically foundational or essential checks. - \\\\`true\\\\`: Only 'implicit' pre-checks will be executed for the specified \\\\`component_type\\\\`. This option requires \\\\`component_type\\\\` to be provided. - \\\\`false\\\\`: (Default) All applicable pre-checks for the specified \\\\`component_type\\\\` will be executed. If \\\\`component_type\\\\` is not provided, this flag must be \\\\`false\\\\` (or omitted) Interaction: If \\\\`run_only_on_enabled_groups\\\\` is also \\\\`true\\\\`, then setting this to \\\\`true\\\\` will result in 'implicit' pre-checks being executed specifically on the 'enabled groups' of the given \\\\`component_type\\\\`. (optional, default to false)
	// @param runOnlyOnEnabledGroupsParam A boolean flag to control the scope of group-based pre-checks. - \\\\`true\\\\`: Universal Pre-checks will be executed on groups that are currently marked as 'enabled' for the specified \\\\`component_type\\\\`. This option requires \\\\`component_type\\\\` to be provided. - \\\\`false\\\\`: (Default) For the specified \\\\`component_type\\\\`, universal pre-checks will be executed on every associated group, whether it's enabled or not. If \\\\`component_type\\\\` is not provided, this flag must be \\\\`false\\\\` (or omitted). (optional, default to false)
	// @param sortAscendingParam If true, results are sorted in ascending order (optional)
	// @param sortByParam Field by which records are sorted (optional)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Executepreupgradechecks(componentTypeParam *string, cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, runOnlyImplicitPreChecksParam *bool, runOnlyOnEnabledGroupsParam *bool, sortAscendingParam *bool, sortByParam *string) error

	// This API enables the rollback of the upgrade coordinator to a previous version. It facilitates downgrading the upgrade coordinator on the other two nodes by utilizing the older version (the 'from' version) of the upgrade coordinator war file stored in the repository. Finally, it performs a self-rollback to the older version on itself.
	//
	// @param componentTypeParam Component type (required)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Rollback(componentTypeParam string) error

	// Upgrade the upgrade coordinator module itself. This call is invoked after user uploads an upgrade bundle. Once this call is invoked, upgrade coordinator stops and gets restarted and target version upgrade coordinator module comes up after restart.
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Upgradeuc() error
}

type upgradeClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewUpgradeClient(connector vapiProtocolClient_.Connector) *upgradeClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx.upgrade")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"abortpreupgradechecks":    vapiCore_.NewMethodIdentifier(interfaceIdentifier, "abortpreupgradechecks"),
		"executepostupgradechecks": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "executepostupgradechecks"),
		"executepreupgradechecks":  vapiCore_.NewMethodIdentifier(interfaceIdentifier, "executepreupgradechecks"),
		"rollback":                 vapiCore_.NewMethodIdentifier(interfaceIdentifier, "rollback"),
		"upgradeuc":                vapiCore_.NewMethodIdentifier(interfaceIdentifier, "upgradeuc"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	uIface := upgradeClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &uIface
}

func (uIface *upgradeClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := uIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (uIface *upgradeClient) Abortpreupgradechecks() error {
	typeConverter := uIface.connector.TypeConverter()
	executionContext := uIface.connector.NewExecutionContext()
	operationRestMetaData := upgradeAbortpreupgradechecksRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(upgradeAbortpreupgradechecksInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := uIface.connector.GetApiProvider().Invoke("com.vmware.nsx.upgrade", "abortpreupgradechecks", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), uIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (uIface *upgradeClient) Executepostupgradechecks(componentTypeParam string) error {
	typeConverter := uIface.connector.TypeConverter()
	executionContext := uIface.connector.NewExecutionContext()
	operationRestMetaData := upgradeExecutepostupgradechecksRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(upgradeExecutepostupgradechecksInputType(), typeConverter)
	sv.AddStructField("ComponentType", componentTypeParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := uIface.connector.GetApiProvider().Invoke("com.vmware.nsx.upgrade", "executepostupgradechecks", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), uIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (uIface *upgradeClient) Executepreupgradechecks(componentTypeParam *string, cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, runOnlyImplicitPreChecksParam *bool, runOnlyOnEnabledGroupsParam *bool, sortAscendingParam *bool, sortByParam *string) error {
	typeConverter := uIface.connector.TypeConverter()
	executionContext := uIface.connector.NewExecutionContext()
	operationRestMetaData := upgradeExecutepreupgradechecksRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(upgradeExecutepreupgradechecksInputType(), typeConverter)
	sv.AddStructField("ComponentType", componentTypeParam)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("RunOnlyImplicitPreChecks", runOnlyImplicitPreChecksParam)
	sv.AddStructField("RunOnlyOnEnabledGroups", runOnlyOnEnabledGroupsParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := uIface.connector.GetApiProvider().Invoke("com.vmware.nsx.upgrade", "executepreupgradechecks", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), uIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (uIface *upgradeClient) Rollback(componentTypeParam string) error {
	typeConverter := uIface.connector.TypeConverter()
	executionContext := uIface.connector.NewExecutionContext()
	operationRestMetaData := upgradeRollbackRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(upgradeRollbackInputType(), typeConverter)
	sv.AddStructField("ComponentType", componentTypeParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := uIface.connector.GetApiProvider().Invoke("com.vmware.nsx.upgrade", "rollback", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), uIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (uIface *upgradeClient) Upgradeuc() error {
	typeConverter := uIface.connector.TypeConverter()
	executionContext := uIface.connector.NewExecutionContext()
	operationRestMetaData := upgradeUpgradeucRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(upgradeUpgradeucInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := uIface.connector.GetApiProvider().Invoke("com.vmware.nsx.upgrade", "upgradeuc", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), uIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}
