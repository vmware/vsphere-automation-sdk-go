/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Hostname
 * Used by client-side stubs.
 */

package dns


// ``Hostname`` interface provides methods Performs operations on Fully Qualified Doman Name.
type HostnameClient interface {

    // Test the Fully Qualified Domain Name.
    //
    // @param nameParam FQDN.
    // @return FQDN status
    // @throws Error Generic error
	Test(nameParam string) (HostnameTestStatusInfo, error)

    // Set the Fully Qualified Domain Name.
    //
    // @param nameParam FQDN.
    // @throws Error Generic error
	Set(nameParam string) error

    // Get the Fully Qualified Doman Name.
    // @return FQDN.
    // @throws Error Generic error
	Get() (string, error)
}
