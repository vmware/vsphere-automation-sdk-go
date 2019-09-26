/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Nodes
 * Used by client-side stubs.
 */

package nodes

import (
)

// The ``Nodes`` interface provides methods to retrieve information about all linked vCenter Server nodes.
type NodesClient interface {


    // This operation retrieves detailed information about all linked vCenter Server nodes.
    // @return List of vCenter node details with. See VcNode.
    // @throws Unauthenticated  if the user can not be authenticated.
    // @throws Error  if there are any errors in retrieving the vCenter data.
    List() ([]VcNode, error) 

}
