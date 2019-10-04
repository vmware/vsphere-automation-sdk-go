/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: AccountLink
 * Used by client-side stubs.
 */

package accountLink

import (
)

type AccountLinkClient interface {


    // Gets a link that can be used on a customer's account to start the linking process.
    //
    // @param orgParam Organization identifier. (required)
    // @throws Error  Generic Error
    Get(orgParam string) error 

}
