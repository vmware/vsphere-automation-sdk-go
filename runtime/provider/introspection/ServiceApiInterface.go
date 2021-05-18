/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

package introspection

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/core"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/lib"
)

type ServiceApiInterface struct {
	interfaceId   core.InterfaceIdentifier
	methodDefs    map[core.MethodIdentifier]*core.MethodDefinition
	introspection APIIntrospection
	methods       map[core.MethodIdentifier]func(ctx *core.ExecutionContext,
		methodDef core.MethodDefinition, inputValue data.DataValue) core.MethodResult
	providerName string
}

func NewServiceApiInterface(name string, introspection APIIntrospection) (*ServiceApiInterface, error) {
	var interfaceId = core.NewInterfaceIdentifier("com.vmware.vapi.std.introspection.service")
	var methodDefs = make(map[core.MethodIdentifier]*core.MethodDefinition)
	var methods = make(map[core.MethodIdentifier]func(ctx *core.ExecutionContext,
		methodDef core.MethodDefinition, inputValue data.DataValue) core.MethodResult)
	var serviceApiInterface = &ServiceApiInterface{interfaceId: interfaceId,
		methodDefs: methodDefs, introspection: introspection, providerName: name, methods: methods}
	serviceApiInterface.registerListMethod()
	serviceApiInterface.registerGetMethod()
	return serviceApiInterface, nil
}

func (serviceApiInterface *ServiceApiInterface) Identifier() core.InterfaceIdentifier {
	return serviceApiInterface.interfaceId
}

func (serviceApiInterface *ServiceApiInterface) Definition() core.InterfaceDefinition {
	var methodIds = serviceApiInterface.getListOfMethodIds()
	return core.NewInterfaceDefinition(serviceApiInterface.interfaceId, methodIds)
}

func (serviceApiInterface *ServiceApiInterface) MethodDefinition(md core.MethodIdentifier) *core.MethodDefinition {
	return serviceApiInterface.methodDefs[md]
}

func (serviceApiInterface *ServiceApiInterface) Invoke(ctx *core.ExecutionContext,
	methodId core.MethodIdentifier, input data.DataValue) core.MethodResult {
	var method = serviceApiInterface.methods[methodId]
	var methodDefinition = serviceApiInterface.methodDefs[methodId]
	return method(ctx, *methodDefinition, input)
}

func (serviceApiInterface *ServiceApiInterface) getListOfMethodIds() map[string]core.MethodIdentifier {
	var methodIdentifiers = make(map[string]core.MethodIdentifier)
	for key, _ := range serviceApiInterface.methods {
		methodIdentifiers[key.Name()] = key
	}
	return methodIdentifiers
}

func (serviceApiInterface *ServiceApiInterface) registerGetMethod() {
	var getMethodId = core.NewMethodIdentifier(serviceApiInterface.interfaceId, "get")
	var outputDefMap = make(map[string]data.DataDefinition, 0)
	var listDef = data.NewListDefinition(data.NewStringDefinition())
	outputDefMap["operations"] = listDef
	var outputDef = data.NewStructDefinition("com.vmware.vapi.std.introspection.service.info", outputDefMap)
	var inputDefMap = make(map[string]data.DataDefinition, 0)
	inputDefMap["id"] = data.NewStringDefinition()
	var inputDef = data.NewStructDefinition(lib.OPERATION_INPUT, inputDefMap)
	var methodDef = core.NewMethodDefinition(getMethodId, inputDef, outputDef, errorDefs)
	serviceApiInterface.methodDefs[getMethodId] = &methodDef
	serviceApiInterface.methods[getMethodId] = serviceApiInterface.get
}

func (serviceApiInterface *ServiceApiInterface) get(ctx *core.ExecutionContext,
	methodDef core.MethodDefinition, inputValue data.DataValue) core.MethodResult {

	var structVal = inputValue.(*data.StructValue)
	var result, error = structVal.String("id")
	if error != nil {
		//TODO
	}
	return serviceApiInterface.introspection.GetServiceInfo(result)
}

func (serviceApiInterface *ServiceApiInterface) registerListMethod() {
	var listMethodId = core.NewMethodIdentifier(serviceApiInterface.interfaceId, "list")
	var outputDef = data.NewListDefinition(data.NewStringDefinition())
	var inputDef = data.NewStructDefinition(lib.OPERATION_INPUT, make(map[string]data.DataDefinition))
	var methodDef = core.NewMethodDefinition(listMethodId, inputDef, outputDef, nil)
	serviceApiInterface.methodDefs[listMethodId] = &methodDef
	serviceApiInterface.methods[listMethodId] = serviceApiInterface.list
}

func (serviceApiInterface *ServiceApiInterface) list(ctx *core.ExecutionContext,
	methodDef core.MethodDefinition, inputValue data.DataValue) core.MethodResult {
	var listValue = data.NewListValue()
	for _, service := range serviceApiInterface.introspection.GetServices() {
		var strValue = data.NewStringValue(service)
		listValue.Add(strValue)
	}
	return core.NewMethodResult(listValue, nil)
}
