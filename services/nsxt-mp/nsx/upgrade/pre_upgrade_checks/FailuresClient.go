// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Failures
// Used by client-side stubs.

package pre_upgrade_checks

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	nsxModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type FailuresClient interface {

	// Get failures resulting from the last execution of pre-upgrade checks. If the execution of checks is in progress, the response has the list of failures observed so far.
	//
	// @param componentTypeParam Component type (optional)
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param filterTextParam Filter text (optional)
	// @param groupIdParam Filter on the group id (optional)
	// @param groupNameParam Filter on the group name (optional)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param needsAckParam Filter based on acknowledgement required (optional)
	// @param originTypeParam Type of origin of failure (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @param type_Param Status of the upgrade check (optional)
	// @param unitIdParam Filter on the unit id (optional)
	// @param unitNameParam Filter on the unit name (optional)
	// @return com.vmware.nsx.model.UpgradeCheckFailureListResult
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(componentTypeParam *string, cursorParam *string, filterTextParam *string, groupIdParam *string, groupNameParam *string, includedFieldsParam *string, needsAckParam *bool, originTypeParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string, type_Param *string, unitIdParam *string, unitNameParam *string) (nsxModel.UpgradeCheckFailureListResult, error)
}

type failuresClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewFailuresClient(connector vapiProtocolClient_.Connector) *failuresClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx.upgrade.pre_upgrade_checks.failures")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"list": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "list"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	fIface := failuresClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &fIface
}

func (fIface *failuresClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := fIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (fIface *failuresClient) List(componentTypeParam *string, cursorParam *string, filterTextParam *string, groupIdParam *string, groupNameParam *string, includedFieldsParam *string, needsAckParam *bool, originTypeParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string, type_Param *string, unitIdParam *string, unitNameParam *string) (nsxModel.UpgradeCheckFailureListResult, error) {
	typeConverter := fIface.connector.TypeConverter()
	executionContext := fIface.connector.NewExecutionContext()
	operationRestMetaData := failuresListRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(failuresListInputType(), typeConverter)
	sv.AddStructField("ComponentType", componentTypeParam)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("FilterText", filterTextParam)
	sv.AddStructField("GroupId", groupIdParam)
	sv.AddStructField("GroupName", groupNameParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("NeedsAck", needsAckParam)
	sv.AddStructField("OriginType", originTypeParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	sv.AddStructField("Type_", type_Param)
	sv.AddStructField("UnitId", unitIdParam)
	sv.AddStructField("UnitName", unitNameParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.UpgradeCheckFailureListResult
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := fIface.connector.GetApiProvider().Invoke("com.vmware.nsx.upgrade.pre_upgrade_checks.failures", "list", inputDataValue, executionContext)
	var emptyOutput nsxModel.UpgradeCheckFailureListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), FailuresListOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.UpgradeCheckFailureListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), fIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
