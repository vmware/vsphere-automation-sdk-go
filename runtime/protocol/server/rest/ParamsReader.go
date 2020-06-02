/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

package rest

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/textproto"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/data/serializers/cleanjson"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/l10n"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/lib"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/log"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/protocol"
)

type ParamsType string

const (
	queryParamsType  ParamsType = "query"
	pathParamsType   ParamsType = "path"
	headerParamsType ParamsType = "header"
)

// ParsePathParams Parse the URL path variable. It takes annotation name to param name map,
// param name to type map, and http request. It generates a map from
// param name its value.
func ParsePathParams(
	request *http.Request,
	metadata protocol.OperationRestMetadata,
	inputDataVal *data.StructValue) []error {

	pathDataValMap, err := PathDataValueMap(request, metadata)
	if err != nil {
		return err
	}

	for k, v := range pathDataValMap {
		tokens := strings.Split(k, ".")
		if len(tokens) > 1 {
			ConvertToDataValueIfNested(inputDataVal, tokens, v, metadata.Fields()[tokens[0]])
		} else {
			inputDataVal.SetField(k, v)
		}
	}

	return nil
}

func PathDataValueMap(request *http.Request,
	metadata protocol.OperationRestMetadata) (map[string]data.DataValue, []error) {

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
		metadata.PathParamsNameMap(),
		metadata.ParamsTypeMap(),
		pathParamsType)
}

func canonicalizeFieldNames(headerParamMap map[string]string) map[string]string {
	res := make(map[string]string)
	for k, v := range headerParamMap {
		res[textproto.CanonicalMIMEHeaderKey(k)] = v
	}
	return res
}

// ParseHeaderParams Parse the http header. It takes annotation name to param name map,
// param name to type map, and http request. It generates a map from
// param name to its value.
func ParseHeaderParams(
	r *http.Request,
	metadata protocol.OperationRestMetadata,
	inputDataVal *data.StructValue) []error {

	headerDataValMap, err := HeaderDataValueMap(r, metadata)
	if err != nil {
		return err
	}

	for k, v := range headerDataValMap {
		tokens := strings.Split(k, ".")
		if len(tokens) > 1 {
			ConvertToDataValueIfNested(inputDataVal, tokens, v, metadata.Fields()[tokens[0]])
		} else {
			inputDataVal.SetField(k, v)
		}
	}

	return nil
}

func HeaderDataValueMap(request *http.Request,
	metadata protocol.OperationRestMetadata) (map[string]data.DataValue, []error) {

	// Get map of header param name to its value (slice)
	headerParamMap := request.Header
	return generateParamDataValueMap(headerParamMap,
		canonicalizeFieldNames(reverseMap(metadata.HeaderParamsNameMap())),
		metadata.ParamsTypeMap(),
		headerParamsType)
}

// ParseQueryParams Parse the http query. It takes annotation name to param name map,
// param name to type map, and http request. It generates a map from
// param name to its value.
func ParseQueryParams(
	r *http.Request,
	metadata protocol.OperationRestMetadata,
	inputDataVal *data.StructValue) []error {

	queryDataValMap, err := QueryDataValueMap(r, metadata)
	if err != nil {
		return err
	}

	for k, v := range queryDataValMap {
		tokens := strings.Split(k, ".")
		if len(tokens) > 1 {
			ConvertToDataValueIfNested(inputDataVal, tokens, v, metadata.Fields()[tokens[0]])
		} else {
			inputDataVal.SetField(k, v)
		}
	}
	return nil
}

/** Parses dispatch params string avaiable in generated rest metadata
 * and returns query field Map value accepted in http standards
 * input : dispatchParam:"a=3&b&c=5&a=2"
 * output : { a:[3,2], b:[], c:[5]}
 */
func parseDispatchParams(dispatchParam string) (map[string][]string, []error) {
	res := map[string][]string{}
	if dispatchParam == "" {
		return res, nil
	}
	parr := strings.Split(dispatchParam, "&")

	for _, val := range parr {
		kvset := strings.Split(val, "=")

		if len(kvset) == 0 {
			continue
		}

		var key string
		var value []string

		if len(kvset) == 1 {
			key, value = kvset[0], []string{}
		} else {
			// `=` is a allowed character so kvset[1:] is complete valid value disjoined by `=`
			key, value = kvset[0], []string{strings.Join(kvset[1:], "=")}
		}

		if key == "" {
			return nil, []error{l10n.NewRuntimeError("vapi.protocol.server.rest.param.invalid_value",
				map[string]string{"errMsg": "default query had invalid key value pair i.e " + val})}
		}

		if _, ok := res[key]; ok {
			res[key] = append(res[key], value...)
		} else {
			res[key] = value
		}
	}
	return res, nil
}

