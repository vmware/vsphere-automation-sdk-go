/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Firewall
 * Used by client-side stubs.
 */

package dashboard

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/vmc/model"
)

type FirewallClient interface {

    // Retrieve firewall dashboard statistics for a management or compute gateway (NSX Edge).
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
