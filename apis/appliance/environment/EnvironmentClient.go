/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Environment
 * Used by client-side stubs.
 */

package environment

import (
)

// The ``Environment`` interface provides methods to get and set appliance environment.
type EnvironmentClient interface {


    // Sets the properties of the appliance environment.
    //
    // @param configParam Structure containing the values of the Environment.
    // @throws Error if any error occurs during the execution of the operation.
    Set(configParam Config) error 


    // Gets the properties of the appliance environment.
    // @return Structure containing the values of the Environment.
    // @throws Error if any error occurs during the execution of the operation.
    Get() (Info, error) 

}
