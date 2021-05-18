// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: CloudServiceVMCOnAWSLinkedVPC
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

type CloudServiceVMCOnAWSLinkedVPCClient interface {

	// Get linked VPC information associated with given ID.
	//
	// @param linkedVpcIdParam (required)
	// @return com.vmware.model.LinkedVpcInfo
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetLinkedVpc(linkedVpcIdParam string) (model.LinkedVpcInfo, error)

	// List connected services to given linked vpc ID. The response consist of all available services along with their status.
	//
	// @param linkedVpcIdParam linked vpc id (required)
	// @return com.vmware.model.ConnectedServiceListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ListConnectedServices(linkedVpcIdParam string) (model.ConnectedServiceListResult, error)

	// List all linked VPC information.
	// @return com.vmware.model.LinkedVpcsListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ListLinkedVpcs() (model.LinkedVpcsListResult, error)

	// Connect/Disconnect the service to the given linked vpc. For example, connect S3. The user will know what services are available through the GET call. If the user is trying to connect/disconnect an unknown service, the POST call will throw a 400 Bad Request error.
	//
	// @param linkedVpcIdParam linked vpc id (required)
	// @param serviceNameParam connected service name, e.g. s3 (required)
	// @param connectedServiceStatusParam (required)
	// @return com.vmware.model.ConnectedServiceStatus
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	UpdateConnectedService(linkedVpcIdParam string, serviceNameParam string, connectedServiceStatusParam model.ConnectedServiceStatus) (model.ConnectedServiceStatus, error)
}

type cloudServiceVMCOnAWSLinkedVPCClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewCloudServiceVMCOnAWSLinkedVPCClient(connector client.Connector) *cloudServiceVMCOnAWSLinkedVPCClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.api.cloud_service_VMC_on_AWS_linked_VPC")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"get_linked_vpc":           core.NewMethodIdentifier(interfaceIdentifier, "get_linked_vpc"),
		"list_connected_services":  core.NewMethodIdentifier(interfaceIdentifier, "list_connected_services"),
		"list_linked_vpcs":         core.NewMethodIdentifier(interfaceIdentifier, "list_linked_vpcs"),
		"update_connected_service": core.NewMethodIdentifier(interfaceIdentifier, "update_connected_service"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	cIface := cloudServiceVMCOnAWSLinkedVPCClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &cIface
}

func (cIface *cloudServiceVMCOnAWSLinkedVPCClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := cIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (cIface *cloudServiceVMCOnAWSLinkedVPCClient) GetLinkedVpc(linkedVpcIdParam string) (model.LinkedVpcInfo, error) {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(cloudServiceVMCOnAWSLinkedVPCGetLinkedVpcInputType(), typeConverter)
	sv.AddStructField("LinkedVpcId", linkedVpcIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.LinkedVpcInfo
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := cloudServiceVMCOnAWSLinkedVPCGetLinkedVpcRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	cIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.api.cloud_service_VMC_on_AWS_linked_VPC", "get_linked_vpc", inputDataValue, executionContext)
	var emptyOutput model.LinkedVpcInfo
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), cloudServiceVMCOnAWSLinkedVPCGetLinkedVpcOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.LinkedVpcInfo), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (cIface *cloudServiceVMCOnAWSLinkedVPCClient) ListConnectedServices(linkedVpcIdParam string) (model.ConnectedServiceListResult, error) {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(cloudServiceVMCOnAWSLinkedVPCListConnectedServicesInputType(), typeConverter)
	sv.AddStructField("LinkedVpcId", linkedVpcIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ConnectedServiceListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := cloudServiceVMCOnAWSLinkedVPCListConnectedServicesRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	cIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.api.cloud_service_VMC_on_AWS_linked_VPC", "list_connected_services", inputDataValue, executionContext)
	var emptyOutput model.ConnectedServiceListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), cloudServiceVMCOnAWSLinkedVPCListConnectedServicesOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ConnectedServiceListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (cIface *cloudServiceVMCOnAWSLinkedVPCClient) ListLinkedVpcs() (model.LinkedVpcsListResult, error) {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(cloudServiceVMCOnAWSLinkedVPCListLinkedVpcsInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.LinkedVpcsListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := cloudServiceVMCOnAWSLinkedVPCListLinkedVpcsRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	cIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.api.cloud_service_VMC_on_AWS_linked_VPC", "list_linked_vpcs", inputDataValue, executionContext)
	var emptyOutput model.LinkedVpcsListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), cloudServiceVMCOnAWSLinkedVPCListLinkedVpcsOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.LinkedVpcsListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (cIface *cloudServiceVMCOnAWSLinkedVPCClient) UpdateConnectedService(linkedVpcIdParam string, serviceNameParam string, connectedServiceStatusParam model.ConnectedServiceStatus) (model.ConnectedServiceStatus, error) {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(cloudServiceVMCOnAWSLinkedVPCUpdateConnectedServiceInputType(), typeConverter)
	sv.AddStructField("LinkedVpcId", linkedVpcIdParam)
	sv.AddStructField("ServiceName", serviceNameParam)
	sv.AddStructField("ConnectedServiceStatus", connectedServiceStatusParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ConnectedServiceStatus
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := cloudServiceVMCOnAWSLinkedVPCUpdateConnectedServiceRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	cIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.api.cloud_service_VMC_on_AWS_linked_VPC", "update_connected_service", inputDataValue, executionContext)
	var emptyOutput model.ConnectedServiceStatus
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), cloudServiceVMCOnAWSLinkedVPCUpdateConnectedServiceOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ConnectedServiceStatus), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
