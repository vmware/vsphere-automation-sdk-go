/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: AccountLink
 * Used by client-side stubs.
 */

package orgs


type AccountLinkClient interface {

    // Gets a link that can be used on a customer's account to start the linking process.
    //
    // @param orgParam Organization identifier (required)
    // @throws Error  Generic Error
	Get(orgParam string) error
}
