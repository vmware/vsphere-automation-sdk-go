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

// The ``Providers`` interface provides methods to create a sync session, get information on Sync. **Warning:** This interface is available as technical preview. It may be changed in a future release.
type ProvidersClient interface {


    // Enumerates the sync providers. **Warning:** This method is available as technical preview. It may be changed in a future release.
    //
    // @param linkParam Unique identifier of the link
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.hvc.Links``.
    // @return The array of sync provider information.
    // @throws Error If list fails.
    // @throws Unauthorized If the user is not authorized to perform this operation.
    List(linkParam string) ([]Summary, error) 


    // Gets Sync information for a sync provider. **Warning:** This method is available as technical preview. It may be changed in a future release.
    //
    // @param linkParam Unique identifier of the link
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.hvc.Links``.
    // @param providerParam Unique identifier of the sync provider.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.hvc.sync.Providers``.
    // @return The Info of sync information for the provider.
    // @throws Error if the system reports an error while responding to the request.
    // @throws NotFound if the sync provider associated with ``provider`` does not exist.
    // @throws Unauthorized if the user is not authorized to perform this operation.
    Get(linkParam string, providerParam string) (Info, error) 


    // Initiates synchronization between the local and remote replicas for the sync provider. **Warning:** This method is available as technical preview. It may be changed in a future release.
    //
    // @param linkParam Unique identifier of the link
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.hvc.Links``.
    // @param providerParam Unique identifier representing the sync provider
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.hvc.sync.Providers``.
    // @param credentialsParam Credentials to use for authentication to the remote system. **Warning:** This parameter is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    // This parameter is currently required for certificate providers and optional for other providers.In the future, it may become optional for all providers.
    // @throws Error if the system reports an error while responding to the request.
    // @throws NotFound if the link associated with ``link`` does not exist if the provider associated with ``provider`` is not registered for sync
    // @throws Unauthorized if the user is not authorized to perform this operation.
    // @throws ResourceBusy if a sync is already running.
    Start(linkParam string, providerParam string, credentialsParam *Credentials) error 

}
