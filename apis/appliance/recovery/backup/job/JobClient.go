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

// The ``Job`` interface provides methods to be performed on a backup job.
type JobClient interface {


    // Cancel the backup job.
    //
    // @param idParam ID (ID of job)
    // The parameter must be an identifier for the resource type: ``com.vmware.appliance.recovery.backup.job``.
    // @return BackupJobStatus Structure
    // @throws NotFound if backup associated with id does not exist.
    // @throws Error if any error occurs during the execution of the operation.
    Cancel(idParam string) (ReturnResult, error) 


    // Initiate backup.
    //
    // @param pieceParam BackupRequest Structure
    // @return BackupJobStatus Structure
    // @throws FeatureInUse A backup or restore is already in progress.
    // @throws Error if any error occurs during the execution of the operation.
    Create(pieceParam BackupRequest) (BackupJobStatus, error) 


    // Get list of backup jobs
    // @return list of BackupJob IDs
    // The return value will contain identifiers for the resource type: ``com.vmware.appliance.recovery.backup.job``.
    // @throws Error if any error occurs during the execution of the operation.
    List() ([]string, error) 


    // See backup job progress/result.
    //
    // @param idParam ID (ID of job)
    // The parameter must be an identifier for the resource type: ``com.vmware.appliance.recovery.backup.job``.
    // @return BackupJobStatus Structure
    // @throws NotFound if backup associated with id does not exist.
    // @throws Error if any error occurs during the execution of the operation.
    Get(idParam string) (BackupJobStatus, error) 

}
