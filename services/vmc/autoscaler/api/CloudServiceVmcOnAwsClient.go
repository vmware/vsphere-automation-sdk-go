/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: CloudServiceVmcOnAws
 * Used by client-side stubs.
 */

package api

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/vmc/autoscaler/model"
)

type CloudServiceVmcOnAwsClient interface {

    // Retrieve the shadow account and linked VPC account information from VMC provider. This API is a live query to VMC provider.
    // @return com.vmware.model.VMCAccounts
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	ListAccounts() (model.VMCAccounts, error)
}
