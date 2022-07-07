// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: ContainerApplicationInstances
// Used by client-side stubs.

package fabric

import (
	"github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
)

const _ = core.SupportedByRuntimeVersion1

type ContainerApplicationInstancesClient interface {

	// Returns information about a specific container application instance.
	//
	// @param containerApplicationInstanceIdParam (required)
	// @return com.vmware.nsx.model.ContainerApplicationInstance
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(containerApplicationInstanceIdParam string) (model.ContainerApplicationInstance, error)

	// Returns information about all container application instance.
	//
	// @param containerApplicationIdParam Identifier of the container application (optional)
	// @param containerClusterIdParam Identifier of the container cluster (optional)
	// @param containerProjectIdParam Identifier of the container project (optional)
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param scopeIdParam Scope id for container cluster (optional)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @return com.vmware.nsx.model.ContainerApplicationInstanceListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(containerApplicationIdParam *string, containerClusterIdParam *string, containerProjectIdParam *string, cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, scopeIdParam *string, sortAscendingParam *bool, sortByParam *string) (model.ContainerApplicationInstanceListResult, error)
}

type containerApplicationInstancesClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewContainerApplicationInstancesClient(connector client.Connector) *containerApplicationInstancesClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.nsx.fabric.container_application_instances")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"get":  core.NewMethodIdentifier(interfaceIdentifier, "get"),
		"list": core.NewMethodIdentifier(interfaceIdentifier, "list"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	cIface := containerApplicationInstancesClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &cIface
}

func (cIface *containerApplicationInstancesClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := cIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (cIface *containerApplicationInstancesClient) Get(containerApplicationInstanceIdParam string) (model.ContainerApplicationInstance, error) {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(containerApplicationInstancesGetInputType(), typeConverter)
	sv.AddStructField("ContainerApplicationInstanceId", containerApplicationInstanceIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ContainerApplicationInstance
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := containerApplicationInstancesGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	cIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.nsx.fabric.container_application_instances", "get", inputDataValue, executionContext)
	var emptyOutput model.ContainerApplicationInstance
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), containerApplicationInstancesGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ContainerApplicationInstance), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (cIface *containerApplicationInstancesClient) List(containerApplicationIdParam *string, containerClusterIdParam *string, containerProjectIdParam *string, cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, scopeIdParam *string, sortAscendingParam *bool, sortByParam *string) (model.ContainerApplicationInstanceListResult, error) {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(containerApplicationInstancesListInputType(), typeConverter)
	sv.AddStructField("ContainerApplicationId", containerApplicationIdParam)
	sv.AddStructField("ContainerClusterId", containerClusterIdParam)
	sv.AddStructField("ContainerProjectId", containerProjectIdParam)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("ScopeId", scopeIdParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ContainerApplicationInstanceListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := containerApplicationInstancesListRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	cIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.nsx.fabric.container_application_instances", "list", inputDataValue, executionContext)
	var emptyOutput model.ContainerApplicationInstanceListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), containerApplicationInstancesListOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ContainerApplicationInstanceListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
