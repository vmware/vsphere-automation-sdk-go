/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: CurrentPeerCertificates.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package providers

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/vapi/std"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)



// The ``Summary`` class contains a summary of the current key server certificates. This class was added in vSphere API 7.0.0.
type CurrentPeerCertificatesSummary struct {
    // Name of the server. This property was added in vSphere API 7.0.0.
	ServerName string
    // Server certificate. This property was added in vSphere API 7.0.0.
	Certificate *string
    // Server certificate retrieval errors. 
    //
    //  Specifies error details when retrieving the remote server certificate fails. This list will be empty when CurrentPeerCertificatesSummary#certificate is map with bool value.. This property was added in vSphere API 7.0.0.
	ErrorMessages []std.LocalizableMessage
    // whether server certificate is already trusted . This property was added in vSphere API 7.0.0.
	Trusted bool
}

// The ``FilterSpec`` class contains properties used to filter the results when listing remote server certificates. This class was added in vSphere API 7.0.0.
type CurrentPeerCertificatesFilterSpec struct {
    // Names that key server must have to match the filter (see CurrentPeerCertificatesSummary#serverName). This property was added in vSphere API 7.0.0.
	ServerNames map[string]bool
    // Trust status that server certificates must have to match the filter (see CurrentPeerCertificatesSummary#trusted). This property was added in vSphere API 7.0.0.
	Trusted *bool
}



func currentPeerCertificatesListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cluster"] = bindings.NewIdType([]string{"ClusterComputeResource"}, "")
	fields["provider"] = bindings.NewIdType([]string{"com.vmware.vcenter.trusted_platform.trusted_clusters.kms.Provider"}, "")
	fields["spec"] = bindings.NewOptionalType(bindings.NewReferenceType(CurrentPeerCertificatesFilterSpecBindingType))
	fieldNameMap["cluster"] = "Cluster"
	fieldNameMap["provider"] = "Provider"
	fieldNameMap["spec"] = "Spec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func currentPeerCertificatesListOutputType() bindings.BindingType {
	return bindings.NewListType(bindings.NewReferenceType(CurrentPeerCertificatesSummaryBindingType), reflect.TypeOf([]CurrentPeerCertificatesSummary{}))
}

func currentPeerCertificatesListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["cluster"] = bindings.NewIdType([]string{"ClusterComputeResource"}, "")
	fields["provider"] = bindings.NewIdType([]string{"com.vmware.vcenter.trusted_platform.trusted_clusters.kms.Provider"}, "")
	fields["spec"] = bindings.NewOptionalType(bindings.NewReferenceType(CurrentPeerCertificatesFilterSpecBindingType))
	fieldNameMap["cluster"] = "Cluster"
	fieldNameMap["provider"] = "Provider"
	fieldNameMap["spec"] = "Spec"
	paramsTypeMap["cluster"] = bindings.NewIdType([]string{"ClusterComputeResource"}, "")
	paramsTypeMap["provider"] = bindings.NewIdType([]string{"com.vmware.vcenter.trusted_platform.trusted_clusters.kms.Provider"}, "")
	paramsTypeMap["spec.server_names"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewStringType(), reflect.TypeOf(map[string]bool{})))
	paramsTypeMap["spec.trusted"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["cluster"] = bindings.NewIdType([]string{"ClusterComputeResource"}, "")
	paramsTypeMap["provider"] = bindings.NewIdType([]string{"com.vmware.vcenter.trusted_platform.trusted_clusters.kms.Provider"}, "")
	pathParams["cluster"] = "cluster"
	pathParams["provider"] = "provider"
	queryParams["spec.server_names"] = "server_names"
	queryParams["spec.trusted"] = "trusted"
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
		"/vcenter/trusted-infrastructure/trust-authority-clusters/{cluster}/kms/providers/{provider}/peer-certs/current",
		resultHeaders,
		202,
		errorHeaders,
		map[string]int{"InvalidArgument": 400,"NotFound": 404,"Unauthenticated": 401,"Error": 500})
}


func CurrentPeerCertificatesSummaryBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["server_name"] = bindings.NewStringType()
	fieldNameMap["server_name"] = "ServerName"
	fields["certificate"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["certificate"] = "Certificate"
	fields["error_messages"] = bindings.NewListType(bindings.NewReferenceType(std.LocalizableMessageBindingType), reflect.TypeOf([]std.LocalizableMessage{}))
	fieldNameMap["error_messages"] = "ErrorMessages"
	fields["trusted"] = bindings.NewBooleanType()
	fieldNameMap["trusted"] = "Trusted"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.trusted_infrastructure.trust_authority_clusters.kms.providers.current_peer_certificates.summary", fields, reflect.TypeOf(CurrentPeerCertificatesSummary{}), fieldNameMap, validators)
}

func CurrentPeerCertificatesFilterSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["server_names"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewStringType(), reflect.TypeOf(map[string]bool{})))
	fieldNameMap["server_names"] = "ServerNames"
	fields["trusted"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["trusted"] = "Trusted"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.trusted_infrastructure.trust_authority_clusters.kms.providers.current_peer_certificates.filter_spec", fields, reflect.TypeOf(CurrentPeerCertificatesFilterSpec{}), fieldNameMap, validators)
}

