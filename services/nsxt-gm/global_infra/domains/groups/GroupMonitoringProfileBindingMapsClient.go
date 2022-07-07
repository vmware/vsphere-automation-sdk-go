// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: GroupMonitoringProfileBindingMaps
// Used by client-side stubs.

package groups

import (
	"github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-gm/model"
)

const _ = core.SupportedByRuntimeVersion1

type GroupMonitoringProfileBindingMapsClient interface {

	// API will get Group Monitoring Profile Binding Map
	//
	// @param domainIdParam Domain-ID (required)
	// @param groupIdParam Group ID (required)
	// @param groupMonitoringProfileBindingMapIdParam Group Monitoring Profile Binding Map ID (required)
	// @return com.vmware.nsx_global_policy.model.GroupMonitoringProfileBindingMap
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(domainIdParam string, groupIdParam string, groupMonitoringProfileBindingMapIdParam string) (model.GroupMonitoringProfileBindingMap, error)

	// API will list all Group Monitoring Profile Binding Maps in current group id.
	//
	// @param domainIdParam (required)
	// @param groupIdParam (required)
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includeMarkForDeleteObjectsParam Include objects that are marked for deletion in results (optional, default to false)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @return com.vmware.nsx_global_policy.model.GroupMonitoringProfileBindingMapListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(domainIdParam string, groupIdParam string, cursorParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.GroupMonitoringProfileBindingMapListResult, error)
}

type groupMonitoringProfileBindingMapsClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewGroupMonitoringProfileBindingMapsClient(connector client.Connector) *groupMonitoringProfileBindingMapsClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.nsx_global_policy.global_infra.domains.groups.group_monitoring_profile_binding_maps")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"get":  core.NewMethodIdentifier(interfaceIdentifier, "get"),
		"list": core.NewMethodIdentifier(interfaceIdentifier, "list"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	gIface := groupMonitoringProfileBindingMapsClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &gIface
}

func (gIface *groupMonitoringProfileBindingMapsClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := gIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (gIface *groupMonitoringProfileBindingMapsClient) Get(domainIdParam string, groupIdParam string, groupMonitoringProfileBindingMapIdParam string) (model.GroupMonitoringProfileBindingMap, error) {
	typeConverter := gIface.connector.TypeConverter()
	executionContext := gIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(groupMonitoringProfileBindingMapsGetInputType(), typeConverter)
	sv.AddStructField("DomainId", domainIdParam)
	sv.AddStructField("GroupId", groupIdParam)
	sv.AddStructField("GroupMonitoringProfileBindingMapId", groupMonitoringProfileBindingMapIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.GroupMonitoringProfileBindingMap
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := groupMonitoringProfileBindingMapsGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	gIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := gIface.connector.GetApiProvider().Invoke("com.vmware.nsx_global_policy.global_infra.domains.groups.group_monitoring_profile_binding_maps", "get", inputDataValue, executionContext)
	var emptyOutput model.GroupMonitoringProfileBindingMap
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), groupMonitoringProfileBindingMapsGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.GroupMonitoringProfileBindingMap), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), gIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (gIface *groupMonitoringProfileBindingMapsClient) List(domainIdParam string, groupIdParam string, cursorParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.GroupMonitoringProfileBindingMapListResult, error) {
	typeConverter := gIface.connector.TypeConverter()
	executionContext := gIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(groupMonitoringProfileBindingMapsListInputType(), typeConverter)
	sv.AddStructField("DomainId", domainIdParam)
	sv.AddStructField("GroupId", groupIdParam)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludeMarkForDeleteObjects", includeMarkForDeleteObjectsParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.GroupMonitoringProfileBindingMapListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := groupMonitoringProfileBindingMapsListRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	gIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := gIface.connector.GetApiProvider().Invoke("com.vmware.nsx_global_policy.global_infra.domains.groups.group_monitoring_profile_binding_maps", "list", inputDataValue, executionContext)
	var emptyOutput model.GroupMonitoringProfileBindingMapListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), groupMonitoringProfileBindingMapsListOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.GroupMonitoringProfileBindingMapListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), gIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
