/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Search
 * Used by client-side stubs.
 */

package ldap_identity_sources

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

type SearchClient interface {

    // Search the LDAP identity source for users and groups that match the given filter_value. In most cases, the LDAP source performs a case-insensitive search.
    //
    // @param ldapIdentitySourceIdParam (required)
    // @param filterValueParam Search filter value (optional)
    // @return com.vmware.nsx_policy.model.LdapIdentitySourceSearchResultList
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Create(ldapIdentitySourceIdParam string, filterValueParam *string) (model.LdapIdentitySourceSearchResultList, error)
}
