// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: KubernetesTools
// Used by client-side stubs.

package deployment

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/core"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/lib"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/protocol/client"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt/model"
)

const _ = core.SupportedByRuntimeVersion1

type KubernetesToolsClient interface {

	// To show information about Kubernetes tools version. It will list kubectl client version, kube-apiserver version. Also show, whether kubectl client version is compatible with kube-apiserver version.kubectl is supported within one minor version (older or newer) of kube-apiserver.
	//
	// @param siteIdParam (required)
	// @return com.vmware.nsx_policy.model.KubernetesToolsInfo
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(siteIdParam string) (model.KubernetesToolsInfo, error)
}

type kubernetesToolsClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewKubernetesToolsClient(connector client.Connector) *kubernetesToolsClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.nsx_policy.infra.sites.napp.deployment.kubernetes_tools")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"get": core.NewMethodIdentifier(interfaceIdentifier, "get"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	kIface := kubernetesToolsClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &kIface
}

func (kIface *kubernetesToolsClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := kIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (kIface *kubernetesToolsClient) Get(siteIdParam string) (model.KubernetesToolsInfo, error) {
	typeConverter := kIface.connector.TypeConverter()
	executionContext := kIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(kubernetesToolsGetInputType(), typeConverter)
	sv.AddStructField("SiteId", siteIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.KubernetesToolsInfo
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := kubernetesToolsGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	kIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := kIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.sites.napp.deployment.kubernetes_tools", "get", inputDataValue, executionContext)
	var emptyOutput model.KubernetesToolsInfo
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), kubernetesToolsGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.KubernetesToolsInfo), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), kIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
