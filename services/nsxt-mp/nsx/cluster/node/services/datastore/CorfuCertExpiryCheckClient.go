// Copyright (c) 2019-2026 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Corfu_cert_expiry_check
// Used by client-side stubs.

package datastore

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	nsxModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type Corfu_cert_expiry_checkClient interface {

	// Enable or Disable Corfu Certificate Expiry Check. Default is enabled
	//
	// @param clusterNodeIdParam (required)
	// @return com.vmware.nsx.model.CorfuCertificateExpiryCheckProperties
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Disable(clusterNodeIdParam string) (nsxModel.CorfuCertificateExpiryCheckProperties, error)

	// Enable or Disable Corfu Certificate Expiry Check. Default is enabled
	//
	// @param clusterNodeIdParam (required)
	// @return com.vmware.nsx.model.CorfuCertificateExpiryCheckProperties
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Enable(clusterNodeIdParam string) (nsxModel.CorfuCertificateExpiryCheckProperties, error)

	// Get the status of Corfu Certificate Expiry Check. Enabled or disabled
	//
	// @param clusterNodeIdParam (required)
	// @return com.vmware.nsx.model.CorfuCertificateExpiryCheckProperties
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(clusterNodeIdParam string) (nsxModel.CorfuCertificateExpiryCheckProperties, error)
}

type corfu_cert_expiry_checkClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewCorfu_cert_expiry_checkClient(connector vapiProtocolClient_.Connector) *corfu_cert_expiry_checkClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx.cluster.node.services.datastore.corfu_cert_expiry_check")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"disable": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "disable"),
		"enable":  vapiCore_.NewMethodIdentifier(interfaceIdentifier, "enable"),
		"get":     vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	cIface := corfu_cert_expiry_checkClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &cIface
}

func (cIface *corfu_cert_expiry_checkClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := cIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (cIface *corfu_cert_expiry_checkClient) Disable(clusterNodeIdParam string) (nsxModel.CorfuCertificateExpiryCheckProperties, error) {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	operationRestMetaData := corfuCertExpiryCheckDisableRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(corfuCertExpiryCheckDisableInputType(), typeConverter)
	sv.AddStructField("ClusterNodeId", clusterNodeIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.CorfuCertificateExpiryCheckProperties
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.nsx.cluster.node.services.datastore.corfu_cert_expiry_check", "disable", inputDataValue, executionContext)
	var emptyOutput nsxModel.CorfuCertificateExpiryCheckProperties
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), CorfuCertExpiryCheckDisableOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.CorfuCertificateExpiryCheckProperties), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (cIface *corfu_cert_expiry_checkClient) Enable(clusterNodeIdParam string) (nsxModel.CorfuCertificateExpiryCheckProperties, error) {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	operationRestMetaData := corfuCertExpiryCheckEnableRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(corfuCertExpiryCheckEnableInputType(), typeConverter)
	sv.AddStructField("ClusterNodeId", clusterNodeIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.CorfuCertificateExpiryCheckProperties
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.nsx.cluster.node.services.datastore.corfu_cert_expiry_check", "enable", inputDataValue, executionContext)
	var emptyOutput nsxModel.CorfuCertificateExpiryCheckProperties
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), CorfuCertExpiryCheckEnableOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.CorfuCertificateExpiryCheckProperties), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (cIface *corfu_cert_expiry_checkClient) Get(clusterNodeIdParam string) (nsxModel.CorfuCertificateExpiryCheckProperties, error) {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	operationRestMetaData := corfuCertExpiryCheckGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(corfuCertExpiryCheckGetInputType(), typeConverter)
	sv.AddStructField("ClusterNodeId", clusterNodeIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.CorfuCertificateExpiryCheckProperties
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.nsx.cluster.node.services.datastore.corfu_cert_expiry_check", "get", inputDataValue, executionContext)
	var emptyOutput nsxModel.CorfuCertificateExpiryCheckProperties
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), CorfuCertExpiryCheckGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.CorfuCertificateExpiryCheckProperties), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
