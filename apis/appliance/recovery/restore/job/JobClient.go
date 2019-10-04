/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Job
 * Used by client-side stubs.
 */

package job

import (
)

// ``Job`` interface provides methods Performs restore operations
type JobClient interface {


    // Cancel the restore job
    // @return RestoreJobStatus Structure
    // @throws Error Generic error
    Cancel() (ReturnResult, error) 


    // Initiate restore.
    //
    // @param pieceParam RestoreRequest Structure
    // @return RestoreJobStatus Structure
    // @throws FeatureInUse A backup or restore is already in progress
    // @throws NotAllowedInCurrentState Restore is allowed only after deployment and before firstboot
    // @throws Error Generic error
    Create(pieceParam RestoreRequest) (RestoreJobStatus, error) 


    // See restore job progress/result.
    // @return RestoreJobStatus Structure
    // @throws Error Generic error
    Get() (RestoreJobStatus, error) 

}
