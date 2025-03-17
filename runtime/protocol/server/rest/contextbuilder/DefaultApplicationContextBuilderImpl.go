// Copyright (c) 2020-2024 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

package contextbuilder

import (
	"net/http"
	"strings"

	"github.com/vmware/vsphere-automation-sdk-go/runtime/common"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib"
)

type DefaultApplicationContextBuilderImpl struct {
}

func NewDefaultApplicationContextBuilderImpl() *DefaultApplicationContextBuilderImpl {
	return &DefaultApplicationContextBuilderImpl{}
}

func (appCtx *DefaultApplicationContextBuilderImpl) canonicalKeyMap(data map[string]string) map[string]string {
	res := make(map[string]string)
	for key, value := range data {
		res[http.CanonicalHeaderKey(key)] = value
	}
	return res
}

func (appCtx *DefaultApplicationContextBuilderImpl) BuildApplicationContext(r *http.Request) (*core.ApplicationContext, error) {
	header := r.Header

	wireData := make(map[string]*string)
	headerTokeyMap := map[string]string{
		lib.REST_OP_ID_HEADER:                                lib.OPID,
		lib.HTTP_ACCEPT_LANGUAGE:                             lib.HTTP_ACCEPT_LANGUAGE,
		lib.VAPI_HEADER_PREFIX + lib.VAPI_L10N_FORMAT_LOCALE: lib.VAPI_L10N_FORMAT_LOCALE,
		lib.VAPI_HEADER_PREFIX + lib.VAPI_L10N_TIMEZONE:      lib.VAPI_L10N_TIMEZONE,
	}

	canonicalHeaderTokeyMap := appCtx.canonicalKeyMap(headerTokeyMap)

	for key, value := range header {
		if wirekey, ok := canonicalHeaderTokeyMap[key]; ok {
			if len(value) == 1 {
				wireData[wirekey] = &value[0]
			} else {
				temp := strings.Join(value, ",")
				wireData[wirekey] = &temp
			}
		}
	}

	if len(wireData) > 0 {
		return core.NewApplicationContext(wireData), nil
	}

	return common.NewDefaultApplicationContext(), nil
}

func (appCtx *DefaultApplicationContextBuilderImpl) CanHandle(r *http.Request) bool {
	/** DefaultApplicationContextBuilderImpl can always return a valid ApplicationContext i.e
	 * either by creating ApplicationContext by extracting data from header or by returning
	 * ApplicationContext created by NewDefaultApplicationContext present in goruntime/common
	 */
	return true
}
