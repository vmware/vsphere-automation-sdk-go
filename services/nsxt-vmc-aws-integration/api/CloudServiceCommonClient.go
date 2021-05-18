// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: CloudServiceCommon
// Used by client-side stubs.

package api

import (
	"github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-vmc-aws-integration/model"
)

const _ = core.SupportedByRuntimeVersion1

type CloudServiceCommonClient interface {

	// This API is used to get external configuration for north-south traffic. For eg., in AWS case, this would refer to Direct Connect config. For Dimension, this would refer to the config at Upstream Intranet interface to Tor.
	// @return com.vmware.model.ExternalConnectivityConfig
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetExternalConnectivityConfig() (model.ExternalConnectivityConfig, error)

	// Get management VM access information.
	//
	// @param vmIdParam (required)
	// @return com.vmware.model.MgmtVmInfo
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetMgmtVmInfo(vmIdParam string) (model.MgmtVmInfo, error)

	// Get the consolidated realized entities given an intent path, specified in the query parameter. The intent object under the intent path is indicated by a specific VMC-App API and can contain multiple objects.
	//
	// @param intentPathParam String Path of the intent object (required)
	// @param sitePathParam String Path of the site (optional)
	// @return com.vmware.model.VmcRealizedEntities
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetRealizedEntities(intentPathParam string, sitePathParam *string) (model.VmcRealizedEntities, error)

	// Get the consolidated status of an intent object, specified by path in query parameter. The intent object is indicated by a specific VMC-App API and can contain multiple objects. For example, /infra/direct-connect/bgp can return the consolidated status of ASN update and route preference update.
	//
	// @param intentPathParam String Path of the intent object (required)
	// @param sitePathParam String Path of the site (optional)
	// @return com.vmware.model.VmcConsolidatedRealizedStatus
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetRealizedStateStatus(intentPathParam string, sitePathParam *string) (model.VmcConsolidatedRealizedStatus, error)

	// Get the user-level SDDC configuration parameters
	// @return com.vmware.model.SddcUserConfiguration
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetSddcUserConfig() (model.SddcUserConfiguration, error)

	// Get vmc environment feature flags
	//
	// @param internalNameParam internal feature name (optional)
	// @param nameParam feature name (optional)
	// @return com.vmware.model.VmcFeatureFlags
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetVmcFeatureFlags(internalNameParam *string, nameParam *string) (model.VmcFeatureFlags, error)

	// List Management VM information.
	// @return com.vmware.model.MgmtVmsListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ListMgmtVms() (model.MgmtVmsListResult, error)

	// This API is used to update intranet configuration for north-south traffic. For eg., in AWS case, this would refer to Direct Connect config. For Dimension, this would refer to the config at Upstream Intranet interface to Tor. Only the intranet MTU can be updated, internet mtu and connected VPC mtu are read-only.
	//
	// @param externalConnectivityConfigParam (required)
	// @return com.vmware.model.ExternalConnectivityConfig
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	UpdateIntranetUplinkMtu(externalConnectivityConfigParam model.ExternalConnectivityConfig) (model.ExternalConnectivityConfig, error)

	// This API is used to update the Vmc app data during upgrade.
	//
	// @param upgradeVersionInfoParam (required)
	// @param actionParam Action performed on the Vmc app upgrade (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	UpgradeVMCAppData(upgradeVersionInfoParam model.UpgradeVersionInfo, actionParam string) error
}

type cloudServiceCommonClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewCloudServiceCommonClient(connector client.Connector) *cloudServiceCommonClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.api.cloud_service_common")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"get_external_connectivity_config": core.NewMethodIdentifier(interfaceIdentifier, "get_external_connectivity_config"),
		"get_mgmt_vm_info":                 core.NewMethodIdentifier(interfaceIdentifier, "get_mgmt_vm_info"),
		"get_realized_entities":            core.NewMethodIdentifier(interfaceIdentifier, "get_realized_entities"),
		"get_realized_state_status":        core.NewMethodIdentifier(interfaceIdentifier, "get_realized_state_status"),
		"get_sddc_user_config":             core.NewMethodIdentifier(interfaceIdentifier, "get_sddc_user_config"),
		"get_vmc_feature_flags":            core.NewMethodIdentifier(interfaceIdentifier, "get_vmc_feature_flags"),
		"list_mgmt_vms":                    core.NewMethodIdentifier(interfaceIdentifier, "list_mgmt_vms"),
		"update_intranet_uplink_mtu":       core.NewMethodIdentifier(interfaceIdentifier, "update_intranet_uplink_mtu"),
		"upgrade_VMC_app_data":             core.NewMethodIdentifier(interfaceIdentifier, "upgrade_VMC_app_data"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	cIface := cloudServiceCommonClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &cIface
}

