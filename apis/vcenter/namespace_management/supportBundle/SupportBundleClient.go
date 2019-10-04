/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: SupportBundle
 * Used by client-side stubs.
 */

package supportBundle

import (
)

// The ``SupportBundle`` interface provides methods to retrieve the cluster's Namespaces-related support bundle download location. **Warning:** This interface is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type SupportBundleClient interface {


    // Returns the location Location information for downloading the Namespaces-related support bundle for the specified cluster. 
    //
    //  Retrieving a support bundle involves two steps: 
    //
    // * Step 1: Invoke method to provision a token and a URI.
    // * Step 2: Make an HTTP GET request using URI and one time used token returned in step 1 to generate the support bundle and return it.
    //
    //  There can only be one valid token per cluster at any given time. If this method is invoked again for the same cluster identifier while a token still valid, the API will return the same Location response. 
    //
    //  The HTTP GET request will: 
    //
    // * return 401 (Not Authorized) if the download URL is recognized, but the token is invalid.
    // * otherwise return 200 (OK), mark the token used (invalidating it for any future use), open a application/tar download stream for the client, and start the bundle process. As part of its work, the API will orchestrate support bundling on the worker nodes of a cluster. If a failure occurs during the collection of support bundle from worker node, the API will not abort the request, but will log a warning and move on to processing other worker nodes' bundles. The API will only abort its operation if the content of the stream has been corrupted. When the API has to abort its operation (and the response stream), it will not provide any indication of failures to the client. The client will need to verify validity of the resultant file based on the format specified in the response's Content-Disposition header.
    //
    // . **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param clusterParam Identifier of cluster for which the Namespaces-related support bundle should be generated.
    // The parameter must be an identifier for the resource type: ``ClusterComputeResource``.
    // @return the download location of the support bundle for the cluster.
    // @throws NotFound if the specified cluster is not registered on this vCenter server.
    // @throws Error if the system reports an error while responding to the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user does not have System.Read privilege.
    Create(clusterParam string) (Location, error) 

}
