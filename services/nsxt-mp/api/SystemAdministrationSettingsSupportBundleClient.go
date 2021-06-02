// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: SystemAdministrationSettingsSupportBundle
// Used by client-side stubs.

package api

import (
	"github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/model"
)

const _ = core.SupportedByRuntimeVersion1

type SystemAdministrationSettingsSupportBundleClient interface {

	// Collect support bundles from registered cluster and fabric nodes.
	//
	// @param supportBundleRequestParam (required)
	// @param overrideAsyncResponseParam Override any existing support bundle async response (optional, default to false)
	// @param requireDeleteOrOverrideAsyncResponseParam Suppress auto-deletion of generated support bundle (optional, default to false)
	// @return com.vmware.model.SupportBundleResult
	// @throws ConcurrentChange  Conflict
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error, Bad Gateway
	// @throws NotFound  Not Found
	CollectSupportBundlesCollect(supportBundleRequestParam model.SupportBundleRequest, overrideAsyncResponseParam *bool, requireDeleteOrOverrideAsyncResponseParam *bool) (model.SupportBundleResult, error)

	// Delete existing support bundles waiting to be downloaded.
	// @throws ConcurrentChange  Conflict
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	DeleteSupportBundlesAsyncResponseDeleteAsyncResponse() error
}

type systemAdministrationSettingsSupportBundleClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewSystemAdministrationSettingsSupportBundleClient(connector client.Connector) *systemAdministrationSettingsSupportBundleClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.api.system_administration_settings_support_bundle")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"collect_support_bundles_collect":                             core.NewMethodIdentifier(interfaceIdentifier, "collect_support_bundles_collect"),
		"delete_support_bundles_async_response_delete_async_response": core.NewMethodIdentifier(interfaceIdentifier, "delete_support_bundles_async_response_delete_async_response"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	sIface := systemAdministrationSettingsSupportBundleClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &sIface
}

func (sIface *systemAdministrationSettingsSupportBundleClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := sIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (sIface *systemAdministrationSettingsSupportBundleClient) CollectSupportBundlesCollect(supportBundleRequestParam model.SupportBundleRequest, overrideAsyncResponseParam *bool, requireDeleteOrOverrideAsyncResponseParam *bool) (model.SupportBundleResult, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationSettingsSupportBundleCollectSupportBundlesCollectInputType(), typeConverter)
	sv.AddStructField("SupportBundleRequest", supportBundleRequestParam)
	sv.AddStructField("OverrideAsyncResponse", overrideAsyncResponseParam)
	sv.AddStructField("RequireDeleteOrOverrideAsyncResponse", requireDeleteOrOverrideAsyncResponseParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.SupportBundleResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationSettingsSupportBundleCollectSupportBundlesCollectRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_settings_support_bundle", "collect_support_bundles_collect", inputDataValue, executionContext)
	var emptyOutput model.SupportBundleResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationSettingsSupportBundleCollectSupportBundlesCollectOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.SupportBundleResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationSettingsSupportBundleClient) DeleteSupportBundlesAsyncResponseDeleteAsyncResponse() error {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationSettingsSupportBundleDeleteSupportBundlesAsyncResponseDeleteAsyncResponseInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationSettingsSupportBundleDeleteSupportBundlesAsyncResponseDeleteAsyncResponseRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_settings_support_bundle", "delete_support_bundles_async_response_delete_async_response", inputDataValue, executionContext)
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
