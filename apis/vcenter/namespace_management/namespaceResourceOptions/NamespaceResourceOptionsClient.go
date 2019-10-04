/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: NamespaceResourceOptions
 * Used by client-side stubs.
 */

package namespaceResourceOptions

import (
)

// The ``NamespaceResourceOptions`` interface provides methods to get the objects used to create and modify resource quotas on a namespace. **Warning:** This interface is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type NamespaceResourceOptionsClient interface {


    // Get the information about the objects used to set and update resource quota keys for version of Kubernetes running on {#link cluster}. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param clusterParam Identifier for the cluster hosting the namespace on which the resource quota needs to be set.
    // The parameter must be an identifier for the resource type: ``ClusterComputeResource``.
    // @return Information about the structures representing the resource spec.
    // @throws NotFound if cluster could not be located.
    // @throws Unsupported if the specified cluster is not enabled for Namespaces.
    // @throws Error if the system reports an error while responding to the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user does not have System.Read privilege.
    Get(clusterParam string) (Info, error) 

}
