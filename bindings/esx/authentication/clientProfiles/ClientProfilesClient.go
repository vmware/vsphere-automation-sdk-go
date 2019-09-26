/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: ClientProfiles
 * Used by client-side stubs.
 */

package clientProfiles

import (
)

// The ``ClientProfiles`` interface provides methods to manage ESX authentication tokens claims. Subject matching is used to provide permission claims described by AccessGrant list.
//
//  When multiple profiles match, union of the AccessGrant elements is applied. When no profiles match, no permission claims are applied.
type ClientProfilesClient interface {


    // List the existing client profiles.
    //
    // @param filterParam A filter for the returned list. If the filter doesn't match, return an empty list.
    // If {\\\\@term.unset} return all profiles.
    // @param projectionParam The type of the returned summary - brief, normal or full.
    // If {\\\\@term.unset} a normal projection will be used.
    // @return The list of current client profiles.
    // @throws Error if there is a problem accessing the stored data.
    // @throws InvalidArgument if the arguments contain invalid data.
    // @throws Unauthenticated if the user can not be authenticated.
    List(filterParam *FilterSpec, projectionParam *SummaryType) ([]Summary, error) 


    // Create a new client profile.
    //
    // @param specParam Settings for the new client profile.
    // @return The new client profile.
    // The return value will be an identifier for the resource type: ``com.vmware.esx.authentication.clientprofile``.
    // @throws AlreadyExists if there is already a ClientProfiles instance with the same subject.
    // @throws InvalidArgument if the CreateSpec contains invalid data.
    // @throws Error if there is a problem storing the data.
    // @throws Unauthenticated if the user can not be authenticated.
    Create(specParam CreateSpec) (string, error) 


    // Get the details of a client profile.
    //
    // @param profileParam The requested client profile identifier.
    // The parameter must be an identifier for the resource type: ``com.vmware.esx.authentication.clientprofile``.
    // @return The requested client profile.
    // @throws NotFound if the profile is not found.
    // @throws Error if there is a problem accessing the stored data.
    // @throws Unauthenticated if the user can not be authenticated.
    Get(profileParam string) (Info, error) 


    // Update the access grants in an existing client profile.
    //
    // @param profileParam The client profile identifier.
    // The parameter must be an identifier for the resource type: ``com.vmware.esx.authentication.clientprofile``.
    // @param specParam The new settings.
    // @throws NotFound if the profile is not found.
    // @throws InvalidArgument if the UpdateSpec contains invalid data.
    // @throws Error if there is a problem storing the data.
    // @throws Unauthenticated if the user can not be authenticated.
    Update(profileParam string, specParam UpdateSpec) error 


    // Delete an existing client profile.
    //
    // @param profileParam The client profile identifier to remove.
    // The parameter must be an identifier for the resource type: ``com.vmware.esx.authentication.clientprofile``.
    // @throws NotFound if the profile is not found.
    // @throws Error if there is a problem storing the data.
    // @throws Unauthenticated if the user can not be authenticated.
    Delete(profileParam string) error 

}
