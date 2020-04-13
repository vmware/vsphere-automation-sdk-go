/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Rules
 * Used by client-side stubs.
 */

package gateway_policies

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/services/nsxt-gm/model"
)

type RulesClient interface {

    // Delete rule
    //
    // @param domainIdParam (required)
    // @param gatewayPolicyIdParam (required)
    // @param ruleIdParam (required)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Delete(domainIdParam string, gatewayPolicyIdParam string, ruleIdParam string) error

    // Read rule
    //
    // @param domainIdParam (required)
    // @param gatewayPolicyIdParam (required)
    // @param ruleIdParam (required)
    // @return com.vmware.nsx_global_policy.model.Rule
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get(domainIdParam string, gatewayPolicyIdParam string, ruleIdParam string) (model.Rule, error)

    // List rules
    //
    // @param domainIdParam (required)
    // @param gatewayPolicyIdParam (required)
    // @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
    // @param includeMarkForDeleteObjectsParam Include objects that are marked for deletion in results (optional, default to false)
    // @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
    // @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
    // @param sortAscendingParam (optional)
    // @param sortByParam Field by which records are sorted (optional)
    // @return com.vmware.nsx_global_policy.model.RuleListResult
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	List(domainIdParam string, gatewayPolicyIdParam string, cursorParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.RuleListResult, error)

    // Update the gateway rule. Create new rule if a rule with the rule-id is not already present.
    //
    // @param domainIdParam (required)
    // @param gatewayPolicyIdParam (required)
    // @param ruleIdParam (required)
    // @param ruleParam (required)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Patch(domainIdParam string, gatewayPolicyIdParam string, ruleIdParam string, ruleParam model.Rule) error

    // This is used to re-order a gateway rule within a gateway policy.
    //
    // @param domainIdParam (required)
    // @param gatewayPolicyIdParam (required)
    // @param ruleIdParam (required)
    // @param ruleParam (required)
    // @param anchorPathParam The security policy/rule path if operation is 'insert_after' or 'insert_before' (optional)
    // @param operationParam Operation (optional, default to insert_top)
    // @return com.vmware.nsx_global_policy.model.Rule
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Revise(domainIdParam string, gatewayPolicyIdParam string, ruleIdParam string, ruleParam model.Rule, anchorPathParam *string, operationParam *string) (model.Rule, error)

    // Update the gateway rule. Create new rule if a rule with the rule-id is not already present.
    //
    // @param domainIdParam (required)
    // @param gatewayPolicyIdParam (required)
    // @param ruleIdParam (required)
    // @param ruleParam (required)
    // @return com.vmware.nsx_global_policy.model.Rule
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Update(domainIdParam string, gatewayPolicyIdParam string, ruleIdParam string, ruleParam model.Rule) (model.Rule, error)
}
