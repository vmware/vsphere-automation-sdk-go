/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Clusters
 * Used by client-side stubs.
 */

package sddcs

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/vmc/model"
)

type ClustersClient interface {

    // Creates a new cluster in customers sddcs with passed clusterConfig.
    //
    // @param orgParam Organization identifier (required)
    // @param sddcParam Sddc identifier (required)
    // @param clusterConfigParam clusterConfig (required)
    // @return com.vmware.vmc.model.Task
    // @throws Unauthenticated  Unauthorized
    // @throws InvalidRequest  The sddc is not in a state that's valid for updates
    // @throws Unauthorized  Access not allowed to the operation for the current user
    // @throws NotFound  Cannot find the cluster with the given identifier
	Create(orgParam string, sddcParam string, clusterConfigParam model.ClusterConfig) (model.Task, error)

    // This is a force operation which will delete the cluster even if there can be a data loss. Before calling this operation, all the VMs should be powered off.
    //
    // @param orgParam Organization identifier (required)
    // @param sddcParam Sddc identifier (required)
    // @param clusterParam cluster identifier (required)
    // @return com.vmware.vmc.model.Task
    // @throws Unauthenticated  Unauthorized
    // @throws InvalidRequest  The sddc is not in a state that's valid for updates
    // @throws Unauthorized  Access not allowed to the operation for the current user
    // @throws NotFound  Cannot find the cluster with the given id
	Delete(orgParam string, sddcParam string, clusterParam string) (model.Task, error)
}
