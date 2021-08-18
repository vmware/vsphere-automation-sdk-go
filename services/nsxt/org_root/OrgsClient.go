// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Orgs
// Used by client-side stubs.

package org_root

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/core"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/lib"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/protocol/client"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt/model"
)

const _ = core.SupportedByRuntimeVersion1

type OrgsClient interface {

	// Delete Org.
	//
	// @param orgIdParam (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Delete(orgIdParam string) error

	// Get Org
	//
	// @param orgIdParam (required)
	// @return com.vmware.nsx_policy.model.Org
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(orgIdParam string) (model.Org, error)

	// Paginated list of Org.
	//
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includeMarkForDeleteObjectsParam Include objects that are marked for deletion in results (optional, default to false)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @return com.vmware.nsx_policy.model.OrgListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(cursorParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.OrgListResult, error)

	// If a Org with the org-1 is not already present, create a new Org. If it already exists, update the Org.
	//
	// @param orgIdParam (required)
	// @param orgParam (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Patch(orgIdParam string, orgParam model.Org) error

	// Update the Org.
	//
	// @param orgIdParam (required)
	// @param orgParam (required)
	// @return com.vmware.nsx_policy.model.Org
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(orgIdParam string, orgParam model.Org) (model.Org, error)
}

type orgsClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewOrgsClient(connector client.Connector) *orgsClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.nsx_policy.org_root.orgs")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"delete": core.NewMethodIdentifier(interfaceIdentifier, "delete"),
		"get":    core.NewMethodIdentifier(interfaceIdentifier, "get"),
		"list":   core.NewMethodIdentifier(interfaceIdentifier, "list"),
		"patch":  core.NewMethodIdentifier(interfaceIdentifier, "patch"),
		"update": core.NewMethodIdentifier(interfaceIdentifier, "update"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	oIface := orgsClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &oIface
}

func (oIface *orgsClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := oIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (oIface *orgsClient) Delete(orgIdParam string) error {
	typeConverter := oIface.connector.TypeConverter()
	executionContext := oIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(orgsDeleteInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := orgsDeleteRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	oIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := oIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.org_root.orgs", "delete", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), oIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (oIface *orgsClient) Get(orgIdParam string) (model.Org, error) {
	typeConverter := oIface.connector.TypeConverter()
	executionContext := oIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(orgsGetInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.Org
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := orgsGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	oIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := oIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.org_root.orgs", "get", inputDataValue, executionContext)
	var emptyOutput model.Org
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), orgsGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.Org), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), oIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (oIface *orgsClient) List(cursorParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.OrgListResult, error) {
	typeConverter := oIface.connector.TypeConverter()
	executionContext := oIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(orgsListInputType(), typeConverter)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludeMarkForDeleteObjects", includeMarkForDeleteObjectsParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.OrgListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := orgsListRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	oIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := oIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.org_root.orgs", "list", inputDataValue, executionContext)
	var emptyOutput model.OrgListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), orgsListOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.OrgListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), oIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (oIface *orgsClient) Patch(orgIdParam string, orgParam model.Org) error {
	typeConverter := oIface.connector.TypeConverter()
	executionContext := oIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(orgsPatchInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("Org", orgParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := orgsPatchRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	oIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := oIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.org_root.orgs", "patch", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), oIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (oIface *orgsClient) Update(orgIdParam string, orgParam model.Org) (model.Org, error) {
	typeConverter := oIface.connector.TypeConverter()
	executionContext := oIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(orgsUpdateInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("Org", orgParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.Org
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := orgsUpdateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	oIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := oIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.org_root.orgs", "update", inputDataValue, executionContext)
	var emptyOutput model.Org
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), orgsUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.Org), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), oIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
