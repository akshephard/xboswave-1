package main
import (
	"fmt"
	"github.com/gtfierro/xboswave/ingester/types"
	xbospb "github.com/gtfierro/xboswave/proto"
)
func has_device(msg xbospb.XBOS) bool {
	return msg.XBOSIoTDeviceState.Rtac!= nil
}
var device_units = map[string]string{
	"island_state":	"",
	"island_type":	"",
	"bess_availability":	"",
	"pge_state":	"",
	"pcc_breaker_state":	"",
	"bess_pv_breaker_state":	"",
	"heartbeat":	"",
	"real_power_setpoint":	"W",
	"reactive_power_setpoint":	"VAR",
	"target_real_power":	"W",
	"target_reactive_power":	"VAR",
	"battery_total_capacity":	"Wh",
	"battery_current_stored_energy":	"Wh",
	"total_actual_real_power":	"W",
	"total_actual_reactive_power":	"VAR",
	"total_actual_apparent_power":	"VA",
	"active_power_output_limit":	"W",
	"current_power_production":	"W",
	"ac_current_phase_a":	"A",
	"ac_current_phase_b":	"A",
	"ac_current_phase_c":	"A",
	"ac_voltage_ab":	"V",
	"ac_voltage_bc":	"V",
	"ac_voltage_ca":	"V",
	"ac_frequency":	"Hz",
	"fault_condition":	"",
	"pge_voltage":	"V",
	"pge_frequency":	"Hz",
}
var device_lookup = map[string]func(msg xbospb.XBOS) (float64, bool){

	"island_state": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Rtac.island_state != nil {
			return float64(msg.XBOSIoTDeviceState.Rtac.island_state.Value), true
		}
		return 0, false
	},
	"island_type": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Rtac.island_type != nil {
			return float64(msg.XBOSIoTDeviceState.Rtac.island_type.Value), true
		}
		return 0, false
	},
	"bess_availability": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Rtac.bess_availability != nil {
			return float64(msg.XBOSIoTDeviceState.Rtac.bess_availability.Value), true
		}
		return 0, false
	},
	"pge_state": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Rtac.pge_state != nil {
			return float64(msg.XBOSIoTDeviceState.Rtac.pge_state.Value), true
		}
		return 0, false
	},
	"pcc_breaker_state": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Rtac.pcc_breaker_state != nil {
			return float64(msg.XBOSIoTDeviceState.Rtac.pcc_breaker_state.Value), true
		}
		return 0, false
	},
	"bess_pv_breaker_state": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Rtac.bess_pv_breaker_state != nil {
			return float64(msg.XBOSIoTDeviceState.Rtac.bess_pv_breaker_state.Value), true
		}
		return 0, false
	},
	"heartbeat": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Rtac.heartbeat != nil {
			return float64(msg.XBOSIoTDeviceState.Rtac.heartbeat.Value), true
		}
		return 0, false
	},
	"real_power_setpoint": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Rtac.real_power_setpoint != nil {
			return float64(msg.XBOSIoTDeviceState.Rtac.real_power_setpoint.Value), true
		}
		return 0, false
	},
	"reactive_power_setpoint": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Rtac.reactive_power_setpoint != nil {
			return float64(msg.XBOSIoTDeviceState.Rtac.reactive_power_setpoint.Value), true
		}
		return 0, false
	},
	"target_real_power": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Rtac.target_real_power != nil {
			return float64(msg.XBOSIoTDeviceState.Rtac.target_real_power.Value), true
		}
		return 0, false
	},
	"target_reactive_power": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Rtac.target_reactive_power != nil {
			return float64(msg.XBOSIoTDeviceState.Rtac.target_reactive_power.Value), true
		}
		return 0, false
	},
	"battery_total_capacity": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Rtac.battery_total_capacity != nil {
			return float64(msg.XBOSIoTDeviceState.Rtac.battery_total_capacity.Value), true
		}
		return 0, false
	},
	"battery_current_stored_energy": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Rtac.battery_current_stored_energy != nil {
			return float64(msg.XBOSIoTDeviceState.Rtac.battery_current_stored_energy.Value), true
		}
		return 0, false
	},
	"total_actual_real_power": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Rtac.total_actual_real_power != nil {
			return float64(msg.XBOSIoTDeviceState.Rtac.total_actual_real_power.Value), true
		}
		return 0, false
	},
	"total_actual_reactive_power": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Rtac.total_actual_reactive_power != nil {
			return float64(msg.XBOSIoTDeviceState.Rtac.total_actual_reactive_power.Value), true
		}
		return 0, false
	},
	"total_actual_apparent_power": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Rtac.total_actual_apparent_power != nil {
			return float64(msg.XBOSIoTDeviceState.Rtac.total_actual_apparent_power.Value), true
		}
		return 0, false
	},
	"active_power_output_limit": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Rtac.active_power_output_limit != nil {
			return float64(msg.XBOSIoTDeviceState.Rtac.active_power_output_limit.Value), true
		}
		return 0, false
	},
	"current_power_production": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Rtac.current_power_production != nil {
			return float64(msg.XBOSIoTDeviceState.Rtac.current_power_production.Value), true
		}
		return 0, false
	},
	"ac_current_phase_a": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Rtac.ac_current_phase_a != nil {
			return float64(msg.XBOSIoTDeviceState.Rtac.ac_current_phase_a.Value), true
		}
		return 0, false
	},
	"ac_current_phase_b": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Rtac.ac_current_phase_b != nil {
			return float64(msg.XBOSIoTDeviceState.Rtac.ac_current_phase_b.Value), true
		}
		return 0, false
	},
	"ac_current_phase_c": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Rtac.ac_current_phase_c != nil {
			return float64(msg.XBOSIoTDeviceState.Rtac.ac_current_phase_c.Value), true
		}
		return 0, false
	},
	"ac_voltage_ab": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Rtac.ac_voltage_ab != nil {
			return float64(msg.XBOSIoTDeviceState.Rtac.ac_voltage_ab.Value), true
		}
		return 0, false
	},
	"ac_voltage_bc": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Rtac.ac_voltage_bc != nil {
			return float64(msg.XBOSIoTDeviceState.Rtac.ac_voltage_bc.Value), true
		}
		return 0, false
	},
	"ac_voltage_ca": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Rtac.ac_voltage_ca != nil {
			return float64(msg.XBOSIoTDeviceState.Rtac.ac_voltage_ca.Value), true
		}
		return 0, false
	},
	"ac_frequency": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Rtac.ac_frequency != nil {
			return float64(msg.XBOSIoTDeviceState.Rtac.ac_frequency.Value), true
		}
		return 0, false
	},
	"fault_condition": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Rtac.fault_condition != nil {
			return float64(msg.XBOSIoTDeviceState.Rtac.fault_condition.Value), true
		}
		return 0, false
	},
	"pge_voltage": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Rtac.pge_voltage != nil {
			return float64(msg.XBOSIoTDeviceState.Rtac.pge_voltage.Value), true
		}
		return 0, false
	},
	"pge_frequency": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Rtac.pge_frequency != nil {
			return float64(msg.XBOSIoTDeviceState.Rtac.pge_frequency.Value), true
		}
		return 0, false
	},
}
func build_device(uri types.SubscriptionURI, name string, msg xbospb.XBOS) types.ExtractedTimeseries {
	
	if extractfunc, found := device_lookup[name]; found {
		if value, found := extractfunc(msg); found {
			var extracted types.ExtractedTimeseries
			time := int64(msg.XBOSIoTDeviceState.Time)
			extracted.Values = append(extracted.Values, value)
			extracted.Times = append(extracted.Times, time)
			extracted.UUID = types.GenerateUUID(uri, []byte(name))
			extracted.Collection = fmt.Sprintf("xbos/%s", uri.Resource)
			extracted.Tags = map[string]string{
				"unit": device_units[name],
				"name": name,
			}
			return extracted
		}
	}
return types.ExtractedTimeseries{}
}

func Extract(uri types.SubscriptionURI, msg xbospb.XBOS, add func(types.ExtractedTimeseries) error) error {
	if msg.XBOSIoTDeviceState != nil {
		if has_device(msg) {
			for name := range device_lookup {
				extracted := build_device(uri, name, msg)
				if err := add(extracted); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
