/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: ResourcePool.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package vcenter

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)

// The resource type for the vCenter resource pool
const ResourcePool_RESOURCE_TYPE = "ResourcePool"


// The ``SharesInfo`` class provides specification of shares. 
//
//  Shares are used to determine relative allocation between resource consumers. In general, a consumer with more shares gets proportionally more of the resource, subject to certain other constraints.. This class was added in vSphere API 7.0.0.
type ResourcePoolSharesInfo struct {
    // The allocation level. It maps to a pre-determined set of numeric values for shares. If the shares value does not map to a predefined size, then the level is set as CUSTOM. This property was added in vSphere API 7.0.0.
	Level ResourcePoolSharesInfoLevel
    // When ResourcePoolSharesInfo#level is set to CUSTOM, it is the number of shares allocated. Otherwise, this value is ignored. 
    //
    //  There is no unit for this value. It is a relative measure based on the settings for other resource pools.. This property was added in vSphere API 7.0.0.
	Shares *int64
}

// The ``Level`` enumeration class defines the possible values for the allocation level. This enumeration was added in vSphere API 7.0.0.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type ResourcePoolSharesInfoLevel string

const (
    // For CPU: Shares = 500 \* number of virtual CPUs.
    //  For Memory: Shares = 5 \* virtual machine memory size in MB.
    // . This constant field was added in vSphere API 7.0.0.
	ResourcePoolSharesInfoLevel_LOW ResourcePoolSharesInfoLevel = "LOW"
    // For CPU: Shares = 1000 \* number of virtual CPUs.
    //  For Memory: Shares = 10 \* virtual machine memory size in MB.
    // . This constant field was added in vSphere API 7.0.0.
	ResourcePoolSharesInfoLevel_NORMAL ResourcePoolSharesInfoLevel = "NORMAL"
    // For CPU: Shares = 2000 \* nmumber of virtual CPUs.
    //  For Memory: Shares = 20 \* virtual machine memory size in MB.
    // . This constant field was added in vSphere API 7.0.0.
	ResourcePoolSharesInfoLevel_HIGH ResourcePoolSharesInfoLevel = "HIGH"
    // If map with bool value, in case there is resource contention the server uses the shares value to determine the resource allocation. This constant field was added in vSphere API 7.0.0.
	ResourcePoolSharesInfoLevel_CUSTOM ResourcePoolSharesInfoLevel = "CUSTOM"
)

func (l ResourcePoolSharesInfoLevel) ResourcePoolSharesInfoLevel() bool {
	switch l {
	case ResourcePoolSharesInfoLevel_LOW:
		return true
	case ResourcePoolSharesInfoLevel_NORMAL:
		return true
	case ResourcePoolSharesInfoLevel_HIGH:
		return true
	case ResourcePoolSharesInfoLevel_CUSTOM:
		return true
	default:
		return false
	}
}


// The ``ResourceAllocationInfo`` class contains resource allocation information of a resource pool. This class was added in vSphere API 7.0.0.
type ResourcePoolResourceAllocationInfo struct {
    // Amount of resource that is guaranteed available to a resource pool. Reserved resources are not wasted if they are not used. If the utilization is less than the reservation, the resources can be utilized by other running virtual machines. Units are MB fo memory, and MHz for CPU. This property was added in vSphere API 7.0.0.
	Reservation int64
    // In a resource pool with an expandable reservation, the reservation can grow beyond the specified value, if the parent resource pool has unreserved resources. A non-expandable reservation is called a fixed reservation. This property was added in vSphere API 7.0.0.
	ExpandableReservation bool
    // The utilization of a resource pool will not exceed this limit, even if there are available resources. This is typically used to ensure a consistent performance of resource pools independent of available resources. If set to -1, then there is no fixed limit on resource usage (only bounded by available resources and shares). Units are MB for memory, and MHz for CPU. This property was added in vSphere API 7.0.0.
	Limit int64
    // Shares are used in case of resource contention. This property was added in vSphere API 7.0.0.
	Shares ResourcePoolSharesInfo
}

