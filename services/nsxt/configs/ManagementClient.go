/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Management
 * Used by client-side stubs.
 */

package configs

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

type ManagementClient interface {

    // Returns the NSX Management nodes global configuration.
    // @return com.vmware.nsx_policy.model.ManagementConfig
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get() (model.ManagementConfig, error)

    // Modifies the NSX Management nodes global configuration.
    //
    // @param managementConfigParam (required)
    // @return com.vmware.nsx_policy.model.ManagementConfig
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Update(managementConfigParam model.ManagementConfig) (model.ManagementConfig, error)
}
