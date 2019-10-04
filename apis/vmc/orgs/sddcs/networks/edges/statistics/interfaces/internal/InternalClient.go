/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Internal
 * Used by client-side stubs.
 */

package internal

import (
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/apis/vmc/model"
)

type InternalClient interface {


    // Retrieve internal interface statistics for a management or compute gateway (NSX Edge).
    //
    // @param orgParam Organization identifier. (required)
    // @param sddcParam Sddc Identifier. (required)
    // @param edgeIdParam Edge Identifier. (required)
    // @param startTimeParam Show statistics from this start time (in milliseconds since epoch). (optional)
    // @param endTimeParam Show statistics until this end time (in milliseconds since epoch). (optional)
    // @return com.vmware.vmc.model.CbmStatistics
    // @throws InvalidRequest  Bad request. Request object passed is invalid.
    // @throws Unauthorized  Forbidden. Authorization header not provided.
    // @throws NotFound  Not found. Requested object not found.
    Get(orgParam string, sddcParam string, edgeIdParam string, startTimeParam *int64, endTimeParam *int64) (model.CbmStatistics, error) 

}
