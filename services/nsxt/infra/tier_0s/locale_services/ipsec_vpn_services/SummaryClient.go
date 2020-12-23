/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Summary
 * Used by client-side stubs.
 */

package ipsec_vpn_services

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

type SummaryClient interface {

    // Summarized view of all tier-0 IPSec VPN sessions for a specified service.
    //
    // @param tier0IdParam (required)
    // @param localeServiceIdParam (required)
    // @param serviceIdParam (required)
    // @param enforcementPointPathParam String Path of the enforcement point (optional)
    // @param sourceParam Data source type. (optional)
    // @return com.vmware.nsx_policy.model.PolicyIpsecVpnIkeServiceSummary
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get(tier0IdParam string, localeServiceIdParam string, serviceIdParam string, enforcementPointPathParam *string, sourceParam *string) (model.PolicyIpsecVpnIkeServiceSummary, error)
}
