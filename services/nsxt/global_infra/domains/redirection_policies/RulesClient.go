/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Rules
 * Used by client-side stubs.
 */

package redirection_policies

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt/model"
)

type RulesClient interface {

    // Delete RedirectionRule
    //
    // @param domainIdParam Domain ID (required)
    // @param redirectionPolicyIdParam Redirection Map ID (required)
    // @param ruleIdParam RedirectionRule ID (required)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Delete(domainIdParam string, redirectionPolicyIdParam string, ruleIdParam string) error

    // Read rule
    //
    // @param domainIdParam Domain id (required)
    // @param redirectionPolicyIdParam Redirection map id (required)
    // @param ruleIdParam Rule id (required)
    // @return com.vmware.nsx_policy.model.RedirectionRule
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get(domainIdParam string, redirectionPolicyIdParam string, ruleIdParam string) (model.RedirectionRule, error)

    // List rules
    //
    // @param domainIdParam Domain id (required)
    // @param redirectionPolicyIdParam Redirection map id (required)
    // @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
    // @param includeMarkForDeleteObjectsParam Include objects that are marked for deletion in results (optional, default to false)
    // @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
    // @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
    // @param sortAscendingParam (optional)
    // @param sortByParam Field by which records are sorted (optional)
    // @return com.vmware.nsx_policy.model.RedirectionRuleListResult
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	List(domainIdParam string, redirectionPolicyIdParam string, cursorParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.RedirectionRuleListResult, error)

    // Create a rule with the rule-id is not already present, otherwise update the rule.
    //
    // @param domainIdParam Domain id (required)
    // @param redirectionPolicyIdParam RedirectionPolicy id (required)
    // @param ruleIdParam rule id (required)
    // @param redirectionRuleParam (required)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Patch(domainIdParam string, redirectionPolicyIdParam string, ruleIdParam string, redirectionRuleParam model.RedirectionRule) error

    // Create a rule with the rule-id is not already present, otherwise update the rule.
    //
    // @param domainIdParam Domain id (required)
    // @param redirectionPolicyIdParam Redirection map id (required)
    // @param ruleIdParam Rule id (required)
    // @param redirectionRuleParam (required)
    // @return com.vmware.nsx_policy.model.RedirectionRule
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Update(domainIdParam string, redirectionPolicyIdParam string, ruleIdParam string, redirectionRuleParam model.RedirectionRule) (model.RedirectionRule, error)
}
