/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Ipv6
 * Used by client-side stubs.
 */

package interfaces


// The ``Ipv6`` interface provides methods to perform IPv6 network configuration for interfaces. This interface was added in vSphere API 6.7.
type Ipv6Client interface {

    // Set IPv6 network configuration for specific interface. This method was added in vSphere API 6.7.
    //
    // @param interfaceNameParam Network interface to update, for example, "nic0".
    // The parameter must be an identifier for the resource type: ``com.vmware.appliance.networking.interfaces``.
    // @param configParam The IPv6 configuration.
    // @throws ResourceBusy The specified NIC is busy.
    // @throws NotFound The specified NIC is not found.
    // @throws Error Generic error.
	Set(interfaceNameParam string, configParam Ipv6Config) error

    // Get IPv6 network configuration for specific interface. This method was added in vSphere API 6.7.
    //
    // @param interfaceNameParam Network interface to query, for example, "nic0".
    // The parameter must be an identifier for the resource type: ``com.vmware.appliance.networking.interfaces``.
    // @return IPv6 configuration.
    // @throws NotFound if the specified NIC is not found.
    // @throws Error Generic error.
	Get(interfaceNameParam string) (Ipv6Info, error)
}
