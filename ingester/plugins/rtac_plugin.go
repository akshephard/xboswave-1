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

	"IslandState": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Rtac.IslandState != nil {
            if msg.XBOSIoTDeviceState.Rtac.IslandState.Value{
                fmt.Println("The value is %v",msg.XBOSIoTDeviceState.Rtac.IslandState.Value)
                return 1, true
            } else {
                fmt.Println("The value is %v",msg.XBOSIoTDeviceState.Rtac.IslandState.Value)
                return 0, true
            }
		}
		return 0, false
	},
    /*
	"IslandType": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Rtac.IslandType != nil {
			return float64(msg.XBOSIoTDeviceState.Rtac.IslandType.Value), true
		}
		return 0, false
	},
	"BessAvailability": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Rtac.BessAvailability != nil {
			return float64(msg.XBOSIoTDeviceState.Rtac.BessAvailability.Value), true
		}
		return 0, false
	},
	"PgeState": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Rtac.PgeState != nil {
			return float64(msg.XBOSIoTDeviceState.Rtac.PgeState.Value), true
		}
		return 0, false
	},
	"PccBreakerState": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Rtac.PccBreakerState != nil {
			return float64(msg.XBOSIoTDeviceState.Rtac.PccBreakerState.Value), true
		}
		return 0, false
	},
	"BessPvBreakerState": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Rtac.BessPvBreakerState != nil {
			return float64(msg.XBOSIoTDeviceState.Rtac.BessPvBreakerState.Value), true
		}
		return 0, false
	},
    */
	"Heartbeat": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Rtac.Heartbeat != nil {
			return float64(msg.XBOSIoTDeviceState.Rtac.Heartbeat.Value), true
		}
		return 0, false
	},
	"RealPowerSetpoint": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Rtac.RealPowerSetpoint != nil {
			return float64(msg.XBOSIoTDeviceState.Rtac.RealPowerSetpoint.Value), true
		}
		return 0, false
	},
	"ReactivePowerSetpoint": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Rtac.ReactivePowerSetpoint != nil {
			return float64(msg.XBOSIoTDeviceState.Rtac.ReactivePowerSetpoint.Value), true
		}
		return 0, false
	},
	"TargetRealPower": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Rtac.TargetRealPower != nil {
			return float64(msg.XBOSIoTDeviceState.Rtac.TargetRealPower.Value), true
		}
		return 0, false
	},
	"TargetReactivePower": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Rtac.TargetReactivePower != nil {
			return float64(msg.XBOSIoTDeviceState.Rtac.TargetReactivePower.Value), true
		}
		return 0, false
	},
	"BatteryTotalCapacity": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Rtac.BatteryTotalCapacity != nil {
			return float64(msg.XBOSIoTDeviceState.Rtac.BatteryTotalCapacity.Value), true
		}
		return 0, false
	},
	"BatteryCurrentStoredEnergy": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Rtac.BatteryCurrentStoredEnergy != nil {
			return float64(msg.XBOSIoTDeviceState.Rtac.BatteryCurrentStoredEnergy.Value), true
		}
		return 0, false
	},
	"TotalActualRealPower": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Rtac.TotalActualRealPower != nil {
			return float64(msg.XBOSIoTDeviceState.Rtac.TotalActualRealPower.Value), true
		}
		return 0, false
	},
	"TotalActualReactivePower": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Rtac.TotalActualReactivePower != nil {
			return float64(msg.XBOSIoTDeviceState.Rtac.TotalActualReactivePower.Value), true
		}
		return 0, false
	},
	"TotalActualApparentPower": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Rtac.TotalActualApparentPower != nil {
			return float64(msg.XBOSIoTDeviceState.Rtac.TotalActualApparentPower.Value), true
		}
		return 0, false
	},
	"ActivePowerOutputLimit": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Rtac.ActivePowerOutputLimit != nil {
			return float64(msg.XBOSIoTDeviceState.Rtac.ActivePowerOutputLimit.Value), true
		}
		return 0, false
	},
	"CurrentPowerProduction": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Rtac.CurrentPowerProduction != nil {
			return float64(msg.XBOSIoTDeviceState.Rtac.CurrentPowerProduction.Value), true
		}
		return 0, false
	},
	"AcCurrentPhaseA": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Rtac.AcCurrentPhaseA != nil {
			return float64(msg.XBOSIoTDeviceState.Rtac.AcCurrentPhaseA.Value), true
		}
		return 0, false
	},
	"AcCurrentPhaseB": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Rtac.AcCurrentPhaseB != nil {
			return float64(msg.XBOSIoTDeviceState.Rtac.AcCurrentPhaseB.Value), true
		}
		return 0, false
	},
	"AcCurrentPhaseC": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Rtac.AcCurrentPhaseC != nil {
			return float64(msg.XBOSIoTDeviceState.Rtac.AcCurrentPhaseC.Value), true
		}
		return 0, false
	},
	"AcVoltageAb": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Rtac.AcVoltageAb != nil {
			return float64(msg.XBOSIoTDeviceState.Rtac.AcVoltageAb.Value), true
		}
		return 0, false
	},
	"AcVoltageBc": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Rtac.AcVoltageBc != nil {
			return float64(msg.XBOSIoTDeviceState.Rtac.AcVoltageBc.Value), true
		}
		return 0, false
	},
	"AcVoltageCa": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Rtac.AcVoltageCa != nil {
			return float64(msg.XBOSIoTDeviceState.Rtac.AcVoltageCa.Value), true
		}
		return 0, false
	},
	"AcFrequency": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Rtac.AcFrequency != nil {
			return float64(msg.XBOSIoTDeviceState.Rtac.AcFrequency.Value), true
		}
		return 0, false
	},
	"FaultCondition": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Rtac.FaultCondition != nil {
			return float64(msg.XBOSIoTDeviceState.Rtac.FaultCondition.Value), true
		}
		return 0, false
	},
	"PgeVoltage": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Rtac.PgeVoltage != nil {
			return float64(msg.XBOSIoTDeviceState.Rtac.PgeVoltage.Value), true
		}
		return 0, false
	},
	"PgeFrequency": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Rtac.PgeFrequency != nil {
			return float64(msg.XBOSIoTDeviceState.Rtac.PgeFrequency.Value), true
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
