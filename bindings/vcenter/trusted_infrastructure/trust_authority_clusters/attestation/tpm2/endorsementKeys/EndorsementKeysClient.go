/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: EndorsementKeys
 * Used by client-side stubs.
 */

package endorsementKeys

import (
)

// The ``EndorsementKeys`` interface provides methods to manage Trusted Platform Module (TPM) Endorsement Keys (EK) on a cluster level.
type EndorsementKeysClient interface {


    // Return a list of configured TPM endorsement keys in a cluster.
    //
    // @param clusterParam The id of the cluster on which the operation will be executed.
    // The parameter must be an identifier for the resource type: ``ClusterComputeResource``.
    // @return A list of configured endorsement keys.
    // @throws Error if there is a generic error.
    // @throws InvalidArgument if the cluster id is empty.
    // @throws NotFound if the cluster is not found.
    // @throws Unauthenticated if the caller is not authenticated.
    List(clusterParam string) ([]Summary, error) 


    // Add a new TPM endorsement key on a cluster.
    //
    // @param clusterParam The id of the cluster on which the operation will be executed.
    // The parameter must be an identifier for the resource type: ``ClusterComputeResource``.
    // @param specParam The configuration.
    // @throws AlreadyExists if the endorsement key name exists.
    // @throws Error if there is a generic error.
    // @throws InvalidArgument if the configuration is invalid or cluster id is empty.
    // @throws NotFound if ``cluster`` doesn't match to any cluster in the vCenter.
    // @throws Unauthenticated if the caller is not authenticated.
    Create(clusterParam string, specParam CreateSpec) error 


    // Remove a TPM endorsement key on a cluster.
    //
    // @param clusterParam The id of the cluster on which the operation will be executed.
    // The parameter must be an identifier for the resource type: ``ClusterComputeResource``.
    // @param nameParam The endorsement key name.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.trusted_platform.trusted_clusters.attestation.tpm2.EndorsementKey``.
    // @throws Error if there is a generic error.
    // @throws InvalidArgument if the name is invalid or cluster id is empty.
    // @throws NotFound if the name is not found or ``cluster`` doesn't match to any cluster in the vCenter.
    // @throws Unauthenticated if the caller is not authenticated.
    Delete(clusterParam string, nameParam string) error 


    // Get the TPM endorsement key details on a cluster.
    //
    // @param clusterParam The id of the cluster on which the operation will be executed.
    // The parameter must be an identifier for the resource type: ``ClusterComputeResource``.
    // @param nameParam The endorsement key name.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.trusted_platform.trusted_clusters.attestation.tpm2.EndorsementKey``.
    // @return The endorsement key info.
    // @throws Error if there is a generic error.
    // @throws InvalidArgument if the name is invalid or cluster id is empty.
    // @throws NotFound if the endorsement key is not found or ``cluster`` doesn't match to any cluster in the vCenter.
    // @throws Unauthenticated if the caller is not authenticated.
    Get(clusterParam string, nameParam string) (Info, error) 

}
