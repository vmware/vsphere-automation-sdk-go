/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: UrlCategorizationService
 * Used by client-side stubs.
 */

package security_cloud_services

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

type UrlCategorizationServiceClient interface {

    // This API returns the details about the URL categorization cloud service. It informs about whether the URL categorization cloud service is up and running or not. It also returns the version of the URL data available on the cloud.
    // @return com.vmware.nsx_policy.model.UrlCategorizationCloudServiceInfo
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get() (model.UrlCategorizationCloudServiceInfo, error)
}
