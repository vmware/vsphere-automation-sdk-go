/* Copyright © 2019-2020 VMware, Inc. All Rights Reserved.
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
	"strings"

	"github.com/gorilla/mux"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data/serializers/cleanjson"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/l10n"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/log"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
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
		// if router is initialized with UseEncodedPath, we'd need to
		// decode parameter value first

		v = PrepareVAPIQuery(v)
		v, err := url.QueryUnescape(v)
		if err != nil {
			return nil, []error{l10n.NewRuntimeError("vapi.protocol.server.rest.param.invalid_value",
				map[string]string{"errMsg": err.Error()})}
		}
		pathParamMap[k] = []string{v}
	}
	return generateParamDataValueMap(converToSliceOfPointers(pathParamMap),
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
	return generateParamDataValueMap(converToSliceOfPointers(headerParamMap),
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

/** Parses dispatch params string available in generated rest metadata
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
	queryParamMap := GetQuery(request)
	requestKeys := []string{}
	for k := range queryParamMap {
		// task query param is vAPI preserved for task related operations
		// should not be possible to be @Query annotated in VMODL
		if k == lib.TaskRESTQueryKey {
			continue
		}
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
		return nil, []error{l10n.NewRuntimeError("vapi.protocol.server.rest.unsupported_property",
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
	pValue []*string,
	pType bindings.BindingType,
	paramsType ParamsType,
	insideOptional bool) (data.DataValue, []error) {

	if paramsType == headerParamsType {
		mergedValue, err := mergeHeaderValues(pValue)
		if err != nil {
			return nil, err
		}
		pValue = mergedValue
	}

	if ok, err := checkParam(pValue, pType); !ok {
		return nil, err
	}

	switch reflect.TypeOf(pType) {
	case bindings.StringBindingType:
		return data.NewStringValue(*pValue[0]), nil
	case bindings.IdBindingType:
		return data.NewStringValue(*pValue[0]), nil
	case bindings.SecretBindingType:
		return data.NewSecretValue(*pValue[0]), nil
	case bindings.BlobBindingType:
		var temp []byte
		var err error
		temp, err = base64.StdEncoding.DecodeString(*pValue[0])
		if err != nil && paramsType == headerParamsType {
			log.Errorf("Error ignored; Header param return valid prefix data: " + err.Error())
		} else if err != nil && insideOptional {
			log.Errorf("Error ignored; BLOB inside Optional: " + err.Error())
		} else if err != nil {
			temp, err = base64.URLEncoding.DecodeString(*pValue[0])
			if err != nil {
				return nil, []error{l10n.NewRuntimeError("vapi.protocol.server.rest.param.invalid_value",
					map[string]string{"errMsg": err.Error()})}
			}
		}
		return data.NewBlobValue(temp), nil
	case bindings.BooleanBindingType:
		var boolVal bool
		var reqVal string
		if paramsType != pathParamsType {
			reqVal = strings.ToLower(*pValue[0])
		} else {
			reqVal = *pValue[0]
		}
		err := json.Unmarshal([]byte(reqVal), &boolVal)
		if err != nil {
			return nil, []error{l10n.NewRuntimeError("vapi.protocol.server.rest.param.invalid_type",
				map[string]string{"paramName": name, "expectedType": "Bool", "errMsg": err.Error()})}
		}
		return data.NewBooleanValue(boolVal), nil
	case bindings.DoubleBindingType:
		var floatVal float64
		err := json.Unmarshal([]byte(*pValue[0]), &floatVal)
		if err != nil {
			return nil, []error{l10n.NewRuntimeError("vapi.protocol.server.rest.param.invalid_type",
				map[string]string{"paramName": name, "expectedType": "Float", "errMsg": err.Error()})}
		}
		return data.NewDoubleValue(floatVal), nil
	case bindings.IntegerBindingType:
		var intVal int64
		err := json.Unmarshal([]byte(*pValue[0]), &intVal)
		if err != nil {
			return nil, []error{l10n.NewRuntimeError("vapi.protocol.server.rest.param.invalid_type",
				map[string]string{"paramName": name, "expectedType": "Int", "errMsg": err.Error()})}
		}
		return data.NewIntegerValue(intVal), nil
	case bindings.ListBindingType:
		eType := pType.(bindings.ListType).ElementType()
		return convertListParamElementsToDataValue(name, pValue, eType, paramsType, insideOptional)
	case bindings.SetBindingType:
		eType := pType.(bindings.SetType).ElementType()
		return convertListParamElementsToDataValue(name, pValue, eType, paramsType, insideOptional)
	case bindings.OptionalBindingType:
		eType := pType.(bindings.OptionalType).ElementType()
		if pValue == nil || pValue[0] == nil {
			return data.NewOptionalValue(nil), nil
		}
		dv, err := convertRequestParamToDataValue(name, pValue, eType, paramsType, true)
		if err != nil {
			return nil, err
		}
		return data.NewOptionalValue(dv), nil
	case bindings.EnumBindingType:
		if paramsType == headerParamsType && *pValue[0] == "" && insideOptional {
			return nil, nil
		}
		return data.NewStringValue(*pValue[0]), nil
	case bindings.DateTimeBindingType:
		return data.NewStringValue(*pValue[0]), nil
	default:
		// Note: Spec state that Map is not supported as query/header parameter
		dataType := fmt.Sprintf("%v", pType)
		return nil, []error{l10n.NewRuntimeError("vapi.protocol.server.rest.param.unsupport_type",
			map[string]string{"dataType": dataType})}
	}
}

func mergeHeaderValues(value []*string) ([]*string, []error) {
	if value == nil || value[0] == nil {
		return nil, []error{l10n.NewRuntimeError("vapi.protocol.server.rest.param.invalid_value",
			map[string]string{"errMsg": "Expected type different from 'nil'"})}
	}

	var headerValues []string
	for _, val := range value {
		valStr := *val
		if val == nil {
			// this should actually never occur but let's put default behavior for just in case
			valStr = ""
		}
		headerValues = append(headerValues, valStr)
	}
	newHeaderValue := strings.Join(headerValues, ",")
	value = []*string{&newHeaderValue}
	return value, nil
}

func checkParam(value []*string, paramType bindings.BindingType) (bool, []error) {
	switch reflect.TypeOf(paramType) {
	case bindings.ListBindingType:
		fallthrough
	case bindings.SetBindingType:
		fallthrough
	case bindings.OptionalBindingType:
		return true, nil
	}

	if value == nil || value[0] == nil {
		return false, []error{l10n.NewRuntimeError("vapi.protocol.server.rest.param.invalid_value",
			map[string]string{"errMsg": "Expected type different from 'nil'"})}
	} else if len(value) > 1 {
		return false, []error{l10n.NewRuntimeError("vapi.protocol.server.rest.param.invalid_value",
			map[string]string{"errMsg": "Expected type different from 'LIST'"})}
	}

	return true, nil
}

// Convert List/Set's element value to a DataValue
func convertListParamElementsToDataValue(
	name string,
	pValue []*string,
	eType bindings.BindingType,
	paramsType ParamsType,
	insideOptional bool) (data.DataValue, []error) {
	result := data.NewListValue()
	for _, value := range pValue {
		dv, err := convertRequestParamToDataValue(name, []*string{value}, eType, paramsType, insideOptional)
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
func generateParamDataValueMap(paramStringMap map[string][]*string,
	paramNameMap map[string]string,
	paramTypeMap map[string]bindings.BindingType, paramsType ParamsType) (map[string]data.DataValue, []error) {

	paramValueMap := map[string]data.DataValue{}

	for annotationName, valueList := range paramStringMap {

		_, found := paramNameMap[annotationName]
		if !found {
			// The query name is not annotated, skip
			continue

		}
		name, ptype, errIn := findParamNameType(annotationName,
			paramNameMap,
			paramTypeMap)

		if errIn == nil {
			paramValueMap[name], errIn = convertRequestParamToDataValue(annotationName, valueList, ptype, paramsType, false)
		}
		if errIn != nil {
			return nil, errIn
		}
	}
	return paramValueMap, nil
}

func converToSliceOfPointers(values map[string][]string) map[string][]*string {
	result := make(map[string][]*string)
	for k, v := range values {
		list := make([]*string, len(v), len(v))
		for index, item := range v {
			itemPtr := item
			list[index] = &itemPtr
		}
		result[k] = list
	}
	return result
}

/*
	Body Field param processing
*/

