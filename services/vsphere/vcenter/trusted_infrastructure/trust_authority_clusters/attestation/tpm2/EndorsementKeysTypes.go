/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: EndorsementKeys.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package tpm2

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/vapi/std"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)

// Resource type for TPM 2.0 endorsement keys. This constant field was added in vSphere API 7.0.0.
const EndorsementKeys_RESOURCE_TYPE = "com.vmware.vcenter.trusted_platform.trusted_clusters.attestation.tpm2.EndorsementKey"


// The ``Health`` enumeration class is indicator for the consistency of the hosts status in the cluster. This enumeration was added in vSphere API 7.0.0.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type EndorsementKeysHealth string

const (
    // No status available. This constant field was added in vSphere API 7.0.0.
	EndorsementKeysHealth_NONE EndorsementKeysHealth = "NONE"
    // Each host in the cluster is in consistent state with the rest hosts in the cluster. This constant field was added in vSphere API 7.0.0.
	EndorsementKeysHealth_OK EndorsementKeysHealth = "OK"
    // Attestation is functioning, however there is an issue that requires attention. This constant field was added in vSphere API 7.0.0.
	EndorsementKeysHealth_WARNING EndorsementKeysHealth = "WARNING"
    // Not all hosts in the cluster are in consistent state. This constant field was added in vSphere API 7.0.0.
	EndorsementKeysHealth_ERROR EndorsementKeysHealth = "ERROR"
)

func (h EndorsementKeysHealth) EndorsementKeysHealth() bool {
	switch h {
	case EndorsementKeysHealth_NONE:
		return true
	case EndorsementKeysHealth_OK:
		return true
	case EndorsementKeysHealth_WARNING:
		return true
	case EndorsementKeysHealth_ERROR:
		return true
	default:
		return false
	}
}


// The ``Summary`` class contains information that summarizes a TPM endorsement key. This class was added in vSphere API 7.0.0.
type EndorsementKeysSummary struct {
    // A unique name for the TPM endorsement key. This property was added in vSphere API 7.0.0.
	Name string
    // A health indicator which indicates whether each host in the cluster has the same endorsement key. This property was added in vSphere API 7.0.0.
	Health EndorsementKeysHealth
}

// The ``Info`` class contains information that describes a TPM endorsement key. This class was added in vSphere API 7.0.0.
type EndorsementKeysInfo struct {
    // TPM public endorsement key in PEM format. This property was added in vSphere API 7.0.0.
	PublicKey string
    // A health indicator which indicates whether each host in the cluster has the same endorsement key. This property was added in vSphere API 7.0.0.
	Health EndorsementKeysHealth
    // Details regarding the health. 
    //
    //  When the ``Health`` is not EndorsementKeysHealth#EndorsementKeysHealth_OK or EndorsementKeysHealth#EndorsementKeysHealth_NONE, this member will provide an actionable description of the issues present.. This property was added in vSphere API 7.0.0.
	Details []std.LocalizableMessage
}

// The ``CreateSpec`` class contains information that describes a TPM endorsement key. 
//
//  Only one of EndorsementKeysCreateSpec#publicKey or EndorsementKeysCreateSpec#certificate must be specified.. This class was added in vSphere API 7.0.0.
type EndorsementKeysCreateSpec struct {
    // A unique name for the TPM endorsement key. 
    //
    //  The unique name should be something that an administrator can use to easily identify the remote system. For example, the hostname, or hardware UUID.. This property was added in vSphere API 7.0.0.
	Name string
    // TPM public endorsement key in PEM format. This property was added in vSphere API 7.0.0.
	PublicKey *string
    // TPM endorsement key certificate in PEM format. 
    //
    //  When a endorsement key certificate is provided, it will be verified against the CA certificate list. Endorsement key certificates that are not signed by one of the CA certificates will be rejected. 
    //
    //  Using this format allows for failures to be caught during configuration rather than later during attestation.. This property was added in vSphere API 7.0.0.
	Certificate *string
}



func endorsementKeysListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cluster"] = bindings.NewIdType([]string{"ClusterComputeResource"}, "")
	fieldNameMap["cluster"] = "Cluster"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func endorsementKeysListOutputType() bindings.BindingType {
	return bindings.NewListType(bindings.NewReferenceType(EndorsementKeysSummaryBindingType), reflect.TypeOf([]EndorsementKeysSummary{}))
}

func endorsementKeysListRestMetadata() protocol.OperationRestMetadata {
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
		"/vcenter/trusted-infrastructure/trust-authority-clusters/{cluster}/attestation/tpm2/endorsement-keys",
		resultHeaders,
		202,
		errorHeaders,
		map[string]int{"Error": 500,"InvalidArgument": 400,"NotFound": 404,"Unauthenticated": 401})
}

func endorsementKeysCreateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cluster"] = bindings.NewIdType([]string{"ClusterComputeResource"}, "")
	fields["spec"] = bindings.NewReferenceType(EndorsementKeysCreateSpecBindingType)
	fieldNameMap["cluster"] = "Cluster"
	fieldNameMap["spec"] = "Spec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func endorsementKeysCreateOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func endorsementKeysCreateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["cluster"] = bindings.NewIdType([]string{"ClusterComputeResource"}, "")
	fields["spec"] = bindings.NewReferenceType(EndorsementKeysCreateSpecBindingType)
	fieldNameMap["cluster"] = "Cluster"
	fieldNameMap["spec"] = "Spec"
	paramsTypeMap["cluster"] = bindings.NewIdType([]string{"ClusterComputeResource"}, "")
	paramsTypeMap["spec"] = bindings.NewReferenceType(EndorsementKeysCreateSpecBindingType)
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
		"/vcenter/trusted-infrastructure/trust-authority-clusters/{cluster}/attestation/tpm2/endorsement-keys",
		resultHeaders,
		202,
		errorHeaders,
		map[string]int{"AlreadyExists": 400,"Error": 500,"InvalidArgument": 400,"NotFound": 404,"Unauthenticated": 401})
}

func endorsementKeysDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cluster"] = bindings.NewIdType([]string{"ClusterComputeResource"}, "")
	fields["name"] = bindings.NewIdType([]string{"com.vmware.vcenter.trusted_platform.trusted_clusters.attestation.tpm2.EndorsementKey"}, "")
	fieldNameMap["cluster"] = "Cluster"
	fieldNameMap["name"] = "Name"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func endorsementKeysDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func endorsementKeysDeleteRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["cluster"] = bindings.NewIdType([]string{"ClusterComputeResource"}, "")
	fields["name"] = bindings.NewIdType([]string{"com.vmware.vcenter.trusted_platform.trusted_clusters.attestation.tpm2.EndorsementKey"}, "")
	fieldNameMap["cluster"] = "Cluster"
	fieldNameMap["name"] = "Name"
	paramsTypeMap["cluster"] = bindings.NewIdType([]string{"ClusterComputeResource"}, "")
	paramsTypeMap["name"] = bindings.NewIdType([]string{"com.vmware.vcenter.trusted_platform.trusted_clusters.attestation.tpm2.EndorsementKey"}, "")
	paramsTypeMap["cluster"] = bindings.NewIdType([]string{"ClusterComputeResource"}, "")
	paramsTypeMap["name"] = bindings.NewIdType([]string{"com.vmware.vcenter.trusted_platform.trusted_clusters.attestation.tpm2.EndorsementKey"}, "")
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
		"/vcenter/trusted-infrastructure/trust-authority-clusters/{cluster}/attestation/tpm2/endorsement-keys/{name}",
		resultHeaders,
		202,
		errorHeaders,
		map[string]int{"Error": 500,"InvalidArgument": 400,"NotFound": 404,"Unauthenticated": 401})
}

func endorsementKeysGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cluster"] = bindings.NewIdType([]string{"ClusterComputeResource"}, "")
	fields["name"] = bindings.NewIdType([]string{"com.vmware.vcenter.trusted_platform.trusted_clusters.attestation.tpm2.EndorsementKey"}, "")
	fieldNameMap["cluster"] = "Cluster"
	fieldNameMap["name"] = "Name"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func endorsementKeysGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(EndorsementKeysInfoBindingType)
}

func endorsementKeysGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["cluster"] = bindings.NewIdType([]string{"ClusterComputeResource"}, "")
	fields["name"] = bindings.NewIdType([]string{"com.vmware.vcenter.trusted_platform.trusted_clusters.attestation.tpm2.EndorsementKey"}, "")
	fieldNameMap["cluster"] = "Cluster"
	fieldNameMap["name"] = "Name"
	paramsTypeMap["cluster"] = bindings.NewIdType([]string{"ClusterComputeResource"}, "")
	paramsTypeMap["name"] = bindings.NewIdType([]string{"com.vmware.vcenter.trusted_platform.trusted_clusters.attestation.tpm2.EndorsementKey"}, "")
	paramsTypeMap["cluster"] = bindings.NewIdType([]string{"ClusterComputeResource"}, "")
	paramsTypeMap["name"] = bindings.NewIdType([]string{"com.vmware.vcenter.trusted_platform.trusted_clusters.attestation.tpm2.EndorsementKey"}, "")
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
		"/vcenter/trusted-infrastructure/trust-authority-clusters/{cluster}/attestation/tpm2/endorsement-keys/{name}",
		resultHeaders,
		202,
		errorHeaders,
		map[string]int{"Error": 500,"InvalidArgument": 400,"NotFound": 404,"Unauthenticated": 401})
}


func EndorsementKeysSummaryBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["name"] = bindings.NewIdType([]string{"com.vmware.vcenter.trusted_platform.trusted_clusters.attestation.tpm2.EndorsementKey"}, "")
	fieldNameMap["name"] = "Name"
	fields["health"] = bindings.NewEnumType("com.vmware.vcenter.trusted_infrastructure.trust_authority_clusters.attestation.tpm2.endorsement_keys.health", reflect.TypeOf(EndorsementKeysHealth(EndorsementKeysHealth_NONE)))
	fieldNameMap["health"] = "Health"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.trusted_infrastructure.trust_authority_clusters.attestation.tpm2.endorsement_keys.summary", fields, reflect.TypeOf(EndorsementKeysSummary{}), fieldNameMap, validators)
}

func EndorsementKeysInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["public_key"] = bindings.NewStringType()
	fieldNameMap["public_key"] = "PublicKey"
	fields["health"] = bindings.NewEnumType("com.vmware.vcenter.trusted_infrastructure.trust_authority_clusters.attestation.tpm2.endorsement_keys.health", reflect.TypeOf(EndorsementKeysHealth(EndorsementKeysHealth_NONE)))
	fieldNameMap["health"] = "Health"
	fields["details"] = bindings.NewListType(bindings.NewReferenceType(std.LocalizableMessageBindingType), reflect.TypeOf([]std.LocalizableMessage{}))
	fieldNameMap["details"] = "Details"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.trusted_infrastructure.trust_authority_clusters.attestation.tpm2.endorsement_keys.info", fields, reflect.TypeOf(EndorsementKeysInfo{}), fieldNameMap, validators)
}

func EndorsementKeysCreateSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["name"] = bindings.NewIdType([]string{"com.vmware.vcenter.trusted_platform.trusted_clusters.attestation.tpm2.EndorsementKey"}, "")
	fieldNameMap["name"] = "Name"
	fields["public_key"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["public_key"] = "PublicKey"
	fields["certificate"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["certificate"] = "Certificate"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.trusted_infrastructure.trust_authority_clusters.attestation.tpm2.endorsement_keys.create_spec", fields, reflect.TypeOf(EndorsementKeysCreateSpec{}), fieldNameMap, validators)
}

