/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Analysis
 * Used by client-side stubs.
 */

package xlb

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/vmc/autoscaler/model"
)

type AnalysisClient interface {

    // Get cross-cluster load-balancer recommendations for the sddc.
    //
    // @param orgParam org identifier (required)
    // @param sddcParam Sddc identifier (required)
    // @param clustersParam List of input cluster uuids. The minimum number of input clusters for cross-cluster load balancer is 2. The maximum number of input clusters for cross-cluster load balancer is 4. Example - \\\\`[\"e823eb1c-d79d-4763-ad6b-b447c14b6cd2\", \"71066e28-3a15-4670-8555-72068c9d5320\"]\\\\` (required)
    // @return com.vmware.model.Task
    // @throws Unauthenticated  Unauthorized
    // @throws InvalidRequest  Invalid action or bad argument
    // @throws Unauthorized  Forbidden
	Post(orgParam string, sddcParam string, clustersParam []string) (model.Task, error)
}
