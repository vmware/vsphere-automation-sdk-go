/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: EdrsPolicy
 * Used by client-side stubs.
 */

package sddcs

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/vmc/autoscaler/model"
)

type EdrsPolicyClient interface {

    // Get the EDRS policy on every cluster in an SDDC.
    //
    // @param orgParam org identifier (required)
    // @param sddcParam Sddc identifier (required)
    // @throws Unauthenticated  Unauthorized
    // @throws InvalidRequest  Invalid action or bad argument
    // @throws Unauthorized  Forbidden
	Get(orgParam string, sddcParam string) ([]model.EdrsClusterInfo, error)
}
