/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Inbound
 * Used by client-side stubs.
 */

package inbound

import (
)

// ``Inbound`` interface provides methods Operations for firewall rules.
type InboundClient interface {


    // Add a firewall rule to allow or deny traffic from incoming IP address.
    //
    // @param posParam Position before which to insert the rule (zero-based). If you try to insert the rule in a position whose number is greater than the number of rules, the firewall rule is inserted at the end of the list.
    // @param ruleParam Firewall IP-based rule.
    // @throws Error Generic error
    Add(posParam int64, ruleParam FirewallAddressRule) error 


    // Set list of inbound IP addresses to allow or deny by firewall. This replaces all existing rules. Firewall rules have no impact on closed ports because these ports are closed for all traffic.
    //
    // @param rulesParam List of address-based firewall rules.
    // @throws Error Generic error
    Set(rulesParam []FirewallAddressRule) error 


    // Get ordered list of inbound IP addresses that are allowed or denied by firewall.
    // @return List of address-based firewall rules.
    // @throws Error Generic error
    List() ([]FirewallAddressRule, error) 


    // Delete specific rule at a given position or delete all rules.
    //
    // @param configParam Delete a firewall rule
    // @throws Error Generic error
    Delete(configParam DeleteFirewallRule) error 

}
