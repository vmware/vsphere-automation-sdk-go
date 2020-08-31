/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Publish
 * Used by client-side stubs.
 */

package msft_licensing

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/vmc/model"
)

type PublishClient interface {

    // Sets the current Microsoft license status of the given SDDC's cluster.
    //
    // @param orgParam Organization identifier (required)
    // @param sddcParam Sddc identifier (required)
    // @param clusterParam cluster identifier (required)
    // @param configChangeParam The license data to set for the clusters. (required)
    // @return com.vmware.vmc.model.Task
    // @throws Unauthorized  Forbidden
	Post(orgParam string, sddcParam string, clusterParam string, configChangeParam model.MsftLicensingConfig) (model.Task, error)
}
