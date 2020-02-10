/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: CounterMetadata.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package vstats

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)



// Counter metadata status. **Warning:** This enumeration is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type CounterMetadataCounterEditionStatus string

const (
    // The counter edition is current and is the default. **Warning:** This constant field is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	CounterMetadataCounterEditionStatus_DEFAULT CounterMetadataCounterEditionStatus = "DEFAULT"
    // The counter edition is current. This implies a support commitment. **Warning:** This constant field is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	CounterMetadataCounterEditionStatus_CURRENT CounterMetadataCounterEditionStatus = "CURRENT"
    // The counter edition is deprecated. It will be decommissioned rather soon. **Warning:** This constant field is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	CounterMetadataCounterEditionStatus_DEPRECATED CounterMetadataCounterEditionStatus = "DEPRECATED"
    // The counter edition is experimental. Consumers shouldn't rely on it for the long haul. **Warning:** This constant field is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	CounterMetadataCounterEditionStatus_EXPERIMENTAL CounterMetadataCounterEditionStatus = "EXPERIMENTAL"
    // The counter edition was removed. **Warning:** This constant field is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	CounterMetadataCounterEditionStatus_REMOVED CounterMetadataCounterEditionStatus = "REMOVED"
)

func (c CounterMetadataCounterEditionStatus) CounterMetadataCounterEditionStatus() bool {
	switch c {
	case CounterMetadataCounterEditionStatus_DEFAULT:
		return true
	case CounterMetadataCounterEditionStatus_CURRENT:
		return true
	case CounterMetadataCounterEditionStatus_DEPRECATED:
		return true
	case CounterMetadataCounterEditionStatus_EXPERIMENTAL:
		return true
	case CounterMetadataCounterEditionStatus_REMOVED:
		return true
	default:
		return false
	}
}


// Type of the sampled data. **Warning:** This enumeration is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type CounterMetadataSampleType string

const (
    // Raw samples. The value unprocessed as-is sampled. **Warning:** This constant field is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	CounterMetadataSampleType_RAW CounterMetadataSampleType = "RAW"
    // Absolute value samples. Represents an actual value of the counter. **Warning:** This constant field is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	CounterMetadataSampleType_ABSOLUTE CounterMetadataSampleType = "ABSOLUTE"
    // Fraction samples. Implies range from 0.00 to 1.00. **Warning:** This constant field is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	CounterMetadataSampleType_FRACTION CounterMetadataSampleType = "FRACTION"
    // Rate samples. Represents a value that has been normalized over the time period. **Warning:** This constant field is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	CounterMetadataSampleType_RATE CounterMetadataSampleType = "RATE"
    // Delta samples. Represents an amount of change for the counter between the current time-stamp and the last time-stamp when the counter was sampled. **Warning:** This constant field is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	CounterMetadataSampleType_DELTA CounterMetadataSampleType = "DELTA"
    // Log(n) samples. A natural logarithm of the value. **Warning:** This constant field is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	CounterMetadataSampleType_LOGN CounterMetadataSampleType = "LOGN"
)

func (s CounterMetadataSampleType) CounterMetadataSampleType() bool {
	switch s {
	case CounterMetadataSampleType_RAW:
		return true
	case CounterMetadataSampleType_ABSOLUTE:
		return true
	case CounterMetadataSampleType_FRACTION:
		return true
	case CounterMetadataSampleType_RATE:
		return true
	case CounterMetadataSampleType_DELTA:
		return true
	case CounterMetadataSampleType_LOGN:
		return true
	default:
		return false
	}
}


// Unit used by a metric. **Warning:** This enumeration is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type CounterMetadataMetricUnits string

