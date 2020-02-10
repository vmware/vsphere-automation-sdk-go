/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Item
 * Used by client-side stubs.
 */

package library


// The ``Item`` interface provides methods for managing library items.
type ItemClient interface {

    // Copies a library item. 
    //
    //  Copying a library item allows a duplicate to be made within the same or different library. The copy occurs by first creating a new library item, whose identifier is returned. The content of the library item is then copied asynchronously. This copy can be tracked as a task. 
    //
    //  If the copy fails, Content Library Service will roll back the copy by deleting any content that was already copied, and removing the new library item. A failure during rollback may require manual cleanup by an administrator. 
    //
    //  A library item cannot be copied into a subscribed library.
    //
    // @param clientTokenParam  A unique token generated on the client for each copy request. The token should be a universally unique identifier (UUID), for example: ``b8a2a2e3-2314-43cd-a871-6ede0f429751``. This token can be used to guarantee idempotent copy.
    // If not specified copy is not idempotent.
    // @param sourceLibraryItemIdParam  Identifier of the existing library item from which the content will be copied.
    // The parameter must be an identifier for the resource type: ``com.vmware.content.library.Item``.
    // @param destinationCreateSpecParam  Specification for the new library item to be created.
    // @return The identifier of the new library item into which the content is being copied.
    // The return value will be an identifier for the resource type: ``com.vmware.content.library.Item``.
    // @throws NotFound  if the library item with ``source_library_item_id`` does not exist, or if the library referenced by the ItemModel#libraryId property of ``destination_create_spec`` does not exist.
    // @throws InvalidArgument  if one of the following is true for the new library item: 
    //
    // * name is empty
    // * name exceeds 80 characters
    // * description exceeds 2000 characters
    // @throws InvalidArgument  if the ``client_token`` does not conform to the UUID format.
    // @throws InvalidElementType  if the ItemModel#libraryId property of ``destination_create_spec`` refers to a subscribed library.
    // @throws ResourceInaccessible  if the copy operation failed because the source or destination library item is not accessible.
    // @throws NotAllowedInCurrentState  if the content of the source library item specified by ``source_library_item_id``, or the content of the target library specified by the library ID (see ItemModel#libraryId) property of ``destination_create_spec`` has been deleted from the storage backings (see null) associated with it.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * The resource ``com.vmware.content.library.Item`` referenced by the parameter ``source_library_item_id`` requires ``System.Read``.
    // * The resource ``com.vmware.content.Library`` referenced by the property ItemModel#libraryId requires ``ContentLibrary.AddLibraryItem``.
	Copy(clientTokenParam *string, sourceLibraryItemIdParam string, destinationCreateSpecParam ItemModel) (string, error)

    // Creates a new library item. 
    //
    //  A new library item is created without any content. After creation, content can be added through the UpdateSession and File interfaces. 
    //
    //  A library item cannot be created in a subscribed library.
    //
    // @param clientTokenParam  A unique token generated on the client for each creation request. The token should be a universally unique identifier (UUID), for example: ``b8a2a2e3-2314-43cd-a871-6ede0f429751``. This token can be used to guarantee idempotent creation.
    // If not specified creation is not idempotent.
    // @param createSpecParam  Specification that defines the properties of the new library item.
    // @return Identifier of the new library item.
    // The return value will be an identifier for the resource type: ``com.vmware.content.library.Item``.
    // @throws NotFound  if the ItemModel#libraryId property of ``create_spec`` refers to a library that does not exist.
    // @throws InvalidArgument  if one of the following is true for the new library item: 
    //
    // * name is empty
    // * name exceeds 80 characters
    // * description exceeds 2000 characters
    // @throws InvalidArgument  if the ``client_token`` does not conform to the UUID format.
    // @throws InvalidElementType  if the ItemModel#libraryId property of ``create_spec`` refers to a subscribed library.
    // @throws NotAllowedInCurrentState  if the content of the library specified by the library ID (see ItemModel#libraryId) property of ``create_spec`` has been deleted from the storage backings (see null) associated with it.
    // @throws AlreadyExists  if there is already a library item with same name in the library.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * The resource ``com.vmware.content.Library`` referenced by the property ItemModel#libraryId requires ``ContentLibrary.AddLibraryItem``.
	Create(clientTokenParam *string, createSpecParam ItemModel) (string, error)

    // Deletes a library item. 
    //
    //  This method will immediately remove the item from the library that owns it. The content of the item will be asynchronously removed from the storage backings. The content deletion can be tracked with a task. In the event that the task fails, an administrator may need to manually remove the files from the storage backing. 
    //
    //  This method cannot be used to delete a library item that is a member of a subscribed library. Removing an item from a subscribed library requires deleting the item from the original published local library and syncing the subscribed library.
    //
    // @param libraryItemIdParam  Identifier of the library item to delete.
    // The parameter must be an identifier for the resource type: ``com.vmware.content.library.Item``.
    // @throws InvalidElementType  if the library item with the given ``library_item_id`` is a member of a subscribed library.
    // @throws NotFound  if the library item with the specified ``library_item_id`` does not exist.
    // @throws NotAllowedInCurrentState  if the library item contains a virtual machine template and a virtual machine is checked out of the library item.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * The resource ``com.vmware.content.library.Item`` referenced by the parameter ``library_item_id`` requires ``ContentLibrary.DeleteLibraryItem``.
	Delete(libraryItemIdParam string) error

    // Returns the ItemModel with the given identifier.
    //
    // @param libraryItemIdParam  Identifier of the library item to return.
    // The parameter must be an identifier for the resource type: ``com.vmware.content.library.Item``.
    // @return The ItemModel instance with the given ``library_item_id``.
    // @throws NotFound  if no item with the given ``library_item_id`` exists.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * The resource ``com.vmware.content.library.Item`` referenced by the parameter ``library_item_id`` requires ``System.Read``.
	Get(libraryItemIdParam string) (ItemModel, error)

    // Returns the identifiers of all items in the given library.
    //
    // @param libraryIdParam  Identifier of the library whose items should be returned.
    // The parameter must be an identifier for the resource type: ``com.vmware.content.Library``.
    // @return The array of identifiers of the items in the library specified by ``library_id``.
    // The return value will contain identifiers for the resource type: ``com.vmware.content.library.Item``.
    // @throws NotFound  if the library associated with ``library_id`` does not exist.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * The resource ``com.vmware.content.Library`` referenced by the parameter ``library_id`` requires ``System.Read``.
	List(libraryIdParam string) ([]string, error)

    // Returns identifiers of all the visible (as determined by authorization policy) library items matching the requested ItemFindSpec.
    //
    // @param specParam  Specification describing what properties to filter on.
    // @return The array of identifiers of all the visible library items matching the given ``spec``.
    // The return value will contain identifiers for the resource type: ``com.vmware.content.library.Item``.
    // @throws InvalidArgument  if no properties are specified in the ``spec``.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``System.Read``.
    // * The resource ``com.vmware.content.Library`` referenced by the property ItemFindSpec#libraryId requires ``System.Read``.
	Find(specParam ItemFindSpec) ([]string, error)

    // Updates the specified properties of a library item. 
    //
    //  This is an incremental update to the library item. Properties that are null in the update specification are left unchanged. 
    //
    //  This method cannot update a library item that is a member of a subscribed library. Those items must be updated in the source published library and synchronized to the subscribed library.
    //
    // @param libraryItemIdParam  Identifier of the library item to update.
    // The parameter must be an identifier for the resource type: ``com.vmware.content.library.Item``.
    // @param updateSpecParam  Specification of the properties to set.
    // @throws NotFound  if the library item specified by ``library_item_id`` does not exist.
    // @throws InvalidElementType  if the library item corresponding to ``library_item_id`` is a member of a subscribed library.
    // @throws InvalidArgument  if one of the following is true for the ``update_spec``: 
    //
    // * name is empty
    // * name exceeds 80 characters
    // * description exceeds 2000 characters
    // * version is not equal to the current version of the library item
    // @throws NotAllowedInCurrentState  if the library item belongs to a published library with JSON persistence enabled (see PublishInfo#persistJsonEnabled) and the content of the library item specified by ``library_item_id`` has been deleted from the storage backings (see null) associated with it.
    // @throws AlreadyExists  if there is already a library item with same name in the library.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * The resource ``com.vmware.content.library.Item`` referenced by the parameter ``library_item_id`` requires ``ContentLibrary.UpdateLibraryItem``.
	Update(libraryItemIdParam string, updateSpecParam ItemModel) error

    // Publishes the library item to specified subscriptions of the library. If no subscriptions are specified, then publishes the library item to all subscriptions of the library. This method was added in vSphere API 6.7.2.
    //
    // @param libraryItemIdParam Library item identifier.
    // The parameter must be an identifier for the resource type: ``com.vmware.content.library.Item``.
    // @param forceSyncContentParam Whether to synchronize file content as well as metadata. This parameter applies only if the subscription is on-demand.
    // @param subscriptionsParam The list of subscriptions to publish this library item to.
    // @throws Error  If the system reports an error while responding to the request.
    // @throws NotFound  If the library item specified by ``library_item_id`` does not exist.
    // @throws InvalidArgument  If one or more arguments in ``subscriptions`` is not valid.
    // @throws InvalidElementType  If the library item specified by ``library_item_id`` is a member of a subscribed library.
    // @throws NotAllowedInCurrentState  If the library item specified by ``library_item_id`` does not belong to a published library.
    // @throws Unauthenticated  If the user that requested the method cannot be authenticated.
    // @throws Unauthorized  If the user that requested the method is not authorized to perform the method.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * The resource ``com.vmware.content.library.Item`` referenced by the parameter ``library_item_id`` requires ``ContentLibrary.PublishLibraryItem``.
	Publish(libraryItemIdParam string, forceSyncContentParam bool, subscriptionsParam []ItemDestinationSpec) error
}
