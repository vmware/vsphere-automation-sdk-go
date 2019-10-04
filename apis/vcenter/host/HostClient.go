/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Host
 * Used by client-side stubs.
 */

package host

import (
)

// The ``Host`` interface provides methods to manage hosts in the vCenter Server.
type HostClient interface {


    // Add a new standalone host in the vCenter inventory. The newly connected host will be in connected state. The vCenter Server will verify the SSL certificate before adding the host to its inventory. In the case where the SSL certificate cannot be verified because the Certificate Authority is not recognized or the certificate is self signed, the vCenter Server will fall back to thumbprint verification mode as defined by CreateSpec_ThumbprintVerification.
    //
    // @param specParam Specification for the new host to be created.
    // @return The newly created identifier of the host in vCenter.
    // The return value will be an identifier for the resource type: ``HostSystem``.
    // @throws AlreadyExists if the host with the same name is already present.
    // @throws Error if installation of VirtualCenter agent on a host fails.
    // @throws Error if the system reports an error while responding to the request.
    // @throws InvalidArgument if the host name is invalid.
    // @throws InvalidArgument if the host folder is invalid.
    // @throws InvalidArgument if the SSL thumbprint specified is invalid.
    // @throws InvalidElementType if the host folder id does not support vSphere compute resource as its children type.
    // @throws NotFound if there is no folder associated with the ``folder`` property in the system.
    // @throws ResourceInUse if the host is already being managed by another vCenter Server
    // @throws UnableToAllocateResource if there are not enough licenses to add the host.
    // @throws Unauthenticated if the user name or password for the administration account on the host are invalid.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unsupported if the software version on the host is not supported.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthorized if the user doesn't have the required privileges.
    Create(specParam CreateSpec) (string, error) 


    // Remove a standalone host from the vCenter Server.
    //
    // @param hostParam Identifier of the host to be deleted.
    // The parameter must be an identifier for the resource type: ``HostSystem``.
    // @throws Error if the system reports an error while responding to the request.
    // @throws NotFound if there is no host associated with ``host`` in the system.
    // @throws ResourceInUse if the host associated with ``host`` is in a vCenter cluster
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
    Delete(hostParam string) error 


    // Returns information about at most 2500 visible (subject to permission checks) hosts in vCenter matching the FilterSpec.
    //
    // @param filterParam Specification of matching hosts for which information should be returned.
    // If null, the behavior is equivalent to a FilterSpec with all properties null which means all hosts match the filter.
    // @return Commonly used information about the hosts matching the FilterSpec.
    // @throws InvalidArgument if the FilterSpec#connectionStates property contains a value that is not supported by the server.
    // @throws UnableToAllocateResource if more than 2500 hosts match the FilterSpec.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
    List(filterParam *FilterSpec) ([]Summary, error) 


    // Connect to the host corresponding to ``host`` previously added to the vCenter server.
    //
    // @param hostParam Identifier of the host to be reconnected.
    // The parameter must be an identifier for the resource type: ``HostSystem``.
    // @throws AlreadyInDesiredState if the host associated with ``host`` is already connected.
    // @throws Error if the system reports an error while responding to the request.
    // @throws NotFound if there is no host associated with ``host`` in the system.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
    Connect(hostParam string) error 


    // Disconnect the host corresponding to ``host`` from the vCenter server
    //
    // @param hostParam Identifier of the host to be disconnected.
    // The parameter must be an identifier for the resource type: ``HostSystem``.
    // @throws AlreadyInDesiredState if the host associated with ``host`` is already disconnected.
    // @throws Error if the system reports an error while responding to the request.
    // @throws NotFound if there is no host associated with ``host`` in the system.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
    Disconnect(hostParam string) error 

}
