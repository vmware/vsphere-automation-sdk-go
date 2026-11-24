// Copyright (c) 2019-2026 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: VlanToVlanBridgeFailureDetectionProfiles
// Used by client-side stubs.

package enforcement_points

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	nsx_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type VlanToVlanBridgeFailureDetectionProfilesClient interface {

	// Deletes the specified profile containing the detection VLAN and BFD heartbeat settings.
	//
	// @param siteIdParam (required)
	// @param enforcementpointIdParam (required)
	// @param vlanToVlanBridgeFailureDetectionProfileIdParam (required)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Delete(siteIdParam string, enforcementpointIdParam string, vlanToVlanBridgeFailureDetectionProfileIdParam string) error

	// Returns the dedicated detection VLAN and BFD heartbeat settings for the specified profile.
	//
	// @param siteIdParam (required)
	// @param enforcementpointIdParam (required)
	// @param vlanToVlanBridgeFailureDetectionProfileIdParam (required)
	// @return com.vmware.nsx_policy.model.PolicyVlanToVlanBridgeFailureDetectionProfile
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(siteIdParam string, enforcementpointIdParam string, vlanToVlanBridgeFailureDetectionProfileIdParam string) (nsx_policyModel.PolicyVlanToVlanBridgeFailureDetectionProfile, error)

	// Returns a paginated list of bridge failure detection profiles for the specified site and enforcement point.
	//
	// @param siteIdParam (required)
	// @param enforcementpointIdParam (required)
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includeConflictsParam If true, resources currently in a sandboxed state due to federation conflicts will be included in the list results alongside active intent. Sandboxed resources will have has_conflict=true in the response. Applicable on LM only. On FNS, all resources including conflicting ones are always returned regardless of this parameter. Default is false. (optional, default to false)
	// @param includeMarkForDeleteObjectsParam If true, resources that are marked for deletion will be included in the results. By default, these resources are not included. (optional, default to false)
	// @param includeSystemOwnedParam Whether the list result contains system resources (optional, default to true)
	// @param includedFieldsParam Note - this parameter currently only works when used with the search APIs /policy/api/v1/search/query and /policy/api/v1/search/dsl. It is ignored for other list APIs. (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam If true, results are sorted in ascending order (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @return com.vmware.nsx_policy.model.PolicyVlanToVlanBridgeFailureDetectionProfileListResult
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(siteIdParam string, enforcementpointIdParam string, cursorParam *string, includeConflictsParam *bool, includeMarkForDeleteObjectsParam *bool, includeSystemOwnedParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (nsx_policyModel.PolicyVlanToVlanBridgeFailureDetectionProfileListResult, error)

	// Creates or updates the dedicated detection VLAN and BFD heartbeat settings. Reference this profile from a PolicyEdgeCluster to apply the configuration.
	//
	// @param siteIdParam (required)
	// @param enforcementpointIdParam (required)
	// @param vlanToVlanBridgeFailureDetectionProfileIdParam (required)
	// @param policyVlanToVlanBridgeFailureDetectionProfileParam (required)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Patch(siteIdParam string, enforcementpointIdParam string, vlanToVlanBridgeFailureDetectionProfileIdParam string, policyVlanToVlanBridgeFailureDetectionProfileParam nsx_policyModel.PolicyVlanToVlanBridgeFailureDetectionProfile) error

	// Creates or updates the dedicated detection VLAN and BFD heartbeat settings. Reference this profile from a PolicyEdgeCluster to apply the configuration.
	//
	// @param siteIdParam (required)
	// @param enforcementpointIdParam (required)
	// @param vlanToVlanBridgeFailureDetectionProfileIdParam (required)
	// @param policyVlanToVlanBridgeFailureDetectionProfileParam (required)
	// @return com.vmware.nsx_policy.model.PolicyVlanToVlanBridgeFailureDetectionProfile
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(siteIdParam string, enforcementpointIdParam string, vlanToVlanBridgeFailureDetectionProfileIdParam string, policyVlanToVlanBridgeFailureDetectionProfileParam nsx_policyModel.PolicyVlanToVlanBridgeFailureDetectionProfile) (nsx_policyModel.PolicyVlanToVlanBridgeFailureDetectionProfile, error)
}

type vlanToVlanBridgeFailureDetectionProfilesClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewVlanToVlanBridgeFailureDetectionProfilesClient(connector vapiProtocolClient_.Connector) *vlanToVlanBridgeFailureDetectionProfilesClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx_policy.infra.sites.enforcement_points.vlan_to_vlan_bridge_failure_detection_profiles")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"delete": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "delete"),
		"get":    vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
		"list":   vapiCore_.NewMethodIdentifier(interfaceIdentifier, "list"),
		"patch":  vapiCore_.NewMethodIdentifier(interfaceIdentifier, "patch"),
		"update": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "update"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	vIface := vlanToVlanBridgeFailureDetectionProfilesClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &vIface
}

