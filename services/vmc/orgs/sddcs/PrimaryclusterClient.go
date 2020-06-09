/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Primarycluster
 * Used by client-side stubs.
 */

package sddcs

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/vmc/model"
)

type PrimaryclusterClient interface {

    // Retrieves the primary cluster in provided customer sddc UUID
    //
    // @param orgParam Organization identifier. (required)
    // @param sddcParam Sddc Identifier. (required)
    // @return com.vmware.vmc.model.Cluster
    // @throws Unauthenticated  Unauthorized
    // @throws InvalidRequest  The sddc is not in a state that's valid for fetching the primary cluster.
    // @throws Unauthorized  Access not allowed to the operation for the current user
    // @throws NotFound  Cannot find the sddc with the given identifier to fetch the primary cluster
	Get(orgParam string, sddcParam string) (model.Cluster, error)
}
