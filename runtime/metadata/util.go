// Copyright (c) 2019-2024 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

package metadata

import "github.com/vmware/vsphere-automation-sdk-go/runtime/l10n"

func getError(msg string) error {
	args := map[string]string{
		"msg": msg,
	}
	return l10n.NewRuntimeError("vapi.metadata.parser.failure", args)
}
