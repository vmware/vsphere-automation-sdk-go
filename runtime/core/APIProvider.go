package core

import "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"

type APIProvider interface {
	/**
	 * Invokes the specified method using the input DataValue and execution context.
	 */
	Invoke(serviceId string, operationId string, inputValue data.DataValue, ctx *ExecutionContext) MethodResult
}