func (cIface *cloudServiceCommonClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := cIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (cIface *cloudServiceCommonClient) GetExternalConnectivityConfig() (model.ExternalConnectivityConfig, error) {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(cloudServiceCommonGetExternalConnectivityConfigInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ExternalConnectivityConfig
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := cloudServiceCommonGetExternalConnectivityConfigRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	cIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.api.cloud_service_common", "get_external_connectivity_config", inputDataValue, executionContext)
	var emptyOutput model.ExternalConnectivityConfig
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), cloudServiceCommonGetExternalConnectivityConfigOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ExternalConnectivityConfig), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (cIface *cloudServiceCommonClient) GetMgmtVmInfo(vmIdParam string) (model.MgmtVmInfo, error) {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(cloudServiceCommonGetMgmtVmInfoInputType(), typeConverter)
	sv.AddStructField("VmId", vmIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.MgmtVmInfo
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := cloudServiceCommonGetMgmtVmInfoRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	cIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.api.cloud_service_common", "get_mgmt_vm_info", inputDataValue, executionContext)
	var emptyOutput model.MgmtVmInfo
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), cloudServiceCommonGetMgmtVmInfoOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.MgmtVmInfo), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (cIface *cloudServiceCommonClient) GetRealizedEntities(intentPathParam string, sitePathParam *string) (model.VmcRealizedEntities, error) {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(cloudServiceCommonGetRealizedEntitiesInputType(), typeConverter)
	sv.AddStructField("IntentPath", intentPathParam)
	sv.AddStructField("SitePath", sitePathParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.VmcRealizedEntities
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := cloudServiceCommonGetRealizedEntitiesRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	cIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.api.cloud_service_common", "get_realized_entities", inputDataValue, executionContext)
	var emptyOutput model.VmcRealizedEntities
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), cloudServiceCommonGetRealizedEntitiesOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.VmcRealizedEntities), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (cIface *cloudServiceCommonClient) GetRealizedStateStatus(intentPathParam string, sitePathParam *string) (model.VmcConsolidatedRealizedStatus, error) {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(cloudServiceCommonGetRealizedStateStatusInputType(), typeConverter)
	sv.AddStructField("IntentPath", intentPathParam)
	sv.AddStructField("SitePath", sitePathParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.VmcConsolidatedRealizedStatus
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := cloudServiceCommonGetRealizedStateStatusRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	cIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.api.cloud_service_common", "get_realized_state_status", inputDataValue, executionContext)
	var emptyOutput model.VmcConsolidatedRealizedStatus
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), cloudServiceCommonGetRealizedStateStatusOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.VmcConsolidatedRealizedStatus), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (cIface *cloudServiceCommonClient) GetSddcUserConfig() (model.SddcUserConfiguration, error) {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(cloudServiceCommonGetSddcUserConfigInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.SddcUserConfiguration
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := cloudServiceCommonGetSddcUserConfigRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	cIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.api.cloud_service_common", "get_sddc_user_config", inputDataValue, executionContext)
	var emptyOutput model.SddcUserConfiguration
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), cloudServiceCommonGetSddcUserConfigOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.SddcUserConfiguration), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (cIface *cloudServiceCommonClient) GetVmcFeatureFlags(internalNameParam *string, nameParam *string) (model.VmcFeatureFlags, error) {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(cloudServiceCommonGetVmcFeatureFlagsInputType(), typeConverter)
	sv.AddStructField("InternalName", internalNameParam)
	sv.AddStructField("Name", nameParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.VmcFeatureFlags
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := cloudServiceCommonGetVmcFeatureFlagsRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	cIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.api.cloud_service_common", "get_vmc_feature_flags", inputDataValue, executionContext)
	var emptyOutput model.VmcFeatureFlags
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), cloudServiceCommonGetVmcFeatureFlagsOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.VmcFeatureFlags), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (cIface *cloudServiceCommonClient) ListMgmtVms() (model.MgmtVmsListResult, error) {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(cloudServiceCommonListMgmtVmsInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.MgmtVmsListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := cloudServiceCommonListMgmtVmsRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	cIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.api.cloud_service_common", "list_mgmt_vms", inputDataValue, executionContext)
	var emptyOutput model.MgmtVmsListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), cloudServiceCommonListMgmtVmsOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.MgmtVmsListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (cIface *cloudServiceCommonClient) UpdateIntranetUplinkMtu(externalConnectivityConfigParam model.ExternalConnectivityConfig) (model.ExternalConnectivityConfig, error) {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(cloudServiceCommonUpdateIntranetUplinkMtuInputType(), typeConverter)
	sv.AddStructField("ExternalConnectivityConfig", externalConnectivityConfigParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ExternalConnectivityConfig
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := cloudServiceCommonUpdateIntranetUplinkMtuRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	cIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.api.cloud_service_common", "update_intranet_uplink_mtu", inputDataValue, executionContext)
	var emptyOutput model.ExternalConnectivityConfig
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), cloudServiceCommonUpdateIntranetUplinkMtuOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ExternalConnectivityConfig), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (cIface *cloudServiceCommonClient) UpgradeVMCAppData(upgradeVersionInfoParam model.UpgradeVersionInfo, actionParam string) error {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(cloudServiceCommonUpgradeVMCAppDataInputType(), typeConverter)
	sv.AddStructField("UpgradeVersionInfo", upgradeVersionInfoParam)
	sv.AddStructField("Action", actionParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := cloudServiceCommonUpgradeVMCAppDataRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	cIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.api.cloud_service_common", "upgrade_VMC_app_data", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}
