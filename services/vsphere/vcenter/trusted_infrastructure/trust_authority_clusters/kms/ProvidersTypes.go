/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: Providers.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package kms

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/vapi/std"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/services/vsphere/vcenter/trusted_infrastructure"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)

// Resource type for a Key Provider. This constant field was added in vSphere API 7.0.0.
const Providers_RESOURCE_TYPE = "com.vmware.vcenter.trusted_platform.trusted_clusters.kms.Provider"


// The ``Health`` enumeration class defines the possible health states. This enumeration was added in vSphere API 7.0.0.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type ProvidersHealth string

const (
    // No status available. This constant field was added in vSphere API 7.0.0.
	ProvidersHealth_NONE ProvidersHealth = "NONE"
    // Health is normal. This constant field was added in vSphere API 7.0.0.
	ProvidersHealth_OK ProvidersHealth = "OK"
    // Health is normal, however there is an issue that requires attention. This constant field was added in vSphere API 7.0.0.
	ProvidersHealth_WARNING ProvidersHealth = "WARNING"
    // Not healthy. This constant field was added in vSphere API 7.0.0.
	ProvidersHealth_ERROR ProvidersHealth = "ERROR"
)

func (h ProvidersHealth) ProvidersHealth() bool {
	switch h {
	case ProvidersHealth_NONE:
		return true
	case ProvidersHealth_OK:
		return true
	case ProvidersHealth_WARNING:
		return true
	case ProvidersHealth_ERROR:
		return true
	default:
		return false
	}
}


// The ``ServerInfo`` class contains properties that describe the status of a key server. This class was added in vSphere API 7.0.0.
type ProvidersServerInfo struct {
    // The connection status health of the server. This property was added in vSphere API 7.0.0.
	Health ProvidersHealth
    // Details regarding the health of the server connection. 
    //
    //  When the service ``Health`` is not ProvidersHealth#ProvidersHealth_OK, this property will provide an actionable description of the issue.. This property was added in vSphere API 7.0.0.
	Details []std.LocalizableMessage
    // Whether this client trusts the server. This property was added in vSphere API 7.0.0.
	ClientTrustServer bool
    // Whether the server trusts this client. This property was added in vSphere API 7.0.0.
	ServerTrustClient bool
    // Name of the server. This property was added in vSphere API 7.0.0.
	Name string
}

// The ``Status`` class contains properties that describe the status of the Key Provider. This class was added in vSphere API 7.0.0.
type ProvidersStatus struct {
    // The health of the provider. This property was added in vSphere API 7.0.0.
	Health ProvidersHealth
    // Details regarding the health of the provider. 
    //
    //  When the service ``Health`` is not ProvidersHealth#ProvidersHealth_OK, this property will provide an actionable description of the issue.. This property was added in vSphere API 7.0.0.
	Details []std.LocalizableMessage
    // Health of the key servers. This property was added in vSphere API 7.0.0.
	Servers []ProvidersServerInfo
}

// The ``Server`` class contains properties that describe a connection endpoint. This class was added in vSphere API 7.0.0.
type ProvidersServer struct {
    // Name of the server. 
    //
    //  A unique string chosen by the client.. This property was added in vSphere API 7.0.0.
	Name string
    // The server's address. This property was added in vSphere API 7.0.0.
	Address trusted_infrastructure.NetworkAddress
}

// The ``KmipServerCreateSpec`` class contains properties that describe Key Management Interoperability Protocol (KMIP) desired key server configuration. This class was added in vSphere API 7.0.0.
type ProvidersKmipServerCreateSpec struct {
    // List of Key Management Interoperability Protocol (KMIP) compliant key servers. 
    //
    //  Key servers must be configured for active-active replication. If the server port is null, a default value for KMIP's port will be used.. This property was added in vSphere API 7.0.0.
	Servers []ProvidersServer
    // Username for authentication. This property was added in vSphere API 7.0.0.
	Username *string
}

