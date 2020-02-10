/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Settings
 * Used by client-side stubs.
 */

package tpm2


// The ``Settings`` interface provides methods to get or update settings related to the TPM 2.0 attestation protocol behavior. This interface was added in vSphere API 7.0.0.
type SettingsClient interface {

    // Return the TPM 2.0 protocol settings. This method was added in vSphere API 7.0.0.
    //
    // @param clusterParam The id of the cluster on which the operation will be executed.
    // The parameter must be an identifier for the resource type: ``ClusterComputeResource``.
    // @return The settings.
    // @throws Error if there is a generic error.
    // @throws InvalidArgument if the cluster id is empty.
    // @throws NotFound if ``cluster`` doesn't match to any cluster in the vCenter.
    // @throws Unauthenticated if the caller is not authenticated.
	Get(clusterParam string) (SettingsInfo, error)

    // Set the TPM 2.0 protocol settings. This method was added in vSphere API 7.0.0.
    //
    // @param clusterParam The id of the cluster on which the operation will be executed.
    // The parameter must be an identifier for the resource type: ``ClusterComputeResource``.
    // @param specParam The settings.
    // @throws Error if there is a generic error.
    // @throws InvalidArgument if the spec is invalid or cluster id is empty.
    // @throws NotFound if ``cluster`` doesn't match to any cluster in the vCenter.
    // @throws Unauthenticated if the caller is not authenticated.
	Update(clusterParam string, specParam SettingsUpdateSpec) error
}
