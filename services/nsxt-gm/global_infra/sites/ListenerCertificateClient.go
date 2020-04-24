/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Listener_certificate
 * Used by client-side stubs.
 */

package sites

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-gm/model"
)

type Listener_certificateClient interface {

    // Connects to the given IP and port, and, if an SSL listener is present, returns the certificate of the listener. Intent of this API is \"Do you trust this certificate?\".
    //
    // @param addressParam Host name or IP address of TLS listener (required)
    // @param portParam TCP port number of the TLS listener (required)
    // @return com.vmware.nsx_global_policy.model.TlsListenerCertificate
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get(addressParam string, portParam int64) (model.TlsListenerCertificate, error)
}
