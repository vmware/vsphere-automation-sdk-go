/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Networking
 * Used by client-side stubs.
 */

package appliance


// The ``Networking`` interface provides methods Get Network configurations. This interface was added in vSphere API 6.7.
type NetworkingClient interface {

    // Get Networking information for all configured interfaces. This method was added in vSphere API 6.7.
    // @return The Map of network configuration info for all interfaces.
    // @throws Error Generic error.
	Get() (NetworkingInfo, error)

    // Enable or Disable ipv6 on all interfaces. This method was added in vSphere API 6.7.
    //
    // @param specParam update spec with optional boolean value
    // @throws Error Generic error.
	Update(specParam NetworkingUpdateSpec) error

    // Reset and restarts network configuration on all interfaces, also this will renew the DHCP lease for DHCP IP address. This method was added in vSphere API 6.7.
    // @throws Error Generic error.
	Reset() error

    // Changes the Hostname/IP of the management network of vCenter appliance. The Hostname/IP change invokes the PNID change process which involves LDAP entry modification, updating registry entries, configuration files modification and network configuration changes. vCenter server is expected to be down for few minutes during these changes. This method was added in vSphere API 6.7.3.
    //
    // @param specParam Information required to change the hostname.
    // @throws Unsupported if it's not embedded node
    // @throws InvalidArgument if passed arguments are invalid.
    // @throws Unauthenticated if the user is not authenticated.
    // @throws NotAllowedInCurrentState if another change task is in progress
	Change(specParam NetworkingChangeSpec) error
}
