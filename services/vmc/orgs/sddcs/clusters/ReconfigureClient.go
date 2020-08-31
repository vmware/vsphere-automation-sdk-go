/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Reconfigure
 * Used by client-side stubs.
 */

package clusters

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/vmc/model"
)

type ReconfigureClient interface {

    // This reconfiguration will handle changing both the number of hosts and the cluster storage capacity.
    //
    // @param orgParam Organization identifier (required)
    // @param sddcParam Sddc identifier (required)
    // @param clusterParam cluster identifier (required)
    // @param clusterReconfigureParamsParam clusterReconfigureParams (required)
    // @return com.vmware.vmc.model.Task
    // @throws Unauthenticated  Unauthorized
    // @throws InvalidRequest  Bad Request. The SDDC is not in a valid state. , Method not allowed
    // @throws Unauthorized  Access not allowed to the operation for the current user
    // @throws NotFound  Cannot find the cluster with the given id
	ClusterReconfig(orgParam string, sddcParam string, clusterParam string, clusterReconfigureParamsParam model.ClusterReconfigureParams) (model.Task, error)
}
