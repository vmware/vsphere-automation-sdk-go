/* Copyright © 2019-2020 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

package introspection

import (
	"crypto/sha512"
	"encoding/hex"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/l10n"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib"
	"sort"
)

var errorDefs = []data.ErrorDefinition{
	bindings.INTERNAL_SERVER_ERROR_DEF,
	bindings.INVALID_ARGUMENT_ERROR_DEF,
	bindings.NOT_FOUND_ERROR_DEF,
	bindings.UNEXPECTED_INPUT_ERROR_DEF,
}

var introspectionTypeMap = map[data.DataType]string{
	data.INTEGER:           "LONG",
	data.DOUBLE:            "DOUBLE",
	data.BOOLEAN:           "BOOLEAN",
	data.STRING:            "STRING",
	data.BLOB:              "BINARY",
	data.LIST:              "LIST",
	data.STRUCTURE:         "STRUCTURE",
	data.OPTIONAL:          "OPTIONAL",
	data.VOID:              "VOID",
	data.OPAQUE:            "OPAQUE",
	data.SECRET:            "SECRET",
	data.ERROR:             "ERROR",
	data.STRUCTURE_REF:     "STRUCTURE_REF",
	data.DYNAMIC_STRUCTURE: "DYNAMIC_STRUCTURE",
	data.ANY_ERROR:         "ANY_ERROR"}

type LocalProviderIntrospector struct {
	introspectionServices []core.ApiInterface
	name                  string
	serviceData           map[string]core.ApiInterface
}

func NewLocalProviderIntrospector(name string) (*LocalProviderIntrospector, error) {
	var localProviderIntrospector = LocalProviderIntrospector{name: name}
	providerAPIInterface, _ := NewProviderApiInterface(name, &localProviderIntrospector)
	serviceAPIInterface, _ := NewServiceApiInterface(name, &localProviderIntrospector)
	operationAPIInterface, _ := NewOperationApiInterface(name, &localProviderIntrospector)
	localProviderIntrospector.introspectionServices = []core.ApiInterface{providerAPIInterface,
		serviceAPIInterface, operationAPIInterface}
	localProviderIntrospector.serviceData = map[string]core.ApiInterface{}
	return &localProviderIntrospector, nil
}

func (localProviderIntrospector *LocalProviderIntrospector) GetCheckSum() string {
	hash := sha512.New()

	// Directly iterating over the map is not guaranteed to be sequential.
	// Which will result in non-unique checkSums
	apiInterfaceNames := make([]string, 0)
	for _, apiInterface := range localProviderIntrospector.serviceData {
		apiInterfaceNames = append(apiInterfaceNames, apiInterface.Identifier().Name())
	}
	sort.Strings(apiInterfaceNames)
	for _, apiInterfaceName := range apiInterfaceNames {
		apiInterface := localProviderIntrospector.serviceData[apiInterfaceName]
		hash.Write([]byte(apiInterface.Identifier().Name()))
		methodIdentifiers := apiInterface.Definition().MethodIdentifiers()
		// Directly iterating over the map is not guaranteed to be sequential.
		// Which will result in non-unique checkSums
		methodNames := make([]string, len(methodIdentifiers))
		for _, methodIdentifier := range methodIdentifiers {
			methodNames = append(methodNames, methodIdentifier.Name())
		}
		sort.Strings(methodNames)
		for _, methodName := range methodNames {
			hash.Write([]byte(methodName))
			methodDefinition := apiInterface.MethodDefinition(methodIdentifiers[methodName])
			if methodDefinition != nil {
				hash.Write([]byte(methodDefinition.String()))
			}
		}
	}
	return hex.EncodeToString(hash.Sum(nil))
}

func (localProviderIntrospector *LocalProviderIntrospector) GetOperations(serviceID string) core.MethodResult {
	var apiInterface = localProviderIntrospector.serviceData[serviceID]
	if apiInterface != nil {
		var methodIds = apiInterface.Definition().MethodIdentifiers()
		var listValue = data.NewListValue()
		for _, method := range methodIds {
			var methodName = method.Name()
			var strValue = data.NewStringValue(methodName)
			listValue.Add(strValue)
		}
		return core.NewMethodResult(listValue, nil)
	}
	return serviceNotFoundMethodResult(serviceID)
}

func (localProviderIntrospector *LocalProviderIntrospector) GetOperationInfo(serviceID string,
	operationID string) core.MethodResult {
	var apiInterface = localProviderIntrospector.serviceData[serviceID]
	if apiInterface != nil {
		var methodIds = apiInterface.Definition().MethodIdentifiers()

		for _, methodID := range methodIds {
			if operationID == methodID.Name() {
				methodDef := apiInterface.MethodDefinition(methodID)
				dataVal := localProviderIntrospector.convertMethodDefToDataValue(*methodDef)
				return core.NewMethodResult(dataVal, nil)
			}

		}
		args := map[string]string{
			"serviceName":   serviceID,
			"operationName": operationID}
		var opNotFoundMsg = l10n.NewRuntimeError("vapi.introspection.operation.not_found", args)
		var errorValue = bindings.CreateErrorValueFromMessages(bindings.NOT_FOUND_ERROR_DEF,
			[]error{opNotFoundMsg})
		return core.NewMethodResult(nil, errorValue)
	}
	return serviceNotFoundMethodResult(serviceID)
}

func (localProviderIntrospector *LocalProviderIntrospector) AddService(serviceId string,
	apiInterface core.ApiInterface) {
	localProviderIntrospector.serviceData[serviceId] = apiInterface
}

func (localProviderIntrospector *LocalProviderIntrospector) GetIntrospectionServices() []core.ApiInterface {
	return localProviderIntrospector.introspectionServices
}
func (localProviderIntrospector *LocalProviderIntrospector) GetServiceInfo(serviceId string) core.MethodResult {
	var apiInterface = localProviderIntrospector.serviceData[serviceId]
	if apiInterface != nil {
		var methodIds = apiInterface.Definition().MethodIdentifiers()
		var listValue = data.NewListValue()
		for _, method := range methodIds {
			var methodName = method.Name()
			var strValue = data.NewStringValue(methodName)
			listValue.Add(strValue)
		}
		var structureFields = make(map[string]data.DataValue, 1)
		structureFields["operations"] = listValue
		var output = data.NewStructValue("com.vmware.vapi.std.introspection.service.info", structureFields)
		return core.NewMethodResult(output, nil)
	}
	return serviceNotFoundMethodResult(serviceId)
}

func (localProviderIntrospector *LocalProviderIntrospector) GetServiceInfoDataValue(serviceId string) core.MethodResult {
	var apiInterface = localProviderIntrospector.serviceData[serviceId]
	if apiInterface != nil {
		var methodIds = apiInterface.Definition().MethodIdentifiers()
		var listValue = data.NewListValue()
		for _, method := range methodIds {
			var methodName = method.Name()
			var strValue = data.NewStringValue(methodName)
			listValue.Add(strValue)
		}
		var structureFields = make(map[string]data.DataValue, 1)
		structureFields["operations"] = listValue
		var output = data.NewStructValue("com.vmware.vapi.std.introspection.service.info", structureFields)
		return core.NewMethodResult(output, nil)
	}
	return serviceNotFoundMethodResult(serviceId)
}

func (localProviderIntrospector *LocalProviderIntrospector) GetOperationsDataValue(serviceId string) core.MethodResult {
	var apiInterface = localProviderIntrospector.serviceData[serviceId]
	if apiInterface != nil {
		var methodIds = apiInterface.Definition().MethodIdentifiers()
		var listValue = data.NewListValue()
		for _, method := range methodIds {
			var methodName = method.Name()
			var strValue = data.NewStringValue(methodName)
			listValue.Add(strValue)
		}
		return core.NewMethodResult(listValue, nil)
	}
	return serviceNotFoundMethodResult(serviceId)
}

func (localProviderIntrospector *LocalProviderIntrospector) convertMethodDefToDataValue(md core.MethodDefinition) data.DataValue {
	var inputDefinition = localProviderIntrospector.convertDataDefinitionToDataValue(md.InputDefinition())
	var outputDefinition = localProviderIntrospector.convertDataDefinitionToDataValue(md.OutputDefinition())
	var errors = data.NewListValue()
	for _, errordef := range md.ErrorDefinitions() {
		errors.Add(localProviderIntrospector.convertDataDefinitionToDataValue(errordef))
	}

	// augment error definitions
	for _, errordef := range errorDefs {
		errors.Add(localProviderIntrospector.convertDataDefinitionToDataValue(errordef))
	}

	var fieldMap = make(map[string]data.DataValue)
	fieldMap["input_definition"] = inputDefinition
	fieldMap["output_definition"] = outputDefinition
	fieldMap["error_definitions"] = errors
	return data.NewStructValue("com.vmware.vapi.std.introspection.operation.info", fieldMap)
}

// TODO: Use type converter to convert from DataDefinition to DataValue
func (localProviderIntrospector *LocalProviderIntrospector) convertDataDefinitionToDataValue(dataDef data.DataDefinition) data.DataValue {
	var result = data.NewStructValue("com.vmware.vapi.std.introspection.operation.data_definition", nil)
	var strValue = data.NewStringValue(introspectionTypeMap[dataDef.Type()])
	result.SetField("type", strValue)

	switch dataDef.Type() {
	case data.STRUCTURE:
		result.SetField("name",
			data.NewOptionalValue(data.NewStringValue((dataDef).(data.StructDefinition).Name())))
	case data.STRUCTURE_REF:
		result.SetField("name",
			data.NewOptionalValue(data.NewStringValue((dataDef).(*data.StructRefDefinition).Name())))
	case data.ERROR:
		result.SetField("name",
			data.NewOptionalValue(data.NewStringValue((dataDef).(data.ErrorDefinition).Name())))
	default:
		result.SetField("name", data.NewOptionalValue(nil))
	}

	switch dataDef.Type() {
	case data.OPTIONAL:
		var elementDef = (dataDef).(data.OptionalDefinition).ElementType()
		var elementValue = localProviderIntrospector.convertDataDefinitionToDataValue(elementDef)
		result.SetField("element_definition", data.NewOptionalValue(elementValue))
	case data.LIST:
		var elementDef = (dataDef).(data.ListDefinition).ElementType()
		var elementValue = localProviderIntrospector.convertDataDefinitionToDataValue(elementDef)
		result.SetField("element_definition", data.NewOptionalValue(elementValue))
	default:
		result.SetField("element_definition", data.NewOptionalValue(nil))
	}

	switch dataDef.Type() {
	case data.STRUCTURE:
		var fields = data.NewListValue()
		var structDef = (dataDef).(data.StructDefinition)
		for _, fieldName := range structDef.FieldNames() {
			var elementDefinition = structDef.Field(fieldName)
			var fieldPair = data.NewStructValue(lib.MAP_ENTRY, nil)
			var strValue = data.NewStringValue(fieldName)
			fieldPair.SetField(lib.MAP_KEY_FIELD, strValue)
			fieldPair.SetField(lib.MAP_VALUE_FIELD, localProviderIntrospector.convertDataDefinitionToDataValue(elementDefinition))
			fields.Add(fieldPair)
		}
		result.SetField(lib.STRUCT_FIELDS, data.NewOptionalValue(fields))
	case data.ERROR:
		var fields = data.NewListValue()
		var errorDefinition = (dataDef).(data.ErrorDefinition)
		for _, fieldName := range errorDefinition.FieldNames() {
			var elementDefinition = errorDefinition.Field(fieldName)
			var fieldPair = data.NewStructValue(lib.MAP_ENTRY, nil)
			var strValue = data.NewStringValue(fieldName)
			fieldPair.SetField(lib.MAP_KEY_FIELD, strValue)
			fieldPair.SetField(lib.MAP_VALUE_FIELD, localProviderIntrospector.convertDataDefinitionToDataValue(elementDefinition))
			fields.Add(fieldPair)
		}
		result.SetField(lib.STRUCT_FIELDS, data.NewOptionalValue(fields))
	default:
		result.SetField(lib.STRUCT_FIELDS, data.NewOptionalValue(nil))
	}

	return result
}

func (localProviderIntrospector *LocalProviderIntrospector) GetOperationInfoDataValue(serviceID string,
	operationID string) core.MethodResult {
	var apiInterface = localProviderIntrospector.serviceData[serviceID]
	if apiInterface != nil {
		var methodIds = apiInterface.Definition().MethodIdentifiers()

		for _, methodID := range methodIds {
			if operationID == methodID.Name() {
				methodDef := apiInterface.MethodDefinition(methodID)
				dataVal := localProviderIntrospector.convertMethodDefToDataValue(*methodDef)
				return core.NewMethodResult(dataVal, nil)
			}

		}
		args := map[string]string{
			"serviceName":   serviceID,
			"operationName": operationID}
		var opNotFoundMsg = l10n.NewRuntimeError("vapi.introspection.operation.not_found", args)
		var errorValue = bindings.CreateErrorValueFromMessages(bindings.NOT_FOUND_ERROR_DEF,
			[]error{opNotFoundMsg})
		return core.NewMethodResult(nil, errorValue)
	}
	return serviceNotFoundMethodResult(serviceID)

}

func (localProviderIntrospector *LocalProviderIntrospector) GetServices() []string {
	var serviceIds = []string{}
	for key, _ := range localProviderIntrospector.serviceData {
		serviceIds = append(serviceIds, key)
	}
	return serviceIds
}

func serviceNotFoundMethodResult(serviceId string) core.MethodResult {
	var msg = l10n.NewRuntimeError("vapi.introspection.service.not_found",
		map[string]string{"serviceName": serviceId})
	var errorValue = bindings.CreateErrorValueFromMessages(bindings.NOT_FOUND_ERROR_DEF,
		[]error{msg})
	return core.NewMethodResult(nil, errorValue)
}
