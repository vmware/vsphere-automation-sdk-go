/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Servers
 * Used by client-side stubs.
 */

package servers

import (
)

// ``Servers`` interface provides methods DNS server configuration.
type ServersClient interface {


    // Test if dns servers are reachable.
    //
    // @param serversParam DNS servers.
    // @return DNS reacable status
    // @throws Error Generic error
    Test(serversParam []string) (TestStatusInfo, error) 


    // Add a DNS server. This method fails if mode argument is "dhcp"
    //
    // @param serverParam DNS server.
    // @throws Error Generic error
    Add(serverParam string) error 


    // Set the DNS server configuration. If you set the mode argument to "DHCP", a DHCP refresh is forced.
    //
    // @param configParam DNS server configuration.
    // @throws Error Generic error
    Set(configParam DNSServerConfig) error 


    // Get DNS server configuration.
    // @return DNS server configuration.
    // @throws Error Generic error
    Get() (DNSServerConfig, error) 

}
