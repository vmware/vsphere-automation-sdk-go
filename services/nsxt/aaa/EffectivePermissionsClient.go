/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: EffectivePermissions
 * Used by client-side stubs.
 */

package aaa

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

type EffectivePermissionsClient interface {

    // Returns none if user doesn't have access or feature_name from required request parameter is empty/invalid/doesn't match with object-path provided.
    //
    // @param featureNameParam Feature name (required)
    // @param objectPathParam Exact object Policy path (required)
    // @return com.vmware.nsx_policy.model.PathPermissionGroup
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get(featureNameParam string, objectPathParam string) (model.PathPermissionGroup, error)
}
