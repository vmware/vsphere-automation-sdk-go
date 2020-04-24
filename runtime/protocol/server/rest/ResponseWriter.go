/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

package rest

import (
	"encoding/base64"
	"net/http"
	"strconv"

	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data/serializers/cleanjson"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/l10n"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib/rest"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
)

// Create http response status based on method result, and
// annotated definition of succeed or exception cases.
func CreateResponseStatus(result core.MethodResult,
	opRestMetaData protocol.OperationRestMetadata) (statusCode int, status string, respError []error) {

	var found bool
	if result.IsSuccess() {
		for _, code := range rest.SUCCESSFUL_STATUS_CODES {
			if opRestMetaData.SuccessCode() == code {
				found = true
				statusCode = code
				break
			}
		}
		if !found {
			// Note: Sucessful response code not annotated is not an error case
			statusCode = http.StatusOK
		}
		status = http.StatusText(statusCode)
		if status == "" {
			status = http.StatusText(http.StatusOK)
		}
	} else {

		errorNameToStatusMap := opRestMetaData.ErrorCodeMap()
		if len(errorNameToStatusMap) > 0 {
			errorName := result.Error().Name()
			if httpErrorCode, ok := errorNameToStatusMap[errorName]; ok {
				statusCode = httpErrorCode
				if httpErrorCode < 500 {
					status = errorName
				} else {
					status = "Internal Server Error"
				}
			} else {
				statusCode = rest.HTTP_500_INTERNAL_SERVER_ERROR
				status = "Internal Server Error"
				return statusCode, status, []error{l10n.NewRuntimeError("vapi.protocol.server.rest.error.not_supported",
					map[string]string{"errorName": errorName})}
			}
		} else {
			errorName := result.Error().Name()
			if httpErrorCode, ok := rest.VAPI_TO_HTTP_ERROR_MAP[errorName]; ok {
				statusCode = httpErrorCode
				status = http.StatusText(statusCode)
			} else {
				statusCode = rest.HTTP_500_INTERNAL_SERVER_ERROR
				status = "Internal Server Error"
				return statusCode, status, []error{l10n.NewRuntimeError("vapi.protocol.server.rest.error.not_supported",
					map[string]string{"errorName": errorName})}
			}

			if status == "" {
				status = http.StatusText(http.StatusInternalServerError)
			}
		}
	}

	return statusCode, status, nil
}

// Create http response header based on method result, and
// annotated definition of result to header name map.
// If the method throws error, it will set header value with
// the annotated field in error value.
func SetResponseHeader(result core.MethodResult,
	resultToHeaderMap map[string]string,
	errorToHeaderMap map[string]string,
	header http.Header) []error {

	if !result.IsSuccess() {
		return setResponseHeader(result.Error(), errorToHeaderMap, header)
	}
	return setResponseHeader(result.Output(), resultToHeaderMap, header)
}

// Create http response body based on method result and
// restmetadata to w.r.t result headers and response body name
func CreateResponseBody(result core.MethodResult,
	metadata protocol.OperationRestMetadata) (string, []error) {

	if !result.IsSuccess() {
		// Only pack the Error field
		// errorDV := result.Error()
		// errorStruct, err := filterBodyFields(errorDV, errorFieldList)
		// if err != nil {
		// 	return "", err
		// }
		return setResponseBody(result.Error())
	}

	respDV := result.Output()

	switch respDV.Type() {
	case data.VOID:
		// Return null if result is void
		return "null", nil
	case data.STRUCTURE:
		// For structure type response, filter out fields with @Header annotation
		resultHeadersNameMap := metadata.ResultHeadersNameMap()
		resStruct, ok := respDV.(*data.StructValue)
		if !ok || (len(resultHeadersNameMap) == 0) {
			return setResponseBody(resStruct)
		}
		structName := resStruct.Name()
		structFields := resStruct.Fields()

		output := data.NewStructValue(structName, nil)

		for field, value := range structFields {
			if _, ok := resultHeadersNameMap[field]; !ok {
				output.SetField(field, value)
			}
		}

		if bodyName := metadata.ResponseBodyName(); bodyName != "" {
			if fieldVal, err := output.Field(bodyName); err == nil {
				return setResponseBody(fieldVal)
			}
		}

		return setResponseBody(output)
	default:
		// For non-structure response, we reply on DataValueToJsonEncoder
		// to serialize the value.
		return setResponseBody(respDV)
	}
}

