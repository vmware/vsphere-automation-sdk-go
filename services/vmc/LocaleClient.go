/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Locale
 * Used by client-side stubs.
 */

package vmc

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/vmc/model"
)

type LocaleClient interface {

    // Sets the locale for the session which is used for translating responses.
    //
    // @param vmcLocaleParam The locale to be set. (required)
    // @return com.vmware.vmc.model.VmcLocale
    // @throws Unauthenticated  Unauthorized
    // @throws Unauthorized  Forbidden
	Set(vmcLocaleParam model.VmcLocale) (model.VmcLocale, error)
}
