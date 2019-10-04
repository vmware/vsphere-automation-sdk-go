/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Snmp.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package snmp

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)



// ``SNMPAuthProto`` enumeration class Defines SNMP authentication protocols
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type SNMPAuthProto string

const (
    // NONE
     SNMPAuthProto_none SNMPAuthProto = "none"
    // SHA1
     SNMPAuthProto_SHA1 SNMPAuthProto = "SHA1"
    // MD5
     SNMPAuthProto_MD5 SNMPAuthProto = "MD5"
)

func (s SNMPAuthProto) SNMPAuthProto() bool {
    switch s {
        case SNMPAuthProto_none:
            return true
        case SNMPAuthProto_SHA1:
            return true
        case SNMPAuthProto_MD5:
            return true
        default:
            return false
    }
}




// ``SNMPPrivProto`` enumeration class Defines SNMP privacy protocols
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type SNMPPrivProto string

const (
    // AES128
     SNMPPrivProto_AES128 SNMPPrivProto = "AES128"
    // NONE
     SNMPPrivProto_none SNMPPrivProto = "none"
)

func (s SNMPPrivProto) SNMPPrivProto() bool {
    switch s {
        case SNMPPrivProto_AES128:
            return true
        case SNMPPrivProto_none:
            return true
        default:
            return false
    }
}




// ``SNMPSecLevel`` enumeration class Defines SNMP decurity levels
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type SNMPSecLevel string

const (
    // none
     SNMPSecLevel_none SNMPSecLevel = "none"
    // auth
     SNMPSecLevel_auth SNMPSecLevel = "auth"
    // priv
     SNMPSecLevel_priv SNMPSecLevel = "priv"
)

func (s SNMPSecLevel) SNMPSecLevel() bool {
    switch s {
        case SNMPSecLevel_none:
            return true
        case SNMPSecLevel_auth:
            return true
        case SNMPSecLevel_priv:
            return true
        default:
            return false
    }
}




// ``SNMPv3Notfication`` enumeration class Defines SNMP v3 notification types
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type SNMPv3Notfication string

const (
    // inform
     SNMPv3Notfication_inform SNMPv3Notfication = "inform"
    // trap
     SNMPv3Notfication_trap SNMPv3Notfication = "trap"
)

func (s SNMPv3Notfication) SNMPv3Notfication() bool {
    switch s {
        case SNMPv3Notfication_inform:
            return true
        case SNMPv3Notfication_trap:
            return true
        default:
            return false
    }
}





// ``SNMPLimits`` class Structure that provides various limits of the SNMP agent.
 type SNMPLimits struct {
    MaxCommunities int64
    MaxTrapDestinationsV1 int64
    MaxDestinationsV3 int64
    MaxNotificationFilters int64
    MaxCommunityLength int64
    MaxBufferSize int64
}






// ``SNMPTestResults`` class Structure to provide operators diagnostics test results.
 type SNMPTestResults struct {
    Success bool
    Message string
}






// ``SNMPStats`` class Structure to provide operators diagnostics on snmp agent itself.
 type SNMPStats struct {
    Sysuptime string
    Worstrtimelast string
    Avgresponsetime string
    Worstresponsetime string
    Inpkts int64
    Outpkts int64
    Usmstatsnotintimewindows int64
    Usmstatsunknownusernames int64
    Usmstatsunknownengineids int64
    Usmstatswrongdigests int64
    Usmstatsdecryptionerrors int64
    Inbadversions int64
    Inbadcommunitynames int64
    Inbadcommunityuses int64
    Inasnparseerrs int64
    Intoobigs int64
    Innosuchnames int64
    Inbadvalues int64
    Ingenerrs int64
    Outtoobigs int64
    Outnosuchnames int64
    Outbadvalues int64
    Outgenerrs int64
    Outtraps int64
    Silentdrops int64
    Avgvarbinds int64
    Maxvarbinds int64
}






