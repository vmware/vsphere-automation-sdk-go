/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: CaCertificates
 * Used by client-side stubs.
 */

package tpm2


// The ``CaCertificates`` interface provides methods to manage Trusted Platform Module (TPM) CA certificates. 
//
//  Endorsement Keys are typically packaged in a certificate that is signed by a certificate authority (CA). This interface allows the CA certificate to be registered with the Attestation Service in order to validate TPM EK certificates when presented at attestation time.. This interface was added in vSphere API 7.0.0.
type CaCertificatesClient interface {

    // Return a list of configured TPM CA certificates on a cluster. This method was added in vSphere API 7.0.0.
    //
    // @param clusterParam The id of the cluster on which the operation will be executed.
    // The parameter must be an identifier for the resource type: ``ClusterComputeResource``.
    // @return A list of configured TPM CA certificates.
    // @throws Error if there is a generic error.
    // @throws InvalidArgument if cluster id is empty.
    // @throws NotFound if the ``cluster`` doesn't match to any cluster in the vCenter or given name is not found.
    // @throws Unauthenticated if the caller is not authenticated.
	List(clusterParam string) ([]CaCertificatesSummary, error)

    // Add a new TPM CA certificate on a cluster. This method was added in vSphere API 7.0.0.
    //
    // @param clusterParam The id of the cluster on which the operation will be executed.
    // The parameter must be an identifier for the resource type: ``ClusterComputeResource``.
    // @param specParam The new CA certificate details.
    // @throws AlreadyExists if the certificate name exists.
    // @throws Error if there is a generic error.
    // @throws InvalidArgument if the configuration is invalid or the cluster id is empty.
    // @throws NotFound if ``cluster`` doesn't match to any cluster in the vCenter.
    // @throws Unauthenticated if the caller is not authenticated.
	Create(clusterParam string, specParam CaCertificatesCreateSpec) error

    // Remove a TPM CA certificate on a cluster. This method was added in vSphere API 7.0.0.
    //
    // @param clusterParam The id of the cluster on which the operation will be executed.
    // The parameter must be an identifier for the resource type: ``ClusterComputeResource``.
    // @param nameParam The CA certificate name.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.trusted_platform.trusted_clusters.attestation.tpm2.CaCertificate``.
    // @throws Error if there is a generic error.
    // @throws InvalidArgument if the name is invalid or cluster id is empty.
    // @throws NotFound if the ``cluster`` doesn't match to any cluster in the vCenter or given name is not found.
    // @throws Unauthenticated if the caller is not authenticated.
	Delete(clusterParam string, nameParam string) error

    // Get the TPM CA certificate details on a cluster. This method was added in vSphere API 7.0.0.
    //
    // @param clusterParam The id of the cluster on which the operation will be executed.
    // The parameter must be an identifier for the resource type: ``ClusterComputeResource``.
    // @param nameParam The CA certificate name.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.trusted_platform.trusted_clusters.attestation.tpm2.CaCertificate``.
    // @return CA certificate info.
    // @throws Error if there is a generic error.
    // @throws InvalidArgument if the name is invalid or cluster id is empty.
    // @throws NotFound if the CA certificate is not found or ``cluster`` doesn't match to any cluster in the vCenter.
    // @throws Unauthenticated if the caller is not authenticated.
	Get(clusterParam string, nameParam string) (CaCertificatesInfo, error)
}
