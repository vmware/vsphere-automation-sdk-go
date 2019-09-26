/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: LocalAccounts
 * Used by client-side stubs.
 */

package localAccounts

import (
)

// The ``LocalAccounts`` interface provides methods to manage local user account.
type LocalAccountsClient interface {


    // Get the local user account information.
    //
    // @param usernameParam User login name
    // @return Local user account information
    // @throws NotFound If the account is not found
    // @throws Error Generic error
    Get(usernameParam string) (Info, error) 


    // Get a list of the local user accounts.
    // @return List of identifiers
    // The return value will contain identifiers for the resource type: ``com.vmware.appliance.local_accounts``.
    // @throws Error Generic error
    List() ([]string, error) 


    // Create a new local user account.
    //
    // @param usernameParam User login name
    // The parameter must be an identifier for the resource type: ``com.vmware.appliance.local_accounts``.
    // @param configParam User configuration
    // @throws AlreadyExists If an account already exists
    // @throws InvalidArgument If a username is invalid (username is validated against [a-zA-Z0-9][a-zA-Z0-9\-\.\\\\@]\*[a-zA-Z0-9] pattern)
    // @throws Error Generic error
    Create(usernameParam string, configParam Config) error 


    // Set local user account properties.
    //
    // @param usernameParam User login name
    // The parameter must be an identifier for the resource type: ``com.vmware.appliance.local_accounts``.
    // @param configParam User configuration
    // @throws NotFound If the account is not found
    // @throws Error Generic error
    Set(usernameParam string, configParam Config) error 


    // Update selected fields in local user account properties.
    //
    // @param usernameParam User login name
    // The parameter must be an identifier for the resource type: ``com.vmware.appliance.local_accounts``.
    // @param configParam User configuration
    // @throws NotFound If the account is not found
    // @throws Error Generic error
    Update(usernameParam string, configParam UpdateConfig) error 


    // Delete a local user account.
    //
    // @param usernameParam User login name
    // The parameter must be an identifier for the resource type: ``com.vmware.appliance.local_accounts``.
    // @throws NotFound If the account is not found
    // @throws Error Generic error
    Delete(usernameParam string) error 

}
