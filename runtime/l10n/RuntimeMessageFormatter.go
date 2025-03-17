// Copyright (c) 2019-2024 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

package l10n

import (
	"bytes"
	"sync"

	"github.com/vmware/vsphere-automation-sdk-go/runtime/l10n/runtime"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/log"
)

var (
	runtimeMessageFormatter     *MessageFormatter
	runtimeMessageFormatterOnce sync.Once
)

// Error formatter with default localization parameters
// backed by message bundle for the runtime
func DefaultMessageFormatter() *MessageFormatter {
	runtimeMessageFormatterOnce.Do(func() {
		runtimeFactory := NewMessageFactory()
		err := runtimeFactory.AddBundle("en", bytes.NewReader(runtime.RuntimeProperties_EN))
		if err != nil {
			log.Error("Error creating runtime message factory")
		}
		formatter := runtimeFactory.GetDefaultFormatter()
		runtimeMessageFormatter = formatter
	})
	return runtimeMessageFormatter
}
