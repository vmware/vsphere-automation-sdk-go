/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

package arch

type archType string

func (at archType) isType() archType {
	return at
}

func (at archType) String() string {
	return string(at)
}
