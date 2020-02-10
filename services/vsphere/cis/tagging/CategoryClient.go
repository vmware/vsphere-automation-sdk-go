/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Category
 * Used by client-side stubs.
 */

package tagging


// The ``Category`` interface provides methods to create, read, update, delete, and enumerate categories.
type CategoryClient interface {

    // Creates a category. To invoke this method, you need the create category privilege.
    //
    // @param createSpecParam  Specification for the new category to be created.
    // @return The identifier of the created category.
    // The return value will be an identifier for the resource type: ``com.vmware.cis.tagging.Category``.
    // @throws AlreadyExists  if the CategoryCreateSpec#name provided in the ``create_spec`` is the name of an already existing category.
    // @throws InvalidArgument  if any of the information in the ``create_spec`` is invalid.
    // @throws Unauthorized  if you do not have the privilege to create a category.
	Create(createSpecParam CategoryCreateSpec) (string, error)

    // Fetches the category information for the given category identifier. In order to view the category information, you need the read privilege on the category.
    //
    // @param categoryIdParam  The identifier of the input category.
    // The parameter must be an identifier for the resource type: ``com.vmware.cis.tagging.Category``.
    // @return The CategoryModel that corresponds to ``category_id``.
    // @throws NotFound  if the category for the given ``category_id`` does not exist in the system.
    // @throws Unauthorized  if you do not have the privilege to read the category.
	Get(categoryIdParam string) (CategoryModel, error)

    // Updates an existing category. To invoke this method, you need the edit privilege on the category.
    //
    // @param categoryIdParam  The identifier of the category to be updated.
    // The parameter must be an identifier for the resource type: ``com.vmware.cis.tagging.Category``.
    // @param updateSpecParam  Specification to update the category.
    // @throws AlreadyExists  if the CategoryUpdateSpec#name provided in the ``update_spec`` is the name of an already existing category.
    // @throws InvalidArgument  if any of the information in the ``update_spec`` is invalid.
    // @throws NotFound  if the category for the given ``category_id`` does not exist in the system.
    // @throws Unauthorized  if you do not have the privilege to update the category.
	Update(categoryIdParam string, updateSpecParam CategoryUpdateSpec) error

    // Deletes an existing category. To invoke this method, you need the delete privilege on the category.
    //
    // @param categoryIdParam  The identifier of category to be deleted.
    // The parameter must be an identifier for the resource type: ``com.vmware.cis.tagging.Category``.
    // @throws NotFound  if the category for the given ``category_id`` does not exist in the system.
    // @throws Unauthorized  if you do not have the privilege to delete the category.
    // @throws Unauthenticated  if the user can not be authenticated.
	Delete(categoryIdParam string) error

    // Enumerates the categories in the system. To invoke this method, you need the read privilege on the individual categories. The array will only contain those categories for which you have read privileges.
    // @return The array of resource identifiers for the categories in the system.
    // The return value will contain identifiers for the resource type: ``com.vmware.cis.tagging.Category``.
	List() ([]string, error)

    // Enumerates all categories for which the ``used_by_entity`` is part of the CategoryModel#usedBy subscribers map with bool value. To invoke this method, you need the read privilege on the individual categories.
    //
    // @param usedByEntityParam  The field on which the results will be filtered.
    // @return The array of resource identifiers for the categories in the system that are used by ``used_by_entity``.
    // The return value will contain identifiers for the resource type: ``com.vmware.cis.tagging.Category``.
	ListUsedCategories(usedByEntityParam string) ([]string, error)

    // Adds the ``used_by_entity`` to the CategoryModel#usedBy subscribers map with bool value for the specified category. If the ``used_by_entity`` is already in the map with bool value, then this becomes an idempotent no-op. To invoke this method, you need the modify CategoryModel#usedBy privilege on the category.
    //
    // @param categoryIdParam  The identifier of the input category.
    // The parameter must be an identifier for the resource type: ``com.vmware.cis.tagging.Category``.
    // @param usedByEntityParam  The name of the user to be added to the CategoryModel#usedBy map with bool value.
    // @throws NotFound  if the category for the given ``category_id`` does not exist in the system.
    // @throws Unauthorized  if you do not have the privilege to add an entity to the CategoryModel#usedBy field.
	AddToUsedBy(categoryIdParam string, usedByEntityParam string) error

    // Removes the ``used_by_entity`` from the CategoryModel#usedBy subscribers map with bool value. If the ``used_by_entity`` is not using this category, then this becomes a no-op. To invoke this method, you need the modify CategoryModel#usedBy privilege on the category.
    //
    // @param categoryIdParam  The identifier of the input category.
    // The parameter must be an identifier for the resource type: ``com.vmware.cis.tagging.Category``.
    // @param usedByEntityParam  The name of the user to be removed from the CategoryModel#usedBy map with bool value.
    // @throws NotFound  if the category for the given ``category_id`` does not exist in the system.
    // @throws Unauthorized  if you do not have the privilege to remove an entity from the CategoryModel#usedBy field.
	RemoveFromUsedBy(categoryIdParam string, usedByEntityParam string) error

    // Revokes all propagating permissions on the given category. You should then attach a direct permission with tagging privileges on the given category. To invoke this method, you need category related privileges (direct or propagating) on the concerned category.
    //
    // @param categoryIdParam  The identifier of the input category.
    // The parameter must be an identifier for the resource type: ``com.vmware.cis.tagging.Category``.
    // @throws NotFound  if the category for the given ``category_id`` does not exist in the system.
    // @throws Unauthorized  if you do not have the privilege to revoke propagating permissions on the category.
	RevokePropagatingPermissions(categoryIdParam string) error
}
