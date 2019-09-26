/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Services
 * Used by client-side stubs.
 */

package services

import (
)

// The ``Service`` interface provides methods to manage a single/set of appliance services.
type ServicesClient interface {


    // Starts a service
    //
    // @param serviceParam identifier of the service to start
    // The parameter must be an identifier for the resource type: ``com.vmware.appliance.services``.
    // @throws NotFound if the service associated with ``service`` does not exist.
    // @throws NotAllowedInCurrentState if the operation is denied in the current state of the service. If a stop or restart operation is in progress, the start operation will not be allowed.
    // @throws NotAllowedInCurrentState if start operation is issued on a service which has startup type null.
    // @throws TimedOut if any timeout occurs during the execution of the start operation. Timeout occurs when the service takes longer than StartTimeout to start.
    // @throws Error if any other error occurs during the execution of the operation.
    Start(serviceParam string) error 


    // Stops a service
    //
    // @param serviceParam identifier of the service to stop
    // The parameter must be an identifier for the resource type: ``com.vmware.appliance.services``.
    // @throws NotFound if the service associated with ``service`` does not exist.
    // @throws Error if any other error occurs during the execution of the operation.
    // @throws NotAllowedInCurrentState if the operation is denied in the current state of the service. If a stop operation is in progress, issuing another stop operation will lead to this error.
    Stop(serviceParam string) error 


    // Restarts a service
    //
    // @param serviceParam identifier of the service to restart
    // The parameter must be an identifier for the resource type: ``com.vmware.appliance.services``.
    // @throws NotFound if the service associated with ``service`` does not exist.
    // @throws TimedOut if any timeout occurs during the execution of the restart operation.
    // @throws NotAllowedInCurrentState if the operation is denied in the current state of the service. If a stop or start operation is in progress, issuing a restart operation will lead to this error.
    // @throws NotAllowedInCurrentState if a restart operation is issued on a service which has startup type null
    // @throws Error if any other error occurs during the execution of the operation.
    Restart(serviceParam string) error 


    // Returns the state of a service.
    //
    // @param serviceParam identifier of the service whose state is being queried.
    // The parameter must be an identifier for the resource type: ``com.vmware.appliance.services``.
    // @return Service Info structure.
    // @throws NotFound if the service associated with ``service`` does not exist.
    // @throws Error if any other error occurs during the execution of the operation.
    Get(serviceParam string) (Info, error) 


    // Lists details of vCenter services.
    // @return Map of service identifiers to service Info structures.
    // The key in the return value map will be an identifier for the resource type: ``com.vmware.appliance.services``.
    // @throws Error if any error occurs during the execution of the operation.
    List() (map[string]Info, error) 

}
