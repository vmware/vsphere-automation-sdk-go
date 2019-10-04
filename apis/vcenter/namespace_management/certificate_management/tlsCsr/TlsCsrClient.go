/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: TlsCsr
 * Used by client-side stubs.
 */

package tlsCsr

import (
)

// The ``TlsCsr`` interface provides methods to generate certificate signing requests. **Warning:** This interface is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type TlsCsrClient interface {


    // Create a Certificate Signing Request for Kubernetes API Server. Certificate issued using this request can be used to update cluster configuration using {name UpdateSpec}. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param clusterParam Identifier for the cluster on which vSphere Namespaces are enabled.
    // The parameter must be an identifier for the resource type: ``ClusterComputeResource``.
    // @param specParam Specification for the new Certificate Signing Request.
    // @return Certificate Signing Request in PEM format.
    // @throws Error if the system reports an error while responding to the request.
    // @throws InvalidArgument if ``spec`` contain any errors.
    // @throws NotFound if cluster could not be located.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user does not have Namespaces.Manage privilege.
    Create(clusterParam string, specParam Spec) (string, error) 


    // Create a Certificate Signing Request used by NSX as a default, fallback certificate for Kubernetes Ingress objects. The certificate issued using this request can be used to update cluster configuration using {name UpdateSpec}. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param clusterParam Identifier for the cluster on which vSphere Namespaces are enabled.
    // The parameter must be an identifier for the resource type: ``ClusterComputeResource``.
    // @param specParam Specification for the new Certificate Signing Request.
    // @return Certificate Signing Request in PEM format.
    // @throws Error if the system reports an error while responding to the request.
    // @throws InvalidArgument if ``spec`` contain any errors.
    // @throws NotFound if cluster could not be located.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user does not have Namespaces.Manage privilege.
    CreateNcpDefaultIngressTls(clusterParam string, specParam Spec) (string, error) 

}
