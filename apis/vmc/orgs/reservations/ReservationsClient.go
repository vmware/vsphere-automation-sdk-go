/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Reservations
 * Used by client-side stubs.
 */

package reservations

import (
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/apis/vmc/model"
)

type ReservationsClient interface {


    // Get all reservations for this org
    //
    // @param orgParam Organization identifier. (required)
    // @throws Unauthenticated  Unauthorized
    // @throws Unauthorized  Access not allowed to the operation for the current user
    List(orgParam string) ([]model.MaintenanceWindowEntry, error) 

}
