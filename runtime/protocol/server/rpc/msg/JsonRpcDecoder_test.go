package msg_test

import (
	"github.com/stretchr/testify/assert"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol/server/rpc/msg"
	"testing"
)

func TestJsonRpcDecoder_DeSerializeApplicationContext(t *testing.T) {
	appCtxWireMap := map[string]interface{}{}
	appCtxWireMap["x"] = "a"

	appCtxWireMap["y"] = nil
	j := msg.NewJsonRpcDecoder()
	deserializedAppContext, err := j.DeSerializeApplicationContext(appCtxWireMap)
	assert.Nil(t, err)
	assert.Equal(t, *deserializedAppContext.GetProperty("x"), "a" )
	var stringPtr *string  =  nil
	assert.Equal(t, deserializedAppContext.GetProperty("y"),  stringPtr)
}
