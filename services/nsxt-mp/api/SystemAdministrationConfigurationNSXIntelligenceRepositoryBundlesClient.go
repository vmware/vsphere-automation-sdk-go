// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: SystemAdministrationConfigurationNSXIntelligenceRepositoryBundles
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

type SystemAdministrationConfigurationNSXIntelligenceRepositoryBundlesClient interface {

	// Cancel upload of bundle. This API works only when bundle upload is in-progress and will not work during post-processing of bundle. If bundle upload is in-progress, then the API call returns http OK response after cancelling the upload and deleting partially uploaded bundle.
	//
	// @param bundleIdParam (required)
	// @param productParam Name of the product (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	CancelBundleUploadCancelUpload(bundleIdParam string, productParam string) error

	// Get list of bundle-ids which are available in repository or in-progress
	//
	// @param fileTypeParam Type of file (required)
	// @param productParam Name of the product (required)
	// @return com.vmware.model.BundleIds
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetBundleIds(fileTypeParam string, productParam string) (model.BundleIds, error)

	// Checks whether bundle upload is allowed on given node for given product. There are different kinds of checks for different products. Some of the checks for Intelligence product are as follows: 1. Is bundle upload-allowed on given node 2. Is bundle upload already in-progress 3. Is Intelliegnce node deployment in-progress 4. Is Intelliegnce node upgrade in-progress
	//
	// @param productParam Name of the product (required)
	// @return com.vmware.model.BundleUploadPermission
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetBundleUploadPermissions(productParam string) (model.BundleUploadPermission, error)

	// Get uploaded bundle upload status
	//
	// @param bundleIdParam (required)
	// @param productParam Name of the product (required)
	// @return com.vmware.model.BundleUploadStatus
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetBundleUploadStatus(bundleIdParam string, productParam string) (model.BundleUploadStatus, error)

	// Get information of the OVF for specified product which is present in repository and will be used to deploy new VM.
	//
	// @param productParam Name of the product (required)
	// @return com.vmware.model.OvfInfo
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetOvfDeployInfo(productParam string) (model.OvfInfo, error)

	//
	//
	// @param remoteBundleUrlParam (required)
	// @param fileTypeParam Type of file (required)
	// @param productParam Name of the product (required)
	// @return com.vmware.model.BundleId
	// @throws Unauthenticated  Unauthorized
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	UploadBundleViaRemoteFile(remoteBundleUrlParam model.RemoteBundleUrl, fileTypeParam string, productParam string) (model.BundleId, error)
}

type systemAdministrationConfigurationNSXIntelligenceRepositoryBundlesClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewSystemAdministrationConfigurationNSXIntelligenceRepositoryBundlesClient(connector client.Connector) *systemAdministrationConfigurationNSXIntelligenceRepositoryBundlesClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.api.system_administration_configuration_NSX_intelligence_repository_bundles")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"cancel_bundle_upload_cancel_upload": core.NewMethodIdentifier(interfaceIdentifier, "cancel_bundle_upload_cancel_upload"),
		"get_bundle_ids":                     core.NewMethodIdentifier(interfaceIdentifier, "get_bundle_ids"),
		"get_bundle_upload_permissions":      core.NewMethodIdentifier(interfaceIdentifier, "get_bundle_upload_permissions"),
		"get_bundle_upload_status":           core.NewMethodIdentifier(interfaceIdentifier, "get_bundle_upload_status"),
		"get_ovf_deploy_info":                core.NewMethodIdentifier(interfaceIdentifier, "get_ovf_deploy_info"),
		"upload_bundle_via_remote_file":      core.NewMethodIdentifier(interfaceIdentifier, "upload_bundle_via_remote_file"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	sIface := systemAdministrationConfigurationNSXIntelligenceRepositoryBundlesClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &sIface
}

