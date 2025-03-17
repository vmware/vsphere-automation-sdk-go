// Copyright (c) 2019-2024 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

package introspection

import (
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/log"
)

type OperationApiInterface struct {
	interfaceID   core.InterfaceIdentifier
	methodDefs    map[core.MethodIdentifier]*core.MethodDefinition
	introspection APIIntrospection
	methods       map[core.MethodIdentifier]func(ctx *core.ExecutionContext, methodId core.MethodIdentifier, input data.DataValue) core.MethodResult
	providerName  string
}

func NewOperationApiInterface(name string, introspection APIIntrospection) (*OperationApiInterface, error) {
	var interfaceId = core.NewInterfaceIdentifier("com.vmware.vapi.std.introspection.operation")
	var methodDefs = make(map[core.MethodIdentifier]*core.MethodDefinition)
	var methods = make(map[core.MethodIdentifier]func(ctx *core.ExecutionContext,
		methodId core.MethodIdentifier, input data.DataValue) core.MethodResult)
	var operationAPIInterface = &OperationApiInterface{interfaceID: interfaceId,
		methodDefs:    methodDefs,
		introspection: introspection,
		methods:       methods,
		providerName:  name}
	operationAPIInterface.registerGetMethod()
	operationAPIInterface.registerListMethod()
	return operationAPIInterface, nil
}

func (operationApiInterface *OperationApiInterface) registerGetMethod() {
	var getMethodID = core.NewMethodIdentifier(operationApiInterface.interfaceID, "get")
	var dataRefDefinition = data.NewStructRefDefinition(
		"com.vmware.vapi.std.introspection.operation.data_definition", nil)
	var fieldMap = make(map[string]data.DataDefinition, 2)
	fieldMap["key"] = data.NewStringDefinition()
	fieldMap["value"] = dataRefDefinition
	var fieldDef = data.NewStructDefinition(lib.MAP_ENTRY, fieldMap)
	var dataDefMap = make(map[string]data.DataDefinition, 4)
	dataDefMap["type"] = data.NewStringDefinition()
	var optDefElementDef = data.NewOptionalDefinition(dataRefDefinition)
	dataDefMap["element_definition"] = optDefElementDef
	var optDefStringDef = data.NewOptionalDefinition(data.NewStringDefinition())
	dataDefMap["name"] = optDefStringDef
	var optDefFields = data.NewOptionalDefinition(data.NewListDefinition(fieldDef))
	dataDefMap[lib.STRUCT_FIELDS] = optDefFields

	var dataDefinition = data.NewStructDefinition(
		"com.vmware.vapi.std.introspection.operation.data_definition", dataDefMap)
	targetError := dataRefDefinition.SetTarget(&dataDefinition)
	if targetError != nil {
		log.Error("Error setting target on dataref ")
	}
	var outputDefMap = make(map[string]data.DataDefinition)
	outputDefMap["input_definition"] = dataDefinition
	outputDefMap["output_definition"] = dataDefinition
	outputDefMap["error_definitions"] = data.NewListDefinition(dataDefinition)
	var outputDef = data.NewStructDefinition("com.vmware.vapi.std.introspection.operation.info", outputDefMap)
	var inputDefMap = make(map[string]data.DataDefinition)
	inputDefMap["service_id"] = data.NewStringDefinition()
	inputDefMap["operation_id"] = data.NewStringDefinition()
	var inputDef = data.NewStructDefinition(lib.OPERATION_INPUT, inputDefMap)
	var getMethodDef = core.NewMethodDefinition(getMethodID, inputDef, outputDef, errorDefs)
	operationApiInterface.methodDefs[getMethodID] = &getMethodDef
	operationApiInterface.methods[getMethodID] = operationApiInterface.get
}

func (operationApiInterface *OperationApiInterface) get(ctx *core.ExecutionContext,
	md core.MethodIdentifier, inputValue data.DataValue) core.MethodResult {
	var structValue = inputValue.(*data.StructValue)
	var serviceID, _ = structValue.StringField("service_id")
	var operationID, _ = structValue.StringField("operation_id")
	return operationApiInterface.introspection.GetOperationInfo(serviceID, operationID)
}

func (operationApiInterface *OperationApiInterface) registerListMethod() {
	var listMethodID = core.NewMethodIdentifier(operationApiInterface.interfaceID, "list")
	var outputDefinition = data.NewListDefinition(data.NewStringDefinition())
	var inputMap = make(map[string]data.DataDefinition)
	inputMap["service_id"] = data.NewStringDefinition()
	var inputDefinition = data.NewStructDefinition(lib.OPERATION_INPUT, inputMap)
	var methodDefinition = core.NewMethodDefinition(listMethodID, inputDefinition, outputDefinition, errorDefs)
	operationApiInterface.methodDefs[listMethodID] = &methodDefinition
	operationApiInterface.methods[listMethodID] = operationApiInterface.list
}

func (operationApiInterface *OperationApiInterface) list(ctx *core.ExecutionContext,
	methodId core.MethodIdentifier, input data.DataValue) core.MethodResult {
	var structValue = (input).(*data.StructValue)
	var serviceID, _ = structValue.StringField("service_id")
	var methodResult = operationApiInterface.introspection.GetOperations(serviceID)
	return methodResult
}

func (operationApiInterface *OperationApiInterface) Identifier() core.InterfaceIdentifier {
	return operationApiInterface.interfaceID
}

func (operationApiInterface *OperationApiInterface) Definition() core.InterfaceDefinition {
	return core.NewInterfaceDefinition(operationApiInterface.interfaceID, operationApiInterface.getListOfMethodIds())
}

func (operationApiInterface *OperationApiInterface) MethodDefinition(md core.MethodIdentifier) *core.MethodDefinition {
	return operationApiInterface.methodDefs[md]
}

func (operationApiInterface *OperationApiInterface) Invoke(ctx *core.ExecutionContext,
	methodID core.MethodIdentifier, input data.DataValue) core.MethodResult {
	var method = operationApiInterface.methods[methodID]
	return method(ctx, methodID, input)
}

// TODO:
// refactor this
func (operationApiInterface *OperationApiInterface) getListOfMethodIds() map[string]core.MethodIdentifier {
	var methods = make(map[string]core.MethodIdentifier)
	for key, _ := range operationApiInterface.methods {
		methods[key.Name()] = key
	}
	return methods
}
