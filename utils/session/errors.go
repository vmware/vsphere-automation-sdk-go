/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

package session

// ArchTypeError represents the error for Architecture-Type/Protocol.
type ArchTypeError struct {
	archType string
}

func (err *ArchTypeError) Error() string {
	return "Session Manager Error: Invalid Session Architecture Type " + err.archType
}

// URLError represents the error for URL Error.
type URLError struct{}

func (err *URLError) Error() string {
	return "Session Manager Error: Invalid URL"
}

// AuthDetailsError represents the error for Authentication Details Error.
type AuthDetailsError struct{}

func (err *AuthDetailsError) Error() string {
	return "Session Manager Error: Invalid Authentication Details"
}
