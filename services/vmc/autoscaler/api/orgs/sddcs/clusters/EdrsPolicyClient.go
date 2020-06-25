/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: EdrsPolicy
 * Used by client-side stubs.
 */

package clusters

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/vmc/autoscaler/model"
)

type EdrsPolicyClient interface {

    // Get the current EDRS policy on a cluster.
    //
    // @param orgParam org identifier (required)
    // @param sddcParam Sddc identifier (required)
    // @param clusterParam cluster identifier (required)
    // @return com.vmware.model.EdrsPolicy
    // @throws Unauthenticated  Unauthorized
    // @throws InvalidRequest  Invalid action or bad argument
    // @throws Unauthorized  Forbidden
	Get(orgParam string, sddcParam string, clusterParam string) (model.EdrsPolicy, error)

    // Change the current EDRS policy on a cluster.
    //
    // @param orgParam org identifier (required)
    // @param sddcParam Sddc identifier (required)
    // @param clusterParam cluster identifier (required)
    // @param edrsPolicyParam EDRS policy (required)
    // @return com.vmware.model.Task
    // @throws Unauthenticated  Unauthorized
    // @throws InvalidRequest  Invalid action or bad argument
    // @throws Unauthorized  Forbidden
	Post(orgParam string, sddcParam string, clusterParam string, edrsPolicyParam model.EdrsPolicy) (model.Task, error)
}
