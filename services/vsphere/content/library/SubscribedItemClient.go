/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: SubscribedItem
 * Used by client-side stubs.
 */

package library


// The ``SubscribedItem`` interface manages the unique features of library items that are members of a subscribed library.
type SubscribedItemClient interface {

    // Evicts the cached content of a library item in a subscribed library. 
    //
    //  This method allows the cached content of a library item to be removed to free up storage capacity. This method will only work when a library item is synchronized on-demand. When a library is not synchronized on-demand, it always attempts to keep its cache up-to-date with the published source. Evicting the library item will set ItemModel#cached to false.
    //
    // @param libraryItemIdParam  Identifier of the library item whose content should be evicted.
    // The parameter must be an identifier for the resource type: ``com.vmware.content.library.Item``.
    // @throws NotFound  if the library item specified by ``library_item_id`` does not exist.
    // @throws InvalidElementType  if the library item specified by ``library_item_id`` is not a member of a subscribed library.
    // @throws InvalidElementConfiguration  if the library item specified by ``library_item_id`` is a member of a subscribed library that does not synchronize on-demand.
    // @throws NotAllowedInCurrentState  if the content of the library item specified by ``library_item_id`` has been deleted from the storage backings (see null) associated with it. 
    //
    //  For instance, this {\\\\@term error) is reported on evicting a library item in an on-demand subscribed library that was restored from backup, and the library item was deleted after backup, thus resulting in its content being deleted from the associated storage backings. In this scenario, the metadata of the library item is present on a restore, while its content has been deleted.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * The resource ``com.vmware.content.library.Item`` referenced by the parameter ``library_item_id`` requires ``ContentLibrary.EvictLibraryItem``.
	Evict(libraryItemIdParam string) error

    // Forces the synchronization of an individual library item in a subscribed library. 
    //
    //  Synchronizing an individual item will update that item's metadata from the remote source. If the source library item on the remote library has been deleted, this method will delete the library item from the subscribed library as well. 
    //
    //  The default behavior of the synchronization is determined by the SubscriptionInfo of the library which owns the library item. 
    //
    // * If SubscriptionInfo#onDemand is true, then the file content is not synchronized by default. In this case, only the library item metadata is synchronized. The file content may still be forcefully synchronized by passing true for the ``force_sync_content`` parameter.
    // * If SubscriptionInfo#onDemand is false, then this call will always synchronize the file content. The ``force_sync_content`` parameter is ignored when the subscription is not on-demand.
    //
    //  When the file content has been synchronized, the ItemModel#cached property will be true. 
    //
    //  This method will return immediately and create an asynchronous task to perform the synchronization.
    //
    // @param libraryItemIdParam  Identifier of the library item to synchronize.
    // The parameter must be an identifier for the resource type: ``com.vmware.content.library.Item``.
    // @param forceSyncContentParam  Whether to synchronize file content as well as metadata. This parameter applies only if the subscription is on-demand.
    // @throws NotFound  if the library item specified by ``library_item_id`` could not be found.
    // @throws InvalidElementType  if the library item specified by ``library_item_id`` is not a member of a subscribed library.
    // @throws NotAllowedInCurrentState  if the content of the library item specified by ``library_item_id`` has been deleted from the storage backings (see null) associated with it. 
    //
    //  For instance, this {\\\\@term error) is reported on synchronizing a library item in a subscribed library that was restored from backup, and the library item was deleted after backup, thus resulting in its content being deleted from the associated storage backings. In this scenario, the metadata of the library item is present on a restore, while its content has been deleted.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * The resource ``com.vmware.content.library.Item`` referenced by the parameter ``library_item_id`` requires ``ContentLibrary.SyncLibraryItem``.
	Sync(libraryItemIdParam string, forceSyncContentParam bool) error
}
