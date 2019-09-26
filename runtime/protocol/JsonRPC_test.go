package protocol_test

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/core"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/lib"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol/server/rpc/msg"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/security"
	"reflect"
	"testing"
)

// TODO
// test blob,secret, double, opaque
var fieldString = "fieldString"
var fieldBooleanTrue = "fieldBooleanTrue"
var fieldBooleanFalse = "fieldBooleanFalse"
var fieldInteger = "fieldInteger"
var fieldVoid = "fieldVoid"
var fieldList = "fieldList"
var fieldOptionalUnset = "fieldOptionalUnset"
var fieldOptionalSet = "fieldOptionalSet"
var fieldOptionalStructure = "fieldOptionalStructure"
var fieldStructureStr = "fieldStructure"
var fieldError = "fieldError"
var fieldDouble = "fieldDouble"
var fieldBinary = "fieldBinary"
var fieldSecret = "fieldSecret"

//https://wiki.eng.vmware.com/VAPI/Specs/Protocol/JSON
func getComprehensiveStructure() *data.StructValue {
	var input = data.NewStructValue("comprehensive-struct", make(map[string]data.DataValue))
	//String
	input.SetField(fieldString, data.NewStringValue("hello"))
	//Boolean
	input.SetField(fieldBooleanTrue, data.NewBooleanValue(true))
	input.SetField(fieldBooleanFalse, data.NewBooleanValue(false))
	//Integer
	input.SetField(fieldInteger, data.NewIntegerValue(10))
	//Void
	input.SetField(fieldVoid, data.NewVoidValue())
	//Binary
	input.SetField(fieldBinary, data.NewBlobValue([]byte("hello")))
	//Secret
	input.SetField(fieldSecret, data.NewSecretValue("mypassword"))
	var f1 float64 = 3
	d1 := data.NewDoubleValue(f1)
	input.SetField(fieldDouble, d1)

	//List
	var listValue = data.NewListValue()
	listValue.Add(data.NewStringValue("hello"))
	listValue.Add(data.NewStringValue("world"))
	input.SetField(fieldList, listValue)

	//Structure

	//Structure with fields
	var structValueMap = make(map[string]data.DataValue)
	structValueMap["inc"] = data.NewIntegerValue(int64(10))
	structValueMap["dec"] = data.NewIntegerValue(int64(20))
	var structValue = data.NewStructValue("sample.first.count.properties", structValueMap)
	input.SetField("fieldStructure", structValue)

	//Structure with no fields
	var structValueNoFields = data.NewStructValue("sample.foo", nil)
	input.SetField("field_structure_with_no_fields", structValueNoFields)

	//OptionalValue

	//OptionalValue UnSet
	var optionalValue = data.NewOptionalValue(nil)
	input.SetField(fieldOptionalUnset, optionalValue)

	//OptionalValue Set
	var optionvalValueSet = data.NewOptionalValue(data.NewIntegerValue(int64(10)))
	input.SetField(fieldOptionalSet, optionvalValueSet)

	//OptionalValue with Struct
	var optionalValWithStruct = data.NewOptionalValue(structValue)
	input.SetField("fieldOptionalStructure", optionalValWithStruct)

	//ErrorValue
	var errorValue = data.NewErrorValue("com.vmware.vapi.std.not_found", nil)
	input.SetField(fieldError, errorValue)

	return input
}

func TestRedactSecretValue(t *testing.T) {
	var input = getComprehensiveStructure()
	var jsonRpcEncoder = msg.NewJsonRpcEncoder()
	jsonRpcEncoder.SetRedactSecretFields(true)
	var result, _ = jsonRpcEncoder.Encode(input)
	var jsonRpcDecoder = msg.NewJsonRpcDecoder()
	var output, error = jsonRpcDecoder.DecodeDataValue(string(result))
	if error != nil {
		t.Errorf("Error Decoding DataValue %s", error)
	}
	var outputStructValue = output.(*data.StructValue)
	var fieldSecretValue, _ = outputStructValue.Field(fieldSecret)
	assert.Equal(t, "*redacted*", fieldSecretValue.(*data.SecretValue).Value())
}

