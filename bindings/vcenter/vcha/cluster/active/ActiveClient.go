/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Active
 * Used by client-side stubs.
 */

package active

import (
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/bindings/vcenter/vcha"
)

// The ``Active`` interface provides methods to get information related to the active vCenter High Availability (VCHA) node.
type ActiveClient interface {


    // Retrieves information about the active node of a VCHA cluster.
    //
    // @param vcSpecParam Contains active node's management vCenter server credentials.
    // If null, then the active vCenter Server instance is assumed to be either self-managed or else in enhanced linked mode and managed by a linked vCenter Server instance.
    // @param partialParam If true, then return only the information that does not require connecting to the Active vCenter Server. 
    //  If false or unset, then return all the information.
    // If null, then return all the information.
    // @return Info Information about the VCHA network and placement of the active node.
    // @throws InvalidArgument If the credentials provided for authentincating with the active node's management vCenter server are invalid.
    // @throws Unauthorized If the user has insufficient privilege to perform the operation. 
    //
    // * If ``partial`` is false or unset, then the operation execution requires the Global.VCServer privilege.
    // * If ``partial`` is true, then the operation execution requires the System.Read privilege.
    // @throws UnverifiedPeer If the SSL certificate of the management vCenter server cannot be validated. 
    //  The value of the data property of errors.Error will be a class that contains all the properties defined in vcha.CertificateInfo.
    // @throws InvalidElementConfiguration If the active node is on more than one datastore.
    // @throws NotFound If the active virtual machine is not managed by the specified vCenter server for the active node.
    // @throws Error If the management interface IP address assignment is not static.
    // @throws Error If any other error occurs.
    Get(vcSpecParam *vcha.CredentialsSpec, partialParam *bool) (Info, error) 

}
