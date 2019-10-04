/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Community
 * Used by client-side stubs.
 */

package community

import (
)

// The ``Community`` interface provides methods to add and remove nodes to a community. **Warning:** This interface is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type CommunityClient interface {


    // Information about the community. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    // @return information about the peers of a community
    // @throws Unauthorized If the caller is not authorized.
    // @throws Error If the system reports an error while responding to the request.
    Get() (Info, error) 


    // Adds the target node to the community. By default, the first member of the community id the local node. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param specParam Specification for adding a node to the community.
    // @throws AlreadyExists If the local and target machines are already in the same community.
    // @throws InvalidArgument If an argument is not valid.
    // @throws Unauthenticated If unable to authenticate to target node.
    // @throws Unauthorized If the caller is not authorized.
    // @throws UnverifiedPeer If the SSL certificate of the target node cannot be validated by comparing with the thumbprint provided in AddSpec#sslThumbprint or if AddSpec#sslThumbprint is null and AddSpec#sslVerify is true. The value of the {\\\\@link UnverifiedPeer#data) property will be a class that contains all the properties defined in CertificateInfo.
    // @throws Error If the system reports an error while responding to the request.
    Add(specParam AddSpec) error 


    // Removes a specified node from a community. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param hostnameParam The hostname of the node to remove.
    // @throws Unauthorized If the caller is not authorized.
    // @throws Error If the system reports an error while responding to the request.
    Remove(hostnameParam string) error 


    // Performs network checks between nodes in the community. This can be called before the add to validate network requirements prior to performing the action or after adding to check network connectivity within a peers of a community. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param specParam Specification for defining how to perform a network check.
    // If null, checks will be done between all pairs of nodes in the community.
    // @return list of CheckInfo objects. Contains results aggregated per pair of nodes.
    // @throws Unauthenticated If unable to authenticate to target node.
    // @throws Unauthorized If the caller is not authorized.
    // @throws UnverifiedPeer If the SSL certificate of the target node cannot be validated by comparing with the thumbprint provided in AddSpec#sslThumbprint or if AddSpec#sslThumbprint is null and AddSpec#sslVerify is true. The value of the {\\\\@link UnverifiedPeer#data) property will be a class that contains all the properties defined in CertificateInfo.
    // @throws Error If the system reports an error while responding to the request.
    Check(specParam *AddCheckSpec) ([]CheckInfo, error) 

}
