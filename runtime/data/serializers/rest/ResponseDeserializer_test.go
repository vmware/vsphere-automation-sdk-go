package rest_test

/* **********************************************************
 * Copyright 2019 VMware, Inc.  All rights reserved.
 *      -- VMware Confidential
 * **********************************************************/

import (
	"github.com/stretchr/testify/assert"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/core"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data/serializers/rest"
	"testing"
)

func TestStringValue(t *testing.T) {
	status := 200
	response := "{\"value\":\"vm-100\"}"

	fields := map[string]data.DataValue{
		"value": data.NewStringValue("vm-100"),
	}
	expected := core.NewMethodResult(data.NewStructValue("", fields), nil)
	actual, _ := rest.DeserializeResponse(status, response)

	assert.Equal(t, expected, actual)
}

func TestStructValue(t *testing.T) {
	status := 200
	response := "{\"value\":{\"id\":\"vm-100\", \"name\":\"Linux VM\"}}"

	subfields := map[string]data.DataValue{
		"id":   data.NewStringValue("vm-100"),
		"name": data.NewStringValue("Linux VM"),
	}
	fields := map[string]data.DataValue{
		"value": data.NewStructValue("", subfields),
	}
	expected := core.NewMethodResult(data.NewStructValue("", fields), nil)
	actual, _ := rest.DeserializeResponse(status, response)

	assert.Equal(t, expected, actual)
}

func TestOptionalValue(t *testing.T) {
	status := 200
	response := "{\"value\":null}"

	fields := map[string]data.DataValue{
		"value": data.NewOptionalValue(nil),
	}
	expected := core.NewMethodResult(data.NewStructValue("", fields), nil)
	actual, _ := rest.DeserializeResponse(status, response)

	assert.Equal(t, expected, actual)
}

func TestVoidValue(t *testing.T) {
	status := 200
	response := ""

	expected := core.NewMethodResult(data.NewVoidValue(), nil)
	actual, _ := rest.DeserializeResponse(status, response)

	assert.Equal(t, expected, actual)
}

func TestNotFoundErrorValue(t *testing.T) {
	status := 404
	response := "{\"type\":\"com.vmware.vapi.std.errors.not_found\",\"value\":{\"messages\":[{\"args\":[\"datacenter-21\"],\"default_message\":\"Datacenter with identifier 'datacenter-21' does not exist.\",\"id\":\"com.vmware.api.vcenter.datacenter.not_found\"}]}}"

	argsList := data.NewListValue()
	msgList := data.NewListValue()

	argsList.Add(data.NewStringValue("datacenter-21"))
	msgfields := map[string]data.DataValue{
		"id":              data.NewStringValue("com.vmware.api.vcenter.datacenter.not_found"),
		"default_message": data.NewStringValue("Datacenter with identifier 'datacenter-21' does not exist."),
		"args":            argsList,
	}
	msgList.Add(data.NewStructValue("", msgfields))
	valfields := map[string]data.DataValue{
		"messages": msgList,
	}

	fields := map[string]data.DataValue{
		"type":  data.NewStringValue("com.vmware.vapi.std.errors.not_found"),
		"value": data.NewStructValue("", valfields),
	}

	fieldsNew := map[string]data.DataValue{}
	fieldsNew["messages"] = data.NewListValue()
	fieldsNew["data"] = data.NewStructValue("", fields)
	fieldsNew["error_type"] = data.NewOptionalValue(data.NewStringValue("NOT_FOUND"))
	expected := core.NewMethodResult(data.DataValue(nil), data.NewErrorValue("com.vmware.vapi.std.errors.not_found", fieldsNew))
	actual, _ := rest.DeserializeResponse(status, response)

	assert.Equal(t, expected, actual)
}

func TestSuccessfulStatus202(t *testing.T) {
	status := 202
	response := "{\"value\":\"vm-100\"}"

	fields := map[string]data.DataValue{
		"value": data.NewStringValue("vm-100"),
	}
	expected := core.NewMethodResult(data.NewStructValue("", fields), nil)
	actual, _ := rest.DeserializeResponse(status, response)

	assert.Equal(t, expected, actual)
}
