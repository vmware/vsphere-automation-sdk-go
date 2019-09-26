package l10n

import (
	"bytes"

	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/l10n/runtime"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/log"
)

var runtimeMessageFormatter *MessageFormatter

// Error formatter with default localization parameters
// backed by message bundle for the runtime
func NewDefaultRuntimeMessageFormatter() *MessageFormatter {
	if runtimeMessageFormatter == nil {
		runtimeFactory := NewMessageFactory()
		err := runtimeFactory.AddBundle("en", bytes.NewReader(runtime.RuntimeProperties_EN))
		if err != nil {
			log.Error("Error creating runtime message factory")
		}
		formatter := runtimeFactory.GetDefaultFormatter()
		runtimeMessageFormatter = formatter
	}
	return runtimeMessageFormatter
}
