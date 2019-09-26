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

func TestProviderDefinition(t *testing.T) {
	pd := core.NewProviderDefinition("id1", "checkSum1")
	assert.Equal(t, pd.Identifier(), "id1")
	assert.Equal(t, pd.CheckSum(), "checkSum1")
}