// Checks if unsupported properties are passed in form of field key-values, if true return list of all unsupported keys.
func checkExtraRequestKeys(requestKeys, requiredKeys []string) ([]string, bool) {
	res := []string{}
	temp := make(map[string]int)

	// Add all the keys sent from client to a map
	for _, v := range requestKeys {
		temp[v]++
	}
	// Remove all the keys that are either required or supported by the operation
	for _, v := range requiredKeys {
		temp[v]--
		if temp[v] <= 0 {
			delete(temp, v)
		}
	}
	// If map is not empty then we are left with residual keys which are not supported
	if len(temp) == 0 {
		return res, false
	}
	// Record the residual keys and send it back to throw unsupported property exception
	for k := range temp {
		res = append(res, k)
	}

	return res, true
}

/** Query data can come from two places:
 * 1. Queries which were passed to be fed as data using @Query
 * 2. Queries used only for routing i.e dispatch parameters
 * While combination of both can exist, example:
 * https://opengrok.eng.vmware.com/source/xref/vapi-main.perforce.1666/vapi-core/idl-toolkit/vmodl-models/src/dist/models/test/vmodl/rest-native/VerbParams.vmodl#53
 */
func QueryDataValueMap(request *http.Request,
	metadata protocol.OperationRestMetadata) (map[string]data.DataValue, []error) {

	// Get map of query param name to its value (slice)
	queryParamMap := request.URL.Query()
	requestKeys := []string{}
	for k := range queryParamMap {
		requestKeys = append(requestKeys, k)
	}

	dispatchparam, err := parseDispatchParams(metadata.DispatchParam())
	if err != nil {
		return nil, err
	}

	// Collect all keys that you have for query field, for validation.
	requiredKeys := []string{}
	for k := range dispatchparam {
		requiredKeys = append(requiredKeys, k)
	}

	reverseQueryParamsName := reverseMap(metadata.QueryParamsNameMap())
	for k := range reverseQueryParamsName {
		requiredKeys = append(requiredKeys, k)
	}

	// Validation: Add check, To see if extra arguments/properties were passed
	if val, ok := checkExtraRequestKeys(requestKeys, requiredKeys); ok {
		return nil, []error{l10n.NewRuntimeError("com.vmware.vapi.rest.unsupported_property",
			map[string]string{"arg": strings.Join(val, ",")})}
	}

	return generateParamDataValueMap(queryParamMap,
		reverseQueryParamsName,
		metadata.ParamsTypeMap(),
		queryParamsType)
}

// Return parameter name and its type by annotation name,
// Parameter's annotationName should always exist in paramNameMap
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

func reverseMap(mapvalue map[string]string) map[string]string {
	res := make(map[string]string)
	for k, v := range mapvalue {
		res[v] = k
	}
	return res
}

