/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Harbor.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package harbor

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/apis/vcenter/content/registries"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/apis/vcenter/content/registries/health"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
    "net/url"
    "time"
)




// The ``StorageSpec`` class contains the specification required to configure storage associated with a Harbor registry. In this version, Harbor registry is created in Kubernetes environment, information in this class will result in storage quotas on a Kubernetes namespace. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type StorageSpec struct {
    Policy string
    Limit *int64
}






// The ``StorageInfo`` class contains the detailed information about storage used by the Harbor registry. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type StorageInfo struct {
    Policy string
    Capacity int64
    Used int64
}






// The ``GarbageCollection`` class contains garbage collection configuration for the Harbor registry. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type GarbageCollection struct {
    Type_ registries.Recurrence
    DayOfWeek *registries.DayOfWeek
    Hour *int64
    Minute *int64
}






// The ``CreateSpec`` class contains the specification required to create a Harbor registry. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type CreateSpec struct {
    Cluster *string
    GarbageCollection *GarbageCollection
    Storage []StorageSpec
}






// The ``Summary`` class contains basic information about a running Harbor registry. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type Summary struct {
    Cluster *string
    Registry string
    Version string
    UiAccessUrl url.URL
}






// The ``Info`` class contains detailed information about a running Harbor registry. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type Info struct {
    Cluster *string
    Namespace *string
    Version string
    CreationTime time.Time
    UiAccessUrl url.URL
    CertChain []string
    GarbageCollection GarbageCollection
    Storage []StorageInfo
    Health health.Info
}









func createInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["client_token"] = bindings.NewOptionalType(bindings.NewStringType())
    fields["spec"] = bindings.NewReferenceType(CreateSpecBindingType)
    fieldNameMap["client_token"] = "ClientToken"
    fieldNameMap["spec"] = "Spec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func createOutputType() bindings.BindingType {
    return bindings.NewIdType([]string {"com.vmware.vcenter.content.Registry"}, "")
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
       map[string]int{"AlreadyExists": 400,"NotFound": 404,"Unsupported": 400,"InvalidArgument": 400,"Error": 500,"Unauthenticated": 401,"Unauthorized": 403})
}


func deleteInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["registry"] = bindings.NewIdType([]string {"com.vmware.vcenter.content.Registry"}, "")
    fieldNameMap["registry"] = "Registry"
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
       map[string]int{"NotFound": 404,"Error": 500,"Unauthenticated": 401,"Unauthorized": 403})
}


func getInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["registry"] = bindings.NewIdType([]string {"com.vmware.vcenter.content.Registry"}, "")
    fieldNameMap["registry"] = "Registry"
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
       map[string]int{"NotFound": 404,"Error": 500,"Unauthenticated": 401,"Unauthorized": 403})
}


func listInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
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
       map[string]int{"Error": 500,"Unauthenticated": 401,"Unauthorized": 403})
}



func StorageSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["policy"] = bindings.NewIdType([]string {"SpsStorageProfile"}, "")
    fieldNameMap["policy"] = "Policy"
    fields["limit"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["limit"] = "Limit"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.content.registries.harbor.storage_spec",fields, reflect.TypeOf(StorageSpec{}), fieldNameMap, validators)
}

func StorageInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["policy"] = bindings.NewIdType([]string {"SpsStorageProfile"}, "")
    fieldNameMap["policy"] = "Policy"
    fields["capacity"] = bindings.NewIntegerType()
    fieldNameMap["capacity"] = "Capacity"
    fields["used"] = bindings.NewIntegerType()
    fieldNameMap["used"] = "Used"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.content.registries.harbor.storage_info",fields, reflect.TypeOf(StorageInfo{}), fieldNameMap, validators)
}

func GarbageCollectionBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["type"] = bindings.NewEnumType("com.vmware.vcenter.content.registries.recurrence", reflect.TypeOf(registries.Recurrence(registries.Recurrence_NONE)))
    fieldNameMap["type"] = "Type_"
    fields["day_of_week"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.content.registries.day_of_week", reflect.TypeOf(registries.DayOfWeek(registries.DayOfWeek_SUNDAY))))
    fieldNameMap["day_of_week"] = "DayOfWeek"
    fields["hour"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["hour"] = "Hour"
    fields["minute"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["minute"] = "Minute"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("type",
        map[string][]bindings.FieldData {
            "WEEKLY": []bindings.FieldData {
                 bindings.NewFieldData("day_of_week", true),
                 bindings.NewFieldData("hour", true),
                 bindings.NewFieldData("minute", true),
            },
            "DAILY": []bindings.FieldData {
                 bindings.NewFieldData("hour", true),
                 bindings.NewFieldData("minute", true),
            },
            "NONE": []bindings.FieldData {},
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.vcenter.content.registries.harbor.garbage_collection",fields, reflect.TypeOf(GarbageCollection{}), fieldNameMap, validators)
}

func CreateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["cluster"] = bindings.NewOptionalType(bindings.NewIdType([]string {"ClusterComputeResource"}, ""))
    fieldNameMap["cluster"] = "Cluster"
    fields["garbage_collection"] = bindings.NewOptionalType(bindings.NewReferenceType(GarbageCollectionBindingType))
    fieldNameMap["garbage_collection"] = "GarbageCollection"
    fields["storage"] = bindings.NewListType(bindings.NewReferenceType(StorageSpecBindingType), reflect.TypeOf([]StorageSpec{}))
    fieldNameMap["storage"] = "Storage"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.content.registries.harbor.create_spec",fields, reflect.TypeOf(CreateSpec{}), fieldNameMap, validators)
}

func SummaryBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["cluster"] = bindings.NewOptionalType(bindings.NewIdType([]string {"ClusterComputeResource"}, ""))
    fieldNameMap["cluster"] = "Cluster"
    fields["registry"] = bindings.NewIdType([]string {"com.vmware.vcenter.content.Registry"}, "")
    fieldNameMap["registry"] = "Registry"
    fields["version"] = bindings.NewStringType()
    fieldNameMap["version"] = "Version"
    fields["ui_access_url"] = bindings.NewUriType()
    fieldNameMap["ui_access_url"] = "UiAccessUrl"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.content.registries.harbor.summary",fields, reflect.TypeOf(Summary{}), fieldNameMap, validators)
}

func InfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["cluster"] = bindings.NewOptionalType(bindings.NewIdType([]string {"ClusterComputeResource"}, ""))
    fieldNameMap["cluster"] = "Cluster"
    fields["namespace"] = bindings.NewOptionalType(bindings.NewIdType([]string {"NamespaceInstance"}, ""))
    fieldNameMap["namespace"] = "Namespace"
    fields["version"] = bindings.NewStringType()
    fieldNameMap["version"] = "Version"
    fields["creation_time"] = bindings.NewDateTimeType()
    fieldNameMap["creation_time"] = "CreationTime"
    fields["ui_access_url"] = bindings.NewUriType()
    fieldNameMap["ui_access_url"] = "UiAccessUrl"
    fields["cert_chain"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
    fieldNameMap["cert_chain"] = "CertChain"
    fields["garbage_collection"] = bindings.NewReferenceType(GarbageCollectionBindingType)
    fieldNameMap["garbage_collection"] = "GarbageCollection"
    fields["storage"] = bindings.NewListType(bindings.NewReferenceType(StorageInfoBindingType), reflect.TypeOf([]StorageInfo{}))
    fieldNameMap["storage"] = "Storage"
    fields["health"] = bindings.NewReferenceType(health.InfoBindingType)
    fieldNameMap["health"] = "Health"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.content.registries.harbor.info",fields, reflect.TypeOf(Info{}), fieldNameMap, validators)
}


