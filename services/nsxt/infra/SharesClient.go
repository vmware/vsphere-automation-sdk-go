// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Shares
// Used by client-side stubs.

package infra

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/core"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/lib"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/protocol/client"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt/model"
)

const _ = core.SupportedByRuntimeVersion1

type SharesClient interface {

	// Get single instance of a share
	//
	// @param shareIdParam Share ID (required)
	// @return com.vmware.nsx_policy.model.Share
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(shareIdParam string) (model.Share, error)

	// List shares
	//
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includeMarkForDeleteObjectsParam Include objects that are marked for deletion in results (optional, default to false)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @return com.vmware.nsx_policy.model.ShareListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(cursorParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.ShareListResult, error)
}

type sharesClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewSharesClient(connector client.Connector) *sharesClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.nsx_policy.infra.shares")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"get":  core.NewMethodIdentifier(interfaceIdentifier, "get"),
		"list": core.NewMethodIdentifier(interfaceIdentifier, "list"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	sIface := sharesClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &sIface
}

func (sIface *sharesClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := sIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (sIface *sharesClient) Get(shareIdParam string) (model.Share, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(sharesGetInputType(), typeConverter)
	sv.AddStructField("ShareId", shareIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.Share
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := sharesGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.shares", "get", inputDataValue, executionContext)
	var emptyOutput model.Share
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), sharesGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.Share), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *sharesClient) List(cursorParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.ShareListResult, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(sharesListInputType(), typeConverter)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludeMarkForDeleteObjects", includeMarkForDeleteObjectsParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ShareListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := sharesListRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.shares", "list", inputDataValue, executionContext)
	var emptyOutput model.ShareListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), sharesListOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ShareListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
