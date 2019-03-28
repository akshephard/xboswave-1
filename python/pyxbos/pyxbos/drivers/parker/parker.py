from pyxbos import *
from modbus_driver import Modbus_Driver
import os,sys
import json
import requests
import yaml
import argparse
import time
from inspect import getmembers


class ParkerDriver(Driver):
    def setup(self, cfg):
        with open('parker.yaml') as f:
            # use safe_load instead load for security reasons
            driverConfig = yaml.safe_load(f)
        self.modbus_device = Modbus_Driver(driverConfig)
        self.modbus_device.initialize_modbus()
        self.service_name = cfg['service_name']

    def read(self, requestid=None):
        output = self.modbus_device.get_data()


        regulator_flag_1 = output['regulator_flag_1']
        output['energy_saving_regulator_flag'] = regulator_flag_1 & 0x0100
        output['energy_saving_real_time_regulator_flag'] = regulator_flag_1 & 0x0200
        output['service_request_regulator_flag'] = regulator_flag_1 & 0x0400

        regulator_flag_2 = output['regulator_flag_2']
        output['on_standby_regulator_flag'] = regulator_flag_2 & 0x0001
        output['new_alarm_to_read_regulator_flag'] = regulator_flag_2 & 0x0080
        output['defrost_status_regulator_flag']	= regulator_flag_2 & 0x0700

        digital_io_status = output['digital_io_status']
        output['door_switch_input_status'] = digital_io_status & 0x0001
        output['multipurpose_input_status'] = digital_io_status & 0x0002
        output['compressor_status'] = digital_io_status & 0x0100
        output['output_defrost_status'] = digital_io_status & 0x0200
        output['fans_status'] = digital_io_status & 0x0400
        output['output_k4_status'] = digital_io_status & 0x0800

        digital_output_flags = output['digital_output_flags']
        output['energy_saving_status'] = digital_output_flags & 0x0100
        output['service_request_status'] =	digital_output_flags & 0x0200
        output['resistors_activated_by_aux_key_status'] = digital_output_flags & 0x001
        output['evaporator_valve_state'] = digital_output_flags & 0x002
        output['output_defrost_state'] = digital_output_flags & 0x004
        output['output_lux_state'] =	digital_output_flags & 0x008
        output['output_aux_state'] =	digital_output_flags & 0x0010
        output['resistors_state'] = digital_output_flags & 0x0020
        output['output_alarm_state'] =	digital_output_flags & 0x0040
        output['second_compressor_state'] =	digital_output_flags & 0x0080

        alarm_status = output['alarm_status']
        #print(format(output['alarm_status'], '#010b'))
        output['probe1_failure_alarm'] = bool(alarm_status & 0x0100)
        output['probe2_failure_alarm'] = bool(alarm_status & 0x0200)
        output['probe3_failure_alarm'] = bool(alarm_status & 0x0400)
        output['minimum_temperature_alarm'] = bool(alarm_status & 0x1000)
        output['maximum_temperture_alarm'] = bool(alarm_status & 0x2000)
        output['condensor_temperature_failure_alarm'] = bool(alarm_status & 0x4000)
        output['condensor_pre_alarm'] = bool(alarm_status & 0x8000)
        output['door_alarm'] = bool(alarm_status & 0x0004)
        output['multipurpose_input_alarm'] = bool(alarm_status & 0x0008)
        output['compressor_blocked_alarm'] = bool(alarm_status & 0x0010)
        output['power_failure_alarm'] = bool(alarm_status & 0x0020)
        output['rtc_error_alarm'] = bool(alarm_status & 0x0080)
        #print(output['rtc_error_alarm'])
        #print(format(output['rtc_error_alarm'], '#010b'))
        print(output)

        msg = xbos_pb2.XBOS(
            XBOSIoTDeviceState = iot_pb2.XBOSIoTDeviceState(
                time = int(time.time()*1e9),
                parker_state = iot_pb2.ParkerState(
                    compressor_working_hours  =   types.Double(value=output['compressor_working_hours']),
                    on_standby_status  =   types.Int64(value=output['on_standby_status']),
                    light_status  =   types.Int64(value=output['light_status']),
                    aux_output_status  =   types.Int64(value=output['aux_output_status']),
                    next_defrost_counter  =   types.Double(value=output['next_defrost_counter']),
                    cabinet_temperature  =   types.Double(value=output['cabinet_temperature']),
                    evaporator_temperature  =   types.Double(value=output['evaporator_temperature']),
                    auxiliary_temperature  =   types.Double(value=output['auxiliary_temperature']),
                    time_until_defrost  =   types.Double(value=output['time_until_defrost']),
                    current_defrost_counter  =   types.Double(value=output['current_defrost_counter']),
                    compressor_delay  =   types.Double(value=output['compressor_delay']),
                    setpoint  =   types.Double(value=output['setpoint']),
                    r1  =   types.Double(value=output['r1']),
                    r2  =   types.Double(value=output['r2']),
                    r4  =   types.Double(value=output['r4']),
                    C0  =   types.Double(value=output['C0']),
                    C1  =   types.Double(value=output['C1']),
                    d0  =   types.Double(value=output['d0']),
                    d3  =   types.Double(value=output['d3']),
                    d5  =   types.Double(value=output['d5']),
                    d7  =   types.Double(value=output['d7']),
                    d8  =   types.Int64(value=output['d8']),
                    A0  =   types.Int64(value=output['A0']),
                    A1  =   types.Double(value=output['A1']),
                    A2  =   types.Int64(value=output['A2']),
                    A3  =   types.Int64(value=output['A3']),
                    A4  =   types.Double(value=output['A4']),
                    A5  =   types.Int64(value=output['A5']),
                    A6  =   types.Double(value=output['A6']),
                    A7  =   types.Double(value=output['A7']),
                    A8  =   types.Double(value=output['A8']),
                    A9  =   types.Double(value=output['A9']),
                    F0  =   types.Int64(value=output['F0']),
                    F1  =   types.Double(value=output['F1']),
                    F2  =   types.Int64(value=output['F2']),
                    F3  =   types.Double(value=output['F3'])

                )
            )
        )
        self.report(self.service_name, msg)


if __name__ == '__main__':
    with open('parker.yaml') as f:
        # use safe_load instead load for security reasons
        driverConfig = yaml.safe_load(f)

    namespace = driverConfig['wavemq']['namespace']
    service_name = driverConfig['xbos']['service_name']
    #driver_cfg = "parker.yaml"
    print(driverConfig)

    xbos_cfg = {
        'wavemq': 'localhost:4516',
        'namespace': namespace,
        'base_resource': 'test/darksky',
        'entity': 'parker.ent',
        'id': 'pyxbos-driver-darksky-1',
        #'rate': 1800, # half hour
        'rate': 900, # 15 min
        'service_name': service_name
    }
    print(getmembers(iot_pb2))
    logging.basicConfig(level="INFO", format='%(asctime)s - %(name)s - %(message)s')
    #e = DarkSkyDriver(cfg)
    e = ParkerDriver(xbos_cfg)
    e.begin()
