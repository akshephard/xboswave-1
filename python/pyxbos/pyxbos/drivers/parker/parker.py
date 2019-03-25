from pyxbos import *
from modbus_driver import Modbus_Driver
import os,sys
import json
import requests
import yaml
import argparse
import time


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
        HAACP_0_0 = output['HAACP_0_0']
        HAACP_0_1 = output['HAACP_0_1']
        HAACP_0_2 = output['HAACP_0_2']
        compressor_working_hours = output['compressor_working_hours']
        defrost_control = output['defrost_control']
        status_on = output['status_on']
        status_light = output['status_light']
        status_aux_output = output['status_aux_output']
        next_defrost_counter = output['next_defrost_counter']
        digital_input_output_status = output['digital_input_output_status']
        cab_probe = output['cab_probe']
        evap_probe = output['evap_probe']
        aux_probe = output['aux_probe']
        alarm_status = output['alarm_status']
        regulator_flag1 = output['regulator_flag1']
        regulator_flag2 = output['regulator_flag2']
        active_set_point = output['active_set_point']
        time_remaining_next_defrost = output['time_remaining_next_defrost']
        current_defrost_counter = output['current_defrost_counter']
        compressor_delay = output['compressor_delay']
        num_HAACP_alarm_historial = output['num_HAACP_alarm_historial']
        digital_output_flag = output['digital_output_flag']
        SP = output['SP']
        CA1 = output['CA1']
        CA2 = output['CA2']
        CA3 = output['CA3']
        p0 = output['p0']
        p1 = output['p1']
        p2 = output['p2']
        p3 = output['p3']
        p4 = output['p4']
        p5 = output['p5']
        p6 = output['p6']
        r0 = output['r0']
        r1 = output['r1']
        r2 = output['r2']
        r3 = output['r3']
        r4 = output['r4']
        c0 = output['c0']
        c1 = output['c1']
        c2 = output['c2']
        c3 = output['c3']
        c4 = output['c4']
        c5 = output['c5']
        c6 = output['c6']
        c7 = output['c7']
        c8 = output['c8']
        c9 = output['c9']
        c10 = output['c10']
        d0 = output['d0']
        d1 = output['d1']
        d2 = output['d2']
        d3 = output['d3']
        d4 = output['d4']
        d5 = output['d5']
        d6 = output['d6']
        d7 = output['d7']
        d8 = output['d8']
        d9 = output['d9']
        dA = output['dA']
        A0 = output['A0']
        A1 = output['A1']
        A2 = output['A2']
        A3 = output['A3']
        A4 = output['A4']
        A5 = output['A5']
        A6 = output['A6']
        A7 = output['A7']
        A8 = output['A8']
        A9 = output['A9']
        AA = output['AA']
        F0 = output['F0']
        F1 = output['F1']
        F2 = output['F2']
        F3 = output['F3']
        i0 = output['i0']
        i1 = output['i1']
        i2 = output['i2']
        i3 = output['i3']
        i4 = output['i4']
        i5 = output['i5']
        i6 = output['i6']
        i7 = output['i7']
        i8 = output['i8']
        i9 = output['i9']
        u1 = output['u1']
        u2 = output['u2']
        u3 = output['u3']
        u4 = output['u4']
        u5 = output['u5']
        u6 = output['u6']
        u7 = output['u7']
        u8 = output['u8']
        HE1 = output['HE1']
        HE2 = output['HE2']
        hd1 = output['hd1']
        hd2 = output['hd2']
        hd3 = output['hd3']
        hd4 = output['hd4']
        hd5 = output['hd5']
        hd6 = output['hd6']

        msg = xbos_pb2.XBOS(
            XBOSIoTDeviceState = iot_pb2.XBOSIoTDeviceState(
                time = int(time.time()*1e9),
                parker_state = iot_pb2.ParkerState(
                    compressor_working_hours  =   types.Double(value=compressor_working_hours),
                    clear_compressor_working_hours  =   types.Int64(value=clear_compressor_working_hours),
                    buzzer_control  =   types.Int64(value=buzzer_control),
                    defrost_control  =   types.Int64(value=defrost_control),
                    start_resistors  =   types.Int64(value=start_resistors),
                    on_standby_status  =   types.Int64(value=on_standby_status),
                    light_status  =   types.Int64(value=light_status),
                    aux_output_status  =   types.Int64(value=aux_output_status),
                    next_defrost_counter  =   types.Double(value=next_defrost_counter),
                    door_switch_input_status  =   types.Bool(value=door_switch_input_status),
                    multipurpose_input_status  =   types.Bool(value=multipurpose_input_status),
                    compressor_status  =   types.Bool(value=compressor_status),
                    output_defrost_status  =   types.Bool(value=output_defrost_status),
                    fans_status  =   types.Bool(value=fans_status),
                    output_k4_status  =   types.Bool(value=output_k4_status),
                    cabinet_temperature  =   types.Double(value=cabinet_temperature),
                    evaporator_temperature  =   types.Double(value=evaporator_temperature),
                    auxiliary_temperature  =   types.Double(value=auxiliary_temperature),
                    probe1_failure_alarm  =   types.Bool(value=probe1_failure_alarm),
                    probe2_failure_alarm  =   types.Bool(value=probe2_failure_alarm),
                    probe3_failure_alarm  =   types.Bool(value=probe3_failure_alarm),
                    minimum_temperature_alarm  =   types.Bool(value=minimum_temperature_alarm),
                    maximum_temperture_alarm  =   types.Bool(value=maximum_temperture_alarm),
                    condensor_temperature_failure_alarm  =   types.Bool(value=condensor_temperature_failure_alarm),
                    condensor_pre_alarm  =   types.Bool(value=condensor_pre_alarm),
                    door_alarm  =   types.Bool(value=door_alarm),
                    multipurpose_input_alarm  =   types.Bool(value=multipurpose_input_alarm),
                    compressor_blocked_alarm  =   types.Bool(value=compressor_blocked_alarm),
                    power_failure_alarm  =   types.Bool(value=power_failure_alarm),
                    rtc_error_alarm  =   types.Bool(value=rtc_error_alarm),
                    energy_saving_regulator_flag  =   types.Bool(value=energy_saving_regulator_flag),
                    energy_saving_real_time_regulator_flag  =   types.Bool(value=energy_saving_real_time_regulator_flag),
                    service_request_regulator_flag  =   types.Bool(value=service_request_regulator_flag),
                    on_standby_regulator_flag  =   types.Bool(value=on_standby_regulator_flag),
                    new_alarm_to_read_regulator_flag  =   types.Bool(value=new_alarm_to_read_regulator_flag),
                    defrost_status_regulator_flag  =   types.Int64(value=defrost_status_regulator_flag),
                    active_setpoint  =   types.Double(value=active_setpoint),
                    time_until_defrost  =   types.Double(value=time_until_defrost),
                    current_defrost_counter  =   types.Double(value=current_defrost_counter),
                    compressor_delay  =   types.Double(value=compressor_delay),
                    num_alarms_in_history  =   types.Int64(value=num_alarms_in_history),
                    energy_saving_status  =   types.Bool(value=energy_saving_status),
                    service_request_status  =   types.Bool(value=service_request_status),
                    resistors_activated_by_aux_key_status  =   types.Bool(value=resistors_activated_by_aux_key_status),
                    evaporator_valve_state  =   types.Bool(value=evaporator_valve_state),
                    output_defrost_state  =   types.Bool(value=output_defrost_state),
                    output_lux_state  =   types.Bool(value=output_lux_state),
                    output_aux_state  =   types.Bool(value=output_aux_state),
                    resistors_state  =   types.Bool(value=resistors_state),
                    output_alarm_state  =   types.Bool(value=output_alarm_state),
                    second_compressor_state  =   types.Bool(value=second_compressor_state),
                    setpoint  =   types.Double(value=setpoint),
                    clear_alarm_history  =   types.Int64(value=clear_alarm_history),
                    clearn_new_alarm_flag  =   types.Int64(value=clearn_new_alarm_flag),
                    HACCP0_datetime  =   types.datetime(value=HACCP0_datetime),
                    HACCP0_alarm_type  =   types.Int64(value=HACCP0_alarm_type),
                    HACCP0_duration  =   types.Double(value=HACCP0_duration),
                    HACCP0_temperature  =   types.Double(value=HACCP0_temperature),
                    HACCP1_datetime  =   types.datetime(value=HACCP1_datetime),
                    HACCP1_alarm_type  =   types.Int64(value=HACCP1_alarm_type),
                    HACCP1_duration  =   types.Double(value=HACCP1_duration),
                    HACCP1_temperature  =   types.Double(value=HACCP1_temperature),
                    HACCP2_datetime  =   types.datetime(value=HACCP2_datetime),
                    HACCP2_alarm_type  =   types.Int64(value=HACCP2_alarm_type),
                    HACCP2_duration  =   types.Double(value=HACCP2_duration),
                    HACCP2_temperature  =   types.Double(value=HACCP2_temperature),
                    HACCP3_datetime  =   types.datetime(value=HACCP3_datetime),
                    HACCP3_alarm_type  =   types.Int64(value=HACCP3_alarm_type),
                    HACCP3_duration  =   types.Double(value=HACCP3_duration),
                    HACCP3_temperature  =   types.Double(value=HACCP3_temperature),
                    HACCP4_datetime  =   types.datetime(value=HACCP4_datetime),
                    HACCP4_alarm_type  =   types.Int64(value=HACCP4_alarm_type),
                    HACCP4_duration  =   types.Double(value=HACCP4_duration),
                    HACCP4_temperature  =   types.Double(value=HACCP4_temperature),
                    HACCP5_datetime  =   types.datetime(value=HACCP5_datetime),
                    HACCP5_alarm_type  =   types.Int64(value=HACCP5_alarm_type),
                    HACCP5_duration  =   types.Double(value=HACCP5_duration),
                    HACCP5_temperature  =   types.Double(value=HACCP5_temperature),
                    HACCP6_datetime  =   types.datetime(value=HACCP6_datetime),
                    HACCP6_alarm_type  =   types.Int64(value=HACCP6_alarm_type),
                    HACCP6_duration  =   types.Double(value=HACCP6_duration),
                    HACCP6_temperature  =   types.Double(value=HACCP6_temperature),
                    HACCP7_datetime  =   types.datetime(value=HACCP7_datetime),
                    HACCP7_alarm_type  =   types.Int64(value=HACCP7_alarm_type),
                    HACCP7_duration  =   types.Double(value=HACCP7_duration),
                    HACCP7_temperature  =   types.Double(value=HACCP7_temperature),
                    HACCP8_datetime  =   types.datetime(value=HACCP8_datetime),
                    HACCP8_alarm_type  =   types.Int64(value=HACCP8_alarm_type),
                    HACCP8_duration  =   types.Double(value=HACCP8_duration),
                    HACCP8_temperature  =   types.Double(value=HACCP8_temperature),
                    r1  =   types.Double(value=r1),
                    r2  =   types.Double(value=r2),
                    r4  =   types.Double(value=r4),
                    C0  =   types.Double(value=C0),
                    C1  =   types.Double(value=C1),
                    d0  =   types.Int64(value=d0),
                    d3  =   types.Double(value=d3),
                    d5  =   types.Double(value=d5),
                    d7  =   types.Double(value=d7),
                    d8  =   types.Int64(value=d8),
                    A0  =   types.Int64(value=A0),
                    A1  =   types.Double(value=A1),
                    A2  =   types.Int64(value=A2),
                    A3  =   types.Int64(value=A3),
                    A4  =   types.Double(value=A4),
                    A5  =   types.Int64(value=A5),
                    A6  =   types.Double(value=A6),
                    A7  =   types.Double(value=A7),
                    A8  =   types.Double(value=A8),
                    A9  =   types.Double(value=A9),
                    F0  =   types.Int64(value=F0),
                    F1  =   types.Double(value=F1),
                    F2  =   types.Int64(value=F2),
                    F3  =   types.Double(value=F3),
                    Hd1  =   types.string(value=Hd1),
                    Hd2  =   types.string(value=Hd2),
                    Hd3  =   types.string(value=Hd3),
                    Hd4  =   types.string(value=Hd4),
                    Hd5  =   types.string(value=Hd5),
                    Hd6  =   types.string(value=Hd6)
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
    logging.basicConfig(level="INFO", format='%(asctime)s - %(name)s - %(message)s')
    #e = DarkSkyDriver(cfg)
    e = ParkerDriver(xbos_cfg)
    e.begin()
