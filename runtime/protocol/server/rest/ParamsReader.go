package rest

/* **********************************************************
 * Copyright 2019 VMware, Inc.  All rights reserved.
 *      -- VMware Confidential
 * **********************************************************/

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data/serializers/cleanjson"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/l10n"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

// Parse the URL path variable. It takes annotation name to param name map,
// param name to type map, and http request. It generates a map from
// param name its value.
func ParsePathParams(
	request *http.Request,
	paramNameMap map[string]string,
	paramTypeMap map[string]bindings.BindingType) (map[string]data.DataValue, []error) {

	// Get map of path param name to its value
	pathStringMap := mux.Vars(request)
	pathParamMap := map[string][]string{}

	// Iterate pathStringMap with type map[string]string to
	// construct pathParamMap with type map[string][]string in order
	// to use the same generateParamDataValueMap() method
	for k, v := range pathStringMap {
		pathParamMap[k] = []string{v}
	}

	return generateParamDataValueMap(pathParamMap,
		paramNameMap,
		paramTypeMap)
}

// Parse the http header. It takes annotation name to param name map,
// param name to type map, and http request. It generates a map from
// param name to its value.
func ParseHeaderParams(
	r *http.Request,
	paramNameMap map[string]string,
	paramTypeMap map[string]bindings.BindingType) (map[string]data.DataValue, []error) {

	// Get map of header param name to its value (slice)
	headerParamMap := r.Header
	return generateParamDataValueMap(headerParamMap,
		paramNameMap,
		paramTypeMap)
}

// Parse the http query. It takes annotation name to param name map,
// param name to type map, and http request. It generates a map from
// param name to its value.
func ParseQueryParams(
	r *http.Request,
	paramNameMap map[string]string,
	paramTypeMap map[string]bindings.BindingType) (map[string]data.DataValue, []error) {

	// Get map of query param name to its value (slice)
	queryParamMap := r.URL.Query()

	return generateParamDataValueMap(queryParamMap,
		paramNameMap,
		paramTypeMap)
}

// Parse the http body. There is at most one @Body annotation per operation.
// It takes param type, and http request. It generates param value.
func ParseBodyParams(
	r *http.Request) (data.DataValue, []error) {
	var err error

	bodyReader := r.Body
	if bodyReader == nil {
		return nil,
			[]error{l10n.NewRuntimeError("vapi.protocol.server.rest.param.body_parse_error",
				map[string]string{"errMsg": "http request body is nil"})}
	}

	bodyByte, err := ioutil.ReadAll(bodyReader)
	if err != nil {
		return nil,
			[]error{l10n.NewRuntimeError("vapi.protocol.server.rest.param.body_parse_error",
				map[string]string{"errMsg": err.Error()})}
	}

	return convertRequestBodyToDataValue(bodyByte)
}

// Return parameter name and its type by annotation name,
// Parameter annotationName should always exist in paramNameMap
func findParamNameType(annotationName string,
	paramNameMap map[string]string,
	paramTypeMap map[string]bindings.BindingType) (string, bindings.BindingType, []error) {

	paramName, _ := paramNameMap[annotationName]

	paramType, exist := paramTypeMap[paramName]
	if !exist {
		pName := fmt.Sprintf("%s", paramName)
		return "", nil, []error{l10n.NewRuntimeError("vapi.protocol.server.rest.param.name_type_map_not_found",
			map[string]string{"paramName": pName})}
	}
	return paramName, paramType, nil
}

