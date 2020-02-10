/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: TagAssociation
 * Used by client-side stubs.
 */

package tagging

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/vapi/std"
)

// The ``TagAssociation`` interface provides methods to attach, detach, and query tags.
type TagAssociationClient interface {

    // Attaches the given tag to the input object. The tag needs to meet the cardinality (CategoryModel#cardinality) and associability (CategoryModel#associableTypes) criteria in order to be eligible for attachment. If the tag is already attached to the object, then this method is a no-op and an error will not be thrown. To invoke this method, you need the attach tag privilege on the tag and the read privilege on the object.
    //
    // @param tagIdParam  The identifier of the input tag.
    // The parameter must be an identifier for the resource type: ``com.vmware.cis.tagging.Tag``.
    // @param objectIdParam  The identifier of the input object.
    // @throws NotFound  if the tag for the given ``tag_id`` does not exist in the system.
    // @throws InvalidArgument  if the input tag is not eligible to be attached to this object or if the ``object_id`` is not valid.
    // @throws Unauthorized  if you do not have the privilege to attach the tag or do not have the privilege to read the object.
    // @throws Unauthenticated  if the user can not be authenticated.
	Attach(tagIdParam string, objectIdParam std.DynamicID) error

    // Attaches the given tags to the input object. If a tag is already attached to the object, then the individual method is a no-op and an error will not be added to TagAssociationBatchResult#errorMessages. To invoke this method, you need the read privilege on the object and the attach tag privilege on each tag. This method was added in vSphere API 6.5.
    //
    // @param objectIdParam  The identifier of the input object.
    // @param tagIdsParam  The identifiers of the input tags.
    // The parameter must contain identifiers for the resource type: ``com.vmware.cis.tagging.Tag``.
    // @return The outcome of the batch method and the array of error messages (TagAssociationBatchResult#errorMessages) describing attachment failures.
    // @throws Unauthorized  if you do not have the privilege to read the object.
    // @throws Unauthenticated  if the user can not be authenticated.
	AttachMultipleTagsToObject(objectIdParam std.DynamicID, tagIdsParam []string) (TagAssociationBatchResult, error)

    // Attaches the given tag to the input objects. If a tag is already attached to the object, then the individual method is a no-op and an error will not be added to TagAssociationBatchResult#errorMessages. To invoke this method, you need the attach tag privilege on the tag and the read privilege on each object. This method was added in vSphere API 6.5.
    //
    // @param tagIdParam  The identifier of the input tag.
    // The parameter must be an identifier for the resource type: ``com.vmware.cis.tagging.Tag``.
    // @param objectIdsParam  The identifiers of the input objects.
    // @return The outcome of the batch method and the array of error messages (TagAssociationBatchResult#errorMessages) describing attachment failures.
    // @throws NotFound  if the tag for the given ``tag_id`` does not exist in the system.
    // @throws Unauthorized  if you do not have the attach tag privilege on the tag.
    // @throws Unauthenticated  if the user can not be authenticated.
	AttachTagToMultipleObjects(tagIdParam string, objectIdsParam []std.DynamicID) (TagAssociationBatchResult, error)

    // Detaches the tag from the given object. If the tag is already removed from the object, then this method is a no-op and an error will not be thrown. To invoke this method, you need the attach tag privilege on the tag and the read privilege on the object.
    //
    // @param tagIdParam  The identifier of the input tag.
    // The parameter must be an identifier for the resource type: ``com.vmware.cis.tagging.Tag``.
    // @param objectIdParam  The identifier of the input object.
    // @throws NotFound  if the tag for the given ``tag_id`` does not exist in the system.
    // @throws Unauthorized  if you do not have the privilege to detach the tag or do not have the privilege to read the given object.
    // @throws Unauthenticated  if the user can not be authenticated.
	Detach(tagIdParam string, objectIdParam std.DynamicID) error

    // Detaches the given tags from the input object. If a tag is already removed from the object, then the individual method is a no-op and an error will not be added to TagAssociationBatchResult#errorMessages. To invoke this method, you need the read privilege on the object and the attach tag privilege each tag. This method was added in vSphere API 6.5.
    //
    // @param objectIdParam  The identifier of the input object.
    // @param tagIdsParam  The identifiers of the input tags.
    // The parameter must contain identifiers for the resource type: ``com.vmware.cis.tagging.Tag``.
    // @return The outcome of the batch method and the array of error messages (TagAssociationBatchResult#errorMessages) describing detachment failures.
    // @throws Unauthorized  if you do not have the privilege to read the object.
    // @throws Unauthenticated  if the user can not be authenticated.
	DetachMultipleTagsFromObject(objectIdParam std.DynamicID, tagIdsParam []string) (TagAssociationBatchResult, error)

    // Detaches the given tag from the input objects. If a tag is already removed from the object, then the individual method is a no-op and an error will not be added to TagAssociationBatchResult#errorMessages. To invoke this method, you need the attach tag privilege on the tag and the read privilege on each object. This method was added in vSphere API 6.5.
    //
    // @param tagIdParam  The identifier of the input tag.
    // The parameter must be an identifier for the resource type: ``com.vmware.cis.tagging.Tag``.
    // @param objectIdsParam  The identifiers of the input objects.
    // @return The outcome of the batch method and the array of error messages (TagAssociationBatchResult#errorMessages) describing detachment failures.
    // @throws NotFound  if the tag for the given tag does not exist in the system.
    // @throws Unauthorized  if you do not have the attach tag privilege on the tag.
    // @throws Unauthenticated  if the user can not be authenticated.
	DetachTagFromMultipleObjects(tagIdParam string, objectIdsParam []std.DynamicID) (TagAssociationBatchResult, error)

    // Fetches the array of attached objects for the given tag. To invoke this method, you need the read privilege on the input tag. Only those objects for which you have the read privilege will be returned.
    //
    // @param tagIdParam  The identifier of the input tag.
    // The parameter must be an identifier for the resource type: ``com.vmware.cis.tagging.Tag``.
    // @return The array of attached object identifiers.
    // @throws NotFound  if the tag for the given ``tag_id`` does not exist in the system.
    // @throws Unauthorized  if you do not have the privilege to read the tag.
    // @throws Unauthenticated  if the user can not be authenticated.
	ListAttachedObjects(tagIdParam string) ([]std.DynamicID, error)

    // Fetches the array of TagAssociationTagToObjects describing the input tag identifiers and the objects they are attached to. To invoke this method, you need the read privilege on each input tag. The TagAssociationTagToObjects#objectIds will only contain those objects for which you have the read privilege. This method was added in vSphere API 6.5.
    //
    // @param tagIdsParam  The identifiers of the input tags.
    // The parameter must contain identifiers for the resource type: ``com.vmware.cis.tagging.Tag``.
    // @return The array of the tag identifiers to all object identifiers that each tag is attached to.
    // @throws Unauthenticated  if the user can not be authenticated.
	ListAttachedObjectsOnTags(tagIdsParam []string) ([]TagAssociationTagToObjects, error)

    // Fetches the array of tags attached to the given object. To invoke this method, you need the read privilege on the input object. The array will only contain those tags for which you have the read privileges.
    //
    // @param objectIdParam  The identifier of the input object.
    // @return The array of all tag identifiers that correspond to the tags attached to the given object.
    // The return value will contain identifiers for the resource type: ``com.vmware.cis.tagging.Tag``.
    // @throws Unauthorized  if you do not have the privilege to read the object.
    // @throws Unauthenticated  if the user can not be authenticated.
	ListAttachedTags(objectIdParam std.DynamicID) ([]string, error)

    // Fetches the array of TagAssociationObjectToTags describing the input object identifiers and the tags attached to each object. To invoke this method, you need the read privilege on each input object. The TagAssociationObjectToTags#tagIds will only contain those tags for which you have the read privilege. This method was added in vSphere API 6.5.
    //
    // @param objectIdsParam  The identifiers of the input objects.
    // @return The array of the object identifiers to all tag identifiers that are attached to that object.
    // @throws Unauthenticated  if the user can not be authenticated.
	ListAttachedTagsOnObjects(objectIdsParam []std.DynamicID) ([]TagAssociationObjectToTags, error)

    // Fetches the array of attachable tags for the given object, omitting the tags that have already been attached. Criteria for attachability is calculated based on tagging cardinality (CategoryModel#cardinality) and associability (CategoryModel#associableTypes) constructs. To invoke this method, you need the read privilege on the input object. The array will only contain those tags for which you have read privileges.
    //
    // @param objectIdParam  The identifier of the input object.
    // @return The array of tag identifiers that are eligible to be attached to the given object.
    // The return value will contain identifiers for the resource type: ``com.vmware.cis.tagging.Tag``.
    // @throws Unauthorized  if you do not have the privilege to read the object.
    // @throws Unauthenticated  if the user can not be authenticated.
	ListAttachableTags(objectIdParam std.DynamicID) ([]string, error)
}
