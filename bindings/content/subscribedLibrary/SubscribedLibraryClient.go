/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: SubscribedLibrary
 * Used by client-side stubs.
 */

package subscribedLibrary

import (
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/bindings/content"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/bindings/content/library"
)

//
type SubscribedLibraryClient interface {


    // Creates a new subscribed library. 
    //
    //  Once created, the subscribed library will be empty. If the content.LibraryModel#subscriptionInfo property is set, the Content Library Service will attempt to synchronize to the remote source. This is an asynchronous operation so the content of the published library may not immediately appear.
    //
    // @param clientTokenParam  Unique token generated on the client for each creation request. The token should be a universally unique identifier (UUID), for example: ``b8a2a2e3-2314-43cd-a871-6ede0f429751``. This token can be used to guarantee idempotent creation.
    // If not specified creation is not idempotent.
    // @param createSpecParam  Specification for the new subscribed library.
    // @return Identifier of the newly created subscribed library.
    // The return value will be an identifier for the resource type: ``com.vmware.content.Library``.
    // @throws InvalidArgument  if the ``create_spec`` is not valid.
    // @throws InvalidArgument  if the ``client_token`` does not conform to the UUID format.
    // @throws Unsupported  if using multiple storage backings.
    // @throws ResourceInaccessible  if subscribing to a published library which cannot be accessed.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``ContentLibrary.CreateSubscribedLibrary``.
    // * The resource ``Datastore`` referenced by the property library.StorageBacking#datastoreId requires ``Datastore.AllocateSpace``.
    Create(clientTokenParam *string, createSpecParam content.LibraryModel) (string, error) 


    // Deletes the specified subscribed library. 
    //
    //  Deleting a subscribed library will remove the entry immediately and begin an asynchronous task to remove all cached content for the library. If the asynchronous task fails, file content may remain on the storage backing. This content will require manual removal.
    //
    // @param libraryIdParam  Identifier of the subscribed library to delete.
    // The parameter must be an identifier for the resource type: ``com.vmware.content.Library``.
    // @throws InvalidElementType  if the library referenced by ``library_id`` is not a subscribed library.
    // @throws NotFound  if the library referenced by ``library_id`` does not exist.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * The resource ``com.vmware.content.Library`` referenced by the parameter ``library_id`` requires ``ContentLibrary.DeleteSubscribedLibrary``.
    Delete(libraryIdParam string) error 


    // Evicts the cached content of an on-demand subscribed library. 
    //
    //  This method allows the cached content of a subscribed library to be removed to free up storage capacity. This method will only work when a subscribed library is synchronized on-demand.
    //
    // @param libraryIdParam  Identifier of the subscribed library whose content should be evicted.
    // The parameter must be an identifier for the resource type: ``com.vmware.content.Library``.
    // @throws NotFound  if the library specified by ``library_id`` does not exist.
    // @throws InvalidElementType  if the library specified by ``library_id`` is not a subscribed library.
    // @throws NotAllowedInCurrentState  if the library specified by ``library_id`` does not synchronize on-demand, or if the content of the library specified by ``library_id`` has been deleted from the storage backings (see content.LibraryModel#storageBackings) associated with it. 
    //
    //  For instance, this {\\\\@term error) is reported on evicting an on-demand subscribed library that was restored from backup, and the library was deleted after the backup was taken, thus resulting in its content being deleted from the associated storage backings. In this scenario, the metadata of the library is present on a restore, while its content has been deleted.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * The resource ``com.vmware.content.Library`` referenced by the parameter ``library_id`` requires ``ContentLibrary.EvictSubscribedLibrary``.
    Evict(libraryIdParam string) error 


    // Returns a given subscribed library.
    //
    // @param libraryIdParam  Identifier of the subscribed library to return.
    // The parameter must be an identifier for the resource type: ``com.vmware.content.Library``.
    // @return The content.LibraryModel instance that corresponds to ``library_id``.
    // @throws NotFound  if the library associated with ``library_id`` does not exist.
    // @throws InvalidElementType  if the library associated with ``library_id`` is not a subscribed library.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * The resource ``com.vmware.content.Library`` referenced by the parameter ``library_id`` requires ``System.Read``.
    Get(libraryIdParam string) (content.LibraryModel, error) 


    // Returns the identifiers of all subscribed libraries in the Content Library.
    // @return The array of identifiers of all subscribed libraries in the Content Library.
    // The return value will contain identifiers for the resource type: ``com.vmware.content.Library``.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``System.Read``.
    List() ([]string, error) 


    // Forces the synchronization of the subscribed library. 
    //
    //  Synchronizing a subscribed library forcefully with this method will perform the same synchronization behavior as would run periodically for the library. The library.SubscriptionInfo#onDemand setting is respected. Calling this method on a library that is already in the process of synchronizing will have no effect.
    //
    // @param libraryIdParam  Identifier of the subscribed library to synchronize.
    // The parameter must be an identifier for the resource type: ``com.vmware.content.Library``.
    // @throws NotFound  if the library specified by ``library_id`` does not exist.
    // @throws InvalidElementType  if the library specified by ``library_id`` is not a subscribed library.
    // @throws NotAllowedInCurrentState  if the content of the library specified by ``library_id`` has been deleted from the storage backings (see content.LibraryModel#storageBackings) associated with it. 
    //
    //  For instance, this {\\\\@term error) is reported on synchronizing a subscribed library that was restored from backup, and the library was deleted after the backup was taken, thus resulting in its content being deleted from the associated storage backings. In this scenario, the metadata of the library is present on a restore, while its content has been deleted.
    // @throws InvalidArgument  if some parameter in the subscribed library subscription info is invalid.
    // @throws ResourceInaccessible  if the published library cannot be contacted or found.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * The resource ``com.vmware.content.Library`` referenced by the parameter ``library_id`` requires ``ContentLibrary.SyncLibrary``.
    Sync(libraryIdParam string) error 


    // Updates the properties of a subscribed library. 
    //
    //  This is an incremental update to the subscribed library. Properties that are null in the update specification will be left unchanged.
    //
    // @param libraryIdParam  Identifier of the subscribed library to update.
    // The parameter must be an identifier for the resource type: ``com.vmware.content.Library``.
    // @param updateSpecParam  Specification of the new property values to set on the subscribed library.
    // @throws NotFound  if the library specified by ``library_id`` does not exist.
    // @throws NotAllowedInCurrentState  if the ``update_spec`` updates the subscription URL (see library.SubscriptionInfo#subscriptionUrl) and the content of the library specified by ``library_id`` has been deleted from the storage backings (see content.LibraryModel#storageBackings) associated with it.
    // @throws InvalidElementType  if the library specified by ``library_id`` is not a subscribed library.
    // @throws InvalidArgument  if the ``update_spec`` is not valid.
    // @throws ResourceInaccessible  if the subscription info is being updated but the published library cannot be contacted or found.
    // @throws ResourceBusy  if the content.LibraryModel#version of ``update_spec`` is null and the library is being concurrently updated by another user.
    // @throws ConcurrentChange  if the content.LibraryModel#version of ``update_spec`` is not equal to the current version of the library.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * The resource ``com.vmware.content.Library`` referenced by the parameter ``library_id`` requires ``ContentLibrary.UpdateSubscribedLibrary``.
    Update(libraryIdParam string, updateSpecParam content.LibraryModel) error 


    // Probes remote library subscription information, including URL, SSL certificate and password. The resulting ProbeResult class describes whether or not the subscription configuration is successful.
    //
    // @param subscriptionInfoParam  The subscription info to be probed.
    // @return The subscription info probe result.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``ContentLibrary.ProbeSubscription``.
    Probe(subscriptionInfoParam library.SubscriptionInfo) (ProbeResult, error) 

}