const (
    // Percent. **Warning:** This constant field is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	CounterMetadataMetricUnits_PERCENT CounterMetadataMetricUnits = "PERCENT"
    // Number. **Warning:** This constant field is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	CounterMetadataMetricUnits_NUMBER CounterMetadataMetricUnits = "NUMBER"
    // Second. **Warning:** This constant field is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	CounterMetadataMetricUnits_SECOND CounterMetadataMetricUnits = "SECOND"
    // Hertz. **Warning:** This constant field is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	CounterMetadataMetricUnits_HERTZ CounterMetadataMetricUnits = "HERTZ"
    // Meter. **Warning:** This constant field is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	CounterMetadataMetricUnits_METER CounterMetadataMetricUnits = "METER"
    // Meters per second. **Warning:** This constant field is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	CounterMetadataMetricUnits_METERSPERSECOND CounterMetadataMetricUnits = "METERSPERSECOND"
    // Meters per second squared. **Warning:** This constant field is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	CounterMetadataMetricUnits_METERSPERSECONDSQUARED CounterMetadataMetricUnits = "METERSPERSECONDSQUARED"
    // Byte. **Warning:** This constant field is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	CounterMetadataMetricUnits_BYTE CounterMetadataMetricUnits = "BYTE"
    // Bit. **Warning:** This constant field is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	CounterMetadataMetricUnits_BIT CounterMetadataMetricUnits = "BIT"
    // Bytes per second. **Warning:** This constant field is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	CounterMetadataMetricUnits_BYTESPERSECOND CounterMetadataMetricUnits = "BYTESPERSECOND"
    // Bits per second. **Warning:** This constant field is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	CounterMetadataMetricUnits_BITSPERSECOND CounterMetadataMetricUnits = "BITSPERSECOND"
    // Kilogram. **Warning:** This constant field is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	CounterMetadataMetricUnits_KILOGRAM CounterMetadataMetricUnits = "KILOGRAM"
    // Gram. **Warning:** This constant field is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	CounterMetadataMetricUnits_GRAM CounterMetadataMetricUnits = "GRAM"
    // Celsius. **Warning:** This constant field is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	CounterMetadataMetricUnits_CELSIUS CounterMetadataMetricUnits = "CELSIUS"
    // Kelvin. **Warning:** This constant field is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	CounterMetadataMetricUnits_KELVIN CounterMetadataMetricUnits = "KELVIN"
    // Joule. **Warning:** This constant field is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	CounterMetadataMetricUnits_JOULE CounterMetadataMetricUnits = "JOULE"
    // Watt. **Warning:** This constant field is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	CounterMetadataMetricUnits_WATT CounterMetadataMetricUnits = "WATT"
    // Volt. **Warning:** This constant field is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	CounterMetadataMetricUnits_VOLT CounterMetadataMetricUnits = "VOLT"
    // Ampere. **Warning:** This constant field is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	CounterMetadataMetricUnits_AMPERE CounterMetadataMetricUnits = "AMPERE"
    // Volt Ampere. **Warning:** This constant field is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	CounterMetadataMetricUnits_VOLTAMPERE CounterMetadataMetricUnits = "VOLTAMPERE"
    // Candela. **Warning:** This constant field is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	CounterMetadataMetricUnits_CANDELA CounterMetadataMetricUnits = "CANDELA"
    // Mole. **Warning:** This constant field is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	CounterMetadataMetricUnits_MOLE CounterMetadataMetricUnits = "MOLE"
)

