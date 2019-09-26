package rest_test

/* **********************************************************
 * Copyright 2019 VMware, Inc.  All rights reserved.
 *      -- VMware Confidential
 * **********************************************************/

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/core"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol/server/rest"
	"strconv"
	"testing"
)

func setResponseStatus(statusCode int, status string, err *data.ErrorValue) (int, string, []error) {
	// Pass in method result structure
	fields := make(map[string]data.DataValue)
	output := data.NewStructValue("testresult", fields)

	result := core.NewMethodResult(output, err)

	statusCodeMap := make(map[string]int)
	if statusCode != 0 {
		statusCodeMap[status] = statusCode
	}

	return rest.CreateResponseStatus(result, statusCodeMap)
}

func TestCreateResponseStatus_With_Defined_OK(t *testing.T) {
	statusCode, status, err := setResponseStatus(201, "SUCCESS", nil)

	assert.Nil(t, err)
	assert.Equal(t, 201, statusCode)
	assert.Equal(t, "Created", status)
}

func TestCreateResponseStatus_With_Custom_OK(t *testing.T) {
	statusCode, status, err := setResponseStatus(299, "SUCCESS", nil)

	assert.Nil(t, err)
	assert.Equal(t, 299, statusCode)
	assert.Equal(t, "OK", status)
}

func TestCreateResponseStatus_Without_OK(t *testing.T) {
	statusCode, status, err := setResponseStatus(0, "", nil)

	assert.Nil(t, err)
	assert.Equal(t, 200, statusCode)
	assert.Equal(t, "OK", status)
}

func TestCreateResponseStatus_With_Defined_Error(t *testing.T) {
	statusCode, status, err := setResponseStatus(403,
		"Forbidden",
		data.NewErrorValue("Forbidden", nil))

	assert.Nil(t, err)
	assert.Equal(t, 403, statusCode)
	assert.Equal(t, "Forbidden", status)
}

func TestCreateResponseStatus_With_Custom_Error(t *testing.T) {
	statusCode, status, err := setResponseStatus(599,
		"MyError",
		data.NewErrorValue("MyError", nil))

	assert.Nil(t, err)
	assert.Equal(t, 599, statusCode)
	assert.Equal(t, "Internal Server Error", status)
}

func TestCreateResponseStatus_Without_Error_Definition(t *testing.T) {
	statusCode, status, err := setResponseStatus(0,
		"",
		data.NewErrorValue("MyError", nil))

	assert.Equal(t,
		"Http status 'MyError' not supported",
		err[0].Error())
	assert.Equal(t, 500, statusCode)
	assert.Equal(t, "Internal Server Error", status)
}

func TestCreateResponseHeader_OK(t *testing.T) {
	// Pass in method result structure
	fields := make(map[string]data.DataValue)
	fields["stringField"] = data.NewStringValue("sfield")
	fields["booleanField"] = data.NewBooleanValue(false)
	fields["intField"] = data.NewIntegerValue(123)
	fields["doubleField"] = data.NewDoubleValue(456.789)
	fields["blobField"] = data.NewBlobValue([]byte("VGhpcyBpcyB0ZXN0Cg"))
	output := data.NewStructValue("testresult", fields)

	result := core.NewMethodResult(output, nil)

	resultToHeaderMap := make(map[string]string)
	// Define result to header name map
	resultToHeaderMap["stringField"] = "X-StringField"
	resultToHeaderMap["booleanField"] = "X-BooleanField"
	resultToHeaderMap["intField"] = "X-IntField"
	resultToHeaderMap["doubleField"] = "X-DoubleField"
	resultToHeaderMap["blobField"] = "X-BlobField"

	header, err := rest.CreateResponseHeader(result, resultToHeaderMap, map[string]string{})

	assert.Nil(t, err)
	assert.Equal(t, "sfield", header.Get("X-StringField"))
	assert.Equal(t, "false", header.Get("X-BooleanField"))
	assert.Equal(t, "123", header.Get("X-IntField"))
	assert.Equal(t, "456.789", header.Get("X-DoubleField"))
	assert.Equal(t, "VGhpcyBpcyB0ZXN0Cg", header.Get("X-BlobField"))
}

