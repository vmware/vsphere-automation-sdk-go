/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: PathExpressions
 * Used by client-side stubs.
 */

package groups

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/services/nsxt-gm/model"
)

type PathExpressionsClient interface {

    // It will add or remove the specified members having path for a given expression of a group.
    //
    // @param domainIdParam (required)
    // @param groupIdParam (required)
    // @param expressionIdParam (required)
    // @param groupMemberListParam (required)
    // @param actionParam Add or Remove group members. (required)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Create(domainIdParam string, groupIdParam string, expressionIdParam string, groupMemberListParam model.GroupMemberList, actionParam string) error

    // Delete Group Path Expression
    //
    // @param domainIdParam Domain ID (required)
    // @param groupIdParam Group ID (required)
    // @param expressionIdParam PathExpression ID (required)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Delete(domainIdParam string, groupIdParam string, expressionIdParam string) error

    // If a group path_expression with the expression-id is not already present, create a new pathexpresison. If it already exists, replace the existing pathexpression.
    //
    // @param domainIdParam Domain ID (required)
    // @param groupIdParam Group ID (required)
    // @param expressionIdParam PathExpression ID (required)
    // @param pathExpressionParam (required)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Patch(domainIdParam string, groupIdParam string, expressionIdParam string, pathExpressionParam model.PathExpression) error
}
