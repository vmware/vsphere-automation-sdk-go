// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: PortDiscoveryProfileBindingMaps
// Used by client-side stubs.

package ports

import (
	"github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

const _ = core.SupportedByRuntimeVersion1

type PortDiscoveryProfileBindingMapsClient interface {

	// API will delete Port Discovery Profile Binding Profile
	//
	// @param orgIdParam (required)
	// @param deploymentGroupIdParam (required)
	// @param vhcIdParam (required)
	// @param tier1IdParam Tier-1 ID (required)
	// @param segmentIdParam Segment ID (required)
	// @param portIdParam Port ID (required)
	// @param portDiscoveryProfileBindingMapIdParam Port Discovery Profile Binding Map ID (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Delete(orgIdParam string, deploymentGroupIdParam string, vhcIdParam string, tier1IdParam string, segmentIdParam string, portIdParam string, portDiscoveryProfileBindingMapIdParam string) error

	// API will get Port Discovery Profile Binding Map
	//
	// @param orgIdParam (required)
	// @param deploymentGroupIdParam (required)
	// @param vhcIdParam (required)
	// @param tier1IdParam Tier-1 ID (required)
	// @param segmentIdParam Segment ID (required)
	// @param portIdParam Port ID (required)
	// @param portDiscoveryProfileBindingMapIdParam Port Discovery Profile Binding Map ID (required)
	// @return com.vmware.nsx_policy.model.PortDiscoveryProfileBindingMap
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(orgIdParam string, deploymentGroupIdParam string, vhcIdParam string, tier1IdParam string, segmentIdParam string, portIdParam string, portDiscoveryProfileBindingMapIdParam string) (model.PortDiscoveryProfileBindingMap, error)

	// API will list all Port Discovery Profile Binding Maps in current port id.
	//
	// @param orgIdParam (required)
	// @param deploymentGroupIdParam (required)
	// @param vhcIdParam (required)
	// @param tier1IdParam (required)
	// @param segmentIdParam (required)
	// @param portIdParam (required)
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includeMarkForDeleteObjectsParam Include objects that are marked for deletion in results (optional, default to false)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @return com.vmware.nsx_policy.model.PortDiscoveryProfileBindingMapListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(orgIdParam string, deploymentGroupIdParam string, vhcIdParam string, tier1IdParam string, segmentIdParam string, portIdParam string, cursorParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.PortDiscoveryProfileBindingMapListResult, error)

	// API will create Port Discovery Profile Binding Map. For objects with no binding maps, default profile is applied.
	//
	// @param orgIdParam (required)
	// @param deploymentGroupIdParam (required)
	// @param vhcIdParam (required)
	// @param tier1IdParam Tier-1 ID (required)
	// @param segmentIdParam Segment ID (required)
	// @param portIdParam Port ID (required)
	// @param portDiscoveryProfileBindingMapIdParam Port Discovery Profile Binding Map ID (required)
	// @param portDiscoveryProfileBindingMapParam (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Patch(orgIdParam string, deploymentGroupIdParam string, vhcIdParam string, tier1IdParam string, segmentIdParam string, portIdParam string, portDiscoveryProfileBindingMapIdParam string, portDiscoveryProfileBindingMapParam model.PortDiscoveryProfileBindingMap) error

	// API will update Port Discovery Profile Binding Map. For objects with no binding maps, default profile is applied.
	//
	// @param orgIdParam (required)
	// @param deploymentGroupIdParam (required)
	// @param vhcIdParam (required)
	// @param tier1IdParam Tier-1 ID (required)
	// @param segmentIdParam Segment ID (required)
	// @param portIdParam Port ID (required)
	// @param portDiscoveryProfileBindingMapIdParam Port Discovery Profile Binding Map ID (required)
	// @param portDiscoveryProfileBindingMapParam (required)
	// @return com.vmware.nsx_policy.model.PortDiscoveryProfileBindingMap
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(orgIdParam string, deploymentGroupIdParam string, vhcIdParam string, tier1IdParam string, segmentIdParam string, portIdParam string, portDiscoveryProfileBindingMapIdParam string, portDiscoveryProfileBindingMapParam model.PortDiscoveryProfileBindingMap) (model.PortDiscoveryProfileBindingMap, error)
}

type portDiscoveryProfileBindingMapsClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewPortDiscoveryProfileBindingMapsClient(connector client.Connector) *portDiscoveryProfileBindingMapsClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.nsx_policy.org_root.orgs.deployment_groups.vhcs.global_infra.tier_1s.segments.ports.port_discovery_profile_binding_maps")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"delete": core.NewMethodIdentifier(interfaceIdentifier, "delete"),
		"get":    core.NewMethodIdentifier(interfaceIdentifier, "get"),
		"list":   core.NewMethodIdentifier(interfaceIdentifier, "list"),
		"patch":  core.NewMethodIdentifier(interfaceIdentifier, "patch"),
		"update": core.NewMethodIdentifier(interfaceIdentifier, "update"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	pIface := portDiscoveryProfileBindingMapsClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &pIface
}

func (pIface *portDiscoveryProfileBindingMapsClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := pIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (pIface *portDiscoveryProfileBindingMapsClient) Delete(orgIdParam string, deploymentGroupIdParam string, vhcIdParam string, tier1IdParam string, segmentIdParam string, portIdParam string, portDiscoveryProfileBindingMapIdParam string) error {
	typeConverter := pIface.connector.TypeConverter()
	executionContext := pIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(portDiscoveryProfileBindingMapsDeleteInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("DeploymentGroupId", deploymentGroupIdParam)
	sv.AddStructField("VhcId", vhcIdParam)
	sv.AddStructField("Tier1Id", tier1IdParam)
	sv.AddStructField("SegmentId", segmentIdParam)
	sv.AddStructField("PortId", portIdParam)
	sv.AddStructField("PortDiscoveryProfileBindingMapId", portDiscoveryProfileBindingMapIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := portDiscoveryProfileBindingMapsDeleteRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	pIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := pIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.org_root.orgs.deployment_groups.vhcs.global_infra.tier_1s.segments.ports.port_discovery_profile_binding_maps", "delete", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), pIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (pIface *portDiscoveryProfileBindingMapsClient) Get(orgIdParam string, deploymentGroupIdParam string, vhcIdParam string, tier1IdParam string, segmentIdParam string, portIdParam string, portDiscoveryProfileBindingMapIdParam string) (model.PortDiscoveryProfileBindingMap, error) {
	typeConverter := pIface.connector.TypeConverter()
	executionContext := pIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(portDiscoveryProfileBindingMapsGetInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("DeploymentGroupId", deploymentGroupIdParam)
	sv.AddStructField("VhcId", vhcIdParam)
	sv.AddStructField("Tier1Id", tier1IdParam)
	sv.AddStructField("SegmentId", segmentIdParam)
	sv.AddStructField("PortId", portIdParam)
	sv.AddStructField("PortDiscoveryProfileBindingMapId", portDiscoveryProfileBindingMapIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.PortDiscoveryProfileBindingMap
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := portDiscoveryProfileBindingMapsGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	pIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := pIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.org_root.orgs.deployment_groups.vhcs.global_infra.tier_1s.segments.ports.port_discovery_profile_binding_maps", "get", inputDataValue, executionContext)
	var emptyOutput model.PortDiscoveryProfileBindingMap
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), portDiscoveryProfileBindingMapsGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.PortDiscoveryProfileBindingMap), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), pIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (pIface *portDiscoveryProfileBindingMapsClient) List(orgIdParam string, deploymentGroupIdParam string, vhcIdParam string, tier1IdParam string, segmentIdParam string, portIdParam string, cursorParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.PortDiscoveryProfileBindingMapListResult, error) {
	typeConverter := pIface.connector.TypeConverter()
	executionContext := pIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(portDiscoveryProfileBindingMapsListInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("DeploymentGroupId", deploymentGroupIdParam)
	sv.AddStructField("VhcId", vhcIdParam)
	sv.AddStructField("Tier1Id", tier1IdParam)
	sv.AddStructField("SegmentId", segmentIdParam)
	sv.AddStructField("PortId", portIdParam)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludeMarkForDeleteObjects", includeMarkForDeleteObjectsParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.PortDiscoveryProfileBindingMapListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := portDiscoveryProfileBindingMapsListRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	pIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := pIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.org_root.orgs.deployment_groups.vhcs.global_infra.tier_1s.segments.ports.port_discovery_profile_binding_maps", "list", inputDataValue, executionContext)
	var emptyOutput model.PortDiscoveryProfileBindingMapListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), portDiscoveryProfileBindingMapsListOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.PortDiscoveryProfileBindingMapListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), pIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (pIface *portDiscoveryProfileBindingMapsClient) Patch(orgIdParam string, deploymentGroupIdParam string, vhcIdParam string, tier1IdParam string, segmentIdParam string, portIdParam string, portDiscoveryProfileBindingMapIdParam string, portDiscoveryProfileBindingMapParam model.PortDiscoveryProfileBindingMap) error {
	typeConverter := pIface.connector.TypeConverter()
	executionContext := pIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(portDiscoveryProfileBindingMapsPatchInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("DeploymentGroupId", deploymentGroupIdParam)
	sv.AddStructField("VhcId", vhcIdParam)
	sv.AddStructField("Tier1Id", tier1IdParam)
	sv.AddStructField("SegmentId", segmentIdParam)
	sv.AddStructField("PortId", portIdParam)
	sv.AddStructField("PortDiscoveryProfileBindingMapId", portDiscoveryProfileBindingMapIdParam)
	sv.AddStructField("PortDiscoveryProfileBindingMap", portDiscoveryProfileBindingMapParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := portDiscoveryProfileBindingMapsPatchRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	pIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := pIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.org_root.orgs.deployment_groups.vhcs.global_infra.tier_1s.segments.ports.port_discovery_profile_binding_maps", "patch", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), pIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (pIface *portDiscoveryProfileBindingMapsClient) Update(orgIdParam string, deploymentGroupIdParam string, vhcIdParam string, tier1IdParam string, segmentIdParam string, portIdParam string, portDiscoveryProfileBindingMapIdParam string, portDiscoveryProfileBindingMapParam model.PortDiscoveryProfileBindingMap) (model.PortDiscoveryProfileBindingMap, error) {
	typeConverter := pIface.connector.TypeConverter()
	executionContext := pIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(portDiscoveryProfileBindingMapsUpdateInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("DeploymentGroupId", deploymentGroupIdParam)
	sv.AddStructField("VhcId", vhcIdParam)
	sv.AddStructField("Tier1Id", tier1IdParam)
	sv.AddStructField("SegmentId", segmentIdParam)
	sv.AddStructField("PortId", portIdParam)
	sv.AddStructField("PortDiscoveryProfileBindingMapId", portDiscoveryProfileBindingMapIdParam)
	sv.AddStructField("PortDiscoveryProfileBindingMap", portDiscoveryProfileBindingMapParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.PortDiscoveryProfileBindingMap
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := portDiscoveryProfileBindingMapsUpdateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	pIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := pIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.org_root.orgs.deployment_groups.vhcs.global_infra.tier_1s.segments.ports.port_discovery_profile_binding_maps", "update", inputDataValue, executionContext)
	var emptyOutput model.PortDiscoveryProfileBindingMap
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), portDiscoveryProfileBindingMapsUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.PortDiscoveryProfileBindingMap), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), pIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
