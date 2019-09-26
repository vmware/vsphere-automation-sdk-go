package bindings_test

/* **********************************************************
 * Copyright 2019 VMware, Inc. All rights reserved.
 *      -- VMware Confidential
 * *********************************************************
 */

import (
	"github.com/stretchr/testify/assert"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"reflect"
	"testing"
)

func TestMapStringString(t *testing.T) {
	bindingType := bindings.NewMapType(bindings.NewStringType(), bindings.NewStringType(), reflect.TypeOf(map[string]string{}))
	gomap := map[string]string{}
	gomap["a"] = "a"
	gomap["b"] = "b"
	typeconverter := bindings.NewTypeConverter()
	typeconverter.SetMode(bindings.REST)
	dataval, err := typeconverter.ConvertToVapi(gomap, bindingType)
	assert.Nil(t, err)
	dataValExpected := data.NewStructValue("map-struct", nil)
	dataValExpected.SetField("a", data.NewStringValue("a"))
	dataValExpected.SetField("b", data.NewStringValue("b"))
	assert.Equal(t, dataValExpected, dataval)
	govalOutput, err := typeconverter.ConvertToGolang(dataval, bindingType)
	assert.Nil(t, err)
	assert.Equal(t, govalOutput, gomap)
}

func TestMapIntegerString(t *testing.T) {
	bindingType := bindings.NewMapType(bindings.NewIntegerType(), bindings.NewStringType(), reflect.TypeOf(map[int64]string{}))
	gomap := map[int64]string{}
	gomap[1] = "a"
	gomap[2] = "b"
	typeconverter := bindings.NewTypeConverter()
	typeconverter.SetMode(bindings.REST)
	dataVal, err := typeconverter.ConvertToVapi(gomap, bindingType)
	dataValExpected := data.NewStructValue("map-struct", nil)
	dataValExpected.SetField("1", data.NewStringValue("a"))
	dataValExpected.SetField("2", data.NewStringValue("b"))
	govalOutput, err := typeconverter.ConvertToGolang(dataVal, bindingType)
	assert.Nil(t, err)
	assert.Equal(t, govalOutput, gomap)

}
