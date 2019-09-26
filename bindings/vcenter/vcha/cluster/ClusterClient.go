/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Cluster
 * Used by client-side stubs.
 */

package cluster

import (
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/bindings/vcenter/vcha"
)

// The ``Cluster`` interface provides methods to deploy and undeploy a vCenter High Availability (VCHA) cluster, failover from the active VCHA node to the passive VCHA node, and retrieve the status of the VCHA cluster.
type ClusterClient interface {


    // Prepares, clones, and configures a VCHA cluster.
    //
    // @param specParam Contains the deploy specification for all three nodes of a VCHA cluster.
    // @throws InvalidArgument If the credentials provided for authenticating with the active node's management vCenter server are invalid.
    // @throws Unauthorized If the user has insufficient privilege to perform the operation. Operation execution requires the Global.VCServer privilege.
    // @throws UnverifiedPeer If the SSL certificate of the management vCenter server cannot be validated.
    //  The value of the data property of errors.Error will be a class that contains all the properties defined in vcha.CertificateInfo.
    // @throws Error If any other error occurs.
    Deploy(specParam DeploySpec) error 


    // Initiates failover from the active vCenter node to the passive node. 
    //
    //  For forced failover, Active node immediately initiates a failover. This may result into a data loss after failover. 
    //
    //  For planned failover, Active node flushes all the state to the Passive node, waits for the flush to complete before causing a failover. After the failover, Passive node starts without any data loss. 
    //
    //  A failover is allowed only in the following cases: 
    //
    // #. Cluster's mode is enabled and all cluster members are present.
    // #. Cluster's mode is maintenance and all cluster members are present.
    //
    // @param plannedParam If false, a failover is initiated immediately and may result in data loss.
    //  If true, a failover is initated after the Active node flushes its state to Passive and there is no data loss.
    // @throws Unauthorized If the user has insufficient privilege to perform the operation. Operation execution requires the Global.VCServer privilege.
    // @throws Error If any other error occurs.
    Failover(plannedParam bool) error 


    // Retrieves the status of a VCHA cluster.
    //
    // @param vcSpecParam Contains active node's management vCenter server credentials.
    // If null, then the active vCenter Server instance is assumed to be either self-managed or else in enhanced linked mode and managed by a linked vCenter Server instance.
    // @param partialParam If true, then return only the information that does not require connecting to the Active vCenter Server.
    //  If false or unset, then return all the information.
    // If null, then return all the information.
    // @return Info structure containing the VCHA configuration and health information.
    // @throws InvalidArgument If the credentials provided for authenticating with the active node's management vCenter server are invalid.
    // @throws Unauthorized If the user has insufficient privilege to perform the operation. 
    //
    // * If ``partial`` is false or unset, then the operation execution requires the Global.VCServer privilege.
    // * If ``partial`` is true, then the operation execution requires the System.Read privilege.
    // @throws UnverifiedPeer If the SSL certificate of the management vCenter server cannot be validated.
    //  The value of the data property of errors.Error will be a class that contains all the properties defined in vcha.CertificateInfo.
    // @throws Error If any other error occurs.
    Get(vcSpecParam *vcha.CredentialsSpec, partialParam *bool) (Info, error) 


    // Destroys the VCHA cluster and removes all VCHA specific information from the VCVA appliance. Optionally, the passive and witness node virtual machines will be deleted only if VCHA was deployed using automatic deployment. The active node in the cluster continues to run as a standalone VCVA appliance after the destroy operation has been performed. 
    //
    //  If the VCHA cluster is in a transition state and not configured, then the VCHA cluster specific information is removed.
    //
    // @param specParam Contains the undeploy specification for a VCHA cluster.
    // @throws InvalidArgument If the credentials provided for authenticating with the active node's management vCenter server are invalid.
    // @throws NotFound If the passive virtual machine is not managed by the specified vCenter server.
    // @throws NotFound If the witness virtual machine is not managed by the specified vCenter server.
    // @throws Unauthorized If the user has insufficient privilege to perform the operation. Operation execution requires the Global.VCServer privilege.
    // @throws UnverifiedPeer If the SSL certificate of the management vCenter server cannot be validated.
    //  The value of the data property of errors.Error will be a class that contains all the properties defined in vcha.CertificateInfo.
    // @throws Error If any other error occurs.
    Undeploy(specParam UndeploySpec) error 

}
