/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Library
 * Used by client-side stubs.
 */

package content


// The ``Library`` interface provides methods to manage and find LibraryModel entities. 
//
//  The ``Library`` interface provides support for generic functionality which can be applied equally to all types of libraries. The functionality provided by this interface will not affect the properties specific to the type of library. See also LocalLibrary and SubscribedLibrary.
type LibraryClient interface {

    // Returns a given LibraryModel.
    //
    // @param libraryIdParam  Identifier of the library to return.
    // The parameter must be an identifier for the resource type: ``com.vmware.content.Library``.
    // @return The LibraryModel instance with the specified ``library_id``.
    // @throws NotFound  if the specified library does not exist.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * The resource ``com.vmware.content.Library`` referenced by the parameter ``library_id`` requires ``System.Read``.
	Get(libraryIdParam string) (LibraryModel, error)

    // Returns the identifiers of all libraries of any type in the Content Library.
    // @return The array of all identifiers of all libraries in the Content Library.
    // The return value will contain identifiers for the resource type: ``com.vmware.content.Library``.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``System.Read``.
	List() ([]string, error)

    // Returns a list of all the visible (as determined by authorization policy) libraries matching the requested LibraryFindSpec.
    //
    // @param specParam  Specification describing what properties to filter on.
    // @return The array of identifiers of all the visible libraries matching the given ``spec``.
    // The return value will contain identifiers for the resource type: ``com.vmware.content.Library``.
    // @throws InvalidArgument  if no properties are specified in the ``spec``.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``System.Read``.
	Find(specParam LibraryFindSpec) ([]string, error)

    // Updates the properties of a library. 
    //
    //  This is an incremental update to the library. Any property in the LibraryModel class that is null will not be modified. 
    //
    //  This method will only update the common properties for all library types. This will not, for example, update the LibraryModel#publishInfo of a local library, nor the LibraryModel#subscriptionInfo of a subscribed library. Specific properties are updated in LocalLibrary#update and SubscribedLibrary#update.
    //
    // @param libraryIdParam  Identifier of the library to update.
    // The parameter must be an identifier for the resource type: ``com.vmware.content.Library``.
    // @param updateSpecParam  Specification of the new property values to set on the library.
    // @throws NotFound  if the library associated with ``library_id`` does not exist.
    // @throws InvalidArgument  if the ``update_spec`` is not valid.
    // @throws InvalidArgument  if the LibraryModel#version of ``update_spec`` is not equal to the current version of the library.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * The resource ``com.vmware.content.Library`` referenced by the parameter ``library_id`` requires ``ContentLibrary.UpdateLibrary``.
	Update(libraryIdParam string, updateSpecParam LibraryModel) error
}
