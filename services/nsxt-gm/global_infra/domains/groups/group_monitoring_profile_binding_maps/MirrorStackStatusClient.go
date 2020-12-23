/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: MirrorStackStatus
 * Used by client-side stubs.
 */

package group_monitoring_profile_binding_maps

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-gm/model"
)

type MirrorStackStatusClient interface {

    // API will get mirror stack status by Group Monitoring Profile Binding Map
    //
    // @param domainIdParam (required)
    // @param groupIdParam (required)
    // @param groupMonitoringProfileBindingMapIdParam (required)
    // @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
    // @param enforcementPointPathParam Binding map path enforcemnt point path to remote L3 mirror session (optional)
    // @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
    // @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
    // @param sortAscendingParam (optional)
    // @param sortByParam Field by which records are sorted (optional)
    // @return com.vmware.nsx_global_policy.model.MirrorStackStatusListResult
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	List(domainIdParam string, groupIdParam string, groupMonitoringProfileBindingMapIdParam string, cursorParam *string, enforcementPointPathParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.MirrorStackStatusListResult, error)
}
