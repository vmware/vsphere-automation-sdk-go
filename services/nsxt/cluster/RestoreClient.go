/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Restore
 * Used by client-side stubs.
 */

package cluster

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt/model"
)

type RestoreClient interface {

    // Advance any currently suspended restore operation. The operation might have been suspended because (1) the user had suspended it previously, or (2) the operation is waiting for user input, to be provided as a part of the POST request body. This operation is only valid when a GET cluster/restore/status returns a status with value SUSPENDED. Otherwise, a 409 response is returned.
    //
    // @param advanceClusterRestoreRequestParam (required)
    // @return com.vmware.nsx_policy.model.ClusterRestoreStatus
    // @throws ConcurrentChange  Conflict
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Advance(advanceClusterRestoreRequestParam model.AdvanceClusterRestoreRequest) (model.ClusterRestoreStatus, error)

    // This operation is only valid when a restore is in suspended state. The UI user can cancel any restore operation when the restore is suspended either due to an error, or for a user input. The API user would need to monitor the progression of a restore by calling periodically \"/api/v1/cluster/restore/status\" API. The response object (ClusterRestoreStatus), contains a field \"endpoints\". The API user can cancel the restore process if 'cancel' action is shown in the endpoint field. This operation is only valid when a GET cluster/restore/status returns a status with value SUSPENDED.
    // @return com.vmware.nsx_policy.model.ClusterRestoreStatus
    // @throws ConcurrentChange  Conflict
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Cancel() (model.ClusterRestoreStatus, error)

    // Retry any currently in-progress, failed restore operation. Only the last step of the multi-step restore operation would have failed,and only that step is retried. This operation is only valid when a GET cluster/restore/status returns a status with value FAILED. Otherwise, a 409 response is returned.
    // @return com.vmware.nsx_policy.model.ClusterRestoreStatus
    // @throws ConcurrentChange  Conflict
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Retry() (model.ClusterRestoreStatus, error)

    // Start the restore of an NSX cluster, from some previously backed-up configuration. This operation is only valid when a GET cluster/restore/status returns a status with value NOT_STARTED. Otherwise, a 409 response is returned.
    //
    // @param initiateClusterRestoreRequestParam (required)
    // @return com.vmware.nsx_policy.model.ClusterRestoreStatus
    // @throws ConcurrentChange  Conflict
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Start(initiateClusterRestoreRequestParam model.InitiateClusterRestoreRequest) (model.ClusterRestoreStatus, error)

    // Suspend any currently running restore operation. The restore operation is made up of a number of steps. When this call is issued, any currently running step is allowed to finish (successfully or with errors), and the next step (and therefore the entire restore operation) is suspended until a subsequent resume or cancel call is issued. This operation is only valid when a GET cluster/restore/status returns a status with value RUNNING. Otherwise, a 409 response is returned.
    // @return com.vmware.nsx_policy.model.ClusterRestoreStatus
    // @throws ConcurrentChange  Conflict
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Suspend() (model.ClusterRestoreStatus, error)
}
