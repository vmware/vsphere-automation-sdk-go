/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: CloudServiceVmcOnAwsTrafficGroup
 * Used by client-side stubs.
 */

package api

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-vmc-aws-integration/model"
)

type CloudServiceVmcOnAwsTrafficGroupClient interface {

    // Delete the Traffic Group, disassociate the prefix list with the traffic group.
    //
    // @param trafficGroupIdParam (required)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	DeleteTrafficGroup(trafficGroupIdParam string) error

    // Delete the specified association map for a traffic group.
    //
    // @param trafficGroupIdParam Traffic group id (required)
    // @param mapIdParam Association map id (required)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	DeleteTrafficGroupAssociationMap(trafficGroupIdParam string, mapIdParam string) error

    // Get the traffic group information by traffic group ID.
    //
    // @param trafficGroupIdParam (required)
    // @return com.vmware.model.TrafficGroup
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	GetTrafficGroup(trafficGroupIdParam string) (model.TrafficGroup, error)

    // @param trafficGroupIdParam Traffic group id (required)
    // @param mapIdParam Association map id (required)
    // @return com.vmware.model.TrafficGroupAssociationMap
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	GetTrafficGroupAssociationMap(trafficGroupIdParam string, mapIdParam string) (model.TrafficGroupAssociationMap, error)

    // Retrieve association maps for a traffic group given its ID.
    //
    // @param trafficGroupIdParam Traffic group id (required)
    // @return com.vmware.model.TrafficGroupAssociationMapsListResult
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	ListTrafficGroupAssociationMaps(trafficGroupIdParam string) (model.TrafficGroupAssociationMapsListResult, error)

    // A list of all traffic groups, which are associated with prefix lists.
    //
    // @param verboseParam Verbose info requested (optional, default to false)
    // @return com.vmware.model.TrafficGroupsListResult
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	ListTrafficGroups(verboseParam *bool) (model.TrafficGroupsListResult, error)

    // Perform an action on a Traffic Group
    //
    // @param trafficGroupIdParam (required)
    // @param actionMessageParam (required)
    // @param actionParam Action performed on the Traffic Group (required)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	TriggerTrafficGroupAction(trafficGroupIdParam string, actionMessageParam model.ActionMessage, actionParam string) error

    // This API is used to create or update a traffic group, a Prefix List is associated to the traffic group.
    //
    // @param trafficGroupIdParam traffic group id (required)
    // @param trafficGroupParam (required)
    // @return com.vmware.model.TrafficGroup
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	UpdateTrafficGroup(trafficGroupIdParam string, trafficGroupParam model.TrafficGroup) (model.TrafficGroup, error)

    // @param trafficGroupIdParam Traffic group id (required)
    // @param mapIdParam Association map id (required)
    // @param trafficGroupAssociationMapParam (required)
    // @return com.vmware.model.TrafficGroupAssociationMap
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	UpdateTrafficGroupAssociationMap(trafficGroupIdParam string, mapIdParam string, trafficGroupAssociationMapParam model.TrafficGroupAssociationMap) (model.TrafficGroupAssociationMap, error)
}
