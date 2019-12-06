/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

package arch

// Type represents the protocol/architecture type for API calls.
type Type interface {
	isType() archType
	// String returns the string value for the Architecture type.
	String() string
}
