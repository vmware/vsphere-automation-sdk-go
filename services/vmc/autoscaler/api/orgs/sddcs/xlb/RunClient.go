/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Run
 * Used by client-side stubs.
 */

package xlb

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/vmc/autoscaler/model"
)

type RunClient interface {

    // Execute cross-cluster load balancing operations on the sddc.
    //
    // @param orgParam org identifier (required)
    // @param sddcParam Sddc identifier (required)
    // @return com.vmware.model.Task
    // @throws Unauthenticated  Unauthorized
    // @throws InvalidRequest  Invalid action or bad argument
    // @throws Unauthorized  Forbidden
	Post(orgParam string, sddcParam string) (model.Task, error)
}
