/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Providers
 * Used by client-side stubs.
 */

package sync


// The ``Providers`` interface provides methods to create a sync session, get information on Sync. Usage beyond VMware Cloud on AWS is not supported. **Warning:** This interface is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
type ProvidersClient interface {

    // Enumerates the sync providers. Usage beyond VMware Cloud on AWS is not supported. **Warning:** This method is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    //
    // @param linkParam Unique identifier of the link
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.hvc.Links``.
    // @return The array of sync provider information.
    // @throws Error If list fails.
    // @throws Unauthorized If the user is not authorized to perform this operation.
	List(linkParam string) ([]ProvidersSummary, error)

    // Gets Sync information for a sync provider. Usage beyond VMware Cloud on AWS is not supported. **Warning:** This method is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    //
    // @param linkParam Unique identifier of the link
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.hvc.Links``.
    // @param providerParam Unique identifier of the sync provider.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.hvc.sync.Providers``.
    // @return The Info of sync information for the provider.
    // @throws Error if the system reports an error while responding to the request.
    // @throws NotFound if the sync provider associated with ``provider`` does not exist.
    // @throws Unauthorized if the user is not authorized to perform this operation.
	Get(linkParam string, providerParam string) (ProvidersInfo, error)

    // Initiates synchronization between the local and remote replicas for the sync provider. Usage beyond VMware Cloud on AWS is not supported. **Warning:** This method is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    //
    // @param linkParam Unique identifier of the link
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.hvc.Links``.
    // @param providerParam Unique identifier representing the sync provider
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.hvc.sync.Providers``.
    // @throws Error if the system reports an error while responding to the request.
    // @throws NotFound if the link associated with ``link`` does not exist if the provider associated with ``provider`` is not registered for sync
    // @throws Unauthorized if the user is not authorized to perform this operation.
    // @throws ResourceBusy if a sync is already running.
	Start(linkParam string, providerParam string) error
}
