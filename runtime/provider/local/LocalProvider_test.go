package local_test

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/common"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/core"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/lib"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/log"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/provider/local"
	"testing"
)

/**Interface with no methods */
var EMPTY_INTERFACE_ID = core.NewInterfaceIdentifier("empty.interface")
var NO_SUCH_INTERFACE_ID = core.NewInterfaceIdentifier("no.such.interface")

func TestLocalProviderConstruction(t *testing.T) {
	var localProvider, error = local.NewLocalProvider("LocalProvider", true)
	if error != nil {
		t.Error("Local provider instantiation failed.")
	}
	log.Info(localProvider)
}

func getInitializedLocalProvider() *local.LocalProvider {
	var localProvider, _ = local.NewLocalProvider("LocalProvider", true)
	var apiInterface = NewMockApiInterface()
	localProvider.AddInterface(apiInterface)
	return localProvider
}

func getInitializedLocalProviderWithMockApiInterface(apiInterface *MockApiInterface) *local.LocalProvider {
	var localProvider, _ = local.NewLocalProvider("LocalProvider", true)
	localProvider.AddInterface(apiInterface)
	return localProvider
}

func TestSuccessfulExecution(t *testing.T) {
	var localProvider = getInitializedLocalProvider()
	var serviceId = "MockupInterface"
	var operationId = "increment"
	var input = data.NewStructValue(lib.OPERATION_INPUT, make(map[string]data.DataValue))
	input.SetField("param", data.NewIntegerValue(10))
	var context = core.NewExecutionContext(common.NewDefaultApplicationContext(), nil)
	var methodResult = localProvider.Invoke(serviceId, operationId, input, context)
	assert.True(t, methodResult.IsSuccess())
	var value = methodResult.Output()
	var intValue = value.(*data.IntegerValue)
	assert.Equal(t, intValue.Value(), int64(11))
}

func TestMissingMethod(t *testing.T) {
	var localProvider = getInitializedLocalProvider()
	var serviceId = "MockupInterface"
	var operationId = "bogus_method"
	var input = data.NewStructValue(lib.OPERATION_INPUT, make(map[string]data.DataValue))
	var context = core.NewExecutionContext(common.NewDefaultApplicationContext(), nil)
	var methodResult = localProvider.Invoke(serviceId, operationId, input, context)
	assert.Nil(t, methodResult.Output())
	assert.False(t, methodResult.IsSuccess())
	assert.Equal(t, "com.vmware.vapi.std.errors.operation_not_found", methodResult.Error().Name())

}

func TestMissingInterface(t *testing.T) {
	var localProvider = getInitializedLocalProvider()
	var serviceId = "bogus_interface"
	var operationId = "test_method"
	var input = data.NewStructValue(lib.OPERATION_INPUT, make(map[string]data.DataValue))
	var context = core.NewExecutionContext(common.NewDefaultApplicationContext(), nil)
	var methodResult = localProvider.Invoke(serviceId, operationId, input, context)
	assert.Nil(t, methodResult.Output())
	assert.False(t, methodResult.IsSuccess())
	var errorValue = methodResult.Error()
	var messages, fieldError = errorValue.Field(bindings.MESSAGES_FIELD_NAME)
	messageListValue := messages.(*data.ListValue)
	assert.Nil(t, fieldError)
	for _, msg := range messageListValue.List() {
		var strMsg = msg.(*data.StructValue)
		var id, err = strMsg.Field("id")
		if err != nil {
			log.Error(err)
		}
		assert.Equal(t, "vapi.method.input.invalid.interface", (id).(*data.StringValue).Value())
	}
}

func TestAddDuplicateInterface(t *testing.T) {
	var localProvider, _ = local.NewLocalProvider("LocalProvider", true)
	var apiInterface = NewMockApiInterface()
	var error = localProvider.AddInterface(apiInterface)
	assert.Nil(t, error)
	var apiInterface2 = NewMockApiInterface()
	error = localProvider.AddInterface(apiInterface2)
	assert.NotNil(t, error)
	assert.Equal(t, fmt.Errorf("Duplicate interface %s found. Registration failed ", apiInterface.interfaceId.Name()), error)
}

