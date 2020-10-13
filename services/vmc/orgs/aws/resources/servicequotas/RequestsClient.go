/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Requests
 * Used by client-side stubs.
 */

package servicequotas

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/vmc/model"
)

type RequestsClient interface {

    // Get Service Quota Request by its Id
    //
    // @param orgParam Organization identifier (required)
    // @param serviceQuotaRequestIdParam The UUID of the service quota request (required)
    // @return com.vmware.vmc.model.ServiceQuotaRequest
    // @throws Unauthenticated  Unauthorized
    // @throws Unauthorized  Forbidden
	Get(orgParam string, serviceQuotaRequestIdParam string) (model.ServiceQuotaRequest, error)

    // List all service quota requests of an org
    //
    // @param orgParam Organization identifier (required)
    // @throws Unauthenticated  Unauthorized
    // @throws Unauthorized  Forbidden
	List(orgParam string) ([]model.ServiceQuotaRequest, error)
}
