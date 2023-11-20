// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: ConnectedServices
// Used by client-side stubs.

package linked_vpcs

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	nsx_vmc_appModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-vmc-aws-integration/nsx_vmc_app/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type ConnectedServicesClient interface {

	// List connected services to given linked vpc ID. The response consist of all available services along with their status.
	//
	// @param linkedVpcIdParam linked vpc id (required)
	// @return com.vmware.nsx_vmc_app.model.ConnectedServiceListResult
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(linkedVpcIdParam string) (nsx_vmc_appModel.ConnectedServiceListResult, error)

	// Connect/Disconnect the service to the given linked vpc. For example, connect S3. The user will know what services are available through the GET call. If the user is trying to connect/disconnect an unknown service, the POST call will throw a 400 Bad Request error.
	//
	// @param linkedVpcIdParam linked vpc id (required)
	// @param serviceNameParam connected service name, e.g. s3 (required)
	// @param connectedServiceStatusParam (required)
	// @return com.vmware.nsx_vmc_app.model.ConnectedServiceStatus
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(linkedVpcIdParam string, serviceNameParam string, connectedServiceStatusParam nsx_vmc_appModel.ConnectedServiceStatus) (nsx_vmc_appModel.ConnectedServiceStatus, error)
}

type connectedServicesClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewConnectedServicesClient(connector vapiProtocolClient_.Connector) *connectedServicesClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx_vmc_app.infra.linked_vpcs.connected_services")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"list":   vapiCore_.NewMethodIdentifier(interfaceIdentifier, "list"),
		"update": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "update"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	cIface := connectedServicesClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &cIface
}

func (cIface *connectedServicesClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := cIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (cIface *connectedServicesClient) List(linkedVpcIdParam string) (nsx_vmc_appModel.ConnectedServiceListResult, error) {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	operationRestMetaData := connectedServicesListRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(connectedServicesListInputType(), typeConverter)
	sv.AddStructField("LinkedVpcId", linkedVpcIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_vmc_appModel.ConnectedServiceListResult
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.nsx_vmc_app.infra.linked_vpcs.connected_services", "list", inputDataValue, executionContext)
	var emptyOutput nsx_vmc_appModel.ConnectedServiceListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), ConnectedServicesListOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_vmc_appModel.ConnectedServiceListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (cIface *connectedServicesClient) Update(linkedVpcIdParam string, serviceNameParam string, connectedServiceStatusParam nsx_vmc_appModel.ConnectedServiceStatus) (nsx_vmc_appModel.ConnectedServiceStatus, error) {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	operationRestMetaData := connectedServicesUpdateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(connectedServicesUpdateInputType(), typeConverter)
	sv.AddStructField("LinkedVpcId", linkedVpcIdParam)
	sv.AddStructField("ServiceName", serviceNameParam)
	sv.AddStructField("ConnectedServiceStatus", connectedServiceStatusParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_vmc_appModel.ConnectedServiceStatus
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.nsx_vmc_app.infra.linked_vpcs.connected_services", "update", inputDataValue, executionContext)
	var emptyOutput nsx_vmc_appModel.ConnectedServiceStatus
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), ConnectedServicesUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_vmc_appModel.ConnectedServiceStatus), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
