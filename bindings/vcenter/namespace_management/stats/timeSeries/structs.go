/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: TimeSeries.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package timeSeries

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)




// A set of timestamps and statistical values representing a time series. The lengths of TimeSeries#timeStamps and TimeSeries#values will always match each other. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type TimeSeries struct {
    Counter string
    TimeStamps []int64
    Values []int64
}






// Pod identifier. These are the fields required to uniquely identify a pod. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type PodIdentifier struct {
    Namespace string
    PodName string
}






// This structure is sent in a request for TimeSeries data and is used to specify what object stats should be returned for. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type Spec struct {
    ObjType Spec_ObjType
    Pod *PodIdentifier
    Namespace *string
    Cluster *string
    Start int64
    End int64
}




    
    // Type of statistics object that this request is operating on. **Warning:** This enumeration is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
     
    type Spec_ObjType string

    const (
        // The CLUSTER object type is used when specifying a vSphere cluster. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
         Spec_ObjType_CLUSTER Spec_ObjType = "CLUSTER"
        // The NAMESPACE object type is used to specify a namespace. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
         Spec_ObjType_NAMESPACE Spec_ObjType = "NAMESPACE"
        // The POD object type is used to specify an individual pod within a namespace. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
         Spec_ObjType_POD Spec_ObjType = "POD"
    )

    func (o Spec_ObjType) Spec_ObjType() bool {
        switch o {
            case Spec_ObjType_CLUSTER:
                return true
            case Spec_ObjType_NAMESPACE:
                return true
            case Spec_ObjType_POD:
                return true
            default:
                return false
        }
    }






func getInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["spec"] = bindings.NewReferenceType(SpecBindingType)
    fieldNameMap["spec"] = "Spec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func getOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewReferenceType(TimeSeriesBindingType), reflect.TypeOf([]TimeSeries{}))
}

func getRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["spec.cluster"] = bindings.NewOptionalType(bindings.NewIdType([]string {"ClusterComputeResource"}, ""))
    paramsTypeMap["spec.obj_type"] = bindings.NewEnumType("com.vmware.vcenter.namespace_management.stats.time_series.spec.obj_type", reflect.TypeOf(Spec_ObjType(Spec_ObjType_CLUSTER)))
    paramsTypeMap["spec.start"] = bindings.NewIntegerType()
    paramsTypeMap["spec.end"] = bindings.NewIntegerType()
    paramsTypeMap["spec.pod.namespace"] = bindings.NewIdType([]string {"com.vmware.vcenter.namespaces.Instance"}, "")
    paramsTypeMap["spec.namespace"] = bindings.NewOptionalType(bindings.NewIdType([]string {"com.vmware.vcenter.namespaces.Instance"}, ""))
    paramsTypeMap["spec.pod.pod_name"] = bindings.NewStringType()
    queryParams["spec.cluster"] = "cluster"
    queryParams["spec.pod.namespace"] = "pod-namespace"
    queryParams["spec.namespace"] = "namespace"
    queryParams["spec.start"] = "start"
    queryParams["spec.obj_type"] = "obj_type"
    queryParams["spec.end"] = "end"
    queryParams["spec.pod.pod_name"] = "pod_name"
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "GET",
      "/vcenter/namespace-management/stats/time-series",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"Error": 500,"InvalidArgument": 400,"NotFound": 404,"Unsupported": 400,"Unauthenticated": 401,"Unauthorized": 403})
}



func TimeSeriesBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["counter"] = bindings.NewStringType()
    fieldNameMap["counter"] = "Counter"
    fields["time_stamps"] = bindings.NewListType(bindings.NewIntegerType(), reflect.TypeOf([]int64{}))
    fieldNameMap["time_stamps"] = "TimeStamps"
    fields["values"] = bindings.NewListType(bindings.NewIntegerType(), reflect.TypeOf([]int64{}))
    fieldNameMap["values"] = "Values"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.namespace_management.stats.time_series.time_series",fields, reflect.TypeOf(TimeSeries{}), fieldNameMap, validators)
}

func PodIdentifierBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["namespace"] = bindings.NewIdType([]string {"com.vmware.vcenter.namespaces.Instance"}, "")
    fieldNameMap["namespace"] = "Namespace"
    fields["pod_name"] = bindings.NewStringType()
    fieldNameMap["pod_name"] = "PodName"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.namespace_management.stats.time_series.pod_identifier",fields, reflect.TypeOf(PodIdentifier{}), fieldNameMap, validators)
}

func SpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["obj_type"] = bindings.NewEnumType("com.vmware.vcenter.namespace_management.stats.time_series.spec.obj_type", reflect.TypeOf(Spec_ObjType(Spec_ObjType_CLUSTER)))
    fieldNameMap["obj_type"] = "ObjType"
    fields["pod"] = bindings.NewOptionalType(bindings.NewReferenceType(PodIdentifierBindingType))
    fieldNameMap["pod"] = "Pod"
    fields["namespace"] = bindings.NewOptionalType(bindings.NewIdType([]string {"com.vmware.vcenter.namespaces.Instance"}, ""))
    fieldNameMap["namespace"] = "Namespace"
    fields["cluster"] = bindings.NewOptionalType(bindings.NewIdType([]string {"ClusterComputeResource"}, ""))
    fieldNameMap["cluster"] = "Cluster"
    fields["start"] = bindings.NewIntegerType()
    fieldNameMap["start"] = "Start"
    fields["end"] = bindings.NewIntegerType()
    fieldNameMap["end"] = "End"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("obj_type",
        map[string][]bindings.FieldData {
            "POD": []bindings.FieldData {
                 bindings.NewFieldData("pod", true),
            },
            "NAMESPACE": []bindings.FieldData {
                 bindings.NewFieldData("namespace", true),
            },
            "CLUSTER": []bindings.FieldData {
                 bindings.NewFieldData("cluster", true),
            },
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.vcenter.namespace_management.stats.time_series.spec",fields, reflect.TypeOf(Spec{}), fieldNameMap, validators)
}


