/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: LibraryItems
 * Used by client-side stubs.
 */

package vm_template


// The ``LibraryItems`` interface provides methods to deploy virtual machines from library items containing virtual machine templates, as well as methods to create library items containing virtual machine templates. The ``LibraryItems`` interface also provides an operation to retrieve information about the template contained in the library item. This interface was added in vSphere API 6.8.
type LibraryItemsClient interface {

    // Creates a library item in content library from a virtual machine. This method creates a library item in content library whose content is a virtual machine template created from the source virtual machine, using the supplied create specification. The virtual machine template is stored in a newly created library item. This method was added in vSphere API 6.8.
    //
    // @param specParam  information used to create the library item from the source virtual machine.
    // @return Identifier of the newly created library item.
    // The return value will be an identifier for the resource type: ``com.vmware.content.library.Item``.
    // @throws AlreadyExists  if an entity with the name specified by LibraryItemsCreateSpec#name already exists in the folder specified by LibraryItemsCreatePlacementSpec#folder.
    // @throws Error  if the system reports an error while responding to the request.
    // @throws InvalidArgument  if LibraryItemsCreateSpec contains invalid arguments.
    // @throws NotAllowedInCurrentState  if the operation cannot be performed because of the source virtual machine's current state.
    // @throws NotFound  if the source virtual machine specified by LibraryItemsCreateSpec#sourceVm does not exist.
    // @throws NotFound  if the library specified by LibraryItemsCreateSpec#library does not exist.
    // @throws ResourceInaccessible  if there was an error accessing a file from the source virtual machine.
    // @throws ResourceInUse  if the source virtual machine is busy.
    // @throws ServiceUnavailable  if any of the services involved in the method are unavailable.
    // @throws UnableToAllocateResource  if any of the resources needed to create the virtual machine template could not be allocated.
    // @throws Unauthenticated  if the user that requested the method cannot be authenticated.
    // @throws Unauthorized  if the user that requested the method is not authorized to perform the method.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``System.Read``.
    // * The resource ``com.vmware.content.Library`` referenced by the property LibraryItemsCreateSpec#library requires ``ContentLibrary.AddLibraryItem``.
    // * The resource ``VirtualMachine`` referenced by the property LibraryItemsCreateSpec#sourceVm requires ``System.Read``.
    // * The resource ``Datastore`` referenced by the property LibraryItemsCreateSpecVmHomeStorage#datastore requires ``System.Read``.
    // * The resource ``com.vmware.spbm.StorageProfile`` referenced by the property LibraryItemsCreateSpecVmHomeStoragePolicy#policy requires ``System.Read``.
    // * The resource ``Datastore`` referenced by the property LibraryItemsCreateSpecDiskStorage#datastore requires ``System.Read``.
    // * The resource ``com.vmware.spbm.StorageProfile`` referenced by the property LibraryItemsCreateSpecDiskStoragePolicy#policy requires ``System.Read``.
    // * The resource ``com.vmware.vcenter.vm.hardware.Disk`` referenced by the map key of property LibraryItemsCreateSpec#diskStorageOverrides requires ``System.Read``.
    // * The resource ``Folder`` referenced by the property LibraryItemsCreatePlacementSpec#folder requires ``System.Read``.
    // * The resource ``ResourcePool`` referenced by the property LibraryItemsCreatePlacementSpec#resourcePool requires ``System.Read``.
    // * The resource ``HostSystem`` referenced by the property LibraryItemsCreatePlacementSpec#host requires ``System.Read``.
    // * The resource ``ClusterComputeResource`` referenced by the property LibraryItemsCreatePlacementSpec#cluster requires ``System.Read``.
	Create(specParam LibraryItemsCreateSpec) (string, error)

    // Deploys a virtual machine as a copy of the source virtual machine template contained in the library item specified by ``template_library_item``. It uses the deployment specification in ``spec``. If LibraryItemsDeploySpec#poweredOn and/or LibraryItemsDeploySpec#guestCustomization are specified, the server triggers the power on and/or guest customization operations, which are executed asynchronously. This method was added in vSphere API 6.8.
    //
    // @param templateLibraryItemParam  identifier of the content library item containing the source virtual machine template to be deployed.
    // The parameter must be an identifier for the resource type: ``com.vmware.content.library.Item``.
    // @param specParam  specification of how the virtual machine should be deployed.
    // @return Identifier of the deployed virtual machine.
    // The return value will be an identifier for the resource type: ``VirtualMachine``.
    // @throws AlreadyExists  if a virtual machine with the name specified by LibraryItemsDeploySpec#name already exists.
    // @throws Error  if the system reports an error while responding to the request.
    // @throws InvalidArgument  if ``spec`` contains invalid arguments.
    // @throws NotAllowedInCurrentState  if either a specified host or a specified datastore is in an invalid state for the deployment, such as maintenance mode.
    // @throws NotFound  if the library item specified by ``template_library_item`` cannot be found.
    // @throws NotFound  if any resource specified by a property of the LibraryItemsDeploySpec class, specified by ``spec`` cannot be found.
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
    // * The resource ``Datastore`` referenced by the property LibraryItemsDeploySpecVmHomeStorage#datastore requires ``System.Read``.
    // * The resource ``com.vmware.spbm.StorageProfile`` referenced by the property LibraryItemsDeploySpecVmHomeStoragePolicy#policy requires ``System.Read``.
    // * The resource ``Datastore`` referenced by the property LibraryItemsDeploySpecDiskStorage#datastore requires ``System.Read``.
    // * The resource ``com.vmware.spbm.StorageProfile`` referenced by the property LibraryItemsDeploySpecDiskStoragePolicy#policy requires ``System.Read``.
    // * The resource ``com.vmware.vcenter.vm.hardware.Disk`` referenced by the map key of property LibraryItemsDeploySpec#diskStorageOverrides requires ``System.Read``.
    // * The resource ``Folder`` referenced by the property LibraryItemsDeployPlacementSpec#folder requires ``System.Read``.
    // * The resource ``ResourcePool`` referenced by the property LibraryItemsDeployPlacementSpec#resourcePool requires ``System.Read``.
    // * The resource ``HostSystem`` referenced by the property LibraryItemsDeployPlacementSpec#host requires ``System.Read``.
    // * The resource ``ClusterComputeResource`` referenced by the property LibraryItemsDeployPlacementSpec#cluster requires ``System.Read``.
    // * The resource ``com.vmware.vcenter.vm.hardware.Ethernet`` referenced by the map key of property LibraryItemsHardwareCustomizationSpec#nics requires ``System.Read``.
    // * The resource ``Network`` referenced by the property LibraryItemsEthernetUpdateSpec#network requires ``System.Read``.
    // * The resource ``com.vmware.vcenter.vm.hardware.Disk`` referenced by the property LibraryItemsHardwareCustomizationSpec#disksToRemove requires ``System.Read``.
    // * The resource ``com.vmware.vcenter.vm.hardware.Disk`` referenced by the map key of property LibraryItemsHardwareCustomizationSpec#disksToUpdate requires ``System.Read``.
	Deploy(templateLibraryItemParam string, specParam LibraryItemsDeploySpec) (string, error)

    // Returns information about a virtual machine template contained in the library item specified by ``template_library_item``. This method was added in vSphere API 6.8.
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
	Get(templateLibraryItemParam string) (*LibraryItemsInfo, error)
}
