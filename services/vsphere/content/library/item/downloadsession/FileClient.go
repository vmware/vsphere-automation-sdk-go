/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: File
 * Used by client-side stubs.
 */

package downloadsession


// The ``File`` interface provides methods for accessing files within a download session. 
//
//  After a download session is created against a library item, the ``File`` interface can be used to retrieve all downloadable content within the library item. Since the content may not be available immediately in a downloadable form on the server side, the client will have to prepare the file and wait for the file status to become FilePrepareStatus#FilePrepareStatus_PREPARED. 
//
//  See DownloadSession.
type FileClient interface {

    // Lists the information of all the files in the library item associated with the download session.
    //
    // @param downloadSessionIdParam  Identifier of the download session.
    // The parameter must be an identifier for the resource type: ``com.vmware.content.library.item.DownloadSession``.
    // @return The array of FileInfo instances.
    // @throws NotFound  if the download session associated with ``download_session_id`` doesn't exist.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``System.Anonymous``.
	List(downloadSessionIdParam string) ([]FileInfo, error)

    // Requests a file to be prepared for download.
    //
    // @param downloadSessionIdParam  Identifier of the download session.
    // The parameter must be an identifier for the resource type: ``com.vmware.content.library.item.DownloadSession``.
    // @param fileNameParam  Name of the file requested for download.
    // @param endpointTypeParam  Endpoint type request, one of HTTPS, DIRECT. This will determine the type of the FileInfo#downloadEndpoint that is generated when the file is prepared. The FileEndpointType#FileEndpointType_DIRECT is only available to users who have the ContentLibrary.ReadStorage privilege.
    // If not specified the default is FileEndpointType#FileEndpointType_HTTPS.
    // @return File information containing the status of the request and the download link to the file.
    // @throws NotFound  if the download session does not exist.
    // @throws InvalidArgument  if there is no file with the specified ``file_name``.
    // @throws Unauthorized  if the the download session wasn't created with the ContentLibrary.ReadStorage privilege and the caller requested a FileEndpointType#FileEndpointType_DIRECT endpoint type.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``System.Anonymous``.
	Prepare(downloadSessionIdParam string, fileNameParam string, endpointTypeParam *FileEndpointType) (FileInfo, error)

    // Retrieves file download information for a specific file.
    //
    // @param downloadSessionIdParam  Identifier of the download session.
    // The parameter must be an identifier for the resource type: ``com.vmware.content.library.item.DownloadSession``.
    // @param fileNameParam  Name of the file requested.
    // @return The FileInfo instance containing the status of the file and its download link if available.
    // @throws NotFound  if the download session associated with ``download_session_id`` does not exist.
    // @throws InvalidArgument  if there is no file with the specified ``file_name``.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``System.Anonymous``.
	Get(downloadSessionIdParam string, fileNameParam string) (FileInfo, error)
}
