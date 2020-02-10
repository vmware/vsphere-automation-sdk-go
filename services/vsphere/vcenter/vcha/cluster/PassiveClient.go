/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Passive
 * Used by client-side stubs.
 */

package cluster


// The ``Passive`` interface provides methods to validate a passive's placement configuration and redeploy the passive node in a vCenter High Availability (VCHA) cluster. This interface was added in vSphere API 6.7.1.
type PassiveClient interface {

    // Validates the specified passive node's placement configuration. This method was added in vSphere API 6.7.1.
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
	Check(specParam PassiveCheckSpec) (PassiveCheckResult, error)

    // Creates the passive node in a degraded cluster with node location information and pre-existing VCHA cluster configuration from the active node. This method was added in vSphere API 6.7.1.
    //
    // @param specParam Contains the passive node's redeploy specification.
    // @throws InvalidArgument If the credentials provided for authentincating with the active node's management vCenter server are invalid.
    // @throws Unauthorized If the user has insufficient privilege to perform the operation. Operation execution requires the Global.VCServer privilege.
    // @throws UnverifiedPeer If the SSL certificate of the management vCenter server cannot be validated.
    //  The value of the data property of errors.Error will be a class that contains all the properties defined in vcha.CertificateInfo.
    // @throws Error If any other error occurs.
	Redeploy(specParam PassiveRedeploySpec) error
}
