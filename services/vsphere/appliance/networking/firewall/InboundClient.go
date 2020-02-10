/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Inbound
 * Used by client-side stubs.
 */

package firewall


// The ``Inbound`` interface provides methods to manage inbound firewall rules. This interface was added in vSphere API 6.7.1.
type InboundClient interface {

    // Set the ordered list of firewall rules to allow or deny traffic from one or more incoming IP addresses. This overwrites the existing firewall rules and creates a new rule list. Within the list of traffic rules, rules are processed in order of appearance, from top to bottom. For example, the list of rules can be as follows: 
    //
    // +------------+--------+----------------+--------+
    // | Address    | Prefix | Interface Name | Policy |
    // +============+========+================+========+
    // | 10.112.0.1 | 0      | \*             | REJECT |
    // +------------+--------+----------------+--------+
    // | 10.112.0.1 | 0      | nic0           | ACCEPT |
    // +------------+--------+----------------+--------+
    //  In the above example, the first rule drops all packets originating from 10.112.0.1 and
    //  the second rule accepts all packets originating from 10.112.0.1 only on nic0. In effect, the second rule is always ignored which is not desired, hence the order has to be swapped. When a connection matches a firewall rule, further processing for the connection stops, and the appliance ignores any additional firewall rules you have set. This method was added in vSphere API 6.7.1.
    //
    // @param rulesParam List of address-based firewall rules.
    // @throws Error Generic error
	Set(rulesParam []InboundRule) error

    // Get the ordered list of firewall rules. Within the list of traffic rules, rules are processed in order of appearance, from top to bottom. When a connection matches a firewall rule, further processing for the connection stops, and the appliance ignores any additional firewall rules you have set. This method was added in vSphere API 6.7.1.
    // @return List of address-based firewall rules.
    // @throws Error Generic error
	Get() ([]InboundRule, error)
}
