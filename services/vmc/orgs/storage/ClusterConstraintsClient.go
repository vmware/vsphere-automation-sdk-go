/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: ClusterConstraints
 * Used by client-side stubs.
 */

package storage

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/vmc/model"
)

type ClusterConstraintsClient interface {

    // Get constraints on cluster storage size for EBS-backed clusters.
    //
    // @param orgParam Organization identifier (required)
    // @param providerParam Cloud storage provider ID (example AWS) (required)
    // @param numHostsParam Number of hosts in cluster (required)
    // @return com.vmware.vmc.model.VsanConfigConstraints
    // @throws Unauthenticated  Unauthorized
    // @throws InvalidRequest  Invalid or missing parameters
    // @throws Unauthorized  Forbidden
	Get(orgParam string, providerParam string, numHostsParam int64) (model.VsanConfigConstraints, error)
}
