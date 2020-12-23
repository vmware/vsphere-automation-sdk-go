/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: MirrorStackStatus
 * Used by client-side stubs.
 */

package segment_monitoring_profile_binding_maps

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

type MirrorStackStatusClient interface {

    // API will get mirror stack status by Segment Monitoring Profile Binding Map.
    //
    // @param tier1IdParam (required)
    // @param segmentIdParam (required)
    // @param segmentMonitoringProfileBindingMapIdParam (required)
    // @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
    // @param enforcementPointPathParam Binding map path enforcemnt point path to remote L3 mirror session (optional)
    // @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
    // @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
    // @param sortAscendingParam (optional)
    // @param sortByParam Field by which records are sorted (optional)
    // @return com.vmware.nsx_policy.model.MirrorStackStatusListResult
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	List(tier1IdParam string, segmentIdParam string, segmentMonitoringProfileBindingMapIdParam string, cursorParam *string, enforcementPointPathParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.MirrorStackStatusListResult, error)
}
