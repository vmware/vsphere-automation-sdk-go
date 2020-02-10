/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Mode
 * Used by client-side stubs.
 */

package cluster


// The Mode interface provides methods to manage the operating mode of a vCenter High Availability Cluster (VCHA Cluster). This interface was added in vSphere API 6.7.1.
type ModeClient interface {

    // Retrieves the current mode of a VCHA cluster. This method was added in vSphere API 6.7.1.
    // @return Info structure containing the mode of the the VCHA cluster.
    // @throws NotAllowedInCurrentState If the VCHA cluster is not configured.
    // @throws Unauthorized If the user has insufficient privilege to perform the operation. Operation execution requires the System.Read privilege.
    // @throws Error If any other error occurs.
	Get() (ModeInfo, error)

    // Manipulates the mode of a VCHA Cluster. Following mode transitions are allowed:
    //  enabled -> disabled - Allowed only in healthy and degraded states.
    //  enabled -> maintenance - Allowed only in healthy state.
    //  disabled -> enabled - Allowed only in healthy state.
    //  maintenance -> enabled - Allowed only in healthy state with all nodes are running the same version.
    //  maintenance -> disabled - Allowed only in healthy state with all nodes are running the same version.
    //  All other transitions are not allowed. 
    //
    //  VCHA Cluster configuration remains intact in any of the cluster modes.. This method was added in vSphere API 6.7.1.
    //
    // @param modeParam Clustermode to change the VCHA cluster mode to.
    // @throws Unauthorized If the user has insufficient privilege to perform the operation. Operation execution requires the Global.VCServer privilege.
    // @throws Error If any other error occurs.
	Set(modeParam ModeClusterMode) error
}
