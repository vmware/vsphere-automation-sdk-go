/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: ResourcePool.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package resourcePool

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
//  Shares are used to determine relative allocation between resource consumers. In general, a consumer with more shares gets proportionally more of the resource, subject to certain other constraints.
 type SharesInfo struct {
    Level SharesInfo_Level
    Shares *int64
}




    
    // The ``Level`` enumeration class defines the possible values for the allocation level.
    //
    // <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
     
    type SharesInfo_Level string

    const (
        // For CPU: Shares = 500 \* number of virtual CPUs.
        //  For Memory: Shares = 5 \* virtual machine memory size in MB.
         SharesInfo_Level_LOW SharesInfo_Level = "LOW"
        // For CPU: Shares = 1000 \* number of virtual CPUs.
        //  For Memory: Shares = 10 \* virtual machine memory size in MB.
         SharesInfo_Level_NORMAL SharesInfo_Level = "NORMAL"
        // For CPU: Shares = 2000 \* nmumber of virtual CPUs.
        //  For Memory: Shares = 20 \* virtual machine memory size in MB.
         SharesInfo_Level_HIGH SharesInfo_Level = "HIGH"
        // If map with bool value, in case there is resource contention the server uses the shares value to determine the resource allocation.
         SharesInfo_Level_CUSTOM SharesInfo_Level = "CUSTOM"
    )

    func (l SharesInfo_Level) SharesInfo_Level() bool {
        switch l {
            case SharesInfo_Level_LOW:
                return true
            case SharesInfo_Level_NORMAL:
                return true
            case SharesInfo_Level_HIGH:
                return true
            case SharesInfo_Level_CUSTOM:
                return true
            default:
                return false
        }
    }



// The ``ResourceAllocationInfo`` class contains resource allocation information of a resource pool.
 type ResourceAllocationInfo struct {
    Reservation int64
    ExpandableReservation bool
    Limit int64
    Shares SharesInfo
}






// The ``Info`` class contains information about a resource pool.
 type Info struct {
    Name string
    ResourcePools map[string]bool
    CpuAllocation *ResourceAllocationInfo
    MemoryAllocation *ResourceAllocationInfo
}






// The ``FilterSpec`` class contains properties used to filter the results when listing resource pools (see ResourcePool#list). If multiple properties are specified, only resource pools matching all of the properties match the filter.
 type FilterSpec struct {
    ResourcePools map[string]bool
    Names map[string]bool
    ParentResourcePools map[string]bool
    Datacenters map[string]bool
    Hosts map[string]bool
    Clusters map[string]bool
}






// The ``Summary`` class contains commonly used information about a resource pool in vCenter Server.
 type Summary struct {
    ResourcePool string
    Name string
}






// The ``ResourceAllocationCreateSpec`` class contains resource allocation information used to create a resource pool, see ResourcePool#create.
 type ResourceAllocationCreateSpec struct {
    Reservation *int64
    ExpandableReservation *bool
    Limit *int64
    Shares *SharesInfo
}






// The ResourcePool.CreateSpec class contains information used to create a resource pool, see ResourcePool#create.
 type CreateSpec struct {
    Name string
    Parent string
    CpuAllocation *ResourceAllocationCreateSpec
    MemoryAllocation *ResourceAllocationCreateSpec
}






// The ``ResourceAllocationUpdateSpec`` class descrives the updates to be made to the resource allocation settings of a resource pool.
 type ResourceAllocationUpdateSpec struct {
    Reservation *int64
    ExpandableReservation *bool
    Limit *int64
    Shares *SharesInfo
}






// The ResourcePool.UpdateSpec class contains specification for updating the configuration of a resource pool.
 type UpdateSpec struct {
    Name *string
    CpuAllocation *ResourceAllocationUpdateSpec
    MemoryAllocation *ResourceAllocationUpdateSpec
}









func getInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["resource_pool"] = bindings.NewIdType([]string {"ResourcePool"}, "")
    fieldNameMap["resource_pool"] = "ResourcePool"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func getOutputType() bindings.BindingType {
    return bindings.NewReferenceType(InfoBindingType)
}

func getRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "null",
      "",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"NotFound": 404,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}


func listInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["filter"] = bindings.NewOptionalType(bindings.NewReferenceType(FilterSpecBindingType))
    fieldNameMap["filter"] = "Filter"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func listOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewReferenceType(SummaryBindingType), reflect.TypeOf([]Summary{}))
}

func listRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "null",
      "",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"UnableToAllocateResource": 500,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}


func createInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["spec"] = bindings.NewReferenceType(CreateSpecBindingType)
    fieldNameMap["spec"] = "Spec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func createOutputType() bindings.BindingType {
    return bindings.NewIdType([]string {"ResourcePool"}, "")
}

func createRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "null",
      "",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"Error": 500,"InvalidArgument": 400,"NotFound": 404,"ResourceInaccessible": 500,"ServiceUnavailable": 503,"UnableToAllocateResource": 500,"Unauthenticated": 401,"Unauthorized": 403})
}


func deleteInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["resource_pool"] = bindings.NewIdType([]string {"ResourcePool"}, "")
    fieldNameMap["resource_pool"] = "ResourcePool"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func deleteOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func deleteRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "null",
      "",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"Error": 500,"NotFound": 404,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403,"Unsupported": 400})
}


func updateInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["resource_pool"] = bindings.NewIdType([]string {"ResourcePool"}, "")
    fields["spec"] = bindings.NewReferenceType(UpdateSpecBindingType)
    fieldNameMap["resource_pool"] = "ResourcePool"
    fieldNameMap["spec"] = "Spec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func updateOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func updateRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "null",
      "",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"Error": 500,"InvalidArgument": 400,"NotFound": 404,"ServiceUnavailable": 503,"UnableToAllocateResource": 500,"Unauthenticated": 401,"Unauthorized": 403})
}



func SharesInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["level"] = bindings.NewEnumType("com.vmware.vcenter.resource_pool.shares_info.level", reflect.TypeOf(SharesInfo_Level(SharesInfo_Level_LOW)))
    fieldNameMap["level"] = "Level"
    fields["shares"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["shares"] = "Shares"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("level",
        map[string][]bindings.FieldData {
            "CUSTOM": []bindings.FieldData {
                 bindings.NewFieldData("shares", true),
            },
            "LOW": []bindings.FieldData {},
            "NORMAL": []bindings.FieldData {},
            "HIGH": []bindings.FieldData {},
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.vcenter.resource_pool.shares_info",fields, reflect.TypeOf(SharesInfo{}), fieldNameMap, validators)
}

func ResourceAllocationInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["reservation"] = bindings.NewIntegerType()
    fieldNameMap["reservation"] = "Reservation"
    fields["expandable_reservation"] = bindings.NewBooleanType()
    fieldNameMap["expandable_reservation"] = "ExpandableReservation"
    fields["limit"] = bindings.NewIntegerType()
    fieldNameMap["limit"] = "Limit"
    fields["shares"] = bindings.NewReferenceType(SharesInfoBindingType)
    fieldNameMap["shares"] = "Shares"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.resource_pool.resource_allocation_info",fields, reflect.TypeOf(ResourceAllocationInfo{}), fieldNameMap, validators)
}

func InfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    fields["resource_pools"] = bindings.NewSetType(bindings.NewIdType([]string {"ResourcePool"}, ""), reflect.TypeOf(map[string]bool{}))
    fieldNameMap["resource_pools"] = "ResourcePools"
    fields["cpu_allocation"] = bindings.NewOptionalType(bindings.NewReferenceType(ResourceAllocationInfoBindingType))
    fieldNameMap["cpu_allocation"] = "CpuAllocation"
    fields["memory_allocation"] = bindings.NewOptionalType(bindings.NewReferenceType(ResourceAllocationInfoBindingType))
    fieldNameMap["memory_allocation"] = "MemoryAllocation"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.resource_pool.info",fields, reflect.TypeOf(Info{}), fieldNameMap, validators)
}

func FilterSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["resource_pools"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string {"ResourcePool"}, ""), reflect.TypeOf(map[string]bool{})))
    fieldNameMap["resource_pools"] = "ResourcePools"
    fields["names"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewStringType(), reflect.TypeOf(map[string]bool{})))
    fieldNameMap["names"] = "Names"
    fields["parent_resource_pools"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string {"ResourcePool"}, ""), reflect.TypeOf(map[string]bool{})))
    fieldNameMap["parent_resource_pools"] = "ParentResourcePools"
    fields["datacenters"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string {"Datacenter"}, ""), reflect.TypeOf(map[string]bool{})))
    fieldNameMap["datacenters"] = "Datacenters"
    fields["hosts"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string {"HostSystem"}, ""), reflect.TypeOf(map[string]bool{})))
    fieldNameMap["hosts"] = "Hosts"
    fields["clusters"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string {"ClusterComputeResource"}, ""), reflect.TypeOf(map[string]bool{})))
    fieldNameMap["clusters"] = "Clusters"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.resource_pool.filter_spec",fields, reflect.TypeOf(FilterSpec{}), fieldNameMap, validators)
}

func SummaryBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["resource_pool"] = bindings.NewIdType([]string {"ResourcePool"}, "")
    fieldNameMap["resource_pool"] = "ResourcePool"
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.resource_pool.summary",fields, reflect.TypeOf(Summary{}), fieldNameMap, validators)
}

func ResourceAllocationCreateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["reservation"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["reservation"] = "Reservation"
    fields["expandable_reservation"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["expandable_reservation"] = "ExpandableReservation"
    fields["limit"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["limit"] = "Limit"
    fields["shares"] = bindings.NewOptionalType(bindings.NewReferenceType(SharesInfoBindingType))
    fieldNameMap["shares"] = "Shares"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.resource_pool.resource_allocation_create_spec",fields, reflect.TypeOf(ResourceAllocationCreateSpec{}), fieldNameMap, validators)
}

func CreateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    fields["parent"] = bindings.NewIdType([]string {"ResourcePool"}, "")
    fieldNameMap["parent"] = "Parent"
    fields["cpu_allocation"] = bindings.NewOptionalType(bindings.NewReferenceType(ResourceAllocationCreateSpecBindingType))
    fieldNameMap["cpu_allocation"] = "CpuAllocation"
    fields["memory_allocation"] = bindings.NewOptionalType(bindings.NewReferenceType(ResourceAllocationCreateSpecBindingType))
    fieldNameMap["memory_allocation"] = "MemoryAllocation"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.resource_pool.create_spec",fields, reflect.TypeOf(CreateSpec{}), fieldNameMap, validators)
}

func ResourceAllocationUpdateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["reservation"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["reservation"] = "Reservation"
    fields["expandable_reservation"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["expandable_reservation"] = "ExpandableReservation"
    fields["limit"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["limit"] = "Limit"
    fields["shares"] = bindings.NewOptionalType(bindings.NewReferenceType(SharesInfoBindingType))
    fieldNameMap["shares"] = "Shares"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.resource_pool.resource_allocation_update_spec",fields, reflect.TypeOf(ResourceAllocationUpdateSpec{}), fieldNameMap, validators)
}

func UpdateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["name"] = "Name"
    fields["cpu_allocation"] = bindings.NewOptionalType(bindings.NewReferenceType(ResourceAllocationUpdateSpecBindingType))
    fieldNameMap["cpu_allocation"] = "CpuAllocation"
    fields["memory_allocation"] = bindings.NewOptionalType(bindings.NewReferenceType(ResourceAllocationUpdateSpecBindingType))
    fieldNameMap["memory_allocation"] = "MemoryAllocation"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.resource_pool.update_spec",fields, reflect.TypeOf(UpdateSpec{}), fieldNameMap, validators)
}