func TestDataValueSerializationDeSerialization(t *testing.T) {

	var input = getComprehensiveStructure()
	//Encode Input
	var jsonRpcEncoder = msg.NewJsonRpcEncoder()
	var result, _ = jsonRpcEncoder.Encode(input)

	//Get output by decoding  Input
	var jsonRpcDecoder = msg.NewJsonRpcDecoder()
	var output, error = jsonRpcDecoder.DecodeDataValue(string(result))
	if error != nil {
		t.Errorf("Error Decoding DataValue %s", error)
	}

	var outputStructValue = output.(*data.StructValue)
	//Compare input and output

	//check name of the struct
	assert.Equal(t, input.Name(), outputStructValue.Name())

	//check string
	var outputFieldString, _ = outputStructValue.String("fieldString")
	assert.Equal(t, outputFieldString, "hello")

	//check list
	var outputFieldList, _ = outputStructValue.List("fieldList")
	assert.Equal(t, outputFieldList.List()[0].(*data.StringValue).Value(), "hello")
	assert.Equal(t, outputFieldList.List()[1].(*data.StringValue).Value(), "world")

	//check boolean
	var outputBooleanTrue, _ = outputStructValue.Field(fieldBooleanTrue)
	assert.Equal(t, outputBooleanTrue.(*data.BooleanValue).Value(), true)
	var outputBooleanFalse, _ = outputStructValue.Field(fieldBooleanFalse)
	assert.Equal(t, outputBooleanFalse.(*data.BooleanValue).Value(), false)

	//check integer
	var outputInteger, _ = outputStructValue.Field(fieldInteger)
	assert.Equal(t, outputInteger.(*data.IntegerValue).Value(), int64(10))

	//check double
	var outputDouble, _ = outputStructValue.Field(fieldDouble)
	assert.Equal(t, outputDouble.(*data.DoubleValue).Value(), float64(3))

	//check void
	var outputVoid, _ = outputStructValue.Field(fieldVoid)
	assert.Equal(t, outputVoid, nil)

	//check structure
	var fieldStructure, _ = outputStructValue.Struct(fieldStructureStr)
	var fieldStructureIncValue, _ = fieldStructure.Field("inc")
	var fieldStructureDecValue, _ = fieldStructure.Field("dec")
	assert.Equal(t, fieldStructureIncValue.(*data.IntegerValue).Value(), int64(10))
	assert.Equal(t, fieldStructureDecValue.(*data.IntegerValue).Value(), int64(20))
	assert.Equal(t, fieldStructure.Name(), "sample.first.count.properties")

	var fieldStructWithNoFields, _ = outputStructValue.Struct("field_structure_with_no_fields")
	assert.Equal(t, fieldStructWithNoFields.Name(), "sample.foo")
	assert.Equal(t, len(fieldStructWithNoFields.Fields()), 0)

	//check optional

	//optional unset
	var fieldOptionalTest, _ = outputStructValue.Optional(fieldOptionalUnset)
	assert.Equal(t, fieldOptionalTest.IsSet(), false)

	//optional set
	var fieldOptionalSetTest, _ = outputStructValue.Optional(fieldOptionalSet)
	assert.Equal(t, fieldOptionalSetTest.Value().(*data.IntegerValue).Value(), int64(10))

	//optional set Structure
	var fieldOptionalStructureTest, _ = outputStructValue.Optional(fieldOptionalStructure)
	var optionalStructValue = fieldOptionalStructureTest.Value().(*data.StructValue)
	assert.Equal(t, optionalStructValue.Name(), "sample.first.count.properties")
	fieldStructureIncValue, _ = optionalStructValue.Field("inc")
	fieldStructureDecValue, _ = optionalStructValue.Field("dec")
	assert.Equal(t, fieldStructureIncValue.(*data.IntegerValue).Value(), int64(10))
	assert.Equal(t, fieldStructureDecValue.(*data.IntegerValue).Value(), int64(20))

	//check error
	var fieldErrorValue, _ = outputStructValue.Error(fieldError)
	assert.Equal(t, fieldErrorValue.Name(), "com.vmware.vapi.std.not_found")

	//check binary
	var fieldBinaryValue, _ = outputStructValue.Field(fieldBinary)
	assert.Equal(t, string(fieldBinaryValue.(*data.BlobValue).Value()), "hello")

	//check secret
	var fieldSecretValue, _ = outputStructValue.Field(fieldSecret)
	assert.Equal(t, fieldSecretValue.(*data.SecretValue).Value(), "mypassword")

}

