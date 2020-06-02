/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: IntrusionServices
 * Used by client-side stubs.
 */

package security

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt/model"
)

type IntrusionServicesClient interface {

    // Intrusion detection system settings.
    // @return com.vmware.nsx_policy.model.IdsSettings
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get() (model.IdsSettings, error)

    // Intrusion detection system settings.
    //
    // @param idsSettingsParam (required)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Patch(idsSettingsParam model.IdsSettings) error

    // Intrusion detection system settings.
    //
    // @param idsSettingsParam (required)
    // @return com.vmware.nsx_policy.model.IdsSettings
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Update(idsSettingsParam model.IdsSettings) (model.IdsSettings, error)
}
