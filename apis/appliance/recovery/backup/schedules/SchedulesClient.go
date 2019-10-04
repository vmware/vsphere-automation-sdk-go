/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Schedules
 * Used by client-side stubs.
 */

package schedules

import (
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/apis/appliance/recovery/backup/job"
)

// The ``Schedules`` interface provides methods to be performed to manage backup schedules.
type SchedulesClient interface {


    // Returns a list of existing schedules with details
    // @return Map of schedule id to Info Structure
    // The key in the return value map will be an identifier for the resource type: ``com.vmware.appliance.recovery.backup.schedule``.
    // @throws Error if any error occurs during the execution of the operation.
    List() (map[string]Info, error) 


    // Initiate backup with the specified schedule
    //
    // @param scheduleParam Identifier of the schedule
    // The parameter must be an identifier for the resource type: ``com.vmware.appliance.recovery.backup.schedule``.
    // @param commentParam field that specifies the description for the backup.
    // If null the backup will have an empty comment.
    // @return BackupJobStatus Structure
    // @throws FeatureInUse if a backup or restore is already in progress.
    // @throws NotFound if schedule associated with id does not exist.
    // @throws Error if any error occurs during the execution of the operation.
    Run(scheduleParam string, commentParam *string) (job.BackupJobStatus, error) 


    // Returns an existing schedule information based on id
    //
    // @param scheduleParam Identifier of the schedule
    // The parameter must be an identifier for the resource type: ``com.vmware.appliance.recovery.backup.schedule``.
    // @return Info Structure
    // @throws NotFound if schedule associated with id does not exist.
    // @throws Error if any error occurs during the execution of the operation.
    Get(scheduleParam string) (Info, error) 


    // Creates a schedule
    //
    // @param scheduleParam Identifier of the schedule
    // The parameter must be an identifier for the resource type: ``com.vmware.appliance.recovery.backup.schedule``.
    // @param specParam CreateSpec Structure
    // @throws InvalidArgument if provided with invalid schedule specification.
    // @throws AlreadyExists if the schedule with the given id already exists.
    // @throws Error if any error occurs during the execution of the operation.
    Create(scheduleParam string, specParam CreateSpec) error 


    // Updates a schedule
    //
    // @param scheduleParam Identifier of the schedule
    // The parameter must be an identifier for the resource type: ``com.vmware.appliance.recovery.backup.schedule``.
    // @param specParam UpdateSpec Structure
    // @throws InvalidArgument if provided with invalid schedule specification.
    // @throws NotFound if schedule associated with id does not exist.
    // @throws Error if any error occurs during the execution of the operation.
    Update(scheduleParam string, specParam UpdateSpec) error 


    // Deletes an existing schedule
    //
    // @param scheduleParam Identifier of the schedule
    // The parameter must be an identifier for the resource type: ``com.vmware.appliance.recovery.backup.schedule``.
    // @throws NotFound if schedule associated with id does not exist.
    // @throws Error if any error occurs during the execution of the operation.
    Delete(scheduleParam string) error 

}
