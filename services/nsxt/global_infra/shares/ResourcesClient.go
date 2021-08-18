// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Resources
// Used by client-side stubs.

package shares

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/core"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/lib"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/protocol/client"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt/model"
)

const _ = core.SupportedByRuntimeVersion1

type ResourcesClient interface {

	// Get the shared resource that represents shared objects.
	//
	// @param shareIdParam Share ID (required)
	// @param sharedResourceIdParam Shared resource id (required)
	// @return com.vmware.nsx_policy.model.SharedResource
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(shareIdParam string, sharedResourceIdParam string) (model.SharedResource, error)

	// Get the list of shared resource in a give n share.
	//
	// @param shareIdParam Share ID (required)
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includeMarkForDeleteObjectsParam Include objects that are marked for deletion in results (optional, default to false)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param resourceTypeParam resource type (optional)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @return com.vmware.nsx_policy.model.SharedResourceListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(shareIdParam string, cursorParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, resourceTypeParam *string, sortAscendingParam *bool, sortByParam *string) (model.SharedResourceListResult, error)

	// If a Shared resource is not already present, create a new shared resource. If it already exists, update the shared resource entity. Presense of the shared resource denotes sharing of the parent object.
	//
	// @param shareIdParam Share ID (required)
	// @param sharedResourceIdParam Shared resource id (required)
	// @param sharedResourceParam (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Patch(shareIdParam string, sharedResourceIdParam string, sharedResourceParam model.SharedResource) error
}

type resourcesClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewResourcesClient(connector client.Connector) *resourcesClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.nsx_policy.global_infra.shares.resources")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"get":   core.NewMethodIdentifier(interfaceIdentifier, "get"),
		"list":  core.NewMethodIdentifier(interfaceIdentifier, "list"),
		"patch": core.NewMethodIdentifier(interfaceIdentifier, "patch"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	rIface := resourcesClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &rIface
}

func (rIface *resourcesClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := rIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (rIface *resourcesClient) Get(shareIdParam string, sharedResourceIdParam string) (model.SharedResource, error) {
	typeConverter := rIface.connector.TypeConverter()
	executionContext := rIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(resourcesGetInputType(), typeConverter)
	sv.AddStructField("ShareId", shareIdParam)
	sv.AddStructField("SharedResourceId", sharedResourceIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.SharedResource
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := resourcesGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	rIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := rIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.global_infra.shares.resources", "get", inputDataValue, executionContext)
	var emptyOutput model.SharedResource
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), resourcesGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.SharedResource), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), rIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (rIface *resourcesClient) List(shareIdParam string, cursorParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, resourceTypeParam *string, sortAscendingParam *bool, sortByParam *string) (model.SharedResourceListResult, error) {
	typeConverter := rIface.connector.TypeConverter()
	executionContext := rIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(resourcesListInputType(), typeConverter)
	sv.AddStructField("ShareId", shareIdParam)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludeMarkForDeleteObjects", includeMarkForDeleteObjectsParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("ResourceType", resourceTypeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.SharedResourceListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := resourcesListRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	rIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := rIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.global_infra.shares.resources", "list", inputDataValue, executionContext)
	var emptyOutput model.SharedResourceListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), resourcesListOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.SharedResourceListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), rIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (rIface *resourcesClient) Patch(shareIdParam string, sharedResourceIdParam string, sharedResourceParam model.SharedResource) error {
	typeConverter := rIface.connector.TypeConverter()
	executionContext := rIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(resourcesPatchInputType(), typeConverter)
	sv.AddStructField("ShareId", shareIdParam)
	sv.AddStructField("SharedResourceId", sharedResourceIdParam)
	sv.AddStructField("SharedResource", sharedResourceParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := resourcesPatchRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	rIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := rIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.global_infra.shares.resources", "patch", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), rIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}
