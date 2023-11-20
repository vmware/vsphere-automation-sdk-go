// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: SupportBundles
// Used by client-side stubs.

package administration

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	nsxModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type SupportBundlesClient interface {

	// Collect support bundles from registered cluster and fabric nodes.
	//
	// @param supportBundleRequestParam (required)
	// @param overrideAsyncResponseParam Override any existing support bundle async response (optional, default to false)
	// @param requireDeleteOrOverrideAsyncResponseParam Suppress auto-deletion of generated support bundle (optional, default to false)
	// @return com.vmware.nsx.model.SupportBundleResult
	//
	// @throws ConcurrentChange  Conflict
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error, Bad Gateway
	// @throws NotFound  Not Found
	Collect(supportBundleRequestParam nsxModel.SupportBundleRequest, overrideAsyncResponseParam *bool, requireDeleteOrOverrideAsyncResponseParam *bool) (nsxModel.SupportBundleResult, error)

	// Delete existing support bundles waiting to be downloaded.
	//
	// @throws ConcurrentChange  Conflict
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Deleteasyncresponse() error
}

type supportBundlesClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewSupportBundlesClient(connector vapiProtocolClient_.Connector) *supportBundlesClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx.administration.support_bundles")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"collect":             vapiCore_.NewMethodIdentifier(interfaceIdentifier, "collect"),
		"deleteasyncresponse": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "deleteasyncresponse"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	sIface := supportBundlesClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &sIface
}

func (sIface *supportBundlesClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := sIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (sIface *supportBundlesClient) Collect(supportBundleRequestParam nsxModel.SupportBundleRequest, overrideAsyncResponseParam *bool, requireDeleteOrOverrideAsyncResponseParam *bool) (nsxModel.SupportBundleResult, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	operationRestMetaData := supportBundlesCollectRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(supportBundlesCollectInputType(), typeConverter)
	sv.AddStructField("SupportBundleRequest", supportBundleRequestParam)
	sv.AddStructField("OverrideAsyncResponse", overrideAsyncResponseParam)
	sv.AddStructField("RequireDeleteOrOverrideAsyncResponse", requireDeleteOrOverrideAsyncResponseParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.SupportBundleResult
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx.administration.support_bundles", "collect", inputDataValue, executionContext)
	var emptyOutput nsxModel.SupportBundleResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), SupportBundlesCollectOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.SupportBundleResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *supportBundlesClient) Deleteasyncresponse() error {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	operationRestMetaData := supportBundlesDeleteasyncresponseRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(supportBundlesDeleteasyncresponseInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx.administration.support_bundles", "deleteasyncresponse", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}
