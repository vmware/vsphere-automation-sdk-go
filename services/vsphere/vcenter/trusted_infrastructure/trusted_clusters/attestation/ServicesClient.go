/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Services
 * Used by client-side stubs.
 */

package attestation


// The ``Services`` interface contains information about the ``Services`` instances a cluster uses. This interface was added in vSphere API 7.0.0.
type ServicesClient interface {

    // Returns the list of all ``Services`` configured for a cluster. This method was added in vSphere API 7.0.0.
    //
    // @param clusterParam The ID of the cluster.
    // The parameter must be an identifier for the resource type: ``ClusterComputeResource``.
    // @param specParam Only return services matching the filters.
    // If {\\\\@term.unset} return all services.
    // @return List of all ``Services`` used by this cluster.
    // @throws Error if there is a generic error.
    // @throws NotFound if the cluster ID is invalid.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``TrustedAdmin.ReadTrustedHosts``.
    // * The resource ``ClusterComputeResource`` referenced by the parameter ``cluster`` requires ``System.View``.
	List(clusterParam string, specParam *ServicesFilterSpec) ([]ServicesSummary, error)

    // Returns detailed information for a ``Services`` configured for a cluster. This method was added in vSphere API 7.0.0.
    //
    // @param clusterParam The ID of the cluster.
    // The parameter must be an identifier for the resource type: ``ClusterComputeResource``.
    // @param serviceParam The ID of the service.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.trusted_platform.attestation.Service``.
    // @return Info for the specified ``Services``.
    // @throws Error if there is a generic error.
    // @throws NotFound if the cluster or the service ID is invalid.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``TrustedAdmin.ReadTrustedHosts``.
    // * The resource ``ClusterComputeResource`` referenced by the parameter ``cluster`` requires ``System.View``.
	Get(clusterParam string, serviceParam string) (ServicesInfo, error)

    // Configures the cluster to use a new ``Services``. This method was added in vSphere API 7.0.0.
    //
    // @param clusterParam The ID of the cluster.
    // The parameter must be an identifier for the resource type: ``ClusterComputeResource``.
    // @param specParam Describes the ``Services``.
    // @return ID of the configured ``Services`` instance.
    // The return value will be an identifier for the resource type: ``com.vmware.vcenter.trusted_platform.attestation.Service``.
    // @throws AlreadyExists if the ``Services`` is already configured for this cluster
    // @throws Error for any other error.
    // @throws InvalidArgument if the CreateSpec is not valid.
    // @throws NotFound if the cluster ID is not valid.
    // @throws UnableToAllocateResource if all the hosts in the cluster do not have VMware vSphere Trust Authority enabled license.
    // @throws Unauthenticated if the user can not be authenticated.
	Create(clusterParam string, specParam ServicesCreateSpec) (string, error)

    // Marks the ``Services`` so that it will no longer be used by the cluster. This method was added in vSphere API 7.0.0.
    //
    // @param clusterParam the unique ID of the cluster.
    // The parameter must be an identifier for the resource type: ``ClusterComputeResource``.
    // @param serviceParam the ``Services`` instance unique identifier.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.trusted_platform.attestation.Service``.
    // @throws Error if there is a generic error.
    // @throws NotFound if the ``Services`` or the cluster are not found.
    // @throws Unauthenticated if the user can not be authenticated.
	Delete(clusterParam string, serviceParam string) error
}
