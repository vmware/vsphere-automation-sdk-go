/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: FirewallSessionTimerProfileBindingMaps
 * Used by client-side stubs.
 */

package groups

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt-gm/model"
)

type FirewallSessionTimerProfileBindingMapsClient interface {

    // API will delete Firewall Session Timer Profile Binding
    //
    // @param domainIdParam Domain ID (required)
    // @param groupIdParam Group ID (required)
    // @param firewallSessionTimerProfileBindingMapIdParam Firewall Session Timer Profile Binding Map ID (required)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Delete(domainIdParam string, groupIdParam string, firewallSessionTimerProfileBindingMapIdParam string) error

    // API will get Firewall Session Timer Profile Binding Map
    //
    // @param domainIdParam Domain-ID (required)
    // @param groupIdParam Group ID (required)
    // @param firewallSessionTimerProfileBindingMapIdParam Firewall Session Timer Profile Binding Map ID (required)
    // @return com.vmware.nsx_global_policy.model.PolicyFirewallSessionTimerProfileBindingMap
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get(domainIdParam string, groupIdParam string, firewallSessionTimerProfileBindingMapIdParam string) (model.PolicyFirewallSessionTimerProfileBindingMap, error)

    // API will list all Firewall Session Timer Profile Binding Maps in current group id.
    //
    // @param domainIdParam (required)
    // @param groupIdParam (required)
    // @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
    // @param includeMarkForDeleteObjectsParam Include objects that are marked for deletion in results (optional, default to false)
    // @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
    // @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
    // @param sortAscendingParam (optional)
    // @param sortByParam Field by which records are sorted (optional)
    // @return com.vmware.nsx_global_policy.model.PolicyFirewallSessionTimerProfileBindingMapListResult
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	List(domainIdParam string, groupIdParam string, cursorParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.PolicyFirewallSessionTimerProfileBindingMapListResult, error)

    // API will create or update Firewall Session Timer profile binding map
    //
    // @param domainIdParam Domain ID (required)
    // @param groupIdParam Group ID (required)
    // @param firewallSessionTimerProfileBindingMapIdParam Firewall Session Timer Profile Binding Map ID (required)
    // @param policyFirewallSessionTimerProfileBindingMapParam (required)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Patch(domainIdParam string, groupIdParam string, firewallSessionTimerProfileBindingMapIdParam string, policyFirewallSessionTimerProfileBindingMapParam model.PolicyFirewallSessionTimerProfileBindingMap) error

    // API will update Firewall Session Timer Profile Binding Map
    //
    // @param domainIdParam DomainID (required)
    // @param groupIdParam Group ID (required)
    // @param firewallSessionTimerProfileBindingMapIdParam Firewall Session Timer Profile Binding Map ID (required)
    // @param policyFirewallSessionTimerProfileBindingMapParam (required)
    // @return com.vmware.nsx_global_policy.model.PolicyFirewallSessionTimerProfileBindingMap
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Update(domainIdParam string, groupIdParam string, firewallSessionTimerProfileBindingMapIdParam string, policyFirewallSessionTimerProfileBindingMapParam model.PolicyFirewallSessionTimerProfileBindingMap) (model.PolicyFirewallSessionTimerProfileBindingMap, error)
}
