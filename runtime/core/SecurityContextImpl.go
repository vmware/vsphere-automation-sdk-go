// Copyright (c) 2019-2024 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

package core

import (
	"encoding/json"
)

type SecurityContextImpl struct {
	contextData map[string]interface{}
}

func NewSecurityContextImpl() *SecurityContextImpl {
	return &SecurityContextImpl{contextData: make(map[string]interface{})}
}

func (s *SecurityContextImpl) Property(key string) interface{} {
	return s.contextData[key]
}

func (s *SecurityContextImpl) SetProperty(key string, value interface{}) {
	s.contextData[key] = value
}

func (s *SecurityContextImpl) SetContextMap(context map[string]interface{}) {
	s.contextData = context
}

func (s *SecurityContextImpl) GetAllProperties() map[string]interface{} {
	return s.contextData
}

func (s *SecurityContextImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.contextData)
}
