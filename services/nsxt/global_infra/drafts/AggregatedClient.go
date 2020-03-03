/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Aggregated
 * Used by client-side stubs.
 */

package drafts

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

type AggregatedClient interface {

    // Get an aggregated configuration that will get applied onto current configuration during publish of this draft. The response is a hierarichal payload containing the aggregated configuration differences from the latest auto draft till the specified draft.
    //
    // @param draftIdParam (required)
    // @return com.vmware.nsx_policy.model.Infra
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get(draftIdParam string) (model.Infra, error)
}
