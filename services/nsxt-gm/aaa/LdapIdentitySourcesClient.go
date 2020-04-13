/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: LdapIdentitySources
 * Used by client-side stubs.
 */

package aaa

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/services/nsxt-gm/model"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
)

type LdapIdentitySourcesClient interface {

    // Delete an LDAP identity source. Users defined in that source will no longer be able to access NSX.
    //
    // @param ldapIdentitySourceIdParam (required)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Delete(ldapIdentitySourceIdParam string) error

    // Attempt to connect to an LDAP server and retrieve the server certificate it presents.
    //
    // @param identitySourceLdapServerEndpointParam (required)
    // @return com.vmware.nsx_global_policy.model.PeerCertificateChain
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Fetchcertificate(identitySourceLdapServerEndpointParam model.IdentitySourceLdapServerEndpoint) (model.PeerCertificateChain, error)

    // Return details about one LDAP identity source
    //
    // @param ldapIdentitySourceIdParam (required)
    // @return com.vmware.nsx_global_policy.model.LdapIdentitySource
    // The return value will contain all the properties defined in model.LdapIdentitySource.
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get(ldapIdentitySourceIdParam string) (*data.StructValue, error)

    // Return a list of all configured LDAP identity sources.
    //
    // @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
    // @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
    // @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
    // @param sortAscendingParam (optional)
    // @param sortByParam Field by which records are sorted (optional)
    // @return com.vmware.nsx_global_policy.model.LdapIdentitySourceListResult
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	List(cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.LdapIdentitySourceListResult, error)

    // Attempt to connect to an existing LDAP identity source and report any errors encountered.
    //
    // @param ldapIdentitySourceIdParam (required)
    // @return com.vmware.nsx_global_policy.model.LdapIdentitySourceProbeResults
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Probe(ldapIdentitySourceIdParam string) (model.LdapIdentitySourceProbeResults, error)

    // Verify that the configuration of an LDAP identity source is correct before actually creating the source.
    //
    // @param ldapIdentitySourceParam (required)
    // The parameter must contain all the properties defined in model.LdapIdentitySource.
    // @return com.vmware.nsx_global_policy.model.LdapIdentitySourceProbeResults
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Probeidentitysource(ldapIdentitySourceParam *data.StructValue) (model.LdapIdentitySourceProbeResults, error)

    // Attempt to connect to an LDAP server and ensure that the server can be contacted using the given URL and authentication credentials.
    //
    // @param identitySourceLdapServerParam (required)
    // @return com.vmware.nsx_global_policy.model.IdentitySourceLdapServerProbeResult
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Probeldapserver(identitySourceLdapServerParam model.IdentitySourceLdapServer) (model.IdentitySourceLdapServerProbeResult, error)

    // Update the configuration of an existing LDAP identity source. You may wish to verify the new configuration using the POST /aaa/ldap-identity-sources?action=probe API before changing the configuration.
    //
    // @param ldapIdentitySourceIdParam (required)
    // @param ldapIdentitySourceParam (required)
    // The parameter must contain all the properties defined in model.LdapIdentitySource.
    // @return com.vmware.nsx_global_policy.model.LdapIdentitySource
    // The return value will contain all the properties defined in model.LdapIdentitySource.
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Update(ldapIdentitySourceIdParam string, ldapIdentitySourceParam *data.StructValue) (*data.StructValue, error)
}
