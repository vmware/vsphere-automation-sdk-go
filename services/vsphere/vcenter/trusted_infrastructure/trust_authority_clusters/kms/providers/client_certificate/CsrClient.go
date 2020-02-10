/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Csr
 * Used by client-side stubs.
 */

package client_certificate


// The ``Csr`` interface provides methods to create a certificate signing request(CSR). This interface was added in vSphere API 7.0.0.
type CsrClient interface {

    // Generate a certificate signing request (CSR) for the client certificate. This overwrites any existing CSR. 
    //
    //  The CSR will allow the certificate to be signed by a third party. Once the CSR is signed, certificate may be added as a client certificate. 
    //
    //  Calling the API repeatedly will result in a generating a new CSR each time.. This method was added in vSphere API 7.0.0.
    //
    // @param clusterParam Identifier of the cluster.
    // The parameter must be an identifier for the resource type: ``ClusterComputeResource``.
    // @param providerParam Identifier of the provider.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.trusted_platform.trusted_clusters.kms.Provider``.
    // @return The client CSR, PEM and host ID which issued it.
    // @throws InvalidArgument If cluster or provider id are empty.
    // @throws NotFound If the provider or cluster is not found.
    // @throws Unauthenticated If the caller is not authenticated.
    // @throws Error If any other error occurs.
	Create(clusterParam string, providerParam string) (CsrInfo, error)

    // Get existing certificate signing request (CSR) for the client certificate. 
    //
    //  The CSR will allow the certificate to be signed by a third party. Once the CSR is signed, certificate may be added as a client certificate. 
    //
    // . This method was added in vSphere API 7.0.0.
    //
    // @param clusterParam Identifier of the cluster.
    // The parameter must be an identifier for the resource type: ``ClusterComputeResource``.
    // @param providerParam Identifier of the provider.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.trusted_platform.trusted_clusters.kms.Provider``.
    // @return The client CSR, PEM and host ID which issued it.
    // @throws InvalidArgument If cluster or provider id are empty.
    // @throws NotFound If the provider or cluster is not found.
    // @throws Unauthenticated If the caller is not authenticated.
    // @throws Error If any other error occurs.
	Get(clusterParam string, providerParam string) (CsrInfo, error)
}
