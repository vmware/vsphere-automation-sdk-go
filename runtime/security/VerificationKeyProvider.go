// Copyright (c) 2022-2024 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

package security

// VerificationKeyProvider provides keys used to validate the authenticity of the JWT token.
type VerificationKeyProvider interface {

	// Retrieves keys from issuer.
	// returns
	//  ([]interface{}, nil) keys utilized for JWT verification
	//  (nil, error) for invalid retrieval
	Get(issuer string) ([]interface{}, error)
}
