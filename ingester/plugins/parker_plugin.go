package main

import (
	"fmt"
	"github.com/gtfierro/xboswave/ingester/types"
	xbospb "github.com/gtfierro/xboswave/proto"
)


func has_parker(msg xbospb.XBOS) bool {
	return msg.XBOSIoTDeviceState.ParkerState != nil
}

var parker_units = map[string]string{
	"CompressorWorkingHours":            "hours",
	"OnStandbyStatus":       "unknown",
    "LightStatus":       "unknown",
    "AuxOutputStatus":       "unknown",
	"NextDefrostCounter": "seconds",
}

var parker_lookup = map[string]func(msg xbospb.XBOS) (float64, bool){
	// XBOSIoTDeviceState.ParkerState
	"CompressorWorkingHours": func(msg xbospb.XBOS) (float64, bool) {
		if has_parker(msg) && msg.XBOSIoTDeviceState.ParkerState.CompressorWorkingHours != nil {
			return float64(msg.XBOSIoTDeviceState.ParkerState.CompressorWorkingHours.Value), true
		}
		return 0, false
	},
	"OnStandbyStatus": func(msg xbospb.XBOS) (float64, bool) {
		if has_parker(msg) && msg.XBOSIoTDeviceState.ParkerState.OnStandbyStatus != nil {
			return float64(msg.XBOSIoTDeviceState.ParkerState.OnStandbyStatus.Value), true
		}
		return 0, false
	},
	"LightStatus": func(msg xbospb.XBOS) (float64, bool) {
		if has_parker(msg) && msg.XBOSIoTDeviceState.ParkerState.LightStatus != nil {
			return float64(msg.XBOSIoTDeviceState.ParkerState.LightStatus.Value), true
		}
		return 0, false
	},
	"AuxOutputStatus": func(msg xbospb.XBOS) (float64, bool) {
		if has_parker(msg) && msg.XBOSIoTDeviceState.ParkerState.AuxOutputStatus != nil {
			return float64(msg.XBOSIoTDeviceState.ParkerState.AuxOutputStatus.Value), true
		}
		return 0, false
	},
	"NextDefrostCounter": func(msg xbospb.XBOS) (float64, bool) {
		if has_parker(msg) && msg.XBOSIoTDeviceState.ParkerState.NextDefrostCounter != nil {
			return float64(msg.XBOSIoTDeviceState.ParkerState.NextDefrostCounter.Value), true
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
