/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Ipsec
 * Used by client-side stubs.
 */

package ipsec

import (
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/apis/vmc/model"
)

type IpsecClient interface {


    // Retrieve ipsec dashboard statistics for a management or compute gateway (NSX Edge).
    //
    // @param orgParam Organization identifier. (required)
    // @param sddcParam Sddc Identifier. (required)
    // @param edgeIdParam Edge Identifier. (required)
    // @param intervalParam 60 min by default, can be given as 1 - 60 min, oneDay, oneWeek, oneMonth, oneYear. (optional)
    // @return com.vmware.vmc.model.DashboardStatistics
    // @throws InvalidRequest  Bad request. Request object passed is invalid.
    // @throws Unauthorized  Forbidden. Authorization header not provided
    // @throws NotFound  Not found. Requested object not found.
    Get(orgParam string, sddcParam string, edgeIdParam string, intervalParam *string) (model.DashboardStatistics, error) 

}