// Parse a parameter value in string presentation to DataValue
// based on the DataType.
// For multiple values with non list type, we pick first value.
// This is an unspecified scenario in spec
func convertRequestParamToDataValue(
	name string,
	pValue []string,
	pType bindings.BindingType,
	paramsType ParamsType,
	insideOptional bool) (data.DataValue, []error) {

	switch reflect.TypeOf(pType) {
	case bindings.StringBindingType:
		// If multiple headers are provided for String Binding then we need to merge them
		if paramsType == headerParamsType {
			return data.NewStringValue(strings.Join(pValue, ",")), nil
		}
		// In query type:
		// 1. If multiple values are provided for query then we fail.
		// 2. If a string value that is passed is empty and it wasn't optional then we propagate empty string forwards. https://opengrok.eng.vmware.com/source/xref/vapi-main.perforce.1666/vapi-core/integration-tests/vapi-verb-integration/src/test/java/com/vmware/vapi/rest/integration/verb/JsonQueryMappingTest.java#176
		// 3. If a string value that is passed is empty and it is optional then we propagate it as null. https://opengrok.eng.vmware.com/source/xref/vapi-main.perforce.1666/vapi-core/integration-tests/vapi-verb-integration/src/test/java/com/vmware/vapi/rest/integration/verb/JsonQueryMappingTest.java#192
		if len(pValue) > 1 {
			return nil, []error{l10n.NewRuntimeError("vapi.protocol.server.rest.param.invalid_value",
				map[string]string{"errMsg": "Expected type STRING but got 'LIST'"})}
		}
		if paramsType == queryParamsType && pValue[0] == "" && insideOptional {
			return nil, nil
		}
		return data.NewStringValue(pValue[0]), nil
	case bindings.IdBindingType:
		if len(pValue) > 1 {
			return nil, []error{l10n.NewRuntimeError("vapi.protocol.server.rest.param.invalid_value",
				map[string]string{"errMsg": "Expected type ID but got 'LIST'"})}
		}
		return data.NewStringValue(pValue[0]), nil
	case bindings.SecretBindingType:
		return data.NewSecretValue(pValue[0]), nil
	case bindings.BlobBindingType:
		if len(pValue) == 1 && string(pValue[0]) == "" {
			return nil, []error{l10n.NewRuntimeError("vapi.protocol.server.rest.param.invalid_value",
				map[string]string{"errMsg": "Empty string passed for Blob type"})}
		}
		var temp []byte
		var err error
		temp, err = base64.StdEncoding.DecodeString(pValue[0])
		if err != nil && paramsType == headerParamsType {
			log.Errorf("Error ignored; Header param return valid prefix data: " + err.Error())
		} else if err != nil && insideOptional {
			log.Errorf("Error ignored; BLOB inside Optional: " + err.Error())
		} else if err != nil {
			temp, err = base64.URLEncoding.DecodeString(pValue[0])
			if err != nil {
				return nil, []error{l10n.NewRuntimeError("vapi.protocol.server.rest.param.invalid_value",
					map[string]string{"errMsg": err.Error()})}
			}
		}
		return data.NewBlobValue(temp), nil
	case bindings.BooleanBindingType:
		b, err := strconv.ParseBool(strings.ToLower(pValue[0]))
		if err != nil {
			return nil, []error{l10n.NewRuntimeError("vapi.protocol.server.rest.param.invalid_type",
				map[string]string{"paramName": name, "expectedType": "Bool", "errMsg": err.Error()})}
		}
		return data.NewBooleanValue(b), nil
	case bindings.DoubleBindingType:
		i, err := strconv.ParseFloat(pValue[0], 64)
		if err != nil {
			return nil, []error{l10n.NewRuntimeError("vapi.protocol.server.rest.param.invalid_type",
				map[string]string{"paramName": name, "expectedType": "Float", "errMsg": err.Error()})}
		}
		return data.NewDoubleValue(i), nil
	case bindings.IntegerBindingType:
		if len(pValue) > 1 {
			return nil, []error{l10n.NewRuntimeError("vapi.protocol.server.rest.param.invalid_value",
				map[string]string{"errMsg": "Expected type Integer but got 'LIST'"})}
		}
		i, err := strconv.ParseInt(pValue[0], 10, 64)
		if err != nil {
			return nil, []error{l10n.NewRuntimeError("vapi.protocol.server.rest.param.invalid_type",
				map[string]string{"paramName": name, "expectedType": "Int", "errMsg": err.Error()})}
		}
		return data.NewIntegerValue(i), nil
	case bindings.ListBindingType:
		eType := pType.(bindings.ListType).ElementType()
		return convertListParamElementsToDataValue(name, pValue, eType, paramsType, insideOptional)
	case bindings.SetBindingType:
		eType := pType.(bindings.SetType).ElementType()
		return convertListParamElementsToDataValue(name, pValue, eType, paramsType, insideOptional)
	case bindings.OptionalBindingType:
		eType := pType.(bindings.OptionalType).ElementType()
		dv, err := convertRequestParamToDataValue(name, pValue, eType, paramsType, true)
		if err != nil {
			return nil, err
		}
		return data.NewOptionalValue(dv), nil
	case bindings.EnumBindingType:
		if paramsType == headerParamsType && pValue[0] == "" && insideOptional {
			return nil, nil
		}
		return data.NewStringValue(pValue[0]), nil
	case bindings.DateTimeBindingType:
		return data.NewStringValue(pValue[0]), nil
	default:
		// Note: Spec state that Map is not supported as query/header parameter
		dataType := fmt.Sprintf("%v", pType)
		return nil, []error{l10n.NewRuntimeError("vapi.protocol.server.rest.param.unsupport_type",
			map[string]string{"dataType": dataType})}
	}
}

