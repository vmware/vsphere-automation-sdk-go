/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Tag
 * Used by client-side stubs.
 */

package tag

import (
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/apis/cis/tagging"
)

// The ``Tag`` interface provides methods to create, read, update, delete, and enumerate tags.
type TagClient interface {


    // Creates a tag. To invoke this method, you need the create tag privilege on the input category.
    //
    // @param createSpecParam  Specification for the new tag to be created.
    // @return The identifier of the created tag.
    // The return value will be an identifier for the resource type: ``com.vmware.cis.tagging.Tag``.
    // @throws AlreadyExists  if the CreateSpec#name provided in the ``create_spec`` is the name of an already existing tag in the input category.
    // @throws InvalidArgument  if any of the input information in the ``create_spec`` is invalid.
    // @throws NotFound  if the category for in the given ``create_spec`` does not exist in the system.
    // @throws Unauthorized  if you do not have the privilege to create tag.
    Create(createSpecParam CreateSpec) (string, error) 


    // Fetches the tag information for the given tag identifier. To invoke this method, you need the read privilege on the tag in order to view the tag info.
    //
    // @param tagIdParam  The identifier of the input tag.
    // The parameter must be an identifier for the resource type: ``com.vmware.cis.tagging.Tag``.
    // @return The tagging.TagModel that corresponds to ``tag_id``.
    // @throws NotFound  if the tag for the given ``tag_id`` does not exist in the system.
    // @throws Unauthorized  if the user does not have the privilege to read the tag.
    Get(tagIdParam string) (tagging.TagModel, error) 


    // Updates an existing tag. To invoke this method, you need the edit privilege on the tag.
    //
    // @param tagIdParam  The identifier of the input tag.
    // The parameter must be an identifier for the resource type: ``com.vmware.cis.tagging.Tag``.
    // @param updateSpecParam  Specification to update the tag.
    // @throws AlreadyExists  if the UpdateSpec#name provided in the ``update_spec`` is the name of an already existing tag in the same category.
    // @throws InvalidArgument  if any of the input information in the ``update_spec`` is invalid.
    // @throws NotFound  if the tag for the given ``tag_id`` does not exist in the system.
    // @throws Unauthorized  if you do not have the privilege to update the tag.
    Update(tagIdParam string, updateSpecParam UpdateSpec) error 


    // Deletes an existing tag. To invoke this method, you need the delete privilege on the tag.
    //
    // @param tagIdParam  The identifier of the input tag.
    // The parameter must be an identifier for the resource type: ``com.vmware.cis.tagging.Tag``.
    // @throws NotFound  if the tag for the given ``tag_id`` does not exist in the system.
    // @throws Unauthorized  if you do not have the privilege to delete the tag.
    // @throws Unauthenticated  if the user can not be authenticated.
    Delete(tagIdParam string) error 


    // Enumerates the tags in the system. To invoke this method, you need read privilege on the individual tags. The array will only contain tags for which you have read privileges.
    // @return The array of resource identifiers for the tags in the system.
    // The return value will contain identifiers for the resource type: ``com.vmware.cis.tagging.Tag``.
    List() ([]string, error) 


    // Enumerates all tags for which the ``used_by_entity`` is part of the tagging.TagModel#usedBy subscribers map with bool value. To invoke this method, you need the read privilege on the individual tags.
    //
    // @param usedByEntityParam  The field on which the results will be filtered.
    // @return The array of resource identifiers for the tags in the system that are used by ``used_by_entity``.
    // The return value will contain identifiers for the resource type: ``com.vmware.cis.tagging.Tag``.
    ListUsedTags(usedByEntityParam string) ([]string, error) 


    // Enumerates all tags for the given category. To invoke this method, you need the read privilege on the given category and the individual tags in that category.
    //
    // @param categoryIdParam  The identifier of the input category.
    // The parameter must be an identifier for the resource type: ``com.vmware.cis.tagging.Category``.
    // @return The array of resource identifiers for the tags in the given input category.
    // The return value will contain identifiers for the resource type: ``com.vmware.cis.tagging.Tag``.
    // @throws NotFound  if the category for the given ``category_id`` does not exist in the system.
    // @throws Unauthorized  if you do not have the privilege to read the category.
    ListTagsForCategory(categoryIdParam string) ([]string, error) 


    // Adds the ``used_by_entity`` to the tagging.TagModel#usedBy subscribers map with bool value. If the ``used_by_entity`` is already in the map with bool value, then this becomes a no-op. To invoke this method, you need the modify tagging.TagModel#usedBy privilege on the tag.
    //
    // @param tagIdParam  The identifier of the input tag.
    // The parameter must be an identifier for the resource type: ``com.vmware.cis.tagging.Tag``.
    // @param usedByEntityParam  The name of the user to be added to the tagging.TagModel#usedBy map with bool value.
    // @throws NotFound  if the tag for the given ``tag_id`` does not exist in the system.
    // @throws Unauthorized  if you do not have the privilege to add an entity to the tagging.TagModel#usedBy field.
    AddToUsedBy(tagIdParam string, usedByEntityParam string) error 


    // Removes the ``used_by_entity`` from the tagging.TagModel#usedBy subscribers set. If the ``used_by_entity`` is not using this tag, then this becomes a no-op. To invoke this method, you need modify tagging.TagModel#usedBy privilege on the tag.
    //
    // @param tagIdParam  The identifier of the input tag.
    // The parameter must be an identifier for the resource type: ``com.vmware.cis.tagging.Tag``.
    // @param usedByEntityParam  The name of the user to be removed from the tagging.TagModel#usedBy map with bool value.
    // @throws NotFound  if the tag for the given ``tag_id`` does not exist in the system.
    // @throws Unauthorized  if you do not have the privilege to remove an entity from the tagging.TagModel#usedBy field.
    RemoveFromUsedBy(tagIdParam string, usedByEntityParam string) error 


    // Revokes all propagating permissions on the given tag. You should then attach a direct permission with tagging privileges on the given tag. To invoke this method, you need tag related privileges (direct or propagating) on the concerned tag.
    //
    // @param tagIdParam  The identifier of the input tag.
    // The parameter must be an identifier for the resource type: ``com.vmware.cis.tagging.Tag``.
    // @throws NotFound  if the tag for the given ``tag_id`` does not exist in the system.
    // @throws Unauthorized  if you do not have the privilege to revoke propagating permissions on the tag.
    RevokePropagatingPermissions(tagIdParam string) error 

}
