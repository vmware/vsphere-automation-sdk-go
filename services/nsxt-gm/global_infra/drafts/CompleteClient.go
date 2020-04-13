/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Complete
 * Used by client-side stubs.
 */

package drafts

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/services/nsxt-gm/model"
)

type CompleteClient interface {

    // Get a preview of a configuration which will be present after publish of a specified draft. The response essentially is a hierarichal payload containing the configuration, which will be in active after a specified draft gets published onto current configuration.
    //
    // @param draftIdParam (required)
    // @return com.vmware.nsx_global_policy.model.Infra
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get(draftIdParam string) (model.Infra, error)
}
