// Copyright (c) 2019-2024 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

package info

type SessionType int

const (
	SESSION_LESS  SessionType = 1 + iota // Depicts if a service is session less
	SESSION_AWARE                        // Depicts if a service is session full
)

func (s SessionType) String() string {
	switch s {
	case SESSION_LESS:
		return "SESSION_LESS"
	case SESSION_AWARE:
		return "SESSION_AWARE"
	}
	return ""
}