// The ``KeyServerCreateSpec`` class contains properties that describe the desired configuration for the key server. This class was added in vSphere API 7.0.0.
type ProvidersKeyServerCreateSpec struct {
    // Type of the key server. This property was added in vSphere API 7.0.0.
	Type_ ProvidersKeyServerCreateSpecType
    // Description of the key server. This property was added in vSphere API 7.0.0.
	Description *string
    // Proxy server configuration. This property was added in vSphere API 7.0.0.
	ProxyServer *trusted_infrastructure.NetworkAddress
    // Connection timeout in seconds. This property was added in vSphere API 7.0.0.
	ConnectionTimeout *int64
    // Configuration information for Key Management Interoperability Protocol (KMIP) based key server. This property was added in vSphere API 7.0.0.
	KmipServer *ProvidersKmipServerCreateSpec
}

// The ``Type`` enumeration class lists the key server types. This enumeration was added in vSphere API 7.0.0.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type ProvidersKeyServerCreateSpecType string

const (
    // Key Management Interoperability Protocol (KMIP) based key management server. This constant field was added in vSphere API 7.0.0.
	ProvidersKeyServerCreateSpecType_KMIP ProvidersKeyServerCreateSpecType = "KMIP"
)

func (t ProvidersKeyServerCreateSpecType) ProvidersKeyServerCreateSpecType() bool {
	switch t {
	case ProvidersKeyServerCreateSpecType_KMIP:
		return true
	default:
		return false
	}
}


// The ``CreateSpec`` class contains properties that describe the desired configuration for a new Key Provider. This class was added in vSphere API 7.0.0.
type ProvidersCreateSpec struct {
    // Name of the provider. 
    //
    //  A unique string chosen by the client.. This property was added in vSphere API 7.0.0.
	Provider string
    // Master key ID created for the provider. 
    //
    //  A unique Key ID.. This property was added in vSphere API 7.0.0.
	MasterKeyId string
    // Key server associated with this Provider. This property was added in vSphere API 7.0.0.
	KeyServer ProvidersKeyServerCreateSpec
}

// The ``KmipServerUpdateSpec`` class contains properties that describe new configuration for KMIP based key server. This class was added in vSphere API 7.0.0.
type ProvidersKmipServerUpdateSpec struct {
    // List of KMIP compliant key servers. 
    //
    //  Key servers must be configured for active-active replication. If the server port is null, a default value for KMIP's port will be used. 
    //
    // . This property was added in vSphere API 7.0.0.
	Servers []ProvidersServer
    // Username for authentication. 
    //
    // . This property was added in vSphere API 7.0.0.
	Username *string
}

// The ``KeyServerUpdateSpec`` class contains properties that describe new configuration for an existing key server. This class was added in vSphere API 7.0.0.
type ProvidersKeyServerUpdateSpec struct {
    // Type of the key server. 
    //
    // . This property was added in vSphere API 7.0.0.
	Type_ *ProvidersKeyServerUpdateSpecType
    // Description of the key server. 
    //
    // . This property was added in vSphere API 7.0.0.
	Description *string
    // Proxy server configuration. 
    //
    // . This property was added in vSphere API 7.0.0.
	ProxyServer *trusted_infrastructure.NetworkAddress
    // Connection timeout in seconds. 
    //
    // . This property was added in vSphere API 7.0.0.
	ConnectionTimeout *int64
    // Configuration information for KMIP based key server. 
    //
    // . This property was added in vSphere API 7.0.0.
	KmipServer *ProvidersKmipServerUpdateSpec
}

// The ``Type`` enumeration class list the key server types. This enumeration was added in vSphere API 7.0.0.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type ProvidersKeyServerUpdateSpecType string

const (
    // Key Management Interoperability Protocol (KMIP) based key management server. This constant field was added in vSphere API 7.0.0.
	ProvidersKeyServerUpdateSpecType_KMIP ProvidersKeyServerUpdateSpecType = "KMIP"
)

func (t ProvidersKeyServerUpdateSpecType) ProvidersKeyServerUpdateSpecType() bool {
	switch t {
	case ProvidersKeyServerUpdateSpecType_KMIP:
		return true
	default:
		return false
	}
}


// The ``UpdateSpec`` class contains properties that describe the new configuration for an existing provider. This class was added in vSphere API 7.0.0.
type ProvidersUpdateSpec struct {
    // Master key identifier created for the provider. 
    //
    //  A unique Key identifier. 
    //
    // . This property was added in vSphere API 7.0.0.
	MasterKeyId *string
    // Key server associated with this provider. 
    //
    // . This property was added in vSphere API 7.0.0.
	KeyServer *ProvidersKeyServerUpdateSpec
}

