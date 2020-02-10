/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Attestation
 * Used by client-side stubs.
 */

package trust_authority_hosts


// The ``Attestation`` interface contains information necessary to connect to the hosts running Attestation Service. This interface was added in vSphere API 7.0.0.
type AttestationClient interface {

    // Returns the connection info about the Attestation Service running on the specified host. This method was added in vSphere API 7.0.0.
    //
    // @param hostParam \\\\@{link com.vmware.vcenter.Host} id.
    // The parameter must be an identifier for the resource type: ``HostSystem``.
    // @return The AttestationInfo instance which contains the information necessary to connect to the Attestation Service.
    // @throws Error if service's TLS certificate chain is not valid.
    // @throws NotFound if ``host`` doesn't match to any Host.
    // @throws ResourceInaccessible if connection to ``host`` failed.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``TrustedAdmin.ReadTrustedHosts``.
    // * The resource ``HostSystem`` referenced by the parameter ``host`` requires ``System.View``.
	Get(hostParam string) (AttestationInfo, error)

    // Returns a list of the hosts running a Attestation Service matching the specified AttestationFilterSpec. This method was added in vSphere API 7.0.0.
    //
    // @param specParam Return details about Attestation Services matching the filter.
    // If {\\\\@term.unset} return all registered Attestation Services.
    // @param projectionParam The type of the returned summary - brief, normal, or full.
    // If {\\\\@term.unset} a normal projection will be used.
    // @return List of AttestationSummary of Attestation Services.
    // @throws Error if there is a generic error.
    // @throws InvalidArgument if the response data will exceed the message limit.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``TrustedAdmin.ReadTrustedHosts``.
    // * The resource ``HostSystem`` referenced by the property AttestationFilterSpec#hosts requires ``System.View``.
    // * The resource ``ClusterComputeResource`` referenced by the property AttestationFilterSpec#clusters requires ``System.View``.
	List(specParam *AttestationFilterSpec, projectionParam *AttestationSummaryType) ([]AttestationSummary, error)
}
