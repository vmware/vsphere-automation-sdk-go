/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: CertificateProfile
 * Used by client-side stubs.
 */

package trust_management

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-gm/model"
)

type CertificateProfileClient interface {

    // Get an available certificate profile
    //
    // @param serviceTypeParam Unique Service Type of the Certificate Profile (required)
    // @return com.vmware.nsx_global_policy.model.CertificateProfile
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get(serviceTypeParam string) (model.CertificateProfile, error)
}