func TestAddInterfaceForNilValue(t *testing.T) {
	var localProvider, _ = local.NewLocalProvider("LocalProvider", true)
	result := localProvider.AddInterface(nil)
	assert.Equal(t, result, errors.New("interface cannot be nil"))
}

func TestAddInterfacesForNilValue(t *testing.T) {
	var localProvider, _ = local.NewLocalProvider("LocalProvider", true)
	result := localProvider.AddInterfaces(nil)
	assert.Equal(t, result, errors.New("interfaces cannot be nil"))
}

func TestInputValidationFailure(t *testing.T) {
	var localProvider = getInitializedLocalProvider()
	var serviceId = "MockupInterface"
	var operationId = invalidInputDefinition
	var input = data.NewStructValue(lib.OPERATION_INPUT, make(map[string]data.DataValue))
	input.SetField("param", data.NewStringValue("helloworld"))
	var context = core.NewExecutionContext(common.NewDefaultApplicationContext(), nil)
	var methodResult = localProvider.Invoke(serviceId, operationId, input, context)
	assert.Nil(t, methodResult.Output())
	assert.False(t, methodResult.IsSuccess())
	var errorValue = methodResult.Error()
	var messages, messageFieldError = errorValue.Field(bindings.MESSAGES_FIELD_NAME)
	assert.Nil(t, messageFieldError)

	for _, msg := range messages.(*data.ListValue).List() {
		var strMsg = msg.(*data.StructValue)
		var id, _ = strMsg.Field("id")
		assert.Equal(t, "vapi.method.input.invalid.definition", (id).(*data.StringValue).Value())
	}
}

func TestOutputValidationFailure(t *testing.T) {
	var localProvider = getInitializedLocalProvider()
	var serviceId = "MockupInterface"
	var operationId = "return_invalid_output"
	var input = data.NewStructValue(lib.OPERATION_INPUT, make(map[string]data.DataValue))
	input.SetField("param", data.NewIntegerValue(10))
	var context = core.NewExecutionContext(common.NewDefaultApplicationContext(), nil)
	var methodResult = localProvider.Invoke(serviceId, operationId, input, context)
	assert.Nil(t, methodResult.Output())
	assert.False(t, methodResult.IsSuccess())
	var errorValue = methodResult.Error()
	var messages, messageFieldError = errorValue.Field(bindings.MESSAGES_FIELD_NAME)
	assert.Nil(t, messageFieldError)
	for _, msg := range messages.(*data.ListValue).List() {
		var strMsg = msg.(*data.StructValue)
		var id, _ = strMsg.Field("id")
		var defaultMsg, _ = strMsg.Field("default_message")
		var args, _ = strMsg.Field("args")
		var argsList = (args).(*data.ListValue).List()
		argsString := []string{argsList[0].(*data.StringValue).Value(), argsList[1].(*data.StringValue).Value()}
		assert.Contains(t, argsString, "INTEGER")
		assert.Contains(t, argsString, "STRUCTURE")
		assert.Equal(t, "Type mismatch - expected an object of type 'INTEGER', but got 'STRUCTURE'",
			defaultMsg.(*data.StringValue).Value())
		assert.Equal(t, "vapi.data.validate.mismatch", (id).(*data.StringValue).Value())
	}
}

