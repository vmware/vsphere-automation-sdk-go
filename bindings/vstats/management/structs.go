/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Management.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package management

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)




// The ``TsdbSizeInfo`` class contains Time-series-database (TSDB) size information.
 type TsdbSizeInfo struct {
    Maximum int64
    Actual *int64
}






// The ``DbsConfig`` class specifies configuration of database connection.
 type DbsConfig struct {
    Name string
    Port int64
    Version string
}






// The ``DbsConfigs`` class specifies configuration information of the databases.
 type DbsConfigs struct {
    CoredbConfig DbsConfig
    SttdbConfig DbsConfig
    TsdbConfig DbsConfig
}






// The ``Info`` class contains information about the database configuration of the vstats service.
 type Info struct {
    DbsConfigs DbsConfigs
    TsDbSize TsdbSizeInfo
}









func getInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
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
      "GET",
      "/stats/management",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"Error": 500,"Unauthenticated": 401,"Unauthorized": 403})
}



func TsdbSizeInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["maximum"] = bindings.NewIntegerType()
    fieldNameMap["maximum"] = "Maximum"
    fields["actual"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["actual"] = "Actual"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vstats.management.tsdb_size_info",fields, reflect.TypeOf(TsdbSizeInfo{}), fieldNameMap, validators)
}

func DbsConfigBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    fields["port"] = bindings.NewIntegerType()
    fieldNameMap["port"] = "Port"
    fields["version"] = bindings.NewStringType()
    fieldNameMap["version"] = "Version"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vstats.management.dbs_config",fields, reflect.TypeOf(DbsConfig{}), fieldNameMap, validators)
}

func DbsConfigsBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["coredb_config"] = bindings.NewReferenceType(DbsConfigBindingType)
    fieldNameMap["coredb_config"] = "CoredbConfig"
    fields["sttdb_config"] = bindings.NewReferenceType(DbsConfigBindingType)
    fieldNameMap["sttdb_config"] = "SttdbConfig"
    fields["tsdb_config"] = bindings.NewReferenceType(DbsConfigBindingType)
    fieldNameMap["tsdb_config"] = "TsdbConfig"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vstats.management.dbs_configs",fields, reflect.TypeOf(DbsConfigs{}), fieldNameMap, validators)
}

func InfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["dbs_configs"] = bindings.NewReferenceType(DbsConfigsBindingType)
    fieldNameMap["dbs_configs"] = "DbsConfigs"
    fields["ts_db_size"] = bindings.NewReferenceType(TsdbSizeInfoBindingType)
    fieldNameMap["ts_db_size"] = "TsDbSize"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vstats.management.info",fields, reflect.TypeOf(Info{}), fieldNameMap, validators)
}


