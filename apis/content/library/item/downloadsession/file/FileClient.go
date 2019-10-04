/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: File
 * Used by client-side stubs.
 */

package file

import (
)

// The ``File`` interface provides methods for accessing files within a download session. 
//
//  After a download session is created against a library item, the ``File`` interface can be used to retrieve all downloadable content within the library item. Since the content may not be available immediately in a downloadable form on the server side, the client will have to prepare the file and wait for the file status to become PrepareStatus#PrepareStatus_PREPARED. 
//
//  See DownloadSession.
type FileClient interface {


    // Lists the information of all the files in the library item associated with the download session.
    //
    // @param downloadSessionIdParam  Identifier of the download session.
    // The parameter must be an identifier for the resource type: ``com.vmware.content.library.item.DownloadSession``.
    // @return The array of Info instances.
    // @throws NotFound  if the download session associated with ``download_session_id`` doesn't exist.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``System.Anonymous``.
    List(downloadSessionIdParam string) ([]Info, error) 


    // Requests a file to be prepared for download.
    //
    // @param downloadSessionIdParam  Identifier of the download session.
    // The parameter must be an identifier for the resource type: ``com.vmware.content.library.item.DownloadSession``.
    // @param fileNameParam  Name of the file requested for download.
    // @param endpointTypeParam  Endpoint type request, one of HTTPS, DIRECT. This will determine the type of the Info#downloadEndpoint that is generated when the file is prepared. The EndpointType#EndpointType_DIRECT is only available to users who have the ContentLibrary.ReadStorage privilege.
    // If not specified the default is EndpointType#EndpointType_HTTPS.
    // @return File information containing the status of the request and the download link to the file.
    // @throws NotFound  if the download session does not exist.
    // @throws InvalidArgument  if there is no file with the specified ``file_name``.
    // @throws Unauthorized  if the the download session wasn't created with the ContentLibrary.ReadStorage privilege and the caller requested a EndpointType#EndpointType_DIRECT endpoint type.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``System.Anonymous``.
    Prepare(downloadSessionIdParam string, fileNameParam string, endpointTypeParam *EndpointType) (Info, error) 


    // Retrieves file download information for a specific file.
    //
    // @param downloadSessionIdParam  Identifier of the download session.
    // The parameter must be an identifier for the resource type: ``com.vmware.content.library.item.DownloadSession``.
    // @param fileNameParam  Name of the file requested.
    // @return The Info instance containing the status of the file and its download link if available.
    // @throws NotFound  if the download session associated with ``download_session_id`` does not exist.
    // @throws InvalidArgument  if there is no file with the specified ``file_name``.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``System.Anonymous``.
    Get(downloadSessionIdParam string, fileNameParam string) (Info, error) 

}
