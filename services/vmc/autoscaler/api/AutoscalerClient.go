/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Autoscaler
 * Used by client-side stubs.
 */

package api

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/vmc/autoscaler/model"
)

type AutoscalerClient interface {

    // Get cross-cluster load-balancer recommendations for the sddc.
    //
    // @param orgParam org identifier (required)
    // @param sddcParam Sddc identifier (required)
    // @param clustersParam List of input cluster uuids. The minimum number of input clusters for cross-cluster load balancer is 2. The maximum number of input clusters for cross-cluster load balancer is 4. Example - \\\\`[\"e823eb1c-d79d-4763-ad6b-b447c14b6cd2\", \"71066e28-3a15-4670-8555-72068c9d5320\"]\\\\` (required)
    // @return com.vmware.model.Task
    // @throws Unauthenticated  Unauthorized
    // @throws InvalidRequest  Invalid action or bad argument
    // @throws Unauthorized  Forbidden
	Analysis(orgParam string, sddcParam string, clustersParam []string) (model.Task, error)

    // Retrieve details of an autoscaler task.
    //
    // @param orgParam org identifier (required)
    // @param taskParam task identifier (required)
    // @return com.vmware.model.Task
    // @throws Unauthenticated  Unauthorized
    // @throws Unauthorized  Forbidden
    // @throws NotFound  Cannot find the task with given identifier
	Get(orgParam string, taskParam string) (model.Task, error)

    // Execute cross-cluster load balancing operations on the sddc.
    //
    // @param orgParam org identifier (required)
    // @param sddcParam Sddc identifier (required)
    // @return com.vmware.model.Task
    // @throws Unauthenticated  Unauthorized
    // @throws InvalidRequest  Invalid action or bad argument
    // @throws Unauthorized  Forbidden
	Run(orgParam string, sddcParam string) (model.Task, error)

    // Stop cross-cluster load balancer initiated xvMotion operations on the sddc.
    //
    // @param orgParam org identifier (required)
    // @param sddcParam Sddc identifier (required)
    // @return com.vmware.model.Task
    // @throws Unauthenticated  Unauthorized
    // @throws InvalidRequest  Invalid action or bad argument
    // @throws Unauthorized  Forbidden
	Stop(orgParam string, sddcParam string) (model.Task, error)
}
