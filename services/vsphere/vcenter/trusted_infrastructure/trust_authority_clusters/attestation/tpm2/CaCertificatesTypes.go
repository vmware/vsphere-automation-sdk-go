/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: CaCertificates.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package tpm2

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/vapi/std"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/services/vsphere/vcenter/trusted_infrastructure"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)

// Resource type for TPM 2.0 CA certificates. This constant field was added in vSphere API 7.0.0.
const CaCertificates_RESOURCE_TYPE = "com.vmware.vcenter.trusted_platform.trusted_clusters.attestation.tpm2.CaCertificate"


// The ``Health`` enumeration class is indicator for the consistency of the hosts status in the cluster. This enumeration was added in vSphere API 7.0.0.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type CaCertificatesHealth string

const (
    // No status available. This constant field was added in vSphere API 7.0.0.
	CaCertificatesHealth_NONE CaCertificatesHealth = "NONE"
    // Each host in the cluster is in consistent state with the rest hosts in the cluster. This constant field was added in vSphere API 7.0.0.
	CaCertificatesHealth_OK CaCertificatesHealth = "OK"
    // Attestation is funtioning, however there is an issue that requires attention. This constant field was added in vSphere API 7.0.0.
	CaCertificatesHealth_WARNING CaCertificatesHealth = "WARNING"
    // Not all hosts in the cluster are in consistent state. This constant field was added in vSphere API 7.0.0.
	CaCertificatesHealth_ERROR CaCertificatesHealth = "ERROR"
)

func (h CaCertificatesHealth) CaCertificatesHealth() bool {
	switch h {
	case CaCertificatesHealth_NONE:
		return true
	case CaCertificatesHealth_OK:
		return true
	case CaCertificatesHealth_WARNING:
		return true
	case CaCertificatesHealth_ERROR:
		return true
	default:
		return false
	}
}


// The ``Summary`` class contains information that summarizes a TPM CA certificate. This class was added in vSphere API 7.0.0.
type CaCertificatesSummary struct {
    // A unique name for the TPM CA certificate. This property was added in vSphere API 7.0.0.
	Name string
    // A health indicator which indicates whether each host in the cluster has the same CA certs. This property was added in vSphere API 7.0.0.
	Health CaCertificatesHealth
}

// The ``Info`` class contains information that describes a TPM CA certificate. This class was added in vSphere API 7.0.0.
type CaCertificatesInfo struct {
    // The CA certificate chain. This property was added in vSphere API 7.0.0.
	CertChain trusted_infrastructure.X509CertChain
    // A health indicator which indicates whether each host in the cluster has the same CA certs. This property was added in vSphere API 7.0.0.
	Health CaCertificatesHealth
    // Details regarding the health. 
    //
    //  When the ``Health`` is not CaCertificatesHealth#CaCertificatesHealth_OK or CaCertificatesHealth#CaCertificatesHealth_NONE, this member will provide an actionable description of the issues present.. This property was added in vSphere API 7.0.0.
	Details []std.LocalizableMessage
}

// The ``CreateSpec`` class contains information that describes a TPM CA certificate. This class was added in vSphere API 7.0.0.
type CaCertificatesCreateSpec struct {
    // A unique name for the TPM CA certificate. This property was added in vSphere API 7.0.0.
	Name string
    // The CA certificate chain. 
    //
    //  Certificates may either be added one at a time, or as a chain. Adding the certificates as a chain allows the group to be managed as a whole. For example, an entire chain can be deleted in one CaCertificates#delete operation. 
    //
    //  When certificates are added one at a time, the order must be root first, followed by any intermediates. The intermediates certificates must also be ordered in the direction from root to leaf. 
    //
    //  Similarly, when added as a chain the list must be ordered in the direction from root to leaf.. This property was added in vSphere API 7.0.0.
	CertChain *trusted_infrastructure.X509CertChain
}



func caCertificatesListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cluster"] = bindings.NewIdType([]string{"ClusterComputeResource"}, "")
	fieldNameMap["cluster"] = "Cluster"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func caCertificatesListOutputType() bindings.BindingType {
	return bindings.NewListType(bindings.NewReferenceType(CaCertificatesSummaryBindingType), reflect.TypeOf([]CaCertificatesSummary{}))
}

func caCertificatesListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["cluster"] = bindings.NewIdType([]string{"ClusterComputeResource"}, "")
	fieldNameMap["cluster"] = "Cluster"
	paramsTypeMap["cluster"] = bindings.NewIdType([]string{"ClusterComputeResource"}, "")
	paramsTypeMap["cluster"] = bindings.NewIdType([]string{"ClusterComputeResource"}, "")
	pathParams["cluster"] = "cluster"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"",
		"GET",
		"/vcenter/trusted-infrastructure/trust-authority-clusters/{cluster}/attestation/tpm2/ca-certificates",
		resultHeaders,
		202,
		errorHeaders,
		map[string]int{"Error": 500,"InvalidArgument": 400,"NotFound": 404,"Unauthenticated": 401})
}

func caCertificatesCreateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cluster"] = bindings.NewIdType([]string{"ClusterComputeResource"}, "")
	fields["spec"] = bindings.NewReferenceType(CaCertificatesCreateSpecBindingType)
	fieldNameMap["cluster"] = "Cluster"
	fieldNameMap["spec"] = "Spec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func caCertificatesCreateOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func caCertificatesCreateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["cluster"] = bindings.NewIdType([]string{"ClusterComputeResource"}, "")
	fields["spec"] = bindings.NewReferenceType(CaCertificatesCreateSpecBindingType)
	fieldNameMap["cluster"] = "Cluster"
	fieldNameMap["spec"] = "Spec"
	paramsTypeMap["cluster"] = bindings.NewIdType([]string{"ClusterComputeResource"}, "")
	paramsTypeMap["spec"] = bindings.NewReferenceType(CaCertificatesCreateSpecBindingType)
	paramsTypeMap["cluster"] = bindings.NewIdType([]string{"ClusterComputeResource"}, "")
	pathParams["cluster"] = "cluster"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"spec",
		"POST",
		"/vcenter/trusted-infrastructure/trust-authority-clusters/{cluster}/attestation/tpm2/ca-certificates",
		resultHeaders,
		202,
		errorHeaders,
		map[string]int{"AlreadyExists": 400,"Error": 500,"InvalidArgument": 400,"NotFound": 404,"Unauthenticated": 401})
}

func caCertificatesDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cluster"] = bindings.NewIdType([]string{"ClusterComputeResource"}, "")
	fields["name"] = bindings.NewIdType([]string{"com.vmware.vcenter.trusted_platform.trusted_clusters.attestation.tpm2.CaCertificate"}, "")
	fieldNameMap["cluster"] = "Cluster"
	fieldNameMap["name"] = "Name"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func caCertificatesDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func caCertificatesDeleteRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["cluster"] = bindings.NewIdType([]string{"ClusterComputeResource"}, "")
	fields["name"] = bindings.NewIdType([]string{"com.vmware.vcenter.trusted_platform.trusted_clusters.attestation.tpm2.CaCertificate"}, "")
	fieldNameMap["cluster"] = "Cluster"
	fieldNameMap["name"] = "Name"
	paramsTypeMap["cluster"] = bindings.NewIdType([]string{"ClusterComputeResource"}, "")
	paramsTypeMap["name"] = bindings.NewIdType([]string{"com.vmware.vcenter.trusted_platform.trusted_clusters.attestation.tpm2.CaCertificate"}, "")
	paramsTypeMap["cluster"] = bindings.NewIdType([]string{"ClusterComputeResource"}, "")
	paramsTypeMap["name"] = bindings.NewIdType([]string{"com.vmware.vcenter.trusted_platform.trusted_clusters.attestation.tpm2.CaCertificate"}, "")
	pathParams["cluster"] = "cluster"
	pathParams["name"] = "name"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"",
		"DELETE",
		"/vcenter/trusted-infrastructure/trust-authority-clusters/{cluster}/attestation/tpm2/ca-certificates/{name}",
		resultHeaders,
		202,
		errorHeaders,
		map[string]int{"Error": 500,"InvalidArgument": 400,"NotFound": 404,"Unauthenticated": 401})
}

func caCertificatesGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cluster"] = bindings.NewIdType([]string{"ClusterComputeResource"}, "")
	fields["name"] = bindings.NewIdType([]string{"com.vmware.vcenter.trusted_platform.trusted_clusters.attestation.tpm2.CaCertificate"}, "")
	fieldNameMap["cluster"] = "Cluster"
	fieldNameMap["name"] = "Name"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func caCertificatesGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(CaCertificatesInfoBindingType)
}

func caCertificatesGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["cluster"] = bindings.NewIdType([]string{"ClusterComputeResource"}, "")
	fields["name"] = bindings.NewIdType([]string{"com.vmware.vcenter.trusted_platform.trusted_clusters.attestation.tpm2.CaCertificate"}, "")
	fieldNameMap["cluster"] = "Cluster"
	fieldNameMap["name"] = "Name"
	paramsTypeMap["cluster"] = bindings.NewIdType([]string{"ClusterComputeResource"}, "")
	paramsTypeMap["name"] = bindings.NewIdType([]string{"com.vmware.vcenter.trusted_platform.trusted_clusters.attestation.tpm2.CaCertificate"}, "")
	paramsTypeMap["cluster"] = bindings.NewIdType([]string{"ClusterComputeResource"}, "")
	paramsTypeMap["name"] = bindings.NewIdType([]string{"com.vmware.vcenter.trusted_platform.trusted_clusters.attestation.tpm2.CaCertificate"}, "")
	pathParams["cluster"] = "cluster"
	pathParams["name"] = "name"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"",
		"GET",
		"/vcenter/trusted-infrastructure/trust-authority-clusters/{cluster}/attestation/tpm2/ca-certificates/{name}",
		resultHeaders,
		202,
		errorHeaders,
		map[string]int{"Error": 500,"InvalidArgument": 400,"NotFound": 404,"Unauthenticated": 401})
}


func CaCertificatesSummaryBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["name"] = bindings.NewIdType([]string{"com.vmware.vcenter.trusted_platform.trusted_clusters.attestation.tpm2.CaCertificate"}, "")
	fieldNameMap["name"] = "Name"
	fields["health"] = bindings.NewEnumType("com.vmware.vcenter.trusted_infrastructure.trust_authority_clusters.attestation.tpm2.ca_certificates.health", reflect.TypeOf(CaCertificatesHealth(CaCertificatesHealth_NONE)))
	fieldNameMap["health"] = "Health"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.trusted_infrastructure.trust_authority_clusters.attestation.tpm2.ca_certificates.summary", fields, reflect.TypeOf(CaCertificatesSummary{}), fieldNameMap, validators)
}

func CaCertificatesInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cert_chain"] = bindings.NewReferenceType(trusted_infrastructure.X509CertChainBindingType)
	fieldNameMap["cert_chain"] = "CertChain"
	fields["health"] = bindings.NewEnumType("com.vmware.vcenter.trusted_infrastructure.trust_authority_clusters.attestation.tpm2.ca_certificates.health", reflect.TypeOf(CaCertificatesHealth(CaCertificatesHealth_NONE)))
	fieldNameMap["health"] = "Health"
	fields["details"] = bindings.NewListType(bindings.NewReferenceType(std.LocalizableMessageBindingType), reflect.TypeOf([]std.LocalizableMessage{}))
	fieldNameMap["details"] = "Details"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.trusted_infrastructure.trust_authority_clusters.attestation.tpm2.ca_certificates.info", fields, reflect.TypeOf(CaCertificatesInfo{}), fieldNameMap, validators)
}

func CaCertificatesCreateSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["name"] = bindings.NewIdType([]string{"com.vmware.vcenter.trusted_platform.trusted_clusters.attestation.tpm2.CaCertificate"}, "")
	fieldNameMap["name"] = "Name"
	fields["cert_chain"] = bindings.NewOptionalType(bindings.NewReferenceType(trusted_infrastructure.X509CertChainBindingType))
	fieldNameMap["cert_chain"] = "CertChain"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.trusted_infrastructure.trust_authority_clusters.attestation.tpm2.ca_certificates.create_spec", fields, reflect.TypeOf(CaCertificatesCreateSpec{}), fieldNameMap, validators)
}

