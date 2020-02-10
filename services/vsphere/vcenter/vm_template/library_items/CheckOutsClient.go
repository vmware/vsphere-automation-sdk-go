/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: CheckOuts
 * Used by client-side stubs.
 */

package library_items


// The ``CheckOuts`` interface provides methods for managing the checkouts of a library item containing a virtual machine template. This interface provides operations to check out a library item to update the virtual machine template, and to check in the library item when the virtual machine changes are complete. This interface was added in vSphere API 6.9.1.
type CheckOutsClient interface {

    // Checks out a library item containing a virtual machine template. This method makes a copy of the source virtual machine template contained in the library item as a virtual machine. The virtual machine is copied with the same storage specification as the source virtual machine template. Changes to the checked out virtual machine do not affect the virtual machine template contained in the library item. To save these changes back into the library item, CheckOuts#checkIn the virtual machine. To discard the changes, CheckOuts#delete the virtual machine. This method was added in vSphere API 6.9.1.
    //
    // @param templateLibraryItemParam  Identifier of the content library item containing the source virtual machine template to be checked out.
    // The parameter must be an identifier for the resource type: ``com.vmware.content.library.Item``.
    // @param specParam Specification used to check out the source virtual machine template as a virtual machine.
    // This parameter is currently required. In the future, if this parameter is null, the system will apply suitable defaults.
    // @return Identifier of the virtual machine that was checked out of the library item.
    // The return value will be an identifier for the resource type: ``VirtualMachine``.
    // @throws AlreadyExists  if a virtual machine with the name specified by CheckOutsCheckOutSpec#name already exists in the folder specified by CheckOutsPlacementSpec#folder.
    // @throws InvalidArgument  if ``spec`` contains invalid arguments.
    // @throws InvalidArgument  if the library item is a member of a subscribed library.
    // @throws NotFound  if the library item specified by ``template_library_item`` cannot be found.
    // @throws ResourceInaccessible  if there is an error accessing the files of the source virtual machine template contained in the library item specified by ``template_library_item``.
    // @throws UnableToAllocateResource  if the limit for the number of virtual machines checked out of a library item (currently 1) has been exceeded.
    // @throws Unauthenticated  if the user that requested the method cannot be authenticated.
    // @throws Unauthorized  if the user that requested the method is not authorized to perform the method.
    // @throws Error  if the system reports an error while responding to the request.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``System.Read``.
    // * The resource ``com.vmware.content.library.Item`` referenced by the parameter ``template_library_item`` requires ``ContentLibrary.CheckOutTemplate``.
    // * The resource ``Folder`` referenced by the property CheckOutsPlacementSpec#folder requires ``System.Read``.
    // * The resource ``ResourcePool`` referenced by the property CheckOutsPlacementSpec#resourcePool requires ``System.Read``.
    // * The resource ``HostSystem`` referenced by the property CheckOutsPlacementSpec#host requires ``System.Read``.
    // * The resource ``ClusterComputeResource`` referenced by the property CheckOutsPlacementSpec#cluster requires ``System.Read``.
	CheckOut(templateLibraryItemParam string, specParam *CheckOutsCheckOutSpec) (string, error)

    // Checks in a virtual machine into the library item. This method updates the library item to contain the virtual machine being checked in as a template. This template becomes the latest version of the library item. The previous virtual machine template contained in the library item is available as a backup and its information can be queried using the ``Versions`` interface. At most one previous version of a virtual machine template is retained in the library item. This method was added in vSphere API 6.9.1.
    //
    // @param templateLibraryItemParam  Identifier of the content library item in which the virtual machine is checked in.
    // The parameter must be an identifier for the resource type: ``com.vmware.content.library.Item``.
    // @param vmParam  Identifier of the virtual machine to check into the library item.
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @param specParam  Specification used to check in the virtual machine into the library item.
    // This parameter is currently required. In the future, if this parameter is null, the system will apply suitable defaults.
    // @return The new version of the library item.
    // The return value will be an identifier for the resource type: ``com.vmware.content.library.item.Version``.
    // @throws InvalidArgument  if any of the specified parameters are invalid.
    // @throws InvalidArgument  if the virtual machine identified by ``vm`` was not checked out of the item specified by ``template_library_item``.
    // @throws NotAllowedInCurrentState  if the method cannot be performed because of the virtual machine's current state. For example, if the virtual machine is not powered off.
    // @throws NotFound  if the item specified by ``template_library_item`` does not exist.
    // @throws NotFound  if the virtual machine specified by ``vm`` does not exist.
    // @throws ResourceInaccessible  if there is an error accessing a file from the virtual machine.
    // @throws Unauthenticated  if the user that requested the method cannot be authenticated.
    // @throws Unauthorized  if the user that requested the method is not authorized to perform the method.
    // @throws Error  if the system reports an error while responding to the request.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``System.Read``.
    // * The resource ``com.vmware.content.library.Item`` referenced by the parameter ``template_library_item`` requires ``ContentLibrary.CheckInTemplate``.
    // * The resource ``VirtualMachine`` referenced by the parameter ``vm`` requires ``System.Read``.
	CheckIn(templateLibraryItemParam string, vmParam string, specParam *CheckOutsCheckInSpec) (string, error)

    // Returns commonly used information about the virtual machines that are checked out of the library item. This method was added in vSphere API 6.9.1.
    //
    // @param templateLibraryItemParam Identifier of the VM template library item.
    // The parameter must be an identifier for the resource type: ``com.vmware.content.library.Item``.
    // @return List of commonly used information about the check outs.
    // @throws NotFound  if the library item is not found.
    // @throws InvalidArgument  if the library item does not contain a virtual machine template.
    // @throws Unauthenticated  if the user that requested the method cannot be authenticated.
    // @throws Unauthorized  if the user that requested the method is not authorized to perform the method.
    // @throws Error  if the system reports an error while responding to the request.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``System.Read``.
    // * The resource ``com.vmware.content.library.Item`` referenced by the parameter ``template_library_item`` requires ``System.Read``.
	List(templateLibraryItemParam string) ([]CheckOutsSummary, error)

    // Returns the information about a checked out virtual machine. This method was added in vSphere API 6.9.1.
    //
    // @param templateLibraryItemParam Identifier of the VM template library item.
    // The parameter must be an identifier for the resource type: ``com.vmware.content.library.Item``.
    // @param vmParam Identifier of the checked out virtual machine.
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @return Information about a check out.
    // @throws NotFound  if the library item or virtual machine is not found.
    // @throws InvalidArgument  if the virtual machine is not checked out of the library item.
    // @throws InvalidArgument  if the library item does not contain a virtual machine template.
    // @throws Unauthenticated  if the user that requested the method cannot be authenticated.
    // @throws Unauthorized  if the user that requested the method is not authorized to perform the method.
    // @throws Error  if the system reports an error while responding to the request.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``System.Read``.
    // * The resource ``com.vmware.content.library.Item`` referenced by the parameter ``template_library_item`` requires ``System.Read``.
    // * The resource ``VirtualMachine`` referenced by the parameter ``vm`` requires ``System.Read``.
	Get(templateLibraryItemParam string, vmParam string) (CheckOutsInfo, error)

    // Deletes the checked out virtual machine. This method was added in vSphere API 6.9.1.
    //
    // @param templateLibraryItemParam Identifier of the VM template library item.
    // The parameter must be an identifier for the resource type: ``com.vmware.content.library.Item``.
    // @param vmParam Identifier of the checked out virtual machine to delete.
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @throws NotFound  if the library item or virtual machine is not found.
    // @throws InvalidArgument  if the virtual machine is not checked out of the library item.
    // @throws NotAllowedInCurrentState  if the virtual machine is running (powered on).
    // @throws ResourceBusy  if the virtual machine is busy performing another operation.
    // @throws ResourceInaccessible  if the virtual machine's configuration state cannot be accessed.
    // @throws Unauthenticated  if the user that requested the method cannot be authenticated.
    // @throws Unauthorized  if the user that requested the method is not authorized to perform the method.
    // @throws Error  if the system reports an error while responding to the request.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``System.Read``.
    // * The resource ``VirtualMachine`` referenced by the parameter ``vm`` requires ``VirtualMachine.Inventory.Delete``.
    // * The resource ``com.vmware.content.library.Item`` referenced by the parameter ``template_library_item`` requires ``System.Read``.
	Delete(templateLibraryItemParam string, vmParam string) error
}
