/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: CloudConnectionState
 * Used by client-side stubs.
 */

package edge_nodes

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

type CloudConnectionStateClient interface {

    // Gets Edge cloud connection state. The cloud_connection_status specifies whether the edge is connected to the URL categorization cloud service or not. The url_data_version specifies the version of the data that is downloaded from the URL categorization cloud service. The last_sync_time specifies when was the last time the data was downloaded from the cloud.
    //
    // @param siteIdParam (required)
    // @param enforcementPointIdParam (required)
    // @param edgeClusterIdParam (required)
    // @param edgeNodeIdParam (required)
    // @return com.vmware.nsx_policy.model.CloudConnectionState
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get(siteIdParam string, enforcementPointIdParam string, edgeClusterIdParam string, edgeNodeIdParam string) (model.CloudConnectionState, error)
}
