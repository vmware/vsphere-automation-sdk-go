// Copyright (c) 2021-2024 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

package serializers

import (
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib"
)

type MethodResultDeserializer interface {
	DeserializeMethodResult(map[string]interface{}) (core.MethodResult, error)
	GetDataValue(interface{}) (data.DataValue, error)
}

type MethodResultDeserializerBase struct {
	Impl MethodResultDeserializer
}

func (m MethodResultDeserializerBase) DeserializeMethodResult(
	methodResultInput map[string]interface{}) (core.MethodResult, error) {

	if val, ok := methodResultInput[lib.METHOD_RESULT_OUTPUT]; ok {
		var output, err = m.Impl.GetDataValue(val)
		if err != nil {
			return nil, err
		}
		return core.NewMethodResult(output, nil), nil
	} else if val, ok := methodResultInput[lib.METHOD_RESULT_ERROR]; ok {
		var methodResultError, err = m.Impl.GetDataValue(val)
		if err != nil {
			return nil, err
		}
		return core.NewMethodResult(nil, methodResultError.(*data.ErrorValue)), nil
	}

	return nil, core.DeserializationError
}
