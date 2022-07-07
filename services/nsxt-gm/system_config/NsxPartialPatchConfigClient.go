// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: NsxPartialPatchConfig
// Used by client-side stubs.

package system_config

import (
	"github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-gm/model"
)

const _ = core.SupportedByRuntimeVersion1

type NsxPartialPatchConfigClient interface {

	// Get Configuration values for nsx-partial-patch. By default partial patch is disabled (i.e false).
	// @return com.vmware.nsx_global_policy.model.PartialPatchConfig
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get() (model.PartialPatchConfig, error)

	// Update partial patch configuration values. Only boolean value is allowed for enable_partial_patch
	//
	// @param partialPatchConfigParam (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Patch(partialPatchConfigParam model.PartialPatchConfig) error
}

type nsxPartialPatchConfigClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewNsxPartialPatchConfigClient(connector client.Connector) *nsxPartialPatchConfigClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.nsx_global_policy.system_config.nsx_partial_patch_config")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"get":   core.NewMethodIdentifier(interfaceIdentifier, "get"),
		"patch": core.NewMethodIdentifier(interfaceIdentifier, "patch"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	nIface := nsxPartialPatchConfigClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &nIface
}

func (nIface *nsxPartialPatchConfigClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := nIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (nIface *nsxPartialPatchConfigClient) Get() (model.PartialPatchConfig, error) {
	typeConverter := nIface.connector.TypeConverter()
	executionContext := nIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(nsxPartialPatchConfigGetInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.PartialPatchConfig
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := nsxPartialPatchConfigGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	nIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := nIface.connector.GetApiProvider().Invoke("com.vmware.nsx_global_policy.system_config.nsx_partial_patch_config", "get", inputDataValue, executionContext)
	var emptyOutput model.PartialPatchConfig
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), nsxPartialPatchConfigGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.PartialPatchConfig), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), nIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (nIface *nsxPartialPatchConfigClient) Patch(partialPatchConfigParam model.PartialPatchConfig) error {
	typeConverter := nIface.connector.TypeConverter()
	executionContext := nIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(nsxPartialPatchConfigPatchInputType(), typeConverter)
	sv.AddStructField("PartialPatchConfig", partialPatchConfigParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := nsxPartialPatchConfigPatchRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	nIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := nIface.connector.GetApiProvider().Invoke("com.vmware.nsx_global_policy.system_config.nsx_partial_patch_config", "patch", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), nIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}
