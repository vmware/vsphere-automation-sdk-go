/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: LdapServers
 * Used by client-side stubs.
 */

package firewall_identity_stores

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt/model"
)

type LdapServersClient interface {

    // The API tests a LDAP server connection for an already configured domain. If the connection is successful, the response will be HTTP status 200. Otherwise the response will be HTTP status 500 and corresponding error message will be returned.
    //
    // @param firewallIdentityStoreIdParam Firewall Identity store identifier (required)
    // @param ldapServerIdParam LDAP server identifier (required)
    // @param actionParam LDAP server test requested (required)
    // @param enforcementPointPathParam String Path of the enforcement point (optional)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Create(firewallIdentityStoreIdParam string, ldapServerIdParam string, actionParam string, enforcementPointPathParam *string) error

    // Delete a LDAP server for Firewall Identity store
    //
    // @param firewallIdentityStoreIdParam Firewall Identity store identifier (required)
    // @param ldapServerIdParam LDAP server identifier (required)
    // @param enforcementPointPathParam String Path of the enforcement point (optional)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Delete(firewallIdentityStoreIdParam string, ldapServerIdParam string, enforcementPointPathParam *string) error

    // Get a specific LDAP server for a given Firewall Identity store
    //
    // @param firewallIdentityStoreIdParam Firewall Identity store identifier (required)
    // @param ldapServerIdParam LDAP server identifier (required)
    // @param enforcementPointPathParam String Path of the enforcement point (optional)
    // @return com.vmware.nsx_policy.model.DirectoryLdapServer
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get(firewallIdentityStoreIdParam string, ldapServerIdParam string, enforcementPointPathParam *string) (model.DirectoryLdapServer, error)

    // List all configured domain LDAP servers
    //
    // @param firewallIdentityStoreIdParam Firewall Identity store identifier (required)
    // @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
    // @param enforcementPointPathParam String Path of the enforcement point (optional)
    // @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
    // @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
    // @param sortAscendingParam (optional)
    // @param sortByParam Field by which records are sorted (optional)
    // @return com.vmware.nsx_policy.model.DirectoryLdapServerListResults
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	List(firewallIdentityStoreIdParam string, cursorParam *string, enforcementPointPathParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.DirectoryLdapServerListResults, error)

    // More than one LDAP server can be created and only one LDAP server is used to synchronize directory objects. If more than one LDAP server is configured, NSX will try all the servers until it is able to successfully connect to one.
    //
    // @param firewallIdentityStoreIdParam Firewall Identity store identifier (required)
    // @param ldapServerIdParam LDAP server identifier (required)
    // @param directoryLdapServerParam (required)
    // @param enforcementPointPathParam String Path of the enforcement point (optional)
    // @return com.vmware.nsx_policy.model.DirectoryLdapServer
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Patch(firewallIdentityStoreIdParam string, ldapServerIdParam string, directoryLdapServerParam model.DirectoryLdapServer, enforcementPointPathParam *string) (model.DirectoryLdapServer, error)

    // Update a LDAP server for Firewall Identity store
    //
    // @param firewallIdentityStoreIdParam Firewall Identity store identifier (required)
    // @param ldapServerIdParam LDAP server identifier (required)
    // @param directoryLdapServerParam (required)
    // @param enforcementPointPathParam String Path of the enforcement point (optional)
    // @return com.vmware.nsx_policy.model.DirectoryLdapServer
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Update(firewallIdentityStoreIdParam string, ldapServerIdParam string, directoryLdapServerParam model.DirectoryLdapServer, enforcementPointPathParam *string) (model.DirectoryLdapServer, error)
}
