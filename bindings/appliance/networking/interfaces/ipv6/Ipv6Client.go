/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Ipv6
 * Used by client-side stubs.
 */

package ipv6

import (
)

// The ``Ipv6`` interface provides methods to perform IPv6 network configuration for interfaces.
type Ipv6Client interface {


    // Set IPv6 network configuration for specific interface.
    //
    // @param interfaceNameParam Network interface to update, for example, "nic0".
    // The parameter must be an identifier for the resource type: ``com.vmware.appliance.networking.interfaces``.
    // @param configParam The IPv6 configuration.
    // @throws ResourceBusy The specified NIC is busy.
    // @throws NotFound The specified NIC is not found.
    // @throws Error Generic error.
    Set(interfaceNameParam string, configParam Config) error 


    // Get IPv6 network configuration for specific interface.
    //
    // @param interfaceNameParam Network interface to query, for example, "nic0".
    // The parameter must be an identifier for the resource type: ``com.vmware.appliance.networking.interfaces``.
    // @return IPv6 configuration.
    // @throws NotFound if the specified NIC is not found.
    // @throws Error Generic error.
    Get(interfaceNameParam string) (Info, error) 

}
