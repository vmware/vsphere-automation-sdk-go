/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Domains
 * Used by client-side stubs.
 */

package dns


// ``Domains`` interface provides methods DNS search domains.
type DomainsClient interface {

    // Add domain to DNS search domains.
    //
    // @param domainParam Domain to add.
    // @throws Error Generic error
	Add(domainParam string) error

    // Set DNS search domains.
    //
    // @param domainsParam List of domains.
    // @throws Error Generic error
	Set(domainsParam []string) error

    // Get list of DNS search domains.
    // @return List of domains.
    // @throws Error Generic error
	List() ([]string, error)
}
