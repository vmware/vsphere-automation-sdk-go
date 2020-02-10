/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Versions
 * Used by client-side stubs.
 */

package library_items


// The ``Versions`` interface provides methods for managing the live versions of the virtual machine templates contained in a library item. Live versions include the latest and previous virtual machine templates that are available on disk. As new versions of virtual machine templates are checked in, old versions of virtual machine templates are automatically purged. Currently, at most one previous virtual machine template version is stored. This interface was added in vSphere API 6.9.1.
type VersionsClient interface {

    // Returns commonly used information about the live versions of a virtual machine template library item. This method was added in vSphere API 6.9.1.
    //
    // @param templateLibraryItemParam Identifier of the VM template library item.
    // The parameter must be an identifier for the resource type: ``com.vmware.content.library.Item``.
    // @return List of commonly used information about the library item versions.
    // @throws NotFound  if the library item is not found.
    // @throws InvalidArgument  if the library item does not contain a virtual machine template.
    // @throws Unauthenticated  if the user that requested the method cannot be authenticated.
    // @throws Unauthorized  if the user that requested the method is not authorized to perform the method.
    // @throws Error  if the system reports an error while responding to the request.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``System.Read``.
    // * The resource ``com.vmware.content.library.Item`` referenced by the parameter ``template_library_item`` requires ``System.Read``.
	List(templateLibraryItemParam string) ([]VersionsSummary, error)

    // Returns information about the live version of a library item containing a virtual machine template. This method was added in vSphere API 6.9.1.
    //
    // @param templateLibraryItemParam Identifier of the VM template library item.
    // The parameter must be an identifier for the resource type: ``com.vmware.content.library.Item``.
    // @param versionParam Version of the library item.
    // The parameter must be an identifier for the resource type: ``com.vmware.content.library.item.Version``.
    // @return Information about the specified library item version.
    // @throws NotFound  if the library item or version is not found.
    // @throws InvalidArgument  if the library item does not contain a virtual machine template.
    // @throws Unauthenticated  if the user that requested the method cannot be authenticated.
    // @throws Unauthorized  if the user that requested the method is not authorized to perform the method.
    // @throws Error  if the system reports an error while responding to the request.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``System.Read``.
    // * The resource ``com.vmware.content.library.Item`` referenced by the parameter ``template_library_item`` requires ``System.Read``.
    // * The resource ``com.vmware.content.library.item.Version`` referenced by the parameter ``version`` requires ``System.Read``.
	Get(templateLibraryItemParam string, versionParam string) (VersionsInfo, error)

    // Rollbacks a library item containing a virtual machine template to a previous version. The virtual machine template at the specified version becomes the latest virtual machine template with a new version identifier. This method was added in vSphere API 6.9.1.
    //
    // @param templateLibraryItemParam Identifier of the VM template library item.
    // The parameter must be an identifier for the resource type: ``com.vmware.content.library.Item``.
    // @param versionParam Version of the library item to rollback.
    // The parameter must be an identifier for the resource type: ``com.vmware.content.library.item.Version``.
    // @param specParam Specification to rollback the library item.
    // This parameter is currently required. In the future, if this parameter is null, the system will apply suitable defaults.
    // @return The new version of the library item.
    // The return value will be an identifier for the resource type: ``com.vmware.content.library.item.Version``.
    // @throws NotFound  if the library item or version is not found.
    // @throws InvalidArgument  if the specified version is the latest version of the library item.
    // @throws InvalidArgument  if the library item does not contain a virtual machine template.
    // @throws NotAllowedInCurrentState  if a virtual machine is checked out of the library item.
    // @throws Unauthenticated  if the user that requested the method cannot be authenticated.
    // @throws Unauthorized  if the user that requested the method is not authorized to perform the method.
    // @throws Error  if the system reports an error while responding to the request.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``System.Read``.
    // * The resource ``com.vmware.content.library.Item`` referenced by the parameter ``template_library_item`` requires ``ContentLibrary.CheckInTemplate``.
    // * The resource ``com.vmware.content.library.item.Version`` referenced by the parameter ``version`` requires ``System.Read``.
	Rollback(templateLibraryItemParam string, versionParam string, specParam *VersionsRollbackSpec) (string, error)

    // Deletes the virtual machine template contained in the library item at the specified version. This method was added in vSphere API 6.9.1.
    //
    // @param templateLibraryItemParam Identifier of the VM template library item.
    // The parameter must be an identifier for the resource type: ``com.vmware.content.library.Item``.
    // @param versionParam Version of the library item to delete.
    // The parameter must be an identifier for the resource type: ``com.vmware.content.library.item.Version``.
    // @throws NotFound  if the library item or version is not found.
    // @throws InvalidArgument  if the specified version is the latest version of the library item.
    // @throws InvalidArgument  if the library item does not contain a virtual machine template.
    // @throws ResourceInaccessible  if the virtual machine template's configuration state cannot be accessed.
    // @throws Unauthenticated  if the user that requested the method cannot be authenticated.
    // @throws Unauthorized  if the user that requested the method is not authorized to perform the method.
    // @throws Error  if the system reports an error while responding to the request.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``System.Read``.
    // * The resource ``com.vmware.content.library.Item`` referenced by the parameter ``template_library_item`` requires ``ContentLibrary.DeleteLibraryItem``.
    // * The resource ``com.vmware.content.library.item.Version`` referenced by the parameter ``version`` requires ``System.Read``.
	Delete(templateLibraryItemParam string, versionParam string) error
}
