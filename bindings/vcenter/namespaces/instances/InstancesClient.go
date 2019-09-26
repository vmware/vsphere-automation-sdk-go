/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Instances
 * Used by client-side stubs.
 */

package instances

import (
)

// The ``Instances`` interface provides methods to create and delete a namespace object. In this version, an Instance is an abstraction around a Kubernetes namespace. **Warning:** This interface is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type InstancesClient interface {


    // Create namespace object in the cluster. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param specParam Specification for setting up the namespace.
    // @throws AlreadyExists if a namespace with the same name exists in vCenter server.
    // @throws Error if the system reports an error while responding to the request.
    // @throws InvalidArgument if ``spec`` contain any errors or if an invalid name is specified.
    // @throws NotAllowedInCurrentState if the associated cluster is deing disabled.
    // @throws NotFound if CreateSpec#cluster is not registered on this vCenter server.
    // @throws Unsupported if CreateSpec#cluster is not enabled for Namespaces.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user does not have Namespaces.Configure privilege.
    Create(specParam CreateSpec) error 


    // Delete the namespace object in the cluster. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param namespaceParam Identifier for the namespace.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.namespaces.Instance``.
    // @throws Error if the system reports an error while responding to the request.
    // @throws NotFound if the specified namespace could not be located.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user does not have Namespaces.Configure privilege.
    Delete(namespaceParam string) error 


    // Returns information about a specific namespace. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param namespaceParam Identifier for the namespace.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.namespaces.Instance``.
    // @return Information about the desired state of the specified namespace.
    // @throws NotFound if namespace could not be located.
    // @throws Error if the system reports an error while responding to the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user does not have System.Read privilege.
    Get(namespaceParam string) (Info, error) 


    // Returns the information about all namespaces on this vCenter. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    // @return The list of summary of all namespaces.
    // @throws Error if the system reports an error while responding to the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user does not have System.Read privilege.
    List() ([]Summary, error) 


    // Set a new configuration on the namespace object. The specified configuration is applied in entirety and will replace the current configuration fully. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param namespaceParam Identifier for the namespace.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.namespaces.Instance``.
    // @param specParam New specification for the namespace.
    // @throws Error if the system reports an error while responding to the request.
    // @throws InvalidArgument if ``spec`` contain any errors.
    // @throws NotAllowedInCurrentState if the namespace is marked for deletion or the associated cluster is being disabled.
    // @throws NotFound if namespace with the name ``namespace`` could not be located.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user does not have Namespaces.Configure privilege.
    Set(namespaceParam string, specParam SetSpec) error 


    // Update the namespace object. The specified configuration is applied partially and null fields in ``spec`` will leave those parts of configuration as-is. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param namespaceParam Identifier for the namespace.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.namespaces.Instance``.
    // @param specParam Specification for updating the namespace.
    // @throws Error if the system reports an error while responding to the request.
    // @throws InvalidArgument if ``spec`` contain any errors.
    // @throws NotAllowedInCurrentState if the namespace is marked for deletion or the associated cluster is being disabled.
    // @throws NotFound if namespace with the name ``namespace`` could not be located.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user does not have Namespaces.Configure privilege.
    Update(namespaceParam string, specParam UpdateSpec) error 

}
