/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Shutdown
 * Used by client-side stubs.
 */

package appliance


// ``Shutdown`` interface provides methods Performs reboot/shutdown operations on appliance. This interface was added in vSphere API 6.7.
type ShutdownClient interface {

    // Cancel pending shutdown action. This method was added in vSphere API 6.7.
    // @throws Error Generic error
	Cancel() error

    // Power off the appliance. This method was added in vSphere API 6.7.
    //
    // @param delayParam Minutes after which poweroff should start. If 0 is specified, poweroff will start immediately.
    // @param reasonParam Reason for peforming poweroff.
    // @throws Error Generic error
	Poweroff(delayParam int64, reasonParam string) error

    // Reboot the appliance. This method was added in vSphere API 6.7.
    //
    // @param delayParam Minutes after which reboot should start. If 0 is specified, reboot will start immediately.
    // @param reasonParam Reason for peforming reboot.
    // @throws Error Generic error
	Reboot(delayParam int64, reasonParam string) error

    // Get details about the pending shutdown action. This method was added in vSphere API 6.7.
    // @return Configuration of pending shutdown action.
    // @throws Error Generic error
	Get() (ShutdownShutdownConfig, error)
}
