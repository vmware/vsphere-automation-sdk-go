/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: SupervisorServices
 * Used by client-side stubs.
 */

package supervisorServices

import (
)

// The ``SupervisorServices`` interface provides methods to set the desired state of Supervisor Services, as well as list them. Supervisor Services are extensions to the vSphere Namespace Supervisor, often provided by 3rd party vendors. They run in a highly privileged mode, and that should be considered before enabling them. **Warning:** This interface is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type SupervisorServicesClient interface {


    // Update the desired state for a Supervisor Service. The desired state is stored and updated in the vCenter Namespaces Supervisor Kubernetes layer. It may also be updated directly within Kubernetes, where the state is authoratively kept. Desired state is a combination of the enabled and version fields. The desired state is set and will be asynchronously remediated. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param clusterParam Identifier for the cluster on which to enable the service.
    // The parameter must be an identifier for the resource type: ``ClusterComputeResource``.
    // @param serviceIDParam ID of the Service that is being changed.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.namespaces.SupervisorService``.
    // @param specParam Specification for the service and its desired state.
    // @throws Error if the system reports an error while responding to the request, e.g. if the Kubernetes cluster is unhealthy or can't be reached.
    // @throws InvalidArgument if ``spec`` contain any errors or if an invalid name is specified.
    // @throws NotFound if ``cluster`` is not registered on this vCenter server.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the Namespaces.Manage privilege.
    Set(clusterParam string, serviceIDParam string, specParam SetSpec) error 

}