func TestCreateResponseHeader_NestedStruct(t *testing.T) {
	// Pass in method result structure
	nestedFields := make(map[string]data.DataValue)
	nestedFields["stringField"] = data.NewStringValue("sfield")
	nestedFields["booleanField"] = data.NewBooleanValue(false)
	nestedFields["intField"] = data.NewIntegerValue(123)
	nestedStruct := data.NewStructValue("nestedStruct", nestedFields)
	fields := make(map[string]data.DataValue)
	fields["nestedStructField"] = nestedStruct
	output := data.NewStructValue("testresult", fields)

	result := core.NewMethodResult(output, nil)

	resultToHeaderMap := make(map[string]string)
	// Define result to header name map
	resultToHeaderMap["nestedStructField"] = "X-StructField"
	resultToHeaderMap["stringField"] = "X-StringField"
	resultToHeaderMap["booleanField"] = "X-BooleanField"
	resultToHeaderMap["intField"] = "X-IntField"

	header, err := rest.CreateResponseHeader(result, resultToHeaderMap, map[string]string{})

	assert.Nil(t, err)
	assert.Equal(t, "sfield", header.Get("X-StringField"))
	assert.Equal(t, "false", header.Get("X-BooleanField"))
	assert.Equal(t, "123", header.Get("X-IntField"))
	assert.Equal(t, "", header.Get("X-StructField"))
}

func TestCreateResponseHeader_NotAnnotation(t *testing.T) {
	// Pass in method result structure
	fields := make(map[string]data.DataValue)
	fields["stringField"] = data.NewStringValue("sfield")
	fields["booleanField"] = data.NewBooleanValue(false)
	fields["intField"] = data.NewIntegerValue(123)
	fields["doubleField"] = data.NewDoubleValue(456.789)
	output := data.NewStructValue("testresult", fields)

	result := core.NewMethodResult(output, nil)

	resultToHeaderMap := make(map[string]string)
	// Define result to header name map
	resultToHeaderMap["booleanField"] = "X-BooleanField"
	resultToHeaderMap["doubleField"] = "X-DoubleField"

	header, err := rest.CreateResponseHeader(result, resultToHeaderMap, map[string]string{})

	assert.Nil(t, err)
	assert.Equal(t, "", header.Get("X-StringField"))
	assert.Equal(t, "false", header.Get("X-BooleanField"))
	assert.Equal(t, "", header.Get("X-IntField"))
	assert.Equal(t, "456.789", header.Get("X-DoubleField"))
}

func TestCreateResponseHeader_NotSupportType(t *testing.T) {
	// Pass in method result structure
	fields := make(map[string]data.DataValue)
	fields["secretField"] = data.NewSecretValue("secret")
	output := data.NewStructValue("testresult", fields)

	result := core.NewMethodResult(output, nil)

	resultToHeaderMap := make(map[string]string)
	// Define result to header name map
	resultToHeaderMap["secretField"] = "X-SecreteField"

	header, err := rest.CreateResponseHeader(result, resultToHeaderMap, map[string]string{})

	assert.Equal(t,
		"Response data type 'SECRET' not supported for field 'X-SecreteField'",
		err[0].Error())
	assert.Equal(t, "", header.Get("X-SecreteField"))
}

func TestCreateResponseHeader_List(t *testing.T) {
	// Pass in method result structure
	fields := make(map[string]data.DataValue)
	listVal := data.NewListValue()
	listVal.Add(data.NewStringValue("hv1"))
	listVal.Add(data.NewStringValue("hv2"))
	listVal.Add(data.NewStringValue("hv3"))
	fields["listField"] = listVal

	output := data.NewStructValue("testresult", fields)

	result := core.NewMethodResult(output, nil)

	resultToHeaderMap := make(map[string]string)
	// Define result to header name map
	resultToHeaderMap["listField"] = "X-ListField"

	header, err := rest.CreateResponseHeader(result, resultToHeaderMap, map[string]string{})

	assert.Nil(t, err)
	actual := header.Get("X-ListField")
	assert.Equal(t, "hv1", actual)

	assert.Equal(t, "hv1", header["X-Listfield"][0])
	assert.Equal(t, "hv2", header["X-Listfield"][1])
	assert.Equal(t, "hv3", header["X-Listfield"][2])
}

func TestCreateResponseHeader_NotSupportTypeUnderList(t *testing.T) {
	// Pass in method result structure
	fields := make(map[string]data.DataValue)
	listVal := data.NewListValue()
	listVal.Add(data.NewOptionalValue(data.NewIntegerValue(123)))
	fields["listField"] = listVal
	output := data.NewStructValue("testresult", fields)

	result := core.NewMethodResult(output, nil)

	resultToHeaderMap := make(map[string]string)
	// Define result to header name map
	resultToHeaderMap["listField"] = "X-ListField"

	header, err := rest.CreateResponseHeader(result, resultToHeaderMap, map[string]string{})

	assert.Equal(t,
		"Response data type 'LIST<OPTIONAL>' not supported for field 'X-ListField'",
		err[0].Error())
	assert.Equal(t, "", header.Get("X-ListField"))
}

