/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: FirewallSections
 * Used by client-side stubs.
 */

package firewalls

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

type FirewallSectionsClient interface {

    // Read a Firewall and the complete tree underneath. Returns the populated Firewall object.
    //
    // @param enforcementPointNameParam Enforcement Point Name (required)
    // @param firewallSectionIdParam Firewall Section Id (required)
    // @return com.vmware.nsx_policy.model.RealizedFirewallSection
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get(enforcementPointNameParam string, firewallSectionIdParam string) (model.RealizedFirewallSection, error)

    // Paginated list of all Firewalls. Returns populated Firewalls.
    //
    // @param enforcementPointNameParam Enforcement Point Name (required)
    // @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
    // @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
    // @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
    // @param sortAscendingParam (optional)
    // @param sortByParam Field by which records are sorted (optional)
    // @return com.vmware.nsx_policy.model.RealizedFirewallSectionListResult
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	List(enforcementPointNameParam string, cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.RealizedFirewallSectionListResult, error)
}