// The ``Info`` class contains information about a resource pool.
type ResourcePoolInfo struct {
    // Name of the vCenter Server resource pool.
	Name string
    // Identifiers of the child resource pools contained in this resource pool.
	ResourcePools map[string]bool
    // Resource allocation information for CPU. This property was added in vSphere API 7.0.0.
	CpuAllocation *ResourcePoolResourceAllocationInfo
    // Resource allocation information for memory. This property was added in vSphere API 7.0.0.
	MemoryAllocation *ResourcePoolResourceAllocationInfo
}

// The ``FilterSpec`` class contains properties used to filter the results when listing resource pools (see ResourcePool#list). If multiple properties are specified, only resource pools matching all of the properties match the filter.
type ResourcePoolFilterSpec struct {
    // Identifiers of resource pools that can match the filter.
	ResourcePools map[string]bool
    // Names that resource pools must have to match the filter (see ResourcePoolInfo#name).
	Names map[string]bool
    // Resource pools that must contain the resource pool for the resource pool to match the filter.
	ParentResourcePools map[string]bool
    // Datacenters that must contain the resource pool for the resource pool to match the filter.
	Datacenters map[string]bool
    // Hosts that must contain the resource pool for the resource pool to match the filter.
	Hosts map[string]bool
    // Clusters that must contain the resource pool for the resource pool to match the filter.
	Clusters map[string]bool
}

// The ``Summary`` class contains commonly used information about a resource pool in vCenter Server.
type ResourcePoolSummary struct {
    // Identifier of the resource pool.
	ResourcePool string
    // Name of the resource pool.
	Name string
}

// The ``ResourceAllocationCreateSpec`` class contains resource allocation information used to create a resource pool, see ResourcePool#create. This class was added in vSphere API 7.0.0.
type ResourcePoolResourceAllocationCreateSpec struct {
    // Amount of resource that is guaranteed available to a resource pool. Reserved resources are not wasted if they are not used. If the utilization is less than the reservation, the resources can be utilized by other running virtual machines. Units are MB fo memory, and MHz for CPU. This property was added in vSphere API 7.0.0.
	Reservation *int64
    // In a resource pool with an expandable reservation, the reservation can grow beyond the specified value, if the parent resource pool has unreserved resources. A non-expandable reservation is called a fixed reservation. This property was added in vSphere API 7.0.0.
	ExpandableReservation *bool
    // The utilization of a resource pool will not exceed this limit, even if there are available resources. This is typically used to ensure a consistent performance of resource pools independent of available resources. If set to -1, then there is no fixed limit on resource usage (only bounded by available resources and shares). Units are MB for memory, and MHz for CPU. This property was added in vSphere API 7.0.0.
	Limit *int64
    // Shares are used in case of resource contention. This property was added in vSphere API 7.0.0.
	Shares *ResourcePoolSharesInfo
}

// The ResourcePool.CreateSpec class contains information used to create a resource pool, see ResourcePool#create. This class was added in vSphere API 7.0.0.
type ResourcePoolCreateSpec struct {
    // Name of the resource pool. This property was added in vSphere API 7.0.0.
	Name string
    // Parent of the created resource pool. This property was added in vSphere API 7.0.0.
	Parent string
    // Resource allocation for CPU. This property was added in vSphere API 7.0.0.
	CpuAllocation *ResourcePoolResourceAllocationCreateSpec
    // Resource allocation for memory. This property was added in vSphere API 7.0.0.
	MemoryAllocation *ResourcePoolResourceAllocationCreateSpec
}

