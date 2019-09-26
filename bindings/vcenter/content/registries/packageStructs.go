/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for package: com.vmware.vcenter.content.registries.
 * Includes binding types of a top level structures and enumerations.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package registries



// The ``DayOfWeek`` enumeration class describes the supported days of the week to run a specific operation for a container registry. **Warning:** This enumeration is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.

type DayOfWeek string

const (
    // Sunday. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     DayOfWeek_SUNDAY DayOfWeek = "SUNDAY"
    // Monday. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     DayOfWeek_MONDAY DayOfWeek = "MONDAY"
    // Tuesday. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     DayOfWeek_TUESDAY DayOfWeek = "TUESDAY"
    // Wednesday. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     DayOfWeek_WEDNESDAY DayOfWeek = "WEDNESDAY"
    // Thursday. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     DayOfWeek_THURSDAY DayOfWeek = "THURSDAY"
    // Friday. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     DayOfWeek_FRIDAY DayOfWeek = "FRIDAY"
    // Saturday. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     DayOfWeek_SATURDAY DayOfWeek = "SATURDAY"
)

func (d DayOfWeek) DayOfWeek() bool {
    switch d {
        case DayOfWeek_SUNDAY:
            return true
        case DayOfWeek_MONDAY:
            return true
        case DayOfWeek_TUESDAY:
            return true
        case DayOfWeek_WEDNESDAY:
            return true
        case DayOfWeek_THURSDAY:
            return true
        case DayOfWeek_FRIDAY:
            return true
        case DayOfWeek_SATURDAY:
            return true
        default:
            return false
    }
}




// The ``Recurrence`` enumeration class defines the supported values for how often to run a specific operation for a container registry. **Warning:** This enumeration is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.

type Recurrence string

const (
    // No operation is scheduled. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     Recurrence_NONE Recurrence = "NONE"
    // An operation occurs on a daily basis. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     Recurrence_DAILY Recurrence = "DAILY"
    // An operation occurs on a weekly basis. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     Recurrence_WEEKLY Recurrence = "WEEKLY"
)

func (r Recurrence) Recurrence() bool {
    switch r {
        case Recurrence_NONE:
            return true
        case Recurrence_DAILY:
            return true
        case Recurrence_WEEKLY:
            return true
        default:
            return false
    }
}










