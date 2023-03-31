/* Copyright © 2021 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

package cleanjson

import (
	"encoding/json"
	"fmt"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data/serializers"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/l10n"
)

type JsonToDataValueDecoder struct {
	serializers.MethodResultDeserializerBase
}

func NewJsonToDataValueDecoder() *JsonToDataValueDecoder {
	i := &JsonToDataValueDecoder{}
	i.Impl = i
	return i
}

func (j *JsonToDataValueDecoder) GetDataValue(value interface{}) (data.DataValue, error) {
	return j.Decode(value)
}

func (j *JsonToDataValueDecoder) Decode(cleanJson interface{}) (data.DataValue, error) {

	switch result := cleanJson.(type) {
	case map[string]interface{}:
		return j.visitJsonDict(cleanJson.(map[string]interface{}))
	case []interface{}:
		return j.visitJsonList(cleanJson.([]interface{}))
	case string:
		return data.NewStringValue(cleanJson.(string)), nil
	case bool:
		return data.NewBooleanValue(cleanJson.(bool)), nil
	case json.Number:
		// By default, json marshaller converts json number to float64.
		// Use the below strategy to distinguish between floating point numbers and int64
		intValue, err := result.Int64()
		if err == nil {
			return data.NewIntegerValue(intValue), nil
		}
		floatValue, err := result.Float64()
		if err == nil {
			return data.NewDoubleValue(floatValue), nil
		}
		jsonNumberInvalid := l10n.NewRuntimeErrorNoParam("vapi.data.invalid.json.number")
		return nil, jsonNumberInvalid
	case nil:
		return data.NewOptionalValue(nil), nil
	default:
		unknownDataValue := l10n.NewRuntimeError("vapi.data.serializers.unsupported.json.type",
			map[string]string{"jsonType": fmt.Sprintf("%T", result)})
		return nil, unknownDataValue
	}
}

func (j *JsonToDataValueDecoder) visitJsonDict(inputData map[string]interface{}) (data.DataValue, error) {
	result := data.NewStructValue("", nil)
	for k, val := range inputData {
		dataVal, err := j.Decode(val)
		if err != nil {
			return nil, err
		}
		result.SetField(k, dataVal)
	}
	return result, nil
}

func (j *JsonToDataValueDecoder) visitJsonList(i []interface{}) (data.DataValue, error) {
	result := data.NewListValue()
	for _, val := range i {
		dataVal, err := j.Decode(val)
		if err != nil {
			return nil, err
		}
		result.Add(dataVal)
	}
	return result, nil
}
