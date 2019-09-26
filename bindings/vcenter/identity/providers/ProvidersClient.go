/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Providers
 * Used by client-side stubs.
 */

package providers

import (
)

// The ``Providers`` interface provides methods to list, read and modify vCenter Server identity providers. **Warning:** This interface is available as technical preview. It may be changed in a future release.
type ProvidersClient interface {


    // Retrieve all identity providers. **Warning:** This method is available as technical preview. It may be changed in a future release.
    // @return Commonly used information about the identity providers.
    // @throws Unauthorized if authorization is not given to caller.
    List() ([]Summary, error) 


    // Retrieve detailed information of the specified identity provider. **Warning:** This method is available as technical preview. It may be changed in a future release.
    //
    // @param providerParam the identifier of the provider
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.identity.Providers``.
    // @return Detailed information of the specified identity provider.
    // @throws Unauthorized if authorization is not given to caller.
    // @throws NotFound if no provider found with the given provider identifier.
    Get(providerParam string) (Info, error) 


    // Create a vCenter Server identity provider. **Warning:** This method is available as technical preview. It may be changed in a future release.
    //
    // @param specParam the CreateSpec contains the information used to create the provider
    // @return The identifier of the created identity provider.
    // The return value will be an identifier for the resource type: ``com.vmware.vcenter.identity.Providers``.
    // @throws Unauthorized if authorization is not given to caller.
    // @throws InvalidArgument if invalid arguments are provided in createSpec.
    Create(specParam CreateSpec) (string, error) 


    // Update a vCenter Server identity provider. **Warning:** This method is available as technical preview. It may be changed in a future release.
    //
    // @param providerParam the identifier of the provider to update
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.identity.Providers``.
    // @param specParam the UpdateSpec contains the information used to update the provider
    // @throws Unauthorized if authorization is not given to caller.
    // @throws InvalidArgument if invalid arguments are provided in updateSpec.
    // @throws NotFound if no provider found with the given provider identifier.
    Update(providerParam string, specParam UpdateSpec) error 


    // Delete a vCenter Server identity provider. **Warning:** This method is available as technical preview. It may be changed in a future release.
    //
    // @param providerParam the identifier of the provider to delete
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.identity.Providers``.
    // @throws Unauthorized if authorization is not given to caller.
    // @throws NotFound if no provider found with the given provider identifier.
    Delete(providerParam string) error 

}
