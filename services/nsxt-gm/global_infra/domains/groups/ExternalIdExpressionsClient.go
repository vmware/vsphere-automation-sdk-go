/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: ExternalIdExpressions
 * Used by client-side stubs.
 */

package groups

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/services/nsxt-gm/model"
)

type ExternalIdExpressionsClient interface {

    // It will add or remove the specified members having external ID for a given expression of a group.
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

    // Delete Group External ID Expression
    //
    // @param domainIdParam Domain ID (required)
    // @param groupIdParam Group ID (required)
    // @param expressionIdParam ExternalIDExpression ID (required)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Delete(domainIdParam string, groupIdParam string, expressionIdParam string) error

    // If a group ExternalIDexpression with the expression-id is not already present, create a new ExternalIDexpresison. If it already exists, replace the existing ExternalIDexpression.
    //
    // @param domainIdParam Domain ID (required)
    // @param groupIdParam Group ID (required)
    // @param expressionIdParam ExternalIDExpression ID (required)
    // @param externalIDExpressionParam (required)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Patch(domainIdParam string, groupIdParam string, expressionIdParam string, externalIDExpressionParam model.ExternalIDExpression) error
}
