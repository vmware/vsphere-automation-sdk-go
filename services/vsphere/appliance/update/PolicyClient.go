/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Policy
 * Used by client-side stubs.
 */

package update


// The ``Policy`` interface provides methods to set/get background check for the new updates. This interface was added in vSphere API 6.7.
type PolicyClient interface {

    // Gets the automatic update checking and staging policy. This method was added in vSphere API 6.7.
    // @return Structure containing the policy for the appliance update.
    // @throws Error Generic error
    // @throws Unauthenticated session is not authenticated
    // @throws Unauthorized session is not authorized to perform this operation
	Get() (PolicyInfo, error)

    // Sets the automatic update checking and staging policy. This method was added in vSphere API 6.7.
    //
    // @param policyParam Info structure containing the policy for the appliance update.
    // @throws Error Generic error
    // @throws Unauthenticated session is not authenticated
    // @throws Unauthorized session is not authorized to perform this operation
	Set(policyParam PolicyConfig) error
}