// ``SNMPConfig`` class Structure that defines the SNMP configuration, provided as input to set(), and never the result of get(). See SNMPConfigReadOnly. This structure is used to configure SNMP v1, v2c, and v3.
 type SNMPConfig struct {
    Authentication SNMPAuthProto
    Communities []string
    Engineid string
    Loglevel string
    Notraps []string
    Port int64
    Privacy SNMPPrivProto
    Remoteusers []string
    Syscontact string
    Syslocation string
    Targets []string
    Users []string
    V3targets []string
}






// ``SNMPUser`` class Structure that defines information associated with an SNMP user. authKey and privKey are localized keys defined in http://tools.ietf.org/html/rfc3826#section-1.2.
 type SNMPUser struct {
    Username string
    SecLevel SNMPSecLevel
    AuthKey string
    PrivKey string
}






// ``SNMPv3Target`` class Structure that defines an SNMP v3 inform or trap target.
 type SNMPv3Target struct {
    Type_ SNMPv3Notfication
    SecLevel SNMPSecLevel
    Ip string
    Port int64
    User string
}






// ``SNMPv1TrapTarget`` class Structure that defines an SNMP v1/v2c trap target.
 type SNMPv1TrapTarget struct {
    Ip string
    Port int64
    Community string
}






// ``SNMPRemoteUser`` class Structure that defines a user at particular remote SNMPv3 entity needed when using informs. auth_key and priv_key contained localized keys as defined in http://tools.ietf.org/html/rfc3826#section-1.2.
 type SNMPRemoteUser struct {
    Username string
    SecLevel SNMPSecLevel
    Authentication SNMPAuthProto
    AuthKey string
    Privacy SNMPPrivProto
    PrivKey string
    Engineid string
}






// ``SNMPConfigReadOnly`` class Structure that defines the SNMP configuration, the result of get(), and never provided as input to set(). This structure differs from SNMPConfig because it contains localized keys (as defined in http://tools.ietf.org/html/rfc3826#section-1.2), instead of raw secret strings. This structure can be used to configure SNMP v1, v2c, and v3. Keep this structure in sync with vmw_snmp.py:_default_config(). Note that if a field if left empty, it is considered unset and will be ignored. Existing array elements below can be unset by sending an element with the string 'reset'.
 type SNMPConfigReadOnly struct {
    Authentication SNMPAuthProto
    Communities []string
    Enable bool
    Engineid string
    Loglevel string
    Notraps []string
    Port int64
    Privacy SNMPPrivProto
    Syscontact string
    Syslocation string
    Targets []SNMPv1TrapTarget
    Users []SNMPUser
    Remoteusers []SNMPRemoteUser
    V3targets []SNMPv3Target
    Pid string
}






// ``SNMPHashConfig`` class Structure to provide up to two secrets to combine with the SNMPv3 engine ID and authentication or privacy protocol to form a localized hash. auth_hash is always required, priv_hash can be empty. By default arguments are paths on the local filesystem, raw_secret takes path to be the actual raw secret. First implementation was in ESXi: esxcli system snmp hash --help
 type SNMPHashConfig struct {
    AuthHash string
    PrivHash string
    RawSecret bool
}






// ``SNMPHashResults`` class Structure to provide operators diagnostics test results.
 type SNMPHashResults struct {
    AuthKey string
    PrivKey string
}









func resetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func resetOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func resetRestMetadata() protocol.OperationRestMetadata {
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


func enableInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func enableOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func enableRestMetadata() protocol.OperationRestMetadata {
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


func hashInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["config"] = bindings.NewReferenceType(SNMPHashConfigBindingType)
    fieldNameMap["config"] = "Config"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func hashOutputType() bindings.BindingType {
    return bindings.NewReferenceType(SNMPHashResultsBindingType)
}

func hashRestMetadata() protocol.OperationRestMetadata {
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


func limitsInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func limitsOutputType() bindings.BindingType {
    return bindings.NewReferenceType(SNMPLimitsBindingType)
}

func limitsRestMetadata() protocol.OperationRestMetadata {
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


func getInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func getOutputType() bindings.BindingType {
    return bindings.NewReferenceType(SNMPConfigReadOnlyBindingType)
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


func disableInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func disableOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func disableRestMetadata() protocol.OperationRestMetadata {
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


func setInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["config"] = bindings.NewReferenceType(SNMPConfigBindingType)
    fieldNameMap["config"] = "Config"
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
       map[string]int{"Error": 500})
}


func testInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func testOutputType() bindings.BindingType {
    return bindings.NewReferenceType(SNMPTestResultsBindingType)
}

func testRestMetadata() protocol.OperationRestMetadata {
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


func statsInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func statsOutputType() bindings.BindingType {
    return bindings.NewReferenceType(SNMPStatsBindingType)
}

func statsRestMetadata() protocol.OperationRestMetadata {
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



func SNMPLimitsBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["max_communities"] = bindings.NewIntegerType()
    fieldNameMap["max_communities"] = "MaxCommunities"
    fields["max_trap_destinations_v1"] = bindings.NewIntegerType()
    fieldNameMap["max_trap_destinations_v1"] = "MaxTrapDestinationsV1"
    fields["max_destinations_v3"] = bindings.NewIntegerType()
    fieldNameMap["max_destinations_v3"] = "MaxDestinationsV3"
    fields["max_notification_filters"] = bindings.NewIntegerType()
    fieldNameMap["max_notification_filters"] = "MaxNotificationFilters"
    fields["max_community_length"] = bindings.NewIntegerType()
    fieldNameMap["max_community_length"] = "MaxCommunityLength"
    fields["max_buffer_size"] = bindings.NewIntegerType()
    fieldNameMap["max_buffer_size"] = "MaxBufferSize"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.techpreview.monitoring.snmp.SNMP_limits",fields, reflect.TypeOf(SNMPLimits{}), fieldNameMap, validators)
}

func SNMPTestResultsBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["success"] = bindings.NewBooleanType()
    fieldNameMap["success"] = "Success"
    fields["message"] = bindings.NewStringType()
    fieldNameMap["message"] = "Message"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.techpreview.monitoring.snmp.SNMP_test_results",fields, reflect.TypeOf(SNMPTestResults{}), fieldNameMap, validators)
}

func SNMPStatsBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["sysuptime"] = bindings.NewStringType()
    fieldNameMap["sysuptime"] = "Sysuptime"
    fields["worstrtimelast"] = bindings.NewStringType()
    fieldNameMap["worstrtimelast"] = "Worstrtimelast"
    fields["avgresponsetime"] = bindings.NewStringType()
    fieldNameMap["avgresponsetime"] = "Avgresponsetime"
    fields["worstresponsetime"] = bindings.NewStringType()
    fieldNameMap["worstresponsetime"] = "Worstresponsetime"
    fields["inpkts"] = bindings.NewIntegerType()
    fieldNameMap["inpkts"] = "Inpkts"
    fields["outpkts"] = bindings.NewIntegerType()
    fieldNameMap["outpkts"] = "Outpkts"
    fields["usmstatsnotintimewindows"] = bindings.NewIntegerType()
    fieldNameMap["usmstatsnotintimewindows"] = "Usmstatsnotintimewindows"
    fields["usmstatsunknownusernames"] = bindings.NewIntegerType()
    fieldNameMap["usmstatsunknownusernames"] = "Usmstatsunknownusernames"
    fields["usmstatsunknownengineids"] = bindings.NewIntegerType()
    fieldNameMap["usmstatsunknownengineids"] = "Usmstatsunknownengineids"
    fields["usmstatswrongdigests"] = bindings.NewIntegerType()
    fieldNameMap["usmstatswrongdigests"] = "Usmstatswrongdigests"
    fields["usmstatsdecryptionerrors"] = bindings.NewIntegerType()
    fieldNameMap["usmstatsdecryptionerrors"] = "Usmstatsdecryptionerrors"
    fields["inbadversions"] = bindings.NewIntegerType()
    fieldNameMap["inbadversions"] = "Inbadversions"
    fields["inbadcommunitynames"] = bindings.NewIntegerType()
    fieldNameMap["inbadcommunitynames"] = "Inbadcommunitynames"
    fields["inbadcommunityuses"] = bindings.NewIntegerType()
    fieldNameMap["inbadcommunityuses"] = "Inbadcommunityuses"
    fields["inasnparseerrs"] = bindings.NewIntegerType()
    fieldNameMap["inasnparseerrs"] = "Inasnparseerrs"
    fields["intoobigs"] = bindings.NewIntegerType()
    fieldNameMap["intoobigs"] = "Intoobigs"
    fields["innosuchnames"] = bindings.NewIntegerType()
    fieldNameMap["innosuchnames"] = "Innosuchnames"
    fields["inbadvalues"] = bindings.NewIntegerType()
    fieldNameMap["inbadvalues"] = "Inbadvalues"
    fields["ingenerrs"] = bindings.NewIntegerType()
    fieldNameMap["ingenerrs"] = "Ingenerrs"
    fields["outtoobigs"] = bindings.NewIntegerType()
    fieldNameMap["outtoobigs"] = "Outtoobigs"
    fields["outnosuchnames"] = bindings.NewIntegerType()
    fieldNameMap["outnosuchnames"] = "Outnosuchnames"
    fields["outbadvalues"] = bindings.NewIntegerType()
    fieldNameMap["outbadvalues"] = "Outbadvalues"
    fields["outgenerrs"] = bindings.NewIntegerType()
    fieldNameMap["outgenerrs"] = "Outgenerrs"
    fields["outtraps"] = bindings.NewIntegerType()
    fieldNameMap["outtraps"] = "Outtraps"
    fields["silentdrops"] = bindings.NewIntegerType()
    fieldNameMap["silentdrops"] = "Silentdrops"
    fields["avgvarbinds"] = bindings.NewIntegerType()
    fieldNameMap["avgvarbinds"] = "Avgvarbinds"
    fields["maxvarbinds"] = bindings.NewIntegerType()
    fieldNameMap["maxvarbinds"] = "Maxvarbinds"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.techpreview.monitoring.snmp.SNMP_stats",fields, reflect.TypeOf(SNMPStats{}), fieldNameMap, validators)
}

func SNMPConfigBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["authentication"] = bindings.NewEnumType("com.vmware.appliance.techpreview.monitoring.snmp.SNMP_auth_proto", reflect.TypeOf(SNMPAuthProto(SNMPAuthProto_none)))
    fieldNameMap["authentication"] = "Authentication"
    fields["communities"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
    fieldNameMap["communities"] = "Communities"
    fields["engineid"] = bindings.NewStringType()
    fieldNameMap["engineid"] = "Engineid"
    fields["loglevel"] = bindings.NewStringType()
    fieldNameMap["loglevel"] = "Loglevel"
    fields["notraps"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
    fieldNameMap["notraps"] = "Notraps"
    fields["port"] = bindings.NewIntegerType()
    fieldNameMap["port"] = "Port"
    fields["privacy"] = bindings.NewEnumType("com.vmware.appliance.techpreview.monitoring.snmp.SNMP_priv_proto", reflect.TypeOf(SNMPPrivProto(SNMPPrivProto_AES128)))
    fieldNameMap["privacy"] = "Privacy"
    fields["remoteusers"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
    fieldNameMap["remoteusers"] = "Remoteusers"
    fields["syscontact"] = bindings.NewStringType()
    fieldNameMap["syscontact"] = "Syscontact"
    fields["syslocation"] = bindings.NewStringType()
    fieldNameMap["syslocation"] = "Syslocation"
    fields["targets"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
    fieldNameMap["targets"] = "Targets"
    fields["users"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
    fieldNameMap["users"] = "Users"
    fields["v3targets"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
    fieldNameMap["v3targets"] = "V3targets"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.techpreview.monitoring.snmp.SNMP_config",fields, reflect.TypeOf(SNMPConfig{}), fieldNameMap, validators)
}

func SNMPUserBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["username"] = bindings.NewStringType()
    fieldNameMap["username"] = "Username"
    fields["sec_level"] = bindings.NewEnumType("com.vmware.appliance.techpreview.monitoring.snmp.SNMP_sec_level", reflect.TypeOf(SNMPSecLevel(SNMPSecLevel_none)))
    fieldNameMap["sec_level"] = "SecLevel"
    fields["auth_key"] = bindings.NewStringType()
    fieldNameMap["auth_key"] = "AuthKey"
    fields["priv_key"] = bindings.NewStringType()
    fieldNameMap["priv_key"] = "PrivKey"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.techpreview.monitoring.snmp.SNMP_user",fields, reflect.TypeOf(SNMPUser{}), fieldNameMap, validators)
}

func SNMPv3TargetBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["type"] = bindings.NewEnumType("com.vmware.appliance.techpreview.monitoring.snmp.SNM_pv3_notfication", reflect.TypeOf(SNMPv3Notfication(SNMPv3Notfication_inform)))
    fieldNameMap["type"] = "Type_"
    fields["sec_level"] = bindings.NewEnumType("com.vmware.appliance.techpreview.monitoring.snmp.SNMP_sec_level", reflect.TypeOf(SNMPSecLevel(SNMPSecLevel_none)))
    fieldNameMap["sec_level"] = "SecLevel"
    fields["ip"] = bindings.NewStringType()
    fieldNameMap["ip"] = "Ip"
    fields["port"] = bindings.NewIntegerType()
    fieldNameMap["port"] = "Port"
    fields["user"] = bindings.NewStringType()
    fieldNameMap["user"] = "User"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.techpreview.monitoring.snmp.SNM_pv3_target",fields, reflect.TypeOf(SNMPv3Target{}), fieldNameMap, validators)
}

func SNMPv1TrapTargetBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["ip"] = bindings.NewStringType()
    fieldNameMap["ip"] = "Ip"
    fields["port"] = bindings.NewIntegerType()
    fieldNameMap["port"] = "Port"
    fields["community"] = bindings.NewStringType()
    fieldNameMap["community"] = "Community"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.techpreview.monitoring.snmp.SNM_pv1_trap_target",fields, reflect.TypeOf(SNMPv1TrapTarget{}), fieldNameMap, validators)
}

func SNMPRemoteUserBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["username"] = bindings.NewStringType()
    fieldNameMap["username"] = "Username"
    fields["sec_level"] = bindings.NewEnumType("com.vmware.appliance.techpreview.monitoring.snmp.SNMP_sec_level", reflect.TypeOf(SNMPSecLevel(SNMPSecLevel_none)))
    fieldNameMap["sec_level"] = "SecLevel"
    fields["authentication"] = bindings.NewEnumType("com.vmware.appliance.techpreview.monitoring.snmp.SNMP_auth_proto", reflect.TypeOf(SNMPAuthProto(SNMPAuthProto_none)))
    fieldNameMap["authentication"] = "Authentication"
    fields["auth_key"] = bindings.NewStringType()
    fieldNameMap["auth_key"] = "AuthKey"
    fields["privacy"] = bindings.NewEnumType("com.vmware.appliance.techpreview.monitoring.snmp.SNMP_priv_proto", reflect.TypeOf(SNMPPrivProto(SNMPPrivProto_AES128)))
    fieldNameMap["privacy"] = "Privacy"
    fields["priv_key"] = bindings.NewStringType()
    fieldNameMap["priv_key"] = "PrivKey"
    fields["engineid"] = bindings.NewStringType()
    fieldNameMap["engineid"] = "Engineid"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.techpreview.monitoring.snmp.SNMP_remote_user",fields, reflect.TypeOf(SNMPRemoteUser{}), fieldNameMap, validators)
}

func SNMPConfigReadOnlyBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["authentication"] = bindings.NewEnumType("com.vmware.appliance.techpreview.monitoring.snmp.SNMP_auth_proto", reflect.TypeOf(SNMPAuthProto(SNMPAuthProto_none)))
    fieldNameMap["authentication"] = "Authentication"
    fields["communities"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
    fieldNameMap["communities"] = "Communities"
    fields["enable"] = bindings.NewBooleanType()
    fieldNameMap["enable"] = "Enable"
    fields["engineid"] = bindings.NewStringType()
    fieldNameMap["engineid"] = "Engineid"
    fields["loglevel"] = bindings.NewStringType()
    fieldNameMap["loglevel"] = "Loglevel"
    fields["notraps"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
    fieldNameMap["notraps"] = "Notraps"
    fields["port"] = bindings.NewIntegerType()
    fieldNameMap["port"] = "Port"
    fields["privacy"] = bindings.NewEnumType("com.vmware.appliance.techpreview.monitoring.snmp.SNMP_priv_proto", reflect.TypeOf(SNMPPrivProto(SNMPPrivProto_AES128)))
    fieldNameMap["privacy"] = "Privacy"
    fields["syscontact"] = bindings.NewStringType()
    fieldNameMap["syscontact"] = "Syscontact"
    fields["syslocation"] = bindings.NewStringType()
    fieldNameMap["syslocation"] = "Syslocation"
    fields["targets"] = bindings.NewListType(bindings.NewReferenceType(SNMPv1TrapTargetBindingType), reflect.TypeOf([]SNMPv1TrapTarget{}))
    fieldNameMap["targets"] = "Targets"
    fields["users"] = bindings.NewListType(bindings.NewReferenceType(SNMPUserBindingType), reflect.TypeOf([]SNMPUser{}))
    fieldNameMap["users"] = "Users"
    fields["remoteusers"] = bindings.NewListType(bindings.NewReferenceType(SNMPRemoteUserBindingType), reflect.TypeOf([]SNMPRemoteUser{}))
    fieldNameMap["remoteusers"] = "Remoteusers"
    fields["v3targets"] = bindings.NewListType(bindings.NewReferenceType(SNMPv3TargetBindingType), reflect.TypeOf([]SNMPv3Target{}))
    fieldNameMap["v3targets"] = "V3targets"
    fields["pid"] = bindings.NewStringType()
    fieldNameMap["pid"] = "Pid"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.techpreview.monitoring.snmp.SNMP_config_read_only",fields, reflect.TypeOf(SNMPConfigReadOnly{}), fieldNameMap, validators)
}

func SNMPHashConfigBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["auth_hash"] = bindings.NewStringType()
    fieldNameMap["auth_hash"] = "AuthHash"
    fields["priv_hash"] = bindings.NewStringType()
    fieldNameMap["priv_hash"] = "PrivHash"
    fields["raw_secret"] = bindings.NewBooleanType()
    fieldNameMap["raw_secret"] = "RawSecret"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.techpreview.monitoring.snmp.SNMP_hash_config",fields, reflect.TypeOf(SNMPHashConfig{}), fieldNameMap, validators)
}

func SNMPHashResultsBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["auth_key"] = bindings.NewStringType()
    fieldNameMap["auth_key"] = "AuthKey"
    fields["priv_key"] = bindings.NewStringType()
    fieldNameMap["priv_key"] = "PrivKey"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.techpreview.monitoring.snmp.SNMP_hash_results",fields, reflect.TypeOf(SNMPHashResults{}), fieldNameMap, validators)
}


