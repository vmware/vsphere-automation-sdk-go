/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: CloudServiceVmcOnAwsPublicIp
 * Used by client-side stubs.
 */

package api

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-vmc-aws-integration/model"
)

type CloudServiceVmcOnAwsPublicIpClient interface {

    // This API is used to create or update a public IP. In creating, the API allocates a new public IP from VMC provider. In updating, only the display name can be modified, the IP is read-only.
    //
    // @param publicIpIdParam (required)
    // @param publicIpParam (required)
    // @return com.vmware.model.PublicIp
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	CreatePublicIp(publicIpIdParam string, publicIpParam model.PublicIp) (model.PublicIp, error)

    // Delete a public IP. The IP will be released in VMC provider.
    //
    // @param publicIpIdParam (required)
    // @param forceParam Force delete the resource even if it is being used somewhere (optional, default to false)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	DeletePublicIp(publicIpIdParam string, forceParam *bool) error

    // Get public IP information associated with given ID.
    //
    // @param publicIpIdParam (required)
    // @return com.vmware.model.PublicIp
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	GetPublicIp(publicIpIdParam string) (model.PublicIp, error)

    // List all public IPs obtained in the SDDC.
    //
    // @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
    // @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
    // @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
    // @param sortAscendingParam (optional)
    // @param sortByParam Field by which records are sorted (optional)
    // @return com.vmware.model.PublicIpsListResult
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	ListPublicIps(cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.PublicIpsListResult, error)
}
