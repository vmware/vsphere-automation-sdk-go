/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Recovery
 * Used by client-side stubs.
 */

package appliance


// The ``Recovery`` interface provides methods to invoke an appliance recovery (backup and restore). This interface was added in vSphere API 6.7.
type RecoveryClient interface {

    // Gets the properties of the appliance Recovery subsystem. This method was added in vSphere API 6.7.
    // @return Structure containing the properties of the Recovery subsystem.
    // @throws Error if any error occurs during the execution of the operation.
	Get() (RecoveryInfo, error)
}
