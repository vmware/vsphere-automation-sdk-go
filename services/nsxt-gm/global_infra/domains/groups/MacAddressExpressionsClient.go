/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: MacAddressExpressions
 * Used by client-side stubs.
 */

package groups

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-gm/model"
)

type MacAddressExpressionsClient interface {

    // It will add or remove the specified MAC Addresses from a given expression of a group.
    //
    // @param domainIdParam (required)
    // @param groupIdParam (required)
    // @param expressionIdParam (required)
    // @param mACAddressListParam (required)
    // @param actionParam Add or Remove group members. (required)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Create(domainIdParam string, groupIdParam string, expressionIdParam string, mACAddressListParam model.MACAddressList, actionParam string) error

    // Delete Group MACAddressExpression
    //
    // @param domainIdParam Domain ID (required)
    // @param groupIdParam Group ID (required)
    // @param expressionIdParam MACAddressExpression ID (required)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Delete(domainIdParam string, groupIdParam string, expressionIdParam string) error

    // If a group MACAddressExpression with the expression-id is not already present, create a new MACAddressExpression. If it already exists, replace the existing MACAddressExpression.
    //
    // @param domainIdParam Domain ID (required)
    // @param groupIdParam Group ID (required)
    // @param expressionIdParam MACAddressExpression ID (required)
    // @param mACAddressExpressionParam (required)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Patch(domainIdParam string, groupIdParam string, expressionIdParam string, mACAddressExpressionParam model.MACAddressExpression) error
}
