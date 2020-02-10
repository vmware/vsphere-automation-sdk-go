/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Timesync
 * Used by client-side stubs.
 */

package appliance


// ``Timesync`` interface provides methods Performs time synchronization configuration. This interface was added in vSphere API 6.7.
type TimesyncClient interface {

    // Set time synchronization mode. This method was added in vSphere API 6.7.
    //
    // @param modeParam Time synchronization mode.
    // @throws Error Generic error
	Set(modeParam TimesyncTimeSyncMode) error

    // Get time synchronization mode. This method was added in vSphere API 6.7.
    // @return Time synchronization mode.
    // @throws Error Generic error
	Get() (TimesyncTimeSyncMode, error)
}
