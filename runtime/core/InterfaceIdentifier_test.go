/* **********************************************************
 * Copyright 2019 VMware, Inc. All rights reserved.
 *      -- VMware Confidential
 * *********************************************************
 */
package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInterfaceEquals(t *testing.T) {
	interfaceIdentifier1 := NewInterfaceIdentifier("name1")
	interfaceIdentifier2 := NewInterfaceIdentifier("name1")
	assert.Equal(t, interfaceIdentifier1, interfaceIdentifier2)
}

func TestInterfaceNotEquals(t *testing.T) {
	interfaceIdentifier1 := NewInterfaceIdentifier("name1")
	interfaceIdentifier2 := NewInterfaceIdentifier("name2")
	assert.NotEqual(t, interfaceIdentifier1, interfaceIdentifier2)
}

func TestInterfaceName(t *testing.T) {
	interfaceIdentifier1 := NewInterfaceIdentifier("name1")
	assert.Equal(t, interfaceIdentifier1.Name(), "name1")
}

func TestInterfaceString(t *testing.T) {
	interfaceIdentifier1 := NewInterfaceIdentifier("name1")
	assert.Equal(t, interfaceIdentifier1.String(), "name1")
}
