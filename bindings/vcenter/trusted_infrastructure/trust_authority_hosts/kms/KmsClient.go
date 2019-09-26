/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Kms
 * Used by client-side stubs.
 */

package kms

import (
)

// The ``Kms`` interface contains information necessary to connect to the hosts running Key Provider Service.
type KmsClient interface {


    // Returns the connection info about the Key Provider Service running on the specified host.
    //
    // @param hostParam \\\\@{link com.vmware.vcenter.Host} id.
    // The parameter must be an identifier for the resource type: ``HostSystem``.
    // @return The Info instance which contains the information necessary to connect to the Key Provider Service.
    // @throws Error if service's TLS certificate chain is not valid.
    // @throws NotFound if ``host`` doesn't match to any Host.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws ResourceInaccessible if connection to ``host`` failed.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``TrustedAdmin.ReadTrustedHosts``.
    // * The resource ``HostSystem`` referenced by the parameter ``host`` requires ``System.View``.
    Get(hostParam string) (Info, error) 


    // Returns a list of the hosts running a Key Provider Service matching the specified FilterSpec.
    //
    // @param specParam Return details about Key Provider Services matching the filter.
    // If {\\\\@term.unset} return all registered Key Provider Services.
    // @param projectionParam The type of the returned summary - brief, normal, or full.
    // If {\\\\@term.unset} a normal projection will be used.
    // @return List of Summary of Key Provider Services.
    // @throws Error if there is a generic error.
    // @throws InvalidArgument if the response data will exceed the message limit.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``TrustedAdmin.ReadTrustedHosts``.
    // * The resource ``HostSystem`` referenced by the property FilterSpec#hosts requires ``System.View``.
    // * The resource ``ClusterComputeResource`` referenced by the property FilterSpec#clusters requires ``System.View``.
    List(specParam *FilterSpec, projectionParam *SummaryType) ([]Summary, error) 

}
