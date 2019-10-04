/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: AlgoParameter.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package algoParameter

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)




// The ``Param`` class contains EDRS algorithm parameters. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type Param struct {
    GenerateRecommendationEvent *bool
    MaxHosts *int64
    MinHosts *int64
    CpuScaleInThreshold *int64
    CpuScaleOutThreshold *int64
    CpuScaleInEwmaWeightPercent *int64
    CpuScaleOutEwmaWeightPercent *int64
    MemoryScaleInThreshold *int64
    MemoryScaleOutThreshold *int64
    MemoryScaleInEwmaWeightPercent *int64
    MemoryScaleOutEwmaWeightPercent *int64
    VsanScaleInThreshold *int64
    VsanScaleOutThreshold *int64
    VsanScaleInEwmaWeightPercent *int64
    VsanScaleOutEwmaWeightPercent *int64
    StorageScaleOutOnly *bool
}









func setInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    fields["params"] = bindings.NewReferenceType(ParamBindingType)
    fieldNameMap["cluster"] = "Cluster"
    fieldNameMap["params"] = "Params"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func setOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func setRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
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
       map[string]int{"Error": 500,"NotFound": 404,"Unauthorized": 403})
}


func getInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    fieldNameMap["cluster"] = "Cluster"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func getOutputType() bindings.BindingType {
    return bindings.NewReferenceType(ParamBindingType)
}

func getRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
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
       map[string]int{"NotFound": 404,"Unauthorized": 403})
}



func ParamBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["generate_recommendation_event"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["generate_recommendation_event"] = "GenerateRecommendationEvent"
    fields["max_hosts"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["max_hosts"] = "MaxHosts"
    fields["min_hosts"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["min_hosts"] = "MinHosts"
    fields["cpu_scale_in_threshold"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["cpu_scale_in_threshold"] = "CpuScaleInThreshold"
    fields["cpu_scale_out_threshold"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["cpu_scale_out_threshold"] = "CpuScaleOutThreshold"
    fields["cpu_scale_in_ewma_weight_percent"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["cpu_scale_in_ewma_weight_percent"] = "CpuScaleInEwmaWeightPercent"
    fields["cpu_scale_out_ewma_weight_percent"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["cpu_scale_out_ewma_weight_percent"] = "CpuScaleOutEwmaWeightPercent"
    fields["memory_scale_in_threshold"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["memory_scale_in_threshold"] = "MemoryScaleInThreshold"
    fields["memory_scale_out_threshold"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["memory_scale_out_threshold"] = "MemoryScaleOutThreshold"
    fields["memory_scale_in_ewma_weight_percent"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["memory_scale_in_ewma_weight_percent"] = "MemoryScaleInEwmaWeightPercent"
    fields["memory_scale_out_ewma_weight_percent"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["memory_scale_out_ewma_weight_percent"] = "MemoryScaleOutEwmaWeightPercent"
    fields["vsan_scale_in_threshold"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["vsan_scale_in_threshold"] = "VsanScaleInThreshold"
    fields["vsan_scale_out_threshold"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["vsan_scale_out_threshold"] = "VsanScaleOutThreshold"
    fields["vsan_scale_in_ewma_weight_percent"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["vsan_scale_in_ewma_weight_percent"] = "VsanScaleInEwmaWeightPercent"
    fields["vsan_scale_out_ewma_weight_percent"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["vsan_scale_out_ewma_weight_percent"] = "VsanScaleOutEwmaWeightPercent"
    fields["storage_scale_out_only"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["storage_scale_out_only"] = "StorageScaleOutOnly"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.cluster.edrs.algo_parameter.param",fields, reflect.TypeOf(Param{}), fieldNameMap, validators)
}


