/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Backup
 * Used by client-side stubs.
 */

package recovery


// ``Backup`` interface provides methods Performs backup restore operations
type BackupClient interface {

    // Check for backup errors without starting backup.
    //
    // @param pieceParam BackupRequest Structure
    // @return ReturnResult Structure
    // @throws FeatureInUse A backup or restore is already in progress
    // @throws Error Generic error
	Validate(pieceParam BackupBackupRequest) (BackupReturnResult, error)
}
