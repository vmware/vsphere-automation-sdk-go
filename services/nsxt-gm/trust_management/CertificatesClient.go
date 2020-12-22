/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Certificates
 * Used by client-side stubs.
 */

package trust_management

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt-gm/model"
)

type CertificatesClient interface {

    // Look up the Certificate Profile matching the service-type and apply the certificate. When the Certificate Profile has cluster_certificate=false, the node_id parameter is required to designate the node where the certificate needs to be applied.
    //
    // @param certIdParam ID of certificate to apply (required)
    // @param serviceTypeParam Supported service types, that are using certificates. (required)
    // @param nodeIdParam Node Id (optional)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Applycertificate(certIdParam string, serviceTypeParam string, nodeIdParam *string) error

    // Removes the specified certificate. The private key associated with the certificate is also deleted.
    //
    // @param certIdParam ID of certificate to delete (required)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Delete(certIdParam string) error

    // Returns information for the specified certificate ID, including the certificate's UUID; resource_type (for example, certificate_self_signed, certificate_ca, or certificate_signed); pem_encoded data; and history of the certificate (who created or modified it and when). For additional information, include the ?details=true modifier at the end of the request URI.
    //
    // @param certIdParam ID of certificate to read (required)
    // @param detailsParam whether to expand the pem data and show all its details (optional, default to false)
    // @return com.vmware.nsx_global_policy.model.Certificate
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get(certIdParam string, detailsParam *bool) (model.Certificate, error)

    // Adds a new private-public certificate or a chain of certificates (CAs) and, optionally, a private key that can be applied to one of the user-facing components (appliance management or edge). The certificate and the key should be stored in PEM format. If no private key is provided, the certificate is used as a client certificate in the trust store.
    //
    // @param trustObjectDataParam (required)
    // @return com.vmware.nsx_global_policy.model.CertificateList
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Importcertificate(trustObjectDataParam model.TrustObjectData) (model.CertificateList, error)

    // Returns all certificate information viewable by the user, including each certificate's UUID; resource_type (for example, certificate_self_signed, certificate_ca, or certificate_signed); pem_encoded data; and history of the certificate (who created or modified it and when). For additional information, include the ?details=true modifier at the end of the request URI.
    //
    // @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
    // @param detailsParam whether to expand the pem data and show all its details (optional, default to false)
    // @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
    // @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
    // @param sortAscendingParam (optional)
    // @param sortByParam Field by which records are sorted (optional)
    // @param type_Param Type of certificate to return (optional)
    // @return com.vmware.nsx_global_policy.model.CertificateList
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	List(cursorParam *string, detailsParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string, type_Param *string) (model.CertificateList, error)

    // Set a certificate that has been imported to be the Appliance Proxy certificate used for communicating with Appliance Proxies on other sites.
    //
    // @param setInterSiteAphCertificateRequestParam (required)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Setapplianceproxycertificateforintersitecommunication(setInterSiteAphCertificateRequestParam model.SetInterSiteAphCertificateRequest) error

    // Set a certificate that has been imported to be either the principal identity certificate for the local cluster with either GM or LM service type. Currently, the service type specified must match the current service type of the local cluster.
    //
    // @param setPrincipalIdentityCertificateForFederationRequestParam (required)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Setpicertificateforfederation(setPrincipalIdentityCertificateForFederationRequestParam model.SetPrincipalIdentityCertificateForFederationRequest) error

    // Checks whether certificate is valid. When the certificate contains a chain, the full chain is validated. The usage parameter can be SERVER (default) or CLIENT. This indicates whether the certificate needs to be validated as a server-auth or a client-auth certificate.
    //
    // @param certIdParam ID of certificate to validate (required)
    // @param usageParam Usage Type of the Certificate, SERVER or CLIENT. Default is SERVER (optional)
    // @return com.vmware.nsx_global_policy.model.CertificateCheckingStatus
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Validate(certIdParam string, usageParam *string) (model.CertificateCheckingStatus, error)
}
