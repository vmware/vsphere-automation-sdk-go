/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Batch
 * Used by client-side stubs.
 */

package batch

import (
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/bindings/cis/tagging"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/vapi/std"
)

// The ``Batch`` interface provides methods to query category, tag, and tag association information in batch. All methods return successfully. Error or insufficient privilege on an individual object will not result in a thrown exception, instead it will just be skipped.
type BatchClient interface {


    // Fetches the category information for the given category identifiers. You need the read privilege on each category to view its information.
    //
    // @param categoryIdsParam  The identifiers of the input categories.
    // The parameter must contain identifiers for the resource type: ``com.vmware.cis.tagging.Category``.
    // @return The array of tagging.CategoryModels that corresponds to ``category_ids``.
    GetCategories(categoryIdsParam []string) ([]tagging.CategoryModel, error) 


    // Fetches all the category information in the system. You need the read privilege on each category to view its information.
    // @return The array of all tagging.CategoryModels in the system.
    GetAllCategories() ([]tagging.CategoryModel, error) 


    // Fetches the tag information for the given tag identifiers. You need the read privilege on each tag to view its information.
    //
    // @param tagIdsParam  The identifiers of the input tags.
    // The parameter must contain identifiers for the resource type: ``com.vmware.cis.tagging.Tag``.
    // @return The array of tagging.TagModels that corresponds to ``tag_ids``.
    GetTags(tagIdsParam []string) ([]tagging.TagModel, error) 


    // Fetches all tag information in the system. You need the read privilege on each tag to view its information.
    // @return The array of all tagging.TagModels in the system.
    GetAllTags() ([]tagging.TagModel, error) 


    // Enumerates all tags for the given \\\\@{term list} of categories. To invoke this method, you need the read privilege on each category and the individual tags in that category.
    //
    // @param categoryIdsParam  The identifiers of the input categories.
    // The parameter must contain identifiers for the resource type: ``com.vmware.cis.tagging.Category``.
    // @return The array of tag identifiers.
    // The return value will contain identifiers for the resource type: ``com.vmware.cis.tagging.Tag``.
    ListTagsForCategories(categoryIdsParam []string) ([]string, error) 


    // Enumerates all tags containing ``tag_name`` in its names. You should note that tag names are only unique under its own category. To invoke this opertaion, you need the read privilege on each tag.
    //
    // @param tagNameParam  The name of the tag.
    // @return The array of tag identifiers for which the tags' names contain ``tag_name``
    // The return value will contain identifiers for the resource type: ``com.vmware.cis.tagging.Tag``.
    FindTagsByName(tagNameParam string) ([]string, error) 


    // Fetches the array of object identifiers to which the input tags are attached. Objects will only appear once in the array regardless of if several input tags are attached to it. To invoke this method, you need the read privilege on each input tag. The array of object identifiers will only contain those objects for which you have read privileges. Note that this method does not preserve the mappings between tag identifier and object identifiers.
    //
    // @param tagIdsParam  The identifiers of the input tags.
    // The parameter must contain identifiers for the resource type: ``com.vmware.cis.tagging.Tag``.
    // @return The array of object identifiers.
    // @throws Unauthenticated  if the user can not be authenticated.
    ListAttachedObjects(tagIdsParam []string) ([]std.DynamicID, error) 


    // Fetches the array of tag identifier to the identifiers of the objects its attached to. To invoke this method, you need the read privilege on each input tag. The array of object identifiers will only contain those objects for which you have read privileges.
    //
    // @param tagIdsParam  The identifiers of the input tags.
    // The parameter must contain identifiers for the resource type: ``com.vmware.cis.tagging.Tag``.
    // @return The array of the tag identifiers to all object identifiers that each is attached to.
    // @throws Unauthenticated  if the user can not be authenticated.
    ListAttachedObjectsOnTags(tagIdsParam []string) ([]TagToObjects, error) 


    // Fetches the array of tag identifier to the identifiers of the objects its attached to for all tags in the system. To invoke this method, you need the read privilege on each tag to be returned. The array of object identifiers will only contain those objects for which you have read privileges.
    // @return The array of the tag identifiers to all object identifiers that each is attached to.
    // @throws Unauthenticated  if the user can not be authenticated.
    ListAllAttachedObjectsOnTags() ([]TagToObjects, error) 


    // Fetches the array of tag identifiers of the tags attached to the input objects. Tags will only appear once in the array regardless of if they are attached to several of the input objects. To invoke this method, you need the read privilege on each input object. The array of tag identifiers will only contain those tags for which you have read privileges. Note that this method does not preserve the mappings between object identifier and tag identifiers.
    //
    // @param objectIdsParam  The identifiers of the input objects.
    // @return The array of the object identifiers to all tag identifiers that are attached to that object.
    // The return value will contain identifiers for the resource type: ``com.vmware.cis.tagging.Tag``.
    // @throws Unauthenticated  if the user can not be authenticated.
    ListAttachedTags(objectIdsParam []std.DynamicID) ([]string, error) 


    // Fetches the array of object identifier to the identifiers of the tags attached to that object. To invoke this method, you need the read privilege on each input object. The array of tag identifiers will only contain those tags for which you have read privileges.
    //
    // @param objectIdsParam  The identifiers of the input objects.
    // @return The array of the object identifiers to all tag identifiers that are attached to that object.
    // @throws Unauthenticated  if the user can not be authenticated.
    ListAttachedTagsOnObjects(objectIdsParam []std.DynamicID) ([]ObjectToTags, error) 

}
