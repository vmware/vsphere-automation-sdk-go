// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: ManagementPlaneAPIGroupingObjectsNSServiceGroups
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

type ManagementPlaneAPIGroupingObjectsNSServiceGroupsClient interface {

	// Deletes the specified NSServiceGroup. By default, if the NSServiceGroup is consumed in a Firewall rule, it won't get deleted. In such situations, pass \"force=true\" as query param to force delete the NSServiceGroup.
	//
	// @param nsServiceGroupIdParam NSServiceGroup Id (required)
	// @param forceParam Force delete the resource even if it is being used somewhere (optional, default to false)
	// @throws ConcurrentChange  Conflict
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	DeleteNSServiceGroup(nsServiceGroupIdParam string, forceParam *bool) error

	// Returns paginated list of NSServiceGroups
	//
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param defaultServiceParam Fetch all default NSServiceGroups (optional)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @return com.vmware.model.NSServiceGroupListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ListNSServiceGroups(cursorParam *string, defaultServiceParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.NSServiceGroupListResult, error)

	// Returns information about the specified NSServiceGroup
	//
	// @param nsServiceGroupIdParam NSServiceGroup Id (required)
	// @return com.vmware.model.NSServiceGroup
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ReadNSServiceGroup(nsServiceGroupIdParam string) (model.NSServiceGroup, error)

	// Updates the specified NSService. Modifiable parameters include the description, display_name and members.
	//
	// @param nsServiceGroupIdParam NSServiceGroup Id (required)
	// @param nsServiceGroupParam (required)
	// @return com.vmware.model.NSServiceGroup
	// @throws ConcurrentChange  Conflict
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	UpdateNSServiceGroup(nsServiceGroupIdParam string, nsServiceGroupParam model.NSServiceGroup) (model.NSServiceGroup, error)
}

type managementPlaneAPIGroupingObjectsNSServiceGroupsClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewManagementPlaneAPIGroupingObjectsNSServiceGroupsClient(connector client.Connector) *managementPlaneAPIGroupingObjectsNSServiceGroupsClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.api.management_plane_API_grouping_objects_NS_service_groups")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"delete_NS_service_group": core.NewMethodIdentifier(interfaceIdentifier, "delete_NS_service_group"),
		"list_NS_service_groups":  core.NewMethodIdentifier(interfaceIdentifier, "list_NS_service_groups"),
		"read_NS_service_group":   core.NewMethodIdentifier(interfaceIdentifier, "read_NS_service_group"),
		"update_NS_service_group": core.NewMethodIdentifier(interfaceIdentifier, "update_NS_service_group"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	mIface := managementPlaneAPIGroupingObjectsNSServiceGroupsClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &mIface
}

func (mIface *managementPlaneAPIGroupingObjectsNSServiceGroupsClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := mIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (mIface *managementPlaneAPIGroupingObjectsNSServiceGroupsClient) DeleteNSServiceGroup(nsServiceGroupIdParam string, forceParam *bool) error {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPIGroupingObjectsNSServiceGroupsDeleteNSServiceGroupInputType(), typeConverter)
	sv.AddStructField("NsServiceGroupId", nsServiceGroupIdParam)
	sv.AddStructField("Force", forceParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPIGroupingObjectsNSServiceGroupsDeleteNSServiceGroupRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_grouping_objects_NS_service_groups", "delete_NS_service_group", inputDataValue, executionContext)
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

func (mIface *managementPlaneAPIGroupingObjectsNSServiceGroupsClient) ListNSServiceGroups(cursorParam *string, defaultServiceParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.NSServiceGroupListResult, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPIGroupingObjectsNSServiceGroupsListNSServiceGroupsInputType(), typeConverter)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("DefaultService", defaultServiceParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.NSServiceGroupListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPIGroupingObjectsNSServiceGroupsListNSServiceGroupsRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_grouping_objects_NS_service_groups", "list_NS_service_groups", inputDataValue, executionContext)
	var emptyOutput model.NSServiceGroupListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPIGroupingObjectsNSServiceGroupsListNSServiceGroupsOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.NSServiceGroupListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPIGroupingObjectsNSServiceGroupsClient) ReadNSServiceGroup(nsServiceGroupIdParam string) (model.NSServiceGroup, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPIGroupingObjectsNSServiceGroupsReadNSServiceGroupInputType(), typeConverter)
	sv.AddStructField("NsServiceGroupId", nsServiceGroupIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.NSServiceGroup
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPIGroupingObjectsNSServiceGroupsReadNSServiceGroupRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_grouping_objects_NS_service_groups", "read_NS_service_group", inputDataValue, executionContext)
	var emptyOutput model.NSServiceGroup
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPIGroupingObjectsNSServiceGroupsReadNSServiceGroupOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.NSServiceGroup), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPIGroupingObjectsNSServiceGroupsClient) UpdateNSServiceGroup(nsServiceGroupIdParam string, nsServiceGroupParam model.NSServiceGroup) (model.NSServiceGroup, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPIGroupingObjectsNSServiceGroupsUpdateNSServiceGroupInputType(), typeConverter)
	sv.AddStructField("NsServiceGroupId", nsServiceGroupIdParam)
	sv.AddStructField("NsServiceGroup", nsServiceGroupParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.NSServiceGroup
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPIGroupingObjectsNSServiceGroupsUpdateNSServiceGroupRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_grouping_objects_NS_service_groups", "update_NS_service_group", inputDataValue, executionContext)
	var emptyOutput model.NSServiceGroup
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPIGroupingObjectsNSServiceGroupsUpdateNSServiceGroupOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.NSServiceGroup), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
