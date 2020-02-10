/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: BaseImages
 * Used by client-side stubs.
 */

package esx


// The ``BaseImages`` interface provides methods to manage trusted instances of ESX software on a cluster level. This interface was added in vSphere API 7.0.0.
type BaseImagesClient interface {

    // Import ESX metadata as a new trusted base image to each host in the cluster. 
    //
    //  Import a boot_imgdb.tgz file which contains metadata that describes a trusted ESX base image. A boot_imgdb.tgz file can be downloaded from a representative host.. This method was added in vSphere API 7.0.0.
    //
    // @param clusterParam The id of the cluster on which the operation will be executed.
    // The parameter must be an identifier for the resource type: ``ClusterComputeResource``.
    // @param imgdbParam ESX metadata on a cluster level.
    // @return The imported imgdb version identifier.
    // The return value will be an identifier for the resource type: ``com.vmware.vcenter.trusted_platform.trusted_clusters.attestation.os.esx.BaseImage``.
    // @throws AlreadyExists if ESX metadata for the same version has already been added.
    // @throws Error if there is a generic error.
    // @throws InvalidArgument if the imgdb is invalid or cluster id is empty.
    // @throws NotFound if the cluster is not found.
    // @throws Unauthenticated if the caller is not authenticated.
	ImportFromImgdb(clusterParam string, imgdbParam []byte) (string, error)

    // Return a list of trusted ESX base images. This method was added in vSphere API 7.0.0.
    //
    // @param clusterParam The id of the cluster on which the operation will be executed.
    // The parameter must be an identifier for the resource type: ``ClusterComputeResource``.
    // @param specParam The search specification.
    // if {\\\\@term.unset} return all information.
    // @return A list of configured trusted ESX base images.
    // @throws Error if there is a generic error.
    // @throws InvalidArgument if the cluster id is empty.
    // @throws NotFound if the cluster is not found.
    // @throws Unauthenticated if the caller is not authenticated.
	List(clusterParam string, specParam *BaseImagesFilterSpec) ([]BaseImagesSummary, error)

    // Remove a trusted ESX base image of each ESX in the cluster. This method was added in vSphere API 7.0.0.
    //
    // @param clusterParam The id of the cluster on which the operation will be executed.
    // The parameter must be an identifier for the resource type: ``ClusterComputeResource``.
    // @param versionParam The ESX base image version.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.trusted_platform.trusted_clusters.attestation.os.esx.BaseImage``.
    // @throws Error if there is a generic error.
    // @throws InvalidArgument if the version is invalid or the cluster id is empty.
    // @throws NotFound if the version or cluster is not found.
    // @throws Unauthenticated if the caller is not authenticated.
	Delete(clusterParam string, versionParam string) error

    // Get the trusted ESX base version details. This method was added in vSphere API 7.0.0.
    //
    // @param clusterParam The id of the cluster on which the operation will be executed.
    // The parameter must be an identifier for the resource type: ``ClusterComputeResource``.
    // @param versionParam The ESX base image version.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.trusted_platform.trusted_clusters.attestation.os.esx.BaseImage``.
    // @return The version info
    // @throws Error if there is a generic error.
    // @throws InvalidArgument if the version is invalid or the cluster id is empty.
    // @throws NotFound if the version or cluster is not found.
    // @throws Unauthenticated if the caller is not authenticated.
	Get(clusterParam string, versionParam string) (BaseImagesInfo, error)
}