func (sIface *systemAdministrationConfigurationNSXIntelligenceRepositoryBundlesClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := sIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (sIface *systemAdministrationConfigurationNSXIntelligenceRepositoryBundlesClient) CancelBundleUploadCancelUpload(bundleIdParam string, productParam string) error {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationNSXIntelligenceRepositoryBundlesCancelBundleUploadCancelUploadInputType(), typeConverter)
	sv.AddStructField("BundleId", bundleIdParam)
	sv.AddStructField("Product", productParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationNSXIntelligenceRepositoryBundlesCancelBundleUploadCancelUploadRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_NSX_intelligence_repository_bundles", "cancel_bundle_upload_cancel_upload", inputDataValue, executionContext)
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

func (sIface *systemAdministrationConfigurationNSXIntelligenceRepositoryBundlesClient) GetBundleIds(fileTypeParam string, productParam string) (model.BundleIds, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationNSXIntelligenceRepositoryBundlesGetBundleIdsInputType(), typeConverter)
	sv.AddStructField("FileType", fileTypeParam)
	sv.AddStructField("Product", productParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.BundleIds
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationNSXIntelligenceRepositoryBundlesGetBundleIdsRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_NSX_intelligence_repository_bundles", "get_bundle_ids", inputDataValue, executionContext)
	var emptyOutput model.BundleIds
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationConfigurationNSXIntelligenceRepositoryBundlesGetBundleIdsOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.BundleIds), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationConfigurationNSXIntelligenceRepositoryBundlesClient) GetBundleUploadPermissions(productParam string) (model.BundleUploadPermission, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationNSXIntelligenceRepositoryBundlesGetBundleUploadPermissionsInputType(), typeConverter)
	sv.AddStructField("Product", productParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.BundleUploadPermission
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationNSXIntelligenceRepositoryBundlesGetBundleUploadPermissionsRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_NSX_intelligence_repository_bundles", "get_bundle_upload_permissions", inputDataValue, executionContext)
	var emptyOutput model.BundleUploadPermission
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationConfigurationNSXIntelligenceRepositoryBundlesGetBundleUploadPermissionsOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.BundleUploadPermission), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationConfigurationNSXIntelligenceRepositoryBundlesClient) GetBundleUploadStatus(bundleIdParam string, productParam string) (model.BundleUploadStatus, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationNSXIntelligenceRepositoryBundlesGetBundleUploadStatusInputType(), typeConverter)
	sv.AddStructField("BundleId", bundleIdParam)
	sv.AddStructField("Product", productParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.BundleUploadStatus
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationNSXIntelligenceRepositoryBundlesGetBundleUploadStatusRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_NSX_intelligence_repository_bundles", "get_bundle_upload_status", inputDataValue, executionContext)
	var emptyOutput model.BundleUploadStatus
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationConfigurationNSXIntelligenceRepositoryBundlesGetBundleUploadStatusOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.BundleUploadStatus), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationConfigurationNSXIntelligenceRepositoryBundlesClient) GetOvfDeployInfo(productParam string) (model.OvfInfo, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationNSXIntelligenceRepositoryBundlesGetOvfDeployInfoInputType(), typeConverter)
	sv.AddStructField("Product", productParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.OvfInfo
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationNSXIntelligenceRepositoryBundlesGetOvfDeployInfoRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_NSX_intelligence_repository_bundles", "get_ovf_deploy_info", inputDataValue, executionContext)
	var emptyOutput model.OvfInfo
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationConfigurationNSXIntelligenceRepositoryBundlesGetOvfDeployInfoOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.OvfInfo), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationConfigurationNSXIntelligenceRepositoryBundlesClient) UploadBundleViaRemoteFile(remoteBundleUrlParam model.RemoteBundleUrl, fileTypeParam string, productParam string) (model.BundleId, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationNSXIntelligenceRepositoryBundlesUploadBundleViaRemoteFileInputType(), typeConverter)
	sv.AddStructField("RemoteBundleUrl", remoteBundleUrlParam)
	sv.AddStructField("FileType", fileTypeParam)
	sv.AddStructField("Product", productParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.BundleId
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationNSXIntelligenceRepositoryBundlesUploadBundleViaRemoteFileRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_NSX_intelligence_repository_bundles", "upload_bundle_via_remote_file", inputDataValue, executionContext)
	var emptyOutput model.BundleId
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationConfigurationNSXIntelligenceRepositoryBundlesUploadBundleViaRemoteFileOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.BundleId), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
