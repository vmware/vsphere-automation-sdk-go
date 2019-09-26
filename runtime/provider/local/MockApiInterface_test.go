package local_test

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/core"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/l10n"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/lib"
)

//Data used by mockup classes and the test methods
var interfaceId = core.NewInterfaceIdentifier("MockupInterface")

//Method names
var missingInputDefinition = "missing_input_definition"
var invalidInputDefinition = "invalid_input_definition"
var missingOutputDefinition = "missing_output_definition"
var invalidOutputDefinition = "invalid_output_definition"
var mockDefinition = "increment"
var reportDeclaredError = "report_declared_error"
var reportUndeclaredError = "report_undeclared_error"
var returnInvalidOutput = "return_invalid_output"

// Method Identifiers
var missingInputDefinitionId = core.NewMethodIdentifier(interfaceId, missingInputDefinition)
var invalidInputDefinitionId = core.NewMethodIdentifier(interfaceId, invalidInputDefinition)
var missingOutputDefinitionId = core.NewMethodIdentifier(interfaceId, missingOutputDefinition)
var invalidOutputDefinitionId = core.NewMethodIdentifier(interfaceId, missingOutputDefinition)
var mockDefinitionId = core.NewMethodIdentifier(interfaceId, mockDefinition)
var reportDeclaredErrorId = core.NewMethodIdentifier(interfaceId, reportDeclaredError)
var reportUnDeclaredErrorId = core.NewMethodIdentifier(interfaceId, reportUndeclaredError)
var reportInvalidOutputId = core.NewMethodIdentifier(interfaceId, returnInvalidOutput)

var NOT_FOUND_ERROR_DEF = bindings.CreateStdErrorDefinition("com.vmware.vapi.std.errors.not_found")
var msg = l10n.NewError("mockup.message.Id", "mockup error message", map[string]string{})
var NOT_FOUND_ERROR_VALUE = bindings.CreateErrorValueFromMessages(NOT_FOUND_ERROR_DEF, []error{msg})

type MockApiInterface struct {
	interfaceId core.InterfaceIdentifier
	methodIdMap map[string]*core.MethodIdentifier
}

func NewMockApiInterface() *MockApiInterface {
	return &MockApiInterface{interfaceId: interfaceId}
}

func (mockApiInterface *MockApiInterface) Identifier() core.InterfaceIdentifier {
	return mockApiInterface.interfaceId
}

func (mockApiInterface *MockApiInterface) Definition() core.InterfaceDefinition {
	var methodIdMap = mockApiInterface.getMethodIdMap()
	var methodNames = []core.MethodIdentifier{}
	for _, value := range methodIdMap {
		methodNames = append(methodNames, *value)
	}
	return core.NewInterfaceDefinition(mockApiInterface.interfaceId, methodNames)
}

func (mockApiInterface *MockApiInterface) MethodDefinition(mi core.MethodIdentifier) *core.MethodDefinition {
	var name = mi.Name()
	if name == "bogus_method" {
		return nil
	}
	var input data.DataDefinition
	var output interface{}
	if name == missingInputDefinition {
		input = nil
	} else if name == invalidInputDefinition {
		input = data.NewStringDefinition()
	} else if name == mockDefinition {
		input = data.NewStructDefinition(lib.OPERATION_INPUT, make(map[string]data.DataDefinition))
	} else {
		var fieldMap = make(map[string]data.DataDefinition, 0)
		fieldMap[paramName] = data.NewIntegerDefinition()
		input = data.NewStructDefinition(lib.OPERATION_INPUT, fieldMap)
	}

	if name == missingOutputDefinition {
		output = nil
	} else if name == invalidInputDefinition {
		output = data.NewStringDefinition()
	} else {
		output = data.NewIntegerDefinition()
	}
	var errors = make([]data.ErrorDefinition, 1)
	if name == reportDeclaredError {
		errors = append(errors, NOT_FOUND_ERROR_DEF)
	}

	if output == nil {
		return nil
	}
	var methodDef = core.NewMethodDefinition(mi, input, output.(data.DataDefinition), errors)
	return &methodDef
}

func (mockApiInterface *MockApiInterface) Invoke(ctx *core.ExecutionContext, methodId core.MethodIdentifier, input data.DataValue) core.MethodResult {
	var name = methodId.Name()
	if name == returnInvalidOutput {
		return core.NewMethodResult(data.NewStructValue("something", nil), nil)
	} else if name == mockDefinition {
		var structValue = input.(*data.StructValue)
		var intValue, _ = structValue.Field("param")
		var output = intValue.(*data.IntegerValue).Value() + 1
		var outputValue = data.NewIntegerValue(output)
		return core.NewMethodResult(outputValue, nil)
	}
	return core.NewMethodResult(nil, NOT_FOUND_ERROR_VALUE)
}

var paramName = "param"

//returns map whose key is methodname and value is methodidentifier
func (mockApiInterface *MockApiInterface) getMethodIdMap() map[string]*core.MethodIdentifier {
	if mockApiInterface.methodIdMap == nil {
		mockApiInterface.methodIdMap = make(map[string]*core.MethodIdentifier)
	}
	mockApiInterface.methodIdMap[missingInputDefinition] = &missingInputDefinitionId
	mockApiInterface.methodIdMap[invalidInputDefinition] = &invalidInputDefinitionId
	mockApiInterface.methodIdMap[missingOutputDefinition] = &missingOutputDefinitionId
	mockApiInterface.methodIdMap[invalidOutputDefinition] = &invalidOutputDefinitionId
	mockApiInterface.methodIdMap[mockDefinition] = &mockDefinitionId
	mockApiInterface.methodIdMap[reportDeclaredError] = &reportDeclaredErrorId
	mockApiInterface.methodIdMap[reportUndeclaredError] = &reportUnDeclaredErrorId
	mockApiInterface.methodIdMap[returnInvalidOutput] = &reportInvalidOutputId
	return mockApiInterface.methodIdMap
}

func (mockApiInterface *MockApiInterface) addMethod(methodName string) {
	var newMethodId = core.NewMethodIdentifier(interfaceId, methodName)
	mockApiInterface.methodIdMap[methodName] = &newMethodId
}
