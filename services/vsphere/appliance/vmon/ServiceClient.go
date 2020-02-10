/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Service
 * Used by client-side stubs.
 */

package vmon


// The ``Service`` interface provides methods to manage a single/set of services that are managed by vMon.
type ServiceClient interface {

    // Starts a service
    //
    // @param serviceParam identifier of the service to start
    // The parameter must be an identifier for the resource type: ``com.vmware.appliance.vmon.Service``.
    // @throws NotFound if the service associated with ``service`` does not exist.
    // @throws NotAllowedInCurrentState if the operation is denied in the current state of the service. If a stop or restart operation is in progress, the start operation will not be allowed.
    // @throws NotAllowedInCurrentState if start operation is issued on a service which has startup type ServiceStartupType#ServiceStartupType_DISABLED.
    // @throws TimedOut if any timeout occurs during the execution of the start operation. Timeout occurs when the service takes longer than StartTimeout to start.
    // @throws Error if any other error occurs during the execution of the operation.
	Start(serviceParam string) error

    // Stops a service
    //
    // @param serviceParam identifier of the service to stop
    // The parameter must be an identifier for the resource type: ``com.vmware.appliance.vmon.Service``.
    // @throws NotFound if the service associated with ``service`` does not exist.
    // @throws Error if any other error occurs during the execution of the operation.
    // @throws NotAllowedInCurrentState if the operation is denied in the current state of the service. If a stop operation is in progress, issuing another stop operation will lead to this error.
	Stop(serviceParam string) error

    // Restarts a service
    //
    // @param serviceParam identifier of the service to restart
    // The parameter must be an identifier for the resource type: ``com.vmware.appliance.vmon.Service``.
    // @throws NotFound if the service associated with ``service`` does not exist.
    // @throws TimedOut if any timeout occurs during the execution of the restart operation.
    // @throws NotAllowedInCurrentState if the operation is denied in the current state of the service. If a stop or start operation is in progress, issuing a restart operation will lead to this error.
    // @throws NotAllowedInCurrentState if a restart operation is issued on a service which has startup type ServiceStartupType#ServiceStartupType_DISABLED
    // @throws Error if any other error occurs during the execution of the operation.
	Restart(serviceParam string) error

    // Returns the state of a service.
    //
    // @param serviceParam identifier of the service whose state is being queried.
    // The parameter must be an identifier for the resource type: ``com.vmware.appliance.vmon.Service``.
    // @return Service Info structure.
    // @throws NotFound if the service associated with ``service`` does not exist.
    // @throws Error if any other error occurs during the execution of the operation.
	Get(serviceParam string) (ServiceInfo, error)

    // Updates the properties of a service.
    //
    // @param serviceParam identifier of the service whose properties are being updated.
    // The parameter must be an identifier for the resource type: ``com.vmware.appliance.vmon.Service``.
    // @param specParam Service Update specification.
    // @throws NotFound if the service associated with ``service`` does not exist.
    // @throws Error if any other error occurs during the execution of the operation.
    // @throws NotAllowedInCurrentState if the operation is denied in the current state of the service. If a start, stop or restart operation is in progress, update operation will fail with this error.
    // @throws NotAllowedInCurrentState if a request to set the ServiceUpdateSpec#startupType property of ``spec`` to ServiceStartupType#ServiceStartupType_DISABLED comes in for a service that is not in ServiceState#ServiceState_STOPPED state.
	Update(serviceParam string, specParam ServiceUpdateSpec) error

    // Lists details of services managed by vMon.
    // @return Map of service identifiers to service Info structures.
    // The key in the return value map will be an identifier for the resource type: ``com.vmware.appliance.vmon.Service``.
    // @throws Error if any error occurs during the execution of the operation.
	ListDetails() (map[string]ServiceInfo, error)
}
