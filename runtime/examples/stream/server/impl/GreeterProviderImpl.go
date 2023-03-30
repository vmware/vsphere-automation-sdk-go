// Copyright © 2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause
package impl

import (
	"context"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/log"
)

type GreeterImpl struct {
}

func NewGreeterServiceImpl() *GreeterImpl {
	return &GreeterImpl{}
}

func (GreeterImpl GreeterImpl) Greet(ctx *core.ExecutionContext) (chan string, chan error) {
	resultChannel := make(chan string)
	errorChannel := make(chan error)

	greetings := []string{
		"Hello",
		"Здравей",
		"Bonjour",
		"Nǐn hǎo",
		"Ciao",
	}

	go func(ctx context.Context, resultChannel chan string, errorChannel chan error) {
		defer close(resultChannel)
		defer close(errorChannel)

		i := 0
		for {
			select {
			case <-ctx.Done():
				log.Debug("Client cancelled the request. exiting")
				break
			default:
				if i > len(greetings)-1 {
					return
				}
				resultChannel <- greetings[i]
				i++
			}
		}
	}(ctx.Context(), resultChannel, errorChannel)

	return resultChannel, errorChannel
}
