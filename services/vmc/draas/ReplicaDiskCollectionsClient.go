/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: ReplicaDiskCollections
 * Used by client-side stubs.
 */

package draas

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/vmc/draas/model"
)

type ReplicaDiskCollectionsClient interface {

    // Query replica disk collections
    //
    // @param orgParam Organization identifier (required)
    // @param sddcParam sddc identifier (required)
    // @param datastoreMoIdParam Represents the datastore moref id to search. (optional)
    // @throws Unauthenticated  Unauthorized
    // @throws Unauthorized  Forbidden
    // @throws NotFound  Not found
	Get(orgParam string, sddcParam string, datastoreMoIdParam *string) ([]model.ReplicaDiskCollection, error)
}