func TestCreateResponseHeader_Optional(t *testing.T) {
	// Pass in method result structure
	fields := make(map[string]data.DataValue)
	optVal := data.NewOptionalValue(data.NewStringValue("hv1"))
	fields["optionalField"] = optVal

	output := data.NewStructValue("testresult", fields)

	result := core.NewMethodResult(output, nil)

	resultToHeaderMap := make(map[string]string)
	// Define result to header name map
	resultToHeaderMap["optionalField"] = "X-OptionalField"

	header, err := rest.CreateResponseHeader(result, resultToHeaderMap, map[string]string{})

	assert.Nil(t, err)
	actual := header.Get("X-OptionalField")
	assert.Equal(t, "hv1", actual)
}

func TestCreateResponseHeader_NotSupportTypeUnderOptional(t *testing.T) {
	// Pass in method result structure
	fields := make(map[string]data.DataValue)
	optVal := data.NewOptionalValue(data.NewErrorValue("err", nil))
	fields["optionalField"] = optVal
	output := data.NewStructValue("testresult", fields)

	result := core.NewMethodResult(output, nil)

	resultToHeaderMap := make(map[string]string)
	// Define result to header name map
	resultToHeaderMap["optionalField"] = "X-OptionalField"

	header, err := rest.CreateResponseHeader(result, resultToHeaderMap, map[string]string{})

	assert.Equal(t,
		"Response data type 'ERROR' not supported for field 'X-OptionalField'",
		err[0].Error())
	assert.Equal(t, "", header.Get("X-OptionalField"))
}

func TestCreateResponseHeader_OptionalList(t *testing.T) {
	// Pass in method result structure
	fields := make(map[string]data.DataValue)
	listVal := data.NewListValue()
	listVal.Add(data.NewStringValue("hv1"))
	listVal.Add(data.NewStringValue("hv2"))
	listVal.Add(data.NewStringValue("hv3"))
	optVal := data.NewOptionalValue(listVal)
	fields["listField"] = optVal

	output := data.NewStructValue("testresult", fields)

	result := core.NewMethodResult(output, nil)

	resultToHeaderMap := make(map[string]string)
	// Define result to header name map
	resultToHeaderMap["listField"] = "X-ListField"

	header, err := rest.CreateResponseHeader(result, resultToHeaderMap, map[string]string{})

	assert.Nil(t, err)
	actual := header.Get("X-ListField")
	assert.Equal(t, "hv1", actual)

	assert.Equal(t, "hv1", header["X-Listfield"][0])
	assert.Equal(t, "hv2", header["X-Listfield"][1])
	assert.Equal(t, "hv3", header["X-Listfield"][2])
}

func TestCreateResponseHeader_Error(t *testing.T) {
	// Pass in method result structure
	fields := make(map[string]data.DataValue)
	fields["reason"] = data.NewStringValue("internal-error")
	fields["code"] = data.NewIntegerValue(500)
	errorDv := data.NewErrorValue("testerror", fields)

	result := core.NewMethodResult(nil, errorDv)

	// Define error to header name map
	errorToHeaderMap := make(map[string]string)
	errorToHeaderMap["reason"] = "X-Reason"
	errorToHeaderMap["code"] = "X-Code"

	header, err := rest.CreateResponseHeader(result, map[string]string{}, errorToHeaderMap)

	assert.Nil(t, err)
	assert.Equal(t, "internal-error", header.Get("X-Reason"))
	assert.Equal(t, "500", header.Get("X-Code"))
}

