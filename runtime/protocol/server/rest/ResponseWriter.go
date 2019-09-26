package rest

/* **********************************************************
 * Copyright 2019 VMware, Inc.  All rights reserved.
 *      -- VMware Confidential
 * **********************************************************/

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/core"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data/serializers/cleanjson"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/l10n"
	"net/http"
	"strconv"
)

// Emitter need to map "SUCCESS" to the annotated operation response http status
const successCode string = "SUCCESS"

// Create http response status based on method result, and
// annotated definition of succeed or exception cases.
// Note: For success case, statusCodeMap["SUCCESS"] = annotated resp code.
//       For error case, statusCodeMap[errorName] = error resp code.
func CreateResponseStatus(result core.MethodResult,
	statusCodeMap map[string]int) (statusCode int, status string, respError []error) {

	var found bool
	if result.IsSuccess() {
		statusCode, found = statusCodeMap[successCode]
		if !found {
			// Note: Sucessful response code not annotated is not an error case
			statusCode = http.StatusOK
		}
		status = http.StatusText(statusCode)
		if status == "" {
			status = http.StatusText(http.StatusOK)
		}
	} else {
		errorName := result.Error().Name()
		statusCode, found = statusCodeMap[errorName]
		if !found {
			// Note: Error response code not annotated, set to default
			//       internal server error.
			statusCode = http.StatusInternalServerError
			status = http.StatusText(statusCode)
			return statusCode, status,
				[]error{l10n.NewRuntimeError("vapi.protocol.server.rest.response.unsupport_http_status",
					map[string]string{"httpStatus": errorName})}
		}

		status = http.StatusText(statusCode)
		if status == "" {
			status = http.StatusText(http.StatusInternalServerError)
		}
	}

	return statusCode, status, nil
}

// Create http response header based on method result, and
// annotated definition of result to header name map.
// If the method throws error, it will set header value with
// the annotated field in error value.
func CreateResponseHeader(result core.MethodResult,
	resultToHeaderMap map[string]string,
	errorToHeaderMap map[string]string) (http.Header, []error) {

	if !result.IsSuccess() {
		return setResponseHeader(result.Error(), errorToHeaderMap)
	}
	return setResponseHeader(result.Output(), resultToHeaderMap)
}

// Create http response body based on method result, and
// annotated definition of result to header name map.
func CreateResponseBody(result core.MethodResult,
	bodyFieldList []string,
	errorFieldList []string) (string, []error) {
	if !result.IsSuccess() {
		// Only pack the Error field
		errorDV := result.Error()
		errorStruct, err := filterBodyFields(errorDV, errorFieldList)
		if err != nil {
			return "", err
		}
		return setResponseBody(errorStruct)
	}

	respDV := result.Output()
	switch respDV.Type() {
	case data.VOID:
		// Return null if result is void
		return "null", nil
	case data.STRUCTURE:
		// For structure type response, filter out fields with @Header annotation
		bodyStruct, err := filterBodyFields(respDV, bodyFieldList)
		if err != nil {
			return "", err
		}
		return setResponseBody(bodyStruct)
	default:
		// For non-structure response, we reply on DataValueToJsonEncoder
		// to serialize the value.
		return setResponseBody(respDV)
	}
}

func convertDVToHeaderValue(name string, respDV data.DataValue, header http.Header) (http.Header, []error) {
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
	case data.BLOB:
		valBlob := respDV.(*data.BlobValue)
		valueString := string(valBlob.Value())
		header.Add(name, valueString)
	case data.LIST:
		valList := respDV.(*data.ListValue)
		valSlice := valList.List()
		for _, elementDv := range valSlice {
			if elementDv.Type() == data.OPTIONAL {
				return nil,
					[]error{l10n.NewRuntimeError("vapi.protocol.server.rest.response.unsupport_type",
						map[string]string{"dataType": "LIST<OPTIONAL>", "fieldName": name})}
			}
			if header, err = convertDVToHeaderValue(name, elementDv, header); err != nil {
				return nil, err
			}
		}
	case data.OPTIONAL:
		optVal := respDV.(*data.OptionalValue)
		if !optVal.IsSet() {
			return header, nil
		}
		elemVal := optVal.Value()
		return convertDVToHeaderValue(name, elemVal, header)
	default:
		return nil,
			[]error{l10n.NewRuntimeError("vapi.protocol.server.rest.response.unsupport_type",
				map[string]string{"dataType": respDV.Type().String(), "fieldName": name})}
	}

	return header, nil
}

// Set the response value to http header
func setResponseHeader(respDV data.DataValue,
	resultToHeaderMap map[string]string) (http.Header, []error) {
	var header http.Header
	var respMap map[string]data.DataValue

	switch respDV.Type() {
	case data.ERROR:
		errorStruct, ok := respDV.(*data.ErrorValue)
		if !ok {
			return nil,
				[]error{l10n.NewRuntimeErrorNoParam("vapi.protocol.server.rest.response.error_not_structure")}
		}
		respMap = errorStruct.Fields()
	case data.STRUCTURE:
		respStruct, ok := respDV.(*data.StructValue)
		if !ok {
			return nil,
				[]error{l10n.NewRuntimeErrorNoParam("vapi.protocol.server.rest.response.not_structure")}
		}
		respMap = respStruct.Fields()
	default:
		return nil,
			[]error{l10n.NewRuntimeErrorNoParam("vapi.protocol.server.rest.response.not_structure")}
	}

	header = make(http.Header)

	var err []error
	var headerName string
	var exist      bool
	for name, val := range respMap {
		if headerName, exist = resultToHeaderMap[name]; !exist {
			continue
		}
		// There is Header annotation for this field
		valType := val.Type()
		switch valType {
		case data.STRUCTURE:
			return setResponseHeader(val, resultToHeaderMap)
		default:
			if header, err = convertDVToHeaderValue(headerName, val, header); err != nil {
				return nil, err
			}
		}
	}
	return header, nil
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