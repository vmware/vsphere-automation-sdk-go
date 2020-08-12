/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: ScaleSddcSize
 * Used by client-side stubs.
 */

package sddcs

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/vmc/autoscaler/model"
)

type ScaleSddcSizeClient interface {

    // Scale SDDC Size.
    //
    // @param orgParam org identifier (required)
    // @param sddcParam Sddc identifier (required)
    // @return com.vmware.model.Task
    // @throws InvalidRequest  Invalid action or bad argument
    // @throws Unauthorized  Forbidden
	Post(orgParam string, sddcParam string) (model.Task, error)
}
