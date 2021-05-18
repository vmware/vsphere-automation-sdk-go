// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: CloudServiceVMCOnAWSPublicIP
// Used by client-side stubs.

package api

import (
	"github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-vmc-aws-integration/model"
)

const _ = core.SupportedByRuntimeVersion1

type CloudServiceVMCOnAWSPublicIPClient interface {

	// This API is used to create or update a public IP. In creating, the API allocates a new public IP from VMC provider. In updating, only the display name can be modified, the IP is read-only.
	//
	// @param publicIpIdParam (required)
	// @param publicIpParam (required)
	// @return com.vmware.model.PublicIp
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	CreatePublicIp(publicIpIdParam string, publicIpParam model.PublicIp) (model.PublicIp, error)

	// Delete a public IP. The IP will be released in VMC provider.
	//
	// @param publicIpIdParam (required)
	// @param forceParam Force delete the resource even if it is being used somewhere (optional, default to false)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	DeletePublicIp(publicIpIdParam string, forceParam *bool) error

	// Get public IP information associated with given ID.
	//
	// @param publicIpIdParam (required)
	// @return com.vmware.model.PublicIp
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetPublicIp(publicIpIdParam string) (model.PublicIp, error)

	// List all public IPs obtained in the SDDC.
	//
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @return com.vmware.model.PublicIpsListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ListPublicIps(cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.PublicIpsListResult, error)
}

type cloudServiceVMCOnAWSPublicIPClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewCloudServiceVMCOnAWSPublicIPClient(connector client.Connector) *cloudServiceVMCOnAWSPublicIPClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.api.cloud_service_VMC_on_AWS_public_IP")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"create_public_ip": core.NewMethodIdentifier(interfaceIdentifier, "create_public_ip"),
		"delete_public_ip": core.NewMethodIdentifier(interfaceIdentifier, "delete_public_ip"),
		"get_public_ip":    core.NewMethodIdentifier(interfaceIdentifier, "get_public_ip"),
		"list_public_ips":  core.NewMethodIdentifier(interfaceIdentifier, "list_public_ips"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	cIface := cloudServiceVMCOnAWSPublicIPClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &cIface
}

func (cIface *cloudServiceVMCOnAWSPublicIPClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := cIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (cIface *cloudServiceVMCOnAWSPublicIPClient) CreatePublicIp(publicIpIdParam string, publicIpParam model.PublicIp) (model.PublicIp, error) {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(cloudServiceVMCOnAWSPublicIPCreatePublicIpInputType(), typeConverter)
	sv.AddStructField("PublicIpId", publicIpIdParam)
	sv.AddStructField("PublicIp", publicIpParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.PublicIp
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := cloudServiceVMCOnAWSPublicIPCreatePublicIpRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	cIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.api.cloud_service_VMC_on_AWS_public_IP", "create_public_ip", inputDataValue, executionContext)
	var emptyOutput model.PublicIp
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), cloudServiceVMCOnAWSPublicIPCreatePublicIpOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.PublicIp), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (cIface *cloudServiceVMCOnAWSPublicIPClient) DeletePublicIp(publicIpIdParam string, forceParam *bool) error {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(cloudServiceVMCOnAWSPublicIPDeletePublicIpInputType(), typeConverter)
	sv.AddStructField("PublicIpId", publicIpIdParam)
	sv.AddStructField("Force", forceParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := cloudServiceVMCOnAWSPublicIPDeletePublicIpRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	cIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.api.cloud_service_VMC_on_AWS_public_IP", "delete_public_ip", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (cIface *cloudServiceVMCOnAWSPublicIPClient) GetPublicIp(publicIpIdParam string) (model.PublicIp, error) {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(cloudServiceVMCOnAWSPublicIPGetPublicIpInputType(), typeConverter)
	sv.AddStructField("PublicIpId", publicIpIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.PublicIp
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := cloudServiceVMCOnAWSPublicIPGetPublicIpRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	cIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.api.cloud_service_VMC_on_AWS_public_IP", "get_public_ip", inputDataValue, executionContext)
	var emptyOutput model.PublicIp
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), cloudServiceVMCOnAWSPublicIPGetPublicIpOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.PublicIp), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (cIface *cloudServiceVMCOnAWSPublicIPClient) ListPublicIps(cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.PublicIpsListResult, error) {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(cloudServiceVMCOnAWSPublicIPListPublicIpsInputType(), typeConverter)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.PublicIpsListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := cloudServiceVMCOnAWSPublicIPListPublicIpsRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	cIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.api.cloud_service_VMC_on_AWS_public_IP", "list_public_ips", inputDataValue, executionContext)
	var emptyOutput model.PublicIpsListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), cloudServiceVMCOnAWSPublicIPListPublicIpsOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.PublicIpsListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
