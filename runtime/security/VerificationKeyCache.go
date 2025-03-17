// Copyright (c) 2022-2024 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

package security

// VerificationKeyCache which caches keys and needs to be refreshed if the signing keys change.
type VerificationKeyCache interface {
	VerificationKeyProvider

	// Refresh of the provider (if the underlying implementation maintains a cache)
	// If the refresh is successful, all subsequent calls to Get must
	// return the refreshed keys.
	Refresh(issuer string) error
}
