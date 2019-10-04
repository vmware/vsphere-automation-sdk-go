/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: LocalLibrary
 * Used by client-side stubs.
 */

package localLibrary

import (
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/apis/content"
)

// The ``LocalLibrary`` interface manages local libraries. 
//
//  The ``LocalLibrary`` interface provides support for creating and maintaining local library instances. A local library may also use the Library interface to manage general library functionality.
type LocalLibraryClient interface {


    // Creates a new local library.
    //
    // @param clientTokenParam  A unique token generated on the client for each creation request. The token should be a universally unique identifier (UUID), for example: ``b8a2a2e3-2314-43cd-a871-6ede0f429751``. This token can be used to guarantee idempotent creation.
    // If not specified creation is not idempotent.
    // @param createSpecParam  Specification for the new local library.
    // @return Identifier of the newly created content.LibraryModel.
    // The return value will be an identifier for the resource type: ``com.vmware.content.Library``.
    // @throws InvalidArgument  if the ``create_spec`` is not valid.
    // @throws InvalidArgument  if the ``client_token`` does not conform to the UUID format.
    // @throws Unsupported  if using multiple storage backings.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``ContentLibrary.CreateLocalLibrary``.
    // * The resource ``Datastore`` referenced by the property library.StorageBacking#datastoreId requires ``Datastore.AllocateSpace``.
    Create(clientTokenParam *string, createSpecParam content.LibraryModel) (string, error) 


    // Deletes the specified local library. 
    //
    //  Deleting a local library will remove the entry immediately and begin an asynchronous task to remove all cached content for the library. If the asynchronous task fails, file content may remain on the storage backing. This content will require manual removal.
    //
    // @param libraryIdParam  Identifier of the local library to delete.
    // The parameter must be an identifier for the resource type: ``com.vmware.content.Library``.
    // @throws InvalidElementType  if the library specified by ``library_id`` is not a local library.
    // @throws NotFound  if the library specified by ``library_id`` does not exist.
    // @throws NotAllowedInCurrentState  if the library contains a library item that cannot be deleted in its current state. For example, the library item contains a virtual machine template and a virtual machine is checked out of the library item.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * The resource ``com.vmware.content.Library`` referenced by the parameter ``library_id`` requires ``ContentLibrary.DeleteLocalLibrary``.
    Delete(libraryIdParam string) error 


    // Returns a given local library.
    //
    // @param libraryIdParam  Identifier of the local library to return.
    // The parameter must be an identifier for the resource type: ``com.vmware.content.Library``.
    // @return The content.LibraryModel instance associated with ``library_id``.
    // @throws NotFound  if the library specified by ``library_id`` does not exist.
    // @throws InvalidElementType  if the library specified by ``library_id`` is not a local library.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * The resource ``com.vmware.content.Library`` referenced by the parameter ``library_id`` requires ``System.Read``.
    Get(libraryIdParam string) (content.LibraryModel, error) 


    // Returns the identifiers of all local libraries in the Content Library.
    // @return The array of identifiers of all local libraries in the Content Library.
    // The return value will contain identifiers for the resource type: ``com.vmware.content.Library``.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``System.Read``.
    List() ([]string, error) 


    // Updates the properties of a local library. 
    //
    //  This is an incremental update to the local library. Properties that are null in the update specification will be left unchanged.
    //
    // @param libraryIdParam  Identifier of the local library to update.
    // The parameter must be an identifier for the resource type: ``com.vmware.content.Library``.
    // @param updateSpecParam  Specification of the new property values to set on the local library.
    // @throws NotFound  if the library specified by ``library_id`` does not exist.
    // @throws NotAllowedInCurrentState  if the library specified by ``library_id`` is a published library with JSON persistence enabled (see library.PublishInfo#persistJsonEnabled) and the content of the library has been deleted from the storage backings (see content.LibraryModel#storageBackings) associated with it.
    // @throws InvalidElementType  if the library specified by ``library_id`` is not a local library.
    // @throws InvalidArgument  if the ``update_spec`` is not valid.
    // @throws InvalidArgument  if the library.PublishInfo#currentPassword in the ``update_spec`` does not match the existing password of the published library.
    // @throws ResourceBusy  if the content.LibraryModel#version of ``update_spec`` is null and the library is being concurrently updated by another user.
    // @throws ConcurrentChange  if the content.LibraryModel#version of ``update_spec`` is not equal to the current version of the library.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * The resource ``com.vmware.content.Library`` referenced by the parameter ``library_id`` requires ``ContentLibrary.UpdateLocalLibrary``.
    Update(libraryIdParam string, updateSpecParam content.LibraryModel) error 


    // Publishes the library to specified subscriptions. If no subscriptions are specified, then publishes the library to all its subscriptions.
    //
    // @param libraryIdParam Identifier of the published library.
    // The parameter must be an identifier for the resource type: ``com.vmware.content.Library``.
    // @param subscriptionsParam The list of subscriptions to publish this library to.
    // @throws Error  If the system reports an error while responding to the request.
    // @throws NotFound  If the library specified by ``library_id`` does not exist.
    // @throws InvalidArgument  If one or more ``subscriptions`` is not valid.
    // @throws InvalidElementType  If the library specified by ``library_id`` is a subscribed library.
    // @throws NotAllowedInCurrentState  If the library specified by ``library_id`` is not a published library.
    // @throws Unauthenticated  If the user that requested the method cannot be authenticated.
    // @throws Unauthorized  If the user that requested the method is not authorized to perform the method.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * The resource ``com.vmware.content.Library`` referenced by the parameter ``library_id`` requires ``ContentLibrary.PublishLibrary``.
    Publish(libraryIdParam string, subscriptionsParam []DestinationSpec) error 

}
