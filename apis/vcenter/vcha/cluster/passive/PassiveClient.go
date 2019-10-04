/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Passive
 * Used by client-side stubs.
 */

package passive

import (
)

// The ``Passive`` interface provides methods to validate a passive's placement configuration and redeploy the passive node in a vCenter High Availability (VCHA) cluster.
type PassiveClient interface {


    // Validates the specified passive node's placement configuration.
    //
    // @param specParam Contains the passive node's placement specification.
    // @return CheckResult structure containing errors and warnings.
    // @throws InvalidArgument If the credentials provided for authentincating with the active node's management vCenter server are invalid.
    // @throws InvalidArgument If the specified resource spec is deemed invalid for the clone operation.
    // @throws UnverifiedPeer If the SSL certificate of the management vCenter server cannot be validated.
    //  The value of the data property of errors.Error will be a class that contains all the properties defined in vcha.CertificateInfo.
    // @throws NotFound If the active virtual machine is not managed by the specified vCenter server for the active node.
    // @throws InvalidElementConfiguration If the active node is on more than one datastore.
    // @throws NotAllowedInCurrentState If the clone operation is not allowed in the current state of the system.
    // @throws Unauthorized If the user has insufficient privilege to perform the operation. Operation execution requires the Global.VCServer privilege.
    // @throws Error If any other error occurs.
    Check(specParam CheckSpec) (CheckResult, error) 


    // Creates the passive node in a degraded cluster with node location information and pre-existing VCHA cluster configuration from the active node.
    //
    // @param specParam Contains the passive node's redeploy specification.
    // @throws InvalidArgument If the credentials provided for authentincating with the active node's management vCenter server are invalid.
    // @throws Unauthorized If the user has insufficient privilege to perform the operation. Operation execution requires the Global.VCServer privilege.
    // @throws UnverifiedPeer If the SSL certificate of the management vCenter server cannot be validated.
    //  The value of the data property of errors.Error will be a class that contains all the properties defined in vcha.CertificateInfo.
    // @throws Error If any other error occurs.
    Redeploy(specParam RedeploySpec) error 

}
