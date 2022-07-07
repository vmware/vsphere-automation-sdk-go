// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Http
// Used by client-side stubs.

package services

import (
	"github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
)

const _ = core.SupportedByRuntimeVersion1

type HttpClient interface {

	// Applies a security certificate to the http service. In the POST request, the CERTIFICATE_ID references a certificate created with the /api/v1/trust-management APIs. If the certificate used is a CA signed certificate,the request fails if the whole chain(leaf, intermediate, root) is not imported.
	//
	// @param certificateIdParam Certificate ID (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Applycertificate(certificateIdParam string) error

	// This API is deprecated. Read the configuration of the http service by calling the GET /api/v1/cluster/api-service API.
	// @return com.vmware.nsx.model.NodeHttpServiceProperties
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get() (model.NodeHttpServiceProperties, error)

	// Restart the http service
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Restart() error

	// Start the http service
	// @return com.vmware.nsx.model.NodeServiceStatusProperties
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Start() (model.NodeServiceStatusProperties, error)

	// Stop the http service
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Stop() error

	// This API is deprecated. Make changes to the http service configuration by calling the PUT /api/v1/cluster/api-service API.
	//
	// @param nodeHttpServicePropertiesParam (required)
	// @return com.vmware.nsx.model.NodeHttpServiceProperties
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(nodeHttpServicePropertiesParam model.NodeHttpServiceProperties) (model.NodeHttpServiceProperties, error)
}

type httpClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewHttpClient(connector client.Connector) *httpClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.nsx.node.services.http")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"applycertificate": core.NewMethodIdentifier(interfaceIdentifier, "applycertificate"),
		"get":              core.NewMethodIdentifier(interfaceIdentifier, "get"),
		"restart":          core.NewMethodIdentifier(interfaceIdentifier, "restart"),
		"start":            core.NewMethodIdentifier(interfaceIdentifier, "start"),
		"stop":             core.NewMethodIdentifier(interfaceIdentifier, "stop"),
		"update":           core.NewMethodIdentifier(interfaceIdentifier, "update"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	hIface := httpClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &hIface
}

func (hIface *httpClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := hIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (hIface *httpClient) Applycertificate(certificateIdParam string) error {
	typeConverter := hIface.connector.TypeConverter()
	executionContext := hIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(httpApplycertificateInputType(), typeConverter)
	sv.AddStructField("CertificateId", certificateIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := httpApplycertificateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	hIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := hIface.connector.GetApiProvider().Invoke("com.vmware.nsx.node.services.http", "applycertificate", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), hIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (hIface *httpClient) Get() (model.NodeHttpServiceProperties, error) {
	typeConverter := hIface.connector.TypeConverter()
	executionContext := hIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(httpGetInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.NodeHttpServiceProperties
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := httpGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	hIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := hIface.connector.GetApiProvider().Invoke("com.vmware.nsx.node.services.http", "get", inputDataValue, executionContext)
	var emptyOutput model.NodeHttpServiceProperties
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), httpGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.NodeHttpServiceProperties), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), hIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (hIface *httpClient) Restart() error {
	typeConverter := hIface.connector.TypeConverter()
	executionContext := hIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(httpRestartInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := httpRestartRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	hIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := hIface.connector.GetApiProvider().Invoke("com.vmware.nsx.node.services.http", "restart", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), hIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (hIface *httpClient) Start() (model.NodeServiceStatusProperties, error) {
	typeConverter := hIface.connector.TypeConverter()
	executionContext := hIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(httpStartInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.NodeServiceStatusProperties
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := httpStartRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	hIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := hIface.connector.GetApiProvider().Invoke("com.vmware.nsx.node.services.http", "start", inputDataValue, executionContext)
	var emptyOutput model.NodeServiceStatusProperties
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), httpStartOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.NodeServiceStatusProperties), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), hIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (hIface *httpClient) Stop() error {
	typeConverter := hIface.connector.TypeConverter()
	executionContext := hIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(httpStopInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := httpStopRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	hIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := hIface.connector.GetApiProvider().Invoke("com.vmware.nsx.node.services.http", "stop", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), hIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (hIface *httpClient) Update(nodeHttpServicePropertiesParam model.NodeHttpServiceProperties) (model.NodeHttpServiceProperties, error) {
	typeConverter := hIface.connector.TypeConverter()
	executionContext := hIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(httpUpdateInputType(), typeConverter)
	sv.AddStructField("NodeHttpServiceProperties", nodeHttpServicePropertiesParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.NodeHttpServiceProperties
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := httpUpdateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	hIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := hIface.connector.GetApiProvider().Invoke("com.vmware.nsx.node.services.http", "update", inputDataValue, executionContext)
	var emptyOutput model.NodeHttpServiceProperties
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), httpUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.NodeHttpServiceProperties), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), hIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
