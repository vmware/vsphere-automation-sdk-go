/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Constraints
 * Used by client-side stubs.
 */

package config

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/vmc/model"
)

type ConstraintsClient interface {

    // This API provides the storage choices available when reconfiguring storage in a cluster. The constraints returned will give vSAN reconfiguration biases and available vSAN capacities per reconfiguration bias. The constraints also indicate the default vSAN capacity per reconfiguration biases as well as the default reconfiguration bias.
    //
    // @param orgParam Organization identifier. (required)
    // @param sddcParam Sddc Identifier. (required)
    // @param clusterParam cluster identifier (required)
    // @param expectedNumHostsParam The expected number of hosts in the cluster. If not specified, returned is based on current number of hosts in the cluster. (optional)
    // @return com.vmware.vmc.model.VsanClusterReconfigConstraints
    // @throws Unauthenticated  Unauthorized
    // @throws InvalidRequest  Invalid or missing parameters
    // @throws Unauthorized  Forbidden
	Get(orgParam string, sddcParam string, clusterParam string, expectedNumHostsParam *int64) (model.VsanClusterReconfigConstraints, error)
}
