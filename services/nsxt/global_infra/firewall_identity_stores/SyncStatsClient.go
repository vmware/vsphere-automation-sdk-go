/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: SyncStats
 * Used by client-side stubs.
 */

package firewall_identity_stores

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

type SyncStatsClient interface {

    // Get Firewall identity store sync statistics for the given identifier
    //
    // @param firewallIdentityStoreIdParam Firewall identity store identifier (required)
    // @param enforcementPointPathParam String Path of the enforcement point (optional)
    // @return com.vmware.nsx_policy.model.DirectoryDomainSyncStats
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get(firewallIdentityStoreIdParam string, enforcementPointPathParam *string) (model.DirectoryDomainSyncStats, error)
}
