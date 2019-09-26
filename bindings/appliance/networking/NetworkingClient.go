/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Networking
 * Used by client-side stubs.
 */

package networking

import (
)

// The ``Networking`` interface provides methods Get Network configurations.
type NetworkingClient interface {


    // Get Networking information for all configured interfaces.
    // @return The Map of network configuration info for all interfaces.
    // @throws Error Generic error.
    Get() (Info, error) 


    // Enable or Disable ipv6 on all interfaces
    //
    // @param specParam update spec with optional boolean value
    // @throws Error Generic error.
    Update(specParam UpdateSpec) error 


    // Reset and restarts network configuration on all interfaces, also this will renew the DHCP lease for DHCP IP address.
    // @throws Error Generic error.
    Reset() error 


    // Changes the Hostname/IP of the management network of vCenter appliance. The Hostname/IP change invokes the PNID change process which involves LDAP entry modification, updating registry entries, configuration files modification and network configuration changes. vCenter server is expected to be down for few minutes during these changes
    //
    // @param specParam Information required to change the hostname.
    // @throws Unsupported if it's not embedded node
    // @throws InvalidArgument if passed arguments are invalid.
    // @throws Unauthenticated if the user is not authenticated.
    // @throws NotAllowedInCurrentState if another change task is in progress
    Change(specParam ChangeSpec) error 

}
