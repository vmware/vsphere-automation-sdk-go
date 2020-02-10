/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: LocalAccounts
 * Used by client-side stubs.
 */

package appliance


// The ``LocalAccounts`` interface provides methods to manage local user account. This interface was added in vSphere API 6.7.
type LocalAccountsClient interface {

    // Get the local user account information. This method was added in vSphere API 6.7.
    //
    // @param usernameParam User login name
    // @return Local user account information
    // @throws NotFound If the account is not found
    // @throws Error Generic error
	Get(usernameParam string) (LocalAccountsInfo, error)

    // Get a list of the local user accounts. This method was added in vSphere API 6.7.
    // @return List of identifiers
    // The return value will contain identifiers for the resource type: ``com.vmware.appliance.local_accounts``.
    // @throws Error Generic error
	List() ([]string, error)

    // Create a new local user account. This method was added in vSphere API 6.7.
    //
    // @param usernameParam User login name
    // The parameter must be an identifier for the resource type: ``com.vmware.appliance.local_accounts``.
    // @param configParam User configuration
    // @throws AlreadyExists If an account already exists
    // @throws InvalidArgument If a username is invalid (username is validated against [a-zA-Z0-9][a-zA-Z0-9\-\.\\\\@]\*[a-zA-Z0-9] pattern)
    // @throws Error Generic error
	Create(usernameParam string, configParam LocalAccountsConfig) error

    // Set local user account properties. This method was added in vSphere API 6.7.
    //
    // @param usernameParam User login name
    // The parameter must be an identifier for the resource type: ``com.vmware.appliance.local_accounts``.
    // @param configParam User configuration
    // @throws NotFound If the account is not found
    // @throws Error Generic error
	Set(usernameParam string, configParam LocalAccountsConfig) error

    // Update selected fields in local user account properties. This method was added in vSphere API 6.7.
    //
    // @param usernameParam User login name
    // The parameter must be an identifier for the resource type: ``com.vmware.appliance.local_accounts``.
    // @param configParam User configuration
    // @throws NotFound If the account is not found
    // @throws Error Generic error
	Update(usernameParam string, configParam LocalAccountsUpdateConfig) error

    // Delete a local user account. This method was added in vSphere API 6.7.
    //
    // @param usernameParam User login name
    // The parameter must be an identifier for the resource type: ``com.vmware.appliance.local_accounts``.
    // @throws NotFound If the account is not found
    // @throws Error Generic error
	Delete(usernameParam string) error
}
