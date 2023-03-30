/* Copyright © 2021 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

package main

import (
	"context"
	"fmt"
	"github.com/buger/jsonparser"
	"github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std"
	"github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/examples/stream/client/bindings/src/generated/vmodl/examples/greeter"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/http"
	"time"
)

// CustomFrameDeserializer prints raw http chunks thus skipping deserialization
type CustomFrameDeserializer struct{}

func NewCustomFrameDeserializer() *CustomFrameDeserializer {
	return &CustomFrameDeserializer{}
}

var _ http.ClientFrameDeserializer = &CustomFrameDeserializer{}

// DeserializeFrames interrupts frames deserialization and prints raw frame response
//
func (c *CustomFrameDeserializer) DeserializeFrames(ctx context.Context, frames chan []byte) (chan core.MonoResult, error) {
	result := make(chan core.MonoResult)

	go func() {
		defer close(result)

		for {
			select {
			case <-ctx.Done():
				err := ctx.Err()
				// context canceled
				if err == context.Canceled {
					return
				}
				// another error occurred
				if err != nil {
					// error populated by go runtime
					internalServerErr := errors.NewInternalServerError()
					internalServerErr.Messages = []std.LocalizableMessage{
						{
							DefaultMessage: err.Error(),
						},
					}
					dataVal, _ := internalServerErr.GetDataValue__()
					result <- core.NewMonoResult(nil, dataVal.(*data.ErrorValue))
				}
				return
			case frame, ok := <-frames:
				if !ok {
					// on closed channel return
					return
				}
				_, _, _, err := jsonparser.Get(frame, "result", "output")
				if err != nil {
					fmt.Println("Raw error frame received: ", string(frame))
					return
				} else {
					fmt.Println("Raw valid frame received: ", string(frame))
				}
			}
		}
	}()

	// note that result channel will be empty in this case
	return result, nil
}

func callServer(connector client.Connector) {
	greeterClient := greeter.NewGreeterClient(connector)
	resultChannel, _, errorChannel := greeterClient.Greet()

	for {
		select {
		case result, ok := <-resultChannel:
			if !ok {
				// if channel closed -> no more data from server, return
				return
			}
			fmt.Println(result)
		case err, ok := <-errorChannel:
			if !ok {
				return
			}
			fmt.Println("An error occurred: ", err.Error())
			return
		case <-time.After(30 * time.Second):
			fmt.Println("No frames returned for 30 seconds. Closing connection...")
			return
		}
	}
}

func main() {
	// make a stream call with default connector option - using Json-RPC underlying protocol
	connector := client.NewConnector(
		"http://localhost:8088")
	callServer(connector)

	// make a stream call using custom response deserializer
	// allowing to catch raw frames before they are deserialized by the runtime
	connector = client.NewConnector(
		"http://localhost:8088",
		client.WithClientFrameDeserializer(NewCustomFrameDeserializer()))
	callServer(connector)

	// make a stream call using custom response deserializer and clean-json wire format
	// actual call not made in this example as go server still does not support clean-json stream protocol
	connector = client.NewConnector(
		"http://localhost:8088",
		client.WithClientFrameDeserializer(NewCustomFrameDeserializer()),
		client.WithStreamingProtocol(lib.VAPI_STREAMING_CLEAN_JSON_CONTENT_TYPE))
	// callServer(connector)
}
