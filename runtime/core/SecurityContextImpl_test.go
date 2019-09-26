/* **********************************************************
 * Copyright 2018-2019 VMware, Inc. All rights reserved.
 *      -- VMware Confidential
 * ********************************************************* */
package core_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/core"
)

func TestSecurityContextSetAndGetProperty(t *testing.T) {
	securityContextImpl := core.NewSecurityContextImpl()
	securityContextImpl.SetProperty("interface1", "Interface1")
	assert.Equal(t, securityContextImpl.Property("interface1"), "Interface1")
}

func TestSecurityContextSetContextMap(t *testing.T) {
	securityContextImpl := core.NewSecurityContextImpl()
	contextData := make(map[string]interface{})
	contextData["interface1"] = "Interface1"
	contextData["Interface2"] = "Interface2"
	securityContextImpl.SetContextMap(contextData)
	assert.Equal(t, securityContextImpl.Property("interface1"), "Interface1")
	assert.Equal(t, securityContextImpl.Property("Interface2"), "Interface2")
}

func TestSecurityContextGetAllProperties(t *testing.T) {
	securityContextImpl := core.NewSecurityContextImpl()
	securityContextImpl.SetProperty("interface1", "Interface1")
	securityContextImpl.SetProperty("interface2", "Interface2")
	assert.Equal(t, securityContextImpl.GetAllProperties()["interface1"], "Interface1")
	assert.Equal(t, securityContextImpl.GetAllProperties()["interface2"], "Interface2")
}

func TestSecurityContextMarshalJSON(t *testing.T) {
	securityContextImpl := core.NewSecurityContextImpl()
	securityContextImpl.SetProperty("interface1", "Interface1")
	securityContextImpl.SetProperty("interface2", "Interface2")
	marshall1, _ := securityContextImpl.MarshalJSON()
	marshall2, _ := json.Marshal(securityContextImpl.GetAllProperties())
	assert.Equal(t, marshall1, marshall2)
}
