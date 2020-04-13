/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Backups
 * Used by client-side stubs.
 */

package cluster

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/services/nsxt-gm/model"
)

type BackupsClient interface {

    // Get SHA256 fingerprint of ECDSA key of remote server. The caller should independently verify that the key is trusted.
    //
    // @param remoteServerFingerprintRequestParam (required)
    // @return com.vmware.nsx_global_policy.model.RemoteServerFingerprint
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Retrievesshfingerprint(remoteServerFingerprintRequestParam model.RemoteServerFingerprintRequest) (model.RemoteServerFingerprint, error)
}
