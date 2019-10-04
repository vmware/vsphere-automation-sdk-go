/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: LibraryItem
 * Used by client-side stubs.
 */

package libraryItem

import (
)

// The ``LibraryItem`` interface provides methods to deploy virtual machines and virtual appliances from library items containing Open Virtualization Format (OVF) packages in content library, as well as methods to create library items in content library from virtual machines and virtual appliances. 
//
//  To deploy a virtual machine or a virtual appliance from a library item: 
//
// #. Create a DeploymentTarget to specify the target deployment type and target deployment designation.
// #. Create a ResourcePoolDeploymentSpec to specify the parameters for the target deployment.
// #. Use the ``deploy()`` method with the created target and parameter specifications, along with the identifier of the specified source content library item. See LibraryItem#deploy.
//
//  
//
//  
//
//  To create a library item in content library from a virtual machine or virtual appliance: 
//
// #. Create a DeployableIdentity to specify the source virtual machine or virtual appliance to be used as the OVF template source.
// #. Create a CreateTarget to specify the target library and library item.
// #. Create a CreateSpec to specify the settings for the OVF package to be created.
// #. Use the ``create()`` method with the created target and parameter specifications, along with the specified source entity. See LibraryItem#create.
//
//  
type LibraryItemClient interface {


    // Deploys an OVF package stored in content library to a newly created virtual machine or virtual appliance. 
    //
    //  This method deploys an OVF package which is stored in the library item specified by ``ovf_library_item_id``. It uses the deployment specification in ``deployment_spec`` to deploy the OVF package to the location specified by ``target``. 
    //
    // @param clientTokenParam  Client-generated token used to retry a request if the client fails to get a response from the server. If the original request succeeded, the result of that request will be returned, otherwise the operation will be retried.
    // If null, the server will create a token.
    // @param ovfLibraryItemIdParam  Identifier of the content library item containing the OVF package to be deployed.
    // The parameter must be an identifier for the resource type: ``com.vmware.content.library.Item``.
    // @param targetParam  Specification of the deployment target.
    // @param deploymentSpecParam  Specification of how the OVF package should be deployed to the target.
    // @return Information about the success or failure of the method, along with the details of the result or failure.
    // @throws InvalidArgument  if ``target`` contains invalid arguments.
    // @throws InvalidArgument  if ``deployment_spec`` contains invalid arguments or has properties that are inconsistent with ``target``.
    // @throws NotFound  if the library item specified by ``ovf_library_item_id`` does not exist.
    // @throws NotFound  if any resource specified by a property of the DeploymentTarget class, specified by ``target``, does not exist.
    // @throws ResourceInaccessible  if there was an error accessing the OVF package stored in the library item specified by ``ovf_library_item_id``.
    // @throws Unauthorized  if you do not have all of the privileges described as follows : 
    //
    // * Method execution requires VirtualMachine.Config.AddNewDisk if the OVF descriptor has a disk drive (type 17) section.
    // * Method execution requires VirtualMachine.Config.AdvancedConfig if the OVF descriptor has an ExtraConfig section.
    // * Method execution requires Extension.Register for specified resource group if the OVF descriptor has a vServiceDependency section.
    // * Method execution requires Network.Assign for target network if specified.
    // * Method execution requires Datastore.AllocateSpace for target datastore if specified.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``System.Read``.
    // * The resource ``com.vmware.content.library.Item`` referenced by the parameter ``ovf_library_item_id`` requires ``System.Read``.
    // * The resource ``HostSystem`` referenced by the property DeploymentTarget#hostId requires ``System.Read``.
    // * The resource ``Network`` referenced by the map value of property ResourcePoolDeploymentSpec#networkMappings requires ``System.Read``.
    // * The resource ``StorageProfile`` referenced by the property ResourcePoolDeploymentSpec#storageProfileId requires ``System.Read``.
    // * The resource ``Datastore`` referenced by the property ResourcePoolDeploymentSpec#defaultDatastoreId requires ``System.Read``.
    // * The resource ``ResourcePool`` referenced by the property DeploymentTarget#resourcePoolId requires ``VApp.Import``.
    // * The resource ``Folder`` referenced by the property DeploymentTarget#folderId requires ``VApp.Import``.
    Deploy(clientTokenParam *string, ovfLibraryItemIdParam string, targetParam DeploymentTarget, deploymentSpecParam ResourcePoolDeploymentSpec) (DeploymentResult, error) 


    // Queries an OVF package stored in content library to retrieve information to use when deploying the package. See LibraryItem#deploy. 
    //
    //  This method retrieves information from the descriptor of the OVF package stored in the library item specified by ``ovf_library_item_id``. The information returned by the method can be used to populate the deployment specification (see ResourcePoolDeploymentSpec when deploying the OVF package to the deployment target specified by ``target``. 
    //
    // @param ovfLibraryItemIdParam  Identifier of the content library item containing the OVF package to query.
    // The parameter must be an identifier for the resource type: ``com.vmware.content.library.Item``.
    // @param targetParam  Specification of the deployment target.
    // @return Information that can be used to populate the deployment specification (see ResourcePoolDeploymentSpec) when deploying the OVF package to the deployment target specified by ``target``.
    // @throws InvalidArgument  if ``target`` contains invalid arguments.
    // @throws NotFound  if the library item specified by ``ovf_library_item_id`` does not exist.
    // @throws NotFound  if any resource specified by a property of the DeploymentTarget class, specified by ``target``, does not exist.
    // @throws ResourceInaccessible  if there was an error accessing the OVF package at the specified ``ovf_library_item_id``.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``System.Read``.
    // * The resource ``com.vmware.content.library.Item`` referenced by the parameter ``ovf_library_item_id`` requires ``System.Read``.
    // * The resource ``ResourcePool`` referenced by the property DeploymentTarget#resourcePoolId requires ``System.Read``.
    // * The resource ``HostSystem`` referenced by the property DeploymentTarget#hostId requires ``System.Read``.
    // * The resource ``Folder`` referenced by the property DeploymentTarget#folderId requires ``System.Read``.
    Filter(ovfLibraryItemIdParam string, targetParam DeploymentTarget) (OvfSummary, error) 


    // Creates a library item in content library from a virtual machine or virtual appliance. 
    //
    //  This method creates a library item in content library whose content is an OVF package derived from a source virtual machine or virtual appliance, using the supplied create specification. The OVF package may be stored as in a newly created library item or in an in an existing library item. For an existing library item whose content is updated by this method, the original content is overwritten. 
    //
    // @param clientTokenParam  Client-generated token used to retry a request if the client fails to get a response from the server. If the original request succeeded, the result of that request will be returned, otherwise the operation will be retried.
    // If null, the server will create a token.
    // @param sourceParam  Identifier of the virtual machine or virtual appliance to use as the source.
    // @param targetParam  Specification of the target content library and library item.
    // @param createSpecParam  Information used to create the OVF package from the source virtual machine or virtual appliance.
    // @return Information about the success or failure of the method, along with the details of the result or failure.
    // @throws InvalidArgument  if ``create_spec`` contains invalid arguments.
    // @throws InvalidArgument  if ``source`` describes an unexpected resource type.
    // @throws NotFound  if the virtual machine or virtual appliance specified by ``source`` does not exist.
    // @throws NotFound  if the library or library item specified by ``target`` does not exist.
    // @throws NotAllowedInCurrentState  if the operation cannot be performed because of the specified virtual machine or virtual appliance's current state. For example, if the virtual machine configuration information is not available, or if the virtual appliance is running.
    // @throws ResourceInaccessible  if there was an error accessing a file from the source virtual machine or virtual appliance.
    // @throws ResourceBusy  if the specified virtual machine or virtual appliance is busy.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``System.Read``.
    // * The resource ``null`` or ``null`` referenced by the property DeployableIdentity#id requires ``VApp.Export``.
    // * The resource ``com.vmware.content.Library`` referenced by the property CreateTarget#libraryId requires ``ContentLibrary.AddLibraryItem``.
    // * The resource ``com.vmware.content.library.Item`` referenced by the property CreateTarget#libraryItemId requires ``System.Read``.
    Create(clientTokenParam *string, sourceParam DeployableIdentity, targetParam CreateTarget, createSpecParam CreateSpec) (CreateResult, error) 

}
