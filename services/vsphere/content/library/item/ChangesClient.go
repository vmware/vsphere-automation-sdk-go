/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Changes
 * Used by client-side stubs.
 */

package item


// The ``Changes`` interface provides methods to get a history of the content changes made to a library item. This interface was added in vSphere API 6.9.1.
type ChangesClient interface {

    // Returns commonly used information about the content changes made to a library item. This method was added in vSphere API 6.9.1.
    //
    // @param libraryItemParam Library item identifier.
    // The parameter must be an identifier for the resource type: ``com.vmware.content.library.Item``.
    // @return List of commonly used information about the library item changes.
    // @throws NotFound  if the library item is not found.
    // @throws Unauthenticated  if the user that requested the method cannot be authenticated.
    // @throws Unauthorized  if the user that requested the method is not authorized to perform the method.
    // @throws Error  if the system reports an error while responding to the request.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * The resource ``com.vmware.content.library.Item`` referenced by the parameter ``library_item`` requires ``System.Read``.
	List(libraryItemParam string) ([]ChangesSummary, error)

    // Returns information about a library item change. This method was added in vSphere API 6.9.1.
    //
    // @param libraryItemParam Library item identifer.
    // The parameter must be an identifier for the resource type: ``com.vmware.content.library.Item``.
    // @param versionParam Library item version.
    // The parameter must be an identifier for the resource type: ``com.vmware.content.library.item.Version``.
    // @return Information about the specified library item change.
    // @throws NotFound  if the library item or version is not found.
    // @throws Unauthenticated  if the user that requested the method cannot be authenticated.
    // @throws Unauthorized  if the user that requested the method is not authorized to perform the method.
    // @throws Error  if the system reports an error while responding to the request.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * The resource ``com.vmware.content.library.Item`` referenced by the parameter ``library_item`` requires ``System.Read``.
	Get(libraryItemParam string, versionParam string) (ChangesInfo, error)
}
