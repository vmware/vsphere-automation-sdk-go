/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: EndpointRules
 * Used by client-side stubs.
 */

package endpoint_policies

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt/model"
)

type EndpointRulesClient interface {

    // Delete EndpointRule
    //
    // @param domainIdParam Domain ID (required)
    // @param endpointPolicyIdParam EndpointPolicy ID (required)
    // @param endpointRuleIdParam EndpointRule ID (required)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Delete(domainIdParam string, endpointPolicyIdParam string, endpointRuleIdParam string) error

    // Read Endpoint rule
    //
    // @param domainIdParam Domain id (required)
    // @param endpointPolicyIdParam Endpoint policy id (required)
    // @param endpointRuleIdParam Endpoint rule id (required)
    // @return com.vmware.nsx_policy.model.EndpointRule
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get(domainIdParam string, endpointPolicyIdParam string, endpointRuleIdParam string) (model.EndpointRule, error)

    // List Endpoint rules
    //
    // @param domainIdParam Domain id (required)
    // @param endpointPolicyIdParam Endpoint policy id (required)
    // @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
    // @param includeMarkForDeleteObjectsParam Include objects that are marked for deletion in results (optional, default to false)
    // @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
    // @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
    // @param sortAscendingParam (optional)
    // @param sortByParam Field by which records are sorted (optional)
    // @return com.vmware.nsx_policy.model.EndpointRuleListResult
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	List(domainIdParam string, endpointPolicyIdParam string, cursorParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.EndpointRuleListResult, error)

    // Create a Endpoint rule with the endpoint-rule-id is not already present, otherwise update the Endpoint Rule.
    //
    // @param domainIdParam Domain id (required)
    // @param endpointPolicyIdParam Endpoint policy id (required)
    // @param endpointRuleIdParam Endpoint rule id (required)
    // @param endpointRuleParam (required)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Patch(domainIdParam string, endpointPolicyIdParam string, endpointRuleIdParam string, endpointRuleParam model.EndpointRule) error

    // Create a Endpoint rule with the endpoint-rule-id is not already present, otherwise update the Endpoint Rule.
    //
    // @param domainIdParam Domain id (required)
    // @param endpointPolicyIdParam Endpoint policy id (required)
    // @param endpointRuleIdParam Endpoint rule id (required)
    // @param endpointRuleParam (required)
    // @return com.vmware.nsx_policy.model.EndpointRule
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Update(domainIdParam string, endpointPolicyIdParam string, endpointRuleIdParam string, endpointRuleParam model.EndpointRule) (model.EndpointRule, error)
}
