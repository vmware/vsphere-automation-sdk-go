/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Recommendation
 * Used by client-side stubs.
 */

package recommendation

import (
)

// The ``Recommendation`` interface provides methods to give scale in (to remove host) / scale out (to add host) recommendation for a cluster. **Warning:** This interface is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type RecommendationClient interface {


    // Gives a recommendation to scale in a cluster when all resources (CPU/memory/storage) are underutilized, or to scale out a cluster when any resource is being heavily utilized. NO_ACTION will be returned if cluster resources are well utilized. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param clusterParam Identifier of cluster that requests EDRS recommenation.
    // The parameter must be an identifier for the resource type: ``ClusterComputeResource``.
    // @return an EDRS recommendation.
    // @throws NotFound if the cluster is not known.
    // @throws Unauthorized if the user doesn't have the required privileges.
    Generate(clusterParam string) (Recommendation, error) 

}
