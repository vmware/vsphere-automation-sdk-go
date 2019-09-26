package protocol

import (
	"fmt"
	"strings"

	"github.com/gorilla/mux"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/core"
)

type RestApiInterface interface {
	RegisterRoutesHandlers(router *mux.Router, provider core.APIProvider)
}

type OperationMetadata struct {
	methodDefinition *core.MethodDefinition
	inputType        bindings.StructType
	outputType       bindings.BindingType
	errorBindingMap  map[string]bindings.BindingType
	restMetadata     OperationRestMetadata
}

func NewOperationMetadata(methodDefinition *core.MethodDefinition,
	inputType bindings.StructType,
	outputType bindings.BindingType,
	errorBindingMap map[string]bindings.BindingType,
	restMetadata OperationRestMetadata) OperationMetadata {
	return OperationMetadata{methodDefinition: methodDefinition,
		inputType:       inputType,
		outputType:      outputType,
		errorBindingMap: errorBindingMap,
		restMetadata:    restMetadata}
}

func (meta OperationMetadata) MethodDefinition() *core.MethodDefinition {
	return meta.methodDefinition
}

func (meta OperationMetadata) InputType() bindings.StructType {
	return meta.inputType
}

func (meta OperationMetadata) OutputType() bindings.BindingType {
	return meta.outputType
}

func (meta OperationMetadata) ErrorBindingMap() map[string]bindings.BindingType {
	return meta.errorBindingMap
}

func (meta OperationMetadata) RestMetadata() OperationRestMetadata {
	return meta.restMetadata
}

// Rest metadata for name and types of query, header and
// body parameters of an operation. Example:
//	meta.ParamsTypeMap["input.nested.bparam"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
//	meta.ParamsTypeMap["input.nested.hparam"] = bindings.NewStringType()
//	meta.ParamsTypeMap["input.nested.qparam"] = bindings.NewStringType()
//	meta.QueryParams["qparam"] = "input.nested.qparam"
//	meta.HeaderParams["Hparam"] = "input.nested.hparam"
//	meta.BodyParam = "input.nested.bparam"
//	httpMethod = "GET|POST|UPDATE|PATCH|DELETE"
//	urlTemplate = "/newannotations/properties/{id}"
type OperationRestMetadata struct {
	// Flattened types of all parameters. Key is fully qualified field name
	paramsTypeMap map[string]bindings.BindingType
	//Names of rest parameter to fully qualified canonical name of the field
	pathParamsNameMap   map[string]string
	queryParamsNameMap  map[string]string
	headerParamsNameMap map[string]string
	//Fully qualified field name canonical name of body param
	bodyParamActualName string
	//HTTP action for the operation
	httpMethod string
	//HTTP URL for the operation
	urlTemplate string
	// HTTP response success code
	successCode int
	// vAPI error name to HTTP response error code mapping
	errorCodeMap map[string]int
	// Map from result field name to http header name
	resultHeadersNameMap map[string]string
	// Map from error field name to http header name
	errorHeadersNameMap map[string]string
}

func NewOperationRestMetadata(
	paramsTypeMap map[string]bindings.BindingType,
	pathParamsNameMap map[string]string,
	queryParamsNameMap map[string]string,
	headerParamsNameMap map[string]string,
	bodyParamActualName string,
	httpMethod string,
	urlTemplate string,
	resultHeadersNameMap map[string]string,
	successCode int,
	errorHeadersNameMap map[string]string,
	errorCodeMap map[string]int) OperationRestMetadata {

	return OperationRestMetadata{
		paramsTypeMap:        paramsTypeMap,
		pathParamsNameMap:    pathParamsNameMap,
		queryParamsNameMap:   queryParamsNameMap,
		headerParamsNameMap:  headerParamsNameMap,
		bodyParamActualName:  bodyParamActualName,
		httpMethod:           httpMethod,
		urlTemplate:          urlTemplate,
		successCode:          successCode,
		errorCodeMap:         errorCodeMap,
		resultHeadersNameMap: resultHeadersNameMap,
		errorHeadersNameMap:  errorHeadersNameMap}
}

func (meta OperationRestMetadata) ParamsTypeMap() map[string]bindings.BindingType {
	return meta.paramsTypeMap
}

func (meta OperationRestMetadata) PathParamsNameMap() map[string]string {
	return meta.pathParamsNameMap
}

func (meta OperationRestMetadata) QueryParamsNameMap() map[string]string {
	return meta.queryParamsNameMap
}

func (meta OperationRestMetadata) HeaderParamsNameMap() map[string]string {
	return meta.headerParamsNameMap
}

func (meta OperationRestMetadata) BodyParamActualName() string {
	return meta.bodyParamActualName
}

func (meta OperationRestMetadata) HttpMethod() string {
	return meta.httpMethod
}

func (meta OperationRestMetadata) UrlTemplate() string {
	return meta.urlTemplate
}

func (meta OperationRestMetadata) SuccessCode() int {
	return meta.successCode
}

func (meta OperationRestMetadata) ErrorCodeMap() map[string]int {
	return meta.errorCodeMap
}

func (meta OperationRestMetadata) ResultHeadersNameMap() map[string]string {
	return meta.resultHeadersNameMap
}
func (meta OperationRestMetadata) ErrorHeadersNameMap() map[string]string {
	return meta.errorHeadersNameMap
}

func (meta OperationRestMetadata) GetUrlPath(
	pathVariableFields map[string]string, queryParamFields map[string]string) string {
	urlPath := meta.urlTemplate
	// Substitute path variables with values in the template
	for fieldName, fieldStr := range pathVariableFields {
		urlPath = strings.Replace(urlPath, fmt.Sprintf("{%s}", meta.pathParamsNameMap[fieldName]), fieldStr, 1)
	}
	// Construct the query params portion of the url
	queryPrams := []string{}
	for fieldName, fieldStr := range queryParamFields {
		queryPrams = append(queryPrams, fmt.Sprintf("%s=%s", meta.queryParamsNameMap[fieldName], fieldStr))
	}
	queryParamStr := strings.Join(queryPrams, "&")

	if queryParamStr != "" {
		// Append the query params portion if it exists
		var connector string
		if strings.ContainsAny(urlPath, "?") {
			connector = "&"
		} else {
			connector = "?"
		}
		urlPath = strings.Join([]string{urlPath, queryParamStr}, connector)
	}
	return urlPath
}

func (meta OperationRestMetadata) PathVariableFieldNames() []string {
	fields := make([]string, 0, len(meta.pathParamsNameMap))
	for k := range meta.pathParamsNameMap {
		fields = append(fields, k)
	}
	return fields
}

func (meta OperationRestMetadata) QueryParamFieldNames() []string {
	fields := make([]string, 0, len(meta.queryParamsNameMap))
	for k := range meta.queryParamsNameMap {
		fields = append(fields, k)
	}
	return fields
}

func (meta OperationRestMetadata) HeaderFieldNames() []string {
	fields := make([]string, 0, len(meta.headerParamsNameMap))
	for k := range meta.headerParamsNameMap {
		fields = append(fields, k)
	}
	return fields
}
