package rest_test

/* **********************************************************
 * Copyright 2019 VMware, Inc.  All rights reserved.
 *      -- VMware Confidential
 * **********************************************************/

import (
	"bytes"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data/serializers/cleanjson"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol/server/rest"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

type testData struct {
	pathParamNameMap   map[string]string
	pathParamTypeMap   map[string]bindings.BindingType
	queryParamNameMap  map[string]string
	queryParamTypeMap  map[string]bindings.BindingType
	headerParamNameMap map[string]string
	headerParamTypeMap map[string]bindings.BindingType
}

func setup() testData {
	pathParamNameMap := map[string]string{
		"path1": "param1",
		"path2": "param2",
		"path3": "param3",
		"path4": "param4",
		"path5": "param5",
	}
	pathParamTypeMap := map[string]bindings.BindingType{
		"param1": bindings.NewStringType(),
		"param2": bindings.NewIdType(nil, ""),
		"param3": bindings.NewIntegerType(),
		"param4": bindings.NewBooleanType(),
		"param5": bindings.NewStringType(),
	}

	queryParamNameMap := map[string]string{
		"query1":  "param6",
		"query2":  "param7",
		"query3":  "param8",
		"query4":  "param9",
		"query5":  "param10",
		"query6":  "param16",
		"query7":  "param17",
		"query8":  "param18",
		"query9":  "param19",
		"query10": "param20",
		"query11": "param21",
		"query12": "param22",
	}
	integerListType := bindings.NewListType(bindings.NewIntegerType(), nil)
	stringElementListType := bindings.NewListType(bindings.NewStringType(), nil)
	optionalStringListType := bindings.NewOptionalType(stringElementListType)
	stringListListType := bindings.NewListType(bindings.NewListType(bindings.NewStringType(), nil), nil)
	stringListSetType := bindings.NewListType(bindings.NewSetType(bindings.NewStringType(), nil), nil)
	queryParamTypeMap := map[string]bindings.BindingType{
		"param6":  bindings.NewStringType(),
		"param7":  bindings.NewBooleanType(),
		"param8":  bindings.NewIntegerType(),
		"param9":  bindings.NewDoubleType(),
		"param10": integerListType,
		"param16": optionalStringListType,
		"param17": bindings.NewIdType([]string{"resource1, resourcex2"}, "holder"),
		"param18": bindings.NewBlobType(),
		"param19": bindings.NewSetType(bindings.NewStringType(), nil),
		"param20": bindings.NewSecretType(),
		"param21": stringListListType,
		"param22": stringListSetType,
	}

	headerParamNameMap := map[string]string{
		"X-Header1": "param11",
		"X-Header2": "param12",
		"X-Header3": "param13",
		"X-Header4": "param14",
		"X-Header5": "param15",
	}
	stringListType := bindings.NewListType(bindings.NewStringType(), nil)
	optionalIntegerType := bindings.NewOptionalType(bindings.NewIntegerType())
	headerParamTypeMap := map[string]bindings.BindingType{
		"param11": bindings.NewStringType(),
		"param12": stringListType,
		"param13": bindings.NewIntegerType(),
		"param14": bindings.NewBooleanType(),
		"param15": optionalIntegerType,
	}

	return testData{pathParamNameMap,
		pathParamTypeMap,
		queryParamNameMap,
		queryParamTypeMap,
		headerParamNameMap,
		headerParamTypeMap}
}

func TestParsePathParams(t *testing.T) {
	// Construct annotation name to param name map and
	// param name to param type map
	setupData := setup()
	paramNameMap := setupData.pathParamNameMap
	paramTypeMap := setupData.pathParamTypeMap

	// Construct http.Request with url path variables
	r, _ := http.NewRequest("GET", "/api/test", nil)
	r = mux.SetURLVars(r, map[string]string{
		"path1":        "pv1",
		"path2":        "id",
		"path3":        "123",
		"path4":        "true",
		"path5":        "pv5",
		"notannotated": "abc",
	})

	// Parse the http url path param based on http.Request
	paramValueMap, err := rest.ParsePathParams(r, paramNameMap, paramTypeMap)
	assert.Nil(t, err)
	assert.Equal(t, data.NewStringValue("pv1"), paramValueMap["param1"])
	assert.Equal(t, data.NewStringValue("id"), paramValueMap["param2"])
	assert.Equal(t, data.NewIntegerValue(123), paramValueMap["param3"])
	assert.Equal(t, data.NewBooleanValue(true), paramValueMap["param4"])
	assert.Equal(t, data.NewStringValue("pv5"), paramValueMap["param5"])
	_, found := paramValueMap["notannotated"]
	assert.Equal(t, false, found)
}

func TestParsePathParams_NotAnnotated(t *testing.T) {
	// Construct annotation name to param name map and
	// param name to param type map
	setupData := setup()
	paramNameMap := setupData.pathParamNameMap
	paramTypeMap := setupData.pathParamTypeMap

	// Construct http.Request with url path variables
	r, _ := http.NewRequest("GET", "/api/test", nil)
	r = mux.SetURLVars(r, map[string]string{
		"notAnnotated": "pv1",
		"path2":        "pv2",
	})

	// Parse the http url path param based on http.Request
	paramValueMap, err := rest.ParsePathParams(r, paramNameMap, paramTypeMap)
	assert.Nil(t, err)
	assert.Equal(t, data.NewStringValue("pv2"), paramValueMap["param2"])
	_, found := paramValueMap["notAnnotated"]
	assert.Equal(t, false, found)
}

func TestParsePathParams_MissingNameMap(t *testing.T) {
	// Construct annotation name to param name map and
	// param name to param type map
	paramNameMap := make(map[string]string)
	paramTypeMap := make(map[string]bindings.BindingType)

	// Construct http.Request with http headers
	r, _ := http.NewRequest("GET", "/api/test", nil)

	// Parse the http header param based on http.Request
	_, err := rest.ParsePathParams(r, paramNameMap, paramTypeMap)
	assert.Nil(t, err)
}

func TestParseHeaderParams(t *testing.T) {
	// Construct annotation name to param name map and
	// param name to param type map
	setupData := setup()
	paramNameMap := setupData.headerParamNameMap
	paramTypeMap := setupData.headerParamTypeMap

	// Construct http.Request with http headers
	r, _ := http.NewRequest("GET", "/api/test", nil)
	r.Header.Add("X-Header1", "hv1")
	r.Header.Add("X-Header2", "hv21")
	r.Header.Add("X-Header2", "hv22")
	r.Header.Add("X-Header3", "123")
	r.Header.Add("X-Header4", "false")
	r.Header.Add("X-Header5", "987")

	// Parse the http header param based on http.Request
	paramValueMap, err := rest.ParseHeaderParams(r, paramNameMap, paramTypeMap)
	assert.Nil(t, err)
	assert.Equal(t, data.NewStringValue("hv1"), paramValueMap["param11"])
	listval := data.NewListValue()
	listval.Add(data.NewStringValue("hv21"))
	listval.Add(data.NewStringValue("hv22"))
	assert.Equal(t, listval, paramValueMap["param12"])
	assert.Equal(t, data.NewIntegerValue(123), paramValueMap["param13"])
	assert.Equal(t, data.NewBooleanValue(false), paramValueMap["param14"])
	assert.Equal(t, data.NewOptionalValue(data.NewIntegerValue(987)), paramValueMap["param15"])
}

func TestParseHeaderParams_List(t *testing.T) {
	// Construct annotation name to param name map and
	// param name to param type map
	setupData := setup()
	paramNameMap := setupData.headerParamNameMap
	paramTypeMap := setupData.headerParamTypeMap

	// Construct http.Request with http headers
	r, _ := http.NewRequest("GET", "/api/test", nil)
	// Note: Set multiple value to the same header, it keep the original order
	r.Header.Add("X-Header1", "hv11")
	r.Header.Add("X-Header1", "hv12")
	r.Header.Add("X-Header2", "hv21")

	// Parse the http header param based on http.Request
	paramValueMap, err := rest.ParseHeaderParams(r, paramNameMap, paramTypeMap)
	assert.Nil(t, err)
	assert.Equal(t, data.NewStringValue("hv11"), paramValueMap["param11"])
	listval := data.NewListValue()
	listval.Add(data.NewStringValue("hv21"))
	assert.Equal(t, listval, paramValueMap["param12"])
}

func TestParseHeaderParams_InvalidType(t *testing.T) {
	// Construct annotation name to param name map and
	// param name to param type map
	setupData := setup()
	paramNameMap := setupData.headerParamNameMap
	paramTypeMap := setupData.headerParamTypeMap

	// Construct http.Request with http headers
	r, _ := http.NewRequest("GET", "/api/test", nil)
	r.Header.Add("X-Header3", "notanumber")

	// Parse the http header param based on http.Request
	paramValueMap, err := rest.ParseHeaderParams(r, paramNameMap, paramTypeMap)
	assert.Equal(t,
		"Invalid value for request parameter 'X-Header3'. Expected a value of type 'Int', but parsing failed because of 'strconv.ParseInt: parsing \"notanumber\": invalid syntax'",
		err[0].Error())
	assert.Equal(t, nil, paramValueMap["param15"])
}

func TestParseHeaderParams_NotAnnotated(t *testing.T) {
	// Construct annotation name to param name map and
	// param name to param type map
	setupData := setup()
	paramNameMap := setupData.headerParamNameMap
	paramTypeMap := setupData.headerParamTypeMap

	// Construct http.Request with http headers
	r, _ := http.NewRequest("GET", "/api/test", nil)
	r.Header.Add("X-NotAnnotated", "abc")
	r.Header.Add("X-Header3", "123")

	// Parse the http header param based on http.Request
	paramValueMap, err := rest.ParseHeaderParams(r, paramNameMap, paramTypeMap)
	assert.Nil(t, err)
	assert.Equal(t, data.NewIntegerValue(123), paramValueMap["param13"])
	_, found := paramValueMap["NotAnnotated"]
	assert.Equal(t, false, found)
}

func TestParseHeaderParams_MissingTypeMap(t *testing.T) {
	// Construct annotation name to param name map and
	// param name to param type map
	setupData := setup()
	paramNameMap := setupData.headerParamNameMap
	paramTypeMap := make(map[string]bindings.BindingType)

	// Construct http.Request with http headers
	r, _ := http.NewRequest("GET", "/api/test", nil)
	r.Header.Add("X-Header3", "123")

	// Parse the http header param based on http.Request
	_, err := rest.ParseHeaderParams(r, paramNameMap, paramTypeMap)
	assert.Equal(t,
		"'param13' is not found in parameter name to type map",
		err[0].Error())

	paramTypeMap = map[string]bindings.BindingType{
		"param12": bindings.NewListType(nil, nil),
		"param15": bindings.NewOptionalType(nil),
	}

	// Construct http.Request with http headers
	r, _ = http.NewRequest("GET", "/api/test", nil)
	r.Header.Add("X-Header2", "123")

	// Parse the http header param based on http.Request
	_, err = rest.ParseHeaderParams(r, paramNameMap, paramTypeMap)
	assert.Equal(t,
		"Request parameter data type '<nil>' not supported",
		err[0].Error())
}

func TestParseQueryParams(t *testing.T) {
	// Construct annotation name to param name map and
	// param name to param type map
	setupData := setup()
	paramNameMap := setupData.queryParamNameMap
	paramTypeMap := setupData.queryParamTypeMap

	// Construct query string, it includes all valid type and list
	query := url.Values{}
	query.Set("query1", "qv1")
	query.Set("query2", "true")
	query.Set("query3", "456")
	query.Set("query4", "123.456")
	query.Add("query5", "1001")
	query.Add("query5", "2002")

	// Construct http.Request
	r, _ := http.NewRequest("GET", "/api/test?"+query.Encode(), nil)

	// Parse the query param based on http.Request
	paramValueMap, err := rest.ParseQueryParams(r, paramNameMap, paramTypeMap)
	assert.Nil(t, err)
	assert.Equal(t, data.NewStringValue("qv1"), paramValueMap["param6"])
	assert.Equal(t, data.NewBooleanValue(true), paramValueMap["param7"])
	assert.Equal(t, data.NewIntegerValue(456), paramValueMap["param8"])
	assert.Equal(t, data.NewDoubleValue(123.456), paramValueMap["param9"])
	listval := data.NewListValue()
	listval.Add(data.NewIntegerValue(1001))
	listval.Add(data.NewIntegerValue(2002))
	assert.Equal(t, listval, paramValueMap["param10"])
}

func TestParseQueryParams_NotAnnotated(t *testing.T) {
	// Construct annotation name to param name map and
	// param name to param type map
	setupData := setup()
	paramNameMap := setupData.queryParamNameMap
	paramTypeMap := setupData.queryParamTypeMap

	// Construct query string, with not exist query name
	query := url.Values{}
	query.Set("NotAnnotated", "abc")
	query.Set("query2", "false")

	// Construct http.Request
	r, _ := http.NewRequest("GET", "/api/test?"+query.Encode(), nil)

	// Parse the query param based on http.Request
	paramValueMap, err := rest.ParseQueryParams(r, paramNameMap, paramTypeMap)
	assert.Nil(t, err)
	assert.Equal(t, data.NewBooleanValue(false), paramValueMap["param7"])
	_, found := paramValueMap["NotAnnotated"]
	assert.Equal(t, false, found)
}

func TestParseQueryParams_InvalidType(t *testing.T) {
	// Construct annotation name to param name map and
	// param name to param type map
	setupData := setup()
	paramNameMap := setupData.queryParamNameMap
	paramTypeMap := setupData.queryParamTypeMap

	// Construct query string, it includes all valid type and list
	query := url.Values{}
	query.Set("query2", "123")
	// Construct http.Request
	r, _ := http.NewRequest("GET", "/api/test?"+query.Encode(), nil)

	// Parse the query param based on http.Request
	_, err := rest.ParseQueryParams(r, paramNameMap, paramTypeMap)
	assert.Equal(t,
		"Invalid value for request parameter 'query2'. Expected a value of type 'Bool', but parsing failed because of 'strconv.ParseBool: parsing \"123\": invalid syntax'",
		err[0].Error())

	// Construct query string, it includes all valid type and list
	query = url.Values{}
	query.Set("query3", "notinteger")
	// Construct http.Request
	r, _ = http.NewRequest("GET", "/api/test?"+query.Encode(), nil)

	// Parse the query param based on http.Request
	_, err = rest.ParseQueryParams(r, paramNameMap, paramTypeMap)
	assert.Equal(t,
		"Invalid value for request parameter 'query3'. Expected a value of type 'Int', but parsing failed because of 'strconv.ParseInt: parsing \"notinteger\": invalid syntax'",
		err[0].Error())

	// Construct query string, it includes all valid type and list
	query = url.Values{}
	query.Set("query4", "notdouble")
	// Construct http.Request
	r, _ = http.NewRequest("GET", "/api/test?"+query.Encode(), nil)

	// Parse the query param based on http.Request
	_, err = rest.ParseQueryParams(r, paramNameMap, paramTypeMap)
	assert.Equal(t,
		"Invalid value for request parameter 'query4'. Expected a value of type 'Float', but parsing failed because of 'strconv.ParseFloat: parsing \"notdouble\": invalid syntax'",
		err[0].Error())

	// Construct query string, it includes all valid type and list
	query = url.Values{}
	query.Add("query5", "notnumber")
	query.Add("query5", "123")
	// Construct http.Request
	r, _ = http.NewRequest("GET", "/api/test?"+query.Encode(), nil)

	// Parse the query param based on http.Request
	_, err = rest.ParseQueryParams(r, paramNameMap, paramTypeMap)
	assert.Equal(t,
		"Invalid value for request parameter 'query5'. Expected a value of type 'Int', but parsing failed because of 'strconv.ParseInt: parsing \"notnumber\": invalid syntax'",
		err[0].Error())
}

func TestParseQueryParams_List(t *testing.T) {
	// Construct annotation name to param name map and
	// param name to param type map
	setupData := setup()
	paramNameMap := setupData.queryParamNameMap
	paramTypeMap := setupData.queryParamTypeMap

	// Construct query string, it includes none list type with multiple values
	// and list type with single value
	query := url.Values{}
	// Note, query Set the order in the actual query param is reversed
	query.Set("query2", "false")
	query.Set("query2", "true")
	query.Add("query5", "123")
	query.Set("query11", "listlist")
	query.Set("query12", "listset")

	// Construct http.Request
	r, _ := http.NewRequest("GET", "/api/test?"+query.Encode(), nil)

	// Parse the query param based on http.Request
	paramValueMap, err := rest.ParseQueryParams(r, paramNameMap, paramTypeMap)
	assert.Nil(t, err)
	assert.Equal(t, data.NewBooleanValue(true), paramValueMap["param7"])
	listval := data.NewListValue()
	listval.Add(data.NewIntegerValue(123))
	assert.Equal(t, listval, paramValueMap["param10"])
	listval = data.NewListValue()
	listval.Add(data.NewStringValue("listlist"))
	listlistval := data.NewListValue()
	listlistval.Add(listval)
	assert.Equal(t, listlistval, paramValueMap["param21"])
	listlistval = data.NewListValue()
	setval := data.NewListValue()
	setval.Add(data.NewStringValue("listset"))
	listlistval.Add(setval)
}

func TestParseQueryParams_Optional(t *testing.T) {
	// Construct annotation name to param name map and
	// param name to param type map
	setupData := setup()
	paramNameMap := setupData.queryParamNameMap
	paramTypeMap := setupData.queryParamTypeMap

	// Construct query string, it includes none list type with multiple values
	// and list type with single value
	query := url.Values{}
	query.Add("query6", "optionallist1")
	query.Add("query6", "optionallist2")

	// Construct http.Request
	r, _ := http.NewRequest("GET", "/api/test?"+query.Encode(), nil)

	// Parse the query param based on http.Request
	paramValueMap, err := rest.ParseQueryParams(r, paramNameMap, paramTypeMap)
	assert.Nil(t, err)
	listval := data.NewListValue()
	listval.Add(data.NewStringValue("optionallist1"))
	listval.Add(data.NewStringValue("optionallist2"))
	optval := data.NewOptionalValue(listval)
	assert.Equal(t, optval, paramValueMap["param16"])
}

func TestParseQueryParams_MiscTypes(t *testing.T) {
	// Construct annotation name to param name map and
	// param name to param type map
	setupData := setup()
	paramNameMap := setupData.queryParamNameMap
	paramTypeMap := setupData.queryParamTypeMap

	// Construct query string,
	query := url.Values{}
	query.Set("query7", "vmID")
	query.Set("query8", "VGhpcyBpcyB0ZXN0")
	query.Add("query9", "setval1")
	query.Add("query9", "setval2")
	query.Set("query10", "secret")

	// Construct http.Request
	r, _ := http.NewRequest("GET", "/api/test?"+query.Encode(), nil)

	// Parse the query param based on http.Request
	paramValueMap, err := rest.ParseQueryParams(r, paramNameMap, paramTypeMap)
	assert.Nil(t, err)
	idval := data.NewStringValue("vmID")
	assert.Equal(t, idval, paramValueMap["param17"])
	blobval := data.NewBlobValue([]byte("VGhpcyBpcyB0ZXN0"))
	assert.Equal(t, blobval, paramValueMap["param18"])
	setval := data.NewListValue()
	setval.Add(data.NewStringValue("setval1"))
	setval.Add(data.NewStringValue("setval2"))
	assert.Equal(t, setval, paramValueMap["param19"])
	secretval := data.NewSecretValue("secret")
	assert.Equal(t, secretval, paramValueMap["param20"])
}

func TestParseQueryParams_MissingTypeMap(t *testing.T) {
	// Construct annotation name to param name map and
	// param name to param type map
	setupData := setup()
	paramNameMap := setupData.queryParamNameMap
	paramTypeMap := make(map[string]bindings.BindingType)

	// Construct query string, it includes all valid type and list
	query := url.Values{}
	query.Set("query4", "true")

	// Construct http.Request
	r, _ := http.NewRequest("GET", "/api/test?"+query.Encode(), nil)

	// Parse the query param based on http.Request
	_, err := rest.ParseQueryParams(r, paramNameMap, paramTypeMap)
	assert.Equal(t,
		"'param9' is not found in parameter name to type map",
		err[0].Error())

	paramTypeMap = map[string]bindings.BindingType{
		"param10": bindings.NewListType(nil, nil),
	}

	// Construct query string, it includes all valid type and list
	query = url.Values{}
	query.Add("query5", "abc")

	// Construct http.Request
	r, _ = http.NewRequest("GET", "/api/test?"+query.Encode(), nil)

	// Parse the query param based on http.Request
	_, err = rest.ParseQueryParams(r, paramNameMap, paramTypeMap)
	assert.Equal(t,
		"Request parameter data type '<nil>' not supported",
		err[0].Error())

	paramTypeMap = map[string]bindings.BindingType{
		"param16": bindings.NewOptionalType(nil),
	}

	// Construct query string, it includes all valid type and list
	query = url.Values{}
	query.Add("query6", "abc")

	// Construct http.Request
	r, _ = http.NewRequest("GET", "/api/test?"+query.Encode(), nil)

	// Parse the query param based on http.Request
	_, err = rest.ParseQueryParams(r, paramNameMap, paramTypeMap)
	assert.Equal(t,
		"Request parameter data type '<nil>' not supported",
		err[0].Error())
}

func TestParseBodyParams_PrimitiveType(t *testing.T) {
	var dataValueToJsonEncoder = cleanjson.NewDataValueToJsonEncoder()

	bodyBytes, _ := dataValueToJsonEncoder.Encode(data.NewStringValue("string body"))
	reader := strings.NewReader(bodyBytes)
	// Construct http.Request
	r, _ := http.NewRequest("POST", "/api/test", reader)
	bodyDV, err := rest.ParseBodyParams(r)
	assert.Nil(t, err)
	assert.Equal(t, "string body", bodyDV.(*data.StringValue).Value())

	bodyBytes, _ = dataValueToJsonEncoder.Encode(data.NewBlobValue([]byte("VGhpcyBpcyB0ZXN0")))
	reader = strings.NewReader(bodyBytes)
	// Construct http.Request
	r, _ = http.NewRequest("POST", "/api/test", reader)
	bodyDV, err = rest.ParseBodyParams(r)
	assert.Nil(t, err)
	assert.Equal(t, "VkdocGN5QnBjeUIwWlhOMA==", bodyDV.(*data.StringValue).Value())

	bodyBytes, _ = dataValueToJsonEncoder.Encode(data.NewSecretValue("secret"))
	reader = strings.NewReader(bodyBytes)
	// Construct http.Request
	r, _ = http.NewRequest("POST", "/api/test", reader)
	bodyDV, err = rest.ParseBodyParams(r)
	assert.Nil(t, err)
	assert.Equal(t, "secret", bodyDV.(*data.StringValue).Value())

	bodyBytes, _ = dataValueToJsonEncoder.Encode(data.NewIntegerValue(123))
	reader = strings.NewReader(bodyBytes)
	// Construct http.Request
	r, _ = http.NewRequest("POST", "/api/test", reader)
	bodyDV, err = rest.ParseBodyParams(r)
	assert.Nil(t, err)
	assert.Equal(t, int64(123), bodyDV.(*data.IntegerValue).Value())

	bodyBytes, _ = dataValueToJsonEncoder.Encode(data.NewDoubleValue(1.23))
	reader = strings.NewReader(bodyBytes)
	// Construct http.Request
	r, _ = http.NewRequest("POST", "/api/test", reader)
	bodyDV, err = rest.ParseBodyParams(r)
	assert.Nil(t, err)
	assert.Equal(t, float64(1.23), bodyDV.(*data.DoubleValue).Value())

	bodyBytes, _ = dataValueToJsonEncoder.Encode(data.NewBooleanValue(true))
	reader = strings.NewReader(bodyBytes)
	// Construct http.Request
	r, _ = http.NewRequest("POST", "/api/test", reader)
	bodyDV, err = rest.ParseBodyParams(r)
	assert.Nil(t, err)
	assert.Equal(t, true, bodyDV.(*data.BooleanValue).Value())
}

func TestParseBodyParams_OptionalType(t *testing.T) {
	var dataValueToJsonEncoder = cleanjson.NewDataValueToJsonEncoder()

	optVal := data.NewOptionalValue(data.NewIntegerValue(123))
	bodyBytes, _ := dataValueToJsonEncoder.Encode(optVal)
	reader := strings.NewReader(bodyBytes)
	// Construct http.Request
	r, _ := http.NewRequest("POST", "/api/test", reader)
	bodyDV, err := rest.ParseBodyParams(r)
	assert.Nil(t, err)
	assert.Equal(t, int64(123), bodyDV.(*data.IntegerValue).Value())

	optVal = data.NewOptionalValue(nil)
	bodyBytes, _ = dataValueToJsonEncoder.Encode(optVal)
	reader = strings.NewReader(bodyBytes)
	// Construct http.Request
	r, _ = http.NewRequest("POST", "/api/test", reader)
	bodyDV, err = rest.ParseBodyParams(r)
	assert.Nil(t, err)
	assert.Equal(t, nil, bodyDV.(*data.OptionalValue).Value())
}

func TestParseBodyParams_ListAndSetType(t *testing.T) {
	var dataValueToJsonEncoder = cleanjson.NewDataValueToJsonEncoder()

	listval := data.NewListValue()
	listval.Add(data.NewStringValue("v1"))
	listval.Add(data.NewStringValue("v2"))

	bodyBytes, _ := dataValueToJsonEncoder.Encode(listval)
	reader := strings.NewReader(bodyBytes)
	// Construct http.Request
	r, _ := http.NewRequest("POST", "/api/test", reader)
	bodyDV, err := rest.ParseBodyParams(r)
	assert.Nil(t, err)
	assert.Equal(t, "v1",
		bodyDV.(*data.ListValue).Get(0).(*data.StringValue).Value())
	assert.Equal(t, "v2",
		bodyDV.(*data.ListValue).Get(1).(*data.StringValue).Value())

	setval := data.NewListValue()
	setval.Add(data.NewDoubleValue(1.23))
	setval.Add(data.NewDoubleValue(4.56))

	bodyBytes, _ = dataValueToJsonEncoder.Encode(setval)
	reader = strings.NewReader(bodyBytes)
	// Construct http.Request
	r, _ = http.NewRequest("POST", "/api/test", reader)
	bodyDV, err = rest.ParseBodyParams(r)
	assert.Nil(t, err)
	assert.Equal(t, 1.23,
		bodyDV.(*data.ListValue).Get(0).(*data.DoubleValue).Value())
	assert.Equal(t, 4.56,
		bodyDV.(*data.ListValue).Get(1).(*data.DoubleValue).Value())
}

func TestParseBodyParams_MapType(t *testing.T) {
	var dataValueToJsonEncoder = cleanjson.NewDataValueToJsonEncoder()

	type TestType struct {
		field1 int
		field2 int
	}

	var fields = make(map[string]data.DataValue)

	fields["field1"] = data.NewIntegerValue(123)
	fields["field2"] = data.NewIntegerValue(456)
	structVal := data.NewStructValue("test", fields)

	bodyBytes, _ := dataValueToJsonEncoder.Encode(structVal)
	reader := strings.NewReader(bodyBytes)
	// Construct http.Request
	r, _ := http.NewRequest("POST", "/api/test", reader)
	bodyDV, err := rest.ParseBodyParams(r)
	assert.Nil(t, err)
	fieldVal, _ := bodyDV.(*data.StructValue).Field("field1")
	assert.Equal(t, int64(123), fieldVal.(*data.IntegerValue).Value())
	fieldVal, _ = bodyDV.(*data.StructValue).Field("field2")
	assert.Equal(t, int64(456), fieldVal.(*data.IntegerValue).Value())
}

func TestParseBodyParams_FailureCases(t *testing.T) {
	//var dataValueToJsonEncoder = cleanjson.NewDataValueToJsonEncoder()

	// Invalid value
	reader := strings.NewReader("number")
	// Construct http.Request
	r, _ := http.NewRequest("POST", "/api/test", reader)
	_, err := rest.ParseBodyParams(r)
	assert.Equal(t, "Error when parsing request body, 'invalid character 'm' in literal null (expecting 'l')'",
		err[0].Error())

	// Request body reader has no data
	r, _ = http.NewRequest("POST", "/api/test", nil)
	_, err = rest.ParseBodyParams(r)
	assert.Equal(t, "Error when parsing request body, 'http request body is nil'",
		err[0].Error())
	r.Body = ioutil.NopCloser(bytes.NewReader([]byte("")))
	_, err = rest.ParseBodyParams(r)
	assert.Equal(t, "Error when parsing request body, 'EOF'",
		err[0].Error())

	// Not supported binding type
	_, err = rest.ParseBodyParams(r)
	assert.Equal(t, "Error when parsing request body, 'EOF'",
		err[0].Error())
}