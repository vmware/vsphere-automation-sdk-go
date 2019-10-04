/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: PrecheckReport.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package precheckReport

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/apis/vcenter/lcm"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
    "time"
)

// Resource type for precheck report
const PrecheckReport_RESOURCE_TYPE = "com.vmware.vcenter.lcm.report"



// The ``Summary`` Class contains the summary of precheck report.
 type ReportSummary struct {
    ErrorCount int64
    WarningCount int64
}






// The ``Report`` class contains estimates of how long it will take an update as well as a list of possible warnings and errors with applying the update.
 type Report struct {
    DateCreated time.Time
    EstimatedTimeToUpdate *int64
    Issues *lcm.Notifications
    Summary ReportSummary
}






// The ``Result`` class contains the precheck report and a link to download the CSV report.
 type Result struct {
    Report Report
    CsvReport *string
}









func createInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["version"] = bindings.NewIdType([]string {"com.vmware.vcenter.lcm.update.pending"}, "")
    fieldNameMap["version"] = "Version"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func createOutputType() bindings.BindingType {
    return bindings.NewReferenceType(ResultBindingType)
}

func createRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["version"] = bindings.NewIdType([]string {"com.vmware.vcenter.lcm.update.pending"}, "")
    paramsTypeMap["version"] = bindings.NewIdType([]string {"com.vmware.vcenter.lcm.update.pending"}, "")
    pathParams["version"] = "version"
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "POST",
      "/vcenter/lcm/update/pending/{version}/precheck-report",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"Unauthenticated": 401,"NotFound": 404,"NotAllowedInCurrentState": 400,"Error": 500})
}



func ReportSummaryBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["error_count"] = bindings.NewIntegerType()
    fieldNameMap["error_count"] = "ErrorCount"
    fields["warning_count"] = bindings.NewIntegerType()
    fieldNameMap["warning_count"] = "WarningCount"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.lcm.update.precheck_report.report_summary",fields, reflect.TypeOf(ReportSummary{}), fieldNameMap, validators)
}

func ReportBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["date_created"] = bindings.NewDateTimeType()
    fieldNameMap["date_created"] = "DateCreated"
    fields["estimated_time_to_update"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["estimated_time_to_update"] = "EstimatedTimeToUpdate"
    fields["issues"] = bindings.NewOptionalType(bindings.NewReferenceType(lcm.NotificationsBindingType))
    fieldNameMap["issues"] = "Issues"
    fields["summary"] = bindings.NewReferenceType(ReportSummaryBindingType)
    fieldNameMap["summary"] = "Summary"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.lcm.update.precheck_report.report",fields, reflect.TypeOf(Report{}), fieldNameMap, validators)
}

func ResultBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["report"] = bindings.NewReferenceType(ReportBindingType)
    fieldNameMap["report"] = "Report"
    fields["csv_report"] = bindings.NewOptionalType(bindings.NewIdType([]string {"com.vmware.vcenter.lcm.report"}, ""))
    fieldNameMap["csv_report"] = "CsvReport"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.lcm.update.precheck_report.result",fields, reflect.TypeOf(Result{}), fieldNameMap, validators)
}