// Convert List/Set's element value to a DataValue
func convertListParamElementsToDataValue(
	name string,
	pValue []string,
	eType bindings.BindingType,
	paramsType ParamsType,
	insideOptional bool) (data.DataValue, []error) {
	result := data.NewListValue()
	for _, value := range pValue {
		dv, err := convertRequestParamToDataValue(name, []string{value}, eType, paramsType, insideOptional)
		if err != nil {
			return nil, err
		} else if dv != nil {
			result.Add(dv)
		}
	}
	if len(result.List()) == 0 {
		return nil, nil
	}
	return result, nil
}

// Create param name to its data value mapping
func generateParamDataValueMap(paramStringMap map[string][]string,
	paramNameMap map[string]string,
	paramTypeMap map[string]bindings.BindingType, paramsType ParamsType) (map[string]data.DataValue, []error) {

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
			paramValueMap[name], errIn = convertRequestParamToDataValue(annotationName, valuelist, ptype, paramsType, false)
		}
		if errIn != nil {
			return nil, errIn
		}
	}
	return paramValueMap, nil
}

/*
	Body Field param processing
*/

func ParseBodyFieldParams(
	r *http.Request,
	metadata protocol.OperationRestMetadata,
	inputDataVal *data.StructValue) []error {
	var err error

	bodyFieldsMap := metadata.BodyFieldsMap()
	if len(bodyFieldsMap) == 0 {
		return nil
	}

	bodyReader := r.Body
	if bodyReader == nil {
		return []error{l10n.NewRuntimeError("vapi.protocol.server.rest.param.body_parse_error",
			map[string]string{"errMsg": "http request body is not present"})}
	}

	bodyByte, err := ioutil.ReadAll(bodyReader)
	if err != nil {
		return []error{l10n.NewRuntimeError("vapi.protocol.server.rest.param.body_parse_error",
			map[string]string{"errMsg": err.Error()})}
	}

	if len(bodyByte) == 0 {
		return nil
	}

	val, errs := getContentAsPerContentType(bodyByte, r)
	if errs != nil {
		return errs
	}

	jsonToDataValueDecoder := cleanjson.NewJsonToDataValueDecoder()
	fieldsBindingMap := metadata.Fields()

	if val == nil {
		for _, bindingType := range fieldsBindingMap {
			if reflect.TypeOf(bindingType) != bindings.OptionalBindingType {
				return []error{l10n.NewRuntimeErrorNoParam("vapi.protocol.server.rest.param.bodyfield_parse_invalid")}
			}
		}
		return nil
	}

	valMap, ok := val.(map[string]interface{})
	if !ok {
		return []error{l10n.NewRuntimeErrorNoParam("vapi.protocol.server.rest.param.bodyfield_parse_invalid")}
	}

	for fieldName, fieldValue := range valMap {
		if paramName, ok := bodyFieldsMap[fieldName]; ok {
			tempData, err := requestBodyToDataValue(fieldValue, fieldsBindingMap[paramName], jsonToDataValueDecoder)
			if err != nil {
				return err
			}
			inputDataVal.SetField(paramName, tempData)
		} else if fieldValue != nil {
			return []error{l10n.NewRuntimeError("com.vmware.vapi.rest.unsupported_property",
				map[string]string{"arg": fieldName})}
		}
	}

	return nil
}

/*
	Body param processing
*/

// ParseBodyParams Parse the http body. There is at most one @Body annotation per operation.
// It takes param type, and http request. It generates param value.
func ParseBodyParams(
	r *http.Request,
	bodyParamName string,
	metadata protocol.OperationRestMetadata,
	inputDataVal *data.StructValue) []error {

	bodyReader := r.Body
	if bodyReader == nil {
		return []error{l10n.NewRuntimeError("vapi.protocol.server.rest.param.body_parse_error",
			map[string]string{"errMsg": "http request body is nil"})}
	}

	var bodyByte []byte
	bodyByte, err := ioutil.ReadAll(bodyReader)
	if err != nil {
		return []error{l10n.NewRuntimeError("vapi.protocol.server.rest.param.body_parse_error",
			map[string]string{"errMsg": err.Error()})}
	}

	tokens := strings.Split(bodyParamName, ".")
	bodyBindingType := metadata.ParamsTypeMap()[bodyParamName]
	if len(tokens) > 1 {
		var bodyDataVal data.DataValue
		if len(bodyByte) == 0 && resolveToCheckIfOptional(metadata.Fields()[tokens[0]], tokens[1:]) {
			return nil
		}
		var er []error
		bodyDataVal, er = convertRequestBodyToDataValue(bodyByte, bodyParamName, bodyBindingType, r)
		if er != nil {
			return er
		}
		ConvertToDataValueIfNested(inputDataVal, tokens, bodyDataVal, metadata.Fields()[tokens[0]])
	} else {
		bodyDataVal, er := convertRequestBodyToDataValue(bodyByte, bodyParamName, bodyBindingType, r)
		if er != nil {
			return er
		}
		inputDataVal.SetField(bodyParamName, bodyDataVal)
	}
	return nil
}

