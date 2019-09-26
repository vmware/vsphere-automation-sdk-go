/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Pending
 * Used by client-side stubs.
 */

package pending

import (
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/bindings/appliance"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/bindings/appliance/update"
)

// The ``Pending`` interface provides methods to manipulate pending updates.
type PendingClient interface {


    // Checks if new updates are available.
    //
    // @param sourceTypeParam type of the source
    // @param urlParam specific URL to check at
    // If null then URL is taken from the policy settings
    // @return List of the update summaries
    // @throws Error Generic error
    // @throws NotFound source is not found
    // @throws Unauthenticated session is not authenticated
    // @throws Unauthorized session is not authorized to perform this operation
    List(sourceTypeParam SourceType, urlParam *string) ([]update.Summary, error) 


    // Gets update information
    //
    // @param versionParam Update version
    // The parameter must be an identifier for the resource type: ``com.vmware.appliance.update.pending``.
    // @return Update
    // @throws Error Generic error
    // @throws Unauthenticated session is not authenticated
    // @throws Unauthorized session is not authorized to perform this operation
    // @throws NotFound the update is not found
    // @throws AlreadyInDesiredState if the update of this version is already installed
    Get(versionParam string) (Info, error) 


    // Runs update precheck
    //
    // @param versionParam Update version
    // The parameter must be an identifier for the resource type: ``com.vmware.appliance.update.pending``.
    // @return PrecheckResult
    // @throws Error Generic error
    // @throws Unauthenticated session is not authenticated
    // @throws Unauthorized session is not authorized to perform this operation
    // @throws NotFound the update is not found
    // @throws AlreadyInDesiredState if this version is already installed
    // @throws NotAllowedInCurrentState if another operation is in progress
    Precheck(versionParam string) (PrecheckResult, error) 


    // Starts staging the appliance update. The updates are searched for in the following order: staged, CDROM, URL
    //
    // @param versionParam Update version
    // The parameter must be an identifier for the resource type: ``com.vmware.appliance.update.pending``.
    // @throws Error Generic error
    // @throws Unauthenticated session is not authenticated
    // @throws Unauthorized session is not authorized to perform this operation
    // @throws NotFound the update is not found
    // @throws AlreadyInDesiredState if the update of this version is already installed
    // @throws AlreadyExists the update is already staged
    // @throws NotAllowedInCurrentState if appliance update state prevents staging
    Stage(versionParam string) error 


    // Validates the user provided data before the update installation.
    //
    // @param versionParam Update version
    // The parameter must be an identifier for the resource type: ``com.vmware.appliance.update.pending``.
    // @param userDataParam map of user provided data with IDs
    // The key in the parameter map must be an identifier for the resource type: ``com.vmware.applicance.update.pending.dataitem``.
    // @return Issues struct with the issues found during the validation
    // @throws Error Generic error
    // @throws Unauthenticated session is not authenticated
    // @throws Unauthorized session is not authorized to perform this operation
    // @throws NotFound if the update is not found
    // @throws AlreadyInDesiredState if the update of this version is already installed
    // @throws NotAllowedInCurrentState if appliance update state prevents running an check
    Validate(versionParam string, userDataParam map[string]string) (appliance.Notifications, error) 


    // Starts operation of installing the appliance update. Will fail is the update is not staged
    //
    // @param versionParam Update version
    // The parameter must be an identifier for the resource type: ``com.vmware.appliance.update.pending``.
    // @param userDataParam map of user provided data with IDs
    // The key in the parameter map must be an identifier for the resource type: ``com.vmware.applicance.update.pending.dataitem``.
    // @throws Error Generic error
    // @throws Error Generic error
    // @throws Unauthenticated session is not authenticated
    // @throws Unauthorized session is not authorized to perform this operation
    // @throws NotFound if the update is not found
    // @throws AlreadyInDesiredState if the update of this version is already installed
    // @throws NotAllowedInCurrentState if appliance update state prevents running an update or not staged
    Install(versionParam string, userDataParam map[string]string) error 


    // Starts operation of installing the appliance update. Will stage update if not already staged The updates are searched for in the following order: staged, CDROM, URL
    //
    // @param versionParam Update version
    // The parameter must be an identifier for the resource type: ``com.vmware.appliance.update.pending``.
    // @param userDataParam map of user provided data with IDs
    // The key in the parameter map must be an identifier for the resource type: ``com.vmware.applicance.update.pending.dataitem``.
    // @throws Error Generic error
    // @throws Error Generic error
    // @throws Unauthenticated session is not authenticated
    // @throws Unauthorized session is not authorized to perform this operation
    // @throws NotFound if the update is not found
    // @throws AlreadyInDesiredState if the update of this version is already installed
    // @throws NotAllowedInCurrentState if appliance update state prevents running an update
    StageAndInstall(versionParam string, userDataParam map[string]string) error 

}
