/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Job
 * Used by client-side stubs.
 */

package restore


// ``Job`` interface provides methods Performs restore operations
type JobClient interface {

    // Cancel the restore job
    // @return RestoreJobStatus Structure
    // @throws Error Generic error
	Cancel() (JobReturnResult, error)

    // Initiate restore.
    //
    // @param pieceParam RestoreRequest Structure
    // @return RestoreJobStatus Structure
    // @throws FeatureInUse A backup or restore is already in progress
    // @throws NotAllowedInCurrentState Restore is allowed only after deployment and before firstboot
    // @throws Error Generic error
	Create(pieceParam JobRestoreRequest) (JobRestoreJobStatus, error)

    // See restore job progress/result.
    // @return RestoreJobStatus Structure
    // @throws Error Generic error
	Get() (JobRestoreJobStatus, error)
}