func getContentAsPerContentType(
	body []byte,
	r *http.Request) (interface{}, []error) {

	contentType := r.Header.Get(lib.HTTP_CONTENT_TYPE_HEADER)
	if contentType == lib.JSON_CONTENT_TYPE {
		return getContentAppJSON(body)
	} else if contentType == lib.TEXT_PLAIN_CONTENT_TYPE {
		return getContentTextPlain(body)
	} else if contentType == lib.FORM_URL_CONTENT_TYPE {
		return getContentAppFormURLEncoded(body)
	}
	return nil, []error{l10n.NewRuntimeErrorNoParam("com.vmware.vapi.rest.unsupported_media_type")}
}

func getContentAppJSON(body []byte) (interface{}, []error) {

	// Validate: Check if its valid Json body
	// Note: json Unmarshal is not sufficient to cover all inputs cases hence only use it to validate
	// the body, json.NewDecoder can even decode non json structures hence can't be used for validation.
	var dummyHolder interface{}
	err := json.Unmarshal(body, &dummyHolder)
	if err != nil {
		return nil, []error{l10n.NewRuntimeErrorNoParam("com.vmware.vapi.rest.unsupported_media_type")}
	}

	return jdecoderBody(body)
}

func getContentTextPlain(body []byte) (interface{}, []error) {

	// TODO: Find more elegant way to parse datetimes
	// json Decode doesn't work with datetime, it parses incomplete information example: 1878-03-04T12:06:07Z is parsed to 1878
	if _, err := time.Parse(bindings.RFC3339Nano_DATETIME_LAYOUT, string(body)); err == nil {
		return string(body), nil
	}
	return jdecoderBody(body)
}

func jdecoderBody(body []byte) (interface{}, []error) {
	var val interface{}
	jdecoder := json.NewDecoder(strings.NewReader(string(body)))
	jdecoder.UseNumber()
	err := jdecoder.Decode(&val)
	if err != nil {
		return nil, []error{l10n.NewRuntimeError("vapi.protocol.server.rest.param.body_parse_error",
			map[string]string{"errMsg": err.Error()})}
	}
	return val, nil
}

func getContentAppFormURLEncoded(body []byte) (interface{}, []error) {

	formURLBody, err := url.ParseQuery(string(body))
	if err != nil {
		return nil, []error{l10n.NewRuntimeError("vapi.protocol.server.rest.param.body_parse_error",
			map[string]string{"errMsg": err.Error()})}
	}

	formURLMap := make(map[string]interface{})
	for key, value := range formURLBody {
		if key == "" {
			return nil, []error{l10n.NewRuntimeError("vapi.protocol.server.rest.param.body_parse_error",
				map[string]string{"errMsg": "Form URL encoded body contained key-value pair with no key."})}
		}
		if len(value) == 1 {
			formURLMap[key] = value[0]
		} else {
			s := make([]interface{}, len(value))
			for i, v := range value {
				s[i] = v
			}
			formURLMap[key] = s
		}
	}
	return formURLMap, nil
}

// Convert request body byte to value in DataValue type
func convertRequestBodyToDataValue(
	bodyByte []byte,
	bodyParamName string,
	bindingType bindings.BindingType,
	r *http.Request) (data.DataValue, []error) {

	if len(bodyByte) == 0 && reflect.TypeOf(bindingType) == bindings.OptionalBindingType {
		return data.NewOptionalValue(nil), nil
	}

	val, errs := getContentAsPerContentType(bodyByte, r)
	if errs != nil {
		return nil, errs
	}
	jsonToDataValueDecoder := cleanjson.NewJsonToDataValueDecoder()
	if reflect.TypeOf(bindingType) == bindings.ReferenceBindingType {
		refVal, valid := val.(map[string]interface{})
		if data, ok := refVal[bodyParamName]; ok && valid {
			return requestBodyToDataValue(data, bindingType, jsonToDataValueDecoder)
		}
	}
	return requestBodyToDataValue(val, bindingType, jsonToDataValueDecoder)
}

