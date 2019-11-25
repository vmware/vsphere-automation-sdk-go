/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: IdsConfig
 * Used by client-side stubs.
 */

package intrusion_services

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/services/nsxt/model"
)

type IdsConfigClient interface {

    // Read intrusion detection system config of standalone hosts.
    // @return com.vmware.nsx_policy.model.IdsConfig
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get() (model.IdsConfig, error)

    // Patch intrusion detection system configuration on standalone hosts.
    //
    // @param idsConfigParam (required)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Patch(idsConfigParam model.IdsConfig) error

    // Update intrusion detection system configuration on standalone hosts.
    //
    // @param idsConfigParam (required)
    // @return com.vmware.nsx_policy.model.IdsConfig
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Update(idsConfigParam model.IdsConfig) (model.IdsConfig, error)
}
