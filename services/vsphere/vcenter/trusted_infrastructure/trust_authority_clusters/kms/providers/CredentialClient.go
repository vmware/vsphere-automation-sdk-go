/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Credential
 * Used by client-side stubs.
 */

package providers


// The ``Credential`` interface provides methods to add a credential for external key management service(s). This interface was added in vSphere API 7.0.0.
type CredentialClient interface {

    // Set the key server credential. This method was added in vSphere API 7.0.0.
    //
    // @param clusterParam Identifier of the cluster.
    // The parameter must be an identifier for the resource type: ``ClusterComputeResource``.
    // @param providerParam Identifier of the provider.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.trusted_platform.trusted_clusters.kms.Provider``.
    // @param credentialParam KMIP KMS password or AWS access key.
    // @throws InvalidArgument If cluster or provider id are empty.
    // @throws NotFound If the provider or cluster is not found.
    // @throws Unauthenticated If the caller is not authenticated.
    // @throws Error If any other error occurs.
	Set(clusterParam string, providerParam string, credentialParam string) error
}
