/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Network
 * Used by client-side stubs.
 */

package inventory


// The ``Network`` interface provides methods to retrieve information about vCenter Server networks.
type NetworkClient interface {

    // Returns network information for the specified vCenter Server networks. The key in the return value map is the network identifier and the value in the map is the network information.
    //
    // @param networksParam  Identifiers of the vCenter Server networks for which information will be returned.
    // The parameter must contain identifiers for the resource type: ``Network``.
    // @return Network information for the specified vCenter Server networks. The key in the return value map is the network identifier and the value in the map is the network information.
    // The key in the return value map will be an identifier for the resource type: ``Network``.
    // @throws NotFound  if no datastore can be found for one or more of the vCenter Server network identifiers in ``networks``
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``System.Read``.
    // * The resource ``Network`` referenced by the parameter ``networks`` requires ``System.Read``.
	Find(networksParam []string) (map[string]*NetworkInfo, error)
}