func TestCreateResponseHeader_MixResultAndError(t *testing.T) {
	// Pass in method result structure
	fieldsStruct := make(map[string]data.DataValue)
	fieldsStruct["common"] = data.NewStringValue("common-in-struct")
	fieldsStruct["booleanField"] = data.NewBooleanValue(false)
	respDV := data.NewStructValue("testresult", fieldsStruct)

	fieldsError := make(map[string]data.DataValue)
	fieldsError["common"] = data.NewStringValue("commmon-in-error")
	fieldsError["reason"] = data.NewStringValue("internal-error")
	fieldsError["code"] = data.NewIntegerValue(500)
	errorDv := data.NewErrorValue("testerror", fieldsError)

	result := core.NewMethodResult(respDV, errorDv)

	// Define result to header name map
	resultToHeaderMap := make(map[string]string)
	resultToHeaderMap["common"] = "X-Common"
	resultToHeaderMap["booleanField"] = "X-BooleanField"

	// Define error to header name map
	errorToHeaderMap := make(map[string]string)
	errorToHeaderMap["common"] = "X-Common"
	errorToHeaderMap["reason"] = "X-Reason"
	errorToHeaderMap["code"] = "X-Code"

	header, err := rest.CreateResponseHeader(result, map[string]string{}, errorToHeaderMap)

	assert.Nil(t, err)
	assert.Equal(t, "commmon-in-error", header.Get("X-Common"))
	assert.Equal(t, "internal-error", header.Get("X-Reason"))
	assert.Equal(t, "500", header.Get("X-Code"))
	assert.Equal(t, "", header.Get("X-BooleanField"))
}

func TestCreateResponseBody_PrimitiveType(t *testing.T) {
	result := core.NewMethodResult(data.NewStringValue("stringResult"), nil)
	body, err := rest.CreateResponseBody(result, []string{}, []string{})
	assert.Nil(t, err)
	actualStringVal, _ := strconv.Unquote(body)
	assert.Equal(t, "stringResult", actualStringVal)

	result = core.NewMethodResult(data.NewIntegerValue(123), nil)
	body, err = rest.CreateResponseBody(result, []string{}, []string{})
	assert.Nil(t, err)
	actualIntVal, _ := strconv.ParseInt(body, 10, 64)
	assert.Equal(t, int64(123), actualIntVal)

	result = core.NewMethodResult(data.NewDoubleValue(1.23), nil)
	body, err = rest.CreateResponseBody(result, []string{}, []string{})
	assert.Nil(t, err)
	actualDoubleVal, _ := strconv.ParseFloat(body, 64)
	assert.Equal(t, 1.23, actualDoubleVal)

	result = core.NewMethodResult(data.NewBooleanValue(false), nil)
	body, err = rest.CreateResponseBody(result, []string{}, []string{})
	assert.Nil(t, err)
	actualBoolVal, _ := strconv.ParseBool(body)
	assert.Equal(t, false, actualBoolVal)

	result = core.NewMethodResult(data.NewVoidDefinition(), nil)
	body, err = rest.CreateResponseBody(result, []string{}, []string{})
	assert.Nil(t, err)
	assert.Equal(t, "null", body)

	result = core.NewMethodResult(data.NewBlobValue([]byte("VGhpcyBpcyB0ZXN0Cg")), nil)
	body, err = rest.CreateResponseBody(result, []string{}, []string{})
	assert.Nil(t, err)
	actualStringVal, _ = strconv.Unquote(body)
	assert.Equal(t, "VkdocGN5QnBjeUIwWlhOMENn", actualStringVal)
}

func TestCreateResponseBody_ListType(t *testing.T) {
	listVal := data.NewListValue()
	listVal.Add(data.NewDoubleValue(1.2))
	listVal.Add(data.NewDoubleValue(3.4))
	result := core.NewMethodResult(listVal, nil)

	body, err := rest.CreateResponseBody(result, []string{}, []string{})
	assert.Nil(t, err)
	assert.Equal(t, "[1.2,3.4]", body)
}

func TestCreateResponseBody_OptionalType(t *testing.T) {
	optionalVal := data.NewOptionalValue(data.NewIntegerValue(123))
	result := core.NewMethodResult(optionalVal, nil)
	body, err := rest.CreateResponseBody(result, []string{}, []string{})
	assert.Nil(t, err)
	actualOptionalIntVal, _ := strconv.ParseInt(body, 10, 64)
	assert.Equal(t, int64(123), actualOptionalIntVal)

	optionalVal = data.NewOptionalValue(nil)
	result = core.NewMethodResult(optionalVal, nil)
	body, err = rest.CreateResponseBody(result, []string{}, []string{})
	assert.Nil(t, err)
	assert.Equal(t, "null", body)

	listVal := data.NewListValue()
	listVal.Add(data.NewDoubleValue(1.2))
	listVal.Add(data.NewDoubleValue(3.4))
	optionalVal = data.NewOptionalValue(listVal)
	result = core.NewMethodResult(optionalVal, nil)
	body, err = rest.CreateResponseBody(result, []string{}, []string{})
	assert.Nil(t, err)
	assert.Equal(t, "[1.2,3.4]", body)
}

