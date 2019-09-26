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

// The ``Nodes`` interface provides methods to manage the state of Kubernetes nodes of a Namespaces cluster. **Warning:** This interface is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type NodesClient interface {


    // Enter Namespaces maintenance mode on the node. This operation removes Kubernetes pods from the given node and prevents further scheduling of pods on the node. This operation can be cancelled by calling exitMaintenanceMode operation on the node. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param clusterParam Identifier for the cluster.
    // The parameter must be an identifier for the resource type: ``ClusterComputeResource``.
    // @param nodeParam Identity for the node.
    // @param actionParam Optional action to be taken when node enters maintenance mode. MaintenanceActionType#MaintenanceActionType_NO_ACTION is used as the default action type when this parameter is not provided.
    // @return The task identifier for the operation. The task is not cancellable.
    // The return value will be an identifier for the resource type: ``Task``.
    // @throws Error if the system reports an error while responding to the request.
    // @throws AlreadyInDesiredState if ``node`` is already in the desired state.
    // @throws NotFound if \\\\@param.name cluster} is not enabled for Namespaces.
    // @throws Unsupported if NodeIdentity#nodeType is not of type
    // @throws InvalidElementConfiguration if ``node`` is not a member of ``cluster``.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user is not vpxd solution user.
    EnterMaintenanceMode(clusterParam string, nodeParam NodeIdentity, actionParam *MaintenanceActionType) (string, error) 


    // Exit Namespaces maintenance mode on the node. This operation enables scheduling of Kubernetes pods on the node. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param clusterParam Identifier for the cluster.
    // The parameter must be an identifier for the resource type: ``ClusterComputeResource``.
    // @param nodeParam Identity for the node.
    // @return The task identifier for the operation. The task is not cancellable.
    // The return value will be an identifier for the resource type: ``Task``.
    // @throws Error if the system reports an error while responding to the request.
    // @throws AlreadyInDesiredState if ``node`` is already in the desired state.
    // @throws NotFound if \\\\@param.name cluster} is not enabled for Namespaces.
    // @throws Unsupported if NodeIdentity#nodeType is not of type
    // @throws InvalidElementConfiguration if ``node`` is not a member of ``cluster``.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user is not vpxd solution user.
    ExitMaintenanceMode(clusterParam string, nodeParam NodeIdentity) (string, error) 


    // Remove Namespaces state from node exiting Namespaces managed cluster. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param clusterParam Identifier for the cluster.
    // The parameter must be an identifier for the resource type: ``ClusterComputeResource``.
    // @param nodeParam Identity for the node.
    // @return The task identifier for the operation. The task is not cancellable.
    // The return value will be an identifier for the resource type: ``Task``.
    // @throws Error if the system reports an error while responding to the request.
    // @throws AlreadyInDesiredState if ``node`` is already in the desired state.
    // @throws NotFound if \\\\@param.name cluster} is not enabled for Namespaces.
    // @throws Unsupported if NodeIdentity#nodeType is not of type
    // @throws InvalidElementConfiguration if ``node`` is not a member of ``cluster``.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user is not vpxd solution user.
    Remove(clusterParam string, nodeParam NodeIdentity) (string, error) 

}
