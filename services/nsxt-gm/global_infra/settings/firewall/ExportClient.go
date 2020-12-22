/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Export
 * Used by client-side stubs.
 */

package firewall

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt-gm/model"
)

type ExportClient interface {

    // This operation cancels an export task. Task needs to be in running state.
    // @return com.vmware.nsx_global_policy.model.ExportTask
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Cancel() (model.ExportTask, error)

    // Invoke export task. There can be only one export task run at any point of time. Hence invocation of another export task will be discarded, when there exist an already running export task. Exported configuration will be in a CSV format. This CSV file will be zipped into a ZIP file, that can be downloaded after the completion of export task.
    //
    // @param exportRequestParameterParam (required)
    // @return com.vmware.nsx_global_policy.model.ExportTask
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Create(exportRequestParameterParam model.ExportRequestParameter) (model.ExportTask, error)

    // Get the information of the latest export task.
    // @return com.vmware.nsx_global_policy.model.ExportTask
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get() (model.ExportTask, error)
}
