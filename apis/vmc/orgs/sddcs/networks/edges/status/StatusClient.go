/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Status
 * Used by client-side stubs.
 */

package status

import (
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/apis/vmc/model"
)

type StatusClient interface {


    // Retrieve the status of the specified management or compute gateway (NSX Edge).
    //
    // @param orgParam Organization identifier. (required)
    // @param sddcParam Sddc Identifier. (required)
    // @param edgeIdParam Edge Identifier. (required)
    // @param getlatestParam If true, retrieve the status live from the gateway (NSX Edge). If false, retrieve the latest available status from database. (optional)
    // @param detailedParam If true, retrieve detailed status per feature. If false, retrieve aggregated summary of status per feature. (optional)
    // @return com.vmware.vmc.model.EdgeStatus
    // @throws InvalidRequest  Bad request. Request object passed is invalid.
    // @throws Unauthorized  Forbidden. Authorization header not provided
    // @throws NotFound  Not found. Requested object not found.
    Get(orgParam string, sddcParam string, edgeIdParam string, getlatestParam *bool, detailedParam *bool) (model.EdgeStatus, error) 

}
