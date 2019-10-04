/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Reservation
 * Used by client-side stubs.
 */

package reservation

import (
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/apis/vmc/model"
)

type ReservationClient interface {


    // Retreive all reservations for all SDDCs in this org
    //
    // @param orgParam Organization identifier. (required)
    // @param sddcStateParam SDDCs and/or states to query (optional)
    // @throws Unauthenticated  Unauthorized
    // @throws InvalidRequest  Bad Call
    // @throws Unauthorized  Forbidden
    Post(orgParam string, sddcStateParam *model.SddcStateRequest) (map[string][]model.ReservationWindow, error) 

}
