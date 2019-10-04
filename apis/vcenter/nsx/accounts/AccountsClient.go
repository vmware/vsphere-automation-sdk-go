/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Accounts
 * Used by client-side stubs.
 */

package accounts

import (
)

// The ``Accounts`` interface represents all the operations details of user accounts on my.vmware.com portal. **Warning:** This interface is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type AccountsClient interface {


    // authenticate user on the my.vmware.com portal. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param specParam usename and password
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
    // @throws Error if there is some unknown internal error. The accompanying error message will give more details about the error.
    Authenticate(specParam Spec) error 


    // Get Entitlement Accounts (EA) attached to the user account on the my.vmware.com portal. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    // @return Accounts
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
    // @throws Error if there is some unknown internal error. The accompanying error message will give more details about the error.
    GetAccounts() ([]AccountsInfo, error) 


    // check if selected user account is entitle to download product binaries from the my.vmware.com portal. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param accountsIdParam account ID.
    // The parameter must be an identifier for the resource type: ``ACCOUNT``.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
    // @throws Error if there is some unknown internal error. The accompanying error message will give more details about the error.
    Check(accountsIdParam string) error 

}
