/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

package basic

// Error represents the error for Basic Authentication.
type Error struct {
	ErrorDetails error
}

func (basicAuthError *Error) Error() string {
	return Name + " Error: UserName and Password are missing."
}

// UsernameError represents Username Error for Basic Authentication.
type UsernameError struct{}

func (userNameError *UsernameError) Error() string {
	return Name + " Error: Username is not available."
}

// PasswordError represents Password Error for Basic Authentication.
type PasswordError struct{}

func (passwordError *PasswordError) Error() string {
	return Name + " Error: Password is not available."
}
