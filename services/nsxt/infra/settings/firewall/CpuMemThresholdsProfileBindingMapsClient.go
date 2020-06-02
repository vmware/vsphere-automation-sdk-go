/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: CpuMemThresholdsProfileBindingMaps
 * Used by client-side stubs.
 */

package firewall

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt/model"
)

type CpuMemThresholdsProfileBindingMapsClient interface {

    // API will delete Firewall CPU Memory Thresholds Profile Binding.
    //
    // @param cpuMemThresholdsProfileBindingMapIdParam Firewall CPU Memory Thresholds Profile Binding Map ID (required)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Delete(cpuMemThresholdsProfileBindingMapIdParam string) error

    // API will get Firewall CPU Memory Thresholds Profile Binding Map.
    //
    // @param cpuMemThresholdsProfileBindingMapIdParam Firewall CPU Memory Thresholds Profile Binding Map ID (required)
    // @return com.vmware.nsx_policy.model.PolicyFirewallCPUMemThresholdsProfileBindingMap
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get(cpuMemThresholdsProfileBindingMapIdParam string) (model.PolicyFirewallCPUMemThresholdsProfileBindingMap, error)

    // API will list all Firewall CPU Memory Thresholds Profile Binding Maps.
    //
    // @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
    // @param includeMarkForDeleteObjectsParam Include objects that are marked for deletion in results (optional, default to false)
    // @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
    // @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
    // @param sortAscendingParam (optional)
    // @param sortByParam Field by which records are sorted (optional)
    // @return com.vmware.nsx_policy.model.PolicyFirewallCPUMemThresholdsProfileBindingMapListResult
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	List(cursorParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.PolicyFirewallCPUMemThresholdsProfileBindingMapListResult, error)

    // API will create or update Firewall CPU Memory Thresholds Profile binding map.
    //
    // @param cpuMemThresholdsProfileBindingMapIdParam Firewall CPU Memory Thresholds Profile Binding Map ID (required)
    // @param policyFirewallCPUMemThresholdsProfileBindingMapParam (required)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Patch(cpuMemThresholdsProfileBindingMapIdParam string, policyFirewallCPUMemThresholdsProfileBindingMapParam model.PolicyFirewallCPUMemThresholdsProfileBindingMap) error

    // API will update Firewall CPU Memory Thresholds Profile Binding Map.
    //
    // @param cpuMemThresholdsProfileBindingMapIdParam Firewall CPU Memory Thresholds Profile Binding Map ID (required)
    // @param policyFirewallCPUMemThresholdsProfileBindingMapParam (required)
    // @return com.vmware.nsx_policy.model.PolicyFirewallCPUMemThresholdsProfileBindingMap
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Update(cpuMemThresholdsProfileBindingMapIdParam string, policyFirewallCPUMemThresholdsProfileBindingMapParam model.PolicyFirewallCPUMemThresholdsProfileBindingMap) (model.PolicyFirewallCPUMemThresholdsProfileBindingMap, error)
}
