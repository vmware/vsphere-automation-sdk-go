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

// The ``Job`` interface provides methods to create and get the status of reconciliation job.
type JobClient interface {


    // Initiate reconciliation.
    //
    // @param specParam CreateSpec Structure
    // @return Info Structure
    // @throws FeatureInUse A backup or restore is already in progress.
    // @throws NotAllowedInCurrentState Reconciliation is allowed only after restore has finished successfully.
    // @throws Error if any error occurs during the execution of the operation.
    Create(specParam CreateSpec) (Info, error) 


    // Get reconciliation job progress/result.
    // @return Info Structure
    // @throws NotFound if there is no running reconciliation job.
    // @throws Error if any error occurs during the execution of the operation.
    Get() (Info, error) 

}