func (m CounterMetadataMetricUnits) CounterMetadataMetricUnits() bool {
	switch m {
	case CounterMetadataMetricUnits_PERCENT:
		return true
	case CounterMetadataMetricUnits_NUMBER:
		return true
	case CounterMetadataMetricUnits_SECOND:
		return true
	case CounterMetadataMetricUnits_HERTZ:
		return true
	case CounterMetadataMetricUnits_METER:
		return true
	case CounterMetadataMetricUnits_METERSPERSECOND:
		return true
	case CounterMetadataMetricUnits_METERSPERSECONDSQUARED:
		return true
	case CounterMetadataMetricUnits_BYTE:
		return true
	case CounterMetadataMetricUnits_BIT:
		return true
	case CounterMetadataMetricUnits_BYTESPERSECOND:
		return true
	case CounterMetadataMetricUnits_BITSPERSECOND:
		return true
	case CounterMetadataMetricUnits_KILOGRAM:
		return true
	case CounterMetadataMetricUnits_GRAM:
		return true
	case CounterMetadataMetricUnits_CELSIUS:
		return true
	case CounterMetadataMetricUnits_KELVIN:
		return true
	case CounterMetadataMetricUnits_JOULE:
		return true
	case CounterMetadataMetricUnits_WATT:
		return true
	case CounterMetadataMetricUnits_VOLT:
		return true
	case CounterMetadataMetricUnits_AMPERE:
		return true
	case CounterMetadataMetricUnits_VOLTAMPERE:
		return true
	case CounterMetadataMetricUnits_CANDELA:
		return true
	case CounterMetadataMetricUnits_MOLE:
		return true
	default:
		return false
	}
}


// Unit factor of measurement. **Warning:** This enumeration is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type CounterMetadataUnitsFactor string

const (
    // Yotta 10^24. **Warning:** This constant field is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	CounterMetadataUnitsFactor_YOTTA CounterMetadataUnitsFactor = "YOTTA"
    // Zetta 10^21. **Warning:** This constant field is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	CounterMetadataUnitsFactor_ZETTA CounterMetadataUnitsFactor = "ZETTA"
    // Exa 10^18. **Warning:** This constant field is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	CounterMetadataUnitsFactor_EXA CounterMetadataUnitsFactor = "EXA"
    // Peta 10^15. **Warning:** This constant field is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	CounterMetadataUnitsFactor_PETA CounterMetadataUnitsFactor = "PETA"
    // Tera 10^12. **Warning:** This constant field is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	CounterMetadataUnitsFactor_TERA CounterMetadataUnitsFactor = "TERA"
    // Giga 10^9. **Warning:** This constant field is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	CounterMetadataUnitsFactor_GIGA CounterMetadataUnitsFactor = "GIGA"
    // Mega 10^6. **Warning:** This constant field is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	CounterMetadataUnitsFactor_MEGA CounterMetadataUnitsFactor = "MEGA"
    // Kilo 10^3. **Warning:** This constant field is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	CounterMetadataUnitsFactor_KILO CounterMetadataUnitsFactor = "KILO"
    // Hecto 10^2. **Warning:** This constant field is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	CounterMetadataUnitsFactor_HECTO CounterMetadataUnitsFactor = "HECTO"
    // Deca 10. **Warning:** This constant field is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	CounterMetadataUnitsFactor_DECA CounterMetadataUnitsFactor = "DECA"
    // One. **Warning:** This constant field is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	CounterMetadataUnitsFactor_ONE CounterMetadataUnitsFactor = "ONE"
    // Deci 10^-1. **Warning:** This constant field is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	CounterMetadataUnitsFactor_DECI CounterMetadataUnitsFactor = "DECI"
    // Centi 10^-2. **Warning:** This constant field is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	CounterMetadataUnitsFactor_CENTI CounterMetadataUnitsFactor = "CENTI"
    // Milli 10^-3. **Warning:** This constant field is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	CounterMetadataUnitsFactor_MILLI CounterMetadataUnitsFactor = "MILLI"
    // Micro 10^-6. **Warning:** This constant field is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	CounterMetadataUnitsFactor_MICRO CounterMetadataUnitsFactor = "MICRO"
    // Nano 10^-9. **Warning:** This constant field is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	CounterMetadataUnitsFactor_NANO CounterMetadataUnitsFactor = "NANO"
    // Pico 10^-12. **Warning:** This constant field is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	CounterMetadataUnitsFactor_PIKO CounterMetadataUnitsFactor = "PIKO"
    // Femto 10^-15. **Warning:** This constant field is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	CounterMetadataUnitsFactor_FEMTO CounterMetadataUnitsFactor = "FEMTO"
    // Atto 10^-18. **Warning:** This constant field is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	CounterMetadataUnitsFactor_ATTO CounterMetadataUnitsFactor = "ATTO"
    // Zepto 10^-21. **Warning:** This constant field is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	CounterMetadataUnitsFactor_ZEPTO CounterMetadataUnitsFactor = "ZEPTO"
    // Yocto 10^-24. **Warning:** This constant field is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	CounterMetadataUnitsFactor_YOCTO CounterMetadataUnitsFactor = "YOCTO"
    // Yobi 2^80, 1024^8. **Warning:** This constant field is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	CounterMetadataUnitsFactor_YOBI CounterMetadataUnitsFactor = "YOBI"
    // Zebi 2^70, 1024^7. **Warning:** This constant field is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	CounterMetadataUnitsFactor_ZEBI CounterMetadataUnitsFactor = "ZEBI"
    // Exbi 2^60, 1024^6. **Warning:** This constant field is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	CounterMetadataUnitsFactor_EXBI CounterMetadataUnitsFactor = "EXBI"
    // Pebi 2^50, 1024^5. **Warning:** This constant field is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	CounterMetadataUnitsFactor_PEBI CounterMetadataUnitsFactor = "PEBI"
    // Tebi 2^40, 1024^4. **Warning:** This constant field is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	CounterMetadataUnitsFactor_TEBI CounterMetadataUnitsFactor = "TEBI"
    // Gibi 2^30, 1024^3. **Warning:** This constant field is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	CounterMetadataUnitsFactor_GIBI CounterMetadataUnitsFactor = "GIBI"
    // Mebi 2^20, 1024^2. **Warning:** This constant field is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	CounterMetadataUnitsFactor_MEBI CounterMetadataUnitsFactor = "MEBI"
    // Kibi 2^10, 1024. **Warning:** This constant field is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	CounterMetadataUnitsFactor_KIBI CounterMetadataUnitsFactor = "KIBI"
)