func getErrorValue() *data.ErrorValue {
	var errorValue = data.NewErrorValue("struct0", nil)
	errorValue.SetField("int", data.NewIntegerValue(7))
	errorValue.SetField("optionalSet", data.NewOptionalValue(data.NewStringValue("helloworld")))
	errorValue.SetField("optionalUnset", data.NewOptionalValue(nil))
	return errorValue
}

func getStructValue() *data.StructValue {
	var structValue = data.NewStructValue("struct0", nil)
	structValue.SetField("int", data.NewIntegerValue(7))
	structValue.SetField("optionalSet", data.NewOptionalValue(data.NewStringValue("helloworld")))
	structValue.SetField("optionalUnset", data.NewOptionalValue(nil))
	return structValue
}

func TestMethodResultSerializationDeSerializationError(t *testing.T) {
	var methodResult = core.NewMethodResult(nil, getErrorValue())
	var jsonRpcEncoder = msg.NewJsonRpcEncoder()
	var result, _ = jsonRpcEncoder.Encode(methodResult)
	decodedMethodResult := map[string]interface{}{}
	json.Unmarshal(result, &decodedMethodResult)
	var jsonRpcDecoder = msg.NewJsonRpcDecoder()
	var deserializedMethodResult, error = jsonRpcDecoder.DeSerializeMethodResult(decodedMethodResult)
	if error != nil {
		t.Error("error deserializing method result")
	}
	assert.Equal(t, deserializedMethodResult.IsSuccess(), false)
}

func TestMethodResultSerializationDeSerializationSuccess(t *testing.T) {
	var methodResult = core.NewMethodResult(getStructValue(), nil)
	var jsonRpcEncoder = msg.NewJsonRpcEncoder()
	var result, _ = jsonRpcEncoder.Encode(methodResult)
	decodedMethodResult := map[string]interface{}{}
	json.Unmarshal(result, &decodedMethodResult)
	var jsonRpcDecoder = msg.NewJsonRpcDecoder()
	var deserializedMethodResult, error = jsonRpcDecoder.DeSerializeMethodResult(decodedMethodResult)
	if error != nil {
		t.Error("error deserializing method result")
	}
	assert.Equal(t, deserializedMethodResult.IsSuccess(), true)
}

func TestJsonRpcPositive(t *testing.T) {
	var requests = [3]string{`{"jsonrpc": "2.0", "method": "subtract", "id": null}`,
		`{"jsonrpc": "2.0", "method": "subtract", "id": 3}`,
		`{"jsonrpc": "2.0", "method": "subtract", "id": "string_id"}`}
	var responses = [3]string{"",
		`{"jsonrpc": "2.0", "result": 0, "id": 3}`,
		`{"jsonrpc": "2.0", "result": 0, "id": "string_id"}`}

	var jsonRpcDecoder = msg.NewJsonRpcDecoder()
	var jsonRpcEncoder = msg.NewJsonRpcEncoder()
	for index, requestString := range requests {
		var request, requestDeserializationError = jsonRpcDecoder.DeSerializeRequest(requestString)
		assert.Nil(t, requestDeserializationError)
		var responseString = responses[index]
		if responseString != "" {
			var response, responseDeserializationError = jsonRpcDecoder.DeSerializeResponse(responseString)
			assert.Nil(t, responseDeserializationError)
			assert.False(t, request.Notification())
			assert.Nil(t, request.ValidateResponse(response))
		}
		var _, encodingError = jsonRpcEncoder.Encode(request)
		assert.Nil(t, encodingError)
	}
	var notifications = [1]string{`{"jsonrpc": "2.0", "method": "subtract"}`}
	for _, notificationString := range notifications {
		var request, requestDeserializationError = jsonRpcDecoder.DeSerializeRequest(notificationString)
		assert.Nil(t, requestDeserializationError)
		assert.True(t, request.Notification())
		var _, encodingError = jsonRpcEncoder.Encode(request)
		assert.Nil(t, encodingError)
	}

}

