/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: CloudServiceVmcOnAwsLinkedVpc
 * Used by client-side stubs.
 */

package api

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-vmc-aws-integration/model"
)

type CloudServiceVmcOnAwsLinkedVpcClient interface {

    // Get linked VPC information associated with given ID.
    //
    // @param linkedVpcIdParam (required)
    // @return com.vmware.model.LinkedVpcInfo
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	GetLinkedVpc(linkedVpcIdParam string) (model.LinkedVpcInfo, error)

    // List connected services to given linked vpc ID. The response consist of all available services along with their status.
    //
    // @param linkedVpcIdParam linked vpc id (required)
    // @return com.vmware.model.ConnectedServiceListResult
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	ListConnectedServices(linkedVpcIdParam string) (model.ConnectedServiceListResult, error)

    // List all linked VPC information.
    // @return com.vmware.model.LinkedVpcsListResult
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	ListLinkedVpcs() (model.LinkedVpcsListResult, error)

    // Connect/Disconnect the service to the given linked vpc. For example, connect S3. The user will know what services are available through the GET call. If the user is trying to connect/disconnect an unknown service, the POST call will throw a 400 Bad Request error.
    //
    // @param linkedVpcIdParam linked vpc id (required)
    // @param serviceNameParam connected service name, e.g. s3 (required)
    // @param connectedServiceStatusParam (required)
    // @return com.vmware.model.ConnectedServiceStatus
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	UpdateConnectedService(linkedVpcIdParam string, serviceNameParam string, connectedServiceStatusParam model.ConnectedServiceStatus) (model.ConnectedServiceStatus, error)
}