func (u CounterMetadataUnitsFactor) CounterMetadataUnitsFactor() bool {
	switch u {
	case CounterMetadataUnitsFactor_YOTTA:
		return true
	case CounterMetadataUnitsFactor_ZETTA:
		return true
	case CounterMetadataUnitsFactor_EXA:
		return true
	case CounterMetadataUnitsFactor_PETA:
		return true
	case CounterMetadataUnitsFactor_TERA:
		return true
	case CounterMetadataUnitsFactor_GIGA:
		return true
	case CounterMetadataUnitsFactor_MEGA:
		return true
	case CounterMetadataUnitsFactor_KILO:
		return true
	case CounterMetadataUnitsFactor_HECTO:
		return true
	case CounterMetadataUnitsFactor_DECA:
		return true
	case CounterMetadataUnitsFactor_ONE:
		return true
	case CounterMetadataUnitsFactor_DECI:
		return true
	case CounterMetadataUnitsFactor_CENTI:
		return true
	case CounterMetadataUnitsFactor_MILLI:
		return true
	case CounterMetadataUnitsFactor_MICRO:
		return true
	case CounterMetadataUnitsFactor_NANO:
		return true
	case CounterMetadataUnitsFactor_PIKO:
		return true
	case CounterMetadataUnitsFactor_FEMTO:
		return true
	case CounterMetadataUnitsFactor_ATTO:
		return true
	case CounterMetadataUnitsFactor_ZEPTO:
		return true
	case CounterMetadataUnitsFactor_YOCTO:
		return true
	case CounterMetadataUnitsFactor_YOBI:
		return true
	case CounterMetadataUnitsFactor_ZEBI:
		return true
	case CounterMetadataUnitsFactor_EXBI:
		return true
	case CounterMetadataUnitsFactor_PEBI:
		return true
	case CounterMetadataUnitsFactor_TEBI:
		return true
	case CounterMetadataUnitsFactor_GIBI:
		return true
	case CounterMetadataUnitsFactor_MEBI:
		return true
	case CounterMetadataUnitsFactor_KIBI:
		return true
	default:
		return false
	}
}


