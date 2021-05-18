// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: SystemAdministrationLifecycleManagementUpgradeBundles
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

type SystemAdministrationLifecycleManagementUpgradeBundlesClient interface {

	// Cancel upload of upgrade bundle. This API works only when bundle upload is in-progress and will not work during post-processing of upgrade bundle. If bundle upload is in-progress, then the API call returns http OK response after cancelling the upload and deleting partially uploaded bundle.
	//
	// @param bundleIdParam (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	CancelUpgradeBundleUploadCancelUpload(bundleIdParam string) error

	//
	//
	// @param upgradeBundleFetchRequestParam (required)
	// @return com.vmware.model.UpgradeBundleId
	// @throws Unauthenticated  Unauthorized
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	FetchUpgradeBundleFromUrl(upgradeBundleFetchRequestParam model.UpgradeBundleFetchRequest) (model.UpgradeBundleId, error)

	// Get uploaded upgrade bundle information
	//
	// @param bundleIdParam (required)
	// @return com.vmware.model.UpgradeBundleInfo
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetUpgradeBundleInfo(bundleIdParam string) (model.UpgradeBundleInfo, error)

	// Get uploaded upgrade bundle upload status
	//
	// @param bundleIdParam (required)
	// @return com.vmware.model.UpgradeBundleUploadStatus
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetUpgradeBundleUploadStatus(bundleIdParam string) (model.UpgradeBundleUploadStatus, error)

	// Upgrade the upgrade coordinator module itself. This call is invoked after user uploads an upgrade bundle. Once this call is invoked, upgrade coordinator stops and gets restarted and target version upgrade coordinator module comes up after restart.
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	TriggerUcUpgradeUpgradeUc() error
}

type systemAdministrationLifecycleManagementUpgradeBundlesClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewSystemAdministrationLifecycleManagementUpgradeBundlesClient(connector client.Connector) *systemAdministrationLifecycleManagementUpgradeBundlesClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.api.system_administration_lifecycle_management_upgrade_bundles")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"cancel_upgrade_bundle_upload_cancel_upload": core.NewMethodIdentifier(interfaceIdentifier, "cancel_upgrade_bundle_upload_cancel_upload"),
		"fetch_upgrade_bundle_from_url":              core.NewMethodIdentifier(interfaceIdentifier, "fetch_upgrade_bundle_from_url"),
		"get_upgrade_bundle_info":                    core.NewMethodIdentifier(interfaceIdentifier, "get_upgrade_bundle_info"),
		"get_upgrade_bundle_upload_status":           core.NewMethodIdentifier(interfaceIdentifier, "get_upgrade_bundle_upload_status"),
		"trigger_uc_upgrade_upgrade_uc":              core.NewMethodIdentifier(interfaceIdentifier, "trigger_uc_upgrade_upgrade_uc"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	sIface := systemAdministrationLifecycleManagementUpgradeBundlesClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &sIface
}

func (sIface *systemAdministrationLifecycleManagementUpgradeBundlesClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := sIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (sIface *systemAdministrationLifecycleManagementUpgradeBundlesClient) CancelUpgradeBundleUploadCancelUpload(bundleIdParam string) error {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationLifecycleManagementUpgradeBundlesCancelUpgradeBundleUploadCancelUploadInputType(), typeConverter)
	sv.AddStructField("BundleId", bundleIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationLifecycleManagementUpgradeBundlesCancelUpgradeBundleUploadCancelUploadRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_lifecycle_management_upgrade_bundles", "cancel_upgrade_bundle_upload_cancel_upload", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (sIface *systemAdministrationLifecycleManagementUpgradeBundlesClient) FetchUpgradeBundleFromUrl(upgradeBundleFetchRequestParam model.UpgradeBundleFetchRequest) (model.UpgradeBundleId, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationLifecycleManagementUpgradeBundlesFetchUpgradeBundleFromUrlInputType(), typeConverter)
	sv.AddStructField("UpgradeBundleFetchRequest", upgradeBundleFetchRequestParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.UpgradeBundleId
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationLifecycleManagementUpgradeBundlesFetchUpgradeBundleFromUrlRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_lifecycle_management_upgrade_bundles", "fetch_upgrade_bundle_from_url", inputDataValue, executionContext)
	var emptyOutput model.UpgradeBundleId
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationLifecycleManagementUpgradeBundlesFetchUpgradeBundleFromUrlOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.UpgradeBundleId), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationLifecycleManagementUpgradeBundlesClient) GetUpgradeBundleInfo(bundleIdParam string) (model.UpgradeBundleInfo, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationLifecycleManagementUpgradeBundlesGetUpgradeBundleInfoInputType(), typeConverter)
	sv.AddStructField("BundleId", bundleIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.UpgradeBundleInfo
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationLifecycleManagementUpgradeBundlesGetUpgradeBundleInfoRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_lifecycle_management_upgrade_bundles", "get_upgrade_bundle_info", inputDataValue, executionContext)
	var emptyOutput model.UpgradeBundleInfo
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationLifecycleManagementUpgradeBundlesGetUpgradeBundleInfoOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.UpgradeBundleInfo), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationLifecycleManagementUpgradeBundlesClient) GetUpgradeBundleUploadStatus(bundleIdParam string) (model.UpgradeBundleUploadStatus, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationLifecycleManagementUpgradeBundlesGetUpgradeBundleUploadStatusInputType(), typeConverter)
	sv.AddStructField("BundleId", bundleIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.UpgradeBundleUploadStatus
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationLifecycleManagementUpgradeBundlesGetUpgradeBundleUploadStatusRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_lifecycle_management_upgrade_bundles", "get_upgrade_bundle_upload_status", inputDataValue, executionContext)
	var emptyOutput model.UpgradeBundleUploadStatus
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationLifecycleManagementUpgradeBundlesGetUpgradeBundleUploadStatusOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.UpgradeBundleUploadStatus), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationLifecycleManagementUpgradeBundlesClient) TriggerUcUpgradeUpgradeUc() error {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationLifecycleManagementUpgradeBundlesTriggerUcUpgradeUpgradeUcInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationLifecycleManagementUpgradeBundlesTriggerUcUpgradeUpgradeUcRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_lifecycle_management_upgrade_bundles", "trigger_uc_upgrade_upgrade_uc", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}
