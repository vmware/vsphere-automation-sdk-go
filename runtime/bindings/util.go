// Copyright (c) 2019-2024 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

package bindings

import (
	"errors"
	"reflect"
)

func isPointer(x interface{}) bool {
	return reflect.ValueOf(x).Kind() == reflect.Ptr
}

func VAPIerrorsToError(errs []error) error {
	message := ""
	for _, e := range errs {
		message = message + " " + e.Error()
	}
	return errors.New(message)
}
