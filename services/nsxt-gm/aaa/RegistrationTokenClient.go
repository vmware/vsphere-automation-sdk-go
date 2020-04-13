/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: RegistrationToken
 * Used by client-side stubs.
 */

package aaa

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/services/nsxt-gm/model"
)

type RegistrationTokenClient interface {

    // The privileges of the registration token will be the same as the caller.
    // @return com.vmware.nsx_global_policy.model.RegistrationToken
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Create() (model.RegistrationToken, error)

    // Delete registration access token
    //
    // @param tokenParam Registration token (required)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Delete(tokenParam string) error

    // Get registration access token
    //
    // @param tokenParam Registration token (required)
    // @return com.vmware.nsx_global_policy.model.RegistrationToken
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get(tokenParam string) (model.RegistrationToken, error)
}