func TestJsonRpcRequestNegative(t *testing.T) {
	var requests = [5]string{``,
		`sjlfskfdl`,
		`{"jsonrpc": "2.0", "method": "foobar, "params": "bar", "baz"]`,
		`{"jsonrpc": "2.0", "method": 1, "params": "bar"}`,
		`{"jsonrpc": "2.0", "method": "a_method", "params": "bar"}`}
	var errors = [5]int64{msg.JSONRPC_PARSE_ERROR,
		msg.JSONRPC_PARSE_ERROR,
		msg.JSONRPC_PARSE_ERROR,
		msg.JSONRPC_INVALID_REQUEST,
		msg.JSONRPC_INVALID_REQUEST}
	var jsonRpcDecoder = msg.NewJsonRpcDecoder()
	for index, requestString := range requests {
		var _, requestDeserializationError = jsonRpcDecoder.DeSerializeRequest(requestString)
		assert.True(t, requestDeserializationError.Code() == errors[index])
	}

	//Negative validate test
	var negValidateRequests = [4]string{`{"jsonrpc": "2.0", "method": "subtract", "params": {"subtrahend": 23, "minuend": 42}, "id": 3}`,
		`{"jsonrpc": "2.0", "method": "subtract", "params":{"subtrahend": 23, "minuend": 42}, "id": 3}`,
		`{"jsonrpc": "2.0", "method": "subtract", "params":{"subtrahend": 23, "minuend": 42}, "id": 3}`,
		`{"jsonrpc": "2.0", "method": "subtract"}`}
	var responses = [4]string{`{"jsonrpc": "1.1", "result": 19, "id": 3}`, //Mismatch version
		`{"jsonrpc": "2.0", "result": 19, "id": 4}`,   //Mismatch id
		`{"jsonrpc": "2.0", "result": 19, "id": "3"}`, //Mismatch id type
		`{"jsonrpc": "2.0", "result": 0}`}             //Notification should have no response
	var validationErrors = [4]int64{msg.JSONRPC_INVALID_PARAMS,
		msg.JSONRPC_INVALID_PARAMS,
		msg.JSONRPC_INVALID_PARAMS,
		msg.JSONRPC_INVALID_PARAMS}

	for index, requestString := range negValidateRequests {
		var request, requestDeserializationError = jsonRpcDecoder.DeSerializeRequest(requestString)
		assert.Nil(t, requestDeserializationError)
		var responseString = responses[index]
		var response, responseDeserializationError = jsonRpcDecoder.DeSerializeResponse(responseString)
		if responseDeserializationError != nil {
			assert.Equal(t, responseDeserializationError.Code(), validationErrors[index])
			continue
		}
		var validationError = request.ValidateResponse(response)
		assert.Equal(t, validationError.Code(), validationErrors[index])
	}
}

func TestJsonRpcResponsePositive(t *testing.T) {
	var error_responses = []string{`{"jsonrpc": "2.0", "error": {"code": -32600}, "id": null}`,
		`{"jsonrpc": "2.0", "error": {"code": -32600, "data": "somedata"}, "id": null}`,
		`{"jsonrpc": "2.0", "error": {"code": -32600, "data": {"key": "somedata"}}, "id": null}`,
		`{"jsonrpc": "2.0", "error": {"code": -32600, "message": "Invalid Request."}, "id": null}`,
		`{"jsonrpc": "2.0", "error": {"code": -32700, "message": "Parse error."}, "id": null}`,
		`{"jsonrpc": "2.0", "error": {"code": -32601, "message": "Method not found."}, "id": "1"}`,
		`{"jsonrpc": "2.0", "error": {"code": -32602, "message": "Invalid params."}, "id": "1"}`,
		`{"jsonrpc": "2.0", "error": {"code": -32603, "message": "Internal error."}, "id": 1}`}

	var error_codes = []int64{msg.JSONRPC_INVALID_REQUEST,
		msg.JSONRPC_INVALID_REQUEST,
		msg.JSONRPC_INVALID_REQUEST,
		msg.JSONRPC_INVALID_REQUEST,
		msg.JSONRPC_PARSE_ERROR,
		msg.JSONRPC_METHOD_NOT_FOUND,
		msg.JSONRPC_INVALID_PARAMS,
		msg.JSONRPC_INTERNAL_ERROR}

	var jsonRpcDecoder = msg.NewJsonRpcDecoder()

	for index, error_response := range error_responses {
		var response, deserializationError = jsonRpcDecoder.DeSerializeResponse(error_response)
		assert.Nil(t, deserializationError)
		code, _ := response.Error()["code"].(json.Number).Int64()
		assert.Equal(t, code, error_codes[index])
	}
}

