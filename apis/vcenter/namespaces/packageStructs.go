/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for package: com.vmware.vcenter.namespaces.
 * Includes binding types of a top level structures and enumerations.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package namespaces

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
)



// A ``Container`` holds the data that describes a container within a pod. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type Container struct {
    Name string
    Image string
    Status string
    StartedAt string
    FinishedAt string
    Reason string
}






// The ``ResourceQuotaOptionsV1`` class represents the resource quota limits which can be applied on the namespace. Refer to ` <https://kubernetes.io/docs/concepts/policy/resource-quotas>`_ for information related to the properties of this object and what they map to. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type ResourceQuotaOptionsV1 struct {
    MemoryLimit *int64
    MemoryLimitDefault *int64
    MemoryRequestDefault *int64
    CpuLimit *int64
    CpuLimitDefault *int64
    CpuRequestDefault *int64
    StorageRequestLimit *int64
    PodCount *int64
    ServiceCount *int64
    DeploymentCount *int64
    DaemonSetCount *int64
    ReplicaSetCount *int64
    ReplicationControllerCount *int64
    StatefulSetCount *int64
    ConfigMapCount *int64
    SecretCount *int64
    PersistentVolumeClaimCount *int64
    JobCount *int64
}






// The ``ResourceQuotaOptionsV1Update`` class represents the changes to resource quota limits which are set on the namespace. Refer to ` <\a> Kubernetes Resource Quota <https://kubernetes.io/docs/concepts/policy/resource-quotas>`_ for information related to the properties of this object and what they map to. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type ResourceQuotaOptionsV1Update struct {
    MemoryLimit *int64
    MemoryLimitUnset *bool
    MemoryLimitDefault *int64
    MemoryLimitDefaultUnset *bool
    MemoryRequestDefault *int64
    MemoryRequestDefaultUnset *bool
    CpuLimit *int64
    CpuLimitUnset *bool
    CpuLimitDefault *int64
    CpuLimitDefaultUnset *bool
    CpuRequestDefault *int64
    CpuRequestDefaultUnset *bool
    StorageRequestLimit *int64
    StorageRequestLimitUnset *bool
    PodCount *int64
    PodCountUnset *bool
    ServiceCount *int64
    ServiceCountUnset *bool
    DeploymentCount *int64
    DeploymentCountUnset *bool
    DaemonSetCount *int64
    DaemonSetCountUnset *bool
    ReplicaSetCount *int64
    ReplicaSetCountUnset *bool
    ReplicationControllerCount *int64
    ReplicationControllerCountUnset *bool
    StatefulSetCount *int64
    StatefulSetCountUnset *bool
    ConfigMapCount *int64
    ConfigMapCountUnset *bool
    SecretCount *int64
    SecretCountUnset *bool
    PersistentVolumeClaimCount *int64
    PersistentVolumeClaimCountUnset *bool
    JobCount *int64
    JobCountUnset *bool
}










func ContainerBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    fields["image"] = bindings.NewStringType()
    fieldNameMap["image"] = "Image"
    fields["status"] = bindings.NewStringType()
    fieldNameMap["status"] = "Status"
    fields["started_at"] = bindings.NewStringType()
    fieldNameMap["started_at"] = "StartedAt"
    fields["finished_at"] = bindings.NewStringType()
    fieldNameMap["finished_at"] = "FinishedAt"
    fields["reason"] = bindings.NewStringType()
    fieldNameMap["reason"] = "Reason"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.namespaces.container",fields, reflect.TypeOf(Container{}), fieldNameMap, validators)
}

