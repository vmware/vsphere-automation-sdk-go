/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: File
 * Used by client-side stubs.
 */

package updatesession


// The ``File`` interface provides methods for accessing files within an update session. 
//
//  After an update session is created against a library item, the ``File`` interface can be used to make changes to the underlying library item metadata as well as the content of the files. The following changes can be made: 
//
// * deleting an existing file within the library item. This deletes both the metadata and the content.
// * updating an existing file with new content.
// * adding a new file to the library item.
//
//  
//
//  The above changes are not applied or visible until the session is completed. See UpdateSession.
type FileClient interface {

    // Validates the files in the update session with the referenced identifier and ensures all necessary files are received. In the case where a file is missing, this method will return its name in the FileValidationResult#missingFiles set. The user can add the missing files and try re-validating. For other type of errors, FileValidationResult#invalidFiles will contain the list of invalid files.
    //
    // @param updateSessionIdParam  Identifier of the update session to validate.
    // The parameter must be an identifier for the resource type: ``com.vmware.content.library.item.UpdateSession``.
    // @return A validation result containing missing files or invalid files and the reason why they are invalid.
    // @throws NotFound  if no update session with the given identifier exists.
    // @throws NotAllowedInCurrentState  if the update session is not in the item.UpdateSessionModelState#State_ACTIVE state, or if some of the files that will be uploaded by the client aren't received correctly.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``System.Anonymous``.
	Validate(updateSessionIdParam string) (FileValidationResult, error)

    // Requests file content to be changed (either created, or updated). Depending on the source type of the file, this method will either return an upload endpoint where the client can push the content, or the server will pull from the provided source endpoint. If a file with the same name already exists in this session, this method will be used to update the content of the existing file. 
    //
    //  When importing a file directly from storage, where the source endpoint is a file or datastore URI, you will need to have the ContentLibrary.ReadStorage privilege on the library item. If the file is located in the same directory as the library storage backing folder, the server will move the file instead of copying it, thereby allowing instantaneous import of files for efficient backup and restore scenarios. In all other cases, a copy is performed rather than a move.
    //
    // @param updateSessionIdParam  Identifier of the update session to be modified.
    // The parameter must be an identifier for the resource type: ``com.vmware.content.library.item.UpdateSession``.
    // @param fileSpecParam  Specification for the file that needs to be added or updated. This includes whether the client wants to push the content or have the server pull it.
    // @return An FileInfo class containing upload links as well as server side state tracking the transfer of the file.
    // @throws InvalidArgument  if the ``file_spec`` is invalid.
    // @throws NotFound  if the update session doesn't exist.
    // @throws Unauthorized  if the caller doesn't have ContentLibrary.ReadStorage privilege on the library item of the update session and source type FileSourceType#FileSourceType_PULL is requested for a file or datastore source endpoint (that is, not HTTP or HTTPs based endpoint).
    // @throws NotAllowedInCurrentState  if the content of the library item associated with the update session has been deleted from the storage backings (see null) associated with it.
    // @throws NotAllowedInCurrentState  if metadata files such as manifest and certificate file are added after the OVF descriptor file. This is applicable to update sessions with library item type OVF only. This error was added in vSphere 6.8.0.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``System.Anonymous``.
	Add(updateSessionIdParam string, fileSpecParam FileAddSpec) (FileInfo, error)

    // Requests a file to be removed. The file will only be effectively removed when the update session is completed.
    //
    // @param updateSessionIdParam  Identifier of the update session.
    // The parameter must be an identifier for the resource type: ``com.vmware.content.library.item.UpdateSession``.
    // @param fileNameParam  Name of the file to be removed.
    // @throws NotFound  if the update session doesn't exist.
    // @throws InvalidArgument  if the file doesn't exist in the library item associated with the update session.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``System.Anonymous``.
	Remove(updateSessionIdParam string, fileNameParam string) error

    // Lists all files in the library item associated with the update session.
    //
    // @param updateSessionIdParam  Identifier of the update session.
    // The parameter must be an identifier for the resource type: ``com.vmware.content.library.item.UpdateSession``.
    // @return The array of the files in the library item associated with the update session. This array may be empty if the caller has removed all the files as part of this session (in which case completing the update session will result in an empty library item).
    // @throws NotFound  if the update session doesn't exist.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``System.Anonymous``.
	List(updateSessionIdParam string) ([]FileInfo, error)

    // Retrieves information about a specific file in the snapshot of the library item at the time when the update session was created.
    //
    // @param updateSessionIdParam  Identifier of the update session.
    // The parameter must be an identifier for the resource type: ``com.vmware.content.library.item.UpdateSession``.
    // @param fileNameParam  Name of the file.
    // @return Information about the file.
    // @throws NotFound  if the update session doesn't exist.
    // @throws InvalidArgument  if the file doesn't exist in the library item associated with the update session.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``System.Anonymous``.
	Get(updateSessionIdParam string, fileNameParam string) (FileInfo, error)
}