// The ``Info`` class contains information about CounterMetadata. It represents edition of the Counter. **Warning:** This class is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
type CounterMetadataInfo struct {
    // Counter Id. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	Cid string
    // Metadata Id. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	Mid string
    // Counter Edition status. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	Status CounterMetadataCounterEditionStatus
    // Numeric properties of the sampled data. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	Type_ CounterMetadataSampleType
    // The units of the measurement. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	Units CounterMetadataMetricUnits
    // Additional multiplier factors to be used with units. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	Scale *CounterMetadataUnitsFactor
    // Human legible localizable text about the counter. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	UserInfo *UserInfo
    // ID of the respective provider. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	Pid *string
}

// The ``FilterSpec`` class is used to filter the counter metadata list. **Warning:** This class is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
type CounterMetadataFilterSpec struct {
    // Counter edition status. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	Status *CounterMetadataCounterEditionStatus
}



func counterMetadataListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cid"] = bindings.NewIdType([]string{"com.vmware.vstats.model.Counter"}, "")
	fields["filter"] = bindings.NewOptionalType(bindings.NewReferenceType(CounterMetadataFilterSpecBindingType))
	fieldNameMap["cid"] = "Cid"
	fieldNameMap["filter"] = "Filter"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func counterMetadataListOutputType() bindings.BindingType {
	return bindings.NewListType(bindings.NewReferenceType(CounterMetadataInfoBindingType), reflect.TypeOf([]CounterMetadataInfo{}))
}

func counterMetadataListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["cid"] = bindings.NewIdType([]string{"com.vmware.vstats.model.Counter"}, "")
	fields["filter"] = bindings.NewOptionalType(bindings.NewReferenceType(CounterMetadataFilterSpecBindingType))
	fieldNameMap["cid"] = "Cid"
	fieldNameMap["filter"] = "Filter"
	paramsTypeMap["filter.status"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vstats.counter_metadata.counter_edition_status", reflect.TypeOf(CounterMetadataCounterEditionStatus(CounterMetadataCounterEditionStatus_DEFAULT))))
	paramsTypeMap["cid"] = bindings.NewIdType([]string{"com.vmware.vstats.model.Counter"}, "")
	paramsTypeMap["cid"] = bindings.NewIdType([]string{"com.vmware.vstats.model.Counter"}, "")
	pathParams["cid"] = "cid"
	queryParams["filter.status"] = "status"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"",
		"GET",
		"/stats/counters/{cid}/metadata",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"Error": 500,"InvalidArgument": 400,"NotFound": 404,"Unauthenticated": 401,"Unauthorized": 403})
}

func counterMetadataGetDefaultInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cid"] = bindings.NewIdType([]string{"com.vmware.vstats.model.Counter"}, "")
	fieldNameMap["cid"] = "Cid"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func counterMetadataGetDefaultOutputType() bindings.BindingType {
	return bindings.NewListType(bindings.NewReferenceType(CounterMetadataInfoBindingType), reflect.TypeOf([]CounterMetadataInfo{}))
}

func counterMetadataGetDefaultRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["cid"] = bindings.NewIdType([]string{"com.vmware.vstats.model.Counter"}, "")
	fieldNameMap["cid"] = "Cid"
	paramsTypeMap["cid"] = bindings.NewIdType([]string{"com.vmware.vstats.model.Counter"}, "")
	paramsTypeMap["cid"] = bindings.NewIdType([]string{"com.vmware.vstats.model.Counter"}, "")
	pathParams["cid"] = "cid"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"",
		"GET",
		"/stats/counters/{cid}/metadata/default",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"Error": 500,"InvalidArgument": 400,"NotFound": 404,"Unauthenticated": 401,"Unauthorized": 403})
}

func counterMetadataGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cid"] = bindings.NewIdType([]string{"com.vmware.vstats.model.Counter"}, "")
	fields["mid"] = bindings.NewIdType([]string{"com.vmware.vstats.model.CounterMetadata"}, "")
	fieldNameMap["cid"] = "Cid"
	fieldNameMap["mid"] = "Mid"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func counterMetadataGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(CounterMetadataInfoBindingType)
}

func counterMetadataGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["cid"] = bindings.NewIdType([]string{"com.vmware.vstats.model.Counter"}, "")
	fields["mid"] = bindings.NewIdType([]string{"com.vmware.vstats.model.CounterMetadata"}, "")
	fieldNameMap["cid"] = "Cid"
	fieldNameMap["mid"] = "Mid"
	paramsTypeMap["mid"] = bindings.NewIdType([]string{"com.vmware.vstats.model.CounterMetadata"}, "")
	paramsTypeMap["cid"] = bindings.NewIdType([]string{"com.vmware.vstats.model.Counter"}, "")
	paramsTypeMap["cid"] = bindings.NewIdType([]string{"com.vmware.vstats.model.Counter"}, "")
	paramsTypeMap["mid"] = bindings.NewIdType([]string{"com.vmware.vstats.model.CounterMetadata"}, "")
	pathParams["mid"] = "mid"
	pathParams["cid"] = "cid"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"",
		"GET",
		"/stats/counters/{cid}/metadata/{mid}",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"Error": 500,"InvalidArgument": 400,"NotFound": 404,"Unauthenticated": 401,"Unauthorized": 403})
}


func CounterMetadataInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cid"] = bindings.NewIdType([]string{"com.vmware.vstats.model.Counter"}, "")
	fieldNameMap["cid"] = "Cid"
	fields["mid"] = bindings.NewIdType([]string{"com.vmware.vstats.model.CounterMetadata"}, "")
	fieldNameMap["mid"] = "Mid"
	fields["status"] = bindings.NewEnumType("com.vmware.vstats.counter_metadata.counter_edition_status", reflect.TypeOf(CounterMetadataCounterEditionStatus(CounterMetadataCounterEditionStatus_DEFAULT)))
	fieldNameMap["status"] = "Status"
	fields["type"] = bindings.NewEnumType("com.vmware.vstats.counter_metadata.sample_type", reflect.TypeOf(CounterMetadataSampleType(CounterMetadataSampleType_RAW)))
	fieldNameMap["type"] = "Type_"
	fields["units"] = bindings.NewEnumType("com.vmware.vstats.counter_metadata.metric_units", reflect.TypeOf(CounterMetadataMetricUnits(CounterMetadataMetricUnits_PERCENT)))
	fieldNameMap["units"] = "Units"
	fields["scale"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vstats.counter_metadata.units_factor", reflect.TypeOf(CounterMetadataUnitsFactor(CounterMetadataUnitsFactor_YOTTA))))
	fieldNameMap["scale"] = "Scale"
	fields["user_info"] = bindings.NewOptionalType(bindings.NewReferenceType(UserInfoBindingType))
	fieldNameMap["user_info"] = "UserInfo"
	fields["pid"] = bindings.NewOptionalType(bindings.NewIdType([]string{"com.vmware.vstats.model.Provider"}, ""))
	fieldNameMap["pid"] = "Pid"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vstats.counter_metadata.info", fields, reflect.TypeOf(CounterMetadataInfo{}), fieldNameMap, validators)
}

func CounterMetadataFilterSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["status"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vstats.counter_metadata.counter_edition_status", reflect.TypeOf(CounterMetadataCounterEditionStatus(CounterMetadataCounterEditionStatus_DEFAULT))))
	fieldNameMap["status"] = "Status"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vstats.counter_metadata.filter_spec", fields, reflect.TypeOf(CounterMetadataFilterSpec{}), fieldNameMap, validators)
}

