/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: ClientCertificate
 * Used by client-side stubs.
 */

package providers


// The ``ClientCertificate`` interface provides methods to add and retrieve client certificate. This interface was added in vSphere API 7.0.0.
type ClientCertificateClient interface {

    // Generate a new self signed client certificate. Existing client certificate is overwritten. The key server will use this certificate to validate the client connection. This method was added in vSphere API 7.0.0.
    //
    // @param clusterParam Identifier of the cluster.
    // The parameter must be an identifier for the resource type: ``ClusterComputeResource``.
    // @param providerParam Identifier of the provider.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.trusted_platform.trusted_clusters.kms.Provider``.
    // @throws InvalidArgument If cluster or provider id are empty.
    // @throws NotFound If the cluster or provider is not found.
    // @throws Unauthenticated If the caller is not authenticated.
    // @throws Error If any other error occurs.
	Create(clusterParam string, providerParam string) error

    // Return the existing client certificate. This method was added in vSphere API 7.0.0.
    //
    // @param clusterParam Identifier of the cluster.
    // The parameter must be an identifier for the resource type: ``ClusterComputeResource``.
    // @param providerParam Identifier of the provider.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.trusted_platform.trusted_clusters.kms.Provider``.
    // @return Client certificate, PEM.
    // @throws InvalidArgument If cluster or provider id are empty.
    // @throws NotFound If the cluster or provider is not found.
    // @throws Unauthenticated If the caller is not authenticated.
    // @throws Error If any other error occurs.
	Get(clusterParam string, providerParam string) (ClientCertificateInfo, error)

    // Update the client certificate. 
    //
    //  The key server will use this certificate to validate the client connection. If a client certificate already exists, it will be replaced. 
    //
    //  An optional private key can be specified if the certificate has already been provisioned.. This method was added in vSphere API 7.0.0.
    //
    // @param clusterParam Identifier of the cluster.
    // The parameter must be an identifier for the resource type: ``ClusterComputeResource``.
    // @param providerParam Identifier of the provider.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.trusted_platform.trusted_clusters.kms.Provider``.
    // @param specParam The update spec.
    // @throws InvalidArgument If the certificate or private key is invalid or cluster/provider id are empty.
    // @throws NotFound If the cluster or provider is not found.
    // @throws Unauthenticated If the caller is not authenticated.
    // @throws Error If any other error occurs.
	Update(clusterParam string, providerParam string, specParam ClientCertificateUpdateSpec) error
}
