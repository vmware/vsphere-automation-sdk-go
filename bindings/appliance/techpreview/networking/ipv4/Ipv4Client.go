/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Ipv4
 * Used by client-side stubs.
 */

package ipv4

import (
)

// ``Ipv4`` interface provides methods Performs IPV4 network configuration for interfaces.
type Ipv4Client interface {


    // Renew IPv4 network configuration on interfaces. If the interface is configured to use DHCP for IP address assignment, the lease of the interface is renewed.
    //
    // @param interfacesParam Interfaces to renew.
    // @throws Error Generic error
    Renew(interfacesParam []string) error 


    // Set IPv4 network configuration.
    //
    // @param configParam List of IPv4 configurations.
    // @throws Error Generic error
    Set(configParam []IPv4Config) error 


    // Get IPv4 network configuration for all configured interfaces.
    // @return IPv4 configuration for each interface.
    // @throws Error Generic error
    List() ([]IPv4ConfigReadOnly, error) 


    // Get IPv4 network configuration for interfaces.
    //
    // @param interfacesParam Network interfaces to query, for example, "nic0".
    // @return IPv4 configuration for each queried interface.
    // @throws Error Generic error
    Get(interfacesParam []string) ([]IPv4ConfigReadOnly, error) 

}
