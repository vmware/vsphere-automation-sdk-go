// Copyright (c) 2021-2024 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

package core

// TaskResultProvider provides tasks result outcome
// Used in bindings
type TaskResultProvider interface {
	GetTaskResult(taskId string) (MonoResult, error)
}
