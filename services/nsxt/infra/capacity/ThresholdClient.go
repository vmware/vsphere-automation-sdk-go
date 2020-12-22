/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Threshold
 * Used by client-side stubs.
 */

package capacity

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt/model"
)

type ThresholdClient interface {

    // Returns warning threshold(s) set for NSX Objects.
    // @return com.vmware.nsx_policy.model.CapacityThresholdList
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get() (model.CapacityThresholdList, error)

    // Updates the warning threshold(s) for NSX Objects specified, does not modify thresholds for any other objects.
    //
    // @param capacityThresholdParam (required)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Patch(capacityThresholdParam model.CapacityThreshold) error

    // Updates the warning threshold(s) for NSX Objects specified, and returns new threshold(s). Threshold list in the request must contain value for GLOBAL_DEFAULT threshold_type which represents global thresholds.
    //
    // @param capacityThresholdListParam (required)
    // @return com.vmware.nsx_policy.model.CapacityThresholdList
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Update(capacityThresholdListParam model.CapacityThresholdList) (model.CapacityThresholdList, error)
}
