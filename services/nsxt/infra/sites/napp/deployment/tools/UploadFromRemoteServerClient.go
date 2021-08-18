// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Upload_from_remote_server
// Used by client-side stubs.

package tools

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/core"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/lib"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/protocol/client"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt/model"
)

const _ = core.SupportedByRuntimeVersion1

type Upload_from_remote_serverClient interface {

	// Upload the Kubernetes tools on NSX Manager via remote url. You need to get this file from VMware website. File contains kubectl binary, helm binary and signature to verify them.
	//
	// @param siteIdParam (required)
	// @param uploadBundleRemoteServerParam (required)
	// @return com.vmware.nsx_policy.model.UploadBundleId
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Create(siteIdParam string, uploadBundleRemoteServerParam model.UploadBundleRemoteServer) (model.UploadBundleId, error)
}

type upload_from_remote_serverClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewUpload_from_remote_serverClient(connector client.Connector) *upload_from_remote_serverClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.nsx_policy.infra.sites.napp.deployment.tools.upload_from_remote_server")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"create": core.NewMethodIdentifier(interfaceIdentifier, "create"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	uIface := upload_from_remote_serverClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &uIface
}

func (uIface *upload_from_remote_serverClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := uIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (uIface *upload_from_remote_serverClient) Create(siteIdParam string, uploadBundleRemoteServerParam model.UploadBundleRemoteServer) (model.UploadBundleId, error) {
	typeConverter := uIface.connector.TypeConverter()
	executionContext := uIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(uploadFromRemoteServerCreateInputType(), typeConverter)
	sv.AddStructField("SiteId", siteIdParam)
	sv.AddStructField("UploadBundleRemoteServer", uploadBundleRemoteServerParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.UploadBundleId
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := uploadFromRemoteServerCreateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	uIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := uIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.sites.napp.deployment.tools.upload_from_remote_server", "create", inputDataValue, executionContext)
	var emptyOutput model.UploadBundleId
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), uploadFromRemoteServerCreateOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.UploadBundleId), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), uIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
