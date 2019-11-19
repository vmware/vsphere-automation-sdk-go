/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Interfaces
 * Used by client-side stubs.
 */

package statistics

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/services/vmc/model"
)

type InterfacesClient interface {

    // Retrieve interface statistics for a management or compute gateway (NSX Edge).
    //
    // @param orgParam Organization identifier. (required)
    // @param sddcParam Sddc Identifier. (required)
    // @param edgeIdParam Edge Identifier. (required)
    // @param startTimeParam Show statistics from this start time (in milliseconds since epoch). (optional)
    // @param endTimeParam Show statistics until this end time (in milliseconds since epoch). (optional)
    // @return com.vmware.vmc.model.CbmStatistics
    // @throws InvalidRequest  Bad request. Request object passed is invalid.
    // @throws Unauthorized  Forbidden. Authorization header not provided.
    // @throws NotFound  Not found. Requested object not found.
	Get(orgParam string, sddcParam string, edgeIdParam string, startTimeParam *int64, endTimeParam *int64) (model.CbmStatistics, error)
}
