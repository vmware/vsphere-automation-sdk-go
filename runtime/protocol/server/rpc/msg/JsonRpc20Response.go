// Copyright (c) 2019-2024 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

package msg

import (
	"encoding/json"
	"errors"

	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib"
)

type JsonRpc20Response struct {
	/**
	 *  This member is REQUIRED.
	 *	It MUST be the same as the value of the id member in the Request Object.
	 *	If there was an error in detecting the id in the Request object (e.g. Parse error/Invalid Request), it MUST be Null.
	 */
	id interface{}

	/**
	 *	This member is REQUIRED on success.
	 *	This member MUST NOT exist if there was an error invoking the method.
	 *	The value of this member is determined by the method invoked on the Server.
	 */
	result interface{}

	/**
	 *	This member is REQUIRED on error.
	 *	This member MUST NOT exist if there was no error triggered during invocation.
	 *	The value for this member MUST be an Object as defined in section 5.1.
	 */
	error map[string]interface{}
}

func NewJsonRpc20Response(id interface{}, result interface{}, error map[string]interface{}) JsonRpc20Response {
	return JsonRpc20Response{id: id, result: result, error: error}
}

func (j JsonRpc20Response) Id() interface{} {
	return j.id
}

func (j JsonRpc20Response) Result() interface{} {
	return j.result
}

func (j JsonRpc20Response) Error() map[string]interface{} {
	return j.error
}

func getJsonRpcVersion(msg map[string]interface{}) (string, error) {
	versionField, ok := msg[lib.JSONRPC]
	if !ok {
		return "", errors.New("field 'jsonrpc' is required")
	}
	version, ok := versionField.(string)
	if !ok {
		return "", errors.New("field 'jsonrpc' must be a string")
	}
	if version != lib.JSONRPC_VERSION {
		return "", errors.New("unsupported 'jsonrpc' version: " + version)
	}
	return version, nil
}

// getJsonRPCIdValue extracts id value from request or response
// The id can be string or int or nil.
func getJsonRPCIdValue(msg map[string]interface{}) (interface{}, error) {
	idField, ok := msg[lib.JSONRPC_ID]
	if !ok || idField == nil {
		return nil, nil
	}
	idString, isString := idField.(string)
	if isString {
		return idString, nil
	}
	jsonNumber, isNumber := idField.(json.Number)
	if !isNumber {
		return nil, errors.New("field 'id' must be string or int")
	}
	idInt, intErr := jsonNumber.Int64()
	if intErr != nil {
		return nil, errors.New("field 'id' must be string or int")
	}
	return idInt, nil
}

// DeSerializeResponse gets JsonRpc20Response object from provided string or byte array json
func DeSerializeResponse(response interface{}) (JsonRpc20Response, *JsonRpc20Error) {
	responseObj, err := DeSerializeJson(response)
	if err != nil {
		return JsonRpc20Response{}, NewJsonRpcErrorParseError(err)
	}
	return GetJsonRpc20Response(responseObj)
}

// GetJsonRpc20Response gets JsonRpc20Response object from unmarshalled json map
func GetJsonRpc20Response(response map[string]interface{}) (JsonRpc20Response, *JsonRpc20Error) {
	_, err := getJsonRpcVersion(response)
	if err != nil {
		return JsonRpc20Response{}, NewJsonRpcErrorInvalidParams(err.Error())
	}
	id, err := getJsonRPCIdValue(response)
	if err != nil {
		return JsonRpc20Response{}, NewJsonRpcErrorInvalidParams(err.Error())
	}
	if result, ok := response[lib.METHOD_RESULT]; ok {
		return NewJsonRpc20Response(id, result, nil), nil
	} else if err, ok := response[lib.METHOD_RESULT_ERROR]; ok {
		errMap, ok := err.(map[string]interface{})
		if !ok {
			return JsonRpc20Response{}, NewJsonRpcErrorInvalidParams("field 'error' must be a JSON object")
		}
		return NewJsonRpc20Response(id, nil, errMap), nil
	} else {
		return JsonRpc20Response{}, NewJsonRpcErrorInvalidParams("response must contain either 'result' or 'error' field")
	}
}