// The ``Summary`` class contains properties that summarize a provider. This class was added in vSphere API 7.0.0.
type ProvidersSummary struct {
    // Name of the provider. 
    //
    //  A unique string chosen by the client.. This property was added in vSphere API 7.0.0.
	Provider string
    // Health of the provider in the cluster. This property was added in vSphere API 7.0.0.
	Health ProvidersHealth
}

// The ``KmipServerInfo`` class contains properties that describe the current configuration of a KMIP based key server. This class was added in vSphere API 7.0.0.
type ProvidersKmipServerInfo struct {
    // List of KMIP compliant key servers. This property was added in vSphere API 7.0.0.
	Servers []ProvidersServer
    // Username for authentication. 
    //
    // . This property was added in vSphere API 7.0.0.
	Username *string
}

// The ``KeyServerInfo`` class contains properties that describe the current configuration of a key server. This class was added in vSphere API 7.0.0.
type ProvidersKeyServerInfo struct {
    // Type of the key server. This property was added in vSphere API 7.0.0.
	Type_ ProvidersKeyServerInfoType
    // Description of the key server. This property was added in vSphere API 7.0.0.
	Description string
    // Proxy server configuration. 
    //
    // . This property was added in vSphere API 7.0.0.
	ProxyServer *trusted_infrastructure.NetworkAddress
    // Connection timeout in seconds. 
    //
    // . This property was added in vSphere API 7.0.0.
	ConnectionTimeout *int64
    // Configuration information for KMIP based key server. This property was added in vSphere API 7.0.0.
	KmipServer *ProvidersKmipServerInfo
}

// The ``Type`` enumeration class list the key server types. This enumeration was added in vSphere API 7.0.0.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type ProvidersKeyServerInfoType string

const (
    // Key Management Interoperability Protocol (KMIP) based key management server. This constant field was added in vSphere API 7.0.0.
	ProvidersKeyServerInfoType_KMIP ProvidersKeyServerInfoType = "KMIP"
)

func (t ProvidersKeyServerInfoType) ProvidersKeyServerInfoType() bool {
	switch t {
	case ProvidersKeyServerInfoType_KMIP:
		return true
	default:
		return false
	}
}


// The ``Info`` class contains properties that describe the current configuration of a provider. This class was added in vSphere API 7.0.0.
type ProvidersInfo struct {
    // Master key identifier created for the provider. 
    //
    //  A unique Key identifier.. This property was added in vSphere API 7.0.0.
	MasterKeyId string
    // Key server associated with this provider. This property was added in vSphere API 7.0.0.
	KeyServer ProvidersKeyServerInfo
    // Status of the provider in the cluster. This property was added in vSphere API 7.0.0.
	Status ProvidersStatus
}



func providersListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cluster"] = bindings.NewIdType([]string{"ClusterComputeResource"}, "")
	fieldNameMap["cluster"] = "Cluster"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func providersListOutputType() bindings.BindingType {
	return bindings.NewListType(bindings.NewReferenceType(ProvidersSummaryBindingType), reflect.TypeOf([]ProvidersSummary{}))
}

func providersListRestMetadata() protocol.OperationRestMetadata {
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
		"/vcenter/trusted-infrastructure/trust-authority-clusters/{cluster}/kms/providers",
		resultHeaders,
		202,
		errorHeaders,
		map[string]int{"InvalidArgument": 400,"NotFound": 404,"Unauthenticated": 401,"Error": 500})
}

func providersCreateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cluster"] = bindings.NewIdType([]string{"ClusterComputeResource"}, "")
	fields["spec"] = bindings.NewReferenceType(ProvidersCreateSpecBindingType)
	fieldNameMap["cluster"] = "Cluster"
	fieldNameMap["spec"] = "Spec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func providersCreateOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func providersCreateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["cluster"] = bindings.NewIdType([]string{"ClusterComputeResource"}, "")
	fields["spec"] = bindings.NewReferenceType(ProvidersCreateSpecBindingType)
	fieldNameMap["cluster"] = "Cluster"
	fieldNameMap["spec"] = "Spec"
	paramsTypeMap["cluster"] = bindings.NewIdType([]string{"ClusterComputeResource"}, "")
	paramsTypeMap["spec"] = bindings.NewReferenceType(ProvidersCreateSpecBindingType)
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
		"/vcenter/trusted-infrastructure/trust-authority-clusters/{cluster}/kms/providers",
		resultHeaders,
		202,
		errorHeaders,
		map[string]int{"AlreadyExists": 400,"InvalidArgument": 400,"NotFound": 404,"Unauthenticated": 401,"Error": 500})
}

func providersUpdateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cluster"] = bindings.NewIdType([]string{"ClusterComputeResource"}, "")
	fields["provider"] = bindings.NewIdType([]string{"com.vmware.vcenter.trusted_platform.trusted_clusters.kms.Provider"}, "")
	fields["spec"] = bindings.NewReferenceType(ProvidersUpdateSpecBindingType)
	fieldNameMap["cluster"] = "Cluster"
	fieldNameMap["provider"] = "Provider"
	fieldNameMap["spec"] = "Spec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func providersUpdateOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func providersUpdateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["cluster"] = bindings.NewIdType([]string{"ClusterComputeResource"}, "")
	fields["provider"] = bindings.NewIdType([]string{"com.vmware.vcenter.trusted_platform.trusted_clusters.kms.Provider"}, "")
	fields["spec"] = bindings.NewReferenceType(ProvidersUpdateSpecBindingType)
	fieldNameMap["cluster"] = "Cluster"
	fieldNameMap["provider"] = "Provider"
	fieldNameMap["spec"] = "Spec"
	paramsTypeMap["cluster"] = bindings.NewIdType([]string{"ClusterComputeResource"}, "")
	paramsTypeMap["provider"] = bindings.NewIdType([]string{"com.vmware.vcenter.trusted_platform.trusted_clusters.kms.Provider"}, "")
	paramsTypeMap["spec"] = bindings.NewReferenceType(ProvidersUpdateSpecBindingType)
	paramsTypeMap["cluster"] = bindings.NewIdType([]string{"ClusterComputeResource"}, "")
	paramsTypeMap["provider"] = bindings.NewIdType([]string{"com.vmware.vcenter.trusted_platform.trusted_clusters.kms.Provider"}, "")
	pathParams["cluster"] = "cluster"
	pathParams["provider"] = "provider"
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
		"PATCH",
		"/vcenter/trusted-infrastructure/trust-authority-clusters/{cluster}/kms/providers/{provider}",
		resultHeaders,
		202,
		errorHeaders,
		map[string]int{"InvalidArgument": 400,"NotFound": 404,"Unauthenticated": 401,"Error": 500})
}

func providersDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cluster"] = bindings.NewIdType([]string{"ClusterComputeResource"}, "")
	fields["provider"] = bindings.NewIdType([]string{"com.vmware.vcenter.trusted_platform.trusted_clusters.kms.Provider"}, "")
	fieldNameMap["cluster"] = "Cluster"
	fieldNameMap["provider"] = "Provider"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func providersDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func providersDeleteRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["cluster"] = bindings.NewIdType([]string{"ClusterComputeResource"}, "")
	fields["provider"] = bindings.NewIdType([]string{"com.vmware.vcenter.trusted_platform.trusted_clusters.kms.Provider"}, "")
	fieldNameMap["cluster"] = "Cluster"
	fieldNameMap["provider"] = "Provider"
	paramsTypeMap["cluster"] = bindings.NewIdType([]string{"ClusterComputeResource"}, "")
	paramsTypeMap["provider"] = bindings.NewIdType([]string{"com.vmware.vcenter.trusted_platform.trusted_clusters.kms.Provider"}, "")
	paramsTypeMap["cluster"] = bindings.NewIdType([]string{"ClusterComputeResource"}, "")
	paramsTypeMap["provider"] = bindings.NewIdType([]string{"com.vmware.vcenter.trusted_platform.trusted_clusters.kms.Provider"}, "")
	pathParams["cluster"] = "cluster"
	pathParams["provider"] = "provider"
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
		"/vcenter/trusted-infrastructure/trust-authority-clusters/{cluster}/kms/providers/{provider}",
		resultHeaders,
		202,
		errorHeaders,
		map[string]int{"InvalidArgument": 400,"NotFound": 404,"Unauthenticated": 401,"Error": 500})
}

func providersGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cluster"] = bindings.NewIdType([]string{"ClusterComputeResource"}, "")
	fields["provider"] = bindings.NewIdType([]string{"com.vmware.vcenter.trusted_platform.trusted_clusters.kms.Provider"}, "")
	fieldNameMap["cluster"] = "Cluster"
	fieldNameMap["provider"] = "Provider"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func providersGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(ProvidersInfoBindingType)
}

func providersGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["cluster"] = bindings.NewIdType([]string{"ClusterComputeResource"}, "")
	fields["provider"] = bindings.NewIdType([]string{"com.vmware.vcenter.trusted_platform.trusted_clusters.kms.Provider"}, "")
	fieldNameMap["cluster"] = "Cluster"
	fieldNameMap["provider"] = "Provider"
	paramsTypeMap["cluster"] = bindings.NewIdType([]string{"ClusterComputeResource"}, "")
	paramsTypeMap["provider"] = bindings.NewIdType([]string{"com.vmware.vcenter.trusted_platform.trusted_clusters.kms.Provider"}, "")
	paramsTypeMap["cluster"] = bindings.NewIdType([]string{"ClusterComputeResource"}, "")
	paramsTypeMap["provider"] = bindings.NewIdType([]string{"com.vmware.vcenter.trusted_platform.trusted_clusters.kms.Provider"}, "")
	pathParams["cluster"] = "cluster"
	pathParams["provider"] = "provider"
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
		"/vcenter/trusted-infrastructure/trust-authority-clusters/{cluster}/kms/providers/{provider}",
		resultHeaders,
		202,
		errorHeaders,
		map[string]int{"InvalidArgument": 400,"NotFound": 404,"Unauthenticated": 401,"Error": 500})
}


func ProvidersServerInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["health"] = bindings.NewEnumType("com.vmware.vcenter.trusted_infrastructure.trust_authority_clusters.kms.providers.health", reflect.TypeOf(ProvidersHealth(ProvidersHealth_NONE)))
	fieldNameMap["health"] = "Health"
	fields["details"] = bindings.NewListType(bindings.NewReferenceType(std.LocalizableMessageBindingType), reflect.TypeOf([]std.LocalizableMessage{}))
	fieldNameMap["details"] = "Details"
	fields["client_trust_server"] = bindings.NewBooleanType()
	fieldNameMap["client_trust_server"] = "ClientTrustServer"
	fields["server_trust_client"] = bindings.NewBooleanType()
	fieldNameMap["server_trust_client"] = "ServerTrustClient"
	fields["name"] = bindings.NewStringType()
	fieldNameMap["name"] = "Name"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.trusted_infrastructure.trust_authority_clusters.kms.providers.server_info", fields, reflect.TypeOf(ProvidersServerInfo{}), fieldNameMap, validators)
}

func ProvidersStatusBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["health"] = bindings.NewEnumType("com.vmware.vcenter.trusted_infrastructure.trust_authority_clusters.kms.providers.health", reflect.TypeOf(ProvidersHealth(ProvidersHealth_NONE)))
	fieldNameMap["health"] = "Health"
	fields["details"] = bindings.NewListType(bindings.NewReferenceType(std.LocalizableMessageBindingType), reflect.TypeOf([]std.LocalizableMessage{}))
	fieldNameMap["details"] = "Details"
	fields["servers"] = bindings.NewListType(bindings.NewReferenceType(ProvidersServerInfoBindingType), reflect.TypeOf([]ProvidersServerInfo{}))
	fieldNameMap["servers"] = "Servers"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.trusted_infrastructure.trust_authority_clusters.kms.providers.status", fields, reflect.TypeOf(ProvidersStatus{}), fieldNameMap, validators)
}

func ProvidersServerBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["name"] = bindings.NewStringType()
	fieldNameMap["name"] = "Name"
	fields["address"] = bindings.NewReferenceType(trusted_infrastructure.NetworkAddressBindingType)
	fieldNameMap["address"] = "Address"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.trusted_infrastructure.trust_authority_clusters.kms.providers.server", fields, reflect.TypeOf(ProvidersServer{}), fieldNameMap, validators)
}

func ProvidersKmipServerCreateSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["servers"] = bindings.NewListType(bindings.NewReferenceType(ProvidersServerBindingType), reflect.TypeOf([]ProvidersServer{}))
	fieldNameMap["servers"] = "Servers"
	fields["username"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["username"] = "Username"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.trusted_infrastructure.trust_authority_clusters.kms.providers.kmip_server_create_spec", fields, reflect.TypeOf(ProvidersKmipServerCreateSpec{}), fieldNameMap, validators)
}

func ProvidersKeyServerCreateSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["type"] = bindings.NewEnumType("com.vmware.vcenter.trusted_infrastructure.trust_authority_clusters.kms.providers.key_server_create_spec.type", reflect.TypeOf(ProvidersKeyServerCreateSpecType(ProvidersKeyServerCreateSpecType_KMIP)))
	fieldNameMap["type"] = "Type_"
	fields["description"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["description"] = "Description"
	fields["proxy_server"] = bindings.NewOptionalType(bindings.NewReferenceType(trusted_infrastructure.NetworkAddressBindingType))
	fieldNameMap["proxy_server"] = "ProxyServer"
	fields["connection_timeout"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["connection_timeout"] = "ConnectionTimeout"
	fields["kmip_server"] = bindings.NewOptionalType(bindings.NewReferenceType(ProvidersKmipServerCreateSpecBindingType))
	fieldNameMap["kmip_server"] = "KmipServer"
	var validators = []bindings.Validator{}
	uv1 := bindings.NewUnionValidator("type",
		map[string][]bindings.FieldData{
			"KMIP": []bindings.FieldData{
				bindings.NewFieldData("kmip_server", true),
			},
		},
	)
	validators = append(validators, uv1)
	return bindings.NewStructType("com.vmware.vcenter.trusted_infrastructure.trust_authority_clusters.kms.providers.key_server_create_spec", fields, reflect.TypeOf(ProvidersKeyServerCreateSpec{}), fieldNameMap, validators)
}

func ProvidersCreateSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["provider"] = bindings.NewIdType([]string{"com.vmware.vcenter.trusted_platform.trusted_clusters.kms.Provider"}, "")
	fieldNameMap["provider"] = "Provider"
	fields["master_key_id"] = bindings.NewStringType()
	fieldNameMap["master_key_id"] = "MasterKeyId"
	fields["key_server"] = bindings.NewReferenceType(ProvidersKeyServerCreateSpecBindingType)
	fieldNameMap["key_server"] = "KeyServer"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.trusted_infrastructure.trust_authority_clusters.kms.providers.create_spec", fields, reflect.TypeOf(ProvidersCreateSpec{}), fieldNameMap, validators)
}

func ProvidersKmipServerUpdateSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["servers"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(ProvidersServerBindingType), reflect.TypeOf([]ProvidersServer{})))
	fieldNameMap["servers"] = "Servers"
	fields["username"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["username"] = "Username"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.trusted_infrastructure.trust_authority_clusters.kms.providers.kmip_server_update_spec", fields, reflect.TypeOf(ProvidersKmipServerUpdateSpec{}), fieldNameMap, validators)
}

func ProvidersKeyServerUpdateSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["type"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.trusted_infrastructure.trust_authority_clusters.kms.providers.key_server_update_spec.type", reflect.TypeOf(ProvidersKeyServerUpdateSpecType(ProvidersKeyServerUpdateSpecType_KMIP))))
	fieldNameMap["type"] = "Type_"
	fields["description"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["description"] = "Description"
	fields["proxy_server"] = bindings.NewOptionalType(bindings.NewReferenceType(trusted_infrastructure.NetworkAddressBindingType))
	fieldNameMap["proxy_server"] = "ProxyServer"
	fields["connection_timeout"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["connection_timeout"] = "ConnectionTimeout"
	fields["kmip_server"] = bindings.NewOptionalType(bindings.NewReferenceType(ProvidersKmipServerUpdateSpecBindingType))
	fieldNameMap["kmip_server"] = "KmipServer"
	var validators = []bindings.Validator{}
	uv1 := bindings.NewUnionValidator("type",
		map[string][]bindings.FieldData{
			"KMIP": []bindings.FieldData{
				bindings.NewFieldData("kmip_server", false),
			},
		},
	)
	validators = append(validators, uv1)
	return bindings.NewStructType("com.vmware.vcenter.trusted_infrastructure.trust_authority_clusters.kms.providers.key_server_update_spec", fields, reflect.TypeOf(ProvidersKeyServerUpdateSpec{}), fieldNameMap, validators)
}

func ProvidersUpdateSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["master_key_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["master_key_id"] = "MasterKeyId"
	fields["key_server"] = bindings.NewOptionalType(bindings.NewReferenceType(ProvidersKeyServerUpdateSpecBindingType))
	fieldNameMap["key_server"] = "KeyServer"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.trusted_infrastructure.trust_authority_clusters.kms.providers.update_spec", fields, reflect.TypeOf(ProvidersUpdateSpec{}), fieldNameMap, validators)
}

func ProvidersSummaryBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["provider"] = bindings.NewIdType([]string{"com.vmware.vcenter.trusted_platform.trusted_clusters.kms.Provider"}, "")
	fieldNameMap["provider"] = "Provider"
	fields["health"] = bindings.NewEnumType("com.vmware.vcenter.trusted_infrastructure.trust_authority_clusters.kms.providers.health", reflect.TypeOf(ProvidersHealth(ProvidersHealth_NONE)))
	fieldNameMap["health"] = "Health"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.trusted_infrastructure.trust_authority_clusters.kms.providers.summary", fields, reflect.TypeOf(ProvidersSummary{}), fieldNameMap, validators)
}

func ProvidersKmipServerInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["servers"] = bindings.NewListType(bindings.NewReferenceType(ProvidersServerBindingType), reflect.TypeOf([]ProvidersServer{}))
	fieldNameMap["servers"] = "Servers"
	fields["username"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["username"] = "Username"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.trusted_infrastructure.trust_authority_clusters.kms.providers.kmip_server_info", fields, reflect.TypeOf(ProvidersKmipServerInfo{}), fieldNameMap, validators)
}

func ProvidersKeyServerInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["type"] = bindings.NewEnumType("com.vmware.vcenter.trusted_infrastructure.trust_authority_clusters.kms.providers.key_server_info.type", reflect.TypeOf(ProvidersKeyServerInfoType(ProvidersKeyServerInfoType_KMIP)))
	fieldNameMap["type"] = "Type_"
	fields["description"] = bindings.NewStringType()
	fieldNameMap["description"] = "Description"
	fields["proxy_server"] = bindings.NewOptionalType(bindings.NewReferenceType(trusted_infrastructure.NetworkAddressBindingType))
	fieldNameMap["proxy_server"] = "ProxyServer"
	fields["connection_timeout"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["connection_timeout"] = "ConnectionTimeout"
	fields["kmip_server"] = bindings.NewOptionalType(bindings.NewReferenceType(ProvidersKmipServerInfoBindingType))
	fieldNameMap["kmip_server"] = "KmipServer"
	var validators = []bindings.Validator{}
	uv1 := bindings.NewUnionValidator("type",
		map[string][]bindings.FieldData{
			"KMIP": []bindings.FieldData{
				bindings.NewFieldData("kmip_server", true),
			},
		},
	)
	validators = append(validators, uv1)
	return bindings.NewStructType("com.vmware.vcenter.trusted_infrastructure.trust_authority_clusters.kms.providers.key_server_info", fields, reflect.TypeOf(ProvidersKeyServerInfo{}), fieldNameMap, validators)
}

func ProvidersInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["master_key_id"] = bindings.NewStringType()
	fieldNameMap["master_key_id"] = "MasterKeyId"
	fields["key_server"] = bindings.NewReferenceType(ProvidersKeyServerInfoBindingType)
	fieldNameMap["key_server"] = "KeyServer"
	fields["status"] = bindings.NewReferenceType(ProvidersStatusBindingType)
	fieldNameMap["status"] = "Status"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.trusted_infrastructure.trust_authority_clusters.kms.providers.info", fields, reflect.TypeOf(ProvidersInfo{}), fieldNameMap, validators)
}

