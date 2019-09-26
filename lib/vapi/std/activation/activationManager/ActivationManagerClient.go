/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: ActivationManager
 * Used by client-side stubs.
 */

package activationManager

import (
)

// **WARNING:** Use only as a sample. The API is experimental and subject to change in future versions. 
//
//  Activation tracking/management service. 
//
//  An activation describes a method invocation in the runtime.
type ActivationManagerClient interface {


    // Asks for cancellation of a running activation. Whether or not the cancellation request will have any effect depends on the implementation of the method that has to be canceled.
    //
    // @param activationIdParam activation identifier
    // @throws NotFound there is no activation with the specified id
    Cancel(activationIdParam string) error 

}