func convertDVToHeaderValue(name string, respDV data.DataValue, header http.Header) []error {
	var err []error

	switch respDV.Type() {
	case data.BOOLEAN:
		valBool := respDV.(*data.BooleanValue)
		valueString := strconv.FormatBool(valBool.Value())
		header.Add(name, valueString)
	case data.INTEGER:
		valInt := respDV.(*data.IntegerValue)
		valueString := strconv.FormatInt(valInt.Value(), 10)
		header.Add(name, valueString)
	case data.DOUBLE:
		valDouble := respDV.(*data.DoubleValue)
		valueString := strconv.FormatFloat(valDouble.Value(), 'f', -1, 64)
		header.Add(name, valueString)
	case data.STRING:
		valString := respDV.(*data.StringValue)
		valueString := valString.Value()
		header.Add(name, valueString)
	case data.SECRET:
		valSecret := respDV.(*data.SecretValue)
		valueString := valSecret.Value()
		header.Add(name, valueString)
	case data.BLOB:
		valBlob := respDV.(*data.BlobValue)
		valueString := base64.StdEncoding.EncodeToString(valBlob.Value())
		header.Add(name, valueString)
	case data.LIST:
		valList := respDV.(*data.ListValue)
		valSlice := valList.List()
		for _, elementDv := range valSlice {
			if elementDv.Type() == data.OPTIONAL {
				return []error{l10n.NewRuntimeError("vapi.protocol.server.rest.response.unsupport_type",
					map[string]string{"dataType": "LIST<OPTIONAL>", "fieldName": name})}
			}
			if err = convertDVToHeaderValue(name, elementDv, header); err != nil {
				return err
			}
		}
	case data.OPTIONAL:
		optVal := respDV.(*data.OptionalValue)
		if !optVal.IsSet() {
			return nil
		}
		elemVal := optVal.Value()
		return convertDVToHeaderValue(name, elemVal, header)
	default:
		return []error{l10n.NewRuntimeError("vapi.protocol.server.rest.response.unsupport_type",
			map[string]string{"dataType": respDV.Type().String(), "fieldName": name})}
	}

	return nil
}

// Set the response value to http header
func setResponseHeader(respDV data.DataValue,
	resultToHeaderMap map[string]string,
	header http.Header) []error {

	var respMap map[string]data.DataValue

	if len(resultToHeaderMap) == 0 {
		return nil
	}

	switch respDV.Type() {
	case data.ERROR:
		errorStruct, ok := respDV.(*data.ErrorValue)
		if !ok {
			return []error{l10n.NewRuntimeErrorNoParam("vapi.protocol.server.rest.response.error_not_structure")}
		}
		respMap = errorStruct.Fields()
	case data.STRUCTURE:
		respStruct, ok := respDV.(*data.StructValue)
		if !ok {
			return []error{l10n.NewRuntimeErrorNoParam("vapi.protocol.server.rest.response.not_structure")}
		}
		respMap = respStruct.Fields()
	default:
		return []error{l10n.NewRuntimeErrorNoParam("vapi.protocol.server.rest.response.not_structure")}
	}

	var err []error
	var headerName string
	var exist bool
	for name, val := range respMap {
		if headerName, exist = resultToHeaderMap[name]; !exist {
			continue
		}
		// There is Header annotation for this field
		valType := val.Type()
		switch valType {
		case data.STRUCTURE:
			return setResponseHeader(val, resultToHeaderMap, header)
		default:
			if err = convertDVToHeaderValue(headerName, val, header); err != nil {
				return err
			}
		}
	}
	return nil
}

// Set response value to http response body
func setResponseBody(respDV data.DataValue) (string, []error) {
	var dataValueToJsonEncoder = cleanjson.NewDataValueToJsonEncoder()

	bodyString, err := dataValueToJsonEncoder.Encode(respDV)

	if err != nil {
		return "",
			[]error{l10n.NewRuntimeErrorNoParam("vapi.protocol.server.rest.response.body_parse_error")}
	}
	return bodyString, nil
}

// Filter out structure fields with @Header annotation
func filterBodyFields(respDV data.DataValue,
	bodyFieldList []string) (*data.StructValue, []error) {
	var respMap map[string]data.DataValue
	var dvName string

	if respDV.Type() == data.ERROR {
		errorStruct, ok := respDV.(*data.ErrorValue)
		if !ok {
			return nil,
				[]error{l10n.NewRuntimeErrorNoParam("vapi.protocol.server.rest.response.error_not_structure")}
		}
		respMap = errorStruct.Fields()
		dvName = errorStruct.Name()
	} else {
		respStruct, ok := respDV.(*data.StructValue)
		if !ok {
			return nil,
				[]error{l10n.NewRuntimeErrorNoParam("vapi.protocol.server.rest.response.not_structure")}
		}
		respMap = respStruct.Fields()
		dvName = respStruct.Name()
	}

	// For structure type response, we iterate through each fields and only
	// assemble fields without @Header annotation
	bodyFields := make(map[string]data.DataValue)
	for k, v := range respMap {
		if isStringInList(k, bodyFieldList) {
			bodyFields[k] = v
		}
	}
	return data.NewStructValue(dvName, bodyFields), nil
}

// Check whether a string is in a slice
func isStringInList(str string, list []string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}
