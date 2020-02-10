/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: Nodes.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package topology

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)



// The ``ApplianceType`` enumeration class defines values for valid appliance types for the vCenter and Platform Services Controller node. See NodesInfo. This enumeration was added in vSphere API 6.7.2.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type NodesApplianceType string

const (
    // vCenter Server Appliance with an embedded Platform Services Controller. This constant field was added in vSphere API 6.7.2.
	NodesApplianceType_VCSA_EMBEDDED NodesApplianceType = "VCSA_EMBEDDED"
    // vCenter Server Appliance with an external Platform Services Controller. This constant field was added in vSphere API 6.7.2.
	NodesApplianceType_VCSA_EXTERNAL NodesApplianceType = "VCSA_EXTERNAL"
    // An external Platform Services Controller. This constant field was added in vSphere API 6.7.2.
	NodesApplianceType_PSC_EXTERNAL NodesApplianceType = "PSC_EXTERNAL"
)

func (a NodesApplianceType) NodesApplianceType() bool {
	switch a {
	case NodesApplianceType_VCSA_EMBEDDED:
		return true
	case NodesApplianceType_VCSA_EXTERNAL:
		return true
	case NodesApplianceType_PSC_EXTERNAL:
		return true
	default:
		return false
	}
}


// The ``Info`` class contains vCenter or Platform Services Controller node details. This class was added in vSphere API 6.7.2.
type NodesInfo struct {
    // Domain name of the node. This property was added in vSphere API 6.7.2.
	Domain string
    // Appliance type of the node. This property was added in vSphere API 6.7.2.
	Type_ NodesApplianceType
    // List of replication partners' node identifiers. Identifiers can be either IP address or DNS resolvable name of the partner node. This property was added in vSphere API 6.7.2.
	ReplicationPartners []string
    // Identifier of the affinitized Platform Services Controller node. Identifier can be either IP address or DNS resolvable name of the affinitized node. This property was added in vSphere API 6.7.2.
	ClientAffinity *string
}

// The ``Summary`` class contains commonly used information of vCenter or Platform Services Controller node. This class was added in vSphere API 6.7.2.
type NodesSummary struct {
    // Identifier for the vCenter or Platform Services Controller node. Identifier can be either IP address or DNS resolvable name of the node. This property was added in vSphere API 6.7.2.
	Node string
    // Appliance type of the node. This property was added in vSphere API 6.7.2.
	Type_ NodesApplianceType
    // List of replication partners' node identifiers. Identifiers can be either IP address or DNS resolvable name of the partner node. This property was added in vSphere API 6.7.2.
	ReplicationPartners []string
    // Identifier of the affinitized Platform Services Controller node. Identifier can be either IP address or DNS resolvable name of the affinitized node. This property was added in vSphere API 6.7.2.
	ClientAffinity *string
}

// The ``FilterSpec`` class contains property used to filter the results when listing vCenter and Platform Services Controller nodes (see Nodes#list). This class was added in vSphere API 6.7.2.
type NodesFilterSpec struct {
    // Types of the appliance that a vCenter and Platform Services Controller node must be to match the filter (see NodesApplianceType. This property was added in vSphere API 6.7.2.
	Types map[NodesApplianceType]bool
}



func nodesListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["filter"] = bindings.NewOptionalType(bindings.NewReferenceType(NodesFilterSpecBindingType))
	fieldNameMap["filter"] = "Filter"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func nodesListOutputType() bindings.BindingType {
	return bindings.NewListType(bindings.NewReferenceType(NodesSummaryBindingType), reflect.TypeOf([]NodesSummary{}))
}

func nodesListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["filter"] = bindings.NewOptionalType(bindings.NewReferenceType(NodesFilterSpecBindingType))
	fieldNameMap["filter"] = "Filter"
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
		"null",
		"",
		resultHeaders,
		0,
		errorHeaders,
		map[string]int{"Unauthenticated": 401,"Unauthorized": 403,"InvalidArgument": 400})
}

func nodesGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["node"] = bindings.NewIdType([]string{"com.vmware.vcenter.VCenter.name"}, "")
	fieldNameMap["node"] = "Node"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func nodesGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(NodesInfoBindingType)
}

func nodesGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["node"] = bindings.NewIdType([]string{"com.vmware.vcenter.VCenter.name"}, "")
	fieldNameMap["node"] = "Node"
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
		"null",
		"",
		resultHeaders,
		0,
		errorHeaders,
		map[string]int{"Unauthenticated": 401,"Unauthorized": 403,"NotFound": 404})
}


func NodesInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["domain"] = bindings.NewStringType()
	fieldNameMap["domain"] = "Domain"
	fields["type"] = bindings.NewEnumType("com.vmware.vcenter.topology.nodes.appliance_type", reflect.TypeOf(NodesApplianceType(NodesApplianceType_VCSA_EMBEDDED)))
	fieldNameMap["type"] = "Type_"
	fields["replication_partners"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewIdType([]string{"com.vmware.vcenter.VCenter.name"}, ""), reflect.TypeOf([]string{})))
	fieldNameMap["replication_partners"] = "ReplicationPartners"
	fields["client_affinity"] = bindings.NewOptionalType(bindings.NewIdType([]string{"com.vmware.vcenter.VCenter.name"}, ""))
	fieldNameMap["client_affinity"] = "ClientAffinity"
	var validators = []bindings.Validator{}
	uv1 := bindings.NewUnionValidator("type",
		map[string][]bindings.FieldData{
			"VCSA_EMBEDDED": []bindings.FieldData{
				bindings.NewFieldData("replication_partners", true),
			},
			"PSC_EXTERNAL": []bindings.FieldData{
				bindings.NewFieldData("replication_partners", true),
			},
			"VCSA_EXTERNAL": []bindings.FieldData{
				bindings.NewFieldData("client_affinity", true),
			},
		},
	)
	validators = append(validators, uv1)
	return bindings.NewStructType("com.vmware.vcenter.topology.nodes.info", fields, reflect.TypeOf(NodesInfo{}), fieldNameMap, validators)
}

func NodesSummaryBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["node"] = bindings.NewIdType([]string{"com.vmware.vcenter.VCenter.name"}, "")
	fieldNameMap["node"] = "Node"
	fields["type"] = bindings.NewEnumType("com.vmware.vcenter.topology.nodes.appliance_type", reflect.TypeOf(NodesApplianceType(NodesApplianceType_VCSA_EMBEDDED)))
	fieldNameMap["type"] = "Type_"
	fields["replication_partners"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewIdType([]string{"com.vmware.vcenter.VCenter.name"}, ""), reflect.TypeOf([]string{})))
	fieldNameMap["replication_partners"] = "ReplicationPartners"
	fields["client_affinity"] = bindings.NewOptionalType(bindings.NewIdType([]string{"com.vmware.vcenter.VCenter.name"}, ""))
	fieldNameMap["client_affinity"] = "ClientAffinity"
	var validators = []bindings.Validator{}
	uv1 := bindings.NewUnionValidator("type",
		map[string][]bindings.FieldData{
			"VCSA_EMBEDDED": []bindings.FieldData{
				bindings.NewFieldData("replication_partners", true),
			},
			"PSC_EXTERNAL": []bindings.FieldData{
				bindings.NewFieldData("replication_partners", true),
			},
			"VCSA_EXTERNAL": []bindings.FieldData{
				bindings.NewFieldData("client_affinity", true),
			},
		},
	)
	validators = append(validators, uv1)
	return bindings.NewStructType("com.vmware.vcenter.topology.nodes.summary", fields, reflect.TypeOf(NodesSummary{}), fieldNameMap, validators)
}

func NodesFilterSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["types"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewEnumType("com.vmware.vcenter.topology.nodes.appliance_type", reflect.TypeOf(NodesApplianceType(NodesApplianceType_VCSA_EMBEDDED))), reflect.TypeOf(map[NodesApplianceType]bool{})))
	fieldNameMap["types"] = "Types"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.topology.nodes.filter_spec", fields, reflect.TypeOf(NodesFilterSpec{}), fieldNameMap, validators)
}

