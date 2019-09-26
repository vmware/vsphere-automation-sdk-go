/* **********************************************************
 * Copyright 2018-2019 VMware, Inc. All rights reserved.
 *      -- VMware Confidential
 * ********************************************************* */
package core_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/core"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
)

func TestMethodDefinitionIdentifier(t *testing.T) {
	var iface = core.NewInterfaceIdentifier("x.y.z")
	var mi = core.NewMethodIdentifier(iface, "method1")
	var input = data.NewStringDefinition()
	var output = data.NewStringDefinition()
	var ddMap = map[string]data.DataDefinition{"data1": input}
	var ed = data.NewErrorDefinition("error1", ddMap)
	var errorDefinitionList = []data.ErrorDefinition{ed}
	newMethodDefinition := core.NewMethodDefinition(mi, input, output, errorDefinitionList)
	assert.Equal(t, newMethodDefinition.Identifier(), mi)
	assert.Equal(t, newMethodDefinition.InputDefinition(), input)
	assert.Equal(t, newMethodDefinition.OutputDefinition(), output)
	assert.Equal(t, newMethodDefinition.ErrorDefinitions()[0].Name(), "error1")
	assert.Equal(t, newMethodDefinition.ErrorDefinition("error1").Name(), "error1")
	assert.True(t, newMethodDefinition.HasErrorDefinition("error1"))
	assert.False(t, newMethodDefinition.HasErrorDefinition("error2"))
}