// The ``ResourceAllocationUpdateSpec`` class descrives the updates to be made to the resource allocation settings of a resource pool. This class was added in vSphere API 7.0.0.
type ResourcePoolResourceAllocationUpdateSpec struct {
    // Amount of resource that is guaranteed available to a resource pool. Reserved resources are not wasted if they are not used. If the utilization is less than the reservation, the resources can be utilized by other running virtual machines. Units are MB fo memory, and MHz for CPU. This property was added in vSphere API 7.0.0.
	Reservation *int64
    // In a resource pool with an expandable reservation, the reservation can grow beyond the specified value, if the parent resource pool has unreserved resources. A non-expandable reservation is called a fixed reservation. This property was added in vSphere API 7.0.0.
	ExpandableReservation *bool
    // The utilization of a resource pool will not exceed this limit, even if there are available resources. This is typically used to ensure a consistent performance of resource pools independent of available resources. If set to -1, then there is no fixed limit on resource usage (only bounded by available resources and shares). Units are MB for memory, and MHz for CPU. This property was added in vSphere API 7.0.0.
	Limit *int64
    // Shares are used in case of resource contention. This property was added in vSphere API 7.0.0.
	Shares *ResourcePoolSharesInfo
}

// The ResourcePool.UpdateSpec class contains specification for updating the configuration of a resource pool. This class was added in vSphere API 7.0.0.
type ResourcePoolUpdateSpec struct {
    // Name of the resource pool. This property was added in vSphere API 7.0.0.
	Name *string
    // Resource allocation for CPU. This property was added in vSphere API 7.0.0.
	CpuAllocation *ResourcePoolResourceAllocationUpdateSpec
    // Resource allocation for memory. This property was added in vSphere API 7.0.0.
	MemoryAllocation *ResourcePoolResourceAllocationUpdateSpec
}



func resourcePoolGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["resource_pool"] = bindings.NewIdType([]string{"ResourcePool"}, "")
	fieldNameMap["resource_pool"] = "ResourcePool"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func resourcePoolGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(ResourcePoolInfoBindingType)
}

func resourcePoolGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["resource_pool"] = bindings.NewIdType([]string{"ResourcePool"}, "")
	fieldNameMap["resource_pool"] = "ResourcePool"
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
		map[string]int{"NotFound": 404,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}

func resourcePoolListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["filter"] = bindings.NewOptionalType(bindings.NewReferenceType(ResourcePoolFilterSpecBindingType))
	fieldNameMap["filter"] = "Filter"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func resourcePoolListOutputType() bindings.BindingType {
	return bindings.NewListType(bindings.NewReferenceType(ResourcePoolSummaryBindingType), reflect.TypeOf([]ResourcePoolSummary{}))
}

func resourcePoolListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["filter"] = bindings.NewOptionalType(bindings.NewReferenceType(ResourcePoolFilterSpecBindingType))
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
		map[string]int{"UnableToAllocateResource": 500,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}

func resourcePoolCreateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["spec"] = bindings.NewReferenceType(ResourcePoolCreateSpecBindingType)
	fieldNameMap["spec"] = "Spec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func resourcePoolCreateOutputType() bindings.BindingType {
	return bindings.NewIdType([]string{"ResourcePool"}, "")
}

func resourcePoolCreateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["spec"] = bindings.NewReferenceType(ResourcePoolCreateSpecBindingType)
	fieldNameMap["spec"] = "Spec"
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
		map[string]int{"Error": 500,"InvalidArgument": 400,"NotFound": 404,"ResourceInaccessible": 500,"ServiceUnavailable": 503,"UnableToAllocateResource": 500,"Unauthenticated": 401,"Unauthorized": 403})
}

func resourcePoolDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["resource_pool"] = bindings.NewIdType([]string{"ResourcePool"}, "")
	fieldNameMap["resource_pool"] = "ResourcePool"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func resourcePoolDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func resourcePoolDeleteRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["resource_pool"] = bindings.NewIdType([]string{"ResourcePool"}, "")
	fieldNameMap["resource_pool"] = "ResourcePool"
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
		map[string]int{"Error": 500,"NotFound": 404,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403,"Unsupported": 400})
}

func resourcePoolUpdateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["resource_pool"] = bindings.NewIdType([]string{"ResourcePool"}, "")
	fields["spec"] = bindings.NewReferenceType(ResourcePoolUpdateSpecBindingType)
	fieldNameMap["resource_pool"] = "ResourcePool"
	fieldNameMap["spec"] = "Spec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func resourcePoolUpdateOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func resourcePoolUpdateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["resource_pool"] = bindings.NewIdType([]string{"ResourcePool"}, "")
	fields["spec"] = bindings.NewReferenceType(ResourcePoolUpdateSpecBindingType)
	fieldNameMap["resource_pool"] = "ResourcePool"
	fieldNameMap["spec"] = "Spec"
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
		map[string]int{"Error": 500,"InvalidArgument": 400,"NotFound": 404,"ServiceUnavailable": 503,"UnableToAllocateResource": 500,"Unauthenticated": 401,"Unauthorized": 403})
}


func ResourcePoolSharesInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["level"] = bindings.NewEnumType("com.vmware.vcenter.resource_pool.shares_info.level", reflect.TypeOf(ResourcePoolSharesInfoLevel(ResourcePoolSharesInfoLevel_LOW)))
	fieldNameMap["level"] = "Level"
	fields["shares"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["shares"] = "Shares"
	var validators = []bindings.Validator{}
	uv1 := bindings.NewUnionValidator("level",
		map[string][]bindings.FieldData{
			"CUSTOM": []bindings.FieldData{
				bindings.NewFieldData("shares", true),
			},
			"LOW": []bindings.FieldData{},
			"NORMAL": []bindings.FieldData{},
			"HIGH": []bindings.FieldData{},
		},
	)
	validators = append(validators, uv1)
	return bindings.NewStructType("com.vmware.vcenter.resource_pool.shares_info", fields, reflect.TypeOf(ResourcePoolSharesInfo{}), fieldNameMap, validators)
}

func ResourcePoolResourceAllocationInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["reservation"] = bindings.NewIntegerType()
	fieldNameMap["reservation"] = "Reservation"
	fields["expandable_reservation"] = bindings.NewBooleanType()
	fieldNameMap["expandable_reservation"] = "ExpandableReservation"
	fields["limit"] = bindings.NewIntegerType()
	fieldNameMap["limit"] = "Limit"
	fields["shares"] = bindings.NewReferenceType(ResourcePoolSharesInfoBindingType)
	fieldNameMap["shares"] = "Shares"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.resource_pool.resource_allocation_info", fields, reflect.TypeOf(ResourcePoolResourceAllocationInfo{}), fieldNameMap, validators)
}

func ResourcePoolInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["name"] = bindings.NewStringType()
	fieldNameMap["name"] = "Name"
	fields["resource_pools"] = bindings.NewSetType(bindings.NewIdType([]string{"ResourcePool"}, ""), reflect.TypeOf(map[string]bool{}))
	fieldNameMap["resource_pools"] = "ResourcePools"
	fields["cpu_allocation"] = bindings.NewOptionalType(bindings.NewReferenceType(ResourcePoolResourceAllocationInfoBindingType))
	fieldNameMap["cpu_allocation"] = "CpuAllocation"
	fields["memory_allocation"] = bindings.NewOptionalType(bindings.NewReferenceType(ResourcePoolResourceAllocationInfoBindingType))
	fieldNameMap["memory_allocation"] = "MemoryAllocation"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.resource_pool.info", fields, reflect.TypeOf(ResourcePoolInfo{}), fieldNameMap, validators)
}

func ResourcePoolFilterSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["resource_pools"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string{"ResourcePool"}, ""), reflect.TypeOf(map[string]bool{})))
	fieldNameMap["resource_pools"] = "ResourcePools"
	fields["names"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewStringType(), reflect.TypeOf(map[string]bool{})))
	fieldNameMap["names"] = "Names"
	fields["parent_resource_pools"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string{"ResourcePool"}, ""), reflect.TypeOf(map[string]bool{})))
	fieldNameMap["parent_resource_pools"] = "ParentResourcePools"
	fields["datacenters"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string{"Datacenter"}, ""), reflect.TypeOf(map[string]bool{})))
	fieldNameMap["datacenters"] = "Datacenters"
	fields["hosts"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string{"HostSystem"}, ""), reflect.TypeOf(map[string]bool{})))
	fieldNameMap["hosts"] = "Hosts"
	fields["clusters"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string{"ClusterComputeResource"}, ""), reflect.TypeOf(map[string]bool{})))
	fieldNameMap["clusters"] = "Clusters"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.resource_pool.filter_spec", fields, reflect.TypeOf(ResourcePoolFilterSpec{}), fieldNameMap, validators)
}

func ResourcePoolSummaryBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["resource_pool"] = bindings.NewIdType([]string{"ResourcePool"}, "")
	fieldNameMap["resource_pool"] = "ResourcePool"
	fields["name"] = bindings.NewStringType()
	fieldNameMap["name"] = "Name"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.resource_pool.summary", fields, reflect.TypeOf(ResourcePoolSummary{}), fieldNameMap, validators)
}

func ResourcePoolResourceAllocationCreateSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["reservation"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["reservation"] = "Reservation"
	fields["expandable_reservation"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["expandable_reservation"] = "ExpandableReservation"
	fields["limit"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["limit"] = "Limit"
	fields["shares"] = bindings.NewOptionalType(bindings.NewReferenceType(ResourcePoolSharesInfoBindingType))
	fieldNameMap["shares"] = "Shares"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.resource_pool.resource_allocation_create_spec", fields, reflect.TypeOf(ResourcePoolResourceAllocationCreateSpec{}), fieldNameMap, validators)
}

func ResourcePoolCreateSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["name"] = bindings.NewStringType()
	fieldNameMap["name"] = "Name"
	fields["parent"] = bindings.NewIdType([]string{"ResourcePool"}, "")
	fieldNameMap["parent"] = "Parent"
	fields["cpu_allocation"] = bindings.NewOptionalType(bindings.NewReferenceType(ResourcePoolResourceAllocationCreateSpecBindingType))
	fieldNameMap["cpu_allocation"] = "CpuAllocation"
	fields["memory_allocation"] = bindings.NewOptionalType(bindings.NewReferenceType(ResourcePoolResourceAllocationCreateSpecBindingType))
	fieldNameMap["memory_allocation"] = "MemoryAllocation"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.resource_pool.create_spec", fields, reflect.TypeOf(ResourcePoolCreateSpec{}), fieldNameMap, validators)
}

func ResourcePoolResourceAllocationUpdateSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["reservation"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["reservation"] = "Reservation"
	fields["expandable_reservation"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["expandable_reservation"] = "ExpandableReservation"
	fields["limit"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["limit"] = "Limit"
	fields["shares"] = bindings.NewOptionalType(bindings.NewReferenceType(ResourcePoolSharesInfoBindingType))
	fieldNameMap["shares"] = "Shares"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.resource_pool.resource_allocation_update_spec", fields, reflect.TypeOf(ResourcePoolResourceAllocationUpdateSpec{}), fieldNameMap, validators)
}

func ResourcePoolUpdateSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["name"] = "Name"
	fields["cpu_allocation"] = bindings.NewOptionalType(bindings.NewReferenceType(ResourcePoolResourceAllocationUpdateSpecBindingType))
	fieldNameMap["cpu_allocation"] = "CpuAllocation"
	fields["memory_allocation"] = bindings.NewOptionalType(bindings.NewReferenceType(ResourcePoolResourceAllocationUpdateSpecBindingType))
	fieldNameMap["memory_allocation"] = "MemoryAllocation"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.resource_pool.update_spec", fields, reflect.TypeOf(ResourcePoolUpdateSpec{}), fieldNameMap, validators)
}

