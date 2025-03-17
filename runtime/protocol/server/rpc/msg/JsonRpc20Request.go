// Copyright (c) 2019-2024 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

package msg

import (
	"errors"
	"reflect"

	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib"
)

type JsonRpc20Request struct {
	//A Structured value that holds the parameter values to be used during the invocation of the method.
	params map[string]interface{}

	id interface{}
}

func NewJsonRpc20Request(params map[string]interface{}, id interface{}) JsonRpc20Request {
	//TODO:sreeshas
	//handle errors
	return JsonRpc20Request{params: params, id: id}
}

func (j JsonRpc20Request) JSON() map[string]interface{} {
	var result = make(map[string]interface{})
	result[lib.JSONRPC] = lib.JSONRPC_VERSION
	result[lib.JSONRPC_METHOD] = lib.JSONRPC_INVOKE
	result[lib.JSONRPC_PARAMS] = j.params
	result[lib.JSONRPC_ID] = j.id
	return result
}

func (j JsonRpc20Request) Id() interface{} {
	return j.id
}

func (j JsonRpc20Request) Params() map[string]interface{} {
	return j.params
}

func getJsonRpcMethod(request map[string]interface{}) (string, error) {
	methodField, ok := request[lib.JSONRPC_METHOD]
	if !ok {
		return "", errors.New("field 'method' is required")
	}
	method, ok := methodField.(string)
	if !ok {
		return "", errors.New("field 'method' must be a string")
	}
	if method != lib.JSONRPC_INVOKE {
		return "", errors.New("unsupported JSON-RPC method: " + method)
	}
	return method, nil
}

// DeSerializeRequest gets JsonRpc20Request object from provided string or byte array json
func DeSerializeRequest(request interface{}) (JsonRpc20Request, *JsonRpc20Error) {
	requestObject, err := DeSerializeJson(request)
	if err != nil {
		return JsonRpc20Request{}, NewJsonRpcErrorParseError(err)
	}
	return getJsonRpc20Request(requestObject)
}

// getJsonRpc20Request gets JsonRpc20Request from unmarshalled json map
func getJsonRpc20Request(request map[string]interface{}) (JsonRpc20Request, *JsonRpc20Error) {
	_, err := getJsonRpcVersion(request)
	if err != nil {
		return JsonRpc20Request{}, NewJsonRpcErrorInvalidRequest(err.Error())
	}
	id, err := getJsonRPCIdValue(request)
	if err != nil {
		return JsonRpc20Request{}, NewJsonRpcErrorInvalidRequest(err.Error())
	}
	if id == nil {
		return JsonRpc20Request{}, NewJsonRpcErrorInvalidRequest("field 'id' is required")
	}
	_, err = getJsonRpcMethod(request)
	if err != nil {
		return JsonRpc20Request{}, NewJsonRpcErrorInvalidRequest(err.Error())
	}

	if paramValue, ok := request[lib.JSONRPC_PARAMS]; ok {
		paramMap, isObject := paramValue.(map[string]interface{})
		if !isObject && reflect.TypeOf(paramValue) != nil {
			return JsonRpc20Request{}, NewJsonRpcErrorInvalidRequest("field 'params' must be a JSON object")
		}
		return NewJsonRpc20Request(paramMap, id), nil
	} else {
		return NewJsonRpc20Request(nil, id), nil
	}
}
