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

func TestInterfaceDefinitionIdentifier(t *testing.T) {
	interfaceIdentifier := core.NewInterfaceIdentifier("id")
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, nil)
	assert.Equal(t, interfaceDefinition.Identifier(), interfaceIdentifier)
}

func TestInterfaceDefinitionMethodIdentifiers(t *testing.T) {
	interfaceIdentifier := core.NewInterfaceIdentifier("id")
	methodIdentifier1 := core.NewMethodIdentifier(interfaceIdentifier, "methodName1")
	methodIdentifier2 := core.NewMethodIdentifier(interfaceIdentifier, "methodName2")
	methodIdentifiers := []core.MethodIdentifier{methodIdentifier1, methodIdentifier2}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	assert.Equal(t, interfaceDefinition.MethodIdentifiers()[0], methodIdentifier1)
	assert.Equal(t, interfaceDefinition.MethodIdentifiers()[1], methodIdentifier2)
}

func TestInterfaceDefinitionEquals(t *testing.T) {
	interfaceIdentifier := core.NewInterfaceIdentifier("id")
	interfaceIdentifier2 := core.NewInterfaceIdentifier("id2")
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, nil)
	interfaceDefinition2 := core.NewInterfaceDefinition(interfaceIdentifier2, nil)
	interfaceDefinition3 := core.NewInterfaceDefinition(interfaceIdentifier, nil)
	assert.False(t, interfaceDefinition.Equals(interfaceDefinition2))
	assert.True(t, interfaceDefinition.Equals(interfaceDefinition3))
}