func TestJsonRpcResponseNegative(t *testing.T) {
	var responses = []string{``,
		`sjlfskfdl`,
		`{"jsonrpc": "2.0", "result": "foobar, "error": "bar"}`,
		`{"jsonrpc": "2.0", "result": "foobar, "error": "bar", "id": null}`,
		`{"jsonrpc": "2.0", "result": "foobar, "error": {"code": -32603}, "id": null}`,
		`{"result": 19, "id": 3}`}
	var errorCodes = []int64{msg.JSONRPC_PARSE_ERROR,
		msg.JSONRPC_PARSE_ERROR,
		msg.JSONRPC_PARSE_ERROR,
		msg.JSONRPC_PARSE_ERROR,
		msg.JSONRPC_PARSE_ERROR,
		msg.JSONRPC_INVALID_PARAMS}

	var jsonRpcDecoder = msg.NewJsonRpcDecoder()
	for index, responseStr := range responses {
		var _, deserializationError = jsonRpcDecoder.DeSerializeResponse(responseStr)
		assert.NotNil(t, deserializationError)
		assert.Equal(t, deserializationError.Code(), errorCodes[index])
	}
}

var validRequest = []byte(`{
   "id":"2",
   "jsonrpc":"2.0",
   "method":"invoke",
   "params":{
      "ctx":{
          "securityCtx":{
             "schemeId": "com.vmware.vapi.std.security.session_id",
             "sessionId": "xyz"
           },
          "appCtx":{
             "opId": "123abc"
          }
      },
     "input":{
         "STRUCTURE":{
             "operation-input": {
                 "vm_id": "vm-123"
             }
          }
      },
      "operationId":"get",
      "serviceId":"sample.first.toy_VM"
   }
}`)

func TestJsonRpc20RequestSerializer_MarshalJSON(t *testing.T) {
	appContextMap := map[string]*string{}
	opId := "123abc"
	appContextMap[lib.OPID] = &opId
	ctx := core.NewExecutionContext(core.NewApplicationContext(appContextMap), security.NewSessionSecurityContext("xyz"))
	params := map[string]interface{}{
		lib.REQUEST_INPUT:        getInput(),
		lib.EXECUTION_CONTEXT:    msg.NewExecutionContextSerializer(ctx),
		lib.REQUEST_OPERATION_ID: "get",
		lib.REQUEST_SERVICE_ID:   "sample.first.toy_VM",
	}
	requestId := "2"
	var jsonRpcRequest = msg.NewJsonRpc20Request(lib.JSONRPC_VERSION, lib.JSONRPC_INVOKE, params, requestId, false)

	var jsonRpcRequestSerializer = msg.NewJsonRpc20RequestSerializer(jsonRpcRequest)
	var result, _ = jsonRpcRequestSerializer.MarshalJSON()
	actualJson := make(map[string]interface{})
	json.Unmarshal(result, &actualJson)
	expectedJson := make(map[string]interface{})
	json.Unmarshal(validRequest, &expectedJson)
	assert.Equal(t, expectedJson, actualJson)
}

func getInput() map[string]interface{} {
	typeConverter := bindings.NewTypeConverter()
	fields := map[string]bindings.BindingType{
		"vm_id": bindings.NewStringType(),
	}
	fieldNameMap := map[string]string{
		"vm_id": "VmId",
	}
	inputBindingType := bindings.NewStructType("operation-input", fields,
		reflect.TypeOf(data.StructValue{}), fieldNameMap, nil)
	sv := bindings.NewStructValueBuilder(inputBindingType, typeConverter)
	sv.AddStructField("VmId", "vm-123")
	inputDataValue, _ := sv.GetStructValue()
	var encodedInput, _ = msg.NewJsonRpcEncoder().Encode(inputDataValue)
	encodedInputMap := make(map[string]interface{})
	json.Unmarshal(encodedInput, &encodedInputMap)
	return encodedInputMap
}
