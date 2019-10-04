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

// ``Ipv6`` interface provides methods Performs IPV4 network configuration for interfaces.
type Ipv6Client interface {


    // Set IPv6 network configuration.
    //
    // @param configParam IPv6 configuration.
    // @throws Error Generic error
    Set(configParam []IPv6Config) error 


    // Get IPv6 network configuration for all configured interfaces.
    // @return IPv6 configuration for each interface.
    // @throws Error Generic error
    List() ([]IPv6ConfigReadOnly, error) 


    // Get IPv6 network configuration for interfaces.
    //
    // @param interfacesParam Network interfaces to query, for example, "nic0".
    // @return IPv6 configuration.
    // @throws Error Generic error
    Get(interfacesParam []string) ([]IPv6ConfigReadOnly, error) 

}