func ParseBodyFieldParams(
	r *http.Request,
	metadata protocol.OperationRestMetadata,
	inputDataVal *data.StructValue) []error {

	bodyFieldsMap := metadata.BodyFieldsMap()
	if len(bodyFieldsMap) == 0 {
		return nil
	}

	if r.ContentLength == 0 {
		return nil
	}

	contentType := r.Header.Get(lib.HTTP_CONTENT_TYPE_HEADER)
	val, errs := getContentAsPerContentType(contentType, r)
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
			tempData, err := requestBodyToDataValue(fieldValue, fieldsBindingMap[paramName], jsonToDataValueDecoder, contentType)
			if err != nil {
				return err
			}
			inputDataVal.SetField(paramName, tempData)
		} else if fieldValue != nil {
			return []error{l10n.NewRuntimeError("vapi.protocol.server.rest.unsupported_property",
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

	tokens := strings.Split(bodyParamName, ".")
	bodyBindingType := metadata.ParamsTypeMap()[bodyParamName]
	if len(tokens) > 1 {
		var bodyDataVal data.DataValue
		if r.ContentLength == 0 && resolveToCheckIfOptional(metadata.Fields()[tokens[0]], tokens[1:]) {
			return nil
		}
		var er []error
		bodyDataVal, er = convertRequestBodyToDataValue(r, bodyParamName, bodyBindingType)
		if er != nil {
			return er
		}
		ConvertToDataValueIfNested(inputDataVal, tokens, bodyDataVal, metadata.Fields()[tokens[0]])
	} else {
		bodyDataVal, er := convertRequestBodyToDataValue(r, bodyParamName, bodyBindingType)
		if er != nil {
			return er
		}
		inputDataVal.SetField(bodyParamName, bodyDataVal)
	}
	return nil
}

func getContentAsPerContentType(
	contentType string,
	r *http.Request) (interface{}, []error) {

	if contentType == lib.JSON_CONTENT_TYPE {
		return getContentAppJSON(r)
	} else if contentType == lib.FORM_URL_CONTENT_TYPE {
		return getContentAppFormURLEncoded(r)
	}
	return nil, []error{l10n.NewRuntimeErrorNoParam("vapi.protocol.server.rest.unsupported_media_type")}
}

func getContentAppJSON(r *http.Request) (interface{}, []error) {

	body, err := getBodyContent(r)
	if err != nil {
		return nil, err
	}

	// Validate: Check if its valid Json body
	// Note: json Unmarshal is not sufficient to cover all inputs cases hence only use it to validate
	// the body, json.NewDecoder can even decode non json structures hence can't be used for validation.
	var dummyHolder interface{}
	jsonErr := json.Unmarshal(body, &dummyHolder)
	if jsonErr != nil {
		return nil, []error{l10n.NewRuntimeError("vapi.protocol.server.rest.invalid_json_content",
			map[string]string{"errMsg": jsonErr.Error()})}
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

func getBodyContent(r *http.Request) ([]byte, []error) {
	if r.Body == nil {
		return nil, []error{l10n.NewRuntimeError("vapi.protocol.server.rest.param.body_parse_error",
			map[string]string{"errMsg": "http request body is not present"})}
	}

	var bodyByte []byte
	bodyByte, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, []error{l10n.NewRuntimeError("vapi.protocol.server.rest.param.body_parse_error",
			map[string]string{"errMsg": err.Error()})}
	}
	return bodyByte, nil
}

func getContentAppFormURLEncoded(r *http.Request) (interface{}, []error) {
	err := r.ParseForm()
	if err != nil {
		return nil, []error{l10n.NewRuntimeError("vapi.protocol.server.rest.param.body_parse_error",
			map[string]string{"errMsg": err.Error()})}
	}

	formURLMap := make(map[string]interface{})
	for key, value := range r.Form {
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
	r *http.Request,
	bodyParamName string,
	bindingType bindings.BindingType) (data.DataValue, []error) {

	if r.ContentLength == 0 && reflect.TypeOf(bindingType) == bindings.OptionalBindingType {
		return data.NewOptionalValue(nil), nil
	}

	contentType := r.Header.Get(lib.HTTP_CONTENT_TYPE_HEADER)
	val, errs := getContentAsPerContentType(contentType, r)
	if errs != nil {
		return nil, errs
	}
	jsonToDataValueDecoder := cleanjson.NewJsonToDataValueDecoder()
	if reflect.TypeOf(bindingType) == bindings.ReferenceBindingType {
		refVal, valid := val.(map[string]interface{})
		if data, ok := refVal[bodyParamName]; ok && valid {
			return requestBodyToDataValue(data, bindingType, jsonToDataValueDecoder, contentType)
		}
	}
	return requestBodyToDataValue(val, bindingType, jsonToDataValueDecoder, contentType)
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
	jsonToDataValueDecoder *cleanjson.JsonToDataValueDecoder,
	contentType string) (data.DataValue, []error) {

	switch reflect.TypeOf(bindingType) {
	case bindings.IntegerBindingType:
		if contentType == lib.JSON_CONTENT_TYPE {
			if _, ok := val.(json.Number); !ok {
				return nil, []error{l10n.NewRuntimeError("vapi.protocol.server.rest.param.body.unexpected",
					map[string]string{"required": "INTEGER", "type": reflect.TypeOf(val).Name()})}
			}
		} else {
			if intStr, ok := val.(string); ok {
				var intVal int64
				err := json.Unmarshal([]byte(intStr), &intVal)
				if err != nil {
					return nil, []error{l10n.NewRuntimeError("vapi.protocol.server.rest.param.invalid_type",
						map[string]string{"paramName": "INTEGER", "expectedType": "Int", "errMsg": err.Error()})}
				}
				return data.NewIntegerValue(intVal), nil
			} else if jsonNum, ok := val.(json.Number); ok {
				intVal, err := jsonNum.Int64()
				if err != nil {
					return nil, []error{l10n.NewRuntimeError("vapi.protocol.server.rest.param.body.unexpected",
						map[string]string{"required": "INTEGER", "type": reflect.TypeOf(jsonNum).Name()})}
				}
				return data.NewIntegerValue(intVal), nil
			}
		}
	case bindings.DoubleBindingType:
		if contentType == lib.JSON_CONTENT_TYPE {
			if _, ok := val.(json.Number); !ok {
				return nil, []error{l10n.NewRuntimeError("vapi.protocol.server.rest.param.body.unexpected",
					map[string]string{"required": "FLOAT", "type": reflect.TypeOf(val).Name()})}
			}
		} else {
			if floatStr, ok := val.(string); ok {
				var floatVal float64
				err := json.Unmarshal([]byte(floatStr), &floatVal)
				if err != nil {
					return nil, []error{l10n.NewRuntimeError("vapi.protocol.server.rest.param.body.unexpected",
						map[string]string{"required": "FLOAT", "type": reflect.TypeOf(val).Name()})}
				}
				return data.NewDoubleValue(floatVal), nil
			} else if jsonNum, ok := val.(json.Number); ok {
				floatVal, err := jsonNum.Float64()
				if err != nil {
					return nil, []error{l10n.NewRuntimeError("vapi.protocol.server.rest.param.body.unexpected",
						map[string]string{"required": "FLOAT", "type": reflect.TypeOf(jsonNum).Name()})}
				}
				return data.NewDoubleValue(floatVal), nil
			}
		}
	case bindings.BooleanBindingType:
		if contentType == lib.JSON_CONTENT_TYPE {
			if _, ok := val.(bool); !ok {
				return nil, []error{l10n.NewRuntimeError("vapi.protocol.server.rest.param.body.unexpected",
					map[string]string{"required": "BOOLEAN", "type": reflect.TypeOf(val).Name()})}
			}
		} else {
			if strVal, ok := val.(string); ok {
				var boolVal bool
				strVal = strings.ToLower(strVal)
				err := json.Unmarshal([]byte(strVal), &boolVal)
				if err != nil {
					return nil, []error{l10n.NewRuntimeError("vapi.protocol.server.rest.param.invalid_type",
						map[string]string{"paramName": "BOOLEAN", "expectedType": "Int", "errMsg": err.Error()})}
				}
				return data.NewBooleanValue(boolVal), nil
			} else if boolVal, ok := val.(bool); ok {
				return data.NewBooleanValue(boolVal), nil
			}
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
		dv, err := requestBodyToDataValue(val, eType, jsonToDataValueDecoder, contentType)
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
			dv, err := requestBodyToDataValue(value, eType, jsonToDataValueDecoder, contentType)
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
				return nil, []error{l10n.NewRuntimeError("vapi.protocol.server.rest.param.invalid_argument",
					map[string]string{"msg": "Invalid SET value passed"})}
			}
			setMap[keyStr] = 1

			dv, err := requestBodyToDataValue(value, eType, jsonToDataValueDecoder, contentType)
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
				dv, err := requestBodyToDataValue(value, eType, jsonToDataValueDecoder, contentType)
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
		return requestBodyToDataValue(val, bindingType.(bindings.ReferenceType).Resolve(), jsonToDataValueDecoder, contentType)
	case bindings.StructBindingType:
		if _, ok := val.(map[string]interface{}); !ok {
			return nil, []error{l10n.NewRuntimeError("vapi.protocol.server.rest.param.body.unexpected",
				map[string]string{"required": "STRUCTURE", "type": reflect.TypeOf(val).Name()})}
		}
		retStructValue := data.NewStructValue(bindingType.(bindings.StructType).Name(), nil)
		structVal := val.(map[string]interface{})

		for _, fieldName := range bindingType.(bindings.StructType).FieldNames() {
			if fieldValue, ok := structVal[fieldName]; ok {
				dv, err := requestBodyToDataValue(fieldValue, bindingType.(bindings.StructType).Field(fieldName), jsonToDataValueDecoder, contentType)
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
