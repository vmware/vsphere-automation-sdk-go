/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Projects
 * Used by client-side stubs.
 */

package projects

import (
)

// The ``Projects`` interface provides methods for managing the lifecycle of Harbor project that stores and distributes container repositories and images. **Warning:** This interface is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type ProjectsClient interface {


    // Creates a project in a Harbor registry using the supplied specification. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param registryParam Identifier of the Registry.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.content.Registry``.
    // @param specParam Information used to create the Harbor project.
    // @return Identifier of the newly created Harbor project.
    // The return value will be an identifier for the resource type: ``com.vmware.vcenter.content.Registry.Harbor.Project``.
    // @throws InvalidArgument if ``spec`` contains any error.
    // @throws NotAllowedInCurrentState if Harbor registry is being deleted.
    // @throws NotFound if a registry specified by ``registry`` could not be found.
    // @throws AlreadyExists if a project with the same name already exists in the registry.
    // @throws Error if the system reports an error while responding to the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user that requested the method is not authorized to perform the method.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``System.Read``.
    // * The resource ``com.vmware.vcenter.content.Registry`` referenced by the parameter ``registry`` requires ``ContentLibrary.CreateHarborProject``.
    Create(registryParam string, specParam CreateSpec) (string, error) 


    // Deletes the specified project from Harbor registry, remove the images in the project from Harbor's storage if possible. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param registryParam Identifier of the registry.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.content.Registry``.
    // @param projectParam Identifier of the Harbor project.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.content.Registry.Harbor.Project``.
    // @throws NotAllowedInCurrentState if Harbor registry is being deleted.
    // @throws NotFound if ``registry`` or ``project`` cannot be found.
    // @throws Error if the system reports an error while responding to the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user that requested the method is not authorized to perform the method.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``System.Read``.
    // * The resource ``com.vmware.vcenter.content.Registry`` referenced by the parameter ``registry`` requires ``ContentLibrary.DeleteHarborProject``.
    // * The resource ``com.vmware.vcenter.content.Registry.Harbor.Project`` referenced by the parameter ``project`` requires ``ContentLibrary.DeleteHarborProject``.
    Delete(registryParam string, projectParam string) error 


    // Returns detailed information about the specified Harbor project. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param registryParam Identifier of the registry.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.content.Registry``.
    // @param projectParam Identifier of the Harbor project.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.content.Registry.Harbor.Project``.
    // @return Detailed information about the specified Harbor project.
    // @throws NotFound if ``registry`` or ``project`` cannot be found.
    // @throws Error if the system reports an error while responding to the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user that requested the method is not authorized to perform the method.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``System.Read``.
    // * The resource ``com.vmware.vcenter.content.Registry`` referenced by the parameter ``registry`` requires ``System.Read``.
    // * The resource ``com.vmware.vcenter.content.Registry.Harbor.Project`` referenced by the parameter ``project`` requires ``System.Read``.
    Get(registryParam string, projectParam string) (Info, error) 


    // Returns basic information of all projects in a Harbor registry. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param registryParam Identifier of the registry.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.content.Registry``.
    // @return The list of summary information of all Harbor projects.
    // @throws NotFound if ``registry`` cannot be found.
    // @throws Error if the system reports an error while responding to the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user is not a member of the Administrators
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``System.Read``.
    // * The resource ``com.vmware.vcenter.content.Registry`` referenced by the parameter ``registry`` requires ``System.Read``.
    List(registryParam string) ([]Summary, error) 


    // Remove all repositories, images and members in the project. Storage space of deleted images in the project will be reclaimed through next scheduled Harbor registry garbage collection. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param registryParam Registry identifier.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.content.Registry``.
    // @param projectParam Identifier of the Harbor project.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.content.Registry.Harbor.Project``.
    // @throws NotFound if ``registry`` or ``project`` cannot be found.
    // @throws NotAllowedInCurrentState if Harbor registry is being deleted or the project is not in ConfigStatus#ConfigStatus_READY status.
    // @throws Error if the system reports an error while responding to the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user that requested the method is not authorized to perform the method.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``System.Read``.
    // * The resource ``com.vmware.vcenter.content.Registry`` referenced by the parameter ``registry`` requires ``ContentLibrary.PurgeHarborProject``.
    // * The resource ``com.vmware.vcenter.content.Registry.Harbor.Project`` referenced by the parameter ``project`` requires ``ContentLibrary.PurgeHarborProject``.
    Purge(registryParam string, projectParam string) error 

}
