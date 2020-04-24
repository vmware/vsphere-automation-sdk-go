/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Nslookup
 * Used by client-side stubs.
 */

package dns_forwarder

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-gm/model"
)

type NslookupClient interface {

    // Query the nameserver for an ip-address or a FQDN of the given an address optionally using an specified DNS server. If the address is a fqdn, nslookup will resolve ip-address with it. If the address is an ip-address, do a reverse lookup and answer fqdn(s). If enforcement point is specified, then DNS forwarder nslookup answer will get fetched from specified enforcement point. Otherwise from all enforcement points.
    //
    // @param tier1IdParam (required)
    // @param addressParam IP address or FQDN for nslookup (optional)
    // @param enforcementPointPathParam String Path of the enforcement point (optional)
    // @return com.vmware.nsx_global_policy.model.AggregatePolicyDnsAnswer
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get(tier1IdParam string, addressParam *string, enforcementPointPathParam *string) (model.AggregatePolicyDnsAnswer, error)
}