func resolveToCheckIfOptional(bindingType bindings.BindingType, tokens []string) bool {
	if reflect.TypeOf(bindingType) == bindings.OptionalBindingType {
		return true
	} else if refBinding, ok := bindingType.(bindings.ReferenceType); ok && len(tokens) > 1 {
		structBinding := refBinding.Resolve().(bindings.StructType)
		return resolveToCheckIfOptional(structBinding.Field(tokens[0]), tokens[1:])
	} else if structBinding, ok := bindingType.(bindings.StructType); ok && len(tokens) > 1 {
		return resolveToCheckIfOptional(structBinding.Field(tokens[0]), tokens[1:])
	}
	return false
}

func requestBodyToDataValue(
	val interface{},
	bindingType bindings.BindingType,
	jsonToDataValueDecoder *cleanjson.JsonToDataValueDecoder) (data.DataValue, []error) {

	switch reflect.TypeOf(bindingType) {
	case bindings.IntegerBindingType:
		if intStr, ok := val.(string); ok {
			i, err := strconv.ParseInt(intStr, 10, 64)
			if err != nil {
				return nil, []error{l10n.NewRuntimeError("vapi.protocol.server.rest.param.invalid_type",
					map[string]string{"paramName": "INTEGER", "expectedType": "Int", "errMsg": err.Error()})}
			}
			return data.NewIntegerValue(i), nil
		} else if jsonNum, ok := val.(json.Number); ok {
			intVal, err := jsonNum.Int64()
			if err != nil {
				return nil, []error{l10n.NewRuntimeError("vapi.protocol.server.rest.param.body.unexpected",
					map[string]string{"required": "INTEGER", "type": reflect.TypeOf(jsonNum).Name()})}
			}
			return data.NewIntegerValue(intVal), nil
		}
	case bindings.DoubleBindingType:
		if floatStr, ok := val.(string); ok {
			i, err := strconv.ParseFloat(floatStr, 64)
			if err != nil {
				return nil, []error{l10n.NewRuntimeError("vapi.protocol.server.rest.param.body.unexpected",
					map[string]string{"required": "FlOAT", "type": reflect.TypeOf(val).Name()})}
			}
			return data.NewDoubleValue(i), nil
		} else if jsonNum, ok := val.(json.Number); ok {
			floatVal, err := jsonNum.Float64()
			if err != nil {
				return nil, []error{l10n.NewRuntimeError("vapi.protocol.server.rest.param.body.unexpected",
					map[string]string{"required": "FlOAT", "type": reflect.TypeOf(jsonNum).Name()})}
			}
			return data.NewDoubleValue(floatVal), nil
		}
	case bindings.BooleanBindingType:
		if boolStr, ok := val.(string); ok {
			b, err := strconv.ParseBool(strings.ToLower(boolStr))
			if err != nil {
				return nil, []error{l10n.NewRuntimeError("vapi.protocol.server.rest.param.body.unexpected",
					map[string]string{"required": "BOOLEAN", "type": reflect.TypeOf(val).Name()})}
			}
			return data.NewBooleanValue(b), nil
		}
	case bindings.SecretBindingType:
		if val == nil {
			return nil, []error{l10n.NewRuntimeError("vapi.protocol.server.rest.param.body.unexpected",
				map[string]string{"required": "SECRET", "type": "null"})}
		}
		if value, ok := val.(string); ok {
			return data.NewSecretValue(value), nil
		}
		return nil, []error{l10n.NewRuntimeError("vapi.protocol.server.rest.param.body.unexpected",
			map[string]string{"required": "SECRET", "type": reflect.TypeOf(val).Name()})}
	case bindings.BlobBindingType:
		if value, ok := val.(string); ok {
			temp, err := base64.StdEncoding.DecodeString(value)
			if err != nil {
				return nil, []error{err}
			}
			return data.NewBlobValue(temp), nil
		}
		return nil, []error{l10n.NewRuntimeError("vapi.protocol.server.rest.param.body.unexpected",
			map[string]string{"required": "BLOB", "type": reflect.TypeOf(val).Name()})}

	case bindings.OptionalBindingType:
		if val == nil {
			return data.NewOptionalValue(nil), nil
		}
		eType := bindingType.(bindings.OptionalType).ElementType()
		dv, err := requestBodyToDataValue(val, eType, jsonToDataValueDecoder)
		if err != nil {
			return nil, err
		}
		return data.NewOptionalValue(dv), nil
	case bindings.ListBindingType:
		if val == nil {
			return nil, []error{l10n.NewRuntimeError("vapi.protocol.server.rest.param.body.unexpected",
				map[string]string{"required": "LIST", "type": "null"})}
		}
		// Check if interface value provided is of List type
		if _, ok := val.([]interface{}); !ok {
			return nil, []error{l10n.NewRuntimeError("vapi.protocol.server.rest.param.body.unexpected",
				map[string]string{"required": "LIST", "type": reflect.TypeOf(val).Name()})}
		}

		eType := bindingType.(bindings.ListType).ElementType()
		result := data.NewListValue()

		for _, value := range val.([]interface{}) {
			dv, err := requestBodyToDataValue(value, eType, jsonToDataValueDecoder)
			if err != nil {
				return nil, err
			}
			result.Add(dv)
		}
		return result, nil
	case bindings.SetBindingType:
		if val == nil {
			return nil, []error{l10n.NewRuntimeError("vapi.protocol.server.rest.param.body.unexpected",
				map[string]string{"required": "SET", "type": "null"})}
		}
		// Check if interface value provided is of SET type
		if _, ok := val.([]interface{}); !ok {
			return nil, []error{l10n.NewRuntimeError("vapi.protocol.server.rest.param.body.unexpected",
				map[string]string{"required": "SET", "type": reflect.TypeOf(val).Name()})}
		}

		eType := bindingType.(bindings.SetType).ElementType()
		result := data.NewListValue()

		setMap := map[string]int{}

		for _, value := range val.([]interface{}) {

			// Validate: If all values are unique, if not throw a internal server error
			keyStr := fmt.Sprintf("%v", value)
			if _, ok := setMap[keyStr]; ok {
				return nil, []error{l10n.NewRuntimeError("vapi.protocol.server.rest.param.internal_server_error",
					map[string]string{"msg": "Invalid SET value passed"})}
			}
			setMap[keyStr] = 1

			dv, err := requestBodyToDataValue(value, eType, jsonToDataValueDecoder)
			if err != nil {
				return nil, err
			}
			result.Add(dv)
		}
		return result, nil
	case bindings.MapBindingType:
		// Validate: value kv of map type
		eType := bindingType.(bindings.MapType).ValueType
		if mapVal, ok := val.(map[string]interface{}); ok {
			result := data.NewStructValue("", nil)
			for key, value := range mapVal {
				if value == nil && reflect.TypeOf(eType) != bindings.OptionalBindingType {
					return nil, []error{l10n.NewRuntimeError("vapi.protocol.server.rest.param.invalid_value",
						map[string]string{"errMsg": "Invalid Map value was provided"})}
				}
				dv, err := requestBodyToDataValue(value, eType, jsonToDataValueDecoder)
				if err != nil {
					return nil, err
				}
				result.SetField(key, dv)
			}
			return result, nil
		}
		return nil, []error{l10n.NewRuntimeError("vapi.protocol.server.rest.param.invalid_value",
			map[string]string{"errMsg": "Invalid MAP key/value was provided"})}
	case bindings.ReferenceBindingType:
		return requestBodyToDataValue(val, bindingType.(bindings.ReferenceType).Resolve(), jsonToDataValueDecoder)
	case bindings.StructBindingType:
		if _, ok := val.(map[string]interface{}); !ok {
			return nil, []error{l10n.NewRuntimeError("vapi.protocol.server.rest.param.body.unexpected",
				map[string]string{"required": "STRUCTURE", "type": reflect.TypeOf(val).Name()})}
		}
		retStructValue := data.NewStructValue(bindingType.(bindings.StructType).Name(), nil)
		structVal := val.(map[string]interface{})

		for _, fieldName := range bindingType.(bindings.StructType).FieldNames() {
			if fieldValue, ok := structVal[fieldName]; ok {
				dv, err := requestBodyToDataValue(fieldValue, bindingType.(bindings.StructType).Field(fieldName), jsonToDataValueDecoder)
				if err != nil {
					return nil, err
				}
				retStructValue.SetField(fieldName, dv)
			}
		}
		return retStructValue, nil
	case bindings.OpaqueBindingType:
		if val == nil {
			return nil, []error{l10n.NewRuntimeError("vapi.protocol.server.rest.param.invalid_value",
				map[string]string{"errMsg": "Invalid value was provided"})}
		}
	}
	bodyDV, err := jsonToDataValueDecoder.Decode(val)
	if err != nil {
		// TODO: will change to err after JsonToDataValueDecoder using l10n
		return nil, []error{err}
	}

	return bodyDV, nil
}

