/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Rules
 * Used by client-side stubs.
 */

package config

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/services/vmc/model"
)

type RulesClient interface {

    // Append firewall rules for a management or compute gateway (NSX Edge).
    //
    // @param orgParam Organization identifier. (required)
    // @param sddcParam Sddc Identifier. (required)
    // @param edgeIdParam Edge Identifier. (required)
    // @param firewallRulesParam (required)
    // @throws InvalidRequest  Bad request. Request object passed is invalid.
    // @throws Unauthorized  Forbidden. Authorization header not provided.
    // @throws NotFound  Not found. Requested object not found.
	Add(orgParam string, sddcParam string, edgeIdParam string, firewallRulesParam model.FirewallRules) error

    // Delete a specific firewall rule for a management or compute gateway (NSX Edge).
    //
    // @param orgParam Organization identifier. (required)
    // @param sddcParam Sddc Identifier. (required)
    // @param edgeIdParam Edge Identifier. (required)
    // @param ruleIdParam Rule Identifier. (required)
    // @throws InvalidRequest  Bad request. Request object passed is invalid.
    // @throws Unauthorized  Forbidden. Authorization header not provided.
    // @throws NotFound  Not found. Requested object not found.
	Delete(orgParam string, sddcParam string, edgeIdParam string, ruleIdParam int64) error

    // Retrieve a specific firewall rule for a management or compute gateway (NSX Edge).
    //
    // @param orgParam Organization identifier. (required)
    // @param sddcParam Sddc Identifier. (required)
    // @param edgeIdParam Edge Identifier. (required)
    // @param ruleIdParam Rule Identifier. (required)
    // @return com.vmware.vmc.model.Nsxfirewallrule
    // @throws InvalidRequest  Bad request. Request object passed is invalid.
    // @throws Unauthorized  Forbidden. Authorization header not provided
    // @throws NotFound  Not found. Requested object not found.
	Get(orgParam string, sddcParam string, edgeIdParam string, ruleIdParam int64) (model.Nsxfirewallrule, error)

    // Modify the specified firewall rule for a management or compute gateway (NSX Edge).
    //
    // @param orgParam Organization identifier. (required)
    // @param sddcParam Sddc Identifier. (required)
    // @param edgeIdParam Edge Identifier. (required)
    // @param ruleIdParam Rule Identifier. (required)
    // @param nsxfirewallruleParam (required)
    // @throws InvalidRequest  Bad request. Request object passed is invalid.
    // @throws Unauthorized  Forbidden. Authorization header not provided.
    // @throws NotFound  Not found. Requested object not found.
	Update(orgParam string, sddcParam string, edgeIdParam string, ruleIdParam int64, nsxfirewallruleParam model.Nsxfirewallrule) error
}
