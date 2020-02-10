/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Update
 * Used by client-side stubs.
 */

package appliance


// The ``Update`` interface provides methods to get the status of the appliance update. This interface was added in vSphere API 6.7.
type UpdateClient interface {

    // Gets the current status of the appliance update. This method was added in vSphere API 6.7.
    // @return Info structure containing the status information about the appliance.
    // @throws Error Generic error
    // @throws Unauthenticated session is not authenticated
    // @throws Unauthorized session is not authorized to perform this operation
	Get() (UpdateInfo, error)

    // Request the cancellation the update operation that is currently in progress. This method was added in vSphere API 6.7.
    // @throws Error Generic error
    // @throws NotAllowedInCurrentState Current task is not cancellable
    // @throws Unauthenticated session is not authenticated
    // @throws Unauthorized session is not authorized to perform this operation
	Cancel() error
}