func ResourceQuotaOptionsV1BindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["memory_limit"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["memory_limit"] = "MemoryLimit"
    fields["memory_limit_default"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["memory_limit_default"] = "MemoryLimitDefault"
    fields["memory_request_default"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["memory_request_default"] = "MemoryRequestDefault"
    fields["cpu_limit"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["cpu_limit"] = "CpuLimit"
    fields["cpu_limit_default"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["cpu_limit_default"] = "CpuLimitDefault"
    fields["cpu_request_default"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["cpu_request_default"] = "CpuRequestDefault"
    fields["storage_request_limit"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["storage_request_limit"] = "StorageRequestLimit"
    fields["pod_count"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["pod_count"] = "PodCount"
    fields["service_count"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["service_count"] = "ServiceCount"
    fields["deployment_count"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["deployment_count"] = "DeploymentCount"
    fields["daemon_set_count"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["daemon_set_count"] = "DaemonSetCount"
    fields["replica_set_count"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["replica_set_count"] = "ReplicaSetCount"
    fields["replication_controller_count"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["replication_controller_count"] = "ReplicationControllerCount"
    fields["stateful_set_count"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["stateful_set_count"] = "StatefulSetCount"
    fields["config_map_count"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["config_map_count"] = "ConfigMapCount"
    fields["secret_count"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["secret_count"] = "SecretCount"
    fields["persistent_volume_claim_count"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["persistent_volume_claim_count"] = "PersistentVolumeClaimCount"
    fields["job_count"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["job_count"] = "JobCount"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.namespaces.resource_quota_options_v1",fields, reflect.TypeOf(ResourceQuotaOptionsV1{}), fieldNameMap, validators)
}

func ResourceQuotaOptionsV1UpdateBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["memory_limit"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["memory_limit"] = "MemoryLimit"
    fields["memory_limit_unset"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["memory_limit_unset"] = "MemoryLimitUnset"
    fields["memory_limit_default"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["memory_limit_default"] = "MemoryLimitDefault"
    fields["memory_limit_default_unset"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["memory_limit_default_unset"] = "MemoryLimitDefaultUnset"
    fields["memory_request_default"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["memory_request_default"] = "MemoryRequestDefault"
    fields["memory_request_default_unset"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["memory_request_default_unset"] = "MemoryRequestDefaultUnset"
    fields["cpu_limit"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["cpu_limit"] = "CpuLimit"
    fields["cpu_limit_unset"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["cpu_limit_unset"] = "CpuLimitUnset"
    fields["cpu_limit_default"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["cpu_limit_default"] = "CpuLimitDefault"
    fields["cpu_limit_default_unset"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["cpu_limit_default_unset"] = "CpuLimitDefaultUnset"
    fields["cpu_request_default"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["cpu_request_default"] = "CpuRequestDefault"
    fields["cpu_request_default_unset"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["cpu_request_default_unset"] = "CpuRequestDefaultUnset"
    fields["storage_request_limit"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["storage_request_limit"] = "StorageRequestLimit"
    fields["storage_request_limit_unset"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["storage_request_limit_unset"] = "StorageRequestLimitUnset"
    fields["pod_count"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["pod_count"] = "PodCount"
    fields["pod_count_unset"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["pod_count_unset"] = "PodCountUnset"
    fields["service_count"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["service_count"] = "ServiceCount"
    fields["service_count_unset"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["service_count_unset"] = "ServiceCountUnset"
    fields["deployment_count"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["deployment_count"] = "DeploymentCount"
    fields["deployment_count_unset"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["deployment_count_unset"] = "DeploymentCountUnset"
    fields["daemon_set_count"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["daemon_set_count"] = "DaemonSetCount"
    fields["daemon_set_count_unset"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["daemon_set_count_unset"] = "DaemonSetCountUnset"
    fields["replica_set_count"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["replica_set_count"] = "ReplicaSetCount"
    fields["replica_set_count_unset"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["replica_set_count_unset"] = "ReplicaSetCountUnset"
    fields["replication_controller_count"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["replication_controller_count"] = "ReplicationControllerCount"
    fields["replication_controller_count_unset"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["replication_controller_count_unset"] = "ReplicationControllerCountUnset"
    fields["stateful_set_count"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["stateful_set_count"] = "StatefulSetCount"
    fields["stateful_set_count_unset"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["stateful_set_count_unset"] = "StatefulSetCountUnset"
    fields["config_map_count"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["config_map_count"] = "ConfigMapCount"
    fields["config_map_count_unset"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["config_map_count_unset"] = "ConfigMapCountUnset"
    fields["secret_count"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["secret_count"] = "SecretCount"
    fields["secret_count_unset"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["secret_count_unset"] = "SecretCountUnset"
    fields["persistent_volume_claim_count"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["persistent_volume_claim_count"] = "PersistentVolumeClaimCount"
    fields["persistent_volume_claim_count_unset"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["persistent_volume_claim_count_unset"] = "PersistentVolumeClaimCountUnset"
    fields["job_count"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["job_count"] = "JobCount"
    fields["job_count_unset"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["job_count_unset"] = "JobCountUnset"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.namespaces.resource_quota_options_v1_update",fields, reflect.TypeOf(ResourceQuotaOptionsV1Update{}), fieldNameMap, validators)
}


