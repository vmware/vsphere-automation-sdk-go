/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Interfaces
 * Used by client-side stubs.
 */

package networking


// ``Interfaces`` interface provides methods Provides information about network interface.
type InterfacesClient interface {

    // Get list of available network interfaces, including those that are not yet configured.
    // @return List of InterfaceInfo structures.
    // @throws Error Generic error
	List() ([]InterfacesInterfaceInfo, error)

    // Get information about a particular network interface.
    //
    // @param interfaceNameParam Network interface, for example, "nic0".
    // The parameter must be an identifier for the resource type: ``com.vmware.appliance.networking.interfaces``.
    // @return Network interface information.
    // @throws NotFound if the specified interface is not found.
    // @throws Error Generic error
	Get(interfaceNameParam string) (InterfacesInterfaceInfo, error)
}
