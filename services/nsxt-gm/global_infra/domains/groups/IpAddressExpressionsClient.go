/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: IpAddressExpressions
 * Used by client-side stubs.
 */

package groups

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-gm/model"
)

type IpAddressExpressionsClient interface {

    // It will add or remove the specified IP Addresses from a given expression of a group.
    //
    // @param domainIdParam (required)
    // @param groupIdParam (required)
    // @param expressionIdParam (required)
    // @param ipAddressListParam (required)
    // @param actionParam Add or Remove group members. (required)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Create(domainIdParam string, groupIdParam string, expressionIdParam string, ipAddressListParam model.IPAddressList, actionParam string) error

    // Delete Group IPAddressExpression
    //
    // @param domainIdParam Domain ID (required)
    // @param groupIdParam Group ID (required)
    // @param expressionIdParam IPAddressExpression ID (required)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Delete(domainIdParam string, groupIdParam string, expressionIdParam string) error

    // If a group IPAddressExpression with the expression-id is not already present, create a new IPAddressExpression. If it already exists, replace the existing IPAddressExpression.
    //
    // @param domainIdParam Domain ID (required)
    // @param groupIdParam Group ID (required)
    // @param expressionIdParam IPAddressExpression ID (required)
    // @param ipAddressExpressionParam (required)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Patch(domainIdParam string, groupIdParam string, expressionIdParam string, ipAddressExpressionParam model.IPAddressExpression) error
}
