/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Links
 * Used by client-side stubs.
 */

package hvc


// The ``Links`` interface provides methods to create, delete, get information, and list hybrid links between the local and foreign Platform Service Controller (PSC). Usage beyond VMware Cloud on AWS is not supported. **Warning:** This interface is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
type LinksClient interface {

    // Creates a new hybrid link between the local and foreign PSC. Usage beyond VMware Cloud on AWS is not supported. **Warning:** This method is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    //
    // @param specParam Specification for the new link to be created.
    // @return The identifier of the newly linked domain.
    // The return value will be an identifier for the resource type: ``com.vmware.vcenter.hvc.Links``.
    // @throws AlreadyExists If the link already exists.
    // @throws InvalidArgument If the values of any of the properties of the ``spec`` parameter are not valid.
    // @throws Unsupported If the PSC or the VC version is not supported.
    // @throws Unauthorized If the user is not authorized.
    // @throws UnverifiedPeer If the SSL certificate of the foreign PSC cannot be validated by comparing with the thumbprint provided in LinksCreateSpec#sslThumbprint or if LinksCreateSpec#sslThumbprint is null. The value of the {\\\\@link UnverifiedPeer#data) property will be a class that contains all the properties defined in LinksCertificateInfo.
    // @throws Error if the system reports an error while responding to the request.
	Create(specParam LinksCreateSpec) (string, error)

    // Deletes an existing hybrid link. Usage beyond VMware Cloud on AWS is not supported. **Warning:** This method is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    //
    // @param linkParam Identifier of the hybrid link.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.hvc.Links``.
    // @throws NotFound If the hybrid link associated with ``link`` does not exist.
    // @throws Unauthorized If the user is not authorized.
    // @throws Error if the system reports an error while responding to the request.
	Delete(linkParam string) error

    // Enumerates the list of registered hybrid links. Usage beyond VMware Cloud on AWS is not supported. **Warning:** This method is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    // @return The array of hybrid link information.
    // @throws Error if the system reports an error while responding to the request.
	List() ([]LinksSummary, error)
}
