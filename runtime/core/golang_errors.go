// Copyright (c) 2021-2024 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

package core

import "errors"

var DeserializationError = errors.New("error de-serializing method result")

var UnacceptableContent = errors.New("response content type doesn't match any of the acceptable types")
