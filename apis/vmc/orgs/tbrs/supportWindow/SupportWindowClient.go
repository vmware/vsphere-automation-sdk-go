/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: SupportWindow
 * Used by client-side stubs.
 */

package supportWindow

import (
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/apis/vmc/model"
)

type SupportWindowClient interface {


    // Get all available support windows
    //
    // @param orgParam Organization identifier. (required)
    // @param minimumSeatsAvailableParam minimum seats available (used as a filter) (optional)
    // @param createdByParam user name which was used to create the support window (used as a filter) (optional)
    // @throws Unauthenticated  Unauthorized
    // @throws InvalidRequest  Invalid request
    // @throws Unauthorized  Forbidden
    // @throws NotFound  No support windows are available
    Get(orgParam string, minimumSeatsAvailableParam *int64, createdByParam *string) ([]model.SupportWindow, error) 


    // Move Sddc to new support window
    //
    // @param orgParam Organization identifier. (required)
    // @param idParam Target Support Window ID (required)
    // @param sddcIdParam SDDC to move (required)
    // @return com.vmware.vmc.model.SupportWindowId
    // @throws Unauthenticated  Unauthorized
    // @throws InvalidRequest  Invalid request
    // @throws Unauthorized  Forbidden
    // @throws NotFound  Feature does not exist
    Put(orgParam string, idParam string, sddcIdParam model.SddcId) (model.SupportWindowId, error) 

}