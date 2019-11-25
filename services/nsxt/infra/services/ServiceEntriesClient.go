/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: ServiceEntries
 * Used by client-side stubs.
 */

package services

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/services/nsxt/model"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
)

type ServiceEntriesClient interface {

    // Delete Service entry
    //
    // @param serviceIdParam Service ID (required)
    // @param serviceEntryIdParam Service entry ID (required)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Delete(serviceIdParam string, serviceEntryIdParam string) error

    // Service entry
    //
    // @param serviceIdParam Service ID (required)
    // @param serviceEntryIdParam Service entry ID (required)
    // @return com.vmware.nsx_policy.model.ServiceEntry
    // The return value will contain all the properties defined in model.ServiceEntry.
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get(serviceIdParam string, serviceEntryIdParam string) (*data.StructValue, error)

    // Paginated list of Service entries for the given service
    //
    // @param serviceIdParam Service ID (required)
    // @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
    // @param includeMarkForDeleteObjectsParam Include objects that are marked for deletion in results (optional, default to false)
    // @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
    // @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
    // @param sortAscendingParam (optional)
    // @param sortByParam Field by which records are sorted (optional)
    // @return com.vmware.nsx_policy.model.ServiceEntryListResult
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	List(serviceIdParam string, cursorParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.ServiceEntryListResult, error)

    // If a service entry with the service-entry-id is not already present, create a new service entry. If it already exists, patch the service entry.
    //
    // @param serviceIdParam Service ID (required)
    // @param serviceEntryIdParam Service entry ID (required)
    // @param serviceEntryParam (required)
    // The parameter must contain all the properties defined in model.ServiceEntry.
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Patch(serviceIdParam string, serviceEntryIdParam string, serviceEntryParam *data.StructValue) error

    // If a service entry with the service-entry-id is not already present, create a new service entry. If it already exists, update the service entry.
    //
    // @param serviceIdParam Service ID (required)
    // @param serviceEntryIdParam Service entry ID (required)
    // @param serviceEntryParam (required)
    // The parameter must contain all the properties defined in model.ServiceEntry.
    // @return com.vmware.nsx_policy.model.ServiceEntry
    // The return value will contain all the properties defined in model.ServiceEntry.
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Update(serviceIdParam string, serviceEntryIdParam string, serviceEntryParam *data.StructValue) (*data.StructValue, error)
}
