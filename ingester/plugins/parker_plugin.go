package main

import (
	"fmt"
	"github.com/gtfierro/xboswave/ingester/types"
	xbospb "github.com/gtfierro/xboswave/proto"
)

func has_parker(msg xbospb.XBOS) bool {
	return msg.XBOSIoTDeviceState.ParkerState != nil
}

var weather_units = map[string]string{
	"compressor_working_hours":            "hours",
	"on_standby_status":       "unknown",
    "light_status":       "unknown",
    "aux_output_status":       "unknown",
	"next_defrost_counter": "seconds",
}

var weather_lookup = map[string]func(msg xbospb.XBOS) (float64, bool){
	// XBOSIoTDeviceState.ParkerState
	"compressor_working_hours": func(msg xbospb.XBOS) (float64, bool) {
		if has_parker(msg) && msg.XBOSIoTDeviceState.ParkerState.compressor_working_hours != nil {
			return float64(msg.XBOSIoTDeviceState.ParkerState.compressor_working_hours.Value), true
		}
		return 0, false
	},
	"on_standby_status": func(msg xbospb.XBOS) (float64, bool) {
		if has_parker(msg) && msg.XBOSIoTDeviceState.ParkerState.on_standby_status != nil {
			return float64(msg.XBOSIoTDeviceState.ParkerState.on_standby_status.Value), true
		}
		return 0, false
	},
	"light_status": func(msg xbospb.XBOS) (float64, bool) {
		if has_parker(msg) && msg.XBOSIoTDeviceState.ParkerState.light_status != nil {
			return float64(msg.XBOSIoTDeviceState.ParkerState.light_status.Value), true
		}
		return 0, false
	},
	"aux_output_status": func(msg xbospb.XBOS) (float64, bool) {
		if has_parker(msg) && msg.XBOSIoTDeviceState.ParkerState.aux_output_status != nil {
			return float64(msg.XBOSIoTDeviceState.ParkerState.aux_output_status.Value), true
		}
		return 0, false
	},
	"next_defrost_counter": func(msg xbospb.XBOS) (float64, bool) {
		if has_parker(msg) && msg.XBOSIoTDeviceState.ParkerState.next_defrost_counter != nil {
			return float64(msg.XBOSIoTDeviceState.ParkerState.next_defrost_counter.Value), true
		}
		return 0, false
	},
}

func build_parker(uri types.SubscriptionURI, name string, msg xbospb.XBOS) types.ExtractedTimeseries {
	if extractfunc, found := parker_lookup[name]; found {
		if value, found := extractfunc(msg); found {
			var extracted types.ExtractedTimeseries
			time := int64(msg.XBOSIoTDeviceState.Time)
			extracted.Values = append(extracted.Values, value)
			extracted.Times = append(extracted.Times, time)
			extracted.UUID = types.GenerateUUID(uri, []byte(name))
			extracted.Collection = fmt.Sprintf("xbos/%s", uri.Resource)
			extracted.Tags = map[string]string{
				"unit": parker_units[name],
				"name": name,
			}
			return extracted
		}
	}

	return types.ExtractedTimeseries{}
}

func Extract(uri types.SubscriptionURI, msg xbospb.XBOS, add func(types.ExtractedTimeseries) error) error {
	if msg.XBOSIoTDeviceState != nil {
		if has_parker(msg) {
			for name := range parker_lookup {
				extracted := build_parker(uri, name, msg)
				if err := add(extracted); err != nil {
					return err
				}
			}
		}

	}
	return nil
}
