/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Package arch defines protocol for API calls.
package arch

import (
	"strings"
)

var prots = []Type{
	JSONRPC,
	REST,
}

// GetByName fetches protocol by protocol name.
func GetByName(name string) Type {
	for _, p := range prots {
		if strings.ToLower(p.String()) == strings.ToLower(name) {
			return p
		}
	}
	return nil
}
