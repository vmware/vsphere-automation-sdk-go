/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: PeerConfig
 * Used by client-side stubs.
 */

package sessions

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

type PeerConfigClient interface {

    // Get peer config for the Tier-0 L2VPN session to configure the remote side of the tunnel. - no enforcement point path specified: L2VPN Session Peer Codes will be evaluated on each enforcement point. - enforcement point paths specified: L2VPN Session Peer Codes are evaluated only on the given enforcement points. API supported only when L2VPN Service is in Server Mode.
    //
    // @param tier0IdParam (required)
    // @param localeServiceIdParam (required)
    // @param serviceIdParam (required)
    // @param sessionIdParam (required)
    // @param enforcementPointPathParam String Path of the enforcement point (optional)
    // @return com.vmware.nsx_policy.model.AggregateL2VPNSessionPeerConfig
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get(tier0IdParam string, localeServiceIdParam string, serviceIdParam string, sessionIdParam string, enforcementPointPathParam *string) (model.AggregateL2VPNSessionPeerConfig, error)
}
