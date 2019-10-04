/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: InteropReport.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package interopReport

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/apis/vcenter/lcm"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/apis/vcenter/lcm/discovery"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
    "net/url"
    "time"
)




// The ``ReleaseInfo`` class contains a product release information.
 type ReleaseInfo struct {
    Version string
    Note *url.URL
}






// The ``ReportRow`` class contains the interoperability between a given product and the target product.
 type ReportRow struct {
    Product discovery.Product
    Compatible bool
    CompatibleReleases []ReleaseInfo
}






// The ``ReportSummary`` class contains a summary of the Report#products. It consists of the count of compatible and incompatible products to the target product.
 type ReportSummary struct {
    CompatibleCount int64
    IncompatibleCount int64
}






// The ``Report`` class contains the interoperability report between the target product and the other registered products in the vCenter Server instance.
 type Report struct {
    DateCreated time.Time
    TargetProduct discovery.Product
    Products []ReportRow
    Issues *lcm.Notifications
    Summary ReportSummary
}






// The ``Result`` class contains the result of interoperability report creation operation.
 type Result struct {
    Report Report
    CsvReport *string
}






// Configuration of report generation.
 type Spec struct {
    TargetVersion string
}









func createInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["spec"] = bindings.NewOptionalType(bindings.NewReferenceType(SpecBindingType))
    fieldNameMap["spec"] = "Spec"
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
    paramsTypeMap["spec"] = bindings.NewOptionalType(bindings.NewReferenceType(SpecBindingType))
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "spec",
      "POST",
      "/vcenter/lcm/discovery/interop-report",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"Unauthenticated": 401,"Error": 500})
}



func ReleaseInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["version"] = bindings.NewStringType()
    fieldNameMap["version"] = "Version"
    fields["note"] = bindings.NewOptionalType(bindings.NewUriType())
    fieldNameMap["note"] = "Note"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.lcm.discovery.interop_report.release_info",fields, reflect.TypeOf(ReleaseInfo{}), fieldNameMap, validators)
}

func ReportRowBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["product"] = bindings.NewReferenceType(discovery.ProductBindingType)
    fieldNameMap["product"] = "Product"
    fields["compatible"] = bindings.NewBooleanType()
    fieldNameMap["compatible"] = "Compatible"
    fields["compatible_releases"] = bindings.NewListType(bindings.NewReferenceType(ReleaseInfoBindingType), reflect.TypeOf([]ReleaseInfo{}))
    fieldNameMap["compatible_releases"] = "CompatibleReleases"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.lcm.discovery.interop_report.report_row",fields, reflect.TypeOf(ReportRow{}), fieldNameMap, validators)
}

func ReportSummaryBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["compatible_count"] = bindings.NewIntegerType()
    fieldNameMap["compatible_count"] = "CompatibleCount"
    fields["incompatible_count"] = bindings.NewIntegerType()
    fieldNameMap["incompatible_count"] = "IncompatibleCount"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.lcm.discovery.interop_report.report_summary",fields, reflect.TypeOf(ReportSummary{}), fieldNameMap, validators)
}

func ReportBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["date_created"] = bindings.NewDateTimeType()
    fieldNameMap["date_created"] = "DateCreated"
    fields["target_product"] = bindings.NewReferenceType(discovery.ProductBindingType)
    fieldNameMap["target_product"] = "TargetProduct"
    fields["products"] = bindings.NewListType(bindings.NewReferenceType(ReportRowBindingType), reflect.TypeOf([]ReportRow{}))
    fieldNameMap["products"] = "Products"
    fields["issues"] = bindings.NewOptionalType(bindings.NewReferenceType(lcm.NotificationsBindingType))
    fieldNameMap["issues"] = "Issues"
    fields["summary"] = bindings.NewReferenceType(ReportSummaryBindingType)
    fieldNameMap["summary"] = "Summary"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.lcm.discovery.interop_report.report",fields, reflect.TypeOf(Report{}), fieldNameMap, validators)
}

func ResultBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["report"] = bindings.NewReferenceType(ReportBindingType)
    fieldNameMap["report"] = "Report"
    fields["csv_report"] = bindings.NewOptionalType(bindings.NewIdType([]string {"com.vmware.vcenter.lcm.report"}, ""))
    fieldNameMap["csv_report"] = "CsvReport"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.lcm.discovery.interop_report.result",fields, reflect.TypeOf(Result{}), fieldNameMap, validators)
}

func SpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["target_version"] = bindings.NewStringType()
    fieldNameMap["target_version"] = "TargetVersion"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.lcm.discovery.interop_report.spec",fields, reflect.TypeOf(Spec{}), fieldNameMap, validators)
}


