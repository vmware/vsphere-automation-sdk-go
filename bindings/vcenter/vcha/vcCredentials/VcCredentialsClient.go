/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: VcCredentials
 * Used by client-side stubs.
 */

package vcCredentials

import (
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/bindings/vcenter/vcha"
)

// The ``VcCredentials`` interface provides methods to validate the credentials of the management vCenter of a vCenter High Availability (VCHA) node.
type VcCredentialsClient interface {


    // Validates the credentials of the management vCenter server of the active node of a VCHA cluster.
    //
    // @param specParam Structure with information to connect to the management vCenter server of the Active Node in the VCHA cluster.
    // @throws InvalidArgument If the credentials provided for authenticating with the active node's management vCenter server are invalid.
    // @throws Unauthorized If the user has insufficient privilege to perform the operation. Operation execution requires the Global.VCServer privilege.
    // @throws UnverifiedPeer If the SSL certificate of the management vCenter server cannot be validated. 
    //  The value of the data property of errors.Error will be a class that contains all the properties defined in vcha.CertificateInfo.
    // @throws NotFound If the active virtual machine is not managed by the specified vCenter server for the active node.
    // @throws Error If any other error occurs.
    Validate(specParam vcha.CredentialsSpec) error 

}
