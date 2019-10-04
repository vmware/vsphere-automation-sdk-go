/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: AlgoParameter
 * Used by client-side stubs.
 */

package algoParameter

import (
)

// The ``AlgoParameter`` interface provides methods to set/get EDRS algorithm parameters for a cluster. **Warning:** This interface is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type AlgoParameterClient interface {


    // Sets EDRS algorithm parameters for a given cluster. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param clusterParam Identifier of cluster whose EDRS need to be configured.
    // The parameter must be an identifier for the resource type: ``ClusterComputeResource``.
    // @param paramsParam EDRS algorithm parameters need to be set.
    // @throws Error if input parameters have issues.
    // @throws NotFound if the cluster is unknown.
    // @throws Unauthorized if the user doesn't have the required privileges.
    Set(clusterParam string, paramsParam Param) error 


    // Gets EDRS algorithm parameters for a given cluster. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param clusterParam Identifier of cluster whose EDRS algorithm parameters need to be returned.
    // The parameter must be an identifier for the resource type: ``ClusterComputeResource``.
    // @return EDRS algorithm parameters for the given cluster.
    // @throws NotFound if the cluster is unknown.
    // @throws Unauthorized if the user doesn't have the required privileges.
    Get(clusterParam string) (Param, error) 

}
