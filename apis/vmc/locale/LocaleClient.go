/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Locale
 * Used by client-side stubs.
 */

package locale

import (
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/apis/vmc/model"
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
