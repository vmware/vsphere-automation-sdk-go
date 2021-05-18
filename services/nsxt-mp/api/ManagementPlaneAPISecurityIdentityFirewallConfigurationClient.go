// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: ManagementPlaneAPISecurityIdentityFirewallConfiguration
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

type ManagementPlaneAPISecurityIdentityFirewallConfigurationClient interface {

	// Delete individual compute collections for IDFW.
	//
	// @param ccExtIdParam (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	DeleteEnabledComputeCollection(ccExtIdParam string) error

	// Get enable/disable status of individual compute collections for IDFW.
	//
	// @param ccExtIdParam (required)
	// @return com.vmware.model.IdfwEnabledComputeCollection
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetEnabledComputeCollection(ccExtIdParam string) (model.IdfwEnabledComputeCollection, error)

	// Fetches IDFW master switch setting to check whether master switch is enabled or disabled
	// @return com.vmware.model.IdfwMasterSwitchSetting
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetMasterSwitchSetting() (model.IdfwMasterSwitchSetting, error)

	// Fetches IDFW standalone hosts switch setting to check whether standalone hosts is enabled or disabled
	// @return com.vmware.model.IdfwStandaloneHostsSwitchSetting
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetStandaloneHostsSwitchSetting() (model.IdfwStandaloneHostsSwitchSetting, error)

	// List all Identity firewall compute collections.
	//
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @return com.vmware.model.IdfwEnabledComputeCollectionListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ListEnabledComputeCollections(cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.IdfwEnabledComputeCollectionListResult, error)

	// Enable/disable individual compute collections for IDFW.
	//
	// @param ccExtIdParam (required)
	// @param idfwEnabledComputeCollectionParam (required)
	// @return com.vmware.model.IdfwEnabledComputeCollection
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	UpdateEnabledComputeCollection(ccExtIdParam string, idfwEnabledComputeCollectionParam model.IdfwEnabledComputeCollection) (model.IdfwEnabledComputeCollection, error)

	// Update Identity Firewall master switch setting (true=enabled / false=disabled). Identity Firewall master switch setting enables or disables Identity Firewall feature across the system. It affects compute collections, hypervisor and virtual machines. This operation is expensive and also has big impact and implication on system perforamce.
	//
	// @param idfwMasterSwitchSettingParam (required)
	// @return com.vmware.model.IdfwMasterSwitchSetting
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	UpdateMasterSwitchSetting(idfwMasterSwitchSettingParam model.IdfwMasterSwitchSetting) (model.IdfwMasterSwitchSetting, error)

	// Update Identity Firewall standalone hosts switch setting (true=enabled / false=disabled).
	//
	// @param idfwStandaloneHostsSwitchSettingParam (required)
	// @return com.vmware.model.IdfwStandaloneHostsSwitchSetting
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	UpdateStandaloneHostsSwitchSetting(idfwStandaloneHostsSwitchSettingParam model.IdfwStandaloneHostsSwitchSetting) (model.IdfwStandaloneHostsSwitchSetting, error)
}

type managementPlaneAPISecurityIdentityFirewallConfigurationClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewManagementPlaneAPISecurityIdentityFirewallConfigurationClient(connector client.Connector) *managementPlaneAPISecurityIdentityFirewallConfigurationClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.api.management_plane_API_security_identity_firewall_configuration")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"delete_enabled_compute_collection":      core.NewMethodIdentifier(interfaceIdentifier, "delete_enabled_compute_collection"),
		"get_enabled_compute_collection":         core.NewMethodIdentifier(interfaceIdentifier, "get_enabled_compute_collection"),
		"get_master_switch_setting":              core.NewMethodIdentifier(interfaceIdentifier, "get_master_switch_setting"),
		"get_standalone_hosts_switch_setting":    core.NewMethodIdentifier(interfaceIdentifier, "get_standalone_hosts_switch_setting"),
		"list_enabled_compute_collections":       core.NewMethodIdentifier(interfaceIdentifier, "list_enabled_compute_collections"),
		"update_enabled_compute_collection":      core.NewMethodIdentifier(interfaceIdentifier, "update_enabled_compute_collection"),
		"update_master_switch_setting":           core.NewMethodIdentifier(interfaceIdentifier, "update_master_switch_setting"),
		"update_standalone_hosts_switch_setting": core.NewMethodIdentifier(interfaceIdentifier, "update_standalone_hosts_switch_setting"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	mIface := managementPlaneAPISecurityIdentityFirewallConfigurationClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &mIface
}