func TestReportUndeclaredError(t *testing.T) {
	var localProvider = getInitializedLocalProvider()
	var serviceId = "MockupInterface"
	var operationId = "report_undeclared_error"
	var input = data.NewStructValue(lib.OPERATION_INPUT, make(map[string]data.DataValue))
	input.SetField("param", data.NewIntegerValue(10))
	var context = core.NewExecutionContext(common.NewDefaultApplicationContext(), nil)
	var methodResult = localProvider.Invoke(serviceId, operationId, input, context)
	assert.Nil(t, methodResult.Output())
	assert.False(t, methodResult.IsSuccess())
	var errorValue = methodResult.Error()
	var messages, fieldError = errorValue.Field(bindings.MESSAGES_FIELD_NAME)
	assert.Nil(t, fieldError)
	var msg1 = messages.(*data.ListValue).List()[0]
	var strMsg1 = msg1.(*data.StructValue)
	var id, _ = strMsg1.Field("id")
	assert.Equal(t, "vapi.method.status.errors.invalid", id.(*data.StringValue).Value())
	var defaultMsg, _ = strMsg1.Field("default_message")
	assert.Equal(t, "Invalid error 'com.vmware.vapi.std.errors.not_found' reported from method 'report_undeclared_error'",
		defaultMsg.(*data.StringValue).Value())
	var args, _ = strMsg1.Field("args")
	var argsListValue = (args).(*data.ListValue)
	argsStrs := []string{argsListValue.List()[0].(*data.StringValue).Value(),
		argsListValue.List()[1].(*data.StringValue).Value()}
	assert.Contains(t, argsStrs, "com.vmware.vapi.std.errors.not_found")
	assert.Contains(t, argsStrs, "report_undeclared_error")
	var msg2 = messages.(*data.ListValue).List()[1]
	var strMsg2 = msg2.(*data.StructValue)
	var id2, _ = strMsg2.Field("id")
	assert.Equal(t, "mockup.message.Id", id2.(*data.StringValue).Value())
}

func TestReportDeclaredError(t *testing.T) {
	var localProvider = getInitializedLocalProvider()
	var serviceId = "MockupInterface"
	var operationId = "report_declared_error"
	var input = data.NewStructValue(lib.OPERATION_INPUT, make(map[string]data.DataValue))
	input.SetField("param", data.NewIntegerValue(10))
	context := core.NewExecutionContext(common.NewDefaultApplicationContext(), nil)
	methodResult := localProvider.Invoke(serviceId, operationId, input, context)
	assert.Nil(t, methodResult.Output())
	assert.False(t, methodResult.IsSuccess())
	errorValue := methodResult.Error()
	assert.Equal(t, "com.vmware.vapi.std.errors.not_found", errorValue.Name())
}

func TestComputeCheckSumConsistent(t *testing.T) {
	mockApiInterface1 := NewMockApiInterface()
	mockApiInterface2 := NewMockApiInterface()
	localProvider1 := getInitializedLocalProviderWithMockApiInterface(mockApiInterface1)
	localProvider2 := getInitializedLocalProviderWithMockApiInterface(mockApiInterface2)
	checkSum1 := localProvider1.ComputeCheckSum()
	checkSum2 := localProvider2.ComputeCheckSum()
	assert.Equal(t, checkSum1, checkSum2)

	localProvider3 := getInitializedLocalProviderWithMockApiInterface(mockApiInterface2)
	localProvider4 := getInitializedLocalProviderWithMockApiInterface(mockApiInterface1)
	checkSum3 := localProvider3.ComputeCheckSum()
	checkSum4 := localProvider4.ComputeCheckSum()
	assert.Equal(t, checkSum3, checkSum4)
}

func TestComputeCheckSumUnique(t *testing.T) {
	mockApiInterface := NewMockApiInterface()
	localProvider := getInitializedLocalProviderWithMockApiInterface(mockApiInterface)
	checkSum1 := localProvider.ComputeCheckSum()

	mockApiInterface.addMethod("NewMethod")
	localProvider = getInitializedLocalProviderWithMockApiInterface(mockApiInterface)
	checkSum2 := localProvider.ComputeCheckSum()
	assert.NotEqual(t, checkSum1, checkSum2)
}