// Parse a parameter value in string presentation to DataValue
// based on the DataType.
// For multiple values with non list type, we pick first value.
// This is an unspecified scenario in spec
func convertRequestParamToDataValue(
	name string,
	pValue []string,
	pType bindings.BindingType) (data.DataValue, []error) {

	switch reflect.TypeOf(pType) {
	case bindings.StringBindingType, bindings.IdBindingType:
		return data.NewStringValue(pValue[0]), nil
	case bindings.SecretBindingType:
		return data.NewSecretValue(pValue[0]), nil
	case bindings.BlobBindingType:
		return data.NewBlobValue([]byte(pValue[0])), nil
	case bindings.BooleanBindingType:
		b, err := strconv.ParseBool(pValue[0])
		if err != nil {
			return nil, []error{l10n.NewRuntimeError("vapi.protocol.server.rest.param.invalid_value",
				map[string]string{"paramName": name, "expectedType": "Bool", "errMsg": err.Error()})}
		}
		return data.NewBooleanValue(b), nil
	case bindings.DoubleBindingType:
		i, err := strconv.ParseFloat(pValue[0], 64)
		if err != nil {
			return nil, []error{l10n.NewRuntimeError("vapi.protocol.server.rest.param.invalid_value",
				map[string]string{"paramName": name, "expectedType": "Float", "errMsg": err.Error()})}
		}
		return data.NewDoubleValue(i), nil
	case bindings.IntegerBindingType:
		i, err := strconv.ParseInt(pValue[0], 10, 64)
		if err != nil {
			return nil, []error{l10n.NewRuntimeError("vapi.protocol.server.rest.param.invalid_value",
				map[string]string{"paramName": name, "expectedType": "Int", "errMsg": err.Error()})}
		}
		return data.NewIntegerValue(i), nil
	case bindings.ListBindingType:
		eType := pType.(bindings.ListType).ElementType()
		return convertParamElementsToDataValue(name, pValue, eType)
	case bindings.SetBindingType:
		eType := pType.(bindings.SetType).ElementType()
		return convertParamElementsToDataValue(name, pValue, eType)
	case bindings.OptionalBindingType:
		eType := pType.(bindings.OptionalType).ElementType()
		dv, err := convertRequestParamToDataValue(name, pValue, eType)
		if err != nil {
			return nil, err
		}
		result := data.NewOptionalValue(dv)
		return result, nil
	default:
		// Note: Spec state that Map is not supported as query/header parameter
		dataType := fmt.Sprintf("%v", pType)
		return nil, []error{l10n.NewRuntimeError("vapi.protocol.server.rest.param.unsupport_type",
			map[string]string{"dataType": dataType})}
	}
}

// Convert List/Set's element value to a DataValue
func convertParamElementsToDataValue(
	name string,
	pValue []string,
	eType bindings.BindingType) (data.DataValue, []error) {
	result := data.NewListValue()
	for _, value := range pValue {
		dv, err := convertRequestParamToDataValue(name, []string{value}, eType)
		if err != nil {
			return nil, err
		} else {
			result.Add(dv)
		}
	}
	return result, nil
}

// Create param name to its data value mapping
func generateParamDataValueMap(paramStringMap map[string][]string,
	paramNameMap map[string]string,
	paramTypeMap map[string]bindings.BindingType) (map[string]data.DataValue, []error) {

	paramValueMap := map[string]data.DataValue{}

	for annotationName, valuelist := range paramStringMap {
		_, found := paramNameMap[annotationName]
		if !found {
			// The query name is not annotated, skip
			continue
		}
		name, ptype, errIn := findParamNameType(annotationName,
			paramNameMap,
			paramTypeMap)

		if errIn == nil {
			paramValueMap[name], errIn = convertRequestParamToDataValue(annotationName, valuelist, ptype)
		}
		if errIn != nil {
			return nil, errIn
		}
	}
	return paramValueMap, nil
}

// Convert request body byte to value in DataValue type
func convertRequestBodyToDataValue(
	bodyByte []byte) (data.DataValue, []error) {
	var err error

	jdecoder := json.NewDecoder(strings.NewReader(string(bodyByte)))
	jdecoder.UseNumber()
	var val interface{}
	err = jdecoder.Decode(&val)

	if err != nil {
		return nil,
			[]error{l10n.NewRuntimeError("vapi.protocol.server.rest.param.body_parse_error",
				map[string]string{"errMsg": err.Error()})}
	}

	jsonToDataValueDecoder := cleanjson.NewJsonToDataValueDecoder()
	bodyDV, err := jsonToDataValueDecoder.Decode(val)
	if err != nil {
		// TODO: will change to err after JsonToDataValueDecoder using l10n
		return nil, []error{err}
	}

	return bodyDV, nil
}