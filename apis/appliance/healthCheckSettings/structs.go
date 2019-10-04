/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: HealthCheckSettings.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package healthCheckSettings

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)




// ``SettingSpec`` class contains specification of vCenter Server health and its corresponding state. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type SettingSpec struct {
    DbHealthCheckStateManualBackup bool
    DbHealthCheckStateScheduledBackup bool
}






 type UpdateSpec struct {
    DbHealthCheckStateManualBackup *bool
    DbHealthCheckStateScheduledBackup *bool
}









func getInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func getOutputType() bindings.BindingType {
    return bindings.NewReferenceType(SettingSpecBindingType)
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
       map[string]int{"Error": 500})
}


func updateInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["health_settings"] = bindings.NewReferenceType(UpdateSpecBindingType)
    fieldNameMap["health_settings"] = "HealthSettings"
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
       map[string]int{"Error": 500,"InvalidArgument": 400})
}



func SettingSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["db_health_check_state_manual_backup"] = bindings.NewBooleanType()
    fieldNameMap["db_health_check_state_manual_backup"] = "DbHealthCheckStateManualBackup"
    fields["db_health_check_state_scheduled_backup"] = bindings.NewBooleanType()
    fieldNameMap["db_health_check_state_scheduled_backup"] = "DbHealthCheckStateScheduledBackup"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.health_check_settings.setting_spec",fields, reflect.TypeOf(SettingSpec{}), fieldNameMap, validators)
}

func UpdateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["db_health_check_state_manual_backup"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["db_health_check_state_manual_backup"] = "DbHealthCheckStateManualBackup"
    fields["db_health_check_state_scheduled_backup"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["db_health_check_state_scheduled_backup"] = "DbHealthCheckStateScheduledBackup"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.health_check_settings.update_spec",fields, reflect.TypeOf(UpdateSpec{}), fieldNameMap, validators)
}


