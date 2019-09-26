/* **********************************************************
 * Copyright 2018-2019 VMware, Inc. All rights reserved.
 *      -- VMware Confidential
 * ********************************************************* */
package core_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/core"
)

func TestFullyQualifiedName(t *testing.T) {
	var iface = core.NewInterfaceIdentifier("x.y.z")
	var methodName = "name"
	var mi = core.NewMethodIdentifier(iface, methodName)
	var result = "x.y.z.name"
	assert.Equal(t, result, mi.FullyQualifiedName())
}

func TestMethodIdentifier(t *testing.T) {
	var interfaceName = "x.y.z"
	var iface = core.NewInterfaceIdentifier(interfaceName)
	var methodName = "name"
	var mi1 = core.NewMethodIdentifier(iface, methodName)
	var methodName1 = "name1"
	var mi2 = core.NewMethodIdentifier(iface, methodName1)
	assert.NotEqual(t, mi1, mi2)
}

func TestInterfaceIdentifier(t *testing.T) {
	var interfaceName = "x.y.z"
	var iface = core.NewInterfaceIdentifier(interfaceName)
	var methodName = "name"
	var mi = core.NewMethodIdentifier(iface, methodName)
	assert.Equal(t, mi.InterfaceIdentifier(), iface)
}
