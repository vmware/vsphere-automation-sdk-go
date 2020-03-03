/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Rules
 * Used by client-side stubs.
 */

package intrusion_service_policies

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

type RulesClient interface {

    // Delete intrusion detection rule.
    //
    // @param domainIdParam Domain ID (required)
    // @param policyIdParam Policy ID (required)
    // @param ruleIdParam Rule ID (required)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Delete(domainIdParam string, policyIdParam string, ruleIdParam string) error

    // Read intrusion detection rule
    //
    // @param domainIdParam Domain ID (required)
    // @param policyIdParam Policy ID (required)
    // @param ruleIdParam Rule ID (required)
    // @return com.vmware.nsx_policy.model.IdsRule
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get(domainIdParam string, policyIdParam string, ruleIdParam string) (model.IdsRule, error)

    // List intrusion detection rules.
    //
    // @param domainIdParam Domain ID (required)
    // @param policyIdParam Policy ID (required)
    // @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
    // @param includeMarkForDeleteObjectsParam Include objects that are marked for deletion in results (optional, default to false)
    // @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
    // @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
    // @param sortAscendingParam (optional)
    // @param sortByParam Field by which records are sorted (optional)
    // @return com.vmware.nsx_policy.model.IdsRuleListResult
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	List(domainIdParam string, policyIdParam string, cursorParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.IdsRuleListResult, error)

    // Patch intrusion detection system rule.
    //
    // @param domainIdParam Domain ID (required)
    // @param policyIdParam Policy ID (required)
    // @param ruleIdParam Rule ID (required)
    // @param idsRuleParam (required)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Patch(domainIdParam string, policyIdParam string, ruleIdParam string, idsRuleParam model.IdsRule) error

    // This is used to re-order a rule within a security policy.
    //
    // @param domainIdParam (required)
    // @param policyIdParam (required)
    // @param ruleIdParam (required)
    // @param idsRuleParam (required)
    // @param anchorPathParam The security policy/rule path if operation is 'insert_after' or 'insert_before' (optional)
    // @param operationParam Operation (optional, default to insert_top)
    // @return com.vmware.nsx_policy.model.IdsRule
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Revise(domainIdParam string, policyIdParam string, ruleIdParam string, idsRuleParam model.IdsRule, anchorPathParam *string, operationParam *string) (model.IdsRule, error)

    // Update intrusion detection system rule.
    //
    // @param domainIdParam Domain ID (required)
    // @param policyIdParam Policy ID (required)
    // @param ruleIdParam Rule ID (required)
    // @param idsRuleParam (required)
    // @return com.vmware.nsx_policy.model.IdsRule
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Update(domainIdParam string, policyIdParam string, ruleIdParam string, idsRuleParam model.IdsRule) (model.IdsRule, error)
}
