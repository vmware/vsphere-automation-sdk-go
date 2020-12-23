/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Rules
 * Used by client-side stubs.
 */

package forwarding_policies

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

type RulesClient interface {

    // Delete ForwardingRule
    //
    // @param domainIdParam Domain ID (required)
    // @param forwardingPolicyIdParam Forwarding Map ID (required)
    // @param ruleIdParam ForwardingRule ID (required)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Delete(domainIdParam string, forwardingPolicyIdParam string, ruleIdParam string) error

    // Read rule
    //
    // @param domainIdParam Domain id (required)
    // @param forwardingPolicyIdParam Forwarding map id (required)
    // @param ruleIdParam Rule id (required)
    // @return com.vmware.nsx_policy.model.ForwardingRule
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get(domainIdParam string, forwardingPolicyIdParam string, ruleIdParam string) (model.ForwardingRule, error)

    // List rules
    //
    // @param domainIdParam Domain id (required)
    // @param forwardingPolicyIdParam Forwarding map id (required)
    // @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
    // @param includeMarkForDeleteObjectsParam Include objects that are marked for deletion in results (optional, default to false)
    // @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
    // @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
    // @param sortAscendingParam (optional)
    // @param sortByParam Field by which records are sorted (optional)
    // @return com.vmware.nsx_policy.model.ForwardingRuleListResult
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	List(domainIdParam string, forwardingPolicyIdParam string, cursorParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.ForwardingRuleListResult, error)

    //
    //
    // @param domainIdParam Domain id (required)
    // @param forwardingPolicyIdParam Forwarding map id (required)
    // @param ruleIdParam Rule id (required)
    // @param forwardingRuleParam (required)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Patch(domainIdParam string, forwardingPolicyIdParam string, ruleIdParam string, forwardingRuleParam model.ForwardingRule) error

    //
    //
    // @param domainIdParam Domain id (required)
    // @param forwardingPolicyIdParam Forwarding map id (required)
    // @param ruleIdParam rule id (required)
    // @param forwardingRuleParam (required)
    // @return com.vmware.nsx_policy.model.ForwardingRule
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Update(domainIdParam string, forwardingPolicyIdParam string, ruleIdParam string, forwardingRuleParam model.ForwardingRule) (model.ForwardingRule, error)
}
