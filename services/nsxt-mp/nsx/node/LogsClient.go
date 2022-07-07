// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Logs
// Used by client-side stubs.

package node

import (
	"github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
)

const _ = core.SupportedByRuntimeVersion1

type LogsClient interface {

	// For a single specified log file, lists the filename, file size, and last-modified time.
	//
	// @param logNameParam Name of log file to read properties (required)
	// @return com.vmware.nsx.model.NodeLogProperties
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(logNameParam string) (model.NodeLogProperties, error)

	// Returns the number of log files and lists the log files that reside on the NSX virtual appliance. The list includes the filename, file size, and last-modified time in milliseconds since epoch (1 January 1970) for each log file. Knowing the last-modified time with millisecond accuracy since epoch is helpful when you are comparing two times, such as the time of a POST request and the end time on a server.
	// @return com.vmware.nsx.model.NodeLogPropertiesListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List() (model.NodeLogPropertiesListResult, error)
}

type logsClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewLogsClient(connector client.Connector) *logsClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.nsx.node.logs")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"get":  core.NewMethodIdentifier(interfaceIdentifier, "get"),
		"list": core.NewMethodIdentifier(interfaceIdentifier, "list"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	lIface := logsClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &lIface
}

func (lIface *logsClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := lIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (lIface *logsClient) Get(logNameParam string) (model.NodeLogProperties, error) {
	typeConverter := lIface.connector.TypeConverter()
	executionContext := lIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(logsGetInputType(), typeConverter)
	sv.AddStructField("LogName", logNameParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.NodeLogProperties
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := logsGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	lIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := lIface.connector.GetApiProvider().Invoke("com.vmware.nsx.node.logs", "get", inputDataValue, executionContext)
	var emptyOutput model.NodeLogProperties
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), logsGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.NodeLogProperties), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), lIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (lIface *logsClient) List() (model.NodeLogPropertiesListResult, error) {
	typeConverter := lIface.connector.TypeConverter()
	executionContext := lIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(logsListInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.NodeLogPropertiesListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := logsListRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	lIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := lIface.connector.GetApiProvider().Invoke("com.vmware.nsx.node.logs", "list", inputDataValue, executionContext)
	var emptyOutput model.NodeLogPropertiesListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), logsListOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.NodeLogPropertiesListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), lIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
