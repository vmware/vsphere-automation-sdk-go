/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Restore
 * Used by client-side stubs.
 */

package recovery


// ``Restore`` interface provides methods Performs restore operations
type RestoreClient interface {

    // Get metadata before restore
    //
    // @param pieceParam RestoreRequest Structure
    // @return Metadata Structure
    // @throws Error Generic error
	Validate(pieceParam RestoreRestoreRequest) (RestoreMetadata, error)
}
