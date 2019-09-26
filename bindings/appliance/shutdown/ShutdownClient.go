/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Shutdown
 * Used by client-side stubs.
 */

package shutdown

import (
)

// ``Shutdown`` interface provides methods Performs reboot/shutdown operations on appliance.
type ShutdownClient interface {


    // Cancel pending shutdown action.
    // @throws Error Generic error
    Cancel() error 


    // Power off the appliance.
    //
    // @param delayParam Minutes after which poweroff should start. If 0 is specified, poweroff will start immediately.
    // @param reasonParam Reason for peforming poweroff.
    // @throws Error Generic error
    Poweroff(delayParam int64, reasonParam string) error 


    // Reboot the appliance.
    //
    // @param delayParam Minutes after which reboot should start. If 0 is specified, reboot will start immediately.
    // @param reasonParam Reason for peforming reboot.
    // @throws Error Generic error
    Reboot(delayParam int64, reasonParam string) error 


    // Get details about the pending shutdown action.
    // @return Configuration of pending shutdown action.
    // @throws Error Generic error
    Get() (ShutdownConfig, error) 

}
