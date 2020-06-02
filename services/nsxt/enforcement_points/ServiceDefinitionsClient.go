/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: ServiceDefinitions
 * Used by client-side stubs.
 */

package enforcement_points

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt/model"
)

type ServiceDefinitionsClient interface {

    // Create a Service Definition on given enforcement point.
    //
    // @param enforcementPointIdParam Enforcement point id (required)
    // @param serviceDefinitionParam (required)
    // @return com.vmware.nsx_policy.model.ServiceDefinition
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Create(enforcementPointIdParam string, serviceDefinitionParam model.ServiceDefinition) (model.ServiceDefinition, error)

    // Delete an existing Service Definition on the given enforcement point.
    //
    // @param enforcementPointIdParam Enforcement point id (required)
    // @param serviceDefinitionIdParam Id of service definition (required)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Delete(enforcementPointIdParam string, serviceDefinitionIdParam string) error

    // Read Service Definition with given service-definition-id.
    //
    // @param enforcementPointIdParam Enforcement point id (required)
    // @param serviceDefinitionIdParam Id of service definition (required)
    // @return com.vmware.nsx_policy.model.ServiceDefinition
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get(enforcementPointIdParam string, serviceDefinitionIdParam string) (model.ServiceDefinition, error)

    // List all Service Definitions registered on given enforcement point.
    //
    // @param enforcementPointIdParam Enforcement point id (required)
    // @return com.vmware.nsx_policy.model.ServiceInsertionServiceListResult
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	List(enforcementPointIdParam string) (model.ServiceInsertionServiceListResult, error)

    // Update an existing Service Definition on the given enforcement point.
    //
    // @param enforcementPointIdParam Enforcement point id (required)
    // @param serviceDefinitionIdParam Id of service definition (required)
    // @param serviceDefinitionParam (required)
    // @return com.vmware.nsx_policy.model.ServiceDefinition
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Update(enforcementPointIdParam string, serviceDefinitionIdParam string, serviceDefinitionParam model.ServiceDefinition) (model.ServiceDefinition, error)
}
