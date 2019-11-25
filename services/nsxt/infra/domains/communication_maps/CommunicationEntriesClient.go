/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: CommunicationEntries
 * Used by client-side stubs.
 */

package communication_maps

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/services/nsxt/model"
)

type CommunicationEntriesClient interface {

    // Delete CommunicationEntry This API is deprecated. Please use the following API instead. DELETE /infra/domains/domain-id/security-policies/security-policy-id/rules/rule-id
    //
    // @param domainIdParam (required)
    // @param communicationMapIdParam (required)
    // @param communicationEntryIdParam (required)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Delete(domainIdParam string, communicationMapIdParam string, communicationEntryIdParam string) error

    // Read CommunicationEntry This API is deprecated. Please use the following API instead. GET /infra/domains/domain-id/security-policies/security-policy-id/rules/rule-id
    //
    // @param domainIdParam (required)
    // @param communicationMapIdParam (required)
    // @param communicationEntryIdParam (required)
    // @return com.vmware.nsx_policy.model.CommunicationEntry
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get(domainIdParam string, communicationMapIdParam string, communicationEntryIdParam string) (model.CommunicationEntry, error)

    // List CommunicationEntries This API is deprecated. Please use the following API instead. GET /infra/domains/domain-id/security-policies/security-policy-id/rules
    //
    // @param domainIdParam (required)
    // @param communicationMapIdParam (required)
    // @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
    // @param includeMarkForDeleteObjectsParam Include objects that are marked for deletion in results (optional, default to false)
    // @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
    // @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
    // @param sortAscendingParam (optional)
    // @param sortByParam Field by which records are sorted (optional)
    // @return com.vmware.nsx_policy.model.CommunicationEntryListResult
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	List(domainIdParam string, communicationMapIdParam string, cursorParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.CommunicationEntryListResult, error)

    // Patch the CommunicationEntry. If a communication entry for the given communication-entry-id is not present, the object will get created and if it is present it will be updated. This is a full replace This API is deprecated. Please use the following API instead. PATCH /infra/domains/domain-id/security-policies/security-policy-id/rules/rule-id
    //
    // @param domainIdParam (required)
    // @param communicationMapIdParam (required)
    // @param communicationEntryIdParam (required)
    // @param communicationEntryParam (required)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Patch(domainIdParam string, communicationMapIdParam string, communicationEntryIdParam string, communicationEntryParam model.CommunicationEntry) error

    // This is used to re-order a communictation entry within a communication map. This API is deprecated. Please use the following API instead. POST /infra/domains/domain-id/security-policies/security-policy-id/rules/rule-id?action=revise
    //
    // @param domainIdParam (required)
    // @param communicationMapIdParam (required)
    // @param communicationEntryIdParam (required)
    // @param communicationEntryParam (required)
    // @param anchorPathParam The communication map/communication entry path if operation is 'insert_after' or 'insert_before' (optional)
    // @param operationParam Operation (optional, default to insert_top)
    // @return com.vmware.nsx_policy.model.CommunicationEntry
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Revise(domainIdParam string, communicationMapIdParam string, communicationEntryIdParam string, communicationEntryParam model.CommunicationEntry, anchorPathParam *string, operationParam *string) (model.CommunicationEntry, error)

    // Update the CommunicationEntry. If a CommunicationEntry with the communication-entry-id is not already present, this API fails with a 404. Creation of CommunicationEntries is not allowed using this API. This API is deprecated. Please use the following API instead PUT /infra/domains/domain-id/security-policies/securit-policy-id/rules/rule-id
    //
    // @param domainIdParam (required)
    // @param communicationMapIdParam (required)
    // @param communicationEntryIdParam (required)
    // @param communicationEntryParam (required)
    // @return com.vmware.nsx_policy.model.CommunicationEntry
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Update(domainIdParam string, communicationMapIdParam string, communicationEntryIdParam string, communicationEntryParam model.CommunicationEntry) (model.CommunicationEntry, error)
}
