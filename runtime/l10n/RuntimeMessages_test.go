package l10n_test

import (
	"github.com/stretchr/testify/assert"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/l10n"
	"testing"
)

func TestValidRuntimeMessage(t *testing.T) {
	formatter := l10n.NewDefaultRuntimeMessageFormatter()

	assert.Equal(t, "Invalid input to method '{methodId}'",
		formatter.GetPositionalArgMessage("vapi.method.input.invalid"))
}

func TestInvalidRuntimeMessageId(t *testing.T) {
	formatter := l10n.NewDefaultRuntimeMessageFormatter()

	assert.Equal(t, "Unknown message ID some.thing requested",
		formatter.GetPositionalArgMessage("some.thing"))
}
