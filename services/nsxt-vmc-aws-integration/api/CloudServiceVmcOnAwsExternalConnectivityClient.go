/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: CloudServiceVmcOnAwsExternalConnectivity
 * Used by client-side stubs.
 */

package api

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-vmc-aws-integration/model"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
)

type CloudServiceVmcOnAwsExternalConnectivityClient interface {

    // Attach the specified VIF (Virtual Interface) to the SDDC.
    //
    // @param vifIdParam (required)
    // @param actionParam Direct Connect VIF (Virtual Interface) action (required)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	AttachVif(vifIdParam string, actionParam string) error

    // Create or update the associated Group connection info for reflecting the SDDC’s association state with SDDC group.
    //
    // @param sddcGroupIdParam (required)
    // @param associatedBaseGroupConnectionInfoParam (required)
    // The parameter must contain all the properties defined in model.AssociatedBaseGroupConnectionInfo.
    // @return com.vmware.model.AssociatedBaseGroupConnectionInfo
    // The return value will contain all the properties defined in model.AssociatedBaseGroupConnectionInfo.
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	CreateAssociatedGroupConnectionInfo(sddcGroupIdParam string, associatedBaseGroupConnectionInfoParam *data.StructValue) (*data.StructValue, error)

    // This API is used to create DX BGP related information.
    //
    // @param directConnectBgpInfoParam (required)
    // @return com.vmware.model.DirectConnectBgpInfo
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	CreateDxBgp(directConnectBgpInfoParam model.DirectConnectBgpInfo) (model.DirectConnectBgpInfo, error)

    // Delete the associated Group connection info for the local SDDC before removing the SDDC from the Group.
    //
    // @param sddcGroupIdParam (required)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	DeleteAssociatedGroupConnectionInfo(sddcGroupIdParam string) error

    // Detach and delete a Direct Connect VIF (Virtual Interface) from the SDDC. Detach and delete are one operation in VMC provider, a Direct Connect VIF will be deleted after detached.
    //
    // @param vifIdParam (required)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	DeleteVifById(vifIdParam string) error

    // The API is used to get the associated Group connection info for the local SDDC. It returns the run time attachment state for the local SDDC, others infomations are from DB.
    //
    // @param sddcGroupIdParam (required)
    // @return com.vmware.model.AssociatedBaseGroupConnectionInfo
    // The return value will contain all the properties defined in model.AssociatedBaseGroupConnectionInfo.
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	GetAssociatedGroupConnectionInfo(sddcGroupIdParam string) (*data.StructValue, error)

    // Get Direct Connect BGP related information, including current Autonomous System Number of the VGW attached to the VPC
    // @return com.vmware.model.DirectConnectBgpInfo
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	GetDxBgpInfo() (model.DirectConnectBgpInfo, error)

    // Retrieve routes that are advertised by the SDDC to the various connectivities. This API is a live query to VMC provider.
    //
    // @param connectivityTypeParam SDDC connectivity (optional)
    // @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
    // @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
    // @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
    // @param sortAscendingParam (optional)
    // @param sortByParam Field by which records are sorted (optional)
    // @return com.vmware.model.ExternalSddcRoutesListResult
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	ListAdvertisedExternalRoutes(connectivityTypeParam *string, cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.ExternalSddcRoutesListResult, error)

    // Retrieve routes that are advertised by the SDDC to the various connectivities in CSV foramt.
    //
    // @param connectivityTypeParam SDDC connectivity (optional)
    // @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
    // @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
    // @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
    // @param sortAscendingParam (optional)
    // @param sortByParam Field by which records are sorted (optional)
    // @return com.vmware.model.ExternalSddcRoutesListResultInCsvFormat
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	ListAdvertisedExternalRoutesInCsvFormatCsv(connectivityTypeParam *string, cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.ExternalSddcRoutesListResultInCsvFormat, error)

    // Retrieve BGP routes that are advertised by Direct Connect from VMC provider to on-premise datacenter. This API is a live query to VMC provider.
    // @return com.vmware.model.BGPAdvertisedRoutes
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	ListAdvertisedRoutes() (model.BGPAdvertisedRoutes, error)

    // This API is used to list all associated Group connection infos for the local SDDC. It returns the run time attachment state for the local SDDC, others infomations are from DB.
    // @return com.vmware.model.AssociatedGroupConnectionInfosListResult
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	ListAssociatedGroupConnectionInfos() (model.AssociatedGroupConnectionInfosListResult, error)

    // Return all non-connected VIFs (with states \"avalible\", \"down\", \"pending\" and \"confirming\") and connected VIFs that are available to the SDDC.
    // @return com.vmware.model.VifsListResult
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	ListDirectConnectVifs() (model.VifsListResult, error)

    // Retrieve routes that are learned for the SDDDC. This API is a db query.
    //
    // @param connectivityTypeParam SDDC connectivity (optional)
    // @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
    // @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
    // @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
    // @param sortAscendingParam (optional)
    // @param sortByParam Field by which records are sorted (optional)
    // @return com.vmware.model.ExternalSddcRoutesListResult
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	ListLearnedExternalRoutes(connectivityTypeParam *string, cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.ExternalSddcRoutesListResult, error)

    // Get TGW learned routes for the SDDDC in CSV format.
    //
    // @param connectivityTypeParam SDDC connectivity (optional)
    // @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
    // @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
    // @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
    // @param sortAscendingParam (optional)
    // @param sortByParam Field by which records are sorted (optional)
    // @return com.vmware.model.ExternalSddcRoutesListResultInCsvFormat
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	ListLearnedExternalRoutesInCsvFormatCsv(connectivityTypeParam *string, cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.ExternalSddcRoutesListResultInCsvFormat, error)

    // Retrieve BGP routes that are learned by Direct Connect from on-premise datacenter. This API is a live query to VMC provider.
    // @return com.vmware.model.BGPLearnedRoutes
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	ListLearnedRoutes() (model.BGPLearnedRoutes, error)

    // This API is used to update DX BGP related information. For ASN update, VIFs should be disconnected from the DX VGW before making this API call. The ASN update operation will be synchronous at this point. In the future the user should make use of the Get RealizationStatus call to check update status. While an ASN update call is in progress, any other DX BGP update request will be rejected.
    //
    // @param directConnectBgpInfoParam (required)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	UpdateDxBgpInfo(directConnectBgpInfoParam model.DirectConnectBgpInfo) error
}
