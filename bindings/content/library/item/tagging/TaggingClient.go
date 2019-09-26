/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Tagging
 * Used by client-side stubs.
 */

package tagging

import (
)

// A service for managing tags on library items. A library item can be tagged by one or multiple key/value pairs. These can be used to add extra information to the library item. Library item tags will be replicated through the VCSP protocol. If a library item belongs to a published library, the subscribed library item will have the same tags once a sync happens. Note that tags cannot be set on library items that are part of a subscribed library.
type TaggingClient interface {


    // Sets a key/value pair on a library item. If the key has been already set in the past, this method will override the old value with the new one.
    //
    // @param libraryItemIdParam  the ID of the library item
    // The parameter must be an identifier for the resource type: ``com.vmware.content.library.Item``.
    // @param keyParam  the key to set
    // @param valueParam  the value to set the key to
    // @throws NotFound  if the library item cannot be found
    // @throws InvalidElementType  if the library item is a member of a subscribed library.
    Create(libraryItemIdParam string, keyParam string, valueParam string) error 


    // Removes the key from the library item tags. If the library item doesn't have such a key this call will be a no-op.
    //
    // @param libraryItemIdParam  the ID of the library item
    // The parameter must be an identifier for the resource type: ``com.vmware.content.library.Item``.
    // @param keyParam  the key to remove
    // @throws NotFound  if the library item cannot be found
    // @throws InvalidElementType  if the library item is a member of a subscribed library.
    Delete(libraryItemIdParam string, keyParam string) error 


    // Lists all tags associated with the library item. If no tag is found, the method will return an empty list.
    //
    // @param libraryItemIdParam  the ID of the library item
    // The parameter must be an identifier for the resource type: ``com.vmware.content.library.Item``.
    // @return the list of tags associated with the library item
    // @throws NotFound  if the library item cannot be found
    List(libraryItemIdParam string) ([]Info, error) 


    // Gets the value of the key associated with the library item. If no key is defined this method will return null.
    //
    // @param libraryItemIdParam  the ID of the library item
    // The parameter must be an identifier for the resource type: ``com.vmware.content.library.Item``.
    // @param keyParam  the key to lookup
    // @return the value of the key or null if unset
    // If a tag key is set with an unspecified value the result will be unspecified.
    // @throws NotFound  if the library item cannot be found
    Get(libraryItemIdParam string, keyParam string) (*string, error) 

}
