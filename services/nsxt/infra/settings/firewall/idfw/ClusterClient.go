/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Cluster
 * Used by client-side stubs.
 */

package idfw

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt/model"
)

type ClusterClient interface {

    // Delete compute cluster identity firewall configuration.
    //
    // @param clusterIdParam Cluster ID (required)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Delete(clusterIdParam string) error

    // Read compute cluster identity firewall configuration
    //
    // @param clusterIdParam Cluster ID (required)
    // @return com.vmware.nsx_policy.model.ComputeClusterIdfwConfiguration
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get(clusterIdParam string) (model.ComputeClusterIdfwConfiguration, error)

    // API will list all compute cluster wise identity firewall configuration
    //
    // @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
    // @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
    // @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
    // @param sortAscendingParam (optional)
    // @param sortByParam Field by which records are sorted (optional)
    // @return com.vmware.nsx_policy.model.ComputeClusterIdfwConfigurationListResult
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	List(cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.ComputeClusterIdfwConfigurationListResult, error)

    // Patch compute cluster identity firewall configuration.
    //
    // @param clusterIdParam Cluster ID (required)
    // @param computeClusterIdfwConfigurationParam (required)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Patch(clusterIdParam string, computeClusterIdfwConfigurationParam model.ComputeClusterIdfwConfiguration) error

    // Update the compute cluster idfw configuration
    //
    // @param clusterIdParam Cluster ID (required)
    // @param computeClusterIdfwConfigurationParam (required)
    // @return com.vmware.nsx_policy.model.ComputeClusterIdfwConfiguration
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Update(clusterIdParam string, computeClusterIdfwConfigurationParam model.ComputeClusterIdfwConfiguration) (model.ComputeClusterIdfwConfiguration, error)
}
