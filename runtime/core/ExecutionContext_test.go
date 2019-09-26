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

type SecurityContextImpl2 struct {
	contextData map[string]interface{}
}

func NewSecurityContextImpl2() *SecurityContextImpl2 {
	return &SecurityContextImpl2{contextData: make(map[string]interface{})}
}

func (s *SecurityContextImpl2) Property(key string) interface{} {
	return s.contextData[key]
}

func (s *SecurityContextImpl2) SetProperty(key string, value interface{}) {
	s.contextData[key] = value
}

func (s *SecurityContextImpl2) GetAllProperties() map[string]interface{} {
	return s.contextData
}

func TestSecurityContext(t *testing.T) {
	securityContextImpl2 := NewSecurityContextImpl2()
	app := make(map[string]*string)
	appContext := core.NewApplicationContext(app)
	executionContext := core.NewExecutionContext(appContext, securityContextImpl2)
	executionContext.SetSecurityContext(securityContextImpl2)
	assert.Equal(t, executionContext.SecurityContext(), securityContextImpl2)
}

func TestAppIsNew2(t *testing.T) {
	securityContextImpl2 := NewSecurityContextImpl2()
	executionContext := core.NewExecutionContext(nil, securityContextImpl2)
	assert.Equal(t, executionContext.SecurityContext(), securityContextImpl2)
}

func TestApplicationContext(t *testing.T) {
	securityContextImpl2 := NewSecurityContextImpl2()
	app := make(map[string]*string)
	appContext := core.NewApplicationContext(app)
	executionContext := core.NewExecutionContext(appContext, securityContextImpl2)
	assert.Equal(t, executionContext.ApplicationContext(), appContext)
}
