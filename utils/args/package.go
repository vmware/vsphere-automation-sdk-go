/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

package args

// InitProperties instantiates and returns Properties.
func InitProperties() Properties {
	return &properties{OptionsMap: make(map[string]*propValue)}
}
