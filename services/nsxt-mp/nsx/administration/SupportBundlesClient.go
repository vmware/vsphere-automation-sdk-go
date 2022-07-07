// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: SupportBundles
// Used by client-side stubs.

package administration

import (
	"github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
)

const _ = core.SupportedByRuntimeVersion1

type SupportBundlesClient interface {

	// Collect support bundles from registered cluster and fabric nodes.
	//
	// @param supportBundleRequestParam (required)
	// @param overrideAsyncResponseParam Override any existing support bundle async response (optional, default to false)
	// @param requireDeleteOrOverrideAsyncResponseParam Suppress auto-deletion of generated support bundle (optional, default to false)
	// @return com.vmware.nsx.model.SupportBundleResult
	// @throws ConcurrentChange  Conflict
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error, Bad Gateway
	// @throws NotFound  Not Found
	Collect(supportBundleRequestParam model.SupportBundleRequest, overrideAsyncResponseParam *bool, requireDeleteOrOverrideAsyncResponseParam *bool) (model.SupportBundleResult, error)

	// Delete existing support bundles waiting to be downloaded.
	// @throws ConcurrentChange  Conflict
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Deleteasyncresponse() error
}

type supportBundlesClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewSupportBundlesClient(connector client.Connector) *supportBundlesClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.nsx.administration.support_bundles")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"collect":             core.NewMethodIdentifier(interfaceIdentifier, "collect"),
		"deleteasyncresponse": core.NewMethodIdentifier(interfaceIdentifier, "deleteasyncresponse"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	sIface := supportBundlesClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &sIface
}

func (sIface *supportBundlesClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := sIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (sIface *supportBundlesClient) Collect(supportBundleRequestParam model.SupportBundleRequest, overrideAsyncResponseParam *bool, requireDeleteOrOverrideAsyncResponseParam *bool) (model.SupportBundleResult, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(supportBundlesCollectInputType(), typeConverter)
	sv.AddStructField("SupportBundleRequest", supportBundleRequestParam)
	sv.AddStructField("OverrideAsyncResponse", overrideAsyncResponseParam)
	sv.AddStructField("RequireDeleteOrOverrideAsyncResponse", requireDeleteOrOverrideAsyncResponseParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.SupportBundleResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := supportBundlesCollectRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx.administration.support_bundles", "collect", inputDataValue, executionContext)
	var emptyOutput model.SupportBundleResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), supportBundlesCollectOutputType())
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

func (sIface *supportBundlesClient) Deleteasyncresponse() error {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(supportBundlesDeleteasyncresponseInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := supportBundlesDeleteasyncresponseRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx.administration.support_bundles", "deleteasyncresponse", inputDataValue, executionContext)
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
