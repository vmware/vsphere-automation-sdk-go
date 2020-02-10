/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Policy
 * Used by client-side stubs.
 */

package local_accounts


// The ``Policy`` interface provides methods to manage local user accounts. This interface was added in vSphere API 6.7.
type PolicyClient interface {

    // Get the global password policy. This method was added in vSphere API 6.7.
    // @return Global password policy
    // @throws Error Generic error
	Get() (PolicyInfo, error)

    // Set the global password policy. This method was added in vSphere API 6.7.
    //
    // @param policyParam Global password policy
    // @throws InvalidArgument if passed policy values are < -1 or > 99999
    // @throws Error Generic error
	Set(policyParam PolicyInfo) error
}
