// Copyright (c) 2019-2024 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

package info

type AuthenticationScheme struct {
	interfaceName  string
	sessionManager string
	sessionType    SessionType
}

// Interface Name
func (auth *AuthenticationScheme) InterfaceName() string {
	return auth.interfaceName
}

func (auth *AuthenticationScheme) SetInterfaceName(interfaceName string) {
	auth.interfaceName = interfaceName
}

// Session Type
func (auth *AuthenticationScheme) SessionType() SessionType {
	return auth.sessionType
}

func (auth *AuthenticationScheme) SetSessionType(sessionType SessionType) {
	auth.sessionType = sessionType
}

// Session Manager
func (auth *AuthenticationScheme) SessionManager() string {
	return auth.sessionManager
}

func (auth *AuthenticationScheme) SetSessionManager(sessionManager string) {
	auth.sessionManager = sessionManager
}
