package main
import (
	"fmt"
	"github.com/gtfierro/xboswave/ingester/types"
	xbospb "github.com/gtfierro/xboswave/proto"
)
func has_device(msg xbospb.XBOS) bool {
	return msg.XBOSIoTDeviceState.Parker!= nil
}
var device_units = map[string]string{
	"compressor_working_hours":	"hours",
	"on_standby_status":	"",
	"light_status":	"",
	"aux_output_status":	"",
	"next_defrost_counter":	"seconds",
	"door_switch_input_status":	"",
	"multipurpose_input_status":	"",
	"compressor_status":	"",
	"output_defrost_status":	"",
	"fans_status":	"",
	"output_k4_status":	"",
	"cabinet_temperature":	"C",
	"evaporator_temperature":	"C",
	"auxiliary_temperature":	"C",
	"probe1_failure_alarm":	"",
	"probe2_failure_alarm":	"",
	"probe3_failure_alarm":	"",
	"minimum_temperature_alarm":	"",
	"maximum_temperture_alarm":	"",
	"condensor_temperature_failure_alarm":	"",
	"condensor_pre_alarm":	"",
	"door_alarm":	"",
	"multipurpose_input_alarm":	"",
	"compressor_blocked_alarm":	"",
	"power_failure_alarm":	"",
	"rtc_error_alarm":	"",
	"energy_saving_regulator_flag":	"",
	"energy_saving_real_time_regulator_flag":	"",
	"service_request_regulator_flag":	"",
	"on_standby_regulator_flag":	"",
	"new_alarm_to_read_regulator_flag":	"",
	"defrost_status_regulator_flag":	"",
	"active_setpoint":	"C",
	"time_until_defrost":	"seconds",
	"current_defrost_counter":	"seconds",
	"compressor_delay":	"seconds",
	"num_alarms_in_history":	"",
	"energy_saving_status":	"",
	"service_request_status":	"",
	"resistors_activated_by_aux_key_status":	"",
	"evaporator_valve_state":	"",
	"output_defrost_state":	"",
	"output_lux_state":	"",
	"output_aux_state":	"",
	"resistors_state":	"",
	"output_alarm_state":	"",
	"second_compressor_state":	"",
	"setpoint":	"",
	"r1":	"C",
	"r2":	"C",
	"r4":	"",
	"C0":	"minutes",
	"C1":	"minutes",
	"d0":	"hours",
	"d3":	"minutes",
	"d5":	"minutes",
	"d7":	"minutes",
	"d8":	"",
	"A0":	"",
	"A1":	"C",
	"A2":	"",
	"A3":	"",
	"A4":	"C",
	"A5":	"",
	"A6":	"minutes",
	"A7":	"minutes",
	"A8":	"minutes",
	"A9":	"minutes",
	"F0":	"",
	"F1":	"C",
	"F2":	"",
	"F3":	"minutes",
	"Hd1":	"hh:mm",
	"Hd2":	"hh:mm",
	"Hd3":	"hh:mm",
	"Hd4":	"hh:mm",
	"Hd5":	"hh:mm",
	"Hd6":	"hh:mm",
}
var device_lookup = map[string]func(msg xbospb.XBOS) (float64, bool){

	"CompressorWorkingHours": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Parker.CompressorWorkingHours != nil {
			return float64(msg.XBOSIoTDeviceState.Parker.CompressorWorkingHours.Value), true
		}
		return 0, false
	},
	"OnStandbyStatus": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Parker.OnStandbyStatus != nil {
			return float64(msg.XBOSIoTDeviceState.Parker.OnStandbyStatus.Value), true
		}
		return 0, false
	},
	"LightStatus": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Parker.LightStatus != nil {
			return float64(msg.XBOSIoTDeviceState.Parker.LightStatus.Value), true
		}
		return 0, false
	},
	"AuxOutputStatus": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Parker.AuxOutputStatus != nil {
			return float64(msg.XBOSIoTDeviceState.Parker.AuxOutputStatus.Value), true
		}
		return 0, false
	},
	"NextDefrostCounter": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Parker.NextDefrostCounter != nil {
			return float64(msg.XBOSIoTDeviceState.Parker.NextDefrostCounter.Value), true
		}
		return 0, false
	},
	"DoorSwitchInputStatus": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Parker.DoorSwitchInputStatus != nil {
			return float64(msg.XBOSIoTDeviceState.Parker.DoorSwitchInputStatus.Value), true
		}
		return 0, false
	},
	"MultipurposeInputStatus": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Parker.MultipurposeInputStatus != nil {
			return float64(msg.XBOSIoTDeviceState.Parker.MultipurposeInputStatus.Value), true
		}
		return 0, false
	},
	"CompressorStatus": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Parker.CompressorStatus != nil {
			return float64(msg.XBOSIoTDeviceState.Parker.CompressorStatus.Value), true
		}
		return 0, false
	},
	"OutputDefrostStatus": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Parker.OutputDefrostStatus != nil {
			return float64(msg.XBOSIoTDeviceState.Parker.OutputDefrostStatus.Value), true
		}
		return 0, false
	},
	"FansStatus": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Parker.FansStatus != nil {
			return float64(msg.XBOSIoTDeviceState.Parker.FansStatus.Value), true
		}
		return 0, false
	},
	"OutputK4Status": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Parker.OutputK4Status != nil {
			return float64(msg.XBOSIoTDeviceState.Parker.OutputK4Status.Value), true
		}
		return 0, false
	},
	"CabinetTemperature": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Parker.CabinetTemperature != nil {
			return float64(msg.XBOSIoTDeviceState.Parker.CabinetTemperature.Value), true
		}
		return 0, false
	},
	"EvaporatorTemperature": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Parker.EvaporatorTemperature != nil {
			return float64(msg.XBOSIoTDeviceState.Parker.EvaporatorTemperature.Value), true
		}
		return 0, false
	},
	"AuxiliaryTemperature": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Parker.AuxiliaryTemperature != nil {
			return float64(msg.XBOSIoTDeviceState.Parker.AuxiliaryTemperature.Value), true
		}
		return 0, false
	},
	"Probe1FailureAlarm": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Parker.Probe1FailureAlarm != nil {
			return float64(msg.XBOSIoTDeviceState.Parker.Probe1FailureAlarm.Value), true
		}
		return 0, false
	},
	"Probe2FailureAlarm": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Parker.Probe2FailureAlarm != nil {
			return float64(msg.XBOSIoTDeviceState.Parker.Probe2FailureAlarm.Value), true
		}
		return 0, false
	},
	"Probe3FailureAlarm": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Parker.Probe3FailureAlarm != nil {
			return float64(msg.XBOSIoTDeviceState.Parker.Probe3FailureAlarm.Value), true
		}
		return 0, false
	},
	"MinimumTemperatureAlarm": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Parker.MinimumTemperatureAlarm != nil {
			return float64(msg.XBOSIoTDeviceState.Parker.MinimumTemperatureAlarm.Value), true
		}
		return 0, false
	},
	"MaximumTempertureAlarm": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Parker.MaximumTempertureAlarm != nil {
			return float64(msg.XBOSIoTDeviceState.Parker.MaximumTempertureAlarm.Value), true
		}
		return 0, false
	},
	"CondensorTemperatureFailureAlarm": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Parker.CondensorTemperatureFailureAlarm != nil {
			return float64(msg.XBOSIoTDeviceState.Parker.CondensorTemperatureFailureAlarm.Value), true
		}
		return 0, false
	},
	"CondensorPreAlarm": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Parker.CondensorPreAlarm != nil {
			return float64(msg.XBOSIoTDeviceState.Parker.CondensorPreAlarm.Value), true
		}
		return 0, false
	},
	"DoorAlarm": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Parker.DoorAlarm != nil {
			return float64(msg.XBOSIoTDeviceState.Parker.DoorAlarm.Value), true
		}
		return 0, false
	},
	"MultipurposeInputAlarm": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Parker.MultipurposeInputAlarm != nil {
			return float64(msg.XBOSIoTDeviceState.Parker.MultipurposeInputAlarm.Value), true
		}
		return 0, false
	},
	"CompressorBlockedAlarm": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Parker.CompressorBlockedAlarm != nil {
			return float64(msg.XBOSIoTDeviceState.Parker.CompressorBlockedAlarm.Value), true
		}
		return 0, false
	},
	"PowerFailureAlarm": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Parker.PowerFailureAlarm != nil {
			return float64(msg.XBOSIoTDeviceState.Parker.PowerFailureAlarm.Value), true
		}
		return 0, false
	},
	"RtcErrorAlarm": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Parker.RtcErrorAlarm != nil {
			return float64(msg.XBOSIoTDeviceState.Parker.RtcErrorAlarm.Value), true
		}
		return 0, false
	},
	"EnergySavingRegulatorFlag": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Parker.EnergySavingRegulatorFlag != nil {
			return float64(msg.XBOSIoTDeviceState.Parker.EnergySavingRegulatorFlag.Value), true
		}
		return 0, false
	},
	"EnergySavingRealTimeRegulatorFlag": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Parker.EnergySavingRealTimeRegulatorFlag != nil {
			return float64(msg.XBOSIoTDeviceState.Parker.EnergySavingRealTimeRegulatorFlag.Value), true
		}
		return 0, false
	},
	"ServiceRequestRegulatorFlag": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Parker.ServiceRequestRegulatorFlag != nil {
			return float64(msg.XBOSIoTDeviceState.Parker.ServiceRequestRegulatorFlag.Value), true
		}
		return 0, false
	},
	"OnStandbyRegulatorFlag": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Parker.OnStandbyRegulatorFlag != nil {
			return float64(msg.XBOSIoTDeviceState.Parker.OnStandbyRegulatorFlag.Value), true
		}
		return 0, false
	},
	"NewAlarmToReadRegulatorFlag": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Parker.NewAlarmToReadRegulatorFlag != nil {
			return float64(msg.XBOSIoTDeviceState.Parker.NewAlarmToReadRegulatorFlag.Value), true
		}
		return 0, false
	},
	"DefrostStatusRegulatorFlag": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Parker.DefrostStatusRegulatorFlag != nil {
			return float64(msg.XBOSIoTDeviceState.Parker.DefrostStatusRegulatorFlag.Value), true
		}
		return 0, false
	},
	"ActiveSetpoint": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Parker.ActiveSetpoint != nil {
			return float64(msg.XBOSIoTDeviceState.Parker.ActiveSetpoint.Value), true
		}
		return 0, false
	},
	"TimeUntilDefrost": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Parker.TimeUntilDefrost != nil {
			return float64(msg.XBOSIoTDeviceState.Parker.TimeUntilDefrost.Value), true
		}
		return 0, false
	},
	"CurrentDefrostCounter": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Parker.CurrentDefrostCounter != nil {
			return float64(msg.XBOSIoTDeviceState.Parker.CurrentDefrostCounter.Value), true
		}
		return 0, false
	},
	"CompressorDelay": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Parker.CompressorDelay != nil {
			return float64(msg.XBOSIoTDeviceState.Parker.CompressorDelay.Value), true
		}
		return 0, false
	},
	"NumAlarmsInHistory": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Parker.NumAlarmsInHistory != nil {
			return float64(msg.XBOSIoTDeviceState.Parker.NumAlarmsInHistory.Value), true
		}
		return 0, false
	},
	"EnergySavingStatus": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Parker.EnergySavingStatus != nil {
			return float64(msg.XBOSIoTDeviceState.Parker.EnergySavingStatus.Value), true
		}
		return 0, false
	},
	"ServiceRequestStatus": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Parker.ServiceRequestStatus != nil {
			return float64(msg.XBOSIoTDeviceState.Parker.ServiceRequestStatus.Value), true
		}
		return 0, false
	},
	"ResistorsActivatedByAuxKeyStatus": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Parker.ResistorsActivatedByAuxKeyStatus != nil {
			return float64(msg.XBOSIoTDeviceState.Parker.ResistorsActivatedByAuxKeyStatus.Value), true
		}
		return 0, false
	},
	"EvaporatorValveState": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Parker.EvaporatorValveState != nil {
			return float64(msg.XBOSIoTDeviceState.Parker.EvaporatorValveState.Value), true
		}
		return 0, false
	},
	"OutputDefrostState": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Parker.OutputDefrostState != nil {
			return float64(msg.XBOSIoTDeviceState.Parker.OutputDefrostState.Value), true
		}
		return 0, false
	},
	"OutputLuxState": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Parker.OutputLuxState != nil {
			return float64(msg.XBOSIoTDeviceState.Parker.OutputLuxState.Value), true
		}
		return 0, false
	},
	"OutputAuxState": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Parker.OutputAuxState != nil {
			return float64(msg.XBOSIoTDeviceState.Parker.OutputAuxState.Value), true
		}
		return 0, false
	},
	"ResistorsState": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Parker.ResistorsState != nil {
			return float64(msg.XBOSIoTDeviceState.Parker.ResistorsState.Value), true
		}
		return 0, false
	},
	"OutputAlarmState": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Parker.OutputAlarmState != nil {
			return float64(msg.XBOSIoTDeviceState.Parker.OutputAlarmState.Value), true
		}
		return 0, false
	},
	"SecondCompressorState": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Parker.SecondCompressorState != nil {
			return float64(msg.XBOSIoTDeviceState.Parker.SecondCompressorState.Value), true
		}
		return 0, false
	},
	"Setpoint": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Parker.Setpoint != nil {
			return float64(msg.XBOSIoTDeviceState.Parker.Setpoint.Value), true
		}
		return 0, false
	},
	"R1": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Parker.R1 != nil {
			return float64(msg.XBOSIoTDeviceState.Parker.R1.Value), true
		}
		return 0, false
	},
	"R2": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Parker.R2 != nil {
			return float64(msg.XBOSIoTDeviceState.Parker.R2.Value), true
		}
		return 0, false
	},
	"R4": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Parker.R4 != nil {
			return float64(msg.XBOSIoTDeviceState.Parker.R4.Value), true
		}
		return 0, false
	},
	"C0": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Parker.C0 != nil {
			return float64(msg.XBOSIoTDeviceState.Parker.C0.Value), true
		}
		return 0, false
	},
	"C1": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Parker.C1 != nil {
			return float64(msg.XBOSIoTDeviceState.Parker.C1.Value), true
		}
		return 0, false
	},
	"D0": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Parker.D0 != nil {
			return float64(msg.XBOSIoTDeviceState.Parker.D0.Value), true
		}
		return 0, false
	},
	"D3": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Parker.D3 != nil {
			return float64(msg.XBOSIoTDeviceState.Parker.D3.Value), true
		}
		return 0, false
	},
	"D5": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Parker.D5 != nil {
			return float64(msg.XBOSIoTDeviceState.Parker.D5.Value), true
		}
		return 0, false
	},
	"D7": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Parker.D7 != nil {
			return float64(msg.XBOSIoTDeviceState.Parker.D7.Value), true
		}
		return 0, false
	},
	"D8": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Parker.D8 != nil {
			return float64(msg.XBOSIoTDeviceState.Parker.D8.Value), true
		}
		return 0, false
	},
	"A0": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Parker.A0 != nil {
			return float64(msg.XBOSIoTDeviceState.Parker.A0.Value), true
		}
		return 0, false
	},
	"A1": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Parker.A1 != nil {
			return float64(msg.XBOSIoTDeviceState.Parker.A1.Value), true
		}
		return 0, false
	},
	"A2": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Parker.A2 != nil {
			return float64(msg.XBOSIoTDeviceState.Parker.A2.Value), true
		}
		return 0, false
	},
	"A3": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Parker.A3 != nil {
			return float64(msg.XBOSIoTDeviceState.Parker.A3.Value), true
		}
		return 0, false
	},
	"A4": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Parker.A4 != nil {
			return float64(msg.XBOSIoTDeviceState.Parker.A4.Value), true
		}
		return 0, false
	},
	"A5": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Parker.A5 != nil {
			return float64(msg.XBOSIoTDeviceState.Parker.A5.Value), true
		}
		return 0, false
	},
	"A6": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Parker.A6 != nil {
			return float64(msg.XBOSIoTDeviceState.Parker.A6.Value), true
		}
		return 0, false
	},
	"A7": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Parker.A7 != nil {
			return float64(msg.XBOSIoTDeviceState.Parker.A7.Value), true
		}
		return 0, false
	},
	"A8": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Parker.A8 != nil {
			return float64(msg.XBOSIoTDeviceState.Parker.A8.Value), true
		}
		return 0, false
	},
	"A9": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Parker.A9 != nil {
			return float64(msg.XBOSIoTDeviceState.Parker.A9.Value), true
		}
		return 0, false
	},
	"F0": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Parker.F0 != nil {
			return float64(msg.XBOSIoTDeviceState.Parker.F0.Value), true
		}
		return 0, false
	},
	"F1": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Parker.F1 != nil {
			return float64(msg.XBOSIoTDeviceState.Parker.F1.Value), true
		}
		return 0, false
	},
	"F2": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Parker.F2 != nil {
			return float64(msg.XBOSIoTDeviceState.Parker.F2.Value), true
		}
		return 0, false
	},
	"F3": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Parker.F3 != nil {
			return float64(msg.XBOSIoTDeviceState.Parker.F3.Value), true
		}
		return 0, false
	},
	"Hd1": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Parker.Hd1 != nil {
			return float64(msg.XBOSIoTDeviceState.Parker.Hd1.Value), true
		}
		return 0, false
	},
	"Hd2": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Parker.Hd2 != nil {
			return float64(msg.XBOSIoTDeviceState.Parker.Hd2.Value), true
		}
		return 0, false
	},
	"Hd3": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Parker.Hd3 != nil {
			return float64(msg.XBOSIoTDeviceState.Parker.Hd3.Value), true
		}
		return 0, false
	},
	"Hd4": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Parker.Hd4 != nil {
			return float64(msg.XBOSIoTDeviceState.Parker.Hd4.Value), true
		}
		return 0, false
	},
	"Hd5": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Parker.Hd5 != nil {
			return float64(msg.XBOSIoTDeviceState.Parker.Hd5.Value), true
		}
		return 0, false
	},
	"Hd6": func(msg xbospb.XBOS) (float64, bool) {
		if has_device(msg) && msg.XBOSIoTDeviceState.Parker.Hd6 != nil {
			return float64(msg.XBOSIoTDeviceState.Parker.Hd6.Value), true
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