func (vIface *vlanToVlanBridgeFailureDetectionProfilesClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := vIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (vIface *vlanToVlanBridgeFailureDetectionProfilesClient) Delete(siteIdParam string, enforcementpointIdParam string, vlanToVlanBridgeFailureDetectionProfileIdParam string) error {
	typeConverter := vIface.connector.TypeConverter()
	executionContext := vIface.connector.NewExecutionContext()
	operationRestMetaData := vlanToVlanBridgeFailureDetectionProfilesDeleteRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(vlanToVlanBridgeFailureDetectionProfilesDeleteInputType(), typeConverter)
	sv.AddStructField("SiteId", siteIdParam)
	sv.AddStructField("EnforcementpointId", enforcementpointIdParam)
	sv.AddStructField("VlanToVlanBridgeFailureDetectionProfileId", vlanToVlanBridgeFailureDetectionProfileIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := vIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.sites.enforcement_points.vlan_to_vlan_bridge_failure_detection_profiles", "delete", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), vIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (vIface *vlanToVlanBridgeFailureDetectionProfilesClient) Get(siteIdParam string, enforcementpointIdParam string, vlanToVlanBridgeFailureDetectionProfileIdParam string) (nsx_policyModel.PolicyVlanToVlanBridgeFailureDetectionProfile, error) {
	typeConverter := vIface.connector.TypeConverter()
	executionContext := vIface.connector.NewExecutionContext()
	operationRestMetaData := vlanToVlanBridgeFailureDetectionProfilesGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(vlanToVlanBridgeFailureDetectionProfilesGetInputType(), typeConverter)
	sv.AddStructField("SiteId", siteIdParam)
	sv.AddStructField("EnforcementpointId", enforcementpointIdParam)
	sv.AddStructField("VlanToVlanBridgeFailureDetectionProfileId", vlanToVlanBridgeFailureDetectionProfileIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.PolicyVlanToVlanBridgeFailureDetectionProfile
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := vIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.sites.enforcement_points.vlan_to_vlan_bridge_failure_detection_profiles", "get", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.PolicyVlanToVlanBridgeFailureDetectionProfile
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), VlanToVlanBridgeFailureDetectionProfilesGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.PolicyVlanToVlanBridgeFailureDetectionProfile), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), vIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (vIface *vlanToVlanBridgeFailureDetectionProfilesClient) List(siteIdParam string, enforcementpointIdParam string, cursorParam *string, includeConflictsParam *bool, includeMarkForDeleteObjectsParam *bool, includeSystemOwnedParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (nsx_policyModel.PolicyVlanToVlanBridgeFailureDetectionProfileListResult, error) {
	typeConverter := vIface.connector.TypeConverter()
	executionContext := vIface.connector.NewExecutionContext()
	operationRestMetaData := vlanToVlanBridgeFailureDetectionProfilesListRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(vlanToVlanBridgeFailureDetectionProfilesListInputType(), typeConverter)
	sv.AddStructField("SiteId", siteIdParam)
	sv.AddStructField("EnforcementpointId", enforcementpointIdParam)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludeConflicts", includeConflictsParam)
	sv.AddStructField("IncludeMarkForDeleteObjects", includeMarkForDeleteObjectsParam)
	sv.AddStructField("IncludeSystemOwned", includeSystemOwnedParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.PolicyVlanToVlanBridgeFailureDetectionProfileListResult
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := vIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.sites.enforcement_points.vlan_to_vlan_bridge_failure_detection_profiles", "list", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.PolicyVlanToVlanBridgeFailureDetectionProfileListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), VlanToVlanBridgeFailureDetectionProfilesListOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.PolicyVlanToVlanBridgeFailureDetectionProfileListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), vIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (vIface *vlanToVlanBridgeFailureDetectionProfilesClient) Patch(siteIdParam string, enforcementpointIdParam string, vlanToVlanBridgeFailureDetectionProfileIdParam string, policyVlanToVlanBridgeFailureDetectionProfileParam nsx_policyModel.PolicyVlanToVlanBridgeFailureDetectionProfile) error {
	typeConverter := vIface.connector.TypeConverter()
	executionContext := vIface.connector.NewExecutionContext()
	operationRestMetaData := vlanToVlanBridgeFailureDetectionProfilesPatchRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(vlanToVlanBridgeFailureDetectionProfilesPatchInputType(), typeConverter)
	sv.AddStructField("SiteId", siteIdParam)
	sv.AddStructField("EnforcementpointId", enforcementpointIdParam)
	sv.AddStructField("VlanToVlanBridgeFailureDetectionProfileId", vlanToVlanBridgeFailureDetectionProfileIdParam)
	sv.AddStructField("PolicyVlanToVlanBridgeFailureDetectionProfile", policyVlanToVlanBridgeFailureDetectionProfileParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := vIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.sites.enforcement_points.vlan_to_vlan_bridge_failure_detection_profiles", "patch", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), vIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (vIface *vlanToVlanBridgeFailureDetectionProfilesClient) Update(siteIdParam string, enforcementpointIdParam string, vlanToVlanBridgeFailureDetectionProfileIdParam string, policyVlanToVlanBridgeFailureDetectionProfileParam nsx_policyModel.PolicyVlanToVlanBridgeFailureDetectionProfile) (nsx_policyModel.PolicyVlanToVlanBridgeFailureDetectionProfile, error) {
	typeConverter := vIface.connector.TypeConverter()
	executionContext := vIface.connector.NewExecutionContext()
	operationRestMetaData := vlanToVlanBridgeFailureDetectionProfilesUpdateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(vlanToVlanBridgeFailureDetectionProfilesUpdateInputType(), typeConverter)
	sv.AddStructField("SiteId", siteIdParam)
	sv.AddStructField("EnforcementpointId", enforcementpointIdParam)
	sv.AddStructField("VlanToVlanBridgeFailureDetectionProfileId", vlanToVlanBridgeFailureDetectionProfileIdParam)
	sv.AddStructField("PolicyVlanToVlanBridgeFailureDetectionProfile", policyVlanToVlanBridgeFailureDetectionProfileParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.PolicyVlanToVlanBridgeFailureDetectionProfile
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := vIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.sites.enforcement_points.vlan_to_vlan_bridge_failure_detection_profiles", "update", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.PolicyVlanToVlanBridgeFailureDetectionProfile
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), VlanToVlanBridgeFailureDetectionProfilesUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.PolicyVlanToVlanBridgeFailureDetectionProfile), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), vIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
