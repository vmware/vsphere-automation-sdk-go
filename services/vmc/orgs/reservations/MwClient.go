/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Mw
 * Used by client-side stubs.
 */

package reservations

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/vmc/model"
)

type MwClient interface {

    // get the maintenance window for this SDDC
    //
    // @param orgParam Organization identifier. (required)
    // @param reservationParam Reservation Identifier (required)
    // @return com.vmware.vmc.model.MaintenanceWindowGet
    // @throws Unauthenticated  Unauthorized
    // @throws Unauthorized  Access not allowed to the operation for the current user
	Get(orgParam string, reservationParam string) (model.MaintenanceWindowGet, error)

    // update the maintenance window for this SDDC
    //
    // @param orgParam Organization identifier. (required)
    // @param reservationParam Reservation Identifier (required)
    // @param windowParam Maintenance Window (required)
    // @return com.vmware.vmc.model.MaintenanceWindow
    // @throws Unauthenticated  Unauthorized
    // @throws ConcurrentChange  Conflict with exiting reservation
    // @throws InvalidRequest  The reservation is not in a state that's valid for updates
    // @throws Unauthorized  Access not allowed to the operation for the current user
	Put(orgParam string, reservationParam string, windowParam model.MaintenanceWindow) (model.MaintenanceWindow, error)
}
