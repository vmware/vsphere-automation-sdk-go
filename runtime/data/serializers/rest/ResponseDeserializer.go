package rest

/* **********************************************************
 * Copyright 2019 VMware, Inc.  All rights reserved.
 *      -- VMware Confidential
 * **********************************************************/

import (
	"encoding/json"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"strings"

	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/core"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data/serializers/cleanjson"
	httpStatus "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/lib/rest"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/log"
)

func isStatusSuccess(status int) bool {
	for _, val := range httpStatus.SUCCESSFUL_STATUS_CODES {
		if status == val {
			return true
		}
	}
	return false
}

func DeSerializeResponseToDataValue(response string) (data.DataValue, error) {
	if response == "" {
		return data.NewVoidValue(), nil
	}
	d := json.NewDecoder(strings.NewReader(response))
	d.UseNumber()
	var jsondata interface{}
	if err := d.Decode(&jsondata); err != nil {
		log.Error(err)
		return nil, err
	}
	jsonToDataValueDecoder := cleanjson.NewJsonToDataValueDecoder()
	responseValue, err := jsonToDataValueDecoder.Decode(jsondata)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return responseValue, nil
}

func DeserializeResponse(status int, response string) (core.MethodResult, error) {
	dataVal, err := DeSerializeResponseToDataValue(response)
	if err != nil {
		return core.MethodResult{}, err
	}
	if isStatusSuccess(status) {
		return core.NewMethodResult(dataVal, nil), nil
	}
	//create errorval
	var errorVal *data.ErrorValue
	vapiStdErrorName, ok := httpStatus.HTTP_TO_VAPI_ERROR_MAP[status]
	if !ok {
		// Use "com.vmware.vapi.std.errors.internal_server_error" if response error code is not present
		vapiStdErrorName = httpStatus.HTTP_TO_VAPI_ERROR_MAP[500]
	}
	errorVal = bindings.CreateErrorValueFromMessages(bindings.ERROR_MAP[vapiStdErrorName], []error{})
	errorVal.SetField("messages", data.NewListValue())
	errorVal.SetField("data", dataVal)
	return core.NewMethodResult(nil, errorVal), nil
}