func TestCreateResponseBody_StructureType(t *testing.T) {
	fields := make(map[string]data.DataValue)
	fields["stringField"] = data.NewStringValue("sfield")
	fields["booleanField"] = data.NewBooleanValue(false)
	fields["intField"] = data.NewIntegerValue(123)
	fields["doubleField"] = data.NewDoubleValue(456.789)
	fields["blobField"] = data.NewBlobValue([]byte("VGhpcyBpcyB0ZXN0Cg"))
	respDV := data.NewStructValue("testresult", fields)
	result := core.NewMethodResult(respDV, nil)

	bodyFieldList := []string{"stringField", "booleanField", "intField",
		"doubleField", "blobField"}

	body, err := rest.CreateResponseBody(result, bodyFieldList, []string{})
	assert.Nil(t, err)
	var jsonBody map[string]interface{}
	json.Unmarshal([]byte(body), &jsonBody)
	assert.Equal(t, "sfield", jsonBody["stringField"])
	assert.Equal(t, false, jsonBody["booleanField"])
	assert.Equal(t, float64(123), jsonBody["intField"])
	assert.Equal(t, 456.789, jsonBody["doubleField"])
	assert.Equal(t, "VkdocGN5QnBjeUIwWlhOMENn", jsonBody["blobField"])

	bodyFieldList = []string{"stringField", "booleanField", "nonexist"}

	body, err = rest.CreateResponseBody(result, bodyFieldList, []string{})
	assert.Nil(t, err)
	jsonBody = make(map[string]interface{})
	json.Unmarshal([]byte(body), &jsonBody)
	assert.Equal(t, "sfield", jsonBody["stringField"])
	assert.Equal(t, false, jsonBody["booleanField"])
	_, exist := jsonBody["nonexist"]
	assert.Equal(t, false, exist)
	_, exist = jsonBody["doubleField"]
	assert.Equal(t, false, exist)
}

func TestCreateResponseBody_ErrorType(t *testing.T) {
	fields := make(map[string]data.DataValue)
	fields["reason"] = data.NewStringValue("internal-error")
	fields["code"] = data.NewIntegerValue(500)
	errorDv := data.NewErrorValue("testerror", fields)

	result := core.NewMethodResult(nil, errorDv)

	errorFieldList := []string{"reason", "code", "nonexist"}

	body, err := rest.CreateResponseBody(result, []string{}, errorFieldList)
	assert.Nil(t, err)
	var jsonBody map[string]interface{}
	json.Unmarshal([]byte(body), &jsonBody)
	assert.Equal(t, "internal-error", jsonBody["reason"])
	assert.Equal(t, float64(500), jsonBody["code"])
	_, exist := jsonBody["nonexist"]
	assert.Equal(t, false, exist)
}

func TestCreateResponseBody_MixStructAndError(t *testing.T) {
	fieldsStruct := make(map[string]data.DataValue)
	fieldsStruct["common"] = data.NewStringValue("common-in-struct")
	fieldsStruct["booleanField"] = data.NewBooleanValue(false)
	respDV := data.NewStructValue("testresult", fieldsStruct)

	fieldsError := make(map[string]data.DataValue)
	fieldsError["common"] = data.NewStringValue("commmon-in-error")
	fieldsError["reason"] = data.NewStringValue("internal-error")
	fieldsError["code"] = data.NewIntegerValue(500)
	errorDv := data.NewErrorValue("testerror", fieldsError)

	result := core.NewMethodResult(respDV, errorDv)

	bodyFieldList := []string{"common", "booleanField", "nonexist"}
	errorFieldList := []string{"common", "reason", "code", "nonexist"}

	// If method result is not successed, only ErrorValue can be packed
	// in response body
	body, err := rest.CreateResponseBody(result, bodyFieldList, errorFieldList)
	assert.Nil(t, err)
	var jsonBody map[string]interface{}
	json.Unmarshal([]byte(body), &jsonBody)
	assert.Equal(t, "commmon-in-error", jsonBody["common"])
	assert.Equal(t, "internal-error", jsonBody["reason"])
	assert.Equal(t, float64(500), jsonBody["code"])
	_, exist := jsonBody["nonexist"]
	assert.Equal(t, false, exist)
	_, exist = jsonBody["booleanField"]
	assert.Equal(t, false, exist)
}