func ConvertToDataValueIfNested(container *data.StructValue, tokens []string, value data.DataValue, dataBinding bindings.BindingType) {
	currentField := tokens[0]
	if value == nil {
		return
	}

	if len(tokens) == 1 {
		container.SetField(currentField, value)
		return
	}

	if container.HasField(currentField) {
		subfield, _ := container.Field(currentField)
		parseExistingNestedParam(subfield, tokens, value, dataBinding)
	} else {
		createNewNestedParam(container, tokens, value, dataBinding)
	}
}

func parseExistingNestedParam(subfield data.DataValue, tokens []string, value data.DataValue, dataBinding bindings.BindingType) {
	if exisitingDataValue, ok := subfield.(*data.StructValue); ok {
		if refBinding, ok := dataBinding.(bindings.ReferenceType); ok {
			structBinding := refBinding.Resolve().(bindings.StructType)
			ConvertToDataValueIfNested(exisitingDataValue, tokens[1:], value, structBinding.Field(tokens[1]))
		} else if structBinding, ok := dataBinding.(bindings.StructType); ok {
			ConvertToDataValueIfNested(exisitingDataValue, tokens[1:], value, structBinding.Field(tokens[1]))
		} else if optBinding, ok := dataBinding.(bindings.OptionalType); ok {
			optElemBinding := optBinding.ElementType()
			if refBinding, ok := optElemBinding.(bindings.ReferenceType); ok {
				structBinding := refBinding.Resolve().(bindings.StructType)
				ConvertToDataValueIfNested(exisitingDataValue, tokens[1:], value, structBinding.Field(tokens[1]))
			} else if structBinding, ok := optElemBinding.(bindings.StructType); ok {
				ConvertToDataValueIfNested(exisitingDataValue, tokens[1:], value, structBinding.Field(tokens[1]))
			} else {
				log.Debugf("This case should never be reached, Optional's element binding was not handled %s : ", reflect.TypeOf(optElemBinding))
			}
		} else {
			log.Debugf("This case should never be reached, SubField type case was not handled %s : ", reflect.TypeOf(subfield))
		}
	} else if exisitingDataValue, ok := subfield.(*data.OptionalValue); ok {
		if exisitingDataValue.IsSet() {
			optionalSubfield := exisitingDataValue.Value()
			parseExistingNestedParam(optionalSubfield, tokens, value, dataBinding)
		}
	} else {
		// TODO: Return an error, but this case should never be reached.
	}
}

