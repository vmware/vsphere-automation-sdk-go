/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

package rest

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data/serializers/cleanjson"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/l10n"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/log"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
	"io/ioutil"
	"net/http"
	"net/textproto"
	"reflect"
	"strconv"
	"strings"
)

type ParamsType string

const (
	queryParamsType  ParamsType = "query"
	pathParamsType   ParamsType = "path"
	headerParamsType ParamsType = "header"
)

// Parse the URL path variable. It takes annotation name to param name map,
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

// Parse the http header. It takes annotation name to param name map,
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

// Parse the http query. It takes annotation name to param name map,
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

func QueryDataValueMap(request *http.Request,
	metadata protocol.OperationRestMetadata) (map[string]data.DataValue, []error) {

	// Get map of query param name to its value (slice)
	queryParamMap := request.URL.Query()
	return generateParamDataValueMap(queryParamMap,
		reverseMap(metadata.QueryParamsNameMap()),
		metadata.ParamsTypeMap(),
		queryParamsType)
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
	paramsType ParamsType) (data.DataValue, []error) {
	switch reflect.TypeOf(pType) {
	case bindings.StringBindingType, bindings.IdBindingType:
		// If multiple headers are provided for String Binding then we need to merge them
		if paramsType == headerParamsType {
			return data.NewStringValue(strings.Join(pValue, ",")), nil
		}
		if len(pValue) > 1 {
			return nil, []error{l10n.NewRuntimeError("vapi.protocol.server.rest.param.invalid_value",
				map[string]string{"errMsg": "Expected type STRING but got 'LIST'"})}
		}
		return data.NewStringValue(pValue[0]), nil
	case bindings.SecretBindingType:
		return data.NewSecretValue(pValue[0]), nil
	case bindings.BlobBindingType:
		temp, err := base64.StdEncoding.DecodeString(pValue[0])
		if err != nil {
			log.Error(l10n.NewRuntimeError("vapi.protocol.server.rest.param.invalid_value",
				map[string]string{"errMsg": err.Error()}))
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
		i, err := strconv.ParseInt(pValue[0], 10, 64)
		if err != nil {
			return nil, []error{l10n.NewRuntimeError("vapi.protocol.server.rest.param.invalid_type",
				map[string]string{"paramName": name, "expectedType": "Int", "errMsg": err.Error()})}
		}
		return data.NewIntegerValue(i), nil
	case bindings.ListBindingType:
		eType := pType.(bindings.ListType).ElementType()
		return convertListParamElementsToDataValue(name, pValue, eType, paramsType)
	case bindings.SetBindingType:
		eType := pType.(bindings.SetType).ElementType()
		return convertListParamElementsToDataValue(name, pValue, eType, paramsType)
	case bindings.OptionalBindingType:
		eType := pType.(bindings.OptionalType).ElementType()
		dv, err := convertRequestParamToDataValue(name, pValue, eType, paramsType)
		if err != nil {
			return nil, err
		}
		result := data.NewOptionalValue(dv)
		return result, nil
	case bindings.EnumBindingType:
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
	paramsType ParamsType) (data.DataValue, []error) {
	result := data.NewListValue()
	for _, value := range pValue {
		dv, err := convertRequestParamToDataValue(name, []string{value}, eType, paramsType)
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
			paramValueMap[name], errIn = convertRequestParamToDataValue(annotationName, valuelist, ptype, paramsType)
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

	fields := listBodyFieldsParams(metadata)
	if len(fields) == 0 {
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

	jdecoder := json.NewDecoder(strings.NewReader(string(bodyByte)))
	jdecoder.UseNumber()

	var val interface{}
	err = jdecoder.Decode(&val)
	if err != nil {
		return []error{l10n.NewRuntimeError("vapi.protocol.server.rest.param.body_parse_error",
			map[string]string{"errMsg": err.Error()})}
	}

	jsonToDataValueDecoder := cleanjson.NewJsonToDataValueDecoder()
	fieldsBindingMap := metadata.Fields()
	fieldNameMap := metadata.FieldNameMap()

	valMap, ok := val.(map[string]interface{})
	if !ok {
		return []error{l10n.NewRuntimeErrorNoParam("vapi.protocol.server.rest.param.bodyfield_parse_invalid")}
	}
	for fieldName, fieldValue := range valMap {
		if _, ok := fieldNameMap[fieldName]; !ok {
			continue
		} else if stringInSlice(fieldName, fields) {
			tempData, err := requestBodyToDataValue(fieldValue, fieldsBindingMap[fieldName], jsonToDataValueDecoder)
			if err != nil {
				return err
			}
			inputDataVal.SetField(fieldName, tempData)
		}
	}

	return nil
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func listBodyFieldsParams(metadata protocol.OperationRestMetadata) []string {

	resMap := map[string]int{}
	res := []string{}

	// Note: We dont have to handle the body key being in fieldsNameMap
	// because Body and BodyField are both mutually exclusive. Hence
	// If its a @Body case then this function would never be called.

	for k, _ := range metadata.FieldNameMap() {
		resMap[k] = 1
	}

	for k, _ := range metadata.HeaderParamsNameMap() {
		delete(resMap, k)
	}

	for k, _ := range metadata.PathParamsNameMap() {
		delete(resMap, k)
	}

	for k, _ := range metadata.QueryParamsNameMap() {
		delete(resMap, k)
	}

	for k, _ := range resMap {
		res = append(res, k)
	}
	return res
}

/*
	Body param processing
*/

// Parse the http body. There is at most one @Body annotation per operation.
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
		bodyDataVal, er = convertRequestBodyToDataValue(bodyByte, bodyBindingType)
		if er != nil {
			return er
		}
		ConvertToDataValueIfNested(inputDataVal, tokens, bodyDataVal, metadata.Fields()[tokens[0]])
	} else {
		bodyDataVal, er := convertRequestBodyToDataValue(bodyByte, bodyBindingType)
		if er != nil {
			return er
		}
		inputDataVal.SetField(bodyParamName, bodyDataVal)
	}
	return nil
}

// Convert request body byte to value in DataValue type
func convertRequestBodyToDataValue(
	bodyByte []byte,
	bindingType bindings.BindingType) (data.DataValue, []error) {
	var err error

	if len(bodyByte) == 0 && reflect.TypeOf(bindingType) == bindings.OptionalBindingType {
		return data.NewOptionalValue(nil), nil
	}

	byteData := strings.NewReader(string(bodyByte))

	jdecoder := json.NewDecoder(byteData)
	jdecoder.UseNumber()
	var val interface{}
	err = jdecoder.Decode(&val)
	if err != nil {
		return nil,
			[]error{l10n.NewRuntimeError("vapi.protocol.server.rest.param.body_parse_error",
				map[string]string{"errMsg": err.Error()})}
	}
	jsonToDataValueDecoder := cleanjson.NewJsonToDataValueDecoder()
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
	case bindings.DoubleBindingType:
		if floatStr, ok := val.(string); ok {
			i, err := strconv.ParseFloat(floatStr, 64)
			if err != nil {
				return nil, []error{l10n.NewRuntimeError("vapi.protocol.server.rest.param.body.unexpected",
					map[string]string{"required": "FlOAT", "type": reflect.TypeOf(val).Name()})}
			}
			return data.NewDoubleValue(i), nil
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
		// Check if interface value provided is not of List type
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
			} else {
				result.Add(dv)
			}
		}
		return result, nil

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