func (mIface *managementPlaneAPISecurityIdentityFirewallConfigurationClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := mIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (mIface *managementPlaneAPISecurityIdentityFirewallConfigurationClient) DeleteEnabledComputeCollection(ccExtIdParam string) error {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPISecurityIdentityFirewallConfigurationDeleteEnabledComputeCollectionInputType(), typeConverter)
	sv.AddStructField("CcExtId", ccExtIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPISecurityIdentityFirewallConfigurationDeleteEnabledComputeCollectionRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_security_identity_firewall_configuration", "delete_enabled_compute_collection", inputDataValue, executionContext)
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

func (mIface *managementPlaneAPISecurityIdentityFirewallConfigurationClient) GetEnabledComputeCollection(ccExtIdParam string) (model.IdfwEnabledComputeCollection, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPISecurityIdentityFirewallConfigurationGetEnabledComputeCollectionInputType(), typeConverter)
	sv.AddStructField("CcExtId", ccExtIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.IdfwEnabledComputeCollection
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPISecurityIdentityFirewallConfigurationGetEnabledComputeCollectionRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_security_identity_firewall_configuration", "get_enabled_compute_collection", inputDataValue, executionContext)
	var emptyOutput model.IdfwEnabledComputeCollection
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPISecurityIdentityFirewallConfigurationGetEnabledComputeCollectionOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.IdfwEnabledComputeCollection), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPISecurityIdentityFirewallConfigurationClient) GetMasterSwitchSetting() (model.IdfwMasterSwitchSetting, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPISecurityIdentityFirewallConfigurationGetMasterSwitchSettingInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.IdfwMasterSwitchSetting
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPISecurityIdentityFirewallConfigurationGetMasterSwitchSettingRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_security_identity_firewall_configuration", "get_master_switch_setting", inputDataValue, executionContext)
	var emptyOutput model.IdfwMasterSwitchSetting
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPISecurityIdentityFirewallConfigurationGetMasterSwitchSettingOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.IdfwMasterSwitchSetting), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPISecurityIdentityFirewallConfigurationClient) GetStandaloneHostsSwitchSetting() (model.IdfwStandaloneHostsSwitchSetting, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPISecurityIdentityFirewallConfigurationGetStandaloneHostsSwitchSettingInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.IdfwStandaloneHostsSwitchSetting
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPISecurityIdentityFirewallConfigurationGetStandaloneHostsSwitchSettingRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_security_identity_firewall_configuration", "get_standalone_hosts_switch_setting", inputDataValue, executionContext)
	var emptyOutput model.IdfwStandaloneHostsSwitchSetting
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPISecurityIdentityFirewallConfigurationGetStandaloneHostsSwitchSettingOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.IdfwStandaloneHostsSwitchSetting), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPISecurityIdentityFirewallConfigurationClient) ListEnabledComputeCollections(cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.IdfwEnabledComputeCollectionListResult, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPISecurityIdentityFirewallConfigurationListEnabledComputeCollectionsInputType(), typeConverter)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.IdfwEnabledComputeCollectionListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPISecurityIdentityFirewallConfigurationListEnabledComputeCollectionsRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_security_identity_firewall_configuration", "list_enabled_compute_collections", inputDataValue, executionContext)
	var emptyOutput model.IdfwEnabledComputeCollectionListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPISecurityIdentityFirewallConfigurationListEnabledComputeCollectionsOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.IdfwEnabledComputeCollectionListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPISecurityIdentityFirewallConfigurationClient) UpdateEnabledComputeCollection(ccExtIdParam string, idfwEnabledComputeCollectionParam model.IdfwEnabledComputeCollection) (model.IdfwEnabledComputeCollection, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPISecurityIdentityFirewallConfigurationUpdateEnabledComputeCollectionInputType(), typeConverter)
	sv.AddStructField("CcExtId", ccExtIdParam)
	sv.AddStructField("IdfwEnabledComputeCollection", idfwEnabledComputeCollectionParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.IdfwEnabledComputeCollection
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPISecurityIdentityFirewallConfigurationUpdateEnabledComputeCollectionRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_security_identity_firewall_configuration", "update_enabled_compute_collection", inputDataValue, executionContext)
	var emptyOutput model.IdfwEnabledComputeCollection
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPISecurityIdentityFirewallConfigurationUpdateEnabledComputeCollectionOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.IdfwEnabledComputeCollection), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPISecurityIdentityFirewallConfigurationClient) UpdateMasterSwitchSetting(idfwMasterSwitchSettingParam model.IdfwMasterSwitchSetting) (model.IdfwMasterSwitchSetting, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPISecurityIdentityFirewallConfigurationUpdateMasterSwitchSettingInputType(), typeConverter)
	sv.AddStructField("IdfwMasterSwitchSetting", idfwMasterSwitchSettingParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.IdfwMasterSwitchSetting
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPISecurityIdentityFirewallConfigurationUpdateMasterSwitchSettingRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_security_identity_firewall_configuration", "update_master_switch_setting", inputDataValue, executionContext)
	var emptyOutput model.IdfwMasterSwitchSetting
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPISecurityIdentityFirewallConfigurationUpdateMasterSwitchSettingOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.IdfwMasterSwitchSetting), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPISecurityIdentityFirewallConfigurationClient) UpdateStandaloneHostsSwitchSetting(idfwStandaloneHostsSwitchSettingParam model.IdfwStandaloneHostsSwitchSetting) (model.IdfwStandaloneHostsSwitchSetting, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPISecurityIdentityFirewallConfigurationUpdateStandaloneHostsSwitchSettingInputType(), typeConverter)
	sv.AddStructField("IdfwStandaloneHostsSwitchSetting", idfwStandaloneHostsSwitchSettingParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.IdfwStandaloneHostsSwitchSetting
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPISecurityIdentityFirewallConfigurationUpdateStandaloneHostsSwitchSettingRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_security_identity_firewall_configuration", "update_standalone_hosts_switch_setting", inputDataValue, executionContext)
	var emptyOutput model.IdfwStandaloneHostsSwitchSetting
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPISecurityIdentityFirewallConfigurationUpdateStandaloneHostsSwitchSettingOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.IdfwStandaloneHostsSwitchSetting), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
