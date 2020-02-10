/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Schedules
 * Used by client-side stubs.
 */

package backup


// The ``Schedules`` interface provides methods to be performed to manage backup schedules. This interface was added in vSphere API 6.7.
type SchedulesClient interface {

    // Returns a list of existing schedules with details. This method was added in vSphere API 6.7.
    // @return Map of schedule id to Info Structure
    // The key in the return value map will be an identifier for the resource type: ``com.vmware.appliance.recovery.backup.schedule``.
    // @throws Error if any error occurs during the execution of the operation.
	List() (map[string]SchedulesInfo, error)

    // Initiate backup with the specified schedule. This method was added in vSphere API 6.7.
    //
    // @param scheduleParam Identifier of the schedule
    // The parameter must be an identifier for the resource type: ``com.vmware.appliance.recovery.backup.schedule``.
    // @param commentParam field that specifies the description for the backup.
    // If null the backup will have an empty comment.
    // @return BackupJobStatus Structure
    // @throws FeatureInUse if a backup or restore is already in progress.
    // @throws NotFound if schedule associated with id does not exist.
    // @throws Error if any error occurs during the execution of the operation.
	Run(scheduleParam string, commentParam *string) (JobBackupJobStatus, error)

    // Returns an existing schedule information based on id. This method was added in vSphere API 6.7.
    //
    // @param scheduleParam Identifier of the schedule
    // The parameter must be an identifier for the resource type: ``com.vmware.appliance.recovery.backup.schedule``.
    // @return Info Structure
    // @throws NotFound if schedule associated with id does not exist.
    // @throws Error if any error occurs during the execution of the operation.
	Get(scheduleParam string) (SchedulesInfo, error)

    // Creates a schedule. This method was added in vSphere API 6.7.
    //
    // @param scheduleParam Identifier of the schedule
    // The parameter must be an identifier for the resource type: ``com.vmware.appliance.recovery.backup.schedule``.
    // @param specParam CreateSpec Structure
    // @throws InvalidArgument if provided with invalid schedule specification.
    // @throws AlreadyExists if the schedule with the given id already exists.
    // @throws Error if any error occurs during the execution of the operation.
	Create(scheduleParam string, specParam SchedulesCreateSpec) error

    // Updates a schedule. This method was added in vSphere API 6.7.
    //
    // @param scheduleParam Identifier of the schedule
    // The parameter must be an identifier for the resource type: ``com.vmware.appliance.recovery.backup.schedule``.
    // @param specParam UpdateSpec Structure
    // @throws InvalidArgument if provided with invalid schedule specification.
    // @throws NotFound if schedule associated with id does not exist.
    // @throws Error if any error occurs during the execution of the operation.
	Update(scheduleParam string, specParam SchedulesUpdateSpec) error

    // Deletes an existing schedule. This method was added in vSphere API 6.7.
    //
    // @param scheduleParam Identifier of the schedule
    // The parameter must be an identifier for the resource type: ``com.vmware.appliance.recovery.backup.schedule``.
    // @throws NotFound if schedule associated with id does not exist.
    // @throws Error if any error occurs during the execution of the operation.
	Delete(scheduleParam string) error
}
