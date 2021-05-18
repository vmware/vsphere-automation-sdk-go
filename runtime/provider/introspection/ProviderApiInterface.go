/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

package introspection

import (
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib"
)

type ProviderApiInterface struct {
	interfaceId   core.InterfaceIdentifier
	methodDefs    map[core.MethodIdentifier]*core.MethodDefinition
	introspection APIIntrospection
	methods       map[core.MethodIdentifier]func(ctx *core.ExecutionContext, methodId core.MethodIdentifier,
		input data.DataValue) core.MethodResult
	providerName string
}

func NewProviderApiInterface(name string, introspection APIIntrospection) (*ProviderApiInterface, error) {
	var interfaceName = "com.vmware.vapi.std.introspection.provider"
	var interfaceId = core.NewInterfaceIdentifier(interfaceName)

	var methodDefs = make(map[core.MethodIdentifier]*core.MethodDefinition)
	var methods = make(map[core.MethodIdentifier]func(ctx *core.ExecutionContext, methodId core.MethodIdentifier,
		input data.DataValue) core.MethodResult)
	var providerApiInterface = ProviderApiInterface{interfaceId: interfaceId, methodDefs: methodDefs,
		introspection: introspection, methods: methods, providerName: name}
	var error = providerApiInterface.registerGetMethod(interfaceId)
	if error != nil {
		return nil, error
	}
	return &providerApiInterface, nil
}

func (providerIntrospectionApiInterface ProviderApiInterface) registerGetMethod(
	interfaceId core.InterfaceIdentifier) error {
	var GET = "get"
	var providerGetMethodId = core.NewMethodIdentifier(interfaceId, GET)

	var inputDefinition = getProviderGetMethodInputDef()
	var outputDefinition = getProviderInfoDef()

	var getMethodDefinition = core.NewMethodDefinition(providerGetMethodId,
		inputDefinition, outputDefinition, nil)
	providerIntrospectionApiInterface.methodDefs[providerGetMethodId] = &getMethodDefinition
	var getMethod = func(ctx *core.ExecutionContext, methodId core.MethodIdentifier,
		input data.DataValue) core.MethodResult {
		var PROVIDER_INFO_NAME = "com.vmware.std.introspection.provider.info"
		var fields = make(map[string]data.DataValue)
		var structValue = data.NewStructValue(PROVIDER_INFO_NAME, fields)
		structValue.SetStringField("id", PROVIDER_INFO_NAME)
		structValue.SetStringField("checksum", providerIntrospectionApiInterface.introspection.GetCheckSum())
		var methodResult = core.NewMethodResult(structValue, nil)
		return methodResult
	}
	providerIntrospectionApiInterface.methods[providerGetMethodId] = getMethod
	return nil
}

func getProviderInfoDef() data.StructDefinition {
	var PROVIDER_INFO_NAME = "com.vmware.std.introspection.provider.info"
	var fields = make(map[string]data.DataDefinition, 0)
	fields["id"] = data.NewStringDefinition()
	fields["checksum"] = data.NewStringDefinition()
	var providerInfoDef = data.NewStructDefinition(PROVIDER_INFO_NAME, fields)
	return providerInfoDef
}

func getProviderGetMethodInputDef() data.StructDefinition {
	var operationInput = lib.OPERATION_INPUT
	var fields = make(map[string]data.DataDefinition, 0)
	var inputDefinition = data.NewStructDefinition(operationInput, fields)
	return inputDefinition
}

func (providerIntrospectionApiInterface *ProviderApiInterface) Identifier() core.InterfaceIdentifier {
	return providerIntrospectionApiInterface.interfaceId
}

func (providerIntrospectionApiInterface *ProviderApiInterface) Definition() core.InterfaceDefinition {
	//refactor this
	var methodIdentifiers = make(map[string]core.MethodIdentifier)
	for key, _ := range providerIntrospectionApiInterface.methodDefs {
		methodIdentifiers[key.Name()] = key
	}
	var interfaceDefinition = core.NewInterfaceDefinition(providerIntrospectionApiInterface.interfaceId, methodIdentifiers)
	return interfaceDefinition
}

func (providerIntrospectionApiInterface ProviderApiInterface) MethodDefinition(
	identifier core.MethodIdentifier) *core.MethodDefinition {

	return providerIntrospectionApiInterface.methodDefs[identifier]
}

func (providerIntrospectionApiInterface ProviderApiInterface) Invoke(ctx *core.ExecutionContext,
	methodId core.MethodIdentifier, input data.DataValue) core.MethodResult {
	var method = providerIntrospectionApiInterface.methods[methodId]
	var methodResult = method(ctx, methodId, input)
	return methodResult
}
