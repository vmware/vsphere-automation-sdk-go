/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Statistics
 * Used by client-side stubs.
 */

package firewall

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/vmc/model"
)

type StatisticsClient interface {

    // Retrieve statistics for a specific firewall rule for a management or compute gateway (NSX Edge).
    //
    // @param orgParam Organization identifier. (required)
    // @param sddcParam Sddc Identifier. (required)
    // @param edgeIdParam Edge Identifier. (required)
    // @param ruleIdParam Rule Identifier. (required)
    // @return com.vmware.vmc.model.FirewallRuleStats
    // @throws InvalidRequest  Bad request. Request object passed is invalid.
    // @throws Unauthorized  Forbidden. Authorization header not provided
    // @throws NotFound  Not found. Requested object not found.
	Get(orgParam string, sddcParam string, edgeIdParam string, ruleIdParam int64) (model.FirewallRuleStats, error)
}
