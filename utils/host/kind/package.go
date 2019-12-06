/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Package kind represents kind of Hosts (VMware Products).
package kind

import (
	"strings"
)

// GetHostTypeByName fetches Host Type Info by Host Type name.
func GetHostTypeByName(name string) Info {
	for _, host := range All {
		if strings.ToLower(host.String()) == strings.ToLower(name) {
			return host
		}
	}
	return nil
}

// GetAllHostKind gives array of names of all Host Types.
func GetAllHostKind() []string {
	var allHostKind []string
	for _, hostKind := range All {
		allHostKind = append(allHostKind, hostKind.String())
	}
	return allHostKind
}
