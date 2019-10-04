/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Recommendation.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package recommendation

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/vapi/std"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)




// The ``UtilizationInfo`` class contains the details of cluster resources utilization and corresponding thresholds. Resource utilization is an estimated value derived from current usage and historical data. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type UtilizationInfo struct {
    CpuScaleInThreshold int64
    CpuScaleInUtilization int64
    CpuScaleOutThreshold int64
    CpuScaleOutUtilization int64
    MemoryScaleInThreshold int64
    MemoryScaleInUtilization int64
    MemoryScaleOutThreshold int64
    MemoryScaleOutUtilization int64
    VsanScaleInThreshold *int64
    VsanScaleInUtilization *int64
    VsanScaleOutThreshold *int64
    VsanScaleOutUtilization *int64
    VsanUtilization *int64
}






// The ``Recommendation`` class contains the recommend action, and all related details. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type Recommendation struct {
    IsEdrsEnabled bool
    Action Recommendation_Action
    HostsToRemove map[string]bool
    NoActionReasons []std.LocalizableMessage
    Hosts map[string]bool
    Utilization UtilizationInfo
}




    
    // The ``Action`` enumeration class defines the recommend action given by EDRS. **Warning:** This enumeration is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
     
    type Recommendation_Action string

    const (
        // Recomend scale out - to add host into the cluster. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
         Recommendation_Action_SCALE_OUT Recommendation_Action = "SCALE_OUT"
        // Recommend scale in - to remove hosts from the cluster. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
         Recommendation_Action_SCALE_IN Recommendation_Action = "SCALE_IN"
        // No action is needed. It could because either cluster resources are well utilized, or scale in recommendation is supressed since no host can be evacuated. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
         Recommendation_Action_NO_ACTION Recommendation_Action = "NO_ACTION"
    )

    func (a Recommendation_Action) Recommendation_Action() bool {
        switch a {
            case Recommendation_Action_SCALE_OUT:
                return true
            case Recommendation_Action_SCALE_IN:
                return true
            case Recommendation_Action_NO_ACTION:
                return true
            default:
                return false
        }
    }






func generateInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    fieldNameMap["cluster"] = "Cluster"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func generateOutputType() bindings.BindingType {
    return bindings.NewReferenceType(RecommendationBindingType)
}

func generateRestMetadata() protocol.OperationRestMetadata {
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



func UtilizationInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["cpu_scale_in_threshold"] = bindings.NewIntegerType()
    fieldNameMap["cpu_scale_in_threshold"] = "CpuScaleInThreshold"
    fields["cpu_scale_in_utilization"] = bindings.NewIntegerType()
    fieldNameMap["cpu_scale_in_utilization"] = "CpuScaleInUtilization"
    fields["cpu_scale_out_threshold"] = bindings.NewIntegerType()
    fieldNameMap["cpu_scale_out_threshold"] = "CpuScaleOutThreshold"
    fields["cpu_scale_out_utilization"] = bindings.NewIntegerType()
    fieldNameMap["cpu_scale_out_utilization"] = "CpuScaleOutUtilization"
    fields["memory_scale_in_threshold"] = bindings.NewIntegerType()
    fieldNameMap["memory_scale_in_threshold"] = "MemoryScaleInThreshold"
    fields["memory_scale_in_utilization"] = bindings.NewIntegerType()
    fieldNameMap["memory_scale_in_utilization"] = "MemoryScaleInUtilization"
    fields["memory_scale_out_threshold"] = bindings.NewIntegerType()
    fieldNameMap["memory_scale_out_threshold"] = "MemoryScaleOutThreshold"
    fields["memory_scale_out_utilization"] = bindings.NewIntegerType()
    fieldNameMap["memory_scale_out_utilization"] = "MemoryScaleOutUtilization"
    fields["vsan_scale_in_threshold"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["vsan_scale_in_threshold"] = "VsanScaleInThreshold"
    fields["vsan_scale_in_utilization"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["vsan_scale_in_utilization"] = "VsanScaleInUtilization"
    fields["vsan_scale_out_threshold"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["vsan_scale_out_threshold"] = "VsanScaleOutThreshold"
    fields["vsan_scale_out_utilization"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["vsan_scale_out_utilization"] = "VsanScaleOutUtilization"
    fields["vsan_utilization"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["vsan_utilization"] = "VsanUtilization"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.cluster.edrs.recommendation.utilization_info",fields, reflect.TypeOf(UtilizationInfo{}), fieldNameMap, validators)
}

func RecommendationBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["is_edrs_enabled"] = bindings.NewBooleanType()
    fieldNameMap["is_edrs_enabled"] = "IsEdrsEnabled"
    fields["action"] = bindings.NewEnumType("com.vmware.vcenter.cluster.edrs.recommendation.recommendation.action", reflect.TypeOf(Recommendation_Action(Recommendation_Action_SCALE_OUT)))
    fieldNameMap["action"] = "Action"
    fields["hosts_to_remove"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string {"HostSystem"}, ""), reflect.TypeOf(map[string]bool{})))
    fieldNameMap["hosts_to_remove"] = "HostsToRemove"
    fields["no_action_reasons"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(std.LocalizableMessageBindingType), reflect.TypeOf([]std.LocalizableMessage{})))
    fieldNameMap["no_action_reasons"] = "NoActionReasons"
    fields["hosts"] = bindings.NewSetType(bindings.NewIdType([]string {"HostSystem"}, ""), reflect.TypeOf(map[string]bool{}))
    fieldNameMap["hosts"] = "Hosts"
    fields["utilization"] = bindings.NewReferenceType(UtilizationInfoBindingType)
    fieldNameMap["utilization"] = "Utilization"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("action",
        map[string][]bindings.FieldData {
            "SCALE_IN": []bindings.FieldData {
                 bindings.NewFieldData("hosts_to_remove", true),
            },
            "NO_ACTION": []bindings.FieldData {
                 bindings.NewFieldData("no_action_reasons", true),
            },
            "SCALE_OUT": []bindings.FieldData {},
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.vcenter.cluster.edrs.recommendation.recommendation",fields, reflect.TypeOf(Recommendation{}), fieldNameMap, validators)
}


