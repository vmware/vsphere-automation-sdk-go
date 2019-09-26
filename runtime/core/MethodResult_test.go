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

func TestMethodResultOutput(t *testing.T) {
	var err *data.ErrorValue = nil
	outputvalue := data.NewStringValue("hello")
	methodResult := core.NewMethodResult(outputvalue, err)
	assert.Equal(t, methodResult.Output(), outputvalue)
	assert.Equal(t, methodResult.Error(), err)
	assert.True(t, methodResult.IsSuccess())
}

func TestMethodResultError(t *testing.T) {
	var output data.DataValue = nil
	errorvalue := data.NewErrorValue("error1", nil)
	methodResult := core.NewMethodResult(output, errorvalue)
	assert.Equal(t, methodResult.Error(), errorvalue)
	assert.Equal(t, methodResult.Output(), output)
	assert.False(t, methodResult.IsSuccess())
}
