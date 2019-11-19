/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Reservation
 * Used by client-side stubs.
 */

package tbrs

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/services/vmc/model"
)

type ReservationClient interface {

    // Retreive all reservations for all SDDCs in this org
    //
    // @param orgParam Organization identifier. (required)
    // @param sddcStateParam SDDCs and/or states to query (optional)
    // @throws Unauthenticated  Unauthorized
    // @throws InvalidRequest  Bad Call
    // @throws Unauthorized  Forbidden
	Post(orgParam string, sddcStateParam *model.SddcStateRequest) ([]model.ReservationWindow, error)
}
