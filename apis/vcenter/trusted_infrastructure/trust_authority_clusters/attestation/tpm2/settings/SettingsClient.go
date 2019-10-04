/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Settings
 * Used by client-side stubs.
 */

package settings

import (
)

// The ``Settings`` interface provides methods to get or update settings related to the TPM 2.0 attestation protocol behavior.
type SettingsClient interface {


    // Return the TPM 2.0 protocol settings.
    //
    // @param clusterParam The id of the cluster on which the operation will be executed.
    // The parameter must be an identifier for the resource type: ``ClusterComputeResource``.
    // @return The settings.
    // @throws Error if there is a generic error.
    // @throws InvalidArgument if the cluster id is empty.
    // @throws NotFound if ``cluster`` doesn't match to any cluster in the vCenter.
    // @throws Unauthenticated if the caller is not authenticated.
    Get(clusterParam string) (Info, error) 


    // Set the TPM 2.0 protocol settings.
    //
    // @param clusterParam The id of the cluster on which the operation will be executed.
    // The parameter must be an identifier for the resource type: ``ClusterComputeResource``.
    // @param specParam The settings.
    // @throws Error if there is a generic error.
    // @throws InvalidArgument if the spec is invalid or cluster id is empty.
    // @throws NotFound if ``cluster`` doesn't match to any cluster in the vCenter.
    // @throws Unauthenticated if the caller is not authenticated.
    Update(clusterParam string, specParam UpdateSpec) error 

}
