// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: ApiCertificate
// Used by client-side stubs.

package cluster

import (
	"github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
)

const _ = core.SupportedByRuntimeVersion1

type ApiCertificateClient interface {

	// Clears the certificate used for the MP cluster. This does not affect the certificate itself. This API is deprecated. Instead use the /api/v1/cluster/api-certificate?action=set_cluster_certificate API to set the cluster certificate to a different one. It just means that from now on, individual certificates will be used on each MP node. This affects all nodes in the cluster.
	//
	// @param certificateIdParam Certificate ID (required)
	// @return com.vmware.nsx.model.ClusterCertificateId
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Clearclustercertificate(certificateIdParam string) (model.ClusterCertificateId, error)

	// Returns the ID of the certificate that is used as the cluster certificate for MP
	// @return com.vmware.nsx.model.ClusterCertificateId
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get() (model.ClusterCertificateId, error)

	// Sets the certificate used for the MP cluster. This affects all nodes in the cluster. If the certificate used is a CA signed certificate,the request fails if the whole chain(leaf, intermediate, root) is not imported.
	//
	// @param certificateIdParam Certificate ID (required)
	// @return com.vmware.nsx.model.ClusterCertificateId
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Setclustercertificate(certificateIdParam string) (model.ClusterCertificateId, error)
}

type apiCertificateClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewApiCertificateClient(connector client.Connector) *apiCertificateClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.nsx.cluster.api_certificate")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"clearclustercertificate": core.NewMethodIdentifier(interfaceIdentifier, "clearclustercertificate"),
		"get":                     core.NewMethodIdentifier(interfaceIdentifier, "get"),
		"setclustercertificate":   core.NewMethodIdentifier(interfaceIdentifier, "setclustercertificate"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	aIface := apiCertificateClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &aIface
}

func (aIface *apiCertificateClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := aIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (aIface *apiCertificateClient) Clearclustercertificate(certificateIdParam string) (model.ClusterCertificateId, error) {
	typeConverter := aIface.connector.TypeConverter()
	executionContext := aIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(apiCertificateClearclustercertificateInputType(), typeConverter)
	sv.AddStructField("CertificateId", certificateIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ClusterCertificateId
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := apiCertificateClearclustercertificateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	aIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := aIface.connector.GetApiProvider().Invoke("com.vmware.nsx.cluster.api_certificate", "clearclustercertificate", inputDataValue, executionContext)
	var emptyOutput model.ClusterCertificateId
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), apiCertificateClearclustercertificateOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ClusterCertificateId), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), aIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (aIface *apiCertificateClient) Get() (model.ClusterCertificateId, error) {
	typeConverter := aIface.connector.TypeConverter()
	executionContext := aIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(apiCertificateGetInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ClusterCertificateId
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := apiCertificateGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	aIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := aIface.connector.GetApiProvider().Invoke("com.vmware.nsx.cluster.api_certificate", "get", inputDataValue, executionContext)
	var emptyOutput model.ClusterCertificateId
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), apiCertificateGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ClusterCertificateId), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), aIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (aIface *apiCertificateClient) Setclustercertificate(certificateIdParam string) (model.ClusterCertificateId, error) {
	typeConverter := aIface.connector.TypeConverter()
	executionContext := aIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(apiCertificateSetclustercertificateInputType(), typeConverter)
	sv.AddStructField("CertificateId", certificateIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ClusterCertificateId
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := apiCertificateSetclustercertificateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	aIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := aIface.connector.GetApiProvider().Invoke("com.vmware.nsx.cluster.api_certificate", "setclustercertificate", inputDataValue, executionContext)
	var emptyOutput model.ClusterCertificateId
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), apiCertificateSetclustercertificateOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ClusterCertificateId), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), aIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
