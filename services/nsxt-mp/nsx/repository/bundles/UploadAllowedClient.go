// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: UploadAllowed
// Used by client-side stubs.

package bundles

import (
	"github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
)

const _ = core.SupportedByRuntimeVersion1

type UploadAllowedClient interface {

	// Checks whether bundle upload is allowed on given node for given appliance. There are different kinds of checks for different appliances. Some of the checks for Intelligence appliance are as follows: 1. Is bundle upload-allowed on given node 2. Is bundle upload already in-progress
	//
	// @param productParam Name of the appliance (required)
	// @return com.vmware.nsx.model.BundleUploadPermission
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(productParam string) (model.BundleUploadPermission, error)
}

type uploadAllowedClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewUploadAllowedClient(connector client.Connector) *uploadAllowedClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.nsx.repository.bundles.upload_allowed")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"get": core.NewMethodIdentifier(interfaceIdentifier, "get"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	uIface := uploadAllowedClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &uIface
}

func (uIface *uploadAllowedClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := uIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (uIface *uploadAllowedClient) Get(productParam string) (model.BundleUploadPermission, error) {
	typeConverter := uIface.connector.TypeConverter()
	executionContext := uIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(uploadAllowedGetInputType(), typeConverter)
	sv.AddStructField("Product", productParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.BundleUploadPermission
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := uploadAllowedGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	uIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := uIface.connector.GetApiProvider().Invoke("com.vmware.nsx.repository.bundles.upload_allowed", "get", inputDataValue, executionContext)
	var emptyOutput model.BundleUploadPermission
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), uploadAllowedGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.BundleUploadPermission), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), uIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