func createNewNestedParam(container *data.StructValue, tokens []string, value data.DataValue, dataBinding bindings.BindingType) {
	currentField := tokens[0]
	var newDataValue data.DataValue
	if refBinding, ok := dataBinding.(bindings.ReferenceType); ok {
		structBinding := refBinding.Resolve().(bindings.StructType)
		newDataValue = data.NewStructValue(structBinding.Name(), nil)
		ConvertToDataValueIfNested(newDataValue.(*data.StructValue), tokens[1:], value, structBinding.Field(tokens[1]))
	} else if structBinding, ok := dataBinding.(bindings.StructType); ok {
		newDataValue = data.NewStructValue(structBinding.Name(), nil)
		ConvertToDataValueIfNested(newDataValue.(*data.StructValue), tokens[1:], value, structBinding.Field(tokens[1]))
	} else if optBinding, ok := dataBinding.(bindings.OptionalType); ok {
		optElemBinding := optBinding.ElementType()
		var newStructValue *data.StructValue
		if refBinding, ok := optElemBinding.(bindings.ReferenceType); ok {
			structBinding := refBinding.Resolve().(bindings.StructType)
			newStructValue = data.NewStructValue(structBinding.Name(), nil)
			ConvertToDataValueIfNested(newStructValue, tokens[1:], value, structBinding.Field(tokens[1]))
		} else if structBinding, ok := optElemBinding.(bindings.StructType); ok {
			newStructValue = data.NewStructValue(structBinding.Name(), nil)
			ConvertToDataValueIfNested(newStructValue, tokens[1:], value, structBinding.Field(tokens[1]))
		} else {
			log.Debugf("This case should never be reached, Optional's element binding was not handled %s : ", reflect.TypeOf(optElemBinding))
		}
		newDataValue = data.NewOptionalValue(newStructValue)
	} else {
		log.Debugf("Unhandled case of sub structure binding %s : ", reflect.TypeOf(dataBinding))
		newDataValue = data.NewStructValue("", nil)
	}
	container.SetField(currentField, newDataValue)
}
