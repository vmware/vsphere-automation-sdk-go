/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: LibraryItems
 * Used by client-side stubs.
 */

package libraryItems

import (
)

// The ``LibraryItems`` interface provides methods to deploy virtual machines from library items containing virtual machine templates, as well as methods to create library items containing virtual machine templates. The ``LibraryItems`` interface also provides an operation to retrieve information about the template contained in the library item.
type LibraryItemsClient interface {


    // Creates a library item in content library from a virtual machine. This method creates a library item in content library whose content is a virtual machine template created from the source virtual machine, using the supplied create specification. The virtual machine template is stored in a newly created library item.
    //
    // @param specParam  information used to create the library item from the source virtual machine.
    // @return Identifier of the newly created library item.
    // The return value will be an identifier for the resource type: ``com.vmware.content.library.Item``.
    // @throws AlreadyExists  if an entity with the name specified by CreateSpec#name already exists in the folder specified by CreatePlacementSpec#folder.
    // @throws Error  if the system reports an error while responding to the request.
    // @throws InvalidArgument  if CreateSpec contains invalid arguments.
    // @throws NotAllowedInCurrentState  if the operation cannot be performed because of the source virtual machine's current state.
    // @throws NotFound  if the source virtual machine specified by CreateSpec#sourceVm does not exist.
    // @throws NotFound  if the library specified by CreateSpec#library does not exist.
    // @throws ResourceInaccessible  if there was an error accessing a file from the source virtual machine.
    // @throws ResourceInUse  if the source virtual machine is busy.
    // @throws ServiceUnavailable  if any of the services involved in the method are unavailable.
    // @throws UnableToAllocateResource  if any of the resources needed to create the virtual machine template could not be allocated.
    // @throws Unauthenticated  if the user that requested the method cannot be authenticated.
    // @throws Unauthorized  if the user that requested the method is not authorized to perform the method.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``System.Read``.
    // * The resource ``com.vmware.content.Library`` referenced by the property CreateSpec#library requires ``ContentLibrary.AddLibraryItem``.
    // * The resource ``VirtualMachine`` referenced by the property CreateSpec#sourceVm requires ``System.Read``.
    // * The resource ``Datastore`` referenced by the property CreateSpecVmHomeStorage#datastore requires ``System.Read``.
    // * The resource ``com.vmware.spbm.StorageProfile`` referenced by the property CreateSpecVmHomeStoragePolicy#policy requires ``System.Read``.
    // * The resource ``Datastore`` referenced by the property CreateSpecDiskStorage#datastore requires ``System.Read``.
    // * The resource ``com.vmware.spbm.StorageProfile`` referenced by the property CreateSpecDiskStoragePolicy#policy requires ``System.Read``.
    // * The resource ``com.vmware.vcenter.vm.hardware.Disk`` referenced by the map key of property CreateSpec#diskStorageOverrides requires ``System.Read``.
    // * The resource ``Folder`` referenced by the property CreatePlacementSpec#folder requires ``System.Read``.
    // * The resource ``ResourcePool`` referenced by the property CreatePlacementSpec#resourcePool requires ``System.Read``.
    // * The resource ``HostSystem`` referenced by the property CreatePlacementSpec#host requires ``System.Read``.
    // * The resource ``ClusterComputeResource`` referenced by the property CreatePlacementSpec#cluster requires ``System.Read``.
    Create(specParam CreateSpec) (string, error) 


    // Deploys a virtual machine as a copy of the source virtual machine template contained in the library item specified by ``template_library_item``. It uses the deployment specification in ``spec``. If DeploySpec#poweredOn and/or DeploySpec#guestCustomization are specified, the server triggers the power on and/or guest customization operations, which are executed asynchronously.
    //
    // @param templateLibraryItemParam  identifier of the content library item containing the source virtual machine template to be deployed.
    // The parameter must be an identifier for the resource type: ``com.vmware.content.library.Item``.
    // @param specParam  specification of how the virtual machine should be deployed.
    // @return Identifier of the deployed virtual machine.
    // The return value will be an identifier for the resource type: ``VirtualMachine``.
    // @throws AlreadyExists  if a virtual machine with the name specified by DeploySpec#name already exists.
    // @throws Error  if the system reports an error while responding to the request.
    // @throws InvalidArgument  if ``spec`` contains invalid arguments.
    // @throws NotAllowedInCurrentState  if either a specified host or a specified datastore is in an invalid state for the deployment, such as maintenance mode.
    // @throws NotFound  if the library item specified by ``template_library_item`` cannot be found.
    // @throws NotFound  if any resource specified by a property of the DeploySpec class, specified by ``spec`` cannot be found.
    // @throws ResourceInaccessible  if there was an error accessing the source virtual machine template contained in the library item specified by ``template_library_item``.
    // @throws ResourceInaccessible  if there an error accessing any of the resources specified in the ``spec``.
    // @throws ServiceUnavailable  if any of the services involved in the method are unavailable.
    // @throws UnableToAllocateResource  if there was an error in allocating any of the resources required by the method.
    // @throws Unauthenticated  if the user that requested the method cannot be authenticated.
    // @throws Unauthorized  if the user that requested the method is not authorized to perform the method.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``System.Read``.
    // * The resource ``com.vmware.content.library.Item`` referenced by the parameter ``template_library_item`` requires ``VirtualMachine.Provisioning.DeployTemplate``.
    // * The resource ``Datastore`` referenced by the property DeploySpecVmHomeStorage#datastore requires ``System.Read``.
    // * The resource ``com.vmware.spbm.StorageProfile`` referenced by the property DeploySpecVmHomeStoragePolicy#policy requires ``System.Read``.
    // * The resource ``Datastore`` referenced by the property DeploySpecDiskStorage#datastore requires ``System.Read``.
    // * The resource ``com.vmware.spbm.StorageProfile`` referenced by the property DeploySpecDiskStoragePolicy#policy requires ``System.Read``.
    // * The resource ``com.vmware.vcenter.vm.hardware.Disk`` referenced by the map key of property DeploySpec#diskStorageOverrides requires ``System.Read``.
    // * The resource ``Folder`` referenced by the property DeployPlacementSpec#folder requires ``System.Read``.
    // * The resource ``ResourcePool`` referenced by the property DeployPlacementSpec#resourcePool requires ``System.Read``.
    // * The resource ``HostSystem`` referenced by the property DeployPlacementSpec#host requires ``System.Read``.
    // * The resource ``ClusterComputeResource`` referenced by the property DeployPlacementSpec#cluster requires ``System.Read``.
    // * The resource ``com.vmware.vcenter.vm.hardware.Ethernet`` referenced by the map key of property HardwareCustomizationSpec#nics requires ``System.Read``.
    // * The resource ``Network`` referenced by the property EthernetUpdateSpec#network requires ``System.Read``.
    // * The resource ``com.vmware.vcenter.vm.hardware.Disk`` referenced by the property HardwareCustomizationSpec#disksToRemove requires ``System.Read``.
    // * The resource ``com.vmware.vcenter.vm.hardware.Disk`` referenced by the map key of property HardwareCustomizationSpec#disksToUpdate requires ``System.Read``.
    Deploy(templateLibraryItemParam string, specParam DeploySpec) (string, error) 


    // Returns information about a virtual machine template contained in the library item specified by ``template_library_item``
    //
    // @param templateLibraryItemParam  identifier of the library item containing the virtual machine template.
    // The parameter must be an identifier for the resource type: ``com.vmware.content.library.Item``.
    // @return Information about the virtual machine template item contained in the library item.
    // If null, the library item specified by ``template_library_item`` does not contain a virtual machine template.
    // @throws Error  if the system reports an error while responding to the request.
    // @throws NotFound  if the library item could not be found.
    // @throws ResourceInaccessible  if the virtual machine template's configuration state cannot be accessed.
    // @throws ServiceUnavailable  if any of the services involved in the method are unavailable.
    // @throws Unauthenticated  if the user that requested the method cannot be authenticated.
    // @throws Unauthorized  if the user that requested the method is not authorized to perform the method.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``System.Read``.
    // * The resource ``com.vmware.content.library.Item`` referenced by the parameter ``template_library_item`` requires ``System.Read``.
    Get(templateLibraryItemParam string) (*Info, error) 

}
